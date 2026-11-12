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
// 查询AI报告
//
// Description:
//
// 查询数据产品附带的 AI 报告状态，并在报告可用时返回临时下载地址。本接口是只读查询，不会再次调用数据服务方，也不会启动或重新生成报告。
//
// callId 与 reportId 必须且只能提供一个，参数名大小写以本页定义为准。按次报告：先调用 InvokeDataProduct 获取 callId，再用 callId 查询；报告状态为 AVAILABLE 时，从 Report 获取下载地址。周期报告：先调用 ListProductReports 获取 ReportId，再将该编号作为 reportId 查询；按周期商品的 callId 查询始终返回 WAITING_PERIOD，即使对应周期报告已经生成，也不会返回下载地址。内部任务编号不是本接口的查询键。
//
// 查询范围限定为当前云账号（主账号）所属的调用和报告；编号不存在或不属于该账号时，均返回 ReportNotFound（HTTP 404），不区分两种原因。使用 RAM 身份调用时，由网关校验 internetbank:GetAiReport 操作权限。
//
// 仅 ReportStatus 为 AVAILABLE 时返回 Report 对象。DownloadUrl 为临时签名下载地址，签发有效期不超过 300 秒，且不晚于 ContentExpiresAt；应在 UrlExpiresAt 前使用，不要公开分享。链接到期而报告内容仍在保留期内时，可以再次查询获取新链接。报告过期或被暂停交付时不返回下载地址；对象核验或签名等交付异常返回 ReportDeliveryFailed（HTTP 500）。参数不合法返回对应的 InvalidParameter 类错误（HTTP 400）。
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
// 查询AI报告
//
// Description:
//
// 查询数据产品附带的 AI 报告状态，并在报告可用时返回临时下载地址。本接口是只读查询，不会再次调用数据服务方，也不会启动或重新生成报告。
//
// callId 与 reportId 必须且只能提供一个，参数名大小写以本页定义为准。按次报告：先调用 InvokeDataProduct 获取 callId，再用 callId 查询；报告状态为 AVAILABLE 时，从 Report 获取下载地址。周期报告：先调用 ListProductReports 获取 ReportId，再将该编号作为 reportId 查询；按周期商品的 callId 查询始终返回 WAITING_PERIOD，即使对应周期报告已经生成，也不会返回下载地址。内部任务编号不是本接口的查询键。
//
// 查询范围限定为当前云账号（主账号）所属的调用和报告；编号不存在或不属于该账号时，均返回 ReportNotFound（HTTP 404），不区分两种原因。使用 RAM 身份调用时，由网关校验 internetbank:GetAiReport 操作权限。
//
// 仅 ReportStatus 为 AVAILABLE 时返回 Report 对象。DownloadUrl 为临时签名下载地址，签发有效期不超过 300 秒，且不晚于 ContentExpiresAt；应在 UrlExpiresAt 前使用，不要公开分享。链接到期而报告内容仍在保留期内时，可以再次查询获取新链接。报告过期或被暂停交付时不返回下载地址；对象核验或签名等交付异常返回 ReportDeliveryFailed（HTTP 500）。参数不合法返回对应的 InvalidParameter 类错误（HTTP 400）。
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
// Provides a unified entry point to invoke data products.
//
// Description:
//
// Provides a unified entry point to invoke data products. After the gateway completes identity parsing, RAM authentication, and product authorization, it forwards the request to the data service provider corresponding to the data product and returns the business data. If the invocation is successful, HTTP 200 is returned, and the response body contains RequestId and Data. If the invocation fails, the corresponding HTTP status code is returned (400 for parameter errors, 403 for no permissions or inactive products, 409 if the first request with the same ClientToken is still being processed, 429 for throttled requests, 503 for temporarily unavailable services, and 500 for internal errors), and the response body includes Code and Message to explain the cause.
//
// This operation accepts only three business parameters: ClientToken, ProductCode, and ServiceParameters. It does not accept TaskId, Operation, parameters starting with Ai, or EnableAiArtifact (including when set to false). If these are passed, an InvalidParameter error is returned. Whether a product includes an AI report is determined by the product configuration corresponding to the ProductCode. For products that include reports, you can call ListProductReports to query the report list, and use GetAiReport to retrieve the report by CallId or ReportId.
//
// Prerequisites for invocation: When calling this operation with a RAM user, you must first grant the RAM user the operation permissions for this operation in RAM. Set Action to internetbank:InvokeDataProduct and Resource to acs:internetbank:*:<your Alibaba Cloud account UID>:*. No additional authorization is required when calling this operation directly with the AccessKey of an Alibaba Cloud account.
//
// Service-linked role: After the preceding authorization is completed, the system performs automatic creation of the service-linked role (SLR) AliyunServiceRoleForInternetBankPredict under your Alibaba Cloud account. This role is used to invoke the corresponding data service provider on behalf of the customer as a service. You can view the invocation records that associate with your account in ActionTrail.
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
// Provides a unified entry point to invoke data products.
//
// Description:
//
// Provides a unified entry point to invoke data products. After the gateway completes identity parsing, RAM authentication, and product authorization, it forwards the request to the data service provider corresponding to the data product and returns the business data. If the invocation is successful, HTTP 200 is returned, and the response body contains RequestId and Data. If the invocation fails, the corresponding HTTP status code is returned (400 for parameter errors, 403 for no permissions or inactive products, 409 if the first request with the same ClientToken is still being processed, 429 for throttled requests, 503 for temporarily unavailable services, and 500 for internal errors), and the response body includes Code and Message to explain the cause.
//
// This operation accepts only three business parameters: ClientToken, ProductCode, and ServiceParameters. It does not accept TaskId, Operation, parameters starting with Ai, or EnableAiArtifact (including when set to false). If these are passed, an InvalidParameter error is returned. Whether a product includes an AI report is determined by the product configuration corresponding to the ProductCode. For products that include reports, you can call ListProductReports to query the report list, and use GetAiReport to retrieve the report by CallId or ReportId.
//
// Prerequisites for invocation: When calling this operation with a RAM user, you must first grant the RAM user the operation permissions for this operation in RAM. Set Action to internetbank:InvokeDataProduct and Resource to acs:internetbank:*:<your Alibaba Cloud account UID>:*. No additional authorization is required when calling this operation directly with the AccessKey of an Alibaba Cloud account.
//
// Service-linked role: After the preceding authorization is completed, the system performs automatic creation of the service-linked role (SLR) AliyunServiceRoleForInternetBankPredict under your Alibaba Cloud account. This role is used to invoke the corresponding data service provider on behalf of the customer as a service. You can view the invocation records that associate with your account in ActionTrail.
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
// 查询商品周期报告列表
//
// Description:
//
// 查询当前云账号（主账号）下指定商品已经生成的可用周期报告。列表只用于发现报告编号和所属周期，不返回下载地址；取得 ReportId 后，调用 GetAiReport，并将该编号作为 reportId 获取最新状态及下载地址。本接口不会调用数据服务方，也不会触发报告生成。
//
// LatestOnly 省略或为 false 时，在最近十个已结束的北京时间自然周范围内查询，按周期从新到旧返回，最多十条报告记录；没有报告的周不会补空记录，不保证返回十份报告。LatestOnly 为 true 时，从所有已结束的历史周期中查找最近一条可用报告，最多返回一条，不受最近十周范围限制。
//
// 自然周按北京时间（Asia/Shanghai）计算，从周一 00:00:00 到下一周周一 00:00:00，区间左闭右开，不包含当前尚未结束的自然周。PeriodStart 与 PeriodEnd 使用 UTC RFC 3339 格式输出。一个自然周可能对应多条报告记录，因此两条记录的周期起止相同不表示重复，使用 ReportId 区分。
//
// 列表依据已发布报告的资产记录判断可用性，不返回已过期或被暂停交付的报告；实际下载状态以 GetAiReport 查询结果为准。没有可用报告时，成功返回空 Reports 数组，不返回 ReportNotFound。本接口只提供 ProductCode 和 LatestOnly 两个业务参数，不支持分页及自定义周期过滤；传入 PeriodStart、PageSize 或 NextToken 时返回对应的 InvalidParameter 类错误（HTTP 400）。
//
// 使用 RAM 身份调用时，由网关校验 internetbank:ListProductReports 操作权限。查询只读取当前账号的既有报告，不再次检查商品商业开通状态。参数错误返回对应的 InvalidParameter 类错误（HTTP 400）；报告资产记录不一致或查询异常返回 ReportDeliveryFailed（HTTP 500）。
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
// 查询商品周期报告列表
//
// Description:
//
// 查询当前云账号（主账号）下指定商品已经生成的可用周期报告。列表只用于发现报告编号和所属周期，不返回下载地址；取得 ReportId 后，调用 GetAiReport，并将该编号作为 reportId 获取最新状态及下载地址。本接口不会调用数据服务方，也不会触发报告生成。
//
// LatestOnly 省略或为 false 时，在最近十个已结束的北京时间自然周范围内查询，按周期从新到旧返回，最多十条报告记录；没有报告的周不会补空记录，不保证返回十份报告。LatestOnly 为 true 时，从所有已结束的历史周期中查找最近一条可用报告，最多返回一条，不受最近十周范围限制。
//
// 自然周按北京时间（Asia/Shanghai）计算，从周一 00:00:00 到下一周周一 00:00:00，区间左闭右开，不包含当前尚未结束的自然周。PeriodStart 与 PeriodEnd 使用 UTC RFC 3339 格式输出。一个自然周可能对应多条报告记录，因此两条记录的周期起止相同不表示重复，使用 ReportId 区分。
//
// 列表依据已发布报告的资产记录判断可用性，不返回已过期或被暂停交付的报告；实际下载状态以 GetAiReport 查询结果为准。没有可用报告时，成功返回空 Reports 数组，不返回 ReportNotFound。本接口只提供 ProductCode 和 LatestOnly 两个业务参数，不支持分页及自定义周期过滤；传入 PeriodStart、PageSize 或 NextToken 时返回对应的 InvalidParameter 类错误（HTTP 400）。
//
// 使用 RAM 身份调用时，由网关校验 internetbank:ListProductReports 操作权限。查询只读取当前账号的既有报告，不再次检查商品商业开通状态。参数错误返回对应的 InvalidParameter 类错误（HTTP 400）；报告资产记录不一致或查询异常返回 ReportDeliveryFailed（HTTP 500）。
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
