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
		"cn-qingdao":                  dara.String("adb.aliyuncs.com"),
		"cn-beijing":                  dara.String("adb.aliyuncs.com"),
		"cn-hangzhou":                 dara.String("adb.aliyuncs.com"),
		"cn-shanghai":                 dara.String("adb.aliyuncs.com"),
		"cn-shenzhen":                 dara.String("adb.aliyuncs.com"),
		"cn-hongkong":                 dara.String("adb.aliyuncs.com"),
		"ap-southeast-1":              dara.String("adb.aliyuncs.com"),
		"us-west-1":                   dara.String("adb.aliyuncs.com"),
		"us-east-1":                   dara.String("adb.aliyuncs.com"),
		"cn-hangzhou-finance":         dara.String("adb.aliyuncs.com"),
		"cn-north-2-gov-1":            dara.String("adb.aliyuncs.com"),
		"ap-northeast-2-pop":          dara.String("adb.ap-northeast-1.aliyuncs.com"),
		"cn-beijing-finance-1":        dara.String("adb.aliyuncs.com"),
		"cn-beijing-finance-pop":      dara.String("adb.aliyuncs.com"),
		"cn-beijing-gov-1":            dara.String("adb.aliyuncs.com"),
		"cn-beijing-nu16-b01":         dara.String("adb.aliyuncs.com"),
		"cn-edge-1":                   dara.String("adb.aliyuncs.com"),
		"cn-fujian":                   dara.String("adb.aliyuncs.com"),
		"cn-haidian-cm12-c01":         dara.String("adb.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          dara.String("adb.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": dara.String("adb.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": dara.String("adb.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": dara.String("adb.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": dara.String("adb.aliyuncs.com"),
		"cn-hangzhou-test-306":        dara.String("adb.aliyuncs.com"),
		"cn-hongkong-finance-pop":     dara.String("adb.aliyuncs.com"),
		"cn-qingdao-nebula":           dara.String("adb.aliyuncs.com"),
		"cn-shanghai-et15-b01":        dara.String("adb.aliyuncs.com"),
		"cn-shanghai-et2-b01":         dara.String("adb.aliyuncs.com"),
		"cn-shanghai-finance-1":       dara.String("adb.aliyuncs.com"),
		"cn-shanghai-inner":           dara.String("adb.aliyuncs.com"),
		"cn-shanghai-internal-test-1": dara.String("adb.aliyuncs.com"),
		"cn-shenzhen-finance-1":       dara.String("adb.aliyuncs.com"),
		"cn-shenzhen-inner":           dara.String("adb.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         dara.String("adb.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        dara.String("adb.aliyuncs.com"),
		"cn-wuhan":                    dara.String("adb.aliyuncs.com"),
		"cn-yushanfang":               dara.String("adb.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        dara.String("adb.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     dara.String("adb.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       dara.String("adb.aliyuncs.com"),
		"eu-west-1-oxs":               dara.String("adb.ap-northeast-1.aliyuncs.com"),
		"me-east-1":                   dara.String("adb.ap-northeast-1.aliyuncs.com"),
		"rus-west-1-pop":              dara.String("adb.ap-northeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("adb"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Applies for a public endpoint for an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - AllocateClusterPublicConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AllocateClusterPublicConnectionResponse
func (client *Client) AllocateClusterPublicConnectionWithOptions(request *AllocateClusterPublicConnectionRequest, runtime *dara.RuntimeOptions) (_result *AllocateClusterPublicConnectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionStringPrefix) {
		query["ConnectionStringPrefix"] = request.ConnectionStringPrefix
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AllocateClusterPublicConnection"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AllocateClusterPublicConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Applies for a public endpoint for an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - AllocateClusterPublicConnectionRequest
//
// @return AllocateClusterPublicConnectionResponse
func (client *Client) AllocateClusterPublicConnection(request *AllocateClusterPublicConnectionRequest) (_result *AllocateClusterPublicConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AllocateClusterPublicConnectionResponse{}
	_body, _err := client.AllocateClusterPublicConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Applies an optimization suggestion.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - ApplyAdviceByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyAdviceByIdResponse
func (client *Client) ApplyAdviceByIdWithOptions(request *ApplyAdviceByIdRequest, runtime *dara.RuntimeOptions) (_result *ApplyAdviceByIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdviceDate) {
		query["AdviceDate"] = request.AdviceDate
	}

	if !dara.IsNil(request.AdviceId) {
		query["AdviceId"] = request.AdviceId
	}

	if !dara.IsNil(request.ApplyType) {
		query["ApplyType"] = request.ApplyType
	}

	if !dara.IsNil(request.BuildImmediately) {
		query["BuildImmediately"] = request.BuildImmediately
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyAdviceById"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyAdviceByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Applies an optimization suggestion.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - ApplyAdviceByIdRequest
//
// @return ApplyAdviceByIdResponse
func (client *Client) ApplyAdviceById(request *ApplyAdviceByIdRequest) (_result *ApplyAdviceByIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ApplyAdviceByIdResponse{}
	_body, _err := client.ApplyAdviceByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Attaches an elastic network interface (ENI) to an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - AttachUserENIRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachUserENIResponse
func (client *Client) AttachUserENIWithOptions(request *AttachUserENIRequest, runtime *dara.RuntimeOptions) (_result *AttachUserENIResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachUserENI"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachUserENIResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attaches an elastic network interface (ENI) to an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - AttachUserENIRequest
//
// @return AttachUserENIResponse
func (client *Client) AttachUserENI(request *AttachUserENIRequest) (_result *AttachUserENIResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachUserENIResponse{}
	_body, _err := client.AttachUserENIWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Applies optimization suggestions.
//
// @param request - BatchApplyAdviceByIdListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchApplyAdviceByIdListResponse
func (client *Client) BatchApplyAdviceByIdListWithOptions(request *BatchApplyAdviceByIdListRequest, runtime *dara.RuntimeOptions) (_result *BatchApplyAdviceByIdListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdviceDate) {
		query["AdviceDate"] = request.AdviceDate
	}

	if !dara.IsNil(request.AdviceIdList) {
		query["AdviceIdList"] = request.AdviceIdList
	}

	if !dara.IsNil(request.ApplyType) {
		query["ApplyType"] = request.ApplyType
	}

	if !dara.IsNil(request.BuildImmediately) {
		query["BuildImmediately"] = request.BuildImmediately
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchApplyAdviceByIdList"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchApplyAdviceByIdListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Applies optimization suggestions.
//
// @param request - BatchApplyAdviceByIdListRequest
//
// @return BatchApplyAdviceByIdListResponse
func (client *Client) BatchApplyAdviceByIdList(request *BatchApplyAdviceByIdListRequest) (_result *BatchApplyAdviceByIdListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchApplyAdviceByIdListResponse{}
	_body, _err := client.BatchApplyAdviceByIdListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates a standard account of an AnalyticDB for MySQL cluster with a Resource Access Management (RAM) user.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - BindAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAccountResponse
func (client *Client) BindAccountWithOptions(request *BindAccountRequest, runtime *dara.RuntimeOptions) (_result *BindAccountResponse, _err error) {
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

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.RamUser) {
		query["RamUser"] = request.RamUser
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindAccount"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a standard account of an AnalyticDB for MySQL cluster with a Resource Access Management (RAM) user.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - BindAccountRequest
//
// @return BindAccountResponse
func (client *Client) BindAccount(request *BindAccountRequest) (_result *BindAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindAccountResponse{}
	_body, _err := client.BindAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates a resource group with a database account.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - BindDBResourceGroupWithUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindDBResourceGroupWithUserResponse
func (client *Client) BindDBResourceGroupWithUserWithOptions(request *BindDBResourceGroupWithUserRequest, runtime *dara.RuntimeOptions) (_result *BindDBResourceGroupWithUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.GroupUser) {
		query["GroupUser"] = request.GroupUser
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindDBResourceGroupWithUser"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindDBResourceGroupWithUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a resource group with a database account.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - BindDBResourceGroupWithUserRequest
//
// @return BindDBResourceGroupWithUserResponse
func (client *Client) BindDBResourceGroupWithUser(request *BindDBResourceGroupWithUserRequest) (_result *BindDBResourceGroupWithUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindDBResourceGroupWithUserResponse{}
	_body, _err := client.BindDBResourceGroupWithUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Terminates part of the code in a Spark job.
//
// @param request - CancelSparkReplStatementRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelSparkReplStatementResponse
func (client *Client) CancelSparkReplStatementWithOptions(request *CancelSparkReplStatementRequest, runtime *dara.RuntimeOptions) (_result *CancelSparkReplStatementResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.StatementId) {
		body["StatementId"] = request.StatementId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelSparkReplStatement"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelSparkReplStatementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Terminates part of the code in a Spark job.
//
// @param request - CancelSparkReplStatementRequest
//
// @return CancelSparkReplStatementResponse
func (client *Client) CancelSparkReplStatement(request *CancelSparkReplStatementRequest) (_result *CancelSparkReplStatementResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelSparkReplStatementResponse{}
	_body, _err := client.CancelSparkReplStatementWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels the execution of a Spark SQL statement.
//
// @param request - CancelSparkWarehouseBatchSQLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelSparkWarehouseBatchSQLResponse
func (client *Client) CancelSparkWarehouseBatchSQLWithOptions(request *CancelSparkWarehouseBatchSQLRequest, runtime *dara.RuntimeOptions) (_result *CancelSparkWarehouseBatchSQLResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Agency) {
		body["Agency"] = request.Agency
	}

	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.QueryId) {
		body["QueryId"] = request.QueryId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelSparkWarehouseBatchSQL"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelSparkWarehouseBatchSQLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels the execution of a Spark SQL statement.
//
// @param request - CancelSparkWarehouseBatchSQLRequest
//
// @return CancelSparkWarehouseBatchSQLResponse
func (client *Client) CancelSparkWarehouseBatchSQL(request *CancelSparkWarehouseBatchSQLRequest) (_result *CancelSparkWarehouseBatchSQLResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelSparkWarehouseBatchSQLResponse{}
	_body, _err := client.CancelSparkWarehouseBatchSQLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether a database account of an AnalyticDB for MySQL cluster is associated with a Resource Access Management (RAM) user.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - CheckBindRamUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckBindRamUserResponse
func (client *Client) CheckBindRamUserWithOptions(request *CheckBindRamUserRequest, runtime *dara.RuntimeOptions) (_result *CheckBindRamUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckBindRamUser"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckBindRamUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether a database account of an AnalyticDB for MySQL cluster is associated with a Resource Access Management (RAM) user.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - CheckBindRamUserRequest
//
// @return CheckBindRamUserResponse
func (client *Client) CheckBindRamUser(request *CheckBindRamUserRequest) (_result *CheckBindRamUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckBindRamUserResponse{}
	_body, _err := client.CheckBindRamUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - CheckSampleDataSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckSampleDataSetResponse
func (client *Client) CheckSampleDataSetWithOptions(request *CheckSampleDataSetRequest, runtime *dara.RuntimeOptions) (_result *CheckSampleDataSetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckSampleDataSet"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckSampleDataSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - CheckSampleDataSetRequest
//
// @return CheckSampleDataSetResponse
func (client *Client) CheckSampleDataSet(request *CheckSampleDataSetRequest) (_result *CheckSampleDataSetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckSampleDataSetResponse{}
	_body, _err := client.CheckSampleDataSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the export destination (SLS or OSS) at the instance level. The configuration is unique per instance and follows the "configure once, use multiple times" principle.
//
// @param tmpReq - ConfigureResultExportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigureResultExportResponse
func (client *Client) ConfigureResultExportWithOptions(tmpReq *ConfigureResultExportRequest, runtime *dara.RuntimeOptions) (_result *ConfigureResultExportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ConfigureResultExportShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OssInfo) {
		request.OssInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OssInfo, dara.String("OssInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SlsInfo) {
		request.SlsInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SlsInfo, dara.String("SlsInfo"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.ExportType) {
		body["ExportType"] = request.ExportType
	}

	if !dara.IsNil(request.OssInfoShrink) {
		body["OssInfo"] = request.OssInfoShrink
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SlsInfoShrink) {
		body["SlsInfo"] = request.SlsInfoShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigureResultExport"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigureResultExportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the export destination (SLS or OSS) at the instance level. The configuration is unique per instance and follows the "configure once, use multiple times" principle.
//
// @param request - ConfigureResultExportRequest
//
// @return ConfigureResultExportResponse
func (client *Client) ConfigureResultExport(request *ConfigureResultExportRequest) (_result *ConfigureResultExportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigureResultExportResponse{}
	_body, _err := client.ConfigureResultExportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an AnalyticDB Pipeline Service (APS) job.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - CreateAPSJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAPSJobResponse
func (client *Client) CreateAPSJobWithOptions(request *CreateAPSJobRequest, runtime *dara.RuntimeOptions) (_result *CreateAPSJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApsJobName) {
		body["ApsJobName"] = request.ApsJobName
	}

	if !dara.IsNil(request.DbList) {
		body["DbList"] = request.DbList
	}

	if !dara.IsNil(request.DestinationEndpointInstanceID) {
		body["DestinationEndpointInstanceID"] = request.DestinationEndpointInstanceID
	}

	if !dara.IsNil(request.DestinationEndpointPassword) {
		body["DestinationEndpointPassword"] = request.DestinationEndpointPassword
	}

	if !dara.IsNil(request.DestinationEndpointUserName) {
		body["DestinationEndpointUserName"] = request.DestinationEndpointUserName
	}

	if !dara.IsNil(request.PartitionList) {
		body["PartitionList"] = request.PartitionList
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SourceEndpointInstanceID) {
		body["SourceEndpointInstanceID"] = request.SourceEndpointInstanceID
	}

	if !dara.IsNil(request.SourceEndpointPassword) {
		body["SourceEndpointPassword"] = request.SourceEndpointPassword
	}

	if !dara.IsNil(request.SourceEndpointRegion) {
		body["SourceEndpointRegion"] = request.SourceEndpointRegion
	}

	if !dara.IsNil(request.SourceEndpointUserName) {
		body["SourceEndpointUserName"] = request.SourceEndpointUserName
	}

	if !dara.IsNil(request.TargetTableMode) {
		body["TargetTableMode"] = request.TargetTableMode
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAPSJob"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAPSJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an AnalyticDB Pipeline Service (APS) job.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - CreateAPSJobRequest
//
// @return CreateAPSJobResponse
func (client *Client) CreateAPSJob(request *CreateAPSJobRequest) (_result *CreateAPSJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAPSJobResponse{}
	_body, _err := client.CreateAPSJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a database account for an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
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

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAccount"),
		Version:     dara.String("2021-12-01"),
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
// Creates a database account for an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
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
// Creates an AnalyticDB Pipeline Service (APS) replication job.
//
// @param request - CreateApsCopyWorkloadRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApsCopyWorkloadResponse
func (client *Client) CreateApsCopyWorkloadWithOptions(request *CreateApsCopyWorkloadRequest, runtime *dara.RuntimeOptions) (_result *CreateApsCopyWorkloadResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DatasourceId) {
		body["DatasourceId"] = request.DatasourceId
	}

	if !dara.IsNil(request.DbName) {
		body["DbName"] = request.DbName
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TableName) {
		body["TableName"] = request.TableName
	}

	if !dara.IsNil(request.WorkloadId) {
		body["WorkloadId"] = request.WorkloadId
	}

	if !dara.IsNil(request.WorkloadType) {
		body["WorkloadType"] = request.WorkloadType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApsCopyWorkload"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApsCopyWorkloadResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an AnalyticDB Pipeline Service (APS) replication job.
//
// @param request - CreateApsCopyWorkloadRequest
//
// @return CreateApsCopyWorkloadResponse
func (client *Client) CreateApsCopyWorkload(request *CreateApsCopyWorkloadRequest) (_result *CreateApsCopyWorkloadResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateApsCopyWorkloadResponse{}
	_body, _err := client.CreateApsCopyWorkloadWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an AnalyticDB Pipeline Service (APS) data source.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param tmpReq - CreateApsDatasoureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApsDatasoureResponse
func (client *Client) CreateApsDatasoureWithOptions(tmpReq *CreateApsDatasoureRequest, runtime *dara.RuntimeOptions) (_result *CreateApsDatasoureResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateApsDatasoureShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DatabricksInfo) {
		request.DatabricksInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DatabricksInfo, dara.String("DatabricksInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.HiveInfo) {
		request.HiveInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HiveInfo, dara.String("HiveInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.KafkaInfo) {
		request.KafkaInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.KafkaInfo, dara.String("KafkaInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PolarDBMysqlInfo) {
		request.PolarDBMysqlInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PolarDBMysqlInfo, dara.String("PolarDBMysqlInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PolarDBXInfo) {
		request.PolarDBXInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PolarDBXInfo, dara.String("PolarDBXInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RdsMysqlInfo) {
		request.RdsMysqlInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RdsMysqlInfo, dara.String("RdsMysqlInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SlsInfo) {
		request.SlsInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SlsInfo, dara.String("SlsInfo"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DatabricksInfoShrink) {
		body["DatabricksInfo"] = request.DatabricksInfoShrink
	}

	if !dara.IsNil(request.DatasourceDescription) {
		body["DatasourceDescription"] = request.DatasourceDescription
	}

	if !dara.IsNil(request.DatasourceName) {
		body["DatasourceName"] = request.DatasourceName
	}

	if !dara.IsNil(request.DatasourceType) {
		body["DatasourceType"] = request.DatasourceType
	}

	if !dara.IsNil(request.HiveInfoShrink) {
		body["HiveInfo"] = request.HiveInfoShrink
	}

	if !dara.IsNil(request.KafkaInfoShrink) {
		body["KafkaInfo"] = request.KafkaInfoShrink
	}

	if !dara.IsNil(request.Mode) {
		body["Mode"] = request.Mode
	}

	if !dara.IsNil(request.PolarDBMysqlInfoShrink) {
		body["PolarDBMysqlInfo"] = request.PolarDBMysqlInfoShrink
	}

	if !dara.IsNil(request.PolarDBXInfoShrink) {
		body["PolarDBXInfo"] = request.PolarDBXInfoShrink
	}

	if !dara.IsNil(request.RdsMysqlInfoShrink) {
		body["RdsMysqlInfo"] = request.RdsMysqlInfoShrink
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SlsInfoShrink) {
		body["SlsInfo"] = request.SlsInfoShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApsDatasoure"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApsDatasoureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an AnalyticDB Pipeline Service (APS) data source.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - CreateApsDatasoureRequest
//
// @return CreateApsDatasoureResponse
func (client *Client) CreateApsDatasoure(request *CreateApsDatasoureRequest) (_result *CreateApsDatasoureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateApsDatasoureResponse{}
	_body, _err := client.CreateApsDatasoureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an AnalyticDB Pipeline Service (APS) job from a Hive data source.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - CreateApsHiveJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApsHiveJobResponse
func (client *Client) CreateApsHiveJobWithOptions(request *CreateApsHiveJobRequest, runtime *dara.RuntimeOptions) (_result *CreateApsHiveJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AdvancedConfig) {
		body["AdvancedConfig"] = request.AdvancedConfig
	}

	if !dara.IsNil(request.ConflictStrategy) {
		body["ConflictStrategy"] = request.ConflictStrategy
	}

	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DatasourceId) {
		body["DatasourceId"] = request.DatasourceId
	}

	if !dara.IsNil(request.FullComputeUnit) {
		body["FullComputeUnit"] = request.FullComputeUnit
	}

	if !dara.IsNil(request.OssLocation) {
		body["OssLocation"] = request.OssLocation
	}

	if !dara.IsNil(request.Parallelism) {
		body["Parallelism"] = request.Parallelism
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroup) {
		body["ResourceGroup"] = request.ResourceGroup
	}

	if !dara.IsNil(request.SyncAllowExpression) {
		body["SyncAllowExpression"] = request.SyncAllowExpression
	}

	if !dara.IsNil(request.SyncDenyExpression) {
		body["SyncDenyExpression"] = request.SyncDenyExpression
	}

	if !dara.IsNil(request.TargetType) {
		body["TargetType"] = request.TargetType
	}

	if !dara.IsNil(request.WorkloadName) {
		body["WorkloadName"] = request.WorkloadName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApsHiveJob"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApsHiveJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an AnalyticDB Pipeline Service (APS) job from a Hive data source.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - CreateApsHiveJobRequest
//
// @return CreateApsHiveJobResponse
func (client *Client) CreateApsHiveJob(request *CreateApsHiveJobRequest) (_result *CreateApsHiveJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateApsHiveJobResponse{}
	_body, _err := client.CreateApsHiveJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a data ingestion task to load data from an Apache Kafka topic into an AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
//
// @param tmpReq - CreateApsKafkaHudiJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApsKafkaHudiJobResponse
func (client *Client) CreateApsKafkaHudiJobWithOptions(tmpReq *CreateApsKafkaHudiJobRequest, runtime *dara.RuntimeOptions) (_result *CreateApsKafkaHudiJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateApsKafkaHudiJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Columns) {
		request.ColumnsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Columns, dara.String("Columns"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PartitionSpecs) {
		request.PartitionSpecsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PartitionSpecs, dara.String("PartitionSpecs"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AcrossRole) {
		body["AcrossRole"] = request.AcrossRole
	}

	if !dara.IsNil(request.AcrossUid) {
		body["AcrossUid"] = request.AcrossUid
	}

	if !dara.IsNil(request.AdvancedConfig) {
		body["AdvancedConfig"] = request.AdvancedConfig
	}

	if !dara.IsNil(request.ColumnsShrink) {
		body["Columns"] = request.ColumnsShrink
	}

	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DataFormatType) {
		body["DataFormatType"] = request.DataFormatType
	}

	if !dara.IsNil(request.DataOutputFormat) {
		body["DataOutputFormat"] = request.DataOutputFormat
	}

	if !dara.IsNil(request.DatasourceId) {
		body["DatasourceId"] = request.DatasourceId
	}

	if !dara.IsNil(request.DbName) {
		body["DbName"] = request.DbName
	}

	if !dara.IsNil(request.FullComputeUnit) {
		body["FullComputeUnit"] = request.FullComputeUnit
	}

	if !dara.IsNil(request.HudiAdvancedConfig) {
		body["HudiAdvancedConfig"] = request.HudiAdvancedConfig
	}

	if !dara.IsNil(request.IncrementalComputeUnit) {
		body["IncrementalComputeUnit"] = request.IncrementalComputeUnit
	}

	if !dara.IsNil(request.JsonParseLevel) {
		body["JsonParseLevel"] = request.JsonParseLevel
	}

	if !dara.IsNil(request.KafkaClusterId) {
		body["KafkaClusterId"] = request.KafkaClusterId
	}

	if !dara.IsNil(request.KafkaTopic) {
		body["KafkaTopic"] = request.KafkaTopic
	}

	if !dara.IsNil(request.LakehouseId) {
		body["LakehouseId"] = request.LakehouseId
	}

	if !dara.IsNil(request.MaxOffsetsPerTrigger) {
		body["MaxOffsetsPerTrigger"] = request.MaxOffsetsPerTrigger
	}

	if !dara.IsNil(request.OssLocation) {
		body["OssLocation"] = request.OssLocation
	}

	if !dara.IsNil(request.OutputFormat) {
		body["OutputFormat"] = request.OutputFormat
	}

	if !dara.IsNil(request.PartitionSpecsShrink) {
		body["PartitionSpecs"] = request.PartitionSpecsShrink
	}

	if !dara.IsNil(request.PrimaryKeyDefinition) {
		body["PrimaryKeyDefinition"] = request.PrimaryKeyDefinition
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroup) {
		body["ResourceGroup"] = request.ResourceGroup
	}

	if !dara.IsNil(request.SourceRegionId) {
		body["SourceRegionId"] = request.SourceRegionId
	}

	if !dara.IsNil(request.StartingOffsets) {
		body["StartingOffsets"] = request.StartingOffsets
	}

	if !dara.IsNil(request.TableName) {
		body["TableName"] = request.TableName
	}

	if !dara.IsNil(request.TargetGenerateRule) {
		body["TargetGenerateRule"] = request.TargetGenerateRule
	}

	if !dara.IsNil(request.TargetType) {
		body["TargetType"] = request.TargetType
	}

	if !dara.IsNil(request.WorkloadName) {
		body["WorkloadName"] = request.WorkloadName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApsKafkaHudiJob"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApsKafkaHudiJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data ingestion task to load data from an Apache Kafka topic into an AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
//
// @param request - CreateApsKafkaHudiJobRequest
//
// @return CreateApsKafkaHudiJobResponse
func (client *Client) CreateApsKafkaHudiJob(request *CreateApsKafkaHudiJobRequest) (_result *CreateApsKafkaHudiJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateApsKafkaHudiJobResponse{}
	_body, _err := client.CreateApsKafkaHudiJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an AnalyticDB Pipeline Service (APS) job from Simple Log Service (SLS) to an AnalyticDB for MySQL Data Warehouse Edition cluster.
//
// @param tmpReq - CreateApsSlsADBJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApsSlsADBJobResponse
func (client *Client) CreateApsSlsADBJobWithOptions(tmpReq *CreateApsSlsADBJobRequest, runtime *dara.RuntimeOptions) (_result *CreateApsSlsADBJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateApsSlsADBJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Columns) {
		request.ColumnsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Columns, dara.String("Columns"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PartitionSpecs) {
		request.PartitionSpecsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PartitionSpecs, dara.String("PartitionSpecs"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UnixTimestampConvert) {
		request.UnixTimestampConvertShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UnixTimestampConvert, dara.String("UnixTimestampConvert"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AcrossRole) {
		body["AcrossRole"] = request.AcrossRole
	}

	if !dara.IsNil(request.AcrossUid) {
		body["AcrossUid"] = request.AcrossUid
	}

	if !dara.IsNil(request.AdvancedConfig) {
		body["AdvancedConfig"] = request.AdvancedConfig
	}

	if !dara.IsNil(request.ColumnsShrink) {
		body["Columns"] = request.ColumnsShrink
	}

	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DatasourceId) {
		body["DatasourceId"] = request.DatasourceId
	}

	if !dara.IsNil(request.DbName) {
		body["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DirtyDataHandleMode) {
		body["DirtyDataHandleMode"] = request.DirtyDataHandleMode
	}

	if !dara.IsNil(request.DirtyDataProcessPattern) {
		body["DirtyDataProcessPattern"] = request.DirtyDataProcessPattern
	}

	if !dara.IsNil(request.ExactlyOnce) {
		body["ExactlyOnce"] = request.ExactlyOnce
	}

	if !dara.IsNil(request.FullComputeUnit) {
		body["FullComputeUnit"] = request.FullComputeUnit
	}

	if !dara.IsNil(request.HudiAdvancedConfig) {
		body["HudiAdvancedConfig"] = request.HudiAdvancedConfig
	}

	if !dara.IsNil(request.IncrementalComputeUnit) {
		body["IncrementalComputeUnit"] = request.IncrementalComputeUnit
	}

	if !dara.IsNil(request.LakehouseId) {
		body["LakehouseId"] = request.LakehouseId
	}

	if !dara.IsNil(request.MaxOffsetsPerTrigger) {
		body["MaxOffsetsPerTrigger"] = request.MaxOffsetsPerTrigger
	}

	if !dara.IsNil(request.OssLocation) {
		body["OssLocation"] = request.OssLocation
	}

	if !dara.IsNil(request.OutputFormat) {
		body["OutputFormat"] = request.OutputFormat
	}

	if !dara.IsNil(request.PartitionSpecsShrink) {
		body["PartitionSpecs"] = request.PartitionSpecsShrink
	}

	if !dara.IsNil(request.Password) {
		body["Password"] = request.Password
	}

	if !dara.IsNil(request.PrimaryKeyDefinition) {
		body["PrimaryKeyDefinition"] = request.PrimaryKeyDefinition
	}

	if !dara.IsNil(request.Project) {
		body["Project"] = request.Project
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroup) {
		body["ResourceGroup"] = request.ResourceGroup
	}

	if !dara.IsNil(request.SourceRegionId) {
		body["SourceRegionId"] = request.SourceRegionId
	}

	if !dara.IsNil(request.StartingOffsets) {
		body["StartingOffsets"] = request.StartingOffsets
	}

	if !dara.IsNil(request.Store) {
		body["Store"] = request.Store
	}

	if !dara.IsNil(request.TableName) {
		body["TableName"] = request.TableName
	}

	if !dara.IsNil(request.TargetGenerateRule) {
		body["TargetGenerateRule"] = request.TargetGenerateRule
	}

	if !dara.IsNil(request.TargetType) {
		body["TargetType"] = request.TargetType
	}

	if !dara.IsNil(request.UnixTimestampConvertShrink) {
		body["UnixTimestampConvert"] = request.UnixTimestampConvertShrink
	}

	if !dara.IsNil(request.UserName) {
		body["UserName"] = request.UserName
	}

	if !dara.IsNil(request.WorkloadName) {
		body["WorkloadName"] = request.WorkloadName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApsSlsADBJob"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApsSlsADBJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an AnalyticDB Pipeline Service (APS) job from Simple Log Service (SLS) to an AnalyticDB for MySQL Data Warehouse Edition cluster.
//
// @param request - CreateApsSlsADBJobRequest
//
// @return CreateApsSlsADBJobResponse
func (client *Client) CreateApsSlsADBJob(request *CreateApsSlsADBJobRequest) (_result *CreateApsSlsADBJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateApsSlsADBJobResponse{}
	_body, _err := client.CreateApsSlsADBJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a new webhook for the specified cluster or task type.
//
// @param tmpReq - CreateApsWebhookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApsWebhookResponse
func (client *Client) CreateApsWebhookWithOptions(tmpReq *CreateApsWebhookRequest, runtime *dara.RuntimeOptions) (_result *CreateApsWebhookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateApsWebhookShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Webhook) {
		request.WebhookShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Webhook, dara.String("Webhook"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.JobType) {
		body["JobType"] = request.JobType
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WebhookShrink) {
		body["Webhook"] = request.WebhookShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApsWebhook"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApsWebhookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a new webhook for the specified cluster or task type.
//
// @param request - CreateApsWebhookRequest
//
// @return CreateApsWebhookResponse
func (client *Client) CreateApsWebhook(request *CreateApsWebhookRequest) (_result *CreateApsWebhookResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateApsWebhookResponse{}
	_body, _err := client.CreateApsWebhookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a data backup for an AnalyticDB for MySQL instance.
//
// Description:
//
// *Before you call this operation, make sure that you fully understand the billing method and [pricing](https://www.aliyun.com/price/product#/ads/detail/ads_pre) of AnalyticDB for MySQL.*	- Temporary backups are the same as regular backups in terms of price and retention period of backup sets.
//
// @param request - CreateBackupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBackupResponse
func (client *Client) CreateBackupWithOptions(request *CreateBackupRequest, runtime *dara.RuntimeOptions) (_result *CreateBackupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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
		Action:      dara.String("CreateBackup"),
		Version:     dara.String("2021-12-01"),
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
// Creates a data backup for an AnalyticDB for MySQL instance.
//
// Description:
//
// *Before you call this operation, make sure that you fully understand the billing method and [pricing](https://www.aliyun.com/price/product#/ads/detail/ads_pre) of AnalyticDB for MySQL.*	- Temporary backups are the same as regular backups in terms of price and retention period of backup sets.
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
// Creates an AnalyticDB for MySQL Data Lakehouse Edition cluster.
//
// Description:
//
// # CreateDBCluster
//
// @param request - CreateDBClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDBClusterResponse
func (client *Client) CreateDBClusterWithOptions(request *CreateDBClusterRequest, runtime *dara.RuntimeOptions) (_result *CreateDBClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupSetId) {
		query["BackupSetId"] = request.BackupSetId
	}

	if !dara.IsNil(request.CloneSourceRegionId) {
		query["CloneSourceRegionId"] = request.CloneSourceRegionId
	}

	if !dara.IsNil(request.ComputeResource) {
		query["ComputeResource"] = request.ComputeResource
	}

	if !dara.IsNil(request.DBClusterDescription) {
		query["DBClusterDescription"] = request.DBClusterDescription
	}

	if !dara.IsNil(request.DBClusterNetworkType) {
		query["DBClusterNetworkType"] = request.DBClusterNetworkType
	}

	if !dara.IsNil(request.DBClusterVersion) {
		query["DBClusterVersion"] = request.DBClusterVersion
	}

	if !dara.IsNil(request.DiskEncryption) {
		query["DiskEncryption"] = request.DiskEncryption
	}

	if !dara.IsNil(request.EnableDefaultResourcePool) {
		query["EnableDefaultResourcePool"] = request.EnableDefaultResourcePool
	}

	if !dara.IsNil(request.EnableSSL) {
		query["EnableSSL"] = request.EnableSSL
	}

	if !dara.IsNil(request.KmsId) {
		query["KmsId"] = request.KmsId
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.ProductForm) {
		query["ProductForm"] = request.ProductForm
	}

	if !dara.IsNil(request.ProductVersion) {
		query["ProductVersion"] = request.ProductVersion
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReservedNodeCount) {
		query["ReservedNodeCount"] = request.ReservedNodeCount
	}

	if !dara.IsNil(request.ReservedNodeSize) {
		query["ReservedNodeSize"] = request.ReservedNodeSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.RestoreToTime) {
		query["RestoreToTime"] = request.RestoreToTime
	}

	if !dara.IsNil(request.RestoreType) {
		query["RestoreType"] = request.RestoreType
	}

	if !dara.IsNil(request.SecondaryVSwitchId) {
		query["SecondaryVSwitchId"] = request.SecondaryVSwitchId
	}

	if !dara.IsNil(request.SecondaryZoneId) {
		query["SecondaryZoneId"] = request.SecondaryZoneId
	}

	if !dara.IsNil(request.SourceDbClusterId) {
		query["SourceDbClusterId"] = request.SourceDbClusterId
	}

	if !dara.IsNil(request.StorageResource) {
		query["StorageResource"] = request.StorageResource
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

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDBCluster"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDBClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an AnalyticDB for MySQL Data Lakehouse Edition cluster.
//
// Description:
//
// # CreateDBCluster
//
// @param request - CreateDBClusterRequest
//
// @return CreateDBClusterResponse
func (client *Client) CreateDBCluster(request *CreateDBClusterRequest) (_result *CreateDBClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDBClusterResponse{}
	_body, _err := client.CreateDBClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a resource group for an AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see Endpoints.
//
// @param tmpReq - CreateDBResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDBResourceGroupResponse
func (client *Client) CreateDBResourceGroupWithOptions(tmpReq *CreateDBResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateDBResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDBResourceGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EngineParams) {
		request.EngineParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EngineParams, dara.String("EngineParams"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.GpuElasticPlan) {
		request.GpuElasticPlanShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GpuElasticPlan, dara.String("GpuElasticPlan"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RayConfig) {
		request.RayConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RayConfig, dara.String("RayConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Rules) {
		request.RulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Rules, dara.String("Rules"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoStopInterval) {
		query["AutoStopInterval"] = request.AutoStopInterval
	}

	if !dara.IsNil(request.ClusterMode) {
		query["ClusterMode"] = request.ClusterMode
	}

	if !dara.IsNil(request.ClusterSizeResource) {
		query["ClusterSizeResource"] = request.ClusterSizeResource
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EnableSpot) {
		query["EnableSpot"] = request.EnableSpot
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.EngineParamsShrink) {
		query["EngineParams"] = request.EngineParamsShrink
	}

	if !dara.IsNil(request.GpuElasticPlanShrink) {
		query["GpuElasticPlan"] = request.GpuElasticPlanShrink
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.GroupType) {
		query["GroupType"] = request.GroupType
	}

	if !dara.IsNil(request.MaxClusterCount) {
		query["MaxClusterCount"] = request.MaxClusterCount
	}

	if !dara.IsNil(request.MaxComputeResource) {
		query["MaxComputeResource"] = request.MaxComputeResource
	}

	if !dara.IsNil(request.MaxGpuQuantity) {
		query["MaxGpuQuantity"] = request.MaxGpuQuantity
	}

	if !dara.IsNil(request.MinClusterCount) {
		query["MinClusterCount"] = request.MinClusterCount
	}

	if !dara.IsNil(request.MinComputeResource) {
		query["MinComputeResource"] = request.MinComputeResource
	}

	if !dara.IsNil(request.MinGpuQuantity) {
		query["MinGpuQuantity"] = request.MinGpuQuantity
	}

	if !dara.IsNil(request.RayConfigShrink) {
		query["RayConfig"] = request.RayConfigShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RulesShrink) {
		query["Rules"] = request.RulesShrink
	}

	if !dara.IsNil(request.SpecName) {
		query["SpecName"] = request.SpecName
	}

	if !dara.IsNil(request.TargetResourceGroupName) {
		query["TargetResourceGroupName"] = request.TargetResourceGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDBResourceGroup"),
		Version:     dara.String("2021-12-01"),
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
// Creates a resource group for an AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see Endpoints.
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
// Creates a scaling plan for an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - CreateElasticPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateElasticPlanResponse
func (client *Client) CreateElasticPlanWithOptions(request *CreateElasticPlanRequest, runtime *dara.RuntimeOptions) (_result *CreateElasticPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoScale) {
		query["AutoScale"] = request.AutoScale
	}

	if !dara.IsNil(request.CronExpression) {
		query["CronExpression"] = request.CronExpression
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.ElasticPlanName) {
		query["ElasticPlanName"] = request.ElasticPlanName
	}

	if !dara.IsNil(request.Enabled) {
		query["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ResourceGroupName) {
		query["ResourceGroupName"] = request.ResourceGroupName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TargetSize) {
		query["TargetSize"] = request.TargetSize
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateElasticPlan"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateElasticPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a scaling plan for an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - CreateElasticPlanRequest
//
// @return CreateElasticPlanResponse
func (client *Client) CreateElasticPlan(request *CreateElasticPlanRequest) (_result *CreateElasticPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateElasticPlanResponse{}
	_body, _err := client.CreateElasticPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a lake storage.
//
// @param tmpReq - CreateLakeStorageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLakeStorageResponse
func (client *Client) CreateLakeStorageWithOptions(tmpReq *CreateLakeStorageRequest, runtime *dara.RuntimeOptions) (_result *CreateLakeStorageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateLakeStorageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Permissions) {
		request.PermissionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Permissions, dara.String("Permissions"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.PermissionsShrink) {
		body["Permissions"] = request.PermissionsShrink
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLakeStorage"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLakeStorageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a lake storage.
//
// @param request - CreateLakeStorageRequest
//
// @return CreateLakeStorageResponse
func (client *Client) CreateLakeStorage(request *CreateLakeStorageRequest) (_result *CreateLakeStorageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLakeStorageResponse{}
	_body, _err := client.CreateLakeStorageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a materialized view recommendation task.
//
// @param request - CreateMaterializedViewRecommendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMaterializedViewRecommendResponse
func (client *Client) CreateMaterializedViewRecommendWithOptions(request *CreateMaterializedViewRecommendRequest, runtime *dara.RuntimeOptions) (_result *CreateMaterializedViewRecommendResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.MinRewriteQueryCount) {
		query["MinRewriteQueryCount"] = request.MinRewriteQueryCount
	}

	if !dara.IsNil(request.MinRewriteQueryPattern) {
		query["MinRewriteQueryPattern"] = request.MinRewriteQueryPattern
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

	if !dara.IsNil(request.ScanQueriesRange) {
		query["ScanQueriesRange"] = request.ScanQueriesRange
	}

	if !dara.IsNil(request.SchedulingDay) {
		query["SchedulingDay"] = request.SchedulingDay
	}

	if !dara.IsNil(request.SchedulingPolicy) {
		query["SchedulingPolicy"] = request.SchedulingPolicy
	}

	if !dara.IsNil(request.SlowQueryThreshold) {
		query["SlowQueryThreshold"] = request.SlowQueryThreshold
	}

	if !dara.IsNil(request.SpecifiedTime) {
		query["SpecifiedTime"] = request.SpecifiedTime
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMaterializedViewRecommend"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMaterializedViewRecommendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a materialized view recommendation task.
//
// @param request - CreateMaterializedViewRecommendRequest
//
// @return CreateMaterializedViewRecommendResponse
func (client *Client) CreateMaterializedViewRecommend(request *CreateMaterializedViewRecommendRequest) (_result *CreateMaterializedViewRecommendResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMaterializedViewRecommendResponse{}
	_body, _err := client.CreateMaterializedViewRecommendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an Object Storage Service (OSS) subdirectory.
//
// Description:
//
//	  General endpoint: `adb.aliyuncs.com`.
//
//		- Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - CreateOssSubDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOssSubDirectoryResponse
func (client *Client) CreateOssSubDirectoryWithOptions(request *CreateOssSubDirectoryRequest, runtime *dara.RuntimeOptions) (_result *CreateOssSubDirectoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Path) {
		body["Path"] = request.Path
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOssSubDirectory"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOssSubDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an Object Storage Service (OSS) subdirectory.
//
// Description:
//
//	  General endpoint: `adb.aliyuncs.com`.
//
//		- Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - CreateOssSubDirectoryRequest
//
// @return CreateOssSubDirectoryResponse
func (client *Client) CreateOssSubDirectory(request *CreateOssSubDirectoryRequest) (_result *CreateOssSubDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateOssSubDirectoryResponse{}
	_body, _err := client.CreateOssSubDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a custom monitoring view.
//
// @param tmpReq - CreatePerformanceViewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePerformanceViewResponse
func (client *Client) CreatePerformanceViewWithOptions(tmpReq *CreatePerformanceViewRequest, runtime *dara.RuntimeOptions) (_result *CreatePerformanceViewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreatePerformanceViewShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ViewDetail) {
		request.ViewDetailShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ViewDetail, dara.String("ViewDetail"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateFromViewType) {
		query["CreateFromViewType"] = request.CreateFromViewType
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.FillOriginViewKeys) {
		query["FillOriginViewKeys"] = request.FillOriginViewKeys
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

	if !dara.IsNil(request.ViewDetailShrink) {
		query["ViewDetail"] = request.ViewDetailShrink
	}

	if !dara.IsNil(request.ViewName) {
		query["ViewName"] = request.ViewName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePerformanceView"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePerformanceViewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom monitoring view.
//
// @param request - CreatePerformanceViewRequest
//
// @return CreatePerformanceViewResponse
func (client *Client) CreatePerformanceView(request *CreatePerformanceViewRequest) (_result *CreatePerformanceViewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePerformanceViewResponse{}
	_body, _err := client.CreatePerformanceViewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Spark application template.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - CreateSparkTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSparkTemplateResponse
func (client *Client) CreateSparkTemplateWithOptions(request *CreateSparkTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateSparkTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ParentId) {
		body["ParentId"] = request.ParentId
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSparkTemplate"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSparkTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Spark application template.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - CreateSparkTemplateRequest
//
// @return CreateSparkTemplateResponse
func (client *Client) CreateSparkTemplate(request *CreateSparkTemplateRequest) (_result *CreateSparkTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSparkTemplateResponse{}
	_body, _err := client.CreateSparkTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a database account from an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
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

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAccount"),
		Version:     dara.String("2021-12-01"),
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
// Deletes a database account from an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
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
// Deletes an AnalyticDB Pipeline Service (APS) data source.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DeleteApsDatasoureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApsDatasoureResponse
func (client *Client) DeleteApsDatasoureWithOptions(request *DeleteApsDatasoureRequest, runtime *dara.RuntimeOptions) (_result *DeleteApsDatasoureResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DatasourceId) {
		body["DatasourceId"] = request.DatasourceId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApsDatasoure"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApsDatasoureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an AnalyticDB Pipeline Service (APS) data source.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DeleteApsDatasoureRequest
//
// @return DeleteApsDatasoureResponse
func (client *Client) DeleteApsDatasoure(request *DeleteApsDatasoureRequest) (_result *DeleteApsDatasoureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteApsDatasoureResponse{}
	_body, _err := client.DeleteApsDatasoureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an AnalyticDB Pipeline Service (APS) job.
//
// Description:
//
//	  Deleting backup sets is an asynchronous operation and may require 10 to 20 minutes to complete.
//
//		- You can delete up to 100 backup sets at a time. If you want to delete more than 100 backup sets, call this operation twice.
//
//		- To ensure data security, the system forcibly retains one valid backup set. If you want to delete the last backup set, the system prohibits your operation.
//
// @param request - DeleteApsJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApsJobResponse
func (client *Client) DeleteApsJobWithOptions(request *DeleteApsJobRequest, runtime *dara.RuntimeOptions) (_result *DeleteApsJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApsJobId) {
		body["ApsJobId"] = request.ApsJobId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApsJob"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApsJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an AnalyticDB Pipeline Service (APS) job.
//
// Description:
//
//	  Deleting backup sets is an asynchronous operation and may require 10 to 20 minutes to complete.
//
//		- You can delete up to 100 backup sets at a time. If you want to delete more than 100 backup sets, call this operation twice.
//
//		- To ensure data security, the system forcibly retains one valid backup set. If you want to delete the last backup set, the system prohibits your operation.
//
// @param request - DeleteApsJobRequest
//
// @return DeleteApsJobResponse
func (client *Client) DeleteApsJob(request *DeleteApsJobRequest) (_result *DeleteApsJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteApsJobResponse{}
	_body, _err := client.DeleteApsJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specific webhook in a specified cluster.
//
// Description:
//
// This API allows users to delete an existing webhook configuration by providing `RegionId`, `DBClusterId`, and `WebhookId`. Make sure that the provided parameter values are accurate to avoid deleting important settings by mistake.
//
// @param request - DeleteApsWebhookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApsWebhookResponse
func (client *Client) DeleteApsWebhookWithOptions(request *DeleteApsWebhookRequest, runtime *dara.RuntimeOptions) (_result *DeleteApsWebhookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WebhookId) {
		body["WebhookId"] = request.WebhookId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApsWebhook"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApsWebhookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specific webhook in a specified cluster.
//
// Description:
//
// This API allows users to delete an existing webhook configuration by providing `RegionId`, `DBClusterId`, and `WebhookId`. Make sure that the provided parameter values are accurate to avoid deleting important settings by mistake.
//
// @param request - DeleteApsWebhookRequest
//
// @return DeleteApsWebhookResponse
func (client *Client) DeleteApsWebhook(request *DeleteApsWebhookRequest) (_result *DeleteApsWebhookResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteApsWebhookResponse{}
	_body, _err := client.DeleteApsWebhookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Manually deletes backup sets.
//
// Description:
//
//	  You can delete up to 100 backup sets at a time. If you want to delete more than 100 backup sets, call this operation twice.
//
//		- To ensure data security, the system forcibly retains one valid backup set. If you want to delete the last backup set, the system prohibits your operation.
//
// @param request - DeleteBackupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBackupsResponse
func (client *Client) DeleteBackupsWithOptions(request *DeleteBackupsRequest, runtime *dara.RuntimeOptions) (_result *DeleteBackupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupIds) {
		query["BackupIds"] = request.BackupIds
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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
		Action:      dara.String("DeleteBackups"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBackupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Manually deletes backup sets.
//
// Description:
//
//	  You can delete up to 100 backup sets at a time. If you want to delete more than 100 backup sets, call this operation twice.
//
//		- To ensure data security, the system forcibly retains one valid backup set. If you want to delete the last backup set, the system prohibits your operation.
//
// @param request - DeleteBackupsRequest
//
// @return DeleteBackupsResponse
func (client *Client) DeleteBackups(request *DeleteBackupsRequest) (_result *DeleteBackupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteBackupsResponse{}
	_body, _err := client.DeleteBackupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an AnalyticDB for MySQL cluster.
//
// Description:
//
// ### [](#)
//
//   - You can delete only pay-as-you-go clusters.
//
//   - The cluster that you want to delete must be in the Running state.
//
//   - For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DeleteDBClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDBClusterResponse
func (client *Client) DeleteDBClusterWithOptions(request *DeleteDBClusterRequest, runtime *dara.RuntimeOptions) (_result *DeleteDBClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDBCluster"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDBClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an AnalyticDB for MySQL cluster.
//
// Description:
//
// ### [](#)
//
//   - You can delete only pay-as-you-go clusters.
//
//   - The cluster that you want to delete must be in the Running state.
//
//   - For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DeleteDBClusterRequest
//
// @return DeleteDBClusterResponse
func (client *Client) DeleteDBCluster(request *DeleteDBClusterRequest) (_result *DeleteDBClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDBClusterResponse{}
	_body, _err := client.DeleteDBClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a resource group from an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DeleteDBResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDBResourceGroupResponse
func (client *Client) DeleteDBResourceGroupWithOptions(request *DeleteDBResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteDBResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDBResourceGroup"),
		Version:     dara.String("2021-12-01"),
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
// Deletes a resource group from an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
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
// Deletes a scaling plan from an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DeleteElasticPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteElasticPlanResponse
func (client *Client) DeleteElasticPlanWithOptions(request *DeleteElasticPlanRequest, runtime *dara.RuntimeOptions) (_result *DeleteElasticPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.ElasticPlanName) {
		query["ElasticPlanName"] = request.ElasticPlanName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteElasticPlan"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteElasticPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a scaling plan from an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DeleteElasticPlanRequest
//
// @return DeleteElasticPlanResponse
func (client *Client) DeleteElasticPlan(request *DeleteElasticPlanRequest) (_result *DeleteElasticPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteElasticPlanResponse{}
	_body, _err := client.DeleteElasticPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a lake storage.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DeleteLakeStorageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLakeStorageResponse
func (client *Client) DeleteLakeStorageWithOptions(request *DeleteLakeStorageRequest, runtime *dara.RuntimeOptions) (_result *DeleteLakeStorageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.LakeStorageId) {
		body["LakeStorageId"] = request.LakeStorageId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLakeStorage"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLakeStorageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a lake storage.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DeleteLakeStorageRequest
//
// @return DeleteLakeStorageResponse
func (client *Client) DeleteLakeStorage(request *DeleteLakeStorageRequest) (_result *DeleteLakeStorageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteLakeStorageResponse{}
	_body, _err := client.DeleteLakeStorageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a materialized view recommendation task.
//
// @param request - DeleteMaterializedViewRecommendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMaterializedViewRecommendResponse
func (client *Client) DeleteMaterializedViewRecommendWithOptions(request *DeleteMaterializedViewRecommendRequest, runtime *dara.RuntimeOptions) (_result *DeleteMaterializedViewRecommendResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMaterializedViewRecommend"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMaterializedViewRecommendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a materialized view recommendation task.
//
// @param request - DeleteMaterializedViewRecommendRequest
//
// @return DeleteMaterializedViewRecommendResponse
func (client *Client) DeleteMaterializedViewRecommend(request *DeleteMaterializedViewRecommendRequest) (_result *DeleteMaterializedViewRecommendResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMaterializedViewRecommendResponse{}
	_body, _err := client.DeleteMaterializedViewRecommendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a monitoring view.
//
// @param request - DeletePerformanceViewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePerformanceViewResponse
func (client *Client) DeletePerformanceViewWithOptions(request *DeletePerformanceViewRequest, runtime *dara.RuntimeOptions) (_result *DeletePerformanceViewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.ViewName) {
		query["ViewName"] = request.ViewName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePerformanceView"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePerformanceViewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a monitoring view.
//
// @param request - DeletePerformanceViewRequest
//
// @return DeletePerformanceViewResponse
func (client *Client) DeletePerformanceView(request *DeletePerformanceViewRequest) (_result *DeletePerformanceViewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePerformanceViewResponse{}
	_body, _err := client.DeletePerformanceViewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes Spark template files.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - DeleteSparkTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSparkTemplateResponse
func (client *Client) DeleteSparkTemplateWithOptions(request *DeleteSparkTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteSparkTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSparkTemplate"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSparkTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes Spark template files.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - DeleteSparkTemplateRequest
//
// @return DeleteSparkTemplateResponse
func (client *Client) DeleteSparkTemplate(request *DeleteSparkTemplateRequest) (_result *DeleteSparkTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSparkTemplateResponse{}
	_body, _err := client.DeleteSparkTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes Spark template files.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - DeleteSparkTemplateFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSparkTemplateFileResponse
func (client *Client) DeleteSparkTemplateFileWithOptions(request *DeleteSparkTemplateFileRequest, runtime *dara.RuntimeOptions) (_result *DeleteSparkTemplateFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSparkTemplateFile"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSparkTemplateFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes Spark template files.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - DeleteSparkTemplateFileRequest
//
// @return DeleteSparkTemplateFileResponse
func (client *Client) DeleteSparkTemplateFile(request *DeleteSparkTemplateFileRequest) (_result *DeleteSparkTemplateFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSparkTemplateFileResponse{}
	_body, _err := client.DeleteSparkTemplateFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of AnalyticDB for MySQL clusters for AnalyticDB Pipeline Service (APS) federated analytics.
//
// Description:
//
// You can call this operation to query the performance data of a cluster over a time range based on performance metrics. The collection granularity is 30 seconds. This operation allows you to query information about slow queries, such as the SQL query duration, number of scanned rows, and amount of scanned data.
//
// @param request - DescribeAPSADBInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAPSADBInstancesResponse
func (client *Client) DescribeAPSADBInstancesWithOptions(request *DescribeAPSADBInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAPSADBInstancesResponse, _err error) {
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
		Action:      dara.String("DescribeAPSADBInstances"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAPSADBInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of AnalyticDB for MySQL clusters for AnalyticDB Pipeline Service (APS) federated analytics.
//
// Description:
//
// You can call this operation to query the performance data of a cluster over a time range based on performance metrics. The collection granularity is 30 seconds. This operation allows you to query information about slow queries, such as the SQL query duration, number of scanned rows, and amount of scanned data.
//
// @param request - DescribeAPSADBInstancesRequest
//
// @return DescribeAPSADBInstancesResponse
func (client *Client) DescribeAPSADBInstances(request *DescribeAPSADBInstancesRequest) (_result *DescribeAPSADBInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAPSADBInstancesResponse{}
	_body, _err := client.DescribeAPSADBInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries abnormal SQL patterns within a time range.
//
// @param request - DescribeAbnormalPatternDetectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAbnormalPatternDetectionResponse
func (client *Client) DescribeAbnormalPatternDetectionWithOptions(request *DescribeAbnormalPatternDetectionRequest, runtime *dara.RuntimeOptions) (_result *DescribeAbnormalPatternDetectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAbnormalPatternDetection"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAbnormalPatternDetectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries abnormal SQL patterns within a time range.
//
// @param request - DescribeAbnormalPatternDetectionRequest
//
// @return DescribeAbnormalPatternDetectionResponse
func (client *Client) DescribeAbnormalPatternDetection(request *DescribeAbnormalPatternDetectionRequest) (_result *DescribeAbnormalPatternDetectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAbnormalPatternDetectionResponse{}
	_body, _err := client.DescribeAbnormalPatternDetectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the permissions of a database account on all permission levels.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeAccountAllPrivilegesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccountAllPrivilegesResponse
func (client *Client) DescribeAccountAllPrivilegesWithOptions(request *DescribeAccountAllPrivilegesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccountAllPrivilegesResponse, _err error) {
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

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Marker) {
		query["Marker"] = request.Marker
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAccountAllPrivileges"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAccountAllPrivilegesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the permissions of a database account on all permission levels.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeAccountAllPrivilegesRequest
//
// @return DescribeAccountAllPrivilegesResponse
func (client *Client) DescribeAccountAllPrivileges(request *DescribeAccountAllPrivilegesRequest) (_result *DescribeAccountAllPrivilegesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAccountAllPrivilegesResponse{}
	_body, _err := client.DescribeAccountAllPrivilegesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the databases, tables, and columns on which a database account has permissions.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeAccountPrivilegeObjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccountPrivilegeObjectsResponse
func (client *Client) DescribeAccountPrivilegeObjectsWithOptions(request *DescribeAccountPrivilegeObjectsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccountPrivilegeObjectsResponse, _err error) {
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

	if !dara.IsNil(request.ColumnPrivilegeObject) {
		query["ColumnPrivilegeObject"] = request.ColumnPrivilegeObject
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DatabasePrivilegeObject) {
		query["DatabasePrivilegeObject"] = request.DatabasePrivilegeObject
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PrivilegeType) {
		query["PrivilegeType"] = request.PrivilegeType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TablePrivilegeObject) {
		query["TablePrivilegeObject"] = request.TablePrivilegeObject
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAccountPrivilegeObjects"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAccountPrivilegeObjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the databases, tables, and columns on which a database account has permissions.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeAccountPrivilegeObjectsRequest
//
// @return DescribeAccountPrivilegeObjectsResponse
func (client *Client) DescribeAccountPrivilegeObjects(request *DescribeAccountPrivilegeObjectsRequest) (_result *DescribeAccountPrivilegeObjectsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAccountPrivilegeObjectsResponse{}
	_body, _err := client.DescribeAccountPrivilegeObjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取某一ADB账户的权限
//
// @param request - DescribeAccountPrivilegesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccountPrivilegesResponse
func (client *Client) DescribeAccountPrivilegesWithOptions(request *DescribeAccountPrivilegesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccountPrivilegesResponse, _err error) {
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

	if !dara.IsNil(request.ColumnPrivilegeObject) {
		query["ColumnPrivilegeObject"] = request.ColumnPrivilegeObject
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DatabasePrivilegeObject) {
		query["DatabasePrivilegeObject"] = request.DatabasePrivilegeObject
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PrivilegeType) {
		query["PrivilegeType"] = request.PrivilegeType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TablePrivilegeObject) {
		query["TablePrivilegeObject"] = request.TablePrivilegeObject
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAccountPrivileges"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAccountPrivilegesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取某一ADB账户的权限
//
// @param request - DescribeAccountPrivilegesRequest
//
// @return DescribeAccountPrivilegesResponse
func (client *Client) DescribeAccountPrivileges(request *DescribeAccountPrivilegesRequest) (_result *DescribeAccountPrivilegesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAccountPrivilegesResponse{}
	_body, _err := client.DescribeAccountPrivilegesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the database accounts of an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
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
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAccounts"),
		Version:     dara.String("2021-12-01"),
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
// Queries the database accounts of an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
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
// Queries the information about table columns for an AnalyticDB for MySQL cluster.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - DescribeAdbMySqlColumnsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAdbMySqlColumnsResponse
func (client *Client) DescribeAdbMySqlColumnsWithOptions(request *DescribeAdbMySqlColumnsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAdbMySqlColumnsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Schema) {
		query["Schema"] = request.Schema
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAdbMySqlColumns"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAdbMySqlColumnsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about table columns for an AnalyticDB for MySQL cluster.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - DescribeAdbMySqlColumnsRequest
//
// @return DescribeAdbMySqlColumnsResponse
func (client *Client) DescribeAdbMySqlColumns(request *DescribeAdbMySqlColumnsRequest) (_result *DescribeAdbMySqlColumnsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAdbMySqlColumnsResponse{}
	_body, _err := client.DescribeAdbMySqlColumnsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about table indexes.
//
// @param request - DescribeAdbMySqlIndexesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAdbMySqlIndexesResponse
func (client *Client) DescribeAdbMySqlIndexesWithOptions(request *DescribeAdbMySqlIndexesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAdbMySqlIndexesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Schema) {
		query["Schema"] = request.Schema
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAdbMySqlIndexes"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAdbMySqlIndexesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about table indexes.
//
// @param request - DescribeAdbMySqlIndexesRequest
//
// @return DescribeAdbMySqlIndexesResponse
func (client *Client) DescribeAdbMySqlIndexes(request *DescribeAdbMySqlIndexesRequest) (_result *DescribeAdbMySqlIndexesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAdbMySqlIndexesResponse{}
	_body, _err := client.DescribeAdbMySqlIndexesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of databases for an AnalyticDB for MySQL cluster.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - DescribeAdbMySqlSchemasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAdbMySqlSchemasResponse
func (client *Client) DescribeAdbMySqlSchemasWithOptions(request *DescribeAdbMySqlSchemasRequest, runtime *dara.RuntimeOptions) (_result *DescribeAdbMySqlSchemasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAdbMySqlSchemas"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAdbMySqlSchemasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of databases for an AnalyticDB for MySQL cluster.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - DescribeAdbMySqlSchemasRequest
//
// @return DescribeAdbMySqlSchemasResponse
func (client *Client) DescribeAdbMySqlSchemas(request *DescribeAdbMySqlSchemasRequest) (_result *DescribeAdbMySqlSchemasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAdbMySqlSchemasResponse{}
	_body, _err := client.DescribeAdbMySqlSchemasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about table metadata.
//
// @param request - DescribeAdbMySqlTableMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAdbMySqlTableMetaResponse
func (client *Client) DescribeAdbMySqlTableMetaWithOptions(request *DescribeAdbMySqlTableMetaRequest, runtime *dara.RuntimeOptions) (_result *DescribeAdbMySqlTableMetaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Schema) {
		query["Schema"] = request.Schema
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAdbMySqlTableMeta"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAdbMySqlTableMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about table metadata.
//
// @param request - DescribeAdbMySqlTableMetaRequest
//
// @return DescribeAdbMySqlTableMetaResponse
func (client *Client) DescribeAdbMySqlTableMeta(request *DescribeAdbMySqlTableMetaRequest) (_result *DescribeAdbMySqlTableMetaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAdbMySqlTableMetaResponse{}
	_body, _err := client.DescribeAdbMySqlTableMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of tables for an AnalyticDB for MySQL cluster.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - DescribeAdbMySqlTablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAdbMySqlTablesResponse
func (client *Client) DescribeAdbMySqlTablesWithOptions(request *DescribeAdbMySqlTablesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAdbMySqlTablesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Schema) {
		query["Schema"] = request.Schema
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAdbMySqlTables"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAdbMySqlTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of tables for an AnalyticDB for MySQL cluster.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - DescribeAdbMySqlTablesRequest
//
// @return DescribeAdbMySqlTablesResponse
func (client *Client) DescribeAdbMySqlTables(request *DescribeAdbMySqlTablesRequest) (_result *DescribeAdbMySqlTablesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAdbMySqlTablesResponse{}
	_body, _err := client.DescribeAdbMySqlTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether the suggestion feature is enabled.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeAdviceServiceEnabledRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAdviceServiceEnabledResponse
func (client *Client) DescribeAdviceServiceEnabledWithOptions(request *DescribeAdviceServiceEnabledRequest, runtime *dara.RuntimeOptions) (_result *DescribeAdviceServiceEnabledResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAdviceServiceEnabled"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAdviceServiceEnabledResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether the suggestion feature is enabled.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeAdviceServiceEnabledRequest
//
// @return DescribeAdviceServiceEnabledResponse
func (client *Client) DescribeAdviceServiceEnabled(request *DescribeAdviceServiceEnabledRequest) (_result *DescribeAdviceServiceEnabledResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAdviceServiceEnabledResponse{}
	_body, _err := client.DescribeAdviceServiceEnabledWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of databases, tables, and columns in an AnalyticDB for MySQL cluster.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - DescribeAllDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAllDataSourceResponse
func (client *Client) DescribeAllDataSourceWithOptions(request *DescribeAllDataSourceRequest, runtime *dara.RuntimeOptions) (_result *DescribeAllDataSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SchemaName) {
		query["SchemaName"] = request.SchemaName
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAllDataSource"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAllDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of databases, tables, and columns in an AnalyticDB for MySQL cluster.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - DescribeAllDataSourceRequest
//
// @return DescribeAllDataSourceResponse
func (client *Client) DescribeAllDataSource(request *DescribeAllDataSourceRequest) (_result *DescribeAllDataSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAllDataSourceResponse{}
	_body, _err := client.DescribeAllDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the applied optimization suggestions for an AnalyticDB for MySQL cluster.
//
// @param request - DescribeAppliedAdvicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppliedAdvicesResponse
func (client *Client) DescribeAppliedAdvicesWithOptions(request *DescribeAppliedAdvicesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppliedAdvicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdviceType) {
		query["AdviceType"] = request.AdviceType
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SchemaTableName) {
		query["SchemaTableName"] = request.SchemaTableName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppliedAdvices"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppliedAdvicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the applied optimization suggestions for an AnalyticDB for MySQL cluster.
//
// @param request - DescribeAppliedAdvicesRequest
//
// @return DescribeAppliedAdvicesResponse
func (client *Client) DescribeAppliedAdvices(request *DescribeAppliedAdvicesRequest) (_result *DescribeAppliedAdvicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAppliedAdvicesResponse{}
	_body, _err := client.DescribeAppliedAdvicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the logs of a real-time data ingestion job for an AnalyticDB for MySQL cluster.
//
// Description:
//
//	  General endpoint: `adb.aliyuncs.com`.
//
//		- Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - DescribeApsActionLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApsActionLogsResponse
func (client *Client) DescribeApsActionLogsWithOptions(request *DescribeApsActionLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeApsActionLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
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

	if !dara.IsNil(request.Stage) {
		query["Stage"] = request.Stage
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	if !dara.IsNil(request.WorkloadId) {
		query["WorkloadId"] = request.WorkloadId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApsActionLogs"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApsActionLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the logs of a real-time data ingestion job for an AnalyticDB for MySQL cluster.
//
// Description:
//
//	  General endpoint: `adb.aliyuncs.com`.
//
//		- Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - DescribeApsActionLogsRequest
//
// @return DescribeApsActionLogsResponse
func (client *Client) DescribeApsActionLogs(request *DescribeApsActionLogsRequest) (_result *DescribeApsActionLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApsActionLogsResponse{}
	_body, _err := client.DescribeApsActionLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an AnalyticDB Pipeline Service (APS) data source.
//
// @param request - DescribeApsDatasourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApsDatasourceResponse
func (client *Client) DescribeApsDatasourceWithOptions(request *DescribeApsDatasourceRequest, runtime *dara.RuntimeOptions) (_result *DescribeApsDatasourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DatasourceId) {
		body["DatasourceId"] = request.DatasourceId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApsDatasource"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApsDatasourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an AnalyticDB Pipeline Service (APS) data source.
//
// @param request - DescribeApsDatasourceRequest
//
// @return DescribeApsDatasourceResponse
func (client *Client) DescribeApsDatasource(request *DescribeApsDatasourceRequest) (_result *DescribeApsDatasourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApsDatasourceResponse{}
	_body, _err := client.DescribeApsDatasourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of AnalyticDB Pipeline Service (APS) data sources.
//
// @param request - DescribeApsDatasourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApsDatasourcesResponse
func (client *Client) DescribeApsDatasourcesWithOptions(request *DescribeApsDatasourcesRequest, runtime *dara.RuntimeOptions) (_result *DescribeApsDatasourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DatasourceName) {
		body["DatasourceName"] = request.DatasourceName
	}

	if !dara.IsNil(request.DatasourceType) {
		body["DatasourceType"] = request.DatasourceType
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApsDatasources"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApsDatasourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of AnalyticDB Pipeline Service (APS) data sources.
//
// @param request - DescribeApsDatasourcesRequest
//
// @return DescribeApsDatasourcesResponse
func (client *Client) DescribeApsDatasources(request *DescribeApsDatasourcesRequest) (_result *DescribeApsDatasourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApsDatasourcesResponse{}
	_body, _err := client.DescribeApsDatasourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an AnalyticDB Pipeline Service (APS) job from a Hive data source.
//
// @param request - DescribeApsHiveWorkloadRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApsHiveWorkloadResponse
func (client *Client) DescribeApsHiveWorkloadWithOptions(request *DescribeApsHiveWorkloadRequest, runtime *dara.RuntimeOptions) (_result *DescribeApsHiveWorkloadResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WorkloadId) {
		body["WorkloadId"] = request.WorkloadId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApsHiveWorkload"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApsHiveWorkloadResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an AnalyticDB Pipeline Service (APS) job from a Hive data source.
//
// @param request - DescribeApsHiveWorkloadRequest
//
// @return DescribeApsHiveWorkloadResponse
func (client *Client) DescribeApsHiveWorkload(request *DescribeApsHiveWorkloadRequest) (_result *DescribeApsHiveWorkloadResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApsHiveWorkloadResponse{}
	_body, _err := client.DescribeApsHiveWorkloadWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an AnalyticDB Pipeline Service (APS) job.
//
// @param request - DescribeApsJobDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApsJobDetailResponse
func (client *Client) DescribeApsJobDetailWithOptions(request *DescribeApsJobDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeApsJobDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApsJobId) {
		body["ApsJobId"] = request.ApsJobId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApsJobDetail"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApsJobDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an AnalyticDB Pipeline Service (APS) job.
//
// @param request - DescribeApsJobDetailRequest
//
// @return DescribeApsJobDetailResponse
func (client *Client) DescribeApsJobDetail(request *DescribeApsJobDetailRequest) (_result *DescribeApsJobDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApsJobDetailResponse{}
	_body, _err := client.DescribeApsJobDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of AnalyticDB Pipeline Service (APS) jobs.
//
// @param request - DescribeApsJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApsJobsResponse
func (client *Client) DescribeApsJobsWithOptions(request *DescribeApsJobsRequest, runtime *dara.RuntimeOptions) (_result *DescribeApsJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApsJobName) {
		body["ApsJobName"] = request.ApsJobName
	}

	if !dara.IsNil(request.CreateTimeEnd) {
		body["CreateTimeEnd"] = request.CreateTimeEnd
	}

	if !dara.IsNil(request.CreateTimeStart) {
		body["CreateTimeStart"] = request.CreateTimeStart
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApsJobs"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApsJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of AnalyticDB Pipeline Service (APS) jobs.
//
// @param request - DescribeApsJobsRequest
//
// @return DescribeApsJobsResponse
func (client *Client) DescribeApsJobs(request *DescribeApsJobsRequest) (_result *DescribeApsJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApsJobsResponse{}
	_body, _err := client.DescribeApsJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the workloads of AnalyticDB Pipeline Service (APS) migration jobs.
//
// @param request - DescribeApsMigrationWorkloadsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApsMigrationWorkloadsResponse
func (client *Client) DescribeApsMigrationWorkloadsWithOptions(request *DescribeApsMigrationWorkloadsRequest, runtime *dara.RuntimeOptions) (_result *DescribeApsMigrationWorkloadsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OssLocation) {
		body["OssLocation"] = request.OssLocation
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.WorkloadName) {
		body["WorkloadName"] = request.WorkloadName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApsMigrationWorkloads"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApsMigrationWorkloadsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the workloads of AnalyticDB Pipeline Service (APS) migration jobs.
//
// @param request - DescribeApsMigrationWorkloadsRequest
//
// @return DescribeApsMigrationWorkloadsResponse
func (client *Client) DescribeApsMigrationWorkloads(request *DescribeApsMigrationWorkloadsRequest) (_result *DescribeApsMigrationWorkloadsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApsMigrationWorkloadsResponse{}
	_body, _err := client.DescribeApsMigrationWorkloadsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the progress of an AnalyticDB Pipeline Service (APS) job.
//
// @param request - DescribeApsProgressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApsProgressResponse
func (client *Client) DescribeApsProgressWithOptions(request *DescribeApsProgressRequest, runtime *dara.RuntimeOptions) (_result *DescribeApsProgressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WorkloadId) {
		body["WorkloadId"] = request.WorkloadId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApsProgress"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApsProgressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the progress of an AnalyticDB Pipeline Service (APS) job.
//
// @param request - DescribeApsProgressRequest
//
// @return DescribeApsProgressResponse
func (client *Client) DescribeApsProgress(request *DescribeApsProgressRequest) (_result *DescribeApsProgressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApsProgressResponse{}
	_body, _err := client.DescribeApsProgressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about resource groups of an AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - DescribeApsResourceGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApsResourceGroupsResponse
func (client *Client) DescribeApsResourceGroupsWithOptions(request *DescribeApsResourceGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeApsResourceGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WorkloadId) {
		body["WorkloadId"] = request.WorkloadId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApsResourceGroups"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApsResourceGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about resource groups of an AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - DescribeApsResourceGroupsRequest
//
// @return DescribeApsResourceGroupsResponse
func (client *Client) DescribeApsResourceGroups(request *DescribeApsResourceGroupsRequest) (_result *DescribeApsResourceGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApsResourceGroupsResponse{}
	_body, _err := client.DescribeApsResourceGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the SQL audit logs of an AnalyticDB for MySQL cluster.
//
// Description:
//
//	  SQL audit logs can be queried only when SQL audit is enabled. Only SQL audit logs within the last 30 days can be queried. If SQL audit was disabled and re-enabled, only SQL audit logs from the time when SQL audit was re-enabled can be queried. The following operations are not recorded in SQL audit logs: **INSERT INTO VALUES**, **REPLACE INTO VALUES**, and **UPSERT INTO VALUES**.
//
//		- For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeAuditLogRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAuditLogRecordsResponse
func (client *Client) DescribeAuditLogRecordsWithOptions(request *DescribeAuditLogRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAuditLogRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.HostAddress) {
		query["HostAddress"] = request.HostAddress
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
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

	if !dara.IsNil(request.ProxyUser) {
		query["ProxyUser"] = request.ProxyUser
	}

	if !dara.IsNil(request.QueryKeyword) {
		query["QueryKeyword"] = request.QueryKeyword
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

	if !dara.IsNil(request.SqlType) {
		query["SqlType"] = request.SqlType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Succeed) {
		query["Succeed"] = request.Succeed
	}

	if !dara.IsNil(request.User) {
		query["User"] = request.User
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAuditLogRecords"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAuditLogRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the SQL audit logs of an AnalyticDB for MySQL cluster.
//
// Description:
//
//	  SQL audit logs can be queried only when SQL audit is enabled. Only SQL audit logs within the last 30 days can be queried. If SQL audit was disabled and re-enabled, only SQL audit logs from the time when SQL audit was re-enabled can be queried. The following operations are not recorded in SQL audit logs: **INSERT INTO VALUES**, **REPLACE INTO VALUES**, and **UPSERT INTO VALUES**.
//
//		- For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeAuditLogRecordsRequest
//
// @return DescribeAuditLogRecordsResponse
func (client *Client) DescribeAuditLogRecords(request *DescribeAuditLogRecordsRequest) (_result *DescribeAuditLogRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAuditLogRecordsResponse{}
	_body, _err := client.DescribeAuditLogRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the available optimization suggestions for an AnalyticDB for MySQL cluster.
//
// @param request - DescribeAvailableAdvicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAvailableAdvicesResponse
func (client *Client) DescribeAvailableAdvicesWithOptions(request *DescribeAvailableAdvicesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAvailableAdvicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdviceDate) {
		query["AdviceDate"] = request.AdviceDate
	}

	if !dara.IsNil(request.AdviceType) {
		query["AdviceType"] = request.AdviceType
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SchemaTableName) {
		query["SchemaTableName"] = request.SchemaTableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAvailableAdvices"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAvailableAdvicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the available optimization suggestions for an AnalyticDB for MySQL cluster.
//
// @param request - DescribeAvailableAdvicesRequest
//
// @return DescribeAvailableAdvicesResponse
func (client *Client) DescribeAvailableAdvices(request *DescribeAvailableAdvicesRequest) (_result *DescribeAvailableAdvicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAvailableAdvicesResponse{}
	_body, _err := client.DescribeAvailableAdvicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看集群备份设置
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
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
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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
		Action:      dara.String("DescribeBackupPolicy"),
		Version:     dara.String("2021-12-01"),
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
// 查看集群备份设置
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
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
// 查询实例备份集
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
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

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.Remote) {
		query["Remote"] = request.Remote
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
		Action:      dara.String("DescribeBackups"),
		Version:     dara.String("2021-12-01"),
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
// 查询实例备份集
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
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
// Queries the bad SQL statements that affect cluster stability within a time range.
//
// @param request - DescribeBadSqlDetectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBadSqlDetectionResponse
func (client *Client) DescribeBadSqlDetectionWithOptions(request *DescribeBadSqlDetectionRequest, runtime *dara.RuntimeOptions) (_result *DescribeBadSqlDetectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBadSqlDetection"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBadSqlDetectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the bad SQL statements that affect cluster stability within a time range.
//
// @param request - DescribeBadSqlDetectionRequest
//
// @return DescribeBadSqlDetectionResponse
func (client *Client) DescribeBadSqlDetection(request *DescribeBadSqlDetectionRequest) (_result *DescribeBadSqlDetectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBadSqlDetectionResponse{}
	_body, _err := client.DescribeBadSqlDetectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the IP address whitelists of an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeClusterAccessWhiteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterAccessWhiteListResponse
func (client *Client) DescribeClusterAccessWhiteListWithOptions(request *DescribeClusterAccessWhiteListRequest, runtime *dara.RuntimeOptions) (_result *DescribeClusterAccessWhiteListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClusterAccessWhiteList"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClusterAccessWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the IP address whitelists of an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeClusterAccessWhiteListRequest
//
// @return DescribeClusterAccessWhiteListResponse
func (client *Client) DescribeClusterAccessWhiteList(request *DescribeClusterAccessWhiteListRequest) (_result *DescribeClusterAccessWhiteListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeClusterAccessWhiteListResponse{}
	_body, _err := client.DescribeClusterAccessWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the network information about an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeClusterNetInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterNetInfoResponse
func (client *Client) DescribeClusterNetInfoWithOptions(request *DescribeClusterNetInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeClusterNetInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClusterNetInfo"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClusterNetInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the network information about an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeClusterNetInfoRequest
//
// @return DescribeClusterNetInfoResponse
func (client *Client) DescribeClusterNetInfo(request *DescribeClusterNetInfoRequest) (_result *DescribeClusterNetInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeClusterNetInfoResponse{}
	_body, _err := client.DescribeClusterNetInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about resource usage of an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeClusterResourceDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterResourceDetailResponse
func (client *Client) DescribeClusterResourceDetailWithOptions(request *DescribeClusterResourceDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeClusterResourceDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClusterResourceDetail"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClusterResourceDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about resource usage of an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeClusterResourceDetailRequest
//
// @return DescribeClusterResourceDetailResponse
func (client *Client) DescribeClusterResourceDetail(request *DescribeClusterResourceDetailRequest) (_result *DescribeClusterResourceDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeClusterResourceDetailResponse{}
	_body, _err := client.DescribeClusterResourceDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the resource usage of an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeClusterResourceUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterResourceUsageResponse
func (client *Client) DescribeClusterResourceUsageWithOptions(request *DescribeClusterResourceUsageRequest, runtime *dara.RuntimeOptions) (_result *DescribeClusterResourceUsageResponse, _err error) {
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
		Action:      dara.String("DescribeClusterResourceUsage"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClusterResourceUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resource usage of an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeClusterResourceUsageRequest
//
// @return DescribeClusterResourceUsageResponse
func (client *Client) DescribeClusterResourceUsage(request *DescribeClusterResourceUsageRequest) (_result *DescribeClusterResourceUsageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeClusterResourceUsageResponse{}
	_body, _err := client.DescribeClusterResourceUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of columns in a table.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - DescribeColumnsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeColumnsResponse
func (client *Client) DescribeColumnsWithOptions(request *DescribeColumnsRequest, runtime *dara.RuntimeOptions) (_result *DescribeColumnsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SchemaName) {
		query["SchemaName"] = request.SchemaName
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeColumns"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeColumnsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of columns in a table.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - DescribeColumnsRequest
//
// @return DescribeColumnsResponse
func (client *Client) DescribeColumns(request *DescribeColumnsRequest) (_result *DescribeColumnsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeColumnsResponse{}
	_body, _err := client.DescribeColumnsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether the remote build feature is enabled in the query acceleration configuration of an AnalyticDB for MySQL cluster.
//
// @param request - DescribeCompactionServiceSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCompactionServiceSwitchResponse
func (client *Client) DescribeCompactionServiceSwitchWithOptions(request *DescribeCompactionServiceSwitchRequest, runtime *dara.RuntimeOptions) (_result *DescribeCompactionServiceSwitchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCompactionServiceSwitch"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCompactionServiceSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether the remote build feature is enabled in the query acceleration configuration of an AnalyticDB for MySQL cluster.
//
// @param request - DescribeCompactionServiceSwitchRequest
//
// @return DescribeCompactionServiceSwitchResponse
func (client *Client) DescribeCompactionServiceSwitch(request *DescribeCompactionServiceSwitchRequest) (_result *DescribeCompactionServiceSwitchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCompactionServiceSwitchResponse{}
	_body, _err := client.DescribeCompactionServiceSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the computing resource usage of a resource group in an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeComputeResourceUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeComputeResourceUsageResponse
func (client *Client) DescribeComputeResourceUsageWithOptions(request *DescribeComputeResourceUsageRequest, runtime *dara.RuntimeOptions) (_result *DescribeComputeResourceUsageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
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
		Action:      dara.String("DescribeComputeResourceUsage"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeComputeResourceUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the computing resource usage of a resource group in an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeComputeResourceUsageRequest
//
// @return DescribeComputeResourceUsageResponse
func (client *Client) DescribeComputeResourceUsage(request *DescribeComputeResourceUsageRequest) (_result *DescribeComputeResourceUsageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeComputeResourceUsageResponse{}
	_body, _err := client.DescribeComputeResourceUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the diagnostic results of the access layer.
//
// @param request - DescribeControllerDetectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeControllerDetectionResponse
func (client *Client) DescribeControllerDetectionWithOptions(request *DescribeControllerDetectionRequest, runtime *dara.RuntimeOptions) (_result *DescribeControllerDetectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeControllerDetection"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeControllerDetectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the diagnostic results of the access layer.
//
// @param request - DescribeControllerDetectionRequest
//
// @return DescribeControllerDetectionResponse
func (client *Client) DescribeControllerDetection(request *DescribeControllerDetectionRequest) (_result *DescribeControllerDetectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeControllerDetectionResponse{}
	_body, _err := client.DescribeControllerDetectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeDBClusterAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterAttributeResponse
func (client *Client) DescribeDBClusterAttributeWithOptions(request *DescribeDBClusterAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClusterAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBClusterAttribute"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBClusterAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeDBClusterAttributeRequest
//
// @return DescribeDBClusterAttributeResponse
func (client *Client) DescribeDBClusterAttribute(request *DescribeDBClusterAttributeRequest) (_result *DescribeDBClusterAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBClusterAttributeResponse{}
	_body, _err := client.DescribeDBClusterAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the health status of an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeDBClusterHealthStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterHealthStatusResponse
func (client *Client) DescribeDBClusterHealthStatusWithOptions(request *DescribeDBClusterHealthStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClusterHealthStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBClusterHealthStatus"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBClusterHealthStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the health status of an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeDBClusterHealthStatusRequest
//
// @return DescribeDBClusterHealthStatusResponse
func (client *Client) DescribeDBClusterHealthStatus(request *DescribeDBClusterHealthStatusRequest) (_result *DescribeDBClusterHealthStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBClusterHealthStatusResponse{}
	_body, _err := client.DescribeDBClusterHealthStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the performance data of an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeDBClusterPerformanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterPerformanceResponse
func (client *Client) DescribeDBClusterPerformanceWithOptions(request *DescribeDBClusterPerformanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClusterPerformanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.ResourcePools) {
		query["ResourcePools"] = request.ResourcePools
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBClusterPerformance"),
		Version:     dara.String("2021-12-01"),
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
// Queries the performance data of an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
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
// Queries the SSL configurations of a cluster.
//
// @param request - DescribeDBClusterSSLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterSSLResponse
func (client *Client) DescribeDBClusterSSLWithOptions(request *DescribeDBClusterSSLRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClusterSSLResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBClusterSSL"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBClusterSSLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the SSL configurations of a cluster.
//
// @param request - DescribeDBClusterSSLRequest
//
// @return DescribeDBClusterSSLResponse
func (client *Client) DescribeDBClusterSSL(request *DescribeDBClusterSSLRequest) (_result *DescribeDBClusterSSLResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBClusterSSLResponse{}
	_body, _err := client.DescribeDBClusterSSLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the storage overview information of an AnalyticDB for MySQL cluster, such as the total data size, hot data size, cold data size, and data growth.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeDBClusterSpaceSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterSpaceSummaryResponse
func (client *Client) DescribeDBClusterSpaceSummaryWithOptions(request *DescribeDBClusterSpaceSummaryRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClusterSpaceSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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
		Action:      dara.String("DescribeDBClusterSpaceSummary"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBClusterSpaceSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the storage overview information of an AnalyticDB for MySQL cluster, such as the total data size, hot data size, cold data size, and data growth.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeDBClusterSpaceSummaryRequest
//
// @return DescribeDBClusterSpaceSummaryResponse
func (client *Client) DescribeDBClusterSpaceSummary(request *DescribeDBClusterSpaceSummaryRequest) (_result *DescribeDBClusterSpaceSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBClusterSpaceSummaryResponse{}
	_body, _err := client.DescribeDBClusterSpaceSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the statuses of AnalyticDB for MySQL clusters within a region.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeDBClusterStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterStatusResponse
func (client *Client) DescribeDBClusterStatusWithOptions(request *DescribeDBClusterStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClusterStatusResponse, _err error) {
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
		Action:      dara.String("DescribeDBClusterStatus"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBClusterStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statuses of AnalyticDB for MySQL clusters within a region.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeDBClusterStatusRequest
//
// @return DescribeDBClusterStatusResponse
func (client *Client) DescribeDBClusterStatus(request *DescribeDBClusterStatusRequest) (_result *DescribeDBClusterStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBClusterStatusResponse{}
	_body, _err := client.DescribeDBClusterStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about AnalyticDB for MySQL Data Lakehouse Edition clusters within a region.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeDBClustersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClustersResponse
func (client *Client) DescribeDBClustersWithOptions(request *DescribeDBClustersRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClustersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterDescription) {
		query["DBClusterDescription"] = request.DBClusterDescription
	}

	if !dara.IsNil(request.DBClusterIds) {
		query["DBClusterIds"] = request.DBClusterIds
	}

	if !dara.IsNil(request.DBClusterStatus) {
		query["DBClusterStatus"] = request.DBClusterStatus
	}

	if !dara.IsNil(request.DBClusterVersion) {
		query["DBClusterVersion"] = request.DBClusterVersion
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBClusters"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about AnalyticDB for MySQL Data Lakehouse Edition clusters within a region.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeDBClustersRequest
//
// @return DescribeDBClustersResponse
func (client *Client) DescribeDBClusters(request *DescribeDBClustersRequest) (_result *DescribeDBClustersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBClustersResponse{}
	_body, _err := client.DescribeDBClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about resource groups of an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeDBResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBResourceGroupResponse
func (client *Client) DescribeDBResourceGroupWithOptions(request *DescribeDBResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.GroupType) {
		query["GroupType"] = request.GroupType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBResourceGroup"),
		Version:     dara.String("2021-12-01"),
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
// Queries the information about resource groups of an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
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
// Queries the deduplicated statistics of resource groups, databases, usernames, and source IP addresses about SQL statements that meet a query condition for an AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeDiagnosisDimensionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDiagnosisDimensionsResponse
func (client *Client) DescribeDiagnosisDimensionsWithOptions(request *DescribeDiagnosisDimensionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDiagnosisDimensionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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
		Action:      dara.String("DescribeDiagnosisDimensions"),
		Version:     dara.String("2021-12-01"),
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
// Queries the deduplicated statistics of resource groups, databases, usernames, and source IP addresses about SQL statements that meet a query condition for an AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
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
// Queries the diagnostic information about SQL statements that meet a query condition for an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see Endpoints.
//
// @param request - DescribeDiagnosisRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDiagnosisRecordsResponse
func (client *Client) DescribeDiagnosisRecordsWithOptions(request *DescribeDiagnosisRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDiagnosisRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientIp) {
		query["ClientIp"] = request.ClientIp
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MaxPeakMemory) {
		query["MaxPeakMemory"] = request.MaxPeakMemory
	}

	if !dara.IsNil(request.MaxScanSize) {
		query["MaxScanSize"] = request.MaxScanSize
	}

	if !dara.IsNil(request.MinPeakMemory) {
		query["MinPeakMemory"] = request.MinPeakMemory
	}

	if !dara.IsNil(request.MinScanSize) {
		query["MinScanSize"] = request.MinScanSize
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

	if !dara.IsNil(request.PatternId) {
		query["PatternId"] = request.PatternId
	}

	if !dara.IsNil(request.QueryCondition) {
		query["QueryCondition"] = request.QueryCondition
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroup) {
		query["ResourceGroup"] = request.ResourceGroup
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDiagnosisRecords"),
		Version:     dara.String("2021-12-01"),
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
// Queries the diagnostic information about SQL statements that meet a query condition for an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see Endpoints.
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
// Queries the execution information about an SQL statement, including the execution plan, execution information, resource usage, and self-diagnostics results.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeDiagnosisSQLInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDiagnosisSQLInfoResponse
func (client *Client) DescribeDiagnosisSQLInfoWithOptions(request *DescribeDiagnosisSQLInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDiagnosisSQLInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ProcessId) {
		query["ProcessId"] = request.ProcessId
	}

	if !dara.IsNil(request.ProcessRcHost) {
		query["ProcessRcHost"] = request.ProcessRcHost
	}

	if !dara.IsNil(request.ProcessStartTime) {
		query["ProcessStartTime"] = request.ProcessStartTime
	}

	if !dara.IsNil(request.ProcessState) {
		query["ProcessState"] = request.ProcessState
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDiagnosisSQLInfo"),
		Version:     dara.String("2021-12-01"),
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
// Queries the execution information about an SQL statement, including the execution plan, execution information, resource usage, and self-diagnostics results.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
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
// Queries the last five SQL query download tasks of an AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeDownloadRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDownloadRecordsResponse
func (client *Client) DescribeDownloadRecordsWithOptions(request *DescribeDownloadRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDownloadRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDownloadRecords"),
		Version:     dara.String("2021-12-01"),
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
// Queries the last five SQL query download tasks of an AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
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
// Queries the information about a scaling plan for an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeElasticPlanAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeElasticPlanAttributeResponse
func (client *Client) DescribeElasticPlanAttributeWithOptions(request *DescribeElasticPlanAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeElasticPlanAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.ElasticPlanName) {
		query["ElasticPlanName"] = request.ElasticPlanName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeElasticPlanAttribute"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeElasticPlanAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a scaling plan for an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeElasticPlanAttributeRequest
//
// @return DescribeElasticPlanAttributeResponse
func (client *Client) DescribeElasticPlanAttribute(request *DescribeElasticPlanAttributeRequest) (_result *DescribeElasticPlanAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeElasticPlanAttributeResponse{}
	_body, _err := client.DescribeElasticPlanAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of scaling plan jobs for an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see Endpoints.
//
// @param request - DescribeElasticPlanJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeElasticPlanJobsResponse
func (client *Client) DescribeElasticPlanJobsWithOptions(request *DescribeElasticPlanJobsRequest, runtime *dara.RuntimeOptions) (_result *DescribeElasticPlanJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.ElasticPlanName) {
		query["ElasticPlanName"] = request.ElasticPlanName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupName) {
		query["ResourceGroupName"] = request.ResourceGroupName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeElasticPlanJobs"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeElasticPlanJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of scaling plan jobs for an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see Endpoints.
//
// @param request - DescribeElasticPlanJobsRequest
//
// @return DescribeElasticPlanJobsResponse
func (client *Client) DescribeElasticPlanJobs(request *DescribeElasticPlanJobsRequest) (_result *DescribeElasticPlanJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeElasticPlanJobsResponse{}
	_body, _err := client.DescribeElasticPlanJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the resource specifications that are supported by different types of scaling plans of an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeElasticPlanSpecificationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeElasticPlanSpecificationsResponse
func (client *Client) DescribeElasticPlanSpecificationsWithOptions(request *DescribeElasticPlanSpecificationsRequest, runtime *dara.RuntimeOptions) (_result *DescribeElasticPlanSpecificationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.ResourceGroupName) {
		query["ResourceGroupName"] = request.ResourceGroupName
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeElasticPlanSpecifications"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeElasticPlanSpecificationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resource specifications that are supported by different types of scaling plans of an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeElasticPlanSpecificationsRequest
//
// @return DescribeElasticPlanSpecificationsResponse
func (client *Client) DescribeElasticPlanSpecifications(request *DescribeElasticPlanSpecificationsRequest) (_result *DescribeElasticPlanSpecificationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeElasticPlanSpecificationsResponse{}
	_body, _err := client.DescribeElasticPlanSpecificationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries scaling plans of an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeElasticPlansRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeElasticPlansResponse
func (client *Client) DescribeElasticPlansWithOptions(request *DescribeElasticPlansRequest, runtime *dara.RuntimeOptions) (_result *DescribeElasticPlansResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.ElasticPlanName) {
		query["ElasticPlanName"] = request.ElasticPlanName
	}

	if !dara.IsNil(request.Enabled) {
		query["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupName) {
		query["ResourceGroupName"] = request.ResourceGroupName
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeElasticPlans"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeElasticPlansResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries scaling plans of an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeElasticPlansRequest
//
// @return DescribeElasticPlansResponse
func (client *Client) DescribeElasticPlans(request *DescribeElasticPlansRequest) (_result *DescribeElasticPlansResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeElasticPlansResponse{}
	_body, _err := client.DescribeElasticPlansWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the permission level and permissions supported for an AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
//
// @param request - DescribeEnabledPrivilegesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEnabledPrivilegesResponse
func (client *Client) DescribeEnabledPrivilegesWithOptions(request *DescribeEnabledPrivilegesRequest, runtime *dara.RuntimeOptions) (_result *DescribeEnabledPrivilegesResponse, _err error) {
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
		Action:      dara.String("DescribeEnabledPrivileges"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEnabledPrivilegesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the permission level and permissions supported for an AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
//
// @param request - DescribeEnabledPrivilegesRequest
//
// @return DescribeEnabledPrivilegesResponse
func (client *Client) DescribeEnabledPrivileges(request *DescribeEnabledPrivilegesRequest) (_result *DescribeEnabledPrivilegesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEnabledPrivilegesResponse{}
	_body, _err := client.DescribeEnabledPrivilegesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the disk cache size in the query acceleration configuration of an AnalyticDB for MySQL cluster.
//
// @param request - DescribeEssdCacheConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEssdCacheConfigResponse
func (client *Client) DescribeEssdCacheConfigWithOptions(request *DescribeEssdCacheConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeEssdCacheConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEssdCacheConfig"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEssdCacheConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the disk cache size in the query acceleration configuration of an AnalyticDB for MySQL cluster.
//
// @param request - DescribeEssdCacheConfigRequest
//
// @return DescribeEssdCacheConfigResponse
func (client *Client) DescribeEssdCacheConfig(request *DescribeEssdCacheConfigRequest) (_result *DescribeEssdCacheConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEssdCacheConfigResponse{}
	_body, _err := client.DescribeEssdCacheConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about tables that have excessive primary key fields in an AnalyticDB for MySQL Data Lakehouse Edition (V5.0) cluster.
//
// @param request - DescribeExcessivePrimaryKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExcessivePrimaryKeysResponse
func (client *Client) DescribeExcessivePrimaryKeysWithOptions(request *DescribeExcessivePrimaryKeysRequest, runtime *dara.RuntimeOptions) (_result *DescribeExcessivePrimaryKeysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
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
		Action:      dara.String("DescribeExcessivePrimaryKeys"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExcessivePrimaryKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about tables that have excessive primary key fields in an AnalyticDB for MySQL Data Lakehouse Edition (V5.0) cluster.
//
// @param request - DescribeExcessivePrimaryKeysRequest
//
// @return DescribeExcessivePrimaryKeysResponse
func (client *Client) DescribeExcessivePrimaryKeys(request *DescribeExcessivePrimaryKeysRequest) (_result *DescribeExcessivePrimaryKeysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeExcessivePrimaryKeysResponse{}
	_body, _err := client.DescribeExcessivePrimaryKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the diagnostic results of the compute layer.
//
// @param request - DescribeExecutorDetectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExecutorDetectionResponse
func (client *Client) DescribeExecutorDetectionWithOptions(request *DescribeExecutorDetectionRequest, runtime *dara.RuntimeOptions) (_result *DescribeExecutorDetectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeExecutorDetection"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExecutorDetectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the diagnostic results of the compute layer.
//
// @param request - DescribeExecutorDetectionRequest
//
// @return DescribeExecutorDetectionResponse
func (client *Client) DescribeExecutorDetection(request *DescribeExecutorDetectionRequest) (_result *DescribeExecutorDetectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeExecutorDetectionResponse{}
	_body, _err := client.DescribeExecutorDetectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieve historical task records.
//
// Description:
//
// Only supports viewing tasks within the last 30 days.
//
// @param request - DescribeHistoryTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHistoryTasksResponse
func (client *Client) DescribeHistoryTasksWithOptions(request *DescribeHistoryTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeHistoryTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
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
		Version:     dara.String("2021-12-01"),
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
// Retrieve historical task records.
//
// Description:
//
// Only supports viewing tasks within the last 30 days.
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
// Queries task statistics.
//
// @param request - DescribeHistoryTasksStatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHistoryTasksStatResponse
func (client *Client) DescribeHistoryTasksStatWithOptions(request *DescribeHistoryTasksStatRequest, runtime *dara.RuntimeOptions) (_result *DescribeHistoryTasksStatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("DescribeHistoryTasksStat"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHistoryTasksStatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries task statistics.
//
// @param request - DescribeHistoryTasksStatRequest
//
// @return DescribeHistoryTasksStatResponse
func (client *Client) DescribeHistoryTasksStat(request *DescribeHistoryTasksStatRequest) (_result *DescribeHistoryTasksStatResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHistoryTasksStatResponse{}
	_body, _err := client.DescribeHistoryTasksStatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the disk usage of all storage nodes.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeInclinedNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInclinedNodesResponse
func (client *Client) DescribeInclinedNodesWithOptions(request *DescribeInclinedNodesRequest, runtime *dara.RuntimeOptions) (_result *DescribeInclinedNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
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
		Action:      dara.String("DescribeInclinedNodes"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInclinedNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the disk usage of all storage nodes.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeInclinedNodesRequest
//
// @return DescribeInclinedNodesResponse
func (client *Client) DescribeInclinedNodes(request *DescribeInclinedNodesRequest) (_result *DescribeInclinedNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInclinedNodesResponse{}
	_body, _err := client.DescribeInclinedNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about skewed tables for an AnalyticDB for MySQL cluster.
//
// @param request - DescribeInclinedTablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInclinedTablesResponse
func (client *Client) DescribeInclinedTablesWithOptions(request *DescribeInclinedTablesRequest, runtime *dara.RuntimeOptions) (_result *DescribeInclinedTablesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
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
		Action:      dara.String("DescribeInclinedTables"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInclinedTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about skewed tables for an AnalyticDB for MySQL cluster.
//
// @param request - DescribeInclinedTablesRequest
//
// @return DescribeInclinedTablesResponse
func (client *Client) DescribeInclinedTables(request *DescribeInclinedTablesRequest) (_result *DescribeInclinedTablesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInclinedTablesResponse{}
	_body, _err := client.DescribeInclinedTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取作业资源使用统计
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeJobResourceUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeJobResourceUsageResponse
func (client *Client) DescribeJobResourceUsageWithOptions(request *DescribeJobResourceUsageRequest, runtime *dara.RuntimeOptions) (_result *DescribeJobResourceUsageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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
		Action:      dara.String("DescribeJobResourceUsage"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeJobResourceUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取作业资源使用统计
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeJobResourceUsageRequest
//
// @return DescribeJobResourceUsageResponse
func (client *Client) DescribeJobResourceUsage(request *DescribeJobResourceUsageRequest) (_result *DescribeJobResourceUsageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeJobResourceUsageResponse{}
	_body, _err := client.DescribeJobResourceUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the minor version of an AnalyticDB for MySQL cluster.
//
// @param request - DescribeKernelVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeKernelVersionResponse
func (client *Client) DescribeKernelVersionWithOptions(request *DescribeKernelVersionRequest, runtime *dara.RuntimeOptions) (_result *DescribeKernelVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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
		Action:      dara.String("DescribeKernelVersion"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeKernelVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the minor version of an AnalyticDB for MySQL cluster.
//
// @param request - DescribeKernelVersionRequest
//
// @return DescribeKernelVersionResponse
func (client *Client) DescribeKernelVersion(request *DescribeKernelVersionRequest) (_result *DescribeKernelVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeKernelVersionResponse{}
	_body, _err := client.DescribeKernelVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the answer by a large language model (LLM) to a user question about the use of AnalyticDB for MySQL.
//
// @param request - DescribeLLMAnswerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLLMAnswerResponse
func (client *Client) DescribeLLMAnswerWithSSE(request *DescribeLLMAnswerRequest, runtime *dara.RuntimeOptions, _yield chan *DescribeLLMAnswerResponse, _yieldErr chan error) {
	defer close(_yield)
	client.describeLLMAnswerWithSSE_opYieldFunc(_yield, _yieldErr, request, runtime)
	return
}

// Summary:
//
// Queries the answer by a large language model (LLM) to a user question about the use of AnalyticDB for MySQL.
//
// @param request - DescribeLLMAnswerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLLMAnswerResponse
func (client *Client) DescribeLLMAnswerWithOptions(request *DescribeLLMAnswerRequest, runtime *dara.RuntimeOptions) (_result *DescribeLLMAnswerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
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
		Action:      dara.String("DescribeLLMAnswer"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLLMAnswerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the answer by a large language model (LLM) to a user question about the use of AnalyticDB for MySQL.
//
// @param request - DescribeLLMAnswerRequest
//
// @return DescribeLLMAnswerResponse
func (client *Client) DescribeLLMAnswer(request *DescribeLLMAnswerRequest) (_result *DescribeLLMAnswerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLLMAnswerResponse{}
	_body, _err := client.DescribeLLMAnswerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of questions similar to a user question.
//
// @param request - DescribeLLMSimilarQuestionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLLMSimilarQuestionsResponse
func (client *Client) DescribeLLMSimilarQuestionsWithOptions(request *DescribeLLMSimilarQuestionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeLLMSimilarQuestionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
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
		Action:      dara.String("DescribeLLMSimilarQuestions"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLLMSimilarQuestionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of questions similar to a user question.
//
// @param request - DescribeLLMSimilarQuestionsRequest
//
// @return DescribeLLMSimilarQuestionsResponse
func (client *Client) DescribeLLMSimilarQuestions(request *DescribeLLMSimilarQuestionsRequest) (_result *DescribeLLMSimilarQuestionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLLMSimilarQuestionsResponse{}
	_body, _err := client.DescribeLLMSimilarQuestionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the lake cache size of an AnalyticDB for MySQL cluster.
//
// @param request - DescribeLakeCacheSizeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLakeCacheSizeResponse
func (client *Client) DescribeLakeCacheSizeWithOptions(request *DescribeLakeCacheSizeRequest, runtime *dara.RuntimeOptions) (_result *DescribeLakeCacheSizeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLakeCacheSize"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLakeCacheSizeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the lake cache size of an AnalyticDB for MySQL cluster.
//
// @param request - DescribeLakeCacheSizeRequest
//
// @return DescribeLakeCacheSizeResponse
func (client *Client) DescribeLakeCacheSize(request *DescribeLakeCacheSizeRequest) (_result *DescribeLakeCacheSizeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLakeCacheSizeResponse{}
	_body, _err := client.DescribeLakeCacheSizeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the result of a recommendation task for a materialized view.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeMVRecommendResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMVRecommendResultsResponse
func (client *Client) DescribeMVRecommendResultsWithOptions(request *DescribeMVRecommendResultsRequest, runtime *dara.RuntimeOptions) (_result *DescribeMVRecommendResultsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionInner) {
		query["ActionInner"] = request.ActionInner
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.From) {
		query["From"] = request.From
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SubQueryId) {
		query["SubQueryId"] = request.SubQueryId
	}

	if !dara.IsNil(request.SubtaskId) {
		query["SubtaskId"] = request.SubtaskId
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMVRecommendResults"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMVRecommendResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the result of a recommendation task for a materialized view.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeMVRecommendResultsRequest
//
// @return DescribeMVRecommendResultsResponse
func (client *Client) DescribeMVRecommendResults(request *DescribeMVRecommendResultsRequest) (_result *DescribeMVRecommendResultsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMVRecommendResultsResponse{}
	_body, _err := client.DescribeMVRecommendResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看物化视图子任务
//
// @param request - DescribeMvRecommendSubTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMvRecommendSubTasksResponse
func (client *Client) DescribeMvRecommendSubTasksWithOptions(request *DescribeMvRecommendSubTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeMvRecommendSubTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionInner) {
		query["ActionInner"] = request.ActionInner
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.From) {
		query["From"] = request.From
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SubtaskId) {
		query["SubtaskId"] = request.SubtaskId
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMvRecommendSubTasks"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMvRecommendSubTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看物化视图子任务
//
// @param request - DescribeMvRecommendSubTasksRequest
//
// @return DescribeMvRecommendSubTasksResponse
func (client *Client) DescribeMvRecommendSubTasks(request *DescribeMvRecommendSubTasksRequest) (_result *DescribeMvRecommendSubTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMvRecommendSubTasksResponse{}
	_body, _err := client.DescribeMvRecommendSubTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看物化视图推荐任务
//
// @param request - DescribeMvRecommendTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMvRecommendTasksResponse
func (client *Client) DescribeMvRecommendTasksWithOptions(request *DescribeMvRecommendTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeMvRecommendTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionInner) {
		query["ActionInner"] = request.ActionInner
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.From) {
		query["From"] = request.From
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

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMvRecommendTasks"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMvRecommendTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看物化视图推荐任务
//
// @param request - DescribeMvRecommendTasksRequest
//
// @return DescribeMvRecommendTasksResponse
func (client *Client) DescribeMvRecommendTasks(request *DescribeMvRecommendTasksRequest) (_result *DescribeMvRecommendTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMvRecommendTasksResponse{}
	_body, _err := client.DescribeMvRecommendTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the service account permissions of an AnalyticDB for MySQL cluster.
//
// @param request - DescribeOperatorPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOperatorPermissionResponse
func (client *Client) DescribeOperatorPermissionWithOptions(request *DescribeOperatorPermissionRequest, runtime *dara.RuntimeOptions) (_result *DescribeOperatorPermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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
		Action:      dara.String("DescribeOperatorPermission"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOperatorPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the service account permissions of an AnalyticDB for MySQL cluster.
//
// @param request - DescribeOperatorPermissionRequest
//
// @return DescribeOperatorPermissionResponse
func (client *Client) DescribeOperatorPermission(request *DescribeOperatorPermissionRequest) (_result *DescribeOperatorPermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeOperatorPermissionResponse{}
	_body, _err := client.DescribeOperatorPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about oversized non-partitioned tables in an AnalyticDB for MySQL cluster.
//
// @param request - DescribeOversizeNonPartitionTableInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOversizeNonPartitionTableInfosResponse
func (client *Client) DescribeOversizeNonPartitionTableInfosWithOptions(request *DescribeOversizeNonPartitionTableInfosRequest, runtime *dara.RuntimeOptions) (_result *DescribeOversizeNonPartitionTableInfosResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
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
		Action:      dara.String("DescribeOversizeNonPartitionTableInfos"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOversizeNonPartitionTableInfosResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about oversized non-partitioned tables in an AnalyticDB for MySQL cluster.
//
// @param request - DescribeOversizeNonPartitionTableInfosRequest
//
// @return DescribeOversizeNonPartitionTableInfosResponse
func (client *Client) DescribeOversizeNonPartitionTableInfos(request *DescribeOversizeNonPartitionTableInfosRequest) (_result *DescribeOversizeNonPartitionTableInfosResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeOversizeNonPartitionTableInfosResponse{}
	_body, _err := client.DescribeOversizeNonPartitionTableInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about performance metrics of an SQL pattern such as the query duration and average memory usage for an AnalyticDB for MySQL cluster within a time range.
//
// Description:
//
//	  General endpoint: `adb.aliyuncs.com`.
//
//		- Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - DescribePatternPerformanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePatternPerformanceResponse
func (client *Client) DescribePatternPerformanceWithOptions(request *DescribePatternPerformanceRequest, runtime *dara.RuntimeOptions) (_result *DescribePatternPerformanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PatternId) {
		query["PatternId"] = request.PatternId
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
		Action:      dara.String("DescribePatternPerformance"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePatternPerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about performance metrics of an SQL pattern such as the query duration and average memory usage for an AnalyticDB for MySQL cluster within a time range.
//
// Description:
//
//	  General endpoint: `adb.aliyuncs.com`.
//
//		- Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - DescribePatternPerformanceRequest
//
// @return DescribePatternPerformanceResponse
func (client *Client) DescribePatternPerformance(request *DescribePatternPerformanceRequest) (_result *DescribePatternPerformanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePatternPerformanceResponse{}
	_body, _err := client.DescribePatternPerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a monitoring view.
//
// @param request - DescribePerformanceViewAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePerformanceViewAttributeResponse
func (client *Client) DescribePerformanceViewAttributeWithOptions(request *DescribePerformanceViewAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribePerformanceViewAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.ViewName) {
		query["ViewName"] = request.ViewName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePerformanceViewAttribute"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePerformanceViewAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a monitoring view.
//
// @param request - DescribePerformanceViewAttributeRequest
//
// @return DescribePerformanceViewAttributeResponse
func (client *Client) DescribePerformanceViewAttribute(request *DescribePerformanceViewAttributeRequest) (_result *DescribePerformanceViewAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePerformanceViewAttributeResponse{}
	_body, _err := client.DescribePerformanceViewAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of monitoring views.
//
// @param request - DescribePerformanceViewsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePerformanceViewsResponse
func (client *Client) DescribePerformanceViewsWithOptions(request *DescribePerformanceViewsRequest, runtime *dara.RuntimeOptions) (_result *DescribePerformanceViewsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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
		Action:      dara.String("DescribePerformanceViews"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePerformanceViewsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of monitoring views.
//
// @param request - DescribePerformanceViewsRequest
//
// @return DescribePerformanceViewsResponse
func (client *Client) DescribePerformanceViews(request *DescribePerformanceViewsRequest) (_result *DescribePerformanceViewsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePerformanceViewsResponse{}
	_body, _err := client.DescribePerformanceViewsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of regions and zones in which AnalyticDB for MySQL Data Lakehouse Edition (V3.0) is available.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
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
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2021-12-01"),
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
// Queries a list of regions and zones in which AnalyticDB for MySQL Data Lakehouse Edition (V3.0) is available.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
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
// Queries the information about resource group specifications.
//
// Description:
//
// ### [](#)
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeResourceGroupSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResourceGroupSpecResponse
func (client *Client) DescribeResourceGroupSpecWithOptions(request *DescribeResourceGroupSpecRequest, runtime *dara.RuntimeOptions) (_result *DescribeResourceGroupSpecResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupType) {
		query["ResourceGroupType"] = request.ResourceGroupType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeResourceGroupSpec"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeResourceGroupSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about resource group specifications.
//
// Description:
//
// ### [](#)
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeResourceGroupSpecRequest
//
// @return DescribeResourceGroupSpecResponse
func (client *Client) DescribeResourceGroupSpec(request *DescribeResourceGroupSpecRequest) (_result *DescribeResourceGroupSpecResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeResourceGroupSpecResponse{}
	_body, _err := client.DescribeResourceGroupSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the user-configured result set export settings.
//
// @param request - DescribeResultExportConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResultExportConfigResponse
func (client *Client) DescribeResultExportConfigWithOptions(request *DescribeResultExportConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeResultExportConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.ExportType) {
		query["ExportType"] = request.ExportType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeResultExportConfig"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeResultExportConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the user-configured result set export settings.
//
// @param request - DescribeResultExportConfigRequest
//
// @return DescribeResultExportConfigResponse
func (client *Client) DescribeResultExportConfig(request *DescribeResultExportConfigRequest) (_result *DescribeResultExportConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeResultExportConfigResponse{}
	_body, _err := client.DescribeResultExportConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of SQL patterns for an AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster within a time range.
//
// Description:
//
//	  General endpoint: `adb.aliyuncs.com`.
//
//		- Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - DescribeSQLPatternsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSQLPatternsResponse
func (client *Client) DescribeSQLPatternsWithOptions(request *DescribeSQLPatternsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSQLPatternsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSQLPatterns"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSQLPatternsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of SQL patterns for an AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster within a time range.
//
// Description:
//
//	  General endpoint: `adb.aliyuncs.com`.
//
//		- Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - DescribeSQLPatternsRequest
//
// @return DescribeSQLPatternsResponse
func (client *Client) DescribeSQLPatterns(request *DescribeSQLPatternsRequest) (_result *DescribeSQLPatternsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSQLPatternsResponse{}
	_body, _err := client.DescribeSQLPatternsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the WebSocket domain name of an AnalyticDB for MySQL cluster.
//
// @param request - DescribeSQLWebSocketDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSQLWebSocketDomainResponse
func (client *Client) DescribeSQLWebSocketDomainWithOptions(request *DescribeSQLWebSocketDomainRequest, runtime *dara.RuntimeOptions) (_result *DescribeSQLWebSocketDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Module) {
		query["Module"] = request.Module
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSQLWebSocketDomain"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSQLWebSocketDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the WebSocket domain name of an AnalyticDB for MySQL cluster.
//
// @param request - DescribeSQLWebSocketDomainRequest
//
// @return DescribeSQLWebSocketDomainResponse
func (client *Client) DescribeSQLWebSocketDomain(request *DescribeSQLWebSocketDomainRequest) (_result *DescribeSQLWebSocketDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSQLWebSocketDomainResponse{}
	_body, _err := client.DescribeSQLWebSocketDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of databases in an AnalyticDB for MySQL cluster.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - DescribeSchemasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSchemasResponse
func (client *Client) DescribeSchemasWithOptions(request *DescribeSchemasRequest, runtime *dara.RuntimeOptions) (_result *DescribeSchemasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSchemas"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSchemasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of databases in an AnalyticDB for MySQL cluster.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - DescribeSchemasRequest
//
// @return DescribeSchemasResponse
func (client *Client) DescribeSchemas(request *DescribeSchemasRequest) (_result *DescribeSchemasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSchemasResponse{}
	_body, _err := client.DescribeSchemasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the diagnostic information about a Spark application.
//
// @param request - DescribeSparkAppDiagnosisInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSparkAppDiagnosisInfoResponse
func (client *Client) DescribeSparkAppDiagnosisInfoWithOptions(request *DescribeSparkAppDiagnosisInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeSparkAppDiagnosisInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSparkAppDiagnosisInfo"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSparkAppDiagnosisInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the diagnostic information about a Spark application.
//
// @param request - DescribeSparkAppDiagnosisInfoRequest
//
// @return DescribeSparkAppDiagnosisInfoResponse
func (client *Client) DescribeSparkAppDiagnosisInfo(request *DescribeSparkAppDiagnosisInfoRequest) (_result *DescribeSparkAppDiagnosisInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSparkAppDiagnosisInfoResponse{}
	_body, _err := client.DescribeSparkAppDiagnosisInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the type of a Spark application.
//
// @param request - DescribeSparkAppTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSparkAppTypeResponse
func (client *Client) DescribeSparkAppTypeWithOptions(request *DescribeSparkAppTypeRequest, runtime *dara.RuntimeOptions) (_result *DescribeSparkAppTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSparkAppType"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSparkAppTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the type of a Spark application.
//
// @param request - DescribeSparkAppTypeRequest
//
// @return DescribeSparkAppTypeResponse
func (client *Client) DescribeSparkAppType(request *DescribeSparkAppTypeRequest) (_result *DescribeSparkAppTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSparkAppTypeResponse{}
	_body, _err := client.DescribeSparkAppTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the SQL audit logs for a Spark Interactive resource group.
//
// Description:
//
// SQL audit logs can be queried only when SQL audit is enabled. Only SQL audit logs within the last 30 days can be queried. If SQL auditing is turned off midway, when it is re-enabled, you can only query the SQL audit logs generated after it was turned back on.
//
// >  You can query only SQL audit logs that are executed by using Spark Interactive Resource Group.
//
// @param request - DescribeSparkAuditLogRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSparkAuditLogRecordsResponse
func (client *Client) DescribeSparkAuditLogRecordsWithOptions(request *DescribeSparkAuditLogRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSparkAuditLogRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientIp) {
		query["ClientIp"] = request.ClientIp
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
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

	if !dara.IsNil(request.ProcessId) {
		query["ProcessId"] = request.ProcessId
	}

	if !dara.IsNil(request.ProxyUser) {
		query["ProxyUser"] = request.ProxyUser
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupName) {
		query["ResourceGroupName"] = request.ResourceGroupName
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SQLText) {
		query["SQLText"] = request.SQLText
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StatementId) {
		query["StatementId"] = request.StatementId
	}

	if !dara.IsNil(request.StatementSource) {
		query["StatementSource"] = request.StatementSource
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TotalTime) {
		query["TotalTime"] = request.TotalTime
	}

	if !dara.IsNil(request.User) {
		query["User"] = request.User
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSparkAuditLogRecords"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSparkAuditLogRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the SQL audit logs for a Spark Interactive resource group.
//
// Description:
//
// SQL audit logs can be queried only when SQL audit is enabled. Only SQL audit logs within the last 30 days can be queried. If SQL auditing is turned off midway, when it is re-enabled, you can only query the SQL audit logs generated after it was turned back on.
//
// >  You can query only SQL audit logs that are executed by using Spark Interactive Resource Group.
//
// @param request - DescribeSparkAuditLogRecordsRequest
//
// @return DescribeSparkAuditLogRecordsResponse
func (client *Client) DescribeSparkAuditLogRecords(request *DescribeSparkAuditLogRecordsRequest) (_result *DescribeSparkAuditLogRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSparkAuditLogRecordsResponse{}
	_body, _err := client.DescribeSparkAuditLogRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the execution logs of Spark code.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - DescribeSparkCodeLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSparkCodeLogResponse
func (client *Client) DescribeSparkCodeLogWithOptions(request *DescribeSparkCodeLogRequest, runtime *dara.RuntimeOptions) (_result *DescribeSparkCodeLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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
		Action:      dara.String("DescribeSparkCodeLog"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSparkCodeLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution logs of Spark code.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - DescribeSparkCodeLogRequest
//
// @return DescribeSparkCodeLogResponse
func (client *Client) DescribeSparkCodeLog(request *DescribeSparkCodeLogRequest) (_result *DescribeSparkCodeLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSparkCodeLogResponse{}
	_body, _err := client.DescribeSparkCodeLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the execution result of Spark code.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - DescribeSparkCodeOutputRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSparkCodeOutputResponse
func (client *Client) DescribeSparkCodeOutputWithOptions(request *DescribeSparkCodeOutputRequest, runtime *dara.RuntimeOptions) (_result *DescribeSparkCodeOutputResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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
		Action:      dara.String("DescribeSparkCodeOutput"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSparkCodeOutputResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution result of Spark code.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - DescribeSparkCodeOutputRequest
//
// @return DescribeSparkCodeOutputResponse
func (client *Client) DescribeSparkCodeOutput(request *DescribeSparkCodeOutputRequest) (_result *DescribeSparkCodeOutputResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSparkCodeOutputResponse{}
	_body, _err := client.DescribeSparkCodeOutputWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the URL of the web UI for a Spark application.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - DescribeSparkCodeWebUiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSparkCodeWebUiResponse
func (client *Client) DescribeSparkCodeWebUiWithOptions(request *DescribeSparkCodeWebUiRequest, runtime *dara.RuntimeOptions) (_result *DescribeSparkCodeWebUiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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
		Action:      dara.String("DescribeSparkCodeWebUi"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSparkCodeWebUiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the URL of the web UI for a Spark application.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - DescribeSparkCodeWebUiRequest
//
// @return DescribeSparkCodeWebUiResponse
func (client *Client) DescribeSparkCodeWebUi(request *DescribeSparkCodeWebUiRequest) (_result *DescribeSparkCodeWebUiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSparkCodeWebUiResponse{}
	_body, _err := client.DescribeSparkCodeWebUiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the diagnostic information about a Spark SQL query.
//
// @param request - DescribeSparkSQLDiagnosisAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSparkSQLDiagnosisAttributeResponse
func (client *Client) DescribeSparkSQLDiagnosisAttributeWithOptions(request *DescribeSparkSQLDiagnosisAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeSparkSQLDiagnosisAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.InnerQueryId) {
		query["InnerQueryId"] = request.InnerQueryId
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSparkSQLDiagnosisAttribute"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSparkSQLDiagnosisAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the diagnostic information about a Spark SQL query.
//
// @param request - DescribeSparkSQLDiagnosisAttributeRequest
//
// @return DescribeSparkSQLDiagnosisAttributeResponse
func (client *Client) DescribeSparkSQLDiagnosisAttribute(request *DescribeSparkSQLDiagnosisAttributeRequest) (_result *DescribeSparkSQLDiagnosisAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSparkSQLDiagnosisAttributeResponse{}
	_body, _err := client.DescribeSparkSQLDiagnosisAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the diagnostic information about Spark SQL queries.
//
// @param request - DescribeSparkSQLDiagnosisListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSparkSQLDiagnosisListResponse
func (client *Client) DescribeSparkSQLDiagnosisListWithOptions(request *DescribeSparkSQLDiagnosisListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSparkSQLDiagnosisListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.MaxStartTime) {
		query["MaxStartTime"] = request.MaxStartTime
	}

	if !dara.IsNil(request.MinStartTime) {
		query["MinStartTime"] = request.MinStartTime
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StatementId) {
		query["StatementId"] = request.StatementId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSparkSQLDiagnosisList"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSparkSQLDiagnosisListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the diagnostic information about Spark SQL queries.
//
// @param request - DescribeSparkSQLDiagnosisListRequest
//
// @return DescribeSparkSQLDiagnosisListResponse
func (client *Client) DescribeSparkSQLDiagnosisList(request *DescribeSparkSQLDiagnosisListRequest) (_result *DescribeSparkSQLDiagnosisListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSparkSQLDiagnosisListResponse{}
	_body, _err := client.DescribeSparkSQLDiagnosisListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about SQL patterns of an AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster within a time range.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeSqlPatternRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSqlPatternResponse
func (client *Client) DescribeSqlPatternWithOptions(request *DescribeSqlPatternRequest, runtime *dara.RuntimeOptions) (_result *DescribeSqlPatternResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SqlPattern) {
		query["SqlPattern"] = request.SqlPattern
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
		Action:      dara.String("DescribeSqlPattern"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSqlPatternResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about SQL patterns of an AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster within a time range.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeSqlPatternRequest
//
// @return DescribeSqlPatternResponse
func (client *Client) DescribeSqlPattern(request *DescribeSqlPatternRequest) (_result *DescribeSqlPatternResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSqlPatternResponse{}
	_body, _err := client.DescribeSqlPatternWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the storage resource usage of an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeStorageResourceUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeStorageResourceUsageResponse
func (client *Client) DescribeStorageResourceUsageWithOptions(request *DescribeStorageResourceUsageRequest, runtime *dara.RuntimeOptions) (_result *DescribeStorageResourceUsageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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
		Action:      dara.String("DescribeStorageResourceUsage"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeStorageResourceUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the storage resource usage of an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeStorageResourceUsageRequest
//
// @return DescribeStorageResourceUsageResponse
func (client *Client) DescribeStorageResourceUsage(request *DescribeStorageResourceUsageRequest) (_result *DescribeStorageResourceUsageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeStorageResourceUsageResponse{}
	_body, _err := client.DescribeStorageResourceUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of accesses to a table or all tables in an AnalyticDB for MySQL cluster on a date.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeTableAccessCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTableAccessCountResponse
func (client *Client) DescribeTableAccessCountWithOptions(request *DescribeTableAccessCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeTableAccessCountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTableAccessCount"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTableAccessCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of accesses to a table or all tables in an AnalyticDB for MySQL cluster on a date.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeTableAccessCountRequest
//
// @return DescribeTableAccessCountResponse
func (client *Client) DescribeTableAccessCount(request *DescribeTableAccessCountRequest) (_result *DescribeTableAccessCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTableAccessCountResponse{}
	_body, _err := client.DescribeTableAccessCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about data distribution among shards of a table.
//
// @param request - DescribeTableDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTableDetailResponse
func (client *Client) DescribeTableDetailWithOptions(request *DescribeTableDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeTableDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SchemaName) {
		query["SchemaName"] = request.SchemaName
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTableDetail"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTableDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about data distribution among shards of a table.
//
// @param request - DescribeTableDetailRequest
//
// @return DescribeTableDetailResponse
func (client *Client) DescribeTableDetail(request *DescribeTableDetailRequest) (_result *DescribeTableDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTableDetailResponse{}
	_body, _err := client.DescribeTableDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about partition diagnostics.
//
// @param request - DescribeTablePartitionDiagnoseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTablePartitionDiagnoseResponse
func (client *Client) DescribeTablePartitionDiagnoseWithOptions(request *DescribeTablePartitionDiagnoseRequest, runtime *dara.RuntimeOptions) (_result *DescribeTablePartitionDiagnoseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTablePartitionDiagnose"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTablePartitionDiagnoseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about partition diagnostics.
//
// @param request - DescribeTablePartitionDiagnoseRequest
//
// @return DescribeTablePartitionDiagnoseResponse
func (client *Client) DescribeTablePartitionDiagnose(request *DescribeTablePartitionDiagnoseRequest) (_result *DescribeTablePartitionDiagnoseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTablePartitionDiagnoseResponse{}
	_body, _err := client.DescribeTablePartitionDiagnoseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the table statistics of an AnalyticDB for MySQL cluster.
//
// @param request - DescribeTableStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTableStatisticsResponse
func (client *Client) DescribeTableStatisticsWithOptions(request *DescribeTableStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeTableStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SchemaName) {
		query["SchemaName"] = request.SchemaName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTableStatistics"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTableStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the table statistics of an AnalyticDB for MySQL cluster.
//
// @param request - DescribeTableStatisticsRequest
//
// @return DescribeTableStatisticsResponse
func (client *Client) DescribeTableStatistics(request *DescribeTableStatisticsRequest) (_result *DescribeTableStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTableStatisticsResponse{}
	_body, _err := client.DescribeTableStatisticsWithOptions(request, runtime)
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
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - DescribeTablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTablesResponse
func (client *Client) DescribeTablesWithOptions(request *DescribeTablesRequest, runtime *dara.RuntimeOptions) (_result *DescribeTablesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SchemaName) {
		query["SchemaName"] = request.SchemaName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTables"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTablesResponse{}
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
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - DescribeTablesRequest
//
// @return DescribeTablesResponse
func (client *Client) DescribeTables(request *DescribeTablesRequest) (_result *DescribeTablesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTablesResponse{}
	_body, _err := client.DescribeTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries available quotas.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeUserQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserQuotaResponse
func (client *Client) DescribeUserQuotaWithOptions(request *DescribeUserQuotaRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserQuotaResponse, _err error) {
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
		Action:      dara.String("DescribeUserQuota"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries available quotas.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DescribeUserQuotaRequest
//
// @return DescribeUserQuotaResponse
func (client *Client) DescribeUserQuota(request *DescribeUserQuotaRequest) (_result *DescribeUserQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserQuotaResponse{}
	_body, _err := client.DescribeUserQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves materialized view refresh tasks.
//
// @param request - DescribeViewJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeViewJobsResponse
func (client *Client) DescribeViewJobsWithOptions(request *DescribeViewJobsRequest, runtime *dara.RuntimeOptions) (_result *DescribeViewJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.FilterOwner) {
		query["FilterOwner"] = request.FilterOwner
	}

	if !dara.IsNil(request.FilterViewName) {
		query["FilterViewName"] = request.FilterViewName
	}

	if !dara.IsNil(request.FilterViewType) {
		query["FilterViewType"] = request.FilterViewType
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SchemaName) {
		query["SchemaName"] = request.SchemaName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeViewJobs"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeViewJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves materialized view refresh tasks.
//
// @param request - DescribeViewJobsRequest
//
// @return DescribeViewJobsResponse
func (client *Client) DescribeViewJobs(request *DescribeViewJobsRequest) (_result *DescribeViewJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeViewJobsResponse{}
	_body, _err := client.DescribeViewJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the diagnostic results of the storage layer.
//
// @param request - DescribeWorkerDetectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWorkerDetectionResponse
func (client *Client) DescribeWorkerDetectionWithOptions(request *DescribeWorkerDetectionRequest, runtime *dara.RuntimeOptions) (_result *DescribeWorkerDetectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWorkerDetection"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWorkerDetectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the diagnostic results of the storage layer.
//
// @param request - DescribeWorkerDetectionRequest
//
// @return DescribeWorkerDetectionResponse
func (client *Client) DescribeWorkerDetection(request *DescribeWorkerDetectionRequest) (_result *DescribeWorkerDetectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeWorkerDetectionResponse{}
	_body, _err := client.DescribeWorkerDetectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 解绑用户弹性网卡
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DetachUserENIRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachUserENIResponse
func (client *Client) DetachUserENIWithOptions(request *DetachUserENIRequest, runtime *dara.RuntimeOptions) (_result *DetachUserENIResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachUserENI"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachUserENIResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解绑用户弹性网卡
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DetachUserENIRequest
//
// @return DetachUserENIResponse
func (client *Client) DetachUserENI(request *DetachUserENIRequest) (_result *DetachUserENIResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachUserENIResponse{}
	_body, _err := client.DetachUserENIWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables the suggestion feature.
//
// @param request - DisableAdviceServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableAdviceServiceResponse
func (client *Client) DisableAdviceServiceWithOptions(request *DisableAdviceServiceRequest, runtime *dara.RuntimeOptions) (_result *DisableAdviceServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableAdviceService"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableAdviceServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the suggestion feature.
//
// @param request - DisableAdviceServiceRequest
//
// @return DisableAdviceServiceResponse
func (client *Client) DisableAdviceService(request *DisableAdviceServiceRequest) (_result *DisableAdviceServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableAdviceServiceResponse{}
	_body, _err := client.DisableAdviceServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables a scaling plan for an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DisableElasticPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableElasticPlanResponse
func (client *Client) DisableElasticPlanWithOptions(request *DisableElasticPlanRequest, runtime *dara.RuntimeOptions) (_result *DisableElasticPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.ElasticPlanName) {
		query["ElasticPlanName"] = request.ElasticPlanName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableElasticPlan"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableElasticPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a scaling plan for an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DisableElasticPlanRequest
//
// @return DisableElasticPlanResponse
func (client *Client) DisableElasticPlan(request *DisableElasticPlanRequest) (_result *DisableElasticPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableElasticPlanResponse{}
	_body, _err := client.DisableElasticPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Downloads the diagnostic information about SQL statements that meet a query condition for an AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - DownloadDiagnosisRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DownloadDiagnosisRecordsResponse
func (client *Client) DownloadDiagnosisRecordsWithOptions(request *DownloadDiagnosisRecordsRequest, runtime *dara.RuntimeOptions) (_result *DownloadDiagnosisRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientIp) {
		query["ClientIp"] = request.ClientIp
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MaxPeakMemory) {
		query["MaxPeakMemory"] = request.MaxPeakMemory
	}

	if !dara.IsNil(request.MaxScanSize) {
		query["MaxScanSize"] = request.MaxScanSize
	}

	if !dara.IsNil(request.MinPeakMemory) {
		query["MinPeakMemory"] = request.MinPeakMemory
	}

	if !dara.IsNil(request.MinScanSize) {
		query["MinScanSize"] = request.MinScanSize
	}

	if !dara.IsNil(request.QueryCondition) {
		query["QueryCondition"] = request.QueryCondition
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroup) {
		query["ResourceGroup"] = request.ResourceGroup
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DownloadDiagnosisRecords"),
		Version:     dara.String("2021-12-01"),
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
// Downloads the diagnostic information about SQL statements that meet a query condition for an AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
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
// Queries the Object Storage Service (OSS) URL of the downloaded certificate authority (CA) certificate that is used to connect to the wide table engine.
//
// @param request - DownloadInstanceCACertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DownloadInstanceCACertificateResponse
func (client *Client) DownloadInstanceCACertificateWithOptions(request *DownloadInstanceCACertificateRequest, runtime *dara.RuntimeOptions) (_result *DownloadInstanceCACertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
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
		Action:      dara.String("DownloadInstanceCACertificate"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DownloadInstanceCACertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Object Storage Service (OSS) URL of the downloaded certificate authority (CA) certificate that is used to connect to the wide table engine.
//
// @param request - DownloadInstanceCACertificateRequest
//
// @return DownloadInstanceCACertificateResponse
func (client *Client) DownloadInstanceCACertificate(request *DownloadInstanceCACertificateRequest) (_result *DownloadInstanceCACertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DownloadInstanceCACertificateResponse{}
	_body, _err := client.DownloadInstanceCACertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the suggestion feature.
//
// @param request - EnableAdviceServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableAdviceServiceResponse
func (client *Client) EnableAdviceServiceWithOptions(request *EnableAdviceServiceRequest, runtime *dara.RuntimeOptions) (_result *EnableAdviceServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableAdviceService"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableAdviceServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the suggestion feature.
//
// @param request - EnableAdviceServiceRequest
//
// @return EnableAdviceServiceResponse
func (client *Client) EnableAdviceService(request *EnableAdviceServiceRequest) (_result *EnableAdviceServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableAdviceServiceResponse{}
	_body, _err := client.EnableAdviceServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables a scaling plan for an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - EnableElasticPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableElasticPlanResponse
func (client *Client) EnableElasticPlanWithOptions(request *EnableElasticPlanRequest, runtime *dara.RuntimeOptions) (_result *EnableElasticPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.ElasticPlanName) {
		query["ElasticPlanName"] = request.ElasticPlanName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableElasticPlan"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableElasticPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a scaling plan for an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - EnableElasticPlanRequest
//
// @return EnableElasticPlanResponse
func (client *Client) EnableElasticPlan(request *EnableElasticPlanRequest) (_result *EnableElasticPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableElasticPlanResponse{}
	_body, _err := client.EnableElasticPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Executes part of the code in a Spark job.
//
// @param request - ExecuteSparkReplStatementRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteSparkReplStatementResponse
func (client *Client) ExecuteSparkReplStatementWithOptions(request *ExecuteSparkReplStatementRequest, runtime *dara.RuntimeOptions) (_result *ExecuteSparkReplStatementResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Code) {
		body["Code"] = request.Code
	}

	if !dara.IsNil(request.CodeType) {
		body["CodeType"] = request.CodeType
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteSparkReplStatement"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteSparkReplStatementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Executes part of the code in a Spark job.
//
// @param request - ExecuteSparkReplStatementRequest
//
// @return ExecuteSparkReplStatementResponse
func (client *Client) ExecuteSparkReplStatement(request *ExecuteSparkReplStatementRequest) (_result *ExecuteSparkReplStatementResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExecuteSparkReplStatementResponse{}
	_body, _err := client.ExecuteSparkReplStatementWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Executes Spark SQL statements in batches.
//
// @param request - ExecuteSparkWarehouseBatchSQLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteSparkWarehouseBatchSQLResponse
func (client *Client) ExecuteSparkWarehouseBatchSQLWithOptions(request *ExecuteSparkWarehouseBatchSQLRequest, runtime *dara.RuntimeOptions) (_result *ExecuteSparkWarehouseBatchSQLResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Agency) {
		body["Agency"] = request.Agency
	}

	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.ExecuteResultLimit) {
		body["ExecuteResultLimit"] = request.ExecuteResultLimit
	}

	if !dara.IsNil(request.ExecuteTimeLimitInSeconds) {
		body["ExecuteTimeLimitInSeconds"] = request.ExecuteTimeLimitInSeconds
	}

	if !dara.IsNil(request.Query) {
		body["Query"] = request.Query
	}

	if !dara.IsNil(request.ResourceGroupName) {
		body["ResourceGroupName"] = request.ResourceGroupName
	}

	if !dara.IsNil(request.RuntimeConfig) {
		body["RuntimeConfig"] = request.RuntimeConfig
	}

	if !dara.IsNil(request.Schema) {
		body["Schema"] = request.Schema
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteSparkWarehouseBatchSQL"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteSparkWarehouseBatchSQLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Executes Spark SQL statements in batches.
//
// @param request - ExecuteSparkWarehouseBatchSQLRequest
//
// @return ExecuteSparkWarehouseBatchSQLResponse
func (client *Client) ExecuteSparkWarehouseBatchSQL(request *ExecuteSparkWarehouseBatchSQLRequest) (_result *ExecuteSparkWarehouseBatchSQLResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExecuteSparkWarehouseBatchSQLResponse{}
	_body, _err := client.ExecuteSparkWarehouseBatchSQLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ExistRunningSQLEngine is deprecated
//
// Summary:
//
// Queries whether a running SQL engine exists.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - ExistRunningSQLEngineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExistRunningSQLEngineResponse
func (client *Client) ExistRunningSQLEngineWithOptions(request *ExistRunningSQLEngineRequest, runtime *dara.RuntimeOptions) (_result *ExistRunningSQLEngineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.ResourceGroupName) {
		body["ResourceGroupName"] = request.ResourceGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExistRunningSQLEngine"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExistRunningSQLEngineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ExistRunningSQLEngine is deprecated
//
// Summary:
//
// Queries whether a running SQL engine exists.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - ExistRunningSQLEngineRequest
//
// @return ExistRunningSQLEngineResponse
// Deprecated
func (client *Client) ExistRunningSQLEngine(request *ExistRunningSQLEngineRequest) (_result *ExistRunningSQLEngineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExistRunningSQLEngineResponse{}
	_body, _err := client.ExistRunningSQLEngineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// View the Spark basic permission diagnosis report of the current user.
//
// Description:
//
// The API diagnosis report contains whether the current user has all permissions required by the AnalyticDB for Spark related features. The scope of the permissions may exceed the minimum requirements of the business. The diagnostic report of the current API is used to quickly initialize the environment of AnalyticDB for Spark. If fine-grained permission configuration is needed, see [Configure fine-grained permissions in AnalyDB for Spark.](https://www.alibabacloud.com/help/zh/analyticdb/analyticdb-for-mysql/user-guide/create-the-aliyunadbsparkprocessingdatarole-role-for-a-ram-user-and-grant-permissions-to-the-role?spm=a2c63.p38356.help-menu-92664.d_2_5_0.48362a487dMzm9#section-y2z-ucd-1ko)
//
// @param request - GetADBSparkNecessaryRAMPermissionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetADBSparkNecessaryRAMPermissionsResponse
func (client *Client) GetADBSparkNecessaryRAMPermissionsWithOptions(request *GetADBSparkNecessaryRAMPermissionsRequest, runtime *dara.RuntimeOptions) (_result *GetADBSparkNecessaryRAMPermissionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetADBSparkNecessaryRAMPermissions"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetADBSparkNecessaryRAMPermissionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// View the Spark basic permission diagnosis report of the current user.
//
// Description:
//
// The API diagnosis report contains whether the current user has all permissions required by the AnalyticDB for Spark related features. The scope of the permissions may exceed the minimum requirements of the business. The diagnostic report of the current API is used to quickly initialize the environment of AnalyticDB for Spark. If fine-grained permission configuration is needed, see [Configure fine-grained permissions in AnalyDB for Spark.](https://www.alibabacloud.com/help/zh/analyticdb/analyticdb-for-mysql/user-guide/create-the-aliyunadbsparkprocessingdatarole-role-for-a-ram-user-and-grant-permissions-to-the-role?spm=a2c63.p38356.help-menu-92664.d_2_5_0.48362a487dMzm9#section-y2z-ucd-1ko)
//
// @param request - GetADBSparkNecessaryRAMPermissionsRequest
//
// @return GetADBSparkNecessaryRAMPermissionsResponse
func (client *Client) GetADBSparkNecessaryRAMPermissions(request *GetADBSparkNecessaryRAMPermissionsRequest) (_result *GetADBSparkNecessaryRAMPermissionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetADBSparkNecessaryRAMPermissionsResponse{}
	_body, _err := client.GetADBSparkNecessaryRAMPermissionsWithOptions(request, runtime)
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
// @param request - GetApsManagedDatabasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApsManagedDatabasesResponse
func (client *Client) GetApsManagedDatabasesWithOptions(request *GetApsManagedDatabasesRequest, runtime *dara.RuntimeOptions) (_result *GetApsManagedDatabasesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApsManagedDatabases"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApsManagedDatabasesResponse{}
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
// @param request - GetApsManagedDatabasesRequest
//
// @return GetApsManagedDatabasesResponse
func (client *Client) GetApsManagedDatabases(request *GetApsManagedDatabasesRequest) (_result *GetApsManagedDatabasesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetApsManagedDatabasesResponse{}
	_body, _err := client.GetApsManagedDatabasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the table creation statement for tables.
//
// @param request - GetCreateTableSQLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCreateTableSQLResponse
func (client *Client) GetCreateTableSQLWithOptions(request *GetCreateTableSQLRequest, runtime *dara.RuntimeOptions) (_result *GetCreateTableSQLResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.SchemaName) {
		query["SchemaName"] = request.SchemaName
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCreateTableSQL"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCreateTableSQLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the table creation statement for tables.
//
// @param request - GetCreateTableSQLRequest
//
// @return GetCreateTableSQLResponse
func (client *Client) GetCreateTableSQL(request *GetCreateTableSQLRequest) (_result *GetCreateTableSQLResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCreateTableSQLResponse{}
	_body, _err := client.GetCreateTableSQLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about databases.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - GetDatabaseObjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatabaseObjectsResponse
func (client *Client) GetDatabaseObjectsWithOptions(request *GetDatabaseObjectsRequest, runtime *dara.RuntimeOptions) (_result *GetDatabaseObjectsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.FilterOwner) {
		query["FilterOwner"] = request.FilterOwner
	}

	if !dara.IsNil(request.FilterSchemaName) {
		query["FilterSchemaName"] = request.FilterSchemaName
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDatabaseObjects"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDatabaseObjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about databases.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - GetDatabaseObjectsRequest
//
// @return GetDatabaseObjectsResponse
func (client *Client) GetDatabaseObjects(request *GetDatabaseObjectsRequest) (_result *GetDatabaseObjectsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDatabaseObjectsResponse{}
	_body, _err := client.GetDatabaseObjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a lake storage.
//
// @param request - GetLakeStorageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLakeStorageResponse
func (client *Client) GetLakeStorageWithOptions(request *GetLakeStorageRequest, runtime *dara.RuntimeOptions) (_result *GetLakeStorageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.LakeStorageId) {
		query["LakeStorageId"] = request.LakeStorageId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLakeStorage"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLakeStorageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a lake storage.
//
// @param request - GetLakeStorageRequest
//
// @return GetLakeStorageResponse
func (client *Client) GetLakeStorage(request *GetLakeStorageRequest) (_result *GetLakeStorageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLakeStorageResponse{}
	_body, _err := client.GetLakeStorageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the retry log of a Spark application.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - GetSparkAppAttemptLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSparkAppAttemptLogResponse
func (client *Client) GetSparkAppAttemptLogWithOptions(request *GetSparkAppAttemptLogRequest, runtime *dara.RuntimeOptions) (_result *GetSparkAppAttemptLogResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.AttemptId) {
		body["AttemptId"] = request.AttemptId
	}

	if !dara.IsNil(request.LogLength) {
		body["LogLength"] = request.LogLength
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSparkAppAttemptLog"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSparkAppAttemptLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the retry log of a Spark application.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - GetSparkAppAttemptLogRequest
//
// @return GetSparkAppAttemptLogResponse
func (client *Client) GetSparkAppAttemptLog(request *GetSparkAppAttemptLogRequest) (_result *GetSparkAppAttemptLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSparkAppAttemptLogResponse{}
	_body, _err := client.GetSparkAppAttemptLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an Spark application.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - GetSparkAppInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSparkAppInfoResponse
func (client *Client) GetSparkAppInfoWithOptions(request *GetSparkAppInfoRequest, runtime *dara.RuntimeOptions) (_result *GetSparkAppInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSparkAppInfo"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSparkAppInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an Spark application.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - GetSparkAppInfoRequest
//
// @return GetSparkAppInfoResponse
func (client *Client) GetSparkAppInfo(request *GetSparkAppInfoRequest) (_result *GetSparkAppInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSparkAppInfoResponse{}
	_body, _err := client.GetSparkAppInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the logs of a Spark application.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - GetSparkAppLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSparkAppLogResponse
func (client *Client) GetSparkAppLogWithOptions(request *GetSparkAppLogRequest, runtime *dara.RuntimeOptions) (_result *GetSparkAppLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.LogLength) {
		body["LogLength"] = request.LogLength
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSparkAppLog"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSparkAppLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the logs of a Spark application.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - GetSparkAppLogRequest
//
// @return GetSparkAppLogResponse
func (client *Client) GetSparkAppLog(request *GetSparkAppLogRequest) (_result *GetSparkAppLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSparkAppLogResponse{}
	_body, _err := client.GetSparkAppLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the metrics of a Spark application.
//
// Description:
//
//	  Before you call this operation, you must call the [PreloadSparkAppMetrics](https://help.aliyun.com/document_detail/612447.html) operation to preload the metrics of a Spark application.
//
//		- Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - GetSparkAppMetricsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSparkAppMetricsResponse
func (client *Client) GetSparkAppMetricsWithOptions(request *GetSparkAppMetricsRequest, runtime *dara.RuntimeOptions) (_result *GetSparkAppMetricsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSparkAppMetrics"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSparkAppMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the metrics of a Spark application.
//
// Description:
//
//	  Before you call this operation, you must call the [PreloadSparkAppMetrics](https://help.aliyun.com/document_detail/612447.html) operation to preload the metrics of a Spark application.
//
//		- Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - GetSparkAppMetricsRequest
//
// @return GetSparkAppMetricsResponse
func (client *Client) GetSparkAppMetrics(request *GetSparkAppMetricsRequest) (_result *GetSparkAppMetricsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSparkAppMetricsResponse{}
	_body, _err := client.GetSparkAppMetricsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of a Spark application.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - GetSparkAppStateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSparkAppStateResponse
func (client *Client) GetSparkAppStateWithOptions(request *GetSparkAppStateRequest, runtime *dara.RuntimeOptions) (_result *GetSparkAppStateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSparkAppState"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSparkAppStateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of a Spark application.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - GetSparkAppStateRequest
//
// @return GetSparkAppStateResponse
func (client *Client) GetSparkAppState(request *GetSparkAppStateRequest) (_result *GetSparkAppStateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSparkAppStateResponse{}
	_body, _err := client.GetSparkAppStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the URL of the web UI for a Spark application.
//
// Description:
//
//	  General endpoint: `adb.aliyuncs.com`.
//
//		- Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - GetSparkAppWebUiAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSparkAppWebUiAddressResponse
func (client *Client) GetSparkAppWebUiAddressWithOptions(request *GetSparkAppWebUiAddressRequest, runtime *dara.RuntimeOptions) (_result *GetSparkAppWebUiAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSparkAppWebUiAddress"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSparkAppWebUiAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the URL of the web UI for a Spark application.
//
// Description:
//
//	  General endpoint: `adb.aliyuncs.com`.
//
//		- Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - GetSparkAppWebUiAddressRequest
//
// @return GetSparkAppWebUiAddressResponse
func (client *Client) GetSparkAppWebUiAddress(request *GetSparkAppWebUiAddressRequest) (_result *GetSparkAppWebUiAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSparkAppWebUiAddressResponse{}
	_body, _err := client.GetSparkAppWebUiAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Spark log configuration of an AnalyticDB for MySQL cluster, including the default Spark log path.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - GetSparkConfigLogPathRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSparkConfigLogPathResponse
func (client *Client) GetSparkConfigLogPathWithOptions(request *GetSparkConfigLogPathRequest, runtime *dara.RuntimeOptions) (_result *GetSparkConfigLogPathResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSparkConfigLogPath"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSparkConfigLogPathResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Spark log configuration of an AnalyticDB for MySQL cluster, including the default Spark log path.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - GetSparkConfigLogPathRequest
//
// @return GetSparkConfigLogPathResponse
func (client *Client) GetSparkConfigLogPath(request *GetSparkConfigLogPathRequest) (_result *GetSparkConfigLogPathResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSparkConfigLogPathResponse{}
	_body, _err := client.GetSparkConfigLogPathWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the results of a Spark log analysis task.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - GetSparkLogAnalyzeTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSparkLogAnalyzeTaskResponse
func (client *Client) GetSparkLogAnalyzeTaskWithOptions(request *GetSparkLogAnalyzeTaskRequest, runtime *dara.RuntimeOptions) (_result *GetSparkLogAnalyzeTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSparkLogAnalyzeTask"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSparkLogAnalyzeTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the results of a Spark log analysis task.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - GetSparkLogAnalyzeTaskRequest
//
// @return GetSparkLogAnalyzeTaskResponse
func (client *Client) GetSparkLogAnalyzeTask(request *GetSparkLogAnalyzeTaskRequest) (_result *GetSparkLogAnalyzeTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSparkLogAnalyzeTaskResponse{}
	_body, _err := client.GetSparkLogAnalyzeTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of a Spark session.
//
// @param request - GetSparkReplSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSparkReplSessionResponse
func (client *Client) GetSparkReplSessionWithOptions(request *GetSparkReplSessionRequest, runtime *dara.RuntimeOptions) (_result *GetSparkReplSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSparkReplSession"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSparkReplSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of a Spark session.
//
// @param request - GetSparkReplSessionRequest
//
// @return GetSparkReplSessionResponse
func (client *Client) GetSparkReplSession(request *GetSparkReplSessionRequest) (_result *GetSparkReplSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSparkReplSessionResponse{}
	_body, _err := client.GetSparkReplSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the execution result of a code block.
//
// @param request - GetSparkReplStatementRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSparkReplStatementResponse
func (client *Client) GetSparkReplStatementWithOptions(request *GetSparkReplStatementRequest, runtime *dara.RuntimeOptions) (_result *GetSparkReplStatementResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.StatementId) {
		body["StatementId"] = request.StatementId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSparkReplStatement"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSparkReplStatementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution result of a code block.
//
// @param request - GetSparkReplStatementRequest
//
// @return GetSparkReplStatementResponse
func (client *Client) GetSparkReplStatement(request *GetSparkReplStatementRequest) (_result *GetSparkReplStatementResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSparkReplStatementResponse{}
	_body, _err := client.GetSparkReplStatementWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetSparkSQLEngineState is deprecated
//
// Summary:
//
// Queries the state information about the Spark SQL engine.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - GetSparkSQLEngineStateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSparkSQLEngineStateResponse
func (client *Client) GetSparkSQLEngineStateWithOptions(request *GetSparkSQLEngineStateRequest, runtime *dara.RuntimeOptions) (_result *GetSparkSQLEngineStateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.ResourceGroupName) {
		body["ResourceGroupName"] = request.ResourceGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSparkSQLEngineState"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSparkSQLEngineStateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetSparkSQLEngineState is deprecated
//
// Summary:
//
// Queries the state information about the Spark SQL engine.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - GetSparkSQLEngineStateRequest
//
// @return GetSparkSQLEngineStateResponse
// Deprecated
func (client *Client) GetSparkSQLEngineState(request *GetSparkSQLEngineStateRequest) (_result *GetSparkSQLEngineStateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSparkSQLEngineStateResponse{}
	_body, _err := client.GetSparkSQLEngineStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the content of a Spark application template.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - GetSparkTemplateFileContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSparkTemplateFileContentResponse
func (client *Client) GetSparkTemplateFileContentWithOptions(request *GetSparkTemplateFileContentRequest, runtime *dara.RuntimeOptions) (_result *GetSparkTemplateFileContentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSparkTemplateFileContent"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSparkTemplateFileContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the content of a Spark application template.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - GetSparkTemplateFileContentRequest
//
// @return GetSparkTemplateFileContentResponse
func (client *Client) GetSparkTemplateFileContent(request *GetSparkTemplateFileContentRequest) (_result *GetSparkTemplateFileContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSparkTemplateFileContentResponse{}
	_body, _err := client.GetSparkTemplateFileContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the directory structure of Spark applications.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - GetSparkTemplateFolderTreeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSparkTemplateFolderTreeResponse
func (client *Client) GetSparkTemplateFolderTreeWithOptions(request *GetSparkTemplateFolderTreeRequest, runtime *dara.RuntimeOptions) (_result *GetSparkTemplateFolderTreeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSparkTemplateFolderTree"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSparkTemplateFolderTreeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the directory structure of Spark applications.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - GetSparkTemplateFolderTreeRequest
//
// @return GetSparkTemplateFolderTreeResponse
func (client *Client) GetSparkTemplateFolderTree(request *GetSparkTemplateFolderTreeRequest) (_result *GetSparkTemplateFolderTreeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSparkTemplateFolderTreeResponse{}
	_body, _err := client.GetSparkTemplateFolderTreeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the directory structure of Spark applications.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - GetSparkTemplateFullTreeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSparkTemplateFullTreeResponse
func (client *Client) GetSparkTemplateFullTreeWithOptions(request *GetSparkTemplateFullTreeRequest, runtime *dara.RuntimeOptions) (_result *GetSparkTemplateFullTreeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSparkTemplateFullTree"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSparkTemplateFullTreeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the directory structure of Spark applications.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - GetSparkTemplateFullTreeRequest
//
// @return GetSparkTemplateFullTreeResponse
func (client *Client) GetSparkTemplateFullTree(request *GetSparkTemplateFullTreeRequest) (_result *GetSparkTemplateFullTreeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSparkTemplateFullTreeResponse{}
	_body, _err := client.GetSparkTemplateFullTreeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the execution result of a Spark SQL statement.
//
// @param request - GetSparkWarehouseBatchSQLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSparkWarehouseBatchSQLResponse
func (client *Client) GetSparkWarehouseBatchSQLWithOptions(request *GetSparkWarehouseBatchSQLRequest, runtime *dara.RuntimeOptions) (_result *GetSparkWarehouseBatchSQLResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Agency) {
		body["Agency"] = request.Agency
	}

	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.QueryId) {
		body["QueryId"] = request.QueryId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSparkWarehouseBatchSQL"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSparkWarehouseBatchSQLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution result of a Spark SQL statement.
//
// @param request - GetSparkWarehouseBatchSQLRequest
//
// @return GetSparkWarehouseBatchSQLResponse
func (client *Client) GetSparkWarehouseBatchSQL(request *GetSparkWarehouseBatchSQLRequest) (_result *GetSparkWarehouseBatchSQLResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSparkWarehouseBatchSQLResponse{}
	_body, _err := client.GetSparkWarehouseBatchSQLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取表
//
// @param request - GetTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableResponse
func (client *Client) GetTableWithOptions(request *GetTableRequest, runtime *dara.RuntimeOptions) (_result *GetTableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
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
		Action:      dara.String("GetTable"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取表
//
// @param request - GetTableRequest
//
// @return GetTableResponse
func (client *Client) GetTable(request *GetTableRequest) (_result *GetTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTableResponse{}
	_body, _err := client.GetTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about columns.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - GetTableColumnsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableColumnsResponse
func (client *Client) GetTableColumnsWithOptions(request *GetTableColumnsRequest, runtime *dara.RuntimeOptions) (_result *GetTableColumnsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ColumnName) {
		query["ColumnName"] = request.ColumnName
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.SchemaName) {
		query["SchemaName"] = request.SchemaName
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTableColumns"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTableColumnsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about columns.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - GetTableColumnsRequest
//
// @return GetTableColumnsResponse
func (client *Client) GetTableColumns(request *GetTableColumnsRequest) (_result *GetTableColumnsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTableColumnsResponse{}
	_body, _err := client.GetTableColumnsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the statement that is used to create a table.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - GetTableDDLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableDDLResponse
func (client *Client) GetTableDDLWithOptions(request *GetTableDDLRequest, runtime *dara.RuntimeOptions) (_result *GetTableDDLResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SchemaName) {
		query["SchemaName"] = request.SchemaName
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTableDDL"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTableDDLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statement that is used to create a table.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - GetTableDDLRequest
//
// @return GetTableDDLResponse
func (client *Client) GetTableDDL(request *GetTableDDLRequest) (_result *GetTableDDLResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTableDDLResponse{}
	_body, _err := client.GetTableDDLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取table概要信息
//
// @param request - GetTableObjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableObjectsResponse
func (client *Client) GetTableObjectsWithOptions(request *GetTableObjectsRequest, runtime *dara.RuntimeOptions) (_result *GetTableObjectsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.FilterDescription) {
		query["FilterDescription"] = request.FilterDescription
	}

	if !dara.IsNil(request.FilterOwner) {
		query["FilterOwner"] = request.FilterOwner
	}

	if !dara.IsNil(request.FilterTblName) {
		query["FilterTblName"] = request.FilterTblName
	}

	if !dara.IsNil(request.FilterTblType) {
		query["FilterTblType"] = request.FilterTblType
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SchemaName) {
		query["SchemaName"] = request.SchemaName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTableObjects"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTableObjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取table概要信息
//
// @param request - GetTableObjectsRequest
//
// @return GetTableObjectsResponse
func (client *Client) GetTableObjects(request *GetTableObjectsRequest) (_result *GetTableObjectsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTableObjectsResponse{}
	_body, _err := client.GetTableObjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the statement that is used to create a view.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - GetViewDDLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetViewDDLResponse
func (client *Client) GetViewDDLWithOptions(request *GetViewDDLRequest, runtime *dara.RuntimeOptions) (_result *GetViewDDLResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SchemaName) {
		query["SchemaName"] = request.SchemaName
	}

	if !dara.IsNil(request.ViewName) {
		query["ViewName"] = request.ViewName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetViewDDL"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetViewDDLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statement that is used to create a view.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - GetViewDDLRequest
//
// @return GetViewDDLResponse
func (client *Client) GetViewDDL(request *GetViewDDLRequest) (_result *GetViewDDLResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetViewDDLResponse{}
	_body, _err := client.GetViewDDLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about views.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - GetViewObjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetViewObjectsResponse
func (client *Client) GetViewObjectsWithOptions(request *GetViewObjectsRequest, runtime *dara.RuntimeOptions) (_result *GetViewObjectsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.FilterOwner) {
		query["FilterOwner"] = request.FilterOwner
	}

	if !dara.IsNil(request.FilterViewName) {
		query["FilterViewName"] = request.FilterViewName
	}

	if !dara.IsNil(request.FilterViewType) {
		query["FilterViewType"] = request.FilterViewType
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SchemaName) {
		query["SchemaName"] = request.SchemaName
	}

	if !dara.IsNil(request.ShowMvBaseTable) {
		query["ShowMvBaseTable"] = request.ShowMvBaseTable
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetViewObjects"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetViewObjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about views.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// @param request - GetViewObjectsRequest
//
// @return GetViewObjectsResponse
func (client *Client) GetViewObjects(request *GetViewObjectsRequest) (_result *GetViewObjectsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetViewObjectsResponse{}
	_body, _err := client.GetViewObjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Grants permissions to the service account of an AnalyticDB for MySQL cluster.
//
// @param request - GrantOperatorPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantOperatorPermissionResponse
func (client *Client) GrantOperatorPermissionWithOptions(request *GrantOperatorPermissionRequest, runtime *dara.RuntimeOptions) (_result *GrantOperatorPermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.ExpiredTime) {
		query["ExpiredTime"] = request.ExpiredTime
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Privileges) {
		query["Privileges"] = request.Privileges
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
		Action:      dara.String("GrantOperatorPermission"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GrantOperatorPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants permissions to the service account of an AnalyticDB for MySQL cluster.
//
// @param request - GrantOperatorPermissionRequest
//
// @return GrantOperatorPermissionResponse
func (client *Client) GrantOperatorPermission(request *GrantOperatorPermissionRequest) (_result *GrantOperatorPermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GrantOperatorPermissionResponse{}
	_body, _err := client.GrantOperatorPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Terminates an ongoing query.
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
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("KillProcess"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
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
// Terminates an ongoing query.
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
// Terminates a Spark application.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - KillSparkAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return KillSparkAppResponse
func (client *Client) KillSparkAppWithOptions(request *KillSparkAppRequest, runtime *dara.RuntimeOptions) (_result *KillSparkAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("KillSparkApp"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &KillSparkAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Terminates a Spark application.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - KillSparkAppRequest
//
// @return KillSparkAppResponse
func (client *Client) KillSparkApp(request *KillSparkAppRequest) (_result *KillSparkAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &KillSparkAppResponse{}
	_body, _err := client.KillSparkAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Terminates a Spark log analysis task and queries the information about the analysis task.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - KillSparkLogAnalyzeTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return KillSparkLogAnalyzeTaskResponse
func (client *Client) KillSparkLogAnalyzeTaskWithOptions(request *KillSparkLogAnalyzeTaskRequest, runtime *dara.RuntimeOptions) (_result *KillSparkLogAnalyzeTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("KillSparkLogAnalyzeTask"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &KillSparkLogAnalyzeTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Terminates a Spark log analysis task and queries the information about the analysis task.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - KillSparkLogAnalyzeTaskRequest
//
// @return KillSparkLogAnalyzeTaskResponse
func (client *Client) KillSparkLogAnalyzeTask(request *KillSparkLogAnalyzeTaskRequest) (_result *KillSparkLogAnalyzeTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &KillSparkLogAnalyzeTaskResponse{}
	_body, _err := client.KillSparkLogAnalyzeTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI KillSparkSQLEngine is deprecated
//
// Summary:
//
// Shuts down a Spark SQL engine.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - KillSparkSQLEngineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return KillSparkSQLEngineResponse
func (client *Client) KillSparkSQLEngineWithOptions(request *KillSparkSQLEngineRequest, runtime *dara.RuntimeOptions) (_result *KillSparkSQLEngineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.ResourceGroupName) {
		body["ResourceGroupName"] = request.ResourceGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("KillSparkSQLEngine"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &KillSparkSQLEngineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI KillSparkSQLEngine is deprecated
//
// Summary:
//
// Shuts down a Spark SQL engine.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - KillSparkSQLEngineRequest
//
// @return KillSparkSQLEngineResponse
// Deprecated
func (client *Client) KillSparkSQLEngine(request *KillSparkSQLEngineRequest) (_result *KillSparkSQLEngineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &KillSparkSQLEngineResponse{}
	_body, _err := client.KillSparkSQLEngineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of lifecycle management policies of an AnalyticDB for MySQL cluster.
//
// @param request - ListApsLifecycleStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApsLifecycleStrategyResponse
func (client *Client) ListApsLifecycleStrategyWithOptions(request *ListApsLifecycleStrategyRequest, runtime *dara.RuntimeOptions) (_result *ListApsLifecycleStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApsLifecycleStrategy"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApsLifecycleStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of lifecycle management policies of an AnalyticDB for MySQL cluster.
//
// @param request - ListApsLifecycleStrategyRequest
//
// @return ListApsLifecycleStrategyResponse
func (client *Client) ListApsLifecycleStrategy(request *ListApsLifecycleStrategyRequest) (_result *ListApsLifecycleStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListApsLifecycleStrategyResponse{}
	_body, _err := client.ListApsLifecycleStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of lake storage optimization policies for an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - ListApsOptimizationStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApsOptimizationStrategyResponse
func (client *Client) ListApsOptimizationStrategyWithOptions(request *ListApsOptimizationStrategyRequest, runtime *dara.RuntimeOptions) (_result *ListApsOptimizationStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApsOptimizationStrategy"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApsOptimizationStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of lake storage optimization policies for an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - ListApsOptimizationStrategyRequest
//
// @return ListApsOptimizationStrategyResponse
func (client *Client) ListApsOptimizationStrategy(request *ListApsOptimizationStrategyRequest) (_result *ListApsOptimizationStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListApsOptimizationStrategyResponse{}
	_body, _err := client.ListApsOptimizationStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of optimization jobs executed based on a lifecycle management policy. The system runs optimization jobs on a regular basis based on lifecycle management policies.
//
// @param request - ListApsOptimizationTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApsOptimizationTasksResponse
func (client *Client) ListApsOptimizationTasksWithOptions(request *ListApsOptimizationTasksRequest, runtime *dara.RuntimeOptions) (_result *ListApsOptimizationTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StrategyType) {
		body["StrategyType"] = request.StrategyType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApsOptimizationTasks"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApsOptimizationTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of optimization jobs executed based on a lifecycle management policy. The system runs optimization jobs on a regular basis based on lifecycle management policies.
//
// @param request - ListApsOptimizationTasksRequest
//
// @return ListApsOptimizationTasksResponse
func (client *Client) ListApsOptimizationTasks(request *ListApsOptimizationTasksRequest) (_result *ListApsOptimizationTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListApsOptimizationTasksResponse{}
	_body, _err := client.ListApsOptimizationTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Webhook configurations of a specified database cluster.
//
// Description:
//
// This API allows you to obtain a list of configured webhooks based on `RegionId`, `DBClusterId`, and optional `JobType`. The `JobType` parameter specifies the task type, such as SLS/OSS export task. If the parameter is provided, webhooks related to the task type are returned. If the parameter is not provided, all types of webhooks are returned.
//
// Note: Make sure that the `RegionId` and `DBClusterId` you provided are correct. Otherwise, the webhook information may not be obtained correctly.
//
// @param request - ListApsWebhookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApsWebhookResponse
func (client *Client) ListApsWebhookWithOptions(request *ListApsWebhookRequest, runtime *dara.RuntimeOptions) (_result *ListApsWebhookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.JobType) {
		body["JobType"] = request.JobType
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApsWebhook"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApsWebhookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Webhook configurations of a specified database cluster.
//
// Description:
//
// This API allows you to obtain a list of configured webhooks based on `RegionId`, `DBClusterId`, and optional `JobType`. The `JobType` parameter specifies the task type, such as SLS/OSS export task. If the parameter is provided, webhooks related to the task type are returned. If the parameter is not provided, all types of webhooks are returned.
//
// Note: Make sure that the `RegionId` and `DBClusterId` you provided are correct. Otherwise, the webhook information may not be obtained correctly.
//
// @param request - ListApsWebhookRequest
//
// @return ListApsWebhookResponse
func (client *Client) ListApsWebhook(request *ListApsWebhookRequest) (_result *ListApsWebhookResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListApsWebhookResponse{}
	_body, _err := client.ListApsWebhookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of lake storages.
//
// @param request - ListLakeStoragesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLakeStoragesResponse
func (client *Client) ListLakeStoragesWithOptions(request *ListLakeStoragesRequest, runtime *dara.RuntimeOptions) (_result *ListLakeStoragesResponse, _err error) {
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
		Action:      dara.String("ListLakeStorages"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLakeStoragesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of lake storages.
//
// @param request - ListLakeStoragesRequest
//
// @return ListLakeStoragesResponse
func (client *Client) ListLakeStorages(request *ListLakeStoragesRequest) (_result *ListLakeStoragesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLakeStoragesResponse{}
	_body, _err := client.ListLakeStoragesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the execution records of result set export jobs of a Resource Access Management (RAM) user.
//
// @param tmpReq - ListResultExportJobHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResultExportJobHistoryResponse
func (client *Client) ListResultExportJobHistoryWithOptions(tmpReq *ListResultExportJobHistoryRequest, runtime *dara.RuntimeOptions) (_result *ListResultExportJobHistoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListResultExportJobHistoryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Order) {
		request.OrderShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Order, dara.String("Order"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.StatusList) {
		request.StatusListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StatusList, dara.String("StatusList"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DatabaseUser) {
		body["DatabaseUser"] = request.DatabaseUser
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OrderShrink) {
		body["Order"] = request.OrderShrink
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroup) {
		body["ResourceGroup"] = request.ResourceGroup
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StatusListShrink) {
		body["StatusList"] = request.StatusListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResultExportJobHistory"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResultExportJobHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution records of result set export jobs of a Resource Access Management (RAM) user.
//
// @param request - ListResultExportJobHistoryRequest
//
// @return ListResultExportJobHistoryResponse
func (client *Client) ListResultExportJobHistory(request *ListResultExportJobHistoryRequest) (_result *ListResultExportJobHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListResultExportJobHistoryResponse{}
	_body, _err := client.ListResultExportJobHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about retry attempts of a Spark application.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - ListSparkAppAttemptsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSparkAppAttemptsResponse
func (client *Client) ListSparkAppAttemptsWithOptions(request *ListSparkAppAttemptsRequest, runtime *dara.RuntimeOptions) (_result *ListSparkAppAttemptsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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
		Action:      dara.String("ListSparkAppAttempts"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSparkAppAttemptsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about retry attempts of a Spark application.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - ListSparkAppAttemptsRequest
//
// @return ListSparkAppAttemptsResponse
func (client *Client) ListSparkAppAttempts(request *ListSparkAppAttemptsRequest) (_result *ListSparkAppAttemptsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSparkAppAttemptsResponse{}
	_body, _err := client.ListSparkAppAttemptsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Spark applications that run on an AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
//
// @param request - ListSparkAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSparkAppsResponse
func (client *Client) ListSparkAppsWithOptions(request *ListSparkAppsRequest, runtime *dara.RuntimeOptions) (_result *ListSparkAppsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Filters) {
		query["Filters"] = request.Filters
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupName) {
		query["ResourceGroupName"] = request.ResourceGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSparkApps"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSparkAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Spark applications that run on an AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
//
// @param request - ListSparkAppsRequest
//
// @return ListSparkAppsResponse
func (client *Client) ListSparkApps(request *ListSparkAppsRequest) (_result *ListSparkAppsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSparkAppsResponse{}
	_body, _err := client.ListSparkAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of Spark log analysis tasks.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - ListSparkLogAnalyzeTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSparkLogAnalyzeTasksResponse
func (client *Client) ListSparkLogAnalyzeTasksWithOptions(request *ListSparkLogAnalyzeTasksRequest, runtime *dara.RuntimeOptions) (_result *ListSparkLogAnalyzeTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSparkLogAnalyzeTasks"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSparkLogAnalyzeTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of Spark log analysis tasks.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - ListSparkLogAnalyzeTasksRequest
//
// @return ListSparkLogAnalyzeTasksResponse
func (client *Client) ListSparkLogAnalyzeTasks(request *ListSparkLogAnalyzeTasksRequest) (_result *ListSparkLogAnalyzeTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSparkLogAnalyzeTasksResponse{}
	_body, _err := client.ListSparkLogAnalyzeTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of Spark template file IDs for AnalyticDB for MySQL clusters.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - ListSparkTemplateFileIdsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSparkTemplateFileIdsResponse
func (client *Client) ListSparkTemplateFileIdsWithOptions(request *ListSparkTemplateFileIdsRequest, runtime *dara.RuntimeOptions) (_result *ListSparkTemplateFileIdsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSparkTemplateFileIds"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSparkTemplateFileIdsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of Spark template file IDs for AnalyticDB for MySQL clusters.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - ListSparkTemplateFileIdsRequest
//
// @return ListSparkTemplateFileIdsResponse
func (client *Client) ListSparkTemplateFileIds(request *ListSparkTemplateFileIdsRequest) (_result *ListSparkTemplateFileIdsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSparkTemplateFileIdsResponse{}
	_body, _err := client.ListSparkTemplateFileIdsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of Spark SQL statements.
//
// @param request - ListSparkWarehouseBatchSQLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSparkWarehouseBatchSQLResponse
func (client *Client) ListSparkWarehouseBatchSQLWithOptions(request *ListSparkWarehouseBatchSQLRequest, runtime *dara.RuntimeOptions) (_result *ListSparkWarehouseBatchSQLResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupName) {
		body["ResourceGroupName"] = request.ResourceGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSparkWarehouseBatchSQL"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSparkWarehouseBatchSQLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of Spark SQL statements.
//
// @param request - ListSparkWarehouseBatchSQLRequest
//
// @return ListSparkWarehouseBatchSQLResponse
func (client *Client) ListSparkWarehouseBatchSQL(request *ListSparkWarehouseBatchSQLRequest) (_result *ListSparkWarehouseBatchSQLResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSparkWarehouseBatchSQLResponse{}
	_body, _err := client.ListSparkWarehouseBatchSQLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags that are added to AnalyticDB for MySQL clusters, or the AnalyticDB for MySQL clusters that have tags added.
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
		Version:     dara.String("2021-12-01"),
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
// Queries the tags that are added to AnalyticDB for MySQL clusters, or the AnalyticDB for MySQL clusters that have tags added.
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
// Loads a built-in dataset.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - LoadSampleDataSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LoadSampleDataSetResponse
func (client *Client) LoadSampleDataSetWithOptions(request *LoadSampleDataSetRequest, runtime *dara.RuntimeOptions) (_result *LoadSampleDataSetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LoadSampleDataSet"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LoadSampleDataSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Loads a built-in dataset.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - LoadSampleDataSetRequest
//
// @return LoadSampleDataSetResponse
func (client *Client) LoadSampleDataSet(request *LoadSampleDataSetRequest) (_result *LoadSampleDataSetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &LoadSampleDataSetResponse{}
	_body, _err := client.LoadSampleDataSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the description of a database account for an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
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
	if !dara.IsNil(request.AccountDescription) {
		query["AccountDescription"] = request.AccountDescription
	}

	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAccountDescription"),
		Version:     dara.String("2021-12-01"),
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
// Modifies the description of a database account for an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
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
// Modifies the permissions of a database account.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param tmpReq - ModifyAccountPrivilegesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccountPrivilegesResponse
func (client *Client) ModifyAccountPrivilegesWithOptions(tmpReq *ModifyAccountPrivilegesRequest, runtime *dara.RuntimeOptions) (_result *ModifyAccountPrivilegesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyAccountPrivilegesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AccountPrivileges) {
		request.AccountPrivilegesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AccountPrivileges, dara.String("AccountPrivileges"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.AccountPrivilegesShrink) {
		query["AccountPrivileges"] = request.AccountPrivilegesShrink
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAccountPrivileges"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAccountPrivilegesResponse{}
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
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - ModifyAccountPrivilegesRequest
//
// @return ModifyAccountPrivilegesResponse
func (client *Client) ModifyAccountPrivileges(request *ModifyAccountPrivilegesRequest) (_result *ModifyAccountPrivilegesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAccountPrivilegesResponse{}
	_body, _err := client.ModifyAccountPrivilegesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies an AnalyticDB Pipeline Service (APS) data source.
//
// Description:
//
// ### [](#)
//
//   - You can call this operation only for AnalyticDB for MySQL clusters in elastic mode for Cluster Edition that have 32 cores or more.
//
//   - You cannot modify the number of nodes for the USER_DEFAULT resource group.
//
// @param tmpReq - ModifyApsDatasoureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApsDatasoureResponse
func (client *Client) ModifyApsDatasoureWithOptions(tmpReq *ModifyApsDatasoureRequest, runtime *dara.RuntimeOptions) (_result *ModifyApsDatasoureResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyApsDatasoureShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.KafkaInfo) {
		request.KafkaInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.KafkaInfo, dara.String("KafkaInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.LakehouseId) {
		request.LakehouseIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LakehouseId, dara.String("LakehouseId"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PolarDBMysqlInfo) {
		request.PolarDBMysqlInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PolarDBMysqlInfo, dara.String("PolarDBMysqlInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RdsMysqlInfo) {
		request.RdsMysqlInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RdsMysqlInfo, dara.String("RdsMysqlInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SlsInfo) {
		request.SlsInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SlsInfo, dara.String("SlsInfo"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DatasourceDescription) {
		body["DatasourceDescription"] = request.DatasourceDescription
	}

	if !dara.IsNil(request.DatasourceId) {
		body["DatasourceId"] = request.DatasourceId
	}

	if !dara.IsNil(request.DatasourceName) {
		body["DatasourceName"] = request.DatasourceName
	}

	if !dara.IsNil(request.KafkaInfoShrink) {
		body["KafkaInfo"] = request.KafkaInfoShrink
	}

	if !dara.IsNil(request.LakehouseIdShrink) {
		body["LakehouseId"] = request.LakehouseIdShrink
	}

	if !dara.IsNil(request.PolarDBMysqlInfoShrink) {
		body["PolarDBMysqlInfo"] = request.PolarDBMysqlInfoShrink
	}

	if !dara.IsNil(request.RdsMysqlInfoShrink) {
		body["RdsMysqlInfo"] = request.RdsMysqlInfoShrink
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SlsInfoShrink) {
		body["SlsInfo"] = request.SlsInfoShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyApsDatasoure"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyApsDatasoureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an AnalyticDB Pipeline Service (APS) data source.
//
// Description:
//
// ### [](#)
//
//   - You can call this operation only for AnalyticDB for MySQL clusters in elastic mode for Cluster Edition that have 32 cores or more.
//
//   - You cannot modify the number of nodes for the USER_DEFAULT resource group.
//
// @param request - ModifyApsDatasoureRequest
//
// @return ModifyApsDatasoureResponse
func (client *Client) ModifyApsDatasoure(request *ModifyApsDatasoureRequest) (_result *ModifyApsDatasoureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyApsDatasoureResponse{}
	_body, _err := client.ModifyApsDatasoureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies an AnalyticDB Pipeline Service (APS) job.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - ModifyApsJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApsJobResponse
func (client *Client) ModifyApsJobWithOptions(request *ModifyApsJobRequest, runtime *dara.RuntimeOptions) (_result *ModifyApsJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApsJobId) {
		body["ApsJobId"] = request.ApsJobId
	}

	if !dara.IsNil(request.DbList) {
		body["DbList"] = request.DbList
	}

	if !dara.IsNil(request.PartitionList) {
		body["PartitionList"] = request.PartitionList
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyApsJob"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyApsJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an AnalyticDB Pipeline Service (APS) job.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - ModifyApsJobRequest
//
// @return ModifyApsJobResponse
func (client *Client) ModifyApsJob(request *ModifyApsJobRequest) (_result *ModifyApsJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyApsJobResponse{}
	_body, _err := client.ModifyApsJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies an AnalyticDB Pipeline Service (APS) job from Simple Log Service (SLS) to an AnalyticDB for MySQL Data Warehouse Edition cluster.
//
// @param tmpReq - ModifyApsSlsADBJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApsSlsADBJobResponse
func (client *Client) ModifyApsSlsADBJobWithOptions(tmpReq *ModifyApsSlsADBJobRequest, runtime *dara.RuntimeOptions) (_result *ModifyApsSlsADBJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyApsSlsADBJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Columns) {
		request.ColumnsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Columns, dara.String("Columns"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ColumnsShrink) {
		body["Columns"] = request.ColumnsShrink
	}

	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DbName) {
		body["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DirtyDataProcessPattern) {
		body["DirtyDataProcessPattern"] = request.DirtyDataProcessPattern
	}

	if !dara.IsNil(request.ExactlyOnce) {
		body["ExactlyOnce"] = request.ExactlyOnce
	}

	if !dara.IsNil(request.Password) {
		body["Password"] = request.Password
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartingOffsets) {
		body["StartingOffsets"] = request.StartingOffsets
	}

	if !dara.IsNil(request.TableName) {
		body["TableName"] = request.TableName
	}

	if !dara.IsNil(request.UnixTimestampConvert) {
		body["UnixTimestampConvert"] = request.UnixTimestampConvert
	}

	if !dara.IsNil(request.UserName) {
		body["UserName"] = request.UserName
	}

	if !dara.IsNil(request.WorkloadId) {
		body["WorkloadId"] = request.WorkloadId
	}

	if !dara.IsNil(request.WorkloadName) {
		body["WorkloadName"] = request.WorkloadName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyApsSlsADBJob"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyApsSlsADBJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an AnalyticDB Pipeline Service (APS) job from Simple Log Service (SLS) to an AnalyticDB for MySQL Data Warehouse Edition cluster.
//
// @param request - ModifyApsSlsADBJobRequest
//
// @return ModifyApsSlsADBJobResponse
func (client *Client) ModifyApsSlsADBJob(request *ModifyApsSlsADBJobRequest) (_result *ModifyApsSlsADBJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyApsSlsADBJobResponse{}
	_body, _err := client.ModifyApsSlsADBJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the AnalyticDB Pipeline Service (APS) workload name.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - ModifyApsWorkloadNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApsWorkloadNameResponse
func (client *Client) ModifyApsWorkloadNameWithOptions(request *ModifyApsWorkloadNameRequest, runtime *dara.RuntimeOptions) (_result *ModifyApsWorkloadNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WorkloadId) {
		body["WorkloadId"] = request.WorkloadId
	}

	if !dara.IsNil(request.WorkloadName) {
		body["WorkloadName"] = request.WorkloadName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyApsWorkloadName"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyApsWorkloadNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the AnalyticDB Pipeline Service (APS) workload name.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - ModifyApsWorkloadNameRequest
//
// @return ModifyApsWorkloadNameResponse
func (client *Client) ModifyApsWorkloadName(request *ModifyApsWorkloadNameRequest) (_result *ModifyApsWorkloadNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyApsWorkloadNameResponse{}
	_body, _err := client.ModifyApsWorkloadNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the SQL audit settings of an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - ModifyAuditLogConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAuditLogConfigResponse
func (client *Client) ModifyAuditLogConfigWithOptions(request *ModifyAuditLogConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyAuditLogConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuditLogStatus) {
		query["AuditLogStatus"] = request.AuditLogStatus
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EngineType) {
		query["EngineType"] = request.EngineType
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
		Action:      dara.String("ModifyAuditLogConfig"),
		Version:     dara.String("2021-12-01"),
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
// Modifies the SQL audit settings of an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
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
// Modifies the backup policy of an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
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

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EnableBackupLog) {
		query["EnableBackupLog"] = request.EnableBackupLog
	}

	if !dara.IsNil(request.LogBackupRetentionPeriod) {
		query["LogBackupRetentionPeriod"] = request.LogBackupRetentionPeriod
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBackupPolicy"),
		Version:     dara.String("2021-12-01"),
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
// Modifies the backup policy of an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
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
// Enables the wide table engine feature for an AnalyticDB for MySQL cluster or modifies the disk cache size of the wide table engine of an AnalyticDB for MySQL cluster for which you enabled the wide table engine feature.
//
// @param request - ModifyClickhouseEngineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyClickhouseEngineResponse
func (client *Client) ModifyClickhouseEngineWithOptions(request *ModifyClickhouseEngineRequest, runtime *dara.RuntimeOptions) (_result *ModifyClickhouseEngineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CacheSize) {
		query["CacheSize"] = request.CacheSize
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Enabled) {
		query["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyClickhouseEngine"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyClickhouseEngineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the wide table engine feature for an AnalyticDB for MySQL cluster or modifies the disk cache size of the wide table engine of an AnalyticDB for MySQL cluster for which you enabled the wide table engine feature.
//
// @param request - ModifyClickhouseEngineRequest
//
// @return ModifyClickhouseEngineResponse
func (client *Client) ModifyClickhouseEngine(request *ModifyClickhouseEngineRequest) (_result *ModifyClickhouseEngineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyClickhouseEngineResponse{}
	_body, _err := client.ModifyClickhouseEngineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the IP address whitelist of an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see Endpoints.
//
// @param request - ModifyClusterAccessWhiteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyClusterAccessWhiteListResponse
func (client *Client) ModifyClusterAccessWhiteListWithOptions(request *ModifyClusterAccessWhiteListRequest, runtime *dara.RuntimeOptions) (_result *ModifyClusterAccessWhiteListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterIPArrayAttribute) {
		query["DBClusterIPArrayAttribute"] = request.DBClusterIPArrayAttribute
	}

	if !dara.IsNil(request.DBClusterIPArrayName) {
		query["DBClusterIPArrayName"] = request.DBClusterIPArrayName
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.ModifyMode) {
		query["ModifyMode"] = request.ModifyMode
	}

	if !dara.IsNil(request.SecurityIps) {
		query["SecurityIps"] = request.SecurityIps
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyClusterAccessWhiteList"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyClusterAccessWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the IP address whitelist of an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see Endpoints.
//
// @param request - ModifyClusterAccessWhiteListRequest
//
// @return ModifyClusterAccessWhiteListResponse
func (client *Client) ModifyClusterAccessWhiteList(request *ModifyClusterAccessWhiteListRequest) (_result *ModifyClusterAccessWhiteListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyClusterAccessWhiteListResponse{}
	_body, _err := client.ModifyClusterAccessWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the public endpoint of an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - ModifyClusterConnectionStringRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyClusterConnectionStringResponse
func (client *Client) ModifyClusterConnectionStringWithOptions(request *ModifyClusterConnectionStringRequest, runtime *dara.RuntimeOptions) (_result *ModifyClusterConnectionStringResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionStringPrefix) {
		query["ConnectionStringPrefix"] = request.ConnectionStringPrefix
	}

	if !dara.IsNil(request.CurrentConnectionString) {
		query["CurrentConnectionString"] = request.CurrentConnectionString
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyClusterConnectionString"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyClusterConnectionStringResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the public endpoint of an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - ModifyClusterConnectionStringRequest
//
// @return ModifyClusterConnectionStringResponse
func (client *Client) ModifyClusterConnectionString(request *ModifyClusterConnectionStringRequest) (_result *ModifyClusterConnectionStringResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyClusterConnectionStringResponse{}
	_body, _err := client.ModifyClusterConnectionStringWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the status of the remote build feature in the query acceleration configuration of an AnalyticDB for MySQL cluster.
//
// @param request - ModifyCompactionServiceSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCompactionServiceSwitchResponse
func (client *Client) ModifyCompactionServiceSwitchWithOptions(request *ModifyCompactionServiceSwitchRequest, runtime *dara.RuntimeOptions) (_result *ModifyCompactionServiceSwitchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EnableCompactionService) {
		query["EnableCompactionService"] = request.EnableCompactionService
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCompactionServiceSwitch"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCompactionServiceSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the status of the remote build feature in the query acceleration configuration of an AnalyticDB for MySQL cluster.
//
// @param request - ModifyCompactionServiceSwitchRequest
//
// @return ModifyCompactionServiceSwitchResponse
func (client *Client) ModifyCompactionServiceSwitch(request *ModifyCompactionServiceSwitchRequest) (_result *ModifyCompactionServiceSwitchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyCompactionServiceSwitchResponse{}
	_body, _err := client.ModifyCompactionServiceSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the configurations of an AnalyticDB for MySQL Data Lakehouse Edition cluster.
//
// Description:
//
// ### [](#)
//
//   - During a scaling event, you are not allowed to execute the `SUBMIT JOB` statement to submit asynchronous jobs. If your business requires asynchronous jobs, perform scaling during appropriate periods.
//
//   - When you scale a cluster, data in the cluster is migrated for redistribution. The amount of time that is required to migrate data is proportional to the data volume. During a scaling event, the services provided by the cluster are not interrupted. When you downgrade cluster specifications, data migration may require up to dozens of hours to complete. Proceed with caution especially if your cluster contains a large amount of data.
//
//   - If the cluster has a built-in dataset loaded, make sure that the cluster has reserved storage resources of at least 24 AnalyticDB compute units (ACUs). Otherwise, the built-in dataset cannot be used.
//
//   - When the scaling process is about to end, transient connections may occur. We recommend that you scale your cluster during off-peak hours or make sure that your application is configured to automatically reconnect to your cluster.
//
//   - You can change an AnalyticDB for MySQL cluster from Data Warehouse Edition to Data Lakehouse Edition, but not the other way around. For more information, see Change a cluster from Data Warehouse Edition to Data Lakehouse Edition.
//
//   - For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - ModifyDBClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterResponse
func (client *Client) ModifyDBClusterWithOptions(request *ModifyDBClusterRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ComputeResource) {
		query["ComputeResource"] = request.ComputeResource
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EnableDefaultResourcePool) {
		query["EnableDefaultResourcePool"] = request.EnableDefaultResourcePool
	}

	if !dara.IsNil(request.ProductForm) {
		query["ProductForm"] = request.ProductForm
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReservedNodeCount) {
		query["ReservedNodeCount"] = request.ReservedNodeCount
	}

	if !dara.IsNil(request.ReservedNodeSize) {
		query["ReservedNodeSize"] = request.ReservedNodeSize
	}

	if !dara.IsNil(request.StorageResource) {
		query["StorageResource"] = request.StorageResource
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBCluster"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the configurations of an AnalyticDB for MySQL Data Lakehouse Edition cluster.
//
// Description:
//
// ### [](#)
//
//   - During a scaling event, you are not allowed to execute the `SUBMIT JOB` statement to submit asynchronous jobs. If your business requires asynchronous jobs, perform scaling during appropriate periods.
//
//   - When you scale a cluster, data in the cluster is migrated for redistribution. The amount of time that is required to migrate data is proportional to the data volume. During a scaling event, the services provided by the cluster are not interrupted. When you downgrade cluster specifications, data migration may require up to dozens of hours to complete. Proceed with caution especially if your cluster contains a large amount of data.
//
//   - If the cluster has a built-in dataset loaded, make sure that the cluster has reserved storage resources of at least 24 AnalyticDB compute units (ACUs). Otherwise, the built-in dataset cannot be used.
//
//   - When the scaling process is about to end, transient connections may occur. We recommend that you scale your cluster during off-peak hours or make sure that your application is configured to automatically reconnect to your cluster.
//
//   - You can change an AnalyticDB for MySQL cluster from Data Warehouse Edition to Data Lakehouse Edition, but not the other way around. For more information, see Change a cluster from Data Warehouse Edition to Data Lakehouse Edition.
//
//   - For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - ModifyDBClusterRequest
//
// @return ModifyDBClusterResponse
func (client *Client) ModifyDBCluster(request *ModifyDBClusterRequest) (_result *ModifyDBClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBClusterResponse{}
	_body, _err := client.ModifyDBClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the description of an AnalyticDB for MySQL cluster to facilitate the maintenance and management of the cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - ModifyDBClusterDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterDescriptionResponse
func (client *Client) ModifyDBClusterDescriptionWithOptions(request *ModifyDBClusterDescriptionRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBClusterDescriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterDescription) {
		query["DBClusterDescription"] = request.DBClusterDescription
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBClusterDescription"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBClusterDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description of an AnalyticDB for MySQL cluster to facilitate the maintenance and management of the cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - ModifyDBClusterDescriptionRequest
//
// @return ModifyDBClusterDescriptionResponse
func (client *Client) ModifyDBClusterDescription(request *ModifyDBClusterDescriptionRequest) (_result *ModifyDBClusterDescriptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBClusterDescriptionResponse{}
	_body, _err := client.ModifyDBClusterDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the maintenance window of an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - ModifyDBClusterMaintainTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterMaintainTimeResponse
func (client *Client) ModifyDBClusterMaintainTimeWithOptions(request *ModifyDBClusterMaintainTimeRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBClusterMaintainTimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.MaintainTime) {
		query["MaintainTime"] = request.MaintainTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBClusterMaintainTime"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBClusterMaintainTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the maintenance window of an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - ModifyDBClusterMaintainTimeRequest
//
// @return ModifyDBClusterMaintainTimeResponse
func (client *Client) ModifyDBClusterMaintainTime(request *ModifyDBClusterMaintainTimeRequest) (_result *ModifyDBClusterMaintainTimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBClusterMaintainTimeResponse{}
	_body, _err := client.ModifyDBClusterMaintainTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the resource group to which an AnalyticDB for MySQL cluster belongs.
//
// @param request - ModifyDBClusterResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterResourceGroupResponse
func (client *Client) ModifyDBClusterResourceGroupWithOptions(request *ModifyDBClusterResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBClusterResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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
		Action:      dara.String("ModifyDBClusterResourceGroup"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBClusterResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the resource group to which an AnalyticDB for MySQL cluster belongs.
//
// @param request - ModifyDBClusterResourceGroupRequest
//
// @return ModifyDBClusterResourceGroupResponse
func (client *Client) ModifyDBClusterResourceGroup(request *ModifyDBClusterResourceGroupRequest) (_result *ModifyDBClusterResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBClusterResourceGroupResponse{}
	_body, _err := client.ModifyDBClusterResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the SSL link configuration of a cluster.
//
// @param request - ModifyDBClusterSSLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterSSLResponse
func (client *Client) ModifyDBClusterSSLWithOptions(request *ModifyDBClusterSSLRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBClusterSSLResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionString) {
		query["ConnectionString"] = request.ConnectionString
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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
		Action:      dara.String("ModifyDBClusterSSL"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBClusterSSLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the SSL link configuration of a cluster.
//
// @param request - ModifyDBClusterSSLRequest
//
// @return ModifyDBClusterSSLResponse
func (client *Client) ModifyDBClusterSSL(request *ModifyDBClusterSSLRequest) (_result *ModifyDBClusterSSLResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBClusterSSLResponse{}
	_body, _err := client.ModifyDBClusterSSLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the virtual IP address (VIP) that is used to connect to an AnalyticDB for MySQL cluster.
//
// @param request - ModifyDBClusterVipRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterVipResponse
func (client *Client) ModifyDBClusterVipWithOptions(request *ModifyDBClusterVipRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBClusterVipResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectString) {
		query["ConnectString"] = request.ConnectString
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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
		Action:      dara.String("ModifyDBClusterVip"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBClusterVipResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the virtual IP address (VIP) that is used to connect to an AnalyticDB for MySQL cluster.
//
// @param request - ModifyDBClusterVipRequest
//
// @return ModifyDBClusterVipResponse
func (client *Client) ModifyDBClusterVip(request *ModifyDBClusterVipRequest) (_result *ModifyDBClusterVipResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBClusterVipResponse{}
	_body, _err := client.ModifyDBClusterVipWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the amount of reserved computing resources for an AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param tmpReq - ModifyDBResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBResourceGroupResponse
func (client *Client) ModifyDBResourceGroupWithOptions(tmpReq *ModifyDBResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyDBResourceGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EngineParams) {
		request.EngineParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EngineParams, dara.String("EngineParams"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.GpuElasticPlan) {
		request.GpuElasticPlanShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GpuElasticPlan, dara.String("GpuElasticPlan"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RayConfig) {
		request.RayConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RayConfig, dara.String("RayConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Rules) {
		request.RulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Rules, dara.String("Rules"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoStopInterval) {
		query["AutoStopInterval"] = request.AutoStopInterval
	}

	if !dara.IsNil(request.ClusterMode) {
		query["ClusterMode"] = request.ClusterMode
	}

	if !dara.IsNil(request.ClusterSizeResource) {
		query["ClusterSizeResource"] = request.ClusterSizeResource
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EnableSpot) {
		query["EnableSpot"] = request.EnableSpot
	}

	if !dara.IsNil(request.EngineParamsShrink) {
		query["EngineParams"] = request.EngineParamsShrink
	}

	if !dara.IsNil(request.GpuElasticPlanShrink) {
		query["GpuElasticPlan"] = request.GpuElasticPlanShrink
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.GroupType) {
		query["GroupType"] = request.GroupType
	}

	if !dara.IsNil(request.MaxClusterCount) {
		query["MaxClusterCount"] = request.MaxClusterCount
	}

	if !dara.IsNil(request.MaxComputeResource) {
		query["MaxComputeResource"] = request.MaxComputeResource
	}

	if !dara.IsNil(request.MaxGpuQuantity) {
		query["MaxGpuQuantity"] = request.MaxGpuQuantity
	}

	if !dara.IsNil(request.MinClusterCount) {
		query["MinClusterCount"] = request.MinClusterCount
	}

	if !dara.IsNil(request.MinComputeResource) {
		query["MinComputeResource"] = request.MinComputeResource
	}

	if !dara.IsNil(request.MinGpuQuantity) {
		query["MinGpuQuantity"] = request.MinGpuQuantity
	}

	if !dara.IsNil(request.RayConfigShrink) {
		query["RayConfig"] = request.RayConfigShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RulesShrink) {
		query["Rules"] = request.RulesShrink
	}

	if !dara.IsNil(request.SpecName) {
		query["SpecName"] = request.SpecName
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TargetResourceGroupName) {
		query["TargetResourceGroupName"] = request.TargetResourceGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBResourceGroup"),
		Version:     dara.String("2021-12-01"),
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
// Modifies the amount of reserved computing resources for an AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
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
// Modifies a scaling plan for an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see Endpoints.
//
// @param request - ModifyElasticPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyElasticPlanResponse
func (client *Client) ModifyElasticPlanWithOptions(request *ModifyElasticPlanRequest, runtime *dara.RuntimeOptions) (_result *ModifyElasticPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CronExpression) {
		query["CronExpression"] = request.CronExpression
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.ElasticPlanName) {
		query["ElasticPlanName"] = request.ElasticPlanName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TargetSize) {
		query["TargetSize"] = request.TargetSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyElasticPlan"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyElasticPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a scaling plan for an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see Endpoints.
//
// @param request - ModifyElasticPlanRequest
//
// @return ModifyElasticPlanResponse
func (client *Client) ModifyElasticPlan(request *ModifyElasticPlanRequest) (_result *ModifyElasticPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyElasticPlanResponse{}
	_body, _err := client.ModifyElasticPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the disk cache size in the query acceleration configuration of an AnalyticDB for MySQL cluster.
//
// @param request - ModifyEssdCacheConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyEssdCacheConfigResponse
func (client *Client) ModifyEssdCacheConfigWithOptions(request *ModifyEssdCacheConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyEssdCacheConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EnableEssdCache) {
		query["EnableEssdCache"] = request.EnableEssdCache
	}

	if !dara.IsNil(request.EssdCacheSize) {
		query["EssdCacheSize"] = request.EssdCacheSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyEssdCacheConfig"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyEssdCacheConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the disk cache size in the query acceleration configuration of an AnalyticDB for MySQL cluster.
//
// @param request - ModifyEssdCacheConfigRequest
//
// @return ModifyEssdCacheConfigResponse
func (client *Client) ModifyEssdCacheConfig(request *ModifyEssdCacheConfigRequest) (_result *ModifyEssdCacheConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyEssdCacheConfigResponse{}
	_body, _err := client.ModifyEssdCacheConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the lake cache size of an AnalyticDB for MySQL cluster.
//
// @param request - ModifyLakeCacheSizeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyLakeCacheSizeResponse
func (client *Client) ModifyLakeCacheSizeWithOptions(request *ModifyLakeCacheSizeRequest, runtime *dara.RuntimeOptions) (_result *ModifyLakeCacheSizeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Capacity) {
		query["Capacity"] = request.Capacity
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EnableLakeCache) {
		query["EnableLakeCache"] = request.EnableLakeCache
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyLakeCacheSize"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyLakeCacheSizeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the lake cache size of an AnalyticDB for MySQL cluster.
//
// @param request - ModifyLakeCacheSizeRequest
//
// @return ModifyLakeCacheSizeResponse
func (client *Client) ModifyLakeCacheSize(request *ModifyLakeCacheSizeRequest) (_result *ModifyLakeCacheSizeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyLakeCacheSizeResponse{}
	_body, _err := client.ModifyLakeCacheSizeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies materialized views.
//
// @param request - ModifyMaterializedViewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMaterializedViewResponse
func (client *Client) ModifyMaterializedViewWithOptions(request *ModifyMaterializedViewRequest, runtime *dara.RuntimeOptions) (_result *ModifyMaterializedViewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.EnableDelayAlert) {
		query["EnableDelayAlert"] = request.EnableDelayAlert
	}

	if !dara.IsNil(request.EnableFailureAlert) {
		query["EnableFailureAlert"] = request.EnableFailureAlert
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.LatencyTolerance) {
		query["LatencyTolerance"] = request.LatencyTolerance
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.QueryWrite) {
		query["QueryWrite"] = request.QueryWrite
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

	if !dara.IsNil(request.ViewName) {
		query["ViewName"] = request.ViewName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyMaterializedView"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyMaterializedViewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies materialized views.
//
// @param request - ModifyMaterializedViewRequest
//
// @return ModifyMaterializedViewResponse
func (client *Client) ModifyMaterializedView(request *ModifyMaterializedViewRequest) (_result *ModifyMaterializedViewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyMaterializedViewResponse{}
	_body, _err := client.ModifyMaterializedViewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a materialized view recommendation task.
//
// @param request - ModifyMaterializedViewRecommendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMaterializedViewRecommendResponse
func (client *Client) ModifyMaterializedViewRecommendWithOptions(request *ModifyMaterializedViewRecommendRequest, runtime *dara.RuntimeOptions) (_result *ModifyMaterializedViewRecommendResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.MinRewriteQueryCount) {
		query["MinRewriteQueryCount"] = request.MinRewriteQueryCount
	}

	if !dara.IsNil(request.MinRewriteQueryPattern) {
		query["MinRewriteQueryPattern"] = request.MinRewriteQueryPattern
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

	if !dara.IsNil(request.ScanQueriesRange) {
		query["ScanQueriesRange"] = request.ScanQueriesRange
	}

	if !dara.IsNil(request.SchedulingDay) {
		query["SchedulingDay"] = request.SchedulingDay
	}

	if !dara.IsNil(request.SchedulingPolicy) {
		query["SchedulingPolicy"] = request.SchedulingPolicy
	}

	if !dara.IsNil(request.SlowQueryThreshold) {
		query["SlowQueryThreshold"] = request.SlowQueryThreshold
	}

	if !dara.IsNil(request.SpecifiedTime) {
		query["SpecifiedTime"] = request.SpecifiedTime
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyMaterializedViewRecommend"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyMaterializedViewRecommendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a materialized view recommendation task.
//
// @param request - ModifyMaterializedViewRecommendRequest
//
// @return ModifyMaterializedViewRecommendResponse
func (client *Client) ModifyMaterializedViewRecommend(request *ModifyMaterializedViewRecommendRequest) (_result *ModifyMaterializedViewRecommendResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyMaterializedViewRecommendResponse{}
	_body, _err := client.ModifyMaterializedViewRecommendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the information about a custom monitoring view.
//
// @param tmpReq - ModifyPerformanceViewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPerformanceViewResponse
func (client *Client) ModifyPerformanceViewWithOptions(tmpReq *ModifyPerformanceViewRequest, runtime *dara.RuntimeOptions) (_result *ModifyPerformanceViewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyPerformanceViewShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ViewDetail) {
		request.ViewDetailShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ViewDetail, dara.String("ViewDetail"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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

	if !dara.IsNil(request.ViewDetailShrink) {
		query["ViewDetail"] = request.ViewDetailShrink
	}

	if !dara.IsNil(request.ViewName) {
		query["ViewName"] = request.ViewName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyPerformanceView"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyPerformanceViewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about a custom monitoring view.
//
// @param request - ModifyPerformanceViewRequest
//
// @return ModifyPerformanceViewResponse
func (client *Client) ModifyPerformanceView(request *ModifyPerformanceViewRequest) (_result *ModifyPerformanceViewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyPerformanceViewResponse{}
	_body, _err := client.ModifyPerformanceViewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the directory location of SQL templates.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - ModifySqlTemplatePositionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySqlTemplatePositionResponse
func (client *Client) ModifySqlTemplatePositionWithOptions(request *ModifySqlTemplatePositionRequest, runtime *dara.RuntimeOptions) (_result *ModifySqlTemplatePositionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TargetTemplateGroupId) {
		query["TargetTemplateGroupId"] = request.TargetTemplateGroupId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySqlTemplatePosition"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySqlTemplatePositionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the directory location of SQL templates.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - ModifySqlTemplatePositionRequest
//
// @return ModifySqlTemplatePositionResponse
func (client *Client) ModifySqlTemplatePosition(request *ModifySqlTemplatePositionRequest) (_result *ModifySqlTemplatePositionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifySqlTemplatePositionResponse{}
	_body, _err := client.ModifySqlTemplatePositionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the vSwitches that are connected to elastic network interfaces (ENIs).
//
// @param request - ModifyUserEniVswitchOptionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyUserEniVswitchOptionsResponse
func (client *Client) ModifyUserEniVswitchOptionsWithOptions(request *ModifyUserEniVswitchOptionsRequest, runtime *dara.RuntimeOptions) (_result *ModifyUserEniVswitchOptionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbClusterId) {
		query["DbClusterId"] = request.DbClusterId
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.VSwitchOptions) {
		body["VSwitchOptions"] = request.VSwitchOptions
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyUserEniVswitchOptions"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyUserEniVswitchOptionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the vSwitches that are connected to elastic network interfaces (ENIs).
//
// @param request - ModifyUserEniVswitchOptionsRequest
//
// @return ModifyUserEniVswitchOptionsResponse
func (client *Client) ModifyUserEniVswitchOptions(request *ModifyUserEniVswitchOptionsRequest) (_result *ModifyUserEniVswitchOptionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyUserEniVswitchOptionsResponse{}
	_body, _err := client.ModifyUserEniVswitchOptionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Preloads metrics for a Spark application.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - PreloadSparkAppMetricsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PreloadSparkAppMetricsResponse
func (client *Client) PreloadSparkAppMetricsWithOptions(request *PreloadSparkAppMetricsRequest, runtime *dara.RuntimeOptions) (_result *PreloadSparkAppMetricsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PreloadSparkAppMetrics"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PreloadSparkAppMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Preloads metrics for a Spark application.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - PreloadSparkAppMetricsRequest
//
// @return PreloadSparkAppMetricsResponse
func (client *Client) PreloadSparkAppMetrics(request *PreloadSparkAppMetricsRequest) (_result *PreloadSparkAppMetricsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PreloadSparkAppMetricsResponse{}
	_body, _err := client.PreloadSparkAppMetricsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases the public endpoint of an AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - ReleaseClusterPublicConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseClusterPublicConnectionResponse
func (client *Client) ReleaseClusterPublicConnectionWithOptions(request *ReleaseClusterPublicConnectionRequest, runtime *dara.RuntimeOptions) (_result *ReleaseClusterPublicConnectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseClusterPublicConnection"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseClusterPublicConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases the public endpoint of an AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - ReleaseClusterPublicConnectionRequest
//
// @return ReleaseClusterPublicConnectionResponse
func (client *Client) ReleaseClusterPublicConnection(request *ReleaseClusterPublicConnectionRequest) (_result *ReleaseClusterPublicConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReleaseClusterPublicConnectionResponse{}
	_body, _err := client.ReleaseClusterPublicConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resets the password of a database account for an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
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
	if !dara.IsNil(request.AccountDescription) {
		query["AccountDescription"] = request.AccountDescription
	}

	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.AccountPassword) {
		query["AccountPassword"] = request.AccountPassword
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetAccountPassword"),
		Version:     dara.String("2021-12-01"),
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
// Resets the password of a database account for an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
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
// Revokes permissions from the service account of an AnalyticDB for MySQL cluster.
//
// @param request - RevokeOperatorPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeOperatorPermissionResponse
func (client *Client) RevokeOperatorPermissionWithOptions(request *RevokeOperatorPermissionRequest, runtime *dara.RuntimeOptions) (_result *RevokeOperatorPermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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
		Action:      dara.String("RevokeOperatorPermission"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeOperatorPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes permissions from the service account of an AnalyticDB for MySQL cluster.
//
// @param request - RevokeOperatorPermissionRequest
//
// @return RevokeOperatorPermissionResponse
func (client *Client) RevokeOperatorPermission(request *RevokeOperatorPermissionRequest) (_result *RevokeOperatorPermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RevokeOperatorPermissionResponse{}
	_body, _err := client.RevokeOperatorPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the Spark log configuration.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - SetSparkAppLogRootPathRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetSparkAppLogRootPathResponse
func (client *Client) SetSparkAppLogRootPathWithOptions(request *SetSparkAppLogRootPathRequest, runtime *dara.RuntimeOptions) (_result *SetSparkAppLogRootPathResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.OssLogPath) {
		body["OssLogPath"] = request.OssLogPath
	}

	if !dara.IsNil(request.UseDefaultOss) {
		body["UseDefaultOss"] = request.UseDefaultOss
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetSparkAppLogRootPath"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetSparkAppLogRootPathResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the Spark log configuration.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - SetSparkAppLogRootPathRequest
//
// @return SetSparkAppLogRootPathResponse
func (client *Client) SetSparkAppLogRootPath(request *SetSparkAppLogRootPathRequest) (_result *SetSparkAppLogRootPathResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetSparkAppLogRootPathResponse{}
	_body, _err := client.SetSparkAppLogRootPathWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts an AnalyticDB Pipeline Service (APS) job.
//
// @param request - StartApsJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartApsJobResponse
func (client *Client) StartApsJobWithOptions(request *StartApsJobRequest, runtime *dara.RuntimeOptions) (_result *StartApsJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApsJobId) {
		body["ApsJobId"] = request.ApsJobId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartApsJob"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartApsJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts an AnalyticDB Pipeline Service (APS) job.
//
// @param request - StartApsJobRequest
//
// @return StartApsJobResponse
func (client *Client) StartApsJob(request *StartApsJobRequest) (_result *StartApsJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartApsJobResponse{}
	_body, _err := client.StartApsJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts a Spark session.
//
// @param request - StartSparkReplSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartSparkReplSessionResponse
func (client *Client) StartSparkReplSessionWithOptions(request *StartSparkReplSessionRequest, runtime *dara.RuntimeOptions) (_result *StartSparkReplSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["Config"] = request.Config
	}

	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.ResourceGroupName) {
		body["ResourceGroupName"] = request.ResourceGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartSparkReplSession"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartSparkReplSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a Spark session.
//
// @param request - StartSparkReplSessionRequest
//
// @return StartSparkReplSessionResponse
func (client *Client) StartSparkReplSession(request *StartSparkReplSessionRequest) (_result *StartSparkReplSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartSparkReplSessionResponse{}
	_body, _err := client.StartSparkReplSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI StartSparkSQLEngine is deprecated
//
// Summary:
//
// Starts the Spark SQL engine.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - StartSparkSQLEngineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartSparkSQLEngineResponse
func (client *Client) StartSparkSQLEngineWithOptions(request *StartSparkSQLEngineRequest, runtime *dara.RuntimeOptions) (_result *StartSparkSQLEngineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["Config"] = request.Config
	}

	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Jars) {
		body["Jars"] = request.Jars
	}

	if !dara.IsNil(request.MaxExecutor) {
		body["MaxExecutor"] = request.MaxExecutor
	}

	if !dara.IsNil(request.MinExecutor) {
		body["MinExecutor"] = request.MinExecutor
	}

	if !dara.IsNil(request.ResourceGroupName) {
		body["ResourceGroupName"] = request.ResourceGroupName
	}

	if !dara.IsNil(request.SlotNum) {
		body["SlotNum"] = request.SlotNum
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartSparkSQLEngine"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartSparkSQLEngineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI StartSparkSQLEngine is deprecated
//
// Summary:
//
// Starts the Spark SQL engine.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - StartSparkSQLEngineRequest
//
// @return StartSparkSQLEngineResponse
// Deprecated
func (client *Client) StartSparkSQLEngine(request *StartSparkSQLEngineRequest) (_result *StartSparkSQLEngineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartSparkSQLEngineResponse{}
	_body, _err := client.StartSparkSQLEngineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits an SQL query and exports a result set.
//
// @param request - SubmitResultExportJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitResultExportJobResponse
func (client *Client) SubmitResultExportJobWithOptions(request *SubmitResultExportJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitResultExportJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Engine) {
		body["Engine"] = request.Engine
	}

	if !dara.IsNil(request.ExportType) {
		body["ExportType"] = request.ExportType
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroup) {
		body["ResourceGroup"] = request.ResourceGroup
	}

	if !dara.IsNil(request.SQL) {
		body["SQL"] = request.SQL
	}

	if !dara.IsNil(request.Schema) {
		body["Schema"] = request.Schema
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitResultExportJob"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitResultExportJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits an SQL query and exports a result set.
//
// @param request - SubmitResultExportJobRequest
//
// @return SubmitResultExportJobResponse
func (client *Client) SubmitResultExportJob(request *SubmitResultExportJobRequest) (_result *SubmitResultExportJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitResultExportJobResponse{}
	_body, _err := client.SubmitResultExportJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a Spark application.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - SubmitSparkAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitSparkAppResponse
func (client *Client) SubmitSparkAppWithOptions(request *SubmitSparkAppRequest, runtime *dara.RuntimeOptions) (_result *SubmitSparkAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentSource) {
		body["AgentSource"] = request.AgentSource
	}

	if !dara.IsNil(request.AgentVersion) {
		body["AgentVersion"] = request.AgentVersion
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Data) {
		body["Data"] = request.Data
	}

	if !dara.IsNil(request.ResourceGroupName) {
		body["ResourceGroupName"] = request.ResourceGroupName
	}

	if !dara.IsNil(request.TemplateFileId) {
		body["TemplateFileId"] = request.TemplateFileId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitSparkApp"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitSparkAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a Spark application.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - SubmitSparkAppRequest
//
// @return SubmitSparkAppResponse
func (client *Client) SubmitSparkApp(request *SubmitSparkAppRequest) (_result *SubmitSparkAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitSparkAppResponse{}
	_body, _err := client.SubmitSparkAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a Spark log analysis task and queries the analysis results.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - SubmitSparkLogAnalyzeTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitSparkLogAnalyzeTaskResponse
func (client *Client) SubmitSparkLogAnalyzeTaskWithOptions(request *SubmitSparkLogAnalyzeTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitSparkLogAnalyzeTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitSparkLogAnalyzeTask"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitSparkLogAnalyzeTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a Spark log analysis task and queries the analysis results.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - SubmitSparkLogAnalyzeTaskRequest
//
// @return SubmitSparkLogAnalyzeTaskResponse
func (client *Client) SubmitSparkLogAnalyzeTask(request *SubmitSparkLogAnalyzeTaskRequest) (_result *SubmitSparkLogAnalyzeTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitSparkLogAnalyzeTaskResponse{}
	_body, _err := client.SubmitSparkLogAnalyzeTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Suspends an AnalyticDB Pipeline Service (APS) job.
//
// @param request - SuspendApsJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SuspendApsJobResponse
func (client *Client) SuspendApsJobWithOptions(request *SuspendApsJobRequest, runtime *dara.RuntimeOptions) (_result *SuspendApsJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApsJobId) {
		body["ApsJobId"] = request.ApsJobId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SuspendApsJob"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SuspendApsJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Suspends an AnalyticDB Pipeline Service (APS) job.
//
// @param request - SuspendApsJobRequest
//
// @return SuspendApsJobResponse
func (client *Client) SuspendApsJob(request *SuspendApsJobRequest) (_result *SuspendApsJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SuspendApsJobResponse{}
	_body, _err := client.SuspendApsJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disassociates a standard account of an AnalyticDB for MySQL cluster from a Resource Access Management (RAM) user.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - UnbindAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindAccountResponse
func (client *Client) UnbindAccountWithOptions(request *UnbindAccountRequest, runtime *dara.RuntimeOptions) (_result *UnbindAccountResponse, _err error) {
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

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindAccount"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates a standard account of an AnalyticDB for MySQL cluster from a Resource Access Management (RAM) user.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - UnbindAccountRequest
//
// @return UnbindAccountResponse
func (client *Client) UnbindAccount(request *UnbindAccountRequest) (_result *UnbindAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnbindAccountResponse{}
	_body, _err := client.UnbindAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disassociates resource groups from database accounts for an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - UnbindDBResourceGroupWithUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindDBResourceGroupWithUserResponse
func (client *Client) UnbindDBResourceGroupWithUserWithOptions(request *UnbindDBResourceGroupWithUserRequest, runtime *dara.RuntimeOptions) (_result *UnbindDBResourceGroupWithUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.GroupUser) {
		query["GroupUser"] = request.GroupUser
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindDBResourceGroupWithUser"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindDBResourceGroupWithUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates resource groups from database accounts for an AnalyticDB for MySQL cluster.
//
// Description:
//
// For information about the endpoints of AnalyticDB for MySQL, see [Endpoints](https://help.aliyun.com/document_detail/612373.html).
//
// @param request - UnbindDBResourceGroupWithUserRequest
//
// @return UnbindDBResourceGroupWithUserResponse
func (client *Client) UnbindDBResourceGroupWithUser(request *UnbindDBResourceGroupWithUserRequest) (_result *UnbindDBResourceGroupWithUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnbindDBResourceGroupWithUserResponse{}
	_body, _err := client.UnbindDBResourceGroupWithUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the webhook configuration of a specified cluster.
//
// @param tmpReq - UpdateApsWebhookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApsWebhookResponse
func (client *Client) UpdateApsWebhookWithOptions(tmpReq *UpdateApsWebhookRequest, runtime *dara.RuntimeOptions) (_result *UpdateApsWebhookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateApsWebhookShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Webhook) {
		request.WebhookShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Webhook, dara.String("Webhook"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WebhookShrink) {
		body["Webhook"] = request.WebhookShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateApsWebhook"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApsWebhookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the webhook configuration of a specified cluster.
//
// @param request - UpdateApsWebhookRequest
//
// @return UpdateApsWebhookResponse
func (client *Client) UpdateApsWebhook(request *UpdateApsWebhookRequest) (_result *UpdateApsWebhookResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateApsWebhookResponse{}
	_body, _err := client.UpdateApsWebhookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a lake storage.
//
// @param tmpReq - UpdateLakeStorageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLakeStorageResponse
func (client *Client) UpdateLakeStorageWithOptions(tmpReq *UpdateLakeStorageRequest, runtime *dara.RuntimeOptions) (_result *UpdateLakeStorageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateLakeStorageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Permissions) {
		request.PermissionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Permissions, dara.String("Permissions"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.LakeStorageId) {
		body["LakeStorageId"] = request.LakeStorageId
	}

	if !dara.IsNil(request.PermissionsShrink) {
		body["Permissions"] = request.PermissionsShrink
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLakeStorage"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLakeStorageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a lake storage.
//
// @param request - UpdateLakeStorageRequest
//
// @return UpdateLakeStorageResponse
func (client *Client) UpdateLakeStorage(request *UpdateLakeStorageRequest) (_result *UpdateLakeStorageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateLakeStorageResponse{}
	_body, _err := client.UpdateLakeStorageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a Spark application template.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - UpdateSparkTemplateFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSparkTemplateFileResponse
func (client *Client) UpdateSparkTemplateFileWithOptions(request *UpdateSparkTemplateFileRequest, runtime *dara.RuntimeOptions) (_result *UpdateSparkTemplateFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.DBClusterId) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.ResourceGroupName) {
		body["ResourceGroupName"] = request.ResourceGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSparkTemplateFile"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSparkTemplateFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a Spark application template.
//
// Description:
//
//	  Regional public endpoint: `adb.<region-id>.aliyuncs.com`. Example: `adb.cn-hangzhou.aliyuncs.com`.
//
//		- Regional Virtual Private Cloud (VPC) endpoint: `adb-vpc.<region-id>.aliyuncs.com`. Example: `adb-vpc.cn-hangzhou.aliyuncs.com`.
//
// >  If HTTP status code 409 is returned when you call this operation in the China (Qingdao), China (Shenzhen), China (Guangzhou), or China (Hong Kong) region, contact technical support.
//
// @param request - UpdateSparkTemplateFileRequest
//
// @return UpdateSparkTemplateFileResponse
func (client *Client) UpdateSparkTemplateFile(request *UpdateSparkTemplateFileRequest) (_result *UpdateSparkTemplateFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSparkTemplateFileResponse{}
	_body, _err := client.UpdateSparkTemplateFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the minor version of an AnalyticDB for MySQL cluster.
//
// @param request - UpgradeKernelVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeKernelVersionResponse
func (client *Client) UpgradeKernelVersionWithOptions(request *UpgradeKernelVersionRequest, runtime *dara.RuntimeOptions) (_result *UpgradeKernelVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBVersion) {
		query["DBVersion"] = request.DBVersion
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

	if !dara.IsNil(request.SwitchMode) {
		query["SwitchMode"] = request.SwitchMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeKernelVersion"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeKernelVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the minor version of an AnalyticDB for MySQL cluster.
//
// @param request - UpgradeKernelVersionRequest
//
// @return UpgradeKernelVersionResponse
func (client *Client) UpgradeKernelVersion(request *UpgradeKernelVersionRequest) (_result *UpgradeKernelVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpgradeKernelVersionResponse{}
	_body, _err := client.UpgradeKernelVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) describeLLMAnswerWithSSE_opYieldFunc(_yield chan *DescribeLLMAnswerResponse, _yieldErr chan error, request *DescribeLLMAnswerRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
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
		Action:      dara.String("DescribeLLMAnswer"),
		Version:     dara.String("2021-12-01"),
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
