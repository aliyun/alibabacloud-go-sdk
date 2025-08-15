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
	client.Endpoint, _err = client.GetEndpoint(dara.String("ossmfd"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 检查服务是否开通
//
// @param request - CheckMfdServiceOpenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckMfdServiceOpenResponse
func (client *Client) CheckMfdServiceOpenWithOptions(runtime *dara.RuntimeOptions) (_result *CheckMfdServiceOpenResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("CheckMfdServiceOpen"),
		Version:     dara.String("2023-10-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckMfdServiceOpenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检查服务是否开通
//
// @return CheckMfdServiceOpenResponse
func (client *Client) CheckMfdServiceOpen() (_result *CheckMfdServiceOpenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckMfdServiceOpenResponse{}
	_body, _err := client.CheckMfdServiceOpenWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建bucket扫描任务
//
// @param request - CreateOssBucketScanTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOssBucketScanTaskResponse
func (client *Client) CreateOssBucketScanTaskWithOptions(request *CreateOssBucketScanTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateOssBucketScanTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllKeyPrefix) {
		query["AllKeyPrefix"] = request.AllKeyPrefix
	}

	if !dara.IsNil(request.BucketNameList) {
		query["BucketNameList"] = request.BucketNameList
	}

	if !dara.IsNil(request.DecompressMaxFileCount) {
		query["DecompressMaxFileCount"] = request.DecompressMaxFileCount
	}

	if !dara.IsNil(request.DecompressMaxLayer) {
		query["DecompressMaxLayer"] = request.DecompressMaxLayer
	}

	if !dara.IsNil(request.DecryptionList) {
		query["DecryptionList"] = request.DecryptionList
	}

	if !dara.IsNil(request.ExcludeKeySuffixList) {
		query["ExcludeKeySuffixList"] = request.ExcludeKeySuffixList
	}

	if !dara.IsNil(request.KeyPrefixList) {
		query["KeyPrefixList"] = request.KeyPrefixList
	}

	if !dara.IsNil(request.KeySuffixList) {
		query["KeySuffixList"] = request.KeySuffixList
	}

	if !dara.IsNil(request.LastModifiedStartTime) {
		query["LastModifiedStartTime"] = request.LastModifiedStartTime
	}

	if !dara.IsNil(request.ScanMode) {
		query["ScanMode"] = request.ScanMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOssBucketScanTask"),
		Version:     dara.String("2023-10-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOssBucketScanTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建bucket扫描任务
//
// @param request - CreateOssBucketScanTaskRequest
//
// @return CreateOssBucketScanTaskResponse
func (client *Client) CreateOssBucketScanTask(request *CreateOssBucketScanTaskRequest) (_result *CreateOssBucketScanTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateOssBucketScanTaskResponse{}
	_body, _err := client.CreateOssBucketScanTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建Bucket检测策略配置
//
// @param request - CreateOssScanConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOssScanConfigResponse
func (client *Client) CreateOssScanConfigWithOptions(request *CreateOssScanConfigRequest, runtime *dara.RuntimeOptions) (_result *CreateOssScanConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllKeyPrefix) {
		query["AllKeyPrefix"] = request.AllKeyPrefix
	}

	if !dara.IsNil(request.BucketName) {
		query["BucketName"] = request.BucketName
	}

	if !dara.IsNil(request.DecompressMaxFileCount) {
		query["DecompressMaxFileCount"] = request.DecompressMaxFileCount
	}

	if !dara.IsNil(request.DecompressMaxLayer) {
		query["DecompressMaxLayer"] = request.DecompressMaxLayer
	}

	if !dara.IsNil(request.DecryptionList) {
		query["DecryptionList"] = request.DecryptionList
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.KeyPrefixList) {
		query["KeyPrefixList"] = request.KeyPrefixList
	}

	if !dara.IsNil(request.KeySuffixList) {
		query["KeySuffixList"] = request.KeySuffixList
	}

	if !dara.IsNil(request.LastModifiedStartTime) {
		query["LastModifiedStartTime"] = request.LastModifiedStartTime
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RealTimeIncr) {
		query["RealTimeIncr"] = request.RealTimeIncr
	}

	if !dara.IsNil(request.ScanDayList) {
		query["ScanDayList"] = request.ScanDayList
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOssScanConfig"),
		Version:     dara.String("2023-10-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOssScanConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建Bucket检测策略配置
//
// @param request - CreateOssScanConfigRequest
//
// @return CreateOssScanConfigResponse
func (client *Client) CreateOssScanConfig(request *CreateOssScanConfigRequest) (_result *CreateOssScanConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateOssScanConfigResponse{}
	_body, _err := client.CreateOssScanConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询导出信息
//
// @param request - DescribeExportInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExportInfoResponse
func (client *Client) DescribeExportInfoWithOptions(request *DescribeExportInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeExportInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExportId) {
		query["ExportId"] = request.ExportId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeExportInfo"),
		Version:     dara.String("2023-10-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExportInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询导出信息
//
// @param request - DescribeExportInfoRequest
//
// @return DescribeExportInfoResponse
func (client *Client) DescribeExportInfo(request *DescribeExportInfoRequest) (_result *DescribeExportInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeExportInfoResponse{}
	_body, _err := client.DescribeExportInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询您是否已创建云安全中心服务关联角色
//
// @param request - DescribeServiceLinkedRoleStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceLinkedRoleStatusResponse
func (client *Client) DescribeServiceLinkedRoleStatusWithOptions(request *DescribeServiceLinkedRoleStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeServiceLinkedRoleStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ServiceLinkedRole) {
		query["ServiceLinkedRole"] = request.ServiceLinkedRole
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeServiceLinkedRoleStatus"),
		Version:     dara.String("2023-10-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeServiceLinkedRoleStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询您是否已创建云安全中心服务关联角色
//
// @param request - DescribeServiceLinkedRoleStatusRequest
//
// @return DescribeServiceLinkedRoleStatusResponse
func (client *Client) DescribeServiceLinkedRoleStatus(request *DescribeServiceLinkedRoleStatusRequest) (_result *DescribeServiceLinkedRoleStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeServiceLinkedRoleStatusResponse{}
	_body, _err := client.DescribeServiceLinkedRoleStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 导出恶意文件列表
//
// @param request - ExportRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportRecordResponse
func (client *Client) ExportRecordWithOptions(request *ExportRecordRequest, runtime *dara.RuntimeOptions) (_result *ExportRecordResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExportType) {
		query["ExportType"] = request.ExportType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Params) {
		query["Params"] = request.Params
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportRecord"),
		Version:     dara.String("2023-10-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出恶意文件列表
//
// @param request - ExportRecordRequest
//
// @return ExportRecordResponse
func (client *Client) ExportRecord(request *ExportRecordRequest) (_result *ExportRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportRecordResponse{}
	_body, _err := client.ExportRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取沙箱检测报告。
//
// @param request - GetFileDetectReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFileDetectReportResponse
func (client *Client) GetFileDetectReportWithOptions(request *GetFileDetectReportRequest, runtime *dara.RuntimeOptions) (_result *GetFileDetectReportResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EventId) {
		query["EventId"] = request.EventId
	}

	if !dara.IsNil(request.Field) {
		query["Field"] = request.Field
	}

	if !dara.IsNil(request.FileHash) {
		query["FileHash"] = request.FileHash
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFileDetectReport"),
		Version:     dara.String("2023-10-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFileDetectReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取沙箱检测报告。
//
// @param request - GetFileDetectReportRequest
//
// @return GetFileDetectReportResponse
func (client *Client) GetFileDetectReport(request *GetFileDetectReportRequest) (_result *GetFileDetectReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetFileDetectReportResponse{}
	_body, _err := client.GetFileDetectReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取bucket检测统计信息
//
// @param request - GetOssBucketScanStatisticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOssBucketScanStatisticResponse
func (client *Client) GetOssBucketScanStatisticWithOptions(request *GetOssBucketScanStatisticRequest, runtime *dara.RuntimeOptions) (_result *GetOssBucketScanStatisticResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BucketNameList) {
		query["BucketNameList"] = request.BucketNameList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOssBucketScanStatistic"),
		Version:     dara.String("2023-10-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOssBucketScanStatisticResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取bucket检测统计信息
//
// @param request - GetOssBucketScanStatisticRequest
//
// @return GetOssBucketScanStatisticResponse
func (client *Client) GetOssBucketScanStatistic(request *GetOssBucketScanStatisticRequest) (_result *GetOssBucketScanStatisticResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOssBucketScanStatisticResponse{}
	_body, _err := client.GetOssBucketScanStatisticWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Bucket检测配置
//
// @param request - GetOssScanConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOssScanConfigResponse
func (client *Client) GetOssScanConfigWithOptions(request *GetOssScanConfigRequest, runtime *dara.RuntimeOptions) (_result *GetOssScanConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BucketName) {
		query["BucketName"] = request.BucketName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOssScanConfig"),
		Version:     dara.String("2023-10-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOssScanConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Bucket检测配置
//
// @param request - GetOssScanConfigRequest
//
// @return GetOssScanConfigResponse
func (client *Client) GetOssScanConfig(request *GetOssScanConfigRequest) (_result *GetOssScanConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOssScanConfigResponse{}
	_body, _err := client.GetOssScanConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文件检测告警列表
//
// @param request - ListObjectScanEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListObjectScanEventResponse
func (client *Client) ListObjectScanEventWithOptions(request *ListObjectScanEventRequest, runtime *dara.RuntimeOptions) (_result *ListObjectScanEventResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BucketName) {
		query["BucketName"] = request.BucketName
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EventName) {
		query["EventName"] = request.EventName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Md5) {
		query["Md5"] = request.Md5
	}

	if !dara.IsNil(request.OssKey) {
		query["OssKey"] = request.OssKey
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ParentEventId) {
		query["ParentEventId"] = request.ParentEventId
	}

	if !dara.IsNil(request.RiskLevel) {
		query["RiskLevel"] = request.RiskLevel
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.TimeEnd) {
		query["TimeEnd"] = request.TimeEnd
	}

	if !dara.IsNil(request.TimeStart) {
		query["TimeStart"] = request.TimeStart
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListObjectScanEvent"),
		Version:     dara.String("2023-10-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListObjectScanEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文件检测告警列表
//
// @param request - ListObjectScanEventRequest
//
// @return ListObjectScanEventResponse
func (client *Client) ListObjectScanEvent(request *ListObjectScanEventRequest) (_result *ListObjectScanEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListObjectScanEventResponse{}
	_body, _err := client.ListObjectScanEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列举用户所有的bucket
//
// @param request - ListOssBucketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOssBucketResponse
func (client *Client) ListOssBucketWithOptions(request *ListOssBucketRequest, runtime *dara.RuntimeOptions) (_result *ListOssBucketResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BucketName) {
		query["BucketName"] = request.BucketName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOssBucket"),
		Version:     dara.String("2023-10-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOssBucketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举用户所有的bucket
//
// @param request - ListOssBucketRequest
//
// @return ListOssBucketResponse
func (client *Client) ListOssBucket(request *ListOssBucketRequest) (_result *ListOssBucketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListOssBucketResponse{}
	_body, _err := client.ListOssBucketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取bucket详情
//
// @param request - ListOssBucketScanInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOssBucketScanInfoResponse
func (client *Client) ListOssBucketScanInfoWithOptions(request *ListOssBucketScanInfoRequest, runtime *dara.RuntimeOptions) (_result *ListOssBucketScanInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BucketName) {
		query["BucketName"] = request.BucketName
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.FuzzBucketName) {
		query["FuzzBucketName"] = request.FuzzBucketName
	}

	if !dara.IsNil(request.HasRisk) {
		query["HasRisk"] = request.HasRisk
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOssBucketScanInfo"),
		Version:     dara.String("2023-10-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOssBucketScanInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取bucket详情
//
// @param request - ListOssBucketScanInfoRequest
//
// @return ListOssBucketScanInfoResponse
func (client *Client) ListOssBucketScanInfo(request *ListOssBucketScanInfoRequest) (_result *ListOssBucketScanInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListOssBucketScanInfoResponse{}
	_body, _err := client.ListOssBucketScanInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取支持的文件后缀
//
// @param request - ListSupportObjectSuffixRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSupportObjectSuffixResponse
func (client *Client) ListSupportObjectSuffixWithOptions(runtime *dara.RuntimeOptions) (_result *ListSupportObjectSuffixResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("ListSupportObjectSuffix"),
		Version:     dara.String("2023-10-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSupportObjectSuffixResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取支持的文件后缀
//
// @return ListSupportObjectSuffixResponse
func (client *Client) ListSupportObjectSuffix() (_result *ListSupportObjectSuffixResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSupportObjectSuffixResponse{}
	_body, _err := client.ListSupportObjectSuffixWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 操作oss检测任务
//
// @param request - OperateBucketScanTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateBucketScanTaskResponse
func (client *Client) OperateBucketScanTaskWithOptions(request *OperateBucketScanTaskRequest, runtime *dara.RuntimeOptions) (_result *OperateBucketScanTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BucketName) {
		query["BucketName"] = request.BucketName
	}

	if !dara.IsNil(request.OperateCode) {
		query["OperateCode"] = request.OperateCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateBucketScanTask"),
		Version:     dara.String("2023-10-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateBucketScanTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 操作oss检测任务
//
// @param request - OperateBucketScanTaskRequest
//
// @return OperateBucketScanTaskResponse
func (client *Client) OperateBucketScanTask(request *OperateBucketScanTaskRequest) (_result *OperateBucketScanTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateBucketScanTaskResponse{}
	_body, _err := client.OperateBucketScanTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改Bucket检测配置
//
// @param request - UpdateOssScanConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOssScanConfigResponse
func (client *Client) UpdateOssScanConfigWithOptions(request *UpdateOssScanConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateOssScanConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllKeyPrefix) {
		query["AllKeyPrefix"] = request.AllKeyPrefix
	}

	if !dara.IsNil(request.BucketName) {
		query["BucketName"] = request.BucketName
	}

	if !dara.IsNil(request.DecompressMaxFileCount) {
		query["DecompressMaxFileCount"] = request.DecompressMaxFileCount
	}

	if !dara.IsNil(request.DecompressMaxLayer) {
		query["DecompressMaxLayer"] = request.DecompressMaxLayer
	}

	if !dara.IsNil(request.DecryptionList) {
		query["DecryptionList"] = request.DecryptionList
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.KeyPrefixList) {
		query["KeyPrefixList"] = request.KeyPrefixList
	}

	if !dara.IsNil(request.KeySuffixList) {
		query["KeySuffixList"] = request.KeySuffixList
	}

	if !dara.IsNil(request.LastModifiedStartTime) {
		query["LastModifiedStartTime"] = request.LastModifiedStartTime
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RealTimeIncr) {
		query["RealTimeIncr"] = request.RealTimeIncr
	}

	if !dara.IsNil(request.ScanDayList) {
		query["ScanDayList"] = request.ScanDayList
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateOssScanConfig"),
		Version:     dara.String("2023-10-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateOssScanConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改Bucket检测配置
//
// @param request - UpdateOssScanConfigRequest
//
// @return UpdateOssScanConfigResponse
func (client *Client) UpdateOssScanConfig(request *UpdateOssScanConfigRequest) (_result *UpdateOssScanConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateOssScanConfigResponse{}
	_body, _err := client.UpdateOssScanConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
