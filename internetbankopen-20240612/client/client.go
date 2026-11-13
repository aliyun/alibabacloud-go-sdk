// This file is auto-generated, don't edit it. Thanks.
package client

import (
	gatewayclient "github.com/alibabacloud-go/alibabacloud-gateway-pop/client"
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
	client.ProductId = dara.String("InternetBankOpen")
	gatewayClient, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = gatewayClient
	client.EndpointRule = dara.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-beijing": dara.String("internetbankopen.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("internetbankopen"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Queries the report status and retrieves a temporary download URL for available reports.
//
// Description:
//
// Queries the status of the AI report attached to a data product and returns a temporary download URL when the report is available. This operation is a read-only query. It does not call the data service provider again, nor does it start or regenerate the report.
//
// You must specify either callId or reportId, but not both. The parameter names are case-sensitive as defined on this page. For per-use reports: first call InvokeDataProduct to obtain the callId, and then use the callId to query. When the report status is AVAILABLE, obtain the download URL from Report. For periodic reports: first call ListProductReports to obtain the ReportId, and then use this ID as reportId to query. Querying by the callId of a periodic product always returns WAITING_PERIOD. Even if the corresponding periodic report has been generated, no download URL is returned. Internal task IDs are not query keys for this operation.
//
// The query scope is limited to the calls and reports that belong to the current Alibaba Cloud account. If the ID does not exist or does not belong to the account, ReportNotFound (HTTP 404) is returned without distinguishing between the two causes. When you call this operation as a RAM user, the gateway verifies the internetbank:GetAiReport operation permission.
//
// The Report object is returned only when ReportStatus is AVAILABLE. DownloadUrl is a temporary signed download URL. The validity period of the signed URL does not exceed 300 seconds and is no later than ContentExpiresAt. Use the URL before UrlExpiresAt and do not share it publicly. If the URL expires but the report content is still within the retention period, you can query again to obtain a new URL. No download URL is returned when the report expires or its delivery is suspended. Delivery exceptions such as object verification or signing failures return ReportDeliveryFailed (HTTP 500). Invalid parameters return the corresponding InvalidParameter errors (HTTP 400).
//
// @param request - GetAiReportRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAiReportResponse
func (client *Client) GetAiReportWithOptions(request *GetAiReportRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAiReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallId) {
		query["callId"] = request.CallId
	}

	if !dara.IsNil(request.ReportId) {
		query["reportId"] = request.ReportId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAiReport"),
		Version:     dara.String("2024-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/api/v1/ai-report"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAiReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the report status and retrieves a temporary download URL for available reports.
//
// Description:
//
// Queries the status of the AI report attached to a data product and returns a temporary download URL when the report is available. This operation is a read-only query. It does not call the data service provider again, nor does it start or regenerate the report.
//
// You must specify either callId or reportId, but not both. The parameter names are case-sensitive as defined on this page. For per-use reports: first call InvokeDataProduct to obtain the callId, and then use the callId to query. When the report status is AVAILABLE, obtain the download URL from Report. For periodic reports: first call ListProductReports to obtain the ReportId, and then use this ID as reportId to query. Querying by the callId of a periodic product always returns WAITING_PERIOD. Even if the corresponding periodic report has been generated, no download URL is returned. Internal task IDs are not query keys for this operation.
//
// The query scope is limited to the calls and reports that belong to the current Alibaba Cloud account. If the ID does not exist or does not belong to the account, ReportNotFound (HTTP 404) is returned without distinguishing between the two causes. When you call this operation as a RAM user, the gateway verifies the internetbank:GetAiReport operation permission.
//
// The Report object is returned only when ReportStatus is AVAILABLE. DownloadUrl is a temporary signed download URL. The validity period of the signed URL does not exceed 300 seconds and is no later than ContentExpiresAt. Use the URL before UrlExpiresAt and do not share it publicly. If the URL expires but the report content is still within the retention period, you can query again to obtain a new URL. No download URL is returned when the report expires or its delivery is suspended. Delivery exceptions such as object verification or signing failures return ReportDeliveryFailed (HTTP 500). Invalid parameters return the corresponding InvalidParameter errors (HTTP 400).
//
// @param request - GetAiReportRequest
//
// @return GetAiReportResponse
func (client *Client) GetAiReport(request *GetAiReportRequest) (_result *GetAiReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAiReportResponse{}
	_body, _err := client.GetAiReportWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Provides a unified entry point for invoking data products.
//
// Description:
//
// Provides a unified entry point for invoking data products. After the gateway completes identity parsing, RAM authentication, and product authorization, it forwards the request to the data service provider corresponding to the data product and returns the business data. A successful call returns an HTTP 200 status code, and the response body contains RequestId and Data. A failed call returns the corresponding HTTP status code (400 for parameter errors, 403 for no permissions or unactivated services, 409 if the initial request with the same ClientToken is still being processed, 429 for throttled requests, 503 for temporarily unavailable services, and 500 for internal errors), and the response body uses Code and Message to explain the cause.
//
// This operation accepts only three business parameters: ClientToken, ProductCode, and ServiceParameters. It does not accept TaskId, Operation, parameters starting with Ai, or EnableAiArtifact (including when set to false). Passing these parameters returns an InvalidParameter error. Whether a product includes an AI report is determined by the product configuration corresponding to the ProductCode. For products that include reports, you can call ListProductReports to query the report list, and use GetAiReport to retrieve the report by CallId or ReportId.
//
// Authorization prerequisites: When calling this operation with a RAM user, you must first grant the RAM user the operation permissions for this operation in RAM. Set Action to internetbank:InvokeDataProduct and Resource to acs:internetbank:*:<your Alibaba Cloud account UID>:*. No additional authorization is required when calling this operation directly with the AccessKey of an Alibaba Cloud account.
//
// Service-linked role: After the preceding authorization is completed, the system automatically creates the service-linked role (SLR) AliyunServiceRoleForInternetBankPredict under your Alibaba Cloud account. This role is used to invoke the corresponding data service provider on behalf of the customer as a service identity. You can view the related invocation records in ActionTrail of your account.
//
// @param request - InvokeDataProductRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvokeDataProductResponse
func (client *Client) InvokeDataProductWithOptions(request *InvokeDataProductRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InvokeDataProductResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ProductCode) {
		query["productCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ServiceParameters) {
		query["serviceParameters"] = request.ServiceParameters
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InvokeDataProduct"),
		Version:     dara.String("2024-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/api/v1/rpc"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InvokeDataProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Provides a unified entry point for invoking data products.
//
// Description:
//
// Provides a unified entry point for invoking data products. After the gateway completes identity parsing, RAM authentication, and product authorization, it forwards the request to the data service provider corresponding to the data product and returns the business data. A successful call returns an HTTP 200 status code, and the response body contains RequestId and Data. A failed call returns the corresponding HTTP status code (400 for parameter errors, 403 for no permissions or unactivated services, 409 if the initial request with the same ClientToken is still being processed, 429 for throttled requests, 503 for temporarily unavailable services, and 500 for internal errors), and the response body uses Code and Message to explain the cause.
//
// This operation accepts only three business parameters: ClientToken, ProductCode, and ServiceParameters. It does not accept TaskId, Operation, parameters starting with Ai, or EnableAiArtifact (including when set to false). Passing these parameters returns an InvalidParameter error. Whether a product includes an AI report is determined by the product configuration corresponding to the ProductCode. For products that include reports, you can call ListProductReports to query the report list, and use GetAiReport to retrieve the report by CallId or ReportId.
//
// Authorization prerequisites: When calling this operation with a RAM user, you must first grant the RAM user the operation permissions for this operation in RAM. Set Action to internetbank:InvokeDataProduct and Resource to acs:internetbank:*:<your Alibaba Cloud account UID>:*. No additional authorization is required when calling this operation directly with the AccessKey of an Alibaba Cloud account.
//
// Service-linked role: After the preceding authorization is completed, the system automatically creates the service-linked role (SLR) AliyunServiceRoleForInternetBankPredict under your Alibaba Cloud account. This role is used to invoke the corresponding data service provider on behalf of the customer as a service identity. You can view the related invocation records in ActionTrail of your account.
//
// @param request - InvokeDataProductRequest
//
// @return InvokeDataProductResponse
func (client *Client) InvokeDataProduct(request *InvokeDataProductRequest) (_result *InvokeDataProductResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InvokeDataProductResponse{}
	_body, _err := client.InvokeDataProductWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the available periodic report IDs and period ranges for a specified product.
//
// Description:
//
// Queries the available periodic reports generated for a specified product under the current Alibaba Cloud account. The list is used only to discover report IDs and their corresponding periods, and does not return download URLs. After obtaining a ReportId, call GetAiReport and pass the ID as the reportId parameter to obtain the latest status and download URL. This operation does not call the data service provider or trigger report generation.
//
// If LatestOnly is omitted or set to false, the operation queries reports within the last ten completed natural weeks in UTC+8 and returns up to ten report records in descending order of epoch. Empty records are not padded for weeks without reports, and ten reports are not guaranteed to be returned. If LatestOnly is set to true, the operation searches for the most recent active report from all completed historical epochs and returns at most one record, without being limited to the last ten weeks.
//
// A natural week is calculated based on UTC+8 (Asia/Shanghai), starting from 00:00:00 on Monday to 00:00:00 on the following Monday. The interval is left-closed and right-open, excluding the current unfinished natural week. PeriodStart and PeriodEnd are output in UTC RFC 3339 format. A single natural week may correspond to multiple report records. Therefore, identical epoch start and end times for two records do not indicate duplication. Use ReportId to distinguish them.
//
// The list determines availability based on the asset records of published reports and does not return expired or suspended reports. The actual download status is subject to the query results of GetAiReport. If no available reports exist, an empty Reports array is returned successfully, and ReportNotFound is not returned. This operation provides only two business parameters, ProductCode and LatestOnly, and does not support pagination or custom period filtering. Passing PeriodStart, PageSize, or NextToken returns a corresponding InvalidParameter error (HTTP 400).
//
// When called with a RAM identity, the gateway verifies the internetbank:ListProductReports operation permission. The query only reads existing reports of the current account and does not recheck the commercial activation status of the product. Parameter errors return a corresponding InvalidParameter error (HTTP 400). Inconsistent report asset records or query exceptions return ReportDeliveryFailed (HTTP 500).
//
// @param request - ListProductReportsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProductReportsResponse
func (client *Client) ListProductReportsWithOptions(request *ListProductReportsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListProductReportsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LatestOnly) {
		query["LatestOnly"] = request.LatestOnly
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProductReports"),
		Version:     dara.String("2024-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/api/v1/product-reports"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProductReportsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the available periodic report IDs and period ranges for a specified product.
//
// Description:
//
// Queries the available periodic reports generated for a specified product under the current Alibaba Cloud account. The list is used only to discover report IDs and their corresponding periods, and does not return download URLs. After obtaining a ReportId, call GetAiReport and pass the ID as the reportId parameter to obtain the latest status and download URL. This operation does not call the data service provider or trigger report generation.
//
// If LatestOnly is omitted or set to false, the operation queries reports within the last ten completed natural weeks in UTC+8 and returns up to ten report records in descending order of epoch. Empty records are not padded for weeks without reports, and ten reports are not guaranteed to be returned. If LatestOnly is set to true, the operation searches for the most recent active report from all completed historical epochs and returns at most one record, without being limited to the last ten weeks.
//
// A natural week is calculated based on UTC+8 (Asia/Shanghai), starting from 00:00:00 on Monday to 00:00:00 on the following Monday. The interval is left-closed and right-open, excluding the current unfinished natural week. PeriodStart and PeriodEnd are output in UTC RFC 3339 format. A single natural week may correspond to multiple report records. Therefore, identical epoch start and end times for two records do not indicate duplication. Use ReportId to distinguish them.
//
// The list determines availability based on the asset records of published reports and does not return expired or suspended reports. The actual download status is subject to the query results of GetAiReport. If no available reports exist, an empty Reports array is returned successfully, and ReportNotFound is not returned. This operation provides only two business parameters, ProductCode and LatestOnly, and does not support pagination or custom period filtering. Passing PeriodStart, PageSize, or NextToken returns a corresponding InvalidParameter error (HTTP 400).
//
// When called with a RAM identity, the gateway verifies the internetbank:ListProductReports operation permission. The query only reads existing reports of the current account and does not recheck the commercial activation status of the product. Parameter errors return a corresponding InvalidParameter error (HTTP 400). Inconsistent report asset records or query exceptions return ReportDeliveryFailed (HTTP 500).
//
// @param request - ListProductReportsRequest
//
// @return ListProductReportsResponse
func (client *Client) ListProductReports(request *ListProductReportsRequest) (_result *ListProductReportsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProductReportsResponse{}
	_body, _err := client.ListProductReportsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
