// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Deprecated: OpenAPI AddAliClusterIdsToPrometheusGlobalView is deprecated
//
// Summary:
//
// Adds a global aggregation instance as a data source in Managed Service for Prometheus.
//
// @param request - AddAliClusterIdsToPrometheusGlobalViewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddAliClusterIdsToPrometheusGlobalViewResponse
func (client *Client) AddAliClusterIdsToPrometheusGlobalViewWithContext(ctx context.Context, request *AddAliClusterIdsToPrometheusGlobalViewRequest, runtime *dara.RuntimeOptions) (_result *AddAliClusterIdsToPrometheusGlobalViewResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterIds) {
		query["ClusterIds"] = request.ClusterIds
	}

	if !dara.IsNil(request.GlobalViewClusterId) {
		query["GlobalViewClusterId"] = request.GlobalViewClusterId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddAliClusterIdsToPrometheusGlobalView"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddAliClusterIdsToPrometheusGlobalViewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI AddGrafana is deprecated
//
// Summary:
//
// Integrates the dashboard of Prometheus Service.
//
// @param request - AddGrafanaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddGrafanaResponse
func (client *Client) AddGrafanaWithContext(ctx context.Context, request *AddGrafanaRequest, runtime *dara.RuntimeOptions) (_result *AddGrafanaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Integration) {
		query["Integration"] = request.Integration
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddGrafana"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddGrafanaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI AddIntegration is deprecated, please use ARMS::2019-08-08::InstallAddon instead.
//
// Summary:
//
// Integrates the dashboard and collection rules of Prometheus Service.
//
// @param request - AddIntegrationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddIntegrationResponse
func (client *Client) AddIntegrationWithContext(ctx context.Context, request *AddIntegrationRequest, runtime *dara.RuntimeOptions) (_result *AddIntegrationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Integration) {
		query["Integration"] = request.Integration
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddIntegration"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddIntegrationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI AddPrometheusGlobalView is deprecated
//
// Summary:
//
// Adds a global aggregation instance in Prometheus Service.
//
// @param request - AddPrometheusGlobalViewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddPrometheusGlobalViewResponse
func (client *Client) AddPrometheusGlobalViewWithContext(ctx context.Context, request *AddPrometheusGlobalViewRequest, runtime *dara.RuntimeOptions) (_result *AddPrometheusGlobalViewResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Clusters) {
		query["Clusters"] = request.Clusters
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
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
		Action:      dara.String("AddPrometheusGlobalView"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddPrometheusGlobalViewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI AddPrometheusGlobalViewByAliClusterIds is deprecated
//
// Summary:
//
// Adds a global aggregation instance in Managed Service for Prometheus.
//
// @param request - AddPrometheusGlobalViewByAliClusterIdsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddPrometheusGlobalViewByAliClusterIdsResponse
func (client *Client) AddPrometheusGlobalViewByAliClusterIdsWithContext(ctx context.Context, request *AddPrometheusGlobalViewByAliClusterIdsRequest, runtime *dara.RuntimeOptions) (_result *AddPrometheusGlobalViewByAliClusterIdsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterIds) {
		query["ClusterIds"] = request.ClusterIds
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddPrometheusGlobalViewByAliClusterIds"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddPrometheusGlobalViewByAliClusterIdsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI AddPrometheusInstance is deprecated, please use ARMS::2019-08-08::CreatePrometheusInstance instead.
//
// Summary:
//
// Creates a Prometheus instance for Remote Write.
//
// @param request - AddPrometheusInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddPrometheusInstanceResponse
func (client *Client) AddPrometheusInstanceWithContext(ctx context.Context, request *AddPrometheusInstanceRequest, runtime *dara.RuntimeOptions) (_result *AddPrometheusInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
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
		Action:      dara.String("AddPrometheusInstance"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddPrometheusInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI AddPrometheusIntegration is deprecated
//
// Summary:
//
// Adds an exporter to a Prometheus instance for Container Service or a Prometheus instance for ECS.
//
// @param request - AddPrometheusIntegrationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddPrometheusIntegrationResponse
func (client *Client) AddPrometheusIntegrationWithContext(ctx context.Context, request *AddPrometheusIntegrationRequest, runtime *dara.RuntimeOptions) (_result *AddPrometheusIntegrationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.IntegrationType) {
		query["IntegrationType"] = request.IntegrationType
	}

	if !dara.IsNil(request.Param) {
		query["Param"] = request.Param
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddPrometheusIntegration"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddPrometheusIntegrationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI AddRecordingRule is deprecated
//
// Summary:
//
// Creates or updates a recording rule of Managed Service for Prometheus.
//
// @param request - AddRecordingRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddRecordingRuleResponse
func (client *Client) AddRecordingRuleWithContext(ctx context.Context, request *AddRecordingRuleRequest, runtime *dara.RuntimeOptions) (_result *AddRecordingRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RuleYaml) {
		query["RuleYaml"] = request.RuleYaml
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddRecordingRule"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddRecordingRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI AddTagToFlinkCluster is deprecated
//
// Summary:
//
// Attaches the workspace ID and workspace name tags to the Prometheus instance corresponding to a Flink workspace.
//
// @param request - AddTagToFlinkClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTagToFlinkClusterResponse
func (client *Client) AddTagToFlinkClusterWithContext(ctx context.Context, request *AddTagToFlinkClusterRequest, runtime *dara.RuntimeOptions) (_result *AddTagToFlinkClusterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.FlinkWorkSpaceId) {
		query["FlinkWorkSpaceId"] = request.FlinkWorkSpaceId
	}

	if !dara.IsNil(request.FlinkWorkSpaceName) {
		query["FlinkWorkSpaceName"] = request.FlinkWorkSpaceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TargetUserId) {
		query["TargetUserId"] = request.TargetUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddTagToFlinkCluster"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddTagToFlinkClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI AppendInstancesToPrometheusGlobalView is deprecated
//
// Summary:
//
// Adds a data source to a global aggregation instance.
//
// @param request - AppendInstancesToPrometheusGlobalViewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AppendInstancesToPrometheusGlobalViewResponse
func (client *Client) AppendInstancesToPrometheusGlobalViewWithContext(ctx context.Context, request *AppendInstancesToPrometheusGlobalViewRequest, runtime *dara.RuntimeOptions) (_result *AppendInstancesToPrometheusGlobalViewResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Clusters) {
		query["Clusters"] = request.Clusters
	}

	if !dara.IsNil(request.GlobalViewClusterId) {
		query["GlobalViewClusterId"] = request.GlobalViewClusterId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AppendInstancesToPrometheusGlobalView"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AppendInstancesToPrometheusGlobalViewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param tmpReq - ApplyScenarioRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyScenarioResponse
func (client *Client) ApplyScenarioWithContext(ctx context.Context, tmpReq *ApplyScenarioRequest, runtime *dara.RuntimeOptions) (_result *ApplyScenarioResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ApplyScenarioShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Config) {
		request.ConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Config, dara.String("Config"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ConfigShrink) {
		query["Config"] = request.ConfigShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Scenario) {
		query["Scenario"] = request.Scenario
	}

	if !dara.IsNil(request.Sign) {
		query["Sign"] = request.Sign
	}

	if !dara.IsNil(request.SnDump) {
		query["SnDump"] = request.SnDump
	}

	if !dara.IsNil(request.SnForce) {
		query["SnForce"] = request.SnForce
	}

	if !dara.IsNil(request.SnStat) {
		query["SnStat"] = request.SnStat
	}

	if !dara.IsNil(request.SnTransfer) {
		query["SnTransfer"] = request.SnTransfer
	}

	if !dara.IsNil(request.UpdateOption) {
		query["UpdateOption"] = request.UpdateOption
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyScenario"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyScenarioResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI BindPrometheusGrafanaInstance is deprecated
//
// Summary:
//
// Binds a Grafana workspace to a Prometheus instance.
//
// @param request - BindPrometheusGrafanaInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindPrometheusGrafanaInstanceResponse
func (client *Client) BindPrometheusGrafanaInstanceWithContext(ctx context.Context, request *BindPrometheusGrafanaInstanceRequest, runtime *dara.RuntimeOptions) (_result *BindPrometheusGrafanaInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.GrafanaInstanceId) {
		query["GrafanaInstanceId"] = request.GrafanaInstanceId
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
		Action:      dara.String("BindPrometheusGrafanaInstance"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindPrometheusGrafanaInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Blocks alert notifications in a time period.
//
// @param request - BlockAlarmNotificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BlockAlarmNotificationResponse
func (client *Client) BlockAlarmNotificationWithContext(ctx context.Context, request *BlockAlarmNotificationRequest, runtime *dara.RuntimeOptions) (_result *BlockAlarmNotificationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlarmId) {
		query["AlarmId"] = request.AlarmId
	}

	if !dara.IsNil(request.HandlerId) {
		query["HandlerId"] = request.HandlerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BlockAlarmNotification"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BlockAlarmNotificationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the severity level of an alert.
//
// @param request - ChangeAlarmSeverityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeAlarmSeverityResponse
func (client *Client) ChangeAlarmSeverityWithContext(ctx context.Context, request *ChangeAlarmSeverityRequest, runtime *dara.RuntimeOptions) (_result *ChangeAlarmSeverityResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlarmId) {
		query["AlarmId"] = request.AlarmId
	}

	if !dara.IsNil(request.HandlerId) {
		query["HandlerId"] = request.HandlerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Severity) {
		query["Severity"] = request.Severity
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeAlarmSeverity"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeAlarmSeverityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves a resource to a specific resource group.
//
// @param request - ChangeResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithContext(ctx context.Context, request *ChangeResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
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
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether Application Real-Time Monitoring Service (ARMS) is available for commercial use in a region.
//
// Description:
//
// You can call this operation to check whether ARMS is available for commercial use in a region.
//
// @param request - CheckCommercialStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckCommercialStatusResponse
func (client *Client) CheckCommercialStatusWithContext(ctx context.Context, request *CheckCommercialStatusRequest, runtime *dara.RuntimeOptions) (_result *CheckCommercialStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckCommercialStatus"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckCommercialStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks the status of a service in the current cluster, such as whether the service is activated and whether the payment is overdue.
//
// @param request - CheckServiceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckServiceStatusResponse
func (client *Client) CheckServiceStatusWithContext(ctx context.Context, request *CheckServiceStatusRequest, runtime *dara.RuntimeOptions) (_result *CheckServiceStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SvcCode) {
		query["SvcCode"] = request.SvcCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckServiceStatus"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckServiceStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Claims an alert. This operation can be used together with escalation policies. When multiple handlers are involved in alert management, each handler can call this operation to claim alerts. After an alert is claimed, the alert enters the Being Processed state.
//
// @param request - ClaimAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClaimAlarmResponse
func (client *Client) ClaimAlarmWithContext(ctx context.Context, request *ClaimAlarmRequest, runtime *dara.RuntimeOptions) (_result *ClaimAlarmResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlarmId) {
		query["AlarmId"] = request.AlarmId
	}

	if !dara.IsNil(request.HandlerId) {
		query["HandlerId"] = request.HandlerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClaimAlarm"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClaimAlarmResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables an alert. Make sure that the alert is resolved before you disable the alert. If an alert is not resolved, new alerts can be triggered even after the alert is disabled.
//
// @param request - CloseAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseAlarmResponse
func (client *Client) CloseAlarmWithContext(ctx context.Context, request *CloseAlarmRequest, runtime *dara.RuntimeOptions) (_result *CloseAlarmResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlarmId) {
		query["AlarmId"] = request.AlarmId
	}

	if !dara.IsNil(request.HandlerId) {
		query["HandlerId"] = request.HandlerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Solution) {
		query["Solution"] = request.Solution
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloseAlarm"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloseAlarmResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Turns on or off the main switch of an ARMS agent, or queries the status of the main switch.
//
// Description:
//
// ***
//
// @param request - ConfigAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigAppResponse
func (client *Client) ConfigAppWithContext(ctx context.Context, request *ConfigAppRequest, runtime *dara.RuntimeOptions) (_result *ConfigAppResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppIds) {
		query["AppIds"] = request.AppIds
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
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
		Action:      dara.String("ConfigApp"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an alert contact.
//
// Description:
//
// This operation is no longer maintained. To create or modify an alert contact, call the CreateOrUpdateContact operation provided by the new version of the Alert Management module.
//
// @param request - CreateAlertContactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAlertContactResponse
func (client *Client) CreateAlertContactWithContext(ctx context.Context, request *CreateAlertContactRequest, runtime *dara.RuntimeOptions) (_result *CreateAlertContactResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactName) {
		query["ContactName"] = request.ContactName
	}

	if !dara.IsNil(request.DingRobotWebhookUrl) {
		query["DingRobotWebhookUrl"] = request.DingRobotWebhookUrl
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.PhoneNum) {
		query["PhoneNum"] = request.PhoneNum
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SystemNoc) {
		query["SystemNoc"] = request.SystemNoc
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAlertContact"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAlertContactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an alert contact group.
//
// Description:
//
// The current API operation is no longer maintained. Call the CreateOrUpdateContactGroup operation of the new Alert Management module to create or modify alert contact groups.
//
// @param request - CreateAlertContactGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAlertContactGroupResponse
func (client *Client) CreateAlertContactGroupWithContext(ctx context.Context, request *CreateAlertContactGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateAlertContactGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactGroupName) {
		query["ContactGroupName"] = request.ContactGroupName
	}

	if !dara.IsNil(request.ContactIds) {
		query["ContactIds"] = request.ContactIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAlertContactGroup"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAlertContactGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a dispatch policy.
//
// @param request - CreateDispatchRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDispatchRuleResponse
func (client *Client) CreateDispatchRuleWithContext(ctx context.Context, request *CreateDispatchRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateDispatchRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DispatchRule) {
		query["DispatchRule"] = request.DispatchRule
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDispatchRule"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDispatchRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom job for an environment.
//
// @param request - CreateEnvCustomJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEnvCustomJobResponse
func (client *Client) CreateEnvCustomJobWithContext(ctx context.Context, request *CreateEnvCustomJobRequest, runtime *dara.RuntimeOptions) (_result *CreateEnvCustomJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunLang) {
		query["AliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.CustomJobName) {
		query["CustomJobName"] = request.CustomJobName
	}

	if !dara.IsNil(request.EnvironmentId) {
		query["EnvironmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigYaml) {
		body["ConfigYaml"] = request.ConfigYaml
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEnvCustomJob"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEnvCustomJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a PodMonitor for an environment.
//
// @param request - CreateEnvPodMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEnvPodMonitorResponse
func (client *Client) CreateEnvPodMonitorWithContext(ctx context.Context, request *CreateEnvPodMonitorRequest, runtime *dara.RuntimeOptions) (_result *CreateEnvPodMonitorResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunLang) {
		query["AliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EnvironmentId) {
		query["EnvironmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigYaml) {
		body["ConfigYaml"] = request.ConfigYaml
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEnvPodMonitor"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEnvPodMonitorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a ServiceMonitor for an environment.
//
// @param request - CreateEnvServiceMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEnvServiceMonitorResponse
func (client *Client) CreateEnvServiceMonitorWithContext(ctx context.Context, request *CreateEnvServiceMonitorRequest, runtime *dara.RuntimeOptions) (_result *CreateEnvServiceMonitorResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunLang) {
		query["AliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EnvironmentId) {
		query["EnvironmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigYaml) {
		body["ConfigYaml"] = request.ConfigYaml
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEnvServiceMonitor"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEnvServiceMonitorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an environment instance.
//
// @param request - CreateEnvironmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEnvironmentResponse
func (client *Client) CreateEnvironmentWithContext(ctx context.Context, request *CreateEnvironmentRequest, runtime *dara.RuntimeOptions) (_result *CreateEnvironmentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunLang) {
		query["AliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.BindResourceId) {
		query["BindResourceId"] = request.BindResourceId
	}

	if !dara.IsNil(request.EnvironmentName) {
		query["EnvironmentName"] = request.EnvironmentName
	}

	if !dara.IsNil(request.EnvironmentSubType) {
		query["EnvironmentSubType"] = request.EnvironmentSubType
	}

	if !dara.IsNil(request.EnvironmentType) {
		query["EnvironmentType"] = request.EnvironmentType
	}

	if !dara.IsNil(request.FeePackage) {
		query["FeePackage"] = request.FeePackage
	}

	if !dara.IsNil(request.GrafanaWorkspaceId) {
		query["GrafanaWorkspaceId"] = request.GrafanaWorkspaceId
	}

	if !dara.IsNil(request.InitEnvironment) {
		query["InitEnvironment"] = request.InitEnvironment
	}

	if !dara.IsNil(request.ManagedType) {
		query["ManagedType"] = request.ManagedType
	}

	if !dara.IsNil(request.PrometheusInstanceId) {
		query["PrometheusInstanceId"] = request.PrometheusInstanceId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEnvironment"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEnvironmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a workspace in Managed Service for Grafana.
//
// Description:
//
// Before you call the operation, make sure that you have learned about the billing methods and [pricing](https://www.alibabacloud.com/help/zh/grafana/product-overview/billing-4?spm=a2c4g.11186623.0.0.14c2d253B3SDbt) of Managed Service for Grafana.
//
// >
//
//   - To create workspaces, you must complete real-name verification.
//
//   - Regular users can create workspaces only in Managed Service for Grafana Developer Edition, Pro Edition, and Advanced Edition. `These editions charge fees.`
//
//   - Internal users can create workspaces only in Managed Service for Grafana Beta Edition and Standard Edition. `These editions do not charge fees.`
//
// @param tmpReq - CreateGrafanaWorkspaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGrafanaWorkspaceResponse
func (client *Client) CreateGrafanaWorkspaceWithContext(ctx context.Context, tmpReq *CreateGrafanaWorkspaceRequest, runtime *dara.RuntimeOptions) (_result *CreateGrafanaWorkspaceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateGrafanaWorkspaceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountNumber) {
		query["AccountNumber"] = request.AccountNumber
	}

	if !dara.IsNil(request.AliyunLang) {
		query["AliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.CustomAccountNumber) {
		query["CustomAccountNumber"] = request.CustomAccountNumber
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.GrafanaVersion) {
		query["GrafanaVersion"] = request.GrafanaVersion
	}

	if !dara.IsNil(request.GrafanaWorkspaceEdition) {
		query["GrafanaWorkspaceEdition"] = request.GrafanaWorkspaceEdition
	}

	if !dara.IsNil(request.GrafanaWorkspaceName) {
		query["GrafanaWorkspaceName"] = request.GrafanaWorkspaceName
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.PricingCycle) {
		query["PricingCycle"] = request.PricingCycle
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGrafanaWorkspace"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGrafanaWorkspaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an alert integration.
//
// @param request - CreateIntegrationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIntegrationResponse
func (client *Client) CreateIntegrationWithContext(ctx context.Context, request *CreateIntegrationRequest, runtime *dara.RuntimeOptions) (_result *CreateIntegrationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoRecover) {
		body["AutoRecover"] = request.AutoRecover
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.IntegrationName) {
		body["IntegrationName"] = request.IntegrationName
	}

	if !dara.IsNil(request.IntegrationProductType) {
		body["IntegrationProductType"] = request.IntegrationProductType
	}

	if !dara.IsNil(request.RecoverTime) {
		body["RecoverTime"] = request.RecoverTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIntegration"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIntegrationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or modifies an alert rule.
//
// @param request - CreateOrUpdateAlertRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOrUpdateAlertRuleResponse
func (client *Client) CreateOrUpdateAlertRuleWithContext(ctx context.Context, request *CreateOrUpdateAlertRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateOrUpdateAlertRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AlertCheckType) {
		body["AlertCheckType"] = request.AlertCheckType
	}

	if !dara.IsNil(request.AlertGroup) {
		body["AlertGroup"] = request.AlertGroup
	}

	if !dara.IsNil(request.AlertId) {
		body["AlertId"] = request.AlertId
	}

	if !dara.IsNil(request.AlertName) {
		body["AlertName"] = request.AlertName
	}

	if !dara.IsNil(request.AlertPiplines) {
		body["AlertPiplines"] = request.AlertPiplines
	}

	if !dara.IsNil(request.AlertRuleContent) {
		body["AlertRuleContent"] = request.AlertRuleContent
	}

	if !dara.IsNil(request.AlertStatus) {
		body["AlertStatus"] = request.AlertStatus
	}

	if !dara.IsNil(request.AlertType) {
		body["AlertType"] = request.AlertType
	}

	if !dara.IsNil(request.Annotations) {
		body["Annotations"] = request.Annotations
	}

	if !dara.IsNil(request.AutoAddNewApplication) {
		body["AutoAddNewApplication"] = request.AutoAddNewApplication
	}

	if !dara.IsNil(request.AutoAddTargetConfig) {
		body["AutoAddTargetConfig"] = request.AutoAddTargetConfig
	}

	if !dara.IsNil(request.CheckCycle) {
		body["CheckCycle"] = request.CheckCycle
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DataConfig) {
		body["DataConfig"] = request.DataConfig
	}

	if !dara.IsNil(request.Duration) {
		body["Duration"] = request.Duration
	}

	if !dara.IsNil(request.Filters) {
		body["Filters"] = request.Filters
	}

	if !dara.IsNil(request.Labels) {
		body["Labels"] = request.Labels
	}

	if !dara.IsNil(request.Level) {
		body["Level"] = request.Level
	}

	if !dara.IsNil(request.MarkTags) {
		body["MarkTags"] = request.MarkTags
	}

	if !dara.IsNil(request.Message) {
		body["Message"] = request.Message
	}

	if !dara.IsNil(request.MetricsKey) {
		body["MetricsKey"] = request.MetricsKey
	}

	if !dara.IsNil(request.MetricsType) {
		body["MetricsType"] = request.MetricsType
	}

	if !dara.IsNil(request.Notice) {
		body["Notice"] = request.Notice
	}

	if !dara.IsNil(request.NotifyMode) {
		body["NotifyMode"] = request.NotifyMode
	}

	if !dara.IsNil(request.NotifyStrategy) {
		body["NotifyStrategy"] = request.NotifyStrategy
	}

	if !dara.IsNil(request.Pids) {
		body["Pids"] = request.Pids
	}

	if !dara.IsNil(request.Product) {
		body["Product"] = request.Product
	}

	if !dara.IsNil(request.PromQL) {
		body["PromQL"] = request.PromQL
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Tags) {
		body["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOrUpdateAlertRule"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOrUpdateAlertRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or modifies an alert contact.
//
// @param request - CreateOrUpdateContactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOrUpdateContactResponse
func (client *Client) CreateOrUpdateContactWithContext(ctx context.Context, request *CreateOrUpdateContactRequest, runtime *dara.RuntimeOptions) (_result *CreateOrUpdateContactResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DingRobotUrl) {
		query["DingRobotUrl"] = request.DingRobotUrl
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ContactId) {
		body["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.ContactName) {
		body["ContactName"] = request.ContactName
	}

	if !dara.IsNil(request.CorpUserId) {
		body["CorpUserId"] = request.CorpUserId
	}

	if !dara.IsNil(request.Email) {
		body["Email"] = request.Email
	}

	if !dara.IsNil(request.IsEmailVerify) {
		body["IsEmailVerify"] = request.IsEmailVerify
	}

	if !dara.IsNil(request.Phone) {
		body["Phone"] = request.Phone
	}

	if !dara.IsNil(request.ReissueSendNotice) {
		body["ReissueSendNotice"] = request.ReissueSendNotice
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOrUpdateContact"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOrUpdateContactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or modifies an alert contact group.
//
// @param request - CreateOrUpdateContactGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOrUpdateContactGroupResponse
func (client *Client) CreateOrUpdateContactGroupWithContext(ctx context.Context, request *CreateOrUpdateContactGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateOrUpdateContactGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ContactGroupId) {
		body["ContactGroupId"] = request.ContactGroupId
	}

	if !dara.IsNil(request.ContactGroupName) {
		body["ContactGroupName"] = request.ContactGroupName
	}

	if !dara.IsNil(request.ContactIds) {
		body["ContactIds"] = request.ContactIds
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOrUpdateContactGroup"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOrUpdateContactGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or modifies an EventBridge integration.
//
// @param request - CreateOrUpdateEventBridgeIntegrationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOrUpdateEventBridgeIntegrationResponse
func (client *Client) CreateOrUpdateEventBridgeIntegrationWithContext(ctx context.Context, request *CreateOrUpdateEventBridgeIntegrationRequest, runtime *dara.RuntimeOptions) (_result *CreateOrUpdateEventBridgeIntegrationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccessKey) {
		body["AccessKey"] = request.AccessKey
	}

	if !dara.IsNil(request.AccessSecret) {
		body["AccessSecret"] = request.AccessSecret
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Endpoint) {
		body["Endpoint"] = request.Endpoint
	}

	if !dara.IsNil(request.EventBusName) {
		body["EventBusName"] = request.EventBusName
	}

	if !dara.IsNil(request.EventBusRegionId) {
		body["EventBusRegionId"] = request.EventBusRegionId
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Source) {
		body["Source"] = request.Source
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOrUpdateEventBridgeIntegration"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOrUpdateEventBridgeIntegrationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or updates an IM chatbot.
//
// @param request - CreateOrUpdateIMRobotRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOrUpdateIMRobotResponse
func (client *Client) CreateOrUpdateIMRobotWithContext(ctx context.Context, request *CreateOrUpdateIMRobotRequest, runtime *dara.RuntimeOptions) (_result *CreateOrUpdateIMRobotResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CardTemplate) {
		body["CardTemplate"] = request.CardTemplate
	}

	if !dara.IsNil(request.DailyNoc) {
		body["DailyNoc"] = request.DailyNoc
	}

	if !dara.IsNil(request.DailyNocTime) {
		body["DailyNocTime"] = request.DailyNocTime
	}

	if !dara.IsNil(request.DingSignKey) {
		body["DingSignKey"] = request.DingSignKey
	}

	if !dara.IsNil(request.EnableOutgoing) {
		body["EnableOutgoing"] = request.EnableOutgoing
	}

	if !dara.IsNil(request.RobotAddress) {
		body["RobotAddress"] = request.RobotAddress
	}

	if !dara.IsNil(request.RobotId) {
		body["RobotId"] = request.RobotId
	}

	if !dara.IsNil(request.RobotName) {
		body["RobotName"] = request.RobotName
	}

	if !dara.IsNil(request.Token) {
		body["Token"] = request.Token
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOrUpdateIMRobot"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOrUpdateIMRobotResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or modifies a notification policy.
//
// @param request - CreateOrUpdateNotificationPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOrUpdateNotificationPolicyResponse
func (client *Client) CreateOrUpdateNotificationPolicyWithContext(ctx context.Context, request *CreateOrUpdateNotificationPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateOrUpdateNotificationPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DirectedMode) {
		body["DirectedMode"] = request.DirectedMode
	}

	if !dara.IsNil(request.EscalationPolicyId) {
		body["EscalationPolicyId"] = request.EscalationPolicyId
	}

	if !dara.IsNil(request.GroupRule) {
		body["GroupRule"] = request.GroupRule
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.IntegrationId) {
		body["IntegrationId"] = request.IntegrationId
	}

	if !dara.IsNil(request.MatchingRules) {
		body["MatchingRules"] = request.MatchingRules
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NotifyRule) {
		body["NotifyRule"] = request.NotifyRule
	}

	if !dara.IsNil(request.NotifyTemplate) {
		body["NotifyTemplate"] = request.NotifyTemplate
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Repeat) {
		body["Repeat"] = request.Repeat
	}

	if !dara.IsNil(request.RepeatInterval) {
		body["RepeatInterval"] = request.RepeatInterval
	}

	if !dara.IsNil(request.SendRecoverMessage) {
		body["SendRecoverMessage"] = request.SendRecoverMessage
	}

	if !dara.IsNil(request.State) {
		body["State"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOrUpdateNotificationPolicy"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOrUpdateNotificationPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or modifies a silence policy.
//
// @param request - CreateOrUpdateSilencePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOrUpdateSilencePolicyResponse
func (client *Client) CreateOrUpdateSilencePolicyWithContext(ctx context.Context, request *CreateOrUpdateSilencePolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateOrUpdateSilencePolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EffectiveTimeType) {
		query["EffectiveTimeType"] = request.EffectiveTimeType
	}

	if !dara.IsNil(request.TimePeriod) {
		query["TimePeriod"] = request.TimePeriod
	}

	if !dara.IsNil(request.TimeSlots) {
		query["TimeSlots"] = request.TimeSlots
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.MatchingRules) {
		body["MatchingRules"] = request.MatchingRules
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.State) {
		body["State"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOrUpdateSilencePolicy"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOrUpdateSilencePolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or modifies a webhook alert contact.
//
// @param request - CreateOrUpdateWebhookContactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOrUpdateWebhookContactResponse
func (client *Client) CreateOrUpdateWebhookContactWithContext(ctx context.Context, request *CreateOrUpdateWebhookContactRequest, runtime *dara.RuntimeOptions) (_result *CreateOrUpdateWebhookContactResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BizHeaders) {
		body["BizHeaders"] = request.BizHeaders
	}

	if !dara.IsNil(request.BizParams) {
		body["BizParams"] = request.BizParams
	}

	if !dara.IsNil(request.Body) {
		body["Body"] = request.Body
	}

	if !dara.IsNil(request.Method) {
		body["Method"] = request.Method
	}

	if !dara.IsNil(request.RecoverBody) {
		body["RecoverBody"] = request.RecoverBody
	}

	if !dara.IsNil(request.Url) {
		body["Url"] = request.Url
	}

	if !dara.IsNil(request.WebhookId) {
		body["WebhookId"] = request.WebhookId
	}

	if !dara.IsNil(request.WebhookName) {
		body["WebhookName"] = request.WebhookName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOrUpdateWebhookContact"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOrUpdateWebhookContactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an alert rule.
//
// @param request - CreatePrometheusAlertRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePrometheusAlertRuleResponse
func (client *Client) CreatePrometheusAlertRuleWithContext(ctx context.Context, request *CreatePrometheusAlertRuleRequest, runtime *dara.RuntimeOptions) (_result *CreatePrometheusAlertRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertName) {
		query["AlertName"] = request.AlertName
	}

	if !dara.IsNil(request.Annotations) {
		query["Annotations"] = request.Annotations
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DispatchRuleId) {
		query["DispatchRuleId"] = request.DispatchRuleId
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.Expression) {
		query["Expression"] = request.Expression
	}

	if !dara.IsNil(request.Labels) {
		query["Labels"] = request.Labels
	}

	if !dara.IsNil(request.Message) {
		query["Message"] = request.Message
	}

	if !dara.IsNil(request.NotifyType) {
		query["NotifyType"] = request.NotifyType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePrometheusAlertRule"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePrometheusAlertRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Prometheus instance.
//
// @param request - CreatePrometheusInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePrometheusInstanceResponse
func (client *Client) CreatePrometheusInstanceWithContext(ctx context.Context, request *CreatePrometheusInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreatePrometheusInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllSubClustersSuccess) {
		query["AllSubClustersSuccess"] = request.AllSubClustersSuccess
	}

	if !dara.IsNil(request.ArchiveDuration) {
		query["ArchiveDuration"] = request.ArchiveDuration
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.ClusterType) {
		query["ClusterType"] = request.ClusterType
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.GrafanaInstanceId) {
		query["GrafanaInstanceId"] = request.GrafanaInstanceId
	}

	if !dara.IsNil(request.PaymentType) {
		query["PaymentType"] = request.PaymentType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.SubClustersJson) {
		query["SubClustersJson"] = request.SubClustersJson
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
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
		Action:      dara.String("CreatePrometheusInstance"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePrometheusInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI CreatePrometheusMonitoring is deprecated
//
// Summary:
//
// Creates a monitoring configuration for a Prometheus instance.
//
// @param request - CreatePrometheusMonitoringRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePrometheusMonitoringResponse
func (client *Client) CreatePrometheusMonitoringWithContext(ctx context.Context, request *CreatePrometheusMonitoringRequest, runtime *dara.RuntimeOptions) (_result *CreatePrometheusMonitoringResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigYaml) {
		body["ConfigYaml"] = request.ConfigYaml
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePrometheusMonitoring"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePrometheusMonitoringResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Browser Monitoring task for an application.
//
// @param request - CreateRetcodeAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRetcodeAppResponse
func (client *Client) CreateRetcodeAppWithContext(ctx context.Context, request *CreateRetcodeAppRequest, runtime *dara.RuntimeOptions) (_result *CreateRetcodeAppResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.RetcodeAppName) {
		query["RetcodeAppName"] = request.RetcodeAppName
	}

	if !dara.IsNil(request.RetcodeAppType) {
		query["RetcodeAppType"] = request.RetcodeAppType
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRetcodeApp"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRetcodeAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a Real User Monitoring (RUM) application.
//
// @param tmpReq - CreateRumAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRumAppResponse
func (client *Client) CreateRumAppWithContext(ctx context.Context, tmpReq *CreateRumAppRequest, runtime *dara.RuntimeOptions) (_result *CreateRumAppResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateRumAppShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppGroup) {
		query["AppGroup"] = request.AppGroup
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.NickName) {
		query["NickName"] = request.NickName
	}

	if !dara.IsNil(request.PackageName) {
		query["PackageName"] = request.PackageName
	}

	if !dara.IsNil(request.RealRegionId) {
		query["RealRegionId"] = request.RealRegionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SiteType) {
		query["SiteType"] = request.SiteType
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRumApp"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRumAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a file upload URL to upload SourceMap files, symbol table files, or dSYM files.
//
// Description:
//
// This operation returns a URL. You can upload files to the URL. For more information, see [Upload local files with signed URLs](https://help.aliyun.com/document_detail/2579659.html).
//
// @param request - CreateRumUploadFileUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRumUploadFileUrlResponse
func (client *Client) CreateRumUploadFileUrlWithContext(ctx context.Context, request *CreateRumUploadFileUrlRequest, runtime *dara.RuntimeOptions) (_result *CreateRumUploadFileUrlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ContentType) {
		query["ContentType"] = request.ContentType
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.Pid) {
		query["Pid"] = request.Pid
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SourcemapType) {
		query["SourcemapType"] = request.SourcemapType
	}

	if !dara.IsNil(request.Uuid) {
		query["Uuid"] = request.Uuid
	}

	if !dara.IsNil(request.VersionId) {
		query["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRumUploadFileUrl"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRumUploadFileUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or updates a synthetic monitoring task.
//
// @param tmpReq - CreateSyntheticTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSyntheticTaskResponse
func (client *Client) CreateSyntheticTaskWithContext(ctx context.Context, tmpReq *CreateSyntheticTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateSyntheticTaskResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateSyntheticTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CommonParam) {
		request.CommonParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CommonParam, dara.String("CommonParam"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Download) {
		request.DownloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Download, dara.String("Download"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ExtendInterval) {
		request.ExtendIntervalShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExtendInterval, dara.String("ExtendInterval"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.MonitorList) {
		request.MonitorListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MonitorList, dara.String("MonitorList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Navigation) {
		request.NavigationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Navigation, dara.String("Navigation"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Net) {
		request.NetShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Net, dara.String("Net"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Protocol) {
		request.ProtocolShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Protocol, dara.String("Protocol"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CommonParamShrink) {
		query["CommonParam"] = request.CommonParamShrink
	}

	if !dara.IsNil(request.DownloadShrink) {
		query["Download"] = request.DownloadShrink
	}

	if !dara.IsNil(request.ExtendIntervalShrink) {
		query["ExtendInterval"] = request.ExtendIntervalShrink
	}

	if !dara.IsNil(request.IntervalTime) {
		query["IntervalTime"] = request.IntervalTime
	}

	if !dara.IsNil(request.IntervalType) {
		query["IntervalType"] = request.IntervalType
	}

	if !dara.IsNil(request.IpType) {
		query["IpType"] = request.IpType
	}

	if !dara.IsNil(request.MonitorListShrink) {
		query["MonitorList"] = request.MonitorListShrink
	}

	if !dara.IsNil(request.NavigationShrink) {
		query["Navigation"] = request.NavigationShrink
	}

	if !dara.IsNil(request.NetShrink) {
		query["Net"] = request.NetShrink
	}

	if !dara.IsNil(request.ProtocolShrink) {
		query["Protocol"] = request.ProtocolShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	if !dara.IsNil(request.UpdateTask) {
		query["UpdateTask"] = request.UpdateTask
	}

	if !dara.IsNil(request.Url) {
		query["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSyntheticTask"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSyntheticTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a scheduled synthetic test task.
//
// @param tmpReq - CreateTimingSyntheticTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTimingSyntheticTaskResponse
func (client *Client) CreateTimingSyntheticTaskWithContext(ctx context.Context, tmpReq *CreateTimingSyntheticTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateTimingSyntheticTaskResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateTimingSyntheticTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AvailableAssertions) {
		request.AvailableAssertionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AvailableAssertions, dara.String("AvailableAssertions"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CommonSetting) {
		request.CommonSettingShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CommonSetting, dara.String("CommonSetting"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CustomPeriod) {
		request.CustomPeriodShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomPeriod, dara.String("CustomPeriod"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.MonitorConf) {
		request.MonitorConfShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MonitorConf, dara.String("MonitorConf"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Monitors) {
		request.MonitorsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Monitors, dara.String("Monitors"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AvailableAssertionsShrink) {
		query["AvailableAssertions"] = request.AvailableAssertionsShrink
	}

	if !dara.IsNil(request.CommonSettingShrink) {
		query["CommonSetting"] = request.CommonSettingShrink
	}

	if !dara.IsNil(request.CustomPeriodShrink) {
		query["CustomPeriod"] = request.CustomPeriodShrink
	}

	if !dara.IsNil(request.Frequency) {
		query["Frequency"] = request.Frequency
	}

	if !dara.IsNil(request.MonitorCategory) {
		query["MonitorCategory"] = request.MonitorCategory
	}

	if !dara.IsNil(request.MonitorConfShrink) {
		query["MonitorConf"] = request.MonitorConfShrink
	}

	if !dara.IsNil(request.MonitorsShrink) {
		query["Monitors"] = request.MonitorsShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTimingSyntheticTask"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTimingSyntheticTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a contact for webhook alerts.
//
// @param request - CreateWebhookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWebhookResponse
func (client *Client) CreateWebhookWithContext(ctx context.Context, request *CreateWebhookRequest, runtime *dara.RuntimeOptions) (_result *CreateWebhookResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Body) {
		query["Body"] = request.Body
	}

	if !dara.IsNil(request.ContactName) {
		query["ContactName"] = request.ContactName
	}

	if !dara.IsNil(request.HttpHeaders) {
		query["HttpHeaders"] = request.HttpHeaders
	}

	if !dara.IsNil(request.HttpParams) {
		query["HttpParams"] = request.HttpParams
	}

	if !dara.IsNil(request.Method) {
		query["Method"] = request.Method
	}

	if !dara.IsNil(request.RecoverBody) {
		query["RecoverBody"] = request.RecoverBody
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Url) {
		query["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWebhook"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWebhookResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables an authentication token.
//
// @param request - DelAuthTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DelAuthTokenResponse
func (client *Client) DelAuthTokenWithContext(ctx context.Context, request *DelAuthTokenRequest, runtime *dara.RuntimeOptions) (_result *DelAuthTokenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("DelAuthToken"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DelAuthTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete AddonRelease data by AddonRelease name.
//
// @param request - DeleteAddonReleaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAddonReleaseResponse
func (client *Client) DeleteAddonReleaseWithContext(ctx context.Context, request *DeleteAddonReleaseRequest, runtime *dara.RuntimeOptions) (_result *DeleteAddonReleaseResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddonName) {
		query["AddonName"] = request.AddonName
	}

	if !dara.IsNil(request.EnvironmentId) {
		query["EnvironmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReleaseName) {
		query["ReleaseName"] = request.ReleaseName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAddonRelease"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAddonReleaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an DeleteAlertContact contact.
//
// Description:
//
// *******
//
// @param request - DeleteAlertContactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAlertContactResponse
func (client *Client) DeleteAlertContactWithContext(ctx context.Context, request *DeleteAlertContactRequest, runtime *dara.RuntimeOptions) (_result *DeleteAlertContactResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactId) {
		query["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAlertContact"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAlertContactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an DeleteAlertContactGroup contact group.
//
// @param request - DeleteAlertContactGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAlertContactGroupResponse
func (client *Client) DeleteAlertContactGroupWithContext(ctx context.Context, request *DeleteAlertContactGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteAlertContactGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactGroupId) {
		query["ContactGroupId"] = request.ContactGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAlertContactGroup"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAlertContactGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an alert rule.
//
// @param request - DeleteAlertRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAlertRuleResponse
func (client *Client) DeleteAlertRuleWithContext(ctx context.Context, request *DeleteAlertRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteAlertRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertId) {
		query["AlertId"] = request.AlertId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAlertRule"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAlertRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes alert rules.
//
// Description:
//
// The current operation is no longer maintained. Call the DeleteAlertRule operation of Alert Management (New) to delete alert rules.
//
// @param request - DeleteAlertRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAlertRulesResponse
func (client *Client) DeleteAlertRulesWithContext(ctx context.Context, request *DeleteAlertRulesRequest, runtime *dara.RuntimeOptions) (_result *DeleteAlertRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertIds) {
		query["AlertIds"] = request.AlertIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAlertRules"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAlertRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes multiple applications at a time based on the process IDs (PIDs).
//
// @param request - DeleteAppListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppListResponse
func (client *Client) DeleteAppListWithContext(ctx context.Context, request *DeleteAppListRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Pids) {
		query["Pids"] = request.Pids
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAppList"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAppListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DeleteCmsExporter is deprecated, please use ARMS::2019-08-08::DeleteAddonRelease instead.
//
// Summary:
//
// Releases a Prometheus instance for Alibaba Cloud services.
//
// @param request - DeleteCmsExporterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCmsExporterResponse
func (client *Client) DeleteCmsExporterWithContext(ctx context.Context, request *DeleteCmsExporterRequest, runtime *dara.RuntimeOptions) (_result *DeleteCmsExporterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("DeleteCmsExporter"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCmsExporterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes alert contacts.
//
// @param request - DeleteContactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteContactResponse
func (client *Client) DeleteContactWithContext(ctx context.Context, request *DeleteContactRequest, runtime *dara.RuntimeOptions) (_result *DeleteContactResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactId) {
		query["ContactId"] = request.ContactId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteContact"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteContactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an alert contact group.
//
// @param request - DeleteContactGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteContactGroupResponse
func (client *Client) DeleteContactGroupWithContext(ctx context.Context, request *DeleteContactGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteContactGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactGroupId) {
		query["ContactGroupId"] = request.ContactGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteContactGroup"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteContactGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the dispatch policy of a specified ID.
//
// @param request - DeleteDispatchRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDispatchRuleResponse
func (client *Client) DeleteDispatchRuleWithContext(ctx context.Context, request *DeleteDispatchRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteDispatchRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDispatchRule"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDispatchRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom job for an environment.
//
// @param request - DeleteEnvCustomJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEnvCustomJobResponse
func (client *Client) DeleteEnvCustomJobWithContext(ctx context.Context, request *DeleteEnvCustomJobRequest, runtime *dara.RuntimeOptions) (_result *DeleteEnvCustomJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomJobName) {
		query["CustomJobName"] = request.CustomJobName
	}

	if !dara.IsNil(request.EnvironmentId) {
		query["EnvironmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEnvCustomJob"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEnvCustomJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the PodMonitor of an environment.
//
// @param request - DeleteEnvPodMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEnvPodMonitorResponse
func (client *Client) DeleteEnvPodMonitorWithContext(ctx context.Context, request *DeleteEnvPodMonitorRequest, runtime *dara.RuntimeOptions) (_result *DeleteEnvPodMonitorResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnvironmentId) {
		query["EnvironmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PodMonitorName) {
		query["PodMonitorName"] = request.PodMonitorName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEnvPodMonitor"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEnvPodMonitorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the ServiceMonitor of an environment.
//
// @param request - DeleteEnvServiceMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEnvServiceMonitorResponse
func (client *Client) DeleteEnvServiceMonitorWithContext(ctx context.Context, request *DeleteEnvServiceMonitorRequest, runtime *dara.RuntimeOptions) (_result *DeleteEnvServiceMonitorResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnvironmentId) {
		query["EnvironmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceMonitorName) {
		query["ServiceMonitorName"] = request.ServiceMonitorName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEnvServiceMonitor"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEnvServiceMonitorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an environment instance.
//
// @param request - DeleteEnvironmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEnvironmentResponse
func (client *Client) DeleteEnvironmentWithContext(ctx context.Context, request *DeleteEnvironmentRequest, runtime *dara.RuntimeOptions) (_result *DeleteEnvironmentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeletePromInstance) {
		query["DeletePromInstance"] = request.DeletePromInstance
	}

	if !dara.IsNil(request.EnvironmentId) {
		query["EnvironmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEnvironment"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEnvironmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a feature.
//
// @param request - DeleteEnvironmentFeatureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEnvironmentFeatureResponse
func (client *Client) DeleteEnvironmentFeatureWithContext(ctx context.Context, request *DeleteEnvironmentFeatureRequest, runtime *dara.RuntimeOptions) (_result *DeleteEnvironmentFeatureResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnvironmentId) {
		query["EnvironmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.FeatureName) {
		query["FeatureName"] = request.FeatureName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEnvironmentFeature"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEnvironmentFeatureResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an EventBridge integration.
//
// @param request - DeleteEventBridgeIntegrationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEventBridgeIntegrationResponse
func (client *Client) DeleteEventBridgeIntegrationWithContext(ctx context.Context, request *DeleteEventBridgeIntegrationRequest, runtime *dara.RuntimeOptions) (_result *DeleteEventBridgeIntegrationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEventBridgeIntegration"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEventBridgeIntegrationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DeleteGrafanaResource is deprecated
//
// Summary:
//
// Deletes Grafana dashboard resources from a Managed Service for Prometheus instance.
//
// @param request - DeleteGrafanaResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGrafanaResourceResponse
func (client *Client) DeleteGrafanaResourceWithContext(ctx context.Context, request *DeleteGrafanaResourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteGrafanaResourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ClusterName) {
		body["ClusterName"] = request.ClusterName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGrafanaResource"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGrafanaResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Managed Service for Prometheus workspace.
//
// Description:
//
//	  You can delete workspaces only in Managed Service for Prometheus Beta Edition, which is `free of charge`.
//
//		- You cannot delete workspaces in Managed Service for Prometheus Developer Edition, Pro Edition, and Advanced Edition. You can go to the [User Center](https://usercenter2.aliyun.com/refund/refund) to unsubscribe from workspaces.
//
// @param request - DeleteGrafanaWorkspaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGrafanaWorkspaceResponse
func (client *Client) DeleteGrafanaWorkspaceWithContext(ctx context.Context, request *DeleteGrafanaWorkspaceRequest, runtime *dara.RuntimeOptions) (_result *DeleteGrafanaWorkspaceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GrafanaWorkspaceId) {
		query["GrafanaWorkspaceId"] = request.GrafanaWorkspaceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGrafanaWorkspace"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGrafanaWorkspaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an instant messaging (IM) chatbot.
//
// @param request - DeleteIMRobotRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIMRobotResponse
func (client *Client) DeleteIMRobotWithContext(ctx context.Context, request *DeleteIMRobotRequest, runtime *dara.RuntimeOptions) (_result *DeleteIMRobotResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RobotId) {
		query["RobotId"] = request.RobotId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteIMRobot"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIMRobotResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DeleteIntegration is deprecated, please use ARMS::2019-08-08::DeleteAddonRelease instead.
//
// Summary:
//
// Deletes collection rules from an integration.
//
// @param request - DeleteIntegrationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIntegrationResponse
func (client *Client) DeleteIntegrationWithContext(ctx context.Context, request *DeleteIntegrationRequest, runtime *dara.RuntimeOptions) (_result *DeleteIntegrationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Integration) {
		query["Integration"] = request.Integration
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteIntegration"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIntegrationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an alert integration.
//
// @param request - DeleteIntegrationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIntegrationsResponse
func (client *Client) DeleteIntegrationsWithContext(ctx context.Context, request *DeleteIntegrationsRequest, runtime *dara.RuntimeOptions) (_result *DeleteIntegrationsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteIntegrations"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIntegrationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a notification policy based on its ID.
//
// @param request - DeleteNotificationPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNotificationPolicyResponse
func (client *Client) DeleteNotificationPolicyWithContext(ctx context.Context, request *DeleteNotificationPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteNotificationPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNotificationPolicy"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNotificationPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an alert rule of Prometheus Service.
//
// @param request - DeletePrometheusAlertRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePrometheusAlertRuleResponse
func (client *Client) DeletePrometheusAlertRuleWithContext(ctx context.Context, request *DeletePrometheusAlertRuleRequest, runtime *dara.RuntimeOptions) (_result *DeletePrometheusAlertRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertId) {
		query["AlertId"] = request.AlertId
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePrometheusAlertRule"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePrometheusAlertRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DeletePrometheusGlobalView is deprecated
//
// Summary:
//
// Deletes a global aggregation instance from Prometheus Service.
//
// @param request - DeletePrometheusGlobalViewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePrometheusGlobalViewResponse
func (client *Client) DeletePrometheusGlobalViewWithContext(ctx context.Context, request *DeletePrometheusGlobalViewRequest, runtime *dara.RuntimeOptions) (_result *DeletePrometheusGlobalViewResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GlobalViewClusterId) {
		query["GlobalViewClusterId"] = request.GlobalViewClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePrometheusGlobalView"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePrometheusGlobalViewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DeletePrometheusIntegration is deprecated
//
// Summary:
//
// Deletes an exporter from a Prometheus instance for Container Service or a Prometheus instance for ECS.
//
// @param request - DeletePrometheusIntegrationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePrometheusIntegrationResponse
func (client *Client) DeletePrometheusIntegrationWithContext(ctx context.Context, request *DeletePrometheusIntegrationRequest, runtime *dara.RuntimeOptions) (_result *DeletePrometheusIntegrationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IntegrationType) {
		query["IntegrationType"] = request.IntegrationType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePrometheusIntegration"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePrometheusIntegrationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DeletePrometheusMonitoring is deprecated
//
// Summary:
//
// Deletes the monitoring configuration of a Prometheus instance.
//
// @param request - DeletePrometheusMonitoringRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePrometheusMonitoringResponse
func (client *Client) DeletePrometheusMonitoringWithContext(ctx context.Context, request *DeletePrometheusMonitoringRequest, runtime *dara.RuntimeOptions) (_result *DeletePrometheusMonitoringResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.MonitoringName) {
		query["MonitoringName"] = request.MonitoringName
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
		Action:      dara.String("DeletePrometheusMonitoring"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePrometheusMonitoringResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Browser Monitoring task.
//
// @param request - DeleteRetcodeAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRetcodeAppResponse
func (client *Client) DeleteRetcodeAppWithContext(ctx context.Context, request *DeleteRetcodeAppRequest, runtime *dara.RuntimeOptions) (_result *DeleteRetcodeAppResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Pid) {
		query["Pid"] = request.Pid
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRetcodeApp"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRetcodeAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a RUM application.
//
// @param request - DeleteRumAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRumAppResponse
func (client *Client) DeleteRumAppWithContext(ctx context.Context, request *DeleteRumAppRequest, runtime *dara.RuntimeOptions) (_result *DeleteRumAppResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppGroup) {
		query["AppGroup"] = request.AppGroup
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.RealRegionId) {
		query["RealRegionId"] = request.RealRegionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRumApp"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRumAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a file such as a symbol table or SourceMap.
//
// Description:
//
// Real User Monitoring (RUM) is available only in the China (Hangzhou), Singapore, and US (Silicon Valley) regions. Select the correct endpoint.
//
// @param request - DeleteRumUploadFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRumUploadFileResponse
func (client *Client) DeleteRumUploadFileWithContext(ctx context.Context, request *DeleteRumUploadFileRequest, runtime *dara.RuntimeOptions) (_result *DeleteRumUploadFileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BatchItems) {
		query["BatchItems"] = request.BatchItems
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.Pid) {
		query["Pid"] = request.Pid
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Uuid) {
		query["Uuid"] = request.Uuid
	}

	if !dara.IsNil(request.VersionId) {
		query["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRumUploadFile"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRumUploadFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteScenarioRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteScenarioResponse
func (client *Client) DeleteScenarioWithContext(ctx context.Context, request *DeleteScenarioRequest, runtime *dara.RuntimeOptions) (_result *DeleteScenarioResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ScenarioId) {
		query["ScenarioId"] = request.ScenarioId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteScenario"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteScenarioResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The ID of the silence policy.
//
// @param request - DeleteSilencePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSilencePolicyResponse
func (client *Client) DeleteSilencePolicyWithContext(ctx context.Context, request *DeleteSilencePolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteSilencePolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSilencePolicy"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSilencePolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the SourceMap files uploaded in Browser Monitoring.
//
// @param tmpReq - DeleteSourceMapRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSourceMapResponse
func (client *Client) DeleteSourceMapWithContext(ctx context.Context, tmpReq *DeleteSourceMapRequest, runtime *dara.RuntimeOptions) (_result *DeleteSourceMapResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DeleteSourceMapShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FidList) {
		request.FidListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FidList, dara.String("FidList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.FidListShrink) {
		query["FidList"] = request.FidListShrink
	}

	if !dara.IsNil(request.Pid) {
		query["Pid"] = request.Pid
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSourceMap"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSourceMapResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes scheduled synthetic monitoring tasks.
//
// @param request - DeleteSyntheticTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSyntheticTaskResponse
func (client *Client) DeleteSyntheticTaskWithContext(ctx context.Context, request *DeleteSyntheticTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteSyntheticTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TaskIds) {
		query["TaskIds"] = request.TaskIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSyntheticTask"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSyntheticTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a scheduled synthetic monitoring task.
//
// @param request - DeleteTimingSyntheticTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTimingSyntheticTaskResponse
func (client *Client) DeleteTimingSyntheticTaskWithContext(ctx context.Context, request *DeleteTimingSyntheticTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteTimingSyntheticTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTimingSyntheticTask"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTimingSyntheticTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an application based on a specified process identifier (PID) and application type.
//
// @param tmpReq - DeleteTraceAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTraceAppResponse
func (client *Client) DeleteTraceAppWithContext(ctx context.Context, tmpReq *DeleteTraceAppRequest, runtime *dara.RuntimeOptions) (_result *DeleteTraceAppResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DeleteTraceAppShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeleteReason) {
		request.DeleteReasonShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeleteReason, dara.String("DeleteReason"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DeleteReasonShrink) {
		query["DeleteReason"] = request.DeleteReasonShrink
	}

	if !dara.IsNil(request.Pid) {
		query["Pid"] = request.Pid
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
		Action:      dara.String("DeleteTraceApp"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTraceAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a webhook alert contact.
//
// @param request - DeleteWebhookContactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWebhookContactResponse
func (client *Client) DeleteWebhookContactWithContext(ctx context.Context, request *DeleteWebhookContactRequest, runtime *dara.RuntimeOptions) (_result *DeleteWebhookContactResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WebhookId) {
		query["WebhookId"] = request.WebhookId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWebhookContact"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWebhookContactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the metric details of a component.
//
// @param request - DescribeAddonMetricsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAddonMetricsResponse
func (client *Client) DescribeAddonMetricsWithContext(ctx context.Context, request *DescribeAddonMetricsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAddonMetricsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddonVersion) {
		query["AddonVersion"] = request.AddonVersion
	}

	if !dara.IsNil(request.AliyunLang) {
		query["AliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.EnvironmentType) {
		query["EnvironmentType"] = request.EnvironmentType
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAddonMetrics"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAddonMetricsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the release information of an add-on by name.
//
// @param request - DescribeAddonReleaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAddonReleaseResponse
func (client *Client) DescribeAddonReleaseWithContext(ctx context.Context, request *DescribeAddonReleaseRequest, runtime *dara.RuntimeOptions) (_result *DescribeAddonReleaseResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnvironmentId) {
		query["EnvironmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReleaseName) {
		query["ReleaseName"] = request.ReleaseName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAddonRelease"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAddonReleaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries an alert contact group.
//
// @param request - DescribeContactGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeContactGroupsResponse
func (client *Client) DescribeContactGroupsWithContext(ctx context.Context, request *DescribeContactGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeContactGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactGroupName) {
		query["ContactGroupName"] = request.ContactGroupName
	}

	if !dara.IsNil(request.GroupIds) {
		query["GroupIds"] = request.GroupIds
	}

	if !dara.IsNil(request.IsDetail) {
		query["IsDetail"] = request.IsDetail
	}

	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Size) {
		query["Size"] = request.Size
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeContactGroups"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeContactGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries alert contacts.
//
// @param request - DescribeContactsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeContactsResponse
func (client *Client) DescribeContactsWithContext(ctx context.Context, request *DescribeContactsRequest, runtime *dara.RuntimeOptions) (_result *DescribeContactsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactIds) {
		query["ContactIds"] = request.ContactIds
	}

	if !dara.IsNil(request.ContactName) {
		query["ContactName"] = request.ContactName
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
	}

	if !dara.IsNil(request.Phone) {
		query["Phone"] = request.Phone
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Size) {
		query["Size"] = request.Size
	}

	if !dara.IsNil(request.Verbose) {
		query["Verbose"] = request.Verbose
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeContacts"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeContactsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a dispatch policy.
//
// @param request - DescribeDispatchRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDispatchRuleResponse
func (client *Client) DescribeDispatchRuleWithContext(ctx context.Context, request *DescribeDispatchRuleRequest, runtime *dara.RuntimeOptions) (_result *DescribeDispatchRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDispatchRule"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDispatchRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a custom job for an environment.
//
// @param request - DescribeEnvCustomJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEnvCustomJobResponse
func (client *Client) DescribeEnvCustomJobWithContext(ctx context.Context, request *DescribeEnvCustomJobRequest, runtime *dara.RuntimeOptions) (_result *DescribeEnvCustomJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomJobName) {
		query["CustomJobName"] = request.CustomJobName
	}

	if !dara.IsNil(request.EncryptYaml) {
		query["EncryptYaml"] = request.EncryptYaml
	}

	if !dara.IsNil(request.EnvironmentId) {
		query["EnvironmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEnvCustomJob"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEnvCustomJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询指定环境实例的废弃指标列表
//
// @param request - DescribeEnvDropMetricsRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEnvDropMetricsRuleResponse
func (client *Client) DescribeEnvDropMetricsRuleWithContext(ctx context.Context, request *DescribeEnvDropMetricsRuleRequest, runtime *dara.RuntimeOptions) (_result *DescribeEnvDropMetricsRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnvironmentId) {
		query["EnvironmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEnvDropMetricsRule"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEnvDropMetricsRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the PodMonitor details of an environment.
//
// @param request - DescribeEnvPodMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEnvPodMonitorResponse
func (client *Client) DescribeEnvPodMonitorWithContext(ctx context.Context, request *DescribeEnvPodMonitorRequest, runtime *dara.RuntimeOptions) (_result *DescribeEnvPodMonitorResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnvironmentId) {
		query["EnvironmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PodMonitorName) {
		query["PodMonitorName"] = request.PodMonitorName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEnvPodMonitor"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEnvPodMonitorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the ServiceMonitor details of an environment.
//
// @param request - DescribeEnvServiceMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEnvServiceMonitorResponse
func (client *Client) DescribeEnvServiceMonitorWithContext(ctx context.Context, request *DescribeEnvServiceMonitorRequest, runtime *dara.RuntimeOptions) (_result *DescribeEnvServiceMonitorResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnvironmentId) {
		query["EnvironmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceMonitorName) {
		query["ServiceMonitorName"] = request.ServiceMonitorName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEnvServiceMonitor"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEnvServiceMonitorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an environment.
//
// @param request - DescribeEnvironmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEnvironmentResponse
func (client *Client) DescribeEnvironmentWithContext(ctx context.Context, request *DescribeEnvironmentRequest, runtime *dara.RuntimeOptions) (_result *DescribeEnvironmentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnvironmentId) {
		query["EnvironmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEnvironment"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEnvironmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a feature.
//
// @param request - DescribeEnvironmentFeatureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEnvironmentFeatureResponse
func (client *Client) DescribeEnvironmentFeatureWithContext(ctx context.Context, request *DescribeEnvironmentFeatureRequest, runtime *dara.RuntimeOptions) (_result *DescribeEnvironmentFeatureResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunLang) {
		query["AliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.EnvironmentId) {
		query["EnvironmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.FeatureName) {
		query["FeatureName"] = request.FeatureName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEnvironmentFeature"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEnvironmentFeatureResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries instant messaging (IM) chatbots.
//
// @param request - DescribeIMRobotsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIMRobotsResponse
func (client *Client) DescribeIMRobotsWithContext(ctx context.Context, request *DescribeIMRobotsRequest, runtime *dara.RuntimeOptions) (_result *DescribeIMRobotsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
	}

	if !dara.IsNil(request.RobotIds) {
		query["RobotIds"] = request.RobotIds
	}

	if !dara.IsNil(request.RobotName) {
		query["RobotName"] = request.RobotName
	}

	if !dara.IsNil(request.Size) {
		query["Size"] = request.Size
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeIMRobots"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIMRobotsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about an alert rule for a Prometheus instance.
//
// @param request - DescribePrometheusAlertRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePrometheusAlertRuleResponse
func (client *Client) DescribePrometheusAlertRuleWithContext(ctx context.Context, request *DescribePrometheusAlertRuleRequest, runtime *dara.RuntimeOptions) (_result *DescribePrometheusAlertRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertId) {
		query["AlertId"] = request.AlertId
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePrometheusAlertRule"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePrometheusAlertRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the license key.
//
// @param request - DescribeTraceLicenseKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTraceLicenseKeyResponse
func (client *Client) DescribeTraceLicenseKeyWithContext(ctx context.Context, request *DescribeTraceLicenseKeyRequest, runtime *dara.RuntimeOptions) (_result *DescribeTraceLicenseKeyResponse, _err error) {
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
		Action:      dara.String("DescribeTraceLicenseKey"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTraceLicenseKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of webhook alert contacts.
//
// @param request - DescribeWebhookContactsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWebhookContactsResponse
func (client *Client) DescribeWebhookContactsWithContext(ctx context.Context, request *DescribeWebhookContactsRequest, runtime *dara.RuntimeOptions) (_result *DescribeWebhookContactsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWebhookContacts"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWebhookContactsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs actions based on the specified module type.
//
// @param request - DoInsightsActionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DoInsightsActionResponse
func (client *Client) DoInsightsActionWithContext(ctx context.Context, request *DoInsightsActionRequest, runtime *dara.RuntimeOptions) (_result *DoInsightsActionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Data) {
		body["Data"] = request.Data
	}

	if !dara.IsNil(request.Module) {
		body["Module"] = request.Module
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DoInsightsAction"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DoInsightsActionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI EnableMetric is deprecated
//
// Summary:
//
// Enables a discarded metric.
//
// @param request - EnableMetricRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableMetricResponse
func (client *Client) EnableMetricWithContext(ctx context.Context, request *EnableMetricRequest, runtime *dara.RuntimeOptions) (_result *EnableMetricResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DropMetric) {
		query["DropMetric"] = request.DropMetric
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableMetric"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableMetricResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the URL for downloading an agent.
//
// @param request - GetAgentDownloadUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentDownloadUrlResponse
func (client *Client) GetAgentDownloadUrlWithContext(ctx context.Context, request *GetAgentDownloadUrlRequest, runtime *dara.RuntimeOptions) (_result *GetAgentDownloadUrlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAgentDownloadUrl"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgentDownloadUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the URL for downloading an agent.
//
// @param request - GetAgentDownloadUrlV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentDownloadUrlV2Response
func (client *Client) GetAgentDownloadUrlV2WithContext(ctx context.Context, request *GetAgentDownloadUrlV2Request, runtime *dara.RuntimeOptions) (_result *GetAgentDownloadUrlV2Response, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentType) {
		query["AgentType"] = request.AgentType
	}

	if !dara.IsNil(request.ArchType) {
		query["ArchType"] = request.ArchType
	}

	if !dara.IsNil(request.OsType) {
		query["OsType"] = request.OsType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAgentDownloadUrlV2"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgentDownloadUrlV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries alert rules.
//
// @param request - GetAlertRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAlertRulesResponse
func (client *Client) GetAlertRulesWithContext(ctx context.Context, request *GetAlertRulesRequest, runtime *dara.RuntimeOptions) (_result *GetAlertRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertIds) {
		query["AlertIds"] = request.AlertIds
	}

	if !dara.IsNil(request.AlertNames) {
		query["AlertNames"] = request.AlertNames
	}

	if !dara.IsNil(request.AlertStatus) {
		query["AlertStatus"] = request.AlertStatus
	}

	if !dara.IsNil(request.AlertType) {
		query["AlertType"] = request.AlertType
	}

	if !dara.IsNil(request.BizRegionId) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Size) {
		query["Size"] = request.Size
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAlertRules"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAlertRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the API operations of application monitoring by page.
//
// @param request - GetAppApiByPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppApiByPageResponse
func (client *Client) GetAppApiByPageWithContext(ctx context.Context, request *GetAppApiByPageRequest, runtime *dara.RuntimeOptions) (_result *GetAppApiByPageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IntervalMills) {
		query["IntervalMills"] = request.IntervalMills
	}

	if !dara.IsNil(request.PId) {
		query["PId"] = request.PId
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
		Action:      dara.String("GetAppApiByPage"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppApiByPageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Obtain the JVM configuration information of each instance of the application
//
// @param request - GetAppJVMConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppJVMConfigResponse
func (client *Client) GetAppJVMConfigWithContext(ctx context.Context, request *GetAppJVMConfigRequest, runtime *dara.RuntimeOptions) (_result *GetAppJVMConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppJVMConfig"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppJVMConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains an authentication token. When you connect a Container Service for Kubernetes (ACK) cluster to Prometheus Service over the Internet, you must use a token for authentication.
//
// @param request - GetAuthTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAuthTokenResponse
func (client *Client) GetAuthTokenWithContext(ctx context.Context, request *GetAuthTokenRequest, runtime *dara.RuntimeOptions) (_result *GetAuthTokenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("GetAuthToken"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAuthTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetCloudClusterAllUrl is deprecated, please use ARMS::2019-08-08::GetRemoteWriteUrl instead.
//
// Summary:
//
// Queries the read and write URLs of a CloudMonitor instance, such as Pushgateway and Grafana URLs.
//
// @param request - GetCloudClusterAllUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCloudClusterAllUrlResponse
func (client *Client) GetCloudClusterAllUrlWithContext(ctx context.Context, request *GetCloudClusterAllUrlRequest, runtime *dara.RuntimeOptions) (_result *GetCloudClusterAllUrlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("GetCloudClusterAllUrl"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCloudClusterAllUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetClusterAllUrl is deprecated, please use ARMS::2019-08-08::GetPrometheusInstance instead.
//
// Summary:
//
// Obtains all the URLs of a cluster, including remote read and write URLs, Pushgateway URLs, and Grafana URLs.
//
// @param request - GetClusterAllUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClusterAllUrlResponse
func (client *Client) GetClusterAllUrlWithContext(ctx context.Context, request *GetClusterAllUrlRequest, runtime *dara.RuntimeOptions) (_result *GetClusterAllUrlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("GetClusterAllUrl"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetClusterAllUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether the current account has activated the commercial edition of a service.
//
// @param request - GetCommercialStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCommercialStatusResponse
func (client *Client) GetCommercialStatusWithContext(ctx context.Context, request *GetCommercialStatusRequest, runtime *dara.RuntimeOptions) (_result *GetCommercialStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CommodityCode) {
		query["CommodityCode"] = request.CommodityCode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCommercialStatus"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCommercialStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetExploreUrl is deprecated
//
// Summary:
//
// Enables the Explore feature of Grafana.
//
// @param request - GetExploreUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExploreUrlResponse
func (client *Client) GetExploreUrlWithContext(ctx context.Context, request *GetExploreUrlRequest, runtime *dara.RuntimeOptions) (_result *GetExploreUrlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Expression) {
		query["Expression"] = request.Expression
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
		Action:      dara.String("GetExploreUrl"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetExploreUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a Grafana workspace.
//
// Description:
//
// Note: The list returned by this operation includes the workspaces of Developer Edition, Expert Edition, and Advanced Edition. The list does not include the workspaces of Shared Edition.
//
// @param request - GetGrafanaWorkspaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGrafanaWorkspaceResponse
func (client *Client) GetGrafanaWorkspaceWithContext(ctx context.Context, request *GetGrafanaWorkspaceRequest, runtime *dara.RuntimeOptions) (_result *GetGrafanaWorkspaceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunLang) {
		query["AliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.GrafanaWorkspaceId) {
		query["GrafanaWorkspaceId"] = request.GrafanaWorkspaceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGrafanaWorkspace"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGrafanaWorkspaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetIntegrationState is deprecated, please use ARMS::2019-08-08::DescribeAddonRelease instead.
//
// Summary:
//
// Queries the integration state of the dashboards and collection rules of Application Real-Time Monitoring Service (ARMS) Prometheus.
//
// @param request - GetIntegrationStateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIntegrationStateResponse
func (client *Client) GetIntegrationStateWithContext(ctx context.Context, request *GetIntegrationStateRequest, runtime *dara.RuntimeOptions) (_result *GetIntegrationStateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Integration) {
		query["Integration"] = request.Integration
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIntegrationState"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIntegrationStateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetManagedPrometheusStatus is deprecated
//
// Summary:
//
// Queries the installation status of a Prometheus agent in a serverless Kubernetes (ASK) cluster or an Elastic Compute Service (ECS) cluster.
//
// @param request - GetManagedPrometheusStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetManagedPrometheusStatusResponse
func (client *Client) GetManagedPrometheusStatusWithContext(ctx context.Context, request *GetManagedPrometheusStatusRequest, runtime *dara.RuntimeOptions) (_result *GetManagedPrometheusStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ClusterType) {
		query["ClusterType"] = request.ClusterType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetManagedPrometheusStatus"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetManagedPrometheusStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of one or more traces.
//
// @param request - GetMultipleTraceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMultipleTraceResponse
func (client *Client) GetMultipleTraceWithContext(ctx context.Context, request *GetMultipleTraceRequest, runtime *dara.RuntimeOptions) (_result *GetMultipleTraceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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

	if !dara.IsNil(request.TraceIDs) {
		query["TraceIDs"] = request.TraceIDs
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMultipleTrace"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMultipleTraceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a scheduling policy.
//
// @param request - GetOnCallSchedulesDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOnCallSchedulesDetailResponse
func (client *Client) GetOnCallSchedulesDetailWithContext(ctx context.Context, request *GetOnCallSchedulesDetailRequest, runtime *dara.RuntimeOptions) (_result *GetOnCallSchedulesDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOnCallSchedulesDetail"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOnCallSchedulesDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the token required for integrating Prometheus Service.
//
// Description:
//
// None.
//
// @param request - GetPrometheusApiTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPrometheusApiTokenResponse
func (client *Client) GetPrometheusApiTokenWithContext(ctx context.Context, request *GetPrometheusApiTokenRequest, runtime *dara.RuntimeOptions) (_result *GetPrometheusApiTokenResponse, _err error) {
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
		Action:      dara.String("GetPrometheusApiToken"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPrometheusApiTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a global aggregation instance.
//
// @param request - GetPrometheusGlobalViewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPrometheusGlobalViewResponse
func (client *Client) GetPrometheusGlobalViewWithContext(ctx context.Context, request *GetPrometheusGlobalViewRequest, runtime *dara.RuntimeOptions) (_result *GetPrometheusGlobalViewResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GlobalViewClusterId) {
		query["GlobalViewClusterId"] = request.GlobalViewClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPrometheusGlobalView"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPrometheusGlobalViewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a Prometheus instance.
//
// @param request - GetPrometheusInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPrometheusInstanceResponse
func (client *Client) GetPrometheusInstanceWithContext(ctx context.Context, request *GetPrometheusInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetPrometheusInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("GetPrometheusInstance"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPrometheusInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetPrometheusIntegration is deprecated
//
// Summary:
//
// Queries the information about an exporter that is integrated into a Prometheus instance for Container Service or a Prometheus instance for ECS.
//
// @param request - GetPrometheusIntegrationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPrometheusIntegrationResponse
func (client *Client) GetPrometheusIntegrationWithContext(ctx context.Context, request *GetPrometheusIntegrationRequest, runtime *dara.RuntimeOptions) (_result *GetPrometheusIntegrationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IntegrationType) {
		query["IntegrationType"] = request.IntegrationType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPrometheusIntegration"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPrometheusIntegrationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetPrometheusMonitoring is deprecated
//
// Summary:
//
// Queries the monitoring configuration of a Prometheus instance.
//
// @param request - GetPrometheusMonitoringRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPrometheusMonitoringResponse
func (client *Client) GetPrometheusMonitoringWithContext(ctx context.Context, request *GetPrometheusMonitoringRequest, runtime *dara.RuntimeOptions) (_result *GetPrometheusMonitoringResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.MonitoringName) {
		query["MonitoringName"] = request.MonitoringName
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
		Action:      dara.String("GetPrometheusMonitoring"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPrometheusMonitoringResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetRecordingRule is deprecated
//
// Summary:
//
// Obtains the recording rule of a cluster.
//
// @param request - GetRecordingRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRecordingRuleResponse
func (client *Client) GetRecordingRuleWithContext(ctx context.Context, request *GetRecordingRuleRequest, runtime *dara.RuntimeOptions) (_result *GetRecordingRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("GetRecordingRule"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRecordingRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the objects of a Browser Monitoring application by process identifier (PID).
//
// @param request - GetRetcodeAppByPidRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRetcodeAppByPidResponse
func (client *Client) GetRetcodeAppByPidWithContext(ctx context.Context, request *GetRetcodeAppByPidRequest, runtime *dara.RuntimeOptions) (_result *GetRetcodeAppByPidResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRetcodeAppByPid"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRetcodeAppByPidResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Browser Monitoring data based on a query statement of Log Service.
//
// @param request - GetRetcodeDataByQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRetcodeDataByQueryResponse
func (client *Client) GetRetcodeDataByQueryWithContext(ctx context.Context, request *GetRetcodeDataByQueryRequest, runtime *dara.RuntimeOptions) (_result *GetRetcodeDataByQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.From) {
		query["From"] = request.From
	}

	if !dara.IsNil(request.Pid) {
		query["Pid"] = request.Pid
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.To) {
		query["To"] = request.To
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRetcodeDataByQuery"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRetcodeDataByQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Log Service project and Logstore that correspond to an application of browser monitoring.
//
// @param request - GetRetcodeLogstoreRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRetcodeLogstoreResponse
func (client *Client) GetRetcodeLogstoreWithContext(ctx context.Context, request *GetRetcodeLogstoreRequest, runtime *dara.RuntimeOptions) (_result *GetRetcodeLogstoreResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Pid) {
		query["Pid"] = request.Pid
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRetcodeLogstore"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRetcodeLogstoreResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the share URL of an application monitored by Browser Monitoring.
//
// @param request - GetRetcodeShareUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRetcodeShareUrlResponse
func (client *Client) GetRetcodeShareUrlWithContext(ctx context.Context, request *GetRetcodeShareUrlRequest, runtime *dara.RuntimeOptions) (_result *GetRetcodeShareUrlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Pid) {
		query["Pid"] = request.Pid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRetcodeShareUrl"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRetcodeShareUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the information about a single application in Browser Monitoring.
//
// Description:
//
// Real User Monitoring (RUM) is available only in the China (Hangzhou), Singapore, and US (Silicon Valley) regions. Select the correct endpoint.
//
// @param request - GetRumAppInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRumAppInfoResponse
func (client *Client) GetRumAppInfoWithContext(ctx context.Context, request *GetRumAppInfoRequest, runtime *dara.RuntimeOptions) (_result *GetRumAppInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppGroup) {
		query["AppGroup"] = request.AppGroup
	}

	if !dara.IsNil(request.Pid) {
		query["Pid"] = request.Pid
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRumAppInfo"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRumAppInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of Real User Monitoring (RUM) applications.
//
// Description:
//
// Real User Monitoring (RUM) is available only in the China (Hangzhou), Singapore, and US (Silicon Valley) regions. Select the correct endpoint.
//
// @param tmpReq - GetRumAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRumAppsResponse
func (client *Client) GetRumAppsWithContext(ctx context.Context, tmpReq *GetRumAppsRequest, runtime *dara.RuntimeOptions) (_result *GetRumAppsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GetRumAppsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppGroup) {
		query["AppGroup"] = request.AppGroup
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.AppType) {
		query["AppType"] = request.AppType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRumApps"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRumAppsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Real User Monitoring (RUM) data by page.
//
// Description:
//
// Real User Monitoring (RUM) is available only in the China (Hangzhou), Singapore, and US (Silicon Valley) regions. Select the correct endpoint.
//
// @param request - GetRumDataForPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRumDataForPageResponse
func (client *Client) GetRumDataForPageWithContext(ctx context.Context, request *GetRumDataForPageRequest, runtime *dara.RuntimeOptions) (_result *GetRumDataForPageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppGroup) {
		query["AppGroup"] = request.AppGroup
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Pid) {
		query["Pid"] = request.Pid
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
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
		Action:      dara.String("GetRumDataForPage"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRumDataForPageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the exception stack information of a Real User Monitoring (RUM) application.
//
// Description:
//
// Real User Monitoring (RUM) is available only in the China (Hangzhou), Singapore, and US (Silicon Valley) regions. Select the correct endpoint.
//
// @param request - GetRumExceptionStackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRumExceptionStackResponse
func (client *Client) GetRumExceptionStackWithContext(ctx context.Context, request *GetRumExceptionStackRequest, runtime *dara.RuntimeOptions) (_result *GetRumExceptionStackResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExceptionBinaryImages) {
		query["ExceptionBinaryImages"] = request.ExceptionBinaryImages
	}

	if !dara.IsNil(request.ExceptionStack) {
		query["ExceptionStack"] = request.ExceptionStack
	}

	if !dara.IsNil(request.ExceptionThreadId) {
		query["ExceptionThreadId"] = request.ExceptionThreadId
	}

	if !dara.IsNil(request.ExtraInfo) {
		query["ExtraInfo"] = request.ExtraInfo
	}

	if !dara.IsNil(request.Pid) {
		query["Pid"] = request.Pid
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SourcemapType) {
		query["SourcemapType"] = request.SourcemapType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRumExceptionStack"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRumExceptionStackResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the observability capacity unit (OCU) usage data of Real User Monitoring (RUM).
//
// Description:
//
// You can query the usage data for the current day at any time. You can query the usage data for the previous day only after 8:00 today.
//
// @param tmpReq - GetRumOcuStatisticDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRumOcuStatisticDataResponse
func (client *Client) GetRumOcuStatisticDataWithContext(ctx context.Context, tmpReq *GetRumOcuStatisticDataRequest, runtime *dara.RuntimeOptions) (_result *GetRumOcuStatisticDataResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GetRumOcuStatisticDataShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Filter) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, dara.String("Filter"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Group) {
		request.GroupShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Group, dara.String("Group"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRumOcuStatisticData"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRumOcuStatisticDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Real User Monitoring (RUM)-related files, such as symbol tables and SourceMap.
//
// Description:
//
// Real User Monitoring (RUM) is available only in the China (Hangzhou), Singapore, and US (Silicon Valley) regions. Select the correct endpoint.
//
// @param request - GetRumUploadFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRumUploadFilesResponse
func (client *Client) GetRumUploadFilesWithContext(ctx context.Context, request *GetRumUploadFilesRequest, runtime *dara.RuntimeOptions) (_result *GetRumUploadFilesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		query["AppType"] = request.AppType
	}

	if !dara.IsNil(request.Pid) {
		query["Pid"] = request.Pid
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VersionId) {
		query["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRumUploadFiles"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRumUploadFilesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details of the SourceMap file uploaded in Browser Monitoring.
//
// @param request - GetSourceMapInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSourceMapInfoResponse
func (client *Client) GetSourceMapInfoWithContext(ctx context.Context, request *GetSourceMapInfoRequest, runtime *dara.RuntimeOptions) (_result *GetSourceMapInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AscendingSequence) {
		query["AscendingSequence"] = request.AscendingSequence
	}

	if !dara.IsNil(request.Edition) {
		query["Edition"] = request.Edition
	}

	if !dara.IsNil(request.ID) {
		query["ID"] = request.ID
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.OrderField) {
		query["OrderField"] = request.OrderField
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSourceMapInfo"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSourceMapInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of a method stack.
//
// @param request - GetStackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStackResponse
func (client *Client) GetStackWithContext(ctx context.Context, request *GetStackRequest, runtime *dara.RuntimeOptions) (_result *GetStackResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Pid) {
		query["Pid"] = request.Pid
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RpcID) {
		query["RpcID"] = request.RpcID
	}

	if !dara.IsNil(request.SpanID) {
		query["SpanID"] = request.SpanID
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TraceID) {
		query["TraceID"] = request.TraceID
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStack"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStackResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains detection points.
//
// @param tmpReq - GetSyntheticMonitorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSyntheticMonitorsResponse
func (client *Client) GetSyntheticMonitorsWithContext(ctx context.Context, tmpReq *GetSyntheticMonitorsRequest, runtime *dara.RuntimeOptions) (_result *GetSyntheticMonitorsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GetSyntheticMonitorsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Filter) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, dara.String("Filter"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSyntheticMonitors"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSyntheticMonitorsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a scheduled synthetic monitoring task.
//
// @param request - GetSyntheticTaskDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSyntheticTaskDetailResponse
func (client *Client) GetSyntheticTaskDetailWithContext(ctx context.Context, request *GetSyntheticTaskDetailRequest, runtime *dara.RuntimeOptions) (_result *GetSyntheticTaskDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSyntheticTaskDetail"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSyntheticTaskDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of scheduled synthetic monitoring tasks.
//
// @param request - GetSyntheticTaskListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSyntheticTaskListResponse
func (client *Client) GetSyntheticTaskListWithContext(ctx context.Context, request *GetSyntheticTaskListRequest, runtime *dara.RuntimeOptions) (_result *GetSyntheticTaskListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
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

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TaskStatus) {
		query["TaskStatus"] = request.TaskStatus
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	if !dara.IsNil(request.Url) {
		query["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSyntheticTaskList"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSyntheticTaskListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the information about synthetic monitoring points.
//
// @param request - GetSyntheticTaskMonitorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSyntheticTaskMonitorsResponse
func (client *Client) GetSyntheticTaskMonitorsWithContext(ctx context.Context, request *GetSyntheticTaskMonitorsRequest, runtime *dara.RuntimeOptions) (_result *GetSyntheticTaskMonitorsResponse, _err error) {
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
		Action:      dara.String("GetSyntheticTaskMonitors"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSyntheticTaskMonitorsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details of a synthetic monitoring task.
//
// @param request - GetTimingSyntheticTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTimingSyntheticTaskResponse
func (client *Client) GetTimingSyntheticTaskWithContext(ctx context.Context, request *GetTimingSyntheticTaskRequest, runtime *dara.RuntimeOptions) (_result *GetTimingSyntheticTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTimingSyntheticTask"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTimingSyntheticTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a trace.
//
// Description:
//
// > You must use Application Real-Time Monitoring Service (ARMS) SDK for Java V2.7.24.
//
// @param request - GetTraceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTraceResponse
func (client *Client) GetTraceWithContext(ctx context.Context, request *GetTraceRequest, runtime *dara.RuntimeOptions) (_result *GetTraceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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

	if !dara.IsNil(request.TraceID) {
		query["TraceID"] = request.TraceID
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTrace"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTraceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an application monitoring task.
//
// @param request - GetTraceAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTraceAppResponse
func (client *Client) GetTraceAppWithContext(ctx context.Context, request *GetTraceAppRequest, runtime *dara.RuntimeOptions) (_result *GetTraceAppResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Pid) {
		query["Pid"] = request.Pid
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTraceApp"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTraceAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all custom settings of an application monitored by Application Monitoring, such as trace sampling settings and agent switches. This operation is applicable only to applications that are monitored by Application Monitoring. It is not applicable to applications that are monitored by Managed Service for OpenTelemetry.
//
// @param request - GetTraceAppConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTraceAppConfigResponse
func (client *Client) GetTraceAppConfigWithContext(ctx context.Context, request *GetTraceAppConfigRequest, runtime *dara.RuntimeOptions) (_result *GetTraceAppConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Pid) {
		query["Pid"] = request.Pid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTraceAppConfig"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTraceAppConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an alert rule based on an alert template.
//
// Description:
//
// >  You can call the **ImportAppAlertRules*	- operation to import only the alert rules that are generated by Application Real-Time Monitoring Service (ARMS) for application monitoring and browser monitoring. This operation cannot be used to import custom alert rules, alert rules for Prometheus monitoring, or default emergency alert rules.
//
// @param request - ImportAppAlertRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportAppAlertRulesResponse
func (client *Client) ImportAppAlertRulesWithContext(ctx context.Context, request *ImportAppAlertRulesRequest, runtime *dara.RuntimeOptions) (_result *ImportAppAlertRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactGroupIds) {
		query["ContactGroupIds"] = request.ContactGroupIds
	}

	if !dara.IsNil(request.IsAutoStart) {
		query["IsAutoStart"] = request.IsAutoStart
	}

	if !dara.IsNil(request.Pids) {
		query["Pids"] = request.Pids
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.TemplageAlertConfig) {
		query["TemplageAlertConfig"] = request.TemplageAlertConfig
	}

	if !dara.IsNil(request.TemplateAlertId) {
		query["TemplateAlertId"] = request.TemplateAlertId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportAppAlertRules"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportAppAlertRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initializes an environment instance.
//
// @param request - InitEnvironmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitEnvironmentResponse
func (client *Client) InitEnvironmentWithContext(ctx context.Context, request *InitEnvironmentRequest, runtime *dara.RuntimeOptions) (_result *InitEnvironmentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunLang) {
		query["AliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.CreateAuthToken) {
		query["CreateAuthToken"] = request.CreateAuthToken
	}

	if !dara.IsNil(request.EnvironmentId) {
		query["EnvironmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.ManagedType) {
		query["ManagedType"] = request.ManagedType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InitEnvironment"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InitEnvironmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Installs an add-on.
//
// @param request - InstallAddonRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallAddonResponse
func (client *Client) InstallAddonWithContext(ctx context.Context, request *InstallAddonRequest, runtime *dara.RuntimeOptions) (_result *InstallAddonResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddonVersion) {
		query["AddonVersion"] = request.AddonVersion
	}

	if !dara.IsNil(request.AliyunLang) {
		query["AliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EnvironmentId) {
		query["EnvironmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReleaseName) {
		query["ReleaseName"] = request.ReleaseName
	}

	if !dara.IsNil(request.Values) {
		query["Values"] = request.Values
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InstallAddon"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InstallAddonResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI InstallCmsExporter is deprecated, please use ARMS::2019-08-08::InstallAddon instead.
//
// Summary:
//
// Installs the cms-exporter collector.
//
// @param request - InstallCmsExporterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallCmsExporterResponse
func (client *Client) InstallCmsExporterWithContext(ctx context.Context, request *InstallCmsExporterRequest, runtime *dara.RuntimeOptions) (_result *InstallCmsExporterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.CmsArgs) {
		query["CmsArgs"] = request.CmsArgs
	}

	if !dara.IsNil(request.DirectArgs) {
		query["DirectArgs"] = request.DirectArgs
	}

	if !dara.IsNil(request.EnableTag) {
		query["EnableTag"] = request.EnableTag
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InstallCmsExporter"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InstallCmsExporterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Installs a feature.
//
// @param request - InstallEnvironmentFeatureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallEnvironmentFeatureResponse
func (client *Client) InstallEnvironmentFeatureWithContext(ctx context.Context, request *InstallEnvironmentFeatureRequest, runtime *dara.RuntimeOptions) (_result *InstallEnvironmentFeatureResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunLang) {
		query["AliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.Config) {
		query["Config"] = request.Config
	}

	if !dara.IsNil(request.EnvironmentId) {
		query["EnvironmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.FeatureName) {
		query["FeatureName"] = request.FeatureName
	}

	if !dara.IsNil(request.FeatureVersion) {
		query["FeatureVersion"] = request.FeatureVersion
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InstallEnvironmentFeature"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InstallEnvironmentFeatureResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI InstallManagedPrometheus is deprecated
//
// Summary:
//
// Installs a Prometheus agent for serverless Kubernetes (ASK) clusters or Elastic Compute Service (ECS) clusters.
//
// Description:
//
// You can call this operation only if the following conditions are met: The resources that you want to monitor are ASK clusters or ECS clusters. No Prometheus agents are installed in the ASK or ECS clusters. Take note that Prometheus agents can be installed only on the cloud service side, not in user clusters.
//
// @param request - InstallManagedPrometheusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallManagedPrometheusResponse
func (client *Client) InstallManagedPrometheusWithContext(ctx context.Context, request *InstallManagedPrometheusRequest, runtime *dara.RuntimeOptions) (_result *InstallManagedPrometheusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.ClusterType) {
		query["ClusterType"] = request.ClusterType
	}

	if !dara.IsNil(request.GrafanaInstanceId) {
		query["GrafanaInstanceId"] = request.GrafanaInstanceId
	}

	if !dara.IsNil(request.KubeConfig) {
		query["KubeConfig"] = request.KubeConfig
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VcExtraInfo) {
		query["VcExtraInfo"] = request.VcExtraInfo
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InstallManagedPrometheus"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InstallManagedPrometheusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the alerts that have been triggered.
//
// @param request - ListActivatedAlertsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListActivatedAlertsResponse
func (client *Client) ListActivatedAlertsWithContext(ctx context.Context, request *ListActivatedAlertsRequest, runtime *dara.RuntimeOptions) (_result *ListActivatedAlertsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
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
		Action:      dara.String("ListActivatedAlerts"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListActivatedAlertsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the add-ons installed in an environment.
//
// @param request - ListAddonReleasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAddonReleasesResponse
func (client *Client) ListAddonReleasesWithContext(ctx context.Context, request *ListAddonReleasesRequest, runtime *dara.RuntimeOptions) (_result *ListAddonReleasesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddonName) {
		query["AddonName"] = request.AddonName
	}

	if !dara.IsNil(request.EnvironmentId) {
		query["EnvironmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAddonReleases"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAddonReleasesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// List of access center products.
//
// @param request - ListAddonsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAddonsResponse
func (client *Client) ListAddonsWithContext(ctx context.Context, request *ListAddonsRequest, runtime *dara.RuntimeOptions) (_result *ListAddonsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunLang) {
		query["AliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.Regexp) {
		query["Regexp"] = request.Regexp
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Search) {
		query["Search"] = request.Search
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAddons"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAddonsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries historical alert events.
//
// @param request - ListAlertEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlertEventsResponse
func (client *Client) ListAlertEventsWithContext(ctx context.Context, request *ListAlertEventsRequest, runtime *dara.RuntimeOptions) (_result *ListAlertEventsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertName) {
		query["AlertName"] = request.AlertName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.MatchingConditions) {
		query["MatchingConditions"] = request.MatchingConditions
	}

	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
	}

	if !dara.IsNil(request.ShowNotificationPolicies) {
		query["ShowNotificationPolicies"] = request.ShowNotificationPolicies
	}

	if !dara.IsNil(request.Size) {
		query["Size"] = request.Size
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
		Action:      dara.String("ListAlertEvents"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAlertEventsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the alert sending history.
//
// @param request - ListAlertsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlertsResponse
func (client *Client) ListAlertsWithContext(ctx context.Context, request *ListAlertsRequest, runtime *dara.RuntimeOptions) (_result *ListAlertsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertName) {
		query["AlertName"] = request.AlertName
	}

	if !dara.IsNil(request.DispatchRuleId) {
		query["DispatchRuleId"] = request.DispatchRuleId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IntegrationType) {
		query["IntegrationType"] = request.IntegrationType
	}

	if !dara.IsNil(request.Owner) {
		query["Owner"] = request.Owner
	}

	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Severity) {
		query["Severity"] = request.Severity
	}

	if !dara.IsNil(request.ShowActivities) {
		query["ShowActivities"] = request.ShowActivities
	}

	if !dara.IsNil(request.ShowEvents) {
		query["ShowEvents"] = request.ShowEvents
	}

	if !dara.IsNil(request.Size) {
		query["Size"] = request.Size
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAlerts"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAlertsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ListClusterFromGrafana is deprecated
//
// Summary:
//
// Queries all Grafana dashboards in a specified region.
//
// @param request - ListClusterFromGrafanaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClusterFromGrafanaResponse
func (client *Client) ListClusterFromGrafanaWithContext(ctx context.Context, request *ListClusterFromGrafanaRequest, runtime *dara.RuntimeOptions) (_result *ListClusterFromGrafanaResponse, _err error) {
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
		Action:      dara.String("ListClusterFromGrafana"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClusterFromGrafanaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ListCmsInstances is deprecated
//
// Summary:
//
// Queries the collection of cloud services.
//
// @param request - ListCmsInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCmsInstancesResponse
func (client *Client) ListCmsInstancesWithContext(ctx context.Context, request *ListCmsInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListCmsInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TypeFilter) {
		query["TypeFilter"] = request.TypeFilter
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCmsInstances"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCmsInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Grafana dashboards of a Container Service for Kubernetes (ACK) cluster.
//
// Description:
//
// None.
//
// @param request - ListDashboardsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDashboardsResponse
func (client *Client) ListDashboardsWithContext(ctx context.Context, request *ListDashboardsRequest, runtime *dara.RuntimeOptions) (_result *ListDashboardsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ClusterType) {
		query["ClusterType"] = request.ClusterType
	}

	if !dara.IsNil(request.DashboardName) {
		query["DashboardName"] = request.DashboardName
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	if !dara.IsNil(request.RecreateSwitch) {
		query["RecreateSwitch"] = request.RecreateSwitch
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDashboards"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDashboardsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ListDashboardsByName is deprecated
//
// Summary:
//
// Uses Loki data sources and other data sources to create a Grafana dashboard in Managed Service for Prometheus.
//
// @param request - ListDashboardsByNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDashboardsByNameResponse
func (client *Client) ListDashboardsByNameWithContext(ctx context.Context, request *ListDashboardsByNameRequest, runtime *dara.RuntimeOptions) (_result *ListDashboardsByNameResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ClusterType) {
		query["ClusterType"] = request.ClusterType
	}

	if !dara.IsNil(request.DashBoardName) {
		query["DashBoardName"] = request.DashBoardName
	}

	if !dara.IsNil(request.DashBoardVersion) {
		query["DashBoardVersion"] = request.DashBoardVersion
	}

	if !dara.IsNil(request.DataSourceType) {
		query["DataSourceType"] = request.DataSourceType
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.OnlyQuery) {
		query["OnlyQuery"] = request.OnlyQuery
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDashboardsByName"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDashboardsByNameResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries notification policies.
//
// Description:
//
// The current API operation is no longer maintained. To query the notification policy information, call the ListNotificationPolicies operation instead.
//
// @param request - ListDispatchRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDispatchRuleResponse
func (client *Client) ListDispatchRuleWithContext(ctx context.Context, request *ListDispatchRuleRequest, runtime *dara.RuntimeOptions) (_result *ListDispatchRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.System) {
		query["System"] = request.System
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDispatchRule"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDispatchRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the custom jobs of an environment.
//
// @param request - ListEnvCustomJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEnvCustomJobsResponse
func (client *Client) ListEnvCustomJobsWithContext(ctx context.Context, request *ListEnvCustomJobsRequest, runtime *dara.RuntimeOptions) (_result *ListEnvCustomJobsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EncryptYaml) {
		query["EncryptYaml"] = request.EncryptYaml
	}

	if !dara.IsNil(request.EnvironmentId) {
		query["EnvironmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEnvCustomJobs"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEnvCustomJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the PodMonitors of an environment.
//
// @param request - ListEnvPodMonitorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEnvPodMonitorsResponse
func (client *Client) ListEnvPodMonitorsWithContext(ctx context.Context, request *ListEnvPodMonitorsRequest, runtime *dara.RuntimeOptions) (_result *ListEnvPodMonitorsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnvironmentId) {
		query["EnvironmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEnvPodMonitors"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEnvPodMonitorsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the ServiceMonitors of an environment.
//
// @param request - ListEnvServiceMonitorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEnvServiceMonitorsResponse
func (client *Client) ListEnvServiceMonitorsWithContext(ctx context.Context, request *ListEnvServiceMonitorsRequest, runtime *dara.RuntimeOptions) (_result *ListEnvServiceMonitorsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnvironmentId) {
		query["EnvironmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEnvServiceMonitors"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEnvServiceMonitorsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 环境addon列表
//
// @param request - ListEnvironmentAddonsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEnvironmentAddonsResponse
func (client *Client) ListEnvironmentAddonsWithContext(ctx context.Context, request *ListEnvironmentAddonsRequest, runtime *dara.RuntimeOptions) (_result *ListEnvironmentAddonsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnvironmentId) {
		query["EnvironmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEnvironmentAddons"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEnvironmentAddonsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the alert groups of an environment instance.
//
// @param request - ListEnvironmentAlertRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEnvironmentAlertRulesResponse
func (client *Client) ListEnvironmentAlertRulesWithContext(ctx context.Context, request *ListEnvironmentAlertRulesRequest, runtime *dara.RuntimeOptions) (_result *ListEnvironmentAlertRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddonName) {
		query["AddonName"] = request.AddonName
	}

	if !dara.IsNil(request.EnvironmentId) {
		query["EnvironmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Scene) {
		query["Scene"] = request.Scene
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEnvironmentAlertRules"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEnvironmentAlertRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a dashboard of an environment instance.
//
// @param request - ListEnvironmentDashboardsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEnvironmentDashboardsResponse
func (client *Client) ListEnvironmentDashboardsWithContext(ctx context.Context, request *ListEnvironmentDashboardsRequest, runtime *dara.RuntimeOptions) (_result *ListEnvironmentDashboardsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddonName) {
		query["AddonName"] = request.AddonName
	}

	if !dara.IsNil(request.EnvironmentId) {
		query["EnvironmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Scene) {
		query["Scene"] = request.Scene
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEnvironmentDashboards"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEnvironmentDashboardsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the features in an environment.
//
// @param request - ListEnvironmentFeaturesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEnvironmentFeaturesResponse
func (client *Client) ListEnvironmentFeaturesWithContext(ctx context.Context, request *ListEnvironmentFeaturesRequest, runtime *dara.RuntimeOptions) (_result *ListEnvironmentFeaturesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunLang) {
		query["AliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.EnvironmentId) {
		query["EnvironmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEnvironmentFeatures"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEnvironmentFeaturesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Kubernetes resources of an environment.
//
// @param tmpReq - ListEnvironmentKubeResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEnvironmentKubeResourcesResponse
func (client *Client) ListEnvironmentKubeResourcesWithContext(ctx context.Context, tmpReq *ListEnvironmentKubeResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListEnvironmentKubeResourcesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListEnvironmentKubeResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LabelSelectors) {
		request.LabelSelectorsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LabelSelectors, dara.String("LabelSelectors"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EnvironmentId) {
		query["EnvironmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.Kind) {
		query["Kind"] = request.Kind
	}

	if !dara.IsNil(request.LabelSelectorsShrink) {
		query["LabelSelectors"] = request.LabelSelectorsShrink
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEnvironmentKubeResources"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEnvironmentKubeResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the targets of an environment.
//
// @param request - ListEnvironmentMetricTargetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEnvironmentMetricTargetsResponse
func (client *Client) ListEnvironmentMetricTargetsWithContext(ctx context.Context, request *ListEnvironmentMetricTargetsRequest, runtime *dara.RuntimeOptions) (_result *ListEnvironmentMetricTargetsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnvironmentId) {
		query["EnvironmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.JobName) {
		query["JobName"] = request.JobName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEnvironmentMetricTargets"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEnvironmentMetricTargetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries environments.
//
// @param tmpReq - ListEnvironmentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEnvironmentsResponse
func (client *Client) ListEnvironmentsWithContext(ctx context.Context, tmpReq *ListEnvironmentsRequest, runtime *dara.RuntimeOptions) (_result *ListEnvironmentsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListEnvironmentsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AddonName) {
		query["AddonName"] = request.AddonName
	}

	if !dara.IsNil(request.BindResourceId) {
		query["BindResourceId"] = request.BindResourceId
	}

	if !dara.IsNil(request.EnvironmentType) {
		query["EnvironmentType"] = request.EnvironmentType
	}

	if !dara.IsNil(request.FeePackage) {
		query["FeePackage"] = request.FeePackage
	}

	if !dara.IsNil(request.FilterRegionIds) {
		query["FilterRegionIds"] = request.FilterRegionIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEnvironments"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEnvironmentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an escalation policy.
//
// @param request - ListEscalationPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEscalationPoliciesResponse
func (client *Client) ListEscalationPoliciesWithContext(ctx context.Context, request *ListEscalationPoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListEscalationPoliciesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEscalationPolicies"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEscalationPoliciesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries an EventBridge integration.
//
// @param request - ListEventBridgeIntegrationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEventBridgeIntegrationsResponse
func (client *Client) ListEventBridgeIntegrationsWithContext(ctx context.Context, request *ListEventBridgeIntegrationsRequest, runtime *dara.RuntimeOptions) (_result *ListEventBridgeIntegrationsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEventBridgeIntegrations"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEventBridgeIntegrationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Obtain Specified Workspace List
//
// Description:
//
// >The list returned by this operation includes the workspaces of Developer Edition, Expert Edition, and Advanced Edition. The list does not include the workspaces of Shared Edition.
//
// @param tmpReq - ListGrafanaWorkspaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGrafanaWorkspaceResponse
func (client *Client) ListGrafanaWorkspaceWithContext(ctx context.Context, tmpReq *ListGrafanaWorkspaceRequest, runtime *dara.RuntimeOptions) (_result *ListGrafanaWorkspaceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListGrafanaWorkspaceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunLang) {
		query["AliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGrafanaWorkspace"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGrafanaWorkspaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the abnormal Insights events within a specified period of time.
//
// @param request - ListInsightsEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInsightsEventsResponse
func (client *Client) ListInsightsEventsWithContext(ctx context.Context, request *ListInsightsEventsRequest, runtime *dara.RuntimeOptions) (_result *ListInsightsEventsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InsightsTypes) {
		query["InsightsTypes"] = request.InsightsTypes
	}

	if !dara.IsNil(request.Pid) {
		query["Pid"] = request.Pid
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
		Action:      dara.String("ListInsightsEvents"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInsightsEventsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Views a list of alert integrations.
//
// @param request - ListIntegrationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntegrationResponse
func (client *Client) ListIntegrationWithContext(ctx context.Context, request *ListIntegrationRequest, runtime *dara.RuntimeOptions) (_result *ListIntegrationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIntegration"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIntegrationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries notification policies based on specified conditions.
//
// @param request - ListNotificationPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNotificationPoliciesResponse
func (client *Client) ListNotificationPoliciesWithContext(ctx context.Context, request *ListNotificationPoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListNotificationPoliciesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectedMode) {
		query["DirectedMode"] = request.DirectedMode
	}

	if !dara.IsNil(request.Ids) {
		query["Ids"] = request.Ids
	}

	if !dara.IsNil(request.IsDetail) {
		query["IsDetail"] = request.IsDetail
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Size) {
		query["Size"] = request.Size
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNotificationPolicies"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNotificationPoliciesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a scheduling policy.
//
// @param request - ListOnCallSchedulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOnCallSchedulesResponse
func (client *Client) ListOnCallSchedulesWithContext(ctx context.Context, request *ListOnCallSchedulesRequest, runtime *dara.RuntimeOptions) (_result *ListOnCallSchedulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOnCallSchedules"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOnCallSchedulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the alert rules created for a Prometheus instance.
//
// @param request - ListPrometheusAlertRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrometheusAlertRulesResponse
func (client *Client) ListPrometheusAlertRulesWithContext(ctx context.Context, request *ListPrometheusAlertRulesRequest, runtime *dara.RuntimeOptions) (_result *ListPrometheusAlertRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.MatchExpressions) {
		query["MatchExpressions"] = request.MatchExpressions
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPrometheusAlertRules"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPrometheusAlertRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the alert templates of Prometheus Service.
//
// @param request - ListPrometheusAlertTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrometheusAlertTemplatesResponse
func (client *Client) ListPrometheusAlertTemplatesWithContext(ctx context.Context, request *ListPrometheusAlertTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListPrometheusAlertTemplatesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("ListPrometheusAlertTemplates"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPrometheusAlertTemplatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a global aggregation instance in Prometheus Service and obtains the list of global aggregation instances.
//
// @param request - ListPrometheusGlobalViewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrometheusGlobalViewResponse
func (client *Client) ListPrometheusGlobalViewWithContext(ctx context.Context, request *ListPrometheusGlobalViewRequest, runtime *dara.RuntimeOptions) (_result *ListPrometheusGlobalViewResponse, _err error) {
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
		Action:      dara.String("ListPrometheusGlobalView"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPrometheusGlobalViewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Prometheus instances by tag and resource group.
//
// @param request - ListPrometheusInstanceByTagAndResourceGroupIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrometheusInstanceByTagAndResourceGroupIdResponse
func (client *Client) ListPrometheusInstanceByTagAndResourceGroupIdWithContext(ctx context.Context, request *ListPrometheusInstanceByTagAndResourceGroupIdRequest, runtime *dara.RuntimeOptions) (_result *ListPrometheusInstanceByTagAndResourceGroupIdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      dara.String("ListPrometheusInstanceByTagAndResourceGroupId"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPrometheusInstanceByTagAndResourceGroupIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains all Prometheus instances in a region.
//
// @param request - ListPrometheusInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrometheusInstancesResponse
func (client *Client) ListPrometheusInstancesWithContext(ctx context.Context, request *ListPrometheusInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListPrometheusInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterType) {
		query["ClusterType"] = request.ClusterType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ShowGlobalView) {
		query["ShowGlobalView"] = request.ShowGlobalView
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPrometheusInstances"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPrometheusInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ListPrometheusIntegration is deprecated
//
// Summary:
//
// Queries a list of exporters that are integrated into a Prometheus instance. Only aliyun-cs and ecs instances are supported.
//
// @param request - ListPrometheusIntegrationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrometheusIntegrationResponse
func (client *Client) ListPrometheusIntegrationWithContext(ctx context.Context, request *ListPrometheusIntegrationRequest, runtime *dara.RuntimeOptions) (_result *ListPrometheusIntegrationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.IntegrationType) {
		query["IntegrationType"] = request.IntegrationType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPrometheusIntegration"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPrometheusIntegrationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ListPrometheusMonitoring is deprecated
//
// Summary:
//
// Queries the monitoring configuration of a Prometheus instance.
//
// @param request - ListPrometheusMonitoringRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrometheusMonitoringResponse
func (client *Client) ListPrometheusMonitoringWithContext(ctx context.Context, request *ListPrometheusMonitoringRequest, runtime *dara.RuntimeOptions) (_result *ListPrometheusMonitoringResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
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
		Action:      dara.String("ListPrometheusMonitoring"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPrometheusMonitoringResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Browser Monitoring tasks in a region.
//
// Description:
//
// ***
//
// @param request - ListRetcodeAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRetcodeAppsResponse
func (client *Client) ListRetcodeAppsWithContext(ctx context.Context, request *ListRetcodeAppsRequest, runtime *dara.RuntimeOptions) (_result *ListRetcodeAppsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRetcodeApps"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRetcodeAppsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListScenarioRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScenarioResponse
func (client *Client) ListScenarioWithContext(ctx context.Context, request *ListScenarioRequest, runtime *dara.RuntimeOptions) (_result *ListScenarioResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Scenario) {
		query["Scenario"] = request.Scenario
	}

	if !dara.IsNil(request.Sign) {
		query["Sign"] = request.Sign
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListScenario"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListScenarioResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of a silence policy list.
//
// @param request - ListSilencePoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSilencePoliciesResponse
func (client *Client) ListSilencePoliciesWithContext(ctx context.Context, request *ListSilencePoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListSilencePoliciesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsDetail) {
		query["IsDetail"] = request.IsDetail
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Size) {
		query["Size"] = request.Size
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSilencePolicies"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSilencePoliciesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the results of one or more synthetic tests.
//
// @param tmpReq - ListSyntheticDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSyntheticDetailResponse
func (client *Client) ListSyntheticDetailWithContext(ctx context.Context, tmpReq *ListSyntheticDetailRequest, runtime *dara.RuntimeOptions) (_result *ListSyntheticDetailResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListSyntheticDetailShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AdvancedFilters) {
		request.AdvancedFiltersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AdvancedFilters, dara.String("AdvancedFilters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ExactFilters) {
		request.ExactFiltersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExactFilters, dara.String("ExactFilters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Filters) {
		request.FiltersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filters, dara.String("Filters"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSyntheticDetail"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSyntheticDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries scheduled synthetic monitoring tasks.
//
// @param tmpReq - ListTimingSyntheticTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTimingSyntheticTasksResponse
func (client *Client) ListTimingSyntheticTasksWithContext(ctx context.Context, tmpReq *ListTimingSyntheticTasksRequest, runtime *dara.RuntimeOptions) (_result *ListTimingSyntheticTasksResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListTimingSyntheticTasksShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Search) {
		request.SearchShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Search, dara.String("Search"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTimingSyntheticTasks"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTimingSyntheticTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all Application Monitoring tasks in a specified region.
//
// @param request - ListTraceAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTraceAppsResponse
func (client *Client) ListTraceAppsWithContext(ctx context.Context, request *ListTraceAppsRequest, runtime *dara.RuntimeOptions) (_result *ListTraceAppsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		query["AppType"] = request.AppType
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTraceApps"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTraceAppsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Activates the service-linked role AliyunServiceRoleForARMS for Application Real-Time Monitoring Service (ARMS).
//
// @param request - OpenArmsDefaultSLRRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenArmsDefaultSLRResponse
func (client *Client) OpenArmsDefaultSLRWithContext(ctx context.Context, request *OpenArmsDefaultSLRRequest, runtime *dara.RuntimeOptions) (_result *OpenArmsDefaultSLRResponse, _err error) {
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
		Action:      dara.String("OpenArmsDefaultSLR"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenArmsDefaultSLRResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Activates a specified pay-as-you-go sub-service of Application Real-Time Monitoring Service (ARMS).
//
// Description:
//
// The **OpenArmsServiceSecondVersion*	- operation supports the following sub-service editions:
//
//   - Application Monitoring: Basic Edition
//
//   - Browser Monitoring: Basic Edition
//
//   - Synthetic Monitoring: Pro Edition (pay-as-you-go)
//
//   - Prometheus Service: Pro Edition
//
// @param request - OpenArmsServiceSecondVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenArmsServiceSecondVersionResponse
func (client *Client) OpenArmsServiceSecondVersionWithContext(ctx context.Context, request *OpenArmsServiceSecondVersionRequest, runtime *dara.RuntimeOptions) (_result *OpenArmsServiceSecondVersionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      dara.String("OpenArmsServiceSecondVersion"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenArmsServiceSecondVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI OpenVCluster is deprecated
//
// Summary:
//
// Activates a virtual cluster.
//
// @param request - OpenVClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenVClusterResponse
func (client *Client) OpenVClusterWithContext(ctx context.Context, request *OpenVClusterRequest, runtime *dara.RuntimeOptions) (_result *OpenVClusterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterType) {
		query["ClusterType"] = request.ClusterType
	}

	if !dara.IsNil(request.Length) {
		query["Length"] = request.Length
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	if !dara.IsNil(request.RecreateSwitch) {
		query["RecreateSwitch"] = request.RecreateSwitch
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OpenVCluster"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenVClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Activates the service-linked role AliyunServiceRoleForXtrace for Tracing Analysis.
//
// @param request - OpenXtraceDefaultSLRRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenXtraceDefaultSLRResponse
func (client *Client) OpenXtraceDefaultSLRWithContext(ctx context.Context, request *OpenXtraceDefaultSLRRequest, runtime *dara.RuntimeOptions) (_result *OpenXtraceDefaultSLRResponse, _err error) {
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
		Action:      dara.String("OpenXtraceDefaultSLR"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenXtraceDefaultSLRResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the encoding mapping content based on the metadata IDs and metadata type.
//
// @param request - QueryAppMetadataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAppMetadataResponse
func (client *Client) QueryAppMetadataWithContext(ctx context.Context, request *QueryAppMetadataRequest, runtime *dara.RuntimeOptions) (_result *QueryAppMetadataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAppMetadata"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAppMetadataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the topology of an application.
//
// @param tmpReq - QueryAppTopologyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAppTopologyResponse
func (client *Client) QueryAppTopologyWithContext(ctx context.Context, tmpReq *QueryAppTopologyRequest, runtime *dara.RuntimeOptions) (_result *QueryAppTopologyResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &QueryAppTopologyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Filters) {
		request.FiltersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filters, dara.String("Filters"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		query["AppType"] = request.AppType
	}

	if !dara.IsNil(request.Db) {
		query["Db"] = request.Db
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.FiltersShrink) {
		query["Filters"] = request.FiltersShrink
	}

	if !dara.IsNil(request.Pid) {
		query["Pid"] = request.Pid
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Rpc) {
		query["Rpc"] = request.Rpc
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
		Action:      dara.String("QueryAppTopology"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAppTopologyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the amount of data written to Application Monitoring, Managed Service for OpenTelemetry, Managed Service for Prometheus, and Real User Monitoring (RUM).
//
// @param request - QueryCommercialUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCommercialUsageResponse
func (client *Client) QueryCommercialUsageWithContext(ctx context.Context, request *QueryCommercialUsageRequest, runtime *dara.RuntimeOptions) (_result *QueryCommercialUsageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdvancedFilters) {
		query["AdvancedFilters"] = request.AdvancedFilters
	}

	if !dara.IsNil(request.Dimensions) {
		query["Dimensions"] = request.Dimensions
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IntervalInSec) {
		query["IntervalInSec"] = request.IntervalInSec
	}

	if !dara.IsNil(request.Measures) {
		query["Measures"] = request.Measures
	}

	if !dara.IsNil(request.Metric) {
		query["Metric"] = request.Metric
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.QueryType) {
		query["QueryType"] = request.QueryType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCommercialUsage"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCommercialUsageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries an Application Monitoring metric or a Browser Monitoring metric.
//
// @param request - QueryMetricByPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMetricByPageResponse
func (client *Client) QueryMetricByPageWithContext(ctx context.Context, request *QueryMetricByPageRequest, runtime *dara.RuntimeOptions) (_result *QueryMetricByPageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.CustomFilters) {
		query["CustomFilters"] = request.CustomFilters
	}

	if !dara.IsNil(request.Dimensions) {
		query["Dimensions"] = request.Dimensions
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Filters) {
		query["Filters"] = request.Filters
	}

	if !dara.IsNil(request.IntervalInSec) {
		query["IntervalInSec"] = request.IntervalInSec
	}

	if !dara.IsNil(request.Measures) {
		query["Measures"] = request.Measures
	}

	if !dara.IsNil(request.Metric) {
		query["Metric"] = request.Metric
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
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
		Action:      dara.String("QueryMetricByPage"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMetricByPageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI QueryPromInstallStatus is deprecated
//
// Summary:
//
// Queries whether the Prometheus agent is installed on a cluster.
//
// @param request - QueryPromInstallStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPromInstallStatusResponse
func (client *Client) QueryPromInstallStatusWithContext(ctx context.Context, request *QueryPromInstallStatusRequest, runtime *dara.RuntimeOptions) (_result *QueryPromInstallStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("QueryPromInstallStatus"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPromInstallStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the metrics that are provided for different versions of a specified Enterprise Distributed Application Service (EDAS) or Kubernetes application.
//
// @param request - QueryReleaseMetricRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryReleaseMetricResponse
func (client *Client) QueryReleaseMetricWithContext(ctx context.Context, request *QueryReleaseMetricRequest, runtime *dara.RuntimeOptions) (_result *QueryReleaseMetricResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChangeOrderId) {
		query["ChangeOrderId"] = request.ChangeOrderId
	}

	if !dara.IsNil(request.CreateTime) {
		query["CreateTime"] = request.CreateTime
	}

	if !dara.IsNil(request.MetricType) {
		query["MetricType"] = request.MetricType
	}

	if !dara.IsNil(request.Pid) {
		query["Pid"] = request.Pid
	}

	if !dara.IsNil(request.ProxyUserId) {
		query["ProxyUserId"] = request.ProxyUserId
	}

	if !dara.IsNil(request.ReleaseEndTime) {
		query["ReleaseEndTime"] = request.ReleaseEndTime
	}

	if !dara.IsNil(request.ReleaseStartTime) {
		query["ReleaseStartTime"] = request.ReleaseStartTime
	}

	if !dara.IsNil(request.Service) {
		query["Service"] = request.Service
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryReleaseMetric"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryReleaseMetricResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI RemoveAliClusterIdsFromPrometheusGlobalView is deprecated
//
// Summary:
//
// Removes data sources from a global aggregation instance in Managed Service for Prometheus.
//
// @param request - RemoveAliClusterIdsFromPrometheusGlobalViewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveAliClusterIdsFromPrometheusGlobalViewResponse
func (client *Client) RemoveAliClusterIdsFromPrometheusGlobalViewWithContext(ctx context.Context, request *RemoveAliClusterIdsFromPrometheusGlobalViewRequest, runtime *dara.RuntimeOptions) (_result *RemoveAliClusterIdsFromPrometheusGlobalViewResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterIds) {
		query["ClusterIds"] = request.ClusterIds
	}

	if !dara.IsNil(request.GlobalViewClusterId) {
		query["GlobalViewClusterId"] = request.GlobalViewClusterId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveAliClusterIdsFromPrometheusGlobalView"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveAliClusterIdsFromPrometheusGlobalViewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI RemoveSourcesFromPrometheusGlobalView is deprecated
//
// Summary:
//
// Removes data sources from a global aggregation instance in Managed Service for Prometheus. You can delete only data sources that are not from Alibaba Cloud.
//
// @param request - RemoveSourcesFromPrometheusGlobalViewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveSourcesFromPrometheusGlobalViewResponse
func (client *Client) RemoveSourcesFromPrometheusGlobalViewWithContext(ctx context.Context, request *RemoveSourcesFromPrometheusGlobalViewRequest, runtime *dara.RuntimeOptions) (_result *RemoveSourcesFromPrometheusGlobalViewResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GlobalViewClusterId) {
		query["GlobalViewClusterId"] = request.GlobalViewClusterId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SourceNames) {
		query["SourceNames"] = request.SourceNames
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveSourcesFromPrometheusGlobalView"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveSourcesFromPrometheusGlobalViewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts a feature.
//
// @param request - RestartEnvironmentFeatureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartEnvironmentFeatureResponse
func (client *Client) RestartEnvironmentFeatureWithContext(ctx context.Context, request *RestartEnvironmentFeatureRequest, runtime *dara.RuntimeOptions) (_result *RestartEnvironmentFeatureResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnvironmentId) {
		query["EnvironmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.FeatureName) {
		query["FeatureName"] = request.FeatureName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartEnvironmentFeature"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartEnvironmentFeatureResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the settings of Application Monitoring, such as trace sampling and agent switch settings.
//
// @param request - SaveTraceAppConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveTraceAppConfigResponse
func (client *Client) SaveTraceAppConfigWithContext(ctx context.Context, request *SaveTraceAppConfigRequest, runtime *dara.RuntimeOptions) (_result *SaveTraceAppConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Pid) {
		query["Pid"] = request.Pid
	}

	if !dara.IsNil(request.Settings) {
		query["Settings"] = request.Settings
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveTraceAppConfig"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveTraceAppConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries alert contacts.
//
// Description:
//
// This operation is no longer maintained. To query alert contacts, call the DescribeContacts operation provided by the new version of Alert Management.
//
// @param request - SearchAlertContactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchAlertContactResponse
func (client *Client) SearchAlertContactWithContext(ctx context.Context, request *SearchAlertContactRequest, runtime *dara.RuntimeOptions) (_result *SearchAlertContactResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactIds) {
		query["ContactIds"] = request.ContactIds
	}

	if !dara.IsNil(request.ContactName) {
		query["ContactName"] = request.ContactName
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Phone) {
		query["Phone"] = request.Phone
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchAlertContact"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchAlertContactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries alert contact groups.
//
// Description:
//
// The operation is no longer maintained. Call the DescribeContactGroups operation in the alert management module to query alert contact groups.
//
// @param request - SearchAlertContactGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchAlertContactGroupResponse
func (client *Client) SearchAlertContactGroupWithContext(ctx context.Context, request *SearchAlertContactGroupRequest, runtime *dara.RuntimeOptions) (_result *SearchAlertContactGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactGroupIds) {
		query["ContactGroupIds"] = request.ContactGroupIds
	}

	if !dara.IsNil(request.ContactGroupName) {
		query["ContactGroupName"] = request.ContactGroupName
	}

	if !dara.IsNil(request.ContactId) {
		query["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.ContactName) {
		query["ContactName"] = request.ContactName
	}

	if !dara.IsNil(request.IsDetail) {
		query["IsDetail"] = request.IsDetail
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchAlertContactGroup"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchAlertContactGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the alert records of an alert rule.
//
// Description:
//
// This operation is no longer maintained. To query alert records, call the ListAlerts operation provided by the new version of Alert Management.
//
// @param request - SearchAlertHistoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchAlertHistoriesResponse
func (client *Client) SearchAlertHistoriesWithContext(ctx context.Context, request *SearchAlertHistoriesRequest, runtime *dara.RuntimeOptions) (_result *SearchAlertHistoriesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertId) {
		query["AlertId"] = request.AlertId
	}

	if !dara.IsNil(request.AlertType) {
		query["AlertType"] = request.AlertType
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
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
		Action:      dara.String("SearchAlertHistories"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchAlertHistoriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries alert rules.
//
// Description:
//
// The current operation is no longer maintained. You can call the GetAlertRules operation of Alert Management (New) to query existing alert rules.
//
// @param request - SearchAlertRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchAlertRulesResponse
func (client *Client) SearchAlertRulesWithContext(ctx context.Context, request *SearchAlertRulesRequest, runtime *dara.RuntimeOptions) (_result *SearchAlertRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertRuleId) {
		query["AlertRuleId"] = request.AlertRuleId
	}

	if !dara.IsNil(request.AppType) {
		query["AppType"] = request.AppType
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Pid) {
		query["Pid"] = request.Pid
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SystemRegionId) {
		query["SystemRegionId"] = request.SystemRegionId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchAlertRules"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchAlertRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries alert event records.
//
// Description:
//
// Alert event records are different from alert notification records. Alert events are recorded every minute after an alert rule filters data. Alert events can be classified based on whether they are triggered or not. If a triggered event is not in the silence period, an alert notification is sent.
//
// @param request - SearchEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchEventsResponse
func (client *Client) SearchEventsWithContext(ctx context.Context, request *SearchEventsRequest, runtime *dara.RuntimeOptions) (_result *SearchEventsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertId) {
		query["AlertId"] = request.AlertId
	}

	if !dara.IsNil(request.AlertType) {
		query["AlertType"] = request.AlertType
	}

	if !dara.IsNil(request.AppType) {
		query["AppType"] = request.AppType
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IsTrigger) {
		query["IsTrigger"] = request.IsTrigger
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Pid) {
		query["Pid"] = request.Pid
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
		Action:      dara.String("SearchEvents"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchEventsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Browser Monitoring tasks by page.
//
// @param request - SearchRetcodeAppByPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchRetcodeAppByPageResponse
func (client *Client) SearchRetcodeAppByPageWithContext(ctx context.Context, request *SearchRetcodeAppByPageRequest, runtime *dara.RuntimeOptions) (_result *SearchRetcodeAppByPageResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.RetcodeAppId) {
		query["RetcodeAppId"] = request.RetcodeAppId
	}

	if !dara.IsNil(request.RetcodeAppName) {
		query["RetcodeAppName"] = request.RetcodeAppName
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchRetcodeAppByPage"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchRetcodeAppByPageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Application Monitoring tasks by application name.
//
// @param request - SearchTraceAppByNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchTraceAppByNameResponse
func (client *Client) SearchTraceAppByNameWithContext(ctx context.Context, request *SearchTraceAppByNameRequest, runtime *dara.RuntimeOptions) (_result *SearchTraceAppByNameResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.TraceAppName) {
		query["TraceAppName"] = request.TraceAppName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchTraceAppByName"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchTraceAppByNameResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries application monitoring tasks by page.
//
// @param request - SearchTraceAppByPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchTraceAppByPageResponse
func (client *Client) SearchTraceAppByPageWithContext(ctx context.Context, request *SearchTraceAppByPageRequest, runtime *dara.RuntimeOptions) (_result *SearchTraceAppByPageResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.TraceAppName) {
		query["TraceAppName"] = request.TraceAppName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchTraceAppByPage"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchTraceAppByPageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries traces by time, application name, IP address, span name, and tag.
//
// Description:
//
// > A maximum of 100 data entries can be returned each time this operation is called. If you want to query all existing traces, we recommend that you call the SearchTracesByPage operation. For more information, see [SearchTracesByPage](https://help.aliyun.com/document_detail/175866.html).
//
// @param request - SearchTracesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchTracesResponse
func (client *Client) SearchTracesWithContext(ctx context.Context, request *SearchTracesRequest, runtime *dara.RuntimeOptions) (_result *SearchTracesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ExclusionFilters) {
		query["ExclusionFilters"] = request.ExclusionFilters
	}

	if !dara.IsNil(request.MinDuration) {
		query["MinDuration"] = request.MinDuration
	}

	if !dara.IsNil(request.OperationName) {
		query["OperationName"] = request.OperationName
	}

	if !dara.IsNil(request.Pid) {
		query["Pid"] = request.Pid
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Reverse) {
		query["Reverse"] = request.Reverse
	}

	if !dara.IsNil(request.ServiceIp) {
		query["ServiceIp"] = request.ServiceIp
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchTraces"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchTracesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries traces by page. You can query traces by time range, application name, IP address, span name, or tag.
//
// @param request - SearchTracesByPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchTracesByPageResponse
func (client *Client) SearchTracesByPageWithContext(ctx context.Context, request *SearchTracesByPageRequest, runtime *dara.RuntimeOptions) (_result *SearchTracesByPageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ExclusionFilters) {
		query["ExclusionFilters"] = request.ExclusionFilters
	}

	if !dara.IsNil(request.IsError) {
		query["IsError"] = request.IsError
	}

	if !dara.IsNil(request.MinDuration) {
		query["MinDuration"] = request.MinDuration
	}

	if !dara.IsNil(request.OperationName) {
		query["OperationName"] = request.OperationName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Pid) {
		query["Pid"] = request.Pid
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Reverse) {
		query["Reverse"] = request.Reverse
	}

	if !dara.IsNil(request.ServiceIp) {
		query["ServiceIp"] = request.ServiceIp
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchTracesByPage"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchTracesByPageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a text message to an alert contact to verify the mobile number of the alert contact.
//
// Description:
//
// After you receive the text message, verify the mobile number as prompted. Before you can specify a mobile phone number in a notification policy, you must verify the mobile phone number.
//
// @param request - SendTTSVerifyLinkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendTTSVerifyLinkResponse
func (client *Client) SendTTSVerifyLinkWithContext(ctx context.Context, request *SendTTSVerifyLinkRequest, runtime *dara.RuntimeOptions) (_result *SendTTSVerifyLinkResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ContactId) {
		body["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.Phone) {
		body["Phone"] = request.Phone
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendTTSVerifyLink"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendTTSVerifyLinkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Turns on or turns off logon-free sharing for an application monitored by Browser Monitoring.
//
// @param request - SetRetcodeShareStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetRetcodeShareStatusResponse
func (client *Client) SetRetcodeShareStatusWithContext(ctx context.Context, request *SetRetcodeShareStatusRequest, runtime *dara.RuntimeOptions) (_result *SetRetcodeShareStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Pid) {
		query["Pid"] = request.Pid
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetRetcodeShareStatus"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetRetcodeShareStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - StartAlertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartAlertResponse
func (client *Client) StartAlertWithContext(ctx context.Context, request *StartAlertRequest, runtime *dara.RuntimeOptions) (_result *StartAlertResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertId) {
		query["AlertId"] = request.AlertId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartAlert"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartAlertResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts scheduled synthetic monitoring tasks.
//
// @param tmpReq - StartTimingSyntheticTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartTimingSyntheticTaskResponse
func (client *Client) StartTimingSyntheticTaskWithContext(ctx context.Context, tmpReq *StartTimingSyntheticTaskRequest, runtime *dara.RuntimeOptions) (_result *StartTimingSyntheticTaskResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &StartTimingSyntheticTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TaskIds) {
		request.TaskIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskIds, dara.String("TaskIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TaskIdsShrink) {
		query["TaskIds"] = request.TaskIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartTimingSyntheticTask"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartTimingSyntheticTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call StartAlert to stop an alert rule.
//
// @param request - StopAlertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopAlertResponse
func (client *Client) StopAlertWithContext(ctx context.Context, request *StopAlertRequest, runtime *dara.RuntimeOptions) (_result *StopAlertResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertId) {
		query["AlertId"] = request.AlertId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopAlert"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopAlertResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops scheduled synthetic monitoring tasks.
//
// @param tmpReq - StopTimingSyntheticTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopTimingSyntheticTaskResponse
func (client *Client) StopTimingSyntheticTaskWithContext(ctx context.Context, tmpReq *StopTimingSyntheticTaskRequest, runtime *dara.RuntimeOptions) (_result *StopTimingSyntheticTaskResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &StopTimingSyntheticTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TaskIds) {
		request.TaskIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskIds, dara.String("TaskIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TaskIdsShrink) {
		query["TaskIds"] = request.TaskIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopTimingSyntheticTask"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopTimingSyntheticTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts or stops a scheduled synthetic monitoring task.
//
// @param request - SwitchSyntheticTaskStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchSyntheticTaskStatusResponse
func (client *Client) SwitchSyntheticTaskStatusWithContext(ctx context.Context, request *SwitchSyntheticTaskStatusRequest, runtime *dara.RuntimeOptions) (_result *SwitchSyntheticTaskStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SwitchStatus) {
		query["SwitchStatus"] = request.SwitchStatus
	}

	if !dara.IsNil(request.TaskIds) {
		query["TaskIds"] = request.TaskIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchSyntheticTaskStatus"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchSyntheticTaskStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI SyncRecordingRules is deprecated
//
// Summary:
//
// Synchronizes the aggregation rule of a cluster to other clusters in a region.
//
// @param request - SyncRecordingRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncRecordingRulesResponse
func (client *Client) SyncRecordingRulesWithContext(ctx context.Context, request *SyncRecordingRulesRequest, runtime *dara.RuntimeOptions) (_result *SyncRecordingRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TargetClusters) {
		query["TargetClusters"] = request.TargetClusters
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SyncRecordingRules"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SyncRecordingRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds tags to ARMS resources.
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithContext(ctx context.Context, request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI UninstallManagedPrometheus is deprecated
//
// Summary:
//
// Uninstalls a managed Prometheus agent for a serverless Kubernetes (ASK) cluster, Distributed Cloud Container Platform for Kubernetes (ACK One) cluster, or Elastic Compute Service (ECS) cluster.
//
// Description:
//
// This operation is available only for ASK, ECS, and ACK One clusters. Before you call this operation, make sure that a managed Prometheus agent is installed for your cluster.
//
// @param request - UninstallManagedPrometheusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UninstallManagedPrometheusResponse
func (client *Client) UninstallManagedPrometheusWithContext(ctx context.Context, request *UninstallManagedPrometheusRequest, runtime *dara.RuntimeOptions) (_result *UninstallManagedPrometheusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ClusterType) {
		query["ClusterType"] = request.ClusterType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UninstallManagedPrometheus"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UninstallManagedPrometheusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases a Prometheus instance.
//
// @param request - UninstallPromClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UninstallPromClusterResponse
func (client *Client) UninstallPromClusterWithContext(ctx context.Context, request *UninstallPromClusterRequest, runtime *dara.RuntimeOptions) (_result *UninstallPromClusterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunLang) {
		query["AliyunLang"] = request.AliyunLang
	}

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
		Action:      dara.String("UninstallPromCluster"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UninstallPromClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes tags from ARMS resources.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithContext(ctx context.Context, request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
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

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an alert contact.
//
// Description:
//
// This operation is no longer maintained. To create or modify an alert contact, call the CreateOrUpdateContact operation provided by the new version of Alert Management.
//
// @param request - UpdateAlertContactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAlertContactResponse
func (client *Client) UpdateAlertContactWithContext(ctx context.Context, request *UpdateAlertContactRequest, runtime *dara.RuntimeOptions) (_result *UpdateAlertContactResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactId) {
		query["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.ContactName) {
		query["ContactName"] = request.ContactName
	}

	if !dara.IsNil(request.DingRobotWebhookUrl) {
		query["DingRobotWebhookUrl"] = request.DingRobotWebhookUrl
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.PhoneNum) {
		query["PhoneNum"] = request.PhoneNum
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SystemNoc) {
		query["SystemNoc"] = request.SystemNoc
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAlertContact"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAlertContactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates UpdateAlertContactGroup alarm contact group.
//
// @param request - UpdateAlertContactGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAlertContactGroupResponse
func (client *Client) UpdateAlertContactGroupWithContext(ctx context.Context, request *UpdateAlertContactGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateAlertContactGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactGroupId) {
		query["ContactGroupId"] = request.ContactGroupId
	}

	if !dara.IsNil(request.ContactGroupName) {
		query["ContactGroupName"] = request.ContactGroupName
	}

	if !dara.IsNil(request.ContactIds) {
		query["ContactIds"] = request.ContactIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAlertContactGroup"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAlertContactGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateAlertRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAlertRuleResponse
func (client *Client) UpdateAlertRuleWithContext(ctx context.Context, request *UpdateAlertRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateAlertRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertId) {
		query["AlertId"] = request.AlertId
	}

	if !dara.IsNil(request.ContactGroupIds) {
		query["ContactGroupIds"] = request.ContactGroupIds
	}

	if !dara.IsNil(request.IsAutoStart) {
		query["IsAutoStart"] = request.IsAutoStart
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TemplageAlertConfig) {
		query["TemplageAlertConfig"] = request.TemplageAlertConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAlertRule"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAlertRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a dispatch policy.
//
// @param request - UpdateDispatchRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDispatchRuleResponse
func (client *Client) UpdateDispatchRuleWithContext(ctx context.Context, request *UpdateDispatchRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateDispatchRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DispatchRule) {
		query["DispatchRule"] = request.DispatchRule
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDispatchRule"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDispatchRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a custom job for an environment.
//
// @param request - UpdateEnvCustomJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEnvCustomJobResponse
func (client *Client) UpdateEnvCustomJobWithContext(ctx context.Context, request *UpdateEnvCustomJobRequest, runtime *dara.RuntimeOptions) (_result *UpdateEnvCustomJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunLang) {
		query["AliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.CustomJobName) {
		query["CustomJobName"] = request.CustomJobName
	}

	if !dara.IsNil(request.EnvironmentId) {
		query["EnvironmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigYaml) {
		body["ConfigYaml"] = request.ConfigYaml
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEnvCustomJob"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEnvCustomJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新环境实例的废弃指标列表
//
// @param request - UpdateEnvDropMetricsRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEnvDropMetricsRuleResponse
func (client *Client) UpdateEnvDropMetricsRuleWithContext(ctx context.Context, request *UpdateEnvDropMetricsRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateEnvDropMetricsRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunLang) {
		query["AliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.EnvironmentId) {
		query["EnvironmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DropMetrics) {
		body["DropMetrics"] = request.DropMetrics
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEnvDropMetricsRule"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEnvDropMetricsRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the PodMonitor of an environment.
//
// @param request - UpdateEnvPodMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEnvPodMonitorResponse
func (client *Client) UpdateEnvPodMonitorWithContext(ctx context.Context, request *UpdateEnvPodMonitorRequest, runtime *dara.RuntimeOptions) (_result *UpdateEnvPodMonitorResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunLang) {
		query["AliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EnvironmentId) {
		query["EnvironmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PodMonitorName) {
		query["PodMonitorName"] = request.PodMonitorName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigYaml) {
		body["ConfigYaml"] = request.ConfigYaml
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEnvPodMonitor"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEnvPodMonitorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the ServiceMonitor of an environment.
//
// @param request - UpdateEnvServiceMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEnvServiceMonitorResponse
func (client *Client) UpdateEnvServiceMonitorWithContext(ctx context.Context, request *UpdateEnvServiceMonitorRequest, runtime *dara.RuntimeOptions) (_result *UpdateEnvServiceMonitorResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunLang) {
		query["AliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EnvironmentId) {
		query["EnvironmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceMonitorName) {
		query["ServiceMonitorName"] = request.ServiceMonitorName
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigYaml) {
		body["ConfigYaml"] = request.ConfigYaml
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEnvServiceMonitor"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEnvServiceMonitorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration of an environment.
//
// @param request - UpdateEnvironmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEnvironmentResponse
func (client *Client) UpdateEnvironmentWithContext(ctx context.Context, request *UpdateEnvironmentRequest, runtime *dara.RuntimeOptions) (_result *UpdateEnvironmentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunLang) {
		query["AliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.EnvironmentId) {
		query["EnvironmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.EnvironmentName) {
		query["EnvironmentName"] = request.EnvironmentName
	}

	if !dara.IsNil(request.FeePackage) {
		query["FeePackage"] = request.FeePackage
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEnvironment"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEnvironmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a Grafana workspace.
//
// @param request - UpdateGrafanaWorkspaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGrafanaWorkspaceResponse
func (client *Client) UpdateGrafanaWorkspaceWithContext(ctx context.Context, request *UpdateGrafanaWorkspaceRequest, runtime *dara.RuntimeOptions) (_result *UpdateGrafanaWorkspaceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunLang) {
		query["AliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GrafanaWorkspaceId) {
		query["GrafanaWorkspaceId"] = request.GrafanaWorkspaceId
	}

	if !dara.IsNil(request.GrafanaWorkspaceName) {
		query["GrafanaWorkspaceName"] = request.GrafanaWorkspaceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGrafanaWorkspace"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGrafanaWorkspaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the version of a Grafana workspace.
//
// Description:
//
// Note: The list returned by this operation includes the workspaces of Developer Edition, Expert Edition, and Advanced Edition. The list does not include the workspaces of Shared Edition.
//
// @param request - UpdateGrafanaWorkspaceVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGrafanaWorkspaceVersionResponse
func (client *Client) UpdateGrafanaWorkspaceVersionWithContext(ctx context.Context, request *UpdateGrafanaWorkspaceVersionRequest, runtime *dara.RuntimeOptions) (_result *UpdateGrafanaWorkspaceVersionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunLang) {
		query["AliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.GrafanaVersion) {
		query["GrafanaVersion"] = request.GrafanaVersion
	}

	if !dara.IsNil(request.GrafanaWorkspaceId) {
		query["GrafanaWorkspaceId"] = request.GrafanaWorkspaceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGrafanaWorkspaceVersion"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGrafanaWorkspaceVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about an integration.
//
// @param request - UpdateIntegrationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIntegrationResponse
func (client *Client) UpdateIntegrationWithContext(ctx context.Context, request *UpdateIntegrationRequest, runtime *dara.RuntimeOptions) (_result *UpdateIntegrationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoRecover) {
		body["AutoRecover"] = request.AutoRecover
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DuplicateKey) {
		body["DuplicateKey"] = request.DuplicateKey
	}

	if !dara.IsNil(request.ExtendedFieldRedefineRules) {
		body["ExtendedFieldRedefineRules"] = request.ExtendedFieldRedefineRules
	}

	if !dara.IsNil(request.FieldRedefineRules) {
		body["FieldRedefineRules"] = request.FieldRedefineRules
	}

	if !dara.IsNil(request.InitiativeRecoverField) {
		body["InitiativeRecoverField"] = request.InitiativeRecoverField
	}

	if !dara.IsNil(request.InitiativeRecoverValue) {
		body["InitiativeRecoverValue"] = request.InitiativeRecoverValue
	}

	if !dara.IsNil(request.IntegrationId) {
		body["IntegrationId"] = request.IntegrationId
	}

	if !dara.IsNil(request.IntegrationName) {
		body["IntegrationName"] = request.IntegrationName
	}

	if !dara.IsNil(request.IntegrationProductType) {
		body["IntegrationProductType"] = request.IntegrationProductType
	}

	if !dara.IsNil(request.Liveness) {
		body["Liveness"] = request.Liveness
	}

	if !dara.IsNil(request.RecoverTime) {
		body["RecoverTime"] = request.RecoverTime
	}

	if !dara.IsNil(request.Stat) {
		body["Stat"] = request.Stat
	}

	if !dara.IsNil(request.State) {
		body["State"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateIntegration"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateIntegrationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI UpdateMetricDrop is deprecated, please use ARMS::2019-08-08::UpdateEnvDropMetricsRule instead.
//
// Summary:
//
// Updates the list of discarded metrics.
//
// @param request - UpdateMetricDropRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMetricDropResponse
func (client *Client) UpdateMetricDropWithContext(ctx context.Context, request *UpdateMetricDropRequest, runtime *dara.RuntimeOptions) (_result *UpdateMetricDropResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.MetricDrop) {
		query["MetricDrop"] = request.MetricDrop
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMetricDrop"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMetricDropResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新Prometheus告警规则
//
// @param request - UpdatePrometheusAlertRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePrometheusAlertRuleResponse
func (client *Client) UpdatePrometheusAlertRuleWithContext(ctx context.Context, request *UpdatePrometheusAlertRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdatePrometheusAlertRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertId) {
		query["AlertId"] = request.AlertId
	}

	if !dara.IsNil(request.AlertName) {
		query["AlertName"] = request.AlertName
	}

	if !dara.IsNil(request.Annotations) {
		query["Annotations"] = request.Annotations
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DispatchRuleId) {
		query["DispatchRuleId"] = request.DispatchRuleId
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.Expression) {
		query["Expression"] = request.Expression
	}

	if !dara.IsNil(request.Labels) {
		query["Labels"] = request.Labels
	}

	if !dara.IsNil(request.Message) {
		query["Message"] = request.Message
	}

	if !dara.IsNil(request.NotifyType) {
		query["NotifyType"] = request.NotifyType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePrometheusAlertRule"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePrometheusAlertRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI UpdatePrometheusGlobalView is deprecated
//
// Summary:
//
// Updates the data sources of Prometheus instance for GlobalView.
//
// @param request - UpdatePrometheusGlobalViewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePrometheusGlobalViewResponse
func (client *Client) UpdatePrometheusGlobalViewWithContext(ctx context.Context, request *UpdatePrometheusGlobalViewRequest, runtime *dara.RuntimeOptions) (_result *UpdatePrometheusGlobalViewResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllSubClustersSuccess) {
		query["AllSubClustersSuccess"] = request.AllSubClustersSuccess
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.MostRegionId) {
		query["MostRegionId"] = request.MostRegionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SubClustersJson) {
		query["SubClustersJson"] = request.SubClustersJson
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePrometheusGlobalView"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePrometheusGlobalViewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a Prometheus instance.
//
// @param request - UpdatePrometheusInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePrometheusInstanceResponse
func (client *Client) UpdatePrometheusInstanceWithContext(ctx context.Context, request *UpdatePrometheusInstanceRequest, runtime *dara.RuntimeOptions) (_result *UpdatePrometheusInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ArchiveDuration) {
		query["ArchiveDuration"] = request.ArchiveDuration
	}

	if !dara.IsNil(request.AuthFreeReadPolicy) {
		query["AuthFreeReadPolicy"] = request.AuthFreeReadPolicy
	}

	if !dara.IsNil(request.AuthFreeWritePolicy) {
		query["AuthFreeWritePolicy"] = request.AuthFreeWritePolicy
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.EnableAuthFreeRead) {
		query["EnableAuthFreeRead"] = request.EnableAuthFreeRead
	}

	if !dara.IsNil(request.EnableAuthFreeWrite) {
		query["EnableAuthFreeWrite"] = request.EnableAuthFreeWrite
	}

	if !dara.IsNil(request.EnableAuthToken) {
		query["EnableAuthToken"] = request.EnableAuthToken
	}

	if !dara.IsNil(request.PaymentType) {
		query["PaymentType"] = request.PaymentType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StorageDuration) {
		query["StorageDuration"] = request.StorageDuration
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePrometheusInstance"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePrometheusInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI UpdatePrometheusIntegration is deprecated
//
// Summary:
//
// Modifies the configurations of an exporter that is integrated into a Prometheus instance for Container Service or a Prometheus instance for ECS.
//
// @param request - UpdatePrometheusIntegrationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePrometheusIntegrationResponse
func (client *Client) UpdatePrometheusIntegrationWithContext(ctx context.Context, request *UpdatePrometheusIntegrationRequest, runtime *dara.RuntimeOptions) (_result *UpdatePrometheusIntegrationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IntegrationType) {
		query["IntegrationType"] = request.IntegrationType
	}

	if !dara.IsNil(request.Param) {
		query["Param"] = request.Param
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePrometheusIntegration"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePrometheusIntegrationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI UpdatePrometheusMonitoring is deprecated
//
// Summary:
//
// Updates the monitoring configuration of a Prometheus instance.
//
// @param request - UpdatePrometheusMonitoringRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePrometheusMonitoringResponse
func (client *Client) UpdatePrometheusMonitoringWithContext(ctx context.Context, request *UpdatePrometheusMonitoringRequest, runtime *dara.RuntimeOptions) (_result *UpdatePrometheusMonitoringResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.MonitoringName) {
		query["MonitoringName"] = request.MonitoringName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigYaml) {
		body["ConfigYaml"] = request.ConfigYaml
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePrometheusMonitoring"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePrometheusMonitoringResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI UpdatePrometheusMonitoringStatus is deprecated
//
// Summary:
//
// Updates the status of the monitoring configuration of a Prometheus instance.
//
// @param request - UpdatePrometheusMonitoringStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePrometheusMonitoringStatusResponse
func (client *Client) UpdatePrometheusMonitoringStatusWithContext(ctx context.Context, request *UpdatePrometheusMonitoringStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdatePrometheusMonitoringStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.MonitoringName) {
		query["MonitoringName"] = request.MonitoringName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePrometheusMonitoringStatus"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePrometheusMonitoringStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a Real User Monitoring (RUM) application.
//
// Description:
//
// Real User Monitoring (RUM) is available only in the China (Hangzhou), Singapore, and US (Silicon Valley) regions. Select the correct endpoint.
//
// @param request - UpdateRumAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRumAppResponse
func (client *Client) UpdateRumAppWithContext(ctx context.Context, request *UpdateRumAppRequest, runtime *dara.RuntimeOptions) (_result *UpdateRumAppResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppConfig) {
		query["AppConfig"] = request.AppConfig
	}

	if !dara.IsNil(request.AutoRestart) {
		query["AutoRestart"] = request.AutoRestart
	}

	if !dara.IsNil(request.BackendServiceTraceRegion) {
		query["BackendServiceTraceRegion"] = request.BackendServiceTraceRegion
	}

	if !dara.IsNil(request.BonreeSDKConfigJson) {
		query["BonreeSDKConfigJson"] = request.BonreeSDKConfigJson
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.IsSubscribe) {
		query["IsSubscribe"] = request.IsSubscribe
	}

	if !dara.IsNil(request.Nickname) {
		query["Nickname"] = request.Nickname
	}

	if !dara.IsNil(request.Pid) {
		query["Pid"] = request.Pid
	}

	if !dara.IsNil(request.RealRegionId) {
		query["RealRegionId"] = request.RealRegionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Restart) {
		query["Restart"] = request.Restart
	}

	if !dara.IsNil(request.ServiceDomainOperationJson) {
		query["ServiceDomainOperationJson"] = request.ServiceDomainOperationJson
	}

	if !dara.IsNil(request.Stop) {
		query["Stop"] = request.Stop
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRumApp"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRumAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the status of a Real User Monitoring (RUM) file. You can call this operation after the RUM file is uploaded.
//
// Description:
//
// Real User Monitoring (RUM) is available only in the China (Hangzhou), Singapore, and US (Silicon Valley) regions. Select the correct endpoint.
//
// @param request - UpdateRumFileStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRumFileStatusResponse
func (client *Client) UpdateRumFileStatusWithContext(ctx context.Context, request *UpdateRumFileStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateRumFileStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.Pid) {
		query["Pid"] = request.Pid
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Size) {
		query["Size"] = request.Size
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Uuid) {
		query["Uuid"] = request.Uuid
	}

	if !dara.IsNil(request.VersionId) {
		query["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRumFileStatus"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRumFileStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a scheduled synthetic test task.
//
// @param tmpReq - UpdateTimingSyntheticTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTimingSyntheticTaskResponse
func (client *Client) UpdateTimingSyntheticTaskWithContext(ctx context.Context, tmpReq *UpdateTimingSyntheticTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateTimingSyntheticTaskResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateTimingSyntheticTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AvailableAssertions) {
		request.AvailableAssertionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AvailableAssertions, dara.String("AvailableAssertions"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CommonSetting) {
		request.CommonSettingShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CommonSetting, dara.String("CommonSetting"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CustomPeriod) {
		request.CustomPeriodShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomPeriod, dara.String("CustomPeriod"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.MonitorConf) {
		request.MonitorConfShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MonitorConf, dara.String("MonitorConf"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Monitors) {
		request.MonitorsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Monitors, dara.String("Monitors"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AvailableAssertionsShrink) {
		query["AvailableAssertions"] = request.AvailableAssertionsShrink
	}

	if !dara.IsNil(request.CommonSettingShrink) {
		query["CommonSetting"] = request.CommonSettingShrink
	}

	if !dara.IsNil(request.CustomPeriodShrink) {
		query["CustomPeriod"] = request.CustomPeriodShrink
	}

	if !dara.IsNil(request.Frequency) {
		query["Frequency"] = request.Frequency
	}

	if !dara.IsNil(request.MonitorConfShrink) {
		query["MonitorConf"] = request.MonitorConfShrink
	}

	if !dara.IsNil(request.MonitorsShrink) {
		query["Monitors"] = request.MonitorsShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTimingSyntheticTask"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTimingSyntheticTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about a webhook alert contact.
//
// Description:
//
// This operation is no longer maintained. Call the CreateOrUpdateWebhookContact operation in the new alter management module to create or modify a webhook alert contact.
//
// @param request - UpdateWebhookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWebhookResponse
func (client *Client) UpdateWebhookWithContext(ctx context.Context, request *UpdateWebhookRequest, runtime *dara.RuntimeOptions) (_result *UpdateWebhookResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Body) {
		query["Body"] = request.Body
	}

	if !dara.IsNil(request.ContactId) {
		query["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.ContactName) {
		query["ContactName"] = request.ContactName
	}

	if !dara.IsNil(request.HttpHeaders) {
		query["HttpHeaders"] = request.HttpHeaders
	}

	if !dara.IsNil(request.HttpParams) {
		query["HttpParams"] = request.HttpParams
	}

	if !dara.IsNil(request.Method) {
		query["Method"] = request.Method
	}

	if !dara.IsNil(request.RecoverBody) {
		query["RecoverBody"] = request.RecoverBody
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Url) {
		query["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWebhook"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWebhookResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the release information of an add-on.
//
// @param request - UpgradeAddonReleaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeAddonReleaseResponse
func (client *Client) UpgradeAddonReleaseWithContext(ctx context.Context, request *UpgradeAddonReleaseRequest, runtime *dara.RuntimeOptions) (_result *UpgradeAddonReleaseResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddonVersion) {
		query["AddonVersion"] = request.AddonVersion
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EnvironmentId) {
		query["EnvironmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReleaseName) {
		query["ReleaseName"] = request.ReleaseName
	}

	if !dara.IsNil(request.Values) {
		query["Values"] = request.Values
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeAddonRelease"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeAddonReleaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the feature information of an environment.
//
// @param request - UpgradeEnvironmentFeatureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeEnvironmentFeatureResponse
func (client *Client) UpgradeEnvironmentFeatureWithContext(ctx context.Context, request *UpgradeEnvironmentFeatureRequest, runtime *dara.RuntimeOptions) (_result *UpgradeEnvironmentFeatureResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunLang) {
		query["AliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.EnvironmentId) {
		query["EnvironmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.FeatureName) {
		query["FeatureName"] = request.FeatureName
	}

	if !dara.IsNil(request.FeatureVersion) {
		query["FeatureVersion"] = request.FeatureVersion
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Values) {
		query["Values"] = request.Values
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeEnvironmentFeature"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeEnvironmentFeatureResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads a SourceMap file to ARMS Browser Monitoring.
//
// @param request - UploadRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadResponse
func (client *Client) UploadWithContext(ctx context.Context, request *UploadRequest, runtime *dara.RuntimeOptions) (_result *UploadResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Edition) {
		query["Edition"] = request.Edition
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.Pid) {
		query["Pid"] = request.Pid
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.File) {
		body["File"] = request.File
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Upload"),
		Version:     dara.String("2019-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
