// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 应用/发布指定的推荐引擎配置
//
// @param request - ApplyEngineConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyEngineConfigResponse
func (client *Client) ApplyEngineConfigWithContext(ctx context.Context, EngineConfigId *string, request *ApplyEngineConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ApplyEngineConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyEngineConfig"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/engineconfigs/" + dara.PercentEncode(dara.StringValue(EngineConfigId)) + "/action/apply"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyEngineConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 特征一致性检查数据回流。
//
// @param request - BackflowFeatureConsistencyCheckJobDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BackflowFeatureConsistencyCheckJobDataResponse
func (client *Client) BackflowFeatureConsistencyCheckJobDataWithContext(ctx context.Context, request *BackflowFeatureConsistencyCheckJobDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BackflowFeatureConsistencyCheckJobDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FeatureConsistencyCheckJobConfigId) {
		body["FeatureConsistencyCheckJobConfigId"] = request.FeatureConsistencyCheckJobConfigId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ItemFeatures) {
		body["ItemFeatures"] = request.ItemFeatures
	}

	if !dara.IsNil(request.LogItemId) {
		body["LogItemId"] = request.LogItemId
	}

	if !dara.IsNil(request.LogRequestId) {
		body["LogRequestId"] = request.LogRequestId
	}

	if !dara.IsNil(request.LogRequestTime) {
		body["LogRequestTime"] = request.LogRequestTime
	}

	if !dara.IsNil(request.LogUserId) {
		body["LogUserId"] = request.LogUserId
	}

	if !dara.IsNil(request.SceneName) {
		body["SceneName"] = request.SceneName
	}

	if !dara.IsNil(request.Scores) {
		body["Scores"] = request.Scores
	}

	if !dara.IsNil(request.ServiceName) {
		body["ServiceName"] = request.ServiceName
	}

	if !dara.IsNil(request.UserFeatures) {
		body["UserFeatures"] = request.UserFeatures
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BackflowFeatureConsistencyCheckJobData"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/featureconsistencycheck/jobs/action/backflowdata"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BackflowFeatureConsistencyCheckJobDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检测实例下配置的资源的连接状态。
//
// @param request - CheckInstanceResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckInstanceResourcesResponse
func (client *Client) CheckInstanceResourcesWithContext(ctx context.Context, InstanceId *string, request *CheckInstanceResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CheckInstanceResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ResourceId) {
		body["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.Uri) {
		body["Uri"] = request.Uri
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckInstanceResources"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/action/checkresources"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckInstanceResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 校验流量调控任务中的表达式
//
// @param request - CheckTrafficControlTaskExpressionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckTrafficControlTaskExpressionResponse
func (client *Client) CheckTrafficControlTaskExpressionWithContext(ctx context.Context, request *CheckTrafficControlTaskExpressionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CheckTrafficControlTaskExpressionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Expression) {
		query["Expression"] = request.Expression
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.TableMetaId) {
		query["TableMetaId"] = request.TableMetaId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckTrafficControlTaskExpression"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltasks/action/checkexpression"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckTrafficControlTaskExpressionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 克隆指定的推荐引擎配置
//
// @param request - CloneEngineConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloneEngineConfigResponse
func (client *Client) CloneEngineConfigWithContext(ctx context.Context, EngineConfigId *string, request *CloneEngineConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CloneEngineConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigValue) {
		body["ConfigValue"] = request.ConfigValue
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Environment) {
		body["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloneEngineConfig"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/engineconfigs/" + dara.PercentEncode(dara.StringValue(EngineConfigId)) + "/action/clone"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CloneEngineConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 克隆实验。
//
// @param request - CloneExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloneExperimentResponse
func (client *Client) CloneExperimentWithContext(ctx context.Context, ExperimentId *string, request *CloneExperimentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CloneExperimentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloneExperiment"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments/" + dara.PercentEncode(dara.StringValue(ExperimentId)) + "/action/clone"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CloneExperimentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 克隆实验组，并克隆实验组下的所有实验至新的实验组中。
//
// @param request - CloneExperimentGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloneExperimentGroupResponse
func (client *Client) CloneExperimentGroupWithContext(ctx context.Context, ExperimentGroupId *string, request *CloneExperimentGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CloneExperimentGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Environment) {
		body["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LayerId) {
		body["LayerId"] = request.LayerId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloneExperimentGroup"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experimentgroups/" + dara.PercentEncode(dara.StringValue(ExperimentGroupId)) + "/action/clone"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CloneExperimentGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 克隆特征一致性检查配置。
//
// @param request - CloneFeatureConsistencyCheckJobConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloneFeatureConsistencyCheckJobConfigResponse
func (client *Client) CloneFeatureConsistencyCheckJobConfigWithContext(ctx context.Context, SourceFeatureConsistencyCheckJobConfigId *string, request *CloneFeatureConsistencyCheckJobConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CloneFeatureConsistencyCheckJobConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloneFeatureConsistencyCheckJobConfig"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/featureconsistencycheck/jobconfigs/" + dara.PercentEncode(dara.StringValue(SourceFeatureConsistencyCheckJobConfigId)) + "/action/clone"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CloneFeatureConsistencyCheckJobConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 克隆实验室。
//
// @param request - CloneLaboratoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloneLaboratoryResponse
func (client *Client) CloneLaboratoryWithContext(ctx context.Context, LaboratoryId *string, request *CloneLaboratoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CloneLaboratoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CloneExperimentGroup) {
		body["CloneExperimentGroup"] = request.CloneExperimentGroup
	}

	if !dara.IsNil(request.Environment) {
		body["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloneLaboratory"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/laboratories/" + dara.PercentEncode(dara.StringValue(LaboratoryId)) + "/action/clone"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CloneLaboratoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 克隆流量调控任务
//
// @param request - CloneTrafficControlTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloneTrafficControlTaskResponse
func (client *Client) CloneTrafficControlTaskWithContext(ctx context.Context, TrafficControlTaskId *string, request *CloneTrafficControlTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CloneTrafficControlTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloneTrafficControlTask"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltasks/" + dara.PercentEncode(dara.StringValue(TrafficControlTaskId)) + "/action/clone"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CloneTrafficControlTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 对比样本一致性任务
//
// @param request - CompareSampleConsistencyJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompareSampleConsistencyJobResponse
func (client *Client) CompareSampleConsistencyJobWithContext(ctx context.Context, SampleConsistencyJobId *string, request *CompareSampleConsistencyJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CompareSampleConsistencyJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CompareSampleConsistencyJob"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/sampleconsistencyjobs/" + dara.PercentEncode(dara.StringValue(SampleConsistencyJobId)) + "/action/compare"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CompareSampleConsistencyJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建AB test实验指标
//
// @param request - CreateABMetricRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateABMetricResponse
func (client *Client) CreateABMetricWithContext(ctx context.Context, request *CreateABMetricRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateABMetricResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Definition) {
		body["Definition"] = request.Definition
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LeftMetricId) {
		body["LeftMetricId"] = request.LeftMetricId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Operator) {
		body["Operator"] = request.Operator
	}

	if !dara.IsNil(request.Realtime) {
		body["Realtime"] = request.Realtime
	}

	if !dara.IsNil(request.ResultResourceId) {
		body["ResultResourceId"] = request.ResultResourceId
	}

	if !dara.IsNil(request.RightMetricId) {
		body["RightMetricId"] = request.RightMetricId
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.StatisticsCycle) {
		body["StatisticsCycle"] = request.StatisticsCycle
	}

	if !dara.IsNil(request.TableMetaId) {
		body["TableMetaId"] = request.TableMetaId
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateABMetric"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/abmetrics"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateABMetricResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建指标组
//
// @param request - CreateABMetricGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateABMetricGroupResponse
func (client *Client) CreateABMetricGroupWithContext(ctx context.Context, request *CreateABMetricGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateABMetricGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ABMetricIds) {
		body["ABMetricIds"] = request.ABMetricIds
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Realtime) {
		body["Realtime"] = request.Realtime
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateABMetricGroup"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/abmetricgroups"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateABMetricGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建AB指标的计算任务。
//
// @param request - CreateCalculationJobsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCalculationJobsResponse
func (client *Client) CreateCalculationJobsWithContext(ctx context.Context, request *CreateCalculationJobsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateCalculationJobsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ABMetricIds) {
		body["ABMetricIds"] = request.ABMetricIds
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCalculationJobs"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/batch/calculationjobs/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCalculationJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建人群。
//
// @param request - CreateCrowdRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCrowdResponse
func (client *Client) CreateCrowdWithContext(ctx context.Context, request *CreateCrowdRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateCrowdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Label) {
		body["Label"] = request.Label
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Source) {
		body["Source"] = request.Source
	}

	if !dara.IsNil(request.Users) {
		body["Users"] = request.Users
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCrowd"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/crowds"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCrowdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建引擎配置
//
// @param request - CreateEngineConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEngineConfigResponse
func (client *Client) CreateEngineConfigWithContext(ctx context.Context, request *CreateEngineConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateEngineConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigValue) {
		body["ConfigValue"] = request.ConfigValue
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Environment) {
		body["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEngineConfig"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/engineconfigs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEngineConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建实验。
//
// @param request - CreateExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExperimentResponse
func (client *Client) CreateExperimentWithContext(ctx context.Context, request *CreateExperimentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateExperimentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["Config"] = request.Config
	}

	if !dara.IsNil(request.DebugCrowdId) {
		body["DebugCrowdId"] = request.DebugCrowdId
	}

	if !dara.IsNil(request.DebugUsers) {
		body["DebugUsers"] = request.DebugUsers
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ExperimentGroupId) {
		body["ExperimentGroupId"] = request.ExperimentGroupId
	}

	if !dara.IsNil(request.FlowPercent) {
		body["FlowPercent"] = request.FlowPercent
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateExperiment"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateExperimentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建实验组。
//
// @param request - CreateExperimentGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExperimentGroupResponse
func (client *Client) CreateExperimentGroupWithContext(ctx context.Context, request *CreateExperimentGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateExperimentGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["Config"] = request.Config
	}

	if !dara.IsNil(request.CrowdId) {
		body["CrowdId"] = request.CrowdId
	}

	if !dara.IsNil(request.CrowdTargetType) {
		body["CrowdTargetType"] = request.CrowdTargetType
	}

	if !dara.IsNil(request.DebugCrowdId) {
		body["DebugCrowdId"] = request.DebugCrowdId
	}

	if !dara.IsNil(request.DebugUsers) {
		body["DebugUsers"] = request.DebugUsers
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DistributionTimeDuration) {
		body["DistributionTimeDuration"] = request.DistributionTimeDuration
	}

	if !dara.IsNil(request.DistributionType) {
		body["DistributionType"] = request.DistributionType
	}

	if !dara.IsNil(request.Filter) {
		body["Filter"] = request.Filter
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LayerId) {
		body["LayerId"] = request.LayerId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NeedAA) {
		body["NeedAA"] = request.NeedAA
	}

	if !dara.IsNil(request.RandomFlow) {
		body["RandomFlow"] = request.RandomFlow
	}

	if !dara.IsNil(request.ReservedBuckets) {
		body["ReservedBuckets"] = request.ReservedBuckets
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateExperimentGroup"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experimentgroups"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateExperimentGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建特征一致性检查任务。
//
// @param request - CreateFeatureConsistencyCheckJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFeatureConsistencyCheckJobResponse
func (client *Client) CreateFeatureConsistencyCheckJobWithContext(ctx context.Context, request *CreateFeatureConsistencyCheckJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateFeatureConsistencyCheckJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Environment) {
		body["Environment"] = request.Environment
	}

	if !dara.IsNil(request.FeatureConsistencyCheckJobConfigId) {
		body["FeatureConsistencyCheckJobConfigId"] = request.FeatureConsistencyCheckJobConfigId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SamplingDuration) {
		body["SamplingDuration"] = request.SamplingDuration
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFeatureConsistencyCheckJob"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/featureconsistencycheck/jobs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFeatureConsistencyCheckJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建特征一致性检查配置。
//
// @param request - CreateFeatureConsistencyCheckJobConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFeatureConsistencyCheckJobConfigResponse
func (client *Client) CreateFeatureConsistencyCheckJobConfigWithContext(ctx context.Context, request *CreateFeatureConsistencyCheckJobConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateFeatureConsistencyCheckJobConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CompareFeature) {
		body["CompareFeature"] = request.CompareFeature
	}

	if !dara.IsNil(request.DatasetId) {
		body["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.DatasetMountPath) {
		body["DatasetMountPath"] = request.DatasetMountPath
	}

	if !dara.IsNil(request.DatasetName) {
		body["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.DatasetType) {
		body["DatasetType"] = request.DatasetType
	}

	if !dara.IsNil(request.DatasetUri) {
		body["DatasetUri"] = request.DatasetUri
	}

	if !dara.IsNil(request.DefaultRoute) {
		body["DefaultRoute"] = request.DefaultRoute
	}

	if !dara.IsNil(request.EasServiceName) {
		body["EasServiceName"] = request.EasServiceName
	}

	if !dara.IsNil(request.EasyRecPackagePath) {
		body["EasyRecPackagePath"] = request.EasyRecPackagePath
	}

	if !dara.IsNil(request.EasyRecVersion) {
		body["EasyRecVersion"] = request.EasyRecVersion
	}

	if !dara.IsNil(request.FeatureDisplayExclude) {
		body["FeatureDisplayExclude"] = request.FeatureDisplayExclude
	}

	if !dara.IsNil(request.FeatureLandingResourceId) {
		body["FeatureLandingResourceId"] = request.FeatureLandingResourceId
	}

	if !dara.IsNil(request.FeaturePriority) {
		body["FeaturePriority"] = request.FeaturePriority
	}

	if !dara.IsNil(request.FeatureStoreItemId) {
		body["FeatureStoreItemId"] = request.FeatureStoreItemId
	}

	if !dara.IsNil(request.FeatureStoreModelId) {
		body["FeatureStoreModelId"] = request.FeatureStoreModelId
	}

	if !dara.IsNil(request.FeatureStoreProjectId) {
		body["FeatureStoreProjectId"] = request.FeatureStoreProjectId
	}

	if !dara.IsNil(request.FeatureStoreProjectName) {
		body["FeatureStoreProjectName"] = request.FeatureStoreProjectName
	}

	if !dara.IsNil(request.FeatureStoreSeqFeatureView) {
		body["FeatureStoreSeqFeatureView"] = request.FeatureStoreSeqFeatureView
	}

	if !dara.IsNil(request.FeatureStoreUserId) {
		body["FeatureStoreUserId"] = request.FeatureStoreUserId
	}

	if !dara.IsNil(request.FgJarVersion) {
		body["FgJarVersion"] = request.FgJarVersion
	}

	if !dara.IsNil(request.FgJsonFileName) {
		body["FgJsonFileName"] = request.FgJsonFileName
	}

	if !dara.IsNil(request.GenerateZip) {
		body["GenerateZip"] = request.GenerateZip
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ItemIdField) {
		body["ItemIdField"] = request.ItemIdField
	}

	if !dara.IsNil(request.ItemTable) {
		body["ItemTable"] = request.ItemTable
	}

	if !dara.IsNil(request.ItemTablePartitionField) {
		body["ItemTablePartitionField"] = request.ItemTablePartitionField
	}

	if !dara.IsNil(request.ItemTablePartitionFieldFormat) {
		body["ItemTablePartitionFieldFormat"] = request.ItemTablePartitionFieldFormat
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.OssResourceId) {
		body["OssResourceId"] = request.OssResourceId
	}

	if !dara.IsNil(request.PredictWorkerCount) {
		body["PredictWorkerCount"] = request.PredictWorkerCount
	}

	if !dara.IsNil(request.PredictWorkerCpu) {
		body["PredictWorkerCpu"] = request.PredictWorkerCpu
	}

	if !dara.IsNil(request.PredictWorkerMemory) {
		body["PredictWorkerMemory"] = request.PredictWorkerMemory
	}

	if !dara.IsNil(request.ResourceConfig) {
		body["ResourceConfig"] = request.ResourceConfig
	}

	if !dara.IsNil(request.SampleRate) {
		body["SampleRate"] = request.SampleRate
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.SecurityGroupId) {
		body["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.ServiceId) {
		body["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.SwitchId) {
		body["SwitchId"] = request.SwitchId
	}

	if !dara.IsNil(request.UseFeatureStore) {
		body["UseFeatureStore"] = request.UseFeatureStore
	}

	if !dara.IsNil(request.UserIdField) {
		body["UserIdField"] = request.UserIdField
	}

	if !dara.IsNil(request.UserTable) {
		body["UserTable"] = request.UserTable
	}

	if !dara.IsNil(request.UserTablePartitionField) {
		body["UserTablePartitionField"] = request.UserTablePartitionField
	}

	if !dara.IsNil(request.UserTablePartitionFieldFormat) {
		body["UserTablePartitionFieldFormat"] = request.UserTablePartitionFieldFormat
	}

	if !dara.IsNil(request.VpcId) {
		body["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.WorkflowName) {
		body["WorkflowName"] = request.WorkflowName
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFeatureConsistencyCheckJobConfig"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/featureconsistencycheck/jobconfigs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFeatureConsistencyCheckJobConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 为指定实例配置创建新的配置资源
//
// @param request - CreateInstanceResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceResourceResponse
func (client *Client) CreateInstanceResourceWithContext(ctx context.Context, InstanceId *string, request *CreateInstanceResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateInstanceResourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		body["Category"] = request.Category
	}

	if !dara.IsNil(request.Group) {
		body["Group"] = request.Group
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.Uri) {
		body["Uri"] = request.Uri
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstanceResource"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/resources"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInstanceResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建实验室
//
// @param request - CreateLaboratoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLaboratoryResponse
func (client *Client) CreateLaboratoryWithContext(ctx context.Context, request *CreateLaboratoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateLaboratoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BucketCount) {
		body["BucketCount"] = request.BucketCount
	}

	if !dara.IsNil(request.BucketType) {
		body["BucketType"] = request.BucketType
	}

	if !dara.IsNil(request.Buckets) {
		body["Buckets"] = request.Buckets
	}

	if !dara.IsNil(request.DebugCrowdId) {
		body["DebugCrowdId"] = request.DebugCrowdId
	}

	if !dara.IsNil(request.DebugUsers) {
		body["DebugUsers"] = request.DebugUsers
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Environment) {
		body["Environment"] = request.Environment
	}

	if !dara.IsNil(request.Filter) {
		body["Filter"] = request.Filter
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLaboratory"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/laboratories"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLaboratoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建层。
//
// @param request - CreateLayerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLayerResponse
func (client *Client) CreateLayerWithContext(ctx context.Context, request *CreateLayerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateLayerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LaboratoryId) {
		body["LaboratoryId"] = request.LaboratoryId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLayer"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/layers"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLayerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建参数。
//
// @param request - CreateParamRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateParamResponse
func (client *Client) CreateParamWithContext(ctx context.Context, request *CreateParamRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateParamResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Environment) {
		body["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.Value) {
		body["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateParam"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/params"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateParamResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建资源规则
//
// @param request - CreateResourceRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceRuleResponse
func (client *Client) CreateResourceRuleWithContext(ctx context.Context, request *CreateResourceRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateResourceRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MetricOperationType) {
		body["MetricOperationType"] = request.MetricOperationType
	}

	if !dara.IsNil(request.MetricPullInfo) {
		body["MetricPullInfo"] = request.MetricPullInfo
	}

	if !dara.IsNil(request.MetricPullPeriod) {
		body["MetricPullPeriod"] = request.MetricPullPeriod
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.RuleComputingDefinition) {
		body["RuleComputingDefinition"] = request.RuleComputingDefinition
	}

	if !dara.IsNil(request.RuleItems) {
		body["RuleItems"] = request.RuleItems
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateResourceRule"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resourcerules"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateResourceRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建资源规则条目
//
// @param request - CreateResourceRuleItemRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceRuleItemResponse
func (client *Client) CreateResourceRuleItemWithContext(ctx context.Context, ResourceRuleId *string, request *CreateResourceRuleItemRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateResourceRuleItemResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxValue) {
		body["MaxValue"] = request.MaxValue
	}

	if !dara.IsNil(request.MinValue) {
		body["MinValue"] = request.MinValue
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Value) {
		body["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateResourceRuleItem"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resourcerules/" + dara.PercentEncode(dara.StringValue(ResourceRuleId)) + "/items"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateResourceRuleItemResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建样本一致性任务
//
// @param request - CreateSampleConsistencyJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSampleConsistencyJobResponse
func (client *Client) CreateSampleConsistencyJobWithContext(ctx context.Context, request *CreateSampleConsistencyJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateSampleConsistencyJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Duration) {
		body["Duration"] = request.Duration
	}

	if !dara.IsNil(request.EasModelServiceName) {
		body["EasModelServiceName"] = request.EasModelServiceName
	}

	if !dara.IsNil(request.FeatureSaveResourceId) {
		body["FeatureSaveResourceId"] = request.FeatureSaveResourceId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ItemIdField) {
		body["ItemIdField"] = request.ItemIdField
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PartitionField) {
		body["PartitionField"] = request.PartitionField
	}

	if !dara.IsNil(request.PartitionFieldFormat) {
		body["PartitionFieldFormat"] = request.PartitionFieldFormat
	}

	if !dara.IsNil(request.RequestIdField) {
		body["RequestIdField"] = request.RequestIdField
	}

	if !dara.IsNil(request.SampleRate) {
		body["SampleRate"] = request.SampleRate
	}

	if !dara.IsNil(request.SampleTableName) {
		body["SampleTableName"] = request.SampleTableName
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.UserIdField) {
		body["UserIdField"] = request.UserIdField
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSampleConsistencyJob"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/sampleconsistencyjobs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSampleConsistencyJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建场景
//
// @param request - CreateSceneRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSceneResponse
func (client *Client) CreateSceneWithContext(ctx context.Context, request *CreateSceneRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateSceneResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Flows) {
		body["Flows"] = request.Flows
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateScene"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/scenes"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSceneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 在指定人群下创建子人群。
//
// @param request - CreateSubCrowdRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSubCrowdResponse
func (client *Client) CreateSubCrowdWithContext(ctx context.Context, CrowdId *string, request *CreateSubCrowdRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateSubCrowdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Source) {
		body["Source"] = request.Source
	}

	if !dara.IsNil(request.Users) {
		body["Users"] = request.Users
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSubCrowd"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/crowds/" + dara.PercentEncode(dara.StringValue(CrowdId)) + "/subcrowds"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSubCrowdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建数据表。
//
// @param request - CreateTableMetaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTableMetaResponse
func (client *Client) CreateTableMetaWithContext(ctx context.Context, request *CreateTableMetaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTableMetaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Fields) {
		body["Fields"] = request.Fields
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Module) {
		body["Module"] = request.Module
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ResourceId) {
		body["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.TableName) {
		body["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTableMeta"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/tablemetas"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTableMetaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建流量调控目标
//
// @param request - CreateTrafficControlTargetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTrafficControlTargetResponse
func (client *Client) CreateTrafficControlTargetWithContext(ctx context.Context, request *CreateTrafficControlTargetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTrafficControlTargetResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Event) {
		body["Event"] = request.Event
	}

	if !dara.IsNil(request.ItemConditionArray) {
		body["ItemConditionArray"] = request.ItemConditionArray
	}

	if !dara.IsNil(request.ItemConditionExpress) {
		body["ItemConditionExpress"] = request.ItemConditionExpress
	}

	if !dara.IsNil(request.ItemConditionType) {
		body["ItemConditionType"] = request.ItemConditionType
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NewProductRegulation) {
		body["NewProductRegulation"] = request.NewProductRegulation
	}

	if !dara.IsNil(request.RecallName) {
		body["RecallName"] = request.RecallName
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StatisPeriod) {
		body["StatisPeriod"] = request.StatisPeriod
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.ToleranceValue) {
		body["ToleranceValue"] = request.ToleranceValue
	}

	if !dara.IsNil(request.TrafficControlTaskId) {
		body["TrafficControlTaskId"] = request.TrafficControlTaskId
	}

	if !dara.IsNil(request.Value) {
		body["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTrafficControlTarget"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltargets"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTrafficControlTargetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建流量调控任务
//
// @param request - CreateTrafficControlTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTrafficControlTaskResponse
func (client *Client) CreateTrafficControlTaskWithContext(ctx context.Context, request *CreateTrafficControlTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTrafficControlTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BehaviorTableMetaId) {
		body["BehaviorTableMetaId"] = request.BehaviorTableMetaId
	}

	if !dara.IsNil(request.ControlGranularity) {
		body["ControlGranularity"] = request.ControlGranularity
	}

	if !dara.IsNil(request.ControlLogic) {
		body["ControlLogic"] = request.ControlLogic
	}

	if !dara.IsNil(request.ControlType) {
		body["ControlType"] = request.ControlType
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EffectiveSceneIds) {
		body["EffectiveSceneIds"] = request.EffectiveSceneIds
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ExecutionTime) {
		body["ExecutionTime"] = request.ExecutionTime
	}

	if !dara.IsNil(request.FlinkResourceId) {
		body["FlinkResourceId"] = request.FlinkResourceId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ItemConditionArray) {
		body["ItemConditionArray"] = request.ItemConditionArray
	}

	if !dara.IsNil(request.ItemConditionExpress) {
		body["ItemConditionExpress"] = request.ItemConditionExpress
	}

	if !dara.IsNil(request.ItemConditionType) {
		body["ItemConditionType"] = request.ItemConditionType
	}

	if !dara.IsNil(request.ItemTableMetaId) {
		body["ItemTableMetaId"] = request.ItemTableMetaId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PreExperimentIds) {
		body["PreExperimentIds"] = request.PreExperimentIds
	}

	if !dara.IsNil(request.ProdExperimentIds) {
		body["ProdExperimentIds"] = request.ProdExperimentIds
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.ServiceId) {
		body["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceIds) {
		body["ServiceIds"] = request.ServiceIds
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StatisBehaviorConditionArray) {
		body["StatisBehaviorConditionArray"] = request.StatisBehaviorConditionArray
	}

	if !dara.IsNil(request.StatisBehaviorConditionExpress) {
		body["StatisBehaviorConditionExpress"] = request.StatisBehaviorConditionExpress
	}

	if !dara.IsNil(request.StatisBehaviorConditionType) {
		body["StatisBehaviorConditionType"] = request.StatisBehaviorConditionType
	}

	if !dara.IsNil(request.TrafficControlTargets) {
		body["TrafficControlTargets"] = request.TrafficControlTargets
	}

	if !dara.IsNil(request.UserConditionArray) {
		body["UserConditionArray"] = request.UserConditionArray
	}

	if !dara.IsNil(request.UserConditionExpress) {
		body["UserConditionExpress"] = request.UserConditionExpress
	}

	if !dara.IsNil(request.UserConditionType) {
		body["UserConditionType"] = request.UserConditionType
	}

	if !dara.IsNil(request.UserTableMetaId) {
		body["UserTableMetaId"] = request.UserTableMetaId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTrafficControlTask"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltasks"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTrafficControlTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 对指定资源规则中的计算公式进行调试
//
// @param tmpReq - DebugResourceRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DebugResourceRuleResponse
func (client *Client) DebugResourceRuleWithContext(ctx context.Context, ResourceRuleId *string, tmpReq *DebugResourceRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DebugResourceRuleResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DebugResourceRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MetricInfo) {
		request.MetricInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MetricInfo, dara.String("MetricInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MetricInfoShrink) {
		query["MetricInfo"] = request.MetricInfoShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DebugResourceRule"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resourcerules/" + dara.PercentEncode(dara.StringValue(ResourceRuleId)) + "/action/debug"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DebugResourceRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定AB实验指标。
//
// @param request - DeleteABMetricRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteABMetricResponse
func (client *Client) DeleteABMetricWithContext(ctx context.Context, ABMetricId *string, request *DeleteABMetricRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteABMetricResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteABMetric"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/abmetrics/" + dara.PercentEncode(dara.StringValue(ABMetricId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteABMetricResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除AB实验指标组。
//
// @param request - DeleteABMetricGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteABMetricGroupResponse
func (client *Client) DeleteABMetricGroupWithContext(ctx context.Context, ABMetricGroupId *string, request *DeleteABMetricGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteABMetricGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteABMetricGroup"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/abmetricgroups/" + dara.PercentEncode(dara.StringValue(ABMetricGroupId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteABMetricGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定人群。
//
// @param request - DeleteCrowdRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCrowdResponse
func (client *Client) DeleteCrowdWithContext(ctx context.Context, CrowdId *string, request *DeleteCrowdRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteCrowdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCrowd"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/crowds/" + dara.PercentEncode(dara.StringValue(CrowdId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCrowdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定推荐引擎配置。
//
// @param request - DeleteEngineConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEngineConfigResponse
func (client *Client) DeleteEngineConfigWithContext(ctx context.Context, EngineConfigId *string, request *DeleteEngineConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteEngineConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEngineConfig"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/engineconfigs/" + dara.PercentEncode(dara.StringValue(EngineConfigId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEngineConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除实验。
//
// @param request - DeleteExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExperimentResponse
func (client *Client) DeleteExperimentWithContext(ctx context.Context, ExperimentId *string, request *DeleteExperimentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteExperimentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteExperiment"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments/" + dara.PercentEncode(dara.StringValue(ExperimentId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteExperimentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定实验组。
//
// @param request - DeleteExperimentGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExperimentGroupResponse
func (client *Client) DeleteExperimentGroupWithContext(ctx context.Context, ExperimentGroupId *string, request *DeleteExperimentGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteExperimentGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteExperimentGroup"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experimentgroups/" + dara.PercentEncode(dara.StringValue(ExperimentGroupId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteExperimentGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定实例下的指定配置资源。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResourceResponse
func (client *Client) DeleteInstanceResourceWithContext(ctx context.Context, InstanceId *string, ResourceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteInstanceResourceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteInstanceResource"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/resources/" + dara.PercentEncode(dara.StringValue(ResourceId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteInstanceResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除实验室。
//
// @param request - DeleteLaboratoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLaboratoryResponse
func (client *Client) DeleteLaboratoryWithContext(ctx context.Context, LaboratoryId *string, request *DeleteLaboratoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteLaboratoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLaboratory"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/laboratories/" + dara.PercentEncode(dara.StringValue(LaboratoryId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLaboratoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除层。
//
// @param request - DeleteLayerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLayerResponse
func (client *Client) DeleteLayerWithContext(ctx context.Context, LayerId *string, request *DeleteLayerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteLayerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLayer"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/layers/" + dara.PercentEncode(dara.StringValue(LayerId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLayerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定参数。
//
// @param request - DeleteParamRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteParamResponse
func (client *Client) DeleteParamWithContext(ctx context.Context, ParamId *string, request *DeleteParamRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteParamResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteParam"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/params/" + dara.PercentEncode(dara.StringValue(ParamId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteParamResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除资源规则
//
// @param request - DeleteResourceRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceRuleResponse
func (client *Client) DeleteResourceRuleWithContext(ctx context.Context, ResourceRuleId *string, request *DeleteResourceRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteResourceRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteResourceRule"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resourcerules/" + dara.PercentEncode(dara.StringValue(ResourceRuleId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteResourceRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除资源规则条目
//
// @param request - DeleteResourceRuleItemRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceRuleItemResponse
func (client *Client) DeleteResourceRuleItemWithContext(ctx context.Context, ResourceRuleId *string, ResourceRuleItemId *string, request *DeleteResourceRuleItemRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteResourceRuleItemResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteResourceRuleItem"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resourcerules/" + dara.PercentEncode(dara.StringValue(ResourceRuleId)) + "/items/" + dara.PercentEncode(dara.StringValue(ResourceRuleItemId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteResourceRuleItemResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定的样本一致性任务
//
// @param request - DeleteSampleConsistencyJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSampleConsistencyJobResponse
func (client *Client) DeleteSampleConsistencyJobWithContext(ctx context.Context, SampleConsistencyJobId *string, request *DeleteSampleConsistencyJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteSampleConsistencyJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSampleConsistencyJob"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/sampleconsistencyjobs/" + dara.PercentEncode(dara.StringValue(SampleConsistencyJobId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSampleConsistencyJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除场景
//
// @param request - DeleteSceneRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSceneResponse
func (client *Client) DeleteSceneWithContext(ctx context.Context, SceneId *string, request *DeleteSceneRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteSceneResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteScene"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/scenes/" + dara.PercentEncode(dara.StringValue(SceneId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSceneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定人群下的指定子人群。
//
// @param request - DeleteSubCrowdRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSubCrowdResponse
func (client *Client) DeleteSubCrowdWithContext(ctx context.Context, CrowdId *string, SubCrowdId *string, request *DeleteSubCrowdRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteSubCrowdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSubCrowd"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/crowds/" + dara.PercentEncode(dara.StringValue(CrowdId)) + "/subcrowds/" + dara.PercentEncode(dara.StringValue(SubCrowdId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSubCrowdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除数据表。
//
// @param request - DeleteTableMetaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTableMetaResponse
func (client *Client) DeleteTableMetaWithContext(ctx context.Context, TableMetaId *string, request *DeleteTableMetaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteTableMetaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTableMeta"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/tablemetas/" + dara.PercentEncode(dara.StringValue(TableMetaId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTableMetaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新流量调控目标
//
// @param request - DeleteTrafficControlTargetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTrafficControlTargetResponse
func (client *Client) DeleteTrafficControlTargetWithContext(ctx context.Context, TrafficControlTargetId *string, request *DeleteTrafficControlTargetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteTrafficControlTargetResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTrafficControlTarget"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltargets/" + dara.PercentEncode(dara.StringValue(TrafficControlTargetId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTrafficControlTargetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定的流量调控任务
//
// @param request - DeleteTrafficControlTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTrafficControlTaskResponse
func (client *Client) DeleteTrafficControlTaskWithContext(ctx context.Context, TrafficControlTaskId *string, request *DeleteTrafficControlTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteTrafficControlTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTrafficControlTask"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltasks/" + dara.PercentEncode(dara.StringValue(TrafficControlTaskId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTrafficControlTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成算法定制脚本
//
// @param request - GenerateAlgorithmCustomizationScriptRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateAlgorithmCustomizationScriptResponse
func (client *Client) GenerateAlgorithmCustomizationScriptWithContext(ctx context.Context, AlgorithmCustomizationId *string, request *GenerateAlgorithmCustomizationScriptRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GenerateAlgorithmCustomizationScriptResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DeployMode) {
		body["DeployMode"] = request.DeployMode
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ModuleFieldTypes) {
		body["ModuleFieldTypes"] = request.ModuleFieldTypes
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateAlgorithmCustomizationScript"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/algorithmcustomizations/" + dara.PercentEncode(dara.StringValue(AlgorithmCustomizationId)) + "/action/generatescript"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateAlgorithmCustomizationScriptResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 产生流量调控的相关代码
//
// @param request - GenerateTrafficControlTaskCodeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateTrafficControlTaskCodeResponse
func (client *Client) GenerateTrafficControlTaskCodeWithContext(ctx context.Context, TrafficControlTaskId *string, request *GenerateTrafficControlTaskCodeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GenerateTrafficControlTaskCodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Environment) {
		body["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateTrafficControlTaskCode"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltasks/" + dara.PercentEncode(dara.StringValue(TrafficControlTaskId)) + "/action/generatecode"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateTrafficControlTaskCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 产生流量调控的相关召回配置
//
// @param request - GenerateTrafficControlTaskConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateTrafficControlTaskConfigResponse
func (client *Client) GenerateTrafficControlTaskConfigWithContext(ctx context.Context, TrafficControlTaskId *string, request *GenerateTrafficControlTaskConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GenerateTrafficControlTaskConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateTrafficControlTaskConfig"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltasks/" + dara.PercentEncode(dara.StringValue(TrafficControlTaskId)) + "/action/generateconfig"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateTrafficControlTaskConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取AB Test实验指标详细信息。
//
// @param request - GetABMetricRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetABMetricResponse
func (client *Client) GetABMetricWithContext(ctx context.Context, ABMetricId *string, request *GetABMetricRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetABMetricResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetABMetric"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/abmetrics/" + dara.PercentEncode(dara.StringValue(ABMetricId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetABMetricResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取AB实验指标组详细信息。
//
// @param request - GetABMetricGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetABMetricGroupResponse
func (client *Client) GetABMetricGroupWithContext(ctx context.Context, ABMetricGroupId *string, request *GetABMetricGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetABMetricGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetABMetricGroup"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/abmetricgroups/" + dara.PercentEncode(dara.StringValue(ABMetricGroupId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetABMetricGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指定计算任务详细信息。
//
// @param request - GetCalculationJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCalculationJobResponse
func (client *Client) GetCalculationJobWithContext(ctx context.Context, CalculationJobId *string, request *GetCalculationJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCalculationJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCalculationJob"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/calculationjobs/" + dara.PercentEncode(dara.StringValue(CalculationJobId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCalculationJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取引擎配置详细信息。
//
// @param request - GetEngineConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEngineConfigResponse
func (client *Client) GetEngineConfigWithContext(ctx context.Context, EngineConfigId *string, request *GetEngineConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEngineConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEngineConfig"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/engineconfigs/" + dara.PercentEncode(dara.StringValue(EngineConfigId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEngineConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实验详细信息。
//
// @param request - GetExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExperimentResponse
func (client *Client) GetExperimentWithContext(ctx context.Context, ExperimentId *string, request *GetExperimentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetExperimentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetExperiment"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments/" + dara.PercentEncode(dara.StringValue(ExperimentId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetExperimentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指定实验组详细信息。
//
// @param request - GetExperimentGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExperimentGroupResponse
func (client *Client) GetExperimentGroupWithContext(ctx context.Context, ExperimentGroupId *string, request *GetExperimentGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetExperimentGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetExperimentGroup"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experimentgroups/" + dara.PercentEncode(dara.StringValue(ExperimentGroupId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetExperimentGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取特征一致性检查任务详细信息。
//
// @param request - GetFeatureConsistencyCheckJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFeatureConsistencyCheckJobResponse
func (client *Client) GetFeatureConsistencyCheckJobWithContext(ctx context.Context, FeatureConsistencyCheckJobId *string, request *GetFeatureConsistencyCheckJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetFeatureConsistencyCheckJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFeatureConsistencyCheckJob"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/featureconsistencycheck/jobs/" + dara.PercentEncode(dara.StringValue(FeatureConsistencyCheckJobId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFeatureConsistencyCheckJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取特征一致性检测配置详情。
//
// @param request - GetFeatureConsistencyCheckJobConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFeatureConsistencyCheckJobConfigResponse
func (client *Client) GetFeatureConsistencyCheckJobConfigWithContext(ctx context.Context, FeatureConsistencyCheckJobConfigId *string, request *GetFeatureConsistencyCheckJobConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetFeatureConsistencyCheckJobConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFeatureConsistencyCheckJobConfig"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/featureconsistencycheck/jobconfigs/" + dara.PercentEncode(dara.StringValue(FeatureConsistencyCheckJobConfigId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFeatureConsistencyCheckJobConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指定推荐全链路深度定制开发平台实例信息。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceResponse
func (client *Client) GetInstanceWithContext(ctx context.Context, InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstance"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指定实例下指定资源的详细信息。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceResourceResponse
func (client *Client) GetInstanceResourceWithContext(ctx context.Context, InstanceId *string, ResourceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetInstanceResourceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceResource"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/resources/" + dara.PercentEncode(dara.StringValue(ResourceId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据源下指定表的详细信息。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceResourceTableResponse
func (client *Client) GetInstanceResourceTableWithContext(ctx context.Context, InstanceId *string, ResourceId *string, TableName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetInstanceResourceTableResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceResourceTable"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/resources/" + dara.PercentEncode(dara.StringValue(ResourceId)) + "/tables/" + dara.PercentEncode(dara.StringValue(TableName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceResourceTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实验室详细信息。
//
// @param request - GetLaboratoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLaboratoryResponse
func (client *Client) GetLaboratoryWithContext(ctx context.Context, LaboratoryId *string, request *GetLaboratoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetLaboratoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLaboratory"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/laboratories/" + dara.PercentEncode(dara.StringValue(LaboratoryId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLaboratoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取层详细信息。
//
// @param request - GetLayerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLayerResponse
func (client *Client) GetLayerWithContext(ctx context.Context, LayerId *string, request *GetLayerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetLayerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLayer"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/layers/" + dara.PercentEncode(dara.StringValue(LayerId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLayerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取资源规则详细信息
//
// @param request - GetResourceRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceRuleResponse
func (client *Client) GetResourceRuleWithContext(ctx context.Context, ResourceRuleId *string, request *GetResourceRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetResourceRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceRule"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resourcerules/" + dara.PercentEncode(dara.StringValue(ResourceRuleId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取样本一致性任务详细信息
//
// @param request - GetSampleConsistencyJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSampleConsistencyJobResponse
func (client *Client) GetSampleConsistencyJobWithContext(ctx context.Context, SampleConsistencyJobId *string, request *GetSampleConsistencyJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSampleConsistencyJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSampleConsistencyJob"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/sampleconsistencyjobs/" + dara.PercentEncode(dara.StringValue(SampleConsistencyJobId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSampleConsistencyJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取场景详细信息
//
// @param request - GetSceneRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSceneResponse
func (client *Client) GetSceneWithContext(ctx context.Context, SceneId *string, request *GetSceneRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSceneResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetScene"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/scenes/" + dara.PercentEncode(dara.StringValue(SceneId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSceneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取服务详细信息。
//
// @param request - GetServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceResponse
func (client *Client) GetServiceWithContext(ctx context.Context, ServiceId *string, request *GetServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetService"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/services/" + dara.PercentEncode(dara.StringValue(ServiceId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指定人群下的指定子人群的详细信息。
//
// @param request - GetSubCrowdRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSubCrowdResponse
func (client *Client) GetSubCrowdWithContext(ctx context.Context, CrowdId *string, SubCrowdId *string, request *GetSubCrowdRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSubCrowdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSubCrowd"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/crowds/" + dara.PercentEncode(dara.StringValue(CrowdId)) + "/subcrowds/" + dara.PercentEncode(dara.StringValue(SubCrowdId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSubCrowdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据表详细信息。
//
// @param request - GetTableMetaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableMetaResponse
func (client *Client) GetTableMetaWithContext(ctx context.Context, TableMetaId *string, request *GetTableMetaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTableMetaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTableMeta"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/tablemetas/" + dara.PercentEncode(dara.StringValue(TableMetaId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTableMetaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取流量调控目标详情
//
// @param request - GetTrafficControlTargetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTrafficControlTargetResponse
func (client *Client) GetTrafficControlTargetWithContext(ctx context.Context, TrafficControlTargetId *string, request *GetTrafficControlTargetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTrafficControlTargetResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTrafficControlTarget"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltargets/" + dara.PercentEncode(dara.StringValue(TrafficControlTargetId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTrafficControlTargetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取流量调控任务详情
//
// @param request - GetTrafficControlTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTrafficControlTaskResponse
func (client *Client) GetTrafficControlTaskWithContext(ctx context.Context, TrafficControlTaskId *string, request *GetTrafficControlTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTrafficControlTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ControlTargetFilter) {
		query["ControlTargetFilter"] = request.ControlTargetFilter
	}

	if !dara.IsNil(request.Environment) {
		query["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTrafficControlTask"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltasks/" + dara.PercentEncode(dara.StringValue(TrafficControlTaskId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTrafficControlTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取流量调控任务的流量详情
//
// @param request - GetTrafficControlTaskTrafficRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTrafficControlTaskTrafficResponse
func (client *Client) GetTrafficControlTaskTrafficWithContext(ctx context.Context, TrafficControlTaskId *string, request *GetTrafficControlTaskTrafficRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTrafficControlTaskTrafficResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Environment) {
		query["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTrafficControlTaskTraffic"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltasks/" + dara.PercentEncode(dara.StringValue(TrafficControlTaskId)) + "/trafficinfo"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTrafficControlTaskTrafficResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取AB Test实验指标组列表。
//
// @param request - ListABMetricGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListABMetricGroupsResponse
func (client *Client) ListABMetricGroupsWithContext(ctx context.Context, request *ListABMetricGroupsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListABMetricGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.Realtime) {
		query["Realtime"] = request.Realtime
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListABMetricGroups"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/abmetricgroups"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListABMetricGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取AB Test实验指标列表。
//
// @param request - ListABMetricsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListABMetricsResponse
func (client *Client) ListABMetricsWithContext(ctx context.Context, request *ListABMetricsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListABMetricsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.Realtime) {
		query["Realtime"] = request.Realtime
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.TableMetaId) {
		query["TableMetaId"] = request.TableMetaId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListABMetrics"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/abmetrics"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListABMetricsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取计算任务列表。
//
// @param request - ListCalculationJobsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCalculationJobsResponse
func (client *Client) ListCalculationJobsWithContext(ctx context.Context, request *ListCalculationJobsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListCalculationJobsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCalculationJobs"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/calculationjobs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCalculationJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取人群下的所有用户。
//
// @param request - ListCrowdUsersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCrowdUsersResponse
func (client *Client) ListCrowdUsersWithContext(ctx context.Context, CrowdId *string, request *ListCrowdUsersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListCrowdUsersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCrowdUsers"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/crowds/" + dara.PercentEncode(dara.StringValue(CrowdId)) + "/users"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCrowdUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取人群列表。
//
// @param request - ListCrowdsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCrowdsResponse
func (client *Client) ListCrowdsWithContext(ctx context.Context, request *ListCrowdsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListCrowdsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCrowds"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/crowds"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCrowdsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取引擎配置列表。
//
// @param request - ListEngineConfigsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEngineConfigsResponse
func (client *Client) ListEngineConfigsWithContext(ctx context.Context, request *ListEngineConfigsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListEngineConfigsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Environment) {
		query["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEngineConfigs"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/engineconfigs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEngineConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实验组列表。
//
// @param request - ListExperimentGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExperimentGroupsResponse
func (client *Client) ListExperimentGroupsWithContext(ctx context.Context, request *ListExperimentGroupsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListExperimentGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LayerId) {
		query["LayerId"] = request.LayerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TimeRangeEnd) {
		query["TimeRangeEnd"] = request.TimeRangeEnd
	}

	if !dara.IsNil(request.TimeRangeStart) {
		query["TimeRangeStart"] = request.TimeRangeStart
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListExperimentGroups"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experimentgroups"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListExperimentGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实验列表。
//
// @param request - ListExperimentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExperimentsResponse
func (client *Client) ListExperimentsWithContext(ctx context.Context, request *ListExperimentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListExperimentsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExperimentGroupId) {
		query["ExperimentGroupId"] = request.ExperimentGroupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListExperiments"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListExperimentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取特征一致性检查配置列表。
//
// @param request - ListFeatureConsistencyCheckJobConfigsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFeatureConsistencyCheckJobConfigsResponse
func (client *Client) ListFeatureConsistencyCheckJobConfigsWithContext(ctx context.Context, request *ListFeatureConsistencyCheckJobConfigsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListFeatureConsistencyCheckJobConfigsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFeatureConsistencyCheckJobConfigs"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/featureconsistencycheck/jobconfigs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFeatureConsistencyCheckJobConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取特征一致性检查任务的特征报表/比对结果。
//
// @param request - ListFeatureConsistencyCheckJobFeatureReportsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFeatureConsistencyCheckJobFeatureReportsResponse
func (client *Client) ListFeatureConsistencyCheckJobFeatureReportsWithContext(ctx context.Context, FeatureConsistencyCheckJobId *string, request *ListFeatureConsistencyCheckJobFeatureReportsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListFeatureConsistencyCheckJobFeatureReportsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LogItemId) {
		query["LogItemId"] = request.LogItemId
	}

	if !dara.IsNil(request.LogRequestId) {
		query["LogRequestId"] = request.LogRequestId
	}

	if !dara.IsNil(request.LogUserId) {
		query["LogUserId"] = request.LogUserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFeatureConsistencyCheckJobFeatureReports"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/featureconsistencycheck/jobs/" + dara.PercentEncode(dara.StringValue(FeatureConsistencyCheckJobId)) + "/featurereports"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFeatureConsistencyCheckJobFeatureReportsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取特征一致性检查任务分数报表/比对结果。
//
// @param tmpReq - ListFeatureConsistencyCheckJobScoreReportsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFeatureConsistencyCheckJobScoreReportsResponse
func (client *Client) ListFeatureConsistencyCheckJobScoreReportsWithContext(ctx context.Context, FeatureConsistencyCheckJobId *string, tmpReq *ListFeatureConsistencyCheckJobScoreReportsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListFeatureConsistencyCheckJobScoreReportsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListFeatureConsistencyCheckJobScoreReportsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ExcludeRequestIds) {
		request.ExcludeRequestIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExcludeRequestIds, dara.String("ExcludeRequestIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ExcludeRequestIdsShrink) {
		query["ExcludeRequestIds"] = request.ExcludeRequestIdsShrink
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFeatureConsistencyCheckJobScoreReports"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/featureconsistencycheck/jobs/" + dara.PercentEncode(dara.StringValue(FeatureConsistencyCheckJobId)) + "/scorereports"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFeatureConsistencyCheckJobScoreReportsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取特征一致性检查任务列表。
//
// @param request - ListFeatureConsistencyCheckJobsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFeatureConsistencyCheckJobsResponse
func (client *Client) ListFeatureConsistencyCheckJobsWithContext(ctx context.Context, request *ListFeatureConsistencyCheckJobsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListFeatureConsistencyCheckJobsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFeatureConsistencyCheckJobs"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/featureconsistencycheck/jobs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFeatureConsistencyCheckJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实例下配置的资源列表。
//
// @param request - ListInstanceResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceResourcesResponse
func (client *Client) ListInstanceResourcesWithContext(ctx context.Context, InstanceId *string, request *ListInstanceResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstanceResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.Group) {
		query["Group"] = request.Group
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstanceResources"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/resources"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取推荐全链路深度定制开发平台实例信息列表。
//
// @param request - ListInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithContext(ctx context.Context, request *ListInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstances"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实验室列表。
//
// @param request - ListLaboratoriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLaboratoriesResponse
func (client *Client) ListLaboratoriesWithContext(ctx context.Context, request *ListLaboratoriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListLaboratoriesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Environment) {
		query["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLaboratories"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/laboratories"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLaboratoriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取层列表。
//
// @param request - ListLayersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLayersResponse
func (client *Client) ListLayersWithContext(ctx context.Context, request *ListLayersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListLayersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LaboratoryId) {
		query["LaboratoryId"] = request.LaboratoryId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLayers"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/layers"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLayersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取参数列表。
//
// @param request - ListParamsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListParamsResponse
func (client *Client) ListParamsWithContext(ctx context.Context, request *ListParamsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListParamsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Encrypted) {
		query["Encrypted"] = request.Encrypted
	}

	if !dara.IsNil(request.Environment) {
		query["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListParams"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/params"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListParamsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取资源规则列表
//
// @param request - ListResourceRulesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceRulesResponse
func (client *Client) ListResourceRulesWithContext(ctx context.Context, request *ListResourceRulesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListResourceRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.ResourceRuleId) {
		query["ResourceRuleId"] = request.ResourceRuleId
	}

	if !dara.IsNil(request.ResourceRuleName) {
		query["ResourceRuleName"] = request.ResourceRuleName
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceRules"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resourcerules"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取样本一致性任务列表
//
// @param request - ListSampleConsistencyJobsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSampleConsistencyJobsResponse
func (client *Client) ListSampleConsistencyJobsWithContext(ctx context.Context, request *ListSampleConsistencyJobsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSampleConsistencyJobsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSampleConsistencyJobs"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/sampleconsistencyjobs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSampleConsistencyJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取场景列表
//
// @param request - ListScenesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScenesResponse
func (client *Client) ListScenesWithContext(ctx context.Context, request *ListScenesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListScenesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListScenes"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/scenes"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListScenesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取人群下的所有子人群。
//
// @param request - ListSubCrowdsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSubCrowdsResponse
func (client *Client) ListSubCrowdsWithContext(ctx context.Context, CrowdId *string, request *ListSubCrowdsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSubCrowdsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSubCrowds"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/crowds/" + dara.PercentEncode(dara.StringValue(CrowdId)) + "/subcrowds"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSubCrowdsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据表列表。
//
// @param request - ListTableMetasRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTableMetasResponse
func (client *Client) ListTableMetasWithContext(ctx context.Context, request *ListTableMetasRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTableMetasResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Module) {
		query["Module"] = request.Module
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

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTableMetas"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/tablemetas"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTableMetasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取流量调控任务流量变更的历史列表
//
// @param request - ListTrafficControlTargetTrafficHistoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTrafficControlTargetTrafficHistoryResponse
func (client *Client) ListTrafficControlTargetTrafficHistoryWithContext(ctx context.Context, TrafficControlTargetId *string, request *ListTrafficControlTargetTrafficHistoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTrafficControlTargetTrafficHistoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Environment) {
		query["Environment"] = request.Environment
	}

	if !dara.IsNil(request.ExperimentGroupId) {
		query["ExperimentGroupId"] = request.ExperimentGroupId
	}

	if !dara.IsNil(request.ExperimentId) {
		query["ExperimentId"] = request.ExperimentId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ItemId) {
		query["ItemId"] = request.ItemId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Threshold) {
		query["Threshold"] = request.Threshold
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTrafficControlTargetTrafficHistory"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltargets/" + dara.PercentEncode(dara.StringValue(TrafficControlTargetId)) + "/traffichistories"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTrafficControlTargetTrafficHistoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取流量调控列表
//
// @param request - ListTrafficControlTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTrafficControlTasksResponse
func (client *Client) ListTrafficControlTasksWithContext(ctx context.Context, request *ListTrafficControlTasksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTrafficControlTasksResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.ControlTargetFilter) {
		query["ControlTargetFilter"] = request.ControlTargetFilter
	}

	if !dara.IsNil(request.Environment) {
		query["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
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

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TrafficControlTaskId) {
		query["TrafficControlTaskId"] = request.TrafficControlTaskId
	}

	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTrafficControlTasks"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltasks"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTrafficControlTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上线实验。
//
// @param request - OfflineExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OfflineExperimentResponse
func (client *Client) OfflineExperimentWithContext(ctx context.Context, ExperimentId *string, request *OfflineExperimentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OfflineExperimentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OfflineExperiment"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments/" + dara.PercentEncode(dara.StringValue(ExperimentId)) + "/action/offline"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &OfflineExperimentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 下线实验组。
//
// @param request - OfflineExperimentGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OfflineExperimentGroupResponse
func (client *Client) OfflineExperimentGroupWithContext(ctx context.Context, ExperimentGroupId *string, request *OfflineExperimentGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OfflineExperimentGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OfflineExperimentGroup"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experimentgroups/" + dara.PercentEncode(dara.StringValue(ExperimentGroupId)) + "/action/offline"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &OfflineExperimentGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 下线实验室。
//
// @param request - OfflineLaboratoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OfflineLaboratoryResponse
func (client *Client) OfflineLaboratoryWithContext(ctx context.Context, LaboratoryId *string, request *OfflineLaboratoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OfflineLaboratoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OfflineLaboratory"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/laboratories/" + dara.PercentEncode(dara.StringValue(LaboratoryId)) + "/action/offline"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &OfflineLaboratoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上线实验
//
// @param request - OnlineExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnlineExperimentResponse
func (client *Client) OnlineExperimentWithContext(ctx context.Context, ExperimentId *string, request *OnlineExperimentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OnlineExperimentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OnlineExperiment"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments/" + dara.PercentEncode(dara.StringValue(ExperimentId)) + "/action/online"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &OnlineExperimentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上线实验组。
//
// @param request - OnlineExperimentGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnlineExperimentGroupResponse
func (client *Client) OnlineExperimentGroupWithContext(ctx context.Context, ExperimentGroupId *string, request *OnlineExperimentGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OnlineExperimentGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OnlineExperimentGroup"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experimentgroups/" + dara.PercentEncode(dara.StringValue(ExperimentGroupId)) + "/action/online"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &OnlineExperimentGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上线实验室。
//
// @param request - OnlineLaboratoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnlineLaboratoryResponse
func (client *Client) OnlineLaboratoryWithContext(ctx context.Context, LaboratoryId *string, request *OnlineLaboratoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OnlineLaboratoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OnlineLaboratory"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/laboratories/" + dara.PercentEncode(dara.StringValue(LaboratoryId)) + "/action/online"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &OnlineLaboratoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 推全。
//
// @param request - PushAllExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushAllExperimentResponse
func (client *Client) PushAllExperimentWithContext(ctx context.Context, ExperimentId *string, request *PushAllExperimentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PushAllExperimentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushAllExperiment"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments/" + dara.PercentEncode(dara.StringValue(ExperimentId)) + "/action/pushall"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PushAllExperimentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 推送指标到指定资源规则
//
// @param tmpReq - PushResourceRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushResourceRuleResponse
func (client *Client) PushResourceRuleWithContext(ctx context.Context, ResourceRuleId *string, tmpReq *PushResourceRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PushResourceRuleResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &PushResourceRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MetricInfo) {
		request.MetricInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MetricInfo, dara.String("MetricInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MetricInfoShrink) {
		query["MetricInfo"] = request.MetricInfoShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushResourceRule"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resourcerules/" + dara.PercentEncode(dara.StringValue(ResourceRuleId)) + "/action/push"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PushResourceRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看样本一致性任务差异的详情
//
// @param request - QuerySampleConsistencyJobDifferenceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySampleConsistencyJobDifferenceResponse
func (client *Client) QuerySampleConsistencyJobDifferenceWithContext(ctx context.Context, SampleConsistencyJobId *string, request *QuerySampleConsistencyJobDifferenceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QuerySampleConsistencyJobDifferenceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FeatureName) {
		query["FeatureName"] = request.FeatureName
	}

	if !dara.IsNil(request.FeatureType) {
		query["FeatureType"] = request.FeatureType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySampleConsistencyJobDifference"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/sampleconsistencyjobs/" + dara.PercentEncode(dara.StringValue(SampleConsistencyJobId)) + "/action/querydifference"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySampleConsistencyJobDifferenceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询流量调控目标的单品调控报表详情。
//
// @param request - QueryTrafficControlTargetItemReportDetailRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTrafficControlTargetItemReportDetailResponse
func (client *Client) QueryTrafficControlTargetItemReportDetailWithContext(ctx context.Context, TrafficControlTargetId *string, request *QueryTrafficControlTargetItemReportDetailRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryTrafficControlTargetItemReportDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Date) {
		query["Date"] = request.Date
	}

	if !dara.IsNil(request.Environment) {
		query["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTrafficControlTargetItemReportDetail"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltargets/" + dara.PercentEncode(dara.StringValue(TrafficControlTargetId)) + "/itemcontrolreportdetail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTrafficControlTargetItemReportDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发布流量调控任务
//
// @param request - ReleaseTrafficControlTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseTrafficControlTaskResponse
func (client *Client) ReleaseTrafficControlTaskWithContext(ctx context.Context, TrafficControlTaskId *string, request *ReleaseTrafficControlTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ReleaseTrafficControlTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Environment) {
		body["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseTrafficControlTask"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltasks/" + dara.PercentEncode(dara.StringValue(TrafficControlTaskId)) + "/action/release"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseTrafficControlTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 对指标组进行报表。
//
// @param request - ReportABMetricGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReportABMetricGroupResponse
func (client *Client) ReportABMetricGroupWithContext(ctx context.Context, ABMetricGroupId *string, request *ReportABMetricGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ReportABMetricGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseExperimentId) {
		body["BaseExperimentId"] = request.BaseExperimentId
	}

	if !dara.IsNil(request.DimensionFields) {
		body["DimensionFields"] = request.DimensionFields
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.ExperimentGroupId) {
		body["ExperimentGroupId"] = request.ExperimentGroupId
	}

	if !dara.IsNil(request.ExperimentIds) {
		body["ExperimentIds"] = request.ExperimentIds
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ReportType) {
		body["ReportType"] = request.ReportType
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.TimeStatisticsMethod) {
		body["TimeStatisticsMethod"] = request.TimeStatisticsMethod
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReportABMetricGroup"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/abmetricgroups/" + dara.PercentEncode(dara.StringValue(ABMetricGroupId)) + "/action/report"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ReportABMetricGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 样本一致性任务报表
//
// @param request - ReportSampleConsistencyJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReportSampleConsistencyJobResponse
func (client *Client) ReportSampleConsistencyJobWithContext(ctx context.Context, SampleConsistencyJobId *string, request *ReportSampleConsistencyJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ReportSampleConsistencyJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReportSampleConsistencyJob"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/sampleconsistencyjobs/" + dara.PercentEncode(dara.StringValue(SampleConsistencyJobId)) + "/action/report"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ReportSampleConsistencyJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 拆分流量调控目标
//
// @param request - SplitTrafficControlTargetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SplitTrafficControlTargetResponse
func (client *Client) SplitTrafficControlTargetWithContext(ctx context.Context, TrafficControlTargetId *string, request *SplitTrafficControlTargetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SplitTrafficControlTargetResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Environment) {
		body["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SetPoints) {
		body["SetPoints"] = request.SetPoints
	}

	if !dara.IsNil(request.SetValues) {
		body["SetValues"] = request.SetValues
	}

	if !dara.IsNil(request.TimePoints) {
		body["TimePoints"] = request.TimePoints
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SplitTrafficControlTarget"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltargets/" + dara.PercentEncode(dara.StringValue(TrafficControlTargetId)) + "/action/split"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SplitTrafficControlTargetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开启流量调控目标
//
// @param request - StartTrafficControlTargetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartTrafficControlTargetResponse
func (client *Client) StartTrafficControlTargetWithContext(ctx context.Context, TrafficControlTargetId *string, request *StartTrafficControlTargetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartTrafficControlTargetResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartTrafficControlTarget"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltargets/" + dara.PercentEncode(dara.StringValue(TrafficControlTargetId)) + "/action/start"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StartTrafficControlTargetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开启流量调控任务
//
// @param request - StartTrafficControlTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartTrafficControlTaskResponse
func (client *Client) StartTrafficControlTaskWithContext(ctx context.Context, TrafficControlTaskId *string, request *StartTrafficControlTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartTrafficControlTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Environment) {
		body["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartTrafficControlTask"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltasks/" + dara.PercentEncode(dara.StringValue(TrafficControlTaskId)) + "/action/start"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StartTrafficControlTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止样本一致性任务
//
// @param request - StopSampleConsistencyJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopSampleConsistencyJobResponse
func (client *Client) StopSampleConsistencyJobWithContext(ctx context.Context, SampleConsistencyJobId *string, request *StopSampleConsistencyJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopSampleConsistencyJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopSampleConsistencyJob"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/sampleconsistencyjobs/" + dara.PercentEncode(dara.StringValue(SampleConsistencyJobId)) + "/action/stop"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopSampleConsistencyJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止流量调控目标
//
// @param request - StopTrafficControlTargetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopTrafficControlTargetResponse
func (client *Client) StopTrafficControlTargetWithContext(ctx context.Context, TrafficControlTargetId *string, request *StopTrafficControlTargetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopTrafficControlTargetResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopTrafficControlTarget"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltargets/" + dara.PercentEncode(dara.StringValue(TrafficControlTargetId)) + "/action/stop"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopTrafficControlTargetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止流量调控任务
//
// @param request - StopTrafficControlTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopTrafficControlTaskResponse
func (client *Client) StopTrafficControlTaskWithContext(ctx context.Context, TrafficControlTaskId *string, request *StopTrafficControlTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopTrafficControlTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Environment) {
		body["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopTrafficControlTask"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltasks/" + dara.PercentEncode(dara.StringValue(TrafficControlTaskId)) + "/action/stop"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopTrafficControlTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 同步特征一致性检测任务重放日志。
//
// @param request - SyncFeatureConsistencyCheckJobReplayLogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncFeatureConsistencyCheckJobReplayLogResponse
func (client *Client) SyncFeatureConsistencyCheckJobReplayLogWithContext(ctx context.Context, request *SyncFeatureConsistencyCheckJobReplayLogRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SyncFeatureConsistencyCheckJobReplayLogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ContextFeatures) {
		body["ContextFeatures"] = request.ContextFeatures
	}

	if !dara.IsNil(request.FeatureConsistencyCheckJobConfigId) {
		body["FeatureConsistencyCheckJobConfigId"] = request.FeatureConsistencyCheckJobConfigId
	}

	if !dara.IsNil(request.GeneratedFeatures) {
		body["GeneratedFeatures"] = request.GeneratedFeatures
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LogItemId) {
		body["LogItemId"] = request.LogItemId
	}

	if !dara.IsNil(request.LogRequestId) {
		body["LogRequestId"] = request.LogRequestId
	}

	if !dara.IsNil(request.LogRequestTime) {
		body["LogRequestTime"] = request.LogRequestTime
	}

	if !dara.IsNil(request.LogUserId) {
		body["LogUserId"] = request.LogUserId
	}

	if !dara.IsNil(request.RawFeatures) {
		body["RawFeatures"] = request.RawFeatures
	}

	if !dara.IsNil(request.SceneName) {
		body["SceneName"] = request.SceneName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SyncFeatureConsistencyCheckJobReplayLog"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/featureconsistencycheck/jobs/action/syncreplaylog"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SyncFeatureConsistencyCheckJobReplayLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消指定特征一致性检查正在运行中的任务。
//
// @param request - TerminateFeatureConsistencyCheckJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TerminateFeatureConsistencyCheckJobResponse
func (client *Client) TerminateFeatureConsistencyCheckJobWithContext(ctx context.Context, FeatureConsistencyCheckJobId *string, request *TerminateFeatureConsistencyCheckJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TerminateFeatureConsistencyCheckJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TerminateFeatureConsistencyCheckJob"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/featureconsistencycheck/jobs/" + dara.PercentEncode(dara.StringValue(FeatureConsistencyCheckJobId)) + "/action/terminate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TerminateFeatureConsistencyCheckJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新AB Test实验指标。
//
// @param request - UpdateABMetricRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateABMetricResponse
func (client *Client) UpdateABMetricWithContext(ctx context.Context, ABMetricId *string, request *UpdateABMetricRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateABMetricResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Definition) {
		body["Definition"] = request.Definition
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LeftMetricId) {
		body["LeftMetricId"] = request.LeftMetricId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Operator) {
		body["Operator"] = request.Operator
	}

	if !dara.IsNil(request.Realtime) {
		body["Realtime"] = request.Realtime
	}

	if !dara.IsNil(request.ResultResourceId) {
		body["ResultResourceId"] = request.ResultResourceId
	}

	if !dara.IsNil(request.RightMetricId) {
		body["RightMetricId"] = request.RightMetricId
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.StatisticsCycle) {
		body["StatisticsCycle"] = request.StatisticsCycle
	}

	if !dara.IsNil(request.TableMetaId) {
		body["TableMetaId"] = request.TableMetaId
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateABMetric"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/abmetrics/" + dara.PercentEncode(dara.StringValue(ABMetricId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateABMetricResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新AB test实验指标组。
//
// @param request - UpdateABMetricGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateABMetricGroupResponse
func (client *Client) UpdateABMetricGroupWithContext(ctx context.Context, ABMetricGroupId *string, request *UpdateABMetricGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateABMetricGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ABMetricIds) {
		body["ABMetricIds"] = request.ABMetricIds
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Realtime) {
		body["Realtime"] = request.Realtime
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateABMetricGroup"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/abmetricgroups/" + dara.PercentEncode(dara.StringValue(ABMetricGroupId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateABMetricGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新指定人群。
//
// @param request - UpdateCrowdRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCrowdResponse
func (client *Client) UpdateCrowdWithContext(ctx context.Context, CrowdId *string, request *UpdateCrowdRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateCrowdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCrowd"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/crowds/" + dara.PercentEncode(dara.StringValue(CrowdId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCrowdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新引擎配置。
//
// @param request - UpdateEngineConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEngineConfigResponse
func (client *Client) UpdateEngineConfigWithContext(ctx context.Context, EngineConfigId *string, request *UpdateEngineConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateEngineConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigValue) {
		body["ConfigValue"] = request.ConfigValue
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Environment) {
		body["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEngineConfig"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/engineconfigs/" + dara.PercentEncode(dara.StringValue(EngineConfigId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEngineConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新实验。
//
// @param request - UpdateExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateExperimentResponse
func (client *Client) UpdateExperimentWithContext(ctx context.Context, ExperimentId *string, request *UpdateExperimentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateExperimentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["Config"] = request.Config
	}

	if !dara.IsNil(request.DebugCrowdId) {
		body["DebugCrowdId"] = request.DebugCrowdId
	}

	if !dara.IsNil(request.DebugUsers) {
		body["DebugUsers"] = request.DebugUsers
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.FlowPercent) {
		body["FlowPercent"] = request.FlowPercent
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateExperiment"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments/" + dara.PercentEncode(dara.StringValue(ExperimentId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateExperimentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新指定实验组。
//
// @param request - UpdateExperimentGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateExperimentGroupResponse
func (client *Client) UpdateExperimentGroupWithContext(ctx context.Context, ExperimentGroupId *string, request *UpdateExperimentGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateExperimentGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["Config"] = request.Config
	}

	if !dara.IsNil(request.CrowdId) {
		body["CrowdId"] = request.CrowdId
	}

	if !dara.IsNil(request.CrowdTargetType) {
		body["CrowdTargetType"] = request.CrowdTargetType
	}

	if !dara.IsNil(request.DebugCrowdId) {
		body["DebugCrowdId"] = request.DebugCrowdId
	}

	if !dara.IsNil(request.DebugUsers) {
		body["DebugUsers"] = request.DebugUsers
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DistributionTimeDuration) {
		body["DistributionTimeDuration"] = request.DistributionTimeDuration
	}

	if !dara.IsNil(request.DistributionType) {
		body["DistributionType"] = request.DistributionType
	}

	if !dara.IsNil(request.Filter) {
		body["Filter"] = request.Filter
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LayerId) {
		body["LayerId"] = request.LayerId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NeedAA) {
		body["NeedAA"] = request.NeedAA
	}

	if !dara.IsNil(request.RandomFlow) {
		body["RandomFlow"] = request.RandomFlow
	}

	if !dara.IsNil(request.ReservcedBuckets) {
		body["ReservcedBuckets"] = request.ReservcedBuckets
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateExperimentGroup"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experimentgroups/" + dara.PercentEncode(dara.StringValue(ExperimentGroupId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateExperimentGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新特征一致性检查配置信息。
//
// @param request - UpdateFeatureConsistencyCheckJobConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFeatureConsistencyCheckJobConfigResponse
func (client *Client) UpdateFeatureConsistencyCheckJobConfigWithContext(ctx context.Context, FeatureConsistencyCheckJobConfigId *string, request *UpdateFeatureConsistencyCheckJobConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateFeatureConsistencyCheckJobConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CompareFeature) {
		body["CompareFeature"] = request.CompareFeature
	}

	if !dara.IsNil(request.DatasetId) {
		body["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.DatasetMountPath) {
		body["DatasetMountPath"] = request.DatasetMountPath
	}

	if !dara.IsNil(request.DatasetName) {
		body["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.DatasetType) {
		body["DatasetType"] = request.DatasetType
	}

	if !dara.IsNil(request.DatasetUri) {
		body["DatasetUri"] = request.DatasetUri
	}

	if !dara.IsNil(request.DefaultRoute) {
		body["DefaultRoute"] = request.DefaultRoute
	}

	if !dara.IsNil(request.EasServiceName) {
		body["EasServiceName"] = request.EasServiceName
	}

	if !dara.IsNil(request.EasyRecPackagePath) {
		body["EasyRecPackagePath"] = request.EasyRecPackagePath
	}

	if !dara.IsNil(request.EasyRecVersion) {
		body["EasyRecVersion"] = request.EasyRecVersion
	}

	if !dara.IsNil(request.FeatureDisplayExclude) {
		body["FeatureDisplayExclude"] = request.FeatureDisplayExclude
	}

	if !dara.IsNil(request.FeatureLandingResourceId) {
		body["FeatureLandingResourceId"] = request.FeatureLandingResourceId
	}

	if !dara.IsNil(request.FeaturePriority) {
		body["FeaturePriority"] = request.FeaturePriority
	}

	if !dara.IsNil(request.FeatureStoreItemId) {
		body["FeatureStoreItemId"] = request.FeatureStoreItemId
	}

	if !dara.IsNil(request.FeatureStoreModelId) {
		body["FeatureStoreModelId"] = request.FeatureStoreModelId
	}

	if !dara.IsNil(request.FeatureStoreProjectId) {
		body["FeatureStoreProjectId"] = request.FeatureStoreProjectId
	}

	if !dara.IsNil(request.FeatureStoreProjectName) {
		body["FeatureStoreProjectName"] = request.FeatureStoreProjectName
	}

	if !dara.IsNil(request.FeatureStoreSeqFeatureView) {
		body["FeatureStoreSeqFeatureView"] = request.FeatureStoreSeqFeatureView
	}

	if !dara.IsNil(request.FeatureStoreUserId) {
		body["FeatureStoreUserId"] = request.FeatureStoreUserId
	}

	if !dara.IsNil(request.FgJarVersion) {
		body["FgJarVersion"] = request.FgJarVersion
	}

	if !dara.IsNil(request.FgJsonFileName) {
		body["FgJsonFileName"] = request.FgJsonFileName
	}

	if !dara.IsNil(request.GenerateZip) {
		body["GenerateZip"] = request.GenerateZip
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IsUseFeatureStore) {
		body["IsUseFeatureStore"] = request.IsUseFeatureStore
	}

	if !dara.IsNil(request.ItemIdField) {
		body["ItemIdField"] = request.ItemIdField
	}

	if !dara.IsNil(request.ItemTable) {
		body["ItemTable"] = request.ItemTable
	}

	if !dara.IsNil(request.ItemTablePartitionField) {
		body["ItemTablePartitionField"] = request.ItemTablePartitionField
	}

	if !dara.IsNil(request.ItemTablePartitionFieldFormat) {
		body["ItemTablePartitionFieldFormat"] = request.ItemTablePartitionFieldFormat
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.OssResourceId) {
		body["OssResourceId"] = request.OssResourceId
	}

	if !dara.IsNil(request.PredictWorkerCount) {
		body["PredictWorkerCount"] = request.PredictWorkerCount
	}

	if !dara.IsNil(request.PredictWorkerCpu) {
		body["PredictWorkerCpu"] = request.PredictWorkerCpu
	}

	if !dara.IsNil(request.PredictWorkerMemory) {
		body["PredictWorkerMemory"] = request.PredictWorkerMemory
	}

	if !dara.IsNil(request.ResourceConfig) {
		body["ResourceConfig"] = request.ResourceConfig
	}

	if !dara.IsNil(request.SampleRate) {
		body["SampleRate"] = request.SampleRate
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.SecurityGroupId) {
		body["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.ServiceId) {
		body["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.SwitchId) {
		body["SwitchId"] = request.SwitchId
	}

	if !dara.IsNil(request.UserIdField) {
		body["UserIdField"] = request.UserIdField
	}

	if !dara.IsNil(request.UserTable) {
		body["UserTable"] = request.UserTable
	}

	if !dara.IsNil(request.UserTablePartitionField) {
		body["UserTablePartitionField"] = request.UserTablePartitionField
	}

	if !dara.IsNil(request.UserTablePartitionFieldFormat) {
		body["UserTablePartitionFieldFormat"] = request.UserTablePartitionFieldFormat
	}

	if !dara.IsNil(request.VpcId) {
		body["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.WorkflowName) {
		body["WorkflowName"] = request.WorkflowName
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFeatureConsistencyCheckJobConfig"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/featureconsistencycheck/jobconfigs/" + dara.PercentEncode(dara.StringValue(FeatureConsistencyCheckJobConfigId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFeatureConsistencyCheckJobConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新指定实例下指定资源的信息。
//
// @param request - UpdateInstanceResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceResourceResponse
func (client *Client) UpdateInstanceResourceWithContext(ctx context.Context, InstanceId *string, ResourceId *string, request *UpdateInstanceResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateInstanceResourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["Config"] = request.Config
	}

	if !dara.IsNil(request.Uri) {
		body["Uri"] = request.Uri
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstanceResource"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/resources/" + dara.PercentEncode(dara.StringValue(ResourceId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstanceResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新实验室。
//
// @param request - UpdateLaboratoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLaboratoryResponse
func (client *Client) UpdateLaboratoryWithContext(ctx context.Context, LaboratoryId *string, request *UpdateLaboratoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateLaboratoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BucketCount) {
		body["BucketCount"] = request.BucketCount
	}

	if !dara.IsNil(request.BucketType) {
		body["BucketType"] = request.BucketType
	}

	if !dara.IsNil(request.Buckets) {
		body["Buckets"] = request.Buckets
	}

	if !dara.IsNil(request.DebugCrowdId) {
		body["DebugCrowdId"] = request.DebugCrowdId
	}

	if !dara.IsNil(request.DebugUsers) {
		body["DebugUsers"] = request.DebugUsers
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Environment) {
		body["Environment"] = request.Environment
	}

	if !dara.IsNil(request.Filter) {
		body["Filter"] = request.Filter
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLaboratory"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/laboratories/" + dara.PercentEncode(dara.StringValue(LaboratoryId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLaboratoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新层。
//
// @param request - UpdateLayerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLayerResponse
func (client *Client) UpdateLayerWithContext(ctx context.Context, LayerId *string, request *UpdateLayerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateLayerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLayer"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/layers/" + dara.PercentEncode(dara.StringValue(LayerId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLayerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新参数。
//
// @param request - UpdateParamRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateParamResponse
func (client *Client) UpdateParamWithContext(ctx context.Context, ParamId *string, request *UpdateParamRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateParamResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Value) {
		body["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateParam"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/params/" + dara.PercentEncode(dara.StringValue(ParamId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateParamResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取资源规则列表
//
// @param request - UpdateResourceRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceRuleResponse
func (client *Client) UpdateResourceRuleWithContext(ctx context.Context, ResourceRuleId *string, request *UpdateResourceRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateResourceRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MetricOperationType) {
		body["MetricOperationType"] = request.MetricOperationType
	}

	if !dara.IsNil(request.MetricPullInfo) {
		body["MetricPullInfo"] = request.MetricPullInfo
	}

	if !dara.IsNil(request.MetricPullPeriod) {
		body["MetricPullPeriod"] = request.MetricPullPeriod
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.RuleComputingDefinition) {
		body["RuleComputingDefinition"] = request.RuleComputingDefinition
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateResourceRule"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resourcerules/" + dara.PercentEncode(dara.StringValue(ResourceRuleId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateResourceRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新资源规则条目
//
// @param request - UpdateResourceRuleItemRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceRuleItemResponse
func (client *Client) UpdateResourceRuleItemWithContext(ctx context.Context, ResourceRuleId *string, ResourceRuleItemId *string, request *UpdateResourceRuleItemRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateResourceRuleItemResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxValue) {
		body["MaxValue"] = request.MaxValue
	}

	if !dara.IsNil(request.MinValue) {
		body["MinValue"] = request.MinValue
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Value) {
		body["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateResourceRuleItem"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resourcerules/" + dara.PercentEncode(dara.StringValue(ResourceRuleId)) + "/items/" + dara.PercentEncode(dara.StringValue(ResourceRuleItemId))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateResourceRuleItemResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新场景
//
// @param request - UpdateSceneRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSceneResponse
func (client *Client) UpdateSceneWithContext(ctx context.Context, SceneId *string, request *UpdateSceneRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateSceneResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Flows) {
		body["Flows"] = request.Flows
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateScene"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/scenes/" + dara.PercentEncode(dara.StringValue(SceneId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSceneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据表详细信息。
//
// @param request - UpdateTableMetaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTableMetaResponse
func (client *Client) UpdateTableMetaWithContext(ctx context.Context, TableMetaId *string, request *UpdateTableMetaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateTableMetaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Fields) {
		body["Fields"] = request.Fields
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Module) {
		body["Module"] = request.Module
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ResourceId) {
		body["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.TableName) {
		body["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTableMeta"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/tablemetas/" + dara.PercentEncode(dara.StringValue(TableMetaId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTableMetaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新流量调控目标
//
// @param request - UpdateTrafficControlTargetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTrafficControlTargetResponse
func (client *Client) UpdateTrafficControlTargetWithContext(ctx context.Context, TrafficControlTargetId *string, request *UpdateTrafficControlTargetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateTrafficControlTargetResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NewParam3) {
		query["new-param-3"] = request.NewParam3
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Event) {
		body["Event"] = request.Event
	}

	if !dara.IsNil(request.ItemConditionArray) {
		body["ItemConditionArray"] = request.ItemConditionArray
	}

	if !dara.IsNil(request.ItemConditionExpress) {
		body["ItemConditionExpress"] = request.ItemConditionExpress
	}

	if !dara.IsNil(request.ItemConditionType) {
		body["ItemConditionType"] = request.ItemConditionType
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NewProductRegulation) {
		body["NewProductRegulation"] = request.NewProductRegulation
	}

	if !dara.IsNil(request.RecallName) {
		body["RecallName"] = request.RecallName
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StatisPeriod) {
		body["StatisPeriod"] = request.StatisPeriod
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.ToleranceValue) {
		body["ToleranceValue"] = request.ToleranceValue
	}

	if !dara.IsNil(request.Value) {
		body["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTrafficControlTarget"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltargets/" + dara.PercentEncode(dara.StringValue(TrafficControlTargetId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTrafficControlTargetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新流量调控任务
//
// @param request - UpdateTrafficControlTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTrafficControlTaskResponse
func (client *Client) UpdateTrafficControlTaskWithContext(ctx context.Context, TrafficControlTaskId *string, request *UpdateTrafficControlTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateTrafficControlTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BehaviorTableMetaId) {
		body["BehaviorTableMetaId"] = request.BehaviorTableMetaId
	}

	if !dara.IsNil(request.ControlGranularity) {
		body["ControlGranularity"] = request.ControlGranularity
	}

	if !dara.IsNil(request.ControlLogic) {
		body["ControlLogic"] = request.ControlLogic
	}

	if !dara.IsNil(request.ControlType) {
		body["ControlType"] = request.ControlType
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EffectiveSceneIds) {
		body["EffectiveSceneIds"] = request.EffectiveSceneIds
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ExecutionTime) {
		body["ExecutionTime"] = request.ExecutionTime
	}

	if !dara.IsNil(request.FlinkResourceId) {
		body["FlinkResourceId"] = request.FlinkResourceId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ItemConditionArray) {
		body["ItemConditionArray"] = request.ItemConditionArray
	}

	if !dara.IsNil(request.ItemConditionExpress) {
		body["ItemConditionExpress"] = request.ItemConditionExpress
	}

	if !dara.IsNil(request.ItemConditionType) {
		body["ItemConditionType"] = request.ItemConditionType
	}

	if !dara.IsNil(request.ItemTableMetaId) {
		body["ItemTableMetaId"] = request.ItemTableMetaId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PreExperimentIds) {
		body["PreExperimentIds"] = request.PreExperimentIds
	}

	if !dara.IsNil(request.ProdExperimentIds) {
		body["ProdExperimentIds"] = request.ProdExperimentIds
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.ServiceId) {
		body["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceIds) {
		body["ServiceIds"] = request.ServiceIds
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StatisBaeaviorConditionArray) {
		body["StatisBaeaviorConditionArray"] = request.StatisBaeaviorConditionArray
	}

	if !dara.IsNil(request.StatisBehaviorConditionExpress) {
		body["StatisBehaviorConditionExpress"] = request.StatisBehaviorConditionExpress
	}

	if !dara.IsNil(request.StatisBehaviorConditionType) {
		body["StatisBehaviorConditionType"] = request.StatisBehaviorConditionType
	}

	if !dara.IsNil(request.TrafficControlTargets) {
		body["TrafficControlTargets"] = request.TrafficControlTargets
	}

	if !dara.IsNil(request.UserConditionArray) {
		body["UserConditionArray"] = request.UserConditionArray
	}

	if !dara.IsNil(request.UserConditionExpress) {
		body["UserConditionExpress"] = request.UserConditionExpress
	}

	if !dara.IsNil(request.UserConditionType) {
		body["UserConditionType"] = request.UserConditionType
	}

	if !dara.IsNil(request.UserTableMetaId) {
		body["UserTableMetaId"] = request.UserTableMetaId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTrafficControlTask"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltasks/" + dara.PercentEncode(dara.StringValue(TrafficControlTaskId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTrafficControlTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新流量调控任务的流量参数
//
// @param request - UpdateTrafficControlTaskTrafficRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTrafficControlTaskTrafficResponse
func (client *Client) UpdateTrafficControlTaskTrafficWithContext(ctx context.Context, TrafficControlTaskId *string, request *UpdateTrafficControlTaskTrafficRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateTrafficControlTaskTrafficResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NewParam3) {
		query["new-param-3"] = request.NewParam3
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Environment) {
		body["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Traffics) {
		body["Traffics"] = request.Traffics
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTrafficControlTaskTraffic"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltasks/" + dara.PercentEncode(dara.StringValue(TrafficControlTaskId)) + "/action/traffic"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTrafficControlTaskTrafficResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上传数据
//
// @param request - UploadRecommendationDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadRecommendationDataResponse
func (client *Client) UploadRecommendationDataWithContext(ctx context.Context, request *UploadRecommendationDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UploadRecommendationDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.DataType) {
		body["DataType"] = request.DataType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadRecommendationData"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/recommendationdata/action/upload"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadRecommendationDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
