// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Adjusts the JMeter load.
//
// @param request - AdjustJMeterSceneSpeedRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AdjustJMeterSceneSpeedResponse
func (client *Client) AdjustJMeterSceneSpeedWithContext(ctx context.Context, request *AdjustJMeterSceneSpeedRequest, runtime *dara.RuntimeOptions) (_result *AdjustJMeterSceneSpeedResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ReportId) {
		query["ReportId"] = request.ReportId
	}

	if !dara.IsNil(request.Speed) {
		query["Speed"] = request.Speed
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AdjustJMeterSceneSpeed"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AdjustJMeterSceneSpeedResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adjusts the stress in a Performance Testing Service (PTS) scenario.
//
// Description:
//
// In concurrency mode, only the concurrency of the first API is passed as that of a session.
//
// In requests per second (RPS) mode, the RPS of each API can be adjusted. Make sure that the RPS decreases in the API order in the same session.
//
// @param tmpReq - AdjustPtsSceneSpeedRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AdjustPtsSceneSpeedResponse
func (client *Client) AdjustPtsSceneSpeedWithContext(ctx context.Context, tmpReq *AdjustPtsSceneSpeedRequest, runtime *dara.RuntimeOptions) (_result *AdjustPtsSceneSpeedResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AdjustPtsSceneSpeedShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApiSpeedList) {
		request.ApiSpeedListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApiSpeedList, dara.String("ApiSpeedList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiSpeedListShrink) {
		query["ApiSpeedList"] = request.ApiSpeedListShrink
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AdjustPtsSceneSpeed"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AdjustPtsSceneSpeedResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a stress testing scenario.
//
// @param request - CreatePtsSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePtsSceneResponse
func (client *Client) CreatePtsSceneWithContext(ctx context.Context, request *CreatePtsSceneRequest, runtime *dara.RuntimeOptions) (_result *CreatePtsSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Scene) {
		query["Scene"] = request.Scene
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePtsScene"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePtsSceneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// null
//
// @param request - CreatePtsSceneBaseLineFromReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePtsSceneBaseLineFromReportResponse
func (client *Client) CreatePtsSceneBaseLineFromReportWithContext(ctx context.Context, request *CreatePtsSceneBaseLineFromReportRequest, runtime *dara.RuntimeOptions) (_result *CreatePtsSceneBaseLineFromReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ReportId) {
		query["ReportId"] = request.ReportId
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePtsSceneBaseLineFromReport"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePtsSceneBaseLineFromReportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Performance Testing Service (PTS) scenario.
//
// @param request - DeletePtsSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePtsSceneResponse
func (client *Client) DeletePtsSceneWithContext(ctx context.Context, request *DeletePtsSceneRequest, runtime *dara.RuntimeOptions) (_result *DeletePtsSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePtsScene"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePtsSceneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeletePtsSceneBaseLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePtsSceneBaseLineResponse
func (client *Client) DeletePtsSceneBaseLineWithContext(ctx context.Context, request *DeletePtsSceneBaseLineRequest, runtime *dara.RuntimeOptions) (_result *DeletePtsSceneBaseLineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePtsSceneBaseLine"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePtsSceneBaseLineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param tmpReq - DeletePtsScenesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePtsScenesResponse
func (client *Client) DeletePtsScenesWithContext(ctx context.Context, tmpReq *DeletePtsScenesRequest, runtime *dara.RuntimeOptions) (_result *DeletePtsScenesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeletePtsScenesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SceneIds) {
		request.SceneIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SceneIds, dara.String("SceneIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.SceneIdsShrink) {
		query["SceneIds"] = request.SceneIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePtsScenes"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePtsScenesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the operational logs of JMeter stress testers. By default, the operational logs of the stress tester identified as number 0 are queried and the total number of stress testers is returned.
//
// @param request - GetJMeterLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJMeterLogsResponse
func (client *Client) GetJMeterLogsWithContext(ctx context.Context, request *GetJMeterLogsRequest, runtime *dara.RuntimeOptions) (_result *GetJMeterLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentIndex) {
		query["AgentIndex"] = request.AgentIndex
	}

	if !dara.IsNil(request.BeginTime) {
		query["BeginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Level) {
		query["Level"] = request.Level
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ReportId) {
		query["ReportId"] = request.ReportId
	}

	if !dara.IsNil(request.Thread) {
		query["Thread"] = request.Thread
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetJMeterLogs"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJMeterLogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details of a JMeter report.
//
// @param request - GetJMeterReportDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJMeterReportDetailsResponse
func (client *Client) GetJMeterReportDetailsWithContext(ctx context.Context, request *GetJMeterReportDetailsRequest, runtime *dara.RuntimeOptions) (_result *GetJMeterReportDetailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ReportId) {
		query["ReportId"] = request.ReportId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetJMeterReportDetails"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJMeterReportDetailsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the metrics of JMeter samplers based on specified conditions.
//
// @param request - GetJMeterSampleMetricsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJMeterSampleMetricsResponse
func (client *Client) GetJMeterSampleMetricsWithContext(ctx context.Context, request *GetJMeterSampleMetricsRequest, runtime *dara.RuntimeOptions) (_result *GetJMeterSampleMetricsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BeginTime) {
		query["BeginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ReportId) {
		query["ReportId"] = request.ReportId
	}

	if !dara.IsNil(request.SamplerId) {
		query["SamplerId"] = request.SamplerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetJMeterSampleMetrics"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJMeterSampleMetricsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the sampling logs of a JMeter sampler.
//
// @param request - GetJMeterSamplingLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJMeterSamplingLogsResponse
func (client *Client) GetJMeterSamplingLogsWithContext(ctx context.Context, request *GetJMeterSamplingLogsRequest, runtime *dara.RuntimeOptions) (_result *GetJMeterSamplingLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.BeginTime) {
		query["BeginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxRT) {
		query["MaxRT"] = request.MaxRT
	}

	if !dara.IsNil(request.MinRT) {
		query["MinRT"] = request.MinRT
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ReportId) {
		query["ReportId"] = request.ReportId
	}

	if !dara.IsNil(request.ResponseCode) {
		query["ResponseCode"] = request.ResponseCode
	}

	if !dara.IsNil(request.SamplerId) {
		query["SamplerId"] = request.SamplerId
	}

	if !dara.IsNil(request.Success) {
		query["Success"] = request.Success
	}

	if !dara.IsNil(request.Thread) {
		query["Thread"] = request.Thread
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetJMeterSamplingLogs"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJMeterSamplingLogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries data that is generated during the stress testing of a JMeter scenario based on the ID of the scenario.
//
// @param request - GetJMeterSceneRunningDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJMeterSceneRunningDataResponse
func (client *Client) GetJMeterSceneRunningDataWithContext(ctx context.Context, request *GetJMeterSceneRunningDataRequest, runtime *dara.RuntimeOptions) (_result *GetJMeterSceneRunningDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetJMeterSceneRunningData"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJMeterSceneRunningDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a JMeter scenario.
//
// @param request - GetOpenJMeterSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOpenJMeterSceneResponse
func (client *Client) GetOpenJMeterSceneWithContext(ctx context.Context, request *GetOpenJMeterSceneRequest, runtime *dara.RuntimeOptions) (_result *GetOpenJMeterSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOpenJMeterScene"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOpenJMeterSceneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the sampling logs for a Performance Testing Service (PTS) debugging task.
//
// @param request - GetPtsDebugSampleLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPtsDebugSampleLogsResponse
func (client *Client) GetPtsDebugSampleLogsWithContext(ctx context.Context, request *GetPtsDebugSampleLogsRequest, runtime *dara.RuntimeOptions) (_result *GetPtsDebugSampleLogsResponse, _err error) {
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

	if !dara.IsNil(request.PlanId) {
		query["PlanId"] = request.PlanId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPtsDebugSampleLogs"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPtsDebugSampleLogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a report for a performance testing task in a Performance Testing Service (PTS) scenario.
//
// @param request - GetPtsReportDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPtsReportDetailsResponse
func (client *Client) GetPtsReportDetailsWithContext(ctx context.Context, request *GetPtsReportDetailsRequest, runtime *dara.RuntimeOptions) (_result *GetPtsReportDetailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PlanId) {
		query["PlanId"] = request.PlanId
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPtsReportDetails"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPtsReportDetailsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all reports of multiple scenarios that are generated during the stress testing in batch.
//
// @param request - GetPtsReportsBySceneIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPtsReportsBySceneIdResponse
func (client *Client) GetPtsReportsBySceneIdWithContext(ctx context.Context, request *GetPtsReportsBySceneIdRequest, runtime *dara.RuntimeOptions) (_result *GetPtsReportsBySceneIdResponse, _err error) {
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

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPtsReportsBySceneId"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPtsReportsBySceneIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the structure and load settings of a Performance Testing Service (PTS) scenario.
//
// @param request - GetPtsSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPtsSceneResponse
func (client *Client) GetPtsSceneWithContext(ctx context.Context, request *GetPtsSceneRequest, runtime *dara.RuntimeOptions) (_result *GetPtsSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPtsScene"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPtsSceneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// null
//
// @param request - GetPtsSceneBaseLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPtsSceneBaseLineResponse
func (client *Client) GetPtsSceneBaseLineWithContext(ctx context.Context, request *GetPtsSceneBaseLineRequest, runtime *dara.RuntimeOptions) (_result *GetPtsSceneBaseLineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPtsSceneBaseLine"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPtsSceneBaseLineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the runtime data of a stress testing or debugging scenario.
//
// @param request - GetPtsSceneRunningDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPtsSceneRunningDataResponse
func (client *Client) GetPtsSceneRunningDataWithContext(ctx context.Context, request *GetPtsSceneRunningDataRequest, runtime *dara.RuntimeOptions) (_result *GetPtsSceneRunningDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PlanId) {
		query["PlanId"] = request.PlanId
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPtsSceneRunningData"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPtsSceneRunningDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the runtime status of a Performance Testing Service (PTS) scenario.
//
// @param request - GetPtsSceneRunningStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPtsSceneRunningStatusResponse
func (client *Client) GetPtsSceneRunningStatusWithContext(ctx context.Context, request *GetPtsSceneRunningStatusRequest, runtime *dara.RuntimeOptions) (_result *GetPtsSceneRunningStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPtsSceneRunningStatus"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPtsSceneRunningStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetUserVpcSecurityGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserVpcSecurityGroupResponse
func (client *Client) GetUserVpcSecurityGroupWithContext(ctx context.Context, request *GetUserVpcSecurityGroupRequest, runtime *dara.RuntimeOptions) (_result *GetUserVpcSecurityGroupResponse, _err error) {
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

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserVpcSecurityGroup"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserVpcSecurityGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetUserVpcVSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserVpcVSwitchResponse
func (client *Client) GetUserVpcVSwitchWithContext(ctx context.Context, request *GetUserVpcVSwitchRequest, runtime *dara.RuntimeOptions) (_result *GetUserVpcVSwitchResponse, _err error) {
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

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserVpcVSwitch"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserVpcVSwitchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetUserVpcsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserVpcsResponse
func (client *Client) GetUserVpcsWithContext(ctx context.Context, request *GetUserVpcsRequest, runtime *dara.RuntimeOptions) (_result *GetUserVpcsResponse, _err error) {
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

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserVpcs"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserVpcsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about JMeter environments.
//
// @param request - ListEnvsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEnvsResponse
func (client *Client) ListEnvsWithContext(ctx context.Context, request *ListEnvsRequest, runtime *dara.RuntimeOptions) (_result *ListEnvsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnvId) {
		query["EnvId"] = request.EnvId
	}

	if !dara.IsNil(request.EnvName) {
		query["EnvName"] = request.EnvName
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
		Action:      dara.String("ListEnvs"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEnvsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries JMeter reports based on specified conditions.
//
// @param request - ListJMeterReportsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJMeterReportsResponse
func (client *Client) ListJMeterReportsWithContext(ctx context.Context, request *ListJMeterReportsRequest, runtime *dara.RuntimeOptions) (_result *ListJMeterReportsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BeginTime) {
		query["BeginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
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

	if !dara.IsNil(request.ReportId) {
		query["ReportId"] = request.ReportId
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListJMeterReports"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListJMeterReportsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries JMeter scenarios based on a specified condition.
//
// @param request - ListOpenJMeterScenesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOpenJMeterScenesResponse
func (client *Client) ListOpenJMeterScenesWithContext(ctx context.Context, request *ListOpenJMeterScenesRequest, runtime *dara.RuntimeOptions) (_result *ListOpenJMeterScenesResponse, _err error) {
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

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.SceneName) {
		query["SceneName"] = request.SceneName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOpenJMeterScenes"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOpenJMeterScenesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Performance Testing Service (PTS) reports based on specified conditions.
//
// @param request - ListPtsReportsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPtsReportsResponse
func (client *Client) ListPtsReportsWithContext(ctx context.Context, request *ListPtsReportsRequest, runtime *dara.RuntimeOptions) (_result *ListPtsReportsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BeginTime) {
		body["BeginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Keyword) {
		body["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ReportId) {
		body["ReportId"] = request.ReportId
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPtsReports"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPtsReportsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Performance Testing Service (PTS) scenarios by page.
//
// @param request - ListPtsSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPtsSceneResponse
func (client *Client) ListPtsSceneWithContext(ctx context.Context, request *ListPtsSceneRequest, runtime *dara.RuntimeOptions) (_result *ListPtsSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyWord) {
		query["KeyWord"] = request.KeyWord
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
		Action:      dara.String("ListPtsScene"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPtsSceneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取vpc配置信息列表
//
// @param request - ListVpcConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVpcConfigsResponse
func (client *Client) ListVpcConfigsWithContext(ctx context.Context, request *ListVpcConfigsRequest, runtime *dara.RuntimeOptions) (_result *ListVpcConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.ConfigName) {
		query["ConfigName"] = request.ConfigName
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

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVpcConfigs"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVpcConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// null
//
// @param request - ModifyPtsSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPtsSceneResponse
func (client *Client) ModifyPtsSceneWithContext(ctx context.Context, request *ModifyPtsSceneRequest, runtime *dara.RuntimeOptions) (_result *ModifyPtsSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Scene) {
		body["Scene"] = request.Scene
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyPtsScene"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyPtsSceneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes the JMeter environment that corresponds to a specific JMeter environment ID.
//
// @param request - RemoveEnvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveEnvResponse
func (client *Client) RemoveEnvWithContext(ctx context.Context, request *RemoveEnvRequest, runtime *dara.RuntimeOptions) (_result *RemoveEnvResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnvId) {
		query["EnvId"] = request.EnvId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveEnv"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveEnvResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a JMeter scenario.
//
// @param request - RemoveOpenJMeterSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveOpenJMeterSceneResponse
func (client *Client) RemoveOpenJMeterSceneWithContext(ctx context.Context, request *RemoveOpenJMeterSceneRequest, runtime *dara.RuntimeOptions) (_result *RemoveOpenJMeterSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveOpenJMeterScene"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveOpenJMeterSceneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or updates a JMeter environment.
//
// @param tmpReq - SaveEnvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveEnvResponse
func (client *Client) SaveEnvWithContext(ctx context.Context, tmpReq *SaveEnvRequest, runtime *dara.RuntimeOptions) (_result *SaveEnvResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SaveEnvShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Env) {
		request.EnvShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Env, dara.String("Env"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EnvShrink) {
		query["Env"] = request.EnvShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveEnv"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveEnvResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or updates a JMeter scenario.
//
// @param tmpReq - SaveOpenJMeterSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveOpenJMeterSceneResponse
func (client *Client) SaveOpenJMeterSceneWithContext(ctx context.Context, tmpReq *SaveOpenJMeterSceneRequest, runtime *dara.RuntimeOptions) (_result *SaveOpenJMeterSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SaveOpenJMeterSceneShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OpenJMeterScene) {
		request.OpenJMeterSceneShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OpenJMeterScene, dara.String("OpenJMeterScene"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpenJMeterSceneShrink) {
		query["OpenJMeterScene"] = request.OpenJMeterSceneShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveOpenJMeterScene"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveOpenJMeterSceneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Saves or modifies a Performance Testing Service (PTS) scenario.
//
// @param tmpReq - SavePtsSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SavePtsSceneResponse
func (client *Client) SavePtsSceneWithContext(ctx context.Context, tmpReq *SavePtsSceneRequest, runtime *dara.RuntimeOptions) (_result *SavePtsSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SavePtsSceneShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Scene) {
		request.SceneShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Scene, dara.String("Scene"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.SceneShrink) {
		query["Scene"] = request.SceneShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SavePtsScene"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SavePtsSceneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts the debugging of a scenario to check whether the settings of the scenario take effect.
//
// @param request - StartDebugPtsSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartDebugPtsSceneResponse
func (client *Client) StartDebugPtsSceneWithContext(ctx context.Context, request *StartDebugPtsSceneRequest, runtime *dara.RuntimeOptions) (_result *StartDebugPtsSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartDebugPtsScene"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartDebugPtsSceneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Debugs a JMeter scenario.
//
// @param request - StartDebuggingJMeterSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartDebuggingJMeterSceneResponse
func (client *Client) StartDebuggingJMeterSceneWithContext(ctx context.Context, request *StartDebuggingJMeterSceneRequest, runtime *dara.RuntimeOptions) (_result *StartDebuggingJMeterSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartDebuggingJMeterScene"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartDebuggingJMeterSceneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a scenario by using its ID.
//
// @param request - StartPtsSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartPtsSceneResponse
func (client *Client) StartPtsSceneWithContext(ctx context.Context, request *StartPtsSceneRequest, runtime *dara.RuntimeOptions) (_result *StartPtsSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartPtsScene"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartPtsSceneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts performance testing in a JMeter scenario.
//
// @param request - StartTestingJMeterSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartTestingJMeterSceneResponse
func (client *Client) StartTestingJMeterSceneWithContext(ctx context.Context, request *StartTestingJMeterSceneRequest, runtime *dara.RuntimeOptions) (_result *StartTestingJMeterSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartTestingJMeterScene"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartTestingJMeterSceneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops the scenario that is in debugging.
//
// @param request - StopDebugPtsSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDebugPtsSceneResponse
func (client *Client) StopDebugPtsSceneWithContext(ctx context.Context, request *StopDebugPtsSceneRequest, runtime *dara.RuntimeOptions) (_result *StopDebugPtsSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PlanId) {
		query["PlanId"] = request.PlanId
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopDebugPtsScene"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopDebugPtsSceneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops the debugging of a JMeter scenario.
//
// @param request - StopDebuggingJMeterSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDebuggingJMeterSceneResponse
func (client *Client) StopDebuggingJMeterSceneWithContext(ctx context.Context, request *StopDebuggingJMeterSceneRequest, runtime *dara.RuntimeOptions) (_result *StopDebuggingJMeterSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopDebuggingJMeterScene"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopDebuggingJMeterSceneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a scenario by using its ID.
//
// @param request - StopPtsSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopPtsSceneResponse
func (client *Client) StopPtsSceneWithContext(ctx context.Context, request *StopPtsSceneRequest, runtime *dara.RuntimeOptions) (_result *StopPtsSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopPtsScene"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopPtsSceneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops performance testing by using a JMeter scenario.
//
// @param request - StopTestingJMeterSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopTestingJMeterSceneResponse
func (client *Client) StopTestingJMeterSceneWithContext(ctx context.Context, request *StopTestingJMeterSceneRequest, runtime *dara.RuntimeOptions) (_result *StopTestingJMeterSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopTestingJMeterScene"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopTestingJMeterSceneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// null
//
// @param tmpReq - UpdatePtsSceneBaseLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePtsSceneBaseLineResponse
func (client *Client) UpdatePtsSceneBaseLineWithContext(ctx context.Context, tmpReq *UpdatePtsSceneBaseLineRequest, runtime *dara.RuntimeOptions) (_result *UpdatePtsSceneBaseLineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdatePtsSceneBaseLineShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApiBaselines) {
		request.ApiBaselinesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApiBaselines, dara.String("ApiBaselines"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SceneBaseline) {
		request.SceneBaselineShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SceneBaseline, dara.String("SceneBaseline"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiBaselinesShrink) {
		query["ApiBaselines"] = request.ApiBaselinesShrink
	}

	if !dara.IsNil(request.SceneBaselineShrink) {
		query["SceneBaseline"] = request.SceneBaselineShrink
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePtsSceneBaseLine"),
		Version:     dara.String("2020-10-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePtsSceneBaseLineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
