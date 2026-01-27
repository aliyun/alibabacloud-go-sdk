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
	client.Endpoint, _err = client.GetEndpoint(dara.String("ecd"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Enables OTA updates for cloud computers.
//
// @param request - ApproveFotaUpdateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApproveFotaUpdateResponse
func (client *Client) ApproveFotaUpdateWithOptions(request *ApproveFotaUpdateRequest, runtime *dara.RuntimeOptions) (_result *ApproveFotaUpdateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppVersion) {
		query["AppVersion"] = request.AppVersion
	}

	if !dara.IsNil(request.ClientId) {
		query["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.LoginToken) {
		query["LoginToken"] = request.LoginToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.TargetStatus) {
		query["TargetStatus"] = request.TargetStatus
	}

	if !dara.IsNil(request.Uuid) {
		query["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApproveFotaUpdate"),
		Version:     dara.String("2020-10-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApproveFotaUpdateResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables OTA updates for cloud computers.
//
// @param request - ApproveFotaUpdateRequest
//
// @return ApproveFotaUpdateResponse
func (client *Client) ApproveFotaUpdate(request *ApproveFotaUpdateRequest) (_result *ApproveFotaUpdateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ApproveFotaUpdateResponse{}
	_body, _err := client.ApproveFotaUpdateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the password of a user account.
//
// @param request - ChangePasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangePasswordResponse
func (client *Client) ChangePasswordWithOptions(request *ChangePasswordRequest, runtime *dara.RuntimeOptions) (_result *ChangePasswordResponse, _err error) {
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

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.LoginToken) {
		query["LoginToken"] = request.LoginToken
	}

	if !dara.IsNil(request.NewPassword) {
		query["NewPassword"] = request.NewPassword
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.OldPassword) {
		query["OldPassword"] = request.OldPassword
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangePassword"),
		Version:     dara.String("2020-10-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangePasswordResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the password of a user account.
//
// @param request - ChangePasswordRequest
//
// @return ChangePasswordResponse
func (client *Client) ChangePassword(request *ChangePasswordRequest) (_result *ChangePasswordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChangePasswordResponse{}
	_body, _err := client.ChangePasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteFingerPrintTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFingerPrintTemplateResponse
func (client *Client) DeleteFingerPrintTemplateWithOptions(request *DeleteFingerPrintTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteFingerPrintTemplateResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Index) {
		query["Index"] = request.Index
	}

	if !dara.IsNil(request.LoginToken) {
		query["LoginToken"] = request.LoginToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFingerPrintTemplate"),
		Version:     dara.String("2020-10-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFingerPrintTemplateResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteFingerPrintTemplateRequest
//
// @return DeleteFingerPrintTemplateResponse
func (client *Client) DeleteFingerPrintTemplate(request *DeleteFingerPrintTemplateRequest) (_result *DeleteFingerPrintTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFingerPrintTemplateResponse{}
	_body, _err := client.DeleteFingerPrintTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries directory details.
//
// @param request - DescribeDirectoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDirectoriesResponse
func (client *Client) DescribeDirectoriesWithOptions(request *DescribeDirectoriesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDirectoriesResponse, _err error) {
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

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDirectories"),
		Version:     dara.String("2020-10-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDirectoriesResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries directory details.
//
// @param request - DescribeDirectoriesRequest
//
// @return DescribeDirectoriesResponse
func (client *Client) DescribeDirectories(request *DescribeDirectoriesRequest) (_result *DescribeDirectoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDirectoriesResponse{}
	_body, _err := client.DescribeDirectoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries fingerprint templates.
//
// @param request - DescribeFingerPrintTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFingerPrintTemplatesResponse
func (client *Client) DescribeFingerPrintTemplatesWithOptions(request *DescribeFingerPrintTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DescribeFingerPrintTemplatesResponse, _err error) {
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

	if !dara.IsNil(request.LoginToken) {
		query["LoginToken"] = request.LoginToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFingerPrintTemplates"),
		Version:     dara.String("2020-10-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFingerPrintTemplatesResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries fingerprint templates.
//
// @param request - DescribeFingerPrintTemplatesRequest
//
// @return DescribeFingerPrintTemplatesResponse
func (client *Client) DescribeFingerPrintTemplates(request *DescribeFingerPrintTemplatesRequest) (_result *DescribeFingerPrintTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFingerPrintTemplatesResponse{}
	_body, _err := client.DescribeFingerPrintTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of cloud computers. Currently, only the region corresponding to the Chinese mainland can be queried (excluding: Nanjing-local region-shutting down).
//
// Description:
//
//	  This API is a centralized domain name. The endpoint is in the China (Shanghai) region. You cannot call this API operation in other regions.
//
//		- The cloud computer status information in this interface has a delay of 1 to 3 seconds from the actual value.
//
// @param request - DescribeGlobalDesktopsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGlobalDesktopsResponse
func (client *Client) DescribeGlobalDesktopsWithOptions(request *DescribeGlobalDesktopsRequest, runtime *dara.RuntimeOptions) (_result *DescribeGlobalDesktopsResponse, _err error) {
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

	if !dara.IsNil(request.DesktopAccessType) {
		query["DesktopAccessType"] = request.DesktopAccessType
	}

	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.DesktopName) {
		query["DesktopName"] = request.DesktopName
	}

	if !dara.IsNil(request.DesktopStatus) {
		query["DesktopStatus"] = request.DesktopStatus
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.LoginRegionId) {
		query["LoginRegionId"] = request.LoginRegionId
	}

	if !dara.IsNil(request.LoginToken) {
		query["LoginToken"] = request.LoginToken
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.QueryFotaUpdate) {
		query["QueryFotaUpdate"] = request.QueryFotaUpdate
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SearchRegionId) {
		query["SearchRegionId"] = request.SearchRegionId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.SortType) {
		query["SortType"] = request.SortType
	}

	if !dara.IsNil(request.WithoutLatency) {
		query["WithoutLatency"] = request.WithoutLatency
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGlobalDesktops"),
		Version:     dara.String("2020-10-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGlobalDesktopsResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of cloud computers. Currently, only the region corresponding to the Chinese mainland can be queried (excluding: Nanjing-local region-shutting down).
//
// Description:
//
//	  This API is a centralized domain name. The endpoint is in the China (Shanghai) region. You cannot call this API operation in other regions.
//
//		- The cloud computer status information in this interface has a delay of 1 to 3 seconds from the actual value.
//
// @param request - DescribeGlobalDesktopsRequest
//
// @return DescribeGlobalDesktopsResponse
func (client *Client) DescribeGlobalDesktops(request *DescribeGlobalDesktopsRequest) (_result *DescribeGlobalDesktopsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGlobalDesktopsResponse{}
	_body, _err := client.DescribeGlobalDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries office networks.
//
// @param request - DescribeOfficeSitesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOfficeSitesResponse
func (client *Client) DescribeOfficeSitesWithOptions(request *DescribeOfficeSitesRequest, runtime *dara.RuntimeOptions) (_result *DescribeOfficeSitesResponse, _err error) {
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

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Uuid) {
		query["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOfficeSites"),
		Version:     dara.String("2020-10-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOfficeSitesResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries office networks.
//
// @param request - DescribeOfficeSitesRequest
//
// @return DescribeOfficeSitesResponse
func (client *Client) DescribeOfficeSites(request *DescribeOfficeSitesRequest) (_result *DescribeOfficeSitesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeOfficeSitesResponse{}
	_body, _err := client.DescribeOfficeSitesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2020-10-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

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
// Queries the snapshots that are created based on a cloud computer and the details of the snapshots.
//
// @param request - DescribeSnapshotsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSnapshotsResponse
func (client *Client) DescribeSnapshotsWithOptions(request *DescribeSnapshotsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSnapshotsResponse, _err error) {
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

	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.LoginToken) {
		query["LoginToken"] = request.LoginToken
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.SnapshotId) {
		query["SnapshotId"] = request.SnapshotId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSnapshots"),
		Version:     dara.String("2020-10-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSnapshotsResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the snapshots that are created based on a cloud computer and the details of the snapshots.
//
// @param request - DescribeSnapshotsRequest
//
// @return DescribeSnapshotsResponse
func (client *Client) DescribeSnapshots(request *DescribeSnapshotsRequest) (_result *DescribeSnapshotsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSnapshotsResponse{}
	_body, _err := client.DescribeSnapshotsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries user resources.
//
// Description:
//
// Before you call this operation, verify supported resource and service types in Alibaba Cloud Workspace.
//
// @param request - DescribeUserResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserResourcesResponse
func (client *Client) DescribeUserResourcesWithOptions(request *DescribeUserResourcesRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessType) {
		query["AccessType"] = request.AccessType
	}

	if !dara.IsNil(request.AutoRefresh) {
		query["AutoRefresh"] = request.AutoRefresh
	}

	if !dara.IsNil(request.CategoryId) {
		query["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.CategoryType) {
		query["CategoryType"] = request.CategoryType
	}

	if !dara.IsNil(request.ClientId) {
		query["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientType) {
		query["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.ClientVersion) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !dara.IsNil(request.DualCenterForward) {
		query["DualCenterForward"] = request.DualCenterForward
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.LoginRegionId) {
		query["LoginRegionId"] = request.LoginRegionId
	}

	if !dara.IsNil(request.LoginToken) {
		query["LoginToken"] = request.LoginToken
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OfficeSiteIds) {
		query["OfficeSiteIds"] = request.OfficeSiteIds
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.ProductTypes) {
		query["ProductTypes"] = request.ProductTypes
	}

	if !dara.IsNil(request.ProtocolType) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !dara.IsNil(request.QueryDesktopDurationList) {
		query["QueryDesktopDurationList"] = request.QueryDesktopDurationList
	}

	if !dara.IsNil(request.QueryDesktopTimers) {
		query["QueryDesktopTimers"] = request.QueryDesktopTimers
	}

	if !dara.IsNil(request.QueryFotaUpdate) {
		query["QueryFotaUpdate"] = request.QueryFotaUpdate
	}

	if !dara.IsNil(request.RefreshFotaUpdate) {
		query["RefreshFotaUpdate"] = request.RefreshFotaUpdate
	}

	if !dara.IsNil(request.ResourceIds) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.ResourceName) {
		query["ResourceName"] = request.ResourceName
	}

	if !dara.IsNil(request.ResourceTypes) {
		query["ResourceTypes"] = request.ResourceTypes
	}

	if !dara.IsNil(request.Scene) {
		query["Scene"] = request.Scene
	}

	if !dara.IsNil(request.SearchRegionId) {
		query["SearchRegionId"] = request.SearchRegionId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.SortType) {
		query["SortType"] = request.SortType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserResources"),
		Version:     dara.String("2020-10-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserResourcesResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries user resources.
//
// Description:
//
// Before you call this operation, verify supported resource and service types in Alibaba Cloud Workspace.
//
// @param request - DescribeUserResourcesRequest
//
// @return DescribeUserResourcesResponse
func (client *Client) DescribeUserResources(request *DescribeUserResourcesRequest) (_result *DescribeUserResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserResourcesResponse{}
	_body, _err := client.DescribeUserResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Encrypts a password.
//
// @param request - EncryptPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EncryptPasswordResponse
func (client *Client) EncryptPasswordWithOptions(request *EncryptPasswordRequest, runtime *dara.RuntimeOptions) (_result *EncryptPasswordResponse, _err error) {
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

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.LoginToken) {
		query["LoginToken"] = request.LoginToken
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EncryptPassword"),
		Version:     dara.String("2020-10-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EncryptPasswordResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Encrypts a password.
//
// @param request - EncryptPasswordRequest
//
// @return EncryptPasswordResponse
func (client *Client) EncryptPassword(request *EncryptPasswordRequest) (_result *EncryptPasswordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EncryptPasswordResponse{}
	_body, _err := client.EncryptPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the logon tokens for enterprise drives.
//
// @param request - GetCloudDriveServiceMountTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCloudDriveServiceMountTokenResponse
func (client *Client) GetCloudDriveServiceMountTokenWithOptions(request *GetCloudDriveServiceMountTokenRequest, runtime *dara.RuntimeOptions) (_result *GetCloudDriveServiceMountTokenResponse, _err error) {
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

	if !dara.IsNil(request.LoginToken) {
		query["LoginToken"] = request.LoginToken
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCloudDriveServiceMountToken"),
		Version:     dara.String("2020-10-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCloudDriveServiceMountTokenResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the logon tokens for enterprise drives.
//
// @param request - GetCloudDriveServiceMountTokenRequest
//
// @return GetCloudDriveServiceMountTokenResponse
func (client *Client) GetCloudDriveServiceMountToken(request *GetCloudDriveServiceMountTokenRequest) (_result *GetCloudDriveServiceMountTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCloudDriveServiceMountTokenResponse{}
	_body, _err := client.GetCloudDriveServiceMountTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the credential that is used to connect to a cloud computer.
//
// Description:
//
// The first time you call this operation, the system returns a task ID in the `TaskID` parameter. Use the task ID indicated in the `TaskID` parameter to continue calling this operation until the value of the `TaskStatus` parameter becomes `FINISHED` or `FAILED`. When `TaskStatus` becomes `FINISHED`, the value of the `Ticket` parameter is the ticket that is used to connect the client to the cloud computer.
//
// @param request - GetConnectionTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConnectionTicketResponse
func (client *Client) GetConnectionTicketWithOptions(request *GetConnectionTicketRequest, runtime *dara.RuntimeOptions) (_result *GetConnectionTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessType) {
		query["AccessType"] = request.AccessType
	}

	if !dara.IsNil(request.ClientId) {
		query["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientOS) {
		query["ClientOS"] = request.ClientOS
	}

	if !dara.IsNil(request.ClientType) {
		query["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.ClientVersion) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !dara.IsNil(request.CommandContent) {
		query["CommandContent"] = request.CommandContent
	}

	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.LoginToken) {
		query["LoginToken"] = request.LoginToken
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

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TicketBlackList) {
		query["TicketBlackList"] = request.TicketBlackList
	}

	if !dara.IsNil(request.Uuid) {
		query["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetConnectionTicket"),
		Version:     dara.String("2020-10-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetConnectionTicketResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the credential that is used to connect to a cloud computer.
//
// Description:
//
// The first time you call this operation, the system returns a task ID in the `TaskID` parameter. Use the task ID indicated in the `TaskID` parameter to continue calling this operation until the value of the `TaskStatus` parameter becomes `FINISHED` or `FAILED`. When `TaskStatus` becomes `FINISHED`, the value of the `Ticket` parameter is the ticket that is used to connect the client to the cloud computer.
//
// @param request - GetConnectionTicketRequest
//
// @return GetConnectionTicketResponse
func (client *Client) GetConnectionTicket(request *GetConnectionTicketRequest) (_result *GetConnectionTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetConnectionTicketResponse{}
	_body, _err := client.GetConnectionTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains logon credentials.
//
// @param tmpReq - GetLoginTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLoginTokenResponse
func (client *Client) GetLoginTokenWithOptions(tmpReq *GetLoginTokenRequest, runtime *dara.RuntimeOptions) (_result *GetLoginTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetLoginTokenShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AvailableFeatures) {
		request.AvailableFeaturesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AvailableFeatures, dara.String("AvailableFeatures"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthenticationCode) {
		query["AuthenticationCode"] = request.AuthenticationCode
	}

	if !dara.IsNil(request.AvailableFeaturesShrink) {
		query["AvailableFeatures"] = request.AvailableFeaturesShrink
	}

	if !dara.IsNil(request.ClientId) {
		query["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientName) {
		query["ClientName"] = request.ClientName
	}

	if !dara.IsNil(request.ClientOS) {
		query["ClientOS"] = request.ClientOS
	}

	if !dara.IsNil(request.ClientType) {
		query["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.ClientVersion) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !dara.IsNil(request.CurrentStage) {
		query["CurrentStage"] = request.CurrentStage
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.KeepAlive) {
		query["KeepAlive"] = request.KeepAlive
	}

	if !dara.IsNil(request.KeepAliveToken) {
		query["KeepAliveToken"] = request.KeepAliveToken
	}

	if !dara.IsNil(request.NewPassword) {
		query["NewPassword"] = request.NewPassword
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.OldPassword) {
		query["OldPassword"] = request.OldPassword
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.TokenCode) {
		query["TokenCode"] = request.TokenCode
	}

	if !dara.IsNil(request.Uuid) {
		query["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLoginToken"),
		Version:     dara.String("2020-10-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLoginTokenResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains logon credentials.
//
// @param request - GetLoginTokenRequest
//
// @return GetLoginTokenResponse
func (client *Client) GetLoginToken(request *GetLoginTokenRequest) (_result *GetLoginTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLoginTokenResponse{}
	_body, _err := client.GetLoginTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies whether the client\\"s logon session is still active.
//
// @param request - IsKeepAliveRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IsKeepAliveResponse
func (client *Client) IsKeepAliveWithOptions(request *IsKeepAliveRequest, runtime *dara.RuntimeOptions) (_result *IsKeepAliveResponse, _err error) {
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

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IsKeepAlive"),
		Version:     dara.String("2020-10-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &IsKeepAliveResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies whether the client\\"s logon session is still active.
//
// @param request - IsKeepAliveRequest
//
// @return IsKeepAliveResponse
func (client *Client) IsKeepAlive(request *IsKeepAliveRequest) (_result *IsKeepAliveResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &IsKeepAliveResponse{}
	_body, _err := client.IsKeepAliveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询Agent需要上报的配置信息
//
// @param request - QueryEdsAgentReportConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryEdsAgentReportConfigResponse
func (client *Client) QueryEdsAgentReportConfigWithOptions(request *QueryEdsAgentReportConfigRequest, runtime *dara.RuntimeOptions) (_result *QueryEdsAgentReportConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliUid) {
		query["AliUid"] = request.AliUid
	}

	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.EcsInstanceId) {
		query["EcsInstanceId"] = request.EcsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryEdsAgentReportConfig"),
		Version:     dara.String("2020-10-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryEdsAgentReportConfigResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Agent需要上报的配置信息
//
// @param request - QueryEdsAgentReportConfigRequest
//
// @return QueryEdsAgentReportConfigResponse
func (client *Client) QueryEdsAgentReportConfig(request *QueryEdsAgentReportConfigRequest) (_result *QueryEdsAgentReportConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryEdsAgentReportConfigResponse{}
	_body, _err := client.QueryEdsAgentReportConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restart cloud computers.
//
// @param request - RebootDesktopsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebootDesktopsResponse
func (client *Client) RebootDesktopsWithOptions(request *RebootDesktopsRequest, runtime *dara.RuntimeOptions) (_result *RebootDesktopsResponse, _err error) {
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

	if !dara.IsNil(request.ClientOS) {
		query["ClientOS"] = request.ClientOS
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ClientVersion) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.LoginToken) {
		query["LoginToken"] = request.LoginToken
	}

	if !dara.IsNil(request.OsUpdate) {
		query["OsUpdate"] = request.OsUpdate
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.SessionToken) {
		query["SessionToken"] = request.SessionToken
	}

	if !dara.IsNil(request.Uuid) {
		query["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RebootDesktops"),
		Version:     dara.String("2020-10-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RebootDesktopsResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restart cloud computers.
//
// @param request - RebootDesktopsRequest
//
// @return RebootDesktopsResponse
func (client *Client) RebootDesktops(request *RebootDesktopsRequest) (_result *RebootDesktopsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RebootDesktopsResponse{}
	_body, _err := client.RebootDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RefreshLoginTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshLoginTokenResponse
func (client *Client) RefreshLoginTokenWithOptions(request *RefreshLoginTokenRequest, runtime *dara.RuntimeOptions) (_result *RefreshLoginTokenResponse, _err error) {
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

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.LoginToken) {
		query["LoginToken"] = request.LoginToken
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RefreshLoginToken"),
		Version:     dara.String("2020-10-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RefreshLoginTokenResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RefreshLoginTokenRequest
//
// @return RefreshLoginTokenResponse
func (client *Client) RefreshLoginToken(request *RefreshLoginTokenRequest) (_result *RefreshLoginTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RefreshLoginTokenResponse{}
	_body, _err := client.RefreshLoginTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 上报edsAgent的信息
//
// @param request - ReportEdsAgentInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReportEdsAgentInfoResponse
func (client *Client) ReportEdsAgentInfoWithOptions(request *ReportEdsAgentInfoRequest, runtime *dara.RuntimeOptions) (_result *ReportEdsAgentInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliUid) {
		query["AliUid"] = request.AliUid
	}

	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.EcsInstanceId) {
		query["EcsInstanceId"] = request.EcsInstanceId
	}

	if !dara.IsNil(request.EdsAgentInfo) {
		query["EdsAgentInfo"] = request.EdsAgentInfo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReportEdsAgentInfo"),
		Version:     dara.String("2020-10-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReportEdsAgentInfoResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上报edsAgent的信息
//
// @param request - ReportEdsAgentInfoRequest
//
// @return ReportEdsAgentInfoResponse
func (client *Client) ReportEdsAgentInfo(request *ReportEdsAgentInfoRequest) (_result *ReportEdsAgentInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReportEdsAgentInfoResponse{}
	_body, _err := client.ReportEdsAgentInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ReportSessionStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReportSessionStatusResponse
func (client *Client) ReportSessionStatusWithOptions(request *ReportSessionStatusRequest, runtime *dara.RuntimeOptions) (_result *ReportSessionStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SessionChangeTime) {
		query["SessionChangeTime"] = request.SessionChangeTime
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.SessionStatus) {
		query["SessionStatus"] = request.SessionStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReportSessionStatus"),
		Version:     dara.String("2020-10-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReportSessionStatusResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ReportSessionStatusRequest
//
// @return ReportSessionStatusResponse
func (client *Client) ReportSessionStatus(request *ReportSessionStatusRequest) (_result *ReportSessionStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReportSessionStatusResponse{}
	_body, _err := client.ReportSessionStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resets a password.
//
// @param request - ResetPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetPasswordResponse
func (client *Client) ResetPasswordWithOptions(request *ResetPasswordRequest, runtime *dara.RuntimeOptions) (_result *ResetPasswordResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Phone) {
		query["phone"] = request.Phone
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetPassword"),
		Version:     dara.String("2020-10-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetPasswordResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets a password.
//
// @param request - ResetPasswordRequest
//
// @return ResetPasswordResponse
func (client *Client) ResetPassword(request *ResetPasswordRequest) (_result *ResetPasswordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetPasswordResponse{}
	_body, _err := client.ResetPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restores the data of a disk from a snapshot.
//
// @param request - ResetSnapshotRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetSnapshotResponse
func (client *Client) ResetSnapshotWithOptions(request *ResetSnapshotRequest, runtime *dara.RuntimeOptions) (_result *ResetSnapshotResponse, _err error) {
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

	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.LoginToken) {
		query["LoginToken"] = request.LoginToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.SnapshotId) {
		query["SnapshotId"] = request.SnapshotId
	}

	if !dara.IsNil(request.StopDesktop) {
		query["StopDesktop"] = request.StopDesktop
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetSnapshot"),
		Version:     dara.String("2020-10-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetSnapshotResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restores the data of a disk from a snapshot.
//
// @param request - ResetSnapshotRequest
//
// @return ResetSnapshotResponse
func (client *Client) ResetSnapshot(request *ResetSnapshotRequest) (_result *ResetSnapshotResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetSnapshotResponse{}
	_body, _err := client.ResetSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends a logon verification code.
//
// @param request - SendTokenCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendTokenCodeResponse
func (client *Client) SendTokenCodeWithOptions(request *SendTokenCodeRequest, runtime *dara.RuntimeOptions) (_result *SendTokenCodeResponse, _err error) {
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

	if !dara.IsNil(request.ClientOS) {
		query["ClientOS"] = request.ClientOS
	}

	if !dara.IsNil(request.ClientVersion) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.LoginToken) {
		query["LoginToken"] = request.LoginToken
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.TokenCode) {
		query["TokenCode"] = request.TokenCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendTokenCode"),
		Version:     dara.String("2020-10-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendTokenCodeResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a logon verification code.
//
// @param request - SendTokenCodeRequest
//
// @return SendTokenCodeResponse
func (client *Client) SendTokenCode(request *SendTokenCodeRequest) (_result *SendTokenCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SendTokenCodeResponse{}
	_body, _err := client.SendTokenCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SetFingerPrintTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetFingerPrintTemplateResponse
func (client *Client) SetFingerPrintTemplateWithOptions(request *SetFingerPrintTemplateRequest, runtime *dara.RuntimeOptions) (_result *SetFingerPrintTemplateResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EncryptedFingerPrintTemplate) {
		query["EncryptedFingerPrintTemplate"] = request.EncryptedFingerPrintTemplate
	}

	if !dara.IsNil(request.EncryptedKey) {
		query["EncryptedKey"] = request.EncryptedKey
	}

	if !dara.IsNil(request.FingerPrintTemplate) {
		query["FingerPrintTemplate"] = request.FingerPrintTemplate
	}

	if !dara.IsNil(request.LoginToken) {
		query["LoginToken"] = request.LoginToken
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetFingerPrintTemplate"),
		Version:     dara.String("2020-10-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetFingerPrintTemplateResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SetFingerPrintTemplateRequest
//
// @return SetFingerPrintTemplateResponse
func (client *Client) SetFingerPrintTemplate(request *SetFingerPrintTemplateRequest) (_result *SetFingerPrintTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetFingerPrintTemplateResponse{}
	_body, _err := client.SetFingerPrintTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SetFingerPrintTemplateDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetFingerPrintTemplateDescriptionResponse
func (client *Client) SetFingerPrintTemplateDescriptionWithOptions(request *SetFingerPrintTemplateDescriptionRequest, runtime *dara.RuntimeOptions) (_result *SetFingerPrintTemplateDescriptionResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Index) {
		query["Index"] = request.Index
	}

	if !dara.IsNil(request.LoginToken) {
		query["LoginToken"] = request.LoginToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetFingerPrintTemplateDescription"),
		Version:     dara.String("2020-10-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetFingerPrintTemplateDescriptionResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SetFingerPrintTemplateDescriptionRequest
//
// @return SetFingerPrintTemplateDescriptionResponse
func (client *Client) SetFingerPrintTemplateDescription(request *SetFingerPrintTemplateDescriptionRequest) (_result *SetFingerPrintTemplateDescriptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetFingerPrintTemplateDescriptionResponse{}
	_body, _err := client.SetFingerPrintTemplateDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Start cloud computers.
//
// Description:
//
// The cloud computers that you want to start must be in the Stopped state. After you call this operation, the cloud computers enter the Running state.
//
// @param request - StartDesktopsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartDesktopsResponse
func (client *Client) StartDesktopsWithOptions(request *StartDesktopsRequest, runtime *dara.RuntimeOptions) (_result *StartDesktopsResponse, _err error) {
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

	if !dara.IsNil(request.ClientOS) {
		query["ClientOS"] = request.ClientOS
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ClientVersion) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.LoginToken) {
		query["LoginToken"] = request.LoginToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Uuid) {
		query["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartDesktops"),
		Version:     dara.String("2020-10-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartDesktopsResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Start cloud computers.
//
// Description:
//
// The cloud computers that you want to start must be in the Stopped state. After you call this operation, the cloud computers enter the Running state.
//
// @param request - StartDesktopsRequest
//
// @return StartDesktopsResponse
func (client *Client) StartDesktops(request *StartDesktopsRequest) (_result *StartDesktopsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartDesktopsResponse{}
	_body, _err := client.StartDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - StartRecordContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartRecordContentResponse
func (client *Client) StartRecordContentWithOptions(request *StartRecordContentRequest, runtime *dara.RuntimeOptions) (_result *StartRecordContentResponse, _err error) {
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

	if !dara.IsNil(request.ClientOS) {
		query["ClientOS"] = request.ClientOS
	}

	if !dara.IsNil(request.ClientVersion) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.FilePath) {
		query["FilePath"] = request.FilePath
	}

	if !dara.IsNil(request.LoginToken) {
		query["LoginToken"] = request.LoginToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartRecordContent"),
		Version:     dara.String("2020-10-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartRecordContentResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - StartRecordContentRequest
//
// @return StartRecordContentResponse
func (client *Client) StartRecordContent(request *StartRecordContentRequest) (_result *StartRecordContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartRecordContentResponse{}
	_body, _err := client.StartRecordContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops cloud computers.
//
// Description:
//
// The cloud computers that you want to stop must be in the Running state. After you call this operation, the cloud computers enter the Stopped state.
//
// @param request - StopDesktopsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDesktopsResponse
func (client *Client) StopDesktopsWithOptions(request *StopDesktopsRequest, runtime *dara.RuntimeOptions) (_result *StopDesktopsResponse, _err error) {
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

	if !dara.IsNil(request.ClientOS) {
		query["ClientOS"] = request.ClientOS
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ClientVersion) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.LoginToken) {
		query["LoginToken"] = request.LoginToken
	}

	if !dara.IsNil(request.OsUpdate) {
		query["OsUpdate"] = request.OsUpdate
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.SessionToken) {
		query["SessionToken"] = request.SessionToken
	}

	if !dara.IsNil(request.Uuid) {
		query["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopDesktops"),
		Version:     dara.String("2020-10-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopDesktopsResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops cloud computers.
//
// Description:
//
// The cloud computers that you want to stop must be in the Running state. After you call this operation, the cloud computers enter the Stopped state.
//
// @param request - StopDesktopsRequest
//
// @return StopDesktopsResponse
func (client *Client) StopDesktops(request *StopDesktopsRequest) (_result *StopDesktopsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopDesktopsResponse{}
	_body, _err := client.StopDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - StopRecordContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopRecordContentResponse
func (client *Client) StopRecordContentWithOptions(request *StopRecordContentRequest, runtime *dara.RuntimeOptions) (_result *StopRecordContentResponse, _err error) {
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

	if !dara.IsNil(request.ClientOS) {
		query["ClientOS"] = request.ClientOS
	}

	if !dara.IsNil(request.ClientVersion) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.LoginToken) {
		query["LoginToken"] = request.LoginToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopRecordContent"),
		Version:     dara.String("2020-10-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopRecordContentResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - StopRecordContentRequest
//
// @return StopRecordContentResponse
func (client *Client) StopRecordContent(request *StopRecordContentRequest) (_result *StopRecordContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopRecordContentResponse{}
	_body, _err := client.StopRecordContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unbinds end users from cloud computers.
//
// @param request - UnbindUserDesktopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindUserDesktopResponse
func (client *Client) UnbindUserDesktopWithOptions(request *UnbindUserDesktopRequest, runtime *dara.RuntimeOptions) (_result *UnbindUserDesktopResponse, _err error) {
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

	if !dara.IsNil(request.ClientType) {
		query["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.LoginToken) {
		query["LoginToken"] = request.LoginToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.UserDesktopId) {
		query["UserDesktopId"] = request.UserDesktopId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindUserDesktop"),
		Version:     dara.String("2020-10-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindUserDesktopResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds end users from cloud computers.
//
// @param request - UnbindUserDesktopRequest
//
// @return UnbindUserDesktopResponse
func (client *Client) UnbindUserDesktop(request *UnbindUserDesktopRequest) (_result *UnbindUserDesktopResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnbindUserDesktopResponse{}
	_body, _err := client.UnbindUserDesktopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Validates credentials to lock on-premises sessions on clients.
//
// @param request - VerifyCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyCredentialResponse
func (client *Client) VerifyCredentialWithOptions(request *VerifyCredentialRequest, runtime *dara.RuntimeOptions) (_result *VerifyCredentialResponse, _err error) {
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

	if !dara.IsNil(request.Credential) {
		query["Credential"] = request.Credential
	}

	if !dara.IsNil(request.CredentialType) {
		query["CredentialType"] = request.CredentialType
	}

	if !dara.IsNil(request.EncryptedKey) {
		query["EncryptedKey"] = request.EncryptedKey
	}

	if !dara.IsNil(request.LoginToken) {
		query["LoginToken"] = request.LoginToken
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyCredential"),
		Version:     dara.String("2020-10-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyCredentialResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Validates credentials to lock on-premises sessions on clients.
//
// @param request - VerifyCredentialRequest
//
// @return VerifyCredentialResponse
func (client *Client) VerifyCredential(request *VerifyCredentialRequest) (_result *VerifyCredentialResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &VerifyCredentialResponse{}
	_body, _err := client.VerifyCredentialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
