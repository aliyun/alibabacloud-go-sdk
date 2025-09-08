// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

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
func (client *Client) AllocateClusterPublicConnectionWithContext(ctx context.Context, request *AllocateClusterPublicConnectionRequest, runtime *dara.RuntimeOptions) (_result *AllocateClusterPublicConnectionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyAdviceByIdResponse
func (client *Client) ApplyAdviceByIdWithContext(ctx context.Context, request *ApplyAdviceByIdRequest, runtime *dara.RuntimeOptions) (_result *ApplyAdviceByIdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachUserENIResponse
func (client *Client) AttachUserENIWithContext(ctx context.Context, request *AttachUserENIRequest, runtime *dara.RuntimeOptions) (_result *AttachUserENIResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchApplyAdviceByIdListResponse
func (client *Client) BatchApplyAdviceByIdListWithContext(ctx context.Context, request *BatchApplyAdviceByIdListRequest, runtime *dara.RuntimeOptions) (_result *BatchApplyAdviceByIdListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAccountResponse
func (client *Client) BindAccountWithContext(ctx context.Context, request *BindAccountRequest, runtime *dara.RuntimeOptions) (_result *BindAccountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindDBResourceGroupWithUserResponse
func (client *Client) BindDBResourceGroupWithUserWithContext(ctx context.Context, request *BindDBResourceGroupWithUserRequest, runtime *dara.RuntimeOptions) (_result *BindDBResourceGroupWithUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelSparkReplStatementResponse
func (client *Client) CancelSparkReplStatementWithContext(ctx context.Context, request *CancelSparkReplStatementRequest, runtime *dara.RuntimeOptions) (_result *CancelSparkReplStatementResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelSparkWarehouseBatchSQLResponse
func (client *Client) CancelSparkWarehouseBatchSQLWithContext(ctx context.Context, request *CancelSparkWarehouseBatchSQLRequest, runtime *dara.RuntimeOptions) (_result *CancelSparkWarehouseBatchSQLResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckBindRamUserResponse
func (client *Client) CheckBindRamUserWithContext(ctx context.Context, request *CheckBindRamUserRequest, runtime *dara.RuntimeOptions) (_result *CheckBindRamUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckSampleDataSetResponse
func (client *Client) CheckSampleDataSetWithContext(ctx context.Context, request *CheckSampleDataSetRequest, runtime *dara.RuntimeOptions) (_result *CheckSampleDataSetResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAPSJobResponse
func (client *Client) CreateAPSJobWithContext(ctx context.Context, request *CreateAPSJobRequest, runtime *dara.RuntimeOptions) (_result *CreateAPSJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccountResponse
func (client *Client) CreateAccountWithContext(ctx context.Context, request *CreateAccountRequest, runtime *dara.RuntimeOptions) (_result *CreateAccountResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApsCopyWorkloadResponse
func (client *Client) CreateApsCopyWorkloadWithContext(ctx context.Context, request *CreateApsCopyWorkloadRequest, runtime *dara.RuntimeOptions) (_result *CreateApsCopyWorkloadResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - CreateApsDatasoureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApsDatasoureResponse
func (client *Client) CreateApsDatasoureWithContext(ctx context.Context, tmpReq *CreateApsDatasoureRequest, runtime *dara.RuntimeOptions) (_result *CreateApsDatasoureResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApsHiveJobResponse
func (client *Client) CreateApsHiveJobWithContext(ctx context.Context, request *CreateApsHiveJobRequest, runtime *dara.RuntimeOptions) (_result *CreateApsHiveJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建Kafka到Huid的APS链路
//
// @param tmpReq - CreateApsKafkaHudiJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApsKafkaHudiJobResponse
func (client *Client) CreateApsKafkaHudiJobWithContext(ctx context.Context, tmpReq *CreateApsKafkaHudiJobRequest, runtime *dara.RuntimeOptions) (_result *CreateApsKafkaHudiJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - CreateApsSlsADBJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApsSlsADBJobResponse
func (client *Client) CreateApsSlsADBJobWithContext(ctx context.Context, tmpReq *CreateApsSlsADBJobRequest, runtime *dara.RuntimeOptions) (_result *CreateApsSlsADBJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDBClusterResponse
func (client *Client) CreateDBClusterWithContext(ctx context.Context, request *CreateDBClusterRequest, runtime *dara.RuntimeOptions) (_result *CreateDBClusterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - CreateDBResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDBResourceGroupResponse
func (client *Client) CreateDBResourceGroupWithContext(ctx context.Context, tmpReq *CreateDBResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateDBResourceGroupResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateDBResourceGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EngineParams) {
		request.EngineParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EngineParams, dara.String("EngineParams"), dara.String("json"))
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateElasticPlanResponse
func (client *Client) CreateElasticPlanWithContext(ctx context.Context, request *CreateElasticPlanRequest, runtime *dara.RuntimeOptions) (_result *CreateElasticPlanResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - CreateLakeStorageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLakeStorageResponse
func (client *Client) CreateLakeStorageWithContext(ctx context.Context, tmpReq *CreateLakeStorageRequest, runtime *dara.RuntimeOptions) (_result *CreateLakeStorageResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOssSubDirectoryResponse
func (client *Client) CreateOssSubDirectoryWithContext(ctx context.Context, request *CreateOssSubDirectoryRequest, runtime *dara.RuntimeOptions) (_result *CreateOssSubDirectoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - CreatePerformanceViewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePerformanceViewResponse
func (client *Client) CreatePerformanceViewWithContext(ctx context.Context, tmpReq *CreatePerformanceViewRequest, runtime *dara.RuntimeOptions) (_result *CreatePerformanceViewResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSparkTemplateResponse
func (client *Client) CreateSparkTemplateWithContext(ctx context.Context, request *CreateSparkTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateSparkTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccountResponse
func (client *Client) DeleteAccountWithContext(ctx context.Context, request *DeleteAccountRequest, runtime *dara.RuntimeOptions) (_result *DeleteAccountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApsDatasoureResponse
func (client *Client) DeleteApsDatasoureWithContext(ctx context.Context, request *DeleteApsDatasoureRequest, runtime *dara.RuntimeOptions) (_result *DeleteApsDatasoureResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApsJobResponse
func (client *Client) DeleteApsJobWithContext(ctx context.Context, request *DeleteApsJobRequest, runtime *dara.RuntimeOptions) (_result *DeleteApsJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBackupsResponse
func (client *Client) DeleteBackupsWithContext(ctx context.Context, request *DeleteBackupsRequest, runtime *dara.RuntimeOptions) (_result *DeleteBackupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDBClusterResponse
func (client *Client) DeleteDBClusterWithContext(ctx context.Context, request *DeleteDBClusterRequest, runtime *dara.RuntimeOptions) (_result *DeleteDBClusterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDBResourceGroupResponse
func (client *Client) DeleteDBResourceGroupWithContext(ctx context.Context, request *DeleteDBResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteDBResourceGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteElasticPlanResponse
func (client *Client) DeleteElasticPlanWithContext(ctx context.Context, request *DeleteElasticPlanRequest, runtime *dara.RuntimeOptions) (_result *DeleteElasticPlanResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLakeStorageResponse
func (client *Client) DeleteLakeStorageWithContext(ctx context.Context, request *DeleteLakeStorageRequest, runtime *dara.RuntimeOptions) (_result *DeleteLakeStorageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePerformanceViewResponse
func (client *Client) DeletePerformanceViewWithContext(ctx context.Context, request *DeletePerformanceViewRequest, runtime *dara.RuntimeOptions) (_result *DeletePerformanceViewResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSparkTemplateResponse
func (client *Client) DeleteSparkTemplateWithContext(ctx context.Context, request *DeleteSparkTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteSparkTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSparkTemplateFileResponse
func (client *Client) DeleteSparkTemplateFileWithContext(ctx context.Context, request *DeleteSparkTemplateFileRequest, runtime *dara.RuntimeOptions) (_result *DeleteSparkTemplateFileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAPSADBInstancesResponse
func (client *Client) DescribeAPSADBInstancesWithContext(ctx context.Context, request *DescribeAPSADBInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAPSADBInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAbnormalPatternDetectionResponse
func (client *Client) DescribeAbnormalPatternDetectionWithContext(ctx context.Context, request *DescribeAbnormalPatternDetectionRequest, runtime *dara.RuntimeOptions) (_result *DescribeAbnormalPatternDetectionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccountAllPrivilegesResponse
func (client *Client) DescribeAccountAllPrivilegesWithContext(ctx context.Context, request *DescribeAccountAllPrivilegesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccountAllPrivilegesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccountPrivilegeObjectsResponse
func (client *Client) DescribeAccountPrivilegeObjectsWithContext(ctx context.Context, request *DescribeAccountPrivilegeObjectsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccountPrivilegeObjectsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccountPrivilegesResponse
func (client *Client) DescribeAccountPrivilegesWithContext(ctx context.Context, request *DescribeAccountPrivilegesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccountPrivilegesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccountsResponse
func (client *Client) DescribeAccountsWithContext(ctx context.Context, request *DescribeAccountsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccountsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAdbMySqlColumnsResponse
func (client *Client) DescribeAdbMySqlColumnsWithContext(ctx context.Context, request *DescribeAdbMySqlColumnsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAdbMySqlColumnsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAdbMySqlIndexesResponse
func (client *Client) DescribeAdbMySqlIndexesWithContext(ctx context.Context, request *DescribeAdbMySqlIndexesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAdbMySqlIndexesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAdbMySqlSchemasResponse
func (client *Client) DescribeAdbMySqlSchemasWithContext(ctx context.Context, request *DescribeAdbMySqlSchemasRequest, runtime *dara.RuntimeOptions) (_result *DescribeAdbMySqlSchemasResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAdbMySqlTableMetaResponse
func (client *Client) DescribeAdbMySqlTableMetaWithContext(ctx context.Context, request *DescribeAdbMySqlTableMetaRequest, runtime *dara.RuntimeOptions) (_result *DescribeAdbMySqlTableMetaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAdbMySqlTablesResponse
func (client *Client) DescribeAdbMySqlTablesWithContext(ctx context.Context, request *DescribeAdbMySqlTablesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAdbMySqlTablesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAdviceServiceEnabledResponse
func (client *Client) DescribeAdviceServiceEnabledWithContext(ctx context.Context, request *DescribeAdviceServiceEnabledRequest, runtime *dara.RuntimeOptions) (_result *DescribeAdviceServiceEnabledResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAllDataSourceResponse
func (client *Client) DescribeAllDataSourceWithContext(ctx context.Context, request *DescribeAllDataSourceRequest, runtime *dara.RuntimeOptions) (_result *DescribeAllDataSourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppliedAdvicesResponse
func (client *Client) DescribeAppliedAdvicesWithContext(ctx context.Context, request *DescribeAppliedAdvicesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppliedAdvicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApsActionLogsResponse
func (client *Client) DescribeApsActionLogsWithContext(ctx context.Context, request *DescribeApsActionLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeApsActionLogsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApsDatasourceResponse
func (client *Client) DescribeApsDatasourceWithContext(ctx context.Context, request *DescribeApsDatasourceRequest, runtime *dara.RuntimeOptions) (_result *DescribeApsDatasourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApsDatasourcesResponse
func (client *Client) DescribeApsDatasourcesWithContext(ctx context.Context, request *DescribeApsDatasourcesRequest, runtime *dara.RuntimeOptions) (_result *DescribeApsDatasourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApsHiveWorkloadResponse
func (client *Client) DescribeApsHiveWorkloadWithContext(ctx context.Context, request *DescribeApsHiveWorkloadRequest, runtime *dara.RuntimeOptions) (_result *DescribeApsHiveWorkloadResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApsJobDetailResponse
func (client *Client) DescribeApsJobDetailWithContext(ctx context.Context, request *DescribeApsJobDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeApsJobDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApsJobsResponse
func (client *Client) DescribeApsJobsWithContext(ctx context.Context, request *DescribeApsJobsRequest, runtime *dara.RuntimeOptions) (_result *DescribeApsJobsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApsMigrationWorkloadsResponse
func (client *Client) DescribeApsMigrationWorkloadsWithContext(ctx context.Context, request *DescribeApsMigrationWorkloadsRequest, runtime *dara.RuntimeOptions) (_result *DescribeApsMigrationWorkloadsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApsProgressResponse
func (client *Client) DescribeApsProgressWithContext(ctx context.Context, request *DescribeApsProgressRequest, runtime *dara.RuntimeOptions) (_result *DescribeApsProgressResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApsResourceGroupsResponse
func (client *Client) DescribeApsResourceGroupsWithContext(ctx context.Context, request *DescribeApsResourceGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeApsResourceGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAuditLogRecordsResponse
func (client *Client) DescribeAuditLogRecordsWithContext(ctx context.Context, request *DescribeAuditLogRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAuditLogRecordsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAvailableAdvicesResponse
func (client *Client) DescribeAvailableAdvicesWithContext(ctx context.Context, request *DescribeAvailableAdvicesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAvailableAdvicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupPolicyResponse
func (client *Client) DescribeBackupPolicyWithContext(ctx context.Context, request *DescribeBackupPolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupsResponse
func (client *Client) DescribeBackupsWithContext(ctx context.Context, request *DescribeBackupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBadSqlDetectionResponse
func (client *Client) DescribeBadSqlDetectionWithContext(ctx context.Context, request *DescribeBadSqlDetectionRequest, runtime *dara.RuntimeOptions) (_result *DescribeBadSqlDetectionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterAccessWhiteListResponse
func (client *Client) DescribeClusterAccessWhiteListWithContext(ctx context.Context, request *DescribeClusterAccessWhiteListRequest, runtime *dara.RuntimeOptions) (_result *DescribeClusterAccessWhiteListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterNetInfoResponse
func (client *Client) DescribeClusterNetInfoWithContext(ctx context.Context, request *DescribeClusterNetInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeClusterNetInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterResourceDetailResponse
func (client *Client) DescribeClusterResourceDetailWithContext(ctx context.Context, request *DescribeClusterResourceDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeClusterResourceDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterResourceUsageResponse
func (client *Client) DescribeClusterResourceUsageWithContext(ctx context.Context, request *DescribeClusterResourceUsageRequest, runtime *dara.RuntimeOptions) (_result *DescribeClusterResourceUsageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeColumnsResponse
func (client *Client) DescribeColumnsWithContext(ctx context.Context, request *DescribeColumnsRequest, runtime *dara.RuntimeOptions) (_result *DescribeColumnsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCompactionServiceSwitchResponse
func (client *Client) DescribeCompactionServiceSwitchWithContext(ctx context.Context, request *DescribeCompactionServiceSwitchRequest, runtime *dara.RuntimeOptions) (_result *DescribeCompactionServiceSwitchResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeComputeResourceUsageResponse
func (client *Client) DescribeComputeResourceUsageWithContext(ctx context.Context, request *DescribeComputeResourceUsageRequest, runtime *dara.RuntimeOptions) (_result *DescribeComputeResourceUsageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeControllerDetectionResponse
func (client *Client) DescribeControllerDetectionWithContext(ctx context.Context, request *DescribeControllerDetectionRequest, runtime *dara.RuntimeOptions) (_result *DescribeControllerDetectionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterAttributeResponse
func (client *Client) DescribeDBClusterAttributeWithContext(ctx context.Context, request *DescribeDBClusterAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClusterAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterHealthStatusResponse
func (client *Client) DescribeDBClusterHealthStatusWithContext(ctx context.Context, request *DescribeDBClusterHealthStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClusterHealthStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterPerformanceResponse
func (client *Client) DescribeDBClusterPerformanceWithContext(ctx context.Context, request *DescribeDBClusterPerformanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClusterPerformanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取SSL配置信息
//
// @param request - DescribeDBClusterSSLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterSSLResponse
func (client *Client) DescribeDBClusterSSLWithContext(ctx context.Context, request *DescribeDBClusterSSLRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClusterSSLResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterSpaceSummaryResponse
func (client *Client) DescribeDBClusterSpaceSummaryWithContext(ctx context.Context, request *DescribeDBClusterSpaceSummaryRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClusterSpaceSummaryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterStatusResponse
func (client *Client) DescribeDBClusterStatusWithContext(ctx context.Context, request *DescribeDBClusterStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClusterStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClustersResponse
func (client *Client) DescribeDBClustersWithContext(ctx context.Context, request *DescribeDBClustersRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClustersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBResourceGroupResponse
func (client *Client) DescribeDBResourceGroupWithContext(ctx context.Context, request *DescribeDBResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBResourceGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDiagnosisDimensionsResponse
func (client *Client) DescribeDiagnosisDimensionsWithContext(ctx context.Context, request *DescribeDiagnosisDimensionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDiagnosisDimensionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDiagnosisRecordsResponse
func (client *Client) DescribeDiagnosisRecordsWithContext(ctx context.Context, request *DescribeDiagnosisRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDiagnosisRecordsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDiagnosisSQLInfoResponse
func (client *Client) DescribeDiagnosisSQLInfoWithContext(ctx context.Context, request *DescribeDiagnosisSQLInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDiagnosisSQLInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDownloadRecordsResponse
func (client *Client) DescribeDownloadRecordsWithContext(ctx context.Context, request *DescribeDownloadRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDownloadRecordsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeElasticPlanAttributeResponse
func (client *Client) DescribeElasticPlanAttributeWithContext(ctx context.Context, request *DescribeElasticPlanAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeElasticPlanAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeElasticPlanJobsResponse
func (client *Client) DescribeElasticPlanJobsWithContext(ctx context.Context, request *DescribeElasticPlanJobsRequest, runtime *dara.RuntimeOptions) (_result *DescribeElasticPlanJobsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeElasticPlanSpecificationsResponse
func (client *Client) DescribeElasticPlanSpecificationsWithContext(ctx context.Context, request *DescribeElasticPlanSpecificationsRequest, runtime *dara.RuntimeOptions) (_result *DescribeElasticPlanSpecificationsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeElasticPlansResponse
func (client *Client) DescribeElasticPlansWithContext(ctx context.Context, request *DescribeElasticPlansRequest, runtime *dara.RuntimeOptions) (_result *DescribeElasticPlansResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEnabledPrivilegesResponse
func (client *Client) DescribeEnabledPrivilegesWithContext(ctx context.Context, request *DescribeEnabledPrivilegesRequest, runtime *dara.RuntimeOptions) (_result *DescribeEnabledPrivilegesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEssdCacheConfigResponse
func (client *Client) DescribeEssdCacheConfigWithContext(ctx context.Context, request *DescribeEssdCacheConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeEssdCacheConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExcessivePrimaryKeysResponse
func (client *Client) DescribeExcessivePrimaryKeysWithContext(ctx context.Context, request *DescribeExcessivePrimaryKeysRequest, runtime *dara.RuntimeOptions) (_result *DescribeExcessivePrimaryKeysResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExecutorDetectionResponse
func (client *Client) DescribeExecutorDetectionWithContext(ctx context.Context, request *DescribeExecutorDetectionRequest, runtime *dara.RuntimeOptions) (_result *DescribeExecutorDetectionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 任务中心任务列表
//
// @param request - DescribeHistoryTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHistoryTasksResponse
func (client *Client) DescribeHistoryTasksWithContext(ctx context.Context, request *DescribeHistoryTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeHistoryTasksResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 任务中心任务统计
//
// @param request - DescribeHistoryTasksStatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHistoryTasksStatResponse
func (client *Client) DescribeHistoryTasksStatWithContext(ctx context.Context, request *DescribeHistoryTasksStatRequest, runtime *dara.RuntimeOptions) (_result *DescribeHistoryTasksStatResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInclinedNodesResponse
func (client *Client) DescribeInclinedNodesWithContext(ctx context.Context, request *DescribeInclinedNodesRequest, runtime *dara.RuntimeOptions) (_result *DescribeInclinedNodesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInclinedTablesResponse
func (client *Client) DescribeInclinedTablesWithContext(ctx context.Context, request *DescribeInclinedTablesRequest, runtime *dara.RuntimeOptions) (_result *DescribeInclinedTablesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeJobResourceUsageResponse
func (client *Client) DescribeJobResourceUsageWithContext(ctx context.Context, request *DescribeJobResourceUsageRequest, runtime *dara.RuntimeOptions) (_result *DescribeJobResourceUsageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeKernelVersionResponse
func (client *Client) DescribeKernelVersionWithContext(ctx context.Context, request *DescribeKernelVersionRequest, runtime *dara.RuntimeOptions) (_result *DescribeKernelVersionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLLMAnswerResponse
func (client *Client) DescribeLLMAnswerWithContext(ctx context.Context, request *DescribeLLMAnswerRequest, runtime *dara.RuntimeOptions) (_result *DescribeLLMAnswerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLLMSimilarQuestionsResponse
func (client *Client) DescribeLLMSimilarQuestionsWithContext(ctx context.Context, request *DescribeLLMSimilarQuestionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeLLMSimilarQuestionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLakeCacheSizeResponse
func (client *Client) DescribeLakeCacheSizeWithContext(ctx context.Context, request *DescribeLakeCacheSizeRequest, runtime *dara.RuntimeOptions) (_result *DescribeLakeCacheSizeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOperatorPermissionResponse
func (client *Client) DescribeOperatorPermissionWithContext(ctx context.Context, request *DescribeOperatorPermissionRequest, runtime *dara.RuntimeOptions) (_result *DescribeOperatorPermissionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOversizeNonPartitionTableInfosResponse
func (client *Client) DescribeOversizeNonPartitionTableInfosWithContext(ctx context.Context, request *DescribeOversizeNonPartitionTableInfosRequest, runtime *dara.RuntimeOptions) (_result *DescribeOversizeNonPartitionTableInfosResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePatternPerformanceResponse
func (client *Client) DescribePatternPerformanceWithContext(ctx context.Context, request *DescribePatternPerformanceRequest, runtime *dara.RuntimeOptions) (_result *DescribePatternPerformanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePerformanceViewAttributeResponse
func (client *Client) DescribePerformanceViewAttributeWithContext(ctx context.Context, request *DescribePerformanceViewAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribePerformanceViewAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePerformanceViewsResponse
func (client *Client) DescribePerformanceViewsWithContext(ctx context.Context, request *DescribePerformanceViewsRequest, runtime *dara.RuntimeOptions) (_result *DescribePerformanceViewsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithContext(ctx context.Context, request *DescribeRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResourceGroupSpecResponse
func (client *Client) DescribeResourceGroupSpecWithContext(ctx context.Context, request *DescribeResourceGroupSpecRequest, runtime *dara.RuntimeOptions) (_result *DescribeResourceGroupSpecResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSQLPatternsResponse
func (client *Client) DescribeSQLPatternsWithContext(ctx context.Context, request *DescribeSQLPatternsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSQLPatternsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSQLWebSocketDomainResponse
func (client *Client) DescribeSQLWebSocketDomainWithContext(ctx context.Context, request *DescribeSQLWebSocketDomainRequest, runtime *dara.RuntimeOptions) (_result *DescribeSQLWebSocketDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSchemasResponse
func (client *Client) DescribeSchemasWithContext(ctx context.Context, request *DescribeSchemasRequest, runtime *dara.RuntimeOptions) (_result *DescribeSchemasResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSparkAppDiagnosisInfoResponse
func (client *Client) DescribeSparkAppDiagnosisInfoWithContext(ctx context.Context, request *DescribeSparkAppDiagnosisInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeSparkAppDiagnosisInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSparkAppTypeResponse
func (client *Client) DescribeSparkAppTypeWithContext(ctx context.Context, request *DescribeSparkAppTypeRequest, runtime *dara.RuntimeOptions) (_result *DescribeSparkAppTypeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSparkCodeLogResponse
func (client *Client) DescribeSparkCodeLogWithContext(ctx context.Context, request *DescribeSparkCodeLogRequest, runtime *dara.RuntimeOptions) (_result *DescribeSparkCodeLogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSparkCodeOutputResponse
func (client *Client) DescribeSparkCodeOutputWithContext(ctx context.Context, request *DescribeSparkCodeOutputRequest, runtime *dara.RuntimeOptions) (_result *DescribeSparkCodeOutputResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSparkCodeWebUiResponse
func (client *Client) DescribeSparkCodeWebUiWithContext(ctx context.Context, request *DescribeSparkCodeWebUiRequest, runtime *dara.RuntimeOptions) (_result *DescribeSparkCodeWebUiResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSparkSQLDiagnosisAttributeResponse
func (client *Client) DescribeSparkSQLDiagnosisAttributeWithContext(ctx context.Context, request *DescribeSparkSQLDiagnosisAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeSparkSQLDiagnosisAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSparkSQLDiagnosisListResponse
func (client *Client) DescribeSparkSQLDiagnosisListWithContext(ctx context.Context, request *DescribeSparkSQLDiagnosisListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSparkSQLDiagnosisListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSqlPatternResponse
func (client *Client) DescribeSqlPatternWithContext(ctx context.Context, request *DescribeSqlPatternRequest, runtime *dara.RuntimeOptions) (_result *DescribeSqlPatternResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeStorageResourceUsageResponse
func (client *Client) DescribeStorageResourceUsageWithContext(ctx context.Context, request *DescribeStorageResourceUsageRequest, runtime *dara.RuntimeOptions) (_result *DescribeStorageResourceUsageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTableAccessCountResponse
func (client *Client) DescribeTableAccessCountWithContext(ctx context.Context, request *DescribeTableAccessCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeTableAccessCountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTableDetailResponse
func (client *Client) DescribeTableDetailWithContext(ctx context.Context, request *DescribeTableDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeTableDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTablePartitionDiagnoseResponse
func (client *Client) DescribeTablePartitionDiagnoseWithContext(ctx context.Context, request *DescribeTablePartitionDiagnoseRequest, runtime *dara.RuntimeOptions) (_result *DescribeTablePartitionDiagnoseResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTableStatisticsResponse
func (client *Client) DescribeTableStatisticsWithContext(ctx context.Context, request *DescribeTableStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeTableStatisticsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTablesResponse
func (client *Client) DescribeTablesWithContext(ctx context.Context, request *DescribeTablesRequest, runtime *dara.RuntimeOptions) (_result *DescribeTablesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserQuotaResponse
func (client *Client) DescribeUserQuotaWithContext(ctx context.Context, request *DescribeUserQuotaRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserQuotaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWorkerDetectionResponse
func (client *Client) DescribeWorkerDetectionWithContext(ctx context.Context, request *DescribeWorkerDetectionRequest, runtime *dara.RuntimeOptions) (_result *DescribeWorkerDetectionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachUserENIResponse
func (client *Client) DetachUserENIWithContext(ctx context.Context, request *DetachUserENIRequest, runtime *dara.RuntimeOptions) (_result *DetachUserENIResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableAdviceServiceResponse
func (client *Client) DisableAdviceServiceWithContext(ctx context.Context, request *DisableAdviceServiceRequest, runtime *dara.RuntimeOptions) (_result *DisableAdviceServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableElasticPlanResponse
func (client *Client) DisableElasticPlanWithContext(ctx context.Context, request *DisableElasticPlanRequest, runtime *dara.RuntimeOptions) (_result *DisableElasticPlanResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DownloadDiagnosisRecordsResponse
func (client *Client) DownloadDiagnosisRecordsWithContext(ctx context.Context, request *DownloadDiagnosisRecordsRequest, runtime *dara.RuntimeOptions) (_result *DownloadDiagnosisRecordsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DownloadInstanceCACertificateResponse
func (client *Client) DownloadInstanceCACertificateWithContext(ctx context.Context, request *DownloadInstanceCACertificateRequest, runtime *dara.RuntimeOptions) (_result *DownloadInstanceCACertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableAdviceServiceResponse
func (client *Client) EnableAdviceServiceWithContext(ctx context.Context, request *EnableAdviceServiceRequest, runtime *dara.RuntimeOptions) (_result *EnableAdviceServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableElasticPlanResponse
func (client *Client) EnableElasticPlanWithContext(ctx context.Context, request *EnableElasticPlanRequest, runtime *dara.RuntimeOptions) (_result *EnableElasticPlanResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteSparkReplStatementResponse
func (client *Client) ExecuteSparkReplStatementWithContext(ctx context.Context, request *ExecuteSparkReplStatementRequest, runtime *dara.RuntimeOptions) (_result *ExecuteSparkReplStatementResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteSparkWarehouseBatchSQLResponse
func (client *Client) ExecuteSparkWarehouseBatchSQLWithContext(ctx context.Context, request *ExecuteSparkWarehouseBatchSQLRequest, runtime *dara.RuntimeOptions) (_result *ExecuteSparkWarehouseBatchSQLResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExistRunningSQLEngineResponse
func (client *Client) ExistRunningSQLEngineWithContext(ctx context.Context, request *ExistRunningSQLEngineRequest, runtime *dara.RuntimeOptions) (_result *ExistRunningSQLEngineResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Spark权限说明,失败时给出配置权限的帮助信息
//
// @param request - GetADBSparkNecessaryRAMPermissionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetADBSparkNecessaryRAMPermissionsResponse
func (client *Client) GetADBSparkNecessaryRAMPermissionsWithContext(ctx context.Context, request *GetADBSparkNecessaryRAMPermissionsRequest, runtime *dara.RuntimeOptions) (_result *GetADBSparkNecessaryRAMPermissionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApsManagedDatabasesResponse
func (client *Client) GetApsManagedDatabasesWithContext(ctx context.Context, request *GetApsManagedDatabasesRequest, runtime *dara.RuntimeOptions) (_result *GetApsManagedDatabasesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCreateTableSQLResponse
func (client *Client) GetCreateTableSQLWithContext(ctx context.Context, request *GetCreateTableSQLRequest, runtime *dara.RuntimeOptions) (_result *GetCreateTableSQLResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatabaseObjectsResponse
func (client *Client) GetDatabaseObjectsWithContext(ctx context.Context, request *GetDatabaseObjectsRequest, runtime *dara.RuntimeOptions) (_result *GetDatabaseObjectsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLakeStorageResponse
func (client *Client) GetLakeStorageWithContext(ctx context.Context, request *GetLakeStorageRequest, runtime *dara.RuntimeOptions) (_result *GetLakeStorageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSparkAppAttemptLogResponse
func (client *Client) GetSparkAppAttemptLogWithContext(ctx context.Context, request *GetSparkAppAttemptLogRequest, runtime *dara.RuntimeOptions) (_result *GetSparkAppAttemptLogResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSparkAppInfoResponse
func (client *Client) GetSparkAppInfoWithContext(ctx context.Context, request *GetSparkAppInfoRequest, runtime *dara.RuntimeOptions) (_result *GetSparkAppInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSparkAppLogResponse
func (client *Client) GetSparkAppLogWithContext(ctx context.Context, request *GetSparkAppLogRequest, runtime *dara.RuntimeOptions) (_result *GetSparkAppLogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSparkAppMetricsResponse
func (client *Client) GetSparkAppMetricsWithContext(ctx context.Context, request *GetSparkAppMetricsRequest, runtime *dara.RuntimeOptions) (_result *GetSparkAppMetricsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSparkAppStateResponse
func (client *Client) GetSparkAppStateWithContext(ctx context.Context, request *GetSparkAppStateRequest, runtime *dara.RuntimeOptions) (_result *GetSparkAppStateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSparkAppWebUiAddressResponse
func (client *Client) GetSparkAppWebUiAddressWithContext(ctx context.Context, request *GetSparkAppWebUiAddressRequest, runtime *dara.RuntimeOptions) (_result *GetSparkAppWebUiAddressResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSparkConfigLogPathResponse
func (client *Client) GetSparkConfigLogPathWithContext(ctx context.Context, request *GetSparkConfigLogPathRequest, runtime *dara.RuntimeOptions) (_result *GetSparkConfigLogPathResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSparkLogAnalyzeTaskResponse
func (client *Client) GetSparkLogAnalyzeTaskWithContext(ctx context.Context, request *GetSparkLogAnalyzeTaskRequest, runtime *dara.RuntimeOptions) (_result *GetSparkLogAnalyzeTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSparkReplSessionResponse
func (client *Client) GetSparkReplSessionWithContext(ctx context.Context, request *GetSparkReplSessionRequest, runtime *dara.RuntimeOptions) (_result *GetSparkReplSessionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSparkReplStatementResponse
func (client *Client) GetSparkReplStatementWithContext(ctx context.Context, request *GetSparkReplStatementRequest, runtime *dara.RuntimeOptions) (_result *GetSparkReplStatementResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSparkSQLEngineStateResponse
func (client *Client) GetSparkSQLEngineStateWithContext(ctx context.Context, request *GetSparkSQLEngineStateRequest, runtime *dara.RuntimeOptions) (_result *GetSparkSQLEngineStateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSparkTemplateFileContentResponse
func (client *Client) GetSparkTemplateFileContentWithContext(ctx context.Context, request *GetSparkTemplateFileContentRequest, runtime *dara.RuntimeOptions) (_result *GetSparkTemplateFileContentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSparkTemplateFolderTreeResponse
func (client *Client) GetSparkTemplateFolderTreeWithContext(ctx context.Context, request *GetSparkTemplateFolderTreeRequest, runtime *dara.RuntimeOptions) (_result *GetSparkTemplateFolderTreeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSparkTemplateFullTreeResponse
func (client *Client) GetSparkTemplateFullTreeWithContext(ctx context.Context, request *GetSparkTemplateFullTreeRequest, runtime *dara.RuntimeOptions) (_result *GetSparkTemplateFullTreeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSparkWarehouseBatchSQLResponse
func (client *Client) GetSparkWarehouseBatchSQLWithContext(ctx context.Context, request *GetSparkWarehouseBatchSQLRequest, runtime *dara.RuntimeOptions) (_result *GetSparkWarehouseBatchSQLResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableResponse
func (client *Client) GetTableWithContext(ctx context.Context, request *GetTableRequest, runtime *dara.RuntimeOptions) (_result *GetTableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableColumnsResponse
func (client *Client) GetTableColumnsWithContext(ctx context.Context, request *GetTableColumnsRequest, runtime *dara.RuntimeOptions) (_result *GetTableColumnsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableDDLResponse
func (client *Client) GetTableDDLWithContext(ctx context.Context, request *GetTableDDLRequest, runtime *dara.RuntimeOptions) (_result *GetTableDDLResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableObjectsResponse
func (client *Client) GetTableObjectsWithContext(ctx context.Context, request *GetTableObjectsRequest, runtime *dara.RuntimeOptions) (_result *GetTableObjectsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetViewDDLResponse
func (client *Client) GetViewDDLWithContext(ctx context.Context, request *GetViewDDLRequest, runtime *dara.RuntimeOptions) (_result *GetViewDDLResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetViewObjectsResponse
func (client *Client) GetViewObjectsWithContext(ctx context.Context, request *GetViewObjectsRequest, runtime *dara.RuntimeOptions) (_result *GetViewObjectsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantOperatorPermissionResponse
func (client *Client) GrantOperatorPermissionWithContext(ctx context.Context, request *GrantOperatorPermissionRequest, runtime *dara.RuntimeOptions) (_result *GrantOperatorPermissionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return KillProcessResponse
func (client *Client) KillProcessWithContext(ctx context.Context, request *KillProcessRequest, runtime *dara.RuntimeOptions) (_result *KillProcessResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return KillSparkAppResponse
func (client *Client) KillSparkAppWithContext(ctx context.Context, request *KillSparkAppRequest, runtime *dara.RuntimeOptions) (_result *KillSparkAppResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return KillSparkLogAnalyzeTaskResponse
func (client *Client) KillSparkLogAnalyzeTaskWithContext(ctx context.Context, request *KillSparkLogAnalyzeTaskRequest, runtime *dara.RuntimeOptions) (_result *KillSparkLogAnalyzeTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return KillSparkSQLEngineResponse
func (client *Client) KillSparkSQLEngineWithContext(ctx context.Context, request *KillSparkSQLEngineRequest, runtime *dara.RuntimeOptions) (_result *KillSparkSQLEngineResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApsLifecycleStrategyResponse
func (client *Client) ListApsLifecycleStrategyWithContext(ctx context.Context, request *ListApsLifecycleStrategyRequest, runtime *dara.RuntimeOptions) (_result *ListApsLifecycleStrategyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApsOptimizationStrategyResponse
func (client *Client) ListApsOptimizationStrategyWithContext(ctx context.Context, request *ListApsOptimizationStrategyRequest, runtime *dara.RuntimeOptions) (_result *ListApsOptimizationStrategyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApsOptimizationTasksResponse
func (client *Client) ListApsOptimizationTasksWithContext(ctx context.Context, request *ListApsOptimizationTasksRequest, runtime *dara.RuntimeOptions) (_result *ListApsOptimizationTasksResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLakeStoragesResponse
func (client *Client) ListLakeStoragesWithContext(ctx context.Context, request *ListLakeStoragesRequest, runtime *dara.RuntimeOptions) (_result *ListLakeStoragesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ListResultExportJobHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResultExportJobHistoryResponse
func (client *Client) ListResultExportJobHistoryWithContext(ctx context.Context, tmpReq *ListResultExportJobHistoryRequest, runtime *dara.RuntimeOptions) (_result *ListResultExportJobHistoryResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSparkAppAttemptsResponse
func (client *Client) ListSparkAppAttemptsWithContext(ctx context.Context, request *ListSparkAppAttemptsRequest, runtime *dara.RuntimeOptions) (_result *ListSparkAppAttemptsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSparkAppsResponse
func (client *Client) ListSparkAppsWithContext(ctx context.Context, request *ListSparkAppsRequest, runtime *dara.RuntimeOptions) (_result *ListSparkAppsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSparkLogAnalyzeTasksResponse
func (client *Client) ListSparkLogAnalyzeTasksWithContext(ctx context.Context, request *ListSparkLogAnalyzeTasksRequest, runtime *dara.RuntimeOptions) (_result *ListSparkLogAnalyzeTasksResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSparkTemplateFileIdsResponse
func (client *Client) ListSparkTemplateFileIdsWithContext(ctx context.Context, request *ListSparkTemplateFileIdsRequest, runtime *dara.RuntimeOptions) (_result *ListSparkTemplateFileIdsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSparkWarehouseBatchSQLResponse
func (client *Client) ListSparkWarehouseBatchSQLWithContext(ctx context.Context, request *ListSparkWarehouseBatchSQLRequest, runtime *dara.RuntimeOptions) (_result *ListSparkWarehouseBatchSQLResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LoadSampleDataSetResponse
func (client *Client) LoadSampleDataSetWithContext(ctx context.Context, request *LoadSampleDataSetRequest, runtime *dara.RuntimeOptions) (_result *LoadSampleDataSetResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccountDescriptionResponse
func (client *Client) ModifyAccountDescriptionWithContext(ctx context.Context, request *ModifyAccountDescriptionRequest, runtime *dara.RuntimeOptions) (_result *ModifyAccountDescriptionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ModifyAccountPrivilegesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccountPrivilegesResponse
func (client *Client) ModifyAccountPrivilegesWithContext(ctx context.Context, tmpReq *ModifyAccountPrivilegesRequest, runtime *dara.RuntimeOptions) (_result *ModifyAccountPrivilegesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ModifyApsDatasoureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApsDatasoureResponse
func (client *Client) ModifyApsDatasoureWithContext(ctx context.Context, tmpReq *ModifyApsDatasoureRequest, runtime *dara.RuntimeOptions) (_result *ModifyApsDatasoureResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApsJobResponse
func (client *Client) ModifyApsJobWithContext(ctx context.Context, request *ModifyApsJobRequest, runtime *dara.RuntimeOptions) (_result *ModifyApsJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ModifyApsSlsADBJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApsSlsADBJobResponse
func (client *Client) ModifyApsSlsADBJobWithContext(ctx context.Context, tmpReq *ModifyApsSlsADBJobRequest, runtime *dara.RuntimeOptions) (_result *ModifyApsSlsADBJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApsWorkloadNameResponse
func (client *Client) ModifyApsWorkloadNameWithContext(ctx context.Context, request *ModifyApsWorkloadNameRequest, runtime *dara.RuntimeOptions) (_result *ModifyApsWorkloadNameResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAuditLogConfigResponse
func (client *Client) ModifyAuditLogConfigWithContext(ctx context.Context, request *ModifyAuditLogConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyAuditLogConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBackupPolicyResponse
func (client *Client) ModifyBackupPolicyWithContext(ctx context.Context, request *ModifyBackupPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyBackupPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyClickhouseEngineResponse
func (client *Client) ModifyClickhouseEngineWithContext(ctx context.Context, request *ModifyClickhouseEngineRequest, runtime *dara.RuntimeOptions) (_result *ModifyClickhouseEngineResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyClusterAccessWhiteListResponse
func (client *Client) ModifyClusterAccessWhiteListWithContext(ctx context.Context, request *ModifyClusterAccessWhiteListRequest, runtime *dara.RuntimeOptions) (_result *ModifyClusterAccessWhiteListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyClusterConnectionStringResponse
func (client *Client) ModifyClusterConnectionStringWithContext(ctx context.Context, request *ModifyClusterConnectionStringRequest, runtime *dara.RuntimeOptions) (_result *ModifyClusterConnectionStringResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCompactionServiceSwitchResponse
func (client *Client) ModifyCompactionServiceSwitchWithContext(ctx context.Context, request *ModifyCompactionServiceSwitchRequest, runtime *dara.RuntimeOptions) (_result *ModifyCompactionServiceSwitchResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterResponse
func (client *Client) ModifyDBClusterWithContext(ctx context.Context, request *ModifyDBClusterRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBClusterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterDescriptionResponse
func (client *Client) ModifyDBClusterDescriptionWithContext(ctx context.Context, request *ModifyDBClusterDescriptionRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBClusterDescriptionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterMaintainTimeResponse
func (client *Client) ModifyDBClusterMaintainTimeWithContext(ctx context.Context, request *ModifyDBClusterMaintainTimeRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBClusterMaintainTimeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterResourceGroupResponse
func (client *Client) ModifyDBClusterResourceGroupWithContext(ctx context.Context, request *ModifyDBClusterResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBClusterResourceGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 配置SSL
//
// @param request - ModifyDBClusterSSLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterSSLResponse
func (client *Client) ModifyDBClusterSSLWithContext(ctx context.Context, request *ModifyDBClusterSSLRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBClusterSSLResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterVipResponse
func (client *Client) ModifyDBClusterVipWithContext(ctx context.Context, request *ModifyDBClusterVipRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBClusterVipResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ModifyDBResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBResourceGroupResponse
func (client *Client) ModifyDBResourceGroupWithContext(ctx context.Context, tmpReq *ModifyDBResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBResourceGroupResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ModifyDBResourceGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EngineParams) {
		request.EngineParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EngineParams, dara.String("EngineParams"), dara.String("json"))
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyElasticPlanResponse
func (client *Client) ModifyElasticPlanWithContext(ctx context.Context, request *ModifyElasticPlanRequest, runtime *dara.RuntimeOptions) (_result *ModifyElasticPlanResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyEssdCacheConfigResponse
func (client *Client) ModifyEssdCacheConfigWithContext(ctx context.Context, request *ModifyEssdCacheConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyEssdCacheConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyLakeCacheSizeResponse
func (client *Client) ModifyLakeCacheSizeWithContext(ctx context.Context, request *ModifyLakeCacheSizeRequest, runtime *dara.RuntimeOptions) (_result *ModifyLakeCacheSizeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ModifyPerformanceViewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPerformanceViewResponse
func (client *Client) ModifyPerformanceViewWithContext(ctx context.Context, tmpReq *ModifyPerformanceViewRequest, runtime *dara.RuntimeOptions) (_result *ModifyPerformanceViewResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyUserEniVswitchOptionsResponse
func (client *Client) ModifyUserEniVswitchOptionsWithContext(ctx context.Context, request *ModifyUserEniVswitchOptionsRequest, runtime *dara.RuntimeOptions) (_result *ModifyUserEniVswitchOptionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PreloadSparkAppMetricsResponse
func (client *Client) PreloadSparkAppMetricsWithContext(ctx context.Context, request *PreloadSparkAppMetricsRequest, runtime *dara.RuntimeOptions) (_result *PreloadSparkAppMetricsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseClusterPublicConnectionResponse
func (client *Client) ReleaseClusterPublicConnectionWithContext(ctx context.Context, request *ReleaseClusterPublicConnectionRequest, runtime *dara.RuntimeOptions) (_result *ReleaseClusterPublicConnectionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetAccountPasswordResponse
func (client *Client) ResetAccountPasswordWithContext(ctx context.Context, request *ResetAccountPasswordRequest, runtime *dara.RuntimeOptions) (_result *ResetAccountPasswordResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeOperatorPermissionResponse
func (client *Client) RevokeOperatorPermissionWithContext(ctx context.Context, request *RevokeOperatorPermissionRequest, runtime *dara.RuntimeOptions) (_result *RevokeOperatorPermissionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetSparkAppLogRootPathResponse
func (client *Client) SetSparkAppLogRootPathWithContext(ctx context.Context, request *SetSparkAppLogRootPathRequest, runtime *dara.RuntimeOptions) (_result *SetSparkAppLogRootPathResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartApsJobResponse
func (client *Client) StartApsJobWithContext(ctx context.Context, request *StartApsJobRequest, runtime *dara.RuntimeOptions) (_result *StartApsJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartSparkReplSessionResponse
func (client *Client) StartSparkReplSessionWithContext(ctx context.Context, request *StartSparkReplSessionRequest, runtime *dara.RuntimeOptions) (_result *StartSparkReplSessionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartSparkSQLEngineResponse
func (client *Client) StartSparkSQLEngineWithContext(ctx context.Context, request *StartSparkSQLEngineRequest, runtime *dara.RuntimeOptions) (_result *StartSparkSQLEngineResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitResultExportJobResponse
func (client *Client) SubmitResultExportJobWithContext(ctx context.Context, request *SubmitResultExportJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitResultExportJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitSparkAppResponse
func (client *Client) SubmitSparkAppWithContext(ctx context.Context, request *SubmitSparkAppRequest, runtime *dara.RuntimeOptions) (_result *SubmitSparkAppResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitSparkLogAnalyzeTaskResponse
func (client *Client) SubmitSparkLogAnalyzeTaskWithContext(ctx context.Context, request *SubmitSparkLogAnalyzeTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitSparkLogAnalyzeTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SuspendApsJobResponse
func (client *Client) SuspendApsJobWithContext(ctx context.Context, request *SuspendApsJobRequest, runtime *dara.RuntimeOptions) (_result *SuspendApsJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindAccountResponse
func (client *Client) UnbindAccountWithContext(ctx context.Context, request *UnbindAccountRequest, runtime *dara.RuntimeOptions) (_result *UnbindAccountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindDBResourceGroupWithUserResponse
func (client *Client) UnbindDBResourceGroupWithUserWithContext(ctx context.Context, request *UnbindDBResourceGroupWithUserRequest, runtime *dara.RuntimeOptions) (_result *UnbindDBResourceGroupWithUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - UpdateLakeStorageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLakeStorageResponse
func (client *Client) UpdateLakeStorageWithContext(ctx context.Context, tmpReq *UpdateLakeStorageRequest, runtime *dara.RuntimeOptions) (_result *UpdateLakeStorageResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSparkTemplateFileResponse
func (client *Client) UpdateSparkTemplateFileWithContext(ctx context.Context, request *UpdateSparkTemplateFileRequest, runtime *dara.RuntimeOptions) (_result *UpdateSparkTemplateFileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeKernelVersionResponse
func (client *Client) UpgradeKernelVersionWithContext(ctx context.Context, request *UpgradeKernelVersionRequest, runtime *dara.RuntimeOptions) (_result *UpgradeKernelVersionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
