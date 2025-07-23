// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Terminates a change order and rolls back the corresponding application.
//
// @param request - AbortAndRollbackChangeOrderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AbortAndRollbackChangeOrderResponse
func (client *Client) AbortAndRollbackChangeOrderWithContext(ctx context.Context, request *AbortAndRollbackChangeOrderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AbortAndRollbackChangeOrderResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChangeOrderId) {
		query["ChangeOrderId"] = request.ChangeOrderId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AbortAndRollbackChangeOrder"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/changeorder/AbortAndRollbackChangeOrder"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AbortAndRollbackChangeOrderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Terminate a change order.
//
// @param request - AbortChangeOrderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AbortChangeOrderResponse
func (client *Client) AbortChangeOrderWithContext(ctx context.Context, request *AbortChangeOrderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AbortChangeOrderResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChangeOrderId) {
		query["ChangeOrderId"] = request.ChangeOrderId
	}

	if !dara.IsNil(request.Rollback) {
		query["Rollback"] = request.Rollback
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AbortChangeOrder"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/changeorder/AbortChangeOrder"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AbortChangeOrderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts multiple applications at a time.
//
// @param request - BatchStartApplicationsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchStartApplicationsResponse
func (client *Client) BatchStartApplicationsWithContext(ctx context.Context, request *BatchStartApplicationsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchStartApplicationsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppIds) {
		query["AppIds"] = request.AppIds
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchStartApplications"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/app/batchStartApplications"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchStartApplicationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stop applications in batches.
//
// @param request - BatchStopApplicationsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchStopApplicationsResponse
func (client *Client) BatchStopApplicationsWithContext(ctx context.Context, request *BatchStopApplicationsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchStopApplicationsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppIds) {
		query["AppIds"] = request.AppIds
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchStopApplications"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/app/batchStopApplications"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchStopApplicationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a Network Load Balancer (NLB) instance with an application.
//
// @param request - BindNlbRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindNlbResponse
func (client *Client) BindNlbWithContext(ctx context.Context, request *BindNlbRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BindNlbResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddressType) {
		query["AddressType"] = request.AddressType
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Listeners) {
		query["Listeners"] = request.Listeners
	}

	if !dara.IsNil(request.NlbId) {
		query["NlbId"] = request.NlbId
	}

	if !dara.IsNil(request.ZoneMappings) {
		query["ZoneMappings"] = request.ZoneMappings
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindNlb"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/app/nlb"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BindNlbResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - BindSlbRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindSlbResponse
func (client *Client) BindSlbWithContext(ctx context.Context, request *BindSlbRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BindSlbResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Internet) {
		query["Internet"] = request.Internet
	}

	if !dara.IsNil(request.InternetSlbChargeType) {
		query["InternetSlbChargeType"] = request.InternetSlbChargeType
	}

	if !dara.IsNil(request.InternetSlbId) {
		query["InternetSlbId"] = request.InternetSlbId
	}

	if !dara.IsNil(request.Intranet) {
		query["Intranet"] = request.Intranet
	}

	if !dara.IsNil(request.IntranetSlbChargeType) {
		query["IntranetSlbChargeType"] = request.IntranetSlbChargeType
	}

	if !dara.IsNil(request.IntranetSlbId) {
		query["IntranetSlbId"] = request.IntranetSlbId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindSlb"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/app/slb"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BindSlbResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Confirms whether to start the next batch.
//
// @param request - ConfirmPipelineBatchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfirmPipelineBatchResponse
func (client *Client) ConfirmPipelineBatchWithContext(ctx context.Context, request *ConfirmPipelineBatchRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ConfirmPipelineBatchResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Confirm) {
		query["Confirm"] = request.Confirm
	}

	if !dara.IsNil(request.PipelineId) {
		query["PipelineId"] = request.PipelineId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfirmPipelineBatch"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/changeorder/ConfirmPipelineBatch"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfirmPipelineBatchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an application.
//
// @param tmpReq - CreateApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApplicationResponse
func (client *Client) CreateApplicationWithContext(ctx context.Context, tmpReq *CreateApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateApplicationResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateApplicationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InitContainersConfig) {
		request.InitContainersConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InitContainersConfig, dara.String("InitContainersConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SidecarContainersConfig) {
		request.SidecarContainersConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SidecarContainersConfig, dara.String("SidecarContainersConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcrAssumeRoleArn) {
		query["AcrAssumeRoleArn"] = request.AcrAssumeRoleArn
	}

	if !dara.IsNil(request.AppDescription) {
		query["AppDescription"] = request.AppDescription
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.AppSource) {
		query["AppSource"] = request.AppSource
	}

	if !dara.IsNil(request.AutoConfig) {
		query["AutoConfig"] = request.AutoConfig
	}

	if !dara.IsNil(request.Command) {
		query["Command"] = request.Command
	}

	if !dara.IsNil(request.CommandArgs) {
		query["CommandArgs"] = request.CommandArgs
	}

	if !dara.IsNil(request.Cpu) {
		query["Cpu"] = request.Cpu
	}

	if !dara.IsNil(request.CustomHostAlias) {
		query["CustomHostAlias"] = request.CustomHostAlias
	}

	if !dara.IsNil(request.CustomImageNetworkType) {
		query["CustomImageNetworkType"] = request.CustomImageNetworkType
	}

	if !dara.IsNil(request.Deploy) {
		query["Deploy"] = request.Deploy
	}

	if !dara.IsNil(request.DiskSize) {
		query["DiskSize"] = request.DiskSize
	}

	if !dara.IsNil(request.Dotnet) {
		query["Dotnet"] = request.Dotnet
	}

	if !dara.IsNil(request.EdasContainerVersion) {
		query["EdasContainerVersion"] = request.EdasContainerVersion
	}

	if !dara.IsNil(request.EnableCpuBurst) {
		query["EnableCpuBurst"] = request.EnableCpuBurst
	}

	if !dara.IsNil(request.EnableEbpf) {
		query["EnableEbpf"] = request.EnableEbpf
	}

	if !dara.IsNil(request.EnableNewArms) {
		query["EnableNewArms"] = request.EnableNewArms
	}

	if !dara.IsNil(request.EnablePrometheus) {
		query["EnablePrometheus"] = request.EnablePrometheus
	}

	if !dara.IsNil(request.Envs) {
		query["Envs"] = request.Envs
	}

	if !dara.IsNil(request.GpuConfig) {
		query["GpuConfig"] = request.GpuConfig
	}

	if !dara.IsNil(request.ImagePullSecrets) {
		query["ImagePullSecrets"] = request.ImagePullSecrets
	}

	if !dara.IsNil(request.ImageUrl) {
		query["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.IsStateful) {
		query["IsStateful"] = request.IsStateful
	}

	if !dara.IsNil(request.JarStartArgs) {
		query["JarStartArgs"] = request.JarStartArgs
	}

	if !dara.IsNil(request.JarStartOptions) {
		query["JarStartOptions"] = request.JarStartOptions
	}

	if !dara.IsNil(request.Jdk) {
		query["Jdk"] = request.Jdk
	}

	if !dara.IsNil(request.KafkaConfigs) {
		query["KafkaConfigs"] = request.KafkaConfigs
	}

	if !dara.IsNil(request.Liveness) {
		query["Liveness"] = request.Liveness
	}

	if !dara.IsNil(request.Memory) {
		query["Memory"] = request.Memory
	}

	if !dara.IsNil(request.MicroRegistration) {
		query["MicroRegistration"] = request.MicroRegistration
	}

	if !dara.IsNil(request.MicroserviceEngineConfig) {
		query["MicroserviceEngineConfig"] = request.MicroserviceEngineConfig
	}

	if !dara.IsNil(request.MountDesc) {
		query["MountDesc"] = request.MountDesc
	}

	if !dara.IsNil(request.MountHost) {
		query["MountHost"] = request.MountHost
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.NasConfigs) {
		query["NasConfigs"] = request.NasConfigs
	}

	if !dara.IsNil(request.NasId) {
		query["NasId"] = request.NasId
	}

	if !dara.IsNil(request.NewSaeVersion) {
		query["NewSaeVersion"] = request.NewSaeVersion
	}

	if !dara.IsNil(request.OidcRoleName) {
		query["OidcRoleName"] = request.OidcRoleName
	}

	if !dara.IsNil(request.PackageType) {
		query["PackageType"] = request.PackageType
	}

	if !dara.IsNil(request.PackageUrl) {
		query["PackageUrl"] = request.PackageUrl
	}

	if !dara.IsNil(request.PackageVersion) {
		query["PackageVersion"] = request.PackageVersion
	}

	if !dara.IsNil(request.PhpArmsConfigLocation) {
		query["PhpArmsConfigLocation"] = request.PhpArmsConfigLocation
	}

	if !dara.IsNil(request.PhpConfigLocation) {
		query["PhpConfigLocation"] = request.PhpConfigLocation
	}

	if !dara.IsNil(request.PostStart) {
		query["PostStart"] = request.PostStart
	}

	if !dara.IsNil(request.PreStop) {
		query["PreStop"] = request.PreStop
	}

	if !dara.IsNil(request.ProgrammingLanguage) {
		query["ProgrammingLanguage"] = request.ProgrammingLanguage
	}

	if !dara.IsNil(request.PvtzDiscoverySvc) {
		query["PvtzDiscoverySvc"] = request.PvtzDiscoverySvc
	}

	if !dara.IsNil(request.Python) {
		query["Python"] = request.Python
	}

	if !dara.IsNil(request.PythonModules) {
		query["PythonModules"] = request.PythonModules
	}

	if !dara.IsNil(request.Readiness) {
		query["Readiness"] = request.Readiness
	}

	if !dara.IsNil(request.Replicas) {
		query["Replicas"] = request.Replicas
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.SaeVersion) {
		query["SaeVersion"] = request.SaeVersion
	}

	if !dara.IsNil(request.SecretMountDesc) {
		query["SecretMountDesc"] = request.SecretMountDesc
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.SlsConfigs) {
		query["SlsConfigs"] = request.SlsConfigs
	}

	if !dara.IsNil(request.StartupProbe) {
		query["StartupProbe"] = request.StartupProbe
	}

	if !dara.IsNil(request.TerminationGracePeriodSeconds) {
		query["TerminationGracePeriodSeconds"] = request.TerminationGracePeriodSeconds
	}

	if !dara.IsNil(request.Timezone) {
		query["Timezone"] = request.Timezone
	}

	if !dara.IsNil(request.TomcatConfig) {
		query["TomcatConfig"] = request.TomcatConfig
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.WarStartOptions) {
		query["WarStartOptions"] = request.WarStartOptions
	}

	if !dara.IsNil(request.WebContainer) {
		query["WebContainer"] = request.WebContainer
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AcrInstanceId) {
		body["AcrInstanceId"] = request.AcrInstanceId
	}

	if !dara.IsNil(request.AssociateEip) {
		body["AssociateEip"] = request.AssociateEip
	}

	if !dara.IsNil(request.BaseAppId) {
		body["BaseAppId"] = request.BaseAppId
	}

	if !dara.IsNil(request.ConfigMapMountDesc) {
		body["ConfigMapMountDesc"] = request.ConfigMapMountDesc
	}

	if !dara.IsNil(request.EnableSidecarResourceIsolated) {
		body["EnableSidecarResourceIsolated"] = request.EnableSidecarResourceIsolated
	}

	if !dara.IsNil(request.InitContainersConfigShrink) {
		body["InitContainersConfig"] = request.InitContainersConfigShrink
	}

	if !dara.IsNil(request.MicroRegistrationConfig) {
		body["MicroRegistrationConfig"] = request.MicroRegistrationConfig
	}

	if !dara.IsNil(request.OssAkId) {
		body["OssAkId"] = request.OssAkId
	}

	if !dara.IsNil(request.OssAkSecret) {
		body["OssAkSecret"] = request.OssAkSecret
	}

	if !dara.IsNil(request.OssMountDescs) {
		body["OssMountDescs"] = request.OssMountDescs
	}

	if !dara.IsNil(request.Php) {
		body["Php"] = request.Php
	}

	if !dara.IsNil(request.PhpConfig) {
		body["PhpConfig"] = request.PhpConfig
	}

	if !dara.IsNil(request.ServiceTags) {
		body["ServiceTags"] = request.ServiceTags
	}

	if !dara.IsNil(request.SidecarContainersConfigShrink) {
		body["SidecarContainersConfig"] = request.SidecarContainersConfigShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApplication"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/app/createApplication"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Null
//
// Description:
//
// The HTTP status code. Take note of the following rules:
//
//   - **2xx**: The call was successful.
//
//   - **3xx**: The call was redirected.
//
//   - **4xx**: The call failed.
//
//   - **5xx**: A server error occurred.
//
// @param request - CreateApplicationScalingRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApplicationScalingRuleResponse
func (client *Client) CreateApplicationScalingRuleWithContext(ctx context.Context, request *CreateApplicationScalingRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateApplicationScalingRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EnableIdle) {
		query["EnableIdle"] = request.EnableIdle
	}

	if !dara.IsNil(request.MinReadyInstanceRatio) {
		query["MinReadyInstanceRatio"] = request.MinReadyInstanceRatio
	}

	if !dara.IsNil(request.MinReadyInstances) {
		query["MinReadyInstances"] = request.MinReadyInstances
	}

	if !dara.IsNil(request.ScalingRuleEnable) {
		query["ScalingRuleEnable"] = request.ScalingRuleEnable
	}

	if !dara.IsNil(request.ScalingRuleMetric) {
		query["ScalingRuleMetric"] = request.ScalingRuleMetric
	}

	if !dara.IsNil(request.ScalingRuleName) {
		query["ScalingRuleName"] = request.ScalingRuleName
	}

	if !dara.IsNil(request.ScalingRuleTimer) {
		query["ScalingRuleTimer"] = request.ScalingRuleTimer
	}

	if !dara.IsNil(request.ScalingRuleType) {
		query["ScalingRuleType"] = request.ScalingRuleType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApplicationScalingRule"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/scale/applicationScalingRule"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApplicationScalingRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a ConfigMap in a namespace.
//
// @param request - CreateConfigMapRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConfigMapResponse
func (client *Client) CreateConfigMapWithContext(ctx context.Context, request *CreateConfigMapRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateConfigMapResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Data) {
		body["Data"] = request.Data
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateConfigMap"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/configmap/configMap"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateConfigMapResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a canary release rule for a Spring Cloud or Dubbo application.
//
// Description:
//
// >  You can configure only one canary release rule for each application.
//
// @param request - CreateGreyTagRouteRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGreyTagRouteResponse
func (client *Client) CreateGreyTagRouteWithContext(ctx context.Context, request *CreateGreyTagRouteRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateGreyTagRouteResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlbRules) {
		query["AlbRules"] = request.AlbRules
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DubboRules) {
		query["DubboRules"] = request.DubboRules
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ScRules) {
		query["ScRules"] = request.ScRules
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGreyTagRoute"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/tagroute/greyTagRoute"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGreyTagRouteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a routing rule.
//
// @param request - CreateIngressRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIngressResponse
func (client *Client) CreateIngressWithContext(ctx context.Context, request *CreateIngressRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateIngressResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddressType) {
		query["AddressType"] = request.AddressType
	}

	if !dara.IsNil(request.CertId) {
		query["CertId"] = request.CertId
	}

	if !dara.IsNil(request.CertIds) {
		query["CertIds"] = request.CertIds
	}

	if !dara.IsNil(request.CorsConfig) {
		query["CorsConfig"] = request.CorsConfig
	}

	if !dara.IsNil(request.DefaultRule) {
		query["DefaultRule"] = request.DefaultRule
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EnableXForwardedFor) {
		query["EnableXForwardedFor"] = request.EnableXForwardedFor
	}

	if !dara.IsNil(request.EnableXForwardedForClientSrcPort) {
		query["EnableXForwardedForClientSrcPort"] = request.EnableXForwardedForClientSrcPort
	}

	if !dara.IsNil(request.EnableXForwardedForProto) {
		query["EnableXForwardedForProto"] = request.EnableXForwardedForProto
	}

	if !dara.IsNil(request.EnableXForwardedForSlbId) {
		query["EnableXForwardedForSlbId"] = request.EnableXForwardedForSlbId
	}

	if !dara.IsNil(request.EnableXForwardedForSlbPort) {
		query["EnableXForwardedForSlbPort"] = request.EnableXForwardedForSlbPort
	}

	if !dara.IsNil(request.IdleTimeout) {
		query["IdleTimeout"] = request.IdleTimeout
	}

	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.ListenerProtocol) {
		query["ListenerProtocol"] = request.ListenerProtocol
	}

	if !dara.IsNil(request.LoadBalanceType) {
		query["LoadBalanceType"] = request.LoadBalanceType
	}

	if !dara.IsNil(request.LoadBalancerEdition) {
		query["LoadBalancerEdition"] = request.LoadBalancerEdition
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.RequestTimeout) {
		query["RequestTimeout"] = request.RequestTimeout
	}

	if !dara.IsNil(request.SecurityPolicyId) {
		query["SecurityPolicyId"] = request.SecurityPolicyId
	}

	if !dara.IsNil(request.SlbId) {
		query["SlbId"] = request.SlbId
	}

	if !dara.IsNil(request.ZoneMappings) {
		query["ZoneMappings"] = request.ZoneMappings
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Rules) {
		body["Rules"] = request.Rules
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIngress"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/ingress/Ingress"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIngressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a job template.
//
// @param request - CreateJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateJobResponse
func (client *Client) CreateJobWithContext(ctx context.Context, request *CreateJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcrAssumeRoleArn) {
		query["AcrAssumeRoleArn"] = request.AcrAssumeRoleArn
	}

	if !dara.IsNil(request.AppDescription) {
		query["AppDescription"] = request.AppDescription
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.AutoConfig) {
		query["AutoConfig"] = request.AutoConfig
	}

	if !dara.IsNil(request.BackoffLimit) {
		query["BackoffLimit"] = request.BackoffLimit
	}

	if !dara.IsNil(request.Command) {
		query["Command"] = request.Command
	}

	if !dara.IsNil(request.CommandArgs) {
		query["CommandArgs"] = request.CommandArgs
	}

	if !dara.IsNil(request.ConcurrencyPolicy) {
		query["ConcurrencyPolicy"] = request.ConcurrencyPolicy
	}

	if !dara.IsNil(request.Cpu) {
		query["Cpu"] = request.Cpu
	}

	if !dara.IsNil(request.CustomHostAlias) {
		query["CustomHostAlias"] = request.CustomHostAlias
	}

	if !dara.IsNil(request.EdasContainerVersion) {
		query["EdasContainerVersion"] = request.EdasContainerVersion
	}

	if !dara.IsNil(request.Envs) {
		query["Envs"] = request.Envs
	}

	if !dara.IsNil(request.ImagePullSecrets) {
		query["ImagePullSecrets"] = request.ImagePullSecrets
	}

	if !dara.IsNil(request.ImageUrl) {
		query["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.JarStartArgs) {
		query["JarStartArgs"] = request.JarStartArgs
	}

	if !dara.IsNil(request.JarStartOptions) {
		query["JarStartOptions"] = request.JarStartOptions
	}

	if !dara.IsNil(request.Jdk) {
		query["Jdk"] = request.Jdk
	}

	if !dara.IsNil(request.Memory) {
		query["Memory"] = request.Memory
	}

	if !dara.IsNil(request.MountDesc) {
		query["MountDesc"] = request.MountDesc
	}

	if !dara.IsNil(request.MountHost) {
		query["MountHost"] = request.MountHost
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.NasId) {
		query["NasId"] = request.NasId
	}

	if !dara.IsNil(request.PackageType) {
		query["PackageType"] = request.PackageType
	}

	if !dara.IsNil(request.PackageUrl) {
		query["PackageUrl"] = request.PackageUrl
	}

	if !dara.IsNil(request.PackageVersion) {
		query["PackageVersion"] = request.PackageVersion
	}

	if !dara.IsNil(request.PhpConfigLocation) {
		query["PhpConfigLocation"] = request.PhpConfigLocation
	}

	if !dara.IsNil(request.PostStart) {
		query["PostStart"] = request.PostStart
	}

	if !dara.IsNil(request.PreStop) {
		query["PreStop"] = request.PreStop
	}

	if !dara.IsNil(request.ProgrammingLanguage) {
		query["ProgrammingLanguage"] = request.ProgrammingLanguage
	}

	if !dara.IsNil(request.Python) {
		query["Python"] = request.Python
	}

	if !dara.IsNil(request.PythonModules) {
		query["PythonModules"] = request.PythonModules
	}

	if !dara.IsNil(request.RefAppId) {
		query["RefAppId"] = request.RefAppId
	}

	if !dara.IsNil(request.Replicas) {
		query["Replicas"] = request.Replicas
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.Slice) {
		query["Slice"] = request.Slice
	}

	if !dara.IsNil(request.SliceEnvs) {
		query["SliceEnvs"] = request.SliceEnvs
	}

	if !dara.IsNil(request.SlsConfigs) {
		query["SlsConfigs"] = request.SlsConfigs
	}

	if !dara.IsNil(request.TerminationGracePeriodSeconds) {
		query["TerminationGracePeriodSeconds"] = request.TerminationGracePeriodSeconds
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	if !dara.IsNil(request.Timezone) {
		query["Timezone"] = request.Timezone
	}

	if !dara.IsNil(request.TomcatConfig) {
		query["TomcatConfig"] = request.TomcatConfig
	}

	if !dara.IsNil(request.TriggerConfig) {
		query["TriggerConfig"] = request.TriggerConfig
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.WarStartOptions) {
		query["WarStartOptions"] = request.WarStartOptions
	}

	if !dara.IsNil(request.WebContainer) {
		query["WebContainer"] = request.WebContainer
	}

	if !dara.IsNil(request.Workload) {
		query["Workload"] = request.Workload
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AcrInstanceId) {
		body["AcrInstanceId"] = request.AcrInstanceId
	}

	if !dara.IsNil(request.ConfigMapMountDesc) {
		body["ConfigMapMountDesc"] = request.ConfigMapMountDesc
	}

	if !dara.IsNil(request.EnableImageAccl) {
		body["EnableImageAccl"] = request.EnableImageAccl
	}

	if !dara.IsNil(request.OssAkId) {
		body["OssAkId"] = request.OssAkId
	}

	if !dara.IsNil(request.OssAkSecret) {
		body["OssAkSecret"] = request.OssAkSecret
	}

	if !dara.IsNil(request.OssMountDescs) {
		body["OssMountDescs"] = request.OssMountDescs
	}

	if !dara.IsNil(request.PhpConfig) {
		body["PhpConfig"] = request.PhpConfig
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateJob"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/job/createJob"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a namespace.
//
// @param request - CreateNamespaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNamespaceResponse
func (client *Client) CreateNamespaceWithContext(ctx context.Context, request *CreateNamespaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateNamespaceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnableMicroRegistration) {
		query["EnableMicroRegistration"] = request.EnableMicroRegistration
	}

	if !dara.IsNil(request.NameSpaceShortId) {
		query["NameSpaceShortId"] = request.NameSpaceShortId
	}

	if !dara.IsNil(request.NamespaceDescription) {
		query["NamespaceDescription"] = request.NamespaceDescription
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.NamespaceName) {
		query["NamespaceName"] = request.NamespaceName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNamespace"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/paas/namespace"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNamespaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建或者更新泳道
//
// @param tmpReq - CreateOrUpdateSwimmingLaneRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOrUpdateSwimmingLaneResponse
func (client *Client) CreateOrUpdateSwimmingLaneWithContext(ctx context.Context, tmpReq *CreateOrUpdateSwimmingLaneRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateOrUpdateSwimmingLaneResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateOrUpdateSwimmingLaneShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AppEntryRule) {
		request.AppEntryRuleShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AppEntryRule, dara.String("AppEntryRule"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.MseGatewayEntryRule) {
		request.MseGatewayEntryRuleShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MseGatewayEntryRule, dara.String("MseGatewayEntryRule"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppEntryRuleShrink) {
		query["AppEntryRule"] = request.AppEntryRuleShrink
	}

	if !dara.IsNil(request.CanaryModel) {
		query["CanaryModel"] = request.CanaryModel
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.LaneId) {
		query["LaneId"] = request.LaneId
	}

	if !dara.IsNil(request.LaneName) {
		query["LaneName"] = request.LaneName
	}

	if !dara.IsNil(request.LaneTag) {
		query["LaneTag"] = request.LaneTag
	}

	if !dara.IsNil(request.MseGatewayEntryRuleShrink) {
		query["MseGatewayEntryRule"] = request.MseGatewayEntryRuleShrink
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOrUpdateSwimmingLane"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/cas/gray/createOrUpdateSwimmingLane"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOrUpdateSwimmingLaneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建或者更新泳道组
//
// @param tmpReq - CreateOrUpdateSwimmingLaneGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOrUpdateSwimmingLaneGroupResponse
func (client *Client) CreateOrUpdateSwimmingLaneGroupWithContext(ctx context.Context, tmpReq *CreateOrUpdateSwimmingLaneGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateOrUpdateSwimmingLaneGroupResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateOrUpdateSwimmingLaneGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AppIds) {
		request.AppIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AppIds, dara.String("AppIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppIdsShrink) {
		query["AppIds"] = request.AppIdsShrink
	}

	if !dara.IsNil(request.EntryAppId) {
		query["EntryAppId"] = request.EntryAppId
	}

	if !dara.IsNil(request.EntryAppType) {
		query["EntryAppType"] = request.EntryAppType
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.SwimVersion) {
		query["SwimVersion"] = request.SwimVersion
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOrUpdateSwimmingLaneGroup"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/cas/gray/createOrUpdateSwimmingLaneGroup"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOrUpdateSwimmingLaneGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Secret in a namespace.
//
// @param tmpReq - CreateSecretRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSecretResponse
func (client *Client) CreateSecretWithContext(ctx context.Context, tmpReq *CreateSecretRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateSecretResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateSecretShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SecretData) {
		request.SecretDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SecretData, dara.String("SecretData"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.SecretDataShrink) {
		query["SecretData"] = request.SecretDataShrink
	}

	if !dara.IsNil(request.SecretName) {
		query["SecretName"] = request.SecretName
	}

	if !dara.IsNil(request.SecretType) {
		query["SecretType"] = request.SecretType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSecret"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/secret/secret"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSecretResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create a web application
//
// Description:
//
// Call the CreateWebApplication operation to create a web application.
//
// @param request - CreateWebApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWebApplicationResponse
func (client *Client) CreateWebApplicationWithContext(ctx context.Context, request *CreateWebApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateWebApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWebApplication"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v2/api/web/applications"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWebApplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a custom domain name for the web application.
//
// Description:
//
// Create a custom domain name for the web application.
//
// @param request - CreateWebCustomDomainRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWebCustomDomainResponse
func (client *Client) CreateWebCustomDomainWithContext(ctx context.Context, request *CreateWebCustomDomainRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateWebCustomDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWebCustomDomain"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v2/api/web/custom-domains"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWebCustomDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an application.
//
// @param request - DeleteApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApplicationResponse
func (client *Client) DeleteApplicationWithContext(ctx context.Context, request *DeleteApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApplication"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/app/deleteApplication"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 7171a6ca-d1cd-4928-8642-7d5cfe69\\*\\*\\*\\*
//
// @param request - DeleteApplicationScalingRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApplicationScalingRuleResponse
func (client *Client) DeleteApplicationScalingRuleWithContext(ctx context.Context, request *DeleteApplicationScalingRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteApplicationScalingRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ScalingRuleName) {
		query["ScalingRuleName"] = request.ScalingRuleName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApplicationScalingRule"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/scale/applicationScalingRule"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApplicationScalingRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a ConfigMap.
//
// @param request - DeleteConfigMapRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConfigMapResponse
func (client *Client) DeleteConfigMapWithContext(ctx context.Context, request *DeleteConfigMapRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteConfigMapResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigMapId) {
		query["ConfigMapId"] = request.ConfigMapId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteConfigMap"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/configmap/configMap"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteConfigMapResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a canary release rule based on the specified rule ID.
//
// @param request - DeleteGreyTagRouteRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGreyTagRouteResponse
func (client *Client) DeleteGreyTagRouteWithContext(ctx context.Context, request *DeleteGreyTagRouteRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteGreyTagRouteResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GreyTagRouteId) {
		query["GreyTagRouteId"] = request.GreyTagRouteId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGreyTagRoute"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/tagroute/greyTagRoute"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGreyTagRouteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a job.
//
// @param request - DeleteHistoryJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHistoryJobResponse
func (client *Client) DeleteHistoryJobWithContext(ctx context.Context, request *DeleteHistoryJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteHistoryJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHistoryJob"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/job/deleteHistoryJob"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHistoryJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a routing rule.
//
// @param request - DeleteIngressRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIngressResponse
func (client *Client) DeleteIngressWithContext(ctx context.Context, request *DeleteIngressRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteIngressResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IngressId) {
		query["IngressId"] = request.IngressId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteIngress"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/ingress/Ingress"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIngressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除实例
//
// @param request - DeleteInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstancesResponse
func (client *Client) DeleteInstancesWithContext(ctx context.Context, request *DeleteInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteInstances"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/app/deleteInstances"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a job template.
//
// @param request - DeleteJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteJobResponse
func (client *Client) DeleteJobWithContext(ctx context.Context, request *DeleteJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteJob"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/job/deleteJob"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNamespaceResponse
func (client *Client) DeleteNamespaceWithContext(ctx context.Context, request *DeleteNamespaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteNamespaceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NameSpaceShortId) {
		query["NameSpaceShortId"] = request.NameSpaceShortId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNamespace"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/paas/namespace"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNamespaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Secret.
//
// @param request - DeleteSecretRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSecretResponse
func (client *Client) DeleteSecretWithContext(ctx context.Context, request *DeleteSecretRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteSecretResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.SecretId) {
		query["SecretId"] = request.SecretId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSecret"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/secret/secret"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSecretResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除泳道组
//
// @param request - DeleteSwimmingLaneGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSwimmingLaneGroupResponse
func (client *Client) DeleteSwimmingLaneGroupWithContext(ctx context.Context, request *DeleteSwimmingLaneGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteSwimmingLaneGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSwimmingLaneGroup"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/cas/gray/deleteSwimmingLaneGroup"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSwimmingLaneGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a web application.
//
// Description:
//
// Call the DeleteWebApplication operation to delete a web application.
//
// @param request - DeleteWebApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWebApplicationResponse
func (client *Client) DeleteWebApplicationWithContext(ctx context.Context, ApplicationId *string, request *DeleteWebApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteWebApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWebApplication"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v2/api/web/applications/" + dara.PercentEncode(dara.StringValue(ApplicationId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWebApplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a web application version.
//
// Description:
//
// Delete a web application version.
//
// @param request - DeleteWebApplicationRevisionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWebApplicationRevisionResponse
func (client *Client) DeleteWebApplicationRevisionWithContext(ctx context.Context, ApplicationId *string, RevisionId *string, request *DeleteWebApplicationRevisionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteWebApplicationRevisionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWebApplicationRevision"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v2/api/web/application-revisions/" + dara.PercentEncode(dara.StringValue(ApplicationId)) + "/revisions/" + dara.PercentEncode(dara.StringValue(RevisionId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWebApplicationRevisionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a custom domain name.
//
// Description:
//
// Delete a custom domain name.
//
// @param request - DeleteWebCustomDomainRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWebCustomDomainResponse
func (client *Client) DeleteWebCustomDomainWithContext(ctx context.Context, DomainName *string, request *DeleteWebCustomDomainRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteWebCustomDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWebCustomDomain"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v2/api/web/custom-domains/" + dara.PercentEncode(dara.StringValue(DomainName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWebCustomDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deploys an application.
//
// @param tmpReq - DeployApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeployApplicationResponse
func (client *Client) DeployApplicationWithContext(ctx context.Context, tmpReq *DeployApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeployApplicationResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DeployApplicationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InitContainersConfig) {
		request.InitContainersConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InitContainersConfig, dara.String("InitContainersConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SidecarContainersConfig) {
		request.SidecarContainersConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SidecarContainersConfig, dara.String("SidecarContainersConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcrAssumeRoleArn) {
		query["AcrAssumeRoleArn"] = request.AcrAssumeRoleArn
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AutoEnableApplicationScalingRule) {
		query["AutoEnableApplicationScalingRule"] = request.AutoEnableApplicationScalingRule
	}

	if !dara.IsNil(request.BatchWaitTime) {
		query["BatchWaitTime"] = request.BatchWaitTime
	}

	if !dara.IsNil(request.ChangeOrderDesc) {
		query["ChangeOrderDesc"] = request.ChangeOrderDesc
	}

	if !dara.IsNil(request.Command) {
		query["Command"] = request.Command
	}

	if !dara.IsNil(request.CommandArgs) {
		query["CommandArgs"] = request.CommandArgs
	}

	if !dara.IsNil(request.Cpu) {
		query["Cpu"] = request.Cpu
	}

	if !dara.IsNil(request.CustomHostAlias) {
		query["CustomHostAlias"] = request.CustomHostAlias
	}

	if !dara.IsNil(request.CustomImageNetworkType) {
		query["CustomImageNetworkType"] = request.CustomImageNetworkType
	}

	if !dara.IsNil(request.Deploy) {
		query["Deploy"] = request.Deploy
	}

	if !dara.IsNil(request.Dotnet) {
		query["Dotnet"] = request.Dotnet
	}

	if !dara.IsNil(request.EdasContainerVersion) {
		query["EdasContainerVersion"] = request.EdasContainerVersion
	}

	if !dara.IsNil(request.EnableAhas) {
		query["EnableAhas"] = request.EnableAhas
	}

	if !dara.IsNil(request.EnableCpuBurst) {
		query["EnableCpuBurst"] = request.EnableCpuBurst
	}

	if !dara.IsNil(request.EnableGreyTagRoute) {
		query["EnableGreyTagRoute"] = request.EnableGreyTagRoute
	}

	if !dara.IsNil(request.EnableNewArms) {
		query["EnableNewArms"] = request.EnableNewArms
	}

	if !dara.IsNil(request.EnablePrometheus) {
		query["EnablePrometheus"] = request.EnablePrometheus
	}

	if !dara.IsNil(request.Envs) {
		query["Envs"] = request.Envs
	}

	if !dara.IsNil(request.GpuConfig) {
		query["GpuConfig"] = request.GpuConfig
	}

	if !dara.IsNil(request.ImagePullSecrets) {
		query["ImagePullSecrets"] = request.ImagePullSecrets
	}

	if !dara.IsNil(request.ImageUrl) {
		query["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.JarStartArgs) {
		query["JarStartArgs"] = request.JarStartArgs
	}

	if !dara.IsNil(request.JarStartOptions) {
		query["JarStartOptions"] = request.JarStartOptions
	}

	if !dara.IsNil(request.Jdk) {
		query["Jdk"] = request.Jdk
	}

	if !dara.IsNil(request.KafkaConfigs) {
		query["KafkaConfigs"] = request.KafkaConfigs
	}

	if !dara.IsNil(request.Liveness) {
		query["Liveness"] = request.Liveness
	}

	if !dara.IsNil(request.Memory) {
		query["Memory"] = request.Memory
	}

	if !dara.IsNil(request.MicroRegistration) {
		query["MicroRegistration"] = request.MicroRegistration
	}

	if !dara.IsNil(request.MicroserviceEngineConfig) {
		query["MicroserviceEngineConfig"] = request.MicroserviceEngineConfig
	}

	if !dara.IsNil(request.MinReadyInstanceRatio) {
		query["MinReadyInstanceRatio"] = request.MinReadyInstanceRatio
	}

	if !dara.IsNil(request.MinReadyInstances) {
		query["MinReadyInstances"] = request.MinReadyInstances
	}

	if !dara.IsNil(request.MountDesc) {
		query["MountDesc"] = request.MountDesc
	}

	if !dara.IsNil(request.MountHost) {
		query["MountHost"] = request.MountHost
	}

	if !dara.IsNil(request.NasConfigs) {
		query["NasConfigs"] = request.NasConfigs
	}

	if !dara.IsNil(request.NasId) {
		query["NasId"] = request.NasId
	}

	if !dara.IsNil(request.NewSaeVersion) {
		query["NewSaeVersion"] = request.NewSaeVersion
	}

	if !dara.IsNil(request.OidcRoleName) {
		query["OidcRoleName"] = request.OidcRoleName
	}

	if !dara.IsNil(request.PackageType) {
		query["PackageType"] = request.PackageType
	}

	if !dara.IsNil(request.PackageUrl) {
		query["PackageUrl"] = request.PackageUrl
	}

	if !dara.IsNil(request.PackageVersion) {
		query["PackageVersion"] = request.PackageVersion
	}

	if !dara.IsNil(request.PhpArmsConfigLocation) {
		query["PhpArmsConfigLocation"] = request.PhpArmsConfigLocation
	}

	if !dara.IsNil(request.PhpConfigLocation) {
		query["PhpConfigLocation"] = request.PhpConfigLocation
	}

	if !dara.IsNil(request.PostStart) {
		query["PostStart"] = request.PostStart
	}

	if !dara.IsNil(request.PreStop) {
		query["PreStop"] = request.PreStop
	}

	if !dara.IsNil(request.PvtzDiscoverySvc) {
		query["PvtzDiscoverySvc"] = request.PvtzDiscoverySvc
	}

	if !dara.IsNil(request.Python) {
		query["Python"] = request.Python
	}

	if !dara.IsNil(request.PythonModules) {
		query["PythonModules"] = request.PythonModules
	}

	if !dara.IsNil(request.Readiness) {
		query["Readiness"] = request.Readiness
	}

	if !dara.IsNil(request.Replicas) {
		query["Replicas"] = request.Replicas
	}

	if !dara.IsNil(request.SecretMountDesc) {
		query["SecretMountDesc"] = request.SecretMountDesc
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.SlsConfigs) {
		query["SlsConfigs"] = request.SlsConfigs
	}

	if !dara.IsNil(request.StartupProbe) {
		query["StartupProbe"] = request.StartupProbe
	}

	if !dara.IsNil(request.TerminationGracePeriodSeconds) {
		query["TerminationGracePeriodSeconds"] = request.TerminationGracePeriodSeconds
	}

	if !dara.IsNil(request.Timezone) {
		query["Timezone"] = request.Timezone
	}

	if !dara.IsNil(request.TomcatConfig) {
		query["TomcatConfig"] = request.TomcatConfig
	}

	if !dara.IsNil(request.UpdateStrategy) {
		query["UpdateStrategy"] = request.UpdateStrategy
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.WarStartOptions) {
		query["WarStartOptions"] = request.WarStartOptions
	}

	if !dara.IsNil(request.WebContainer) {
		query["WebContainer"] = request.WebContainer
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AcrInstanceId) {
		body["AcrInstanceId"] = request.AcrInstanceId
	}

	if !dara.IsNil(request.AssociateEip) {
		body["AssociateEip"] = request.AssociateEip
	}

	if !dara.IsNil(request.ConfigMapMountDesc) {
		body["ConfigMapMountDesc"] = request.ConfigMapMountDesc
	}

	if !dara.IsNil(request.EnableSidecarResourceIsolated) {
		body["EnableSidecarResourceIsolated"] = request.EnableSidecarResourceIsolated
	}

	if !dara.IsNil(request.InitContainersConfigShrink) {
		body["InitContainersConfig"] = request.InitContainersConfigShrink
	}

	if !dara.IsNil(request.MicroRegistrationConfig) {
		body["MicroRegistrationConfig"] = request.MicroRegistrationConfig
	}

	if !dara.IsNil(request.OssAkId) {
		body["OssAkId"] = request.OssAkId
	}

	if !dara.IsNil(request.OssAkSecret) {
		body["OssAkSecret"] = request.OssAkSecret
	}

	if !dara.IsNil(request.OssMountDescs) {
		body["OssMountDescs"] = request.OssMountDescs
	}

	if !dara.IsNil(request.Php) {
		body["Php"] = request.Php
	}

	if !dara.IsNil(request.PhpConfig) {
		body["PhpConfig"] = request.PhpConfig
	}

	if !dara.IsNil(request.ServiceTags) {
		body["ServiceTags"] = request.ServiceTags
	}

	if !dara.IsNil(request.SidecarContainersConfigShrink) {
		body["SidecarContainersConfig"] = request.SidecarContainersConfigShrink
	}

	if !dara.IsNil(request.SwimlanePvtzDiscoverySvc) {
		body["SwimlanePvtzDiscoverySvc"] = request.SwimlanePvtzDiscoverySvc
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeployApplication"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/app/deployApplication"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeployApplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the metadata details of the service of an application.
//
// @param request - DescribeAppServiceDetailRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppServiceDetailResponse
func (client *Client) DescribeAppServiceDetailWithContext(ctx context.Context, request *DescribeAppServiceDetailRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeAppServiceDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.NacosInstanceId) {
		query["NacosInstanceId"] = request.NacosInstanceId
	}

	if !dara.IsNil(request.NacosNamespaceId) {
		query["NacosNamespaceId"] = request.NacosNamespaceId
	}

	if !dara.IsNil(request.ServiceGroup) {
		query["ServiceGroup"] = request.ServiceGroup
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	if !dara.IsNil(request.ServiceType) {
		query["ServiceType"] = request.ServiceType
	}

	if !dara.IsNil(request.ServiceVersion) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppServiceDetail"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/service/describeAppServiceDetail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppServiceDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of an application.
//
// @param request - DescribeApplicationConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApplicationConfigResponse
func (client *Client) DescribeApplicationConfigWithContext(ctx context.Context, request *DescribeApplicationConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeApplicationConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.VersionId) {
		query["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApplicationConfig"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/app/describeApplicationConfig"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApplicationConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the instance groups of an application.
//
// @param request - DescribeApplicationGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApplicationGroupsResponse
func (client *Client) DescribeApplicationGroupsWithContext(ctx context.Context, request *DescribeApplicationGroupsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeApplicationGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApplicationGroups"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/app/describeApplicationGroups"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApplicationGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the image of an application.
//
// @param request - DescribeApplicationImageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApplicationImageResponse
func (client *Client) DescribeApplicationImageWithContext(ctx context.Context, request *DescribeApplicationImageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeApplicationImageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ImageUrl) {
		query["ImageUrl"] = request.ImageUrl
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApplicationImage"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/container/describeApplicationImage"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApplicationImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of application instances.
//
// @param request - DescribeApplicationInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApplicationInstancesResponse
func (client *Client) DescribeApplicationInstancesWithContext(ctx context.Context, request *DescribeApplicationInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeApplicationInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PipelineId) {
		query["PipelineId"] = request.PipelineId
	}

	if !dara.IsNil(request.Reverse) {
		query["Reverse"] = request.Reverse
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApplicationInstances"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/app/describeApplicationInstances"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApplicationInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Network Load Balancer (NLB) instances bound to an application and their listeners.
//
// @param request - DescribeApplicationNlbsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApplicationNlbsResponse
func (client *Client) DescribeApplicationNlbsWithContext(ctx context.Context, request *DescribeApplicationNlbsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeApplicationNlbsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApplicationNlbs"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/app/nlb"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApplicationNlbsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries an Auto Scaling policy of an application.
//
// @param request - DescribeApplicationScalingRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApplicationScalingRuleResponse
func (client *Client) DescribeApplicationScalingRuleWithContext(ctx context.Context, request *DescribeApplicationScalingRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeApplicationScalingRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ScalingRuleName) {
		query["ScalingRuleName"] = request.ScalingRuleName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApplicationScalingRule"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/scale/applicationScalingRule"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApplicationScalingRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the auto scaling policies of an application.
//
// @param request - DescribeApplicationScalingRulesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApplicationScalingRulesResponse
func (client *Client) DescribeApplicationScalingRulesWithContext(ctx context.Context, request *DescribeApplicationScalingRulesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeApplicationScalingRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApplicationScalingRules"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/scale/applicationScalingRules"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApplicationScalingRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 017f39b8-dfa4-4e16-a84b-1dcee4b1\\*\\*\\*\\*
//
// @param request - DescribeApplicationSlbsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApplicationSlbsResponse
func (client *Client) DescribeApplicationSlbsWithContext(ctx context.Context, request *DescribeApplicationSlbsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeApplicationSlbsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApplicationSlbs"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/app/slb"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApplicationSlbsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of an application.
//
// @param request - DescribeApplicationStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApplicationStatusResponse
func (client *Client) DescribeApplicationStatusWithContext(ctx context.Context, request *DescribeApplicationStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeApplicationStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApplicationStatus"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/app/describeApplicationStatus"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApplicationStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of a change order.
//
// @param request - DescribeChangeOrderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChangeOrderResponse
func (client *Client) DescribeChangeOrderWithContext(ctx context.Context, request *DescribeChangeOrderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeChangeOrderResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChangeOrderId) {
		query["ChangeOrderId"] = request.ChangeOrderId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeChangeOrder"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/changeorder/DescribeChangeOrder"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeChangeOrderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the version of the component that is required when you create and deploy an application.
//
// @param request - DescribeComponentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeComponentsResponse
func (client *Client) DescribeComponentsWithContext(ctx context.Context, request *DescribeComponentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeComponentsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeComponents"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/resource/components"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeComponentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a ConfigMap.
//
// @param request - DescribeConfigMapRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeConfigMapResponse
func (client *Client) DescribeConfigMapWithContext(ctx context.Context, request *DescribeConfigMapRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeConfigMapResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigMapId) {
		query["ConfigMapId"] = request.ConfigMapId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeConfigMap"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/configmap/configMap"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeConfigMapResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query configuration price.
//
// @param request - DescribeConfigurationPriceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeConfigurationPriceResponse
func (client *Client) DescribeConfigurationPriceWithContext(ctx context.Context, request *DescribeConfigurationPriceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeConfigurationPriceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cpu) {
		query["Cpu"] = request.Cpu
	}

	if !dara.IsNil(request.Memory) {
		query["Memory"] = request.Memory
	}

	if !dara.IsNil(request.NewSaeVersion) {
		query["NewSaeVersion"] = request.NewSaeVersion
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Workload) {
		query["Workload"] = request.Workload
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeConfigurationPrice"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/paas/configurationPrice"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeConfigurationPriceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the container components of a microservices application.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEdasContainersResponse
func (client *Client) DescribeEdasContainersWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeEdasContainersResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEdasContainers"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/resource/edasContainers"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEdasContainersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a canary release rule based on the specified rule ID.
//
// @param request - DescribeGreyTagRouteRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGreyTagRouteResponse
func (client *Client) DescribeGreyTagRouteWithContext(ctx context.Context, request *DescribeGreyTagRouteRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeGreyTagRouteResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GreyTagRouteId) {
		query["GreyTagRouteId"] = request.GreyTagRouteId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGreyTagRoute"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/tagroute/greyTagRoute"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGreyTagRouteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call the DescribeIngress operation to query the details of an Ingress.
//
// @param request - DescribeIngressRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIngressResponse
func (client *Client) DescribeIngressWithContext(ctx context.Context, request *DescribeIngressRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeIngressResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IngressId) {
		query["IngressId"] = request.IngressId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeIngress"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/ingress/Ingress"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIngressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the logs of a sidecar container instance.
//
// @param request - DescribeInstanceLogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceLogResponse
func (client *Client) DescribeInstanceLogWithContext(ctx context.Context, request *DescribeInstanceLogRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeInstanceLogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContainerId) {
		query["ContainerId"] = request.ContainerId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceLog"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/instance/describeInstanceLog"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all instance types.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceSpecificationsResponse
func (client *Client) DescribeInstanceSpecificationsWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeInstanceSpecificationsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceSpecifications"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/paas/quota/instanceSpecifications"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceSpecificationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of a job template.
//
// @param request - DescribeJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeJobResponse
func (client *Client) DescribeJobWithContext(ctx context.Context, request *DescribeJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeJob"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/job/describeJob"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the information about jobs.
//
// @param request - DescribeJobHistoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeJobHistoryResponse
func (client *Client) DescribeJobHistoryWithContext(ctx context.Context, request *DescribeJobHistoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeJobHistoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeJobHistory"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/job/describeJobHistory"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeJobHistoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of a job.
//
// @param request - DescribeJobStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeJobStatusResponse
func (client *Client) DescribeJobStatusWithContext(ctx context.Context, request *DescribeJobStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeJobStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeJobStatus"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/job/describeJobStatus"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeJobStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a namespace.
//
// @param request - DescribeNamespaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNamespaceResponse
func (client *Client) DescribeNamespaceWithContext(ctx context.Context, request *DescribeNamespaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeNamespaceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NameSpaceShortId) {
		query["NameSpaceShortId"] = request.NameSpaceShortId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNamespace"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/paas/namespace"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNamespaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param request - DescribeNamespaceListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNamespaceListResponse
func (client *Client) DescribeNamespaceListWithContext(ctx context.Context, request *DescribeNamespaceListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeNamespaceListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContainCustom) {
		query["ContainCustom"] = request.ContainCustom
	}

	if !dara.IsNil(request.HybridCloudExclude) {
		query["HybridCloudExclude"] = request.HybridCloudExclude
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNamespaceList"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/namespace/describeNamespaceList"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNamespaceListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the information about resources in a namespace.
//
// @param request - DescribeNamespaceResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNamespaceResourcesResponse
func (client *Client) DescribeNamespaceResourcesWithContext(ctx context.Context, request *DescribeNamespaceResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeNamespaceResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NameSpaceShortId) {
		query["NameSpaceShortId"] = request.NameSpaceShortId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNamespaceResources"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/namespace/describeNamespaceResources"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNamespaceResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param request - DescribeNamespacesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNamespacesResponse
func (client *Client) DescribeNamespacesWithContext(ctx context.Context, request *DescribeNamespacesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeNamespacesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNamespaces"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/paas/namespaces"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNamespacesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # View batch information
//
// @param request - DescribePipelineRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePipelineResponse
func (client *Client) DescribePipelineWithContext(ctx context.Context, request *DescribePipelineRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribePipelineResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PipelineId) {
		query["PipelineId"] = request.PipelineId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePipeline"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/changeorder/DescribePipeline"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePipelineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries available regions.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/paas/regionConfig"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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
// Queries the details of a Secret.
//
// @param request - DescribeSecretRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSecretResponse
func (client *Client) DescribeSecretWithContext(ctx context.Context, request *DescribeSecretRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeSecretResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.SecretId) {
		query["SecretId"] = request.SecretId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSecret"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/secret/secret"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSecretResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询泳道详情
//
// @param request - DescribeSwimmingLaneRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSwimmingLaneResponse
func (client *Client) DescribeSwimmingLaneWithContext(ctx context.Context, request *DescribeSwimmingLaneRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeSwimmingLaneResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.LaneId) {
		query["LaneId"] = request.LaneId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSwimmingLane"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/cas/gray/describeSwimmingLane"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSwimmingLaneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query web applications.
//
// Description:
//
// Call the DescribeWebApplication operation to query web applications.
//
// @param request - DescribeWebApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWebApplicationResponse
func (client *Client) DescribeWebApplicationWithContext(ctx context.Context, ApplicationId *string, request *DescribeWebApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeWebApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWebApplication"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v2/api/web/applications/" + dara.PercentEncode(dara.StringValue(ApplicationId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWebApplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the resource usage of a web application.
//
// Description:
//
// Query the resource usage of a web application.
//
// @param request - DescribeWebApplicationResourceStaticsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWebApplicationResourceStaticsResponse
func (client *Client) DescribeWebApplicationResourceStaticsWithContext(ctx context.Context, ApplicationId *string, request *DescribeWebApplicationResourceStaticsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeWebApplicationResourceStaticsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWebApplicationResourceStatics"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v2/api/web/applications-observability/" + dara.PercentEncode(dara.StringValue(ApplicationId)) + "/resource"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWebApplicationResourceStaticsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Describe a web application version.
//
// Description:
//
// Describe a web application version.
//
// @param request - DescribeWebApplicationRevisionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWebApplicationRevisionResponse
func (client *Client) DescribeWebApplicationRevisionWithContext(ctx context.Context, ApplicationId *string, RevisionId *string, request *DescribeWebApplicationRevisionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeWebApplicationRevisionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWebApplicationRevision"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v2/api/web/application-revisions/" + dara.PercentEncode(dara.StringValue(ApplicationId)) + "/revisions/" + dara.PercentEncode(dara.StringValue(RevisionId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWebApplicationRevisionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Describe the scaling configuration of a web application.
//
// Description:
//
// Call the DescribeWebApplicationScalingConfig operation to obtain the scaling configuration of a web application.
//
// @param request - DescribeWebApplicationScalingConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWebApplicationScalingConfigResponse
func (client *Client) DescribeWebApplicationScalingConfigWithContext(ctx context.Context, ApplicationId *string, request *DescribeWebApplicationScalingConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeWebApplicationScalingConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWebApplicationScalingConfig"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v2/api/web/application-scaling/" + dara.PercentEncode(dara.StringValue(ApplicationId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWebApplicationScalingConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the traffic configurations of a web application.
//
// Description:
//
// Call the DescribeWebApplicationTrafficConfig operation to query the traffic configurations of a web application.
//
// @param request - DescribeWebApplicationTrafficConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWebApplicationTrafficConfigResponse
func (client *Client) DescribeWebApplicationTrafficConfigWithContext(ctx context.Context, ApplicationId *string, request *DescribeWebApplicationTrafficConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeWebApplicationTrafficConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWebApplicationTrafficConfig"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v2/api/web/application-traffic/" + dara.PercentEncode(dara.StringValue(ApplicationId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWebApplicationTrafficConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the details of a custom domain name for a web application.
//
// Description:
//
// Query the details of a custom domain name for a web application.
//
// @param request - DescribeWebCustomDomainRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWebCustomDomainResponse
func (client *Client) DescribeWebCustomDomainWithContext(ctx context.Context, DomainName *string, request *DescribeWebCustomDomainRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeWebCustomDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWebCustomDomain"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v2/api/web/custom-domains/" + dara.PercentEncode(dara.StringValue(DomainName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWebCustomDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the logs of web application instances.
//
// Description:
//
// Obtain the logs of web application instances.
//
// @param request - DescribeWebInstanceLogsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWebInstanceLogsResponse
func (client *Client) DescribeWebInstanceLogsWithContext(ctx context.Context, ApplicationId *string, InstanceId *string, request *DescribeWebInstanceLogsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeWebInstanceLogsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWebInstanceLogs"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v2/api/web/applications-observability/" + dara.PercentEncode(dara.StringValue(ApplicationId)) + "/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/logs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWebInstanceLogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DisableApplicationScalingRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableApplicationScalingRuleResponse
func (client *Client) DisableApplicationScalingRuleWithContext(ctx context.Context, request *DisableApplicationScalingRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DisableApplicationScalingRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ScalingRuleName) {
		query["ScalingRuleName"] = request.ScalingRuleName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableApplicationScalingRule"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/scale/disableApplicationScalingRule"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableApplicationScalingRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the advanced monitoring feature of Application Real-Time Monitoring Service (ARMS).
//
// @param request - DowngradeApplicationApmServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DowngradeApplicationApmServiceResponse
func (client *Client) DowngradeApplicationApmServiceWithContext(ctx context.Context, request *DowngradeApplicationApmServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DowngradeApplicationApmServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DowngradeApplicationApmService"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/app/applicationApmService"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DowngradeApplicationApmServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables an auto scaling policy for an application.
//
// @param request - EnableApplicationScalingRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableApplicationScalingRuleResponse
func (client *Client) EnableApplicationScalingRuleWithContext(ctx context.Context, request *EnableApplicationScalingRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EnableApplicationScalingRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ScalingRuleName) {
		query["ScalingRuleName"] = request.ScalingRuleName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableApplicationScalingRule"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/scale/enableApplicationScalingRule"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableApplicationScalingRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ExecJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecJobResponse
func (client *Client) ExecJobWithContext(ctx context.Context, request *ExecJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Command) {
		query["Command"] = request.Command
	}

	if !dara.IsNil(request.CommandArgs) {
		query["CommandArgs"] = request.CommandArgs
	}

	if !dara.IsNil(request.Envs) {
		query["Envs"] = request.Envs
	}

	if !dara.IsNil(request.EventId) {
		query["EventId"] = request.EventId
	}

	if !dara.IsNil(request.JarStartArgs) {
		query["JarStartArgs"] = request.JarStartArgs
	}

	if !dara.IsNil(request.JarStartOptions) {
		query["JarStartOptions"] = request.JarStartOptions
	}

	if !dara.IsNil(request.Replicas) {
		query["Replicas"] = request.Replicas
	}

	if !dara.IsNil(request.Time) {
		query["Time"] = request.Time
	}

	if !dara.IsNil(request.WarStartOptions) {
		query["WarStartOptions"] = request.WarStartOptions
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecJob"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/job/execJob"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the basic information of an application.
//
// @param request - GetApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationResponse
func (client *Client) GetApplicationWithContext(ctx context.Context, request *GetApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetApplicationResponse, _err error) {
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

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApplication"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/app/getApplication"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The application name.
//
// @param request - GetArmsTopNMetricRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetArmsTopNMetricResponse
func (client *Client) GetArmsTopNMetricWithContext(ctx context.Context, request *GetArmsTopNMetricRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetArmsTopNMetricResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppSource) {
		query["AppSource"] = request.AppSource
	}

	if !dara.IsNil(request.CpuStrategy) {
		query["CpuStrategy"] = request.CpuStrategy
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetArmsTopNMetric"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/getArmsTopNMetric"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetArmsTopNMetricResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the top N applications in which abnormal instances exist. The applications are sorted by the total number of abnormal instances.
//
// @param request - GetAvailabilityMetricRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAvailabilityMetricResponse
func (client *Client) GetAvailabilityMetricWithContext(ctx context.Context, request *GetAvailabilityMetricRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAvailabilityMetricResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppSource) {
		query["AppSource"] = request.AppSource
	}

	if !dara.IsNil(request.CpuStrategy) {
		query["CpuStrategy"] = request.CpuStrategy
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAvailabilityMetric"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/getAvailabilityMetric"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAvailabilityMetricResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries top N applications in abnormal change orders.
//
// @param request - GetChangeOrderMetricRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChangeOrderMetricResponse
func (client *Client) GetChangeOrderMetricWithContext(ctx context.Context, request *GetChangeOrderMetricRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetChangeOrderMetricResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppSource) {
		query["AppSource"] = request.AppSource
	}

	if !dara.IsNil(request.CoType) {
		query["CoType"] = request.CoType
	}

	if !dara.IsNil(request.CpuStrategy) {
		query["CpuStrategy"] = request.CpuStrategy
	}

	if !dara.IsNil(request.CreateTime) {
		query["CreateTime"] = request.CreateTime
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetChangeOrderMetric"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/getChangeOrderMetric"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetChangeOrderMetricResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the top N applications in which auto scaling takes effect.
//
// @param request - GetScaleAppMetricRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetScaleAppMetricResponse
func (client *Client) GetScaleAppMetricWithContext(ctx context.Context, request *GetScaleAppMetricRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetScaleAppMetricResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppSource) {
		query["AppSource"] = request.AppSource
	}

	if !dara.IsNil(request.CpuStrategy) {
		query["CpuStrategy"] = request.CpuStrategy
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetScaleAppMetric"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/getScaleAppMetric"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetScaleAppMetricResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The number of Warning events.
//
// @param request - GetWarningEventMetricRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWarningEventMetricResponse
func (client *Client) GetWarningEventMetricWithContext(ctx context.Context, request *GetWarningEventMetricRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetWarningEventMetricResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppSource) {
		query["AppSource"] = request.AppSource
	}

	if !dara.IsNil(request.CpuStrategy) {
		query["CpuStrategy"] = request.CpuStrategy
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWarningEventMetric"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/getWarningEventMetric"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWarningEventMetricResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the token used to remotely log on to the Webshell of an instance.
//
// @param request - GetWebshellTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWebshellTokenResponse
func (client *Client) GetWebshellTokenWithContext(ctx context.Context, request *GetWebshellTokenRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetWebshellTokenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ContainerName) {
		query["ContainerName"] = request.ContainerName
	}

	if !dara.IsNil(request.PodName) {
		query["PodName"] = request.PodName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWebshellToken"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/instance/webshellToken"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWebshellTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询所有泳道组
//
// @param request - ListAllSwimmingLaneGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAllSwimmingLaneGroupsResponse
func (client *Client) ListAllSwimmingLaneGroupsWithContext(ctx context.Context, request *ListAllSwimmingLaneGroupsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAllSwimmingLaneGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAllSwimmingLaneGroups"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/cas/gray/listSwimmingLaneGroups"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAllSwimmingLaneGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询所有泳道
//
// @param request - ListAllSwimmingLanesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAllSwimmingLanesResponse
func (client *Client) ListAllSwimmingLanesWithContext(ctx context.Context, request *ListAllSwimmingLanesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAllSwimmingLanesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAllSwimmingLanes"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/cas/gray/listSwimmingLanes"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAllSwimmingLanesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the events that occurred in an application.
//
// @param request - ListAppEventsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppEventsResponse
func (client *Client) ListAppEventsWithContext(ctx context.Context, request *ListAppEventsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAppEventsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EventType) {
		query["EventType"] = request.EventType
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.ObjectKind) {
		query["ObjectKind"] = request.ObjectKind
	}

	if !dara.IsNil(request.ObjectName) {
		query["ObjectName"] = request.ObjectName
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Reason) {
		query["Reason"] = request.Reason
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAppEvents"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/app/listAppEvents"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAppEventsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of microservices.
//
// @param request - ListAppServicesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppServicesResponse
func (client *Client) ListAppServicesWithContext(ctx context.Context, request *ListAppServicesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAppServicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.NacosInstanceId) {
		query["NacosInstanceId"] = request.NacosInstanceId
	}

	if !dara.IsNil(request.NacosNamespaceId) {
		query["NacosNamespaceId"] = request.NacosNamespaceId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegistryType) {
		query["RegistryType"] = request.RegistryType
	}

	if !dara.IsNil(request.ServiceType) {
		query["ServiceType"] = request.ServiceType
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAppServices"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/service/listAppServices"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAppServicesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the services of an application.
//
// @param request - ListAppServicesPageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppServicesPageResponse
func (client *Client) ListAppServicesPageWithContext(ctx context.Context, request *ListAppServicesPageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAppServicesPageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ServiceType) {
		query["ServiceType"] = request.ServiceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAppServicesPage"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/service/listAppServicesPage"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAppServicesPageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the deployment versions of an application.
//
// @param request - ListAppVersionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppVersionsResponse
func (client *Client) ListAppVersionsWithContext(ctx context.Context, request *ListAppVersionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAppVersionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAppVersions"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/app/listAppVersions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAppVersionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of applications.
//
// @param request - ListApplicationsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationsResponse
func (client *Client) ListApplicationsWithContext(ctx context.Context, request *ListApplicationsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListApplicationsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.AppSource) {
		query["AppSource"] = request.AppSource
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.FieldType) {
		query["FieldType"] = request.FieldType
	}

	if !dara.IsNil(request.FieldValue) {
		query["FieldValue"] = request.FieldValue
	}

	if !dara.IsNil(request.IsStateful) {
		query["IsStateful"] = request.IsStateful
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Reverse) {
		query["Reverse"] = request.Reverse
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApplications"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/app/listApplications"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApplicationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取应用列表，供全链路灰度拉取应用列表
//
// @param request - ListApplicationsForSwimmingLaneRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationsForSwimmingLaneResponse
func (client *Client) ListApplicationsForSwimmingLaneWithContext(ctx context.Context, request *ListApplicationsForSwimmingLaneRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListApplicationsForSwimmingLaneResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApplicationsForSwimmingLane"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/cas/gray/listApplicationsForSwimmingLane"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApplicationsForSwimmingLaneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query a list of change orders.
//
// @param request - ListChangeOrdersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListChangeOrdersResponse
func (client *Client) ListChangeOrdersWithContext(ctx context.Context, request *ListChangeOrdersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListChangeOrdersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.CoStatus) {
		query["CoStatus"] = request.CoStatus
	}

	if !dara.IsNil(request.CoType) {
		query["CoType"] = request.CoType
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Reverse) {
		query["Reverse"] = request.Reverse
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListChangeOrders"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/changeorder/ListChangeOrders"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListChangeOrdersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of microservices that are subscribed.
//
// @param request - ListConsumedServicesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConsumedServicesResponse
func (client *Client) ListConsumedServicesWithContext(ctx context.Context, request *ListConsumedServicesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListConsumedServicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListConsumedServices"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/service/listConsumedServices"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConsumedServicesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a canary release rule based on an application ID.
//
// Description:
//
// >  You can configure only one canary release rule for each application.
//
// @param request - ListGreyTagRouteRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGreyTagRouteResponse
func (client *Client) ListGreyTagRouteWithContext(ctx context.Context, request *ListGreyTagRouteRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListGreyTagRouteResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGreyTagRoute"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/tagroute/greyTagRouteList"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGreyTagRouteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Use ListIngress API call to query Ingress list
//
// @param request - ListIngressesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIngressesResponse
func (client *Client) ListIngressesWithContext(ctx context.Context, request *ListIngressesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIngressesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIngresses"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/ingress/IngressList"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIngressesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about job templates.
//
// @param request - ListJobsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobsResponse
func (client *Client) ListJobsWithContext(ctx context.Context, request *ListJobsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListJobsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.FieldType) {
		query["FieldType"] = request.FieldType
	}

	if !dara.IsNil(request.FieldValue) {
		query["FieldValue"] = request.FieldValue
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Reverse) {
		query["Reverse"] = request.Reverse
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.Workload) {
		query["Workload"] = request.Workload
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListJobs"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/job/listJobs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of application logs.
//
// @param request - ListLogConfigsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLogConfigsResponse
func (client *Client) ListLogConfigsWithContext(ctx context.Context, request *ListLogConfigsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListLogConfigsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLogConfigs"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/log/listLogConfigs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLogConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of change orders in a namespace.
//
// @param request - ListNamespaceChangeOrdersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNamespaceChangeOrdersResponse
func (client *Client) ListNamespaceChangeOrdersWithContext(ctx context.Context, request *ListNamespaceChangeOrdersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListNamespaceChangeOrdersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CoStatus) {
		query["CoStatus"] = request.CoStatus
	}

	if !dara.IsNil(request.CoType) {
		query["CoType"] = request.CoType
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNamespaceChangeOrders"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/changeorder/listNamespaceChangeOrders"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNamespaceChangeOrdersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the ConfigMap instances in a namespace.
//
// @param request - ListNamespacedConfigMapsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNamespacedConfigMapsResponse
func (client *Client) ListNamespacedConfigMapsWithContext(ctx context.Context, request *ListNamespacedConfigMapsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListNamespacedConfigMapsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNamespacedConfigMaps"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/configmap/listNamespacedConfigMaps"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNamespacedConfigMapsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of microservices that are published.
//
// @param request - ListPublishedServicesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPublishedServicesResponse
func (client *Client) ListPublishedServicesWithContext(ctx context.Context, request *ListPublishedServicesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPublishedServicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPublishedServices"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/service/listPublishedServices"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPublishedServicesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about Secrets in a namespace.
//
// @param request - ListSecretsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSecretsResponse
func (client *Client) ListSecretsWithContext(ctx context.Context, request *ListSecretsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSecretsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSecrets"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/secret/secrets"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSecretsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询泳道可选的网关路由
//
// @param request - ListSwimmingLaneGatewayRoutesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSwimmingLaneGatewayRoutesResponse
func (client *Client) ListSwimmingLaneGatewayRoutesWithContext(ctx context.Context, request *ListSwimmingLaneGatewayRoutesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSwimmingLaneGatewayRoutesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSwimmingLaneGatewayRoutes"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/cas/gray/listSwimmingLaneGatewayRoutes"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSwimmingLaneGatewayRoutesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询所有泳道标签列表
//
// @param request - ListSwimmingLaneGroupTagsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSwimmingLaneGroupTagsResponse
func (client *Client) ListSwimmingLaneGroupTagsWithContext(ctx context.Context, request *ListSwimmingLaneGroupTagsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSwimmingLaneGroupTagsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSwimmingLaneGroupTags"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/cas/gray/listSwimmingLaneGroupTags"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSwimmingLaneGroupTagsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the mapping relationships between applications and tags.
//
// @param request - ListTagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithContext(ctx context.Context, request *ListTagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
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

	if !dara.IsNil(request.ResourceIds) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tags"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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
// Query the list of web application instances.
//
// Description:
//
// Query the list of web application instances.
//
// @param tmpReq - ListWebApplicationInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWebApplicationInstancesResponse
func (client *Client) ListWebApplicationInstancesWithContext(ctx context.Context, ApplicationId *string, tmpReq *ListWebApplicationInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListWebApplicationInstancesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListWebApplicationInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Statuses) {
		request.StatusesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Statuses, dara.String("Statuses"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.VersionIds) {
		request.VersionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VersionIds, dara.String("VersionIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StatusesShrink) {
		query["Statuses"] = request.StatusesShrink
	}

	if !dara.IsNil(request.VersionIdsShrink) {
		query["VersionIds"] = request.VersionIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWebApplicationInstances"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v2/api/web/applications-observability/" + dara.PercentEncode(dara.StringValue(ApplicationId)) + "/instances"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWebApplicationInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the list of web application versions.
//
// Description:
//
// Query the list of web application versions.
//
// @param request - ListWebApplicationRevisionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWebApplicationRevisionsResponse
func (client *Client) ListWebApplicationRevisionsWithContext(ctx context.Context, ApplicationId *string, request *ListWebApplicationRevisionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListWebApplicationRevisionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWebApplicationRevisions"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v2/api/web/application-revisions/" + dara.PercentEncode(dara.StringValue(ApplicationId)) + "/revisions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWebApplicationRevisionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the list of web applications.
//
// Description:
//
// Call the ListWebApplications operation to query the list of web applications.
//
// @param request - ListWebApplicationsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWebApplicationsResponse
func (client *Client) ListWebApplicationsWithContext(ctx context.Context, request *ListWebApplicationsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListWebApplicationsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Prefix) {
		query["Prefix"] = request.Prefix
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWebApplications"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v2/api/web/applications"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWebApplicationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query available custom domain names.
//
// Description:
//
// Query available custom domain names.
//
// @param request - ListWebCustomDomainsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWebCustomDomainsResponse
func (client *Client) ListWebCustomDomainsWithContext(ctx context.Context, request *ListWebCustomDomainsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListWebCustomDomainsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Prefix) {
		query["Prefix"] = request.Prefix
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWebCustomDomains"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v2/api/web/custom-domains"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWebCustomDomainsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Activates Serverless App Engine (SAE) for free.
//
// Description:
//
// > Make sure that your account balance is greater than 0. Otherwise, the SAE service cannot be activated.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenSaeServiceResponse
func (client *Client) OpenSaeServiceWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OpenSaeServiceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("OpenSaeService"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/service/open"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenSaeServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publish a web application version.
//
// Description:
//
// Publish a web application version. You can change the configurations of the version and create a new version.
//
// @param request - PublishWebApplicationRevisionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishWebApplicationRevisionResponse
func (client *Client) PublishWebApplicationRevisionWithContext(ctx context.Context, ApplicationId *string, request *PublishWebApplicationRevisionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PublishWebApplicationRevisionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishWebApplicationRevision"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v2/api/web/application-revisions/" + dara.PercentEncode(dara.StringValue(ApplicationId)) + "/revisions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishWebApplicationRevisionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resource usage of an application.
//
// @param request - QueryResourceStaticsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryResourceStaticsResponse
func (client *Client) QueryResourceStaticsWithContext(ctx context.Context, request *QueryResourceStaticsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryResourceStaticsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryResourceStatics"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/paas/quota/queryResourceStatics"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryResourceStaticsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Scales in an application based on instance IDs.
//
// @param request - ReduceApplicationCapacityByInstanceIdsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReduceApplicationCapacityByInstanceIdsResponse
func (client *Client) ReduceApplicationCapacityByInstanceIdsWithContext(ctx context.Context, request *ReduceApplicationCapacityByInstanceIdsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ReduceApplicationCapacityByInstanceIdsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReduceApplicationCapacityByInstanceIds"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/app/ScaleInApplicationWithInstanceIds"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ReduceApplicationCapacityByInstanceIdsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Rescale an application.
//
// @param request - RescaleApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RescaleApplicationResponse
func (client *Client) RescaleApplicationWithContext(ctx context.Context, request *RescaleApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RescaleApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AutoEnableApplicationScalingRule) {
		query["AutoEnableApplicationScalingRule"] = request.AutoEnableApplicationScalingRule
	}

	if !dara.IsNil(request.MinReadyInstanceRatio) {
		query["MinReadyInstanceRatio"] = request.MinReadyInstanceRatio
	}

	if !dara.IsNil(request.MinReadyInstances) {
		query["MinReadyInstances"] = request.MinReadyInstances
	}

	if !dara.IsNil(request.Replicas) {
		query["Replicas"] = request.Replicas
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RescaleApplication"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/app/rescaleApplication"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RescaleApplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the instance specifications of an application.
//
// @param request - RescaleApplicationVerticallyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RescaleApplicationVerticallyResponse
func (client *Client) RescaleApplicationVerticallyWithContext(ctx context.Context, request *RescaleApplicationVerticallyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RescaleApplicationVerticallyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Cpu) {
		query["Cpu"] = request.Cpu
	}

	if !dara.IsNil(request.DiskSize) {
		query["DiskSize"] = request.DiskSize
	}

	if !dara.IsNil(request.Memory) {
		query["Memory"] = request.Memory
	}

	if !dara.IsNil(request.AutoEnableApplicationScalingRule) {
		query["autoEnableApplicationScalingRule"] = request.AutoEnableApplicationScalingRule
	}

	if !dara.IsNil(request.MinReadyInstanceRatio) {
		query["minReadyInstanceRatio"] = request.MinReadyInstanceRatio
	}

	if !dara.IsNil(request.MinReadyInstances) {
		query["minReadyInstances"] = request.MinReadyInstances
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RescaleApplicationVertically"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/app/rescaleApplicationVertically"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RescaleApplicationVerticallyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts an application.
//
// @param request - RestartApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartApplicationResponse
func (client *Client) RestartApplicationWithContext(ctx context.Context, request *RestartApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RestartApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AutoEnableApplicationScalingRule) {
		query["AutoEnableApplicationScalingRule"] = request.AutoEnableApplicationScalingRule
	}

	if !dara.IsNil(request.MinReadyInstanceRatio) {
		query["MinReadyInstanceRatio"] = request.MinReadyInstanceRatio
	}

	if !dara.IsNil(request.MinReadyInstances) {
		query["MinReadyInstances"] = request.MinReadyInstances
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartApplication"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/app/restartApplication"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartApplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts one or more instances in an application.
//
// @param request - RestartInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartInstancesResponse
func (client *Client) RestartInstancesWithContext(ctx context.Context, request *RestartInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RestartInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartInstances"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/app/restartInstances"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Rolls back an application.
//
// @param request - RollbackApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RollbackApplicationResponse
func (client *Client) RollbackApplicationWithContext(ctx context.Context, request *RollbackApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RollbackApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AutoEnableApplicationScalingRule) {
		query["AutoEnableApplicationScalingRule"] = request.AutoEnableApplicationScalingRule
	}

	if !dara.IsNil(request.BatchWaitTime) {
		query["BatchWaitTime"] = request.BatchWaitTime
	}

	if !dara.IsNil(request.MinReadyInstanceRatio) {
		query["MinReadyInstanceRatio"] = request.MinReadyInstanceRatio
	}

	if !dara.IsNil(request.MinReadyInstances) {
		query["MinReadyInstances"] = request.MinReadyInstances
	}

	if !dara.IsNil(request.UpdateStrategy) {
		query["UpdateStrategy"] = request.UpdateStrategy
	}

	if !dara.IsNil(request.VersionId) {
		query["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RollbackApplication"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/app/rollbackApplication"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RollbackApplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts an application.
//
// @param request - StartApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartApplicationResponse
func (client *Client) StartApplicationWithContext(ctx context.Context, request *StartApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartApplication"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/app/startApplication"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StartApplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Start a web application.
//
// Description:
//
// Call the StartWebApplication operation to start a web application.
//
// @param request - StartWebApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartWebApplicationResponse
func (client *Client) StartWebApplicationWithContext(ctx context.Context, ApplicationId *string, request *StartWebApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartWebApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartWebApplication"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v2/api/web/application-ops/" + dara.PercentEncode(dara.StringValue(ApplicationId)) + "/start"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StartWebApplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops an application.
//
// @param request - StopApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopApplicationResponse
func (client *Client) StopApplicationWithContext(ctx context.Context, request *StopApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopApplication"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/app/stopApplication"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopApplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stop a web application.
//
// Description:
//
// Call the StopWebApplication operation to stop a web application.
//
// @param request - StopWebApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopWebApplicationResponse
func (client *Client) StopWebApplicationWithContext(ctx context.Context, ApplicationId *string, request *StopWebApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopWebApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopWebApplication"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v2/api/web/application-ops/" + dara.PercentEncode(dara.StringValue(ApplicationId)) + "/stop"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopWebApplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Suspends one or more jobs.
//
// @param request - SuspendJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SuspendJobResponse
func (client *Client) SuspendJobWithContext(ctx context.Context, request *SuspendJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SuspendJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Suspend) {
		query["Suspend"] = request.Suspend
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SuspendJob"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/job/suspendJob"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SuspendJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds tags to resources.
//
// @param request - TagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithContext(ctx context.Context, request *TagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceIds) {
		body["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tags) {
		body["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tags"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
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

// Summary:
//
// # Calls the UnbindNlb operation to delete an NLB listener bound for application access
//
// @param request - UnbindNlbRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindNlbResponse
func (client *Client) UnbindNlbWithContext(ctx context.Context, request *UnbindNlbRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UnbindNlbResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.NlbId) {
		query["NlbId"] = request.NlbId
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindNlb"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/app/nlb"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindNlbResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates an internal-facing or Internet-facing SLB instance from an application.
//
// @param request - UnbindSlbRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindSlbResponse
func (client *Client) UnbindSlbWithContext(ctx context.Context, request *UnbindSlbRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UnbindSlbResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Internet) {
		query["Internet"] = request.Internet
	}

	if !dara.IsNil(request.Intranet) {
		query["Intranet"] = request.Intranet
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindSlb"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/app/slb"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindSlbResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes tags from resources.
//
// @param request - UntagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithContext(ctx context.Context, request *UntagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeleteAll) {
		query["DeleteAll"] = request.DeleteAll
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceIds) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKeys) {
		query["TagKeys"] = request.TagKeys
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tags"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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
// 应用闲置模式更新
//
// @param request - UpdateAppModeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAppModeResponse
func (client *Client) UpdateAppModeWithContext(ctx context.Context, request *UpdateAppModeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAppModeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppIds) {
		query["AppIds"] = request.AppIds
	}

	if !dara.IsNil(request.EnableIdle) {
		query["EnableIdle"] = request.EnableIdle
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAppMode"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/app/updateAppMode"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAppModeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the security group of an application.
//
// @param request - UpdateAppSecurityGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAppSecurityGroupResponse
func (client *Client) UpdateAppSecurityGroupWithContext(ctx context.Context, request *UpdateAppSecurityGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAppSecurityGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAppSecurityGroup"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/app/updateAppSecurityGroup"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAppSecurityGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the description of an application.
//
// @param request - UpdateApplicationDescriptionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApplicationDescriptionResponse
func (client *Client) UpdateApplicationDescriptionWithContext(ctx context.Context, request *UpdateApplicationDescriptionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateApplicationDescriptionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppDescription) {
		query["AppDescription"] = request.AppDescription
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateApplicationDescription"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/app/updateAppDescription"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApplicationDescriptionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the auto scaling policy of an application.
//
// Description:
//
// ##
//
// If you want to configure more than 50 instances for an application, you must submit a [ticket](https://workorder.console.aliyun.com/#/ticket/createIndex) to add your account to the whitelist.
//
// @param request - UpdateApplicationScalingRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApplicationScalingRuleResponse
func (client *Client) UpdateApplicationScalingRuleWithContext(ctx context.Context, request *UpdateApplicationScalingRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateApplicationScalingRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EnableIdle) {
		query["EnableIdle"] = request.EnableIdle
	}

	if !dara.IsNil(request.MinReadyInstanceRatio) {
		query["MinReadyInstanceRatio"] = request.MinReadyInstanceRatio
	}

	if !dara.IsNil(request.MinReadyInstances) {
		query["MinReadyInstances"] = request.MinReadyInstances
	}

	if !dara.IsNil(request.ScalingRuleMetric) {
		query["ScalingRuleMetric"] = request.ScalingRuleMetric
	}

	if !dara.IsNil(request.ScalingRuleName) {
		query["ScalingRuleName"] = request.ScalingRuleName
	}

	if !dara.IsNil(request.ScalingRuleTimer) {
		query["ScalingRuleTimer"] = request.ScalingRuleTimer
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateApplicationScalingRule"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/scale/applicationScalingRule"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApplicationScalingRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update the configuration of a vSwitch.
//
// @param request - UpdateApplicationVswitchesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApplicationVswitchesResponse
func (client *Client) UpdateApplicationVswitchesWithContext(ctx context.Context, request *UpdateApplicationVswitchesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateApplicationVswitchesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateApplicationVswitches"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/app/updateAppVswitches"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApplicationVswitchesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a ConfigMap instance.
//
// @param request - UpdateConfigMapRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConfigMapResponse
func (client *Client) UpdateConfigMapWithContext(ctx context.Context, request *UpdateConfigMapRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateConfigMapResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigMapId) {
		query["ConfigMapId"] = request.ConfigMapId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Data) {
		body["Data"] = request.Data
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateConfigMap"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/configmap/configMap"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateConfigMapResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a canary release rule.
//
// @param request - UpdateGreyTagRouteRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGreyTagRouteResponse
func (client *Client) UpdateGreyTagRouteWithContext(ctx context.Context, request *UpdateGreyTagRouteRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateGreyTagRouteResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlbRules) {
		query["AlbRules"] = request.AlbRules
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DubboRules) {
		query["DubboRules"] = request.DubboRules
	}

	if !dara.IsNil(request.GreyTagRouteId) {
		query["GreyTagRouteId"] = request.GreyTagRouteId
	}

	if !dara.IsNil(request.ScRules) {
		query["ScRules"] = request.ScRules
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGreyTagRoute"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/tagroute/greyTagRoute"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGreyTagRouteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update the configurations of an Ingress instance.
//
// @param request - UpdateIngressRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIngressResponse
func (client *Client) UpdateIngressWithContext(ctx context.Context, request *UpdateIngressRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateIngressResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertId) {
		query["CertId"] = request.CertId
	}

	if !dara.IsNil(request.CertIds) {
		query["CertIds"] = request.CertIds
	}

	if !dara.IsNil(request.CorsConfig) {
		query["CorsConfig"] = request.CorsConfig
	}

	if !dara.IsNil(request.DefaultRule) {
		query["DefaultRule"] = request.DefaultRule
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EnableXForwardedFor) {
		query["EnableXForwardedFor"] = request.EnableXForwardedFor
	}

	if !dara.IsNil(request.EnableXForwardedForClientSrcPort) {
		query["EnableXForwardedForClientSrcPort"] = request.EnableXForwardedForClientSrcPort
	}

	if !dara.IsNil(request.EnableXForwardedForProto) {
		query["EnableXForwardedForProto"] = request.EnableXForwardedForProto
	}

	if !dara.IsNil(request.EnableXForwardedForSlbId) {
		query["EnableXForwardedForSlbId"] = request.EnableXForwardedForSlbId
	}

	if !dara.IsNil(request.EnableXForwardedForSlbPort) {
		query["EnableXForwardedForSlbPort"] = request.EnableXForwardedForSlbPort
	}

	if !dara.IsNil(request.IdleTimeout) {
		query["IdleTimeout"] = request.IdleTimeout
	}

	if !dara.IsNil(request.IngressId) {
		query["IngressId"] = request.IngressId
	}

	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.ListenerProtocol) {
		query["ListenerProtocol"] = request.ListenerProtocol
	}

	if !dara.IsNil(request.LoadBalanceType) {
		query["LoadBalanceType"] = request.LoadBalanceType
	}

	if !dara.IsNil(request.RequestTimeout) {
		query["RequestTimeout"] = request.RequestTimeout
	}

	if !dara.IsNil(request.SecurityPolicyId) {
		query["SecurityPolicyId"] = request.SecurityPolicyId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Rules) {
		body["Rules"] = request.Rules
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateIngress"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/ingress/Ingress"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateIngressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a job template.
//
// @param request - UpdateJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateJobResponse
func (client *Client) UpdateJobWithContext(ctx context.Context, request *UpdateJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcrAssumeRoleArn) {
		query["AcrAssumeRoleArn"] = request.AcrAssumeRoleArn
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.BackoffLimit) {
		query["BackoffLimit"] = request.BackoffLimit
	}

	if !dara.IsNil(request.Command) {
		query["Command"] = request.Command
	}

	if !dara.IsNil(request.CommandArgs) {
		query["CommandArgs"] = request.CommandArgs
	}

	if !dara.IsNil(request.ConcurrencyPolicy) {
		query["ConcurrencyPolicy"] = request.ConcurrencyPolicy
	}

	if !dara.IsNil(request.CustomHostAlias) {
		query["CustomHostAlias"] = request.CustomHostAlias
	}

	if !dara.IsNil(request.EdasContainerVersion) {
		query["EdasContainerVersion"] = request.EdasContainerVersion
	}

	if !dara.IsNil(request.Envs) {
		query["Envs"] = request.Envs
	}

	if !dara.IsNil(request.ImagePullSecrets) {
		query["ImagePullSecrets"] = request.ImagePullSecrets
	}

	if !dara.IsNil(request.ImageUrl) {
		query["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.JarStartArgs) {
		query["JarStartArgs"] = request.JarStartArgs
	}

	if !dara.IsNil(request.JarStartOptions) {
		query["JarStartOptions"] = request.JarStartOptions
	}

	if !dara.IsNil(request.Jdk) {
		query["Jdk"] = request.Jdk
	}

	if !dara.IsNil(request.MountDesc) {
		query["MountDesc"] = request.MountDesc
	}

	if !dara.IsNil(request.MountHost) {
		query["MountHost"] = request.MountHost
	}

	if !dara.IsNil(request.NasId) {
		query["NasId"] = request.NasId
	}

	if !dara.IsNil(request.PackageUrl) {
		query["PackageUrl"] = request.PackageUrl
	}

	if !dara.IsNil(request.PackageVersion) {
		query["PackageVersion"] = request.PackageVersion
	}

	if !dara.IsNil(request.PhpConfigLocation) {
		query["PhpConfigLocation"] = request.PhpConfigLocation
	}

	if !dara.IsNil(request.PostStart) {
		query["PostStart"] = request.PostStart
	}

	if !dara.IsNil(request.PreStop) {
		query["PreStop"] = request.PreStop
	}

	if !dara.IsNil(request.ProgrammingLanguage) {
		query["ProgrammingLanguage"] = request.ProgrammingLanguage
	}

	if !dara.IsNil(request.Python) {
		query["Python"] = request.Python
	}

	if !dara.IsNil(request.PythonModules) {
		query["PythonModules"] = request.PythonModules
	}

	if !dara.IsNil(request.RefAppId) {
		query["RefAppId"] = request.RefAppId
	}

	if !dara.IsNil(request.Replicas) {
		query["Replicas"] = request.Replicas
	}

	if !dara.IsNil(request.Slice) {
		query["Slice"] = request.Slice
	}

	if !dara.IsNil(request.SliceEnvs) {
		query["SliceEnvs"] = request.SliceEnvs
	}

	if !dara.IsNil(request.SlsConfigs) {
		query["SlsConfigs"] = request.SlsConfigs
	}

	if !dara.IsNil(request.TerminationGracePeriodSeconds) {
		query["TerminationGracePeriodSeconds"] = request.TerminationGracePeriodSeconds
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	if !dara.IsNil(request.Timezone) {
		query["Timezone"] = request.Timezone
	}

	if !dara.IsNil(request.TomcatConfig) {
		query["TomcatConfig"] = request.TomcatConfig
	}

	if !dara.IsNil(request.TriggerConfig) {
		query["TriggerConfig"] = request.TriggerConfig
	}

	if !dara.IsNil(request.WarStartOptions) {
		query["WarStartOptions"] = request.WarStartOptions
	}

	if !dara.IsNil(request.WebContainer) {
		query["WebContainer"] = request.WebContainer
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AcrInstanceId) {
		body["AcrInstanceId"] = request.AcrInstanceId
	}

	if !dara.IsNil(request.ConfigMapMountDesc) {
		body["ConfigMapMountDesc"] = request.ConfigMapMountDesc
	}

	if !dara.IsNil(request.EnableImageAccl) {
		body["EnableImageAccl"] = request.EnableImageAccl
	}

	if !dara.IsNil(request.OssAkId) {
		body["OssAkId"] = request.OssAkId
	}

	if !dara.IsNil(request.OssAkSecret) {
		body["OssAkSecret"] = request.OssAkSecret
	}

	if !dara.IsNil(request.OssMountDescs) {
		body["OssMountDescs"] = request.OssMountDescs
	}

	if !dara.IsNil(request.Php) {
		body["Php"] = request.Php
	}

	if !dara.IsNil(request.PhpConfig) {
		body["PhpConfig"] = request.PhpConfig
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateJob"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/job/updateJob"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a namespace.
//
// @param request - UpdateNamespaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNamespaceResponse
func (client *Client) UpdateNamespaceWithContext(ctx context.Context, request *UpdateNamespaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateNamespaceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnableMicroRegistration) {
		query["EnableMicroRegistration"] = request.EnableMicroRegistration
	}

	if !dara.IsNil(request.NameSpaceShortId) {
		query["NameSpaceShortId"] = request.NameSpaceShortId
	}

	if !dara.IsNil(request.NamespaceDescription) {
		query["NamespaceDescription"] = request.NamespaceDescription
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.NamespaceName) {
		query["NamespaceName"] = request.NamespaceName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateNamespace"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/paas/namespace"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateNamespaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// cn-beijing:test
//
// @param request - UpdateNamespaceVpcRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNamespaceVpcResponse
func (client *Client) UpdateNamespaceVpcWithContext(ctx context.Context, request *UpdateNamespaceVpcRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateNamespaceVpcResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NameSpaceShortId) {
		query["NameSpaceShortId"] = request.NameSpaceShortId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateNamespaceVpc"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/namespace/updateNamespaceVpc"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateNamespaceVpcResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The HTTP status code. Valid values:
//
//   - **2xx**: The call was successful.
//
//   - **3xx**: The call was redirected.
//
//   - **4xx**: The call failed.
//
//   - **5xx**: A server error occurred.
//
// @param tmpReq - UpdateSecretRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSecretResponse
func (client *Client) UpdateSecretWithContext(ctx context.Context, tmpReq *UpdateSecretRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateSecretResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateSecretShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SecretData) {
		request.SecretDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SecretData, dara.String("SecretData"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.SecretDataShrink) {
		query["SecretData"] = request.SecretDataShrink
	}

	if !dara.IsNil(request.SecretId) {
		query["SecretId"] = request.SecretId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSecret"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/secret/secret"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSecretResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新泳道的启用属性
//
// @param request - UpdateSwimmingLaneEnableAttributeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSwimmingLaneEnableAttributeResponse
func (client *Client) UpdateSwimmingLaneEnableAttributeWithContext(ctx context.Context, request *UpdateSwimmingLaneEnableAttributeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateSwimmingLaneEnableAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.LaneId) {
		query["LaneId"] = request.LaneId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSwimmingLaneEnableAttribute"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/cas/gray/updateSwimmingLaneEnableAttribute"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSwimmingLaneEnableAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration at the web application level.
//
// Description:
//
// You can call the UpdateWebApplication operation to update the configuration at the web application level.
//
// @param request - UpdateWebApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWebApplicationResponse
func (client *Client) UpdateWebApplicationWithContext(ctx context.Context, ApplicationId *string, request *UpdateWebApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateWebApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWebApplication"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v2/api/web/applications/" + dara.PercentEncode(dara.StringValue(ApplicationId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWebApplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update the scaling configuration of a web application.
//
// Description:
//
// You can call the UpdateWebApplicationScalingConfig operation to update the scaling configurations of a web application.
//
// @param request - UpdateWebApplicationScalingConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWebApplicationScalingConfigResponse
func (client *Client) UpdateWebApplicationScalingConfigWithContext(ctx context.Context, ApplicationId *string, request *UpdateWebApplicationScalingConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateWebApplicationScalingConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWebApplicationScalingConfig"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v2/api/web/application-scaling/" + dara.PercentEncode(dara.StringValue(ApplicationId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWebApplicationScalingConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update the traffic configurations of a web application.
//
// Description:
//
// Call the UpdateWebApplicationTrafficConfig operation to update the traffic configurations of a web application.
//
// @param request - UpdateWebApplicationTrafficConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWebApplicationTrafficConfigResponse
func (client *Client) UpdateWebApplicationTrafficConfigWithContext(ctx context.Context, ApplicationId *string, request *UpdateWebApplicationTrafficConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateWebApplicationTrafficConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWebApplicationTrafficConfig"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v2/api/web/application-traffic/" + dara.PercentEncode(dara.StringValue(ApplicationId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWebApplicationTrafficConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update a custom domain name.
//
// Description:
//
// Update a custom domain name.
//
// @param request - UpdateWebCustomDomainRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWebCustomDomainResponse
func (client *Client) UpdateWebCustomDomainWithContext(ctx context.Context, DomainName *string, request *UpdateWebCustomDomainRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateWebCustomDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWebCustomDomain"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v2/api/web/custom-domains/" + dara.PercentEncode(dara.StringValue(DomainName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWebCustomDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the advanced monitoring feature of Application Real-Time Monitoring Service (ARMS).
//
// Description:
//
// You are charged when you use the ARMS advanced monitoring feature. Enable this feature based on your business requirements. For more information, see [Billing overview](https://icms.alibaba-inc.com/content/arms/arms?l=1\\&m=16992\\&n=3183148).
//
// @param request - UpgradeApplicationApmServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeApplicationApmServiceResponse
func (client *Client) UpgradeApplicationApmServiceWithContext(ctx context.Context, request *UpgradeApplicationApmServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpgradeApplicationApmServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeApplicationApmService"),
		Version:     dara.String("2019-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/sam/app/applicationApmService"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeApplicationApmServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
