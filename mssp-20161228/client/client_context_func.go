// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// # Confirm Receipt of Security Assessment Report
//
// @param request - ConfirmDjbhReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfirmDjbhReportResponse
func (client *Client) ConfirmDjbhReportWithContext(ctx context.Context, request *ConfirmDjbhReportRequest, runtime *dara.RuntimeOptions) (_result *ConfirmDjbhReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfirmDjbhReport"),
		Version:     dara.String("2016-12-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfirmDjbhReportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Service-Linked Role
//
// @param request - CreateServiceLinkedRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceLinkedRoleResponse
func (client *Client) CreateServiceLinkedRoleWithContext(ctx context.Context, request *CreateServiceLinkedRoleRequest, runtime *dara.RuntimeOptions) (_result *CreateServiceLinkedRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
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
		Action:      dara.String("CreateServiceLinkedRole"),
		Version:     dara.String("2016-12-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateServiceLinkedRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Service Work Order
//
// @param request - CreateServiceWorkOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceWorkOrderResponse
func (client *Client) CreateServiceWorkOrderWithContext(ctx context.Context, request *CreateServiceWorkOrderRequest, runtime *dara.RuntimeOptions) (_result *CreateServiceWorkOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Creator) {
		body["Creator"] = request.Creator
	}

	if !dara.IsNil(request.CustomerId) {
		body["CustomerId"] = request.CustomerId
	}

	if !dara.IsNil(request.DurationDay) {
		body["DurationDay"] = request.DurationDay
	}

	if !dara.IsNil(request.IsAttachment) {
		body["IsAttachment"] = request.IsAttachment
	}

	if !dara.IsNil(request.IsMilestone) {
		body["IsMilestone"] = request.IsMilestone
	}

	if !dara.IsNil(request.IsWorkOrderNotify) {
		body["IsWorkOrderNotify"] = request.IsWorkOrderNotify
	}

	if !dara.IsNil(request.NotifyDay) {
		body["NotifyDay"] = request.NotifyDay
	}

	if !dara.IsNil(request.NotifyId) {
		body["NotifyId"] = request.NotifyId
	}

	if !dara.IsNil(request.OperateRemark) {
		body["OperateRemark"] = request.OperateRemark
	}

	if !dara.IsNil(request.OperateType) {
		body["OperateType"] = request.OperateType
	}

	if !dara.IsNil(request.Operator) {
		body["Operator"] = request.Operator
	}

	if !dara.IsNil(request.OwnerId) {
		body["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.WorkOrderDetail) {
		body["WorkOrderDetail"] = request.WorkOrderDetail
	}

	if !dara.IsNil(request.WorkOrderName) {
		body["WorkOrderName"] = request.WorkOrderName
	}

	if !dara.IsNil(request.WorkOrderSource) {
		body["WorkOrderSource"] = request.WorkOrderSource
	}

	if !dara.IsNil(request.WorkOrderStatus) {
		body["WorkOrderStatus"] = request.WorkOrderStatus
	}

	if !dara.IsNil(request.WorkOrderType) {
		body["WorkOrderType"] = request.WorkOrderType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateServiceWorkOrder"),
		Version:     dara.String("2016-12-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateServiceWorkOrderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Security Assessment Report
//
// @param request - DeleteDjbhReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDjbhReportResponse
func (client *Client) DeleteDjbhReportWithContext(ctx context.Context, request *DeleteDjbhReportRequest, runtime *dara.RuntimeOptions) (_result *DeleteDjbhReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDjbhReport"),
		Version:     dara.String("2016-12-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDjbhReportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Service-Linked Role
//
// @param request - DescribeServiceLinkedRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceLinkedRoleResponse
func (client *Client) DescribeServiceLinkedRoleWithContext(ctx context.Context, request *DescribeServiceLinkedRoleRequest, runtime *dara.RuntimeOptions) (_result *DescribeServiceLinkedRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
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
		Action:      dara.String("DescribeServiceLinkedRole"),
		Version:     dara.String("2016-12-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeServiceLinkedRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Process Service Work Order
//
// @param request - DisposeServiceWorkOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisposeServiceWorkOrderResponse
func (client *Client) DisposeServiceWorkOrderWithContext(ctx context.Context, request *DisposeServiceWorkOrderRequest, runtime *dara.RuntimeOptions) (_result *DisposeServiceWorkOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AttachmentName) {
		body["AttachmentName"] = request.AttachmentName
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ForwardOwnerId) {
		body["ForwardOwnerId"] = request.ForwardOwnerId
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.IsAttachment) {
		body["IsAttachment"] = request.IsAttachment
	}

	if !dara.IsNil(request.IsWorkOrderNotify) {
		body["IsWorkOrderNotify"] = request.IsWorkOrderNotify
	}

	if !dara.IsNil(request.NotifyId) {
		body["NotifyId"] = request.NotifyId
	}

	if !dara.IsNil(request.OperateRemark) {
		body["OperateRemark"] = request.OperateRemark
	}

	if !dara.IsNil(request.OperateType) {
		body["OperateType"] = request.OperateType
	}

	if !dara.IsNil(request.Operator) {
		body["Operator"] = request.Operator
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.UpgradeOwnerId) {
		body["UpgradeOwnerId"] = request.UpgradeOwnerId
	}

	if !dara.IsNil(request.WorkOrderDetail) {
		body["WorkOrderDetail"] = request.WorkOrderDetail
	}

	if !dara.IsNil(request.WorkOrderName) {
		body["WorkOrderName"] = request.WorkOrderName
	}

	if !dara.IsNil(request.WorkOrderStatus) {
		body["WorkOrderStatus"] = request.WorkOrderStatus
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisposeServiceWorkOrder"),
		Version:     dara.String("2016-12-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisposeServiceWorkOrderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Handle Alert Work Order
//
// @param tmpReq - DisposeWorkTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisposeWorkTaskResponse
func (client *Client) DisposeWorkTaskWithContext(ctx context.Context, tmpReq *DisposeWorkTaskRequest, runtime *dara.RuntimeOptions) (_result *DisposeWorkTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DisposeWorkTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.WorkTaskAnalysisResults) {
		request.WorkTaskAnalysisResultsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WorkTaskAnalysisResults, dara.String("WorkTaskAnalysisResults"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Operator) {
		body["Operator"] = request.Operator
	}

	if !dara.IsNil(request.OptRemark) {
		body["OptRemark"] = request.OptRemark
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskIds) {
		body["TaskIds"] = request.TaskIds
	}

	if !dara.IsNil(request.WorkTaskAnalysisResultsShrink) {
		body["WorkTaskAnalysisResults"] = request.WorkTaskAnalysisResultsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisposeWorkTask"),
		Version:     dara.String("2016-12-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisposeWorkTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Alarm Details
//
// @param request - GetAlarmDetailByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAlarmDetailByIdResponse
func (client *Client) GetAlarmDetailByIdWithContext(ctx context.Context, request *GetAlarmDetailByIdRequest, runtime *dara.RuntimeOptions) (_result *GetAlarmDetailByIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAlarmDetailById"),
		Version:     dara.String("2016-12-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAlarmDetailByIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Trend of Attacked Asset Convergence
//
// @param request - GetAttackedAssetDealRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAttackedAssetDealResponse
func (client *Client) GetAttackedAssetDealWithContext(ctx context.Context, request *GetAttackedAssetDealRequest, runtime *dara.RuntimeOptions) (_result *GetAttackedAssetDealResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DateType) {
		body["DateType"] = request.DateType
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.SuspEventSource) {
		body["SuspEventSource"] = request.SuspEventSource
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAttackedAssetDeal"),
		Version:     dara.String("2016-12-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAttackedAssetDealResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Compliance Risk Convergence Trend
//
// @param request - GetBaselineSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBaselineSummaryResponse
func (client *Client) GetBaselineSummaryWithContext(ctx context.Context, request *GetBaselineSummaryRequest, runtime *dara.RuntimeOptions) (_result *GetBaselineSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DateType) {
		body["DateType"] = request.DateType
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.SuspEventSource) {
		body["SuspEventSource"] = request.SuspEventSource
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBaselineSummary"),
		Version:     dara.String("2016-12-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBaselineSummaryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Console Score
//
// @param request - GetConsoleScoreRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConsoleScoreResponse
func (client *Client) GetConsoleScoreWithContext(ctx context.Context, request *GetConsoleScoreRequest, runtime *dara.RuntimeOptions) (_result *GetConsoleScoreResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DateType) {
		body["DateType"] = request.DateType
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.SuspEventSource) {
		body["SuspEventSource"] = request.SuspEventSource
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetConsoleScore"),
		Version:     dara.String("2016-12-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetConsoleScoreResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Risk Details
//
// @param request - GetDetailByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDetailByIdResponse
func (client *Client) GetDetailByIdWithContext(ctx context.Context, request *GetDetailByIdRequest, runtime *dara.RuntimeOptions) (_result *GetDetailByIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDetailById"),
		Version:     dara.String("2016-12-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDetailByIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Single Service Report Download
//
// @param request - GetDocumentDownloadUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentDownloadUrlResponse
func (client *Client) GetDocumentDownloadUrlWithContext(ctx context.Context, request *GetDocumentDownloadUrlRequest, runtime *dara.RuntimeOptions) (_result *GetDocumentDownloadUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileKey) {
		body["FileKey"] = request.FileKey
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.ReportType) {
		body["ReportType"] = request.ReportType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDocumentDownloadUrl"),
		Version:     dara.String("2016-12-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDocumentDownloadUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Service Report Query
//
// @param request - GetDocumentPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentPageResponse
func (client *Client) GetDocumentPageWithContext(ctx context.Context, request *GetDocumentPageRequest, runtime *dara.RuntimeOptions) (_result *GetDocumentPageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DeliveredBy) {
		body["DeliveredBy"] = request.DeliveredBy
	}

	if !dara.IsNil(request.DocumentName) {
		body["DocumentName"] = request.DocumentName
	}

	if !dara.IsNil(request.DocumentType) {
		body["DocumentType"] = request.DocumentType
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ReportType) {
		body["ReportType"] = request.ReportType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDocumentPage"),
		Version:     dara.String("2016-12-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDocumentPageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Service Report Home Page Statistics Acquisition
//
// @param request - GetDocumentSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentSummaryResponse
func (client *Client) GetDocumentSummaryWithContext(ctx context.Context, request *GetDocumentSummaryRequest, runtime *dara.RuntimeOptions) (_result *GetDocumentSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ReportType) {
		body["ReportType"] = request.ReportType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDocumentSummary"),
		Version:     dara.String("2016-12-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDocumentSummaryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Recently Uploaded Service Reports
//
// @param request - GetRecentDocumentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRecentDocumentResponse
func (client *Client) GetRecentDocumentWithContext(ctx context.Context, request *GetRecentDocumentRequest, runtime *dara.RuntimeOptions) (_result *GetRecentDocumentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DateType) {
		body["DateType"] = request.DateType
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.SuspEventSource) {
		body["SuspEventSource"] = request.SuspEventSource
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRecentDocument"),
		Version:     dara.String("2016-12-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRecentDocumentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Safety Coverage
//
// @param request - GetSafetyCoverRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSafetyCoverResponse
func (client *Client) GetSafetyCoverWithContext(ctx context.Context, request *GetSafetyCoverRequest, runtime *dara.RuntimeOptions) (_result *GetSafetyCoverResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DateType) {
		body["DateType"] = request.DateType
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.SuspEventSource) {
		body["SuspEventSource"] = request.SuspEventSource
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSafetyCover"),
		Version:     dara.String("2016-12-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSafetyCoverResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get SOW List
//
// @param request - GetSowListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSowListResponse
func (client *Client) GetSowListWithContext(ctx context.Context, request *GetSowListRequest, runtime *dara.RuntimeOptions) (_result *GetSowListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DateType) {
		body["DateType"] = request.DateType
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.SuspEventSource) {
		body["SuspEventSource"] = request.SuspEventSource
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSowList"),
		Version:     dara.String("2016-12-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSowListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Alarm Disposal Query
//
// @param request - GetSuspEventPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSuspEventPageResponse
func (client *Client) GetSuspEventPageWithContext(ctx context.Context, request *GetSuspEventPageRequest, runtime *dara.RuntimeOptions) (_result *GetSuspEventPageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AlarmEndTime) {
		body["AlarmEndTime"] = request.AlarmEndTime
	}

	if !dara.IsNil(request.AlarmStartTime) {
		body["AlarmStartTime"] = request.AlarmStartTime
	}

	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Source) {
		body["Source"] = request.Source
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSuspEventPage"),
		Version:     dara.String("2016-12-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSuspEventPageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Alert Statistics
//
// @param request - GetSuspEventSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSuspEventSummaryResponse
func (client *Client) GetSuspEventSummaryWithContext(ctx context.Context, request *GetSuspEventSummaryRequest, runtime *dara.RuntimeOptions) (_result *GetSuspEventSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DateType) {
		body["DateType"] = request.DateType
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.SuspEventSource) {
		body["SuspEventSource"] = request.SuspEventSource
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSuspEventSummary"),
		Version:     dara.String("2016-12-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSuspEventSummaryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Risk Query
//
// @param request - GetVulItemPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVulItemPageResponse
func (client *Client) GetVulItemPageWithContext(ctx context.Context, request *GetVulItemPageRequest, runtime *dara.RuntimeOptions) (_result *GetVulItemPageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AliasName) {
		body["AliasName"] = request.AliasName
	}

	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Dealed) {
		body["Dealed"] = request.Dealed
	}

	if !dara.IsNil(request.Level) {
		body["Level"] = request.Level
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ScanType) {
		body["ScanType"] = request.ScanType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVulItemPage"),
		Version:     dara.String("2016-12-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVulItemPageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query processed details
//
// @param request - GetVulListByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVulListByIdResponse
func (client *Client) GetVulListByIdWithContext(ctx context.Context, request *GetVulListByIdRequest, runtime *dara.RuntimeOptions) (_result *GetVulListByIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Dealed) {
		body["Dealed"] = request.Dealed
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Necessity) {
		body["Necessity"] = request.Necessity
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Remark) {
		body["Remark"] = request.Remark
	}

	if !dara.IsNil(request.Uuids) {
		body["Uuids"] = request.Uuids
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVulListById"),
		Version:     dara.String("2016-12-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVulListByIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Risk Statistics
//
// @param request - GetVulSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVulSummaryResponse
func (client *Client) GetVulSummaryWithContext(ctx context.Context, request *GetVulSummaryRequest, runtime *dara.RuntimeOptions) (_result *GetVulSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DateType) {
		body["DateType"] = request.DateType
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.SuspEventSource) {
		body["SuspEventSource"] = request.SuspEventSource
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVulSummary"),
		Version:     dara.String("2016-12-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVulSummaryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get the First Line Work Order Statistics
//
// @param request - GetWorkTaskSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkTaskSummaryResponse
func (client *Client) GetWorkTaskSummaryWithContext(ctx context.Context, request *GetWorkTaskSummaryRequest, runtime *dara.RuntimeOptions) (_result *GetWorkTaskSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DateType) {
		body["DateType"] = request.DateType
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.SuspEventSource) {
		body["SuspEventSource"] = request.SuspEventSource
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWorkTaskSummary"),
		Version:     dara.String("2016-12-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkTaskSummaryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Service Customer Information Query
//
// @param request - PageServiceCustomerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PageServiceCustomerResponse
func (client *Client) PageServiceCustomerWithContext(ctx context.Context, request *PageServiceCustomerRequest, runtime *dara.RuntimeOptions) (_result *PageServiceCustomerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AuthStatus) {
		body["AuthStatus"] = request.AuthStatus
	}

	if !dara.IsNil(request.CmAuthStatus) {
		body["CmAuthStatus"] = request.CmAuthStatus
	}

	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.MonitorAuthStatus) {
		body["MonitorAuthStatus"] = request.MonitorAuthStatus
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PageServiceCustomer"),
		Version:     dara.String("2016-12-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PageServiceCustomerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Send Custom Alert Event
//
// @param request - SendCustomEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendCustomEventResponse
func (client *Client) SendCustomEventWithContext(ctx context.Context, request *SendCustomEventRequest, runtime *dara.RuntimeOptions) (_result *SendCustomEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CustomerId) {
		body["CustomerId"] = request.CustomerId
	}

	if !dara.IsNil(request.DataSource) {
		body["DataSource"] = request.DataSource
	}

	if !dara.IsNil(request.EventDescription) {
		body["EventDescription"] = request.EventDescription
	}

	if !dara.IsNil(request.EventDetails) {
		body["EventDetails"] = request.EventDetails
	}

	if !dara.IsNil(request.EventMarkdown) {
		body["EventMarkdown"] = request.EventMarkdown
	}

	if !dara.IsNil(request.EventName) {
		body["EventName"] = request.EventName
	}

	if !dara.IsNil(request.EventType) {
		body["EventType"] = request.EventType
	}

	if !dara.IsNil(request.FindTime) {
		body["FindTime"] = request.FindTime
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceName) {
		body["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.IsSend) {
		body["IsSend"] = request.IsSend
	}

	if !dara.IsNil(request.Level) {
		body["Level"] = request.Level
	}

	if !dara.IsNil(request.OccurrenceTime) {
		body["OccurrenceTime"] = request.OccurrenceTime
	}

	if !dara.IsNil(request.OwnerId) {
		body["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Product) {
		body["Product"] = request.Product
	}

	if !dara.IsNil(request.UniqueId) {
		body["UniqueId"] = request.UniqueId
	}

	if !dara.IsNil(request.Uuid) {
		body["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendCustomEvent"),
		Version:     dara.String("2016-12-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendCustomEventResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
