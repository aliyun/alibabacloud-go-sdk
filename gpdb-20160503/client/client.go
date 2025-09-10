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
		"cn-beijing":            dara.String("gpdb.aliyuncs.com"),
		"cn-hangzhou":           dara.String("gpdb.aliyuncs.com"),
		"cn-shanghai":           dara.String("gpdb.aliyuncs.com"),
		"cn-shenzhen":           dara.String("gpdb.aliyuncs.com"),
		"cn-hongkong":           dara.String("gpdb.aliyuncs.com"),
		"ap-southeast-1":        dara.String("gpdb.aliyuncs.com"),
		"us-west-1":             dara.String("gpdb.aliyuncs.com"),
		"us-east-1":             dara.String("gpdb.aliyuncs.com"),
		"cn-hangzhou-finance":   dara.String("gpdb.aliyuncs.com"),
		"cn-shanghai-finance-1": dara.String("gpdb.aliyuncs.com"),
		"cn-shenzhen-finance-1": dara.String("gpdb.aliyuncs.com"),
		"cn-qingdao":            dara.String("gpdb.aliyuncs.com"),
		"cn-north-2-gov-1":      dara.String("gpdb.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("gpdb"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) _postOSSObject(bucketName *string, form map[string]interface{}) (_result map[string]interface{}, _err error) {
	request_ := dara.NewRequest()
	boundary := dara.GetBoundary()
	request_.Protocol = dara.String("HTTPS")
	request_.Method = dara.String("POST")
	request_.Pathname = dara.String("/")
	request_.Headers = map[string]*string{
		"host":       dara.String(dara.ToString(form["host"])),
		"date":       openapiutil.GetDateUTCString(),
		"user-agent": openapiutil.GetUserAgent(dara.String("")),
	}
	request_.Headers["content-type"] = dara.String("multipart/form-data; boundary=" + boundary)
	request_.Body = dara.ToFileForm(form, boundary)
	response_, _err := dara.DoRequest(request_, nil)
	if _err != nil {
		return nil, _err
	}

	_result, _err = _postOSSObject_opResponse(response_)
	if _err != nil {
		return nil, _err
	}

	return _result, nil
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
// Allocates a public endpoint for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to apply for a public endpoint for an AnalyticDB for PostgreSQL instance. Both the primary and instance endpoints of an AnalyticDB for PostgreSQL instance can be public endpoints. For more information, see [Endpoints of an instance and its primary coordinator node](https://help.aliyun.com/document_detail/204879.html).
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
	if !dara.IsNil(request.AddressType) {
		query["AddressType"] = request.AddressType
	}

	if !dara.IsNil(request.ConnectionStringPrefix) {
		query["ConnectionStringPrefix"] = request.ConnectionStringPrefix
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AllocateInstancePublicConnection"),
		Version:     dara.String("2016-05-03"),
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
// Allocates a public endpoint for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to apply for a public endpoint for an AnalyticDB for PostgreSQL instance. Both the primary and instance endpoints of an AnalyticDB for PostgreSQL instance can be public endpoints. For more information, see [Endpoints of an instance and its primary coordinator node](https://help.aliyun.com/document_detail/204879.html).
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// Binds a resource group to a database role.
//
// @param tmpReq - BindDBResourceGroupWithRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindDBResourceGroupWithRoleResponse
func (client *Client) BindDBResourceGroupWithRoleWithOptions(tmpReq *BindDBResourceGroupWithRoleRequest, runtime *dara.RuntimeOptions) (_result *BindDBResourceGroupWithRoleResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &BindDBResourceGroupWithRoleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RoleList) {
		request.RoleListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoleList, dara.String("RoleList"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceGroupName) {
		query["ResourceGroupName"] = request.ResourceGroupName
	}

	if !dara.IsNil(request.RoleListShrink) {
		query["RoleList"] = request.RoleListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindDBResourceGroupWithRole"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindDBResourceGroupWithRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds a resource group to a database role.
//
// @param request - BindDBResourceGroupWithRoleRequest
//
// @return BindDBResourceGroupWithRoleResponse
func (client *Client) BindDBResourceGroupWithRole(request *BindDBResourceGroupWithRoleRequest) (_result *BindDBResourceGroupWithRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindDBResourceGroupWithRoleResponse{}
	_body, _err := client.BindDBResourceGroupWithRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消创建索引任务
//
// @param request - CancelCreateIndexJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelCreateIndexJobResponse
func (client *Client) CancelCreateIndexJobWithOptions(request *CancelCreateIndexJobRequest, runtime *dara.RuntimeOptions) (_result *CancelCreateIndexJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Collection) {
		query["Collection"] = request.Collection
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespacePassword) {
		query["NamespacePassword"] = request.NamespacePassword
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelCreateIndexJob"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelCreateIndexJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消创建索引任务
//
// @param request - CancelCreateIndexJobRequest
//
// @return CancelCreateIndexJobResponse
func (client *Client) CancelCreateIndexJob(request *CancelCreateIndexJobRequest) (_result *CancelCreateIndexJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelCreateIndexJobResponse{}
	_body, _err := client.CancelCreateIndexJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels an asynchronous document upload job based on the job ID.
//
// Description:
//
// This operation is related to the UploadDocumentAsync operation. You can call this operation to cancel a document upload job.
//
// >  If the canceling operation is complete, failed, or is canceled, you cannot call the operation again. The canceling operation only interrupts the document upload job. To remove the uploaded data, you must manually remove it or call the DeleteCollectionData operation. You can also call the document upload operation to overwrite the data by using the same FileName parameter.
//
// @param request - CancelUploadDocumentJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelUploadDocumentJobResponse
func (client *Client) CancelUploadDocumentJobWithOptions(request *CancelUploadDocumentJobRequest, runtime *dara.RuntimeOptions) (_result *CancelUploadDocumentJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Collection) {
		body["Collection"] = request.Collection
	}

	if !dara.IsNil(request.JobId) {
		body["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespacePassword) {
		body["NamespacePassword"] = request.NamespacePassword
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelUploadDocumentJob"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelUploadDocumentJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels an asynchronous document upload job based on the job ID.
//
// Description:
//
// This operation is related to the UploadDocumentAsync operation. You can call this operation to cancel a document upload job.
//
// >  If the canceling operation is complete, failed, or is canceled, you cannot call the operation again. The canceling operation only interrupts the document upload job. To remove the uploaded data, you must manually remove it or call the DeleteCollectionData operation. You can also call the document upload operation to overwrite the data by using the same FileName parameter.
//
// @param request - CancelUploadDocumentJobRequest
//
// @return CancelUploadDocumentJobResponse
func (client *Client) CancelUploadDocumentJob(request *CancelUploadDocumentJobRequest) (_result *CancelUploadDocumentJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelUploadDocumentJobResponse{}
	_body, _err := client.CancelUploadDocumentJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels an asynchronous vector data upload job by using a job ID.
//
// Description:
//
// This operation is related to the `UpsertCollectionDataAsync` operation. You can call this operation to cancel an upload job.
//
// >  If the canceling operation is complete, failed, or is canceled, you cannot call the operation again. The canceling operation only interrupts the upload job. To remove the uploaded data, you must manually remove it or call the DeleteCollectionData operation.
//
// @param request - CancelUpsertCollectionDataJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelUpsertCollectionDataJobResponse
func (client *Client) CancelUpsertCollectionDataJobWithOptions(request *CancelUpsertCollectionDataJobRequest, runtime *dara.RuntimeOptions) (_result *CancelUpsertCollectionDataJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Collection) {
		body["Collection"] = request.Collection
	}

	if !dara.IsNil(request.JobId) {
		body["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespacePassword) {
		body["NamespacePassword"] = request.NamespacePassword
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelUpsertCollectionDataJob"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelUpsertCollectionDataJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels an asynchronous vector data upload job by using a job ID.
//
// Description:
//
// This operation is related to the `UpsertCollectionDataAsync` operation. You can call this operation to cancel an upload job.
//
// >  If the canceling operation is complete, failed, or is canceled, you cannot call the operation again. The canceling operation only interrupts the upload job. To remove the uploaded data, you must manually remove it or call the DeleteCollectionData operation.
//
// @param request - CancelUpsertCollectionDataJobRequest
//
// @return CancelUpsertCollectionDataJobResponse
func (client *Client) CancelUpsertCollectionDataJob(request *CancelUpsertCollectionDataJobRequest) (_result *CancelUpsertCollectionDataJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelUpsertCollectionDataJobResponse{}
	_body, _err := client.CancelUpsertCollectionDataJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks the configurations of a Hadoop data source.
//
// @param request - CheckHadoopDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckHadoopDataSourceResponse
func (client *Client) CheckHadoopDataSourceWithOptions(request *CheckHadoopDataSourceRequest, runtime *dara.RuntimeOptions) (_result *CheckHadoopDataSourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CheckDir) {
		query["CheckDir"] = request.CheckDir
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DataSourceId) {
		query["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckHadoopDataSource"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckHadoopDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks the configurations of a Hadoop data source.
//
// @param request - CheckHadoopDataSourceRequest
//
// @return CheckHadoopDataSourceResponse
func (client *Client) CheckHadoopDataSource(request *CheckHadoopDataSourceRequest) (_result *CheckHadoopDataSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckHadoopDataSourceResponse{}
	_body, _err := client.CheckHadoopDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Check Hadoop Cluster Network Connectivity
//
// @param request - CheckHadoopNetConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckHadoopNetConnectionResponse
func (client *Client) CheckHadoopNetConnectionWithOptions(request *CheckHadoopNetConnectionRequest, runtime *dara.RuntimeOptions) (_result *CheckHadoopNetConnectionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DataSourceId) {
		query["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.EmrInstanceId) {
		query["EmrInstanceId"] = request.EmrInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckHadoopNetConnection"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckHadoopNetConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Check Hadoop Cluster Network Connectivity
//
// @param request - CheckHadoopNetConnectionRequest
//
// @return CheckHadoopNetConnectionResponse
func (client *Client) CheckHadoopNetConnection(request *CheckHadoopNetConnectionRequest) (_result *CheckHadoopNetConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckHadoopNetConnectionResponse{}
	_body, _err := client.CheckHadoopNetConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Check the network connectivity of the JDBC connection string
//
// @param request - CheckJDBCSourceNetConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckJDBCSourceNetConnectionResponse
func (client *Client) CheckJDBCSourceNetConnectionWithOptions(request *CheckJDBCSourceNetConnectionRequest, runtime *dara.RuntimeOptions) (_result *CheckJDBCSourceNetConnectionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DataSourceId) {
		query["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.JdbcConnectionString) {
		query["JdbcConnectionString"] = request.JdbcConnectionString
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckJDBCSourceNetConnection"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckJDBCSourceNetConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Check the network connectivity of the JDBC connection string
//
// @param request - CheckJDBCSourceNetConnectionRequest
//
// @return CheckJDBCSourceNetConnectionResponse
func (client *Client) CheckJDBCSourceNetConnection(request *CheckJDBCSourceNetConnectionRequest) (_result *CheckJDBCSourceNetConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckJDBCSourceNetConnectionResponse{}
	_body, _err := client.CheckJDBCSourceNetConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether a service-linked role is created.
//
// @param request - CheckServiceLinkedRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckServiceLinkedRoleResponse
func (client *Client) CheckServiceLinkedRoleWithOptions(request *CheckServiceLinkedRoleRequest, runtime *dara.RuntimeOptions) (_result *CheckServiceLinkedRoleResponse, _err error) {
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
		Action:      dara.String("CheckServiceLinkedRole"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckServiceLinkedRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether a service-linked role is created.
//
// @param request - CheckServiceLinkedRoleRequest
//
// @return CheckServiceLinkedRoleResponse
func (client *Client) CheckServiceLinkedRole(request *CheckServiceLinkedRoleRequest) (_result *CheckServiceLinkedRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckServiceLinkedRoleResponse{}
	_body, _err := client.CheckServiceLinkedRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 恢复数据至指定实例
//
// @param request - CloneDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloneDBInstanceResponse
func (client *Client) CloneDBInstanceWithOptions(request *CloneDBInstanceRequest, runtime *dara.RuntimeOptions) (_result *CloneDBInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupId) {
		query["BackupId"] = request.BackupId
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.SrcDbInstanceName) {
		query["SrcDbInstanceName"] = request.SrcDbInstanceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloneDBInstance"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloneDBInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 恢复数据至指定实例
//
// @param request - CloneDBInstanceRequest
//
// @return CloneDBInstanceResponse
func (client *Client) CloneDBInstance(request *CloneDBInstanceRequest) (_result *CloneDBInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloneDBInstanceResponse{}
	_body, _err := client.CloneDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an initial account for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
//	  Before you can use an AnalyticDB for PostgreSQL instance, you must create an initial account for the instance.
//
//		- You can call this operation to create only initial accounts. For information about how to create other types of accounts, see [Create a database account](https://help.aliyun.com/document_detail/50206.html).
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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

	if !dara.IsNil(request.AccountType) {
		query["AccountType"] = request.AccountType
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DatabaseName) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAccount"),
		Version:     dara.String("2016-05-03"),
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
// Creates an initial account for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
//	  Before you can use an AnalyticDB for PostgreSQL instance, you must create an initial account for the instance.
//
//		- You can call this operation to create only initial accounts. For information about how to create other types of accounts, see [Create a database account](https://help.aliyun.com/document_detail/50206.html).
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// 创建备份
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
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBackup"),
		Version:     dara.String("2016-05-03"),
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
// 创建备份
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
// Creates a vector collection.
//
// @param tmpReq - CreateCollectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCollectionResponse
func (client *Client) CreateCollectionWithOptions(tmpReq *CreateCollectionRequest, runtime *dara.RuntimeOptions) (_result *CreateCollectionResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateCollectionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SparseVectorIndexConfig) {
		request.SparseVectorIndexConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SparseVectorIndexConfig, dara.String("SparseVectorIndexConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Collection) {
		query["Collection"] = request.Collection
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Dimension) {
		query["Dimension"] = request.Dimension
	}

	if !dara.IsNil(request.ExternalStorage) {
		query["ExternalStorage"] = request.ExternalStorage
	}

	if !dara.IsNil(request.FullTextRetrievalFields) {
		query["FullTextRetrievalFields"] = request.FullTextRetrievalFields
	}

	if !dara.IsNil(request.HnswEfConstruction) {
		query["HnswEfConstruction"] = request.HnswEfConstruction
	}

	if !dara.IsNil(request.HnswM) {
		query["HnswM"] = request.HnswM
	}

	if !dara.IsNil(request.ManagerAccount) {
		query["ManagerAccount"] = request.ManagerAccount
	}

	if !dara.IsNil(request.ManagerAccountPassword) {
		query["ManagerAccountPassword"] = request.ManagerAccountPassword
	}

	if !dara.IsNil(request.Metadata) {
		query["Metadata"] = request.Metadata
	}

	if !dara.IsNil(request.MetadataIndices) {
		query["MetadataIndices"] = request.MetadataIndices
	}

	if !dara.IsNil(request.Metrics) {
		query["Metrics"] = request.Metrics
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Parser) {
		query["Parser"] = request.Parser
	}

	if !dara.IsNil(request.PqEnable) {
		query["PqEnable"] = request.PqEnable
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SparseVectorIndexConfigShrink) {
		query["SparseVectorIndexConfig"] = request.SparseVectorIndexConfigShrink
	}

	if !dara.IsNil(request.SupportSparse) {
		query["SupportSparse"] = request.SupportSparse
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCollection"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCollectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a vector collection.
//
// @param request - CreateCollectionRequest
//
// @return CreateCollectionResponse
func (client *Client) CreateCollection(request *CreateCollectionRequest) (_result *CreateCollectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCollectionResponse{}
	_body, _err := client.CreateCollectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Instance
//
// Description:
//
// Before using this interface, please make sure you have fully understood the [billing method](https://help.aliyun.com/document_detail/35406.html) and <props="china">[pricing](https://www.aliyun.com/price/product#/gpdb/detail/GreenplumPost)<props="intl">[pricing](https://www.alibabacloud.com/zh/product/hybriddb-postgresql/pricing) of the AnalyticDB for PostgreSQL product.
//
// @param request - CreateDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDBInstanceResponse
func (client *Client) CreateDBInstanceWithOptions(request *CreateDBInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateDBInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AINodeSpecInfos) {
		query["AINodeSpecInfos"] = request.AINodeSpecInfos
	}

	if !dara.IsNil(request.BackupId) {
		query["BackupId"] = request.BackupId
	}

	if !dara.IsNil(request.CacheStorageSize) {
		query["CacheStorageSize"] = request.CacheStorageSize
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CreateSampleData) {
		query["CreateSampleData"] = request.CreateSampleData
	}

	if !dara.IsNil(request.DBInstanceCategory) {
		query["DBInstanceCategory"] = request.DBInstanceCategory
	}

	if !dara.IsNil(request.DBInstanceClass) {
		query["DBInstanceClass"] = request.DBInstanceClass
	}

	if !dara.IsNil(request.DBInstanceDescription) {
		query["DBInstanceDescription"] = request.DBInstanceDescription
	}

	if !dara.IsNil(request.DBInstanceGroupCount) {
		query["DBInstanceGroupCount"] = request.DBInstanceGroupCount
	}

	if !dara.IsNil(request.DBInstanceMode) {
		query["DBInstanceMode"] = request.DBInstanceMode
	}

	if !dara.IsNil(request.DeployMode) {
		query["DeployMode"] = request.DeployMode
	}

	if !dara.IsNil(request.EnableSSL) {
		query["EnableSSL"] = request.EnableSSL
	}

	if !dara.IsNil(request.EncryptionKey) {
		query["EncryptionKey"] = request.EncryptionKey
	}

	if !dara.IsNil(request.EncryptionType) {
		query["EncryptionType"] = request.EncryptionType
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.IdleTime) {
		query["IdleTime"] = request.IdleTime
	}

	if !dara.IsNil(request.InstanceNetworkType) {
		query["InstanceNetworkType"] = request.InstanceNetworkType
	}

	if !dara.IsNil(request.InstanceSpec) {
		query["InstanceSpec"] = request.InstanceSpec
	}

	if !dara.IsNil(request.MasterAISpec) {
		query["MasterAISpec"] = request.MasterAISpec
	}

	if !dara.IsNil(request.MasterCU) {
		query["MasterCU"] = request.MasterCU
	}

	if !dara.IsNil(request.MasterNodeNum) {
		query["MasterNodeNum"] = request.MasterNodeNum
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PrivateIpAddress) {
		query["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !dara.IsNil(request.ProdType) {
		query["ProdType"] = request.ProdType
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

	if !dara.IsNil(request.SegDiskPerformanceLevel) {
		query["SegDiskPerformanceLevel"] = request.SegDiskPerformanceLevel
	}

	if !dara.IsNil(request.SegNodeNum) {
		query["SegNodeNum"] = request.SegNodeNum
	}

	if !dara.IsNil(request.SegStorageType) {
		query["SegStorageType"] = request.SegStorageType
	}

	if !dara.IsNil(request.ServerlessMode) {
		query["ServerlessMode"] = request.ServerlessMode
	}

	if !dara.IsNil(request.ServerlessResource) {
		query["ServerlessResource"] = request.ServerlessResource
	}

	if !dara.IsNil(request.SrcDbInstanceName) {
		query["SrcDbInstanceName"] = request.SrcDbInstanceName
	}

	if !dara.IsNil(request.StandbyVSwitchId) {
		query["StandbyVSwitchId"] = request.StandbyVSwitchId
	}

	if !dara.IsNil(request.StandbyZoneId) {
		query["StandbyZoneId"] = request.StandbyZoneId
	}

	if !dara.IsNil(request.StorageSize) {
		query["StorageSize"] = request.StorageSize
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
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

	if !dara.IsNil(request.VectorConfigurationStatus) {
		query["VectorConfigurationStatus"] = request.VectorConfigurationStatus
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDBInstance"),
		Version:     dara.String("2016-05-03"),
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
// # Create Instance
//
// Description:
//
// Before using this interface, please make sure you have fully understood the [billing method](https://help.aliyun.com/document_detail/35406.html) and <props="china">[pricing](https://www.aliyun.com/price/product#/gpdb/detail/GreenplumPost)<props="intl">[pricing](https://www.alibabacloud.com/zh/product/hybriddb-postgresql/pricing) of the AnalyticDB for PostgreSQL product.
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
// Creates a plan for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
//	  The plan management feature is supported only for pay-as-you-go instances.
//
//		- When you change the compute node specifications or change the number of compute nodes, transient connections may occur. We recommend that you perform these operations during off-peak hours.
//
// Before you call this operation, make sure that you are familiar with the billing of AnalyticDB for PostgreSQL. For more information, see [Billing methods](https://help.aliyun.com/document_detail/35406.html) and [AnalyticDB for PostgreSQL pricing](https://www.alibabacloud.com/zh/product/hybriddb-postgresql/pricing).
//
// @param request - CreateDBInstancePlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDBInstancePlanResponse
func (client *Client) CreateDBInstancePlanWithOptions(request *CreateDBInstancePlanRequest, runtime *dara.RuntimeOptions) (_result *CreateDBInstancePlanResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PlanConfig) {
		query["PlanConfig"] = request.PlanConfig
	}

	if !dara.IsNil(request.PlanDesc) {
		query["PlanDesc"] = request.PlanDesc
	}

	if !dara.IsNil(request.PlanEndDate) {
		query["PlanEndDate"] = request.PlanEndDate
	}

	if !dara.IsNil(request.PlanName) {
		query["PlanName"] = request.PlanName
	}

	if !dara.IsNil(request.PlanScheduleType) {
		query["PlanScheduleType"] = request.PlanScheduleType
	}

	if !dara.IsNil(request.PlanStartDate) {
		query["PlanStartDate"] = request.PlanStartDate
	}

	if !dara.IsNil(request.PlanType) {
		query["PlanType"] = request.PlanType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDBInstancePlan"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDBInstancePlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a plan for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
//	  The plan management feature is supported only for pay-as-you-go instances.
//
//		- When you change the compute node specifications or change the number of compute nodes, transient connections may occur. We recommend that you perform these operations during off-peak hours.
//
// Before you call this operation, make sure that you are familiar with the billing of AnalyticDB for PostgreSQL. For more information, see [Billing methods](https://help.aliyun.com/document_detail/35406.html) and [AnalyticDB for PostgreSQL pricing](https://www.alibabacloud.com/zh/product/hybriddb-postgresql/pricing).
//
// @param request - CreateDBInstancePlanRequest
//
// @return CreateDBInstancePlanResponse
func (client *Client) CreateDBInstancePlan(request *CreateDBInstancePlanRequest) (_result *CreateDBInstancePlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDBInstancePlanResponse{}
	_body, _err := client.CreateDBInstancePlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a resource group.
//
// @param request - CreateDBResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDBResourceGroupResponse
func (client *Client) CreateDBResourceGroupWithOptions(request *CreateDBResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateDBResourceGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceGroupConfig) {
		query["ResourceGroupConfig"] = request.ResourceGroupConfig
	}

	if !dara.IsNil(request.ResourceGroupName) {
		query["ResourceGroupName"] = request.ResourceGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDBResourceGroup"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDBResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a resource group.
//
// @param request - CreateDBResourceGroupRequest
//
// @return CreateDBResourceGroupResponse
func (client *Client) CreateDBResourceGroup(request *CreateDBResourceGroupRequest) (_result *CreateDBResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDBResourceGroupResponse{}
	_body, _err := client.CreateDBResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a document collection.
//
// @param tmpReq - CreateDocumentCollectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDocumentCollectionResponse
func (client *Client) CreateDocumentCollectionWithOptions(tmpReq *CreateDocumentCollectionRequest, runtime *dara.RuntimeOptions) (_result *CreateDocumentCollectionResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateDocumentCollectionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EntityTypes) {
		request.EntityTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EntityTypes, dara.String("EntityTypes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RelationshipTypes) {
		request.RelationshipTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RelationshipTypes, dara.String("RelationshipTypes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Collection) {
		query["Collection"] = request.Collection
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Dimension) {
		query["Dimension"] = request.Dimension
	}

	if !dara.IsNil(request.EmbeddingModel) {
		query["EmbeddingModel"] = request.EmbeddingModel
	}

	if !dara.IsNil(request.EnableGraph) {
		query["EnableGraph"] = request.EnableGraph
	}

	if !dara.IsNil(request.EntityTypesShrink) {
		query["EntityTypes"] = request.EntityTypesShrink
	}

	if !dara.IsNil(request.ExternalStorage) {
		query["ExternalStorage"] = request.ExternalStorage
	}

	if !dara.IsNil(request.FullTextRetrievalFields) {
		query["FullTextRetrievalFields"] = request.FullTextRetrievalFields
	}

	if !dara.IsNil(request.HnswEfConstruction) {
		query["HnswEfConstruction"] = request.HnswEfConstruction
	}

	if !dara.IsNil(request.HnswM) {
		query["HnswM"] = request.HnswM
	}

	if !dara.IsNil(request.LLMModel) {
		query["LLMModel"] = request.LLMModel
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.ManagerAccount) {
		query["ManagerAccount"] = request.ManagerAccount
	}

	if !dara.IsNil(request.ManagerAccountPassword) {
		query["ManagerAccountPassword"] = request.ManagerAccountPassword
	}

	if !dara.IsNil(request.Metadata) {
		query["Metadata"] = request.Metadata
	}

	if !dara.IsNil(request.MetadataIndices) {
		query["MetadataIndices"] = request.MetadataIndices
	}

	if !dara.IsNil(request.Metrics) {
		query["Metrics"] = request.Metrics
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Parser) {
		query["Parser"] = request.Parser
	}

	if !dara.IsNil(request.PqEnable) {
		query["PqEnable"] = request.PqEnable
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RelationshipTypesShrink) {
		query["RelationshipTypes"] = request.RelationshipTypesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDocumentCollection"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDocumentCollectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a document collection.
//
// @param request - CreateDocumentCollectionRequest
//
// @return CreateDocumentCollectionResponse
func (client *Client) CreateDocumentCollection(request *CreateDocumentCollectionRequest) (_result *CreateDocumentCollectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDocumentCollectionResponse{}
	_body, _err := client.CreateDocumentCollectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Install extensions.
//
// @param request - CreateExtensionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExtensionsResponse
func (client *Client) CreateExtensionsWithOptions(request *CreateExtensionsRequest, runtime *dara.RuntimeOptions) (_result *CreateExtensionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBNames) {
		query["DBNames"] = request.DBNames
	}

	if !dara.IsNil(request.Extensions) {
		query["Extensions"] = request.Extensions
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateExtensions"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateExtensionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Install extensions.
//
// @param request - CreateExtensionsRequest
//
// @return CreateExtensionsResponse
func (client *Client) CreateExtensions(request *CreateExtensionsRequest) (_result *CreateExtensionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateExtensionsResponse{}
	_body, _err := client.CreateExtensionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create External Data Service
//
// @param request - CreateExternalDataServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExternalDataServiceResponse
func (client *Client) CreateExternalDataServiceWithOptions(request *CreateExternalDataServiceRequest, runtime *dara.RuntimeOptions) (_result *CreateExternalDataServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceDescription) {
		query["ServiceDescription"] = request.ServiceDescription
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	if !dara.IsNil(request.ServiceSpec) {
		query["ServiceSpec"] = request.ServiceSpec
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateExternalDataService"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateExternalDataServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create External Data Service
//
// @param request - CreateExternalDataServiceRequest
//
// @return CreateExternalDataServiceResponse
func (client *Client) CreateExternalDataService(request *CreateExternalDataServiceRequest) (_result *CreateExternalDataServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateExternalDataServiceResponse{}
	_body, _err := client.CreateExternalDataServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Hadoop data source configuration
//
// @param request - CreateHadoopDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHadoopDataSourceResponse
func (client *Client) CreateHadoopDataSourceWithOptions(request *CreateHadoopDataSourceRequest, runtime *dara.RuntimeOptions) (_result *CreateHadoopDataSourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DataSourceDescription) {
		query["DataSourceDescription"] = request.DataSourceDescription
	}

	if !dara.IsNil(request.DataSourceName) {
		query["DataSourceName"] = request.DataSourceName
	}

	if !dara.IsNil(request.DataSourceType) {
		query["DataSourceType"] = request.DataSourceType
	}

	if !dara.IsNil(request.EmrInstanceId) {
		query["EmrInstanceId"] = request.EmrInstanceId
	}

	if !dara.IsNil(request.HDFSConf) {
		query["HDFSConf"] = request.HDFSConf
	}

	if !dara.IsNil(request.HadoopCoreConf) {
		query["HadoopCoreConf"] = request.HadoopCoreConf
	}

	if !dara.IsNil(request.HadoopCreateType) {
		query["HadoopCreateType"] = request.HadoopCreateType
	}

	if !dara.IsNil(request.HadoopHostsAddress) {
		query["HadoopHostsAddress"] = request.HadoopHostsAddress
	}

	if !dara.IsNil(request.HiveConf) {
		query["HiveConf"] = request.HiveConf
	}

	if !dara.IsNil(request.MapReduceConf) {
		query["MapReduceConf"] = request.MapReduceConf
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.YarnConf) {
		query["YarnConf"] = request.YarnConf
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHadoopDataSource"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHadoopDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Hadoop data source configuration
//
// @param request - CreateHadoopDataSourceRequest
//
// @return CreateHadoopDataSourceResponse
func (client *Client) CreateHadoopDataSource(request *CreateHadoopDataSourceRequest) (_result *CreateHadoopDataSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateHadoopDataSourceResponse{}
	_body, _err := client.CreateHadoopDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建索引
//
// @param request - CreateIndexRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIndexResponse
func (client *Client) CreateIndexWithOptions(request *CreateIndexRequest, runtime *dara.RuntimeOptions) (_result *CreateIndexResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Collection) {
		query["Collection"] = request.Collection
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.IndexConfig) {
		query["IndexConfig"] = request.IndexConfig
	}

	if !dara.IsNil(request.IndexField) {
		query["IndexField"] = request.IndexField
	}

	if !dara.IsNil(request.IndexName) {
		query["IndexName"] = request.IndexName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespacePassword) {
		query["NamespacePassword"] = request.NamespacePassword
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIndex"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建索引
//
// @param request - CreateIndexRequest
//
// @return CreateIndexResponse
func (client *Client) CreateIndex(request *CreateIndexRequest) (_result *CreateIndexResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateIndexResponse{}
	_body, _err := client.CreateIndexWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a JDBC data source.
//
// @param request - CreateJDBCDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateJDBCDataSourceResponse
func (client *Client) CreateJDBCDataSourceWithOptions(request *CreateJDBCDataSourceRequest, runtime *dara.RuntimeOptions) (_result *CreateJDBCDataSourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DataSourceDescription) {
		query["DataSourceDescription"] = request.DataSourceDescription
	}

	if !dara.IsNil(request.DataSourceName) {
		query["DataSourceName"] = request.DataSourceName
	}

	if !dara.IsNil(request.DataSourceType) {
		query["DataSourceType"] = request.DataSourceType
	}

	if !dara.IsNil(request.JDBCConnectionString) {
		query["JDBCConnectionString"] = request.JDBCConnectionString
	}

	if !dara.IsNil(request.JDBCPassword) {
		query["JDBCPassword"] = request.JDBCPassword
	}

	if !dara.IsNil(request.JDBCUserName) {
		query["JDBCUserName"] = request.JDBCUserName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateJDBCDataSource"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateJDBCDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a JDBC data source.
//
// @param request - CreateJDBCDataSourceRequest
//
// @return CreateJDBCDataSourceResponse
func (client *Client) CreateJDBCDataSource(request *CreateJDBCDataSourceRequest) (_result *CreateJDBCDataSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateJDBCDataSourceResponse{}
	_body, _err := client.CreateJDBCDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a vector namespace.
//
// @param request - CreateNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNamespaceResponse
func (client *Client) CreateNamespaceWithOptions(request *CreateNamespaceRequest, runtime *dara.RuntimeOptions) (_result *CreateNamespaceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.ManagerAccount) {
		query["ManagerAccount"] = request.ManagerAccount
	}

	if !dara.IsNil(request.ManagerAccountPassword) {
		query["ManagerAccountPassword"] = request.ManagerAccountPassword
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespacePassword) {
		query["NamespacePassword"] = request.NamespacePassword
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNamespace"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a vector namespace.
//
// @param request - CreateNamespaceRequest
//
// @return CreateNamespaceResponse
func (client *Client) CreateNamespace(request *CreateNamespaceRequest) (_result *CreateNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateNamespaceResponse{}
	_body, _err := client.CreateNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Homogeneous Data Source
//
// @param request - CreateRemoteADBDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRemoteADBDataSourceResponse
func (client *Client) CreateRemoteADBDataSourceWithOptions(request *CreateRemoteADBDataSourceRequest, runtime *dara.RuntimeOptions) (_result *CreateRemoteADBDataSourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataSourceName) {
		query["DataSourceName"] = request.DataSourceName
	}

	if !dara.IsNil(request.LocalDBInstanceId) {
		query["LocalDBInstanceId"] = request.LocalDBInstanceId
	}

	if !dara.IsNil(request.LocalDatabase) {
		query["LocalDatabase"] = request.LocalDatabase
	}

	if !dara.IsNil(request.ManagerUserName) {
		query["ManagerUserName"] = request.ManagerUserName
	}

	if !dara.IsNil(request.ManagerUserPassword) {
		query["ManagerUserPassword"] = request.ManagerUserPassword
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RemoteDBInstanceId) {
		query["RemoteDBInstanceId"] = request.RemoteDBInstanceId
	}

	if !dara.IsNil(request.RemoteDatabase) {
		query["RemoteDatabase"] = request.RemoteDatabase
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	if !dara.IsNil(request.UserPassword) {
		query["UserPassword"] = request.UserPassword
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRemoteADBDataSource"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRemoteADBDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Homogeneous Data Source
//
// @param request - CreateRemoteADBDataSourceRequest
//
// @return CreateRemoteADBDataSourceResponse
func (client *Client) CreateRemoteADBDataSource(request *CreateRemoteADBDataSourceRequest) (_result *CreateRemoteADBDataSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRemoteADBDataSourceResponse{}
	_body, _err := client.CreateRemoteADBDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a sample dataset for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
//	  You can call this operation to create a sample dataset for an AnalyticDB for PostgreSQL instance. Then, you can execute query statements on the sample dataset to experience or test your instance. For more information about query statements, see [Dataset information and query examples](https://help.aliyun.com/document_detail/452277.html).
//
//		- This operation is supported only for AnalyticDB for PostgreSQL V6.3.8.8 to 6.3.8.x, V6.3.10.3, and later.
//
//		- Versions from V6.3.9.0 to V6.3.10.2 are not supported.
//
// @param request - CreateSampleDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSampleDataResponse
func (client *Client) CreateSampleDataWithOptions(request *CreateSampleDataRequest, runtime *dara.RuntimeOptions) (_result *CreateSampleDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSampleData"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSampleDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a sample dataset for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
//	  You can call this operation to create a sample dataset for an AnalyticDB for PostgreSQL instance. Then, you can execute query statements on the sample dataset to experience or test your instance. For more information about query statements, see [Dataset information and query examples](https://help.aliyun.com/document_detail/452277.html).
//
//		- This operation is supported only for AnalyticDB for PostgreSQL V6.3.8.8 to 6.3.8.x, V6.3.10.3, and later.
//
//		- Versions from V6.3.9.0 to V6.3.10.2 are not supported.
//
// @param request - CreateSampleDataRequest
//
// @return CreateSampleDataResponse
func (client *Client) CreateSampleData(request *CreateSampleDataRequest) (_result *CreateSampleDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSampleDataResponse{}
	_body, _err := client.CreateSampleDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an access credential for an AnalyticDB for PostgreSQL instance by using the name and password of a database account.
//
// @param request - CreateSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSecretResponse
func (client *Client) CreateSecretWithOptions(request *CreateSecretRequest, runtime *dara.RuntimeOptions) (_result *CreateSecretResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecretName) {
		query["SecretName"] = request.SecretName
	}

	if !dara.IsNil(request.TestConnection) {
		query["TestConnection"] = request.TestConnection
	}

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSecret"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an access credential for an AnalyticDB for PostgreSQL instance by using the name and password of a database account.
//
// @param request - CreateSecretRequest
//
// @return CreateSecretResponse
func (client *Client) CreateSecret(request *CreateSecretRequest) (_result *CreateSecretResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSecretResponse{}
	_body, _err := client.CreateSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a service-linked role.
//
// @param request - CreateServiceLinkedRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceLinkedRoleResponse
func (client *Client) CreateServiceLinkedRoleWithOptions(request *CreateServiceLinkedRoleRequest, runtime *dara.RuntimeOptions) (_result *CreateServiceLinkedRoleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateServiceLinkedRole"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateServiceLinkedRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a service-linked role.
//
// @param request - CreateServiceLinkedRoleRequest
//
// @return CreateServiceLinkedRoleResponse
func (client *Client) CreateServiceLinkedRole(request *CreateServiceLinkedRoleRequest) (_result *CreateServiceLinkedRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateServiceLinkedRoleResponse{}
	_body, _err := client.CreateServiceLinkedRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create External Data Source Configuration
//
// @param request - CreateStreamingDataServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStreamingDataServiceResponse
func (client *Client) CreateStreamingDataServiceWithOptions(request *CreateStreamingDataServiceRequest, runtime *dara.RuntimeOptions) (_result *CreateStreamingDataServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceDescription) {
		query["ServiceDescription"] = request.ServiceDescription
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	if !dara.IsNil(request.ServiceSpec) {
		query["ServiceSpec"] = request.ServiceSpec
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateStreamingDataService"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStreamingDataServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create External Data Source Configuration
//
// @param request - CreateStreamingDataServiceRequest
//
// @return CreateStreamingDataServiceResponse
func (client *Client) CreateStreamingDataService(request *CreateStreamingDataServiceRequest) (_result *CreateStreamingDataServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateStreamingDataServiceResponse{}
	_body, _err := client.CreateStreamingDataServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create External Data Source Configuration
//
// @param request - CreateStreamingDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStreamingDataSourceResponse
func (client *Client) CreateStreamingDataSourceWithOptions(request *CreateStreamingDataSourceRequest, runtime *dara.RuntimeOptions) (_result *CreateStreamingDataSourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DataSourceConfig) {
		query["DataSourceConfig"] = request.DataSourceConfig
	}

	if !dara.IsNil(request.DataSourceDescription) {
		query["DataSourceDescription"] = request.DataSourceDescription
	}

	if !dara.IsNil(request.DataSourceName) {
		query["DataSourceName"] = request.DataSourceName
	}

	if !dara.IsNil(request.DataSourceType) {
		query["DataSourceType"] = request.DataSourceType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateStreamingDataSource"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStreamingDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create External Data Source Configuration
//
// @param request - CreateStreamingDataSourceRequest
//
// @return CreateStreamingDataSourceResponse
func (client *Client) CreateStreamingDataSource(request *CreateStreamingDataSourceRequest) (_result *CreateStreamingDataSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateStreamingDataSourceResponse{}
	_body, _err := client.CreateStreamingDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create External Data Source Configuration
//
// @param tmpReq - CreateStreamingJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStreamingJobResponse
func (client *Client) CreateStreamingJobWithOptions(tmpReq *CreateStreamingJobRequest, runtime *dara.RuntimeOptions) (_result *CreateStreamingJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateStreamingJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DestColumns) {
		request.DestColumnsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DestColumns, dara.String("DestColumns"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.MatchColumns) {
		request.MatchColumnsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MatchColumns, dara.String("MatchColumns"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SrcColumns) {
		request.SrcColumnsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SrcColumns, dara.String("SrcColumns"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UpdateColumns) {
		request.UpdateColumnsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateColumns, dara.String("UpdateColumns"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Account) {
		query["Account"] = request.Account
	}

	if !dara.IsNil(request.Consistency) {
		query["Consistency"] = request.Consistency
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DataSourceId) {
		query["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.DestColumnsShrink) {
		query["DestColumns"] = request.DestColumnsShrink
	}

	if !dara.IsNil(request.DestDatabase) {
		query["DestDatabase"] = request.DestDatabase
	}

	if !dara.IsNil(request.DestSchema) {
		query["DestSchema"] = request.DestSchema
	}

	if !dara.IsNil(request.DestTable) {
		query["DestTable"] = request.DestTable
	}

	if !dara.IsNil(request.ErrorLimitCount) {
		query["ErrorLimitCount"] = request.ErrorLimitCount
	}

	if !dara.IsNil(request.FallbackOffset) {
		query["FallbackOffset"] = request.FallbackOffset
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.JobConfig) {
		query["JobConfig"] = request.JobConfig
	}

	if !dara.IsNil(request.JobDescription) {
		query["JobDescription"] = request.JobDescription
	}

	if !dara.IsNil(request.JobName) {
		query["JobName"] = request.JobName
	}

	if !dara.IsNil(request.MatchColumnsShrink) {
		query["MatchColumns"] = request.MatchColumnsShrink
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SrcColumnsShrink) {
		query["SrcColumns"] = request.SrcColumnsShrink
	}

	if !dara.IsNil(request.TryRun) {
		query["TryRun"] = request.TryRun
	}

	if !dara.IsNil(request.UpdateColumnsShrink) {
		query["UpdateColumns"] = request.UpdateColumnsShrink
	}

	if !dara.IsNil(request.WriteMode) {
		query["WriteMode"] = request.WriteMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateStreamingJob"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStreamingJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create External Data Source Configuration
//
// @param request - CreateStreamingJobRequest
//
// @return CreateStreamingJobResponse
func (client *Client) CreateStreamingJob(request *CreateStreamingJobRequest) (_result *CreateStreamingJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateStreamingJobResponse{}
	_body, _err := client.CreateStreamingJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建supabase project
//
// @param request - CreateSupabaseProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSupabaseProjectResponse
func (client *Client) CreateSupabaseProjectWithOptions(request *CreateSupabaseProjectRequest, runtime *dara.RuntimeOptions) (_result *CreateSupabaseProjectResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountPassword) {
		query["AccountPassword"] = request.AccountPassword
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DiskPerformanceLevel) {
		query["DiskPerformanceLevel"] = request.DiskPerformanceLevel
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.ProjectSpec) {
		query["ProjectSpec"] = request.ProjectSpec
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityIPList) {
		query["SecurityIPList"] = request.SecurityIPList
	}

	if !dara.IsNil(request.StorageSize) {
		query["StorageSize"] = request.StorageSize
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
		Action:      dara.String("CreateSupabaseProject"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSupabaseProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建supabase project
//
// @param request - CreateSupabaseProjectRequest
//
// @return CreateSupabaseProjectResponse
func (client *Client) CreateSupabaseProject(request *CreateSupabaseProjectRequest) (_result *CreateSupabaseProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSupabaseProjectResponse{}
	_body, _err := client.CreateSupabaseProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Vector Index
//
// @param request - CreateVectorIndexRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVectorIndexResponse
func (client *Client) CreateVectorIndexWithOptions(request *CreateVectorIndexRequest, runtime *dara.RuntimeOptions) (_result *CreateVectorIndexResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Collection) {
		query["Collection"] = request.Collection
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Dimension) {
		query["Dimension"] = request.Dimension
	}

	if !dara.IsNil(request.ExternalStorage) {
		query["ExternalStorage"] = request.ExternalStorage
	}

	if !dara.IsNil(request.HnswEfConstruction) {
		query["HnswEfConstruction"] = request.HnswEfConstruction
	}

	if !dara.IsNil(request.HnswM) {
		query["HnswM"] = request.HnswM
	}

	if !dara.IsNil(request.ManagerAccount) {
		query["ManagerAccount"] = request.ManagerAccount
	}

	if !dara.IsNil(request.ManagerAccountPassword) {
		query["ManagerAccountPassword"] = request.ManagerAccountPassword
	}

	if !dara.IsNil(request.Metrics) {
		query["Metrics"] = request.Metrics
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PqEnable) {
		query["PqEnable"] = request.PqEnable
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVectorIndex"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVectorIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Vector Index
//
// @param request - CreateVectorIndexRequest
//
// @return CreateVectorIndexResponse
func (client *Client) CreateVectorIndex(request *CreateVectorIndexRequest) (_result *CreateVectorIndexResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateVectorIndexResponse{}
	_body, _err := client.CreateVectorIndexWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除数据库账号
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

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAccount"),
		Version:     dara.String("2016-05-03"),
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
// 删除数据库账号
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
// 删除备份
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

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBackup"),
		Version:     dara.String("2016-05-03"),
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
// 删除备份
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
// Deletes a vector collection.
//
// @param request - DeleteCollectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCollectionResponse
func (client *Client) DeleteCollectionWithOptions(request *DeleteCollectionRequest, runtime *dara.RuntimeOptions) (_result *DeleteCollectionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Collection) {
		query["Collection"] = request.Collection
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespacePassword) {
		query["NamespacePassword"] = request.NamespacePassword
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCollection"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCollectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a vector collection.
//
// @param request - DeleteCollectionRequest
//
// @return DeleteCollectionResponse
func (client *Client) DeleteCollection(request *DeleteCollectionRequest) (_result *DeleteCollectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCollectionResponse{}
	_body, _err := client.DeleteCollectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes vector data.
//
// @param request - DeleteCollectionDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCollectionDataResponse
func (client *Client) DeleteCollectionDataWithOptions(request *DeleteCollectionDataRequest, runtime *dara.RuntimeOptions) (_result *DeleteCollectionDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Collection) {
		query["Collection"] = request.Collection
	}

	if !dara.IsNil(request.CollectionData) {
		query["CollectionData"] = request.CollectionData
	}

	if !dara.IsNil(request.CollectionDataFilter) {
		query["CollectionDataFilter"] = request.CollectionDataFilter
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespacePassword) {
		query["NamespacePassword"] = request.NamespacePassword
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCollectionData"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCollectionDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes vector data.
//
// @param request - DeleteCollectionDataRequest
//
// @return DeleteCollectionDataResponse
func (client *Client) DeleteCollectionData(request *DeleteCollectionDataRequest) (_result *DeleteCollectionDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCollectionDataResponse{}
	_body, _err := client.DeleteCollectionDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases a pay-as-you-go AnalyticDB for PostgreSQL instance.
//
// Description:
//
//	  Subscription instances cannot be manually released. They are automatically released when they expire.
//
//		- You can call this operation to release pay-as-you-go instances only when they are in the **Running*	- state.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
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
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDBInstance"),
		Version:     dara.String("2016-05-03"),
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
// Releases a pay-as-you-go AnalyticDB for PostgreSQL instance.
//
// Description:
//
//	  Subscription instances cannot be manually released. They are automatically released when they expire.
//
//		- You can call this operation to release pay-as-you-go instances only when they are in the **Running*	- state.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// Deletes a plan from an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// If you no longer need a plan, you can call this operation to delete the plan. The plan management feature is supported only for AnalyticDB for PostgreSQL instances in Serverless mode.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteDBInstancePlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDBInstancePlanResponse
func (client *Client) DeleteDBInstancePlanWithOptions(request *DeleteDBInstancePlanRequest, runtime *dara.RuntimeOptions) (_result *DeleteDBInstancePlanResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PlanId) {
		query["PlanId"] = request.PlanId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDBInstancePlan"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDBInstancePlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a plan from an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// If you no longer need a plan, you can call this operation to delete the plan. The plan management feature is supported only for AnalyticDB for PostgreSQL instances in Serverless mode.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteDBInstancePlanRequest
//
// @return DeleteDBInstancePlanResponse
func (client *Client) DeleteDBInstancePlan(request *DeleteDBInstancePlanRequest) (_result *DeleteDBInstancePlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDBInstancePlanResponse{}
	_body, _err := client.DeleteDBInstancePlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a resource group.
//
// @param request - DeleteDBResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDBResourceGroupResponse
func (client *Client) DeleteDBResourceGroupWithOptions(request *DeleteDBResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteDBResourceGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceGroupName) {
		query["ResourceGroupName"] = request.ResourceGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDBResourceGroup"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDBResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a resource group.
//
// @param request - DeleteDBResourceGroupRequest
//
// @return DeleteDBResourceGroupResponse
func (client *Client) DeleteDBResourceGroup(request *DeleteDBResourceGroupRequest) (_result *DeleteDBResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDBResourceGroupResponse{}
	_body, _err := client.DeleteDBResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Document
//
// @param request - DeleteDocumentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDocumentResponse
func (client *Client) DeleteDocumentWithOptions(request *DeleteDocumentRequest, runtime *dara.RuntimeOptions) (_result *DeleteDocumentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Collection) {
		query["Collection"] = request.Collection
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespacePassword) {
		query["NamespacePassword"] = request.NamespacePassword
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDocument"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDocumentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Document
//
// @param request - DeleteDocumentRequest
//
// @return DeleteDocumentResponse
func (client *Client) DeleteDocument(request *DeleteDocumentRequest) (_result *DeleteDocumentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDocumentResponse{}
	_body, _err := client.DeleteDocumentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Knowledge Base
//
// @param request - DeleteDocumentCollectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDocumentCollectionResponse
func (client *Client) DeleteDocumentCollectionWithOptions(request *DeleteDocumentCollectionRequest, runtime *dara.RuntimeOptions) (_result *DeleteDocumentCollectionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Collection) {
		query["Collection"] = request.Collection
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespacePassword) {
		query["NamespacePassword"] = request.NamespacePassword
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDocumentCollection"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDocumentCollectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Knowledge Base
//
// @param request - DeleteDocumentCollectionRequest
//
// @return DeleteDocumentCollectionResponse
func (client *Client) DeleteDocumentCollection(request *DeleteDocumentCollectionRequest) (_result *DeleteDocumentCollectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDocumentCollectionResponse{}
	_body, _err := client.DeleteDocumentCollectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uninstall an extension.
//
// @param request - DeleteExtensionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExtensionResponse
func (client *Client) DeleteExtensionWithOptions(request *DeleteExtensionRequest, runtime *dara.RuntimeOptions) (_result *DeleteExtensionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBNames) {
		query["DBNames"] = request.DBNames
	}

	if !dara.IsNil(request.Extension) {
		query["Extension"] = request.Extension
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteExtension"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteExtensionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uninstall an extension.
//
// @param request - DeleteExtensionRequest
//
// @return DeleteExtensionResponse
func (client *Client) DeleteExtension(request *DeleteExtensionRequest) (_result *DeleteExtensionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteExtensionResponse{}
	_body, _err := client.DeleteExtensionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete External Data Service
//
// @param request - DeleteExternalDataServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExternalDataServiceResponse
func (client *Client) DeleteExternalDataServiceWithOptions(request *DeleteExternalDataServiceRequest, runtime *dara.RuntimeOptions) (_result *DeleteExternalDataServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteExternalDataService"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteExternalDataServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete External Data Service
//
// @param request - DeleteExternalDataServiceRequest
//
// @return DeleteExternalDataServiceResponse
func (client *Client) DeleteExternalDataService(request *DeleteExternalDataServiceRequest) (_result *DeleteExternalDataServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteExternalDataServiceResponse{}
	_body, _err := client.DeleteExternalDataServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除hadoop数据源
//
// @param request - DeleteHadoopDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHadoopDataSourceResponse
func (client *Client) DeleteHadoopDataSourceWithOptions(request *DeleteHadoopDataSourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteHadoopDataSourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DataSourceId) {
		query["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHadoopDataSource"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHadoopDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除hadoop数据源
//
// @param request - DeleteHadoopDataSourceRequest
//
// @return DeleteHadoopDataSourceResponse
func (client *Client) DeleteHadoopDataSource(request *DeleteHadoopDataSourceRequest) (_result *DeleteHadoopDataSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteHadoopDataSourceResponse{}
	_body, _err := client.DeleteHadoopDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除索引
//
// @param request - DeleteIndexRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIndexResponse
func (client *Client) DeleteIndexWithOptions(request *DeleteIndexRequest, runtime *dara.RuntimeOptions) (_result *DeleteIndexResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Collection) {
		query["Collection"] = request.Collection
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.IndexName) {
		query["IndexName"] = request.IndexName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespacePassword) {
		query["NamespacePassword"] = request.NamespacePassword
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteIndex"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除索引
//
// @param request - DeleteIndexRequest
//
// @return DeleteIndexResponse
func (client *Client) DeleteIndex(request *DeleteIndexRequest) (_result *DeleteIndexResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteIndexResponse{}
	_body, _err := client.DeleteIndexWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete JDBC data source
//
// @param request - DeleteJDBCDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteJDBCDataSourceResponse
func (client *Client) DeleteJDBCDataSourceWithOptions(request *DeleteJDBCDataSourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteJDBCDataSourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DataSourceId) {
		query["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteJDBCDataSource"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteJDBCDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete JDBC data source
//
// @param request - DeleteJDBCDataSourceRequest
//
// @return DeleteJDBCDataSourceResponse
func (client *Client) DeleteJDBCDataSource(request *DeleteJDBCDataSourceRequest) (_result *DeleteJDBCDataSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteJDBCDataSourceResponse{}
	_body, _err := client.DeleteJDBCDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a namespace.
//
// @param request - DeleteNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNamespaceResponse
func (client *Client) DeleteNamespaceWithOptions(request *DeleteNamespaceRequest, runtime *dara.RuntimeOptions) (_result *DeleteNamespaceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.ManagerAccount) {
		query["ManagerAccount"] = request.ManagerAccount
	}

	if !dara.IsNil(request.ManagerAccountPassword) {
		query["ManagerAccountPassword"] = request.ManagerAccountPassword
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNamespace"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a namespace.
//
// @param request - DeleteNamespaceRequest
//
// @return DeleteNamespaceResponse
func (client *Client) DeleteNamespace(request *DeleteNamespaceRequest) (_result *DeleteNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteNamespaceResponse{}
	_body, _err := client.DeleteNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a remote AnalyticDB data source.
//
// @param request - DeleteRemoteADBDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRemoteADBDataSourceResponse
func (client *Client) DeleteRemoteADBDataSourceWithOptions(request *DeleteRemoteADBDataSourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteRemoteADBDataSourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataSourceId) {
		query["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.LocalDBInstanceId) {
		query["LocalDBInstanceId"] = request.LocalDBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRemoteADBDataSource"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRemoteADBDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a remote AnalyticDB data source.
//
// @param request - DeleteRemoteADBDataSourceRequest
//
// @return DeleteRemoteADBDataSourceResponse
func (client *Client) DeleteRemoteADBDataSource(request *DeleteRemoteADBDataSourceRequest) (_result *DeleteRemoteADBDataSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRemoteADBDataSourceResponse{}
	_body, _err := client.DeleteRemoteADBDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the access credentials of an AnalyticDB for PostgreSQL instance.
//
// @param request - DeleteSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSecretResponse
func (client *Client) DeleteSecretWithOptions(request *DeleteSecretRequest, runtime *dara.RuntimeOptions) (_result *DeleteSecretResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecretArn) {
		query["SecretArn"] = request.SecretArn
	}

	if !dara.IsNil(request.SecretName) {
		query["SecretName"] = request.SecretName
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSecret"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the access credentials of an AnalyticDB for PostgreSQL instance.
//
// @param request - DeleteSecretRequest
//
// @return DeleteSecretResponse
func (client *Client) DeleteSecret(request *DeleteSecretRequest) (_result *DeleteSecretResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSecretResponse{}
	_body, _err := client.DeleteSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a real-time data service.
//
// @param request - DeleteStreamingDataServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStreamingDataServiceResponse
func (client *Client) DeleteStreamingDataServiceWithOptions(request *DeleteStreamingDataServiceRequest, runtime *dara.RuntimeOptions) (_result *DeleteStreamingDataServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteStreamingDataService"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteStreamingDataServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a real-time data service.
//
// @param request - DeleteStreamingDataServiceRequest
//
// @return DeleteStreamingDataServiceResponse
func (client *Client) DeleteStreamingDataService(request *DeleteStreamingDataServiceRequest) (_result *DeleteStreamingDataServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteStreamingDataServiceResponse{}
	_body, _err := client.DeleteStreamingDataServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a real-time data source.
//
// @param request - DeleteStreamingDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStreamingDataSourceResponse
func (client *Client) DeleteStreamingDataSourceWithOptions(request *DeleteStreamingDataSourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteStreamingDataSourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DataSourceId) {
		query["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteStreamingDataSource"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteStreamingDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a real-time data source.
//
// @param request - DeleteStreamingDataSourceRequest
//
// @return DeleteStreamingDataSourceResponse
func (client *Client) DeleteStreamingDataSource(request *DeleteStreamingDataSourceRequest) (_result *DeleteStreamingDataSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteStreamingDataSourceResponse{}
	_body, _err := client.DeleteStreamingDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a real-time data service job.
//
// @param request - DeleteStreamingJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStreamingJobResponse
func (client *Client) DeleteStreamingJobWithOptions(request *DeleteStreamingJobRequest, runtime *dara.RuntimeOptions) (_result *DeleteStreamingJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteStreamingJob"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteStreamingJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a real-time data service job.
//
// @param request - DeleteStreamingJobRequest
//
// @return DeleteStreamingJobResponse
func (client *Client) DeleteStreamingJob(request *DeleteStreamingJobRequest) (_result *DeleteStreamingJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteStreamingJobResponse{}
	_body, _err := client.DeleteStreamingJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除Supabase实例
//
// @param request - DeleteSupabaseProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSupabaseProjectResponse
func (client *Client) DeleteSupabaseProjectWithOptions(request *DeleteSupabaseProjectRequest, runtime *dara.RuntimeOptions) (_result *DeleteSupabaseProjectResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      dara.String("DeleteSupabaseProject"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSupabaseProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除Supabase实例
//
// @param request - DeleteSupabaseProjectRequest
//
// @return DeleteSupabaseProjectResponse
func (client *Client) DeleteSupabaseProject(request *DeleteSupabaseProjectRequest) (_result *DeleteSupabaseProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSupabaseProjectResponse{}
	_body, _err := client.DeleteSupabaseProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a vector index.
//
// @param request - DeleteVectorIndexRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVectorIndexResponse
func (client *Client) DeleteVectorIndexWithOptions(request *DeleteVectorIndexRequest, runtime *dara.RuntimeOptions) (_result *DeleteVectorIndexResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Collection) {
		query["Collection"] = request.Collection
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.ManagerAccount) {
		query["ManagerAccount"] = request.ManagerAccount
	}

	if !dara.IsNil(request.ManagerAccountPassword) {
		query["ManagerAccountPassword"] = request.ManagerAccountPassword
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVectorIndex"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVectorIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a vector index.
//
// @param request - DeleteVectorIndexRequest
//
// @return DeleteVectorIndexResponse
func (client *Client) DeleteVectorIndex(request *DeleteVectorIndexRequest) (_result *DeleteVectorIndexResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteVectorIndexResponse{}
	_body, _err := client.DeleteVectorIndexWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about database accounts for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// This operation is called to query the information of the privileged account in an AnalyticDB for PostgreSQL instance, such as its state, description, and the instance.
//
// ## Limit
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered and may affect your business. We recommend that you take note of the limit when you call this operation.
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

	if !dara.IsNil(request.AccountType) {
		query["AccountType"] = request.AccountType
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAccounts"),
		Version:     dara.String("2016-05-03"),
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
// Queries the information about database accounts for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// This operation is called to query the information of the privileged account in an AnalyticDB for PostgreSQL instance, such as its state, description, and the instance.
//
// ## Limit
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered and may affect your business. We recommend that you take note of the limit when you call this operation.
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
// Queries active SQL records.
//
// @param request - DescribeActiveSQLRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeActiveSQLRecordsResponse
func (client *Client) DescribeActiveSQLRecordsWithOptions(request *DescribeActiveSQLRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeActiveSQLRecordsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Database) {
		query["Database"] = request.Database
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxDuration) {
		query["MaxDuration"] = request.MaxDuration
	}

	if !dara.IsNil(request.MinDuration) {
		query["MinDuration"] = request.MinDuration
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.User) {
		query["User"] = request.User
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeActiveSQLRecords"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeActiveSQLRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries active SQL records.
//
// @param request - DescribeActiveSQLRecordsRequest
//
// @return DescribeActiveSQLRecordsResponse
func (client *Client) DescribeActiveSQLRecords(request *DescribeActiveSQLRecordsRequest) (_result *DescribeActiveSQLRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeActiveSQLRecordsResponse{}
	_body, _err := client.DescribeActiveSQLRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about available resources of AnalyticDB for PostgreSQL.
//
// Description:
//
// When you create an AnalyticDB for PostgreSQL instance, you can call this operation to query the available resources within a zone.
//
// @param request - DescribeAvailableResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAvailableResourcesResponse
func (client *Client) DescribeAvailableResourcesWithOptions(request *DescribeAvailableResourcesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAvailableResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAvailableResources"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAvailableResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about available resources of AnalyticDB for PostgreSQL.
//
// Description:
//
// When you create an AnalyticDB for PostgreSQL instance, you can call this operation to query the available resources within a zone.
//
// @param request - DescribeAvailableResourcesRequest
//
// @return DescribeAvailableResourcesResponse
func (client *Client) DescribeAvailableResources(request *DescribeAvailableResourcesRequest) (_result *DescribeAvailableResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAvailableResourcesResponse{}
	_body, _err := client.DescribeAvailableResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取备份任务详情
//
// @param request - DescribeBackupJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupJobResponse
func (client *Client) DescribeBackupJobWithOptions(request *DescribeBackupJobRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupJobId) {
		query["BackupJobId"] = request.BackupJobId
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupJob"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取备份任务详情
//
// @param request - DescribeBackupJobRequest
//
// @return DescribeBackupJobResponse
func (client *Client) DescribeBackupJob(request *DescribeBackupJobRequest) (_result *DescribeBackupJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBackupJobResponse{}
	_body, _err := client.DescribeBackupJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the backup policy of an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to query the backup settings of an AnalyticDB for PostgreSQL instance in elastic storage mode. Periodically backing data can prevent data loss. For more information about how to modify backup policies, see [ModifyBackupPolicy](https://help.aliyun.com/document_detail/210095.html).
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupPolicy"),
		Version:     dara.String("2016-05-03"),
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
// Queries the backup policy of an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to query the backup settings of an AnalyticDB for PostgreSQL instance in elastic storage mode. Periodically backing data can prevent data loss. For more information about how to modify backup policies, see [ModifyBackupPolicy](https://help.aliyun.com/document_detail/210095.html).
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// Queries the information about a vector collection.
//
// @param request - DescribeCollectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCollectionResponse
func (client *Client) DescribeCollectionWithOptions(request *DescribeCollectionRequest, runtime *dara.RuntimeOptions) (_result *DescribeCollectionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Collection) {
		query["Collection"] = request.Collection
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespacePassword) {
		query["NamespacePassword"] = request.NamespacePassword
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCollection"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCollectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a vector collection.
//
// @param request - DescribeCollectionRequest
//
// @return DescribeCollectionResponse
func (client *Client) DescribeCollection(request *DescribeCollectionRequest) (_result *DescribeCollectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCollectionResponse{}
	_body, _err := client.DescribeCollectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取创建索引任务
//
// @param request - DescribeCreateIndexJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCreateIndexJobResponse
func (client *Client) DescribeCreateIndexJobWithOptions(request *DescribeCreateIndexJobRequest, runtime *dara.RuntimeOptions) (_result *DescribeCreateIndexJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Collection) {
		query["Collection"] = request.Collection
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespacePassword) {
		query["NamespacePassword"] = request.NamespacePassword
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCreateIndexJob"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCreateIndexJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取创建索引任务
//
// @param request - DescribeCreateIndexJobRequest
//
// @return DescribeCreateIndexJobResponse
func (client *Client) DescribeCreateIndexJob(request *DescribeCreateIndexJobRequest) (_result *DescribeCreateIndexJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCreateIndexJobResponse{}
	_body, _err := client.DescribeCreateIndexJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of nodes in an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// ##
//
// You can call this operation to query the information about coordinator and compute nodes in an AnalyticDB for PostgreSQL instance in elastic storage mode or Serverless mode.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDBClusterNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterNodeResponse
func (client *Client) DescribeDBClusterNodeWithOptions(request *DescribeDBClusterNodeRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClusterNodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.NodeType) {
		query["NodeType"] = request.NodeType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBClusterNode"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBClusterNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of nodes in an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// ##
//
// You can call this operation to query the information about coordinator and compute nodes in an AnalyticDB for PostgreSQL instance in elastic storage mode or Serverless mode.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDBClusterNodeRequest
//
// @return DescribeDBClusterNodeResponse
func (client *Client) DescribeDBClusterNode(request *DescribeDBClusterNodeRequest) (_result *DescribeDBClusterNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBClusterNodeResponse{}
	_body, _err := client.DescribeDBClusterNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about performance metrics of an AnalyticDB for PostgreSQL instance within a time range.
//
// Description:
//
// You can query monitoring information only within the last 30 days.
//
// @param request - DescribeDBClusterPerformanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterPerformanceResponse
func (client *Client) DescribeDBClusterPerformanceWithOptions(request *DescribeDBClusterPerformanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClusterPerformanceResponse, _err error) {
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

	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.NodeType) {
		query["NodeType"] = request.NodeType
	}

	if !dara.IsNil(request.Nodes) {
		query["Nodes"] = request.Nodes
	}

	if !dara.IsNil(request.ResourceGroupName) {
		query["ResourceGroupName"] = request.ResourceGroupName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBClusterPerformance"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBClusterPerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about performance metrics of an AnalyticDB for PostgreSQL instance within a time range.
//
// Description:
//
// You can query monitoring information only within the last 30 days.
//
// @param request - DescribeDBClusterPerformanceRequest
//
// @return DescribeDBClusterPerformanceResponse
func (client *Client) DescribeDBClusterPerformance(request *DescribeDBClusterPerformanceRequest) (_result *DescribeDBClusterPerformanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBClusterPerformanceResponse{}
	_body, _err := client.DescribeDBClusterPerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query detailed information about the instance.
//
// Description:
//
// ## Usage Instructions
//
// This interface is generally used to view information such as the specifications, network type, and instance status of AnalyticDB for PostgreSQL instances.
//
// ## QPS Limitation
//
// The default single-user QPS limit for this interface is 1000 times/second. If the limit is exceeded, API calls will be throttled, which may affect your business. Please use it reasonably.
//
// <props="china">The QPS in this document is only a default reference value. For accurate information, please refer to the [API Rate Quota List](https://quotas.console.aliyun.com/flow-control-products/gpdb/quotas).
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
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstanceAttribute"),
		Version:     dara.String("2016-05-03"),
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
// Query detailed information about the instance.
//
// Description:
//
// ## Usage Instructions
//
// This interface is generally used to view information such as the specifications, network type, and instance status of AnalyticDB for PostgreSQL instances.
//
// ## QPS Limitation
//
// The default single-user QPS limit for this interface is 1000 times/second. If the limit is exceeded, API calls will be throttled, which may affect your business. Please use it reasonably.
//
// <props="china">The QPS in this document is only a default reference value. For accurate information, please refer to the [API Rate Quota List](https://quotas.console.aliyun.com/flow-control-products/gpdb/quotas).
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
// Queries the information about data bloat for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to query the details of data bloat on an AnalyticDB for PostgreSQL instance in elastic storage mode. The minor version of the instance must be V6.3.10.1 or later. For more information about how to view and update the minor version of an instance, see [View the minor engine version](https://help.aliyun.com/document_detail/277424.html) and [Update the minor engine version](https://help.aliyun.com/document_detail/139271.html).
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDBInstanceDataBloatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceDataBloatResponse
func (client *Client) DescribeDBInstanceDataBloatWithOptions(request *DescribeDBInstanceDataBloatRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceDataBloatResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Database) {
		query["Database"] = request.Database
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
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
		Action:      dara.String("DescribeDBInstanceDataBloat"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceDataBloatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about data bloat for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to query the details of data bloat on an AnalyticDB for PostgreSQL instance in elastic storage mode. The minor version of the instance must be V6.3.10.1 or later. For more information about how to view and update the minor version of an instance, see [View the minor engine version](https://help.aliyun.com/document_detail/277424.html) and [Update the minor engine version](https://help.aliyun.com/document_detail/139271.html).
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDBInstanceDataBloatRequest
//
// @return DescribeDBInstanceDataBloatResponse
func (client *Client) DescribeDBInstanceDataBloat(request *DescribeDBInstanceDataBloatRequest) (_result *DescribeDBInstanceDataBloatResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBInstanceDataBloatResponse{}
	_body, _err := client.DescribeDBInstanceDataBloatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about data skew for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// To prevent data skew from affecting your database service, you can call this operation to query the details about data skew on an AnalyticDB for PostgreSQL instance.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDBInstanceDataSkewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceDataSkewResponse
func (client *Client) DescribeDBInstanceDataSkewWithOptions(request *DescribeDBInstanceDataSkewRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceDataSkewResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Database) {
		query["Database"] = request.Database
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
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
		Action:      dara.String("DescribeDBInstanceDataSkew"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceDataSkewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about data skew for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// To prevent data skew from affecting your database service, you can call this operation to query the details about data skew on an AnalyticDB for PostgreSQL instance.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDBInstanceDataSkewRequest
//
// @return DescribeDBInstanceDataSkewResponse
func (client *Client) DescribeDBInstanceDataSkew(request *DescribeDBInstanceDataSkewRequest) (_result *DescribeDBInstanceDataSkewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBInstanceDataSkewResponse{}
	_body, _err := client.DescribeDBInstanceDataSkewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about nodes in an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to query the distribution and states of coordinator and compute nodes in an AnalyticDB for PostgreSQL instance.
//
// @param request - DescribeDBInstanceDiagnosisSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceDiagnosisSummaryResponse
func (client *Client) DescribeDBInstanceDiagnosisSummaryWithOptions(request *DescribeDBInstanceDiagnosisSummaryRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceDiagnosisSummaryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

	if !dara.IsNil(request.RolePreferd) {
		query["RolePreferd"] = request.RolePreferd
	}

	if !dara.IsNil(request.StartStatus) {
		query["StartStatus"] = request.StartStatus
	}

	if !dara.IsNil(request.SyncMode) {
		query["SyncMode"] = request.SyncMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstanceDiagnosisSummary"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceDiagnosisSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about nodes in an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to query the distribution and states of coordinator and compute nodes in an AnalyticDB for PostgreSQL instance.
//
// @param request - DescribeDBInstanceDiagnosisSummaryRequest
//
// @return DescribeDBInstanceDiagnosisSummaryResponse
func (client *Client) DescribeDBInstanceDiagnosisSummary(request *DescribeDBInstanceDiagnosisSummaryRequest) (_result *DescribeDBInstanceDiagnosisSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBInstanceDiagnosisSummaryResponse{}
	_body, _err := client.DescribeDBInstanceDiagnosisSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the error logs of an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to query the error logs of an AnalyticDB for PostgreSQL instance.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDBInstanceErrorLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceErrorLogResponse
func (client *Client) DescribeDBInstanceErrorLogWithOptions(request *DescribeDBInstanceErrorLogRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceErrorLogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Database) {
		query["Database"] = request.Database
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Host) {
		query["Host"] = request.Host
	}

	if !dara.IsNil(request.Keywords) {
		query["Keywords"] = request.Keywords
	}

	if !dara.IsNil(request.LogLevel) {
		query["LogLevel"] = request.LogLevel
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

	if !dara.IsNil(request.User) {
		query["User"] = request.User
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstanceErrorLog"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceErrorLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the error logs of an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to query the error logs of an AnalyticDB for PostgreSQL instance.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDBInstanceErrorLogRequest
//
// @return DescribeDBInstanceErrorLogResponse
func (client *Client) DescribeDBInstanceErrorLog(request *DescribeDBInstanceErrorLogRequest) (_result *DescribeDBInstanceErrorLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBInstanceErrorLogResponse{}
	_body, _err := client.DescribeDBInstanceErrorLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the whitelists of IP addresses that are allowed to access an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to query the whitelists of IP addresses that are allowed to access an AnalyticDB for PostgreSQL instance.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDBInstanceIPArrayListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceIPArrayListResponse
func (client *Client) DescribeDBInstanceIPArrayListWithOptions(request *DescribeDBInstanceIPArrayListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceIPArrayListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceIPArrayName) {
		query["DBInstanceIPArrayName"] = request.DBInstanceIPArrayName
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstanceIPArrayList"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceIPArrayListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the whitelists of IP addresses that are allowed to access an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to query the whitelists of IP addresses that are allowed to access an AnalyticDB for PostgreSQL instance.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDBInstanceIPArrayListRequest
//
// @return DescribeDBInstanceIPArrayListResponse
func (client *Client) DescribeDBInstanceIPArrayList(request *DescribeDBInstanceIPArrayListRequest) (_result *DescribeDBInstanceIPArrayListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBInstanceIPArrayListResponse{}
	_body, _err := client.DescribeDBInstanceIPArrayListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the index usage of an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// Appropriate indexes can accelerate database queries. You can call this operation to query the index usage of an AnalyticDB for PostgreSQL instance.
//
// This operation is available only for instances of V6.3.10.1 or later in elastic storage mode. For information about how to view and update the minor version of an instance, see [View the minor engine version](https://help.aliyun.com/document_detail/277424.html) and [Update the minor engine version](https://help.aliyun.com/document_detail/139271.html).
//
// @param request - DescribeDBInstanceIndexUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceIndexUsageResponse
func (client *Client) DescribeDBInstanceIndexUsageWithOptions(request *DescribeDBInstanceIndexUsageRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceIndexUsageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Database) {
		query["Database"] = request.Database
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
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
		Action:      dara.String("DescribeDBInstanceIndexUsage"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceIndexUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the index usage of an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// Appropriate indexes can accelerate database queries. You can call this operation to query the index usage of an AnalyticDB for PostgreSQL instance.
//
// This operation is available only for instances of V6.3.10.1 or later in elastic storage mode. For information about how to view and update the minor version of an instance, see [View the minor engine version](https://help.aliyun.com/document_detail/277424.html) and [Update the minor engine version](https://help.aliyun.com/document_detail/139271.html).
//
// @param request - DescribeDBInstanceIndexUsageRequest
//
// @return DescribeDBInstanceIndexUsageResponse
func (client *Client) DescribeDBInstanceIndexUsage(request *DescribeDBInstanceIndexUsageRequest) (_result *DescribeDBInstanceIndexUsageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBInstanceIndexUsageResponse{}
	_body, _err := client.DescribeDBInstanceIndexUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the connection information of an instance.
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
	if !dara.IsNil(request.ConnectionString) {
		query["ConnectionString"] = request.ConnectionString
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstanceNetInfo"),
		Version:     dara.String("2016-05-03"),
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
// Queries the connection information of an instance.
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
// Queries the information about performance metrics of an AnalyticDB for PostgreSQL instance within a time range.
//
// @param request - DescribeDBInstancePerformanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstancePerformanceResponse
func (client *Client) DescribeDBInstancePerformanceWithOptions(request *DescribeDBInstancePerformanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstancePerformanceResponse, _err error) {
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

	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstancePerformance"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstancePerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about performance metrics of an AnalyticDB for PostgreSQL instance within a time range.
//
// @param request - DescribeDBInstancePerformanceRequest
//
// @return DescribeDBInstancePerformanceResponse
func (client *Client) DescribeDBInstancePerformance(request *DescribeDBInstancePerformanceRequest) (_result *DescribeDBInstancePerformanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBInstancePerformanceResponse{}
	_body, _err := client.DescribeDBInstancePerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about plans for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to query the details of plans for an AnalyticDB for PostgreSQL instance in Serverless mode.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDBInstancePlansRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstancePlansResponse
func (client *Client) DescribeDBInstancePlansWithOptions(request *DescribeDBInstancePlansRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstancePlansResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PlanCreateDate) {
		query["PlanCreateDate"] = request.PlanCreateDate
	}

	if !dara.IsNil(request.PlanDesc) {
		query["PlanDesc"] = request.PlanDesc
	}

	if !dara.IsNil(request.PlanId) {
		query["PlanId"] = request.PlanId
	}

	if !dara.IsNil(request.PlanScheduleType) {
		query["PlanScheduleType"] = request.PlanScheduleType
	}

	if !dara.IsNil(request.PlanType) {
		query["PlanType"] = request.PlanType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstancePlans"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstancePlansResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about plans for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to query the details of plans for an AnalyticDB for PostgreSQL instance in Serverless mode.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDBInstancePlansRequest
//
// @return DescribeDBInstancePlansResponse
func (client *Client) DescribeDBInstancePlans(request *DescribeDBInstancePlansRequest) (_result *DescribeDBInstancePlansResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBInstancePlansResponse{}
	_body, _err := client.DescribeDBInstancePlansWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the SSL information about an AnalyticDB for PostgreSQL instance.
//
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
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstanceSSL"),
		Version:     dara.String("2016-05-03"),
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

// Summary:
//
// Queries the SSL information about an AnalyticDB for PostgreSQL instance.
//
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

// Summary:
//
// Queries the maximum performance of an AnalyticDB for PostgreSQL instance.
//
// @param request - DescribeDBInstanceSupportMaxPerformanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceSupportMaxPerformanceResponse
func (client *Client) DescribeDBInstanceSupportMaxPerformanceWithOptions(request *DescribeDBInstanceSupportMaxPerformanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceSupportMaxPerformanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstanceSupportMaxPerformance"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceSupportMaxPerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the maximum performance of an AnalyticDB for PostgreSQL instance.
//
// @param request - DescribeDBInstanceSupportMaxPerformanceRequest
//
// @return DescribeDBInstanceSupportMaxPerformanceResponse
func (client *Client) DescribeDBInstanceSupportMaxPerformance(request *DescribeDBInstanceSupportMaxPerformanceRequest) (_result *DescribeDBInstanceSupportMaxPerformanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBInstanceSupportMaxPerformanceResponse{}
	_body, _err := client.DescribeDBInstanceSupportMaxPerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of AnalyticDB for PostgreSQL instances.
//
// Description:
//
// ##
//
// You can call this operation to query the instance types, network types, and states of AnalyticDB for PostgreSQL instances within a region.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param tmpReq - DescribeDBInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstancesResponse
func (client *Client) DescribeDBInstancesWithOptions(tmpReq *DescribeDBInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstancesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DescribeDBInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DBInstanceCategories) {
		request.DBInstanceCategoriesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DBInstanceCategories, dara.String("DBInstanceCategories"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.DBInstanceModes) {
		request.DBInstanceModesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DBInstanceModes, dara.String("DBInstanceModes"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.DBInstanceStatuses) {
		request.DBInstanceStatusesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DBInstanceStatuses, dara.String("DBInstanceStatuses"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.InstanceDeployTypes) {
		request.InstanceDeployTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceDeployTypes, dara.String("InstanceDeployTypes"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceCategoriesShrink) {
		query["DBInstanceCategories"] = request.DBInstanceCategoriesShrink
	}

	if !dara.IsNil(request.DBInstanceDescription) {
		query["DBInstanceDescription"] = request.DBInstanceDescription
	}

	if !dara.IsNil(request.DBInstanceIds) {
		query["DBInstanceIds"] = request.DBInstanceIds
	}

	if !dara.IsNil(request.DBInstanceModesShrink) {
		query["DBInstanceModes"] = request.DBInstanceModesShrink
	}

	if !dara.IsNil(request.DBInstanceStatusesShrink) {
		query["DBInstanceStatuses"] = request.DBInstanceStatusesShrink
	}

	if !dara.IsNil(request.InstanceDeployTypesShrink) {
		query["InstanceDeployTypes"] = request.InstanceDeployTypesShrink
	}

	if !dara.IsNil(request.InstanceNetworkType) {
		query["InstanceNetworkType"] = request.InstanceNetworkType
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstances"),
		Version:     dara.String("2016-05-03"),
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
// Queries a list of AnalyticDB for PostgreSQL instances.
//
// Description:
//
// ##
//
// You can call this operation to query the instance types, network types, and states of AnalyticDB for PostgreSQL instances within a region.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// Queries the information about resource groups.
//
// @param request - DescribeDBResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBResourceGroupResponse
func (client *Client) DescribeDBResourceGroupWithOptions(request *DescribeDBResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBResourceGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceGroupName) {
		query["ResourceGroupName"] = request.ResourceGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBResourceGroup"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about resource groups.
//
// @param request - DescribeDBResourceGroupRequest
//
// @return DescribeDBResourceGroupResponse
func (client *Client) DescribeDBResourceGroup(request *DescribeDBResourceGroupRequest) (_result *DescribeDBResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBResourceGroupResponse{}
	_body, _err := client.DescribeDBResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the resource management mode of an AnalyticDB for PostgreSQL instance.
//
// @param request - DescribeDBResourceManagementModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBResourceManagementModeResponse
func (client *Client) DescribeDBResourceManagementModeWithOptions(request *DescribeDBResourceManagementModeRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBResourceManagementModeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBResourceManagementMode"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBResourceManagementModeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resource management mode of an AnalyticDB for PostgreSQL instance.
//
// @param request - DescribeDBResourceManagementModeRequest
//
// @return DescribeDBResourceManagementModeResponse
func (client *Client) DescribeDBResourceManagementMode(request *DescribeDBResourceManagementModeRequest) (_result *DescribeDBResourceManagementModeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBResourceManagementModeResponse{}
	_body, _err := client.DescribeDBResourceManagementModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about minor versions of AnalyticDB for PostgreSQL instances.
//
// @param request - DescribeDBVersionInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBVersionInfosResponse
func (client *Client) DescribeDBVersionInfosWithOptions(request *DescribeDBVersionInfosRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBVersionInfosResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceMode) {
		query["DBInstanceMode"] = request.DBInstanceMode
	}

	if !dara.IsNil(request.DBVersion) {
		query["DBVersion"] = request.DBVersion
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBVersionInfos"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBVersionInfosResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about minor versions of AnalyticDB for PostgreSQL instances.
//
// @param request - DescribeDBVersionInfosRequest
//
// @return DescribeDBVersionInfosResponse
func (client *Client) DescribeDBVersionInfos(request *DescribeDBVersionInfosRequest) (_result *DescribeDBVersionInfosResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBVersionInfosResponse{}
	_body, _err := client.DescribeDBVersionInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of backup sets of full backup or point-in-time backup for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to query a list of backup sets and backup details only for instances in elastic storage mode.
//
// @param request - DescribeDataBackupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataBackupsResponse
func (client *Client) DescribeDataBackupsWithOptions(request *DescribeDataBackupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataBackupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupId) {
		query["BackupId"] = request.BackupId
	}

	if !dara.IsNil(request.BackupMode) {
		query["BackupMode"] = request.BackupMode
	}

	if !dara.IsNil(request.BackupStatus) {
		query["BackupStatus"] = request.BackupStatus
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DataType) {
		query["DataType"] = request.DataType
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDataBackups"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDataBackupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of backup sets of full backup or point-in-time backup for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to query a list of backup sets and backup details only for instances in elastic storage mode.
//
// @param request - DescribeDataBackupsRequest
//
// @return DescribeDataBackupsResponse
func (client *Client) DescribeDataBackups(request *DescribeDataBackupsRequest) (_result *DescribeDataBackupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDataBackupsResponse{}
	_body, _err := client.DescribeDataBackupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the data redistribution information about an AnalyticDB for PostgreSQL V6.0 instance in elastic storage mode.
//
// @param request - DescribeDataReDistributeInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataReDistributeInfoResponse
func (client *Client) DescribeDataReDistributeInfoWithOptions(request *DescribeDataReDistributeInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataReDistributeInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDataReDistributeInfo"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDataReDistributeInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the data redistribution information about an AnalyticDB for PostgreSQL V6.0 instance in elastic storage mode.
//
// @param request - DescribeDataReDistributeInfoRequest
//
// @return DescribeDataReDistributeInfoResponse
func (client *Client) DescribeDataReDistributeInfo(request *DescribeDataReDistributeInfoRequest) (_result *DescribeDataReDistributeInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDataReDistributeInfoResponse{}
	_body, _err := client.DescribeDataReDistributeInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the state of data sharing for AnalyticDB for PostgreSQL instances.
//
// Description:
//
// Data sharing is supported only for instances in Serverless mode.
//
// @param request - DescribeDataShareInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataShareInstancesResponse
func (client *Client) DescribeDataShareInstancesWithOptions(request *DescribeDataShareInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataShareInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SearchValue) {
		query["SearchValue"] = request.SearchValue
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDataShareInstances"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDataShareInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the state of data sharing for AnalyticDB for PostgreSQL instances.
//
// Description:
//
// Data sharing is supported only for instances in Serverless mode.
//
// @param request - DescribeDataShareInstancesRequest
//
// @return DescribeDataShareInstancesResponse
func (client *Client) DescribeDataShareInstances(request *DescribeDataShareInstancesRequest) (_result *DescribeDataShareInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDataShareInstancesResponse{}
	_body, _err := client.DescribeDataShareInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about data sharing performance metrics.
//
// Description:
//
// You can call this operation to query the details of data sharing performance metrics for an AnalyticDB for PostgreSQL instance in Serverless mode, such as the number of shared topics and the amount of data shared.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDataSharePerformanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataSharePerformanceResponse
func (client *Client) DescribeDataSharePerformanceWithOptions(request *DescribeDataSharePerformanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataSharePerformanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDataSharePerformance"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDataSharePerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about data sharing performance metrics.
//
// Description:
//
// You can call this operation to query the details of data sharing performance metrics for an AnalyticDB for PostgreSQL instance in Serverless mode, such as the number of shared topics and the amount of data shared.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDataSharePerformanceRequest
//
// @return DescribeDataSharePerformanceResponse
func (client *Client) DescribeDataSharePerformance(request *DescribeDataSharePerformanceRequest) (_result *DescribeDataSharePerformanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDataSharePerformanceResponse{}
	_body, _err := client.DescribeDataSharePerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all databases and database accounts for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// To facilitate management, you can call this operation to query all databases and database accounts on an AnalyticDB for PostgreSQL instance.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDiagnosisDimensionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDiagnosisDimensionsResponse
func (client *Client) DescribeDiagnosisDimensionsWithOptions(request *DescribeDiagnosisDimensionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDiagnosisDimensionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDiagnosisDimensions"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDiagnosisDimensionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all databases and database accounts for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// To facilitate management, you can call this operation to query all databases and database accounts on an AnalyticDB for PostgreSQL instance.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDiagnosisDimensionsRequest
//
// @return DescribeDiagnosisDimensionsResponse
func (client *Client) DescribeDiagnosisDimensions(request *DescribeDiagnosisDimensionsRequest) (_result *DescribeDiagnosisDimensionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDiagnosisDimensionsResponse{}
	_body, _err := client.DescribeDiagnosisDimensionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of query execution on an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to query the details of query execution on an AnalyticDB for PostgreSQL instance in elastic storage mode within a specified time range.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDiagnosisMonitorPerformanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDiagnosisMonitorPerformanceResponse
func (client *Client) DescribeDiagnosisMonitorPerformanceWithOptions(request *DescribeDiagnosisMonitorPerformanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDiagnosisMonitorPerformanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Database) {
		query["Database"] = request.Database
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.QueryCondition) {
		query["QueryCondition"] = request.QueryCondition
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.User) {
		query["User"] = request.User
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDiagnosisMonitorPerformance"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDiagnosisMonitorPerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of query execution on an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to query the details of query execution on an AnalyticDB for PostgreSQL instance in elastic storage mode within a specified time range.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDiagnosisMonitorPerformanceRequest
//
// @return DescribeDiagnosisMonitorPerformanceResponse
func (client *Client) DescribeDiagnosisMonitorPerformance(request *DescribeDiagnosisMonitorPerformanceRequest) (_result *DescribeDiagnosisMonitorPerformanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDiagnosisMonitorPerformanceResponse{}
	_body, _err := client.DescribeDiagnosisMonitorPerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about SQL queries for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to query the details of SQL queries on an AnalyticDB for PostgreSQL instance within a specified time range.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDiagnosisRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDiagnosisRecordsResponse
func (client *Client) DescribeDiagnosisRecordsWithOptions(request *DescribeDiagnosisRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDiagnosisRecordsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Database) {
		query["Database"] = request.Database
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
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

	if !dara.IsNil(request.QueryCondition) {
		query["QueryCondition"] = request.QueryCondition
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.User) {
		query["User"] = request.User
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDiagnosisRecords"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDiagnosisRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about SQL queries for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to query the details of SQL queries on an AnalyticDB for PostgreSQL instance within a specified time range.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDiagnosisRecordsRequest
//
// @return DescribeDiagnosisRecordsResponse
func (client *Client) DescribeDiagnosisRecords(request *DescribeDiagnosisRecordsRequest) (_result *DescribeDiagnosisRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDiagnosisRecordsResponse{}
	_body, _err := client.DescribeDiagnosisRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a query for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to query the information about a query for an AnalyticDB for PostgreSQL instance, including the SQL statement, execution plan text, and execution plan tree.
//
// This operation is available only for instances of V6.3.10.1 or later in elastic storage mode. For information about how to view and update the minor version of an instance, see [View the minor engine version](https://help.aliyun.com/document_detail/277424.html) and [Update the minor engine version](https://help.aliyun.com/document_detail/139271.html).
//
// @param request - DescribeDiagnosisSQLInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDiagnosisSQLInfoResponse
func (client *Client) DescribeDiagnosisSQLInfoWithOptions(request *DescribeDiagnosisSQLInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDiagnosisSQLInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Database) {
		query["Database"] = request.Database
	}

	if !dara.IsNil(request.QueryID) {
		query["QueryID"] = request.QueryID
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDiagnosisSQLInfo"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDiagnosisSQLInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a query for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to query the information about a query for an AnalyticDB for PostgreSQL instance, including the SQL statement, execution plan text, and execution plan tree.
//
// This operation is available only for instances of V6.3.10.1 or later in elastic storage mode. For information about how to view and update the minor version of an instance, see [View the minor engine version](https://help.aliyun.com/document_detail/277424.html) and [Update the minor engine version](https://help.aliyun.com/document_detail/139271.html).
//
// @param request - DescribeDiagnosisSQLInfoRequest
//
// @return DescribeDiagnosisSQLInfoResponse
func (client *Client) DescribeDiagnosisSQLInfo(request *DescribeDiagnosisSQLInfoRequest) (_result *DescribeDiagnosisSQLInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDiagnosisSQLInfoResponse{}
	_body, _err := client.DescribeDiagnosisSQLInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Document Details
//
// @param request - DescribeDocumentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDocumentResponse
func (client *Client) DescribeDocumentWithOptions(request *DescribeDocumentRequest, runtime *dara.RuntimeOptions) (_result *DescribeDocumentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Collection) {
		query["Collection"] = request.Collection
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespacePassword) {
		query["NamespacePassword"] = request.NamespacePassword
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDocument"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDocumentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Document Details
//
// @param request - DescribeDocumentRequest
//
// @return DescribeDocumentResponse
func (client *Client) DescribeDocument(request *DescribeDocumentRequest) (_result *DescribeDocumentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDocumentResponse{}
	_body, _err := client.DescribeDocumentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the download records of query diagnostic information for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You must call the [DownloadDiagnosisRecords](https://help.aliyun.com/document_detail/447700.html) operation to download the query diagnostic information before you can call this operation to query the download records and download URLs.
//
// This operation is available only for instances of V6.3.10.1 or later in elastic storage mode. For information about how to view and update the minor version of an instance, see [View the minor engine version](https://help.aliyun.com/document_detail/277424.html) and [Update the minor engine version](https://help.aliyun.com/document_detail/139271.html).
//
// @param request - DescribeDownloadRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDownloadRecordsResponse
func (client *Client) DescribeDownloadRecordsWithOptions(request *DescribeDownloadRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDownloadRecordsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DownloadTaskType) {
		query["DownloadTaskType"] = request.DownloadTaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDownloadRecords"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDownloadRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the download records of query diagnostic information for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You must call the [DownloadDiagnosisRecords](https://help.aliyun.com/document_detail/447700.html) operation to download the query diagnostic information before you can call this operation to query the download records and download URLs.
//
// This operation is available only for instances of V6.3.10.1 or later in elastic storage mode. For information about how to view and update the minor version of an instance, see [View the minor engine version](https://help.aliyun.com/document_detail/277424.html) and [Update the minor engine version](https://help.aliyun.com/document_detail/139271.html).
//
// @param request - DescribeDownloadRecordsRequest
//
// @return DescribeDownloadRecordsResponse
func (client *Client) DescribeDownloadRecords(request *DescribeDownloadRecordsRequest) (_result *DescribeDownloadRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDownloadRecordsResponse{}
	_body, _err := client.DescribeDownloadRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get download records
//
// @param request - DescribeDownloadSQLLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDownloadSQLLogsResponse
func (client *Client) DescribeDownloadSQLLogsWithOptions(request *DescribeDownloadSQLLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDownloadSQLLogsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDownloadSQLLogs"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDownloadSQLLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get download records
//
// @param request - DescribeDownloadSQLLogsRequest
//
// @return DescribeDownloadSQLLogsResponse
func (client *Client) DescribeDownloadSQLLogs(request *DescribeDownloadSQLLogsRequest) (_result *DescribeDownloadSQLLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDownloadSQLLogsResponse{}
	_body, _err := client.DescribeDownloadSQLLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an external data service.
//
// @param request - DescribeExternalDataServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExternalDataServiceResponse
func (client *Client) DescribeExternalDataServiceWithOptions(request *DescribeExternalDataServiceRequest, runtime *dara.RuntimeOptions) (_result *DescribeExternalDataServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeExternalDataService"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExternalDataServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an external data service.
//
// @param request - DescribeExternalDataServiceRequest
//
// @return DescribeExternalDataServiceResponse
func (client *Client) DescribeExternalDataService(request *DescribeExternalDataServiceRequest) (_result *DescribeExternalDataServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeExternalDataServiceResponse{}
	_body, _err := client.DescribeExternalDataServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries E-MapReduce (EMR) Hadoop clusters in a specific virtual private cloud (VPC).
//
// @param request - DescribeHadoopClustersInSameNetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHadoopClustersInSameNetResponse
func (client *Client) DescribeHadoopClustersInSameNetWithOptions(request *DescribeHadoopClustersInSameNetRequest, runtime *dara.RuntimeOptions) (_result *DescribeHadoopClustersInSameNetResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("DescribeHadoopClustersInSameNet"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHadoopClustersInSameNetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries E-MapReduce (EMR) Hadoop clusters in a specific virtual private cloud (VPC).
//
// @param request - DescribeHadoopClustersInSameNetRequest
//
// @return DescribeHadoopClustersInSameNetResponse
func (client *Client) DescribeHadoopClustersInSameNet(request *DescribeHadoopClustersInSameNetRequest) (_result *DescribeHadoopClustersInSameNetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHadoopClustersInSameNetResponse{}
	_body, _err := client.DescribeHadoopClustersInSameNetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configuration information about a Hadoop cluster.
//
// @param request - DescribeHadoopConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHadoopConfigsResponse
func (client *Client) DescribeHadoopConfigsWithOptions(request *DescribeHadoopConfigsRequest, runtime *dara.RuntimeOptions) (_result *DescribeHadoopConfigsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigName) {
		query["ConfigName"] = request.ConfigName
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.EmrInstanceId) {
		query["EmrInstanceId"] = request.EmrInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHadoopConfigs"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHadoopConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration information about a Hadoop cluster.
//
// @param request - DescribeHadoopConfigsRequest
//
// @return DescribeHadoopConfigsResponse
func (client *Client) DescribeHadoopConfigs(request *DescribeHadoopConfigsRequest) (_result *DescribeHadoopConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHadoopConfigsResponse{}
	_body, _err := client.DescribeHadoopConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the configurations of a Hadoop data source.
//
// @param request - DescribeHadoopDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHadoopDataSourceResponse
func (client *Client) DescribeHadoopDataSourceWithOptions(request *DescribeHadoopDataSourceRequest, runtime *dara.RuntimeOptions) (_result *DescribeHadoopDataSourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DataSourceId) {
		query["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHadoopDataSource"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHadoopDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the configurations of a Hadoop data source.
//
// @param request - DescribeHadoopDataSourceRequest
//
// @return DescribeHadoopDataSourceResponse
func (client *Client) DescribeHadoopDataSource(request *DescribeHadoopDataSourceRequest) (_result *DescribeHadoopDataSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHadoopDataSourceResponse{}
	_body, _err := client.DescribeHadoopDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the health status of an AnalyticDB for PostgreSQL instance and its nodes.
//
// Description:
//
// This operation is called to query the health status of an AnalyticDB for PostgreSQL instance in elastic storage mode or Serverless mode and its coordinator and compute nodes.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeHealthStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHealthStatusResponse
func (client *Client) DescribeHealthStatusWithOptions(request *DescribeHealthStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeHealthStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHealthStatus"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHealthStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the health status of an AnalyticDB for PostgreSQL instance and its nodes.
//
// Description:
//
// This operation is called to query the health status of an AnalyticDB for PostgreSQL instance in elastic storage mode or Serverless mode and its coordinator and compute nodes.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeHealthStatusRequest
//
// @return DescribeHealthStatusResponse
func (client *Client) DescribeHealthStatus(request *DescribeHealthStatusRequest) (_result *DescribeHealthStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHealthStatusResponse{}
	_body, _err := client.DescribeHealthStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about real-time materialized views of an AnalyticDB for PostgreSQL instance.
//
// @param request - DescribeIMVInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIMVInfosResponse
func (client *Client) DescribeIMVInfosWithOptions(request *DescribeIMVInfosRequest, runtime *dara.RuntimeOptions) (_result *DescribeIMVInfosResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Database) {
		query["Database"] = request.Database
	}

	if !dara.IsNil(request.MVName) {
		query["MVName"] = request.MVName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeIMVInfos"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIMVInfosResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about real-time materialized views of an AnalyticDB for PostgreSQL instance.
//
// @param request - DescribeIMVInfosRequest
//
// @return DescribeIMVInfosResponse
func (client *Client) DescribeIMVInfos(request *DescribeIMVInfosRequest) (_result *DescribeIMVInfosResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeIMVInfosResponse{}
	_body, _err := client.DescribeIMVInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取索引详情
//
// @param request - DescribeIndexRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIndexResponse
func (client *Client) DescribeIndexWithOptions(request *DescribeIndexRequest, runtime *dara.RuntimeOptions) (_result *DescribeIndexResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Collection) {
		query["Collection"] = request.Collection
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.IndexName) {
		query["IndexName"] = request.IndexName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespacePassword) {
		query["NamespacePassword"] = request.NamespacePassword
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeIndex"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取索引详情
//
// @param request - DescribeIndexRequest
//
// @return DescribeIndexResponse
func (client *Client) DescribeIndex(request *DescribeIndexRequest) (_result *DescribeIndexResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeIndexResponse{}
	_body, _err := client.DescribeIndexWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of a Java Database Connectivity (JDBC) data source.
//
// @param request - DescribeJDBCDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeJDBCDataSourceResponse
func (client *Client) DescribeJDBCDataSourceWithOptions(request *DescribeJDBCDataSourceRequest, runtime *dara.RuntimeOptions) (_result *DescribeJDBCDataSourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DataSourceId) {
		query["DataSourceId"] = request.DataSourceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeJDBCDataSource"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeJDBCDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of a Java Database Connectivity (JDBC) data source.
//
// @param request - DescribeJDBCDataSourceRequest
//
// @return DescribeJDBCDataSourceResponse
func (client *Client) DescribeJDBCDataSource(request *DescribeJDBCDataSourceRequest) (_result *DescribeJDBCDataSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeJDBCDataSourceResponse{}
	_body, _err := client.DescribeJDBCDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of log backups.
//
// @param request - DescribeLogBackupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLogBackupsResponse
func (client *Client) DescribeLogBackupsWithOptions(request *DescribeLogBackupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeLogBackupsResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLogBackups"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLogBackupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of log backups.
//
// @param request - DescribeLogBackupsRequest
//
// @return DescribeLogBackupsResponse
func (client *Client) DescribeLogBackups(request *DescribeLogBackupsRequest) (_result *DescribeLogBackupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLogBackupsResponse{}
	_body, _err := client.DescribeLogBackupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the parameter modification logs of an AnalyticDB for PostgreSQL instance.
//
// @param request - DescribeModifyParameterLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeModifyParameterLogResponse
func (client *Client) DescribeModifyParameterLogWithOptions(request *DescribeModifyParameterLogRequest, runtime *dara.RuntimeOptions) (_result *DescribeModifyParameterLogResponse, _err error) {
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeModifyParameterLog"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeModifyParameterLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the parameter modification logs of an AnalyticDB for PostgreSQL instance.
//
// @param request - DescribeModifyParameterLogRequest
//
// @return DescribeModifyParameterLogResponse
func (client *Client) DescribeModifyParameterLog(request *DescribeModifyParameterLogRequest) (_result *DescribeModifyParameterLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeModifyParameterLogResponse{}
	_body, _err := client.DescribeModifyParameterLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a namespace.
//
// @param request - DescribeNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNamespaceResponse
func (client *Client) DescribeNamespaceWithOptions(request *DescribeNamespaceRequest, runtime *dara.RuntimeOptions) (_result *DescribeNamespaceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.ManagerAccount) {
		query["ManagerAccount"] = request.ManagerAccount
	}

	if !dara.IsNil(request.ManagerAccountPassword) {
		query["ManagerAccountPassword"] = request.ManagerAccountPassword
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNamespace"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a namespace.
//
// @param request - DescribeNamespaceRequest
//
// @return DescribeNamespaceResponse
func (client *Client) DescribeNamespace(request *DescribeNamespaceRequest) (_result *DescribeNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNamespaceResponse{}
	_body, _err := client.DescribeNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about configuration parameters for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// This operation can be called to query the details of parameters in an AnalyticDB for PostgreSQL instance in elastic storage mode or Serverless mode.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered and may affect your business. We recommend that you take note of the limit when you call this operation.
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeParameters"),
		Version:     dara.String("2016-05-03"),
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
// Queries the information about configuration parameters for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// This operation can be called to query the details of parameters in an AnalyticDB for PostgreSQL instance in elastic storage mode or Serverless mode.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered and may affect your business. We recommend that you take note of the limit when you call this operation.
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
// Queries a list of vSwitches.
//
// Description:
//
// When you create AnalyticDB for PostgreSQL instances, you can call this operation to query the details of vSwitches within a specified region or zone.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeRdsVSwitchsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRdsVSwitchsResponse
func (client *Client) DescribeRdsVSwitchsWithOptions(request *DescribeRdsVSwitchsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRdsVSwitchsResponse, _err error) {
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
		Action:      dara.String("DescribeRdsVSwitchs"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRdsVSwitchsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of vSwitches.
//
// Description:
//
// When you create AnalyticDB for PostgreSQL instances, you can call this operation to query the details of vSwitches within a specified region or zone.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeRdsVSwitchsRequest
//
// @return DescribeRdsVSwitchsResponse
func (client *Client) DescribeRdsVSwitchs(request *DescribeRdsVSwitchsRequest) (_result *DescribeRdsVSwitchsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRdsVSwitchsResponse{}
	_body, _err := client.DescribeRdsVSwitchsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of VPCs.
//
// Description:
//
// When you create an AnalyticDB for PostgreSQL instance, you can call this operation to query the available VPCs within a specified region or zone.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeRdsVpcsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRdsVpcsResponse
func (client *Client) DescribeRdsVpcsWithOptions(request *DescribeRdsVpcsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRdsVpcsResponse, _err error) {
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
		Action:      dara.String("DescribeRdsVpcs"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRdsVpcsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of VPCs.
//
// Description:
//
// When you create an AnalyticDB for PostgreSQL instance, you can call this operation to query the available VPCs within a specified region or zone.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeRdsVpcsRequest
//
// @return DescribeRdsVpcsResponse
func (client *Client) DescribeRdsVpcs(request *DescribeRdsVpcsRequest) (_result *DescribeRdsVpcsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRdsVpcsResponse{}
	_body, _err := client.DescribeRdsVpcsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of regions and zones where AnalyticDB for PostgreSQL is available.
//
// Description:
//
// Before you create an AnalyticDB for PostgreSQL instance, you must call this operation to query available regions and zones.
//
// ## Limit
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered and may affect your business. We recommend that you take note of the limit when you call this operation.
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
	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2016-05-03"),
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
// Queries a list of regions and zones where AnalyticDB for PostgreSQL is available.
//
// Description:
//
// Before you create an AnalyticDB for PostgreSQL instance, you must call this operation to query available regions and zones.
//
// ## Limit
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered and may affect your business. We recommend that you take note of the limit when you call this operation.
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
// Queries a list of roles.
//
// @param request - DescribeRolesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRolesResponse
func (client *Client) DescribeRolesWithOptions(request *DescribeRolesRequest, runtime *dara.RuntimeOptions) (_result *DescribeRolesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRoles"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRolesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of roles.
//
// @param request - DescribeRolesRequest
//
// @return DescribeRolesResponse
func (client *Client) DescribeRoles(request *DescribeRolesRequest) (_result *DescribeRolesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRolesResponse{}
	_body, _err := client.DescribeRolesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of audit logs for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// This operation is not available for instances in reserved storage mode.
//
// @param request - DescribeSQLLogCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSQLLogCountResponse
func (client *Client) DescribeSQLLogCountWithOptions(request *DescribeSQLLogCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeSQLLogCountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Database) {
		query["Database"] = request.Database
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ExecuteCost) {
		query["ExecuteCost"] = request.ExecuteCost
	}

	if !dara.IsNil(request.ExecuteState) {
		query["ExecuteState"] = request.ExecuteState
	}

	if !dara.IsNil(request.MaxExecuteCost) {
		query["MaxExecuteCost"] = request.MaxExecuteCost
	}

	if !dara.IsNil(request.MinExecuteCost) {
		query["MinExecuteCost"] = request.MinExecuteCost
	}

	if !dara.IsNil(request.OperationClass) {
		query["OperationClass"] = request.OperationClass
	}

	if !dara.IsNil(request.OperationType) {
		query["OperationType"] = request.OperationType
	}

	if !dara.IsNil(request.QueryKeywords) {
		query["QueryKeywords"] = request.QueryKeywords
	}

	if !dara.IsNil(request.SourceIP) {
		query["SourceIP"] = request.SourceIP
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.User) {
		query["User"] = request.User
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSQLLogCount"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSQLLogCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of audit logs for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// This operation is not available for instances in reserved storage mode.
//
// @param request - DescribeSQLLogCountRequest
//
// @return DescribeSQLLogCountResponse
func (client *Client) DescribeSQLLogCount(request *DescribeSQLLogCountRequest) (_result *DescribeSQLLogCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSQLLogCountResponse{}
	_body, _err := client.DescribeSQLLogCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the SQL execution logs of an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// > This operation is no longer used. To query SQL execution logs, call the [DescribeSQLLogsV2](https://help.aliyun.com/document_detail/453722.html) operation.
//
// @param request - DescribeSQLLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSQLLogsResponse
func (client *Client) DescribeSQLLogsWithOptions(request *DescribeSQLLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSQLLogsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Database) {
		query["Database"] = request.Database
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ExecuteCost) {
		query["ExecuteCost"] = request.ExecuteCost
	}

	if !dara.IsNil(request.ExecuteState) {
		query["ExecuteState"] = request.ExecuteState
	}

	if !dara.IsNil(request.MaxExecuteCost) {
		query["MaxExecuteCost"] = request.MaxExecuteCost
	}

	if !dara.IsNil(request.MinExecuteCost) {
		query["MinExecuteCost"] = request.MinExecuteCost
	}

	if !dara.IsNil(request.OperationClass) {
		query["OperationClass"] = request.OperationClass
	}

	if !dara.IsNil(request.OperationType) {
		query["OperationType"] = request.OperationType
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

	if !dara.IsNil(request.SourceIP) {
		query["SourceIP"] = request.SourceIP
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.User) {
		query["User"] = request.User
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSQLLogs"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSQLLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the SQL execution logs of an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// > This operation is no longer used. To query SQL execution logs, call the [DescribeSQLLogsV2](https://help.aliyun.com/document_detail/453722.html) operation.
//
// @param request - DescribeSQLLogsRequest
//
// @return DescribeSQLLogsResponse
func (client *Client) DescribeSQLLogs(request *DescribeSQLLogsRequest) (_result *DescribeSQLLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSQLLogsResponse{}
	_body, _err := client.DescribeSQLLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries SQL logs within a specific time range.
//
// Description:
//
// You can call this operation to query SQL logs of an AnalyticDB for PostgreSQL instance within a specific time range.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeSQLLogsV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSQLLogsV2Response
func (client *Client) DescribeSQLLogsV2WithOptions(request *DescribeSQLLogsV2Request, runtime *dara.RuntimeOptions) (_result *DescribeSQLLogsV2Response, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Database) {
		query["Database"] = request.Database
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ExecuteCost) {
		query["ExecuteCost"] = request.ExecuteCost
	}

	if !dara.IsNil(request.ExecuteState) {
		query["ExecuteState"] = request.ExecuteState
	}

	if !dara.IsNil(request.MaxExecuteCost) {
		query["MaxExecuteCost"] = request.MaxExecuteCost
	}

	if !dara.IsNil(request.MinExecuteCost) {
		query["MinExecuteCost"] = request.MinExecuteCost
	}

	if !dara.IsNil(request.OperationClass) {
		query["OperationClass"] = request.OperationClass
	}

	if !dara.IsNil(request.OperationType) {
		query["OperationType"] = request.OperationType
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SourceIP) {
		query["SourceIP"] = request.SourceIP
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.User) {
		query["User"] = request.User
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSQLLogsV2"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSQLLogsV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries SQL logs within a specific time range.
//
// Description:
//
// You can call this operation to query SQL logs of an AnalyticDB for PostgreSQL instance within a specific time range.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeSQLLogsV2Request
//
// @return DescribeSQLLogsV2Response
func (client *Client) DescribeSQLLogsV2(request *DescribeSQLLogsV2Request) (_result *DescribeSQLLogsV2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSQLLogsV2Response{}
	_body, _err := client.DescribeSQLLogsV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether a sample dataset is loaded to an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeSampleDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSampleDataResponse
func (client *Client) DescribeSampleDataWithOptions(request *DescribeSampleDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeSampleDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSampleData"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSampleDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether a sample dataset is loaded to an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeSampleDataRequest
//
// @return DescribeSampleDataResponse
func (client *Client) DescribeSampleData(request *DescribeSampleDataRequest) (_result *DescribeSampleDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSampleDataResponse{}
	_body, _err := client.DescribeSampleDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a real-time data service.
//
// @param request - DescribeStreamingDataServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeStreamingDataServiceResponse
func (client *Client) DescribeStreamingDataServiceWithOptions(request *DescribeStreamingDataServiceRequest, runtime *dara.RuntimeOptions) (_result *DescribeStreamingDataServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeStreamingDataService"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeStreamingDataServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a real-time data service.
//
// @param request - DescribeStreamingDataServiceRequest
//
// @return DescribeStreamingDataServiceResponse
func (client *Client) DescribeStreamingDataService(request *DescribeStreamingDataServiceRequest) (_result *DescribeStreamingDataServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeStreamingDataServiceResponse{}
	_body, _err := client.DescribeStreamingDataServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get external data source configuration information
//
// @param request - DescribeStreamingDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeStreamingDataSourceResponse
func (client *Client) DescribeStreamingDataSourceWithOptions(request *DescribeStreamingDataSourceRequest, runtime *dara.RuntimeOptions) (_result *DescribeStreamingDataSourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DataSourceId) {
		query["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeStreamingDataSource"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeStreamingDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get external data source configuration information
//
// @param request - DescribeStreamingDataSourceRequest
//
// @return DescribeStreamingDataSourceResponse
func (client *Client) DescribeStreamingDataSource(request *DescribeStreamingDataSourceRequest) (_result *DescribeStreamingDataSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeStreamingDataSourceResponse{}
	_body, _err := client.DescribeStreamingDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete External Data Source Configuration
//
// @param request - DescribeStreamingJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeStreamingJobResponse
func (client *Client) DescribeStreamingJobWithOptions(request *DescribeStreamingJobRequest, runtime *dara.RuntimeOptions) (_result *DescribeStreamingJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeStreamingJob"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeStreamingJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete External Data Source Configuration
//
// @param request - DescribeStreamingJobRequest
//
// @return DescribeStreamingJobResponse
func (client *Client) DescribeStreamingJob(request *DescribeStreamingJobRequest) (_result *DescribeStreamingJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeStreamingJobResponse{}
	_body, _err := client.DescribeStreamingJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the features that are supported by an AnalyticDB for PostgreSQL instance.
//
// @param request - DescribeSupportFeaturesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSupportFeaturesResponse
func (client *Client) DescribeSupportFeaturesWithOptions(request *DescribeSupportFeaturesRequest, runtime *dara.RuntimeOptions) (_result *DescribeSupportFeaturesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSupportFeatures"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSupportFeaturesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the features that are supported by an AnalyticDB for PostgreSQL instance.
//
// @param request - DescribeSupportFeaturesRequest
//
// @return DescribeSupportFeaturesResponse
func (client *Client) DescribeSupportFeatures(request *DescribeSupportFeaturesRequest) (_result *DescribeSupportFeaturesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSupportFeaturesResponse{}
	_body, _err := client.DescribeSupportFeaturesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a table.
//
// @param request - DescribeTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTableResponse
func (client *Client) DescribeTableWithOptions(request *DescribeTableRequest, runtime *dara.RuntimeOptions) (_result *DescribeTableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Database) {
		query["Database"] = request.Database
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Schema) {
		query["Schema"] = request.Schema
	}

	if !dara.IsNil(request.SecretArn) {
		query["SecretArn"] = request.SecretArn
	}

	if !dara.IsNil(request.Table) {
		query["Table"] = request.Table
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTable"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a table.
//
// @param request - DescribeTableRequest
//
// @return DescribeTableResponse
func (client *Client) DescribeTable(request *DescribeTableRequest) (_result *DescribeTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTableResponse{}
	_body, _err := client.DescribeTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of tags for AnalyticDB for PostgreSQL instances.
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

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTags"),
		Version:     dara.String("2016-05-03"),
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
// Queries a list of tags for AnalyticDB for PostgreSQL instances.
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

// Summary:
//
// Queries a list of Key Management Service (KMS) keys.
//
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
		Action:      dara.String("DescribeUserEncryptionKeyList"),
		Version:     dara.String("2016-05-03"),
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

// Summary:
//
// Queries a list of Key Management Service (KMS) keys.
//
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

// Summary:
//
// Queries the information about a lock-waiting query for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to query the details of a lock-waiting query only for an AnalyticDB for PostgreSQL V6.0 instance in elastic storage mode.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeWaitingSQLInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWaitingSQLInfoResponse
func (client *Client) DescribeWaitingSQLInfoWithOptions(request *DescribeWaitingSQLInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeWaitingSQLInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Database) {
		query["Database"] = request.Database
	}

	if !dara.IsNil(request.PID) {
		query["PID"] = request.PID
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWaitingSQLInfo"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWaitingSQLInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a lock-waiting query for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to query the details of a lock-waiting query only for an AnalyticDB for PostgreSQL V6.0 instance in elastic storage mode.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeWaitingSQLInfoRequest
//
// @return DescribeWaitingSQLInfoResponse
func (client *Client) DescribeWaitingSQLInfo(request *DescribeWaitingSQLInfoRequest) (_result *DescribeWaitingSQLInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeWaitingSQLInfoResponse{}
	_body, _err := client.DescribeWaitingSQLInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the lock diagnostic records of an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to query the lock diagnostics records only for an AnalyticDB for PostgreSQL V6.0 instance in elastic storage mode.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeWaitingSQLRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWaitingSQLRecordsResponse
func (client *Client) DescribeWaitingSQLRecordsWithOptions(request *DescribeWaitingSQLRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeWaitingSQLRecordsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Database) {
		query["Database"] = request.Database
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
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

	if !dara.IsNil(request.QueryCondition) {
		query["QueryCondition"] = request.QueryCondition
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.User) {
		query["User"] = request.User
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWaitingSQLRecords"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWaitingSQLRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the lock diagnostic records of an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to query the lock diagnostics records only for an AnalyticDB for PostgreSQL V6.0 instance in elastic storage mode.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeWaitingSQLRecordsRequest
//
// @return DescribeWaitingSQLRecordsResponse
func (client *Client) DescribeWaitingSQLRecords(request *DescribeWaitingSQLRecordsRequest) (_result *DescribeWaitingSQLRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeWaitingSQLRecordsResponse{}
	_body, _err := client.DescribeWaitingSQLRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables resource group management for an AnalyticDB for PostgreSQL V6.0 instance in elastic storage mode. After you disable resource group management, the resource management method of the instance switches from resource group management to resource queue management.
//
// Description:
//
//	  You can call this operation only for AnalyticDB for PostgreSQL V6.0 instances in elastic storage mode whose minor version is V6.6.1.0 or later.
//
//		- You can call this operation to disable resource group management only for AnalyticDB for PostgreSQL instances that are in the **Running*	- state.
//
//		- **Note: When the resource management method is switched, your AnalyticDB for PostgreSQL instance restarts and becomes unavailable for approximately 5 minutes. To prevent your business from being affected, call this operation during off-peak hours.
//
// @param request - DisableDBResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableDBResourceGroupResponse
func (client *Client) DisableDBResourceGroupWithOptions(request *DisableDBResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *DisableDBResourceGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableDBResourceGroup"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableDBResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables resource group management for an AnalyticDB for PostgreSQL V6.0 instance in elastic storage mode. After you disable resource group management, the resource management method of the instance switches from resource group management to resource queue management.
//
// Description:
//
//	  You can call this operation only for AnalyticDB for PostgreSQL V6.0 instances in elastic storage mode whose minor version is V6.6.1.0 or later.
//
//		- You can call this operation to disable resource group management only for AnalyticDB for PostgreSQL instances that are in the **Running*	- state.
//
//		- **Note: When the resource management method is switched, your AnalyticDB for PostgreSQL instance restarts and becomes unavailable for approximately 5 minutes. To prevent your business from being affected, call this operation during off-peak hours.
//
// @param request - DisableDBResourceGroupRequest
//
// @return DisableDBResourceGroupResponse
func (client *Client) DisableDBResourceGroup(request *DisableDBResourceGroupRequest) (_result *DisableDBResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableDBResourceGroupResponse{}
	_body, _err := client.DisableDBResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Downloads the query diagnostic information of an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to download the query diagnostic information of an AnalyticDB for PostgreSQL instance. After the download is complete, you can call the [DescribeDownloadRecords](https://help.aliyun.com/document_detail/447712.html) operation to query download records and download URLs.
//
// This operation is available only for instances of V6.3.10.1 or later in elastic storage mode. For more information about how to view and update the minor engine version of an instance, see [View the minor engine version](https://help.aliyun.com/document_detail/277424.html) and [Update the minor engine version](https://help.aliyun.com/document_detail/139271.html).
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DownloadDiagnosisRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DownloadDiagnosisRecordsResponse
func (client *Client) DownloadDiagnosisRecordsWithOptions(request *DownloadDiagnosisRecordsRequest, runtime *dara.RuntimeOptions) (_result *DownloadDiagnosisRecordsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Database) {
		query["Database"] = request.Database
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.QueryCondition) {
		query["QueryCondition"] = request.QueryCondition
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.User) {
		query["User"] = request.User
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DownloadDiagnosisRecords"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DownloadDiagnosisRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Downloads the query diagnostic information of an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to download the query diagnostic information of an AnalyticDB for PostgreSQL instance. After the download is complete, you can call the [DescribeDownloadRecords](https://help.aliyun.com/document_detail/447712.html) operation to query download records and download URLs.
//
// This operation is available only for instances of V6.3.10.1 or later in elastic storage mode. For more information about how to view and update the minor engine version of an instance, see [View the minor engine version](https://help.aliyun.com/document_detail/277424.html) and [Update the minor engine version](https://help.aliyun.com/document_detail/139271.html).
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DownloadDiagnosisRecordsRequest
//
// @return DownloadDiagnosisRecordsResponse
func (client *Client) DownloadDiagnosisRecords(request *DownloadDiagnosisRecordsRequest) (_result *DownloadDiagnosisRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DownloadDiagnosisRecordsResponse{}
	_body, _err := client.DownloadDiagnosisRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Download the slow query logs of an AnalyticDB for PostgreSQL instance.
//
// @param request - DownloadSQLLogsRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DownloadSQLLogsRecordsResponse
func (client *Client) DownloadSQLLogsRecordsWithOptions(request *DownloadSQLLogsRecordsRequest, runtime *dara.RuntimeOptions) (_result *DownloadSQLLogsRecordsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Database) {
		query["Database"] = request.Database
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ExecuteCost) {
		query["ExecuteCost"] = request.ExecuteCost
	}

	if !dara.IsNil(request.ExecuteState) {
		query["ExecuteState"] = request.ExecuteState
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MaxExecuteCost) {
		query["MaxExecuteCost"] = request.MaxExecuteCost
	}

	if !dara.IsNil(request.MinExecuteCost) {
		query["MinExecuteCost"] = request.MinExecuteCost
	}

	if !dara.IsNil(request.OperationClass) {
		query["OperationClass"] = request.OperationClass
	}

	if !dara.IsNil(request.OperationType) {
		query["OperationType"] = request.OperationType
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

	if !dara.IsNil(request.SourceIP) {
		query["SourceIP"] = request.SourceIP
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.User) {
		query["User"] = request.User
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DownloadSQLLogsRecords"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DownloadSQLLogsRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Download the slow query logs of an AnalyticDB for PostgreSQL instance.
//
// @param request - DownloadSQLLogsRecordsRequest
//
// @return DownloadSQLLogsRecordsResponse
func (client *Client) DownloadSQLLogsRecords(request *DownloadSQLLogsRecordsRequest) (_result *DownloadSQLLogsRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DownloadSQLLogsRecordsResponse{}
	_body, _err := client.DownloadSQLLogsRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables resource group management for an AnalyticDB for PostgreSQL V6.0 instance in elastic storage mode. After resource group management is enabled, the resource management mode of the instance is changed from resource queue to resource group.
//
// Description:
//
//	  You can call this operation only for AnalyticDB for PostgreSQL V6.0 instances in elastic storage mode whose minor version is V6.6.1.0 or later.
//
//		- You can call this operation to enable resource group management only for AnalyticDB for PostgreSQL instances that are in the **Running*	- state.
//
//		- **Note: When the resource management mode is changed, your AnalyticDB for PostgreSQL instance is restarted and remains unavailable within 5 minutes. To prevent your business from being affected, call this operation during off-peak hours.
//
// @param request - EnableDBResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableDBResourceGroupResponse
func (client *Client) EnableDBResourceGroupWithOptions(request *EnableDBResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *EnableDBResourceGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableDBResourceGroup"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableDBResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables resource group management for an AnalyticDB for PostgreSQL V6.0 instance in elastic storage mode. After resource group management is enabled, the resource management mode of the instance is changed from resource queue to resource group.
//
// Description:
//
//	  You can call this operation only for AnalyticDB for PostgreSQL V6.0 instances in elastic storage mode whose minor version is V6.6.1.0 or later.
//
//		- You can call this operation to enable resource group management only for AnalyticDB for PostgreSQL instances that are in the **Running*	- state.
//
//		- **Note: When the resource management mode is changed, your AnalyticDB for PostgreSQL instance is restarted and remains unavailable within 5 minutes. To prevent your business from being affected, call this operation during off-peak hours.
//
// @param request - EnableDBResourceGroupRequest
//
// @return EnableDBResourceGroupResponse
func (client *Client) EnableDBResourceGroup(request *EnableDBResourceGroupRequest) (_result *EnableDBResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableDBResourceGroupResponse{}
	_body, _err := client.EnableDBResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Executes SQL statements.
//
// @param tmpReq - ExecuteStatementRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteStatementResponse
func (client *Client) ExecuteStatementWithOptions(tmpReq *ExecuteStatementRequest, runtime *dara.RuntimeOptions) (_result *ExecuteStatementResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ExecuteStatementShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Parameters) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, dara.String("Parameters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RagWorkspaceCollection) {
		request.RagWorkspaceCollectionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RagWorkspaceCollection, dara.String("RagWorkspaceCollection"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Sqls) {
		request.SqlsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sqls, dara.String("Sqls"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Database) {
		query["Database"] = request.Database
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RagWorkspaceCollectionShrink) {
		query["RagWorkspaceCollection"] = request.RagWorkspaceCollectionShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RunType) {
		query["RunType"] = request.RunType
	}

	if !dara.IsNil(request.SecretArn) {
		query["SecretArn"] = request.SecretArn
	}

	if !dara.IsNil(request.StatementName) {
		query["StatementName"] = request.StatementName
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ParametersShrink) {
		body["Parameters"] = request.ParametersShrink
	}

	if !dara.IsNil(request.Sql) {
		body["Sql"] = request.Sql
	}

	if !dara.IsNil(request.SqlsShrink) {
		body["Sqls"] = request.SqlsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteStatement"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Executes SQL statements.
//
// @param request - ExecuteStatementRequest
//
// @return ExecuteStatementResponse
func (client *Client) ExecuteStatement(request *ExecuteStatementRequest) (_result *ExecuteStatementResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExecuteStatementResponse{}
	_body, _err := client.ExecuteStatementWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取特定的账号信息
//
// @param request - GetAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccountResponse
func (client *Client) GetAccountWithOptions(request *GetAccountRequest, runtime *dara.RuntimeOptions) (_result *GetAccountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAccount"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取特定的账号信息
//
// @param request - GetAccountRequest
//
// @return GetAccountResponse
func (client *Client) GetAccount(request *GetAccountRequest) (_result *GetAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAccountResponse{}
	_body, _err := client.GetAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an access credential.
//
// @param request - GetSecretValueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecretValueResponse
func (client *Client) GetSecretValueWithOptions(request *GetSecretValueRequest, runtime *dara.RuntimeOptions) (_result *GetSecretValueResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecretArn) {
		query["SecretArn"] = request.SecretArn
	}

	if !dara.IsNil(request.SecretName) {
		query["SecretName"] = request.SecretName
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSecretValue"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSecretValueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an access credential.
//
// @param request - GetSecretValueRequest
//
// @return GetSecretValueResponse
func (client *Client) GetSecretValue(request *GetSecretValueRequest) (_result *GetSecretValueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSecretValueResponse{}
	_body, _err := client.GetSecretValueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Asynchronous SQL Execution Result
//
// @param request - GetStatementResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStatementResultResponse
func (client *Client) GetStatementResultWithOptions(request *GetStatementResultRequest, runtime *dara.RuntimeOptions) (_result *GetStatementResultResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Database) {
		query["Database"] = request.Database
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecretArn) {
		query["SecretArn"] = request.SecretArn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStatementResult"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStatementResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Asynchronous SQL Execution Result
//
// @param request - GetStatementResultRequest
//
// @return GetStatementResultResponse
func (client *Client) GetStatementResult(request *GetStatementResultRequest) (_result *GetStatementResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetStatementResultResponse{}
	_body, _err := client.GetStatementResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询Supabase实例详情
//
// @param request - GetSupabaseProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSupabaseProjectResponse
func (client *Client) GetSupabaseProjectWithOptions(request *GetSupabaseProjectRequest, runtime *dara.RuntimeOptions) (_result *GetSupabaseProjectResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      dara.String("GetSupabaseProject"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSupabaseProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Supabase实例详情
//
// @param request - GetSupabaseProjectRequest
//
// @return GetSupabaseProjectResponse
func (client *Client) GetSupabaseProject(request *GetSupabaseProjectRequest) (_result *GetSupabaseProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSupabaseProjectResponse{}
	_body, _err := client.GetSupabaseProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询Supabase实例 API Keys
//
// @param request - GetSupabaseProjectApiKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSupabaseProjectApiKeysResponse
func (client *Client) GetSupabaseProjectApiKeysWithOptions(request *GetSupabaseProjectApiKeysRequest, runtime *dara.RuntimeOptions) (_result *GetSupabaseProjectApiKeysResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      dara.String("GetSupabaseProjectApiKeys"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSupabaseProjectApiKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Supabase实例 API Keys
//
// @param request - GetSupabaseProjectApiKeysRequest
//
// @return GetSupabaseProjectApiKeysResponse
func (client *Client) GetSupabaseProjectApiKeys(request *GetSupabaseProjectApiKeysRequest) (_result *GetSupabaseProjectApiKeysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSupabaseProjectApiKeysResponse{}
	_body, _err := client.GetSupabaseProjectApiKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询Supabase项目dashboard账号信息
//
// @param request - GetSupabaseProjectDashboardAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSupabaseProjectDashboardAccountResponse
func (client *Client) GetSupabaseProjectDashboardAccountWithOptions(request *GetSupabaseProjectDashboardAccountRequest, runtime *dara.RuntimeOptions) (_result *GetSupabaseProjectDashboardAccountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      dara.String("GetSupabaseProjectDashboardAccount"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSupabaseProjectDashboardAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Supabase项目dashboard账号信息
//
// @param request - GetSupabaseProjectDashboardAccountRequest
//
// @return GetSupabaseProjectDashboardAccountResponse
func (client *Client) GetSupabaseProjectDashboardAccount(request *GetSupabaseProjectDashboardAccountRequest) (_result *GetSupabaseProjectDashboardAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSupabaseProjectDashboardAccountResponse{}
	_body, _err := client.GetSupabaseProjectDashboardAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the progress and result of an asynchronous document upload job based on the job ID.
//
// Description:
//
// This operation is related to the UploadDocumentAsync operation. You can call the UploadDocumentAsync operation to create an upload job and obtain the job ID, and then call the GetUploadDocumentJob operation to query the execution information of the job.
//
// >  Suggestions:
//
//   - Determine whether the document upload job times out based on the document complexity and the number of tokens after chunking. In most cases, a job that lasts more than 2 hours is considered timeout.
//
// @param request - GetUploadDocumentJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUploadDocumentJobResponse
func (client *Client) GetUploadDocumentJobWithOptions(request *GetUploadDocumentJobRequest, runtime *dara.RuntimeOptions) (_result *GetUploadDocumentJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Collection) {
		body["Collection"] = request.Collection
	}

	if !dara.IsNil(request.JobId) {
		body["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespacePassword) {
		body["NamespacePassword"] = request.NamespacePassword
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUploadDocumentJob"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUploadDocumentJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the progress and result of an asynchronous document upload job based on the job ID.
//
// Description:
//
// This operation is related to the UploadDocumentAsync operation. You can call the UploadDocumentAsync operation to create an upload job and obtain the job ID, and then call the GetUploadDocumentJob operation to query the execution information of the job.
//
// >  Suggestions:
//
//   - Determine whether the document upload job times out based on the document complexity and the number of tokens after chunking. In most cases, a job that lasts more than 2 hours is considered timeout.
//
// @param request - GetUploadDocumentJobRequest
//
// @return GetUploadDocumentJobResponse
func (client *Client) GetUploadDocumentJob(request *GetUploadDocumentJobRequest) (_result *GetUploadDocumentJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUploadDocumentJobResponse{}
	_body, _err := client.GetUploadDocumentJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the progress and result of an asynchronous vector data upload job by using a job ID.
//
// Description:
//
// This operation is related to the `UpsertCollectionDataAsync` operation. You can call the `UpsertCollectionDataAsync` operation to create an upload job and obtain a job ID, and then call the GetUpsertCollectionDataJob operation to query the execution information of the job.
//
// >  We recommend that you evaluate the amount of time required for the upload job based on 1,000 data entries every second, and then query the job progress every 5 seconds. The timeout period can be set to 30 minutes after the evaluated amount of time.
//
// @param request - GetUpsertCollectionDataJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUpsertCollectionDataJobResponse
func (client *Client) GetUpsertCollectionDataJobWithOptions(request *GetUpsertCollectionDataJobRequest, runtime *dara.RuntimeOptions) (_result *GetUpsertCollectionDataJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Collection) {
		body["Collection"] = request.Collection
	}

	if !dara.IsNil(request.JobId) {
		body["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespacePassword) {
		body["NamespacePassword"] = request.NamespacePassword
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUpsertCollectionDataJob"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUpsertCollectionDataJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the progress and result of an asynchronous vector data upload job by using a job ID.
//
// Description:
//
// This operation is related to the `UpsertCollectionDataAsync` operation. You can call the `UpsertCollectionDataAsync` operation to create an upload job and obtain a job ID, and then call the GetUpsertCollectionDataJob operation to query the execution information of the job.
//
// >  We recommend that you evaluate the amount of time required for the upload job based on 1,000 data entries every second, and then query the job progress every 5 seconds. The timeout period can be set to 30 minutes after the evaluated amount of time.
//
// @param request - GetUpsertCollectionDataJobRequest
//
// @return GetUpsertCollectionDataJobResponse
func (client *Client) GetUpsertCollectionDataJob(request *GetUpsertCollectionDataJobRequest) (_result *GetUpsertCollectionDataJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUpsertCollectionDataJobResponse{}
	_body, _err := client.GetUpsertCollectionDataJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Grants vector collection permissions to a namespace.
//
// @param request - GrantCollectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantCollectionResponse
func (client *Client) GrantCollectionWithOptions(request *GrantCollectionRequest, runtime *dara.RuntimeOptions) (_result *GrantCollectionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Collection) {
		query["Collection"] = request.Collection
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.GrantToNamespace) {
		query["GrantToNamespace"] = request.GrantToNamespace
	}

	if !dara.IsNil(request.GrantType) {
		query["GrantType"] = request.GrantType
	}

	if !dara.IsNil(request.ManagerAccount) {
		query["ManagerAccount"] = request.ManagerAccount
	}

	if !dara.IsNil(request.ManagerAccountPassword) {
		query["ManagerAccountPassword"] = request.ManagerAccountPassword
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GrantCollection"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GrantCollectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants vector collection permissions to a namespace.
//
// @param request - GrantCollectionRequest
//
// @return GrantCollectionResponse
func (client *Client) GrantCollection(request *GrantCollectionRequest) (_result *GrantCollectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GrantCollectionResponse{}
	_body, _err := client.GrantCollectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Processes active queries.
//
// @param request - HandleActiveSQLRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HandleActiveSQLRecordResponse
func (client *Client) HandleActiveSQLRecordWithOptions(request *HandleActiveSQLRecordRequest, runtime *dara.RuntimeOptions) (_result *HandleActiveSQLRecordResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OperateType) {
		query["OperateType"] = request.OperateType
	}

	if !dara.IsNil(request.Pids) {
		query["Pids"] = request.Pids
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HandleActiveSQLRecord"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &HandleActiveSQLRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Processes active queries.
//
// @param request - HandleActiveSQLRecordRequest
//
// @return HandleActiveSQLRecordResponse
func (client *Client) HandleActiveSQLRecord(request *HandleActiveSQLRecordRequest) (_result *HandleActiveSQLRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &HandleActiveSQLRecordResponse{}
	_body, _err := client.HandleActiveSQLRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Initializes vector databases.
//
// @param request - InitVectorDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitVectorDatabaseResponse
func (client *Client) InitVectorDatabaseWithOptions(request *InitVectorDatabaseRequest, runtime *dara.RuntimeOptions) (_result *InitVectorDatabaseResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.ManagerAccount) {
		query["ManagerAccount"] = request.ManagerAccount
	}

	if !dara.IsNil(request.ManagerAccountPassword) {
		query["ManagerAccountPassword"] = request.ManagerAccountPassword
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InitVectorDatabase"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InitVectorDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initializes vector databases.
//
// @param request - InitVectorDatabaseRequest
//
// @return InitVectorDatabaseResponse
func (client *Client) InitVectorDatabase(request *InitVectorDatabaseRequest) (_result *InitVectorDatabaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InitVectorDatabaseResponse{}
	_body, _err := client.InitVectorDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取备份任务列表
//
// @param request - ListBackupJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBackupJobsResponse
func (client *Client) ListBackupJobsWithOptions(request *ListBackupJobsRequest, runtime *dara.RuntimeOptions) (_result *ListBackupJobsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupMode) {
		query["BackupMode"] = request.BackupMode
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBackupJobs"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBackupJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取备份任务列表
//
// @param request - ListBackupJobsRequest
//
// @return ListBackupJobsResponse
func (client *Client) ListBackupJobs(request *ListBackupJobsRequest) (_result *ListBackupJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBackupJobsResponse{}
	_body, _err := client.ListBackupJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of vector collections.
//
// @param request - ListCollectionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCollectionsResponse
func (client *Client) ListCollectionsWithOptions(request *ListCollectionsRequest, runtime *dara.RuntimeOptions) (_result *ListCollectionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespacePassword) {
		query["NamespacePassword"] = request.NamespacePassword
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCollections"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCollectionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of vector collections.
//
// @param request - ListCollectionsRequest
//
// @return ListCollectionsResponse
func (client *Client) ListCollections(request *ListCollectionsRequest) (_result *ListCollectionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCollectionsResponse{}
	_body, _err := client.ListCollectionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of databases.
//
// @param request - ListDatabasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatabasesResponse
func (client *Client) ListDatabasesWithOptions(request *ListDatabasesRequest, runtime *dara.RuntimeOptions) (_result *ListDatabasesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Database) {
		query["Database"] = request.Database
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecretArn) {
		query["SecretArn"] = request.SecretArn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDatabases"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Queries a list of databases.
//
// @param request - ListDatabasesRequest
//
// @return ListDatabasesResponse
func (client *Client) ListDatabases(request *ListDatabasesRequest) (_result *ListDatabasesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDatabasesResponse{}
	_body, _err := client.ListDatabasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of document collections.
//
// @param request - ListDocumentCollectionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDocumentCollectionsResponse
func (client *Client) ListDocumentCollectionsWithOptions(request *ListDocumentCollectionsRequest, runtime *dara.RuntimeOptions) (_result *ListDocumentCollectionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespacePassword) {
		query["NamespacePassword"] = request.NamespacePassword
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDocumentCollections"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDocumentCollectionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of document collections.
//
// @param request - ListDocumentCollectionsRequest
//
// @return ListDocumentCollectionsResponse
func (client *Client) ListDocumentCollections(request *ListDocumentCollectionsRequest) (_result *ListDocumentCollectionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDocumentCollectionsResponse{}
	_body, _err := client.ListDocumentCollectionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of documents in a collection.
//
// @param request - ListDocumentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDocumentsResponse
func (client *Client) ListDocumentsWithOptions(request *ListDocumentsRequest, runtime *dara.RuntimeOptions) (_result *ListDocumentsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Collection) {
		query["Collection"] = request.Collection
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespacePassword) {
		query["NamespacePassword"] = request.NamespacePassword
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDocuments"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDocumentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of documents in a collection.
//
// @param request - ListDocumentsRequest
//
// @return ListDocumentsResponse
func (client *Client) ListDocuments(request *ListDocumentsRequest) (_result *ListDocumentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDocumentsResponse{}
	_body, _err := client.ListDocumentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of external data services.
//
// @param request - ListExternalDataServicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExternalDataServicesResponse
func (client *Client) ListExternalDataServicesWithOptions(request *ListExternalDataServicesRequest, runtime *dara.RuntimeOptions) (_result *ListExternalDataServicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("ListExternalDataServices"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListExternalDataServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of external data services.
//
// @param request - ListExternalDataServicesRequest
//
// @return ListExternalDataServicesResponse
func (client *Client) ListExternalDataServices(request *ListExternalDataServicesRequest) (_result *ListExternalDataServicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListExternalDataServicesResponse{}
	_body, _err := client.ListExternalDataServicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实例外表配置列表
//
// @param request - ListExternalDataSourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExternalDataSourcesResponse
func (client *Client) ListExternalDataSourcesWithOptions(request *ListExternalDataSourcesRequest, runtime *dara.RuntimeOptions) (_result *ListExternalDataSourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("ListExternalDataSources"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListExternalDataSourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实例外表配置列表
//
// @param request - ListExternalDataSourcesRequest
//
// @return ListExternalDataSourcesResponse
func (client *Client) ListExternalDataSources(request *ListExternalDataSourcesRequest) (_result *ListExternalDataSourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListExternalDataSourcesResponse{}
	_body, _err := client.ListExternalDataSourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取索引列表
//
// @param request - ListIndicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIndicesResponse
func (client *Client) ListIndicesWithOptions(request *ListIndicesRequest, runtime *dara.RuntimeOptions) (_result *ListIndicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Collection) {
		query["Collection"] = request.Collection
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespacePassword) {
		query["NamespacePassword"] = request.NamespacePassword
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIndices"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIndicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取索引列表
//
// @param request - ListIndicesRequest
//
// @return ListIndicesResponse
func (client *Client) ListIndices(request *ListIndicesRequest) (_result *ListIndicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListIndicesResponse{}
	_body, _err := client.ListIndicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of extensions.
//
// @param request - ListInstanceExtensionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceExtensionsResponse
func (client *Client) ListInstanceExtensionsWithOptions(request *ListInstanceExtensionsRequest, runtime *dara.RuntimeOptions) (_result *ListInstanceExtensionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Extension) {
		query["Extension"] = request.Extension
	}

	if !dara.IsNil(request.InstallStatus) {
		query["InstallStatus"] = request.InstallStatus
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
		Action:      dara.String("ListInstanceExtensions"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceExtensionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of extensions.
//
// @param request - ListInstanceExtensionsRequest
//
// @return ListInstanceExtensionsResponse
func (client *Client) ListInstanceExtensions(request *ListInstanceExtensionsRequest) (_result *ListInstanceExtensionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInstanceExtensionsResponse{}
	_body, _err := client.ListInstanceExtensionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of namespaces.
//
// @param request - ListNamespacesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNamespacesResponse
func (client *Client) ListNamespacesWithOptions(request *ListNamespacesRequest, runtime *dara.RuntimeOptions) (_result *ListNamespacesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.ManagerAccount) {
		query["ManagerAccount"] = request.ManagerAccount
	}

	if !dara.IsNil(request.ManagerAccountPassword) {
		query["ManagerAccountPassword"] = request.ManagerAccountPassword
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNamespaces"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNamespacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of namespaces.
//
// @param request - ListNamespacesRequest
//
// @return ListNamespacesResponse
func (client *Client) ListNamespaces(request *ListNamespacesRequest) (_result *ListNamespacesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNamespacesResponse{}
	_body, _err := client.ListNamespacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Homogeneous Data Source
//
// @param request - ListRemoteADBDataSourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRemoteADBDataSourcesResponse
func (client *Client) ListRemoteADBDataSourcesWithOptions(request *ListRemoteADBDataSourcesRequest, runtime *dara.RuntimeOptions) (_result *ListRemoteADBDataSourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DataSourceId) {
		query["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRemoteADBDataSources"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRemoteADBDataSourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Homogeneous Data Source
//
// @param request - ListRemoteADBDataSourcesRequest
//
// @return ListRemoteADBDataSourcesResponse
func (client *Client) ListRemoteADBDataSources(request *ListRemoteADBDataSourcesRequest) (_result *ListRemoteADBDataSourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRemoteADBDataSourcesResponse{}
	_body, _err := client.ListRemoteADBDataSourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of schemas.
//
// @param request - ListSchemasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSchemasResponse
func (client *Client) ListSchemasWithOptions(request *ListSchemasRequest, runtime *dara.RuntimeOptions) (_result *ListSchemasResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Database) {
		query["Database"] = request.Database
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SchemaPattern) {
		query["SchemaPattern"] = request.SchemaPattern
	}

	if !dara.IsNil(request.SecretArn) {
		query["SecretArn"] = request.SecretArn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSchemas"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSchemasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of schemas.
//
// @param request - ListSchemasRequest
//
// @return ListSchemasResponse
func (client *Client) ListSchemas(request *ListSchemasRequest) (_result *ListSchemasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSchemasResponse{}
	_body, _err := client.ListSchemasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of access credentials.
//
// @param request - ListSecretsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSecretsResponse
func (client *Client) ListSecretsWithOptions(request *ListSecretsRequest, runtime *dara.RuntimeOptions) (_result *ListSecretsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSecrets"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSecretsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of access credentials.
//
// @param request - ListSecretsRequest
//
// @return ListSecretsResponse
func (client *Client) ListSecrets(request *ListSecretsRequest) (_result *ListSecretsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSecretsResponse{}
	_body, _err := client.ListSecretsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create External Data Source Configuration
//
// @param request - ListStreamingDataServicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListStreamingDataServicesResponse
func (client *Client) ListStreamingDataServicesWithOptions(request *ListStreamingDataServicesRequest, runtime *dara.RuntimeOptions) (_result *ListStreamingDataServicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("ListStreamingDataServices"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListStreamingDataServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create External Data Source Configuration
//
// @param request - ListStreamingDataServicesRequest
//
// @return ListStreamingDataServicesResponse
func (client *Client) ListStreamingDataServices(request *ListStreamingDataServicesRequest) (_result *ListStreamingDataServicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListStreamingDataServicesResponse{}
	_body, _err := client.ListStreamingDataServicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries real-time service data sources.
//
// @param request - ListStreamingDataSourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListStreamingDataSourcesResponse
func (client *Client) ListStreamingDataSourcesWithOptions(request *ListStreamingDataSourcesRequest, runtime *dara.RuntimeOptions) (_result *ListStreamingDataSourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("ListStreamingDataSources"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListStreamingDataSourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries real-time service data sources.
//
// @param request - ListStreamingDataSourcesRequest
//
// @return ListStreamingDataSourcesResponse
func (client *Client) ListStreamingDataSources(request *ListStreamingDataSourcesRequest) (_result *ListStreamingDataSourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListStreamingDataSourcesResponse{}
	_body, _err := client.ListStreamingDataSourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries real-time data synchronization jobs.
//
// @param request - ListStreamingJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListStreamingJobsResponse
func (client *Client) ListStreamingJobsWithOptions(request *ListStreamingJobsRequest, runtime *dara.RuntimeOptions) (_result *ListStreamingJobsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("ListStreamingJobs"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListStreamingJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries real-time data synchronization jobs.
//
// @param request - ListStreamingJobsRequest
//
// @return ListStreamingJobsResponse
func (client *Client) ListStreamingJobs(request *ListStreamingJobsRequest) (_result *ListStreamingJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListStreamingJobsResponse{}
	_body, _err := client.ListStreamingJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询Supabase实例列表
//
// @param request - ListSupabaseProjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSupabaseProjectsResponse
func (client *Client) ListSupabaseProjectsWithOptions(request *ListSupabaseProjectsRequest, runtime *dara.RuntimeOptions) (_result *ListSupabaseProjectsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSupabaseProjects"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSupabaseProjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Supabase实例列表
//
// @param request - ListSupabaseProjectsRequest
//
// @return ListSupabaseProjectsResponse
func (client *Client) ListSupabaseProjects(request *ListSupabaseProjectsRequest) (_result *ListSupabaseProjectsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSupabaseProjectsResponse{}
	_body, _err := client.ListSupabaseProjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of tables in a database.
//
// @param request - ListTablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTablesResponse
func (client *Client) ListTablesWithOptions(request *ListTablesRequest, runtime *dara.RuntimeOptions) (_result *ListTablesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Database) {
		query["Database"] = request.Database
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Schema) {
		query["Schema"] = request.Schema
	}

	if !dara.IsNil(request.SecretArn) {
		query["SecretArn"] = request.SecretArn
	}

	if !dara.IsNil(request.TablePattern) {
		query["TablePattern"] = request.TablePattern
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTables"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of tables in a database.
//
// @param request - ListTablesRequest
//
// @return ListTablesResponse
func (client *Client) ListTables(request *ListTablesRequest) (_result *ListTablesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTablesResponse{}
	_body, _err := client.ListTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of tags that are added to AnalyticDB for PostgreSQL instances.
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
		Version:     dara.String("2016-05-03"),
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
// Queries a list of tags that are added to AnalyticDB for PostgreSQL instances.
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
// Modifies the description of a database account for an AnalyticDB for PostgreSQL instance.
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAccountDescription"),
		Version:     dara.String("2016-05-03"),
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
// Modifies the description of a database account for an AnalyticDB for PostgreSQL instance.
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
// Configures the backup policy of an AnalyticDB for PostgreSQL instance.
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

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.EnableRecoveryPoint) {
		query["EnableRecoveryPoint"] = request.EnableRecoveryPoint
	}

	if !dara.IsNil(request.PreferredBackupPeriod) {
		query["PreferredBackupPeriod"] = request.PreferredBackupPeriod
	}

	if !dara.IsNil(request.PreferredBackupTime) {
		query["PreferredBackupTime"] = request.PreferredBackupTime
	}

	if !dara.IsNil(request.RecoveryPointPeriod) {
		query["RecoveryPointPeriod"] = request.RecoveryPointPeriod
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBackupPolicy"),
		Version:     dara.String("2016-05-03"),
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
// Configures the backup policy of an AnalyticDB for PostgreSQL instance.
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
// 更新Collection
//
// @param request - ModifyCollectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCollectionResponse
func (client *Client) ModifyCollectionWithOptions(request *ModifyCollectionRequest, runtime *dara.RuntimeOptions) (_result *ModifyCollectionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Collection) {
		query["Collection"] = request.Collection
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Metadata) {
		query["Metadata"] = request.Metadata
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespacePassword) {
		query["NamespacePassword"] = request.NamespacePassword
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCollection"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCollectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新Collection
//
// @param request - ModifyCollectionRequest
//
// @return ModifyCollectionResponse
func (client *Client) ModifyCollection(request *ModifyCollectionRequest) (_result *ModifyCollectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyCollectionResponse{}
	_body, _err := client.ModifyCollectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the threshold of computing resources and the wait period of idle resources for an AnalyticDB for PostgreSQL instance in Serverless automatic scheduling mode.
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
	if !dara.IsNil(request.DBInstanceDescription) {
		query["DBInstanceDescription"] = request.DBInstanceDescription
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.IdleTime) {
		query["IdleTime"] = request.IdleTime
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ServerlessResource) {
		query["ServerlessResource"] = request.ServerlessResource
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceConfig"),
		Version:     dara.String("2016-05-03"),
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
// Changes the threshold of computing resources and the wait period of idle resources for an AnalyticDB for PostgreSQL instance in Serverless automatic scheduling mode.
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
// Changes the endpoint of an AnalyticDB for PostgreSQL instance.
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
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConnectionStringPrefix) {
		query["ConnectionStringPrefix"] = request.ConnectionStringPrefix
	}

	if !dara.IsNil(request.CurrentConnectionString) {
		query["CurrentConnectionString"] = request.CurrentConnectionString
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceConnectionString"),
		Version:     dara.String("2016-05-03"),
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
// Changes the endpoint of an AnalyticDB for PostgreSQL instance.
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
// 修改实例部署模式
//
// @param request - ModifyDBInstanceDeploymentModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceDeploymentModeResponse
func (client *Client) ModifyDBInstanceDeploymentModeWithOptions(request *ModifyDBInstanceDeploymentModeRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceDeploymentModeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DeployMode) {
		query["DeployMode"] = request.DeployMode
	}

	if !dara.IsNil(request.StandbyVSwitchId) {
		query["StandbyVSwitchId"] = request.StandbyVSwitchId
	}

	if !dara.IsNil(request.StandbyZoneId) {
		query["StandbyZoneId"] = request.StandbyZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceDeploymentMode"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceDeploymentModeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改实例部署模式
//
// @param request - ModifyDBInstanceDeploymentModeRequest
//
// @return ModifyDBInstanceDeploymentModeResponse
func (client *Client) ModifyDBInstanceDeploymentMode(request *ModifyDBInstanceDeploymentModeRequest) (_result *ModifyDBInstanceDeploymentModeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBInstanceDeploymentModeResponse{}
	_body, _err := client.ModifyDBInstanceDeploymentModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the description of an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// To make it easy to identify AnalyticDB for PostgreSQL instances, you can call this operation to modify the description of instances.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
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

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceDescription"),
		Version:     dara.String("2016-05-03"),
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

// Summary:
//
// Changes the description of an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// To make it easy to identify AnalyticDB for PostgreSQL instances, you can call this operation to modify the description of instances.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
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

// Summary:
//
// Modifies the maintenance window of an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// The system maintains AnalyticDB for PostgreSQL instances during the maintenance window that you specify. We recommend that you set the maintenance window to off-peak hours to minimize the impact on your business.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyDBInstanceMaintainTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceMaintainTimeResponse
func (client *Client) ModifyDBInstanceMaintainTimeWithOptions(request *ModifyDBInstanceMaintainTimeRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceMaintainTimeResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceMaintainTime"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceMaintainTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the maintenance window of an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// The system maintains AnalyticDB for PostgreSQL instances during the maintenance window that you specify. We recommend that you set the maintenance window to off-peak hours to minimize the impact on your business.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyDBInstanceMaintainTimeRequest
//
// @return ModifyDBInstanceMaintainTimeResponse
func (client *Client) ModifyDBInstanceMaintainTime(request *ModifyDBInstanceMaintainTimeRequest) (_result *ModifyDBInstanceMaintainTimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBInstanceMaintainTimeResponse{}
	_body, _err := client.ModifyDBInstanceMaintainTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the network type of an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// ##
//
// This operation is available only for AnalyticDB for PostgreSQL instances in reserved storage mode.
//
// ## QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyDBInstanceNetworkTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceNetworkTypeResponse
func (client *Client) ModifyDBInstanceNetworkTypeWithOptions(request *ModifyDBInstanceNetworkTypeRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceNetworkTypeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.InstanceNetworkType) {
		query["InstanceNetworkType"] = request.InstanceNetworkType
	}

	if !dara.IsNil(request.PrivateIpAddress) {
		query["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !dara.IsNil(request.VPCId) {
		query["VPCId"] = request.VPCId
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceNetworkType"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceNetworkTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the network type of an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// ##
//
// This operation is available only for AnalyticDB for PostgreSQL instances in reserved storage mode.
//
// ## QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyDBInstanceNetworkTypeRequest
//
// @return ModifyDBInstanceNetworkTypeResponse
func (client *Client) ModifyDBInstanceNetworkType(request *ModifyDBInstanceNetworkTypeRequest) (_result *ModifyDBInstanceNetworkTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBInstanceNetworkTypeResponse{}
	_body, _err := client.ModifyDBInstanceNetworkTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 包年包月/按量付费转换改造
//
// @param request - ModifyDBInstancePayTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstancePayTypeResponse
func (client *Client) ModifyDBInstancePayTypeWithOptions(request *ModifyDBInstancePayTypeRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstancePayTypeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.UsedTime) {
		query["UsedTime"] = request.UsedTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstancePayType"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstancePayTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 包年包月/按量付费转换改造
//
// @param request - ModifyDBInstancePayTypeRequest
//
// @return ModifyDBInstancePayTypeResponse
func (client *Client) ModifyDBInstancePayType(request *ModifyDBInstancePayTypeRequest) (_result *ModifyDBInstancePayTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBInstancePayTypeResponse{}
	_body, _err := client.ModifyDBInstancePayTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Moves an AnalyticDB for PostgreSQL instance to a resource group.
//
// Description:
//
// Resource Management allows you to build an organizational structure for resources based on your business requirements. You can use resource directories, folders, accounts, and resource groups to hierarchically organize and manage resources. For more information, see [What is Resource Management?](https://help.aliyun.com/document_detail/94475.html)
//
// @param request - ModifyDBInstanceResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceResourceGroupResponse
func (client *Client) ModifyDBInstanceResourceGroupWithOptions(request *ModifyDBInstanceResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceResourceGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.NewResourceGroupId) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceResourceGroup"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves an AnalyticDB for PostgreSQL instance to a resource group.
//
// Description:
//
// Resource Management allows you to build an organizational structure for resources based on your business requirements. You can use resource directories, folders, accounts, and resource groups to hierarchically organize and manage resources. For more information, see [What is Resource Management?](https://help.aliyun.com/document_detail/94475.html)
//
// @param request - ModifyDBInstanceResourceGroupRequest
//
// @return ModifyDBInstanceResourceGroupResponse
func (client *Client) ModifyDBInstanceResourceGroup(request *ModifyDBInstanceResourceGroupRequest) (_result *ModifyDBInstanceResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBInstanceResourceGroupResponse{}
	_body, _err := client.ModifyDBInstanceResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables, disables, or updates SSL encryption for an AnalyticDB for PostgreSQL instance.
//
// @param request - ModifyDBInstanceSSLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceSSLResponse
func (client *Client) ModifyDBInstanceSSLWithOptions(request *ModifyDBInstanceSSLRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceSSLResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionString) {
		query["ConnectionString"] = request.ConnectionString
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.SSLEnabled) {
		query["SSLEnabled"] = request.SSLEnabled
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceSSL"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceSSLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables, disables, or updates SSL encryption for an AnalyticDB for PostgreSQL instance.
//
// @param request - ModifyDBInstanceSSLRequest
//
// @return ModifyDBInstanceSSLResponse
func (client *Client) ModifyDBInstanceSSL(request *ModifyDBInstanceSSLRequest) (_result *ModifyDBInstanceSSLResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBInstanceSSLResponse{}
	_body, _err := client.ModifyDBInstanceSSLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a resource group.
//
// @param tmpReq - ModifyDBResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBResourceGroupResponse
func (client *Client) ModifyDBResourceGroupWithOptions(tmpReq *ModifyDBResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBResourceGroupResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ModifyDBResourceGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceGroupItems) {
		request.ResourceGroupItemsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceGroupItems, dara.String("ResourceGroupItems"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceGroupItemsShrink) {
		query["ResourceGroupItems"] = request.ResourceGroupItemsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBResourceGroup"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a resource group.
//
// @param request - ModifyDBResourceGroupRequest
//
// @return ModifyDBResourceGroupResponse
func (client *Client) ModifyDBResourceGroup(request *ModifyDBResourceGroupRequest) (_result *ModifyDBResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBResourceGroupResponse{}
	_body, _err := client.ModifyDBResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Modify External Data Service
//
// @param request - ModifyExternalDataServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyExternalDataServiceResponse
func (client *Client) ModifyExternalDataServiceWithOptions(request *ModifyExternalDataServiceRequest, runtime *dara.RuntimeOptions) (_result *ModifyExternalDataServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceDescription) {
		query["ServiceDescription"] = request.ServiceDescription
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceSpec) {
		query["ServiceSpec"] = request.ServiceSpec
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyExternalDataService"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyExternalDataServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify External Data Service
//
// @param request - ModifyExternalDataServiceRequest
//
// @return ModifyExternalDataServiceResponse
func (client *Client) ModifyExternalDataService(request *ModifyExternalDataServiceRequest) (_result *ModifyExternalDataServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyExternalDataServiceResponse{}
	_body, _err := client.ModifyExternalDataServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a Hadoop data source.
//
// @param request - ModifyHadoopDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHadoopDataSourceResponse
func (client *Client) ModifyHadoopDataSourceWithOptions(request *ModifyHadoopDataSourceRequest, runtime *dara.RuntimeOptions) (_result *ModifyHadoopDataSourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DataSourceDescription) {
		query["DataSourceDescription"] = request.DataSourceDescription
	}

	if !dara.IsNil(request.DataSourceId) {
		query["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.DataSourceType) {
		query["DataSourceType"] = request.DataSourceType
	}

	if !dara.IsNil(request.EmrInstanceId) {
		query["EmrInstanceId"] = request.EmrInstanceId
	}

	if !dara.IsNil(request.HDFSConf) {
		query["HDFSConf"] = request.HDFSConf
	}

	if !dara.IsNil(request.HadoopCoreConf) {
		query["HadoopCoreConf"] = request.HadoopCoreConf
	}

	if !dara.IsNil(request.HadoopCreateType) {
		query["HadoopCreateType"] = request.HadoopCreateType
	}

	if !dara.IsNil(request.HadoopHostsAddress) {
		query["HadoopHostsAddress"] = request.HadoopHostsAddress
	}

	if !dara.IsNil(request.HiveConf) {
		query["HiveConf"] = request.HiveConf
	}

	if !dara.IsNil(request.MapReduceConf) {
		query["MapReduceConf"] = request.MapReduceConf
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.YarnConf) {
		query["YarnConf"] = request.YarnConf
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHadoopDataSource"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHadoopDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a Hadoop data source.
//
// @param request - ModifyHadoopDataSourceRequest
//
// @return ModifyHadoopDataSourceResponse
func (client *Client) ModifyHadoopDataSource(request *ModifyHadoopDataSourceRequest) (_result *ModifyHadoopDataSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyHadoopDataSourceResponse{}
	_body, _err := client.ModifyHadoopDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a Java Database Connectivity (JDBC) data source.
//
// @param request - ModifyJDBCDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyJDBCDataSourceResponse
func (client *Client) ModifyJDBCDataSourceWithOptions(request *ModifyJDBCDataSourceRequest, runtime *dara.RuntimeOptions) (_result *ModifyJDBCDataSourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DataSourceDescription) {
		query["DataSourceDescription"] = request.DataSourceDescription
	}

	if !dara.IsNil(request.DataSourceId) {
		query["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.DataSourceType) {
		query["DataSourceType"] = request.DataSourceType
	}

	if !dara.IsNil(request.JDBCConnectionString) {
		query["JDBCConnectionString"] = request.JDBCConnectionString
	}

	if !dara.IsNil(request.JDBCPassword) {
		query["JDBCPassword"] = request.JDBCPassword
	}

	if !dara.IsNil(request.JDBCUserName) {
		query["JDBCUserName"] = request.JDBCUserName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyJDBCDataSource"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyJDBCDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a Java Database Connectivity (JDBC) data source.
//
// @param request - ModifyJDBCDataSourceRequest
//
// @return ModifyJDBCDataSourceResponse
func (client *Client) ModifyJDBCDataSource(request *ModifyJDBCDataSourceRequest) (_result *ModifyJDBCDataSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyJDBCDataSourceResponse{}
	_body, _err := client.ModifyJDBCDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the specifications of coordinator node resources for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// This operation is not available for instances in reserved storage mode.
//
// Before you call this operation, make sure that you are familiar with the billing of AnalyticDB for PostgreSQL. For more information, see [Billing methods](https://help.aliyun.com/document_detail/35406.html) and [AnalyticDB for PostgreSQL pricing](https://www.alibabacloud.com/zh/product/hybriddb-postgresql/pricing).
//
// @param request - ModifyMasterSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMasterSpecResponse
func (client *Client) ModifyMasterSpecWithOptions(request *ModifyMasterSpecRequest, runtime *dara.RuntimeOptions) (_result *ModifyMasterSpecResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceDescription) {
		query["DBInstanceDescription"] = request.DBInstanceDescription
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.MasterAISpec) {
		query["MasterAISpec"] = request.MasterAISpec
	}

	if !dara.IsNil(request.MasterCU) {
		query["MasterCU"] = request.MasterCU
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyMasterSpec"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyMasterSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the specifications of coordinator node resources for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// This operation is not available for instances in reserved storage mode.
//
// Before you call this operation, make sure that you are familiar with the billing of AnalyticDB for PostgreSQL. For more information, see [Billing methods](https://help.aliyun.com/document_detail/35406.html) and [AnalyticDB for PostgreSQL pricing](https://www.alibabacloud.com/zh/product/hybriddb-postgresql/pricing).
//
// @param request - ModifyMasterSpecRequest
//
// @return ModifyMasterSpecResponse
func (client *Client) ModifyMasterSpec(request *ModifyMasterSpecRequest) (_result *ModifyMasterSpecResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyMasterSpecResponse{}
	_body, _err := client.ModifyMasterSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configuration parameters of an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// This operation can be called to modify parameters of an AnalyticDB for PostgreSQL instance in elastic storage mode or Serverless mode.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered and may affect your business. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyParametersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyParametersResponse
func (client *Client) ModifyParametersWithOptions(request *ModifyParametersRequest, runtime *dara.RuntimeOptions) (_result *ModifyParametersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.ForceRestartInstance) {
		query["ForceRestartInstance"] = request.ForceRestartInstance
	}

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyParameters"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyParametersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration parameters of an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// This operation can be called to modify parameters of an AnalyticDB for PostgreSQL instance in elastic storage mode or Serverless mode.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered and may affect your business. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyParametersRequest
//
// @return ModifyParametersResponse
func (client *Client) ModifyParameters(request *ModifyParametersRequest) (_result *ModifyParametersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyParametersResponse{}
	_body, _err := client.ModifyParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Modify Homogeneous Data Source
//
// @param request - ModifyRemoteADBDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRemoteADBDataSourceResponse
func (client *Client) ModifyRemoteADBDataSourceWithOptions(request *ModifyRemoteADBDataSourceRequest, runtime *dara.RuntimeOptions) (_result *ModifyRemoteADBDataSourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataSourceId) {
		query["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.DataSourceName) {
		query["DataSourceName"] = request.DataSourceName
	}

	if !dara.IsNil(request.LocalDBInstanceId) {
		query["LocalDBInstanceId"] = request.LocalDBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	if !dara.IsNil(request.UserPassword) {
		query["UserPassword"] = request.UserPassword
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyRemoteADBDataSource"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyRemoteADBDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify Homogeneous Data Source
//
// @param request - ModifyRemoteADBDataSourceRequest
//
// @return ModifyRemoteADBDataSourceResponse
func (client *Client) ModifyRemoteADBDataSource(request *ModifyRemoteADBDataSourceRequest) (_result *ModifyRemoteADBDataSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyRemoteADBDataSourceResponse{}
	_body, _err := client.ModifyRemoteADBDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables the SQL Explorer feature for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
//	  You can call this operation only for AnalyticDB for PostgreSQL instances in reserved storage mode.
//
//		- You can call this operation only for AnalyticDB for PostgreSQL instances in Serverless automatic scheduling mode.
//
// @param request - ModifySQLCollectorPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySQLCollectorPolicyResponse
func (client *Client) ModifySQLCollectorPolicyWithOptions(request *ModifySQLCollectorPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifySQLCollectorPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.SQLCollectorStatus) {
		query["SQLCollectorStatus"] = request.SQLCollectorStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySQLCollectorPolicy"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySQLCollectorPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the SQL Explorer feature for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
//	  You can call this operation only for AnalyticDB for PostgreSQL instances in reserved storage mode.
//
//		- You can call this operation only for AnalyticDB for PostgreSQL instances in Serverless automatic scheduling mode.
//
// @param request - ModifySQLCollectorPolicyRequest
//
// @return ModifySQLCollectorPolicyResponse
func (client *Client) ModifySQLCollectorPolicy(request *ModifySQLCollectorPolicyRequest) (_result *ModifySQLCollectorPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifySQLCollectorPolicyResponse{}
	_body, _err := client.ModifySQLCollectorPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the IP address whitelist of an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// To ensure the security and stability of AnalyticDB for PostgreSQL instances, the system denies all external IP addresses to access AnalyticDB for PostgreSQL instances by default. Before you can use an AnalyticDB for PostgreSQL instance, you must add the IP address or CIDR block of your client to the IP address whitelist of the instance.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
	if !dara.IsNil(request.DBInstanceIPArrayAttribute) {
		query["DBInstanceIPArrayAttribute"] = request.DBInstanceIPArrayAttribute
	}

	if !dara.IsNil(request.DBInstanceIPArrayName) {
		query["DBInstanceIPArrayName"] = request.DBInstanceIPArrayName
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.ModifyMode) {
		query["ModifyMode"] = request.ModifyMode
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SecurityIPList) {
		query["SecurityIPList"] = request.SecurityIPList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySecurityIps"),
		Version:     dara.String("2016-05-03"),
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
// Modifies the IP address whitelist of an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// To ensure the security and stability of AnalyticDB for PostgreSQL instances, the system denies all external IP addresses to access AnalyticDB for PostgreSQL instances by default. Before you can use an AnalyticDB for PostgreSQL instance, you must add the IP address or CIDR block of your client to the IP address whitelist of the instance.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// Modifies a real-time data service.
//
// @param request - ModifyStreamingDataServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyStreamingDataServiceResponse
func (client *Client) ModifyStreamingDataServiceWithOptions(request *ModifyStreamingDataServiceRequest, runtime *dara.RuntimeOptions) (_result *ModifyStreamingDataServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceDescription) {
		query["ServiceDescription"] = request.ServiceDescription
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceSpec) {
		query["ServiceSpec"] = request.ServiceSpec
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyStreamingDataService"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyStreamingDataServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a real-time data service.
//
// @param request - ModifyStreamingDataServiceRequest
//
// @return ModifyStreamingDataServiceResponse
func (client *Client) ModifyStreamingDataService(request *ModifyStreamingDataServiceRequest) (_result *ModifyStreamingDataServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyStreamingDataServiceResponse{}
	_body, _err := client.ModifyStreamingDataServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a real-time service data source.
//
// @param request - ModifyStreamingDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyStreamingDataSourceResponse
func (client *Client) ModifyStreamingDataSourceWithOptions(request *ModifyStreamingDataSourceRequest, runtime *dara.RuntimeOptions) (_result *ModifyStreamingDataSourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DataSourceConfig) {
		query["DataSourceConfig"] = request.DataSourceConfig
	}

	if !dara.IsNil(request.DataSourceDescription) {
		query["DataSourceDescription"] = request.DataSourceDescription
	}

	if !dara.IsNil(request.DataSourceId) {
		query["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyStreamingDataSource"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyStreamingDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a real-time service data source.
//
// @param request - ModifyStreamingDataSourceRequest
//
// @return ModifyStreamingDataSourceResponse
func (client *Client) ModifyStreamingDataSource(request *ModifyStreamingDataSourceRequest) (_result *ModifyStreamingDataSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyStreamingDataSourceResponse{}
	_body, _err := client.ModifyStreamingDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create External Data Source Configuration
//
// @param tmpReq - ModifyStreamingJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyStreamingJobResponse
func (client *Client) ModifyStreamingJobWithOptions(tmpReq *ModifyStreamingJobRequest, runtime *dara.RuntimeOptions) (_result *ModifyStreamingJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ModifyStreamingJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DestColumns) {
		request.DestColumnsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DestColumns, dara.String("DestColumns"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.MatchColumns) {
		request.MatchColumnsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MatchColumns, dara.String("MatchColumns"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SrcColumns) {
		request.SrcColumnsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SrcColumns, dara.String("SrcColumns"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UpdateColumns) {
		request.UpdateColumnsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateColumns, dara.String("UpdateColumns"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Account) {
		query["Account"] = request.Account
	}

	if !dara.IsNil(request.Consistency) {
		query["Consistency"] = request.Consistency
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DestColumnsShrink) {
		query["DestColumns"] = request.DestColumnsShrink
	}

	if !dara.IsNil(request.DestDatabase) {
		query["DestDatabase"] = request.DestDatabase
	}

	if !dara.IsNil(request.DestSchema) {
		query["DestSchema"] = request.DestSchema
	}

	if !dara.IsNil(request.DestTable) {
		query["DestTable"] = request.DestTable
	}

	if !dara.IsNil(request.ErrorLimitCount) {
		query["ErrorLimitCount"] = request.ErrorLimitCount
	}

	if !dara.IsNil(request.FallbackOffset) {
		query["FallbackOffset"] = request.FallbackOffset
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.JobConfig) {
		query["JobConfig"] = request.JobConfig
	}

	if !dara.IsNil(request.JobDescription) {
		query["JobDescription"] = request.JobDescription
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.MatchColumnsShrink) {
		query["MatchColumns"] = request.MatchColumnsShrink
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SrcColumnsShrink) {
		query["SrcColumns"] = request.SrcColumnsShrink
	}

	if !dara.IsNil(request.TryRun) {
		query["TryRun"] = request.TryRun
	}

	if !dara.IsNil(request.UpdateColumnsShrink) {
		query["UpdateColumns"] = request.UpdateColumnsShrink
	}

	if !dara.IsNil(request.WriteMode) {
		query["WriteMode"] = request.WriteMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyStreamingJob"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyStreamingJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create External Data Source Configuration
//
// @param request - ModifyStreamingJobRequest
//
// @return ModifyStreamingJobResponse
func (client *Client) ModifyStreamingJob(request *ModifyStreamingJobRequest) (_result *ModifyStreamingJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyStreamingJobResponse{}
	_body, _err := client.ModifyStreamingJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改supabase项目白名单
//
// @param request - ModifySupabaseProjectSecurityIpsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySupabaseProjectSecurityIpsResponse
func (client *Client) ModifySupabaseProjectSecurityIpsWithOptions(request *ModifySupabaseProjectSecurityIpsRequest, runtime *dara.RuntimeOptions) (_result *ModifySupabaseProjectSecurityIpsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
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
		Action:      dara.String("ModifySupabaseProjectSecurityIps"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySupabaseProjectSecurityIpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改supabase项目白名单
//
// @param request - ModifySupabaseProjectSecurityIpsRequest
//
// @return ModifySupabaseProjectSecurityIpsResponse
func (client *Client) ModifySupabaseProjectSecurityIps(request *ModifySupabaseProjectSecurityIpsRequest) (_result *ModifySupabaseProjectSecurityIpsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifySupabaseProjectSecurityIpsResponse{}
	_body, _err := client.ModifySupabaseProjectSecurityIpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the vector engine optimization configuration of an AnalyticDB for PostgreSQL instance.
//
// @param request - ModifyVectorConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyVectorConfigurationResponse
func (client *Client) ModifyVectorConfigurationWithOptions(request *ModifyVectorConfigurationRequest, runtime *dara.RuntimeOptions) (_result *ModifyVectorConfigurationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.VectorConfigurationStatus) {
		query["VectorConfigurationStatus"] = request.VectorConfigurationStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyVectorConfiguration"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyVectorConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the vector engine optimization configuration of an AnalyticDB for PostgreSQL instance.
//
// @param request - ModifyVectorConfigurationRequest
//
// @return ModifyVectorConfigurationResponse
func (client *Client) ModifyVectorConfiguration(request *ModifyVectorConfigurationRequest) (_result *ModifyVectorConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyVectorConfigurationResponse{}
	_body, _err := client.ModifyVectorConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Pauses data redistribution.
//
// @param request - PauseDataRedistributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PauseDataRedistributeResponse
func (client *Client) PauseDataRedistributeWithOptions(request *PauseDataRedistributeRequest, runtime *dara.RuntimeOptions) (_result *PauseDataRedistributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PauseDataRedistribute"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PauseDataRedistributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Pauses data redistribution.
//
// @param request - PauseDataRedistributeRequest
//
// @return PauseDataRedistributeResponse
func (client *Client) PauseDataRedistribute(request *PauseDataRedistributeRequest) (_result *PauseDataRedistributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PauseDataRedistributeResponse{}
	_body, _err := client.PauseDataRedistributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Pauses an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to pause an AnalyticDB for PostgreSQL instance that is in the **Running*	- state.
//
// This operation is available only for AnalyticDB for PostgreSQL instances in Serverless mode that run V1.0.2.1 or later. For more information about how to view and update the minor engine version of an instance, see [View the minor engine version](https://help.aliyun.com/document_detail/277424.html) and [Update the minor engine version](https://help.aliyun.com/document_detail/139271.html).
//
// >  Before you call this operation, make sure that you are familiar with the billing methods and pricing of AnalyticDB for PostgreSQL instances. For more information, see [Billing methods](https://help.aliyun.com/document_detail/35406.html) and [AnalyticDB for PostgreSQL pricing](https://www.alibabacloud.com/zh/product/hybriddb-postgresql/pricing).
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - PauseInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PauseInstanceResponse
func (client *Client) PauseInstanceWithOptions(request *PauseInstanceRequest, runtime *dara.RuntimeOptions) (_result *PauseInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PauseInstance"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PauseInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Pauses an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to pause an AnalyticDB for PostgreSQL instance that is in the **Running*	- state.
//
// This operation is available only for AnalyticDB for PostgreSQL instances in Serverless mode that run V1.0.2.1 or later. For more information about how to view and update the minor engine version of an instance, see [View the minor engine version](https://help.aliyun.com/document_detail/277424.html) and [Update the minor engine version](https://help.aliyun.com/document_detail/139271.html).
//
// >  Before you call this operation, make sure that you are familiar with the billing methods and pricing of AnalyticDB for PostgreSQL instances. For more information, see [Billing methods](https://help.aliyun.com/document_detail/35406.html) and [AnalyticDB for PostgreSQL pricing](https://www.alibabacloud.com/zh/product/hybriddb-postgresql/pricing).
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - PauseInstanceRequest
//
// @return PauseInstanceResponse
func (client *Client) PauseInstance(request *PauseInstanceRequest) (_result *PauseInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PauseInstanceResponse{}
	_body, _err := client.PauseInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Vector Data
//
// @param tmpReq - QueryCollectionDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCollectionDataResponse
func (client *Client) QueryCollectionDataWithOptions(tmpReq *QueryCollectionDataRequest, runtime *dara.RuntimeOptions) (_result *QueryCollectionDataResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &QueryCollectionDataShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.HybridSearchArgs) {
		request.HybridSearchArgsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HybridSearchArgs, dara.String("HybridSearchArgs"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RelationalTableFilter) {
		request.RelationalTableFilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RelationalTableFilter, dara.String("RelationalTableFilter"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SparseVector) {
		request.SparseVectorShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SparseVector, dara.String("SparseVector"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Vector) {
		request.VectorShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Vector, dara.String("Vector"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Collection) {
		query["Collection"] = request.Collection
	}

	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.HybridSearch) {
		query["HybridSearch"] = request.HybridSearch
	}

	if !dara.IsNil(request.HybridSearchArgsShrink) {
		query["HybridSearchArgs"] = request.HybridSearchArgsShrink
	}

	if !dara.IsNil(request.IncludeMetadataFields) {
		query["IncludeMetadataFields"] = request.IncludeMetadataFields
	}

	if !dara.IsNil(request.IncludeValues) {
		query["IncludeValues"] = request.IncludeValues
	}

	if !dara.IsNil(request.Metrics) {
		query["Metrics"] = request.Metrics
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespacePassword) {
		query["NamespacePassword"] = request.NamespacePassword
	}

	if !dara.IsNil(request.Offset) {
		query["Offset"] = request.Offset
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RelationalTableFilterShrink) {
		query["RelationalTableFilter"] = request.RelationalTableFilterShrink
	}

	if !dara.IsNil(request.SparseVectorShrink) {
		query["SparseVector"] = request.SparseVectorShrink
	}

	if !dara.IsNil(request.TopK) {
		query["TopK"] = request.TopK
	}

	if !dara.IsNil(request.VectorShrink) {
		query["Vector"] = request.VectorShrink
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
		Action:      dara.String("QueryCollectionData"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCollectionDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Vector Data
//
// @param request - QueryCollectionDataRequest
//
// @return QueryCollectionDataResponse
func (client *Client) QueryCollectionData(request *QueryCollectionDataRequest) (_result *QueryCollectionDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryCollectionDataResponse{}
	_body, _err := client.QueryCollectionDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query
//
// @param tmpReq - QueryContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryContentResponse
func (client *Client) QueryContentWithOptions(tmpReq *QueryContentRequest, runtime *dara.RuntimeOptions) (_result *QueryContentResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &QueryContentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.GraphSearchArgs) {
		request.GraphSearchArgsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GraphSearchArgs, dara.String("GraphSearchArgs"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.HybridSearchArgs) {
		request.HybridSearchArgsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HybridSearchArgs, dara.String("HybridSearchArgs"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RecallWindow) {
		request.RecallWindowShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RecallWindow, dara.String("RecallWindow"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Collection) {
		query["Collection"] = request.Collection
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileUrl) {
		query["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.GraphEnhance) {
		query["GraphEnhance"] = request.GraphEnhance
	}

	if !dara.IsNil(request.GraphSearchArgsShrink) {
		query["GraphSearchArgs"] = request.GraphSearchArgsShrink
	}

	if !dara.IsNil(request.HybridSearch) {
		query["HybridSearch"] = request.HybridSearch
	}

	if !dara.IsNil(request.HybridSearchArgsShrink) {
		query["HybridSearchArgs"] = request.HybridSearchArgsShrink
	}

	if !dara.IsNil(request.IncludeFileUrl) {
		query["IncludeFileUrl"] = request.IncludeFileUrl
	}

	if !dara.IsNil(request.IncludeMetadataFields) {
		query["IncludeMetadataFields"] = request.IncludeMetadataFields
	}

	if !dara.IsNil(request.IncludeVector) {
		query["IncludeVector"] = request.IncludeVector
	}

	if !dara.IsNil(request.Metrics) {
		query["Metrics"] = request.Metrics
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespacePassword) {
		query["NamespacePassword"] = request.NamespacePassword
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RecallWindowShrink) {
		query["RecallWindow"] = request.RecallWindowShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RerankFactor) {
		query["RerankFactor"] = request.RerankFactor
	}

	if !dara.IsNil(request.TopK) {
		query["TopK"] = request.TopK
	}

	if !dara.IsNil(request.UrlExpiration) {
		query["UrlExpiration"] = request.UrlExpiration
	}

	if !dara.IsNil(request.UseFullTextRetrieval) {
		query["UseFullTextRetrieval"] = request.UseFullTextRetrieval
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryContent"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query
//
// @param request - QueryContentRequest
//
// @return QueryContentResponse
func (client *Client) QueryContent(request *QueryContentRequest) (_result *QueryContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryContentResponse{}
	_body, _err := client.QueryContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryContentAdvance(request *QueryContentAdvanceRequest, runtime *dara.RuntimeOptions) (_result *QueryContentResponse, _err error) {
	// Step 0: init client
	if dara.IsNil(client.Credential) {
		_err = &openapi.ClientError{
			Code:    dara.String("InvalidCredentials"),
			Message: dara.String("Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details."),
		}
		return _result, _err
	}

	credentialModel, _err := client.Credential.GetCredential()
	if _err != nil {
		return _result, _err
	}

	accessKeyId := dara.StringValue(credentialModel.AccessKeyId)
	accessKeySecret := dara.StringValue(credentialModel.AccessKeySecret)
	securityToken := dara.StringValue(credentialModel.SecurityToken)
	credentialType := dara.StringValue(credentialModel.Type)
	openPlatformEndpoint := dara.StringValue(client.OpenPlatformEndpoint)
	if dara.IsNil(dara.String(openPlatformEndpoint)) || openPlatformEndpoint == "" {
		openPlatformEndpoint = "openplatform.aliyuncs.com"
	}

	if dara.IsNil(dara.String(credentialType)) {
		credentialType = "access_key"
	}

	authConfig := &openapiutil.Config{
		AccessKeyId:     dara.String(accessKeyId),
		AccessKeySecret: dara.String(accessKeySecret),
		SecurityToken:   dara.String(securityToken),
		Type:            dara.String(credentialType),
		Endpoint:        dara.String(openPlatformEndpoint),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openapi.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := map[string]*string{
		"Product":  dara.String("gpdb"),
		"RegionId": client.RegionId,
	}
	authReq := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(authRequest),
	}
	authParams := &openapiutil.Params{
		Action:      dara.String("AuthorizeFileUpload"),
		Version:     dara.String("2019-12-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	authResponse := map[string]interface{}{}
	fileObj := &dara.FileField{}
	ossHeader := map[string]interface{}{}
	tmpBody := map[string]interface{}{}
	useAccelerate := false
	authResponseBody := make(map[string]*string)
	queryContentReq := &QueryContentRequest{}
	openapiutil.Convert(request, queryContentReq)
	if !dara.IsNil(request.FileUrlObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.FileUrlObject,
			ContentType: dara.String(""),
		}
		ossHeader = map[string]interface{}{
			"host":                  dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(openapiutil.GetEndpoint(authResponseBody["Endpoint"], dara.Bool(useAccelerate), client.EndpointType)),
			"OSSAccessKeyId":        dara.StringValue(authResponseBody["AccessKeyId"]),
			"policy":                dara.StringValue(authResponseBody["EncodedPolicy"]),
			"Signature":             dara.StringValue(authResponseBody["Signature"]),
			"key":                   dara.StringValue(authResponseBody["ObjectKey"]),
			"file":                  fileObj,
			"success_action_status": "201",
		}
		_, _err = client._postOSSObject(authResponseBody["Bucket"], ossHeader)
		if _err != nil {
			return _result, _err
		}
		queryContentReq.FileUrl = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	queryContentResp, _err := client.QueryContentWithOptions(queryContentReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = queryContentResp
	return _result, _err
}

// Summary:
//
// Rebalances an AnalyticDB for PostgreSQL instance.
//
// @param request - RebalanceDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebalanceDBInstanceResponse
func (client *Client) RebalanceDBInstanceWithOptions(request *RebalanceDBInstanceRequest, runtime *dara.RuntimeOptions) (_result *RebalanceDBInstanceResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RebalanceDBInstance"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RebalanceDBInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Rebalances an AnalyticDB for PostgreSQL instance.
//
// @param request - RebalanceDBInstanceRequest
//
// @return RebalanceDBInstanceResponse
func (client *Client) RebalanceDBInstance(request *RebalanceDBInstanceRequest) (_result *RebalanceDBInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RebalanceDBInstanceResponse{}
	_body, _err := client.RebalanceDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases the public endpoint of an AnalyticDB for PostgreSQL instance.
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
	if !dara.IsNil(request.AddressType) {
		query["AddressType"] = request.AddressType
	}

	if !dara.IsNil(request.CurrentConnectionString) {
		query["CurrentConnectionString"] = request.CurrentConnectionString
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseInstancePublicConnection"),
		Version:     dara.String("2016-05-03"),
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
// Releases the public endpoint of an AnalyticDB for PostgreSQL instance.
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
// # Score and re-order documents using a model
//
// @param tmpReq - RerankRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RerankResponse
func (client *Client) RerankWithOptions(tmpReq *RerankRequest, runtime *dara.RuntimeOptions) (_result *RerankResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &RerankShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Documents) {
		request.DocumentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Documents, dara.String("Documents"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DocumentsShrink) {
		body["Documents"] = request.DocumentsShrink
	}

	if !dara.IsNil(request.MaxChunksPerDoc) {
		body["MaxChunksPerDoc"] = request.MaxChunksPerDoc
	}

	if !dara.IsNil(request.Model) {
		body["Model"] = request.Model
	}

	if !dara.IsNil(request.Query) {
		body["Query"] = request.Query
	}

	if !dara.IsNil(request.ReturnDocuments) {
		body["ReturnDocuments"] = request.ReturnDocuments
	}

	if !dara.IsNil(request.TopK) {
		body["TopK"] = request.TopK
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Rerank"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RerankResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Score and re-order documents using a model
//
// @param request - RerankRequest
//
// @return RerankResponse
func (client *Client) Rerank(request *RerankRequest) (_result *RerankResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RerankResponse{}
	_body, _err := client.RerankWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resets the password of a database account for an AnalyticDB for PostgreSQL instance.
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

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetAccountPassword"),
		Version:     dara.String("2016-05-03"),
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
// Resets the password of a database account for an AnalyticDB for PostgreSQL instance.
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
// Resets the IMV statistics.
//
// @param request - ResetIMVMonitorDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetIMVMonitorDataResponse
func (client *Client) ResetIMVMonitorDataWithOptions(request *ResetIMVMonitorDataRequest, runtime *dara.RuntimeOptions) (_result *ResetIMVMonitorDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Database) {
		query["Database"] = request.Database
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetIMVMonitorData"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetIMVMonitorDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the IMV statistics.
//
// @param request - ResetIMVMonitorDataRequest
//
// @return ResetIMVMonitorDataResponse
func (client *Client) ResetIMVMonitorData(request *ResetIMVMonitorDataRequest) (_result *ResetIMVMonitorDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetIMVMonitorDataResponse{}
	_body, _err := client.ResetIMVMonitorDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重置supabase数据库密码
//
// @param request - ResetSupabaseProjectPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetSupabaseProjectPasswordResponse
func (client *Client) ResetSupabaseProjectPasswordWithOptions(request *ResetSupabaseProjectPasswordRequest, runtime *dara.RuntimeOptions) (_result *ResetSupabaseProjectPasswordResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountPassword) {
		query["AccountPassword"] = request.AccountPassword
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
		Action:      dara.String("ResetSupabaseProjectPassword"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetSupabaseProjectPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重置supabase数据库密码
//
// @param request - ResetSupabaseProjectPasswordRequest
//
// @return ResetSupabaseProjectPasswordResponse
func (client *Client) ResetSupabaseProjectPassword(request *ResetSupabaseProjectPasswordRequest) (_result *ResetSupabaseProjectPasswordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetSupabaseProjectPasswordResponse{}
	_body, _err := client.ResetSupabaseProjectPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restarts an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// A restart takes about 3 to 30 minutes. During the restart, services are unavailable. We recommend that you restart the instance during off-peak hours. After the instance is restarted and enters the running state, you can access the instance.
//
// ## Limit
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered and may affect your business. We recommend that you take note of the limit when you call this operation.
//
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
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartDBInstance"),
		Version:     dara.String("2016-05-03"),
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
// Restarts an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// A restart takes about 3 to 30 minutes. During the restart, services are unavailable. We recommend that you restart the instance during off-peak hours. After the instance is restarted and enters the running state, you can access the instance.
//
// ## Limit
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered and may affect your business. We recommend that you take note of the limit when you call this operation.
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
// Resumes data redistribution.
//
// @param request - ResumeDataRedistributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeDataRedistributeResponse
func (client *Client) ResumeDataRedistributeWithOptions(request *ResumeDataRedistributeRequest, runtime *dara.RuntimeOptions) (_result *ResumeDataRedistributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumeDataRedistribute"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResumeDataRedistributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resumes data redistribution.
//
// @param request - ResumeDataRedistributeRequest
//
// @return ResumeDataRedistributeResponse
func (client *Client) ResumeDataRedistribute(request *ResumeDataRedistributeRequest) (_result *ResumeDataRedistributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResumeDataRedistributeResponse{}
	_body, _err := client.ResumeDataRedistributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resumes an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to resume an AnalyticDB for PostgreSQL instance that is in the **Paused*	- state.
//
// This operation is available only for AnalyticDB for PostgreSQL instances in Serverless mode that run V1.0.2.1 or later. For more information about how to view and update the minor engine version of an instance, see [View the minor engine version](https://help.aliyun.com/document_detail/277424.html) and [Update the minor engine version](https://help.aliyun.com/document_detail/139271.html).
//
// >  Before you call this operation, make sure that you are familiar with the billing methods and pricing of AnalyticDB for PostgreSQL instances. For more information, see [Billing methods](https://help.aliyun.com/document_detail/35406.html) and [AnalyticDB for PostgreSQL pricing](https://www.alibabacloud.com/zh/product/hybriddb-postgresql/pricing).
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ResumeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeInstanceResponse
func (client *Client) ResumeInstanceWithOptions(request *ResumeInstanceRequest, runtime *dara.RuntimeOptions) (_result *ResumeInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumeInstance"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Resumes an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to resume an AnalyticDB for PostgreSQL instance that is in the **Paused*	- state.
//
// This operation is available only for AnalyticDB for PostgreSQL instances in Serverless mode that run V1.0.2.1 or later. For more information about how to view and update the minor engine version of an instance, see [View the minor engine version](https://help.aliyun.com/document_detail/277424.html) and [Update the minor engine version](https://help.aliyun.com/document_detail/139271.html).
//
// >  Before you call this operation, make sure that you are familiar with the billing methods and pricing of AnalyticDB for PostgreSQL instances. For more information, see [Billing methods](https://help.aliyun.com/document_detail/35406.html) and [AnalyticDB for PostgreSQL pricing](https://www.alibabacloud.com/zh/product/hybriddb-postgresql/pricing).
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ResumeInstanceRequest
//
// @return ResumeInstanceResponse
func (client *Client) ResumeInstance(request *ResumeInstanceRequest) (_result *ResumeInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResumeInstanceResponse{}
	_body, _err := client.ResumeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables a plan for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to enable or disable a specified plan. The plan management feature is supported only for AnalyticDB for PostgreSQL instances in Serverless mode.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - SetDBInstancePlanStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDBInstancePlanStatusResponse
func (client *Client) SetDBInstancePlanStatusWithOptions(request *SetDBInstancePlanStatusRequest, runtime *dara.RuntimeOptions) (_result *SetDBInstancePlanStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PlanId) {
		query["PlanId"] = request.PlanId
	}

	if !dara.IsNil(request.PlanStatus) {
		query["PlanStatus"] = request.PlanStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDBInstancePlanStatus"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDBInstancePlanStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables a plan for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to enable or disable a specified plan. The plan management feature is supported only for AnalyticDB for PostgreSQL instances in Serverless mode.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - SetDBInstancePlanStatusRequest
//
// @return SetDBInstancePlanStatusResponse
func (client *Client) SetDBInstancePlanStatus(request *SetDBInstancePlanStatusRequest) (_result *SetDBInstancePlanStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDBInstancePlanStatusResponse{}
	_body, _err := client.SetDBInstancePlanStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables data sharing for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// This operation is called to enable or disable data sharing for an AnalyticDB for PostgreSQL instance in Serverless mode.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param tmpReq - SetDataShareInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDataShareInstanceResponse
func (client *Client) SetDataShareInstanceWithOptions(tmpReq *SetDataShareInstanceRequest, runtime *dara.RuntimeOptions) (_result *SetDataShareInstanceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SetDataShareInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceList) {
		request.InstanceListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceList, dara.String("InstanceList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceListShrink) {
		query["InstanceList"] = request.InstanceListShrink
	}

	if !dara.IsNil(request.OperationType) {
		query["OperationType"] = request.OperationType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDataShareInstance"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDataShareInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables data sharing for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// This operation is called to enable or disable data sharing for an AnalyticDB for PostgreSQL instance in Serverless mode.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - SetDataShareInstanceRequest
//
// @return SetDataShareInstanceResponse
func (client *Client) SetDataShareInstance(request *SetDataShareInstanceRequest) (_result *SetDataShareInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDataShareInstanceResponse{}
	_body, _err := client.SetDataShareInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Switches between the internal and public endpoints of an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// This operation is not supported for AnalyticDB for PostgreSQL instances in elastic storage mode or Serverless mode.
//
// @param request - SwitchDBInstanceNetTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchDBInstanceNetTypeResponse
func (client *Client) SwitchDBInstanceNetTypeWithOptions(request *SwitchDBInstanceNetTypeRequest, runtime *dara.RuntimeOptions) (_result *SwitchDBInstanceNetTypeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionStringPrefix) {
		query["ConnectionStringPrefix"] = request.ConnectionStringPrefix
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchDBInstanceNetType"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchDBInstanceNetTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Switches between the internal and public endpoints of an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// This operation is not supported for AnalyticDB for PostgreSQL instances in elastic storage mode or Serverless mode.
//
// @param request - SwitchDBInstanceNetTypeRequest
//
// @return SwitchDBInstanceNetTypeResponse
func (client *Client) SwitchDBInstanceNetType(request *SwitchDBInstanceNetTypeRequest) (_result *SwitchDBInstanceNetTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SwitchDBInstanceNetTypeResponse{}
	_body, _err := client.SwitchDBInstanceNetTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates and adds tags to AnalyticDB for PostgreSQL instances.
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
		Version:     dara.String("2016-05-03"),
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
// Creates and adds tags to AnalyticDB for PostgreSQL instances.
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
// 通过模型对文本文档进行向量化
//
// @param tmpReq - TextEmbeddingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TextEmbeddingResponse
func (client *Client) TextEmbeddingWithOptions(tmpReq *TextEmbeddingRequest, runtime *dara.RuntimeOptions) (_result *TextEmbeddingResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &TextEmbeddingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Input) {
		request.InputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Input, dara.String("Input"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Dimension) {
		query["Dimension"] = request.Dimension
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InputShrink) {
		body["Input"] = request.InputShrink
	}

	if !dara.IsNil(request.Model) {
		body["Model"] = request.Model
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TextEmbedding"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TextEmbeddingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过模型对文本文档进行向量化
//
// @param request - TextEmbeddingRequest
//
// @return TextEmbeddingResponse
func (client *Client) TextEmbedding(request *TextEmbeddingRequest) (_result *TextEmbeddingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TextEmbeddingResponse{}
	_body, _err := client.TextEmbeddingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unbinds database roles from a resource group.
//
// @param tmpReq - UnbindDBResourceGroupWithRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindDBResourceGroupWithRoleResponse
func (client *Client) UnbindDBResourceGroupWithRoleWithOptions(tmpReq *UnbindDBResourceGroupWithRoleRequest, runtime *dara.RuntimeOptions) (_result *UnbindDBResourceGroupWithRoleResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UnbindDBResourceGroupWithRoleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RoleList) {
		request.RoleListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoleList, dara.String("RoleList"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceGroupName) {
		query["ResourceGroupName"] = request.ResourceGroupName
	}

	if !dara.IsNil(request.RoleListShrink) {
		query["RoleList"] = request.RoleListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindDBResourceGroupWithRole"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindDBResourceGroupWithRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds database roles from a resource group.
//
// @param request - UnbindDBResourceGroupWithRoleRequest
//
// @return UnbindDBResourceGroupWithRoleResponse
func (client *Client) UnbindDBResourceGroupWithRole(request *UnbindDBResourceGroupWithRoleRequest) (_result *UnbindDBResourceGroupWithRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnbindDBResourceGroupWithRoleResponse{}
	_body, _err := client.UnbindDBResourceGroupWithRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases a sample dataset from an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to release a sample dataset from an AnalyticDB for PostgreSQL instance. You must have already loaded the sample dataset.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - UnloadSampleDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnloadSampleDataResponse
func (client *Client) UnloadSampleDataWithOptions(request *UnloadSampleDataRequest, runtime *dara.RuntimeOptions) (_result *UnloadSampleDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnloadSampleData"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnloadSampleDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases a sample dataset from an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to release a sample dataset from an AnalyticDB for PostgreSQL instance. You must have already loaded the sample dataset.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - UnloadSampleDataRequest
//
// @return UnloadSampleDataResponse
func (client *Client) UnloadSampleData(request *UnloadSampleDataRequest) (_result *UnloadSampleDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnloadSampleDataResponse{}
	_body, _err := client.UnloadSampleDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Remove resource tags
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
		Version:     dara.String("2016-05-03"),
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
// # Remove resource tags
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

// Summary:
//
// # Update Metadata of Collection Data
//
// @param tmpReq - UpdateCollectionDataMetadataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCollectionDataMetadataResponse
func (client *Client) UpdateCollectionDataMetadataWithOptions(tmpReq *UpdateCollectionDataMetadataRequest, runtime *dara.RuntimeOptions) (_result *UpdateCollectionDataMetadataResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateCollectionDataMetadataShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Metadata) {
		request.MetadataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Metadata, dara.String("Metadata"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Collection) {
		query["Collection"] = request.Collection
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.IdsShrink) {
		query["Ids"] = request.IdsShrink
	}

	if !dara.IsNil(request.MetadataShrink) {
		query["Metadata"] = request.MetadataShrink
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespacePassword) {
		query["NamespacePassword"] = request.NamespacePassword
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCollectionDataMetadata"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCollectionDataMetadataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update Metadata of Collection Data
//
// @param request - UpdateCollectionDataMetadataRequest
//
// @return UpdateCollectionDataMetadataResponse
func (client *Client) UpdateCollectionDataMetadata(request *UpdateCollectionDataMetadataRequest) (_result *UpdateCollectionDataMetadataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCollectionDataMetadataResponse{}
	_body, _err := client.UpdateCollectionDataMetadataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a plan for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to modify a plan for an AnalyticDB for PostgreSQL instance in Serverless mode. For example, you can modify a plan for periodically pausing and resuming an instance or scaling an instance.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - UpdateDBInstancePlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDBInstancePlanResponse
func (client *Client) UpdateDBInstancePlanWithOptions(request *UpdateDBInstancePlanRequest, runtime *dara.RuntimeOptions) (_result *UpdateDBInstancePlanResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PlanConfig) {
		query["PlanConfig"] = request.PlanConfig
	}

	if !dara.IsNil(request.PlanDesc) {
		query["PlanDesc"] = request.PlanDesc
	}

	if !dara.IsNil(request.PlanEndDate) {
		query["PlanEndDate"] = request.PlanEndDate
	}

	if !dara.IsNil(request.PlanId) {
		query["PlanId"] = request.PlanId
	}

	if !dara.IsNil(request.PlanName) {
		query["PlanName"] = request.PlanName
	}

	if !dara.IsNil(request.PlanStartDate) {
		query["PlanStartDate"] = request.PlanStartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDBInstancePlan"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDBInstancePlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a plan for an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// You can call this operation to modify a plan for an AnalyticDB for PostgreSQL instance in Serverless mode. For example, you can modify a plan for periodically pausing and resuming an instance or scaling an instance.
//
// ## Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - UpdateDBInstancePlanRequest
//
// @return UpdateDBInstancePlanResponse
func (client *Client) UpdateDBInstancePlan(request *UpdateDBInstancePlanRequest) (_result *UpdateDBInstancePlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDBInstancePlanResponse{}
	_body, _err := client.UpdateDBInstancePlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the configurations of an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// This operation is not available for instances in reserved storage mode.
//
// Before you call this operation, make sure that you are familiar with the billing of AnalyticDB for PostgreSQL. For more information, see [Billing methods](https://help.aliyun.com/document_detail/35406.html) and [AnalyticDB for PostgreSQL pricing](https://www.alibabacloud.com/zh/product/hybriddb-postgresql/pricing).
//
// @param request - UpgradeDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeDBInstanceResponse
func (client *Client) UpgradeDBInstanceWithOptions(request *UpgradeDBInstanceRequest, runtime *dara.RuntimeOptions) (_result *UpgradeDBInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CacheStorageSize) {
		query["CacheStorageSize"] = request.CacheStorageSize
	}

	if !dara.IsNil(request.DBInstanceClass) {
		query["DBInstanceClass"] = request.DBInstanceClass
	}

	if !dara.IsNil(request.DBInstanceGroupCount) {
		query["DBInstanceGroupCount"] = request.DBInstanceGroupCount
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.InstanceSpec) {
		query["InstanceSpec"] = request.InstanceSpec
	}

	if !dara.IsNil(request.MasterNodeNum) {
		query["MasterNodeNum"] = request.MasterNodeNum
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SegDiskPerformanceLevel) {
		query["SegDiskPerformanceLevel"] = request.SegDiskPerformanceLevel
	}

	if !dara.IsNil(request.SegNodeNum) {
		query["SegNodeNum"] = request.SegNodeNum
	}

	if !dara.IsNil(request.SegStorageType) {
		query["SegStorageType"] = request.SegStorageType
	}

	if !dara.IsNil(request.ServerlessResource) {
		query["ServerlessResource"] = request.ServerlessResource
	}

	if !dara.IsNil(request.StorageSize) {
		query["StorageSize"] = request.StorageSize
	}

	if !dara.IsNil(request.UpgradeType) {
		query["UpgradeType"] = request.UpgradeType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeDBInstance"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeDBInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the configurations of an AnalyticDB for PostgreSQL instance.
//
// Description:
//
// This operation is not available for instances in reserved storage mode.
//
// Before you call this operation, make sure that you are familiar with the billing of AnalyticDB for PostgreSQL. For more information, see [Billing methods](https://help.aliyun.com/document_detail/35406.html) and [AnalyticDB for PostgreSQL pricing](https://www.alibabacloud.com/zh/product/hybriddb-postgresql/pricing).
//
// @param request - UpgradeDBInstanceRequest
//
// @return UpgradeDBInstanceResponse
func (client *Client) UpgradeDBInstance(request *UpgradeDBInstanceRequest) (_result *UpgradeDBInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpgradeDBInstanceResponse{}
	_body, _err := client.UpgradeDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Upgrades the minor version of an AnalyticDB for PostgreSQL instance.
//
// @param request - UpgradeDBVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeDBVersionResponse
func (client *Client) UpgradeDBVersionWithOptions(request *UpgradeDBVersionRequest, runtime *dara.RuntimeOptions) (_result *UpgradeDBVersionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.MajorVersion) {
		query["MajorVersion"] = request.MajorVersion
	}

	if !dara.IsNil(request.MinorVersion) {
		query["MinorVersion"] = request.MinorVersion
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeDBVersion"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeDBVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades the minor version of an AnalyticDB for PostgreSQL instance.
//
// @param request - UpgradeDBVersionRequest
//
// @return UpgradeDBVersionResponse
func (client *Client) UpgradeDBVersion(request *UpgradeDBVersionRequest) (_result *UpgradeDBVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpgradeDBVersionResponse{}
	_body, _err := client.UpgradeDBVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates extensions.
//
// @param request - UpgradeExtensionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeExtensionsResponse
func (client *Client) UpgradeExtensionsWithOptions(request *UpgradeExtensionsRequest, runtime *dara.RuntimeOptions) (_result *UpgradeExtensionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DatabaseName) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.Extensions) {
		query["Extensions"] = request.Extensions
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeExtensions"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeExtensionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates extensions.
//
// @param request - UpgradeExtensionsRequest
//
// @return UpgradeExtensionsResponse
func (client *Client) UpgradeExtensions(request *UpgradeExtensionsRequest) (_result *UpgradeExtensionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpgradeExtensionsResponse{}
	_body, _err := client.UpgradeExtensionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Asynchronous Document Upload
//
// Description:
//
// The server loads and chunks a document based on the file extension, performs vectorization by using the embedding model that is specified when you call the CreateDocumentCollection operation, and then writes the document to the specified document collection. This operation supports multi-modal embedding for various formats of text and images.
//
// Related operations:
//
//   - You can call the GetUploadDocumentJob operation to query the progress and result of a document upload job.
//
//   - You can call the CancelUploadDocumentJob operation to cancel a document upload job.
//
// >
//
//   - After a document upload request is submitted, the request is queued for processing. Up to 20 documents in the Pending and Running states can be processed within a Resource Access Management (RAM) user or Alibaba Cloud account.
//
//   - A text document can be split into up to 100,000 chunks.
//
//   - If a document collection uses the OnePeace model, each RAM user or Alibaba Cloud account can upload and query up to 10,000 images.
//
// @param tmpReq - UploadDocumentAsyncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadDocumentAsyncResponse
func (client *Client) UploadDocumentAsyncWithOptions(tmpReq *UploadDocumentAsyncRequest, runtime *dara.RuntimeOptions) (_result *UploadDocumentAsyncResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UploadDocumentAsyncShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Metadata) {
		request.MetadataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Metadata, dara.String("Metadata"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Separators) {
		request.SeparatorsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Separators, dara.String("Separators"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ChunkOverlap) {
		body["ChunkOverlap"] = request.ChunkOverlap
	}

	if !dara.IsNil(request.ChunkSize) {
		body["ChunkSize"] = request.ChunkSize
	}

	if !dara.IsNil(request.Collection) {
		body["Collection"] = request.Collection
	}

	if !dara.IsNil(request.DocumentLoaderName) {
		body["DocumentLoaderName"] = request.DocumentLoaderName
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileName) {
		body["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileUrl) {
		body["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.MetadataShrink) {
		body["Metadata"] = request.MetadataShrink
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespacePassword) {
		body["NamespacePassword"] = request.NamespacePassword
	}

	if !dara.IsNil(request.SeparatorsShrink) {
		body["Separators"] = request.SeparatorsShrink
	}

	if !dara.IsNil(request.SplitterModel) {
		body["SplitterModel"] = request.SplitterModel
	}

	if !dara.IsNil(request.TextSplitterName) {
		body["TextSplitterName"] = request.TextSplitterName
	}

	if !dara.IsNil(request.VlEnhance) {
		body["VlEnhance"] = request.VlEnhance
	}

	if !dara.IsNil(request.ZhTitleEnhance) {
		body["ZhTitleEnhance"] = request.ZhTitleEnhance
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadDocumentAsync"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadDocumentAsyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Asynchronous Document Upload
//
// Description:
//
// The server loads and chunks a document based on the file extension, performs vectorization by using the embedding model that is specified when you call the CreateDocumentCollection operation, and then writes the document to the specified document collection. This operation supports multi-modal embedding for various formats of text and images.
//
// Related operations:
//
//   - You can call the GetUploadDocumentJob operation to query the progress and result of a document upload job.
//
//   - You can call the CancelUploadDocumentJob operation to cancel a document upload job.
//
// >
//
//   - After a document upload request is submitted, the request is queued for processing. Up to 20 documents in the Pending and Running states can be processed within a Resource Access Management (RAM) user or Alibaba Cloud account.
//
//   - A text document can be split into up to 100,000 chunks.
//
//   - If a document collection uses the OnePeace model, each RAM user or Alibaba Cloud account can upload and query up to 10,000 images.
//
// @param request - UploadDocumentAsyncRequest
//
// @return UploadDocumentAsyncResponse
func (client *Client) UploadDocumentAsync(request *UploadDocumentAsyncRequest) (_result *UploadDocumentAsyncResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadDocumentAsyncResponse{}
	_body, _err := client.UploadDocumentAsyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadDocumentAsyncAdvance(request *UploadDocumentAsyncAdvanceRequest, runtime *dara.RuntimeOptions) (_result *UploadDocumentAsyncResponse, _err error) {
	// Step 0: init client
	if dara.IsNil(client.Credential) {
		_err = &openapi.ClientError{
			Code:    dara.String("InvalidCredentials"),
			Message: dara.String("Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details."),
		}
		return _result, _err
	}

	credentialModel, _err := client.Credential.GetCredential()
	if _err != nil {
		return _result, _err
	}

	accessKeyId := dara.StringValue(credentialModel.AccessKeyId)
	accessKeySecret := dara.StringValue(credentialModel.AccessKeySecret)
	securityToken := dara.StringValue(credentialModel.SecurityToken)
	credentialType := dara.StringValue(credentialModel.Type)
	openPlatformEndpoint := dara.StringValue(client.OpenPlatformEndpoint)
	if dara.IsNil(dara.String(openPlatformEndpoint)) || openPlatformEndpoint == "" {
		openPlatformEndpoint = "openplatform.aliyuncs.com"
	}

	if dara.IsNil(dara.String(credentialType)) {
		credentialType = "access_key"
	}

	authConfig := &openapiutil.Config{
		AccessKeyId:     dara.String(accessKeyId),
		AccessKeySecret: dara.String(accessKeySecret),
		SecurityToken:   dara.String(securityToken),
		Type:            dara.String(credentialType),
		Endpoint:        dara.String(openPlatformEndpoint),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openapi.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := map[string]*string{
		"Product":  dara.String("gpdb"),
		"RegionId": client.RegionId,
	}
	authReq := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(authRequest),
	}
	authParams := &openapiutil.Params{
		Action:      dara.String("AuthorizeFileUpload"),
		Version:     dara.String("2019-12-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	authResponse := map[string]interface{}{}
	fileObj := &dara.FileField{}
	ossHeader := map[string]interface{}{}
	tmpBody := map[string]interface{}{}
	useAccelerate := false
	authResponseBody := make(map[string]*string)
	uploadDocumentAsyncReq := &UploadDocumentAsyncRequest{}
	openapiutil.Convert(request, uploadDocumentAsyncReq)
	if !dara.IsNil(request.FileUrlObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.FileUrlObject,
			ContentType: dara.String(""),
		}
		ossHeader = map[string]interface{}{
			"host":                  dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(openapiutil.GetEndpoint(authResponseBody["Endpoint"], dara.Bool(useAccelerate), client.EndpointType)),
			"OSSAccessKeyId":        dara.StringValue(authResponseBody["AccessKeyId"]),
			"policy":                dara.StringValue(authResponseBody["EncodedPolicy"]),
			"Signature":             dara.StringValue(authResponseBody["Signature"]),
			"key":                   dara.StringValue(authResponseBody["ObjectKey"]),
			"file":                  fileObj,
			"success_action_status": "201",
		}
		_, _err = client._postOSSObject(authResponseBody["Bucket"], ossHeader)
		if _err != nil {
			return _result, _err
		}
		uploadDocumentAsyncReq.FileUrl = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	uploadDocumentAsyncResp, _err := client.UploadDocumentAsyncWithOptions(uploadDocumentAsyncReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = uploadDocumentAsyncResp
	return _result, _err
}

// Summary:
//
// # Upload split text
//
// Description:
//
// The vectorization algorithm for the document is specified by the CreateDocumentCollection API.
//
// @param tmpReq - UpsertChunksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpsertChunksResponse
func (client *Client) UpsertChunksWithOptions(tmpReq *UpsertChunksRequest, runtime *dara.RuntimeOptions) (_result *UpsertChunksResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpsertChunksShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TextChunks) {
		request.TextChunksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TextChunks, dara.String("TextChunks"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AllowInsertWithFilter) {
		query["AllowInsertWithFilter"] = request.AllowInsertWithFilter
	}

	if !dara.IsNil(request.Collection) {
		query["Collection"] = request.Collection
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespacePassword) {
		query["NamespacePassword"] = request.NamespacePassword
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ShouldReplaceFile) {
		query["ShouldReplaceFile"] = request.ShouldReplaceFile
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TextChunksShrink) {
		body["TextChunks"] = request.TextChunksShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpsertChunks"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpsertChunksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Upload split text
//
// Description:
//
// The vectorization algorithm for the document is specified by the CreateDocumentCollection API.
//
// @param request - UpsertChunksRequest
//
// @return UpsertChunksResponse
func (client *Client) UpsertChunks(request *UpsertChunksRequest) (_result *UpsertChunksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpsertChunksResponse{}
	_body, _err := client.UpsertChunksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uploads vector data to a vector collection.
//
// @param tmpReq - UpsertCollectionDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpsertCollectionDataResponse
func (client *Client) UpsertCollectionDataWithOptions(tmpReq *UpsertCollectionDataRequest, runtime *dara.RuntimeOptions) (_result *UpsertCollectionDataResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpsertCollectionDataShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Rows) {
		request.RowsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Rows, dara.String("Rows"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Collection) {
		query["Collection"] = request.Collection
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespacePassword) {
		query["NamespacePassword"] = request.NamespacePassword
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RowsShrink) {
		body["Rows"] = request.RowsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpsertCollectionData"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpsertCollectionDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads vector data to a vector collection.
//
// @param request - UpsertCollectionDataRequest
//
// @return UpsertCollectionDataResponse
func (client *Client) UpsertCollectionData(request *UpsertCollectionDataRequest) (_result *UpsertCollectionDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpsertCollectionDataResponse{}
	_body, _err := client.UpsertCollectionDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uploads vector data in an asynchronous manner by using an on-premises file or a password-free Internet-accessible file URL. The vector data can be up to 200 MB in size.
//
// Description:
//
// This operation is the asynchronous operation of `UpsertCollectionData`. The `UpsertCollectionData` operation supports up to 10 MB of data, and this operation supports up to 200 MB of data.
//
// >  Related operations:
//
//   - You can call the GetUpsertCollectionDataJob operation to query the progress and result of an upload job.
//
//   - You can call the CancelUpsertCollectionDataJob operation to cancel an upload job.
//
// > You can upload data for the same collection only in a serial manner.
//
// @param request - UpsertCollectionDataAsyncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpsertCollectionDataAsyncResponse
func (client *Client) UpsertCollectionDataAsyncWithOptions(request *UpsertCollectionDataAsyncRequest, runtime *dara.RuntimeOptions) (_result *UpsertCollectionDataAsyncResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Collection) {
		body["Collection"] = request.Collection
	}

	if !dara.IsNil(request.FileUrl) {
		body["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespacePassword) {
		body["NamespacePassword"] = request.NamespacePassword
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpsertCollectionDataAsync"),
		Version:     dara.String("2016-05-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpsertCollectionDataAsyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads vector data in an asynchronous manner by using an on-premises file or a password-free Internet-accessible file URL. The vector data can be up to 200 MB in size.
//
// Description:
//
// This operation is the asynchronous operation of `UpsertCollectionData`. The `UpsertCollectionData` operation supports up to 10 MB of data, and this operation supports up to 200 MB of data.
//
// >  Related operations:
//
//   - You can call the GetUpsertCollectionDataJob operation to query the progress and result of an upload job.
//
//   - You can call the CancelUpsertCollectionDataJob operation to cancel an upload job.
//
// > You can upload data for the same collection only in a serial manner.
//
// @param request - UpsertCollectionDataAsyncRequest
//
// @return UpsertCollectionDataAsyncResponse
func (client *Client) UpsertCollectionDataAsync(request *UpsertCollectionDataAsyncRequest) (_result *UpsertCollectionDataAsyncResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpsertCollectionDataAsyncResponse{}
	_body, _err := client.UpsertCollectionDataAsyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpsertCollectionDataAsyncAdvance(request *UpsertCollectionDataAsyncAdvanceRequest, runtime *dara.RuntimeOptions) (_result *UpsertCollectionDataAsyncResponse, _err error) {
	// Step 0: init client
	if dara.IsNil(client.Credential) {
		_err = &openapi.ClientError{
			Code:    dara.String("InvalidCredentials"),
			Message: dara.String("Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details."),
		}
		return _result, _err
	}

	credentialModel, _err := client.Credential.GetCredential()
	if _err != nil {
		return _result, _err
	}

	accessKeyId := dara.StringValue(credentialModel.AccessKeyId)
	accessKeySecret := dara.StringValue(credentialModel.AccessKeySecret)
	securityToken := dara.StringValue(credentialModel.SecurityToken)
	credentialType := dara.StringValue(credentialModel.Type)
	openPlatformEndpoint := dara.StringValue(client.OpenPlatformEndpoint)
	if dara.IsNil(dara.String(openPlatformEndpoint)) || openPlatformEndpoint == "" {
		openPlatformEndpoint = "openplatform.aliyuncs.com"
	}

	if dara.IsNil(dara.String(credentialType)) {
		credentialType = "access_key"
	}

	authConfig := &openapiutil.Config{
		AccessKeyId:     dara.String(accessKeyId),
		AccessKeySecret: dara.String(accessKeySecret),
		SecurityToken:   dara.String(securityToken),
		Type:            dara.String(credentialType),
		Endpoint:        dara.String(openPlatformEndpoint),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openapi.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := map[string]*string{
		"Product":  dara.String("gpdb"),
		"RegionId": client.RegionId,
	}
	authReq := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(authRequest),
	}
	authParams := &openapiutil.Params{
		Action:      dara.String("AuthorizeFileUpload"),
		Version:     dara.String("2019-12-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	authResponse := map[string]interface{}{}
	fileObj := &dara.FileField{}
	ossHeader := map[string]interface{}{}
	tmpBody := map[string]interface{}{}
	useAccelerate := false
	authResponseBody := make(map[string]*string)
	upsertCollectionDataAsyncReq := &UpsertCollectionDataAsyncRequest{}
	openapiutil.Convert(request, upsertCollectionDataAsyncReq)
	if !dara.IsNil(request.FileUrlObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.FileUrlObject,
			ContentType: dara.String(""),
		}
		ossHeader = map[string]interface{}{
			"host":                  dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(openapiutil.GetEndpoint(authResponseBody["Endpoint"], dara.Bool(useAccelerate), client.EndpointType)),
			"OSSAccessKeyId":        dara.StringValue(authResponseBody["AccessKeyId"]),
			"policy":                dara.StringValue(authResponseBody["EncodedPolicy"]),
			"Signature":             dara.StringValue(authResponseBody["Signature"]),
			"key":                   dara.StringValue(authResponseBody["ObjectKey"]),
			"file":                  fileObj,
			"success_action_status": "201",
		}
		_, _err = client._postOSSObject(authResponseBody["Bucket"], ossHeader)
		if _err != nil {
			return _result, _err
		}
		upsertCollectionDataAsyncReq.FileUrl = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	upsertCollectionDataAsyncResp, _err := client.UpsertCollectionDataAsyncWithOptions(upsertCollectionDataAsyncReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = upsertCollectionDataAsyncResp
	return _result, _err
}

func _postOSSObject_opResponse(response_ *dara.Response) (_result map[string]interface{}, _err error) {
	var respMap map[string]interface{}
	bodyStr, _err := dara.ReadAsString(response_.Body)
	if _err != nil {
		return _result, _err
	}

	if (dara.IntValue(response_.StatusCode) >= 400) && (dara.IntValue(response_.StatusCode) < 600) {
		respMap = dara.ParseXml(bodyStr, nil)
		err := dara.ToMap(respMap["Error"])
		_err = &openapi.ClientError{
			Code:    dara.String(dara.ToString(err["Code"])),
			Message: dara.String(dara.ToString(err["Message"])),
			Data: map[string]interface{}{
				"httpCode":  dara.IntValue(response_.StatusCode),
				"requestId": dara.ToString(err["RequestId"]),
				"hostId":    dara.ToString(err["HostId"]),
			},
		}
		return _result, _err
	}

	respMap = dara.ParseXml(bodyStr, nil)
	_result = make(map[string]interface{})
	_err = dara.Convert(dara.ToMap(respMap), &_result)

	return _result, _err
}
