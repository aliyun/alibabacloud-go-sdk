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
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("winnexo"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) _postOSSObject(bucketName *string, form map[string]interface{}, runtime *dara.RuntimeOptions) (_result map[string]interface{}, _err error) {
	_runtime := dara.NewRuntimeObject(map[string]interface{}{
		"key":            dara.ToString(dara.Default(dara.StringValue(runtime.Key), dara.StringValue(client.Key))),
		"cert":           dara.ToString(dara.Default(dara.StringValue(runtime.Cert), dara.StringValue(client.Cert))),
		"ca":             dara.ToString(dara.Default(dara.StringValue(runtime.Ca), dara.StringValue(client.Ca))),
		"readTimeout":    dara.ForceInt(dara.Default(dara.IntValue(runtime.ReadTimeout), dara.IntValue(client.ReadTimeout))),
		"connectTimeout": dara.ForceInt(dara.Default(dara.IntValue(runtime.ConnectTimeout), dara.IntValue(client.ConnectTimeout))),
		"httpProxy":      dara.ToString(dara.Default(dara.StringValue(runtime.HttpProxy), dara.StringValue(client.HttpProxy))),
		"httpsProxy":     dara.ToString(dara.Default(dara.StringValue(runtime.HttpsProxy), dara.StringValue(client.HttpsProxy))),
		"noProxy":        dara.ToString(dara.Default(dara.StringValue(runtime.NoProxy), dara.StringValue(client.NoProxy))),
		"socks5Proxy":    dara.ToString(dara.Default(dara.StringValue(runtime.Socks5Proxy), dara.StringValue(client.Socks5Proxy))),
		"socks5NetWork":  dara.ToString(dara.Default(dara.StringValue(runtime.Socks5NetWork), dara.StringValue(client.Socks5NetWork))),
		"maxIdleConns":   dara.ForceInt(dara.Default(dara.IntValue(runtime.MaxIdleConns), dara.IntValue(client.MaxIdleConns))),
		"retryOptions":   client.RetryOptions,
		"ignoreSSL":      dara.ForceBoolean(dara.Default(dara.BoolValue(runtime.IgnoreSSL), false)),
		"tlsMinVersion":  dara.StringValue(client.TlsMinVersion),
	})

	var retryPolicyContext *dara.RetryPolicyContext
	var request_ *dara.Request
	var response_ *dara.Response
	var _resultErr error
	retriesAttempted := int(0)
	retryPolicyContext = &dara.RetryPolicyContext{
		RetriesAttempted: retriesAttempted,
	}

	_result = make(map[string]interface{})
	for dara.ShouldRetry(_runtime.RetryOptions, retryPolicyContext) {
		_resultErr = nil
		_backoffDelayTime := dara.GetBackoffDelay(_runtime.RetryOptions, retryPolicyContext)
		dara.Sleep(_backoffDelayTime)

		request_ = dara.NewRequest()
		boundary := dara.GetBoundary()
		tmp := dara.ToString(form["host"])
		host := dara.StringValue(bucketName) + "." + tmp
		request_.Protocol = dara.String("HTTPS")
		request_.Method = dara.String("POST")
		request_.Pathname = dara.String("/")
		request_.Headers = map[string]*string{
			"host":       dara.String(host),
			"date":       openapiutil.GetDateUTCString(),
			"user-agent": openapiutil.GetUserAgent(dara.String("")),
		}
		request_.Headers["content-type"] = dara.String("multipart/form-data; boundary=" + boundary)
		request_.Body = dara.ToFileForm(form, boundary)
		response_, _err = dara.DoRequest(request_, _runtime)
		if _err != nil {
			retriesAttempted++
			retryPolicyContext = &dara.RetryPolicyContext{
				RetriesAttempted: retriesAttempted,
				HttpRequest:      request_,
				HttpResponse:     response_,
				Exception:        _err,
			}
			_resultErr = _err
			continue
		}

		_result, _err = _postOSSObject_opResponse(response_)
		if _err != nil {
			retriesAttempted++
			retryPolicyContext = &dara.RetryPolicyContext{
				RetriesAttempted: retriesAttempted,
				HttpRequest:      request_,
				HttpResponse:     response_,
				Exception:        _err,
			}
			_resultErr = _err
			continue
		}

		return _result, _err
	}
	if dara.BoolValue(client.DisableSDKError) != true {
		_resultErr = dara.TeaSDKError(_resultErr)
	}
	return _result, _resultErr
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
// Adds multiple tenant members to a specified user group in a single request.
//
// Description:
//
// ## Request description
//
// - This operation supports batch addition of members by providing a user group ID and one or more user IDs.
//
// - Duplicate entries in the user ID list do not cause errors. The system automatically handles duplicates to ensure each user is added only once.
//
// - The caller must have the required permissions to perform this operation.
//
// - This operation is applicable to scenarios that require quick team structure management or access control policy adjustments.
//
// @param tmpReq - AddUserGroupMembersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserGroupMembersResponse
func (client *Client) AddUserGroupMembersWithOptions(tmpReq *AddUserGroupMembersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddUserGroupMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddUserGroupMembersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserIds) {
		request.UserIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserIds, dara.String("userIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UserGroupId) {
		body["userGroupId"] = request.UserGroupId
	}

	if !dara.IsNil(request.UserIdsShrink) {
		body["userIds"] = request.UserIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddUserGroupMembers"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/addUserGroupMembers"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddUserGroupMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds multiple tenant members to a specified user group in a single request.
//
// Description:
//
// ## Request description
//
// - This operation supports batch addition of members by providing a user group ID and one or more user IDs.
//
// - Duplicate entries in the user ID list do not cause errors. The system automatically handles duplicates to ensure each user is added only once.
//
// - The caller must have the required permissions to perform this operation.
//
// - This operation is applicable to scenarios that require quick team structure management or access control policy adjustments.
//
// @param request - AddUserGroupMembersRequest
//
// @return AddUserGroupMembersResponse
func (client *Client) AddUserGroupMembers(request *AddUserGroupMembersRequest) (_result *AddUserGroupMembersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddUserGroupMembersResponse{}
	_body, _err := client.AddUserGroupMembersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Batch cancels digital employee favorites for specific object types.
//
// Description:
//
// Idempotently cancels favorites across three independent dimensions: graphName, operatingObjectName, and objectType. The input array accepts 1 to 200 items per request. Each item must be a non-empty string with a maximum length of 128 characters. The server validates and deduplicates items while preserving order. Non-string values, values that exceed the length limit, or arrays that exceed the size limit are rejected. Deletion, per-item status updates, and remaining valid count are completed within a single transaction. To safely cancel all favorites, you must also call ClearOperatingObjectFavorites to clean up historical records, MISSING records, or permission-hidden records that are not visible in the list. Then read back the result to confirm that total is 0.
//
// @param tmpReq - BatchRemoveOperatingObjectFavoritesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchRemoveOperatingObjectFavoritesResponse
func (client *Client) BatchRemoveOperatingObjectFavoritesWithOptions(tmpReq *BatchRemoveOperatingObjectFavoritesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchRemoveOperatingObjectFavoritesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchRemoveOperatingObjectFavoritesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ObjectIds) {
		request.ObjectIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ObjectIds, dara.String("objectIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.GraphName) {
		body["graphName"] = request.GraphName
	}

	if !dara.IsNil(request.ObjectIdsShrink) {
		body["objectIds"] = request.ObjectIdsShrink
	}

	if !dara.IsNil(request.ObjectType) {
		body["objectType"] = request.ObjectType
	}

	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchRemoveOperatingObjectFavorites"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/batchRemoveOperatingObjectFavorites"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchRemoveOperatingObjectFavoritesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Batch cancels digital employee favorites for specific object types.
//
// Description:
//
// Idempotently cancels favorites across three independent dimensions: graphName, operatingObjectName, and objectType. The input array accepts 1 to 200 items per request. Each item must be a non-empty string with a maximum length of 128 characters. The server validates and deduplicates items while preserving order. Non-string values, values that exceed the length limit, or arrays that exceed the size limit are rejected. Deletion, per-item status updates, and remaining valid count are completed within a single transaction. To safely cancel all favorites, you must also call ClearOperatingObjectFavorites to clean up historical records, MISSING records, or permission-hidden records that are not visible in the list. Then read back the result to confirm that total is 0.
//
// @param request - BatchRemoveOperatingObjectFavoritesRequest
//
// @return BatchRemoveOperatingObjectFavoritesResponse
func (client *Client) BatchRemoveOperatingObjectFavorites(request *BatchRemoveOperatingObjectFavoritesRequest) (_result *BatchRemoveOperatingObjectFavoritesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchRemoveOperatingObjectFavoritesResponse{}
	_body, _err := client.BatchRemoveOperatingObjectFavoritesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a service health check.
//
// @param request - CheckHealthRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckHealthResponse
func (client *Client) CheckHealthWithOptions(request *CheckHealthRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CheckHealthResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckHealth"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/checkHealth"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckHealthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a service health check.
//
// @param request - CheckHealthRequest
//
// @return CheckHealthResponse
func (client *Client) CheckHealth(request *CheckHealthRequest) (_result *CheckHealthResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckHealthResponse{}
	_body, _err := client.CheckHealthWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Clears all follows of a specific object type for a digital employee.
//
// Description:
//
// Clears all persisted follows for the current calling user across three independent dimensions: graphName, operatingObjectName, and objectType. This includes historical records, MISSING records, and permission-hidden records that are not visible in the list. The operation does not return invisible object IDs and verifies that the remaining physical record count is zero within the same transaction.
//
// @param request - ClearOperatingObjectFavoritesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClearOperatingObjectFavoritesResponse
func (client *Client) ClearOperatingObjectFavoritesWithOptions(request *ClearOperatingObjectFavoritesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ClearOperatingObjectFavoritesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.GraphName) {
		body["graphName"] = request.GraphName
	}

	if !dara.IsNil(request.ObjectType) {
		body["objectType"] = request.ObjectType
	}

	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClearOperatingObjectFavorites"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/clearOperatingObjectFavorites"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClearOperatingObjectFavoritesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Clears all follows of a specific object type for a digital employee.
//
// Description:
//
// Clears all persisted follows for the current calling user across three independent dimensions: graphName, operatingObjectName, and objectType. This includes historical records, MISSING records, and permission-hidden records that are not visible in the list. The operation does not return invisible object IDs and verifies that the remaining physical record count is zero within the same transaction.
//
// @param request - ClearOperatingObjectFavoritesRequest
//
// @return ClearOperatingObjectFavoritesResponse
func (client *Client) ClearOperatingObjectFavorites(request *ClearOperatingObjectFavoritesRequest) (_result *ClearOperatingObjectFavoritesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ClearOperatingObjectFavoritesResponse{}
	_body, _err := client.ClearOperatingObjectFavoritesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a service notice.
//
// Description:
//
// ## Request description
//
// Creates a service notice. The caller must be mapped to a real platform user in the system O&M tenant and must have announcement management permissions.
//
// - `priority`: The importance level of the notice. Valid values: URGENT, IMPORTANT, and GENERAL.
//
// - `targetTenantIds` / `targetRoleCodes`: Used only when the corresponding target mode is set to SPECIFIED. Pass values as a JSON array.
//
// - `effectiveStart` / `effectiveEnd`: ISO 8601 time with time zone.
//
// - `publishNow`: If set to true, the notice is published immediately after creation. Otherwise, it is saved as a draft.
//
// @param tmpReq - CreateAnnouncementRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAnnouncementResponse
func (client *Client) CreateAnnouncementWithOptions(tmpReq *CreateAnnouncementRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAnnouncementResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAnnouncementShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TargetRoleCodes) {
		request.TargetRoleCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TargetRoleCodes, dara.String("targetRoleCodes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TargetTenantIds) {
		request.TargetTenantIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TargetTenantIds, dara.String("targetTenantIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	if !dara.IsNil(request.DisplayPage) {
		body["displayPage"] = request.DisplayPage
	}

	if !dara.IsNil(request.DisplayType) {
		body["displayType"] = request.DisplayType
	}

	if !dara.IsNil(request.EffectiveEnd) {
		body["effectiveEnd"] = request.EffectiveEnd
	}

	if !dara.IsNil(request.EffectiveStart) {
		body["effectiveStart"] = request.EffectiveStart
	}

	if !dara.IsNil(request.Priority) {
		body["priority"] = request.Priority
	}

	if !dara.IsNil(request.PublishNow) {
		body["publishNow"] = request.PublishNow
	}

	if !dara.IsNil(request.TargetRoleCodesShrink) {
		body["targetRoleCodes"] = request.TargetRoleCodesShrink
	}

	if !dara.IsNil(request.TargetRoleMode) {
		body["targetRoleMode"] = request.TargetRoleMode
	}

	if !dara.IsNil(request.TargetTenantIdsShrink) {
		body["targetTenantIds"] = request.TargetTenantIdsShrink
	}

	if !dara.IsNil(request.TargetTenantMode) {
		body["targetTenantMode"] = request.TargetTenantMode
	}

	if !dara.IsNil(request.Title) {
		body["title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAnnouncement"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/createAnnouncement"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAnnouncementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a service notice.
//
// Description:
//
// ## Request description
//
// Creates a service notice. The caller must be mapped to a real platform user in the system O&M tenant and must have announcement management permissions.
//
// - `priority`: The importance level of the notice. Valid values: URGENT, IMPORTANT, and GENERAL.
//
// - `targetTenantIds` / `targetRoleCodes`: Used only when the corresponding target mode is set to SPECIFIED. Pass values as a JSON array.
//
// - `effectiveStart` / `effectiveEnd`: ISO 8601 time with time zone.
//
// - `publishNow`: If set to true, the notice is published immediately after creation. Otherwise, it is saved as a draft.
//
// @param request - CreateAnnouncementRequest
//
// @return CreateAnnouncementResponse
func (client *Client) CreateAnnouncement(request *CreateAnnouncementRequest) (_result *CreateAnnouncementResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAnnouncementResponse{}
	_body, _err := client.CreateAnnouncementWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a session.
//
// @param tmpReq - CreateConversationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConversationResponse
func (client *Client) CreateConversationWithOptions(tmpReq *CreateConversationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateConversationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateConversationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OperatingObjectName) {
		request.OperatingObjectNameShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OperatingObjectName, dara.String("operatingObjectName"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Metadata) {
		body["metadata"] = request.Metadata
	}

	if !dara.IsNil(request.ObjectId) {
		body["objectId"] = request.ObjectId
	}

	if !dara.IsNil(request.OperatingObjectNameShrink) {
		body["operatingObjectName"] = request.OperatingObjectNameShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateConversation"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/createConversation"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateConversationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a session.
//
// @param request - CreateConversationRequest
//
// @return CreateConversationResponse
func (client *Client) CreateConversation(request *CreateConversationRequest) (_result *CreateConversationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateConversationResponse{}
	_body, _err := client.CreateConversationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Registers a custom-only organization.
//
// Description:
//
// Registers a custom-only organization for subsequent department tree push through syncOrgStructure.
//
//	Registration logic:
//
//	1. Validates the corpId format (must start with a lowercase letter or digit, 3-64 characters, hyphens allowed).
//
//	2. Delegates to OrgSyncAuthorizedService to execute registration (includes permission verification and tenant-level uniqueness check).
//
//	3. Returns the registration result.
//
//	Note: Custom-only organizations support only department tree synchronization. Member relationship synchronization is not supported.
//
// @param request - CreateCustomOrgRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomOrgResponse
func (client *Client) CreateCustomOrgWithOptions(request *CreateCustomOrgRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateCustomOrgResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CorpId) {
		body["corpId"] = request.CorpId
	}

	if !dara.IsNil(request.CorpName) {
		body["corpName"] = request.CorpName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCustomOrg"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/createCustomOrg"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCustomOrgResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Registers a custom-only organization.
//
// Description:
//
// Registers a custom-only organization for subsequent department tree push through syncOrgStructure.
//
//	Registration logic:
//
//	1. Validates the corpId format (must start with a lowercase letter or digit, 3-64 characters, hyphens allowed).
//
//	2. Delegates to OrgSyncAuthorizedService to execute registration (includes permission verification and tenant-level uniqueness check).
//
//	3. Returns the registration result.
//
//	Note: Custom-only organizations support only department tree synchronization. Member relationship synchronization is not supported.
//
// @param request - CreateCustomOrgRequest
//
// @return CreateCustomOrgResponse
func (client *Client) CreateCustomOrg(request *CreateCustomOrgRequest) (_result *CreateCustomOrgResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateCustomOrgResponse{}
	_body, _err := client.CreateCustomOrgWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a DingTalk group chat knowledge resource in a group.
//
// Description:
//
// ## Operation description
//
// - Connects a specified DingTalk group chat to a group knowledge base that the caller has joined.
//
// - The resource type is fixed to ALI_DING, the scope is fixed to GROUP, and the owning user is resolved from the gateway authentication identity.
//
// - groupId, chatId, and historyStartTime are required.
//
// - updateFrequency can be configured by using a preset or a five-field cron expression for subsequent synchronization frequency.
//
// - The server verifies the caller\\"s group membership, the target group directory permissions, and the uniqueness of chatId within the scope.
//
// @param tmpReq - CreateGroupAliDingChatRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGroupAliDingChatResponse
func (client *Client) CreateGroupAliDingChatWithOptions(tmpReq *CreateGroupAliDingChatRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateGroupAliDingChatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateGroupAliDingChatShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateFrequency) {
		request.UpdateFrequencyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateFrequency, dara.String("updateFrequency"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ChatId) {
		body["chatId"] = request.ChatId
	}

	if !dara.IsNil(request.ChatName) {
		body["chatName"] = request.ChatName
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DirectoryId) {
		body["directoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.GroupId) {
		body["groupId"] = request.GroupId
	}

	if !dara.IsNil(request.HistoryStartTime) {
		body["historyStartTime"] = request.HistoryStartTime
	}

	if !dara.IsNil(request.Notes) {
		body["notes"] = request.Notes
	}

	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	if !dara.IsNil(request.SourceTags) {
		body["sourceTags"] = request.SourceTags
	}

	if !dara.IsNil(request.UpdateFrequencyShrink) {
		body["updateFrequency"] = request.UpdateFrequencyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGroupAliDingChat"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/createGroupAliDingChat"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGroupAliDingChatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a DingTalk group chat knowledge resource in a group.
//
// Description:
//
// ## Operation description
//
// - Connects a specified DingTalk group chat to a group knowledge base that the caller has joined.
//
// - The resource type is fixed to ALI_DING, the scope is fixed to GROUP, and the owning user is resolved from the gateway authentication identity.
//
// - groupId, chatId, and historyStartTime are required.
//
// - updateFrequency can be configured by using a preset or a five-field cron expression for subsequent synchronization frequency.
//
// - The server verifies the caller\\"s group membership, the target group directory permissions, and the uniqueness of chatId within the scope.
//
// @param request - CreateGroupAliDingChatRequest
//
// @return CreateGroupAliDingChatResponse
func (client *Client) CreateGroupAliDingChat(request *CreateGroupAliDingChatRequest) (_result *CreateGroupAliDingChatResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateGroupAliDingChatResponse{}
	_body, _err := client.CreateGroupAliDingChatWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates knowledge from a standard DingTalk group chat for a group.
//
// Description:
//
// ## Request description
//
// - Connects a specified standard DingTalk group chat to the group knowledge base that the caller has joined.
//
// - The resource type is fixed to DINGTALK, the scope is fixed to GROUP, and the owning user is parsed from the gateway authentication identity.
//
// - groupId, chatId, and historyStartTime are required.
//
// - updateFrequency can be configured through preset or a five-segment cron expression for subsequent synchronization frequency.
//
// - The server verifies the caller\\"s group member identity and target group directory permissions. The same group chat can be created as different Sources.
//
// @param tmpReq - CreateGroupDingtalkChatRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGroupDingtalkChatResponse
func (client *Client) CreateGroupDingtalkChatWithOptions(tmpReq *CreateGroupDingtalkChatRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateGroupDingtalkChatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateGroupDingtalkChatShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateFrequency) {
		request.UpdateFrequencyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateFrequency, dara.String("updateFrequency"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ChatId) {
		body["chatId"] = request.ChatId
	}

	if !dara.IsNil(request.ChatName) {
		body["chatName"] = request.ChatName
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DirectoryId) {
		body["directoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.GroupId) {
		body["groupId"] = request.GroupId
	}

	if !dara.IsNil(request.HistoryStartTime) {
		body["historyStartTime"] = request.HistoryStartTime
	}

	if !dara.IsNil(request.Notes) {
		body["notes"] = request.Notes
	}

	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	if !dara.IsNil(request.SourceTags) {
		body["sourceTags"] = request.SourceTags
	}

	if !dara.IsNil(request.UpdateFrequencyShrink) {
		body["updateFrequency"] = request.UpdateFrequencyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGroupDingtalkChat"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/createGroupDingtalkChat"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGroupDingtalkChatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates knowledge from a standard DingTalk group chat for a group.
//
// Description:
//
// ## Request description
//
// - Connects a specified standard DingTalk group chat to the group knowledge base that the caller has joined.
//
// - The resource type is fixed to DINGTALK, the scope is fixed to GROUP, and the owning user is parsed from the gateway authentication identity.
//
// - groupId, chatId, and historyStartTime are required.
//
// - updateFrequency can be configured through preset or a five-segment cron expression for subsequent synchronization frequency.
//
// - The server verifies the caller\\"s group member identity and target group directory permissions. The same group chat can be created as different Sources.
//
// @param request - CreateGroupDingtalkChatRequest
//
// @return CreateGroupDingtalkChatResponse
func (client *Client) CreateGroupDingtalkChat(request *CreateGroupDingtalkChatRequest) (_result *CreateGroupDingtalkChatResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateGroupDingtalkChatResponse{}
	_body, _err := client.CreateGroupDingtalkChatWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a group knowledge resource from a single Lark online document using the current user\\"s Lark authorization.
//
// Description:
//
// ## Request description\\n\\nFixed as `ONLINE_DOC + FEISHU + GROUP`. `groupId` is required. If `directoryId` is omitted, the root directory of the group knowledge base is used. Group membership and directory write permissions are verified by the backend.
//
// @param tmpReq - CreateGroupFeishuDocRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGroupFeishuDocResponse
func (client *Client) CreateGroupFeishuDocWithOptions(tmpReq *CreateGroupFeishuDocRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateGroupFeishuDocResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateGroupFeishuDocShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ObjectBindings) {
		request.ObjectBindingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ObjectBindings, dara.String("objectBindings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SyncConfig) {
		request.SyncConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SyncConfig, dara.String("syncConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DirectoryId) {
		body["directoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.DocUrl) {
		body["docUrl"] = request.DocUrl
	}

	if !dara.IsNil(request.GroupId) {
		body["groupId"] = request.GroupId
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Notes) {
		body["notes"] = request.Notes
	}

	if !dara.IsNil(request.ObjectBindingsShrink) {
		body["objectBindings"] = request.ObjectBindingsShrink
	}

	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	if !dara.IsNil(request.SourceTags) {
		body["sourceTags"] = request.SourceTags
	}

	if !dara.IsNil(request.SyncConfigShrink) {
		body["syncConfig"] = request.SyncConfigShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGroupFeishuDoc"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/createGroupFeishuDoc"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGroupFeishuDocResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a group knowledge resource from a single Lark online document using the current user\\"s Lark authorization.
//
// Description:
//
// ## Request description\\n\\nFixed as `ONLINE_DOC + FEISHU + GROUP`. `groupId` is required. If `directoryId` is omitted, the root directory of the group knowledge base is used. Group membership and directory write permissions are verified by the backend.
//
// @param request - CreateGroupFeishuDocRequest
//
// @return CreateGroupFeishuDocResponse
func (client *Client) CreateGroupFeishuDoc(request *CreateGroupFeishuDocRequest) (_result *CreateGroupFeishuDocResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateGroupFeishuDocResponse{}
	_body, _err := client.CreateGroupFeishuDocWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uploads an AliDing online document to the enterprise knowledge base. Management permissions are required.
//
// Description:
//
// ## Request description
//
// - This operation adds an AliDing online document to a specified enterprise knowledge base.
//
// - The caller must have the `DEVELOPMENT_KB_MANAGE` permission.
//
// - `source_type` is fixed to `ONLINE_DOC`, `platform` is fixed to `ALI_DING`, and `scope` is fixed to `TENANT`.
//
// - If `directoryId` is not provided, the document is bound to the root directory of the current digital employee by default. If provided, it must be a valid directory ID under the current tenant.
//
// - The `filePublicUrl` parameter is required and specifies the publicly accessible URL of the AliDing online document to upload.
//
// - Optional parameters include `operatingObjectName` (digital employee name), `description` (resource description), `knowledgeId` (knowledge base ID), and `sourceTags` (resource tags).
//
// - A successful response returns information about the newly created resource, such as `sourceId`, `name`, `status`, `directoryId`, and creation time.
//
// @param request - CreateKnowledgeBaseAliDingDocRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateKnowledgeBaseAliDingDocResponse
func (client *Client) CreateKnowledgeBaseAliDingDocWithOptions(request *CreateKnowledgeBaseAliDingDocRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateKnowledgeBaseAliDingDocResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DirectoryId) {
		body["directoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.FilePublicUrl) {
		body["filePublicUrl"] = request.FilePublicUrl
	}

	if !dara.IsNil(request.KnowledgeId) {
		body["knowledgeId"] = request.KnowledgeId
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	if !dara.IsNil(request.SourceTags) {
		body["sourceTags"] = request.SourceTags
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateKnowledgeBaseAliDingDoc"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/createKnowledgeBaseAlidingDoc"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateKnowledgeBaseAliDingDocResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads an AliDing online document to the enterprise knowledge base. Management permissions are required.
//
// Description:
//
// ## Request description
//
// - This operation adds an AliDing online document to a specified enterprise knowledge base.
//
// - The caller must have the `DEVELOPMENT_KB_MANAGE` permission.
//
// - `source_type` is fixed to `ONLINE_DOC`, `platform` is fixed to `ALI_DING`, and `scope` is fixed to `TENANT`.
//
// - If `directoryId` is not provided, the document is bound to the root directory of the current digital employee by default. If provided, it must be a valid directory ID under the current tenant.
//
// - The `filePublicUrl` parameter is required and specifies the publicly accessible URL of the AliDing online document to upload.
//
// - Optional parameters include `operatingObjectName` (digital employee name), `description` (resource description), `knowledgeId` (knowledge base ID), and `sourceTags` (resource tags).
//
// - A successful response returns information about the newly created resource, such as `sourceId`, `name`, `status`, `directoryId`, and creation time.
//
// @param request - CreateKnowledgeBaseAliDingDocRequest
//
// @return CreateKnowledgeBaseAliDingDocResponse
func (client *Client) CreateKnowledgeBaseAliDingDoc(request *CreateKnowledgeBaseAliDingDocRequest) (_result *CreateKnowledgeBaseAliDingDocResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateKnowledgeBaseAliDingDocResponse{}
	_body, _err := client.CreateKnowledgeBaseAliDingDocWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a category directory in the enterprise knowledge base.
//
// Description:
//
// ## Request description
//
// - This operation allows users with the `DEVELOPMENT_KB_MANAGE` permission to create a category for the enterprise knowledge base.
//
// - You can specify a parent category ID when creating a category. If no parent category ID is specified, the new category is mounted directly under the root directory of the enterprise knowledge base.
//
// - The system automatically checks for name conflicts and directory depth limits.
//
// - `tenant_id` and `user_id` are obtained only through authentication. These parameters are ignored even if they are provided in the request body.
//
// - Ensure that the specified `parentDirectoryId` (if any) belongs to the current tenant.
//
// @param request - CreateKnowledgeBaseDirectoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateKnowledgeBaseDirectoryResponse
func (client *Client) CreateKnowledgeBaseDirectoryWithOptions(request *CreateKnowledgeBaseDirectoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateKnowledgeBaseDirectoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.ParentDirectoryId) {
		body["parentDirectoryId"] = request.ParentDirectoryId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateKnowledgeBaseDirectory"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/createKnowledgeBaseDirectory"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateKnowledgeBaseDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a category directory in the enterprise knowledge base.
//
// Description:
//
// ## Request description
//
// - This operation allows users with the `DEVELOPMENT_KB_MANAGE` permission to create a category for the enterprise knowledge base.
//
// - You can specify a parent category ID when creating a category. If no parent category ID is specified, the new category is mounted directly under the root directory of the enterprise knowledge base.
//
// - The system automatically checks for name conflicts and directory depth limits.
//
// - `tenant_id` and `user_id` are obtained only through authentication. These parameters are ignored even if they are provided in the request body.
//
// - Ensure that the specified `parentDirectoryId` (if any) belongs to the current tenant.
//
// @param request - CreateKnowledgeBaseDirectoryRequest
//
// @return CreateKnowledgeBaseDirectoryResponse
func (client *Client) CreateKnowledgeBaseDirectory(request *CreateKnowledgeBaseDirectoryRequest) (_result *CreateKnowledgeBaseDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateKnowledgeBaseDirectoryResponse{}
	_body, _err := client.CreateKnowledgeBaseDirectoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a single Lark online document in the enterprise knowledge base using the current user\\"s Lark authorization.
//
// Description:
//
// ## Request description\\n\\nFixed as `ONLINE_DOC + FEISHU + TENANT`. `directoryId` is required. The invoker must have the enterprise knowledge base feature permission and knowledge base management permission on the target knowledge base.
//
// @param tmpReq - CreateKnowledgeBaseFeishuDocRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateKnowledgeBaseFeishuDocResponse
func (client *Client) CreateKnowledgeBaseFeishuDocWithOptions(tmpReq *CreateKnowledgeBaseFeishuDocRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateKnowledgeBaseFeishuDocResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateKnowledgeBaseFeishuDocShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ObjectBindings) {
		request.ObjectBindingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ObjectBindings, dara.String("objectBindings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SyncConfig) {
		request.SyncConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SyncConfig, dara.String("syncConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DirectoryId) {
		body["directoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.DocUrl) {
		body["docUrl"] = request.DocUrl
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Notes) {
		body["notes"] = request.Notes
	}

	if !dara.IsNil(request.ObjectBindingsShrink) {
		body["objectBindings"] = request.ObjectBindingsShrink
	}

	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	if !dara.IsNil(request.SourceTags) {
		body["sourceTags"] = request.SourceTags
	}

	if !dara.IsNil(request.SyncConfigShrink) {
		body["syncConfig"] = request.SyncConfigShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateKnowledgeBaseFeishuDoc"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/createKnowledgeBaseFeishuDoc"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateKnowledgeBaseFeishuDocResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a single Lark online document in the enterprise knowledge base using the current user\\"s Lark authorization.
//
// Description:
//
// ## Request description\\n\\nFixed as `ONLINE_DOC + FEISHU + TENANT`. `directoryId` is required. The invoker must have the enterprise knowledge base feature permission and knowledge base management permission on the target knowledge base.
//
// @param request - CreateKnowledgeBaseFeishuDocRequest
//
// @return CreateKnowledgeBaseFeishuDocResponse
func (client *Client) CreateKnowledgeBaseFeishuDoc(request *CreateKnowledgeBaseFeishuDocRequest) (_result *CreateKnowledgeBaseFeishuDocResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateKnowledgeBaseFeishuDocResponse{}
	_body, _err := client.CreateKnowledgeBaseFeishuDocWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uploads a specified file to the enterprise knowledge base. Management permissions are required.
//
// Description:
//
// ## Operation description
//
// - This operation uploads a file to the enterprise knowledge base.
//
// - The `DEVELOPMENT_KB_MANAGE` permission is required to call this operation.
//
// - You must provide the Object Storage Service (OSS) persistent address (`filePath`) of the file when uploading.
//
// - Optional parameters include the public access URL and original file name to enhance the completeness of file information.
//
// - If `directoryId` is specified, the file is placed in the corresponding enterprise knowledge base directory. Otherwise, the file is bound to the default root directory of the current digital employee.
//
// - You can use `sourceTags` to add labels to resources for subsequent management and retrieval.
//
// - This operation initiates a billing item (UNSTRUCTURED_PARSE). Ensure that your account balance is sufficient.
//
// @param request - CreateKnowledgeBaseFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateKnowledgeBaseFileResponse
func (client *Client) CreateKnowledgeBaseFileWithOptions(request *CreateKnowledgeBaseFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateKnowledgeBaseFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DirectoryId) {
		body["directoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.FileExt) {
		body["fileExt"] = request.FileExt
	}

	if !dara.IsNil(request.FileName) {
		body["fileName"] = request.FileName
	}

	if !dara.IsNil(request.FilePath) {
		body["filePath"] = request.FilePath
	}

	if !dara.IsNil(request.FilePublicUrl) {
		body["filePublicUrl"] = request.FilePublicUrl
	}

	if !dara.IsNil(request.FileRecordId) {
		body["fileRecordId"] = request.FileRecordId
	}

	if !dara.IsNil(request.KnowledgeId) {
		body["knowledgeId"] = request.KnowledgeId
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	if !dara.IsNil(request.SourceTags) {
		body["sourceTags"] = request.SourceTags
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateKnowledgeBaseFile"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/createKnowledgeBaseFile"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateKnowledgeBaseFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads a specified file to the enterprise knowledge base. Management permissions are required.
//
// Description:
//
// ## Operation description
//
// - This operation uploads a file to the enterprise knowledge base.
//
// - The `DEVELOPMENT_KB_MANAGE` permission is required to call this operation.
//
// - You must provide the Object Storage Service (OSS) persistent address (`filePath`) of the file when uploading.
//
// - Optional parameters include the public access URL and original file name to enhance the completeness of file information.
//
// - If `directoryId` is specified, the file is placed in the corresponding enterprise knowledge base directory. Otherwise, the file is bound to the default root directory of the current digital employee.
//
// - You can use `sourceTags` to add labels to resources for subsequent management and retrieval.
//
// - This operation initiates a billing item (UNSTRUCTURED_PARSE). Ensure that your account balance is sufficient.
//
// @param request - CreateKnowledgeBaseFileRequest
//
// @return CreateKnowledgeBaseFileResponse
func (client *Client) CreateKnowledgeBaseFile(request *CreateKnowledgeBaseFileRequest) (_result *CreateKnowledgeBaseFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateKnowledgeBaseFileResponse{}
	_body, _err := client.CreateKnowledgeBaseFileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds plain text content to a specified enterprise knowledge base.
//
// Description:
//
// ## Operation description
//
// - This API is used to upload plain text information to an enterprise knowledge base. The caller must have the required management permissions.
//
// - The `textContent` field is required and represents the plain text content to upload.
//
// - Optional parameters include the digital employee name (`operatingObjectName`) and resource description (`description`), which allow users to customize additional details.
//
// - If `directoryId` is provided, the uploaded text is attached to the specified knowledge base folder. If not provided, the text is attached to the root folder of the current digital employee by default.
//
// - You can use `sourceTags` to add labels to resources for easier management and retrieval.
//
// - Before invoking this operation, make sure that you have correctly configured the authentication method (AK, BearerToken, and APP authentication are supported) and have the `DEVELOPMENT_KB_MANAGE` permission.
//
// @param request - CreateKnowledgeBaseTextRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateKnowledgeBaseTextResponse
func (client *Client) CreateKnowledgeBaseTextWithOptions(request *CreateKnowledgeBaseTextRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateKnowledgeBaseTextResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DirectoryId) {
		body["directoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.KnowledgeId) {
		body["knowledgeId"] = request.KnowledgeId
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	if !dara.IsNil(request.SourceTags) {
		body["sourceTags"] = request.SourceTags
	}

	if !dara.IsNil(request.TextContent) {
		body["textContent"] = request.TextContent
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateKnowledgeBaseText"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/createKnowledgeBaseText"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateKnowledgeBaseTextResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds plain text content to a specified enterprise knowledge base.
//
// Description:
//
// ## Operation description
//
// - This API is used to upload plain text information to an enterprise knowledge base. The caller must have the required management permissions.
//
// - The `textContent` field is required and represents the plain text content to upload.
//
// - Optional parameters include the digital employee name (`operatingObjectName`) and resource description (`description`), which allow users to customize additional details.
//
// - If `directoryId` is provided, the uploaded text is attached to the specified knowledge base folder. If not provided, the text is attached to the root folder of the current digital employee by default.
//
// - You can use `sourceTags` to add labels to resources for easier management and retrieval.
//
// - Before invoking this operation, make sure that you have correctly configured the authentication method (AK, BearerToken, and APP authentication are supported) and have the `DEVELOPMENT_KB_MANAGE` permission.
//
// @param request - CreateKnowledgeBaseTextRequest
//
// @return CreateKnowledgeBaseTextResponse
func (client *Client) CreateKnowledgeBaseText(request *CreateKnowledgeBaseTextRequest) (_result *CreateKnowledgeBaseTextResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateKnowledgeBaseTextResponse{}
	_body, _err := client.CreateKnowledgeBaseTextWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a personal AliDing group chat knowledge resource.
//
// Description:
//
// ## Request description
//
// - Connects the specified AliDing group chat to the current user\\"s personal knowledge base.
//
// - The resource type is fixed to ALI_DING, the scope is fixed to PERSONAL, and the owning user is parsed from the gateway authentication identity.
//
// - historyStartTime is required and supports YYYY-MM-DD or YYYY-MM-DD HH:MM:SS format.
//
// - updateFrequency can be configured with a preset or a five-field cron expression for subsequent synchronization frequency.
//
// - chatId must be unique within the target personal scope.
//
// @param tmpReq - CreatePersonalAliDingChatRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePersonalAliDingChatResponse
func (client *Client) CreatePersonalAliDingChatWithOptions(tmpReq *CreatePersonalAliDingChatRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePersonalAliDingChatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreatePersonalAliDingChatShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateFrequency) {
		request.UpdateFrequencyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateFrequency, dara.String("updateFrequency"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ChatId) {
		body["chatId"] = request.ChatId
	}

	if !dara.IsNil(request.ChatName) {
		body["chatName"] = request.ChatName
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DirectoryId) {
		body["directoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.HistoryStartTime) {
		body["historyStartTime"] = request.HistoryStartTime
	}

	if !dara.IsNil(request.Notes) {
		body["notes"] = request.Notes
	}

	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	if !dara.IsNil(request.SourceTags) {
		body["sourceTags"] = request.SourceTags
	}

	if !dara.IsNil(request.UpdateFrequencyShrink) {
		body["updateFrequency"] = request.UpdateFrequencyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePersonalAliDingChat"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/createPersonalAliDingChat"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePersonalAliDingChatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a personal AliDing group chat knowledge resource.
//
// Description:
//
// ## Request description
//
// - Connects the specified AliDing group chat to the current user\\"s personal knowledge base.
//
// - The resource type is fixed to ALI_DING, the scope is fixed to PERSONAL, and the owning user is parsed from the gateway authentication identity.
//
// - historyStartTime is required and supports YYYY-MM-DD or YYYY-MM-DD HH:MM:SS format.
//
// - updateFrequency can be configured with a preset or a five-field cron expression for subsequent synchronization frequency.
//
// - chatId must be unique within the target personal scope.
//
// @param request - CreatePersonalAliDingChatRequest
//
// @return CreatePersonalAliDingChatResponse
func (client *Client) CreatePersonalAliDingChat(request *CreatePersonalAliDingChatRequest) (_result *CreatePersonalAliDingChatResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePersonalAliDingChatResponse{}
	_body, _err := client.CreatePersonalAliDingChatWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uploads an AliDing meeting file to the personal resource library of the current digital employee.
//
// Description:
//
// ## Request description
//
// - This API uploads AliDing meeting materials (such as audio/video files and Shanji links) to the "My Resources" section of a specified digital employee.
//
// - The `source_type` is fixed to `ALI_DING_MEETING`, and the `scope` is fixed to `PERSONAL`.
//
// - You must provide a public audio/video OSS URL (`ossUrl`) and the original Shanji link (`shanjiUrl`).
//
// - Optionally, you can specify a target personal directory ID (`directoryId`). If not specified, the resource is automatically bound to the default root directory of the current digital employee.
//
// - You can add a resource description (`description`) and meeting notes (`notes`). The meeting notes can be used for auxiliary analysis.
//
// - This operation requires authentication. AK, BearerToken, and APP authentication methods are supported.
//
// @param request - CreatePersonalAliDingMeetingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePersonalAliDingMeetingResponse
func (client *Client) CreatePersonalAliDingMeetingWithOptions(request *CreatePersonalAliDingMeetingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePersonalAliDingMeetingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DirectoryId) {
		body["directoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Notes) {
		body["notes"] = request.Notes
	}

	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	if !dara.IsNil(request.ShanjiUrl) {
		body["shanjiUrl"] = request.ShanjiUrl
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePersonalAliDingMeeting"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/createPersonalAliDingMeeting"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePersonalAliDingMeetingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads an AliDing meeting file to the personal resource library of the current digital employee.
//
// Description:
//
// ## Request description
//
// - This API uploads AliDing meeting materials (such as audio/video files and Shanji links) to the "My Resources" section of a specified digital employee.
//
// - The `source_type` is fixed to `ALI_DING_MEETING`, and the `scope` is fixed to `PERSONAL`.
//
// - You must provide a public audio/video OSS URL (`ossUrl`) and the original Shanji link (`shanjiUrl`).
//
// - Optionally, you can specify a target personal directory ID (`directoryId`). If not specified, the resource is automatically bound to the default root directory of the current digital employee.
//
// - You can add a resource description (`description`) and meeting notes (`notes`). The meeting notes can be used for auxiliary analysis.
//
// - This operation requires authentication. AK, BearerToken, and APP authentication methods are supported.
//
// @param request - CreatePersonalAliDingMeetingRequest
//
// @return CreatePersonalAliDingMeetingResponse
func (client *Client) CreatePersonalAliDingMeeting(request *CreatePersonalAliDingMeetingRequest) (_result *CreatePersonalAliDingMeetingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePersonalAliDingMeetingResponse{}
	_body, _err := client.CreatePersonalAliDingMeetingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uploads an AliDing online document to the personal resources of the current digital employee.
//
// Description:
//
// ## Request description
//
// - This API is used to add an AliDing online document to the "My Resources" section of a specified digital employee.
//
// - Fixed parameters include `source_type=ONLINE_DOC`, `platform=ALI_DING`, and `scope=PERSONAL`.
//
// - If `directoryId` is not provided, the document is attached to the root folder of the current digital employee by default. If provided, ensure that the folder belongs to the current user and exists under the current digital employee.
//
// - During the invoke process, metering is started and related operation logs are recorded.
//
// - For security purposes, `tenant_id` and `user_id` are obtained only from the authentication identity. Values provided by the caller for these fields are ignored.
//
// - Any validation or execute failure is thrown as an exception by the service and transformed into a POP error code returned to the caller.
//
// @param request - CreatePersonalAlidingDocRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePersonalAlidingDocResponse
func (client *Client) CreatePersonalAlidingDocWithOptions(request *CreatePersonalAlidingDocRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePersonalAlidingDocResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DirectoryId) {
		body["directoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.FilePublicUrl) {
		body["filePublicUrl"] = request.FilePublicUrl
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePersonalAlidingDoc"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/createPersonalAliDingDoc"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePersonalAlidingDocResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads an AliDing online document to the personal resources of the current digital employee.
//
// Description:
//
// ## Request description
//
// - This API is used to add an AliDing online document to the "My Resources" section of a specified digital employee.
//
// - Fixed parameters include `source_type=ONLINE_DOC`, `platform=ALI_DING`, and `scope=PERSONAL`.
//
// - If `directoryId` is not provided, the document is attached to the root folder of the current digital employee by default. If provided, ensure that the folder belongs to the current user and exists under the current digital employee.
//
// - During the invoke process, metering is started and related operation logs are recorded.
//
// - For security purposes, `tenant_id` and `user_id` are obtained only from the authentication identity. Values provided by the caller for these fields are ignored.
//
// - Any validation or execute failure is thrown as an exception by the service and transformed into a POP error code returned to the caller.
//
// @param request - CreatePersonalAlidingDocRequest
//
// @return CreatePersonalAlidingDocResponse
func (client *Client) CreatePersonalAlidingDoc(request *CreatePersonalAlidingDocRequest) (_result *CreatePersonalAlidingDocResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePersonalAlidingDocResponse{}
	_body, _err := client.CreatePersonalAlidingDocWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds an AliDing knowledge base to the personal resources of the current digital employee.
//
// Description:
//
// ## Request description
//
// - This API creates an AliDing knowledge base and mounts it to the personal resource directory of the specified digital employee.
//
// - `platform` is fixed to `ALI_DING`, and `directory_type` is fixed to `PERSONAL`.
//
// - If `directoryId` is provided, the system verifies that the directory exists and belongs to the current tenant and is of the personal type.
//
// - During creation, the knowledge base root directory is initialized (with the status set to `RUNNING`), and background tasks are dispatched based on the provided synchronization configuration to pull the remote directory tree and create child nodes.
//
// - For security purposes, `tenant_id` and `user_id` are obtained only from the authenticated identity. These fields in the request body are ignored.
//
// - The synchronization configuration is optional. If enabled, a cron expression must be provided. If not provided or disabled, scheduled synchronization is not performed by default.
//
// - The knowledge base name can be customized. If not provided, it is automatically populated after background synchronization.
//
// - Multi-value object binding is supported. Related information is serialized and stored in the knowledge base metadata.
//
// @param tmpReq - CreatePersonalAlidingKnowledgeBaseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePersonalAlidingKnowledgeBaseResponse
func (client *Client) CreatePersonalAlidingKnowledgeBaseWithOptions(tmpReq *CreatePersonalAlidingKnowledgeBaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePersonalAlidingKnowledgeBaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreatePersonalAlidingKnowledgeBaseShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ObjectBindings) {
		request.ObjectBindingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ObjectBindings, dara.String("objectBindings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SyncConfig) {
		request.SyncConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SyncConfig, dara.String("syncConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		body["directoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.KbName) {
		body["kbName"] = request.KbName
	}

	if !dara.IsNil(request.KbUrl) {
		body["kbUrl"] = request.KbUrl
	}

	if !dara.IsNil(request.ObjectBindingsShrink) {
		body["objectBindings"] = request.ObjectBindingsShrink
	}

	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	if !dara.IsNil(request.SyncConfigShrink) {
		body["syncConfig"] = request.SyncConfigShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePersonalAlidingKnowledgeBase"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/createPersonalAliDingKnowledgeBase"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePersonalAlidingKnowledgeBaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an AliDing knowledge base to the personal resources of the current digital employee.
//
// Description:
//
// ## Request description
//
// - This API creates an AliDing knowledge base and mounts it to the personal resource directory of the specified digital employee.
//
// - `platform` is fixed to `ALI_DING`, and `directory_type` is fixed to `PERSONAL`.
//
// - If `directoryId` is provided, the system verifies that the directory exists and belongs to the current tenant and is of the personal type.
//
// - During creation, the knowledge base root directory is initialized (with the status set to `RUNNING`), and background tasks are dispatched based on the provided synchronization configuration to pull the remote directory tree and create child nodes.
//
// - For security purposes, `tenant_id` and `user_id` are obtained only from the authenticated identity. These fields in the request body are ignored.
//
// - The synchronization configuration is optional. If enabled, a cron expression must be provided. If not provided or disabled, scheduled synchronization is not performed by default.
//
// - The knowledge base name can be customized. If not provided, it is automatically populated after background synchronization.
//
// - Multi-value object binding is supported. Related information is serialized and stored in the knowledge base metadata.
//
// @param request - CreatePersonalAlidingKnowledgeBaseRequest
//
// @return CreatePersonalAlidingKnowledgeBaseResponse
func (client *Client) CreatePersonalAlidingKnowledgeBase(request *CreatePersonalAlidingKnowledgeBaseRequest) (_result *CreatePersonalAlidingKnowledgeBaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePersonalAlidingKnowledgeBaseResponse{}
	_body, _err := client.CreatePersonalAlidingKnowledgeBaseWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a personal DingTalk group chat knowledge source.
//
// Description:
//
// ## Operation description
//
// - Connects a specified DingTalk group chat to the personal knowledge base of the current user.
//
// - The resource type is fixed to DINGTALK, the scope is fixed to PERSONAL, and the owning user is parsed from the gateway authentication identity.
//
// - historyStartTime is required and supports YYYY-MM-DD or YYYY-MM-DD HH:MM:SS format.
//
// - updateFrequency can be configured with a preset or a five-field cron expression for subsequent synchronization frequency.
//
// - The same group chat can be created as different sources. Each source is isolated by sourceId.
//
// @param tmpReq - CreatePersonalDingtalkChatRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePersonalDingtalkChatResponse
func (client *Client) CreatePersonalDingtalkChatWithOptions(tmpReq *CreatePersonalDingtalkChatRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePersonalDingtalkChatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreatePersonalDingtalkChatShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateFrequency) {
		request.UpdateFrequencyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateFrequency, dara.String("updateFrequency"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ChatId) {
		body["chatId"] = request.ChatId
	}

	if !dara.IsNil(request.ChatName) {
		body["chatName"] = request.ChatName
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DirectoryId) {
		body["directoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.HistoryStartTime) {
		body["historyStartTime"] = request.HistoryStartTime
	}

	if !dara.IsNil(request.Notes) {
		body["notes"] = request.Notes
	}

	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	if !dara.IsNil(request.SourceTags) {
		body["sourceTags"] = request.SourceTags
	}

	if !dara.IsNil(request.UpdateFrequencyShrink) {
		body["updateFrequency"] = request.UpdateFrequencyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePersonalDingtalkChat"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/createPersonalDingtalkChat"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePersonalDingtalkChatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a personal DingTalk group chat knowledge source.
//
// Description:
//
// ## Operation description
//
// - Connects a specified DingTalk group chat to the personal knowledge base of the current user.
//
// - The resource type is fixed to DINGTALK, the scope is fixed to PERSONAL, and the owning user is parsed from the gateway authentication identity.
//
// - historyStartTime is required and supports YYYY-MM-DD or YYYY-MM-DD HH:MM:SS format.
//
// - updateFrequency can be configured with a preset or a five-field cron expression for subsequent synchronization frequency.
//
// - The same group chat can be created as different sources. Each source is isolated by sourceId.
//
// @param request - CreatePersonalDingtalkChatRequest
//
// @return CreatePersonalDingtalkChatResponse
func (client *Client) CreatePersonalDingtalkChat(request *CreatePersonalDingtalkChatRequest) (_result *CreatePersonalDingtalkChatResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePersonalDingtalkChatResponse{}
	_body, _err := client.CreatePersonalDingtalkChatWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI CreatePersonalDingtalkMeeting is deprecated
//
// Summary:
//
// Uploads a DingTalk meeting to the personal resource library of the current digital employee.
//
// Description:
//
// ## Request description
//
// - This operation uploads a DingTalk meeting as a resource to the "My Resources" section of a specified digital employee.
//
// - `source_type` is fixed to `DINGTALK_MEETING`, and `scope` is fixed to `PERSONAL`.
//
// - If `credentialId` is not provided, the system default configurations are used.
//
// - If `directoryId` is not specified, the resource is automatically attached to the default root folder of the current digital employee. If specified, it must be an existing personal folder of the invoker under the digital employee.
//
// - The optional parameters `description` and `notes` are used to describe the resource and record meeting notes, respectively. The `notes` value is used for auxiliary analysis.
//
// @param request - CreatePersonalDingtalkMeetingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePersonalDingtalkMeetingResponse
func (client *Client) CreatePersonalDingtalkMeetingWithOptions(request *CreatePersonalDingtalkMeetingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePersonalDingtalkMeetingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CredentialId) {
		body["credentialId"] = request.CredentialId
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DirectoryId) {
		body["directoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Notes) {
		body["notes"] = request.Notes
	}

	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	if !dara.IsNil(request.RoomCode) {
		body["roomCode"] = request.RoomCode
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePersonalDingtalkMeeting"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/createPersonalDingtalkMeeting"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePersonalDingtalkMeetingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI CreatePersonalDingtalkMeeting is deprecated
//
// Summary:
//
// Uploads a DingTalk meeting to the personal resource library of the current digital employee.
//
// Description:
//
// ## Request description
//
// - This operation uploads a DingTalk meeting as a resource to the "My Resources" section of a specified digital employee.
//
// - `source_type` is fixed to `DINGTALK_MEETING`, and `scope` is fixed to `PERSONAL`.
//
// - If `credentialId` is not provided, the system default configurations are used.
//
// - If `directoryId` is not specified, the resource is automatically attached to the default root folder of the current digital employee. If specified, it must be an existing personal folder of the invoker under the digital employee.
//
// - The optional parameters `description` and `notes` are used to describe the resource and record meeting notes, respectively. The `notes` value is used for auxiliary analysis.
//
// @param request - CreatePersonalDingtalkMeetingRequest
//
// @return CreatePersonalDingtalkMeetingResponse
// Deprecated
func (client *Client) CreatePersonalDingtalkMeeting(request *CreatePersonalDingtalkMeetingRequest) (_result *CreatePersonalDingtalkMeetingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePersonalDingtalkMeetingResponse{}
	_body, _err := client.CreatePersonalDingtalkMeetingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uploads a meeting to the current user\\"s personal knowledge base by using a standard DingTalk Shanji URL.
//
// Description:
//
// ## Request description
//
// - This API creates a meeting resource by using a standard DingTalk Shanji link. The collection method is fixed to the DWS corresponding to personal OAuth.
//
// - `source_type` is fixed to `DINGTALK_MEETING`, and `scope` is fixed to `PERSONAL`.
//
// - You must provide a standard DingTalk Shanji link or taskUuid (`shanjiUrl`).
//
// - Optionally specify a target personal directory ID (`directoryId`). If not specified, the default root directory of the current digital employee is used.
//
// - You can add a resource description (`description`) and meeting notes (`notes`).
//
// - This operation supports one of the following authentication methods: AK, BearerToken, or APP.
//
// @param request - CreatePersonalDingtalkMinutesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePersonalDingtalkMinutesResponse
func (client *Client) CreatePersonalDingtalkMinutesWithOptions(request *CreatePersonalDingtalkMinutesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePersonalDingtalkMinutesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DirectoryId) {
		body["directoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Notes) {
		body["notes"] = request.Notes
	}

	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	if !dara.IsNil(request.ShanjiUrl) {
		body["shanjiUrl"] = request.ShanjiUrl
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePersonalDingtalkMinutes"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/createPersonalDingtalkMinutes"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePersonalDingtalkMinutesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads a meeting to the current user\\"s personal knowledge base by using a standard DingTalk Shanji URL.
//
// Description:
//
// ## Request description
//
// - This API creates a meeting resource by using a standard DingTalk Shanji link. The collection method is fixed to the DWS corresponding to personal OAuth.
//
// - `source_type` is fixed to `DINGTALK_MEETING`, and `scope` is fixed to `PERSONAL`.
//
// - You must provide a standard DingTalk Shanji link or taskUuid (`shanjiUrl`).
//
// - Optionally specify a target personal directory ID (`directoryId`). If not specified, the default root directory of the current digital employee is used.
//
// - You can add a resource description (`description`) and meeting notes (`notes`).
//
// - This operation supports one of the following authentication methods: AK, BearerToken, or APP.
//
// @param request - CreatePersonalDingtalkMinutesRequest
//
// @return CreatePersonalDingtalkMinutesResponse
func (client *Client) CreatePersonalDingtalkMinutes(request *CreatePersonalDingtalkMinutesRequest) (_result *CreatePersonalDingtalkMinutesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePersonalDingtalkMinutesResponse{}
	_body, _err := client.CreatePersonalDingtalkMinutesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a personal folder (category) under My Resources.
//
// Description:
//
// ## Request description
//
// - This API is used to create a personal folder (category) under "My Resources".
//
// - If `parentDirectoryId` is not specified, the system automatically uses or creates the default root folder of the current digital employee as the parent folder.
//
// - If `parentDirectoryId` is specified, it must be an existing personal folder of the current user under the current digital employee.
//
// - `tenant_id` and `user_id` are derived from the authenticated identity only. These fields are ignored if passed in the request body.
//
// @param request - CreatePersonalDirectoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePersonalDirectoryResponse
func (client *Client) CreatePersonalDirectoryWithOptions(request *CreatePersonalDirectoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePersonalDirectoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	if !dara.IsNil(request.ParentDirectoryId) {
		body["parentDirectoryId"] = request.ParentDirectoryId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePersonalDirectory"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/createPersonalDirectory"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePersonalDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a personal folder (category) under My Resources.
//
// Description:
//
// ## Request description
//
// - This API is used to create a personal folder (category) under "My Resources".
//
// - If `parentDirectoryId` is not specified, the system automatically uses or creates the default root folder of the current digital employee as the parent folder.
//
// - If `parentDirectoryId` is specified, it must be an existing personal folder of the current user under the current digital employee.
//
// - `tenant_id` and `user_id` are derived from the authenticated identity only. These fields are ignored if passed in the request body.
//
// @param request - CreatePersonalDirectoryRequest
//
// @return CreatePersonalDirectoryResponse
func (client *Client) CreatePersonalDirectory(request *CreatePersonalDirectoryRequest) (_result *CreatePersonalDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePersonalDirectoryResponse{}
	_body, _err := client.CreatePersonalDirectoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a personal knowledge resource from a single Lark group chat using the current user\\"s Lark authorization.
//
// Description:
//
// ## Request description
//
// This API uses the Lark application connection managed by the user corresponding to the current OpenAPI identity. It pulls the name and historical messages of the specified group chat through the built-in CLI of the project and creates a knowledge resource in the user\\"s personal knowledge base.
//
// - `chatId`: The Lark group chat ID. Must start with `oc_`.
//
// - `directoryId` (optional): The target personal directory ID. If omitted, the current user\\"s default personal root directory is used.
//
// - `historyStartTime` (optional): The start time for historical messages. Supports `YYYY-MM-DD` or `YYYY-MM-DD HH:MM:SS`.
//
// - `updateFrequency` (optional): The Source-level scheduled synchronization configuration. Supports preset frequencies or five-field cron expressions.
//
// - `description`, `operatingObjectName`, `notes`, `sourceTags`: Optional Source metadata.
//
// Security constraints: The Source Type is fixed to FEISHU, and the knowledge scope is fixed to PERSONAL. The Lark connector user is determined by the POP trusted identity. Credentials or user IDs passed by the caller are not accepted.
//
// @param tmpReq - CreatePersonalFeishuChatRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePersonalFeishuChatResponse
func (client *Client) CreatePersonalFeishuChatWithOptions(tmpReq *CreatePersonalFeishuChatRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePersonalFeishuChatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreatePersonalFeishuChatShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateFrequency) {
		request.UpdateFrequencyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateFrequency, dara.String("updateFrequency"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ChatId) {
		body["chatId"] = request.ChatId
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DirectoryId) {
		body["directoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.HistoryStartTime) {
		body["historyStartTime"] = request.HistoryStartTime
	}

	if !dara.IsNil(request.Notes) {
		body["notes"] = request.Notes
	}

	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	if !dara.IsNil(request.SourceTags) {
		body["sourceTags"] = request.SourceTags
	}

	if !dara.IsNil(request.UpdateFrequencyShrink) {
		body["updateFrequency"] = request.UpdateFrequencyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePersonalFeishuChat"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/createPersonalFeishuChat"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePersonalFeishuChatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a personal knowledge resource from a single Lark group chat using the current user\\"s Lark authorization.
//
// Description:
//
// ## Request description
//
// This API uses the Lark application connection managed by the user corresponding to the current OpenAPI identity. It pulls the name and historical messages of the specified group chat through the built-in CLI of the project and creates a knowledge resource in the user\\"s personal knowledge base.
//
// - `chatId`: The Lark group chat ID. Must start with `oc_`.
//
// - `directoryId` (optional): The target personal directory ID. If omitted, the current user\\"s default personal root directory is used.
//
// - `historyStartTime` (optional): The start time for historical messages. Supports `YYYY-MM-DD` or `YYYY-MM-DD HH:MM:SS`.
//
// - `updateFrequency` (optional): The Source-level scheduled synchronization configuration. Supports preset frequencies or five-field cron expressions.
//
// - `description`, `operatingObjectName`, `notes`, `sourceTags`: Optional Source metadata.
//
// Security constraints: The Source Type is fixed to FEISHU, and the knowledge scope is fixed to PERSONAL. The Lark connector user is determined by the POP trusted identity. Credentials or user IDs passed by the caller are not accepted.
//
// @param request - CreatePersonalFeishuChatRequest
//
// @return CreatePersonalFeishuChatResponse
func (client *Client) CreatePersonalFeishuChat(request *CreatePersonalFeishuChatRequest) (_result *CreatePersonalFeishuChatResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePersonalFeishuChatResponse{}
	_body, _err := client.CreatePersonalFeishuChatWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a personal knowledge resource from a single Lark online document using the current user\\"s Lark authorization.
//
// Description:
//
// ## Request description\\n\\nFixed as `ONLINE_DOC + FEISHU + PERSONAL`. The Lark connector user is determined by the trusted OpenAPI identity. If `directoryId` is omitted, the current user\\"s default personal root directory is used.
//
// @param tmpReq - CreatePersonalFeishuDocRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePersonalFeishuDocResponse
func (client *Client) CreatePersonalFeishuDocWithOptions(tmpReq *CreatePersonalFeishuDocRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePersonalFeishuDocResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreatePersonalFeishuDocShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ObjectBindings) {
		request.ObjectBindingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ObjectBindings, dara.String("objectBindings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SyncConfig) {
		request.SyncConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SyncConfig, dara.String("syncConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DirectoryId) {
		body["directoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.DocUrl) {
		body["docUrl"] = request.DocUrl
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Notes) {
		body["notes"] = request.Notes
	}

	if !dara.IsNil(request.ObjectBindingsShrink) {
		body["objectBindings"] = request.ObjectBindingsShrink
	}

	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	if !dara.IsNil(request.SourceTags) {
		body["sourceTags"] = request.SourceTags
	}

	if !dara.IsNil(request.SyncConfigShrink) {
		body["syncConfig"] = request.SyncConfigShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePersonalFeishuDoc"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/createPersonalFeishuDoc"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePersonalFeishuDocResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a personal knowledge resource from a single Lark online document using the current user\\"s Lark authorization.
//
// Description:
//
// ## Request description\\n\\nFixed as `ONLINE_DOC + FEISHU + PERSONAL`. The Lark connector user is determined by the trusted OpenAPI identity. If `directoryId` is omitted, the current user\\"s default personal root directory is used.
//
// @param request - CreatePersonalFeishuDocRequest
//
// @return CreatePersonalFeishuDocResponse
func (client *Client) CreatePersonalFeishuDoc(request *CreatePersonalFeishuDocRequest) (_result *CreatePersonalFeishuDocResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePersonalFeishuDocResponse{}
	_body, _err := client.CreatePersonalFeishuDocWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uploads a Lark Minutes meeting file to the personal resource library of the current digital employee.
//
// Description:
//
// ## Request description
//
// This API uploads a meeting record from Lark Minutes to the "My Resources" section of a specified digital employee. By providing the required parameters, such as the unique identifier of the Lark Minutes record (`minuteToken`) and the credential ID (`credentialId`), you can migrate and save meeting content. If no target directory is specified, the resource is bound to the root directory of the current digital employee by default.
//
// - `operatingObjectName`: The name of the digital employee that performs the operation.
//
// - `name`: The display name of the uploaded resource in the system.
//
// - `minuteToken`: The unique identifier of the meeting from the Lark Minutes platform.
//
// - `credentialId`: The ID associated with specific authentication information, used to verify the validity of the request.
//
// - `directoryId` (optional): The ID of the target personal directory where the resource is stored. If this field is omitted, the resource is automatically placed in the default location.
//
// - `description` (optional): A brief description or note about the uploaded resource.
//
// Precautions:
//
// - Ensure that the provided `minuteToken` and `credentialId` are valid.
//
// - If `directoryId` is specified, confirm that it belongs to one of the available personal directories of the caller in the current digital employee environment.
//
// @param request - CreatePersonalFeishuMinuteRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePersonalFeishuMinuteResponse
func (client *Client) CreatePersonalFeishuMinuteWithOptions(request *CreatePersonalFeishuMinuteRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePersonalFeishuMinuteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CredentialId) {
		body["credentialId"] = request.CredentialId
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DirectoryId) {
		body["directoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.MinuteToken) {
		body["minuteToken"] = request.MinuteToken
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePersonalFeishuMinute"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/createPersonalFeishuMinute"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePersonalFeishuMinuteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads a Lark Minutes meeting file to the personal resource library of the current digital employee.
//
// Description:
//
// ## Request description
//
// This API uploads a meeting record from Lark Minutes to the "My Resources" section of a specified digital employee. By providing the required parameters, such as the unique identifier of the Lark Minutes record (`minuteToken`) and the credential ID (`credentialId`), you can migrate and save meeting content. If no target directory is specified, the resource is bound to the root directory of the current digital employee by default.
//
// - `operatingObjectName`: The name of the digital employee that performs the operation.
//
// - `name`: The display name of the uploaded resource in the system.
//
// - `minuteToken`: The unique identifier of the meeting from the Lark Minutes platform.
//
// - `credentialId`: The ID associated with specific authentication information, used to verify the validity of the request.
//
// - `directoryId` (optional): The ID of the target personal directory where the resource is stored. If this field is omitted, the resource is automatically placed in the default location.
//
// - `description` (optional): A brief description or note about the uploaded resource.
//
// Precautions:
//
// - Ensure that the provided `minuteToken` and `credentialId` are valid.
//
// - If `directoryId` is specified, confirm that it belongs to one of the available personal directories of the caller in the current digital employee environment.
//
// @param request - CreatePersonalFeishuMinuteRequest
//
// @return CreatePersonalFeishuMinuteResponse
func (client *Client) CreatePersonalFeishuMinute(request *CreatePersonalFeishuMinuteRequest) (_result *CreatePersonalFeishuMinuteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePersonalFeishuMinuteResponse{}
	_body, _err := client.CreatePersonalFeishuMinuteWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uploads a file to the personal resource library of the current digital employee.
//
// Description:
//
// ## Request description
//
// - This API is used to upload a file to the "My Resources" section of a specified digital employee.
//
// - `source_type` is fixed to `FILE`, `scope` is fixed to `PERSONAL`, and `platform` is fixed to `LOCAL`.
//
// - The file must include an OSS persistent address (`filePath`). Other information such as the public access URL and original file name is optional.
//
// - If the target folder ID (`directoryId`) is not specified, the file is automatically attached to the default root folder of the current digital employee. If specified, ensure that the folder belongs to the personal folder of the caller.
//
// - Security authentication is supported through multiple authentication methods (AK, BearerToken, and APP) to authenticate requests.
//
// - The operation type is write (`write`), and operation logs are recorded for subsequent auditing.
//
// To invoke this operation, you can use AK, BearerToken, or APP authentication.
//
// @param request - CreatePersonalFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePersonalFileResponse
func (client *Client) CreatePersonalFileWithOptions(request *CreatePersonalFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePersonalFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DirectoryId) {
		body["directoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.FileExt) {
		body["fileExt"] = request.FileExt
	}

	if !dara.IsNil(request.FileName) {
		body["fileName"] = request.FileName
	}

	if !dara.IsNil(request.FilePath) {
		body["filePath"] = request.FilePath
	}

	if !dara.IsNil(request.FilePublicUrl) {
		body["filePublicUrl"] = request.FilePublicUrl
	}

	if !dara.IsNil(request.FileRecordId) {
		body["fileRecordId"] = request.FileRecordId
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePersonalFile"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/createPersonalFile"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePersonalFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads a file to the personal resource library of the current digital employee.
//
// Description:
//
// ## Request description
//
// - This API is used to upload a file to the "My Resources" section of a specified digital employee.
//
// - `source_type` is fixed to `FILE`, `scope` is fixed to `PERSONAL`, and `platform` is fixed to `LOCAL`.
//
// - The file must include an OSS persistent address (`filePath`). Other information such as the public access URL and original file name is optional.
//
// - If the target folder ID (`directoryId`) is not specified, the file is automatically attached to the default root folder of the current digital employee. If specified, ensure that the folder belongs to the personal folder of the caller.
//
// - Security authentication is supported through multiple authentication methods (AK, BearerToken, and APP) to authenticate requests.
//
// - The operation type is write (`write`), and operation logs are recorded for subsequent auditing.
//
// To invoke this operation, you can use AK, BearerToken, or APP authentication.
//
// @param request - CreatePersonalFileRequest
//
// @return CreatePersonalFileResponse
func (client *Client) CreatePersonalFile(request *CreatePersonalFileRequest) (_result *CreatePersonalFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePersonalFileResponse{}
	_body, _err := client.CreatePersonalFileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uploads plain text content to the personal resource library of the current digital employee.
//
// Description:
//
// ## Operation description
//
// - This API is used to add plain text content to the personal resources of a specified digital employee.
//
// - `source_type` is fixed to `TEXT`, and `scope` is fixed to `PERSONAL`.
//
// - If `directoryId` is not provided, the content is bound to the root directory of the current digital employee by default. If provided, it must be an existing personal directory of the caller under the digital employee.
//
// - `tenant_id` and `user_id` can only be obtained from the authentication identity information. These parameters are ignored if passed in the request body.
//
// - The call initiates metering and generates a corresponding `billing_id`.
//
// - The text content is written to `unstructured_docs`, and an initial resource record is generated.
//
// - Any validation or execution failure throws a `RobjectException`, which is converted to a POP error code by the global middleware and returned to the caller.
//
// @param request - CreatePersonalTextRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePersonalTextResponse
func (client *Client) CreatePersonalTextWithOptions(request *CreatePersonalTextRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePersonalTextResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DirectoryId) {
		body["directoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	if !dara.IsNil(request.TextContent) {
		body["textContent"] = request.TextContent
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePersonalText"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/createPersonalText"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePersonalTextResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads plain text content to the personal resource library of the current digital employee.
//
// Description:
//
// ## Operation description
//
// - This API is used to add plain text content to the personal resources of a specified digital employee.
//
// - `source_type` is fixed to `TEXT`, and `scope` is fixed to `PERSONAL`.
//
// - If `directoryId` is not provided, the content is bound to the root directory of the current digital employee by default. If provided, it must be an existing personal directory of the caller under the digital employee.
//
// - `tenant_id` and `user_id` can only be obtained from the authentication identity information. These parameters are ignored if passed in the request body.
//
// - The call initiates metering and generates a corresponding `billing_id`.
//
// - The text content is written to `unstructured_docs`, and an initial resource record is generated.
//
// - Any validation or execution failure throws a `RobjectException`, which is converted to a POP error code by the global middleware and returned to the caller.
//
// @param request - CreatePersonalTextRequest
//
// @return CreatePersonalTextResponse
func (client *Client) CreatePersonalText(request *CreatePersonalTextRequest) (_result *CreatePersonalTextResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePersonalTextResponse{}
	_body, _err := client.CreatePersonalTextWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uploads an offline meeting audio file to the personal resources of the current digital employee.
//
// Description:
//
// ## Operation description
//
// - This API operation uploads an offline meeting audio file to the "My Resources" section of a specified digital employee.
//
// - `source_type` is fixed to `VOICE_MEETING`, `scope` is fixed to `PERSONAL`, and `voice_meeting_type` is fixed to `OFFLINE`.
//
// - If `directoryId` is not provided in the request body, the resource is automatically bound to the default root directory. If `directoryId` is provided, it must be an existing personal directory of the current user under the current digital employee.
//
// - Calling this operation starts a background process to transcribe the audio file and returns information about the newly created resource.
//
// - For security purposes, `tenant_id` and `user_id` are obtained only from the authenticated identity. These fields are ignored even if they are included in the request body.
//
// - Any validation or execution failure throws a `RobjectException`, which is converted to a POP error code through the global middleware.
//
// @param request - CreatePersonalVoiceMeetingRequest
//
// @param headers - CreatePersonalVoiceMeetingHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePersonalVoiceMeetingResponse
func (client *Client) CreatePersonalVoiceMeetingWithOptions(request *CreatePersonalVoiceMeetingRequest, headers *CreatePersonalVoiceMeetingHeaders, runtime *dara.RuntimeOptions) (_result *CreatePersonalVoiceMeetingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DirectoryId) {
		body["directoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.FileUrl) {
		body["fileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.RequestId) {
		realHeaders["requestId"] = dara.String(dara.ToString(dara.StringValue(headers.RequestId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePersonalVoiceMeeting"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/createPersonalVoiceMeeting"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePersonalVoiceMeetingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads an offline meeting audio file to the personal resources of the current digital employee.
//
// Description:
//
// ## Operation description
//
// - This API operation uploads an offline meeting audio file to the "My Resources" section of a specified digital employee.
//
// - `source_type` is fixed to `VOICE_MEETING`, `scope` is fixed to `PERSONAL`, and `voice_meeting_type` is fixed to `OFFLINE`.
//
// - If `directoryId` is not provided in the request body, the resource is automatically bound to the default root directory. If `directoryId` is provided, it must be an existing personal directory of the current user under the current digital employee.
//
// - Calling this operation starts a background process to transcribe the audio file and returns information about the newly created resource.
//
// - For security purposes, `tenant_id` and `user_id` are obtained only from the authenticated identity. These fields are ignored even if they are included in the request body.
//
// - Any validation or execution failure throws a `RobjectException`, which is converted to a POP error code through the global middleware.
//
// @param request - CreatePersonalVoiceMeetingRequest
//
// @return CreatePersonalVoiceMeetingResponse
func (client *Client) CreatePersonalVoiceMeeting(request *CreatePersonalVoiceMeetingRequest) (_result *CreatePersonalVoiceMeetingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreatePersonalVoiceMeetingHeaders{}
	_result = &CreatePersonalVoiceMeetingResponse{}
	_body, _err := client.CreatePersonalVoiceMeetingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a scheduled task.
//
// Description:
//
// ## Operation description
//
// - This operation is used to upload files to an enterprise knowledge base.
//
// - You must have the `DEVELOPMENT_KB_MANAGE` permission to call this API operation.
//
// - You must provide the OSS persistent address (`filePath`) of the file when uploading.
//
// - Optional parameters include the public access URL of the file and the original file name to enhance the completeness of file information.
//
// - If `directoryId` is specified, the file is placed in the corresponding enterprise knowledge base directory. Otherwise, the file is bound to the default root directory of the current digital employee by default.
//
// - You can add tags to resources by using `sourceTags` for subsequent management and retrieval.
//
// - This operation initiates a billing item (UNSTRUCTURED_PARSE). Make sure that your account balance is sufficient.
//
// @param tmpReq - CreateScheduledTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScheduledTaskResponse
func (client *Client) CreateScheduledTaskWithOptions(tmpReq *CreateScheduledTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateScheduledTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateScheduledTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Description) {
		request.DescriptionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Description, dara.String("description"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DigitalEmployeeName) {
		request.DigitalEmployeeNameShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DigitalEmployeeName, dara.String("digitalEmployeeName"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Segments) {
		request.SegmentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Segments, dara.String("segments"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TaskDetail) {
		request.TaskDetailShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskDetail, dara.String("taskDetail"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TriggerConfig) {
		request.TriggerConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TriggerConfig, dara.String("triggerConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.VisibleMemberUserIds) {
		request.VisibleMemberUserIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VisibleMemberUserIds, dara.String("visibleMemberUserIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CollaborationGroupId) {
		body["collaborationGroupId"] = request.CollaborationGroupId
	}

	if !dara.IsNil(request.DescriptionShrink) {
		body["description"] = request.DescriptionShrink
	}

	if !dara.IsNil(request.DigitalEmployeeNameShrink) {
		body["digitalEmployeeName"] = request.DigitalEmployeeNameShrink
	}

	if !dara.IsNil(request.IsOpen) {
		body["isOpen"] = request.IsOpen
	}

	if !dara.IsNil(request.Model) {
		body["model"] = request.Model
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.SegmentsShrink) {
		body["segments"] = request.SegmentsShrink
	}

	if !dara.IsNil(request.TaskDetailShrink) {
		body["taskDetail"] = request.TaskDetailShrink
	}

	if !dara.IsNil(request.TriggerConfigShrink) {
		body["triggerConfig"] = request.TriggerConfigShrink
	}

	if !dara.IsNil(request.Visibility) {
		body["visibility"] = request.Visibility
	}

	if !dara.IsNil(request.VisibleMemberUserIdsShrink) {
		body["visibleMemberUserIds"] = request.VisibleMemberUserIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateScheduledTask"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/createScheduledTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateScheduledTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a scheduled task.
//
// Description:
//
// ## Operation description
//
// - This operation is used to upload files to an enterprise knowledge base.
//
// - You must have the `DEVELOPMENT_KB_MANAGE` permission to call this API operation.
//
// - You must provide the OSS persistent address (`filePath`) of the file when uploading.
//
// - Optional parameters include the public access URL of the file and the original file name to enhance the completeness of file information.
//
// - If `directoryId` is specified, the file is placed in the corresponding enterprise knowledge base directory. Otherwise, the file is bound to the default root directory of the current digital employee by default.
//
// - You can add tags to resources by using `sourceTags` for subsequent management and retrieval.
//
// - This operation initiates a billing item (UNSTRUCTURED_PARSE). Make sure that your account balance is sufficient.
//
// @param request - CreateScheduledTaskRequest
//
// @return CreateScheduledTaskResponse
func (client *Client) CreateScheduledTask(request *CreateScheduledTaskRequest) (_result *CreateScheduledTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateScheduledTaskResponse{}
	_body, _err := client.CreateScheduledTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a knowledge base and knowledge base categories.
//
// Description:
//
// ## Request description
//
// - This API is used to create a new enterprise knowledge base folder under a specified tenant.
//
// - You can set the `parentId` parameter to specify the parent folder of the new folder. If this parameter is not specified, the folder is created as a root folder by default.
//
// - The `path` parameter is optional. If this parameter is not specified, the system automatically calculates the path based on the parent folder.
//
// - Calling this operation requires the corresponding permissions. Multiple authentication methods are supported, including AK, BearerToken, and APP authentication.
//
// - After the folder is created, the related information about the new folder is returned, such as the folder ID and name.
//
// @param request - CreateTenantDirectoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTenantDirectoryResponse
func (client *Client) CreateTenantDirectoryWithOptions(request *CreateTenantDirectoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTenantDirectoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.ParentId) {
		body["parentId"] = request.ParentId
	}

	if !dara.IsNil(request.Path) {
		body["path"] = request.Path
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTenantDirectory"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/createTenantDirectory"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTenantDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a knowledge base and knowledge base categories.
//
// Description:
//
// ## Request description
//
// - This API is used to create a new enterprise knowledge base folder under a specified tenant.
//
// - You can set the `parentId` parameter to specify the parent folder of the new folder. If this parameter is not specified, the folder is created as a root folder by default.
//
// - The `path` parameter is optional. If this parameter is not specified, the system automatically calculates the path based on the parent folder.
//
// - Calling this operation requires the corresponding permissions. Multiple authentication methods are supported, including AK, BearerToken, and APP authentication.
//
// - After the folder is created, the related information about the new folder is returned, such as the folder ID and name.
//
// @param request - CreateTenantDirectoryRequest
//
// @return CreateTenantDirectoryResponse
func (client *Client) CreateTenantDirectory(request *CreateTenantDirectoryRequest) (_result *CreateTenantDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTenantDirectoryResponse{}
	_body, _err := client.CreateTenantDirectoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a user and adds the user to a tenant.
//
// Description:
//
// Creates a user by using OpenAPI.
//
//	Business orchestration:
//
//	1. Parses roleCodes into role_ids (validates against system role enumerations).
//
//	2. Checks whether the user already exists (used to return the isNewUser flag).
//
//	3. Calls UserManagementService.add_tenant_member to create or add the user (the password must be passed by the caller as an RSA ciphertext).
//
//	4. Returns the creation result (including the isNewUser flag).
//
//	Error codes:
//
//	- ERR.User.DeactivatedInTenant: The user is deactivated in the tenant. Use updateUser to restore the user.
//
//	- ERR.User.AlreadyInTenant: The user is already an active member of the tenant.
//
//	- ERR.User.DisplayNameDuplicateInTenant: The display name is duplicate within the tenant.
//
// @param tmpReq - CreateUserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserResponse
func (client *Client) CreateUserWithOptions(tmpReq *CreateUserRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateUserShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RoleCodes) {
		request.RoleCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoleCodes, dara.String("roleCodes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DisplayName) {
		body["displayName"] = request.DisplayName
	}

	if !dara.IsNil(request.PasswordEncrypted) {
		body["passwordEncrypted"] = request.PasswordEncrypted
	}

	if !dara.IsNil(request.RoleCodesShrink) {
		body["roleCodes"] = request.RoleCodesShrink
	}

	if !dara.IsNil(request.SsoProvider) {
		body["ssoProvider"] = request.SsoProvider
	}

	if !dara.IsNil(request.WnAccountId) {
		body["wnAccountId"] = request.WnAccountId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUser"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/createUser"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a user and adds the user to a tenant.
//
// Description:
//
// Creates a user by using OpenAPI.
//
//	Business orchestration:
//
//	1. Parses roleCodes into role_ids (validates against system role enumerations).
//
//	2. Checks whether the user already exists (used to return the isNewUser flag).
//
//	3. Calls UserManagementService.add_tenant_member to create or add the user (the password must be passed by the caller as an RSA ciphertext).
//
//	4. Returns the creation result (including the isNewUser flag).
//
//	Error codes:
//
//	- ERR.User.DeactivatedInTenant: The user is deactivated in the tenant. Use updateUser to restore the user.
//
//	- ERR.User.AlreadyInTenant: The user is already an active member of the tenant.
//
//	- ERR.User.DisplayNameDuplicateInTenant: The display name is duplicate within the tenant.
//
// @param request - CreateUserRequest
//
// @return CreateUserResponse
func (client *Client) CreateUser(request *CreateUserRequest) (_result *CreateUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateUserResponse{}
	_body, _err := client.CreateUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a user group under the tenant to which the authenticated identity belongs.
//
// Description:
//
// WinNexo user management OpenAPI: Creates a user group. The tenant identity is derived from the authentication context.
//
// @param request - CreateUserGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserGroupResponse
func (client *Client) CreateUserGroupWithOptions(request *CreateUserGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.ParentId) {
		body["parentId"] = request.ParentId
	}

	if !dara.IsNil(request.UserGroupName) {
		body["userGroupName"] = request.UserGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUserGroup"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/createUserGroup"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a user group under the tenant to which the authenticated identity belongs.
//
// Description:
//
// WinNexo user management OpenAPI: Creates a user group. The tenant identity is derived from the authentication context.
//
// @param request - CreateUserGroupRequest
//
// @return CreateUserGroupResponse
func (client *Client) CreateUserGroup(request *CreateUserGroupRequest) (_result *CreateUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateUserGroupResponse{}
	_body, _err := client.CreateUserGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a WINNEXO user in the current tenant and assigns roles and user groups to the user.
//
// Description:
//
// ## Request description
//
// - This operation creates a WINNEXO user under a specified tenant and optionally assigns system roles and user groups to the user.
//
// - The `accountId` parameter serves as the logon account for the user and must be unique.
//
// - The `displayName` parameter specifies the display name of the user, which must also be unique within the tenant and cannot exceed 100 characters in length.
//
// - The optional `roleCodes` parameter specifies a list of roles for the user. By default, the `APPLICATION_USER` role is assigned.
//
// - The `userGroupIds` parameter allows you to add up to 100 user group IDs to the new user. Make sure that all specified user groups belong to the same tenant.
//
// - The password must be encrypted by using the RSA-OAEP-SHA256 algorithm and submitted in Base64 format.
//
// - This operation supports calls over HTTPS and requires the request body in JSON format.
//
// - For security authentication, AK, BearerToken, and APP are supported.
//
// @param tmpReq - CreateUserWithGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserWithGroupsResponse
func (client *Client) CreateUserWithGroupsWithOptions(tmpReq *CreateUserWithGroupsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateUserWithGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateUserWithGroupsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RoleCodes) {
		request.RoleCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoleCodes, dara.String("roleCodes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserGroupIds) {
		request.UserGroupIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserGroupIds, dara.String("userGroupIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DisplayName) {
		body["displayName"] = request.DisplayName
	}

	if !dara.IsNil(request.PasswordEncrypted) {
		body["passwordEncrypted"] = request.PasswordEncrypted
	}

	if !dara.IsNil(request.RoleCodesShrink) {
		body["roleCodes"] = request.RoleCodesShrink
	}

	if !dara.IsNil(request.UserGroupIdsShrink) {
		body["userGroupIds"] = request.UserGroupIdsShrink
	}

	if !dara.IsNil(request.WnAccountId) {
		body["wnAccountId"] = request.WnAccountId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUserWithGroups"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/createUserWithGroups"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUserWithGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a WINNEXO user in the current tenant and assigns roles and user groups to the user.
//
// Description:
//
// ## Request description
//
// - This operation creates a WINNEXO user under a specified tenant and optionally assigns system roles and user groups to the user.
//
// - The `accountId` parameter serves as the logon account for the user and must be unique.
//
// - The `displayName` parameter specifies the display name of the user, which must also be unique within the tenant and cannot exceed 100 characters in length.
//
// - The optional `roleCodes` parameter specifies a list of roles for the user. By default, the `APPLICATION_USER` role is assigned.
//
// - The `userGroupIds` parameter allows you to add up to 100 user group IDs to the new user. Make sure that all specified user groups belong to the same tenant.
//
// - The password must be encrypted by using the RSA-OAEP-SHA256 algorithm and submitted in Base64 format.
//
// - This operation supports calls over HTTPS and requires the request body in JSON format.
//
// - For security authentication, AK, BearerToken, and APP are supported.
//
// @param request - CreateUserWithGroupsRequest
//
// @return CreateUserWithGroupsResponse
func (client *Client) CreateUserWithGroups(request *CreateUserWithGroupsRequest) (_result *CreateUserWithGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateUserWithGroupsResponse{}
	_body, _err := client.CreateUserWithGroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a session.
//
// Description:
//
// ## Request description
//
// - This API is used to upload a file to the "My Resources" section of a specified digital employee.
//
// - `source_type` is fixed to `FILE`, `scope` is fixed to `PERSONAL`, and `platform` is fixed to `LOCAL`.
//
// - The file must include an OSS persistent address (`filePath`). Other information such as the public access URL and original file name is optional.
//
// - If the target directory ID (`directoryId`) is not specified, the file is automatically bound to the default root directory of the current digital employee. If specified, ensure that the directory belongs to the caller\\"s personal directory.
//
// - Multiple authentication methods (AK, BearerToken, APP) are supported for security authentication.
//
// - The operation type is write (`write`), and operation logs are recorded for subsequent auditing.
//
// @param request - DeleteChatSessionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteChatSessionResponse
func (client *Client) DeleteChatSessionWithOptions(request *DeleteChatSessionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteChatSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SessionId) {
		query["sessionId"] = request.SessionId
	}

	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteChatSession"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/deleteChatSession"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteChatSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a session.
//
// Description:
//
// ## Request description
//
// - This API is used to upload a file to the "My Resources" section of a specified digital employee.
//
// - `source_type` is fixed to `FILE`, `scope` is fixed to `PERSONAL`, and `platform` is fixed to `LOCAL`.
//
// - The file must include an OSS persistent address (`filePath`). Other information such as the public access URL and original file name is optional.
//
// - If the target directory ID (`directoryId`) is not specified, the file is automatically bound to the default root directory of the current digital employee. If specified, ensure that the directory belongs to the caller\\"s personal directory.
//
// - Multiple authentication methods (AK, BearerToken, APP) are supported for security authentication.
//
// - The operation type is write (`write`), and operation logs are recorded for subsequent auditing.
//
// @param request - DeleteChatSessionRequest
//
// @return DeleteChatSessionResponse
func (client *Client) DeleteChatSession(request *DeleteChatSessionRequest) (_result *DeleteChatSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteChatSessionResponse{}
	_body, _err := client.DeleteChatSessionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specified source.
//
// Description:
//
// ## Operation description
//
// - `tenantId` is derived from the authenticated identity only. Any value passed by the caller is ignored.
//
// - `sourceId` is passed through the request body. The registration path is the flat URI `/openapi/deleteSource` and does not contain a `{sourceId}` path template. Do not append the resource ID as a path segment. The gateway performs exact routing based on the flat URI and returns `InvalidAction.NotFound` if the path does not match.
//
// - Deletion is irreversible. The parsing results and bindings associated with the resource are invalidated.
//
// @param request - DeleteSourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSourceResponse
func (client *Client) DeleteSourceWithOptions(request *DeleteSourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SourceId) {
		body["sourceId"] = request.SourceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSource"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/deleteSource"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified source.
//
// Description:
//
// ## Operation description
//
// - `tenantId` is derived from the authenticated identity only. Any value passed by the caller is ignored.
//
// - `sourceId` is passed through the request body. The registration path is the flat URI `/openapi/deleteSource` and does not contain a `{sourceId}` path template. Do not append the resource ID as a path segment. The gateway performs exact routing based on the flat URI and returns `InvalidAction.NotFound` if the path does not match.
//
// - Deletion is irreversible. The parsing results and bindings associated with the resource are invalidated.
//
// @param request - DeleteSourceRequest
//
// @return DeleteSourceResponse
func (client *Client) DeleteSource(request *DeleteSourceRequest) (_result *DeleteSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSourceResponse{}
	_body, _err := client.DeleteSourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an enterprise knowledge base and its subdirectories.
//
// Description:
//
// ## Request description
//
// - This API allows you to delete a specific enterprise knowledge base directory.
//
// - Set the `deleteMode` parameter to select different deletion strategies, including reject deletion (reject), recursive deletion (recursive), or move the directory to the root directory (move_to_root).
//
// - If `deleteMode` is not provided, the default behavior is to reject deletion.
//
// - The enterprise directory boundary is validated before the deletion operation.
//
// @param request - DeleteTenantDirectoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTenantDirectoryResponse
func (client *Client) DeleteTenantDirectoryWithOptions(request *DeleteTenantDirectoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteTenantDirectoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeleteMode) {
		body["deleteMode"] = request.DeleteMode
	}

	if !dara.IsNil(request.DirectoryId) {
		body["directoryId"] = request.DirectoryId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTenantDirectory"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/deleteTenantDirectory"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTenantDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an enterprise knowledge base and its subdirectories.
//
// Description:
//
// ## Request description
//
// - This API allows you to delete a specific enterprise knowledge base directory.
//
// - Set the `deleteMode` parameter to select different deletion strategies, including reject deletion (reject), recursive deletion (recursive), or move the directory to the root directory (move_to_root).
//
// - If `deleteMode` is not provided, the default behavior is to reject deletion.
//
// - The enterprise directory boundary is validated before the deletion operation.
//
// @param request - DeleteTenantDirectoryRequest
//
// @return DeleteTenantDirectoryResponse
func (client *Client) DeleteTenantDirectory(request *DeleteTenantDirectoryRequest) (_result *DeleteTenantDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTenantDirectoryResponse{}
	_body, _err := client.DeleteTenantDirectoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables an API token.
//
// Description:
//
// Disables the INSTANCE token of a user.
//
//	Business logic:
//
//	1. Retrieves user_id from identity (caller_type=user is enforced).
//
//	2. Constructs an AuthContext and delegates permission verification to UserTokenAuthorizedService.
//
//	3. Calls disable_token (ACTIVE → INACTIVE).
//
//	4. Returns disabled=True.
//
//	Idempotence: If no ACTIVE token exists, deactivate_all affects 0 rows and does not return an error.
//
// @param request - DisableTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableTokenResponse
func (client *Client) DisableTokenWithOptions(request *DisableTokenRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DisableTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.WnUserId) {
		body["wnUserId"] = request.WnUserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableToken"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/disableToken"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables an API token.
//
// Description:
//
// Disables the INSTANCE token of a user.
//
//	Business logic:
//
//	1. Retrieves user_id from identity (caller_type=user is enforced).
//
//	2. Constructs an AuthContext and delegates permission verification to UserTokenAuthorizedService.
//
//	3. Calls disable_token (ACTIVE → INACTIVE).
//
//	4. Returns disabled=True.
//
//	Idempotence: If no ACTIVE token exists, deactivate_all affects 0 rows and does not return an error.
//
// @param request - DisableTokenRequest
//
// @return DisableTokenResponse
func (client *Client) DisableToken(request *DisableTokenRequest) (_result *DisableTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DisableTokenResponse{}
	_body, _err := client.DisableTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables an API token.
//
// Description:
//
// Enables the INSTANCE token for a user (idempotent).
//
//	Business logic:
//
//	1. Retrieves user_id from identity (caller_type=user is required).
//
//	2. Constructs an AuthContext and delegates permission verification to UserTokenAuthorizedService.
//
//	3. Calls enable_token:
//
//	   - If an ACTIVE token exists, returns idempotently (only the masked value is returned, and the plaintext is not issued again).
//
//	   - If an INACTIVE token exists, reactivates it (returns the plaintext).
//
//	   - If no token exists, creates one (returns the plaintext).
//
//	Security constraint: The token plaintext is returned only once when the token is first enabled. Subsequent idempotent calls do not return the plaintext.
//
// @param request - EnableTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableTokenResponse
func (client *Client) EnableTokenWithOptions(request *EnableTokenRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EnableTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.WnUserId) {
		body["wnUserId"] = request.WnUserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableToken"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/enableToken"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables an API token.
//
// Description:
//
// Enables the INSTANCE token for a user (idempotent).
//
//	Business logic:
//
//	1. Retrieves user_id from identity (caller_type=user is required).
//
//	2. Constructs an AuthContext and delegates permission verification to UserTokenAuthorizedService.
//
//	3. Calls enable_token:
//
//	   - If an ACTIVE token exists, returns idempotently (only the masked value is returned, and the plaintext is not issued again).
//
//	   - If an INACTIVE token exists, reactivates it (returns the plaintext).
//
//	   - If no token exists, creates one (returns the plaintext).
//
//	Security constraint: The token plaintext is returned only once when the token is first enabled. Subsequent idempotent calls do not return the plaintext.
//
// @param request - EnableTokenRequest
//
// @return EnableTokenResponse
func (client *Client) EnableToken(request *EnableTokenRequest) (_result *EnableTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EnableTokenResponse{}
	_body, _err := client.EnableTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves session details.
//
// Description:
//
// ## Request description
//
// - This API uploads a file to the "My Resources" section of a specified digital employee.
//
// - `source_type` is fixed to `FILE`, `scope` is fixed to `PERSONAL`, and `platform` is fixed to `LOCAL`.
//
// - The file must include an OSS persistent address (`filePath`). Other information such as the public access URL and original file name is optional.
//
// - If no target folder ID (`directoryId`) is specified, the file is automatically attached to the default root folder of the current digital employee. If specified, ensure that the folder belongs to the invoker\\"s personal folder.
//
// - Multiple authentication methods (AK, BearerToken, APP) are supported to authenticate requests.
//
// - The operation type is write (`write`), and operation logs are recorded for subsequent auditing.
//
// @param request - GetChatSessionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChatSessionResponse
func (client *Client) GetChatSessionWithOptions(request *GetChatSessionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetChatSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Limit) {
		query["limit"] = request.Limit
	}

	if !dara.IsNil(request.SessionId) {
		query["sessionId"] = request.SessionId
	}

	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetChatSession"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/getChatSession"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetChatSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves session details.
//
// Description:
//
// ## Request description
//
// - This API uploads a file to the "My Resources" section of a specified digital employee.
//
// - `source_type` is fixed to `FILE`, `scope` is fixed to `PERSONAL`, and `platform` is fixed to `LOCAL`.
//
// - The file must include an OSS persistent address (`filePath`). Other information such as the public access URL and original file name is optional.
//
// - If no target folder ID (`directoryId`) is specified, the file is automatically attached to the default root folder of the current digital employee. If specified, ensure that the folder belongs to the invoker\\"s personal folder.
//
// - Multiple authentication methods (AK, BearerToken, APP) are supported to authenticate requests.
//
// - The operation type is write (`write`), and operation logs are recorded for subsequent auditing.
//
// @param request - GetChatSessionRequest
//
// @return GetChatSessionResponse
func (client *Client) GetChatSession(request *GetChatSessionRequest) (_result *GetChatSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetChatSessionResponse{}
	_body, _err := client.GetChatSessionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the active Graph Schema readable by the current user.
//
// Description:
//
// Reads the active schema_content and securely trims it based on the token user\\"s semantic resource READ permissions.
//
// @param request - GetGraphSchemaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGraphSchemaResponse
func (client *Client) GetGraphSchemaWithOptions(request *GetGraphSchemaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetGraphSchemaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.GraphName) {
		body["graphName"] = request.GraphName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGraphSchema"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/getGraphSchema"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGraphSchemaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the active Graph Schema readable by the current user.
//
// Description:
//
// Reads the active schema_content and securely trims it based on the token user\\"s semantic resource READ permissions.
//
// @param request - GetGraphSchemaRequest
//
// @return GetGraphSchemaResponse
func (client *Client) GetGraphSchema(request *GetGraphSchemaRequest) (_result *GetGraphSchemaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetGraphSchemaResponse{}
	_body, _err := client.GetGraphSchemaWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the expiration time of the most recently created standard package instance for a tenant.
//
// Description:
//
// ## Operation description
//
// - This API operation queries the expiration time of the most recently created standard package instance for a specified tenant.
//
// - If no standard package instance is found, the `found` field returns `False`.
//
// - You can use the `tenantId` parameter to specify the tenant ID. By default, the tenant ID of the caller is used.
//
// - The request method is POST and must be called over HTTPS.
//
// - Valid authentication information (such as AK, BearerToken, or APP) is required to complete the request.
//
// @param request - GetInstanceExpireTimeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceExpireTimeResponse
func (client *Client) GetInstanceExpireTimeWithOptions(request *GetInstanceExpireTimeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetInstanceExpireTimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceExpireTime"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/getInstanceExpireTime"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceExpireTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the expiration time of the most recently created standard package instance for a tenant.
//
// Description:
//
// ## Operation description
//
// - This API operation queries the expiration time of the most recently created standard package instance for a specified tenant.
//
// - If no standard package instance is found, the `found` field returns `False`.
//
// - You can use the `tenantId` parameter to specify the tenant ID. By default, the tenant ID of the caller is used.
//
// - The request method is POST and must be called over HTTPS.
//
// - Valid authentication information (such as AK, BearerToken, or APP) is required to complete the request.
//
// @param request - GetInstanceExpireTimeRequest
//
// @return GetInstanceExpireTimeResponse
func (client *Client) GetInstanceExpireTime(request *GetInstanceExpireTimeRequest) (_result *GetInstanceExpireTimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceExpireTimeResponse{}
	_body, _err := client.GetInstanceExpireTimeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specified knowledge item in the enterprise knowledge base.
//
// Description:
//
// ## Operation description
//
// - This API operation retrieves the details of a specific knowledge item in the enterprise knowledge base.
//
// - Calling this operation requires the `DEVELOPMENT_KB_VIEW` feature permission.
//
// - Knowledge details include but are not limited to the knowledge type, name, and description.
//
// - The `sourceId` parameter is required to identify the knowledge item to query.
//
// - `tenantId` is an optional parameter. The tenant ID of the caller is used by default.
//
// - Authentication is supported through `AK`, `BearerToken`, or `APP` methods.
//
// - Security constraint: `tenant_id` and `user_id` can only be derived from the authenticated identity.
//
// @param request - GetKnowledgeBaseSourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetKnowledgeBaseSourceResponse
func (client *Client) GetKnowledgeBaseSourceWithOptions(request *GetKnowledgeBaseSourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetKnowledgeBaseSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SourceId) {
		body["sourceId"] = request.SourceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetKnowledgeBaseSource"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/getKnowledgeBaseSource"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetKnowledgeBaseSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified knowledge item in the enterprise knowledge base.
//
// Description:
//
// ## Operation description
//
// - This API operation retrieves the details of a specific knowledge item in the enterprise knowledge base.
//
// - Calling this operation requires the `DEVELOPMENT_KB_VIEW` feature permission.
//
// - Knowledge details include but are not limited to the knowledge type, name, and description.
//
// - The `sourceId` parameter is required to identify the knowledge item to query.
//
// - `tenantId` is an optional parameter. The tenant ID of the caller is used by default.
//
// - Authentication is supported through `AK`, `BearerToken`, or `APP` methods.
//
// - Security constraint: `tenant_id` and `user_id` can only be derived from the authenticated identity.
//
// @param request - GetKnowledgeBaseSourceRequest
//
// @return GetKnowledgeBaseSourceResponse
func (client *Client) GetKnowledgeBaseSource(request *GetKnowledgeBaseSourceRequest) (_result *GetKnowledgeBaseSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetKnowledgeBaseSourceResponse{}
	_body, _err := client.GetKnowledgeBaseSourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the execution details of a scheduled task.
//
// Description:
//
// ## Operation description
//
// - This operation uploads a file to an enterprise knowledge base.
//
// - You must have the `DEVELOPMENT_KB_MANAGE` permission to call this API operation.
//
// - You must provide the OSS persistent address (`filePath`) of the file when uploading.
//
// - Optional parameters include the public access URL and original file name to enhance the completeness of file information.
//
// - If `directoryId` is specified, the file is placed in the corresponding enterprise knowledge base directory. Otherwise, the file is bound to the default root directory of the current digital employee.
//
// - You can add tags to the resource by using `sourceTags` for subsequent management and retrieval.
//
// - This operation initiates a billing item (UNSTRUCTURED_PARSE). Make sure your account balance is sufficient.
//
// @param request - GetScheduledTaskExecutionDetailRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetScheduledTaskExecutionDetailResponse
func (client *Client) GetScheduledTaskExecutionDetailWithOptions(request *GetScheduledTaskExecutionDetailRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetScheduledTaskExecutionDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExecutionId) {
		query["executionId"] = request.ExecutionId
	}

	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetScheduledTaskExecutionDetail"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/getScheduledTaskExecutionDetail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetScheduledTaskExecutionDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the execution details of a scheduled task.
//
// Description:
//
// ## Operation description
//
// - This operation uploads a file to an enterprise knowledge base.
//
// - You must have the `DEVELOPMENT_KB_MANAGE` permission to call this API operation.
//
// - You must provide the OSS persistent address (`filePath`) of the file when uploading.
//
// - Optional parameters include the public access URL and original file name to enhance the completeness of file information.
//
// - If `directoryId` is specified, the file is placed in the corresponding enterprise knowledge base directory. Otherwise, the file is bound to the default root directory of the current digital employee.
//
// - You can add tags to the resource by using `sourceTags` for subsequent management and retrieval.
//
// - This operation initiates a billing item (UNSTRUCTURED_PARSE). Make sure your account balance is sufficient.
//
// @param request - GetScheduledTaskExecutionDetailRequest
//
// @return GetScheduledTaskExecutionDetailResponse
func (client *Client) GetScheduledTaskExecutionDetail(request *GetScheduledTaskExecutionDetailRequest) (_result *GetScheduledTaskExecutionDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetScheduledTaskExecutionDetailResponse{}
	_body, _err := client.GetScheduledTaskExecutionDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves execution records of scheduled tasks.
//
// Description:
//
// ## Operation description
//
// - This operation uploads a file to the enterprise knowledge base.
//
// - The `DEVELOPMENT_KB_MANAGE` feature permission is required to call this API.
//
// - You must provide the OSS persistent address (`filePath`) of the file when uploading.
//
// - Optional parameters include the public access URL and original file name to enhance the completeness of file information.
//
// - If `directoryId` is specified, the file is placed in the corresponding enterprise knowledge base directory. Otherwise, the file is bound to the default root directory of the current digital employee.
//
// - You can add tags to the resource by using `sourceTags` for subsequent management and retrieval.
//
// - This operation initiates a billing item (UNSTRUCTURED_PARSE). Ensure that your account balance is sufficient.
//
// @param request - GetScheduledTaskExecutionRecordsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetScheduledTaskExecutionRecordsResponse
func (client *Client) GetScheduledTaskExecutionRecordsWithOptions(request *GetScheduledTaskExecutionRecordsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetScheduledTaskExecutionRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CollaborationGroupId) {
		query["collaborationGroupId"] = request.CollaborationGroupId
	}

	if !dara.IsNil(request.InitiatorUserId) {
		query["initiatorUserId"] = request.InitiatorUserId
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	if !dara.IsNil(request.TaskId) {
		query["taskId"] = request.TaskId
	}

	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetScheduledTaskExecutionRecords"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/getScheduledTaskExecutionRecords"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetScheduledTaskExecutionRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves execution records of scheduled tasks.
//
// Description:
//
// ## Operation description
//
// - This operation uploads a file to the enterprise knowledge base.
//
// - The `DEVELOPMENT_KB_MANAGE` feature permission is required to call this API.
//
// - You must provide the OSS persistent address (`filePath`) of the file when uploading.
//
// - Optional parameters include the public access URL and original file name to enhance the completeness of file information.
//
// - If `directoryId` is specified, the file is placed in the corresponding enterprise knowledge base directory. Otherwise, the file is bound to the default root directory of the current digital employee.
//
// - You can add tags to the resource by using `sourceTags` for subsequent management and retrieval.
//
// - This operation initiates a billing item (UNSTRUCTURED_PARSE). Ensure that your account balance is sufficient.
//
// @param request - GetScheduledTaskExecutionRecordsRequest
//
// @return GetScheduledTaskExecutionRecordsResponse
func (client *Client) GetScheduledTaskExecutionRecords(request *GetScheduledTaskExecutionRecordsRequest) (_result *GetScheduledTaskExecutionRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetScheduledTaskExecutionRecordsResponse{}
	_body, _err := client.GetScheduledTaskExecutionRecordsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the push configuration options for scheduled tasks.
//
// Description:
//
// Queries the channels and methods available to the current user for scheduled task push notifications.
//
// @param request - GetScheduledTaskPushOptionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetScheduledTaskPushOptionsResponse
func (client *Client) GetScheduledTaskPushOptionsWithOptions(request *GetScheduledTaskPushOptionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetScheduledTaskPushOptionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CollaborationGroupId) {
		body["collaborationGroupId"] = request.CollaborationGroupId
	}

	if !dara.IsNil(request.DigitalEmployeeName) {
		body["digitalEmployeeName"] = request.DigitalEmployeeName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetScheduledTaskPushOptions"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/getScheduledTaskPushOptions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetScheduledTaskPushOptionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the push configuration options for scheduled tasks.
//
// Description:
//
// Queries the channels and methods available to the current user for scheduled task push notifications.
//
// @param request - GetScheduledTaskPushOptionsRequest
//
// @return GetScheduledTaskPushOptionsResponse
func (client *Client) GetScheduledTaskPushOptions(request *GetScheduledTaskPushOptionsRequest) (_result *GetScheduledTaskPushOptionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetScheduledTaskPushOptionsResponse{}
	_body, _err := client.GetScheduledTaskPushOptionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetScheduledTaskUnderstandDetail is deprecated
//
// Summary:
//
// Retrieves the details of scheduled task understanding.
//
// Description:
//
// ## Operation description
//
// - This operation uploads a file to the enterprise knowledge base.
//
// - The `DEVELOPMENT_KB_MANAGE` feature permission is required to call this API.
//
// - The OSS persistent address (`filePath`) of the file must be provided during upload.
//
// - Optional parameters include the public access URL and original file name to enhance the completeness of file information.
//
// - If `directoryId` is specified, the file is placed in the corresponding enterprise knowledge base directory. Otherwise, the file is bound to the default root directory of the current digital employee.
//
// - You can add tags to the resource by using `sourceTags` for subsequent management and retrieval.
//
// - This operation initiates a billing item (UNSTRUCTURED_PARSE). Ensure that your account balance is sufficient.
//
// @param tmpReq - GetScheduledTaskUnderstandDetailRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetScheduledTaskUnderstandDetailResponse
func (client *Client) GetScheduledTaskUnderstandDetailWithOptions(tmpReq *GetScheduledTaskUnderstandDetailRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetScheduledTaskUnderstandDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetScheduledTaskUnderstandDetailShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DigitalEmployeeName) {
		request.DigitalEmployeeNameShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DigitalEmployeeName, dara.String("digitalEmployeeName"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Segments) {
		request.SegmentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Segments, dara.String("segments"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CollaborationGroupId) {
		query["collaborationGroupId"] = request.CollaborationGroupId
	}

	if !dara.IsNil(request.DigitalEmployeeNameShrink) {
		query["digitalEmployeeName"] = request.DigitalEmployeeNameShrink
	}

	if !dara.IsNil(request.SegmentsShrink) {
		query["segments"] = request.SegmentsShrink
	}

	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	if !dara.IsNil(request.UserInput) {
		query["userInput"] = request.UserInput
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetScheduledTaskUnderstandDetail"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/getScheduledTaskUnderstandDetail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetScheduledTaskUnderstandDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetScheduledTaskUnderstandDetail is deprecated
//
// Summary:
//
// Retrieves the details of scheduled task understanding.
//
// Description:
//
// ## Operation description
//
// - This operation uploads a file to the enterprise knowledge base.
//
// - The `DEVELOPMENT_KB_MANAGE` feature permission is required to call this API.
//
// - The OSS persistent address (`filePath`) of the file must be provided during upload.
//
// - Optional parameters include the public access URL and original file name to enhance the completeness of file information.
//
// - If `directoryId` is specified, the file is placed in the corresponding enterprise knowledge base directory. Otherwise, the file is bound to the default root directory of the current digital employee.
//
// - You can add tags to the resource by using `sourceTags` for subsequent management and retrieval.
//
// - This operation initiates a billing item (UNSTRUCTURED_PARSE). Ensure that your account balance is sufficient.
//
// @param request - GetScheduledTaskUnderstandDetailRequest
//
// @return GetScheduledTaskUnderstandDetailResponse
// Deprecated
func (client *Client) GetScheduledTaskUnderstandDetail(request *GetScheduledTaskUnderstandDetailRequest) (_result *GetScheduledTaskUnderstandDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetScheduledTaskUnderstandDetailResponse{}
	_body, _err := client.GetScheduledTaskUnderstandDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves skill details.
//
// Description:
//
// ## Request description
//
// Queries skill details by SkillCode or SkillName, including metadata, input parameter schema, and SKILL.md summary.
//
// - **TenantId**: Optional common parameter passed through by the gateway to the backend header. If not specified, the default tenant of the current caller is used.
//
// - **SkillCode**: Mutually exclusive with SkillName. If both are specified, SkillCode takes precedence.
//
// - **SkillName**: Mutually exclusive with SkillCode. If the name is not unique within the tenant, `ERR.SkillHub.SkillNameAmbiguous` is returned.
//
// - **ViewMode**: Optional. Valid values: `draft` (draft/editing view) or `published` (published view, default).
//
// - **IncludeSkillFiles**: Optional. Specifies whether to return the complete skill file tree (SKILL.md / scripts / templates). Default value: `false`.
//
// @param request - GetSkillRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSkillResponse
func (client *Client) GetSkillWithOptions(request *GetSkillRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.IncludeSkillFiles) {
		body["includeSkillFiles"] = request.IncludeSkillFiles
	}

	if !dara.IsNil(request.SkillCode) {
		body["skillCode"] = request.SkillCode
	}

	if !dara.IsNil(request.SkillName) {
		body["skillName"] = request.SkillName
	}

	if !dara.IsNil(request.ViewMode) {
		body["viewMode"] = request.ViewMode
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSkill"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/getSkill"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSkillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves skill details.
//
// Description:
//
// ## Request description
//
// Queries skill details by SkillCode or SkillName, including metadata, input parameter schema, and SKILL.md summary.
//
// - **TenantId**: Optional common parameter passed through by the gateway to the backend header. If not specified, the default tenant of the current caller is used.
//
// - **SkillCode**: Mutually exclusive with SkillName. If both are specified, SkillCode takes precedence.
//
// - **SkillName**: Mutually exclusive with SkillCode. If the name is not unique within the tenant, `ERR.SkillHub.SkillNameAmbiguous` is returned.
//
// - **ViewMode**: Optional. Valid values: `draft` (draft/editing view) or `published` (published view, default).
//
// - **IncludeSkillFiles**: Optional. Specifies whether to return the complete skill file tree (SKILL.md / scripts / templates). Default value: `false`.
//
// @param request - GetSkillRequest
//
// @return GetSkillResponse
func (client *Client) GetSkill(request *GetSkillRequest) (_result *GetSkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSkillResponse{}
	_body, _err := client.GetSkillWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the execution result of a skill.
//
// Description:
//
// ## Request description
//
// Queries the current status and result of an asynchronous task by `RunId`.
//
// - **State machine**: Running (PENDING/RUNNING) → Succeeded / Failed / Cancelled
//
// - **TenantId**: An optional common parameter passed through by the gateway. The backend verifies that the RunId belongs to the current tenant. Otherwise, `ERR.SkillHub.RunNotFound` is returned to avoid exposing existence information.
//
// - **IncludeLogs**: Optional. Specifies whether to return execution logs. Default value: `false`.
//
// When execution succeeds, `Result.Content[]` is an MCP-style Content block array (Text / File / Image).
//
// @param request - GetSkillRunRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSkillRunResponse
func (client *Client) GetSkillRunWithOptions(request *GetSkillRunRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSkillRunResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.IncludeLogs) {
		body["includeLogs"] = request.IncludeLogs
	}

	if !dara.IsNil(request.RunId) {
		body["runId"] = request.RunId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSkillRun"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/getSkillRun"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSkillRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution result of a skill.
//
// Description:
//
// ## Request description
//
// Queries the current status and result of an asynchronous task by `RunId`.
//
// - **State machine**: Running (PENDING/RUNNING) → Succeeded / Failed / Cancelled
//
// - **TenantId**: An optional common parameter passed through by the gateway. The backend verifies that the RunId belongs to the current tenant. Otherwise, `ERR.SkillHub.RunNotFound` is returned to avoid exposing existence information.
//
// - **IncludeLogs**: Optional. Specifies whether to return execution logs. Default value: `false`.
//
// When execution succeeds, `Result.Content[]` is an MCP-style Content block array (Text / File / Image).
//
// @param request - GetSkillRunRequest
//
// @return GetSkillRunResponse
func (client *Client) GetSkillRun(request *GetSkillRunRequest) (_result *GetSkillRunResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSkillRunResponse{}
	_body, _err := client.GetSkillRunWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specified resource (knowledge), with support for returning large detail fields on demand.
//
// Description:
//
// ## Operation description
//
// - `tenant_id` is derived from the authenticated identity only. Any value passed in the body is ignored.
//
// - Response parameters do not expose audit fields such as `creator` or `modifier`. The `unstructured_docs[ ].content` field is not returned by default to avoid large responses.
//
// - Set the `includeDetails` parameter to `True` to retrieve additional details including `settings`, `notes`, `structuredTables`, and `unstructuredDocs`.
//
// @param request - GetSourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSourceResponse
func (client *Client) GetSourceWithOptions(request *GetSourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.IncludeDetails) {
		body["includeDetails"] = request.IncludeDetails
	}

	if !dara.IsNil(request.SourceId) {
		body["sourceId"] = request.SourceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSource"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/getSource"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified resource (knowledge), with support for returning large detail fields on demand.
//
// Description:
//
// ## Operation description
//
// - `tenant_id` is derived from the authenticated identity only. Any value passed in the body is ignored.
//
// - Response parameters do not expose audit fields such as `creator` or `modifier`. The `unstructured_docs[ ].content` field is not returned by default to avoid large responses.
//
// - Set the `includeDetails` parameter to `True` to retrieve additional details including `settings`, `notes`, `structuredTables`, and `unstructuredDocs`.
//
// @param request - GetSourceRequest
//
// @return GetSourceResponse
func (client *Client) GetSource(request *GetSourceRequest) (_result *GetSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSourceResponse{}
	_body, _err := client.GetSourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates a signed URL for directly uploading files to OSS.
//
// Description:
//
// ## Operation description
//
// This API allows callers to obtain a signed URL for directly uploading files to Alibaba Cloud Object Storage Service (OSS) based on the provided file name and other information. With this URL, users can upload files directly to the specified OSS location without routing through an intermediate server, which improves efficiency and security.
//
// - **Security constraint**: `tenant_id`/`user_id` are derived only from the authenticated identity. Values provided in the request body are ignored.
//
// - **Default value**: If the `expires` parameter is not specified, the default expiration time is 3600 seconds (1 hour).
//
// - **Content-Type**: If `contentType` is not provided, the system attempts to automatically infer the file type.
//
// - **Scope**: The `scope` parameter defines whether the data source belongs to a personal or enterprise knowledge base. In most cases, this does not need to be set.
//
// @param request - GetSourceUploadSignatureRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSourceUploadSignatureResponse
func (client *Client) GetSourceUploadSignatureWithOptions(request *GetSourceUploadSignatureRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSourceUploadSignatureResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ContentType) {
		body["contentType"] = request.ContentType
	}

	if !dara.IsNil(request.Expires) {
		body["expires"] = request.Expires
	}

	if !dara.IsNil(request.Filename) {
		body["filename"] = request.Filename
	}

	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	if !dara.IsNil(request.Scope) {
		body["scope"] = request.Scope
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSourceUploadSignature"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/getSourceUploadSignature"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSourceUploadSignatureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a signed URL for directly uploading files to OSS.
//
// Description:
//
// ## Operation description
//
// This API allows callers to obtain a signed URL for directly uploading files to Alibaba Cloud Object Storage Service (OSS) based on the provided file name and other information. With this URL, users can upload files directly to the specified OSS location without routing through an intermediate server, which improves efficiency and security.
//
// - **Security constraint**: `tenant_id`/`user_id` are derived only from the authenticated identity. Values provided in the request body are ignored.
//
// - **Default value**: If the `expires` parameter is not specified, the default expiration time is 3600 seconds (1 hour).
//
// - **Content-Type**: If `contentType` is not provided, the system attempts to automatically infer the file type.
//
// - **Scope**: The `scope` parameter defines whether the data source belongs to a personal or enterprise knowledge base. In most cases, this does not need to be set.
//
// @param request - GetSourceUploadSignatureRequest
//
// @return GetSourceUploadSignatureResponse
func (client *Client) GetSourceUploadSignature(request *GetSourceUploadSignatureRequest) (_result *GetSourceUploadSignatureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSourceUploadSignatureResponse{}
	_body, _err := client.GetSourceUploadSignatureWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves an API token and ensures that it is active.
//
// Description:
//
// Retrieves the INSTANCE token for a user and ensures that it is in an active state (idempotent).
//
//	Business logic:
//
//	1. Obtains user_id from identity (caller_type=user is enforced).
//
//	2. Constructs an AuthContext and delegates permission verification to UserTokenAuthorizedService.
//
//	3. Calls ensure_active_token:
//
//	   - If an ACTIVE token exists, returns the token in plaintext as-is (no reset, no key rotation).
//
//	   - If an INACTIVE token exists, automatically re-enables it and returns the plaintext.
//
//	   - If no token exists (or only expired RESET records exist), creates a new token and returns the plaintext.
//
//	Difference from EnableToken: When an ACTIVE token already exists, EnableToken returns only the masked value. This operation guarantees that a usable plaintext credential is returned without destroying the existing token.
//
// @param request - GetTokenEnsureEnableRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTokenEnsureEnableResponse
func (client *Client) GetTokenEnsureEnableWithOptions(request *GetTokenEnsureEnableRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTokenEnsureEnableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.WnUserId) {
		body["wnUserId"] = request.WnUserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTokenEnsureEnable"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/getTokenEnsureEnable"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTokenEnsureEnableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves an API token and ensures that it is active.
//
// Description:
//
// Retrieves the INSTANCE token for a user and ensures that it is in an active state (idempotent).
//
//	Business logic:
//
//	1. Obtains user_id from identity (caller_type=user is enforced).
//
//	2. Constructs an AuthContext and delegates permission verification to UserTokenAuthorizedService.
//
//	3. Calls ensure_active_token:
//
//	   - If an ACTIVE token exists, returns the token in plaintext as-is (no reset, no key rotation).
//
//	   - If an INACTIVE token exists, automatically re-enables it and returns the plaintext.
//
//	   - If no token exists (or only expired RESET records exist), creates a new token and returns the plaintext.
//
//	Difference from EnableToken: When an ACTIVE token already exists, EnableToken returns only the masked value. This operation guarantees that a usable plaintext credential is returned without destroying the existing token.
//
// @param request - GetTokenEnsureEnableRequest
//
// @return GetTokenEnsureEnableResponse
func (client *Client) GetTokenEnsureEnable(request *GetTokenEnsureEnableRequest) (_result *GetTokenEnsureEnableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTokenEnsureEnableResponse{}
	_body, _err := client.GetTokenEnsureEnableWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the token status of a user.
//
// Description:
//
// Queries the INSTANCE token status of a user.
//
//	Business logic:
//
//	1. Retrieves user_id from identity (caller_type=user is required).
//
//	2. Constructs an AuthContext and delegates permission verification to UserTokenAuthorizedService.
//
//	3. Queries the ACTIVE INSTANCE token.
//
//	4. If the token exists, returns enabled=True with the masked value and creation time.
//
//	5. If the token does not exist, returns enabled=False.
//
// @param request - GetTokenInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTokenInfoResponse
func (client *Client) GetTokenInfoWithOptions(request *GetTokenInfoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTokenInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.WnUserId) {
		body["wnUserId"] = request.WnUserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTokenInfo"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/getTokenInfo"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTokenInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the token status of a user.
//
// Description:
//
// Queries the INSTANCE token status of a user.
//
//	Business logic:
//
//	1. Retrieves user_id from identity (caller_type=user is required).
//
//	2. Constructs an AuthContext and delegates permission verification to UserTokenAuthorizedService.
//
//	3. Queries the ACTIVE INSTANCE token.
//
//	4. If the token exists, returns enabled=True with the masked value and creation time.
//
//	5. If the token does not exist, returns enabled=False.
//
// @param request - GetTokenInfoRequest
//
// @return GetTokenInfoResponse
func (client *Client) GetTokenInfo(request *GetTokenInfoRequest) (_result *GetTokenInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTokenInfoResponse{}
	_body, _err := client.GetTokenInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries user details.
//
// Description:
//
// Queries user details through OpenAPI.
//
//	Business orchestration:
//
//	1. Locate the user by wnUserId or accountId.
//
//	2. Query the user mapping information in the current tenant (status, join time, and last logon time).
//
//	3. Query the role list of the user in the current tenant.
//
//	4. Query the user group list of the user in the current tenant.
//
//	5. Assemble the response.
//
//	Error codes:
//
//	- ERR.User.NotFound: The user does not exist.
//
//	- ERR.User.NotInTenant: The user does not belong to the current tenant.
//
// @param request - GetUserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserResponse
func (client *Client) GetUserWithOptions(request *GetUserRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WnAccountId) {
		query["wnAccountId"] = request.WnAccountId
	}

	if !dara.IsNil(request.WnUserId) {
		query["wnUserId"] = request.WnUserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUser"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/getUser"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries user details.
//
// Description:
//
// Queries user details through OpenAPI.
//
//	Business orchestration:
//
//	1. Locate the user by wnUserId or accountId.
//
//	2. Query the user mapping information in the current tenant (status, join time, and last logon time).
//
//	3. Query the role list of the user in the current tenant.
//
//	4. Query the user group list of the user in the current tenant.
//
//	5. Assemble the response.
//
//	Error codes:
//
//	- ERR.User.NotFound: The user does not exist.
//
//	- ERR.User.NotInTenant: The user does not belong to the current tenant.
//
// @param request - GetUserRequest
//
// @return GetUserResponse
func (client *Client) GetUser(request *GetUserRequest) (_result *GetUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the real-time credit consumption, limit, and remaining balance of the current logged-on user.
//
// Description:
//
// ## Request description
//
// - This API is used to retrieve the credit usage details of the current logged-on user, including the credit limit, consumed credits, and remaining credits.
//
// - Data is sourced from a real-time Redis cache, ensuring information immediacy.
//
// - You can specify a tenant ID to query the credit usage of a user under a specific tenant. By default, the caller\\"s default tenant is used.
//
// - You can optionally provide a `RequestId` as a request identifier, but this is not required.
//
// @param request - GetUserCreditUsageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserCreditUsageResponse
func (client *Client) GetUserCreditUsageWithOptions(request *GetUserCreditUsageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetUserCreditUsageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserCreditUsage"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/getUserCreditUsage"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserCreditUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the real-time credit consumption, limit, and remaining balance of the current logged-on user.
//
// Description:
//
// ## Request description
//
// - This API is used to retrieve the credit usage details of the current logged-on user, including the credit limit, consumed credits, and remaining credits.
//
// - Data is sourced from a real-time Redis cache, ensuring information immediacy.
//
// - You can specify a tenant ID to query the credit usage of a user under a specific tenant. By default, the caller\\"s default tenant is used.
//
// - You can optionally provide a `RequestId` as a request identifier, but this is not required.
//
// @param request - GetUserCreditUsageRequest
//
// @return GetUserCreditUsageResponse
func (client *Client) GetUserCreditUsage(request *GetUserCreditUsageRequest) (_result *GetUserCreditUsageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUserCreditUsageResponse{}
	_body, _err := client.GetUserCreditUsageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specified user group, including its parent group, child groups, and members.
//
// Description:
//
// ## Operation description
//
// - This operation retrieves the details of a specified user group, including the basic information of the user group, parent user group information, direct child user group list, and direct member list.
//
// - `userGroupId` is a required parameter that must be provided in the request body.
//
// - `tenantId` is an optional parameter that can be passed through the query string.
//
// - The operation supports multiple authentication methods, including AK, BearerToken, and APP authentication.
//
// - The content type for both requests and responses is `application/json`.
//
// - Ensure that you have the required permissions (such as `winnexo:GetUserGroup`) before calling this operation.
//
// @param request - GetUserGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserGroupResponse
func (client *Client) GetUserGroupWithOptions(request *GetUserGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UserGroupId) {
		body["userGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserGroup"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/getUserGroup"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified user group, including its parent group, child groups, and members.
//
// Description:
//
// ## Operation description
//
// - This operation retrieves the details of a specified user group, including the basic information of the user group, parent user group information, direct child user group list, and direct member list.
//
// - `userGroupId` is a required parameter that must be provided in the request body.
//
// - `tenantId` is an optional parameter that can be passed through the query string.
//
// - The operation supports multiple authentication methods, including AK, BearerToken, and APP authentication.
//
// - The content type for both requests and responses is `application/json`.
//
// - Ensure that you have the required permissions (such as `winnexo:GetUserGroup`) before calling this operation.
//
// @param request - GetUserGroupRequest
//
// @return GetUserGroupResponse
func (client *Client) GetUserGroup(request *GetUserGroupRequest) (_result *GetUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUserGroupResponse{}
	_body, _err := client.GetUserGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the complete information of the authenticated user through OpenAPI, including basic information and tenant list.
//
// Description:
//
// ## Request description
//
// - This operation returns the detailed information of the current authenticated user.
//
// - If the tenant information is invalid, the corresponding error message is returned.
//
// - `tenantId` is an optional parameter. If not provided, the default tenant ID of the caller is used.
//
// - Multiple authentication methods are supported: AK, BearerToken, and APP authentication.
//
// - The returned data includes the user profile (such as username and profile picture URL), role preference settings, and details of all tenants to which the user belongs.
//
// - If the current logon tenant is the system tenant (that is, `tenantId=10000`), this is explicitly indicated in the response.
//
// @param request - GetUserInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserInfoResponse
func (client *Client) GetUserInfoWithOptions(request *GetUserInfoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetUserInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserInfo"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/getUserInfo"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the complete information of the authenticated user through OpenAPI, including basic information and tenant list.
//
// Description:
//
// ## Request description
//
// - This operation returns the detailed information of the current authenticated user.
//
// - If the tenant information is invalid, the corresponding error message is returned.
//
// - `tenantId` is an optional parameter. If not provided, the default tenant ID of the caller is used.
//
// - Multiple authentication methods are supported: AK, BearerToken, and APP authentication.
//
// - The returned data includes the user profile (such as username and profile picture URL), role preference settings, and details of all tenants to which the user belongs.
//
// - If the current logon tenant is the system tenant (that is, `tenantId=10000`), this is explicitly indicated in the response.
//
// @param request - GetUserInfoRequest
//
// @return GetUserInfoResponse
func (client *Client) GetUserInfo(request *GetUserInfoRequest) (_result *GetUserInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUserInfoResponse{}
	_body, _err := client.GetUserInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Grants authorization to authorized users or user groups to use a digital human.
//
// Description:
//
// Grants authorization to authorized users or user groups to use a specified digital human.
//
//	Business logic:
//
//	1. Constructs an AuthContext from identity.
//
//	2. Performs mutual exclusion validation on the request body: specify either userIds or userGroupIds.
//
//	3. Delegates to AgentAuthorizationAuthorizedService.grant_authorization to execute.
//
//	4. Pre-validation: verifies MANAGE permission and agent existence (performed at the AuthorizedService layer, which performs authentication first before it exposes existence).
//
//	5. Existing authorization records are updated (expire_date / permissions).
//
// @param tmpReq - GrantAgentUsersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantAgentUsersResponse
func (client *Client) GrantAgentUsersWithOptions(tmpReq *GrantAgentUsersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GrantAgentUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GrantAgentUsersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Permissions) {
		request.PermissionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Permissions, dara.String("permissions"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserGroupIds) {
		request.UserGroupIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserGroupIds, dara.String("userGroupIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserIds) {
		request.UserIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserIds, dara.String("userIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ExpireDate) {
		body["expireDate"] = request.ExpireDate
	}

	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	if !dara.IsNil(request.PermissionsShrink) {
		body["permissions"] = request.PermissionsShrink
	}

	if !dara.IsNil(request.UserGroupIdsShrink) {
		body["userGroupIds"] = request.UserGroupIdsShrink
	}

	if !dara.IsNil(request.UserIdsShrink) {
		body["userIds"] = request.UserIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GrantAgentUsers"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/grantAgentUsers"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GrantAgentUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants authorization to authorized users or user groups to use a digital human.
//
// Description:
//
// Grants authorization to authorized users or user groups to use a specified digital human.
//
//	Business logic:
//
//	1. Constructs an AuthContext from identity.
//
//	2. Performs mutual exclusion validation on the request body: specify either userIds or userGroupIds.
//
//	3. Delegates to AgentAuthorizationAuthorizedService.grant_authorization to execute.
//
//	4. Pre-validation: verifies MANAGE permission and agent existence (performed at the AuthorizedService layer, which performs authentication first before it exposes existence).
//
//	5. Existing authorization records are updated (expire_date / permissions).
//
// @param request - GrantAgentUsersRequest
//
// @return GrantAgentUsersResponse
func (client *Client) GrantAgentUsers(request *GrantAgentUsersRequest) (_result *GrantAgentUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GrantAgentUsersResponse{}
	_body, _err := client.GrantAgentUsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries currently effective service notices.
//
// Description:
//
// ## Operation description
//
// Performs a paging query for published platform announcements that are effective within the current database time window. The caller must be a real user in the system O&M tenant who has the permission to view announcements.
//
// @param request - ListActiveAnnouncementsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListActiveAnnouncementsResponse
func (client *Client) ListActiveAnnouncementsWithOptions(request *ListActiveAnnouncementsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListActiveAnnouncementsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		body["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListActiveAnnouncements"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/listActiveAnnouncements"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListActiveAnnouncementsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries currently effective service notices.
//
// Description:
//
// ## Operation description
//
// Performs a paging query for published platform announcements that are effective within the current database time window. The caller must be a real user in the system O&M tenant who has the permission to view announcements.
//
// @param request - ListActiveAnnouncementsRequest
//
// @return ListActiveAnnouncementsResponse
func (client *Client) ListActiveAnnouncements(request *ListActiveAnnouncementsRequest) (_result *ListActiveAnnouncementsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListActiveAnnouncementsResponse{}
	_body, _err := client.ListActiveAnnouncementsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries or drills down into the enterprise knowledge base list of a tenant.
//
// Description:
//
// ## Operation description
//
// - This API supports two modes: when `directoryId` is empty or set to \\"root\\", the top-level knowledge base list is returned. When `directoryId` has a specific value, a drill-down operation is performed to return subdirectories and resources under the specified directory.
//
// - `tenantId` is a common parameter. If not provided, the caller\\"s tenant ID is used by default.
//
// - In drill-down mode (when `directoryId` is not empty), use the `sourceTypes` parameter to filter resources by specific types.
//
// - The sort field (`sortField`) and sort order (`sortOrder`) can be customized. Invalid values are reset to default settings.
//
// - The search feature is only effective when retrieving the top-level list and supports only fuzzy matching on names or descriptions.
//
// - For security purposes, `tenant_id` is strictly obtained from the authenticated identity and cannot be passed through the request body.
//
// @param tmpReq - ListAdminKnowledgeBasesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAdminKnowledgeBasesResponse
func (client *Client) ListAdminKnowledgeBasesWithOptions(tmpReq *ListAdminKnowledgeBasesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAdminKnowledgeBasesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListAdminKnowledgeBasesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SourceTypes) {
		request.SourceTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceTypes, dara.String("sourceTypes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		body["directoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.Keyword) {
		body["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Page) {
		body["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortField) {
		body["sortField"] = request.SortField
	}

	if !dara.IsNil(request.SortOrder) {
		body["sortOrder"] = request.SortOrder
	}

	if !dara.IsNil(request.SourceTypesShrink) {
		body["sourceTypes"] = request.SourceTypesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAdminKnowledgeBases"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/listAdminKnowledgeBases"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAdminKnowledgeBasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries or drills down into the enterprise knowledge base list of a tenant.
//
// Description:
//
// ## Operation description
//
// - This API supports two modes: when `directoryId` is empty or set to \\"root\\", the top-level knowledge base list is returned. When `directoryId` has a specific value, a drill-down operation is performed to return subdirectories and resources under the specified directory.
//
// - `tenantId` is a common parameter. If not provided, the caller\\"s tenant ID is used by default.
//
// - In drill-down mode (when `directoryId` is not empty), use the `sourceTypes` parameter to filter resources by specific types.
//
// - The sort field (`sortField`) and sort order (`sortOrder`) can be customized. Invalid values are reset to default settings.
//
// - The search feature is only effective when retrieving the top-level list and supports only fuzzy matching on names or descriptions.
//
// - For security purposes, `tenant_id` is strictly obtained from the authenticated identity and cannot be passed through the request body.
//
// @param request - ListAdminKnowledgeBasesRequest
//
// @return ListAdminKnowledgeBasesResponse
func (client *Client) ListAdminKnowledgeBases(request *ListAdminKnowledgeBasesRequest) (_result *ListAdminKnowledgeBasesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAdminKnowledgeBasesResponse{}
	_body, _err := client.ListAdminKnowledgeBasesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the full list of digital employees for a tenant.
//
// Description:
//
// Queries the full list of digital employees under a tenant, including deactivated ones.
//
//	Business logic:
//
//	1. Constructs AuthContext from identity.
//
//	2. Delegates to AgentAuthorizationAuthorizedService.list_agents to complete permission verification (APPLICATION_AGENT_VIEW).
//
//	3. Returns rich fields for all digital employees of the tenant (operatingObjectName / displayName / authMode / isActive).
//
//	4. System-level tokens are automatically allowed through ctx.skip_permission.
//
//	Difference from listAuthorizedAgents: This operation returns all digital employees of the tenant (including deactivated ones, without authorization filtering) and includes rich fields such as displayName and isActive for management console display.
//
// @param request - ListAgentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAgentsResponse
func (client *Client) ListAgentsWithOptions(request *ListAgentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAgentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAgents"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/listAgents"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAgentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the full list of digital employees for a tenant.
//
// Description:
//
// Queries the full list of digital employees under a tenant, including deactivated ones.
//
//	Business logic:
//
//	1. Constructs AuthContext from identity.
//
//	2. Delegates to AgentAuthorizationAuthorizedService.list_agents to complete permission verification (APPLICATION_AGENT_VIEW).
//
//	3. Returns rich fields for all digital employees of the tenant (operatingObjectName / displayName / authMode / isActive).
//
//	4. System-level tokens are automatically allowed through ctx.skip_permission.
//
//	Difference from listAuthorizedAgents: This operation returns all digital employees of the tenant (including deactivated ones, without authorization filtering) and includes rich fields such as displayName and isActive for management console display.
//
// @param request - ListAgentsRequest
//
// @return ListAgentsResponse
func (client *Client) ListAgents(request *ListAgentsRequest) (_result *ListAgentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAgentsResponse{}
	_body, _err := client.ListAgentsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of digital human names for which the caller has specified permissions.
//
// Description:
//
// Queries the list of digital human names for which the current caller (or a specified target user) has specified permissions (USE/MANAGE).
//
//	Business logic:
//
//	1. Constructs an AuthContext from the identity.
//
//	2. Delegates to AgentAuthorizationAuthorizedService.list_authorized_agents to execute the query.
//
//	3. When skip_permission=True, returns all active agents for the tenant.
//
//	4. Regular users are filtered based on authorization records and auth_mode.
//
//	5. When targetUserId is specified (querying on behalf of another user), the APPLICATION_AGENT_VIEW gate is required, and the query is restricted to the current tenant. If the target user is not a member of the current tenant, a USER_NOT_IN_TENANT error is thrown (an empty list is not silently returned).
//
// @param request - ListAuthorizedAgentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuthorizedAgentsResponse
func (client *Client) ListAuthorizedAgentsWithOptions(request *ListAuthorizedAgentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAuthorizedAgentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Permission) {
		body["permission"] = request.Permission
	}

	if !dara.IsNil(request.TargetUserId) {
		body["targetUserId"] = request.TargetUserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAuthorizedAgents"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/listAuthorizedAgents"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAuthorizedAgentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of digital human names for which the caller has specified permissions.
//
// Description:
//
// Queries the list of digital human names for which the current caller (or a specified target user) has specified permissions (USE/MANAGE).
//
//	Business logic:
//
//	1. Constructs an AuthContext from the identity.
//
//	2. Delegates to AgentAuthorizationAuthorizedService.list_authorized_agents to execute the query.
//
//	3. When skip_permission=True, returns all active agents for the tenant.
//
//	4. Regular users are filtered based on authorization records and auth_mode.
//
//	5. When targetUserId is specified (querying on behalf of another user), the APPLICATION_AGENT_VIEW gate is required, and the query is restricted to the current tenant. If the target user is not a member of the current tenant, a USER_NOT_IN_TENANT error is thrown (an empty list is not silently returned).
//
// @param request - ListAuthorizedAgentsRequest
//
// @return ListAuthorizedAgentsResponse
func (client *Client) ListAuthorizedAgents(request *ListAuthorizedAgentsRequest) (_result *ListAuthorizedAgentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAuthorizedAgentsResponse{}
	_body, _err := client.ListAuthorizedAgentsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of authorized users or user groups for a digital employee.
//
// Description:
//
// Queries the list of authorized users or user groups for a specified digital employee.
//
//	Business logic:
//
//	1. Constructs an AuthContext from the identity.
//
//	2. Delegates to AgentAuthorizationAuthorizedService.list_authorized_users to execute the query.
//
//	3. Permission verification is performed at the AuthorizedService layer by @require_permission(APPLICATION_AGENT_VIEW).
//
//	4. When auth_mode=ALL_USERS, only records with MANAGE permissions are displayed.
//
// @param request - ListAuthorizedUsersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuthorizedUsersResponse
func (client *Client) ListAuthorizedUsersWithOptions(request *ListAuthorizedUsersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAuthorizedUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.GranteeType) {
		body["granteeType"] = request.GranteeType
	}

	if !dara.IsNil(request.Keyword) {
		body["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	if !dara.IsNil(request.Permission) {
		body["permission"] = request.Permission
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAuthorizedUsers"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/listAuthorizedUsers"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAuthorizedUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of authorized users or user groups for a digital employee.
//
// Description:
//
// Queries the list of authorized users or user groups for a specified digital employee.
//
//	Business logic:
//
//	1. Constructs an AuthContext from the identity.
//
//	2. Delegates to AgentAuthorizationAuthorizedService.list_authorized_users to execute the query.
//
//	3. Permission verification is performed at the AuthorizedService layer by @require_permission(APPLICATION_AGENT_VIEW).
//
//	4. When auth_mode=ALL_USERS, only records with MANAGE permissions are displayed.
//
// @param request - ListAuthorizedUsersRequest
//
// @return ListAuthorizedUsersResponse
func (client *Client) ListAuthorizedUsers(request *ListAuthorizedUsersRequest) (_result *ListAuthorizedUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAuthorizedUsersResponse{}
	_body, _err := client.ListAuthorizedUsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enumerates available organization synchronization configurations.
//
// Description:
//
// Enumerates all available organization synchronization configurations under the current tenant.
//
//	Returns a unified configs list covering four platform types:
//
//	- **wecom**: Retrieves active WeCom SSO configurations from SsoProviderRegistry.
//
//	- **saml**: Retrieves active SAML SSO configurations from SsoProviderRegistry. The corpId is set to idpEntityId.
//
//	- **oauth2**: Retrieves active OAuth2 SSO configurations from SsoProviderRegistry. The corpId is set to clientId.
//
//	- **custom**: Queries the database for pure custom organizations registered under the tenant.
//
//	The client distinguishes processing logic based on the returned platformType. The corpId is a required parameter for subsequent synchronization operations.
//
// @param request - ListAvailableConfigsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAvailableConfigsResponse
func (client *Client) ListAvailableConfigsWithOptions(request *ListAvailableConfigsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAvailableConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAvailableConfigs"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/listAvailableConfigs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAvailableConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enumerates available organization synchronization configurations.
//
// Description:
//
// Enumerates all available organization synchronization configurations under the current tenant.
//
//	Returns a unified configs list covering four platform types:
//
//	- **wecom**: Retrieves active WeCom SSO configurations from SsoProviderRegistry.
//
//	- **saml**: Retrieves active SAML SSO configurations from SsoProviderRegistry. The corpId is set to idpEntityId.
//
//	- **oauth2**: Retrieves active OAuth2 SSO configurations from SsoProviderRegistry. The corpId is set to clientId.
//
//	- **custom**: Queries the database for pure custom organizations registered under the tenant.
//
//	The client distinguishes processing logic based on the returned platformType. The corpId is a required parameter for subsequent synchronization operations.
//
// @param request - ListAvailableConfigsRequest
//
// @return ListAvailableConfigsResponse
func (client *Client) ListAvailableConfigs(request *ListAvailableConfigsRequest) (_result *ListAvailableConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAvailableConfigsResponse{}
	_body, _err := client.ListAvailableConfigsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries and filters the bill list through OpenAPI with support for multiple filter conditions.
//
// Description:
//
// ## Request description
//
// - This operation queries the bill list based on specified conditions.
//
// - Supports filtering by tenant, user, operation type, status, time range, business source, and other conditions.
//
// - Returns bill data in pages. The default page size is 20 records.
//
// - You can choose whether to filter out bills with zero credit consumption. By default, such bills are filtered out.
//
// - Authentication information (such as AK, BearerToken, or APP authentication) is required for the request.
//
// @param request - ListBillingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBillingResponse
func (client *Client) ListBillingWithOptions(request *ListBillingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListBillingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		body["bizId"] = request.BizId
	}

	if !dara.IsNil(request.BizType) {
		body["bizType"] = request.BizType
	}

	if !dara.IsNil(request.EndTime) {
		body["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.IgnoreZero) {
		body["ignoreZero"] = request.IgnoreZero
	}

	if !dara.IsNil(request.Operation) {
		body["operation"] = request.Operation
	}

	if !dara.IsNil(request.Page) {
		body["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		body["startTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	if !dara.IsNil(request.WnUserId) {
		body["wnUserId"] = request.WnUserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBilling"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/listBilling"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBillingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries and filters the bill list through OpenAPI with support for multiple filter conditions.
//
// Description:
//
// ## Request description
//
// - This operation queries the bill list based on specified conditions.
//
// - Supports filtering by tenant, user, operation type, status, time range, business source, and other conditions.
//
// - Returns bill data in pages. The default page size is 20 records.
//
// - You can choose whether to filter out bills with zero credit consumption. By default, such bills are filtered out.
//
// - Authentication information (such as AK, BearerToken, or APP authentication) is required for the request.
//
// @param request - ListBillingRequest
//
// @return ListBillingResponse
func (client *Client) ListBilling(request *ListBillingRequest) (_result *ListBillingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListBillingResponse{}
	_body, _err := client.ListBillingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists chat sessions of the current user in reverse chronological order by creation time.
//
// Description:
//
// ## Operation description
//
// - This API supports filtering and sorting by multiple parameters, including tenant ID, page size, pagination token, keyword search, digital employee name, and update time range.
//
// - By default, results are sorted in descending order by the `UpdatedAt` field.
//
// - If an invalid `NextToken` is provided or `PageSize` exceeds the allowed range (1-100), the API returns a 400 error.
//
// @param request - ListChatSessionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListChatSessionsResponse
func (client *Client) ListChatSessionsWithOptions(request *ListChatSessionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListChatSessionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DigitalEmployeeName) {
		query["digitalEmployeeName"] = request.DigitalEmployeeName
	}

	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListChatSessions"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/listChatSessions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListChatSessionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists chat sessions of the current user in reverse chronological order by creation time.
//
// Description:
//
// ## Operation description
//
// - This API supports filtering and sorting by multiple parameters, including tenant ID, page size, pagination token, keyword search, digital employee name, and update time range.
//
// - By default, results are sorted in descending order by the `UpdatedAt` field.
//
// - If an invalid `NextToken` is provided or `PageSize` exceeds the allowed range (1-100), the API returns a 400 error.
//
// @param request - ListChatSessionsRequest
//
// @return ListChatSessionsResponse
func (client *Client) ListChatSessions(request *ListChatSessionsRequest) (_result *ListChatSessionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListChatSessionsResponse{}
	_body, _err := client.ListChatSessionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of knowledge graphs available for semantic queries under a tenant.
//
// Description:
//
// Lists published knowledge graphs under an identity tenant.
//
//	CLI mapping: ``winnexo graph list``. ``tenantId`` is a required common parameter and is not included in the request body.
//
//	The returned ``graphName`` can be used directly in ``querySemanticKnowledge``. This query is consistent with the existing frontend knowledge graph list and does not apply digital worker permission filtering. Specific semantic queries still verify agent USE permissions.
//
//	Database exceptions go directly into unified 5xx error handling and are not disguised as a successful empty list.
//
// @param request - ListGraphsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGraphsResponse
func (client *Client) ListGraphsWithOptions(request *ListGraphsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListGraphsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGraphs"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/listGraphs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGraphsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of knowledge graphs available for semantic queries under a tenant.
//
// Description:
//
// Lists published knowledge graphs under an identity tenant.
//
//	CLI mapping: ``winnexo graph list``. ``tenantId`` is a required common parameter and is not included in the request body.
//
//	The returned ``graphName`` can be used directly in ``querySemanticKnowledge``. This query is consistent with the existing frontend knowledge graph list and does not apply digital worker permission filtering. Specific semantic queries still verify agent USE permissions.
//
//	Database exceptions go directly into unified 5xx error handling and are not disguised as a successful empty list.
//
// @param request - ListGraphsRequest
//
// @return ListGraphsResponse
func (client *Client) ListGraphs(request *ListGraphsRequest) (_result *ListGraphsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListGraphsResponse{}
	_body, _err := client.ListGraphsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the category directory tree of an enterprise knowledge base, with support for sorting by a specified field.
//
// Description:
//
// ## Request description
//
// - This API retrieves the category list (subdirectory tree) of an enterprise knowledge base. You must have the knowledge base view permission.
//
// - If the `directoryId` parameter is not provided, the API returns all category trees under the root directory of the enterprise knowledge base. If `directoryId` is provided, the API returns the subdirectory tree rooted at the specified directory.
//
// - You can sort results by using the `sortField` and `sortOrder` parameters. By default, results are sorted by creation time in descending order.
//
// - Security constraints: `tenant_id` and `user_id` are derived only from the authenticated identity, and the caller must have the `DEVELOPMENT_KB_VIEW` feature permission.
//
// @param request - ListKnowledgeBaseDirectoriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListKnowledgeBaseDirectoriesResponse
func (client *Client) ListKnowledgeBaseDirectoriesWithOptions(request *ListKnowledgeBaseDirectoriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListKnowledgeBaseDirectoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		body["directoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.SortField) {
		body["sortField"] = request.SortField
	}

	if !dara.IsNil(request.SortOrder) {
		body["sortOrder"] = request.SortOrder
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListKnowledgeBaseDirectories"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/listKnowledgeBaseDirectories"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListKnowledgeBaseDirectoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the category directory tree of an enterprise knowledge base, with support for sorting by a specified field.
//
// Description:
//
// ## Request description
//
// - This API retrieves the category list (subdirectory tree) of an enterprise knowledge base. You must have the knowledge base view permission.
//
// - If the `directoryId` parameter is not provided, the API returns all category trees under the root directory of the enterprise knowledge base. If `directoryId` is provided, the API returns the subdirectory tree rooted at the specified directory.
//
// - You can sort results by using the `sortField` and `sortOrder` parameters. By default, results are sorted by creation time in descending order.
//
// - Security constraints: `tenant_id` and `user_id` are derived only from the authenticated identity, and the caller must have the `DEVELOPMENT_KB_VIEW` feature permission.
//
// @param request - ListKnowledgeBaseDirectoriesRequest
//
// @return ListKnowledgeBaseDirectoriesResponse
func (client *Client) ListKnowledgeBaseDirectories(request *ListKnowledgeBaseDirectoriesRequest) (_result *ListKnowledgeBaseDirectoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListKnowledgeBaseDirectoriesResponse{}
	_body, _err := client.ListKnowledgeBaseDirectoriesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the precise object type follows of a digital employee by page.
//
// Description:
//
// Queries follows by three independent dimensions: graphName, operatingObjectName, and objectType. Supports primary objects and explicit first-level associated objects. Uses opaque cursor pagination and is not limited by the 1000-item display window of the follow panel.
//
// @param request - ListOperatingObjectFavoritesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOperatingObjectFavoritesResponse
func (client *Client) ListOperatingObjectFavoritesWithOptions(request *ListOperatingObjectFavoritesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListOperatingObjectFavoritesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.GraphName) {
		body["graphName"] = request.GraphName
	}

	if !dara.IsNil(request.NextToken) {
		body["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ObjectType) {
		body["objectType"] = request.ObjectType
	}

	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOperatingObjectFavorites"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/listOperatingObjectFavorites"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOperatingObjectFavoritesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the precise object type follows of a digital employee by page.
//
// Description:
//
// Queries follows by three independent dimensions: graphName, operatingObjectName, and objectType. Supports primary objects and explicit first-level associated objects. Uses opaque cursor pagination and is not limited by the 1000-item display window of the follow panel.
//
// @param request - ListOperatingObjectFavoritesRequest
//
// @return ListOperatingObjectFavoritesResponse
func (client *Client) ListOperatingObjectFavorites(request *ListOperatingObjectFavoritesRequest) (_result *ListOperatingObjectFavoritesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListOperatingObjectFavoritesResponse{}
	_body, _err := client.ListOperatingObjectFavoritesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the output list of the current user, with support for conditional filtering and pagination.
//
// Description:
//
// ## Operation description
//
// - This API operation queries the output list of the current logged-in user.
//
// - `tenantId` is a common parameter. If not specified, the default tenant of the caller is used.
//
// - Supports filtering by parameters such as `operatingObjectName`, `itemType`, and `keyword`.
//
// - Set `sharedOnly` to `true` to display only outputs with sharing enabled.
//
// - Pagination is controlled by `page` (page number) and `pageSize` (number of items per page). By default, results start from page 1 with 20 records per page.
//
// - Results are sorted by update time in descending order by default.
//
// - The `tenant_id` or `user_id` passed in the request body by the caller is ignored. This information is derived only from the authenticated identity.
//
// @param request - ListOutputFilesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOutputFilesResponse
func (client *Client) ListOutputFilesWithOptions(request *ListOutputFilesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListOutputFilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ItemType) {
		body["itemType"] = request.ItemType
	}

	if !dara.IsNil(request.Keyword) {
		body["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	if !dara.IsNil(request.Page) {
		body["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SharedOnly) {
		body["sharedOnly"] = request.SharedOnly
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOutputFiles"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/listOutputFiles"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOutputFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the output list of the current user, with support for conditional filtering and pagination.
//
// Description:
//
// ## Operation description
//
// - This API operation queries the output list of the current logged-in user.
//
// - `tenantId` is a common parameter. If not specified, the default tenant of the caller is used.
//
// - Supports filtering by parameters such as `operatingObjectName`, `itemType`, and `keyword`.
//
// - Set `sharedOnly` to `true` to display only outputs with sharing enabled.
//
// - Pagination is controlled by `page` (page number) and `pageSize` (number of items per page). By default, results start from page 1 with 20 records per page.
//
// - Results are sorted by update time in descending order by default.
//
// - The `tenant_id` or `user_id` passed in the request body by the caller is ignored. This information is derived only from the authenticated identity.
//
// @param request - ListOutputFilesRequest
//
// @return ListOutputFilesResponse
func (client *Client) ListOutputFiles(request *ListOutputFilesRequest) (_result *ListOutputFilesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListOutputFilesResponse{}
	_body, _err := client.ListOutputFilesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries subdirectories and resources under a specified digital employee resource directory.
//
// Description:
//
// ## Operation description
//
// - This API is used to drill down and query subdirectories and resources under the "My Resources" directory.
//
// - When `directoryId` is set to \\"root\\", the service automatically resolves and returns the content under the current digital employee\\"s default root directory. If a specific directory ID is provided, the subdirectories and resources under that directory are returned.
//
// - Security constraint: `tenant_id` and `user_id` can only come from the authenticated identity information. These fields provided by the caller in the request body are ignored.
//
// - You can use the `sourceTypes` parameter to filter resources of specific types. When this parameter has a value, only resources that match the type condition are returned, and subdirectories are not included.
//
// - Sorting supports ascending or descending order by name (`name`), creation time (`gmt_create`), or modification time (`gmt_modified`).
//
// - The pagination feature allows you to customize the number of items displayed per page (maximum 100) and the current page number.
//
// @param tmpReq - ListPersonalDirectoryContentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPersonalDirectoryContentsResponse
func (client *Client) ListPersonalDirectoryContentsWithOptions(tmpReq *ListPersonalDirectoryContentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPersonalDirectoryContentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListPersonalDirectoryContentsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SourceTypes) {
		request.SourceTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceTypes, dara.String("sourceTypes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		body["directoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	if !dara.IsNil(request.Page) {
		body["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortField) {
		body["sortField"] = request.SortField
	}

	if !dara.IsNil(request.SortOrder) {
		body["sortOrder"] = request.SortOrder
	}

	if !dara.IsNil(request.SourceTypesShrink) {
		body["sourceTypes"] = request.SourceTypesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPersonalDirectoryContents"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/listPersonalDirectoryContents"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPersonalDirectoryContentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries subdirectories and resources under a specified digital employee resource directory.
//
// Description:
//
// ## Operation description
//
// - This API is used to drill down and query subdirectories and resources under the "My Resources" directory.
//
// - When `directoryId` is set to \\"root\\", the service automatically resolves and returns the content under the current digital employee\\"s default root directory. If a specific directory ID is provided, the subdirectories and resources under that directory are returned.
//
// - Security constraint: `tenant_id` and `user_id` can only come from the authenticated identity information. These fields provided by the caller in the request body are ignored.
//
// - You can use the `sourceTypes` parameter to filter resources of specific types. When this parameter has a value, only resources that match the type condition are returned, and subdirectories are not included.
//
// - Sorting supports ascending or descending order by name (`name`), creation time (`gmt_create`), or modification time (`gmt_modified`).
//
// - The pagination feature allows you to customize the number of items displayed per page (maximum 100) and the current page number.
//
// @param request - ListPersonalDirectoryContentsRequest
//
// @return ListPersonalDirectoryContentsResponse
func (client *Client) ListPersonalDirectoryContents(request *ListPersonalDirectoryContentsRequest) (_result *ListPersonalDirectoryContentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPersonalDirectoryContentsResponse{}
	_body, _err := client.ListPersonalDirectoryContentsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of system built-in roles.
//
// Description:
//
// Queries the list of system built-in roles.
//
//	Business logic:
//
//	1. Constructs AuthContext from identity.
//
//	2. Delegates to UserManagementAuthorizedService.list_system_roles for permission verification (PLATFORM_USER_VIEW).
//
//	3. Renders role names and descriptions based on the request Accept-Language header.
//
//	4. Returns a fixed set of 7 system built-in roles.
//
//	The returned roleCode field can be directly used as the roleCodes parameter for createUser or updateUser.
//
// @param request - ListRolesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRolesResponse
func (client *Client) ListRolesWithOptions(request *ListRolesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRolesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRoles"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/listRoles"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRolesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of system built-in roles.
//
// Description:
//
// Queries the list of system built-in roles.
//
//	Business logic:
//
//	1. Constructs AuthContext from identity.
//
//	2. Delegates to UserManagementAuthorizedService.list_system_roles for permission verification (PLATFORM_USER_VIEW).
//
//	3. Renders role names and descriptions based on the request Accept-Language header.
//
//	4. Returns a fixed set of 7 system built-in roles.
//
//	The returned roleCode field can be directly used as the roleCodes parameter for createUser or updateUser.
//
// @param request - ListRolesRequest
//
// @return ListRolesResponse
func (client *Client) ListRoles(request *ListRolesRequest) (_result *ListRolesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRolesResponse{}
	_body, _err := client.ListRolesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of scheduled tasks.
//
// Description:
//
// ## Operation description
//
// - This operation uploads a file to an enterprise knowledge base.
//
// - The DEVELOPMENT_KB_MANAGE permission is required to call this operation.
//
// - You must provide the OSS persistent address (`filePath`) of the file when uploading.
//
// - Optional parameters include the public access URL and original file name to enhance the completeness of file information.
//
// - If `directoryId` is specified, the file is placed in the corresponding enterprise knowledge base directory. Otherwise, the file is bound to the default root directory of the current digital employee.
//
// - You can add tags to the resource by using `sourceTags` for subsequent management and retrieval.
//
// - This operation initiates a billing item (UNSTRUCTURED_PARSE). Make sure your account balance is sufficient.
//
// @param tmpReq - ListScheduledTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScheduledTasksResponse
func (client *Client) ListScheduledTasksWithOptions(tmpReq *ListScheduledTasksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListScheduledTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListScheduledTasksShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Visibilities) {
		request.VisibilitiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Visibilities, dara.String("visibilities"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CollaborationGroupId) {
		query["collaborationGroupId"] = request.CollaborationGroupId
	}

	if !dara.IsNil(request.CreatorOnly) {
		query["creatorOnly"] = request.CreatorOnly
	}

	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	if !dara.IsNil(request.VisibilitiesShrink) {
		query["visibilities"] = request.VisibilitiesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListScheduledTasks"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/listScheduledTasks"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListScheduledTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of scheduled tasks.
//
// Description:
//
// ## Operation description
//
// - This operation uploads a file to an enterprise knowledge base.
//
// - The DEVELOPMENT_KB_MANAGE permission is required to call this operation.
//
// - You must provide the OSS persistent address (`filePath`) of the file when uploading.
//
// - Optional parameters include the public access URL and original file name to enhance the completeness of file information.
//
// - If `directoryId` is specified, the file is placed in the corresponding enterprise knowledge base directory. Otherwise, the file is bound to the default root directory of the current digital employee.
//
// - You can add tags to the resource by using `sourceTags` for subsequent management and retrieval.
//
// - This operation initiates a billing item (UNSTRUCTURED_PARSE). Make sure your account balance is sufficient.
//
// @param request - ListScheduledTasksRequest
//
// @return ListScheduledTasksResponse
func (client *Client) ListScheduledTasks(request *ListScheduledTasksRequest) (_result *ListScheduledTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListScheduledTasksResponse{}
	_body, _err := client.ListScheduledTasksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the skills visible to the current tenant.
//
// Description:
//
// ## Request description
//
// This API retrieves all visible skills under the current tenant. It supports filtering by digital employee binding relationship, skill source, tags, and keywords, and supports pagination.
//
// ### Request parameters
//
// - **TenantId**: Optional. A common parameter passed through by the gateway to the backend header. If not specified, the default tenant of the current caller is used.
//
// - **FilterType**: Optional. The skill filtering dimension. Valid values: `ALL` (all published), `BUILTIN` (built-in published), `CUSTOM` (custom published), `DRAFT` (drafts, including published skills with unpublished modifications). Default value: `ALL`.
//
// - **Tags**: Optional. Filters by tags. A match is returned if any tag in the array is hit.
//
// - **Keyword**: Optional. Performs a fuzzy match on the skill name or description.
//
// - **Page**: Optional. The page number. Minimum value: 1. Default value: 1.
//
// - **PageSize**: Optional. The number of entries per page. Valid values: 1 to 100. Default value: 20.
//
// - **OperatingObjectName**: Optional. The digital employee name. If specified, results are filtered by binding relationship. Must be used together with `BindStatus`.
//
// - **BindStatus**: Optional. The binding status. Valid values: `BOUND` (bound), `UNBOUND` (unbound global skills).
//
// ### Response parameters
//
// The response contains the skill list `items`, total count `total`, current page `page`, and page size `pageSize`.
//
// @param tmpReq - ListSkillsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSkillsResponse
func (client *Client) ListSkillsWithOptions(tmpReq *ListSkillsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSkillsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListSkillsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BindStatus) {
		body["bindStatus"] = request.BindStatus
	}

	if !dara.IsNil(request.FilterType) {
		body["filterType"] = request.FilterType
	}

	if !dara.IsNil(request.Keyword) {
		body["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	if !dara.IsNil(request.Page) {
		body["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TagsShrink) {
		body["tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSkills"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/listSkills"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSkillsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the skills visible to the current tenant.
//
// Description:
//
// ## Request description
//
// This API retrieves all visible skills under the current tenant. It supports filtering by digital employee binding relationship, skill source, tags, and keywords, and supports pagination.
//
// ### Request parameters
//
// - **TenantId**: Optional. A common parameter passed through by the gateway to the backend header. If not specified, the default tenant of the current caller is used.
//
// - **FilterType**: Optional. The skill filtering dimension. Valid values: `ALL` (all published), `BUILTIN` (built-in published), `CUSTOM` (custom published), `DRAFT` (drafts, including published skills with unpublished modifications). Default value: `ALL`.
//
// - **Tags**: Optional. Filters by tags. A match is returned if any tag in the array is hit.
//
// - **Keyword**: Optional. Performs a fuzzy match on the skill name or description.
//
// - **Page**: Optional. The page number. Minimum value: 1. Default value: 1.
//
// - **PageSize**: Optional. The number of entries per page. Valid values: 1 to 100. Default value: 20.
//
// - **OperatingObjectName**: Optional. The digital employee name. If specified, results are filtered by binding relationship. Must be used together with `BindStatus`.
//
// - **BindStatus**: Optional. The binding status. Valid values: `BOUND` (bound), `UNBOUND` (unbound global skills).
//
// ### Response parameters
//
// The response contains the skill list `items`, total count `total`, current page `page`, and page size `pageSize`.
//
// @param request - ListSkillsRequest
//
// @return ListSkillsResponse
func (client *Client) ListSkills(request *ListSkillsRequest) (_result *ListSkillsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSkillsResponse{}
	_body, _err := client.ListSkillsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the list of knowledge bases.
//
// Description:
//
// ## Request description
//
// - This API is used to perform a paging query on the folder content and resources in an enterprise knowledge base.
//
// - Multiple parameters are supported for filtering and sorting, such as `directoryId`, `page`, `pageSize`, `sortField`, `sortOrder`, and others.
//
// - The `sourceTypes` parameter allows you to filter by resource type. Separate multiple types with commas.
//
// - When `directoryId` is not specified or set to `root`, the root folder list of the knowledge base is queried by default.
//
// - The default sort field is `name`, and the default sort order is ascending (`asc`).
//
// @param request - ListTenantDirectoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTenantDirectoryResponse
func (client *Client) ListTenantDirectoryWithOptions(request *ListTenantDirectoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTenantDirectoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		body["directoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.Page) {
		body["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortField) {
		body["sortField"] = request.SortField
	}

	if !dara.IsNil(request.SortOrder) {
		body["sortOrder"] = request.SortOrder
	}

	if !dara.IsNil(request.SourceTypes) {
		body["sourceTypes"] = request.SourceTypes
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTenantDirectory"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/listTenantDirectory"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTenantDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of knowledge bases.
//
// Description:
//
// ## Request description
//
// - This API is used to perform a paging query on the folder content and resources in an enterprise knowledge base.
//
// - Multiple parameters are supported for filtering and sorting, such as `directoryId`, `page`, `pageSize`, `sortField`, `sortOrder`, and others.
//
// - The `sourceTypes` parameter allows you to filter by resource type. Separate multiple types with commas.
//
// - When `directoryId` is not specified or set to `root`, the root folder list of the knowledge base is queried by default.
//
// - The default sort field is `name`, and the default sort order is ascending (`asc`).
//
// @param request - ListTenantDirectoryRequest
//
// @return ListTenantDirectoryResponse
func (client *Client) ListTenantDirectory(request *ListTenantDirectoryRequest) (_result *ListTenantDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTenantDirectoryResponse{}
	_body, _err := client.ListTenantDirectoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Returns the multi-level user group tree for the current tenant.
//
// Description:
//
// ## Request description
//
// This API is used to query the complete user group hierarchy under a specified tenant, including the basic information of each user group and its direct child user group list. Use the `tenantId` parameter to specify the tenant ID to query. If this parameter is not provided, the caller\\"s tenant ID is used by default.
//
// ### Precautions
//
// - This operation returns only the direct member count and direct child user group count. It does not include information about indirect members or child groups.
//
// - The external synchronization status field is empty when data is normal. It is populated with relevant information only when data is out of sync between an external system (such as WeCom) and the internal system.
//
// @param request - ListUserGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserGroupsResponse
func (client *Client) ListUserGroupsWithOptions(request *ListUserGroupsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListUserGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserGroups"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/listUserGroups"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Returns the multi-level user group tree for the current tenant.
//
// Description:
//
// ## Request description
//
// This API is used to query the complete user group hierarchy under a specified tenant, including the basic information of each user group and its direct child user group list. Use the `tenantId` parameter to specify the tenant ID to query. If this parameter is not provided, the caller\\"s tenant ID is used by default.
//
// ### Precautions
//
// - This operation returns only the direct member count and direct child user group count. It does not include information about indirect members or child groups.
//
// - The external synchronization status field is empty when data is normal. It is populated with relevant information only when data is out of sync between an external system (such as WeCom) and the internal system.
//
// @param request - ListUserGroupsRequest
//
// @return ListUserGroupsResponse
func (client *Client) ListUserGroups(request *ListUserGroupsRequest) (_result *ListUserGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListUserGroupsResponse{}
	_body, _err := client.ListUserGroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the knowledge base directory content visible to the current OpenAPI user.
//
// Description:
//
// ## Operation description
//
// - This operation returns subdirectories and READY resources under the specified directory based on the enterprise knowledge base frontend scope.
//
// - The user identity and directory visibility scope are derived from the OpenAPI authentication context.
//
// - When `sourceTypes` has a value, only resources are returned. `keyword` searches only the current directory level.
//
// @param request - ListUserVisibleKnowledgeBaseContentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserVisibleKnowledgeBaseContentsResponse
func (client *Client) ListUserVisibleKnowledgeBaseContentsWithOptions(request *ListUserVisibleKnowledgeBaseContentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListUserVisibleKnowledgeBaseContentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		body["directoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.Keyword) {
		body["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Page) {
		body["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortField) {
		body["sortField"] = request.SortField
	}

	if !dara.IsNil(request.SortOrder) {
		body["sortOrder"] = request.SortOrder
	}

	if !dara.IsNil(request.SourceTypes) {
		body["sourceTypes"] = request.SourceTypes
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserVisibleKnowledgeBaseContents"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/listUserVisibleKnowledgeBaseContents"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserVisibleKnowledgeBaseContentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the knowledge base directory content visible to the current OpenAPI user.
//
// Description:
//
// ## Operation description
//
// - This operation returns subdirectories and READY resources under the specified directory based on the enterprise knowledge base frontend scope.
//
// - The user identity and directory visibility scope are derived from the OpenAPI authentication context.
//
// - When `sourceTypes` has a value, only resources are returned. `keyword` searches only the current directory level.
//
// @param request - ListUserVisibleKnowledgeBaseContentsRequest
//
// @return ListUserVisibleKnowledgeBaseContentsResponse
func (client *Client) ListUserVisibleKnowledgeBaseContents(request *ListUserVisibleKnowledgeBaseContentsRequest) (_result *ListUserVisibleKnowledgeBaseContentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListUserVisibleKnowledgeBaseContentsResponse{}
	_body, _err := client.ListUserVisibleKnowledgeBaseContentsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of enterprise knowledge bases visible to the current OpenAPI user.
//
// Description:
//
// ## Operation description
//
// - This operation queries the enterprise knowledge bases visible to the platform user mapped from the OpenAPI authentication identity.
//
// - Both the tenant and user identities are determined by the authentication context. Callers cannot expand the visible scope through business parameters.
//
// - `tenantId` is an optional common parameter. `keyword` can filter by knowledge base name or description.
//
// @param request - ListUserVisibleKnowledgeBasesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserVisibleKnowledgeBasesResponse
func (client *Client) ListUserVisibleKnowledgeBasesWithOptions(request *ListUserVisibleKnowledgeBasesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListUserVisibleKnowledgeBasesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		body["keyword"] = request.Keyword
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserVisibleKnowledgeBases"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/listUserVisibleKnowledgeBases"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserVisibleKnowledgeBasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of enterprise knowledge bases visible to the current OpenAPI user.
//
// Description:
//
// ## Operation description
//
// - This operation queries the enterprise knowledge bases visible to the platform user mapped from the OpenAPI authentication identity.
//
// - Both the tenant and user identities are determined by the authentication context. Callers cannot expand the visible scope through business parameters.
//
// - `tenantId` is an optional common parameter. `keyword` can filter by knowledge base name or description.
//
// @param request - ListUserVisibleKnowledgeBasesRequest
//
// @return ListUserVisibleKnowledgeBasesResponse
func (client *Client) ListUserVisibleKnowledgeBases(request *ListUserVisibleKnowledgeBasesRequest) (_result *ListUserVisibleKnowledgeBasesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListUserVisibleKnowledgeBasesResponse{}
	_body, _err := client.ListUserVisibleKnowledgeBasesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries tenant members by paging.
//
// Description:
//
// Queries a paged list of tenant members by using OpenAPI.
//
//	Business orchestration:
//
//	1. Parse filter conditions (roleCodes → role_ids).
//
//	2. Call UserTenantMappingRepository.query_paged_tenant_members to perform a paged query.
//
//	3. Convert role_id in the results to roleCode and assemble the response.
//
//	Error codes:
//
//	- An error is thrown when an invalid roleCode parameter is specified.
//
// @param tmpReq - ListUsersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersResponse
func (client *Client) ListUsersWithOptions(tmpReq *ListUsersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListUsersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AccountIds) {
		request.AccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AccountIds, dara.String("accountIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RoleCodes) {
		request.RoleCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoleCodes, dara.String("roleCodes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountIdsShrink) {
		body["accountIds"] = request.AccountIdsShrink
	}

	if !dara.IsNil(request.IsActive) {
		body["isActive"] = request.IsActive
	}

	if !dara.IsNil(request.Keyword) {
		body["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Page) {
		body["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RoleCodesShrink) {
		body["roleCodes"] = request.RoleCodesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUsers"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/listUsers"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries tenant members by paging.
//
// Description:
//
// Queries a paged list of tenant members by using OpenAPI.
//
//	Business orchestration:
//
//	1. Parse filter conditions (roleCodes → role_ids).
//
//	2. Call UserTenantMappingRepository.query_paged_tenant_members to perform a paged query.
//
//	3. Convert role_id in the results to roleCode and assemble the response.
//
//	Error codes:
//
//	- An error is thrown when an invalid roleCode parameter is specified.
//
// @param request - ListUsersRequest
//
// @return ListUsersResponse
func (client *Client) ListUsers(request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Drills down to query subdirectories and resources under a specified knowledge base directory visible to a specified digital employee.
//
// Description:
//
// ## Operation description
//
// - This operation queries all subdirectories and resources under a specified knowledge base directory for a specific digital employee.
//
// - The user must have the USE permission on the target digital employee, and the digital employee must have access to the directory and its subdirectories specified in the request.
//
// - You must provide the digital employee name (`operatingObjectName`) and the directory ID (`directoryId`) to query. Other parameters such as pagination information and sorting method are optional.
//
// - The response includes the list of subdirectories and resources under the directory, and supports pagination.
//
// - The `sourceStatus` field filters only resources in the `READY` state.
//
// - For security purposes, `tenant_id` and `user_id` are obtained only from the authenticated identity. Values passed in the request body by the caller are ignored.
//
// @param tmpReq - ListVisibleKnowledgeBaseContentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVisibleKnowledgeBaseContentsResponse
func (client *Client) ListVisibleKnowledgeBaseContentsWithOptions(tmpReq *ListVisibleKnowledgeBaseContentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListVisibleKnowledgeBaseContentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListVisibleKnowledgeBaseContentsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SourceTypes) {
		request.SourceTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceTypes, dara.String("sourceTypes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		body["directoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	if !dara.IsNil(request.Page) {
		body["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortField) {
		body["sortField"] = request.SortField
	}

	if !dara.IsNil(request.SortOrder) {
		body["sortOrder"] = request.SortOrder
	}

	if !dara.IsNil(request.SourceTypesShrink) {
		body["sourceTypes"] = request.SourceTypesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVisibleKnowledgeBaseContents"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/listVisibleKnowledgeBaseContents"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVisibleKnowledgeBaseContentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Drills down to query subdirectories and resources under a specified knowledge base directory visible to a specified digital employee.
//
// Description:
//
// ## Operation description
//
// - This operation queries all subdirectories and resources under a specified knowledge base directory for a specific digital employee.
//
// - The user must have the USE permission on the target digital employee, and the digital employee must have access to the directory and its subdirectories specified in the request.
//
// - You must provide the digital employee name (`operatingObjectName`) and the directory ID (`directoryId`) to query. Other parameters such as pagination information and sorting method are optional.
//
// - The response includes the list of subdirectories and resources under the directory, and supports pagination.
//
// - The `sourceStatus` field filters only resources in the `READY` state.
//
// - For security purposes, `tenant_id` and `user_id` are obtained only from the authenticated identity. Values passed in the request body by the caller are ignored.
//
// @param request - ListVisibleKnowledgeBaseContentsRequest
//
// @return ListVisibleKnowledgeBaseContentsResponse
func (client *Client) ListVisibleKnowledgeBaseContents(request *ListVisibleKnowledgeBaseContentsRequest) (_result *ListVisibleKnowledgeBaseContentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListVisibleKnowledgeBaseContentsResponse{}
	_body, _err := client.ListVisibleKnowledgeBaseContentsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the top-level directories of enterprise knowledge bases accessible to a digital employee.
//
// Description:
//
// ## Request description
//
// - This API operation retrieves the list of top-level knowledge base directories visible to a specified digital employee (operating object) within the enterprise.
//
// @param request - ListVisibleKnowledgeBasesRequest
//
// @param headers - ListVisibleKnowledgeBasesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVisibleKnowledgeBasesResponse
func (client *Client) ListVisibleKnowledgeBasesWithOptions(request *ListVisibleKnowledgeBasesRequest, headers *ListVisibleKnowledgeBasesHeaders, runtime *dara.RuntimeOptions) (_result *ListVisibleKnowledgeBasesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.RequestId) {
		realHeaders["requestId"] = dara.String(dara.ToString(dara.StringValue(headers.RequestId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVisibleKnowledgeBases"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/listVisibleKnowledgeBases"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVisibleKnowledgeBasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the top-level directories of enterprise knowledge bases accessible to a digital employee.
//
// Description:
//
// ## Request description
//
// - This API operation retrieves the list of top-level knowledge base directories visible to a specified digital employee (operating object) within the enterprise.
//
// @param request - ListVisibleKnowledgeBasesRequest
//
// @return ListVisibleKnowledgeBasesResponse
func (client *Client) ListVisibleKnowledgeBases(request *ListVisibleKnowledgeBasesRequest) (_result *ListVisibleKnowledgeBasesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListVisibleKnowledgeBasesHeaders{}
	_result = &ListVisibleKnowledgeBasesResponse{}
	_body, _err := client.ListVisibleKnowledgeBasesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Moves a specified resource between enterprise knowledge base directories. Management permissions are required.
//
// Description:
//
// ## Operation description
//
// - **Authentication flow**:
//
//  1. Basic authentication is performed by the root router (`request.state.openapi_identity`).
//
//  2. This handler checks the `DEVELOPMENT_KB_MANAGE` feature permission.
//
// - **Procedure**:
//
//  1. Check that the source directory and target directory are not the same.
//
//  2. Confirm that the target directory exists.
//
//  3. Verify that the resource to be moved is in the source directory.
//
//  4. Update the directory binding of the resource.
//
//  5. Best-effort update of `source.settings["knowledge_id"]` to the target knowledge base ID.
//
//  6. Best-effort notification to DocumentAgent to sync `knowledge_id` and `update_time`.
//
// - **Security constraints**:
//
//   - `tenant_id` and `user_id` must come from the authenticated identity.
//
//   - The caller must have KB management permissions.
//
// @param request - MoveKnowledgeBaseResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveKnowledgeBaseResourceResponse
func (client *Client) MoveKnowledgeBaseResourceWithOptions(request *MoveKnowledgeBaseResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *MoveKnowledgeBaseResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.KnowledgeId) {
		body["knowledgeId"] = request.KnowledgeId
	}

	if !dara.IsNil(request.SourceDirectoryId) {
		body["sourceDirectoryId"] = request.SourceDirectoryId
	}

	if !dara.IsNil(request.SourceId) {
		body["sourceId"] = request.SourceId
	}

	if !dara.IsNil(request.TargetDirectoryId) {
		body["targetDirectoryId"] = request.TargetDirectoryId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveKnowledgeBaseResource"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/moveKnowledgeBaseResource"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveKnowledgeBaseResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves a specified resource between enterprise knowledge base directories. Management permissions are required.
//
// Description:
//
// ## Operation description
//
// - **Authentication flow**:
//
//  1. Basic authentication is performed by the root router (`request.state.openapi_identity`).
//
//  2. This handler checks the `DEVELOPMENT_KB_MANAGE` feature permission.
//
// - **Procedure**:
//
//  1. Check that the source directory and target directory are not the same.
//
//  2. Confirm that the target directory exists.
//
//  3. Verify that the resource to be moved is in the source directory.
//
//  4. Update the directory binding of the resource.
//
//  5. Best-effort update of `source.settings["knowledge_id"]` to the target knowledge base ID.
//
//  6. Best-effort notification to DocumentAgent to sync `knowledge_id` and `update_time`.
//
// - **Security constraints**:
//
//   - `tenant_id` and `user_id` must come from the authenticated identity.
//
//   - The caller must have KB management permissions.
//
// @param request - MoveKnowledgeBaseResourceRequest
//
// @return MoveKnowledgeBaseResourceResponse
func (client *Client) MoveKnowledgeBaseResource(request *MoveKnowledgeBaseResourceRequest) (_result *MoveKnowledgeBaseResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &MoveKnowledgeBaseResourceResponse{}
	_body, _err := client.MoveKnowledgeBaseResourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Moves a specified resource between personal directories of a user.
//
// Description:
//
// ## Request description
//
// - **The source directory and target directory cannot be the same**. Otherwise, the `ERR.Robject.UserDirectory.InvalidOperation` error is returned.
//
// - **The target directory must exist**. If it does not exist, the `ERR.Robject.UserDirectory.DirectoryNotFound` error is returned.
//
// - **The resource to be moved must exist in the source directory**. If it is not in the source directory, the `ERR.Robject.UserDirectory.ResourceNotInDirectory` error is returned.
//
// - After a successful move, the system attempts to notify DocumentAgent to update the new path (`source_path`) of the resource. This step is best-effort. Even if it fails, the overall operation success status is not affected. Only an error log is recorded.
//
// - For security purposes, the value of `tenant_id` can only be derived from the authenticated identity information.
//
// @param request - MoveResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveResourceResponse
func (client *Client) MoveResourceWithOptions(request *MoveResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *MoveResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SourceDirectoryId) {
		body["sourceDirectoryId"] = request.SourceDirectoryId
	}

	if !dara.IsNil(request.SourceId) {
		body["sourceId"] = request.SourceId
	}

	if !dara.IsNil(request.TargetDirectoryId) {
		body["targetDirectoryId"] = request.TargetDirectoryId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveResource"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/moveResource"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves a specified resource between personal directories of a user.
//
// Description:
//
// ## Request description
//
// - **The source directory and target directory cannot be the same**. Otherwise, the `ERR.Robject.UserDirectory.InvalidOperation` error is returned.
//
// - **The target directory must exist**. If it does not exist, the `ERR.Robject.UserDirectory.DirectoryNotFound` error is returned.
//
// - **The resource to be moved must exist in the source directory**. If it is not in the source directory, the `ERR.Robject.UserDirectory.ResourceNotInDirectory` error is returned.
//
// - After a successful move, the system attempts to notify DocumentAgent to update the new path (`source_path`) of the resource. This step is best-effort. Even if it fails, the overall operation success status is not affected. Only an error log is recorded.
//
// - For security purposes, the value of `tenant_id` can only be derived from the authenticated identity information.
//
// @param request - MoveResourceRequest
//
// @return MoveResourceResponse
func (client *Client) MoveResource(request *MoveResourceRequest) (_result *MoveResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &MoveResourceResponse{}
	_body, _err := client.MoveResourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Takes a service notice offline.
//
// Description:
//
// ## Operation description
//
// Idempotently takes a platform announcement offline by announcement ID. Returns `changed=true` when a PUBLISHED announcement is taken offline for the first time. Returns `changed=false` when the announcement is already offline or expired.
//
// The caller must belong to the system operations tenant and have announcement management permissions.
//
// @param request - OfflineAnnouncementRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OfflineAnnouncementResponse
func (client *Client) OfflineAnnouncementWithOptions(request *OfflineAnnouncementRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OfflineAnnouncementResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AnnouncementId) {
		body["announcementId"] = request.AnnouncementId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OfflineAnnouncement"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/offlineAnnouncement"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OfflineAnnouncementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Takes a service notice offline.
//
// Description:
//
// ## Operation description
//
// Idempotently takes a platform announcement offline by announcement ID. Returns `changed=true` when a PUBLISHED announcement is taken offline for the first time. Returns `changed=false` when the announcement is already offline or expired.
//
// The caller must belong to the system operations tenant and have announcement management permissions.
//
// @param request - OfflineAnnouncementRequest
//
// @return OfflineAnnouncementResponse
func (client *Client) OfflineAnnouncement(request *OfflineAnnouncementRequest) (_result *OfflineAnnouncementResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OfflineAnnouncementResponse{}
	_body, _err := client.OfflineAnnouncementWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Previews the knowledge content in a specified enterprise knowledge base.
//
// Description:
//
// ## Operation description
//
// - This operation previews the content of a specified knowledge entry in an enterprise knowledge base.
//
// - The `DEVELOPMENT_KB_VIEW` permission is required to call this API.
//
// - `sourceId` is a required parameter that identifies the knowledge entry to preview.
//
// - The optional parameter `tenantId` specifies the tenant ID. If not provided, the default tenant ID of the caller is used.
//
// - Multiple preview types are supported, including but not limited to images, audio, video, and text.
//
// @param request - PreviewKnowledgeBaseSourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PreviewKnowledgeBaseSourceResponse
func (client *Client) PreviewKnowledgeBaseSourceWithOptions(request *PreviewKnowledgeBaseSourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PreviewKnowledgeBaseSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SourceId) {
		body["sourceId"] = request.SourceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PreviewKnowledgeBaseSource"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/previewKnowledgeBaseSource"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PreviewKnowledgeBaseSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Previews the knowledge content in a specified enterprise knowledge base.
//
// Description:
//
// ## Operation description
//
// - This operation previews the content of a specified knowledge entry in an enterprise knowledge base.
//
// - The `DEVELOPMENT_KB_VIEW` permission is required to call this API.
//
// - `sourceId` is a required parameter that identifies the knowledge entry to preview.
//
// - The optional parameter `tenantId` specifies the tenant ID. If not provided, the default tenant ID of the caller is used.
//
// - Multiple preview types are supported, including but not limited to images, audio, video, and text.
//
// @param request - PreviewKnowledgeBaseSourceRequest
//
// @return PreviewKnowledgeBaseSourceResponse
func (client *Client) PreviewKnowledgeBaseSource(request *PreviewKnowledgeBaseSourceRequest) (_result *PreviewKnowledgeBaseSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PreviewKnowledgeBaseSourceResponse{}
	_body, _err := client.PreviewKnowledgeBaseSourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Allows a user to preview specified knowledge content in their personal directory.
//
// Description:
//
// ## Request description
//
// - This operation only allows a user to preview resources in their own personal directory.
//
// - The authentication process includes basic authentication and data source ownership verification to ensure that the requester can only access knowledge in their personal directory.
//
// - You must provide the unique identifier `sourceId` of the knowledge content in the request. The system queries and returns the corresponding preview information based on this ID and the user\\"s tenant information.
//
// - Multiple preview types are supported, such as image, audio, and video. The system returns the corresponding preview URL or direct content display based on the type.
//
// @param request - PreviewPersonalSourceRequest
//
// @param headers - PreviewPersonalSourceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PreviewPersonalSourceResponse
func (client *Client) PreviewPersonalSourceWithOptions(request *PreviewPersonalSourceRequest, headers *PreviewPersonalSourceHeaders, runtime *dara.RuntimeOptions) (_result *PreviewPersonalSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SourceId) {
		body["sourceId"] = request.SourceId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.RequestId) {
		realHeaders["requestId"] = dara.String(dara.ToString(dara.StringValue(headers.RequestId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PreviewPersonalSource"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/previewPersonalSource"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PreviewPersonalSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Allows a user to preview specified knowledge content in their personal directory.
//
// Description:
//
// ## Request description
//
// - This operation only allows a user to preview resources in their own personal directory.
//
// - The authentication process includes basic authentication and data source ownership verification to ensure that the requester can only access knowledge in their personal directory.
//
// - You must provide the unique identifier `sourceId` of the knowledge content in the request. The system queries and returns the corresponding preview information based on this ID and the user\\"s tenant information.
//
// - Multiple preview types are supported, such as image, audio, and video. The system returns the corresponding preview URL or direct content display based on the type.
//
// @param request - PreviewPersonalSourceRequest
//
// @return PreviewPersonalSourceResponse
func (client *Client) PreviewPersonalSource(request *PreviewPersonalSourceRequest) (_result *PreviewPersonalSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &PreviewPersonalSourceHeaders{}
	_result = &PreviewPersonalSourceResponse{}
	_body, _err := client.PreviewPersonalSourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries primary object data by operating object name with pagination, and supports filtering and searching.
//
// Description:
//
// ## Request description
//
// - This API queries primary object data with pagination based on a specified operating object name (such as `customer_1`).
//
// - Supports keyword-based searching and allows you to specify whether to return only objects marked as favorites.
//
// - Complex filter conditions can be used to further refine results, including but not limited to logical operators such as equal to, not equal to, greater than, and less than.
//
// - If no primary object type is configured, an empty result set is returned.
//
// - Data included in the request undergoes authentication and filtering to ensure security and accuracy.
//
// @param request - QueryPrimaryObjectDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPrimaryObjectDataResponse
func (client *Client) QueryPrimaryObjectDataWithOptions(request *QueryPrimaryObjectDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryPrimaryObjectDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		body["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.OnlyFavorites) {
		body["onlyFavorites"] = request.OnlyFavorites
	}

	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	if !dara.IsNil(request.Page) {
		body["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryPrimaryObjectData"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/queryPrimaryObjectData"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPrimaryObjectDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries primary object data by operating object name with pagination, and supports filtering and searching.
//
// Description:
//
// ## Request description
//
// - This API queries primary object data with pagination based on a specified operating object name (such as `customer_1`).
//
// - Supports keyword-based searching and allows you to specify whether to return only objects marked as favorites.
//
// - Complex filter conditions can be used to further refine results, including but not limited to logical operators such as equal to, not equal to, greater than, and less than.
//
// - If no primary object type is configured, an empty result set is returned.
//
// - Data included in the request undergoes authentication and filtering to ensure security and accuracy.
//
// @param request - QueryPrimaryObjectDataRequest
//
// @return QueryPrimaryObjectDataResponse
func (client *Client) QueryPrimaryObjectData(request *QueryPrimaryObjectDataRequest) (_result *QueryPrimaryObjectDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryPrimaryObjectDataResponse{}
	_body, _err := client.QueryPrimaryObjectDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries semantic knowledge related to a user question.
//
// Description:
//
// Exposes only the schema_knowledge semantic recall capability of smart-query.
//
//	CLI mapping: ``winnexo semantic query``. ``tenantId`` is passed through common parameters. ``userId``
//
//	is read only from the Token identity and cannot be overridden by the request body. The service validates
//
//	the ownership of ``graphName + agentName``, active graph status, digital human enablement status, and
//
//	the current user\\"s USE permission. A cross-graph agent with the same name will fail and be closed.
//
//	Then ``outputs=[schema_knowledge]`` is fixed.
//
// @param request - QuerySemanticKnowledgeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySemanticKnowledgeResponse
func (client *Client) QuerySemanticKnowledgeWithOptions(request *QuerySemanticKnowledgeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QuerySemanticKnowledgeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentName) {
		body["agentName"] = request.AgentName
	}

	if !dara.IsNil(request.GraphName) {
		body["graphName"] = request.GraphName
	}

	if !dara.IsNil(request.Query) {
		body["query"] = request.Query
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySemanticKnowledge"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/querySemanticKnowledge"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySemanticKnowledgeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries semantic knowledge related to a user question.
//
// Description:
//
// Exposes only the schema_knowledge semantic recall capability of smart-query.
//
//	CLI mapping: ``winnexo semantic query``. ``tenantId`` is passed through common parameters. ``userId``
//
//	is read only from the Token identity and cannot be overridden by the request body. The service validates
//
//	the ownership of ``graphName + agentName``, active graph status, digital human enablement status, and
//
//	the current user\\"s USE permission. A cross-graph agent with the same name will fail and be closed.
//
//	Then ``outputs=[schema_knowledge]`` is fixed.
//
// @param request - QuerySemanticKnowledgeRequest
//
// @return QuerySemanticKnowledgeResponse
func (client *Client) QuerySemanticKnowledge(request *QuerySemanticKnowledgeRequest) (_result *QuerySemanticKnowledgeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QuerySemanticKnowledgeResponse{}
	_body, _err := client.QuerySemanticKnowledgeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the result of an organization synchronization task.
//
// Description:
//
// Queries the execution status and result of an organization synchronization task based on the task ID.
//
//	Task status transitions: PENDING → RUNNING → COMPLETED / FAILED / TIMEOUT / CANCELED
//
//	Recommended client polling interval: 3 to 5 seconds.
//
// @param request - QuerySyncResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySyncResultResponse
func (client *Client) QuerySyncResultWithOptions(request *QuerySyncResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QuerySyncResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["taskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySyncResult"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/querySyncResult"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySyncResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the result of an organization synchronization task.
//
// Description:
//
// Queries the execution status and result of an organization synchronization task based on the task ID.
//
//	Task status transitions: PENDING → RUNNING → COMPLETED / FAILED / TIMEOUT / CANCELED
//
//	Recommended client polling interval: 3 to 5 seconds.
//
// @param request - QuerySyncResultRequest
//
// @return QuerySyncResultResponse
func (client *Client) QuerySyncResult(request *QuerySyncResultRequest) (_result *QuerySyncResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QuerySyncResultResponse{}
	_body, _err := client.QuerySyncResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates next-step recommendations for a session.
//
// Description:
//
// ## Request description
//
// Based on the most recent N messages in a session and the skills attached to the agent, this operation invokes an LLM to generate 0 to 3 next-step recommendations (follow-up questions or recommended skills to execute).
//
// - `sessionId`: The session ID. Required. Only sessions that the currently authenticated user has permission to access are allowed.
//
// - `recentMessageCount`: The number of recent messages used to assemble contextual information. Valid values: 1 to 30. Default value: 10 (approximately 5 rounds of user+assistant conversation).
//
// - `customPrompt`: A custom recommendation instruction (up to 10,000 characters). This is injected into the default recommendation template as a custom instruction (before the output format constraints). The output is still subject to the JSON format and type constraints of the template.
//
// - `outputType`: The output type filter. followUpOnly = follow-up recommendations only (default). skillOnly = skill recommendations only. both = generate both types.
//
// Unlike internal endpoints, API calls are not restricted by the next-step recommendation toggle in user personal settings and always execute recommendation generation.
//
// @param request - RecommendNextActionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecommendNextActionsResponse
func (client *Client) RecommendNextActionsWithOptions(request *RecommendNextActionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RecommendNextActionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CustomPrompt) {
		body["customPrompt"] = request.CustomPrompt
	}

	if !dara.IsNil(request.OutputType) {
		body["outputType"] = request.OutputType
	}

	if !dara.IsNil(request.RecentMessageCount) {
		body["recentMessageCount"] = request.RecentMessageCount
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RecommendNextActions"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/recommendNextActions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RecommendNextActionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates next-step recommendations for a session.
//
// Description:
//
// ## Request description
//
// Based on the most recent N messages in a session and the skills attached to the agent, this operation invokes an LLM to generate 0 to 3 next-step recommendations (follow-up questions or recommended skills to execute).
//
// - `sessionId`: The session ID. Required. Only sessions that the currently authenticated user has permission to access are allowed.
//
// - `recentMessageCount`: The number of recent messages used to assemble contextual information. Valid values: 1 to 30. Default value: 10 (approximately 5 rounds of user+assistant conversation).
//
// - `customPrompt`: A custom recommendation instruction (up to 10,000 characters). This is injected into the default recommendation template as a custom instruction (before the output format constraints). The output is still subject to the JSON format and type constraints of the template.
//
// - `outputType`: The output type filter. followUpOnly = follow-up recommendations only (default). skillOnly = skill recommendations only. both = generate both types.
//
// Unlike internal endpoints, API calls are not restricted by the next-step recommendation toggle in user personal settings and always execute recommendation generation.
//
// @param request - RecommendNextActionsRequest
//
// @return RecommendNextActionsResponse
func (client *Client) RecommendNextActions(request *RecommendNextActionsRequest) (_result *RecommendNextActionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RecommendNextActionsResponse{}
	_body, _err := client.RecommendNextActionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a user from a tenant.
//
// Description:
//
// Removes a user from a tenant.
//
//	Business orchestration:
//
//	1. Obtains tenant_id from identity.
//
//	2. Calls delete_user_from_tenant (includes last admin protection).
//
//	3. Returns success.
//
//	This operation:
//
//	- Removes all role associations of the user under the tenant.
//
//	- Removes all user group associations of the user under the tenant.
//
//	- Revokes all digital employee usage authorizations of the user under the tenant.
//
//	- Deletes the user-tenant mapping.
//
// @param request - RemoveUserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUserResponse
func (client *Client) RemoveUserWithOptions(request *RemoveUserRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RemoveUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WnUserId) {
		query["wnUserId"] = request.WnUserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveUser"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/removeUser"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a user from a tenant.
//
// Description:
//
// Removes a user from a tenant.
//
//	Business orchestration:
//
//	1. Obtains tenant_id from identity.
//
//	2. Calls delete_user_from_tenant (includes last admin protection).
//
//	3. Returns success.
//
//	This operation:
//
//	- Removes all role associations of the user under the tenant.
//
//	- Removes all user group associations of the user under the tenant.
//
//	- Revokes all digital employee usage authorizations of the user under the tenant.
//
//	- Deletes the user-tenant mapping.
//
// @param request - RemoveUserRequest
//
// @return RemoveUserResponse
func (client *Client) RemoveUser(request *RemoveUserRequest) (_result *RemoveUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveUserResponse{}
	_body, _err := client.RemoveUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes direct member relationships in bulk from a specified user group.
//
// Description:
//
// ## Request description
//
// - This operation supports batch removal of direct member relationships between users and a specified user group by providing the user group ID and one or more user IDs.
//
// - The `userIds` parameter accepts an integer array that represents the list of platform user IDs to be removed.
//
// - If a user you attempt to remove is not a direct member of the user group, the final result count is not affected.
//
// - After a successful call, the response returns information such as the number of members actually removed and the number of members before the request was processed.
//
// - This operation requires appropriate permission authentication and is recorded in operation logs.
//
// @param tmpReq - RemoveUserGroupMembersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUserGroupMembersResponse
func (client *Client) RemoveUserGroupMembersWithOptions(tmpReq *RemoveUserGroupMembersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RemoveUserGroupMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RemoveUserGroupMembersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserIds) {
		request.UserIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserIds, dara.String("userIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UserGroupId) {
		body["userGroupId"] = request.UserGroupId
	}

	if !dara.IsNil(request.UserIdsShrink) {
		body["userIds"] = request.UserIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveUserGroupMembers"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/removeUserGroupMembers"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveUserGroupMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes direct member relationships in bulk from a specified user group.
//
// Description:
//
// ## Request description
//
// - This operation supports batch removal of direct member relationships between users and a specified user group by providing the user group ID and one or more user IDs.
//
// - The `userIds` parameter accepts an integer array that represents the list of platform user IDs to be removed.
//
// - If a user you attempt to remove is not a direct member of the user group, the final result count is not affected.
//
// - After a successful call, the response returns information such as the number of members actually removed and the number of members before the request was processed.
//
// - This operation requires appropriate permission authentication and is recorded in operation logs.
//
// @param request - RemoveUserGroupMembersRequest
//
// @return RemoveUserGroupMembersResponse
func (client *Client) RemoveUserGroupMembers(request *RemoveUserGroupMembersRequest) (_result *RemoveUserGroupMembersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveUserGroupMembersResponse{}
	_body, _err := client.RemoveUserGroupMembersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Renames a data source in a specified enterprise knowledge base.
//
// Description:
//
// ## Request description
//
// - This operation allows users with the required permissions to modify the name of a specific data source in an enterprise knowledge base.
//
// - You must provide the ID of the data source to be renamed (sourceId) and the new name (newName).
//
// - The rename operation only updates the name field of the data source and does not trigger other processing flows.
//
// - After successful execution, the system publishes a `SOURCE_CHANGED` event for frontend display refresh and attempts to notify DocumentAgent to synchronize the latest source_name information. However, if this step fails, it does not affect the completion status of the main flow.
//
// - If the specified sourceId does not exist, the error code `ERR.Robject.Source.NotFound` is returned.
//
// - To invoke this API, you must have the `DEVELOPMENT_KB_MANAGE` feature permission.
//
// - Identity verification is supported through AccessKey, BearerToken, or APP methods to authenticate requests.
//
// @param request - RenameKnowledgeBaseSourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenameKnowledgeBaseSourceResponse
func (client *Client) RenameKnowledgeBaseSourceWithOptions(request *RenameKnowledgeBaseSourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RenameKnowledgeBaseSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.NewName) {
		body["newName"] = request.NewName
	}

	if !dara.IsNil(request.SourceId) {
		body["sourceId"] = request.SourceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenameKnowledgeBaseSource"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/renameKnowledgeBaseSource"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenameKnowledgeBaseSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renames a data source in a specified enterprise knowledge base.
//
// Description:
//
// ## Request description
//
// - This operation allows users with the required permissions to modify the name of a specific data source in an enterprise knowledge base.
//
// - You must provide the ID of the data source to be renamed (sourceId) and the new name (newName).
//
// - The rename operation only updates the name field of the data source and does not trigger other processing flows.
//
// - After successful execution, the system publishes a `SOURCE_CHANGED` event for frontend display refresh and attempts to notify DocumentAgent to synchronize the latest source_name information. However, if this step fails, it does not affect the completion status of the main flow.
//
// - If the specified sourceId does not exist, the error code `ERR.Robject.Source.NotFound` is returned.
//
// - To invoke this API, you must have the `DEVELOPMENT_KB_MANAGE` feature permission.
//
// - Identity verification is supported through AccessKey, BearerToken, or APP methods to authenticate requests.
//
// @param request - RenameKnowledgeBaseSourceRequest
//
// @return RenameKnowledgeBaseSourceResponse
func (client *Client) RenameKnowledgeBaseSource(request *RenameKnowledgeBaseSourceRequest) (_result *RenameKnowledgeBaseSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RenameKnowledgeBaseSourceResponse{}
	_body, _err := client.RenameKnowledgeBaseSourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Renames a specified data source. This is a lightweight operation.
//
// Description:
//
// ## Operation description
//
// - This API updates only the `name` field of the data source and does not trigger `process_source`.
//
// - After a successful update, a `SOURCE_CHANGED` event is published for the frontend to refresh the display.
//
// - The system makes a best-effort attempt to notify DocumentAgent to sync the new `source_name`. Even if the sync fails, the main process is not blocked.
//
// - If the specified data source does not exist, the `ERR.Robject.Source.NotFound` error is returned. The global middleware converts this error into a POP error code.
//
// - Security constraint: `tenant_id` and `user_id` must be derived from the authenticated identity.
//
// @param request - RenameSourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenameSourceResponse
func (client *Client) RenameSourceWithOptions(request *RenameSourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RenameSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.NewName) {
		body["newName"] = request.NewName
	}

	if !dara.IsNil(request.SourceId) {
		body["sourceId"] = request.SourceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenameSource"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/renameSource"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenameSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renames a specified data source. This is a lightweight operation.
//
// Description:
//
// ## Operation description
//
// - This API updates only the `name` field of the data source and does not trigger `process_source`.
//
// - After a successful update, a `SOURCE_CHANGED` event is published for the frontend to refresh the display.
//
// - The system makes a best-effort attempt to notify DocumentAgent to sync the new `source_name`. Even if the sync fails, the main process is not blocked.
//
// - If the specified data source does not exist, the `ERR.Robject.Source.NotFound` error is returned. The global middleware converts this error into a POP error code.
//
// - Security constraint: `tenant_id` and `user_id` must be derived from the authenticated identity.
//
// @param request - RenameSourceRequest
//
// @return RenameSourceResponse
func (client *Client) RenameSource(request *RenameSourceRequest) (_result *RenameSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RenameSourceResponse{}
	_body, _err := client.RenameSourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Re-parses a resource.
//
// Description:
//
// ## Operation description
//
// This API operation re-parses a specified data source. You can choose synchronous or asynchronous execution. You must provide the data source ID in the request. You can optionally specify whether to synchronously wait for parsing to complete. By default, the request is processed asynchronously by being added to a queue. You can also use the `tenantId` parameter to specify a tenant ID, but this parameter is optional.
//
// - **forceSync**: If set to `true`, the operation synchronously waits for the re-parsing to complete. Default value: `false`, which indicates that the request is processed asynchronously.
//
// - When the service returns `None`, it is converted to a `SourceNotFound` exception. Other exceptions are handled by the OpenAPI global exception chain.
//
// @param request - ReparseSourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReparseSourceResponse
func (client *Client) ReparseSourceWithOptions(request *ReparseSourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ReparseSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ForceSync) {
		body["forceSync"] = request.ForceSync
	}

	if !dara.IsNil(request.SourceId) {
		body["sourceId"] = request.SourceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReparseSource"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/reparseSource"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReparseSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Re-parses a resource.
//
// Description:
//
// ## Operation description
//
// This API operation re-parses a specified data source. You can choose synchronous or asynchronous execution. You must provide the data source ID in the request. You can optionally specify whether to synchronously wait for parsing to complete. By default, the request is processed asynchronously by being added to a queue. You can also use the `tenantId` parameter to specify a tenant ID, but this parameter is optional.
//
// - **forceSync**: If set to `true`, the operation synchronously waits for the re-parsing to complete. Default value: `false`, which indicates that the request is processed asynchronously.
//
// - When the service returns `None`, it is converted to a `SourceNotFound` exception. Other exceptions are handled by the OpenAPI global exception chain.
//
// @param request - ReparseSourceRequest
//
// @return ReparseSourceResponse
func (client *Client) ReparseSource(request *ReparseSourceRequest) (_result *ReparseSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReparseSourceResponse{}
	_body, _err := client.ReparseSourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Replaces a FILE resource in a specified enterprise knowledge base and triggers re-parsing.
//
// Description:
//
// ## Request description
//
// This API allows you to update a specific FILE-type data source in a self-built enterprise knowledge base and trigger the system to re-parse the data source by providing a new file path and public access URL. Operations can be performed in synchronous or asynchronous mode. In synchronous mode, the client waits until the parsing process is complete.
//
// - The **forceSync*	- parameter controls whether the request is processed synchronously. The default value is `false`, which indicates asynchronous processing.
//
// - If **fileName*	- is not provided or its value is empty, the newly uploaded file retains the original file name.
//
// - Ensure that the provided **filePath*	- and **filePublicUrl*	- are valid and point to the same file entity.
//
// @param request - ReplaceKnowledgeBaseSourceFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReplaceKnowledgeBaseSourceFileResponse
func (client *Client) ReplaceKnowledgeBaseSourceFileWithOptions(request *ReplaceKnowledgeBaseSourceFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ReplaceKnowledgeBaseSourceFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		body["fileName"] = request.FileName
	}

	if !dara.IsNil(request.FilePath) {
		body["filePath"] = request.FilePath
	}

	if !dara.IsNil(request.FilePublicUrl) {
		body["filePublicUrl"] = request.FilePublicUrl
	}

	if !dara.IsNil(request.FileRecordId) {
		body["fileRecordId"] = request.FileRecordId
	}

	if !dara.IsNil(request.ForceSync) {
		body["forceSync"] = request.ForceSync
	}

	if !dara.IsNil(request.SourceId) {
		body["sourceId"] = request.SourceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReplaceKnowledgeBaseSourceFile"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/replaceKnowledgeBaseSourceFile"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReplaceKnowledgeBaseSourceFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Replaces a FILE resource in a specified enterprise knowledge base and triggers re-parsing.
//
// Description:
//
// ## Request description
//
// This API allows you to update a specific FILE-type data source in a self-built enterprise knowledge base and trigger the system to re-parse the data source by providing a new file path and public access URL. Operations can be performed in synchronous or asynchronous mode. In synchronous mode, the client waits until the parsing process is complete.
//
// - The **forceSync*	- parameter controls whether the request is processed synchronously. The default value is `false`, which indicates asynchronous processing.
//
// - If **fileName*	- is not provided or its value is empty, the newly uploaded file retains the original file name.
//
// - Ensure that the provided **filePath*	- and **filePublicUrl*	- are valid and point to the same file entity.
//
// @param request - ReplaceKnowledgeBaseSourceFileRequest
//
// @return ReplaceKnowledgeBaseSourceFileResponse
func (client *Client) ReplaceKnowledgeBaseSourceFile(request *ReplaceKnowledgeBaseSourceFileRequest) (_result *ReplaceKnowledgeBaseSourceFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReplaceKnowledgeBaseSourceFileResponse{}
	_body, _err := client.ReplaceKnowledgeBaseSourceFileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Replaces all object binding information under a specified data source.
//
// Description:
//
// ## Operation description
//
// This API performs a full replacement of object bindings for a specified data source (deletes existing bindings first, then inserts new bindings). If an empty list is passed, all bindings are cleared.
//
// - **Security constraints**: `tenant_id` and `user_id` must come from the authenticated identity.
//
// - **Error handling**: If the specified data source does not exist, an `ERR.Robject.InvalidParameter` error is thrown and converted to a POP error code by the global middleware.
//
// - **Synchronous notification**: After a successful replacement, the system makes a best-effort synchronous notification to DocumentAgent to update `semantics.object_bindings`. However, failures are only logged and do not block the main process.
//
// @param tmpReq - ReplaceObjectBindingsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReplaceObjectBindingsResponse
func (client *Client) ReplaceObjectBindingsWithOptions(tmpReq *ReplaceObjectBindingsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ReplaceObjectBindingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ReplaceObjectBindingsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ObjectBindings) {
		request.ObjectBindingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ObjectBindings, dara.String("objectBindings"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ObjectBindingsShrink) {
		body["objectBindings"] = request.ObjectBindingsShrink
	}

	if !dara.IsNil(request.SourceId) {
		body["sourceId"] = request.SourceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReplaceObjectBindings"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/replaceObjectBindings"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReplaceObjectBindingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Replaces all object binding information under a specified data source.
//
// Description:
//
// ## Operation description
//
// This API performs a full replacement of object bindings for a specified data source (deletes existing bindings first, then inserts new bindings). If an empty list is passed, all bindings are cleared.
//
// - **Security constraints**: `tenant_id` and `user_id` must come from the authenticated identity.
//
// - **Error handling**: If the specified data source does not exist, an `ERR.Robject.InvalidParameter` error is thrown and converted to a POP error code by the global middleware.
//
// - **Synchronous notification**: After a successful replacement, the system makes a best-effort synchronous notification to DocumentAgent to update `semantics.object_bindings`. However, failures are only logged and do not block the main process.
//
// @param request - ReplaceObjectBindingsRequest
//
// @return ReplaceObjectBindingsResponse
func (client *Client) ReplaceObjectBindings(request *ReplaceObjectBindingsRequest) (_result *ReplaceObjectBindingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReplaceObjectBindingsResponse{}
	_body, _err := client.ReplaceObjectBindingsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Allows a user to replace a personal file resource that they created and triggers the system to re-parse the file.
//
// Description:
//
// ## Operation description
//
// - This API operation replaces a personal FILE resource created by the current platform user and triggers the system to re-parse the file.
//
// - The `tenant_id`, operator, and creator constraints are read only from the authenticated identity. Requests without a platform user are rejected to prevent bypassing ownership verification.
//
// - If the server returns `None`, it is converted to a `NotFound` exception. Other exceptions are handled by the OpenAPI global exception chain.
//
// - This operation supports synchronous or asynchronous waiting for re-parsing to complete. The default behavior is asynchronous queuing (controlled by the `forceSync` parameter).
//
// @param request - ReplaceSourceFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReplaceSourceFileResponse
func (client *Client) ReplaceSourceFileWithOptions(request *ReplaceSourceFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ReplaceSourceFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		body["fileName"] = request.FileName
	}

	if !dara.IsNil(request.FilePath) {
		body["filePath"] = request.FilePath
	}

	if !dara.IsNil(request.FilePublicUrl) {
		body["filePublicUrl"] = request.FilePublicUrl
	}

	if !dara.IsNil(request.FileRecordId) {
		body["fileRecordId"] = request.FileRecordId
	}

	if !dara.IsNil(request.ForceSync) {
		body["forceSync"] = request.ForceSync
	}

	if !dara.IsNil(request.SourceId) {
		body["sourceId"] = request.SourceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReplaceSourceFile"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/replaceSourceFile"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReplaceSourceFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Allows a user to replace a personal file resource that they created and triggers the system to re-parse the file.
//
// Description:
//
// ## Operation description
//
// - This API operation replaces a personal FILE resource created by the current platform user and triggers the system to re-parse the file.
//
// - The `tenant_id`, operator, and creator constraints are read only from the authenticated identity. Requests without a platform user are rejected to prevent bypassing ownership verification.
//
// - If the server returns `None`, it is converted to a `NotFound` exception. Other exceptions are handled by the OpenAPI global exception chain.
//
// - This operation supports synchronous or asynchronous waiting for re-parsing to complete. The default behavior is asynchronous queuing (controlled by the `forceSync` parameter).
//
// @param request - ReplaceSourceFileRequest
//
// @return ReplaceSourceFileResponse
func (client *Client) ReplaceSourceFile(request *ReplaceSourceFileRequest) (_result *ReplaceSourceFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReplaceSourceFileResponse{}
	_body, _err := client.ReplaceSourceFileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resets the password of a user.
//
// Description:
//
// Resets the password of a user through OpenAPI.
//
//	Business orchestration:
//
//	1. Call UserManagementService.reset_member_password with password_encrypted (required).
//
//	   The service internally performs RSA decryption, complexity validation, bcrypt hashing, and writes the result.
//
//	2. Returns the reset result.
//
//	Error codes:
//
//	- ERR.User.NotFound: The user does not exist.
//
//	- ERR.User.NotInTenant: The user does not belong to the current tenant.
//
//	- ERR.User.WinnexoPasswordRequired: The user does not have password credentials (non-WINNEXO type).
//
// @param request - ResetPasswordRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetPasswordResponse
func (client *Client) ResetPasswordWithOptions(request *ResetPasswordRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ResetPasswordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PasswordEncrypted) {
		body["passwordEncrypted"] = request.PasswordEncrypted
	}

	if !dara.IsNil(request.WnUserId) {
		body["wnUserId"] = request.WnUserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetPassword"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/resetPassword"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the password of a user.
//
// Description:
//
// Resets the password of a user through OpenAPI.
//
//	Business orchestration:
//
//	1. Call UserManagementService.reset_member_password with password_encrypted (required).
//
//	   The service internally performs RSA decryption, complexity validation, bcrypt hashing, and writes the result.
//
//	2. Returns the reset result.
//
//	Error codes:
//
//	- ERR.User.NotFound: The user does not exist.
//
//	- ERR.User.NotInTenant: The user does not belong to the current tenant.
//
//	- ERR.User.WinnexoPasswordRequired: The user does not have password credentials (non-WINNEXO type).
//
// @param request - ResetPasswordRequest
//
// @return ResetPasswordResponse
func (client *Client) ResetPassword(request *ResetPasswordRequest) (_result *ResetPasswordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ResetPasswordResponse{}
	_body, _err := client.ResetPasswordWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resets an API token.
//
// Description:
//
// Resets the token of a user.
//
//	Business logic:
//
//	1. Retrieves user_id from identity (caller_type=user is required).
//
//	2. Constructs an AuthContext and delegates permission verification to UserTokenAuthorizedService.
//
//	3. Calls reset_token:
//
//	   - Changes the old ACTIVE token to RESET (permanently invalidated).
//
//	   - Generates a new ACTIVE token.
//
//	4. Returns the new token in plaintext and the masked value.
//
//	Note: After the reset, the old token is permanently invalidated and cannot be recovered. The new token in plaintext is returned only in this response.
//
// @param request - ResetTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetTokenResponse
func (client *Client) ResetTokenWithOptions(request *ResetTokenRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ResetTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.WnUserId) {
		body["wnUserId"] = request.WnUserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetToken"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/resetToken"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets an API token.
//
// Description:
//
// Resets the token of a user.
//
//	Business logic:
//
//	1. Retrieves user_id from identity (caller_type=user is required).
//
//	2. Constructs an AuthContext and delegates permission verification to UserTokenAuthorizedService.
//
//	3. Calls reset_token:
//
//	   - Changes the old ACTIVE token to RESET (permanently invalidated).
//
//	   - Generates a new ACTIVE token.
//
//	4. Returns the new token in plaintext and the masked value.
//
//	Note: After the reset, the old token is permanently invalidated and cannot be recovered. The new token in plaintext is returned only in this response.
//
// @param request - ResetTokenRequest
//
// @return ResetTokenResponse
func (client *Client) ResetToken(request *ResetTokenRequest) (_result *ResetTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ResetTokenResponse{}
	_body, _err := client.ResetTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retries all failed data sources in a specified folder and its subfolders in batch.
//
// Description:
//
// ## Request description
//
// This API retrieves and retries data sources with a FAILED status in the specified personal folder of a user (including all subfolders). The request returns immediately, and the actual retry tasks are executed asynchronously in the background. Only resources that the current logged-in user has access to and that were created by the user can be retried.
//
// ### Security and permissions
//
// - This operation requires appropriate RAM permissions.
//
// - You can only operate on resources within the tenant to which the current user belongs.
//
// - Ensure that `tenantId` and `userId` come from verified identity information.
//
// ### Precautions
//
// - `directoryId` is a required parameter that specifies the target folder in which to check and retry failed data sources.
//
// - If `tenantId` is not provided, the caller\\"s tenant ID is used by default.
//
// - The API supports multiple authentication methods, including AccessKey, BearerToken, and APP authentication.
//
// @param request - RetryDirectoryFailedSourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetryDirectoryFailedSourcesResponse
func (client *Client) RetryDirectoryFailedSourcesWithOptions(request *RetryDirectoryFailedSourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RetryDirectoryFailedSourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		body["directoryId"] = request.DirectoryId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RetryDirectoryFailedSources"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/retryDirectoryFailedSources"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RetryDirectoryFailedSourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retries all failed data sources in a specified folder and its subfolders in batch.
//
// Description:
//
// ## Request description
//
// This API retrieves and retries data sources with a FAILED status in the specified personal folder of a user (including all subfolders). The request returns immediately, and the actual retry tasks are executed asynchronously in the background. Only resources that the current logged-in user has access to and that were created by the user can be retried.
//
// ### Security and permissions
//
// - This operation requires appropriate RAM permissions.
//
// - You can only operate on resources within the tenant to which the current user belongs.
//
// - Ensure that `tenantId` and `userId` come from verified identity information.
//
// ### Precautions
//
// - `directoryId` is a required parameter that specifies the target folder in which to check and retry failed data sources.
//
// - If `tenantId` is not provided, the caller\\"s tenant ID is used by default.
//
// - The API supports multiple authentication methods, including AccessKey, BearerToken, and APP authentication.
//
// @param request - RetryDirectoryFailedSourcesRequest
//
// @return RetryDirectoryFailedSourcesResponse
func (client *Client) RetryDirectoryFailedSources(request *RetryDirectoryFailedSourcesRequest) (_result *RetryDirectoryFailedSourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RetryDirectoryFailedSourcesResponse{}
	_body, _err := client.RetryDirectoryFailedSourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retries all data sources in failed status under a specified directory in batches.
//
// Description:
//
// ## Operation description
//
// This API retrieves and retries all data sources in FAILED status under a specified enterprise knowledge base directory (including its subdirectories). The request returns immediately, and the actual retry operations are executed asynchronously in the background.
//
// - **Authentication**: In addition to basic authentication, the `DEVELOPMENT_KB_MANAGE` permission is required.
//
// - **Security constraints**: Only callers with the corresponding tenant and user identity are allowed access, and KB management permission is required. Administrators can initiate retries for failed resources of any user.
//
// - **Parameters**:
//
//   - `directoryId` (required): The ID of the enterprise knowledge base directory to check and retry failed data sources.
//
//   - `tenantId` (optional): The tenant ID. The default tenant of the caller is used if this parameter is not specified.
//
// - **Response**: On success, returns the number of data sources enqueued for retry and related details.
//
// @param request - RetryKnowledgeBaseFailedSourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetryKnowledgeBaseFailedSourcesResponse
func (client *Client) RetryKnowledgeBaseFailedSourcesWithOptions(request *RetryKnowledgeBaseFailedSourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RetryKnowledgeBaseFailedSourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		body["directoryId"] = request.DirectoryId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RetryKnowledgeBaseFailedSources"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/retryKnowledgeBaseFailedSources"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RetryKnowledgeBaseFailedSourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retries all data sources in failed status under a specified directory in batches.
//
// Description:
//
// ## Operation description
//
// This API retrieves and retries all data sources in FAILED status under a specified enterprise knowledge base directory (including its subdirectories). The request returns immediately, and the actual retry operations are executed asynchronously in the background.
//
// - **Authentication**: In addition to basic authentication, the `DEVELOPMENT_KB_MANAGE` permission is required.
//
// - **Security constraints**: Only callers with the corresponding tenant and user identity are allowed access, and KB management permission is required. Administrators can initiate retries for failed resources of any user.
//
// - **Parameters**:
//
//   - `directoryId` (required): The ID of the enterprise knowledge base directory to check and retry failed data sources.
//
//   - `tenantId` (optional): The tenant ID. The default tenant of the caller is used if this parameter is not specified.
//
// - **Response**: On success, returns the number of data sources enqueued for retry and related details.
//
// @param request - RetryKnowledgeBaseFailedSourcesRequest
//
// @return RetryKnowledgeBaseFailedSourcesResponse
func (client *Client) RetryKnowledgeBaseFailedSources(request *RetryKnowledgeBaseFailedSourcesRequest) (_result *RetryKnowledgeBaseFailedSourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RetryKnowledgeBaseFailedSourcesResponse{}
	_body, _err := client.RetryKnowledgeBaseFailedSourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes the usage permissions of a user or user group on a digital human.
//
// Description:
//
// Revokes the usage permissions of a user or user group on a specified digital human.
//
//	Business logic:
//
//	1. Constructs an AuthContext from identity.
//
//	2. Performs mutual exclusion validation on the request body: either userIds or userGroupIds must be specified.
//
//	3. Delegates to AgentAuthorizationAuthorizedService.revoke_authorization for execution.
//
//	4. Pre-validation: MANAGE permission + agent existence check (performed by the AuthorizedService layer, which authenticates before exposing existence).
//
//	5. After direct user authorization is revoked, the user may still have access through user group authorization.
//
// @param tmpReq - RevokeAgentUsersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeAgentUsersResponse
func (client *Client) RevokeAgentUsersWithOptions(tmpReq *RevokeAgentUsersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RevokeAgentUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RevokeAgentUsersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserGroupIds) {
		request.UserGroupIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserGroupIds, dara.String("userGroupIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserIds) {
		request.UserIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserIds, dara.String("userIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	if !dara.IsNil(request.UserGroupIdsShrink) {
		body["userGroupIds"] = request.UserGroupIdsShrink
	}

	if !dara.IsNil(request.UserIdsShrink) {
		body["userIds"] = request.UserIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeAgentUsers"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/revokeAgentUsers"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeAgentUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes the usage permissions of a user or user group on a digital human.
//
// Description:
//
// Revokes the usage permissions of a user or user group on a specified digital human.
//
//	Business logic:
//
//	1. Constructs an AuthContext from identity.
//
//	2. Performs mutual exclusion validation on the request body: either userIds or userGroupIds must be specified.
//
//	3. Delegates to AgentAuthorizationAuthorizedService.revoke_authorization for execution.
//
//	4. Pre-validation: MANAGE permission + agent existence check (performed by the AuthorizedService layer, which authenticates before exposing existence).
//
//	5. After direct user authorization is revoked, the user may still have access through user group authorization.
//
// @param request - RevokeAgentUsersRequest
//
// @return RevokeAgentUsersResponse
func (client *Client) RevokeAgentUsers(request *RevokeAgentUsersRequest) (_result *RevokeAgentUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RevokeAgentUsersResponse{}
	_body, _err := client.RevokeAgentUsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Asynchronously triggers skill execution and immediately returns a RunId.
//
// Description:
//
// ## Request description
//
// This operation supports only asynchronous mode. After submission, the operation immediately returns a `RunId` and `Status=Running`. The client polls for the final result by calling `GetSkillRun`.
//
// - **TenantId**: An optional common parameter that the gateway passes through to the backend header.
//
// - **SkillCode*	- / **SkillName**: Specify one of the two parameters. SkillCode takes priority. If SkillName is not unique, `ERR.SkillHub.SkillNameAmbiguous` is returned.
//
// - **Arguments**: Required. The skill input parameter object. The structure is described by the inputConfig returned by `GetSkill`.
//
// - **ClientToken**: An optional idempotency key. In the current version, this value is only recorded in the task metadata and is not used for strict idempotency deduplication.
//
// Note: Synchronous mode (Async=false), Stream, and CallbackUrl are not supported in the first release and will be available in later versions.
//
// @param tmpReq - RunSkillRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunSkillResponse
func (client *Client) RunSkillWithOptions(tmpReq *RunSkillRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunSkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunSkillShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Arguments) {
		request.ArgumentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Arguments, dara.String("arguments"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ArgumentsShrink) {
		body["arguments"] = request.ArgumentsShrink
	}

	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Model) {
		body["model"] = request.Model
	}

	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	if !dara.IsNil(request.SkillCode) {
		body["skillCode"] = request.SkillCode
	}

	if !dara.IsNil(request.SkillName) {
		body["skillName"] = request.SkillName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunSkill"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/runSkill"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunSkillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Asynchronously triggers skill execution and immediately returns a RunId.
//
// Description:
//
// ## Request description
//
// This operation supports only asynchronous mode. After submission, the operation immediately returns a `RunId` and `Status=Running`. The client polls for the final result by calling `GetSkillRun`.
//
// - **TenantId**: An optional common parameter that the gateway passes through to the backend header.
//
// - **SkillCode*	- / **SkillName**: Specify one of the two parameters. SkillCode takes priority. If SkillName is not unique, `ERR.SkillHub.SkillNameAmbiguous` is returned.
//
// - **Arguments**: Required. The skill input parameter object. The structure is described by the inputConfig returned by `GetSkill`.
//
// - **ClientToken**: An optional idempotency key. In the current version, this value is only recorded in the task metadata and is not used for strict idempotency deduplication.
//
// Note: Synchronous mode (Async=false), Stream, and CallbackUrl are not supported in the first release and will be available in later versions.
//
// @param request - RunSkillRequest
//
// @return RunSkillResponse
func (client *Client) RunSkill(request *RunSkillRequest) (_result *RunSkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunSkillResponse{}
	_body, _err := client.RunSkillWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Saves group outputs in batches to the collaboration group repository.
//
// Description:
//
// ## Request description
//
// - Saves specified group outputs to the repository directory of the same collaboration group.
//
// - Supports two modes: `link` (maintains output association) and `copy` (creates an independent snapshot).
//
// - The caller must be a platform user and a member of the target group. The caller can archive group outputs visible to them, including outputs created by other members.
//
// - If `directoryId` is not specified, the default repository directory of the target group is used.
//
// - A maximum of 50 outputs can be processed per batch. All entries are validated before saving. If any entry does not exist, is not visible, or cannot be operated on, the entire batch fails.
//
// - After unified validation passes, entries are saved one by one. The response results maintain the same order as `itemIds`. A failure of a single entry does not affect other entries.
//
// @param tmpReq - SaveGroupOutputFileToGroupResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveGroupOutputFileToGroupResourceResponse
func (client *Client) SaveGroupOutputFileToGroupResourceWithOptions(tmpReq *SaveGroupOutputFileToGroupResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SaveGroupOutputFileToGroupResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SaveGroupOutputFileToGroupResourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ItemIds) {
		request.ItemIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ItemIds, dara.String("itemIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		body["directoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.GroupId) {
		body["groupId"] = request.GroupId
	}

	if !dara.IsNil(request.ItemIdsShrink) {
		body["itemIds"] = request.ItemIdsShrink
	}

	if !dara.IsNil(request.Mode) {
		body["mode"] = request.Mode
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveGroupOutputFileToGroupResource"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/saveGroupOutputFileToGroupResource"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveGroupOutputFileToGroupResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Saves group outputs in batches to the collaboration group repository.
//
// Description:
//
// ## Request description
//
// - Saves specified group outputs to the repository directory of the same collaboration group.
//
// - Supports two modes: `link` (maintains output association) and `copy` (creates an independent snapshot).
//
// - The caller must be a platform user and a member of the target group. The caller can archive group outputs visible to them, including outputs created by other members.
//
// - If `directoryId` is not specified, the default repository directory of the target group is used.
//
// - A maximum of 50 outputs can be processed per batch. All entries are validated before saving. If any entry does not exist, is not visible, or cannot be operated on, the entire batch fails.
//
// - After unified validation passes, entries are saved one by one. The response results maintain the same order as `itemIds`. A failure of a single entry does not affect other entries.
//
// @param request - SaveGroupOutputFileToGroupResourceRequest
//
// @return SaveGroupOutputFileToGroupResourceResponse
func (client *Client) SaveGroupOutputFileToGroupResource(request *SaveGroupOutputFileToGroupResourceRequest) (_result *SaveGroupOutputFileToGroupResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SaveGroupOutputFileToGroupResourceResponse{}
	_body, _err := client.SaveGroupOutputFileToGroupResourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Batch saves group outputs to the current operator\\"s personal knowledge base.
//
// Description:
//
// ## Request description
//
// - Saves specified group outputs to the current operator\\"s personal knowledge base.
//
// - Supports two modes: `link` (maintains output association) and `copy` (creates an independent snapshot).
//
// - The caller must be a member of the target group who is associated with a platform user. Regular members can only archive outputs they created, while group administrators can archive visible outputs from other members. Personal ownership is always derived from the gateway authentication identity.
//
// - If `directoryId` is not specified, the current operator\\"s default personal directory is used.
//
// - A maximum of 50 outputs can be processed per batch. All entries are validated before saving. The entire batch fails if any entry does not exist, is not visible, or cannot be operated on.
//
// - After unified validation passes, entries are saved one by one. The response results maintain the same order as `itemIds`. A failure to save a single entry does not affect other entries.
//
// @param tmpReq - SaveGroupOutputFileToPersonalResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveGroupOutputFileToPersonalResourceResponse
func (client *Client) SaveGroupOutputFileToPersonalResourceWithOptions(tmpReq *SaveGroupOutputFileToPersonalResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SaveGroupOutputFileToPersonalResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SaveGroupOutputFileToPersonalResourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ItemIds) {
		request.ItemIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ItemIds, dara.String("itemIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		body["directoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.GroupId) {
		body["groupId"] = request.GroupId
	}

	if !dara.IsNil(request.ItemIdsShrink) {
		body["itemIds"] = request.ItemIdsShrink
	}

	if !dara.IsNil(request.Mode) {
		body["mode"] = request.Mode
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveGroupOutputFileToPersonalResource"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/saveGroupOutputFileToPersonalResource"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveGroupOutputFileToPersonalResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Batch saves group outputs to the current operator\\"s personal knowledge base.
//
// Description:
//
// ## Request description
//
// - Saves specified group outputs to the current operator\\"s personal knowledge base.
//
// - Supports two modes: `link` (maintains output association) and `copy` (creates an independent snapshot).
//
// - The caller must be a member of the target group who is associated with a platform user. Regular members can only archive outputs they created, while group administrators can archive visible outputs from other members. Personal ownership is always derived from the gateway authentication identity.
//
// - If `directoryId` is not specified, the current operator\\"s default personal directory is used.
//
// - A maximum of 50 outputs can be processed per batch. All entries are validated before saving. The entire batch fails if any entry does not exist, is not visible, or cannot be operated on.
//
// - After unified validation passes, entries are saved one by one. The response results maintain the same order as `itemIds`. A failure to save a single entry does not affect other entries.
//
// @param request - SaveGroupOutputFileToPersonalResourceRequest
//
// @return SaveGroupOutputFileToPersonalResourceResponse
func (client *Client) SaveGroupOutputFileToPersonalResource(request *SaveGroupOutputFileToPersonalResourceRequest) (_result *SaveGroupOutputFileToPersonalResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SaveGroupOutputFileToPersonalResourceResponse{}
	_body, _err := client.SaveGroupOutputFileToPersonalResourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Saves output details in batch as personal resources. Supports link or copy mode.
//
// Description:
//
// ## Operation description
//
// - This API saves a batch of output details as personal resources for the user.
//
// - Two save modes are supported: `link` and `copy`. When `link` is selected, edits to the output are synchronized to the resource. When `copy` is selected, a snapshot is created with no limit on the number of copies.
//
// - `tenant_id` and `user_id` are derived only from the authenticated identity.
//
// - If `operating_object` values are inconsistent within the batch and `directoryId` is not specified, the entire batch fails with a pre-check error.
//
// - The processing result of a single record does not affect other records. Failure information for individual records is returned in the response.
//
// - A maximum of 50 records are supported per batch operation.
//
// - Batch-level pre-check failures are returned in a POP-compatible error format by the global exception middleware.
//
// @param tmpReq - SaveOutputFileToResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveOutputFileToResourceResponse
func (client *Client) SaveOutputFileToResourceWithOptions(tmpReq *SaveOutputFileToResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SaveOutputFileToResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SaveOutputFileToResourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ItemIds) {
		request.ItemIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ItemIds, dara.String("itemIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		body["directoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.ItemIdsShrink) {
		body["itemIds"] = request.ItemIdsShrink
	}

	if !dara.IsNil(request.Mode) {
		body["mode"] = request.Mode
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveOutputFileToResource"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/saveOutputFileToResource"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveOutputFileToResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Saves output details in batch as personal resources. Supports link or copy mode.
//
// Description:
//
// ## Operation description
//
// - This API saves a batch of output details as personal resources for the user.
//
// - Two save modes are supported: `link` and `copy`. When `link` is selected, edits to the output are synchronized to the resource. When `copy` is selected, a snapshot is created with no limit on the number of copies.
//
// - `tenant_id` and `user_id` are derived only from the authenticated identity.
//
// - If `operating_object` values are inconsistent within the batch and `directoryId` is not specified, the entire batch fails with a pre-check error.
//
// - The processing result of a single record does not affect other records. Failure information for individual records is returned in the response.
//
// - A maximum of 50 records are supported per batch operation.
//
// - Batch-level pre-check failures are returned in a POP-compatible error format by the global exception middleware.
//
// @param request - SaveOutputFileToResourceRequest
//
// @return SaveOutputFileToResourceResponse
func (client *Client) SaveOutputFileToResource(request *SaveOutputFileToResourceRequest) (_result *SaveOutputFileToResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SaveOutputFileToResourceResponse{}
	_body, _err := client.SaveOutputFileToResourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Asynchronously sends a session message.
//
// Description:
//
// Asynchronously sends a session message.
//
// @param tmpReq - SendAsyncChatMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendAsyncChatMessageResponse
func (client *Client) SendAsyncChatMessageWithOptions(tmpReq *SendAsyncChatMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SendAsyncChatMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SendAsyncChatMessageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DigitalEmployeeName) {
		request.DigitalEmployeeNameShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DigitalEmployeeName, dara.String("digitalEmployeeName"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Files) {
		request.FilesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Files, dara.String("files"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TaskExecution) {
		request.TaskExecutionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskExecution, dara.String("taskExecution"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	if !dara.IsNil(request.ContentType) {
		body["contentType"] = request.ContentType
	}

	if !dara.IsNil(request.DigitalEmployeeNameShrink) {
		body["digitalEmployeeName"] = request.DigitalEmployeeNameShrink
	}

	if !dara.IsNil(request.DirectChat) {
		body["directChat"] = request.DirectChat
	}

	if !dara.IsNil(request.FilesShrink) {
		body["files"] = request.FilesShrink
	}

	if !dara.IsNil(request.Model) {
		body["model"] = request.Model
	}

	if !dara.IsNil(request.ReuseLastSession) {
		body["reuseLastSession"] = request.ReuseLastSession
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	if !dara.IsNil(request.TaskExecutionShrink) {
		body["taskExecution"] = request.TaskExecutionShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendAsyncChatMessage"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/sendAsyncChatMessage"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendAsyncChatMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Asynchronously sends a session message.
//
// Description:
//
// Asynchronously sends a session message.
//
// @param request - SendAsyncChatMessageRequest
//
// @return SendAsyncChatMessageResponse
func (client *Client) SendAsyncChatMessage(request *SendAsyncChatMessageRequest) (_result *SendAsyncChatMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SendAsyncChatMessageResponse{}
	_body, _err := client.SendAsyncChatMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends a message.
//
// Description:
//
// ## Request description
//
// - This API is used to upload a file to the "My Resources" section of a specified digital employee.
//
// - `source_type` is fixed to `FILE`, `scope` is fixed to `PERSONAL`, and `platform` is fixed to `LOCAL`.
//
// - The file must include an OSS persistent address (`filePath`). Other information such as the public access URL and original file name is optional.
//
// - If no target folder ID (`directoryId`) is specified, the file is automatically attached to the default root folder of the current digital employee. If specified, ensure that the folder belongs to the invoker\\"s personal folder.
//
// - Multiple authentication methods (AK, BearerToken, APP) are supported to authenticate requests.
//
// - The operation type is write, and operation logs are recorded for subsequent auditing.
//
// @param tmpReq - SendChatMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendChatMessageResponse
func (client *Client) SendChatMessageWithSSE(tmpReq *SendChatMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *SendChatMessageResponse, _yieldErr chan error) {
	defer close(_yield)
	client.sendChatMessageWithSSE_opYieldFunc(_yield, _yieldErr, tmpReq, headers, runtime)
	return
}

// Summary:
//
// Sends a message.
//
// Description:
//
// ## Request description
//
// - This API is used to upload a file to the "My Resources" section of a specified digital employee.
//
// - `source_type` is fixed to `FILE`, `scope` is fixed to `PERSONAL`, and `platform` is fixed to `LOCAL`.
//
// - The file must include an OSS persistent address (`filePath`). Other information such as the public access URL and original file name is optional.
//
// - If no target folder ID (`directoryId`) is specified, the file is automatically attached to the default root folder of the current digital employee. If specified, ensure that the folder belongs to the invoker\\"s personal folder.
//
// - Multiple authentication methods (AK, BearerToken, APP) are supported to authenticate requests.
//
// - The operation type is write, and operation logs are recorded for subsequent auditing.
//
// @param tmpReq - SendChatMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendChatMessageResponse
func (client *Client) SendChatMessageWithOptions(tmpReq *SendChatMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SendChatMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SendChatMessageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DigitalEmployeeName) {
		request.DigitalEmployeeNameShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DigitalEmployeeName, dara.String("digitalEmployeeName"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Files) {
		request.FilesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Files, dara.String("files"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TaskExecution) {
		request.TaskExecutionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskExecution, dara.String("taskExecution"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	if !dara.IsNil(request.ContentType) {
		body["contentType"] = request.ContentType
	}

	if !dara.IsNil(request.DigitalEmployeeNameShrink) {
		body["digitalEmployeeName"] = request.DigitalEmployeeNameShrink
	}

	if !dara.IsNil(request.DirectChat) {
		body["directChat"] = request.DirectChat
	}

	if !dara.IsNil(request.FilesShrink) {
		body["files"] = request.FilesShrink
	}

	if !dara.IsNil(request.Model) {
		body["model"] = request.Model
	}

	if !dara.IsNil(request.ReuseLastSession) {
		body["reuseLastSession"] = request.ReuseLastSession
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	if !dara.IsNil(request.TaskExecutionShrink) {
		body["taskExecution"] = request.TaskExecutionShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendChatMessage"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/sendChatMessage"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendChatMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a message.
//
// Description:
//
// ## Request description
//
// - This API is used to upload a file to the "My Resources" section of a specified digital employee.
//
// - `source_type` is fixed to `FILE`, `scope` is fixed to `PERSONAL`, and `platform` is fixed to `LOCAL`.
//
// - The file must include an OSS persistent address (`filePath`). Other information such as the public access URL and original file name is optional.
//
// - If no target folder ID (`directoryId`) is specified, the file is automatically attached to the default root folder of the current digital employee. If specified, ensure that the folder belongs to the invoker\\"s personal folder.
//
// - Multiple authentication methods (AK, BearerToken, APP) are supported to authenticate requests.
//
// - The operation type is write, and operation logs are recorded for subsequent auditing.
//
// @param request - SendChatMessageRequest
//
// @return SendChatMessageResponse
func (client *Client) SendChatMessage(request *SendChatMessageRequest) (_result *SendChatMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SendChatMessageResponse{}
	_body, _err := client.SendChatMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops conversation generation.
//
// @param request - StopChatMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopChatMessageResponse
func (client *Client) StopChatMessageWithOptions(request *StopChatMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopChatMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SessionId) {
		query["sessionId"] = request.SessionId
	}

	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopChatMessage"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/stopChatMessage"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopChatMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops conversation generation.
//
// @param request - StopChatMessageRequest
//
// @return StopChatMessageResponse
func (client *Client) StopChatMessage(request *StopChatMessageRequest) (_result *StopChatMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopChatMessageResponse{}
	_body, _err := client.StopChatMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Subscribes to a conversation message stream.
//
// Description:
//
// Subscribes to a conversation message stream.
//
// @param request - StreamChatMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StreamChatMessageResponse
func (client *Client) StreamChatMessageWithSSE(messageId *string, request *StreamChatMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *StreamChatMessageResponse, _yieldErr chan error) {
	defer close(_yield)
	client.streamChatMessageWithSSE_opYieldFunc(_yield, _yieldErr, messageId, request, headers, runtime)
	return
}

// Summary:
//
// Subscribes to a conversation message stream.
//
// Description:
//
// Subscribes to a conversation message stream.
//
// @param request - StreamChatMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StreamChatMessageResponse
func (client *Client) StreamChatMessageWithOptions(messageId *string, request *StreamChatMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StreamChatMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LastEventId) {
		query["lastEventId"] = request.LastEventId
	}

	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StreamChatMessage"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/streamChatMessage/" + dara.PercentEncode(dara.StringValue(messageId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StreamChatMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Subscribes to a conversation message stream.
//
// Description:
//
// Subscribes to a conversation message stream.
//
// @param request - StreamChatMessageRequest
//
// @return StreamChatMessageResponse
func (client *Client) StreamChatMessage(messageId *string, request *StreamChatMessageRequest) (_result *StreamChatMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StreamChatMessageResponse{}
	_body, _err := client.StreamChatMessageWithOptions(messageId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Pushes organizational structure synchronization.
//
// Description:
//
// Accepts a department tree and member relationships pushed from the client and creates an asynchronous synchronization task.
//
//	Processing flow:
//
//	1. Validates platformType (only saml, oauth2, or custom are allowed).
//
//	2. Validates data volume limits (departments + members <= 50000).
//
//	3. Validates the compatibility between syncMembers and platformType.
//
//	4. SAML/OAuth2 scenario: Parses or automatically derives ssoSettingsId.
//
//	5. Custom scenario: Validates that corpId has been registered through createCustomOrg.
//
//	6. Delegates to OrgSyncAuthorizedService to create the task (which includes permission verification).
//
//	7. Returns taskId for polling.
//
// @param tmpReq - SyncOrgStructureRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncOrgStructureResponse
func (client *Client) SyncOrgStructureWithOptions(tmpReq *SyncOrgStructureRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SyncOrgStructureResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SyncOrgStructureShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Departments) {
		request.DepartmentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Departments, dara.String("departments"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Members) {
		request.MembersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Members, dara.String("members"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CorpId) {
		body["corpId"] = request.CorpId
	}

	if !dara.IsNil(request.DepartmentsShrink) {
		body["departments"] = request.DepartmentsShrink
	}

	if !dara.IsNil(request.MembersShrink) {
		body["members"] = request.MembersShrink
	}

	if !dara.IsNil(request.PlatformType) {
		body["platformType"] = request.PlatformType
	}

	if !dara.IsNil(request.SsoSettingsId) {
		body["ssoSettingsId"] = request.SsoSettingsId
	}

	if !dara.IsNil(request.SyncMembers) {
		body["syncMembers"] = request.SyncMembers
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SyncOrgStructure"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/syncOrgStructure"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SyncOrgStructureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Pushes organizational structure synchronization.
//
// Description:
//
// Accepts a department tree and member relationships pushed from the client and creates an asynchronous synchronization task.
//
//	Processing flow:
//
//	1. Validates platformType (only saml, oauth2, or custom are allowed).
//
//	2. Validates data volume limits (departments + members <= 50000).
//
//	3. Validates the compatibility between syncMembers and platformType.
//
//	4. SAML/OAuth2 scenario: Parses or automatically derives ssoSettingsId.
//
//	5. Custom scenario: Validates that corpId has been registered through createCustomOrg.
//
//	6. Delegates to OrgSyncAuthorizedService to create the task (which includes permission verification).
//
//	7. Returns taskId for polling.
//
// @param request - SyncOrgStructureRequest
//
// @return SyncOrgStructureResponse
func (client *Client) SyncOrgStructure(request *SyncOrgStructureRequest) (_result *SyncOrgStructureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SyncOrgStructureResponse{}
	_body, _err := client.SyncOrgStructureWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Follows or unfollows a specified primary object.
//
// Description:
//
// ## Request description
//
// - **Precheck**:
//
//  1. When adding a follow: The system checks whether the primary object is already followed to prevent duplicates, and authenticates that the primary object exists.
//
//  2. When unfollowing: This is an idempotent operation. Regardless of whether the user has previously followed the object, `success=true` is returned.
//
// - **Security**: Three authentication methods are supported: AK, BearerToken, and APP.
//
// - **Request frequency limit**: A maximum of 100 requests can be send per second.
//
// - **Response log**: The response log record feature is enabled.
//
// - **Tenant relevance**: This API is associated with a specific tenant. The tenant ID of the invoker is used by default.
//
// - **Operation type**: Write operation.
//
// - **Backend service**: Requests are forwarded to an internal service for processing. The timeout period is 3 seconds.
//
// @param tmpReq - TogglePrimaryObjectFavoriteRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TogglePrimaryObjectFavoriteResponse
func (client *Client) TogglePrimaryObjectFavoriteWithOptions(tmpReq *TogglePrimaryObjectFavoriteRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TogglePrimaryObjectFavoriteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &TogglePrimaryObjectFavoriteShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ObjectIds) {
		request.ObjectIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ObjectIds, dara.String("objectIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Action) {
		body["action"] = request.Action
	}

	if !dara.IsNil(request.ObjectIdsShrink) {
		body["objectIds"] = request.ObjectIdsShrink
	}

	if !dara.IsNil(request.ObjectType) {
		body["objectType"] = request.ObjectType
	}

	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TogglePrimaryObjectFavorite"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/togglePrimaryObjectFavorite"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TogglePrimaryObjectFavoriteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Follows or unfollows a specified primary object.
//
// Description:
//
// ## Request description
//
// - **Precheck**:
//
//  1. When adding a follow: The system checks whether the primary object is already followed to prevent duplicates, and authenticates that the primary object exists.
//
//  2. When unfollowing: This is an idempotent operation. Regardless of whether the user has previously followed the object, `success=true` is returned.
//
// - **Security**: Three authentication methods are supported: AK, BearerToken, and APP.
//
// - **Request frequency limit**: A maximum of 100 requests can be send per second.
//
// - **Response log**: The response log record feature is enabled.
//
// - **Tenant relevance**: This API is associated with a specific tenant. The tenant ID of the invoker is used by default.
//
// - **Operation type**: Write operation.
//
// - **Backend service**: Requests are forwarded to an internal service for processing. The timeout period is 3 seconds.
//
// @param request - TogglePrimaryObjectFavoriteRequest
//
// @return TogglePrimaryObjectFavoriteResponse
func (client *Client) TogglePrimaryObjectFavorite(request *TogglePrimaryObjectFavoriteRequest) (_result *TogglePrimaryObjectFavoriteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TogglePrimaryObjectFavoriteResponse{}
	_body, _err := client.TogglePrimaryObjectFavoriteWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the authorization mode for digital employee usage permissions.
//
// Description:
//
// Switches the authorization mode for digital employee usage permissions.
//
//	Business logic:
//
//	1. Constructs an AuthContext from the identity.
//
//	2. Delegates to AgentAuthorizationAuthorizedService.update_auth_mode for execution.
//
//	3. Pre-validation: MANAGE permission + agent existence check (performed by the AuthorizedService layer, which authenticates before exposing existence).
//
//	4. SPECIFIED_USERS: Explicit authorization is required before usage.
//
//	5. ALL_USERS: All users can use the digital employee without authorization (management permissions are not affected).
//
// @param request - UpdateAgentAuthModeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAgentAuthModeResponse
func (client *Client) UpdateAgentAuthModeWithOptions(request *UpdateAgentAuthModeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAgentAuthModeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AuthMode) {
		body["authMode"] = request.AuthMode
	}

	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAgentAuthMode"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/updateAgentAuthMode"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAgentAuthModeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the authorization mode for digital employee usage permissions.
//
// Description:
//
// Switches the authorization mode for digital employee usage permissions.
//
//	Business logic:
//
//	1. Constructs an AuthContext from the identity.
//
//	2. Delegates to AgentAuthorizationAuthorizedService.update_auth_mode for execution.
//
//	3. Pre-validation: MANAGE permission + agent existence check (performed by the AuthorizedService layer, which authenticates before exposing existence).
//
//	4. SPECIFIED_USERS: Explicit authorization is required before usage.
//
//	5. ALL_USERS: All users can use the digital employee without authorization (management permissions are not affected).
//
// @param request - UpdateAgentAuthModeRequest
//
// @return UpdateAgentAuthModeResponse
func (client *Client) UpdateAgentAuthMode(request *UpdateAgentAuthModeRequest) (_result *UpdateAgentAuthModeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAgentAuthModeResponse{}
	_body, _err := client.UpdateAgentAuthModeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a session.
//
// Description:
//
// Updates a session.
//
// @param request - UpdateChatSessionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateChatSessionResponse
func (client *Client) UpdateChatSessionWithOptions(request *UpdateChatSessionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateChatSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Model) {
		body["model"] = request.Model
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Title) {
		body["title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateChatSession"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/updateChatSession"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateChatSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a session.
//
// Description:
//
// Updates a session.
//
// @param request - UpdateChatSessionRequest
//
// @return UpdateChatSessionResponse
func (client *Client) UpdateChatSession(request *UpdateChatSessionRequest) (_result *UpdateChatSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateChatSessionResponse{}
	_body, _err := client.UpdateChatSessionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the basic information of a user\\"s personal directory.
//
// Description:
//
// ## Operation description
//
// This API operation updates the personal directory information of a specified user, including the name, description, and parent directory. Ensure that the directory corresponding to the provided `directoryId` exists and belongs to the current user. If the `name` or `path` of the directory is changed, the system automatically and recursively updates the paths of all subdirectories to maintain consistency. When adjusting the parent directory, ensure the validity of the new parent directory (that is, it is not the directory itself and does not cause a circular dependency).
//
// - **Security constraints**: `tenant_id` and `user_id` must be derived from the authenticated identity.
//
// - **Permission requirements**: Corresponding RAM permissions are required to perform this operation.
//
// - **Input parameters**:
//
//   - `directoryId`: Required. The unique identifier of the directory to update.
//
//   - `name`: Optional. The new directory name.
//
//   - `description`: Optional. The new directory description.
//
//   - `parentId`: Optional. The ID of the new parent directory.
//
//   - `path`: Optional. When specified, the system cascades the update to the paths of the current directory and all its subdirectories.
//
// @param request - UpdateDirectoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDirectoryResponse
func (client *Client) UpdateDirectoryWithOptions(request *UpdateDirectoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDirectoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DirectoryId) {
		body["directoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.ParentId) {
		body["parentId"] = request.ParentId
	}

	if !dara.IsNil(request.Path) {
		body["path"] = request.Path
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDirectory"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/updateDirectory"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the basic information of a user\\"s personal directory.
//
// Description:
//
// ## Operation description
//
// This API operation updates the personal directory information of a specified user, including the name, description, and parent directory. Ensure that the directory corresponding to the provided `directoryId` exists and belongs to the current user. If the `name` or `path` of the directory is changed, the system automatically and recursively updates the paths of all subdirectories to maintain consistency. When adjusting the parent directory, ensure the validity of the new parent directory (that is, it is not the directory itself and does not cause a circular dependency).
//
// - **Security constraints**: `tenant_id` and `user_id` must be derived from the authenticated identity.
//
// - **Permission requirements**: Corresponding RAM permissions are required to perform this operation.
//
// - **Input parameters**:
//
//   - `directoryId`: Required. The unique identifier of the directory to update.
//
//   - `name`: Optional. The new directory name.
//
//   - `description`: Optional. The new directory description.
//
//   - `parentId`: Optional. The ID of the new parent directory.
//
//   - `path`: Optional. When specified, the system cascades the update to the paths of the current directory and all its subdirectories.
//
// @param request - UpdateDirectoryRequest
//
// @return UpdateDirectoryResponse
func (client *Client) UpdateDirectory(request *UpdateDirectoryRequest) (_result *UpdateDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDirectoryResponse{}
	_body, _err := client.UpdateDirectoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information of a specified enterprise knowledge base directory, including the name, description, and parent directory.
//
// Description:
//
// ## Operation description
//
// - This operation modifies a directory in the enterprise knowledge base.
//
// - You must have the `DEVELOPMENT_KB_MANAGE` permission to call this API operation.
//
// - The `tenantId` parameter is optional. If not provided, the tenant ID of the caller is used by default.
//
// - You must specify the `directoryId` of the directory to modify. The `name`, `description`, and `parentDirectoryId` parameters are optional. If not provided, the corresponding fields remain unchanged.
//
// - When a new `parentDirectoryId` is specified, the system checks whether the new parent directory belongs to the current tenant and does not cause a circular dependency.
//
// - This API operation supports multiple authentication methods (AK, BearerToken, APP) and has RAM permission control and operation auditing enabled.
//
// @param request - UpdateKnowledgeBaseDirectoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateKnowledgeBaseDirectoryResponse
func (client *Client) UpdateKnowledgeBaseDirectoryWithOptions(request *UpdateKnowledgeBaseDirectoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateKnowledgeBaseDirectoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DirectoryId) {
		body["directoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.ParentDirectoryId) {
		body["parentDirectoryId"] = request.ParentDirectoryId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateKnowledgeBaseDirectory"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/updateKnowledgeBaseDirectory"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateKnowledgeBaseDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information of a specified enterprise knowledge base directory, including the name, description, and parent directory.
//
// Description:
//
// ## Operation description
//
// - This operation modifies a directory in the enterprise knowledge base.
//
// - You must have the `DEVELOPMENT_KB_MANAGE` permission to call this API operation.
//
// - The `tenantId` parameter is optional. If not provided, the tenant ID of the caller is used by default.
//
// - You must specify the `directoryId` of the directory to modify. The `name`, `description`, and `parentDirectoryId` parameters are optional. If not provided, the corresponding fields remain unchanged.
//
// - When a new `parentDirectoryId` is specified, the system checks whether the new parent directory belongs to the current tenant and does not cause a circular dependency.
//
// - This API operation supports multiple authentication methods (AK, BearerToken, APP) and has RAM permission control and operation auditing enabled.
//
// @param request - UpdateKnowledgeBaseDirectoryRequest
//
// @return UpdateKnowledgeBaseDirectoryResponse
func (client *Client) UpdateKnowledgeBaseDirectory(request *UpdateKnowledgeBaseDirectoryRequest) (_result *UpdateKnowledgeBaseDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateKnowledgeBaseDirectoryResponse{}
	_body, _err := client.UpdateKnowledgeBaseDirectoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Edits the body content of a resource in an enterprise self-built knowledge base and triggers re-parsing.
//
// Description:
//
// ## Request description
//
// This operation allows you to update the body content of a specified enterprise knowledge base data source and optionally wait synchronously for parsing to complete. By setting the `forceSync` parameter, you can control whether the parsing process is executed synchronously or asynchronously. The default is asynchronous processing.
//
// - **Note**: When the `content` field is an empty string, the original content is cleared.
//
// - **Permission requirement**: Calling this operation requires the corresponding RAM action permission (`winnexo:UpdateKnowledgeBaseSourceContent`).
//
// @param request - UpdateKnowledgeBaseSourceContentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateKnowledgeBaseSourceContentResponse
func (client *Client) UpdateKnowledgeBaseSourceContentWithOptions(request *UpdateKnowledgeBaseSourceContentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateKnowledgeBaseSourceContentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	if !dara.IsNil(request.ForceSync) {
		body["forceSync"] = request.ForceSync
	}

	if !dara.IsNil(request.SourceId) {
		body["sourceId"] = request.SourceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateKnowledgeBaseSourceContent"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/updateKnowledgeBaseSourceContent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateKnowledgeBaseSourceContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Edits the body content of a resource in an enterprise self-built knowledge base and triggers re-parsing.
//
// Description:
//
// ## Request description
//
// This operation allows you to update the body content of a specified enterprise knowledge base data source and optionally wait synchronously for parsing to complete. By setting the `forceSync` parameter, you can control whether the parsing process is executed synchronously or asynchronously. The default is asynchronous processing.
//
// - **Note**: When the `content` field is an empty string, the original content is cleared.
//
// - **Permission requirement**: Calling this operation requires the corresponding RAM action permission (`winnexo:UpdateKnowledgeBaseSourceContent`).
//
// @param request - UpdateKnowledgeBaseSourceContentRequest
//
// @return UpdateKnowledgeBaseSourceContentResponse
func (client *Client) UpdateKnowledgeBaseSourceContent(request *UpdateKnowledgeBaseSourceContentRequest) (_result *UpdateKnowledgeBaseSourceContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateKnowledgeBaseSourceContentResponse{}
	_body, _err := client.UpdateKnowledgeBaseSourceContentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the resource tags of a specified data source in an enterprise knowledge base.
//
// Description:
//
// ## Request description
//
// - This operation updates the labels of a specific data source in an enterprise knowledge base.
//
// - You must have knowledge base management permissions to invoke this operation.
//
// - The `sourceTags` parameter accepts a JSON character string list, such as `["tagA", "tagB"]`. If you set this parameter to `null`, all existing labels are cleared.
//
// - The update operation affects only the `sourceTags` and `gmt_modified` fields and does not trigger the `process_source` workflow.
//
// - If the specified data source does not exist, the `ERR.Robject.Source.NotFound` fault is returned.
//
// - This operation supports authentication through AccessKey, BearerToken, or APP methods.
//
// - When you invoke this operation, make sure that `tenant_id` and `user_id` are from valid authentication identity information.
//
// @param request - UpdateKnowledgeBaseSourceTagsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateKnowledgeBaseSourceTagsResponse
func (client *Client) UpdateKnowledgeBaseSourceTagsWithOptions(request *UpdateKnowledgeBaseSourceTagsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateKnowledgeBaseSourceTagsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SourceId) {
		body["sourceId"] = request.SourceId
	}

	if !dara.IsNil(request.SourceTags) {
		body["sourceTags"] = request.SourceTags
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateKnowledgeBaseSourceTags"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/updateKnowledgeBaseSourceTags"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateKnowledgeBaseSourceTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the resource tags of a specified data source in an enterprise knowledge base.
//
// Description:
//
// ## Request description
//
// - This operation updates the labels of a specific data source in an enterprise knowledge base.
//
// - You must have knowledge base management permissions to invoke this operation.
//
// - The `sourceTags` parameter accepts a JSON character string list, such as `["tagA", "tagB"]`. If you set this parameter to `null`, all existing labels are cleared.
//
// - The update operation affects only the `sourceTags` and `gmt_modified` fields and does not trigger the `process_source` workflow.
//
// - If the specified data source does not exist, the `ERR.Robject.Source.NotFound` fault is returned.
//
// - This operation supports authentication through AccessKey, BearerToken, or APP methods.
//
// - When you invoke this operation, make sure that `tenant_id` and `user_id` are from valid authentication identity information.
//
// @param request - UpdateKnowledgeBaseSourceTagsRequest
//
// @return UpdateKnowledgeBaseSourceTagsResponse
func (client *Client) UpdateKnowledgeBaseSourceTags(request *UpdateKnowledgeBaseSourceTagsRequest) (_result *UpdateKnowledgeBaseSourceTagsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateKnowledgeBaseSourceTagsResponse{}
	_body, _err := client.UpdateKnowledgeBaseSourceTagsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a scheduled task.
//
// Description:
//
// Updates a scheduled task.
//
// @param tmpReq - UpdateScheduledTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateScheduledTaskResponse
func (client *Client) UpdateScheduledTaskWithOptions(tmpReq *UpdateScheduledTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateScheduledTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateScheduledTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Description) {
		request.DescriptionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Description, dara.String("description"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DigitalEmployeeName) {
		request.DigitalEmployeeNameShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DigitalEmployeeName, dara.String("digitalEmployeeName"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Segments) {
		request.SegmentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Segments, dara.String("segments"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TaskDetail) {
		request.TaskDetailShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskDetail, dara.String("taskDetail"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TriggerConfig) {
		request.TriggerConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TriggerConfig, dara.String("triggerConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.VisibleMemberUserIds) {
		request.VisibleMemberUserIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VisibleMemberUserIds, dara.String("visibleMemberUserIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DescriptionShrink) {
		body["description"] = request.DescriptionShrink
	}

	if !dara.IsNil(request.DigitalEmployeeNameShrink) {
		body["digitalEmployeeName"] = request.DigitalEmployeeNameShrink
	}

	if !dara.IsNil(request.IsOpen) {
		body["isOpen"] = request.IsOpen
	}

	if !dara.IsNil(request.Model) {
		body["model"] = request.Model
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.SegmentsShrink) {
		body["segments"] = request.SegmentsShrink
	}

	if !dara.IsNil(request.TaskDetailShrink) {
		body["taskDetail"] = request.TaskDetailShrink
	}

	if !dara.IsNil(request.TaskId) {
		body["taskId"] = request.TaskId
	}

	if !dara.IsNil(request.TriggerConfigShrink) {
		body["triggerConfig"] = request.TriggerConfigShrink
	}

	if !dara.IsNil(request.Visibility) {
		body["visibility"] = request.Visibility
	}

	if !dara.IsNil(request.VisibleMemberUserIdsShrink) {
		body["visibleMemberUserIds"] = request.VisibleMemberUserIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateScheduledTask"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/updateScheduledTask"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateScheduledTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a scheduled task.
//
// Description:
//
// Updates a scheduled task.
//
// @param request - UpdateScheduledTaskRequest
//
// @return UpdateScheduledTaskResponse
func (client *Client) UpdateScheduledTask(request *UpdateScheduledTaskRequest) (_result *UpdateScheduledTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateScheduledTaskResponse{}
	_body, _err := client.UpdateScheduledTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the content of an editable data source within a tenant and triggers re-parsing.
//
// Description:
//
// ## Operation description
//
// - This API operation updates the content of a data source within a specified tenant and triggers synchronous or asynchronous re-parsing of the data source as needed.
//
// - `tenant_id` and `user_id` are used only for authentication and are not involved in actual business logic processing.
//
// - When the provided content is an empty string, the system performs the operation according to the existing service contract.
//
// - If the specified data source does not exist, a standard NotFound error is returned. Other exceptions are handled by the global exception chain.
//
// - Set the `forceSync` parameter to determine whether to wait for the parsing process to complete. The default behavior is asynchronous queuing.
//
// @param request - UpdateSourceContentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSourceContentResponse
func (client *Client) UpdateSourceContentWithOptions(request *UpdateSourceContentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateSourceContentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	if !dara.IsNil(request.ForceSync) {
		body["forceSync"] = request.ForceSync
	}

	if !dara.IsNil(request.SourceId) {
		body["sourceId"] = request.SourceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSourceContent"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/updateSourceContent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSourceContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the content of an editable data source within a tenant and triggers re-parsing.
//
// Description:
//
// ## Operation description
//
// - This API operation updates the content of a data source within a specified tenant and triggers synchronous or asynchronous re-parsing of the data source as needed.
//
// - `tenant_id` and `user_id` are used only for authentication and are not involved in actual business logic processing.
//
// - When the provided content is an empty string, the system performs the operation according to the existing service contract.
//
// - If the specified data source does not exist, a standard NotFound error is returned. Other exceptions are handled by the global exception chain.
//
// - Set the `forceSync` parameter to determine whether to wait for the parsing process to complete. The default behavior is asynchronous queuing.
//
// @param request - UpdateSourceContentRequest
//
// @return UpdateSourceContentResponse
func (client *Client) UpdateSourceContent(request *UpdateSourceContentRequest) (_result *UpdateSourceContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateSourceContentResponse{}
	_body, _err := client.UpdateSourceContentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies knowledge base folder information.
//
// Description:
//
// Modifies knowledge base folder information.
//
// @param request - UpdateTenantDirectoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTenantDirectoryResponse
func (client *Client) UpdateTenantDirectoryWithOptions(request *UpdateTenantDirectoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateTenantDirectoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DirectoryId) {
		body["directoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.ParentId) {
		body["parentId"] = request.ParentId
	}

	if !dara.IsNil(request.Path) {
		body["path"] = request.Path
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTenantDirectory"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/updateTenantDirectory"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTenantDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies knowledge base folder information.
//
// Description:
//
// Modifies knowledge base folder information.
//
// @param request - UpdateTenantDirectoryRequest
//
// @return UpdateTenantDirectoryResponse
func (client *Client) UpdateTenantDirectory(request *UpdateTenantDirectoryRequest) (_result *UpdateTenantDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTenantDirectoryResponse{}
	_body, _err := client.UpdateTenantDirectoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies user information, including enabling or disabling the user.
//
// Description:
//
// Modifies user information through OpenAPI.
//
//	Business orchestration:
//
//	1. Parse roleCodes → role_ids
//
//	2. If isActive has changed, perform the status switch first (including last super admin protection)
//
//	3. Call update_tenant_member to modify other fields (displayName / roleCodes / userGroupIds)
//
//	4. Return HTTP 200 if all steps succeed
//
//	Execution order notes:
//
//	- The isActive status change is performed before other field writes. The two steps are not in the same transaction.
//
//	- If validation fails (such as last super admin protection) → an exception is thrown and subsequent steps are not executed.
//
//	- If the isActive change has been persisted but a subsequent step fails, the isActive change is not rolled back.
//
// @param tmpReq - UpdateUserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserResponse
func (client *Client) UpdateUserWithOptions(tmpReq *UpdateUserRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateUserShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RoleCodes) {
		request.RoleCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoleCodes, dara.String("roleCodes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserGroupIds) {
		request.UserGroupIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserGroupIds, dara.String("userGroupIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DisplayName) {
		body["displayName"] = request.DisplayName
	}

	if !dara.IsNil(request.IsActive) {
		body["isActive"] = request.IsActive
	}

	if !dara.IsNil(request.RoleCodesShrink) {
		body["roleCodes"] = request.RoleCodesShrink
	}

	if !dara.IsNil(request.UserGroupIdsShrink) {
		body["userGroupIds"] = request.UserGroupIdsShrink
	}

	if !dara.IsNil(request.WnUserId) {
		body["wnUserId"] = request.WnUserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUser"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/updateUser"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies user information, including enabling or disabling the user.
//
// Description:
//
// Modifies user information through OpenAPI.
//
//	Business orchestration:
//
//	1. Parse roleCodes → role_ids
//
//	2. If isActive has changed, perform the status switch first (including last super admin protection)
//
//	3. Call update_tenant_member to modify other fields (displayName / roleCodes / userGroupIds)
//
//	4. Return HTTP 200 if all steps succeed
//
//	Execution order notes:
//
//	- The isActive status change is performed before other field writes. The two steps are not in the same transaction.
//
//	- If validation fails (such as last super admin protection) → an exception is thrown and subsequent steps are not executed.
//
//	- If the isActive change has been persisted but a subsequent step fails, the isActive change is not rolled back.
//
// @param request - UpdateUserRequest
//
// @return UpdateUserResponse
func (client *Client) UpdateUser(request *UpdateUserRequest) (_result *UpdateUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateUserResponse{}
	_body, _err := client.UpdateUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the name, description, and parent relationship of a specified user group.
//
// Description:
//
// WinNexo user management OpenAPI: updates a user group. The tenant identity is obtained from the authentication context.
//
// @param request - UpdateUserGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserGroupResponse
func (client *Client) UpdateUserGroupWithOptions(request *UpdateUserGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.MoveToRoot) {
		body["moveToRoot"] = request.MoveToRoot
	}

	if !dara.IsNil(request.ParentId) {
		body["parentId"] = request.ParentId
	}

	if !dara.IsNil(request.UserGroupId) {
		body["userGroupId"] = request.UserGroupId
	}

	if !dara.IsNil(request.UserGroupName) {
		body["userGroupName"] = request.UserGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUserGroup"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/updateUserGroup"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the name, description, and parent relationship of a specified user group.
//
// Description:
//
// WinNexo user management OpenAPI: updates a user group. The tenant identity is obtained from the authentication context.
//
// @param request - UpdateUserGroupRequest
//
// @return UpdateUserGroupResponse
func (client *Client) UpdateUserGroup(request *UpdateUserGroupRequest) (_result *UpdateUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateUserGroupResponse{}
	_body, _err := client.UpdateUserGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates partial fields of the current user information and returns the complete user information.
//
// Description:
//
// ## Request description
//
// - This API allows the caller to update some or all optional fields of a specified user. Fields that are not provided retain their original values.
//
// - Use the `tenantId` parameter to specify a tenant ID. If omitted, the default tenant of the caller is used.
//
// - After a successful update, the response body contains the complete user information object.
//
// - This operation requires authentication and supports AK, BearerToken, and APP security schemes.
//
// - The request content type is JSON, and the operation is available only over HTTPS.
//
// - Note: The `profileRoleInfo` field is valid only when the user role is set to Others. It describes the specific role information of the user.
//
// @param request - UpdateUserInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserInfoResponse
func (client *Client) UpdateUserInfoWithOptions(request *UpdateUserInfoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateUserInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Avatar) {
		body["avatar"] = request.Avatar
	}

	if !dara.IsNil(request.LanguagePreference) {
		body["languagePreference"] = request.LanguagePreference
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Offering) {
		body["offering"] = request.Offering
	}

	if !dara.IsNil(request.ProfileRoleInfo) {
		body["profileRoleInfo"] = request.ProfileRoleInfo
	}

	if !dara.IsNil(request.SelfIntroduction) {
		body["selfIntroduction"] = request.SelfIntroduction
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUserInfo"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/updateUserInfo"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates partial fields of the current user information and returns the complete user information.
//
// Description:
//
// ## Request description
//
// - This API allows the caller to update some or all optional fields of a specified user. Fields that are not provided retain their original values.
//
// - Use the `tenantId` parameter to specify a tenant ID. If omitted, the default tenant of the caller is used.
//
// - After a successful update, the response body contains the complete user information object.
//
// - This operation requires authentication and supports AK, BearerToken, and APP security schemes.
//
// - The request content type is JSON, and the operation is available only over HTTPS.
//
// - Note: The `profileRoleInfo` field is valid only when the user role is set to Others. It describes the specific role information of the user.
//
// @param request - UpdateUserInfoRequest
//
// @return UpdateUserInfoResponse
func (client *Client) UpdateUserInfo(request *UpdateUserInfoRequest) (_result *UpdateUserInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateUserInfoResponse{}
	_body, _err := client.UpdateUserInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uploads a local file for a session.
//
// Description:
//
// ## Operation description
//
// This API operation uploads a temporary temporary file by using the **file transfer upload*	- mode (`fileTransfer`). The file binary data is not transmitted in the request body of this API operation. Instead, the file is first uploaded to Object Storage Service (OSS), and then the OSS address is passed to the backend through the `FileUrl` parameter. The backend retrieves the bytes from that address, writes them to its own OSS bucket, and creates a temporary temporary file record.
//
// ### How to call
//
// - **Recommended**: Use the `UploadChatFileAdvance` method generated by the SDK. Pass in the local file stream, and the SDK automatically completes the transfer upload and populates the `FileUrl` parameter.
//
// - **Direct upload**: Upload the file to an OSS address accessible by the server, and then call this API operation directly with the `FileUrl` parameter.
//
// ### Request parameters
//
// - **FileUrl**: Required. The OSS address of the file. When you use the Advance method, the SDK automatically populates this parameter. You do not need to manually set it.
//
// - **FileName**: Required. The original file name including the extension, such as `report.pdf`. The OSS address generated during the transfer does not carry the original file name. The backend uses this parameter to determine the file extension and display name. Therefore, you must explicitly pass in this parameter.
//
// - **ContentType**: Optional. The MIME type of the file. If this parameter is not specified, `application/octet-stream` is used.
//
// - **OperatingObjectName**: Optional. The agent namespace identifier that determines the file storage path.
//
// ### Response parameters
//
// The response includes the OSS object path `objectName`, the storage address `fileUrl`, the publicly accessible address `filePublicUrl` (valid for 1 hour), and the file record ID `fileRecordId`. The `uploadSignatureUrl` parameter is always empty in this mode.
//
// @param request - UploadChatFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadChatFileResponse
func (client *Client) UploadChatFileWithOptions(request *UploadChatFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UploadChatFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ContentType) {
		body["contentType"] = request.ContentType
	}

	if !dara.IsNil(request.FileName) {
		body["fileName"] = request.FileName
	}

	if !dara.IsNil(request.FileUrl) {
		body["fileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.OperatingObjectName) {
		body["operatingObjectName"] = request.OperatingObjectName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadChatFile"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/uploadChatFile"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadChatFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads a local file for a session.
//
// Description:
//
// ## Operation description
//
// This API operation uploads a temporary temporary file by using the **file transfer upload*	- mode (`fileTransfer`). The file binary data is not transmitted in the request body of this API operation. Instead, the file is first uploaded to Object Storage Service (OSS), and then the OSS address is passed to the backend through the `FileUrl` parameter. The backend retrieves the bytes from that address, writes them to its own OSS bucket, and creates a temporary temporary file record.
//
// ### How to call
//
// - **Recommended**: Use the `UploadChatFileAdvance` method generated by the SDK. Pass in the local file stream, and the SDK automatically completes the transfer upload and populates the `FileUrl` parameter.
//
// - **Direct upload**: Upload the file to an OSS address accessible by the server, and then call this API operation directly with the `FileUrl` parameter.
//
// ### Request parameters
//
// - **FileUrl**: Required. The OSS address of the file. When you use the Advance method, the SDK automatically populates this parameter. You do not need to manually set it.
//
// - **FileName**: Required. The original file name including the extension, such as `report.pdf`. The OSS address generated during the transfer does not carry the original file name. The backend uses this parameter to determine the file extension and display name. Therefore, you must explicitly pass in this parameter.
//
// - **ContentType**: Optional. The MIME type of the file. If this parameter is not specified, `application/octet-stream` is used.
//
// - **OperatingObjectName**: Optional. The agent namespace identifier that determines the file storage path.
//
// ### Response parameters
//
// The response includes the OSS object path `objectName`, the storage address `fileUrl`, the publicly accessible address `filePublicUrl` (valid for 1 hour), and the file record ID `fileRecordId`. The `uploadSignatureUrl` parameter is always empty in this mode.
//
// @param request - UploadChatFileRequest
//
// @return UploadChatFileResponse
func (client *Client) UploadChatFile(request *UploadChatFileRequest) (_result *UploadChatFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UploadChatFileResponse{}
	_body, _err := client.UploadChatFileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadChatFileAdvance(request *UploadChatFileAdvanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UploadChatFileResponse, _err error) {
	// Step 0: init client
	if dara.IsNil(client.Credential) {
		_err = &openapi.ClientError{
			Code:    dara.String("InvalidCredentials"),
			Message: dara.String("Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details."),
		}
		return _result, _err
	}

	credentialModel, _err := client.Credential.GetCredential()
	if _err != nil {
		return _result, _err
	}

	accessKeyId := dara.StringValue(credentialModel.AccessKeyId)
	accessKeySecret := dara.StringValue(credentialModel.AccessKeySecret)
	securityToken := dara.StringValue(credentialModel.SecurityToken)
	credentialType := dara.StringValue(credentialModel.Type)
	openPlatformEndpoint := dara.StringValue(client.OpenPlatformEndpoint)
	if dara.IsNil(dara.String(openPlatformEndpoint)) || openPlatformEndpoint == "" {
		openPlatformEndpoint = "openplatform.aliyuncs.com"
	}

	if dara.IsNil(dara.String(credentialType)) {
		credentialType = "access_key"
	}

	authConfig := &openapiutil.Config{
		AccessKeyId:     dara.String(accessKeyId),
		AccessKeySecret: dara.String(accessKeySecret),
		SecurityToken:   dara.String(securityToken),
		Type:            dara.String(credentialType),
		Endpoint:        dara.String(openPlatformEndpoint),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openapi.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := map[string]*string{
		"Product":  dara.String("WinNexo"),
		"RegionId": client.RegionId,
	}
	authReq := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(authRequest),
	}
	authParams := &openapiutil.Params{
		Action:      dara.String("AuthorizeFileUpload"),
		Version:     dara.String("2019-12-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	authResponse := map[string]interface{}{}
	fileObj := &dara.FileField{}
	ossHeader := map[string]interface{}{}
	tmpBody := map[string]interface{}{}
	useAccelerate := false
	authResponseBody := make(map[string]*string)
	uploadChatFileReq := &UploadChatFileRequest{}
	openapiutil.Convert(request, uploadChatFileReq)
	if !dara.IsNil(request.FileUrlObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.FileUrlObject,
			ContentType: dara.String(""),
		}
		ossHeader = map[string]interface{}{
			"host":                  dara.StringValue(openapiutil.GetEndpoint(authResponseBody["Endpoint"], dara.Bool(useAccelerate), client.EndpointType)),
			"OSSAccessKeyId":        dara.StringValue(authResponseBody["AccessKeyId"]),
			"policy":                dara.StringValue(authResponseBody["EncodedPolicy"]),
			"Signature":             dara.StringValue(authResponseBody["Signature"]),
			"key":                   dara.StringValue(authResponseBody["ObjectKey"]),
			"file":                  fileObj,
			"success_action_status": "201",
		}
		_, _err = client._postOSSObject(authResponseBody["Bucket"], ossHeader, runtime)
		if _err != nil {
			return _result, _err
		}
		uploadChatFileReq.FileUrl = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	uploadChatFileResp, _err := client.UploadChatFileWithOptions(uploadChatFileReq, headers, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = uploadChatFileResp
	return _result, _err
}

func _postOSSObject_opResponse(response_ *dara.Response) (_result map[string]interface{}, _err error) {
	var respMap map[string]interface{}
	bodyStr, _err := dara.ReadAsString(response_.Body)
	if _err != nil {
		return _result, _err
	}

	if (dara.IntValue(response_.StatusCode) >= 400) && (dara.IntValue(response_.StatusCode) < 600) {
		respMap = dara.ParseXml(bodyStr, nil)
		err := dara.ToMap(respMap["Error"])
		_err = &openapi.ClientError{
			Code:    dara.String(dara.ToString(err["Code"])),
			Message: dara.String(dara.ToString(err["Message"])),
			Data: map[string]interface{}{
				"httpCode":  dara.IntValue(response_.StatusCode),
				"requestId": dara.ToString(err["RequestId"]),
				"hostId":    dara.ToString(err["HostId"]),
			},
		}
		return _result, _err
	}

	respMap = dara.ParseXml(bodyStr, nil)
	_result = make(map[string]interface{})
	_err = dara.Convert(dara.ToMap(respMap), &_result)

	return _result, _err
}

func (client *Client) sendChatMessageWithSSE_opYieldFunc(_yield chan *SendChatMessageResponse, _yieldErr chan error, tmpReq *SendChatMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &SendChatMessageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DigitalEmployeeName) {
		request.DigitalEmployeeNameShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DigitalEmployeeName, dara.String("digitalEmployeeName"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Files) {
		request.FilesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Files, dara.String("files"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TaskExecution) {
		request.TaskExecutionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskExecution, dara.String("taskExecution"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	if !dara.IsNil(request.ContentType) {
		body["contentType"] = request.ContentType
	}

	if !dara.IsNil(request.DigitalEmployeeNameShrink) {
		body["digitalEmployeeName"] = request.DigitalEmployeeNameShrink
	}

	if !dara.IsNil(request.DirectChat) {
		body["directChat"] = request.DirectChat
	}

	if !dara.IsNil(request.FilesShrink) {
		body["files"] = request.FilesShrink
	}

	if !dara.IsNil(request.Model) {
		body["model"] = request.Model
	}

	if !dara.IsNil(request.ReuseLastSession) {
		body["reuseLastSession"] = request.ReuseLastSession
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	if !dara.IsNil(request.TaskExecutionShrink) {
		body["taskExecution"] = request.TaskExecutionShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendChatMessage"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/sendChatMessage"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		if !dara.IsNil(resp.Event) && !dara.IsNil(resp.Event.Data) {
			data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
			_err := dara.ConvertChan(map[string]interface{}{
				"statusCode": dara.IntValue(resp.StatusCode),
				"headers":    resp.Headers,
				"id":         dara.StringValue(resp.Event.Id),
				"event":      dara.StringValue(resp.Event.Event),
				"body":       data,
			}, _yield)
			if _err != nil {
				_yieldErr <- _err
				return
			}
		}

	}
}

func (client *Client) streamChatMessageWithSSE_opYieldFunc(_yield chan *StreamChatMessageResponse, _yieldErr chan error, messageId *string, request *StreamChatMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LastEventId) {
		query["lastEventId"] = request.LastEventId
	}

	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StreamChatMessage"),
		Version:     dara.String("2026-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/streamChatMessage/" + dara.PercentEncode(dara.StringValue(messageId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		if !dara.IsNil(resp.Event) && !dara.IsNil(resp.Event.Data) {
			data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
			_err := dara.ConvertChan(map[string]interface{}{
				"statusCode": dara.IntValue(resp.StatusCode),
				"headers":    resp.Headers,
				"id":         dara.StringValue(resp.Event.Id),
				"event":      dara.StringValue(resp.Event.Event),
				"body":       data,
			}, _yield)
			if _err != nil {
				_yieldErr <- _err
				return
			}
		}

	}
}
