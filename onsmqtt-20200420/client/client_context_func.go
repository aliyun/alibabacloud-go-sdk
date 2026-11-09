// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// ApsaraMQ for MQTT supports X.509 certificate authentication, which allows clients to authenticate by using device certificates. A CA certificate is used to issue device certificates for clients and verify the validity of the device certificates. A CA certificate must be active to be used. By default, a CA certificate is active after registration. If a CA certificate is changed to the inactive state, you can invoke this operation to reactivate it.
//
// Description:
//
// - Only Platinum<props="china"> and Professional instances support the ActiveCaCertificate operation.
//
// - The request frequency limit for a single user is 500 requests per second. For special requirements, contact ApsaraMQ for MQTT technical support. DingTalk group ID: 35228338.
//
// - The ActiveCaCertificate operation only supports CA certificates that are registered on the ApsaraMQ for MQTT server. You can invoke the [ListCaCertificate](https://help.aliyun.com/document_detail/2604958.html) operation to query the registered CA certificates under a specified instance.
//
// @param request - ActiveCaCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActiveCaCertificateResponse
func (client *Client) ActiveCaCertificateWithContext(ctx context.Context, request *ActiveCaCertificateRequest, runtime *dara.RuntimeOptions) (_result *ActiveCaCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MqttInstanceId) {
		query["MqttInstanceId"] = request.MqttInstanceId
	}

	if !dara.IsNil(request.Sn) {
		query["Sn"] = request.Sn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ActiveCaCertificate"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ActiveCaCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Activates a device certificate. A device certificate is a digital certificate issued by a CA root certificate to a client device for security authentication when the client connects to the server. When a client device uses a device certificate for access authentication, the server verifies the device certificate based on the registered CA certificate. If the CA certificate and the device certificate match, the client authentication succeeds, and the system automatically registers the device certificate on the server. By default, a device certificate is in the activated state after it is registered on the server. If a device certificate that you want to use is deactivated, you can call this operation to reactivate it.
//
// Description:
//
// - Only Platinum<props="china"> and Professional edition instances support the ActiveDeviceCertificate operation.
//
// - The request frequency limit for a single user is 500 calls per second. If you have special requirements, contact the technical support team of MSMQ for MQTT. DingTalk group ID: 35228338.
//
// @param request - ActiveDeviceCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActiveDeviceCertificateResponse
func (client *Client) ActiveDeviceCertificateWithContext(ctx context.Context, request *ActiveDeviceCertificateRequest, runtime *dara.RuntimeOptions) (_result *ActiveDeviceCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CaSn) {
		query["CaSn"] = request.CaSn
	}

	if !dara.IsNil(request.DeviceSn) {
		query["DeviceSn"] = request.DeviceSn
	}

	if !dara.IsNil(request.MqttInstanceId) {
		query["MqttInstanceId"] = request.MqttInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ActiveDeviceCertificate"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ActiveDeviceCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a connection blacklist to disable connections from a specific device.
//
// @param request - AddCustomAuthConnectBlackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCustomAuthConnectBlackResponse
func (client *Client) AddCustomAuthConnectBlackWithContext(ctx context.Context, request *AddCustomAuthConnectBlackRequest, runtime *dara.RuntimeOptions) (_result *AddCustomAuthConnectBlackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		body["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddCustomAuthConnectBlack"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddCustomAuthConnectBlackResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds account authentication information, which can be specified at the client level.
//
// @param request - AddCustomAuthIdentityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCustomAuthIdentityResponse
func (client *Client) AddCustomAuthIdentityWithContext(ctx context.Context, request *AddCustomAuthIdentityRequest, runtime *dara.RuntimeOptions) (_result *AddCustomAuthIdentityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		body["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.IdentityType) {
		body["IdentityType"] = request.IdentityType
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Secret) {
		body["Secret"] = request.Secret
	}

	if !dara.IsNil(request.SignMode) {
		body["SignMode"] = request.SignMode
	}

	if !dara.IsNil(request.Username) {
		body["Username"] = request.Username
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddCustomAuthIdentity"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddCustomAuthIdentityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds authorization information for a topic resource. A first-level topic must be created in the console in advance.
//
// @param request - AddCustomAuthPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCustomAuthPermissionResponse
func (client *Client) AddCustomAuthPermissionWithContext(ctx context.Context, request *AddCustomAuthPermissionRequest, runtime *dara.RuntimeOptions) (_result *AddCustomAuthPermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Effect) {
		body["Effect"] = request.Effect
	}

	if !dara.IsNil(request.Identity) {
		body["Identity"] = request.Identity
	}

	if !dara.IsNil(request.IdentityType) {
		body["IdentityType"] = request.IdentityType
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PermitAction) {
		body["PermitAction"] = request.PermitAction
	}

	if !dara.IsNil(request.Topic) {
		body["Topic"] = request.Topic
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddCustomAuthPermission"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddCustomAuthPermissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// When token-based authentication is used, the ApsaraMQ for MQTT broker issues a token (temporary credential) to the client, and the client must include the token when accessing the ApsaraMQ for MQTT broker to authenticate permissions. This API operation is used to apply for a temporary token from ApsaraMQ for MQTT.
//
// Description:
//
// - The request frequency limit for a single user is 100 calls per second. If you have special requirements, contact the technical support team of ApsaraMQ for MQTT. DingTalk group ID: 35228338.
//
// - Each time you successfully call the **ApplyToken*	- API operation, it is counted as one message TPS, which affects your billing. For billing details, see [Billing overview](https://help.aliyun.com/document_detail/52819.html).
//
// @param request - ApplyTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyTokenResponse
func (client *Client) ApplyTokenWithContext(ctx context.Context, request *ApplyTokenRequest, runtime *dara.RuntimeOptions) (_result *ApplyTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Actions) {
		query["Actions"] = request.Actions
	}

	if !dara.IsNil(request.ExpireTime) {
		query["ExpireTime"] = request.ExpireTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Resources) {
		query["Resources"] = request.Resources
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyToken"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the online status of clients in batches by client ID.
//
// Description:
//
// - The QPS limit for the **BatchQuerySessionByClientIds*	- operation is 100. For more information, see [QPS limits](https://help.aliyun.com/document_detail/163047.html).
//
// - You can query the online status of up to 10 client IDs each time you call the **BatchQuerySessionByClientIds*	- operation.
//
// - Each successful call to the **BatchQuerySessionByClientIds*	- operation is counted as one message TPS, which affects your billing. For billing details, see [Billing](https://help.aliyun.com/document_detail/52819.html).
//
// - Considering the performance impact, the **BatchQuerySessionByClientIds*	- operation uses an internal cache for queries, which may result in data inconsistency. For exact queries, use the [QuerySessionByClientId](https://help.aliyun.com/document_detail/2603924.html) operation.
//
// @param request - BatchQuerySessionByClientIdsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchQuerySessionByClientIdsResponse
func (client *Client) BatchQuerySessionByClientIdsWithContext(ctx context.Context, request *BatchQuerySessionByClientIdsRequest, runtime *dara.RuntimeOptions) (_result *BatchQuerySessionByClientIdsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientIdList) {
		query["ClientIdList"] = request.ClientIdList
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchQuerySessionByClientIds"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchQuerySessionByClientIdsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Closes an active connection. However, the device may reconnect depending on the client reconnection mechanism.
//
// @param request - CloseConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseConnectionResponse
func (client *Client) CloseConnectionWithContext(ctx context.Context, request *CloseConnectionRequest, runtime *dara.RuntimeOptions) (_result *CloseConnectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		query["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloseConnection"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloseConnectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// When message producers and consumers connect to the ApsaraMQ for MQTT server to send and receive messages, they must specify a client ID as the unique identity of a client. A client ID consists of two parts in the <GroupID>@@@<DeviceID> format. The DeviceID is defined by the client, and the GroupID must be created in advance in the ApsaraMQ for MQTT server.
//
// Description:
//
// Each successful call to the **CreateGroupId*	- operation is counted as one message TPS, which affects your billing. For more information, see [Billing overview](https://help.aliyun.com/document_detail/52819.html).
//
// @param request - CreateGroupIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGroupIdResponse
func (client *Client) CreateGroupIdWithContext(ctx context.Context, request *CreateGroupIdRequest, runtime *dara.RuntimeOptions) (_result *CreateGroupIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGroupId"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGroupIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// ApsaraMQ for MQTT supports X.509 certificate authentication, allowing clients to use device certificates for access authentication. CA certificates are used to issue device certificates for clients and verify the validity of device certificates. Before you use a CA certificate, you must register it with the ApsaraMQ for MQTT server. If you no longer need a CA certificate, you can invoke the DeleteCaCertificate operation to delete it from the server records.
//
// Description:
//
// - Only Platinum<props="china"> and Professional instances support the DeleteCaCertificate operation.
//
// - The request frequency limit for a single user is 500 requests per second. For special requirements, contact ApsaraMQ for MQTT technical support. DingTalk group ID: 35228338.
//
// - The DeleteCaCertificate operation only supports CA certificates that are already registered with the ApsaraMQ for MQTT server. You can invoke the [ListCaCertificate](https://help.aliyun.com/document_detail/2604958.html) operation to query registered CA certificates under a specific instance.
//
// - When you delete a specified CA certificate, all device certificates issued by this CA certificate and registered with the ApsaraMQ for MQTT server are also automatically deleted.
//
// @param request - DeleteCaCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCaCertificateResponse
func (client *Client) DeleteCaCertificateWithContext(ctx context.Context, request *DeleteCaCertificateRequest, runtime *dara.RuntimeOptions) (_result *DeleteCaCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MqttInstanceId) {
		query["MqttInstanceId"] = request.MqttInstanceId
	}

	if !dara.IsNil(request.Sn) {
		query["Sn"] = request.Sn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCaCertificate"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCaCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a connection blacklist.
//
// @param request - DeleteCustomAuthConnectBlackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomAuthConnectBlackResponse
func (client *Client) DeleteCustomAuthConnectBlackWithContext(ctx context.Context, request *DeleteCustomAuthConnectBlackRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomAuthConnectBlackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		body["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomAuthConnectBlack"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomAuthConnectBlackResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes custom authentication.
//
// @param request - DeleteCustomAuthIdentityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomAuthIdentityResponse
func (client *Client) DeleteCustomAuthIdentityWithContext(ctx context.Context, request *DeleteCustomAuthIdentityRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomAuthIdentityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		body["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.IdentityType) {
		body["IdentityType"] = request.IdentityType
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Username) {
		body["Username"] = request.Username
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomAuthIdentity"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomAuthIdentityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes topic authorization.
//
// @param request - DeleteCustomAuthPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomAuthPermissionResponse
func (client *Client) DeleteCustomAuthPermissionWithContext(ctx context.Context, request *DeleteCustomAuthPermissionRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomAuthPermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Identity) {
		body["Identity"] = request.Identity
	}

	if !dara.IsNil(request.IdentityType) {
		body["IdentityType"] = request.IdentityType
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Topic) {
		body["Topic"] = request.Topic
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomAuthPermission"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomAuthPermissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the registration information of a specified device certificate from the server storage. A device certificate is a digital certificate issued by a certification authority (CA) root certificate to a client device. It is used for security authentication when the client connects to the server. After authentication, the server and the client can communicate securely by using the encryption key in the certificate. If the authentication fails, the server rejects the client connection.
//
// Description:
//
// - Only Platinum<props="china"> and Professional edition instances support the DeleteDeviceCertificate operation.
//
// - The request frequency limit for a single user is 500 calls per second. If you have special requirements, contact the technical support team of ApsaraMQ for MQTT. DingTalk group ID: 35228338.
//
// @param request - DeleteDeviceCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDeviceCertificateResponse
func (client *Client) DeleteDeviceCertificateWithContext(ctx context.Context, request *DeleteDeviceCertificateRequest, runtime *dara.RuntimeOptions) (_result *DeleteDeviceCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CaSn) {
		query["CaSn"] = request.CaSn
	}

	if !dara.IsNil(request.DeviceSn) {
		query["DeviceSn"] = request.DeviceSn
	}

	if !dara.IsNil(request.MqttInstanceId) {
		query["MqttInstanceId"] = request.MqttInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDeviceCertificate"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDeviceCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a group ID.
//
// Description:
//
// Each successful call to the **DeleteGroupId*	- operation is counted as one message TPS, which affects your billing. For more information about billing, see [Billing overview](https://help.aliyun.com/document_detail/52819.html).
//
// @param request - DeleteGroupIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGroupIdResponse
func (client *Client) DeleteGroupIdWithContext(ctx context.Context, request *DeleteGroupIdRequest, runtime *dara.RuntimeOptions) (_result *DeleteGroupIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGroupId"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGroupIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes wildcard subscriptions in batches.
//
// Description:
//
// > Each successful call to the **DeleteTopic*	- operation counts as one message TPS, which affects your billing. For more information, see [Billing overview](https://help.aliyun.com/document_detail/52819.html).
//
// @param request - DeleteWildcardByTopicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWildcardByTopicResponse
func (client *Client) DeleteWildcardByTopicWithContext(ctx context.Context, request *DeleteWildcardByTopicRequest, runtime *dara.RuntimeOptions) (_result *DeleteWildcardByTopicResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FirstTopic) {
		query["FirstTopic"] = request.FirstTopic
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Wildcards) {
		query["Wildcards"] = request.Wildcards
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWildcardByTopic"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWildcardByTopicResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Downgrades the public VIP access of a specific instance for multi-instance disaster recovery. This operation is supported only for Platinum instances and requires you to submit a ticket.
//
// @param request - DisasterDowngradeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisasterDowngradeResponse
func (client *Client) DisasterDowngradeWithContext(ctx context.Context, request *DisasterDowngradeRequest, runtime *dara.RuntimeOptions) (_result *DisasterDowngradeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DowngradeInstanceId) {
		body["DowngradeInstanceId"] = request.DowngradeInstanceId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisasterDowngrade"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisasterDowngradeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restores Internet access for an instance using its virtual IP address (VIP) for graceful recovery in a multi-instance disaster recovery architecture. This operation is available only for Platinum Edition instances. To use this operation, submit a ticket.
//
// @param request - DisasterRecoveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisasterRecoveryResponse
func (client *Client) DisasterRecoveryWithContext(ctx context.Context, request *DisasterRecoveryRequest, runtime *dara.RuntimeOptions) (_result *DisasterRecoveryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RecoveryInstanceId) {
		body["RecoveryInstanceId"] = request.RecoveryInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisasterRecovery"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisasterRecoveryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a certificate authority (CA) certificate, such as the content and status of the certificate. ApsaraMQ for MQTT allows you to use X.509 certificates for authentication. When you connect an ApsaraMQ for MQTT client to an ApsaraMQ for MQTT broker, you can use the device certificate to implement authentication. CA certificates are used to issue device certificates to clients and validate the device certificates.
//
// Description:
//
//	  Only ApsaraMQ for MQTT Enterprise Platinum Edition instances support this operation.
//
//		- You can call this operation up to 500 times per second per Alibaba Cloud account. If you want to increase the limit, join the DingTalk group (ID: 35228338) to contact ApsaraMQ for MQTT technical support.
//
// @param request - GetCaCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCaCertificateResponse
func (client *Client) GetCaCertificateWithContext(ctx context.Context, request *GetCaCertificateRequest, runtime *dara.RuntimeOptions) (_result *GetCaCertificateResponse, _err error) {
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
		Action:      dara.String("GetCaCertificate"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCaCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a device certificate. Device certificates are digital certificates issued to clients by certificate authority (CA) root certificates. When you connect an ApsaraMQ for MQTT client to an ApsaraMQ for MQTT broker, the broker uses the device certificate to authenticate the client. If the client passes the authentication, the client and the broker can communicate with each other based on the encrypted private key in the device certificate. If the client fails the authentication, access requests from the client are denied by the client.
//
// Description:
//
//	  Only ApsaraMQ for MQTT Enterprise Platinum Edition instances support this operation.
//
//		- You can call this operation up to 500 times per second per Alibaba Cloud account. If you want to increase the limit, join the DingTalk group (ID: 35228338) to contact ApsaraMQ for MQTT technical support.
//
// @param request - GetDeviceCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeviceCertificateResponse
func (client *Client) GetDeviceCertificateWithContext(ctx context.Context, request *GetDeviceCertificateRequest, runtime *dara.RuntimeOptions) (_result *GetDeviceCertificateResponse, _err error) {
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
		Action:      dara.String("GetDeviceCertificate"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeviceCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the access credential of a device. If unique-certificate-per-device authentication is used as the authentication method on an ApsaraMQ for MQTT broker, an access credential that you apply for in advance is required for authentication when you connect your device to the broker. The connection can be established only after the authentication is passed.
//
// Description:
//
// - You can call this operation up to 500 times per second per account. If the limit is exceeded, throttling is triggered. We recommend that you take note of this limit when you call this operation. For more information, see [Limits on QPS](https://help.aliyun.com/document_detail/163047.html).
//
// - Each successful call to the **GetDeviceCredential*	- operation increases the messaging transactions per second (TPS) by one. This affects the billing of your instance. For more information, see [Billing rules](https://help.aliyun.com/document_detail/52819.html).
//
// @param request - GetDeviceCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeviceCredentialResponse
func (client *Client) GetDeviceCredentialWithContext(ctx context.Context, request *GetDeviceCredentialRequest, runtime *dara.RuntimeOptions) (_result *GetDeviceCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		query["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeviceCredential"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeviceCredentialResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the registration code of a specific certificate authority (CA) certificate. When you register a CA certificate with an ApsaraMQ for MQTT broker, you must upload the validation certificate of the CA certificate to verify whether you have the private key of the CA certificate. The validation certificate of a CA certificate must be generated by using the registration code of the CA certificate.
//
// Description:
//
//	  Only ApsaraMQ for MQTT Enterprise Platinum Edition instances support this operation.
//
//		- You can call this operation up to 500 times per second per Alibaba Cloud account. If you want to increase the limit, join the DingTalk group (ID: 35228338) to contact ApsaraMQ for MQTT technical support.
//
// @param request - GetRegisterCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRegisterCodeResponse
func (client *Client) GetRegisterCodeWithContext(ctx context.Context, request *GetRegisterCodeRequest, runtime *dara.RuntimeOptions) (_result *GetRegisterCodeResponse, _err error) {
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
		Action:      dara.String("GetRegisterCode"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRegisterCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// ApsaraMQ for MQTT supports X.509 certificate-based device verification. Clients use device certificates for connection authentication. A Certificate Authority (CA) certificate is used to issue device certificates to clients and to verify the validity of those certificates. If you do not need to use a CA certificate for a period of time, call this operation to unregister it. This changes the status of the CA certificate to inactive. When you need to use the CA certificate again, you can call the ActiveCaCertificate operation to reactivate it.
//
// Description:
//
// - Only Platinum Edition<props="china"> and Professional Edition instances support the InactivateCaCertificate operation.
//
// - The request rate for a single user is limited to 500 per second. If you have special requirements, contact ApsaraMQ for MQTT technical support. The DingTalk group number is 35228338.
//
// - You can call this operation only for CA certificates that are registered with ApsaraMQ for MQTT. To query the registered CA certificates for an instance, call the [ListCaCertificate](https://help.aliyun.com/document_detail/2604958.html) operation.
//
// @param request - InactivateCaCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InactivateCaCertificateResponse
func (client *Client) InactivateCaCertificateWithContext(ctx context.Context, request *InactivateCaCertificateRequest, runtime *dara.RuntimeOptions) (_result *InactivateCaCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MqttInstanceId) {
		query["MqttInstanceId"] = request.MqttInstanceId
	}

	if !dara.IsNil(request.Sn) {
		query["Sn"] = request.Sn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InactivateCaCertificate"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InactivateCaCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// A device certificate is a digital certificate that a root Certificate Authority (CA) issues to a client device. The service uses this certificate to authenticate the client during connection. If authentication is successful, the service and the client can communicate securely using the encryption key in the certificate. If authentication fails, the service denies the client\\"s connection. Should you need to temporarily disable a device certificate, you can call this operation to inactivate it, setting its status to inactive.
//
// Description:
//
// - Only Platinum Edition<props="china">and Professional Edition instances support the InactivateDeviceCertificate operation.
//
// - The request frequency for a single user is limited to 500 calls per second. For special requirements, join the DingTalk group (ID: 35228338) to contact ApsaraMQ for MQTT technical support.
//
// @param request - InactivateDeviceCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InactivateDeviceCertificateResponse
func (client *Client) InactivateDeviceCertificateWithContext(ctx context.Context, request *InactivateDeviceCertificateRequest, runtime *dara.RuntimeOptions) (_result *InactivateDeviceCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CaSn) {
		query["CaSn"] = request.CaSn
	}

	if !dara.IsNil(request.DeviceSn) {
		query["DeviceSn"] = request.DeviceSn
	}

	if !dara.IsNil(request.MqttInstanceId) {
		query["MqttInstanceId"] = request.MqttInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InactivateDeviceCertificate"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InactivateDeviceCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all certificate authority (CA) certificates that are registered with an ApsaraMQ for MQTT instance. ApsaraMQ for MQTT allows you to use X.509 certificates for authentication. When you connect an ApsaraMQ for MQTT client to an ApsaraMQ for MQTT broker, you can use the device certificate to implement authentication. CA certificates are used to issue device certificates to clients and validate the device certificates.
//
// Description:
//
//	  Only ApsaraMQ for MQTT Enterprise Platinum Edition instances support this operation.
//
//		- You can call this operation up to 500 times per second per Alibaba Cloud account. If you want to increase the limit, join the DingTalk group (ID: 35228338) to contact ApsaraMQ for MQTT technical support.
//
// @param request - ListCaCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCaCertificateResponse
func (client *Client) ListCaCertificateWithContext(ctx context.Context, request *ListCaCertificateRequest, runtime *dara.RuntimeOptions) (_result *ListCaCertificateResponse, _err error) {
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
		Action:      dara.String("ListCaCertificate"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCaCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all device certificates that are registered with an ApsaraMQ for MQTT instance. Device certificates are digital certificates issued to clients by certificate authority (CA) root certificates. When you connect an ApsaraMQ for MQTT client to an ApsaraMQ for MQTT broker, the broker uses the device certificate to authenticate the client. If the client passes the authentication, the client and the broker can communicate with each other based on the encrypted private key in the device certificate. If the client fails the authentication, access requests from the client are denied by the client.
//
// Description:
//
//	  Only ApsaraMQ for MQTT Enterprise Platinum Edition instances support this operation.
//
//		- You can call this operation up to 500 times per second per Alibaba Cloud account. If you want to increase the limit, join the DingTalk group (ID: 35228338) to contact ApsaraMQ for MQTT technical support.
//
// @param request - ListDeviceCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeviceCertificateResponse
func (client *Client) ListDeviceCertificateWithContext(ctx context.Context, request *ListDeviceCertificateRequest, runtime *dara.RuntimeOptions) (_result *ListDeviceCertificateResponse, _err error) {
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
		Action:      dara.String("ListDeviceCertificate"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeviceCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all device certificates that are issued by a certificate authority (CA) certificate and registered with ApsaraMQ for MQTT brokers. Device certificates are digital certificates issued to clients by CA root certificates. When you connect an ApsaraMQ for MQTT client to an ApsaraMQ for MQTT broker, the broker uses the device certificate to authenticate the client. If the client passes the authentication, the client and the broker can communicate with each other based on the encrypted private key in the device certificate. If the client fails the authentication, access requests from the client are denied by the client.
//
// Description:
//
//	  Only ApsaraMQ for MQTT Enterprise Platinum Edition instances support this operation.
//
//		- You can call this operation up to 500 times per second per Alibaba Cloud account. If you want to increase the limit, join the DingTalk group (ID: 35228338) to contact ApsaraMQ for MQTT technical support.
//
// @param request - ListDeviceCertificateByCaSnRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeviceCertificateByCaSnResponse
func (client *Client) ListDeviceCertificateByCaSnWithContext(ctx context.Context, request *ListDeviceCertificateByCaSnRequest, runtime *dara.RuntimeOptions) (_result *ListDeviceCertificateByCaSnResponse, _err error) {
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
		Action:      dara.String("ListDeviceCertificateByCaSn"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeviceCertificateByCaSnResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries clients that have applied for access credentials in unique-certificate-per-device authentication mode in an ApsaraMQ for MQTT instance.
//
// @param request - ListDeviceCredentialClientIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeviceCredentialClientIdResponse
func (client *Client) ListDeviceCredentialClientIdWithContext(ctx context.Context, request *ListDeviceCredentialClientIdRequest, runtime *dara.RuntimeOptions) (_result *ListDeviceCredentialClientIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDeviceCredentialClientId"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeviceCredentialClientIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the ListGroupId operation to query all Group IDs in an instance.
//
// Description:
//
// Each successful call to the **ListGroupId*	- operation is counted as one message transaction and is billed based on transactions per second (TPS). For more information, see [billing details](https://help.aliyun.com/document_detail/52819.html).
//
// @param request - ListGroupIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupIdResponse
func (client *Client) ListGroupIdWithContext(ctx context.Context, request *ListGroupIdRequest, runtime *dara.RuntimeOptions) (_result *ListGroupIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGroupId"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGroupIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to query a list of instances.
//
// @param request - ListInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithContext(ctx context.Context, request *ListInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstances"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Lists the tags that are attached to one or more resources.
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
		Version:     dara.String("2020-04-20"),
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
// Queries a client ID in a connection blacklist.
//
// @param request - QueryCustomAuthConnectBlackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCustomAuthConnectBlackResponse
func (client *Client) QueryCustomAuthConnectBlackWithContext(ctx context.Context, request *QueryCustomAuthConnectBlackRequest, runtime *dara.RuntimeOptions) (_result *QueryCustomAuthConnectBlackResponse, _err error) {
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
		Action:      dara.String("QueryCustomAuthConnectBlack"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCustomAuthConnectBlackResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries custom identity authentication.
//
// @param request - QueryCustomAuthIdentityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCustomAuthIdentityResponse
func (client *Client) QueryCustomAuthIdentityWithContext(ctx context.Context, request *QueryCustomAuthIdentityRequest, runtime *dara.RuntimeOptions) (_result *QueryCustomAuthIdentityResponse, _err error) {
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
		Action:      dara.String("QueryCustomAuthIdentity"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCustomAuthIdentityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the authorization information about a topic.
//
// @param request - QueryCustomAuthPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCustomAuthPermissionResponse
func (client *Client) QueryCustomAuthPermissionWithContext(ctx context.Context, request *QueryCustomAuthPermissionRequest, runtime *dara.RuntimeOptions) (_result *QueryCustomAuthPermissionResponse, _err error) {
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
		Action:      dara.String("QueryCustomAuthPermission"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCustomAuthPermissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the device trace of an ApsaraMQ for MQTT client by paging. When a device encounters an exception, you can call this operation to query the connection history of the specified device to help you quickly identify the cause of the exception.
//
// Description:
//
// - Each successful call to the **QueryMqttTraceDevice*	- operation counts as one message TPS, which affects your billing. For more information about billing, see [Billing](https://help.aliyun.com/document_detail/52819.html).
//
// - The queries per second (QPS) limit for a single user on this operation is 500. Exceeding this limit will trigger throttling, which may affect your business. Call this operation appropriately. For more information, see [QPS limits](https://help.aliyun.com/document_detail/163047.html).
//
// @param request - QueryMqttTraceDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMqttTraceDeviceResponse
func (client *Client) QueryMqttTraceDeviceWithContext(ctx context.Context, request *QueryMqttTraceDeviceRequest, runtime *dara.RuntimeOptions) (_result *QueryMqttTraceDeviceResponse, _err error) {
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

	if !dara.IsNil(request.ClientId) {
		query["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MqttRegionId) {
		query["MqttRegionId"] = request.MqttRegionId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Reverse) {
		query["Reverse"] = request.Reverse
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMqttTraceDevice"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMqttTraceDeviceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries messages on a device within a specific period of time. If a message is not sent or received as expected, you can call this operation to query the messaging status of the message to efficiently troubleshoot issues.
//
// Description:
//
//	  Each successful call to the **QueryMqttTraceMessageOfClient*	- operation increases the messaging transactions per second (TPS) by one. This affects the billing of your instance. For more information, see [Billing rules](https://help.aliyun.com/document_detail/52819.html).
//
//		- You can call this operation up to 500 times per second per account. If the limit is exceeded, throttling is triggered. This may affect your business. We recommend that you take note of this limit when you call this operation. For more information, see [Limits on QPS](https://help.aliyun.com/document_detail/163047.html).
//
// @param request - QueryMqttTraceMessageOfClientRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMqttTraceMessageOfClientResponse
func (client *Client) QueryMqttTraceMessageOfClientWithContext(ctx context.Context, request *QueryMqttTraceMessageOfClientRequest, runtime *dara.RuntimeOptions) (_result *QueryMqttTraceMessageOfClientResponse, _err error) {
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

	if !dara.IsNil(request.ClientId) {
		query["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MqttRegionId) {
		query["MqttRegionId"] = request.MqttRegionId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Reverse) {
		query["Reverse"] = request.Reverse
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMqttTraceMessageOfClient"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMqttTraceMessageOfClientResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the trace of a message. If a message is not sent or received as expected, you can call this operation to view the message details to troubleshoot the issue. For example, you can query the time when the message is published and the client that publishes the message.
//
// Description:
//
//	  Each successful call to the **QueryMqttTraceMessagePublish*	- operation increases the messaging transactions per second (TPS). This affects the billing of your instance. For more information, see [Billing rules](https://help.aliyun.com/document_detail/52819.html).
//
//		- You can call this operation up to 500 times per second per account. If the limit is exceeded, throttling is triggered. This may affect your business. We recommend that you take note of this limit when you call this operation. For more information, see [Limits on QPS](https://help.aliyun.com/document_detail/163047.html).
//
// @param request - QueryMqttTraceMessagePublishRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMqttTraceMessagePublishResponse
func (client *Client) QueryMqttTraceMessagePublishWithContext(ctx context.Context, request *QueryMqttTraceMessagePublishRequest, runtime *dara.RuntimeOptions) (_result *QueryMqttTraceMessagePublishResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MqttRegionId) {
		query["MqttRegionId"] = request.MqttRegionId
	}

	if !dara.IsNil(request.MsgId) {
		query["MsgId"] = request.MsgId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMqttTraceMessagePublish"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMqttTraceMessagePublishResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the delivery trace of a message. If a message is not sent or received as expected, you can call this operation to view the details about the message. For example, you can query the clients that subscribe to the message and the time when the message is delivered. This operation helps you locate the problem and identify the cause of the problem.
//
// Description:
//
// - Each successful call to the **QueryMqttTraceMessageSubscribe*	- operation increases the messaging transactions per second (TPS) by one. This affects the billing of your instance. For more information, see [Billing rules](https://help.aliyun.com/document_detail/52819.html).
//
// - You can call this operation up to 500 times per second per account. If the limit is exceeded, throttling is triggered. This may affect your business. We recommend that you take note of this limit when you call this operation. For more information, see [Limits on QPS](https://help.aliyun.com/document_detail/163047.html).
//
// @param request - QueryMqttTraceMessageSubscribeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMqttTraceMessageSubscribeResponse
func (client *Client) QueryMqttTraceMessageSubscribeWithContext(ctx context.Context, request *QueryMqttTraceMessageSubscribeRequest, runtime *dara.RuntimeOptions) (_result *QueryMqttTraceMessageSubscribeResponse, _err error) {
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

	if !dara.IsNil(request.ClientId) {
		query["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MqttRegionId) {
		query["MqttRegionId"] = request.MqttRegionId
	}

	if !dara.IsNil(request.MsgId) {
		query["MsgId"] = request.MsgId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Reverse) {
		query["Reverse"] = request.Reverse
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMqttTraceMessageSubscribe"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMqttTraceMessageSubscribeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the running status of an ApsaraMQ for MQTT client. You can troubleshoot issues based on the queried results. You can enter the ID of an ApsaraMQ for MQTT client to check the connection status and IP address of the device.
//
// Description:
//
// - You can call this operation up to 500 times per second.\\*\\*\\*\\	- For more information, see [Limits on QPS](https://help.aliyun.com/document_detail/163047.html).
//
// - Each successful call to the **QuerySessionByClientId*	- operation increases the messaging transactions per second (TPS) by one. This affects the billing of your instance. For more information, see [Billing rules](https://help.aliyun.com/document_detail/52819.html).
//
// @param request - QuerySessionByClientIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySessionByClientIdResponse
func (client *Client) QuerySessionByClientIdWithContext(ctx context.Context, request *QuerySessionByClientIdRequest, runtime *dara.RuntimeOptions) (_result *QuerySessionByClientIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		query["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySessionByClientId"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySessionByClientIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of a token. If token-based authentication is used for permission authentication on an ApsaraMQ for MQTT broker, a token that is issued by the broker is required for authentication each time a client is connected to the broker. A token is a temporary credential and is valid only within a specific period of time. You can call this operation to query whether a token expires.
//
// Description:
//
//	  You can call this operation up to 100 times per second per account. If you want to increase the limit, join the DingTalk group 35228338 to contact ApsaraMQ for MQTT technical support.
//
//		- Each successful call to the **QueryToken*	- operation increases the messaging transactions per second (TPS) by one. This affects the billing of your instance. For more information, see [Billing rules](https://help.aliyun.com/document_detail/52819.html).
//
// @param request - QueryTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTokenResponse
func (client *Client) QueryTokenWithContext(ctx context.Context, request *QueryTokenRequest, runtime *dara.RuntimeOptions) (_result *QueryTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryToken"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all wildcards and the upper limit under a specific topic.
//
// Description:
//
// > Each successful call to the **DeleteTopic*	- operation counts as one message TPS, which affects your billing. For more information, see [Billing overview](https://help.aliyun.com/document_detail/52819.html).
//
// @param request - QueryWildcardByTopicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryWildcardByTopicResponse
func (client *Client) QueryWildcardByTopicWithContext(ctx context.Context, request *QueryWildcardByTopicRequest, runtime *dara.RuntimeOptions) (_result *QueryWildcardByTopicResponse, _err error) {
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
		Action:      dara.String("QueryWildcardByTopic"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryWildcardByTopicResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the access credential of a device.
//
// Description:
//
// ## [](#)Limits
//
// You can call this operation up to 500 times per second per account. If the limit is exceeded, throttling is triggered. This may affect your business. We recommend that you take note of this limit when you call this operation. For more information, see [Limits on QPS](https://help.aliyun.com/document_detail/163047.html).
//
// > Each successful call to the **RefreshDeviceCredential*	- operation increases the messaging transactions per second (TPS) by one. This affects the billing of your instance. For more information, see [Billing rules](https://help.aliyun.com/document_detail/52819.html).
//
// @param request - RefreshDeviceCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshDeviceCredentialResponse
func (client *Client) RefreshDeviceCredentialWithContext(ctx context.Context, request *RefreshDeviceCredentialRequest, runtime *dara.RuntimeOptions) (_result *RefreshDeviceCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		query["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RefreshDeviceCredential"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RefreshDeviceCredentialResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # RegisterCaCertificate
//
// Description:
//
// - Only ApsaraMQ for MQTT Enterprise Platinum Edition instances support this operation.
//
// - You can call this operation up to 500 times per second per Alibaba Cloud account. If you want to increase the limit, join the DingTalk group (ID: 35228338) to contact ApsaraMQ for MQTT technical support.
//
// @param request - RegisterCaCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterCaCertificateResponse
func (client *Client) RegisterCaCertificateWithContext(ctx context.Context, request *RegisterCaCertificateRequest, runtime *dara.RuntimeOptions) (_result *RegisterCaCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CaContent) {
		query["CaContent"] = request.CaContent
	}

	if !dara.IsNil(request.CaName) {
		query["CaName"] = request.CaName
	}

	if !dara.IsNil(request.MqttInstanceId) {
		query["MqttInstanceId"] = request.MqttInstanceId
	}

	if !dara.IsNil(request.VerificationContent) {
		query["VerificationContent"] = request.VerificationContent
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RegisterCaCertificate"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RegisterCaCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Registers an access credential for a device. In unique-certificate-per-device authentication mode, an application server applies a unique access credential for each device from the corresponding ApsaraMQ for MQTT broker. The access credential of a device consists of the client ID, AccessKey ID, and AccessKey secret of the device. When you connect a device to ApsaraMQ for MQTT, you must configure Username and Password based on the access credential of the device for authentication. You can activate the device and transfer data between the device and ApsaraMQ for MQTT only after the authentication is passed.
//
// Description:
//
// - You can call this operation up to 500 times per second per account. If the limit is exceeded, throttling is triggered. This may affect your business. We recommend that you take note of this limit when you call this operation. For more information, see [Limits on QPS](https://help.aliyun.com/document_detail/163047.html).
//
// - Each successful call to the **RegisterDeviceCredential*	- operation increases the messaging transactions per second (TPS) by one. This affects the billing of your instance. For more information, see [Billing rules](https://help.aliyun.com/document_detail/52819.html).
//
// @param request - RegisterDeviceCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterDeviceCredentialResponse
func (client *Client) RegisterDeviceCredentialWithContext(ctx context.Context, request *RegisterDeviceCredentialRequest, runtime *dara.RuntimeOptions) (_result *RegisterDeviceCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		query["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RegisterDeviceCredential"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RegisterDeviceCredentialResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes a token.
//
// Description:
//
//	  You can call this operation up to 5 times per second per account. If you want to increase the limit, join the DingTalk group 35228338 to contact ApsaraMQ for MQTT technical support.
//
//		- Each successful call to the **RevokeToken*	- operation increases the messaging transactions per second (TPS). This affects the billing of your instance. For more information, see [Billing rules](https://help.aliyun.com/document_detail/52819.html).
//
// @param request - RevokeTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeTokenResponse
func (client *Client) RevokeTokenWithContext(ctx context.Context, request *RevokeTokenRequest, runtime *dara.RuntimeOptions) (_result *RevokeTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeToken"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a single message from an application on a cloud server to ApsaraMQ for MQTT.
//
// Description:
//
//	  The **SendMessage*	- operation is called by an application on a cloud server. This operation is complementary to the operation that is called to send a message from an ApsaraMQ for MQTT client. For information about the differences between the scenarios of sending messages from applications on cloud servers and the scenarios of sending messages from ApsaraMQ for MQTT clients, see [Developer guide](https://help.aliyun.com/document_detail/179160.html).
//
//		- Messages that are sent by calling the **SendMessage*	- operation cannot be forwarded to ApsaraMQ for RocketMQ. If you want to use an ApsaraMQ for MQTT broker to forward messages to ApsaraMQ for RocketMQ, use [an SDK to send the messages](https://help.aliyun.com/document_detail/174527.html). The **SendMessage*	- operation supports up to 1,000 queries per second (QPS). For more information, see [QPS limits](https://help.aliyun.com/document_detail/163047.html).
//
//		- Each successful call to the **SendMessage*	- operation is calculated as a message transaction per second (TPS). This affects the billing of your instance. For more information, see [Billing rules](https://help.aliyun.com/document_detail/52819.html).
//
// @param request - SendMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendMessageResponse
func (client *Client) SendMessageWithContext(ctx context.Context, request *SendMessageRequest, runtime *dara.RuntimeOptions) (_result *SendMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MqttTopic) {
		query["MqttTopic"] = request.MqttTopic
	}

	if !dara.IsNil(request.Payload) {
		query["Payload"] = request.Payload
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendMessage"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendMessageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a multi-domain certificate.
//
// @param request - SetSniConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetSniConfigResponse
func (client *Client) SetSniConfigWithContext(ctx context.Context, request *SetSniConfigRequest, runtime *dara.RuntimeOptions) (_result *SetSniConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DefaultCertificate) {
		query["DefaultCertificate"] = request.DefaultCertificate
	}

	if !dara.IsNil(request.MqttInstanceId) {
		query["MqttInstanceId"] = request.MqttInstanceId
	}

	if !dara.IsNil(request.SniConfig) {
		query["SniConfig"] = request.SniConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetSniConfig"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetSniConfigResponse{}
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
		Version:     dara.String("2020-04-20"),
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
// Deregisters the access credential of a device. After the access credential of a device is deregistered, you can no longer use the access credential to authenticate the device on the ApsaraMQ for MQTT broker.
//
// Description:
//
// - You can call this operation up to 500 times per second per account. If the limit is exceeded, throttling is triggered. This may affect your business. We recommend that you take note of this limit when you call this operation. For more information, see [Limits on QPS](https://help.aliyun.com/document_detail/163047.html).
//
// - Each successful call to the **UnRegisterDeviceCredential*	- operation increases the number of transactions per second (TPS) by one. This affects the billing of your instance. For more information, see [Billing rules](https://help.aliyun.com/document_detail/52819.html).
//
// @param request - UnRegisterDeviceCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnRegisterDeviceCredentialResponse
func (client *Client) UnRegisterDeviceCredentialWithContext(ctx context.Context, request *UnRegisterDeviceCredentialRequest, runtime *dara.RuntimeOptions) (_result *UnRegisterDeviceCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		query["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnRegisterDeviceCredential"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnRegisterDeviceCredentialResponse{}
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
		Version:     dara.String("2020-04-20"),
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
// Updates custom authentication.
//
// @param request - UpdateCustomAuthIdentityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCustomAuthIdentityResponse
func (client *Client) UpdateCustomAuthIdentityWithContext(ctx context.Context, request *UpdateCustomAuthIdentityRequest, runtime *dara.RuntimeOptions) (_result *UpdateCustomAuthIdentityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		body["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.IdentityType) {
		body["IdentityType"] = request.IdentityType
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Secret) {
		body["Secret"] = request.Secret
	}

	if !dara.IsNil(request.SignMode) {
		body["SignMode"] = request.SignMode
	}

	if !dara.IsNil(request.Username) {
		body["Username"] = request.Username
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCustomAuthIdentity"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCustomAuthIdentityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the permissions on a topic.
//
// @param request - UpdateCustomAuthPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCustomAuthPermissionResponse
func (client *Client) UpdateCustomAuthPermissionWithContext(ctx context.Context, request *UpdateCustomAuthPermissionRequest, runtime *dara.RuntimeOptions) (_result *UpdateCustomAuthPermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Effect) {
		body["Effect"] = request.Effect
	}

	if !dara.IsNil(request.Identity) {
		body["Identity"] = request.Identity
	}

	if !dara.IsNil(request.IdentityType) {
		body["IdentityType"] = request.IdentityType
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PermitAction) {
		body["PermitAction"] = request.PermitAction
	}

	if !dara.IsNil(request.Topic) {
		body["Topic"] = request.Topic
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCustomAuthPermission"),
		Version:     dara.String("2020-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCustomAuthPermissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
