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
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("hsm"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Enables or disables the audit log feature and delivers audit logs to buckets.
//
// Description:
//
//	  The region of the bucket must be the same as the region where the security audit feature is enabled.
//
//		- If the security audit feature is enabled, do not delete Object Storage Service (OSS) buckets. If you delete OSS buckets, audit logs fail to be delivered.
//
//		- Only electronic virtual security modules (EVSMs) and general virtual security modules (GVSMs) within the Chinese mainland support the security audit feature.
//
// @param request - ConfigAuditLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigAuditLogResponse
func (client *Client) ConfigAuditLogWithOptions(request *ConfigAuditLogRequest, runtime *dara.RuntimeOptions) (_result *ConfigAuditLogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuditAction) {
		query["AuditAction"] = request.AuditAction
	}

	if !dara.IsNil(request.AuditOssBucket) {
		query["AuditOssBucket"] = request.AuditOssBucket
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigAuditLog"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigAuditLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the audit log feature and delivers audit logs to buckets.
//
// Description:
//
//	  The region of the bucket must be the same as the region where the security audit feature is enabled.
//
//		- If the security audit feature is enabled, do not delete Object Storage Service (OSS) buckets. If you delete OSS buckets, audit logs fail to be delivered.
//
//		- Only electronic virtual security modules (EVSMs) and general virtual security modules (GVSMs) within the Chinese mainland support the security audit feature.
//
// @param request - ConfigAuditLogRequest
//
// @return ConfigAuditLogResponse
func (client *Client) ConfigAuditLog(request *ConfigAuditLogRequest) (_result *ConfigAuditLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigAuditLogResponse{}
	_body, _err := client.ConfigAuditLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the name and description of a backup.
//
// @param request - ConfigBackupRemarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigBackupRemarkResponse
func (client *Client) ConfigBackupRemarkWithOptions(request *ConfigBackupRemarkRequest, runtime *dara.RuntimeOptions) (_result *ConfigBackupRemarkResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupId) {
		query["BackupId"] = request.BackupId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigBackupRemark"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigBackupRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the name and description of a backup.
//
// @param request - ConfigBackupRemarkRequest
//
// @return ConfigBackupRemarkResponse
func (client *Client) ConfigBackupRemark(request *ConfigBackupRemarkRequest) (_result *ConfigBackupRemarkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigBackupRemarkResponse{}
	_body, _err := client.ConfigBackupRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the execution mode of a backup task.
//
// Description:
//
// Only hardware security modules (HSMs) in the Chinese mainland support the operation.
//
// @param request - ConfigBackupTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigBackupTaskResponse
func (client *Client) ConfigBackupTaskWithOptions(request *ConfigBackupTaskRequest, runtime *dara.RuntimeOptions) (_result *ConfigBackupTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupHourInDay) {
		query["BackupHourInDay"] = request.BackupHourInDay
	}

	if !dara.IsNil(request.BackupId) {
		query["BackupId"] = request.BackupId
	}

	if !dara.IsNil(request.BackupPeriod) {
		query["BackupPeriod"] = request.BackupPeriod
	}

	if !dara.IsNil(request.Manual2PeriodicList) {
		query["Manual2PeriodicList"] = request.Manual2PeriodicList
	}

	if !dara.IsNil(request.Periodic2ManualList) {
		query["Periodic2ManualList"] = request.Periodic2ManualList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigBackupTask"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigBackupTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the execution mode of a backup task.
//
// Description:
//
// Only hardware security modules (HSMs) in the Chinese mainland support the operation.
//
// @param request - ConfigBackupTaskRequest
//
// @return ConfigBackupTaskResponse
func (client *Client) ConfigBackupTask(request *ConfigBackupTaskRequest) (_result *ConfigBackupTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigBackupTaskResponse{}
	_body, _err := client.ConfigBackupTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures a certificate for a cluster of hardware security modules (HSMs) outside the Chinese mainland.
//
// Description:
//
// For more information about how to create a self-signed certificate and a cluster certificate on an Elastic Compute Service (ECS) instance, see [Create a NIST FIPS-validated GVSM cluster](https://help.aliyun.com/document_detail/293585.html).
//
// @param request - ConfigClusterCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigClusterCertificateResponse
func (client *Client) ConfigClusterCertificateWithOptions(request *ConfigClusterCertificateRequest, runtime *dara.RuntimeOptions) (_result *ConfigClusterCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterCertificate) {
		body["ClusterCertificate"] = request.ClusterCertificate
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.IssuerCertificate) {
		body["IssuerCertificate"] = request.IssuerCertificate
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigClusterCertificate"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigClusterCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a certificate for a cluster of hardware security modules (HSMs) outside the Chinese mainland.
//
// Description:
//
// For more information about how to create a self-signed certificate and a cluster certificate on an Elastic Compute Service (ECS) instance, see [Create a NIST FIPS-validated GVSM cluster](https://help.aliyun.com/document_detail/293585.html).
//
// @param request - ConfigClusterCertificateRequest
//
// @return ConfigClusterCertificateResponse
func (client *Client) ConfigClusterCertificate(request *ConfigClusterCertificateRequest) (_result *ConfigClusterCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigClusterCertificateResponse{}
	_body, _err := client.ConfigClusterCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the name of a cluster.
//
// @param request - ConfigClusterNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigClusterNameResponse
func (client *Client) ConfigClusterNameWithOptions(request *ConfigClusterNameRequest, runtime *dara.RuntimeOptions) (_result *ConfigClusterNameResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ClusterName) {
		body["ClusterName"] = request.ClusterName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigClusterName"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigClusterNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the name of a cluster.
//
// @param request - ConfigClusterNameRequest
//
// @return ConfigClusterNameResponse
func (client *Client) ConfigClusterName(request *ConfigClusterNameRequest) (_result *ConfigClusterNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigClusterNameResponse{}
	_body, _err := client.ConfigClusterNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a list of vSwitches that are associated with a hardware security module (HSM) cluster.
//
// Description:
//
// You can call the operation to configure all vSwitches that are associated with a HSM cluster. You can only add new vSwitches. You cannot delete vSwitches.
//
// @param tmpReq - ConfigClusterSubnetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigClusterSubnetResponse
func (client *Client) ConfigClusterSubnetWithOptions(tmpReq *ConfigClusterSubnetRequest, runtime *dara.RuntimeOptions) (_result *ConfigClusterSubnetResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ConfigClusterSubnetShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.VSwitchIds) {
		request.VSwitchIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VSwitchIds, dara.String("VSwitchIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VSwitchIdsShrink) {
		body["VSwitchIds"] = request.VSwitchIdsShrink
	}

	if !dara.IsNil(request.VpcId) {
		body["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigClusterSubnet"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigClusterSubnetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a list of vSwitches that are associated with a hardware security module (HSM) cluster.
//
// Description:
//
// You can call the operation to configure all vSwitches that are associated with a HSM cluster. You can only add new vSwitches. You cannot delete vSwitches.
//
// @param request - ConfigClusterSubnetRequest
//
// @return ConfigClusterSubnetResponse
func (client *Client) ConfigClusterSubnet(request *ConfigClusterSubnetRequest) (_result *ConfigClusterSubnetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigClusterSubnetResponse{}
	_body, _err := client.ConfigClusterSubnetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the IP address whitelist of a cluster.
//
// Description:
//
// The IP address whitelist of a cluster has a higher priority than the IP address whitelist of a hardware security module (HSM) in the cluster. In cluster mode, we recommend that you create an IP address whitelist for your cluster. You do not need to create an IP address for the HSM in the cluster.
//
// @param request - ConfigClusterWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigClusterWhitelistResponse
func (client *Client) ConfigClusterWhitelistWithOptions(request *ConfigClusterWhitelistRequest, runtime *dara.RuntimeOptions) (_result *ConfigClusterWhitelistResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Whitelist) {
		body["Whitelist"] = request.Whitelist
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigClusterWhitelist"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigClusterWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the IP address whitelist of a cluster.
//
// Description:
//
// The IP address whitelist of a cluster has a higher priority than the IP address whitelist of a hardware security module (HSM) in the cluster. In cluster mode, we recommend that you create an IP address whitelist for your cluster. You do not need to create an IP address for the HSM in the cluster.
//
// @param request - ConfigClusterWhitelistRequest
//
// @return ConfigClusterWhitelistResponse
func (client *Client) ConfigClusterWhitelist(request *ConfigClusterWhitelistRequest) (_result *ConfigClusterWhitelistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigClusterWhitelistResponse{}
	_body, _err := client.ConfigClusterWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the description of an image.
//
// @param request - ConfigImageRemarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigImageRemarkResponse
func (client *Client) ConfigImageRemarkWithOptions(request *ConfigImageRemarkRequest, runtime *dara.RuntimeOptions) (_result *ConfigImageRemarkResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigImageRemark"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigImageRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description of an image.
//
// @param request - ConfigImageRemarkRequest
//
// @return ConfigImageRemarkResponse
func (client *Client) ConfigImageRemark(request *ConfigImageRemarkRequest) (_result *ConfigImageRemarkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigImageRemarkResponse{}
	_body, _err := client.ConfigImageRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the virtual private cloud (VPC) endpoint of a hardware security module (HSM).
//
// Description:
//
// After you add an HSM to a cluster, you cannot modify the VPC endpoint of the HSM.
//
// @param request - ConfigInstanceIpAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigInstanceIpAddressResponse
func (client *Client) ConfigInstanceIpAddressWithOptions(request *ConfigInstanceIpAddressRequest, runtime *dara.RuntimeOptions) (_result *ConfigInstanceIpAddressResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Ip) {
		body["Ip"] = request.Ip
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VSwitchId) {
		body["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VpcId) {
		body["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigInstanceIpAddress"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigInstanceIpAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the virtual private cloud (VPC) endpoint of a hardware security module (HSM).
//
// Description:
//
// After you add an HSM to a cluster, you cannot modify the VPC endpoint of the HSM.
//
// @param request - ConfigInstanceIpAddressRequest
//
// @return ConfigInstanceIpAddressResponse
func (client *Client) ConfigInstanceIpAddress(request *ConfigInstanceIpAddressRequest) (_result *ConfigInstanceIpAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigInstanceIpAddressResponse{}
	_body, _err := client.ConfigInstanceIpAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the description of a hardware security module (HSM).
//
// @param request - ConfigInstanceRemarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigInstanceRemarkResponse
func (client *Client) ConfigInstanceRemarkWithOptions(request *ConfigInstanceRemarkRequest, runtime *dara.RuntimeOptions) (_result *ConfigInstanceRemarkResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Remark) {
		body["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigInstanceRemark"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigInstanceRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description of a hardware security module (HSM).
//
// @param request - ConfigInstanceRemarkRequest
//
// @return ConfigInstanceRemarkResponse
func (client *Client) ConfigInstanceRemark(request *ConfigInstanceRemarkRequest) (_result *ConfigInstanceRemarkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigInstanceRemarkResponse{}
	_body, _err := client.ConfigInstanceRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the IP address whitelist of a hardware security module (HSM).
//
// Description:
//
// You can configure the IP address whitelist for HSMs that are not added to a cluster and are in the ACTIVE state.
//
// @param request - ConfigInstanceWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigInstanceWhitelistResponse
func (client *Client) ConfigInstanceWhitelistWithOptions(request *ConfigInstanceWhitelistRequest, runtime *dara.RuntimeOptions) (_result *ConfigInstanceWhitelistResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Whitelist) {
		body["Whitelist"] = request.Whitelist
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigInstanceWhitelist"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigInstanceWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the IP address whitelist of a hardware security module (HSM).
//
// Description:
//
// You can configure the IP address whitelist for HSMs that are not added to a cluster and are in the ACTIVE state.
//
// @param request - ConfigInstanceWhitelistRequest
//
// @return ConfigInstanceWhitelistResponse
func (client *Client) ConfigInstanceWhitelist(request *ConfigInstanceWhitelistRequest) (_result *ConfigInstanceWhitelistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigInstanceWhitelistResponse{}
	_body, _err := client.ConfigInstanceWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Copies an image to another region.
//
// Description:
//
// This operation requires that the destination region does not have the same image. This operation is available only for hardware security modules (HSMs) outside the Chinese mainland.
//
// @param request - CopyImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyImageResponse
func (client *Client) CopyImageWithOptions(request *CopyImageRequest, runtime *dara.RuntimeOptions) (_result *CopyImageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageUid) {
		body["ImageUid"] = request.ImageUid
	}

	if !dara.IsNil(request.TargetRegionId) {
		body["TargetRegionId"] = request.TargetRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CopyImage"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CopyImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Copies an image to another region.
//
// Description:
//
// This operation requires that the destination region does not have the same image. This operation is available only for hardware security modules (HSMs) outside the Chinese mainland.
//
// @param request - CopyImageRequest
//
// @return CopyImageResponse
func (client *Client) CopyImage(request *CopyImageRequest) (_result *CopyImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CopyImageResponse{}
	_body, _err := client.CopyImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a cluster by specifying a hardware security module (HSM) as the master HSM.
//
// Description:
//
// The master HSM that you specify to create a cluster must be in the ACTIVE state.
//
// @param request - CreateClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClusterResponse
func (client *Client) CreateClusterWithOptions(request *CreateClusterRequest, runtime *dara.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterName) {
		body["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.MasterInstanceId) {
		body["MasterInstanceId"] = request.MasterInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCluster"),
		Version:     dara.String("2023-11-13"),
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
// Creates a cluster by specifying a hardware security module (HSM) as the master HSM.
//
// Description:
//
// The master HSM that you specify to create a cluster must be in the ACTIVE state.
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
// Deletes a hardware security module (HSM) cluster.
//
// Description:
//
// You can delete an HSM only if the cluster does not contain HSMs.
//
// @param request - DeleteClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClusterResponse
func (client *Client) DeleteClusterWithOptions(request *DeleteClusterRequest, runtime *dara.RuntimeOptions) (_result *DeleteClusterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCluster"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a hardware security module (HSM) cluster.
//
// Description:
//
// You can delete an HSM only if the cluster does not contain HSMs.
//
// @param request - DeleteClusterRequest
//
// @return DeleteClusterResponse
func (client *Client) DeleteCluster(request *DeleteClusterRequest) (_result *DeleteClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteClusterResponse{}
	_body, _err := client.DeleteClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the regions that are supported by Cloud Hardware Security Module.
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
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2023-11-13"),
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
// Queries the regions that are supported by Cloud Hardware Security Module.
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
// 下载集群托管证书
//
// Description:
//
// ## 请求说明
//
// - 该API允许用户获取特定集群的管理证书。
//
// - 返回的数据是经过base64编码的证书内容。
//
// @param request - DownloadClusterManagedCertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DownloadClusterManagedCertResponse
func (client *Client) DownloadClusterManagedCertWithOptions(request *DownloadClusterManagedCertRequest, runtime *dara.RuntimeOptions) (_result *DownloadClusterManagedCertResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DownloadClusterManagedCert"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DownloadClusterManagedCertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 下载集群托管证书
//
// Description:
//
// ## 请求说明
//
// - 该API允许用户获取特定集群的管理证书。
//
// - 返回的数据是经过base64编码的证书内容。
//
// @param request - DownloadClusterManagedCertRequest
//
// @return DownloadClusterManagedCertResponse
func (client *Client) DownloadClusterManagedCert(request *DownloadClusterManagedCertRequest) (_result *DownloadClusterManagedCertResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DownloadClusterManagedCertResponse{}
	_body, _err := client.DownloadClusterManagedCertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds a backup to a specified hardware security module (HSM).
//
// Description:
//
// This operation is available only for backups in the Chinese mainland.
//
// @param request - EnableBackupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableBackupResponse
func (client *Client) EnableBackupWithOptions(request *EnableBackupRequest, runtime *dara.RuntimeOptions) (_result *EnableBackupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupId) {
		query["BackupId"] = request.BackupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableBackup"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableBackupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds a backup to a specified hardware security module (HSM).
//
// Description:
//
// This operation is available only for backups in the Chinese mainland.
//
// @param request - EnableBackupRequest
//
// @return EnableBackupResponse
func (client *Client) EnableBackup(request *EnableBackupRequest) (_result *EnableBackupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableBackupResponse{}
	_body, _err := client.EnableBackupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Exports the image for a specified hardware security module (HSM).
//
// @param request - ExportImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportImageResponse
func (client *Client) ExportImageWithOptions(request *ExportImageRequest, runtime *dara.RuntimeOptions) (_result *ExportImageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportImage"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Exports the image for a specified hardware security module (HSM).
//
// @param request - ExportImageRequest
//
// @return ExportImageResponse
func (client *Client) ExportImage(request *ExportImageRequest) (_result *ExportImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportImageResponse{}
	_body, _err := client.ExportImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of the audit log feature in the current region.
//
// @param request - GetAuditLogStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAuditLogStatusResponse
func (client *Client) GetAuditLogStatusWithOptions(request *GetAuditLogStatusRequest, runtime *dara.RuntimeOptions) (_result *GetAuditLogStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GetOssBucket) {
		query["GetOssBucket"] = request.GetOssBucket
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAuditLogStatus"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAuditLogStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the audit log feature in the current region.
//
// @param request - GetAuditLogStatusRequest
//
// @return GetAuditLogStatusResponse
func (client *Client) GetAuditLogStatus(request *GetAuditLogStatusRequest) (_result *GetAuditLogStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAuditLogStatusResponse{}
	_body, _err := client.GetAuditLogStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a specified backup.
//
// @param request - GetBackupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBackupResponse
func (client *Client) GetBackupWithOptions(request *GetBackupRequest, runtime *dara.RuntimeOptions) (_result *GetBackupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupId) {
		query["BackupId"] = request.BackupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBackup"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBackupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a specified backup.
//
// @param request - GetBackupRequest
//
// @return GetBackupResponse
func (client *Client) GetBackup(request *GetBackupRequest) (_result *GetBackupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBackupResponse{}
	_body, _err := client.GetBackupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about a specified cluster.
//
// @param request - GetClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClusterResponse
func (client *Client) GetClusterWithOptions(request *GetClusterRequest, runtime *dara.RuntimeOptions) (_result *GetClusterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCluster"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a specified cluster.
//
// @param request - GetClusterRequest
//
// @return GetClusterResponse
func (client *Client) GetCluster(request *GetClusterRequest) (_result *GetClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetClusterResponse{}
	_body, _err := client.GetClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about an image.
//
// @param request - GetImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageResponse
func (client *Client) GetImageWithOptions(request *GetImageRequest, runtime *dara.RuntimeOptions) (_result *GetImageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetImage"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about an image.
//
// @param request - GetImageRequest
//
// @return GetImageResponse
func (client *Client) GetImage(request *GetImageRequest) (_result *GetImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetImageResponse{}
	_body, _err := client.GetImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about a specified hardware security module (HSM).
//
// @param request - GetInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceResponse
func (client *Client) GetInstanceWithOptions(request *GetInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstance"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a specified hardware security module (HSM).
//
// @param request - GetInstanceRequest
//
// @return GetInstanceResponse
func (client *Client) GetInstance(request *GetInstanceRequest) (_result *GetInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceResponse{}
	_body, _err := client.GetInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an asynchronous task.
//
// @param request - GetJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobResponse
func (client *Client) GetJobWithOptions(request *GetJobRequest, runtime *dara.RuntimeOptions) (_result *GetJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetJob"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an asynchronous task.
//
// @param request - GetJobRequest
//
// @return GetJobResponse
func (client *Client) GetJob(request *GetJobRequest) (_result *GetJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetJobResponse{}
	_body, _err := client.GetJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Authorizes Cloud Hardware Security Module to deliver logs.
//
// @param request - InitializeAuditLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitializeAuditLogResponse
func (client *Client) InitializeAuditLogWithOptions(runtime *dara.RuntimeOptions) (_result *InitializeAuditLogResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("InitializeAuditLog"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InitializeAuditLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Authorizes Cloud Hardware Security Module to deliver logs.
//
// @return InitializeAuditLogResponse
func (client *Client) InitializeAuditLog() (_result *InitializeAuditLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InitializeAuditLogResponse{}
	_body, _err := client.InitializeAuditLogWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Initializes a cluster.
//
// Description:
//
//	  The cluster is not initialized, but the master hardware security module (HSM) of the cluster is initialized.
//
//		- Two or more vSwitches are configured for the cluster.
//
// @param request - InitializeClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitializeClusterResponse
func (client *Client) InitializeClusterWithOptions(request *InitializeClusterRequest, runtime *dara.RuntimeOptions) (_result *InitializeClusterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InitializeCluster"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InitializeClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initializes a cluster.
//
// Description:
//
//	  The cluster is not initialized, but the master hardware security module (HSM) of the cluster is initialized.
//
//		- Two or more vSwitches are configured for the cluster.
//
// @param request - InitializeClusterRequest
//
// @return InitializeClusterResponse
func (client *Client) InitializeCluster(request *InitializeClusterRequest) (_result *InitializeClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InitializeClusterResponse{}
	_body, _err := client.InitializeClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a hardware security module (HSM) to the current cluster.
//
// Description:
//
//	  You can add the HSM to only the cluster that is in the INITIALIZED state.
//
//		- The HSM that you want to add to the cluster is enabled or disabled and is not initialized.
//
// @param request - JoinClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return JoinClusterResponse
func (client *Client) JoinClusterWithOptions(request *JoinClusterRequest, runtime *dara.RuntimeOptions) (_result *JoinClusterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("JoinCluster"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &JoinClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a hardware security module (HSM) to the current cluster.
//
// Description:
//
//	  You can add the HSM to only the cluster that is in the INITIALIZED state.
//
//		- The HSM that you want to add to the cluster is enabled or disabled and is not initialized.
//
// @param request - JoinClusterRequest
//
// @return JoinClusterResponse
func (client *Client) JoinCluster(request *JoinClusterRequest) (_result *JoinClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &JoinClusterResponse{}
	_body, _err := client.JoinClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a hardware security module (HSM) from the current cluster.
//
// Description:
//
//	  If non-master HSMs exist in a cluster, the master HSM cannot be removed from the cluster.
//
//		- After the master HSM is removed from a cluster, the cluster enters the TO_DELETE state and cannot be restored to the available state. Proceed with caution.
//
// @param request - LeaveClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LeaveClusterResponse
func (client *Client) LeaveClusterWithOptions(request *LeaveClusterRequest, runtime *dara.RuntimeOptions) (_result *LeaveClusterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LeaveCluster"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LeaveClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a hardware security module (HSM) from the current cluster.
//
// Description:
//
//	  If non-master HSMs exist in a cluster, the master HSM cannot be removed from the cluster.
//
//		- After the master HSM is removed from a cluster, the cluster enters the TO_DELETE state and cannot be restored to the available state. Proceed with caution.
//
// @param request - LeaveClusterRequest
//
// @return LeaveClusterResponse
func (client *Client) LeaveCluster(request *LeaveClusterRequest) (_result *LeaveClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &LeaveClusterResponse{}
	_body, _err := client.LeaveClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the backups that meet the query conditions.
//
// @param request - ListBackupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBackupsResponse
func (client *Client) ListBackupsWithOptions(request *ListBackupsRequest, runtime *dara.RuntimeOptions) (_result *ListBackupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupId) {
		query["BackupId"] = request.BackupId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
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
		Action:      dara.String("ListBackups"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBackupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the backups that meet the query conditions.
//
// @param request - ListBackupsRequest
//
// @return ListBackupsResponse
func (client *Client) ListBackups(request *ListBackupsRequest) (_result *ListBackupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBackupsResponse{}
	_body, _err := client.ListBackupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the clusters that meet the query conditions.
//
// @param request - ListClustersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClustersResponse
func (client *Client) ListClustersWithOptions(request *ListClustersRequest, runtime *dara.RuntimeOptions) (_result *ListClustersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
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
		Action:      dara.String("ListClusters"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the clusters that meet the query conditions.
//
// @param request - ListClustersRequest
//
// @return ListClustersResponse
func (client *Client) ListClusters(request *ListClustersRequest) (_result *ListClustersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListClustersResponse{}
	_body, _err := client.ListClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the images that meet the specified conditions.
//
// @param request - ListImagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListImagesResponse
func (client *Client) ListImagesWithOptions(request *ListImagesRequest, runtime *dara.RuntimeOptions) (_result *ListImagesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupId) {
		query["BackupId"] = request.BackupId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
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
		Action:      dara.String("ListImages"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the images that meet the specified conditions.
//
// @param request - ListImagesRequest
//
// @return ListImagesResponse
func (client *Client) ListImages(request *ListImagesRequest) (_result *ListImagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListImagesResponse{}
	_body, _err := client.ListImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the hardware security modules (HSMs) that meet the query conditions.
//
// @param request - ListInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithOptions(request *ListInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TenantIsolationType) {
		body["TenantIsolationType"] = request.TenantIsolationType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstances"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the hardware security modules (HSMs) that meet the query conditions.
//
// @param request - ListInstancesRequest
//
// @return ListInstancesResponse
func (client *Client) ListInstances(request *ListInstancesRequest) (_result *ListInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInstancesResponse{}
	_body, _err := client.ListInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Moves a resource to a new resource group.
//
// @param request - MoveResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroupWithOptions(request *MoveResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
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
		Action:      dara.String("MoveResourceGroup"),
		Version:     dara.String("2023-11-13"),
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
// Moves a resource to a new resource group.
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
// Disables a hardware security module (HSM).
//
// Description:
//
// After you disable an HSM, the relevant service operations fail. Proceed with caution.
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
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PauseInstance"),
		Version:     dara.String("2023-11-13"),
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
// Disables a hardware security module (HSM).
//
// Description:
//
// After you disable an HSM, the relevant service operations fail. Proceed with caution.
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
// 快速部署集群
//
// @param tmpReq - QuickDeployClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuickDeployClusterResponse
func (client *Client) QuickDeployClusterWithOptions(tmpReq *QuickDeployClusterRequest, runtime *dara.RuntimeOptions) (_result *QuickDeployClusterResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &QuickDeployClusterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceList) {
		request.InstanceListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceList, dara.String("InstanceList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.VSwitchIdList) {
		request.VSwitchIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VSwitchIdList, dara.String("VSwitchIdList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.WhiteList) {
		request.WhiteListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WhiteList, dara.String("WhiteList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CertManaged) {
		query["CertManaged"] = request.CertManaged
	}

	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.InstanceListShrink) {
		query["InstanceList"] = request.InstanceListShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VSwitchIdListShrink) {
		query["VSwitchIdList"] = request.VSwitchIdListShrink
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.WhiteListShrink) {
		query["WhiteList"] = request.WhiteListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuickDeployCluster"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuickDeployClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 快速部署集群
//
// @param request - QuickDeployClusterRequest
//
// @return QuickDeployClusterResponse
func (client *Client) QuickDeployCluster(request *QuickDeployClusterRequest) (_result *QuickDeployClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuickDeployClusterResponse{}
	_body, _err := client.QuickDeployClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Initializes a hardware security module (HSM).
//
// Description:
//
// This operation is supported only for general virtual security modules (GVSMs) in the Chinese mainland.
//
// @param request - QuickInitInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuickInitInstanceResponse
func (client *Client) QuickInitInstanceWithOptions(request *QuickInitInstanceRequest, runtime *dara.RuntimeOptions) (_result *QuickInitInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuickInitInstance"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuickInitInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initializes a hardware security module (HSM).
//
// Description:
//
// This operation is supported only for general virtual security modules (GVSMs) in the Chinese mainland.
//
// @param request - QuickInitInstanceRequest
//
// @return QuickInitInstanceResponse
func (client *Client) QuickInitInstance(request *QuickInitInstanceRequest) (_result *QuickInitInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuickInitInstanceResponse{}
	_body, _err := client.QuickInitInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disassociates a backup from a hardware security module (HSM).
//
// Description:
//
// This operation is available only for HSMs in the Chinese mainland.
//
// @param request - ResetBackupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetBackupResponse
func (client *Client) ResetBackupWithOptions(request *ResetBackupRequest, runtime *dara.RuntimeOptions) (_result *ResetBackupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupId) {
		query["BackupId"] = request.BackupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetBackup"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetBackupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates a backup from a hardware security module (HSM).
//
// Description:
//
// This operation is available only for HSMs in the Chinese mainland.
//
// @param request - ResetBackupRequest
//
// @return ResetBackupResponse
func (client *Client) ResetBackup(request *ResetBackupRequest) (_result *ResetBackupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetBackupResponse{}
	_body, _err := client.ResetBackupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resets a hardware security module (HSM).
//
// Description:
//
// After an HSM is reset, all related data is deleted and cannot be recovered. Proceed with caution.
//
// @param request - ResetInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetInstanceResponse
func (client *Client) ResetInstanceWithOptions(request *ResetInstanceRequest, runtime *dara.RuntimeOptions) (_result *ResetInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetInstance"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets a hardware security module (HSM).
//
// Description:
//
// After an HSM is reset, all related data is deleted and cannot be recovered. Proceed with caution.
//
// @param request - ResetInstanceRequest
//
// @return ResetInstanceResponse
func (client *Client) ResetInstance(request *ResetInstanceRequest) (_result *ResetInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetInstanceResponse{}
	_body, _err := client.ResetInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restores a hardware security module (HSM) by using an image.
//
// Description:
//
// You can use images to restore only HSMs that are paused or disabled.
//
// @param request - RestoreInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestoreInstanceResponse
func (client *Client) RestoreInstanceWithOptions(request *RestoreInstanceRequest, runtime *dara.RuntimeOptions) (_result *RestoreInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestoreInstance"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestoreInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restores a hardware security module (HSM) by using an image.
//
// Description:
//
// You can use images to restore only HSMs that are paused or disabled.
//
// @param request - RestoreInstanceRequest
//
// @return RestoreInstanceResponse
func (client *Client) RestoreInstance(request *RestoreInstanceRequest) (_result *RestoreInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RestoreInstanceResponse{}
	_body, _err := client.RestoreInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resumes a disabled hardware security module (HSM).
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
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumeInstance"),
		Version:     dara.String("2023-11-13"),
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
// Resumes a disabled hardware security module (HSM).
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
// 轮转集群托管证书
//
// Description:
//
// ## 请求说明
//
// 该API用于触发指定集群的管理证书轮转过程。通过提供`ClusterId`参数，可以指定需要进行证书轮转的集群。此操作有助于提高集群的安全性，建议定期执行。
//
// ### 注意事项
//
// - 确保提供的`ClusterId`是有效的，并且用户具有对该集群的操作权限。
//
// - 证书轮转可能会影响依赖于旧证书的服务，请在适当的时间窗口内执行此操作。
//
// @param request - RotateClusterManagedCertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RotateClusterManagedCertResponse
func (client *Client) RotateClusterManagedCertWithOptions(request *RotateClusterManagedCertRequest, runtime *dara.RuntimeOptions) (_result *RotateClusterManagedCertResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RotateClusterManagedCert"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RotateClusterManagedCertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 轮转集群托管证书
//
// Description:
//
// ## 请求说明
//
// 该API用于触发指定集群的管理证书轮转过程。通过提供`ClusterId`参数，可以指定需要进行证书轮转的集群。此操作有助于提高集群的安全性，建议定期执行。
//
// ### 注意事项
//
// - 确保提供的`ClusterId`是有效的，并且用户具有对该集群的操作权限。
//
// - 证书轮转可能会影响依赖于旧证书的服务，请在适当的时间窗口内执行此操作。
//
// @param request - RotateClusterManagedCertRequest
//
// @return RotateClusterManagedCertResponse
func (client *Client) RotateClusterManagedCert(request *RotateClusterManagedCertRequest) (_result *RotateClusterManagedCertResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RotateClusterManagedCertResponse{}
	_body, _err := client.RotateClusterManagedCertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Promotes a slave hardware security module (HSM) to the master HSM within the cluster. Clusters that are manually synchronized in the Chinese Mainland do not support this operation.
//
// @param request - SwitchClusterMasterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchClusterMasterResponse
func (client *Client) SwitchClusterMasterWithOptions(request *SwitchClusterMasterRequest, runtime *dara.RuntimeOptions) (_result *SwitchClusterMasterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchClusterMaster"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchClusterMasterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Promotes a slave hardware security module (HSM) to the master HSM within the cluster. Clusters that are manually synchronized in the Chinese Mainland do not support this operation.
//
// @param request - SwitchClusterMasterRequest
//
// @return SwitchClusterMasterResponse
func (client *Client) SwitchClusterMaster(request *SwitchClusterMasterRequest) (_result *SwitchClusterMasterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SwitchClusterMasterResponse{}
	_body, _err := client.SwitchClusterMasterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Synchronizes the data of hardware security modules (HSMs) in a cluster.
//
// Description:
//
//	  This operation is used for manually synchronizing data within clusters located in the Chinese Mainland. For clusters outside the Chinese Mainland, automatic data synchronization is supported, and this operation is unnecessary. If you attempt to use this operation, a 400 error code will be returned.
//
//		- The data synchronization takes approximately 5 minutes. To avoid service interruptions, we recommend performing this operation during off-peak hours.
//
// @param request - SyncClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncClusterResponse
func (client *Client) SyncClusterWithOptions(request *SyncClusterRequest, runtime *dara.RuntimeOptions) (_result *SyncClusterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SyncCluster"),
		Version:     dara.String("2023-11-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SyncClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Synchronizes the data of hardware security modules (HSMs) in a cluster.
//
// Description:
//
//	  This operation is used for manually synchronizing data within clusters located in the Chinese Mainland. For clusters outside the Chinese Mainland, automatic data synchronization is supported, and this operation is unnecessary. If you attempt to use this operation, a 400 error code will be returned.
//
//		- The data synchronization takes approximately 5 minutes. To avoid service interruptions, we recommend performing this operation during off-peak hours.
//
// @param request - SyncClusterRequest
//
// @return SyncClusterResponse
func (client *Client) SyncCluster(request *SyncClusterRequest) (_result *SyncClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SyncClusterResponse{}
	_body, _err := client.SyncClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
