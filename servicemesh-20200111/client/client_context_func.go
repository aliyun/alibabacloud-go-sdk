// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Adds a cluster to an ASM instance.
//
// @param request - AddClusterIntoServiceMeshRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddClusterIntoServiceMeshResponse
func (client *Client) AddClusterIntoServiceMeshWithContext(ctx context.Context, request *AddClusterIntoServiceMeshRequest, runtime *dara.RuntimeOptions) (_result *AddClusterIntoServiceMeshResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DiscoveryOnly) {
		body["DiscoveryOnly"] = request.DiscoveryOnly
	}

	if !dara.IsNil(request.IgnoreNamespaceCheck) {
		body["IgnoreNamespaceCheck"] = request.IgnoreNamespaceCheck
	}

	if !dara.IsNil(request.Kubeconfig) {
		body["Kubeconfig"] = request.Kubeconfig
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddClusterIntoServiceMesh"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddClusterIntoServiceMeshResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI AddVMIntoServiceMesh is deprecated
//
// Summary:
//
// Adds a virtual machine (VM) to a Service Mesh (ASM) instance.
//
// @param request - AddVMIntoServiceMeshRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddVMIntoServiceMeshResponse
func (client *Client) AddVMIntoServiceMeshWithContext(ctx context.Context, request *AddVMIntoServiceMeshRequest, runtime *dara.RuntimeOptions) (_result *AddVMIntoServiceMeshResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EcsId) {
		query["EcsId"] = request.EcsId
	}

	if !dara.IsNil(request.ServiceMeshId) {
		query["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddVMIntoServiceMesh"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddVMIntoServiceMeshResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Service Mesh (ASM) gateway.
//
// @param request - CreateASMGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateASMGatewayResponse
func (client *Client) CreateASMGatewayWithContext(ctx context.Context, request *CreateASMGatewayRequest, runtime *dara.RuntimeOptions) (_result *CreateASMGatewayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Body) {
		body["Body"] = request.Body
	}

	if !dara.IsNil(request.IstioGatewayName) {
		body["IstioGatewayName"] = request.IstioGatewayName
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateASMGateway"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateASMGatewayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a secret for a Service Mesh (ASM) gateway.
//
// @param request - CreateGatewaySecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGatewaySecretResponse
func (client *Client) CreateGatewaySecretWithContext(ctx context.Context, request *CreateGatewaySecretRequest, runtime *dara.RuntimeOptions) (_result *CreateGatewaySecretResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Cert) {
		body["Cert"] = request.Cert
	}

	if !dara.IsNil(request.IstioGatewayName) {
		body["IstioGatewayName"] = request.IstioGatewayName
	}

	if !dara.IsNil(request.Key) {
		body["Key"] = request.Key
	}

	if !dara.IsNil(request.SecretName) {
		body["SecretName"] = request.SecretName
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGatewaySecret"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGatewaySecretResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds domain names for a Service Mesh (ASM) gateway.
//
// @param request - CreateIstioGatewayDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIstioGatewayDomainsResponse
func (client *Client) CreateIstioGatewayDomainsWithContext(ctx context.Context, request *CreateIstioGatewayDomainsRequest, runtime *dara.RuntimeOptions) (_result *CreateIstioGatewayDomainsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Credential) {
		body["Credential"] = request.Credential
	}

	if !dara.IsNil(request.ForceHttps) {
		body["ForceHttps"] = request.ForceHttps
	}

	if !dara.IsNil(request.Hosts) {
		body["Hosts"] = request.Hosts
	}

	if !dara.IsNil(request.IstioGatewayName) {
		body["IstioGatewayName"] = request.IstioGatewayName
	}

	if !dara.IsNil(request.Limit) {
		body["Limit"] = request.Limit
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Number) {
		body["Number"] = request.Number
	}

	if !dara.IsNil(request.PortName) {
		body["PortName"] = request.PortName
	}

	if !dara.IsNil(request.Protocol) {
		body["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIstioGatewayDomains"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIstioGatewayDomainsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a routing rule for a Service Mesh (ASM) gateway.
//
// @param tmpReq - CreateIstioGatewayRoutesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIstioGatewayRoutesResponse
func (client *Client) CreateIstioGatewayRoutesWithContext(ctx context.Context, tmpReq *CreateIstioGatewayRoutesRequest, runtime *dara.RuntimeOptions) (_result *CreateIstioGatewayRoutesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateIstioGatewayRoutesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.GatewayRoute) {
		request.GatewayRouteShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GatewayRoute, dara.String("GatewayRoute"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.GatewayRouteShrink) {
		body["GatewayRoute"] = request.GatewayRouteShrink
	}

	if !dara.IsNil(request.IstioGatewayName) {
		body["IstioGatewayName"] = request.IstioGatewayName
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIstioGatewayRoutes"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIstioGatewayRoutesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Service Mesh (ASM) instance.
//
// @param request - CreateServiceMeshRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceMeshResponse
func (client *Client) CreateServiceMeshWithContext(ctx context.Context, request *CreateServiceMeshRequest, runtime *dara.RuntimeOptions) (_result *CreateServiceMeshResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccessLogEnabled) {
		body["AccessLogEnabled"] = request.AccessLogEnabled
	}

	if !dara.IsNil(request.AccessLogFile) {
		body["AccessLogFile"] = request.AccessLogFile
	}

	if !dara.IsNil(request.AccessLogFormat) {
		body["AccessLogFormat"] = request.AccessLogFormat
	}

	if !dara.IsNil(request.AccessLogProject) {
		body["AccessLogProject"] = request.AccessLogProject
	}

	if !dara.IsNil(request.AccessLogServiceEnabled) {
		body["AccessLogServiceEnabled"] = request.AccessLogServiceEnabled
	}

	if !dara.IsNil(request.AccessLogServiceHost) {
		body["AccessLogServiceHost"] = request.AccessLogServiceHost
	}

	if !dara.IsNil(request.AccessLogServicePort) {
		body["AccessLogServicePort"] = request.AccessLogServicePort
	}

	if !dara.IsNil(request.ApiServerLoadBalancerSpec) {
		body["ApiServerLoadBalancerSpec"] = request.ApiServerLoadBalancerSpec
	}

	if !dara.IsNil(request.ApiServerPublicEip) {
		body["ApiServerPublicEip"] = request.ApiServerPublicEip
	}

	if !dara.IsNil(request.AuditProject) {
		body["AuditProject"] = request.AuditProject
	}

	if !dara.IsNil(request.AutoRenew) {
		body["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.AutoRenewPeriod) {
		body["AutoRenewPeriod"] = request.AutoRenewPeriod
	}

	if !dara.IsNil(request.CRAggregationEnabled) {
		body["CRAggregationEnabled"] = request.CRAggregationEnabled
	}

	if !dara.IsNil(request.CertChain) {
		body["CertChain"] = request.CertChain
	}

	if !dara.IsNil(request.ChargeType) {
		body["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.ClusterDomain) {
		body["ClusterDomain"] = request.ClusterDomain
	}

	if !dara.IsNil(request.ClusterSpec) {
		body["ClusterSpec"] = request.ClusterSpec
	}

	if !dara.IsNil(request.ConfigSourceEnabled) {
		body["ConfigSourceEnabled"] = request.ConfigSourceEnabled
	}

	if !dara.IsNil(request.ConfigSourceNacosID) {
		body["ConfigSourceNacosID"] = request.ConfigSourceNacosID
	}

	if !dara.IsNil(request.ControlPlaneLogEnabled) {
		body["ControlPlaneLogEnabled"] = request.ControlPlaneLogEnabled
	}

	if !dara.IsNil(request.ControlPlaneLogProject) {
		body["ControlPlaneLogProject"] = request.ControlPlaneLogProject
	}

	if !dara.IsNil(request.CustomizedPrometheus) {
		body["CustomizedPrometheus"] = request.CustomizedPrometheus
	}

	if !dara.IsNil(request.CustomizedZipkin) {
		body["CustomizedZipkin"] = request.CustomizedZipkin
	}

	if !dara.IsNil(request.DNSProxyingEnabled) {
		body["DNSProxyingEnabled"] = request.DNSProxyingEnabled
	}

	if !dara.IsNil(request.DubboFilterEnabled) {
		body["DubboFilterEnabled"] = request.DubboFilterEnabled
	}

	if !dara.IsNil(request.Edition) {
		body["Edition"] = request.Edition
	}

	if !dara.IsNil(request.EnableACMG) {
		body["EnableACMG"] = request.EnableACMG
	}

	if !dara.IsNil(request.EnableAmbient) {
		body["EnableAmbient"] = request.EnableAmbient
	}

	if !dara.IsNil(request.EnableAudit) {
		body["EnableAudit"] = request.EnableAudit
	}

	if !dara.IsNil(request.EnableCRHistory) {
		body["EnableCRHistory"] = request.EnableCRHistory
	}

	if !dara.IsNil(request.EnableSDSServer) {
		body["EnableSDSServer"] = request.EnableSDSServer
	}

	if !dara.IsNil(request.ExcludeIPRanges) {
		body["ExcludeIPRanges"] = request.ExcludeIPRanges
	}

	if !dara.IsNil(request.ExcludeInboundPorts) {
		body["ExcludeInboundPorts"] = request.ExcludeInboundPorts
	}

	if !dara.IsNil(request.ExcludeOutboundPorts) {
		body["ExcludeOutboundPorts"] = request.ExcludeOutboundPorts
	}

	if !dara.IsNil(request.ExistingCaCert) {
		body["ExistingCaCert"] = request.ExistingCaCert
	}

	if !dara.IsNil(request.ExistingCaKey) {
		body["ExistingCaKey"] = request.ExistingCaKey
	}

	if !dara.IsNil(request.ExistingCaType) {
		body["ExistingCaType"] = request.ExistingCaType
	}

	if !dara.IsNil(request.ExistingRootCaCert) {
		body["ExistingRootCaCert"] = request.ExistingRootCaCert
	}

	if !dara.IsNil(request.ExistingRootCaKey) {
		body["ExistingRootCaKey"] = request.ExistingRootCaKey
	}

	if !dara.IsNil(request.FilterGatewayClusterConfig) {
		body["FilterGatewayClusterConfig"] = request.FilterGatewayClusterConfig
	}

	if !dara.IsNil(request.GatewayAPIEnabled) {
		body["GatewayAPIEnabled"] = request.GatewayAPIEnabled
	}

	if !dara.IsNil(request.GuestCluster) {
		body["GuestCluster"] = request.GuestCluster
	}

	if !dara.IsNil(request.IncludeIPRanges) {
		body["IncludeIPRanges"] = request.IncludeIPRanges
	}

	if !dara.IsNil(request.IstioVersion) {
		body["IstioVersion"] = request.IstioVersion
	}

	if !dara.IsNil(request.KialiEnabled) {
		body["KialiEnabled"] = request.KialiEnabled
	}

	if !dara.IsNil(request.LocalityLBConf) {
		body["LocalityLBConf"] = request.LocalityLBConf
	}

	if !dara.IsNil(request.LocalityLoadBalancing) {
		body["LocalityLoadBalancing"] = request.LocalityLoadBalancing
	}

	if !dara.IsNil(request.MSEEnabled) {
		body["MSEEnabled"] = request.MSEEnabled
	}

	if !dara.IsNil(request.MultiBufferEnabled) {
		body["MultiBufferEnabled"] = request.MultiBufferEnabled
	}

	if !dara.IsNil(request.MultiBufferPollDelay) {
		body["MultiBufferPollDelay"] = request.MultiBufferPollDelay
	}

	if !dara.IsNil(request.MysqlFilterEnabled) {
		body["MysqlFilterEnabled"] = request.MysqlFilterEnabled
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.OPALimitCPU) {
		body["OPALimitCPU"] = request.OPALimitCPU
	}

	if !dara.IsNil(request.OPALimitMemory) {
		body["OPALimitMemory"] = request.OPALimitMemory
	}

	if !dara.IsNil(request.OPALogLevel) {
		body["OPALogLevel"] = request.OPALogLevel
	}

	if !dara.IsNil(request.OPARequestCPU) {
		body["OPARequestCPU"] = request.OPARequestCPU
	}

	if !dara.IsNil(request.OPARequestMemory) {
		body["OPARequestMemory"] = request.OPARequestMemory
	}

	if !dara.IsNil(request.OpaEnabled) {
		body["OpaEnabled"] = request.OpaEnabled
	}

	if !dara.IsNil(request.OpenAgentPolicy) {
		body["OpenAgentPolicy"] = request.OpenAgentPolicy
	}

	if !dara.IsNil(request.Period) {
		body["Period"] = request.Period
	}

	if !dara.IsNil(request.PilotLoadBalancerSpec) {
		body["PilotLoadBalancerSpec"] = request.PilotLoadBalancerSpec
	}

	if !dara.IsNil(request.PlaygroundScene) {
		body["PlaygroundScene"] = request.PlaygroundScene
	}

	if !dara.IsNil(request.PrometheusUrl) {
		body["PrometheusUrl"] = request.PrometheusUrl
	}

	if !dara.IsNil(request.ProxyLimitCPU) {
		body["ProxyLimitCPU"] = request.ProxyLimitCPU
	}

	if !dara.IsNil(request.ProxyLimitMemory) {
		body["ProxyLimitMemory"] = request.ProxyLimitMemory
	}

	if !dara.IsNil(request.ProxyRequestCPU) {
		body["ProxyRequestCPU"] = request.ProxyRequestCPU
	}

	if !dara.IsNil(request.ProxyRequestMemory) {
		body["ProxyRequestMemory"] = request.ProxyRequestMemory
	}

	if !dara.IsNil(request.RedisFilterEnabled) {
		body["RedisFilterEnabled"] = request.RedisFilterEnabled
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Telemetry) {
		body["Telemetry"] = request.Telemetry
	}

	if !dara.IsNil(request.ThriftFilterEnabled) {
		body["ThriftFilterEnabled"] = request.ThriftFilterEnabled
	}

	if !dara.IsNil(request.TraceSampling) {
		body["TraceSampling"] = request.TraceSampling
	}

	if !dara.IsNil(request.Tracing) {
		body["Tracing"] = request.Tracing
	}

	if !dara.IsNil(request.UseExistingCA) {
		body["UseExistingCA"] = request.UseExistingCA
	}

	if !dara.IsNil(request.VSwitches) {
		body["VSwitches"] = request.VSwitches
	}

	if !dara.IsNil(request.VpcId) {
		body["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.WebAssemblyFilterEnabled) {
		body["WebAssemblyFilterEnabled"] = request.WebAssemblyFilterEnabled
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateServiceMesh"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateServiceMeshResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a lane.
//
// @param request - CreateSwimLaneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSwimLaneResponse
func (client *Client) CreateSwimLaneWithContext(ctx context.Context, request *CreateSwimLaneRequest, runtime *dara.RuntimeOptions) (_result *CreateSwimLaneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		body["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.LabelSelectorKey) {
		body["LabelSelectorKey"] = request.LabelSelectorKey
	}

	if !dara.IsNil(request.LabelSelectorValue) {
		body["LabelSelectorValue"] = request.LabelSelectorValue
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	if !dara.IsNil(request.ServicesList) {
		body["ServicesList"] = request.ServicesList
	}

	if !dara.IsNil(request.SwimLaneName) {
		body["SwimLaneName"] = request.SwimLaneName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSwimLane"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSwimLaneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a lane group.
//
// @param request - CreateSwimLaneGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSwimLaneGroupResponse
func (client *Client) CreateSwimLaneGroupWithContext(ctx context.Context, request *CreateSwimLaneGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateSwimLaneGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		body["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.IngressGatewayName) {
		body["IngressGatewayName"] = request.IngressGatewayName
	}

	if !dara.IsNil(request.IngressGatewayNamespace) {
		body["IngressGatewayNamespace"] = request.IngressGatewayNamespace
	}

	if !dara.IsNil(request.IngressType) {
		body["IngressType"] = request.IngressType
	}

	if !dara.IsNil(request.IsPermissive) {
		body["IsPermissive"] = request.IsPermissive
	}

	if !dara.IsNil(request.RouteHeader) {
		body["RouteHeader"] = request.RouteHeader
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	if !dara.IsNil(request.ServicesList) {
		body["ServicesList"] = request.ServicesList
	}

	if !dara.IsNil(request.TraceHeader) {
		body["TraceHeader"] = request.TraceHeader
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSwimLaneGroup"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSwimLaneGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建Waypoint
//
// @param request - CreateWaypointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWaypointResponse
func (client *Client) CreateWaypointWithContext(ctx context.Context, request *CreateWaypointRequest, runtime *dara.RuntimeOptions) (_result *CreateWaypointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.HPAEnabled) {
		body["HPAEnabled"] = request.HPAEnabled
	}

	if !dara.IsNil(request.HPAMaxReplicas) {
		body["HPAMaxReplicas"] = request.HPAMaxReplicas
	}

	if !dara.IsNil(request.HPAMinReplicas) {
		body["HPAMinReplicas"] = request.HPAMinReplicas
	}

	if !dara.IsNil(request.HPATargetCPU) {
		body["HPATargetCPU"] = request.HPATargetCPU
	}

	if !dara.IsNil(request.HPATargetMemory) {
		body["HPATargetMemory"] = request.HPATargetMemory
	}

	if !dara.IsNil(request.LimitCPU) {
		body["LimitCPU"] = request.LimitCPU
	}

	if !dara.IsNil(request.LimitMemory) {
		body["LimitMemory"] = request.LimitMemory
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PreferECI) {
		body["PreferECI"] = request.PreferECI
	}

	if !dara.IsNil(request.Replicas) {
		body["Replicas"] = request.Replicas
	}

	if !dara.IsNil(request.RequestCPU) {
		body["RequestCPU"] = request.RequestCPU
	}

	if !dara.IsNil(request.RequestMemory) {
		body["RequestMemory"] = request.RequestMemory
	}

	if !dara.IsNil(request.ServiceAccount) {
		body["ServiceAccount"] = request.ServiceAccount
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWaypoint"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWaypointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a routing rule for a Service Mesh (ASM) gateway.
//
// @param request - DeleteGatewayRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGatewayRouteResponse
func (client *Client) DeleteGatewayRouteWithContext(ctx context.Context, request *DeleteGatewayRouteRequest, runtime *dara.RuntimeOptions) (_result *DeleteGatewayRouteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IstioGatewayName) {
		body["IstioGatewayName"] = request.IstioGatewayName
	}

	if !dara.IsNil(request.RouteName) {
		body["RouteName"] = request.RouteName
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGatewayRoute"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGatewayRouteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a secret for a Service Mesh (ASM) gateway.
//
// @param request - DeleteGatewaySecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGatewaySecretResponse
func (client *Client) DeleteGatewaySecretWithContext(ctx context.Context, request *DeleteGatewaySecretRequest, runtime *dara.RuntimeOptions) (_result *DeleteGatewaySecretResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IstioGatewayName) {
		body["IstioGatewayName"] = request.IstioGatewayName
	}

	if !dara.IsNil(request.SecretName) {
		body["SecretName"] = request.SecretName
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGatewaySecret"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGatewaySecretResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes one or more domain names for a Service Mesh (ASM) gateway.
//
// @param request - DeleteIstioGatewayDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIstioGatewayDomainsResponse
func (client *Client) DeleteIstioGatewayDomainsWithContext(ctx context.Context, request *DeleteIstioGatewayDomainsRequest, runtime *dara.RuntimeOptions) (_result *DeleteIstioGatewayDomainsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Hosts) {
		body["Hosts"] = request.Hosts
	}

	if !dara.IsNil(request.IstioGatewayName) {
		body["IstioGatewayName"] = request.IstioGatewayName
	}

	if !dara.IsNil(request.Limit) {
		body["Limit"] = request.Limit
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PortName) {
		body["PortName"] = request.PortName
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteIstioGatewayDomains"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIstioGatewayDomainsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Service Mesh (ASM) instance.
//
// @param request - DeleteServiceMeshRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServiceMeshResponse
func (client *Client) DeleteServiceMeshWithContext(ctx context.Context, request *DeleteServiceMeshRequest, runtime *dara.RuntimeOptions) (_result *DeleteServiceMeshResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Force) {
		body["Force"] = request.Force
	}

	if !dara.IsNil(request.RetainResources) {
		body["RetainResources"] = request.RetainResources
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteServiceMesh"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteServiceMeshResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a lane.
//
// @param request - DeleteSwimLaneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSwimLaneResponse
func (client *Client) DeleteSwimLaneWithContext(ctx context.Context, request *DeleteSwimLaneRequest, runtime *dara.RuntimeOptions) (_result *DeleteSwimLaneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		body["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	if !dara.IsNil(request.SwimLaneName) {
		body["SwimLaneName"] = request.SwimLaneName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSwimLane"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSwimLaneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a lane group. If a lane group is deleted, the lanes in the group and the traffic routing rules attached to the lanes are deleted.
//
// @param request - DeleteSwimLaneGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSwimLaneGroupResponse
func (client *Client) DeleteSwimLaneGroupWithContext(ctx context.Context, request *DeleteSwimLaneGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteSwimLaneGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		body["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSwimLaneGroup"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSwimLaneGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除Waypoint资源
//
// @param request - DeleteWaypointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWaypointResponse
func (client *Client) DeleteWaypointWithContext(ctx context.Context, request *DeleteWaypointRequest, runtime *dara.RuntimeOptions) (_result *DeleteWaypointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWaypoint"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWaypointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about imported services on a Service Mesh (ASM) gateway.
//
// @param request - DescribeASMGatewayImportedServicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeASMGatewayImportedServicesResponse
func (client *Client) DescribeASMGatewayImportedServicesWithContext(ctx context.Context, request *DescribeASMGatewayImportedServicesRequest, runtime *dara.RuntimeOptions) (_result *DescribeASMGatewayImportedServicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ASMGatewayName) {
		body["ASMGatewayName"] = request.ASMGatewayName
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	if !dara.IsNil(request.ServiceNamespace) {
		body["ServiceNamespace"] = request.ServiceNamespace
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeASMGatewayImportedServices"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeASMGatewayImportedServicesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the versions of the Cloud Controller Manager (CCM) component.
//
// @param request - DescribeCCMVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCCMVersionResponse
func (client *Client) DescribeCCMVersionWithContext(ctx context.Context, request *DescribeCCMVersionRequest, runtime *dara.RuntimeOptions) (_result *DescribeCCMVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ServiceMeshId) {
		query["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCCMVersion"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCCMVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the network connectivity between clusters that are deployed across virtual private clouds (VPCs) in a Service Mesh (ASM) instance.
//
// @param request - DescribeCensRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCensResponse
func (client *Client) DescribeCensWithContext(ctx context.Context, request *DescribeCensRequest, runtime *dara.RuntimeOptions) (_result *DescribeCensResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ServiceMeshId) {
		query["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCens"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCensResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about Grafana dashboards of a cluster in a Service Mesh (ASM) instance.
//
// @param request - DescribeClusterGrafanaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterGrafanaResponse
func (client *Client) DescribeClusterGrafanaWithContext(ctx context.Context, request *DescribeClusterGrafanaRequest, runtime *dara.RuntimeOptions) (_result *DescribeClusterGrafanaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.K8sClusterId) {
		query["K8sClusterId"] = request.K8sClusterId
	}

	if !dara.IsNil(request.ReAddPrometheusIntegration) {
		query["ReAddPrometheusIntegration"] = request.ReAddPrometheusIntegration
	}

	if !dara.IsNil(request.ServiceMeshId) {
		query["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClusterGrafana"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClusterGrafanaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the public endpoint of the Prometheus service that is used to monitor a cluster in an Alibaba Cloud Service Mesh (ASM) instance.
//
// @param request - DescribeClusterPrometheusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterPrometheusResponse
func (client *Client) DescribeClusterPrometheusWithContext(ctx context.Context, request *DescribeClusterPrometheusRequest, runtime *dara.RuntimeOptions) (_result *DescribeClusterPrometheusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.K8sClusterId) {
		query["K8sClusterId"] = request.K8sClusterId
	}

	if !dara.IsNil(request.K8sClusterRegionId) {
		query["K8sClusterRegionId"] = request.K8sClusterRegionId
	}

	if !dara.IsNil(request.ServiceMeshId) {
		query["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClusterPrometheus"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClusterPrometheusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about clusters in a Service Mesh (ASM) instance.
//
// @param request - DescribeClustersInServiceMeshRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClustersInServiceMeshResponse
func (client *Client) DescribeClustersInServiceMeshWithContext(ctx context.Context, request *DescribeClustersInServiceMeshRequest, runtime *dara.RuntimeOptions) (_result *DescribeClustersInServiceMeshResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ServiceMeshId) {
		query["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClustersInServiceMesh"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClustersInServiceMeshResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the common YAML templates of Istio resources used by Service Mesh (ASM) instances.
//
// @param request - DescribeCrTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCrTemplatesResponse
func (client *Client) DescribeCrTemplatesWithContext(ctx context.Context, request *DescribeCrTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCrTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IstioVersion) {
		body["IstioVersion"] = request.IstioVersion
	}

	if !dara.IsNil(request.Kind) {
		body["Kind"] = request.Kind
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCrTemplates"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCrTemplatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DescribeEipResources
//
// @param request - DescribeEipResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEipResourcesResponse
func (client *Client) DescribeEipResourcesWithContext(ctx context.Context, request *DescribeEipResourcesRequest, runtime *dara.RuntimeOptions) (_result *DescribeEipResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PageNum) {
		body["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEipResources"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEipResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the detailed information about a secret of a Service Mesh (ASM) gateway.
//
// @param request - DescribeGatewaySecretDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGatewaySecretDetailsResponse
func (client *Client) DescribeGatewaySecretDetailsWithContext(ctx context.Context, request *DescribeGatewaySecretDetailsRequest, runtime *dara.RuntimeOptions) (_result *DescribeGatewaySecretDetailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IstioGatewayName) {
		body["IstioGatewayName"] = request.IstioGatewayName
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGatewaySecretDetails"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGatewaySecretDetailsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the access log dashboards of a cluster on the data plane.
//
// @param request - DescribeGuestClusterAccessLogDashboardsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGuestClusterAccessLogDashboardsResponse
func (client *Client) DescribeGuestClusterAccessLogDashboardsWithContext(ctx context.Context, request *DescribeGuestClusterAccessLogDashboardsRequest, runtime *dara.RuntimeOptions) (_result *DescribeGuestClusterAccessLogDashboardsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.K8sClusterId) {
		body["K8sClusterId"] = request.K8sClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGuestClusterAccessLogDashboards"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGuestClusterAccessLogDashboardsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of the namespaces of a Kubernetes cluster that is added to a Service Mesh (ASM) instance.
//
// @param request - DescribeGuestClusterNamespacesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGuestClusterNamespacesResponse
func (client *Client) DescribeGuestClusterNamespacesWithContext(ctx context.Context, request *DescribeGuestClusterNamespacesRequest, runtime *dara.RuntimeOptions) (_result *DescribeGuestClusterNamespacesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.GuestClusterID) {
		body["GuestClusterID"] = request.GuestClusterID
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	if !dara.IsNil(request.ShowNsLabels) {
		body["ShowNsLabels"] = request.ShowNsLabels
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGuestClusterNamespaces"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGuestClusterNamespacesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of the pods in a specified namespace of a Kubernetes cluster that is added to a Service Mesh (ASM) instance.
//
// @param request - DescribeGuestClusterPodsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGuestClusterPodsResponse
func (client *Client) DescribeGuestClusterPodsWithContext(ctx context.Context, request *DescribeGuestClusterPodsRequest, runtime *dara.RuntimeOptions) (_result *DescribeGuestClusterPodsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.GuestClusterID) {
		body["GuestClusterID"] = request.GuestClusterID
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGuestClusterPods"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGuestClusterPodsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of the imported services on a Service Mesh (ASM) gateway.
//
// @param request - DescribeImportedServicesDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeImportedServicesDetailResponse
func (client *Client) DescribeImportedServicesDetailWithContext(ctx context.Context, request *DescribeImportedServicesDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeImportedServicesDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ASMGatewayName) {
		body["ASMGatewayName"] = request.ASMGatewayName
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	if !dara.IsNil(request.ServiceNamespace) {
		body["ServiceNamespace"] = request.ServiceNamespace
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeImportedServicesDetail"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeImportedServicesDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of the domain names that are exposed by a Service Mesh (ASM) gateway.
//
// @param request - DescribeIstioGatewayDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIstioGatewayDomainsResponse
func (client *Client) DescribeIstioGatewayDomainsWithContext(ctx context.Context, request *DescribeIstioGatewayDomainsRequest, runtime *dara.RuntimeOptions) (_result *DescribeIstioGatewayDomainsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IstioGatewayName) {
		body["IstioGatewayName"] = request.IstioGatewayName
	}

	if !dara.IsNil(request.Limit) {
		body["Limit"] = request.Limit
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeIstioGatewayDomains"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIstioGatewayDomainsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the detailed information about a routing rule of an Alibaba Cloud Service Mesh (ASM) gateway.
//
// @param request - DescribeIstioGatewayRouteDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIstioGatewayRouteDetailResponse
func (client *Client) DescribeIstioGatewayRouteDetailWithContext(ctx context.Context, request *DescribeIstioGatewayRouteDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeIstioGatewayRouteDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IstioGatewayName) {
		body["IstioGatewayName"] = request.IstioGatewayName
	}

	if !dara.IsNil(request.RouteName) {
		body["RouteName"] = request.RouteName
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeIstioGatewayRouteDetail"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIstioGatewayRouteDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the routing rules for a Service Mesh (ASM) gateway.
//
// @param request - DescribeIstioGatewayRoutesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIstioGatewayRoutesResponse
func (client *Client) DescribeIstioGatewayRoutesWithContext(ctx context.Context, request *DescribeIstioGatewayRoutesRequest, runtime *dara.RuntimeOptions) (_result *DescribeIstioGatewayRoutesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IstioGatewayName) {
		body["IstioGatewayName"] = request.IstioGatewayName
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeIstioGatewayRoutes"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIstioGatewayRoutesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取ASM网格网络分区设置
//
// @param request - DescribeMeshMultiClusterNetworkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMeshMultiClusterNetworkResponse
func (client *Client) DescribeMeshMultiClusterNetworkWithContext(ctx context.Context, request *DescribeMeshMultiClusterNetworkRequest, runtime *dara.RuntimeOptions) (_result *DescribeMeshMultiClusterNetworkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMeshMultiClusterNetwork"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMeshMultiClusterNetworkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of sidecar proxies at the namespace level.
//
// @param request - DescribeNamespaceScopeSidecarConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNamespaceScopeSidecarConfigResponse
func (client *Client) DescribeNamespaceScopeSidecarConfigWithContext(ctx context.Context, request *DescribeNamespaceScopeSidecarConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeNamespaceScopeSidecarConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNamespaceScopeSidecarConfig"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNamespaceScopeSidecarConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the instance types of nodes on the data plane and whether the instance types support Multi-Buffer acceleration.
//
// @param request - DescribeNodesInstanceTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNodesInstanceTypeResponse
func (client *Client) DescribeNodesInstanceTypeWithContext(ctx context.Context, request *DescribeNodesInstanceTypeRequest, runtime *dara.RuntimeOptions) (_result *DescribeNodesInstanceTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNodesInstanceType"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNodesInstanceTypeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Server Load Balancer (SLB) instances that can be reused.
//
// @param request - DescribeReusableSlbRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeReusableSlbResponse
func (client *Client) DescribeReusableSlbWithContext(ctx context.Context, request *DescribeReusableSlbRequest, runtime *dara.RuntimeOptions) (_result *DescribeReusableSlbResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.K8sClusterId) {
		body["K8sClusterId"] = request.K8sClusterId
	}

	if !dara.IsNil(request.LbType) {
		body["LbType"] = request.LbType
	}

	if !dara.IsNil(request.NetworkType) {
		body["NetworkType"] = request.NetworkType
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeReusableSlb"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeReusableSlbResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the check results of a Service Mesh (ASM) instance.
//
// @param request - DescribeServiceMeshAdditionalStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceMeshAdditionalStatusResponse
func (client *Client) DescribeServiceMeshAdditionalStatusWithContext(ctx context.Context, request *DescribeServiceMeshAdditionalStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeServiceMeshAdditionalStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CheckMode) {
		body["CheckMode"] = request.CheckMode
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeServiceMeshAdditionalStatus"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeServiceMeshAdditionalStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the clusters that can be added to a Service Mesh (ASM) instance.
//
// @param request - DescribeServiceMeshClustersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceMeshClustersResponse
func (client *Client) DescribeServiceMeshClustersWithContext(ctx context.Context, request *DescribeServiceMeshClustersRequest, runtime *dara.RuntimeOptions) (_result *DescribeServiceMeshClustersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ServiceMeshId) {
		query["ServiceMeshId"] = request.ServiceMeshId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Limit) {
		body["Limit"] = request.Limit
	}

	if !dara.IsNil(request.Offset) {
		body["Offset"] = request.Offset
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeServiceMeshClusters"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeServiceMeshClustersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a Service Mesh (ASM) instance.
//
// @param request - DescribeServiceMeshDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceMeshDetailResponse
func (client *Client) DescribeServiceMeshDetailWithContext(ctx context.Context, request *DescribeServiceMeshDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeServiceMeshDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeServiceMeshDetail"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeServiceMeshDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the content of the kubeconfig file of a Service Mesh (ASM) instance.
//
// @param request - DescribeServiceMeshKubeconfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceMeshKubeconfigResponse
func (client *Client) DescribeServiceMeshKubeconfigWithContext(ctx context.Context, request *DescribeServiceMeshKubeconfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeServiceMeshKubeconfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PrivateIpAddress) {
		query["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !dara.IsNil(request.ServiceMeshId) {
		query["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeServiceMeshKubeconfig"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeServiceMeshKubeconfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the logs of a Service Mesh (ASM) instance.
//
// @param request - DescribeServiceMeshLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceMeshLogsResponse
func (client *Client) DescribeServiceMeshLogsWithContext(ctx context.Context, request *DescribeServiceMeshLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeServiceMeshLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeServiceMeshLogs"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeServiceMeshLogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of proxies on the data plane of a Service Mesh (ASM) instance.
//
// @param request - DescribeServiceMeshProxyStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceMeshProxyStatusResponse
func (client *Client) DescribeServiceMeshProxyStatusWithContext(ctx context.Context, request *DescribeServiceMeshProxyStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeServiceMeshProxyStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeServiceMeshProxyStatus"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeServiceMeshProxyStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the upgrade details of a Service Mesh (ASM) instance and its ingress gateways.
//
// @param request - DescribeServiceMeshUpgradeStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceMeshUpgradeStatusResponse
func (client *Client) DescribeServiceMeshUpgradeStatusWithContext(ctx context.Context, request *DescribeServiceMeshUpgradeStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeServiceMeshUpgradeStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ServiceMeshId) {
		query["ServiceMeshId"] = request.ServiceMeshId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AllIstioGatewayFullNames) {
		body["AllIstioGatewayFullNames"] = request.AllIstioGatewayFullNames
	}

	if !dara.IsNil(request.GuestClusterIds) {
		body["GuestClusterIds"] = request.GuestClusterIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeServiceMeshUpgradeStatus"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeServiceMeshUpgradeStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeServiceMeshVMs is deprecated
//
// Summary:
//
// Queries the Elastic Compute Service (ECS) instances that reside in the same virtual private cloud (VPC) as a Service Mesh (ASM) instance.
//
// @param request - DescribeServiceMeshVMsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceMeshVMsResponse
func (client *Client) DescribeServiceMeshVMsWithContext(ctx context.Context, request *DescribeServiceMeshVMsRequest, runtime *dara.RuntimeOptions) (_result *DescribeServiceMeshVMsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ServiceMeshId) {
		query["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeServiceMeshVMs"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeServiceMeshVMsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries ASM instances.
//
// @param request - DescribeServiceMeshesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceMeshesResponse
func (client *Client) DescribeServiceMeshesWithContext(ctx context.Context, request *DescribeServiceMeshesRequest, runtime *dara.RuntimeOptions) (_result *DescribeServiceMeshesResponse, _err error) {
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
		Action:      dara.String("DescribeServiceMeshes"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeServiceMeshesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the update status of a Service Mesh (ASM) instance.
//
// @param request - DescribeUpgradeVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUpgradeVersionResponse
func (client *Client) DescribeUpgradeVersionWithContext(ctx context.Context, request *DescribeUpgradeVersionRequest, runtime *dara.RuntimeOptions) (_result *DescribeUpgradeVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ServiceMeshId) {
		query["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUpgradeVersion"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUpgradeVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains role-based access control (RBAC) permissions.
//
// @param request - DescribeUserPermissionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserPermissionsResponse
func (client *Client) DescribeUserPermissionsWithContext(ctx context.Context, request *DescribeUserPermissionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserPermissionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SubAccountUserId) {
		body["SubAccountUserId"] = request.SubAccountUserId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserPermissions"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserPermissionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the IDs of all RAM users or RAM roles to which a Role-based Access Control (RBAC) role is assigned.
//
// @param request - DescribeUsersWithPermissionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUsersWithPermissionsResponse
func (client *Client) DescribeUsersWithPermissionsWithContext(ctx context.Context, request *DescribeUsersWithPermissionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeUsersWithPermissionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.UserType) {
		body["UserType"] = request.UserType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUsersWithPermissions"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUsersWithPermissionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeVMsInServiceMesh is deprecated
//
// Summary:
//
// Queries the virtual machines (VMs) that are added to a Service Mesh (ASM) instance.
//
// @param request - DescribeVMsInServiceMeshRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVMsInServiceMeshResponse
func (client *Client) DescribeVMsInServiceMeshWithContext(ctx context.Context, request *DescribeVMsInServiceMeshRequest, runtime *dara.RuntimeOptions) (_result *DescribeVMsInServiceMeshResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ServiceMeshId) {
		query["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVMsInServiceMesh"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVMsInServiceMeshResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of vSwitches that are deployed in a specified virtual private cloud (VPC) in a region.
//
// @param request - DescribeVSwitchesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVSwitchesResponse
func (client *Client) DescribeVSwitchesWithContext(ctx context.Context, request *DescribeVSwitchesRequest, runtime *dara.RuntimeOptions) (_result *DescribeVSwitchesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VpcId) {
		body["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVSwitches"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVSwitchesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the virtual private clouds (VPCs) that are available in a specified region.
//
// @param request - DescribeVpcsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcsResponse
func (client *Client) DescribeVpcsWithContext(ctx context.Context, request *DescribeVpcsRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVpcs"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVpcsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a certificate issued by a certificate authority (CA).
//
// @param request - GetCaCertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCaCertResponse
func (client *Client) GetCaCertWithContext(ctx context.Context, request *GetCaCertRequest, runtime *dara.RuntimeOptions) (_result *GetCaCertResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCaCert"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCaCertResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of workloads specified by a label selector.
//
// @param tmpReq - GetDeploymentBySelectorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeploymentBySelectorResponse
func (client *Client) GetDeploymentBySelectorWithContext(ctx context.Context, tmpReq *GetDeploymentBySelectorRequest, runtime *dara.RuntimeOptions) (_result *GetDeploymentBySelectorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetDeploymentBySelectorShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LabelSelector) {
		request.LabelSelectorShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LabelSelector, dara.String("LabelSelector"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.GuestCluster) {
		body["GuestCluster"] = request.GuestCluster
	}

	if !dara.IsNil(request.LabelSelectorShrink) {
		body["LabelSelector"] = request.LabelSelectorShrink
	}

	if !dara.IsNil(request.Limit) {
		body["Limit"] = request.Limit
	}

	if !dara.IsNil(request.Mark) {
		body["Mark"] = request.Mark
	}

	if !dara.IsNil(request.NameSpace) {
		body["NameSpace"] = request.NameSpace
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeploymentBySelector"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeploymentBySelectorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Grafana dashboard URL from Application Real-Time Monitoring Service (ARMS).
//
// @param request - GetGrafanaDashboardUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGrafanaDashboardUrlResponse
func (client *Client) GetGrafanaDashboardUrlWithContext(ctx context.Context, request *GetGrafanaDashboardUrlRequest, runtime *dara.RuntimeOptions) (_result *GetGrafanaDashboardUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.K8sClusterId) {
		body["K8sClusterId"] = request.K8sClusterId
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGrafanaDashboardUrl"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGrafanaDashboardUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 描述ServiceEndpoints信息
//
// @param request - GetRegisteredServiceEndpointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRegisteredServiceEndpointsResponse
func (client *Client) GetRegisteredServiceEndpointsWithContext(ctx context.Context, request *GetRegisteredServiceEndpointsRequest, runtime *dara.RuntimeOptions) (_result *GetRegisteredServiceEndpointsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterIds) {
		body["ClusterIds"] = request.ClusterIds
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	if !dara.IsNil(request.ServiceType) {
		body["ServiceType"] = request.ServiceType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRegisteredServiceEndpoints"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRegisteredServiceEndpointsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetRegisteredServiceNamespacesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRegisteredServiceNamespacesResponse
func (client *Client) GetRegisteredServiceNamespacesWithContext(ctx context.Context, request *GetRegisteredServiceNamespacesRequest, runtime *dara.RuntimeOptions) (_result *GetRegisteredServiceNamespacesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRegisteredServiceNamespaces"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRegisteredServiceNamespacesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries detailed information about a lane.
//
// @param request - GetSwimLaneDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSwimLaneDetailResponse
func (client *Client) GetSwimLaneDetailWithContext(ctx context.Context, request *GetSwimLaneDetailRequest, runtime *dara.RuntimeOptions) (_result *GetSwimLaneDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		body["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	if !dara.IsNil(request.SwimLaneName) {
		body["SwimLaneName"] = request.SwimLaneName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSwimLaneDetail"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSwimLaneDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of all lane groups in an Alibaba Cloud Service Mesh (ASM) instance.
//
// @param request - GetSwimLaneGroupListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSwimLaneGroupListResponse
func (client *Client) GetSwimLaneGroupListWithContext(ctx context.Context, request *GetSwimLaneGroupListRequest, runtime *dara.RuntimeOptions) (_result *GetSwimLaneGroupListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSwimLaneGroupList"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSwimLaneGroupListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of all the lanes in a lane group.
//
// @param request - GetSwimLaneListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSwimLaneListResponse
func (client *Client) GetSwimLaneListWithContext(ctx context.Context, request *GetSwimLaneListRequest, runtime *dara.RuntimeOptions) (_result *GetSwimLaneListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		body["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSwimLaneList"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSwimLaneListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetVmAppMeshInfo is deprecated
//
// Summary:
//
// Queries the information about VMs that are added to a Service Mesh (ASM) instance.
//
// @param request - GetVmAppMeshInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVmAppMeshInfoResponse
func (client *Client) GetVmAppMeshInfoWithContext(ctx context.Context, request *GetVmAppMeshInfoRequest, runtime *dara.RuntimeOptions) (_result *GetVmAppMeshInfoResponse, _err error) {
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
		Action:      dara.String("GetVmAppMeshInfo"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVmAppMeshInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetVmMeta is deprecated
//
// Summary:
//
// Queries the metadata that is required to add a non-containerized application to a Service Mesh (ASM) instance.
//
// @param request - GetVmMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVmMetaResponse
func (client *Client) GetVmMetaWithContext(ctx context.Context, request *GetVmMetaRequest, runtime *dara.RuntimeOptions) (_result *GetVmMetaResponse, _err error) {
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
		Action:      dara.String("GetVmMeta"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVmMetaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants permissions to a Resource Access Management (RAM) user.
//
// @param tmpReq - GrantUserPermissionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantUserPermissionsResponse
func (client *Client) GrantUserPermissionsWithContext(ctx context.Context, tmpReq *GrantUserPermissionsRequest, runtime *dara.RuntimeOptions) (_result *GrantUserPermissionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GrantUserPermissionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SubAccountUserIds) {
		request.SubAccountUserIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SubAccountUserIds, dara.String("SubAccountUserIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Permissions) {
		body["Permissions"] = request.Permissions
	}

	if !dara.IsNil(request.SubAccountUserId) {
		body["SubAccountUserId"] = request.SubAccountUserId
	}

	if !dara.IsNil(request.SubAccountUserIdsShrink) {
		body["SubAccountUserIds"] = request.SubAccountUserIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GrantUserPermissions"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GrantUserPermissionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举所有服务账号
//
// @param request - ListServiceAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceAccountsResponse
func (client *Client) ListServiceAccountsWithContext(ctx context.Context, request *ListServiceAccountsRequest, runtime *dara.RuntimeOptions) (_result *ListServiceAccountsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServiceAccounts"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServiceAccountsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries user tags on a Service Mesh (ASM) instance.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithContext(ctx context.Context, request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
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
		Version:     dara.String("2020-01-11"),
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
// Queries the configurations of all waypoint proxies in a namespace of a cluster on the data plane.
//
// @param request - ListWaypointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWaypointsResponse
func (client *Client) ListWaypointsWithContext(ctx context.Context, request *ListWaypointsRequest, runtime *dara.RuntimeOptions) (_result *ListWaypointsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Continue) {
		body["Continue"] = request.Continue
	}

	if !dara.IsNil(request.Limit) {
		body["Limit"] = request.Limit
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWaypoints"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWaypointsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ModifyApiServerEipResource
//
// @param request - ModifyApiServerEipResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApiServerEipResourceResponse
func (client *Client) ModifyApiServerEipResourceWithContext(ctx context.Context, request *ModifyApiServerEipResourceRequest, runtime *dara.RuntimeOptions) (_result *ModifyApiServerEipResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiServerEipId) {
		body["ApiServerEipId"] = request.ApiServerEipId
	}

	if !dara.IsNil(request.Operation) {
		body["Operation"] = request.Operation
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyApiServerEipResource"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyApiServerEipResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ModifyPilotEipResource
//
// @param request - ModifyPilotEipResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPilotEipResourceResponse
func (client *Client) ModifyPilotEipResourceWithContext(ctx context.Context, request *ModifyPilotEipResourceRequest, runtime *dara.RuntimeOptions) (_result *ModifyPilotEipResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EipId) {
		body["EipId"] = request.EipId
	}

	if !dara.IsNil(request.IsCanary) {
		body["IsCanary"] = request.IsCanary
	}

	if !dara.IsNil(request.Operation) {
		body["Operation"] = request.Operation
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyPilotEipResource"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyPilotEipResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name of a Service Mesh (ASM) instance.
//
// @param request - ModifyServiceMeshNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyServiceMeshNameResponse
func (client *Client) ModifyServiceMeshNameWithContext(ctx context.Context, request *ModifyServiceMeshNameRequest, runtime *dara.RuntimeOptions) (_result *ModifyServiceMeshNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyServiceMeshName"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyServiceMeshNameResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Recreates a project that is used to store audit logs. After mesh audit is enabled, if you delete the log project that stores audit logs by mistake, you can recreate a project for storing audit logs.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods of Simple Log Service. For more information, visit the [pricing page](https://www.alibabacloud.com/zh/pricing-calculator?_p_lc=1\\&spm=a2796.7960336.3034855210.1.44e6b91aaSp2M7#/commodity/vm_intl).
//
// @param request - ReActivateAuditRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReActivateAuditResponse
func (client *Client) ReActivateAuditWithContext(ctx context.Context, request *ReActivateAuditRequest, runtime *dara.RuntimeOptions) (_result *ReActivateAuditResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EnableAudit) {
		body["EnableAudit"] = request.EnableAudit
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReActivateAudit"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReActivateAuditResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a cluster from a Service Mesh (ASM) instance.
//
// @param request - RemoveClusterFromServiceMeshRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveClusterFromServiceMeshResponse
func (client *Client) RemoveClusterFromServiceMeshWithContext(ctx context.Context, request *RemoveClusterFromServiceMeshRequest, runtime *dara.RuntimeOptions) (_result *RemoveClusterFromServiceMeshResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ReserveNamespace) {
		body["ReserveNamespace"] = request.ReserveNamespace
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveClusterFromServiceMesh"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveClusterFromServiceMeshResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI RemoveVMFromServiceMesh is deprecated
//
// Summary:
//
// Removes a virtual machine (VM) from a Service Mesh (ASM) instance.
//
// @param request - RemoveVMFromServiceMeshRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveVMFromServiceMeshResponse
func (client *Client) RemoveVMFromServiceMeshWithContext(ctx context.Context, request *RemoveVMFromServiceMeshRequest, runtime *dara.RuntimeOptions) (_result *RemoveVMFromServiceMeshResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EcsId) {
		query["EcsId"] = request.EcsId
	}

	if !dara.IsNil(request.ServiceMeshId) {
		query["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveVMFromServiceMesh"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveVMFromServiceMeshResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes the kubeconfig file of a Service Mesh (ASM) instance and generates a new kubeconfig file.
//
// @param request - RevokeKubeconfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeKubeconfigResponse
func (client *Client) RevokeKubeconfigWithContext(ctx context.Context, request *RevokeKubeconfigRequest, runtime *dara.RuntimeOptions) (_result *RevokeKubeconfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PrivateIpAddress) {
		body["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeKubeconfig"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeKubeconfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds or modifies user tags on a resource.
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithContext(ctx context.Context, request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
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
		Version:     dara.String("2020-01-11"),
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

// Summary:
//
// Deletes user tags on a Service Mesh (ASM) instance.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithContext(ctx context.Context, request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
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
		Version:     dara.String("2020-01-11"),
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
// Updates a Service Mesh (ASM) gateway.
//
// @param request - UpdateASMGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateASMGatewayResponse
func (client *Client) UpdateASMGatewayWithContext(ctx context.Context, request *UpdateASMGatewayRequest, runtime *dara.RuntimeOptions) (_result *UpdateASMGatewayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Body) {
		body["Body"] = request.Body
	}

	if !dara.IsNil(request.IstioGatewayName) {
		body["IstioGatewayName"] = request.IstioGatewayName
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateASMGateway"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateASMGatewayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates imported services on a Service Mesh (ASM) gateway to import or delete upstream services associated with the gateway.
//
// @param request - UpdateASMGatewayImportedServicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateASMGatewayImportedServicesResponse
func (client *Client) UpdateASMGatewayImportedServicesWithContext(ctx context.Context, request *UpdateASMGatewayImportedServicesRequest, runtime *dara.RuntimeOptions) (_result *UpdateASMGatewayImportedServicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ASMGatewayName) {
		body["ASMGatewayName"] = request.ASMGatewayName
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	if !dara.IsNil(request.ServiceNames) {
		body["ServiceNames"] = request.ServiceNames
	}

	if !dara.IsNil(request.ServiceNamespace) {
		body["ServiceNamespace"] = request.ServiceNamespace
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateASMGatewayImportedServices"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateASMGatewayImportedServicesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Synchronizes namespaces of a Kubernetes cluster that is added to a Service Mesh (ASM) instance.
//
// @param request - UpdateASMNamespaceFromGuestClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateASMNamespaceFromGuestClusterResponse
func (client *Client) UpdateASMNamespaceFromGuestClusterWithContext(ctx context.Context, request *UpdateASMNamespaceFromGuestClusterRequest, runtime *dara.RuntimeOptions) (_result *UpdateASMNamespaceFromGuestClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.K8sClusterId) {
		body["K8sClusterId"] = request.K8sClusterId
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateASMNamespaceFromGuestCluster"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateASMNamespaceFromGuestClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration for collecting control plane logs.
//
// @param request - UpdateControlPlaneLogConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateControlPlaneLogConfigResponse
func (client *Client) UpdateControlPlaneLogConfigWithContext(ctx context.Context, request *UpdateControlPlaneLogConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateControlPlaneLogConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Enabled) {
		body["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.LogTTLInDay) {
		body["LogTTLInDay"] = request.LogTTLInDay
	}

	if !dara.IsNil(request.Project) {
		body["Project"] = request.Project
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateControlPlaneLogConfig"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateControlPlaneLogConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新Guest Cluster配置
//
// @param request - UpdateGuestClusterConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGuestClusterConfigResponse
func (client *Client) UpdateGuestClusterConfigWithContext(ctx context.Context, request *UpdateGuestClusterConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateGuestClusterConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.GuestClusterId) {
		body["GuestClusterId"] = request.GuestClusterId
	}

	if !dara.IsNil(request.SMCEnabled) {
		body["SMCEnabled"] = request.SMCEnabled
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGuestClusterConfig"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGuestClusterConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a routing rule for a Service Mesh (ASM) gateway.
//
// @param tmpReq - UpdateIstioGatewayRoutesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIstioGatewayRoutesResponse
func (client *Client) UpdateIstioGatewayRoutesWithContext(ctx context.Context, tmpReq *UpdateIstioGatewayRoutesRequest, runtime *dara.RuntimeOptions) (_result *UpdateIstioGatewayRoutesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateIstioGatewayRoutesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.GatewayRoute) {
		request.GatewayRouteShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GatewayRoute, dara.String("GatewayRoute"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.GatewayRouteShrink) {
		body["GatewayRoute"] = request.GatewayRouteShrink
	}

	if !dara.IsNil(request.IstioGatewayName) {
		body["IstioGatewayName"] = request.IstioGatewayName
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateIstioGatewayRoutes"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateIstioGatewayRoutesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateIstioInjectionConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIstioInjectionConfigResponse
func (client *Client) UpdateIstioInjectionConfigWithContext(ctx context.Context, request *UpdateIstioInjectionConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateIstioInjectionConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DataPlaneMode) {
		body["DataPlaneMode"] = request.DataPlaneMode
	}

	if !dara.IsNil(request.EnableIstioInjection) {
		body["EnableIstioInjection"] = request.EnableIstioInjection
	}

	if !dara.IsNil(request.EnableSidecarSetInjection) {
		body["EnableSidecarSetInjection"] = request.EnableSidecarSetInjection
	}

	if !dara.IsNil(request.IstioRev) {
		body["IstioRev"] = request.IstioRev
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateIstioInjectionConfig"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateIstioInjectionConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a routing rule for a Service Mesh (ASM) gateway.
//
// @param request - UpdateIstioRouteAdditionalStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIstioRouteAdditionalStatusResponse
func (client *Client) UpdateIstioRouteAdditionalStatusWithContext(ctx context.Context, request *UpdateIstioRouteAdditionalStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateIstioRouteAdditionalStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RouteName) {
		query["RouteName"] = request.RouteName
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.IstioGatewayName) {
		body["IstioGatewayName"] = request.IstioGatewayName
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateIstioRouteAdditionalStatus"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateIstioRouteAdditionalStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the settings of whether to enable the Kubernetes API on the data plane to access Istio resources.
//
// @param request - UpdateMeshCRAggregationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMeshCRAggregationResponse
func (client *Client) UpdateMeshCRAggregationWithContext(ctx context.Context, request *UpdateMeshCRAggregationRequest, runtime *dara.RuntimeOptions) (_result *UpdateMeshCRAggregationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CPULimit) {
		body["CPULimit"] = request.CPULimit
	}

	if !dara.IsNil(request.CPURequirement) {
		body["CPURequirement"] = request.CPURequirement
	}

	if !dara.IsNil(request.Enabled) {
		body["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.MemoryLimit) {
		body["MemoryLimit"] = request.MemoryLimit
	}

	if !dara.IsNil(request.MemoryRequirement) {
		body["MemoryRequirement"] = request.MemoryRequirement
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	if !dara.IsNil(request.UsePublicApiServer) {
		body["UsePublicApiServer"] = request.UsePublicApiServer
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMeshCRAggregation"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMeshCRAggregationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration of a Service Mesh (ASM) instance.
//
// @param request - UpdateMeshFeatureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMeshFeatureResponse
func (client *Client) UpdateMeshFeatureWithContext(ctx context.Context, request *UpdateMeshFeatureRequest, runtime *dara.RuntimeOptions) (_result *UpdateMeshFeatureResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessLogGatewayEnabled) {
		query["AccessLogGatewayEnabled"] = request.AccessLogGatewayEnabled
	}

	if !dara.IsNil(request.AccessLogSidecarEnabled) {
		query["AccessLogSidecarEnabled"] = request.AccessLogSidecarEnabled
	}

	if !dara.IsNil(request.LabelsForOffloadedWorkloads) {
		query["LabelsForOffloadedWorkloads"] = request.LabelsForOffloadedWorkloads
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccessLogEnabled) {
		body["AccessLogEnabled"] = request.AccessLogEnabled
	}

	if !dara.IsNil(request.AccessLogFile) {
		body["AccessLogFile"] = request.AccessLogFile
	}

	if !dara.IsNil(request.AccessLogFormat) {
		body["AccessLogFormat"] = request.AccessLogFormat
	}

	if !dara.IsNil(request.AccessLogGatewayLifecycle) {
		body["AccessLogGatewayLifecycle"] = request.AccessLogGatewayLifecycle
	}

	if !dara.IsNil(request.AccessLogProject) {
		body["AccessLogProject"] = request.AccessLogProject
	}

	if !dara.IsNil(request.AccessLogServiceEnabled) {
		body["AccessLogServiceEnabled"] = request.AccessLogServiceEnabled
	}

	if !dara.IsNil(request.AccessLogServiceHost) {
		body["AccessLogServiceHost"] = request.AccessLogServiceHost
	}

	if !dara.IsNil(request.AccessLogServicePort) {
		body["AccessLogServicePort"] = request.AccessLogServicePort
	}

	if !dara.IsNil(request.AccessLogSidecarLifecycle) {
		body["AccessLogSidecarLifecycle"] = request.AccessLogSidecarLifecycle
	}

	if !dara.IsNil(request.AuditProject) {
		body["AuditProject"] = request.AuditProject
	}

	if !dara.IsNil(request.AutoInjectionPolicyEnabled) {
		body["AutoInjectionPolicyEnabled"] = request.AutoInjectionPolicyEnabled
	}

	if !dara.IsNil(request.CRAggregationEnabled) {
		body["CRAggregationEnabled"] = request.CRAggregationEnabled
	}

	if !dara.IsNil(request.CertChain) {
		body["CertChain"] = request.CertChain
	}

	if !dara.IsNil(request.ClusterSpec) {
		body["ClusterSpec"] = request.ClusterSpec
	}

	if !dara.IsNil(request.CniEnabled) {
		body["CniEnabled"] = request.CniEnabled
	}

	if !dara.IsNil(request.CniExcludeNamespaces) {
		body["CniExcludeNamespaces"] = request.CniExcludeNamespaces
	}

	if !dara.IsNil(request.Concurrency) {
		body["Concurrency"] = request.Concurrency
	}

	if !dara.IsNil(request.ConfigSourceEnabled) {
		body["ConfigSourceEnabled"] = request.ConfigSourceEnabled
	}

	if !dara.IsNil(request.ConfigSourceNacosID) {
		body["ConfigSourceNacosID"] = request.ConfigSourceNacosID
	}

	if !dara.IsNil(request.CustomizedPrometheus) {
		body["CustomizedPrometheus"] = request.CustomizedPrometheus
	}

	if !dara.IsNil(request.CustomizedZipkin) {
		body["CustomizedZipkin"] = request.CustomizedZipkin
	}

	if !dara.IsNil(request.DNSProxyingEnabled) {
		body["DNSProxyingEnabled"] = request.DNSProxyingEnabled
	}

	if !dara.IsNil(request.DefaultComponentsScheduleConfig) {
		body["DefaultComponentsScheduleConfig"] = request.DefaultComponentsScheduleConfig
	}

	if !dara.IsNil(request.DiscoverySelectors) {
		body["DiscoverySelectors"] = request.DiscoverySelectors
	}

	if !dara.IsNil(request.DubboFilterEnabled) {
		body["DubboFilterEnabled"] = request.DubboFilterEnabled
	}

	if !dara.IsNil(request.EnableAudit) {
		body["EnableAudit"] = request.EnableAudit
	}

	if !dara.IsNil(request.EnableAutoDiagnosis) {
		body["EnableAutoDiagnosis"] = request.EnableAutoDiagnosis
	}

	if !dara.IsNil(request.EnableBootstrapXdsAgent) {
		body["EnableBootstrapXdsAgent"] = request.EnableBootstrapXdsAgent
	}

	if !dara.IsNil(request.EnableCRHistory) {
		body["EnableCRHistory"] = request.EnableCRHistory
	}

	if !dara.IsNil(request.EnableNamespacesByDefault) {
		body["EnableNamespacesByDefault"] = request.EnableNamespacesByDefault
	}

	if !dara.IsNil(request.EnableSDSServer) {
		body["EnableSDSServer"] = request.EnableSDSServer
	}

	if !dara.IsNil(request.ExcludeIPRanges) {
		body["ExcludeIPRanges"] = request.ExcludeIPRanges
	}

	if !dara.IsNil(request.ExcludeInboundPorts) {
		body["ExcludeInboundPorts"] = request.ExcludeInboundPorts
	}

	if !dara.IsNil(request.ExcludeOutboundPorts) {
		body["ExcludeOutboundPorts"] = request.ExcludeOutboundPorts
	}

	if !dara.IsNil(request.ExistingCaCert) {
		body["ExistingCaCert"] = request.ExistingCaCert
	}

	if !dara.IsNil(request.ExistingCaKey) {
		body["ExistingCaKey"] = request.ExistingCaKey
	}

	if !dara.IsNil(request.ExistingRootCaCert) {
		body["ExistingRootCaCert"] = request.ExistingRootCaCert
	}

	if !dara.IsNil(request.FilterGatewayClusterConfig) {
		body["FilterGatewayClusterConfig"] = request.FilterGatewayClusterConfig
	}

	if !dara.IsNil(request.GatewayAPIEnabled) {
		body["GatewayAPIEnabled"] = request.GatewayAPIEnabled
	}

	if !dara.IsNil(request.HoldApplicationUntilProxyStarts) {
		body["HoldApplicationUntilProxyStarts"] = request.HoldApplicationUntilProxyStarts
	}

	if !dara.IsNil(request.Http10Enabled) {
		body["Http10Enabled"] = request.Http10Enabled
	}

	if !dara.IsNil(request.IncludeIPRanges) {
		body["IncludeIPRanges"] = request.IncludeIPRanges
	}

	if !dara.IsNil(request.IncludeInboundPorts) {
		body["IncludeInboundPorts"] = request.IncludeInboundPorts
	}

	if !dara.IsNil(request.IncludeOutboundPorts) {
		body["IncludeOutboundPorts"] = request.IncludeOutboundPorts
	}

	if !dara.IsNil(request.IntegrateKiali) {
		body["IntegrateKiali"] = request.IntegrateKiali
	}

	if !dara.IsNil(request.InterceptionMode) {
		body["InterceptionMode"] = request.InterceptionMode
	}

	if !dara.IsNil(request.KialiArmsAuthTokens) {
		body["KialiArmsAuthTokens"] = request.KialiArmsAuthTokens
	}

	if !dara.IsNil(request.KialiEnabled) {
		body["KialiEnabled"] = request.KialiEnabled
	}

	if !dara.IsNil(request.KialiServiceAnnotations) {
		body["KialiServiceAnnotations"] = request.KialiServiceAnnotations
	}

	if !dara.IsNil(request.Lifecycle) {
		body["Lifecycle"] = request.Lifecycle
	}

	if !dara.IsNil(request.LocalityLBConf) {
		body["LocalityLBConf"] = request.LocalityLBConf
	}

	if !dara.IsNil(request.LocalityLoadBalancing) {
		body["LocalityLoadBalancing"] = request.LocalityLoadBalancing
	}

	if !dara.IsNil(request.LogLevel) {
		body["LogLevel"] = request.LogLevel
	}

	if !dara.IsNil(request.MSEEnabled) {
		body["MSEEnabled"] = request.MSEEnabled
	}

	if !dara.IsNil(request.MultiBufferEnabled) {
		body["MultiBufferEnabled"] = request.MultiBufferEnabled
	}

	if !dara.IsNil(request.MultiBufferPollDelay) {
		body["MultiBufferPollDelay"] = request.MultiBufferPollDelay
	}

	if !dara.IsNil(request.MysqlFilterEnabled) {
		body["MysqlFilterEnabled"] = request.MysqlFilterEnabled
	}

	if !dara.IsNil(request.NFDEnabled) {
		body["NFDEnabled"] = request.NFDEnabled
	}

	if !dara.IsNil(request.NFDLabelPruned) {
		body["NFDLabelPruned"] = request.NFDLabelPruned
	}

	if !dara.IsNil(request.OPAInjectorCPULimit) {
		body["OPAInjectorCPULimit"] = request.OPAInjectorCPULimit
	}

	if !dara.IsNil(request.OPAInjectorCPURequirement) {
		body["OPAInjectorCPURequirement"] = request.OPAInjectorCPURequirement
	}

	if !dara.IsNil(request.OPAInjectorMemoryLimit) {
		body["OPAInjectorMemoryLimit"] = request.OPAInjectorMemoryLimit
	}

	if !dara.IsNil(request.OPAInjectorMemoryRequirement) {
		body["OPAInjectorMemoryRequirement"] = request.OPAInjectorMemoryRequirement
	}

	if !dara.IsNil(request.OPALimitCPU) {
		body["OPALimitCPU"] = request.OPALimitCPU
	}

	if !dara.IsNil(request.OPALimitMemory) {
		body["OPALimitMemory"] = request.OPALimitMemory
	}

	if !dara.IsNil(request.OPALogLevel) {
		body["OPALogLevel"] = request.OPALogLevel
	}

	if !dara.IsNil(request.OPARequestCPU) {
		body["OPARequestCPU"] = request.OPARequestCPU
	}

	if !dara.IsNil(request.OPARequestMemory) {
		body["OPARequestMemory"] = request.OPARequestMemory
	}

	if !dara.IsNil(request.OPAScopeInjected) {
		body["OPAScopeInjected"] = request.OPAScopeInjected
	}

	if !dara.IsNil(request.OpaEnabled) {
		body["OpaEnabled"] = request.OpaEnabled
	}

	if !dara.IsNil(request.OpenAgentPolicy) {
		body["OpenAgentPolicy"] = request.OpenAgentPolicy
	}

	if !dara.IsNil(request.OutboundTrafficPolicy) {
		body["OutboundTrafficPolicy"] = request.OutboundTrafficPolicy
	}

	if !dara.IsNil(request.PilotEnableQuicListeners) {
		body["PilotEnableQuicListeners"] = request.PilotEnableQuicListeners
	}

	if !dara.IsNil(request.PrometheusUrl) {
		body["PrometheusUrl"] = request.PrometheusUrl
	}

	if !dara.IsNil(request.ProxyInitCPUResourceLimit) {
		body["ProxyInitCPUResourceLimit"] = request.ProxyInitCPUResourceLimit
	}

	if !dara.IsNil(request.ProxyInitCPUResourceRequest) {
		body["ProxyInitCPUResourceRequest"] = request.ProxyInitCPUResourceRequest
	}

	if !dara.IsNil(request.ProxyInitMemoryResourceLimit) {
		body["ProxyInitMemoryResourceLimit"] = request.ProxyInitMemoryResourceLimit
	}

	if !dara.IsNil(request.ProxyInitMemoryResourceRequest) {
		body["ProxyInitMemoryResourceRequest"] = request.ProxyInitMemoryResourceRequest
	}

	if !dara.IsNil(request.ProxyLimitCPU) {
		body["ProxyLimitCPU"] = request.ProxyLimitCPU
	}

	if !dara.IsNil(request.ProxyLimitMemory) {
		body["ProxyLimitMemory"] = request.ProxyLimitMemory
	}

	if !dara.IsNil(request.ProxyRequestCPU) {
		body["ProxyRequestCPU"] = request.ProxyRequestCPU
	}

	if !dara.IsNil(request.ProxyRequestMemory) {
		body["ProxyRequestMemory"] = request.ProxyRequestMemory
	}

	if !dara.IsNil(request.ProxyStatsMatcher) {
		body["ProxyStatsMatcher"] = request.ProxyStatsMatcher
	}

	if !dara.IsNil(request.RedisFilterEnabled) {
		body["RedisFilterEnabled"] = request.RedisFilterEnabled
	}

	if !dara.IsNil(request.SMCEnabled) {
		body["SMCEnabled"] = request.SMCEnabled
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	if !dara.IsNil(request.SidecarInjectorLimitCPU) {
		body["SidecarInjectorLimitCPU"] = request.SidecarInjectorLimitCPU
	}

	if !dara.IsNil(request.SidecarInjectorLimitMemory) {
		body["SidecarInjectorLimitMemory"] = request.SidecarInjectorLimitMemory
	}

	if !dara.IsNil(request.SidecarInjectorRequestCPU) {
		body["SidecarInjectorRequestCPU"] = request.SidecarInjectorRequestCPU
	}

	if !dara.IsNil(request.SidecarInjectorRequestMemory) {
		body["SidecarInjectorRequestMemory"] = request.SidecarInjectorRequestMemory
	}

	if !dara.IsNil(request.SidecarInjectorWebhookAsYaml) {
		body["SidecarInjectorWebhookAsYaml"] = request.SidecarInjectorWebhookAsYaml
	}

	if !dara.IsNil(request.Telemetry) {
		body["Telemetry"] = request.Telemetry
	}

	if !dara.IsNil(request.TerminationDrainDuration) {
		body["TerminationDrainDuration"] = request.TerminationDrainDuration
	}

	if !dara.IsNil(request.ThriftFilterEnabled) {
		body["ThriftFilterEnabled"] = request.ThriftFilterEnabled
	}

	if !dara.IsNil(request.TraceCustomTags) {
		body["TraceCustomTags"] = request.TraceCustomTags
	}

	if !dara.IsNil(request.TraceMaxPathTagLength) {
		body["TraceMaxPathTagLength"] = request.TraceMaxPathTagLength
	}

	if !dara.IsNil(request.TraceSampling) {
		body["TraceSampling"] = request.TraceSampling
	}

	if !dara.IsNil(request.Tracing) {
		body["Tracing"] = request.Tracing
	}

	if !dara.IsNil(request.TracingOnExtZipkinLimitCPU) {
		body["TracingOnExtZipkinLimitCPU"] = request.TracingOnExtZipkinLimitCPU
	}

	if !dara.IsNil(request.TracingOnExtZipkinLimitMemory) {
		body["TracingOnExtZipkinLimitMemory"] = request.TracingOnExtZipkinLimitMemory
	}

	if !dara.IsNil(request.TracingOnExtZipkinReplicaCount) {
		body["TracingOnExtZipkinReplicaCount"] = request.TracingOnExtZipkinReplicaCount
	}

	if !dara.IsNil(request.TracingOnExtZipkinRequestCPU) {
		body["TracingOnExtZipkinRequestCPU"] = request.TracingOnExtZipkinRequestCPU
	}

	if !dara.IsNil(request.TracingOnExtZipkinRequestMemory) {
		body["TracingOnExtZipkinRequestMemory"] = request.TracingOnExtZipkinRequestMemory
	}

	if !dara.IsNil(request.WebAssemblyFilterEnabled) {
		body["WebAssemblyFilterEnabled"] = request.WebAssemblyFilterEnabled
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMeshFeature"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMeshFeatureResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the network configurations of multiple Kubernetes clusters in a Service Mesh (ASM) instance.
//
// @param tmpReq - UpdateMeshMultiClusterNetworkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMeshMultiClusterNetworkResponse
func (client *Client) UpdateMeshMultiClusterNetworkWithContext(ctx context.Context, tmpReq *UpdateMeshMultiClusterNetworkRequest, runtime *dara.RuntimeOptions) (_result *UpdateMeshMultiClusterNetworkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateMeshMultiClusterNetworkShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MultiClusterNetworks) {
		request.MultiClusterNetworksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MultiClusterNetworks, dara.String("MultiClusterNetworks"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MultiClusterNetworksShrink) {
		body["MultiClusterNetworks"] = request.MultiClusterNetworksShrink
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMeshMultiClusterNetwork"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMeshMultiClusterNetworkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configurations of sidecar proxies at the namespace level.
//
// @param tmpReq - UpdateNamespaceScopeSidecarConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNamespaceScopeSidecarConfigResponse
func (client *Client) UpdateNamespaceScopeSidecarConfigWithContext(ctx context.Context, tmpReq *UpdateNamespaceScopeSidecarConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateNamespaceScopeSidecarConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateNamespaceScopeSidecarConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ScaledSidecarResource) {
		request.ScaledSidecarResourceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ScaledSidecarResource, dara.String("ScaledSidecarResource"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Concurrency) {
		body["Concurrency"] = request.Concurrency
	}

	if !dara.IsNil(request.EnableCoreDump) {
		body["EnableCoreDump"] = request.EnableCoreDump
	}

	if !dara.IsNil(request.ExcludeIPRanges) {
		body["ExcludeIPRanges"] = request.ExcludeIPRanges
	}

	if !dara.IsNil(request.ExcludeInboundPorts) {
		body["ExcludeInboundPorts"] = request.ExcludeInboundPorts
	}

	if !dara.IsNil(request.ExcludeOutboundPorts) {
		body["ExcludeOutboundPorts"] = request.ExcludeOutboundPorts
	}

	if !dara.IsNil(request.HoldApplicationUntilProxyStarts) {
		body["HoldApplicationUntilProxyStarts"] = request.HoldApplicationUntilProxyStarts
	}

	if !dara.IsNil(request.IncludeIPRanges) {
		body["IncludeIPRanges"] = request.IncludeIPRanges
	}

	if !dara.IsNil(request.IncludeInboundPorts) {
		body["IncludeInboundPorts"] = request.IncludeInboundPorts
	}

	if !dara.IsNil(request.IncludeOutboundPorts) {
		body["IncludeOutboundPorts"] = request.IncludeOutboundPorts
	}

	if !dara.IsNil(request.InterceptionMode) {
		body["InterceptionMode"] = request.InterceptionMode
	}

	if !dara.IsNil(request.IstioDNSProxyEnabled) {
		body["IstioDNSProxyEnabled"] = request.IstioDNSProxyEnabled
	}

	if !dara.IsNil(request.Lifecycle) {
		body["Lifecycle"] = request.Lifecycle
	}

	if !dara.IsNil(request.LogLevel) {
		body["LogLevel"] = request.LogLevel
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PostStart) {
		body["PostStart"] = request.PostStart
	}

	if !dara.IsNil(request.PreStop) {
		body["PreStop"] = request.PreStop
	}

	if !dara.IsNil(request.Privileged) {
		body["Privileged"] = request.Privileged
	}

	if !dara.IsNil(request.ProxyInitAckSloCPUResourceLimit) {
		body["ProxyInitAckSloCPUResourceLimit"] = request.ProxyInitAckSloCPUResourceLimit
	}

	if !dara.IsNil(request.ProxyInitAckSloCPUResourceRequest) {
		body["ProxyInitAckSloCPUResourceRequest"] = request.ProxyInitAckSloCPUResourceRequest
	}

	if !dara.IsNil(request.ProxyInitAckSloMemoryResourceLimit) {
		body["ProxyInitAckSloMemoryResourceLimit"] = request.ProxyInitAckSloMemoryResourceLimit
	}

	if !dara.IsNil(request.ProxyInitAckSloMemoryResourceRequest) {
		body["ProxyInitAckSloMemoryResourceRequest"] = request.ProxyInitAckSloMemoryResourceRequest
	}

	if !dara.IsNil(request.ProxyInitCPUResourceLimit) {
		body["ProxyInitCPUResourceLimit"] = request.ProxyInitCPUResourceLimit
	}

	if !dara.IsNil(request.ProxyInitCPUResourceRequest) {
		body["ProxyInitCPUResourceRequest"] = request.ProxyInitCPUResourceRequest
	}

	if !dara.IsNil(request.ProxyInitMemoryResourceLimit) {
		body["ProxyInitMemoryResourceLimit"] = request.ProxyInitMemoryResourceLimit
	}

	if !dara.IsNil(request.ProxyInitMemoryResourceRequest) {
		body["ProxyInitMemoryResourceRequest"] = request.ProxyInitMemoryResourceRequest
	}

	if !dara.IsNil(request.ProxyMetadata) {
		body["ProxyMetadata"] = request.ProxyMetadata
	}

	if !dara.IsNil(request.ProxyStatsMatcher) {
		body["ProxyStatsMatcher"] = request.ProxyStatsMatcher
	}

	if !dara.IsNil(request.ReadinessFailureThreshold) {
		body["ReadinessFailureThreshold"] = request.ReadinessFailureThreshold
	}

	if !dara.IsNil(request.ReadinessInitialDelaySeconds) {
		body["ReadinessInitialDelaySeconds"] = request.ReadinessInitialDelaySeconds
	}

	if !dara.IsNil(request.ReadinessPeriodSeconds) {
		body["ReadinessPeriodSeconds"] = request.ReadinessPeriodSeconds
	}

	if !dara.IsNil(request.RuntimeValues) {
		body["RuntimeValues"] = request.RuntimeValues
	}

	if !dara.IsNil(request.SMCEnabled) {
		body["SMCEnabled"] = request.SMCEnabled
	}

	if !dara.IsNil(request.ScaledSidecarResourceShrink) {
		body["ScaledSidecarResource"] = request.ScaledSidecarResourceShrink
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	if !dara.IsNil(request.SidecarProxyAckSloCPUResourceLimit) {
		body["SidecarProxyAckSloCPUResourceLimit"] = request.SidecarProxyAckSloCPUResourceLimit
	}

	if !dara.IsNil(request.SidecarProxyAckSloCPUResourceRequest) {
		body["SidecarProxyAckSloCPUResourceRequest"] = request.SidecarProxyAckSloCPUResourceRequest
	}

	if !dara.IsNil(request.SidecarProxyAckSloMemoryResourceLimit) {
		body["SidecarProxyAckSloMemoryResourceLimit"] = request.SidecarProxyAckSloMemoryResourceLimit
	}

	if !dara.IsNil(request.SidecarProxyAckSloMemoryResourceRequest) {
		body["SidecarProxyAckSloMemoryResourceRequest"] = request.SidecarProxyAckSloMemoryResourceRequest
	}

	if !dara.IsNil(request.SidecarProxyCPUResourceLimit) {
		body["SidecarProxyCPUResourceLimit"] = request.SidecarProxyCPUResourceLimit
	}

	if !dara.IsNil(request.SidecarProxyCPUResourceRequest) {
		body["SidecarProxyCPUResourceRequest"] = request.SidecarProxyCPUResourceRequest
	}

	if !dara.IsNil(request.SidecarProxyMemoryResourceLimit) {
		body["SidecarProxyMemoryResourceLimit"] = request.SidecarProxyMemoryResourceLimit
	}

	if !dara.IsNil(request.SidecarProxyMemoryResourceRequest) {
		body["SidecarProxyMemoryResourceRequest"] = request.SidecarProxyMemoryResourceRequest
	}

	if !dara.IsNil(request.TerminationDrainDuration) {
		body["TerminationDrainDuration"] = request.TerminationDrainDuration
	}

	if !dara.IsNil(request.Tracing) {
		body["Tracing"] = request.Tracing
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateNamespaceScopeSidecarConfig"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateNamespaceScopeSidecarConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a lane.
//
// @param request - UpdateSwimLaneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSwimLaneResponse
func (client *Client) UpdateSwimLaneWithContext(ctx context.Context, request *UpdateSwimLaneRequest, runtime *dara.RuntimeOptions) (_result *UpdateSwimLaneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		body["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.LabelSelectorKey) {
		body["LabelSelectorKey"] = request.LabelSelectorKey
	}

	if !dara.IsNil(request.LabelSelectorValue) {
		body["LabelSelectorValue"] = request.LabelSelectorValue
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	if !dara.IsNil(request.ServicesList) {
		body["ServicesList"] = request.ServicesList
	}

	if !dara.IsNil(request.SwimLaneName) {
		body["SwimLaneName"] = request.SwimLaneName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSwimLane"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSwimLaneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information of a lane group.
//
// @param request - UpdateSwimLaneGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSwimLaneGroupResponse
func (client *Client) UpdateSwimLaneGroupWithContext(ctx context.Context, request *UpdateSwimLaneGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateSwimLaneGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FallbackTarget) {
		body["FallbackTarget"] = request.FallbackTarget
	}

	if !dara.IsNil(request.GroupName) {
		body["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.IngressRoutingStrategy) {
		body["IngressRoutingStrategy"] = request.IngressRoutingStrategy
	}

	if !dara.IsNil(request.ServiceLevelFallbackTarget) {
		body["ServiceLevelFallbackTarget"] = request.ServiceLevelFallbackTarget
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	if !dara.IsNil(request.ServicesList) {
		body["ServicesList"] = request.ServicesList
	}

	if !dara.IsNil(request.WeightedIngressRule) {
		body["WeightedIngressRule"] = request.WeightedIngressRule
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSwimLaneGroup"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSwimLaneGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新Waypoint
//
// @param request - UpdateWaypointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWaypointResponse
func (client *Client) UpdateWaypointWithContext(ctx context.Context, request *UpdateWaypointRequest, runtime *dara.RuntimeOptions) (_result *UpdateWaypointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.HPAEnabled) {
		body["HPAEnabled"] = request.HPAEnabled
	}

	if !dara.IsNil(request.HPAMaxReplicas) {
		body["HPAMaxReplicas"] = request.HPAMaxReplicas
	}

	if !dara.IsNil(request.HPAMinReplicas) {
		body["HPAMinReplicas"] = request.HPAMinReplicas
	}

	if !dara.IsNil(request.HPATargetCPU) {
		body["HPATargetCPU"] = request.HPATargetCPU
	}

	if !dara.IsNil(request.HPATargetMemory) {
		body["HPATargetMemory"] = request.HPATargetMemory
	}

	if !dara.IsNil(request.LimitCPU) {
		body["LimitCPU"] = request.LimitCPU
	}

	if !dara.IsNil(request.LimitMemory) {
		body["LimitMemory"] = request.LimitMemory
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PreferECI) {
		body["PreferECI"] = request.PreferECI
	}

	if !dara.IsNil(request.Replicas) {
		body["Replicas"] = request.Replicas
	}

	if !dara.IsNil(request.RequestCPU) {
		body["RequestCPU"] = request.RequestCPU
	}

	if !dara.IsNil(request.RequestMemory) {
		body["RequestMemory"] = request.RequestMemory
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWaypoint"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWaypointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades a Service Mesh (ASM) instance to Professional Edition that is commercially released.
//
// @param request - UpgradeMeshEditionPartiallyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeMeshEditionPartiallyResponse
func (client *Client) UpgradeMeshEditionPartiallyWithContext(ctx context.Context, request *UpgradeMeshEditionPartiallyRequest, runtime *dara.RuntimeOptions) (_result *UpgradeMeshEditionPartiallyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ASMGatewayContinue) {
		body["ASMGatewayContinue"] = request.ASMGatewayContinue
	}

	if !dara.IsNil(request.ExpectedVersion) {
		body["ExpectedVersion"] = request.ExpectedVersion
	}

	if !dara.IsNil(request.PreCheck) {
		body["PreCheck"] = request.PreCheck
	}

	if !dara.IsNil(request.ServiceMeshId) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	if !dara.IsNil(request.SwitchToPro) {
		body["SwitchToPro"] = request.SwitchToPro
	}

	if !dara.IsNil(request.UpgradeGatewayRecords) {
		body["UpgradeGatewayRecords"] = request.UpgradeGatewayRecords
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeMeshEditionPartially"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeMeshEditionPartiallyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the version of a Service Mesh (ASM) instance.
//
// @param request - UpgradeMeshVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeMeshVersionResponse
func (client *Client) UpgradeMeshVersionWithContext(ctx context.Context, request *UpgradeMeshVersionRequest, runtime *dara.RuntimeOptions) (_result *UpgradeMeshVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PreCheck) {
		query["PreCheck"] = request.PreCheck
	}

	if !dara.IsNil(request.ServiceMeshId) {
		query["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeMeshVersion"),
		Version:     dara.String("2020-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeMeshVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
