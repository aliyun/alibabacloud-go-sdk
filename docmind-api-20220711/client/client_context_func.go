// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 整票识别
//
// @param tmpReq - AyncTradeDocumentPackageExtractSmartAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AyncTradeDocumentPackageExtractSmartAppResponse
func (client *Client) AyncTradeDocumentPackageExtractSmartAppWithContext(ctx context.Context, tmpReq *AyncTradeDocumentPackageExtractSmartAppRequest, runtime *dara.RuntimeOptions) (_result *AyncTradeDocumentPackageExtractSmartAppResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &AyncTradeDocumentPackageExtractSmartAppShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CustomExtractionRange) {
		request.CustomExtractionRangeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomExtractionRange, dara.String("CustomExtractionRange"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomExtractionRangeShrink) {
		query["CustomExtractionRange"] = request.CustomExtractionRangeShrink
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileUrl) {
		query["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.Option) {
		query["Option"] = request.Option
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AyncTradeDocumentPackageExtractSmartApp"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AyncTradeDocumentPackageExtractSmartAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档结构化流式接口
//
// @param request - GetDocParserResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocParserResultResponse
func (client *Client) GetDocParserResultWithContext(ctx context.Context, request *GetDocParserResultRequest, runtime *dara.RuntimeOptions) (_result *GetDocParserResultResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.LayoutNum) {
		query["LayoutNum"] = request.LayoutNum
	}

	if !dara.IsNil(request.LayoutStepSize) {
		query["LayoutStepSize"] = request.LayoutStepSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDocParserResult"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDocParserResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档智能解析结果查询
//
// @param request - GetDocStructureResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocStructureResultResponse
func (client *Client) GetDocStructureResultWithContext(ctx context.Context, request *GetDocStructureResultRequest, runtime *dara.RuntimeOptions) (_result *GetDocStructureResultResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.ImageStrategy) {
		query["ImageStrategy"] = request.ImageStrategy
	}

	if !dara.IsNil(request.RevealMarkdown) {
		query["RevealMarkdown"] = request.RevealMarkdown
	}

	if !dara.IsNil(request.UseUrlResponseBody) {
		query["UseUrlResponseBody"] = request.UseUrlResponseBody
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDocStructureResult"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDocStructureResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档对比结果查询
//
// @param request - GetDocumentCompareResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentCompareResultResponse
func (client *Client) GetDocumentCompareResultWithContext(ctx context.Context, request *GetDocumentCompareResultRequest, runtime *dara.RuntimeOptions) (_result *GetDocumentCompareResultResponse, _err error) {
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
		Action:      dara.String("GetDocumentCompareResult"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDocumentCompareResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档转换结果查询
//
// @param request - GetDocumentConvertResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentConvertResultResponse
func (client *Client) GetDocumentConvertResultWithContext(ctx context.Context, request *GetDocumentConvertResultRequest, runtime *dara.RuntimeOptions) (_result *GetDocumentConvertResultResponse, _err error) {
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
		Action:      dara.String("GetDocumentConvertResult"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDocumentConvertResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档抽取结果查询
//
// @param request - GetDocumentExtractResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentExtractResultResponse
func (client *Client) GetDocumentExtractResultWithContext(ctx context.Context, request *GetDocumentExtractResultRequest, runtime *dara.RuntimeOptions) (_result *GetDocumentExtractResultResponse, _err error) {
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
		Action:      dara.String("GetDocumentExtractResult"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDocumentExtractResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// openmind
//
// @param request - GetPageNumRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPageNumResponse
func (client *Client) GetPageNumWithContext(ctx context.Context, request *GetPageNumRequest, runtime *dara.RuntimeOptions) (_result *GetPageNumResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPageNum"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPageNumResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 表格智能解析结果查询
//
// @param request - GetTableUnderstandingResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableUnderstandingResultResponse
func (client *Client) GetTableUnderstandingResultWithContext(ctx context.Context, request *GetTableUnderstandingResultRequest, runtime *dara.RuntimeOptions) (_result *GetTableUnderstandingResultResponse, _err error) {
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
		Action:      dara.String("GetTableUnderstandingResult"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTableUnderstandingResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文档智能解析处理状态
//
// @param request - QueryDocParserStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDocParserStatusResponse
func (client *Client) QueryDocParserStatusWithContext(ctx context.Context, request *QueryDocParserStatusRequest, runtime *dara.RuntimeOptions) (_result *QueryDocParserStatusResponse, _err error) {
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
		Action:      dara.String("QueryDocParserStatus"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDocParserStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 图片转excel
//
// @param tmpReq - SubmitConvertImageToExcelJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitConvertImageToExcelJobResponse
func (client *Client) SubmitConvertImageToExcelJobWithContext(ctx context.Context, tmpReq *SubmitConvertImageToExcelJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitConvertImageToExcelJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SubmitConvertImageToExcelJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ImageNames) {
		request.ImageNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImageNames, dara.String("ImageNames"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.ImageUrls) {
		request.ImageUrlsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImageUrls, dara.String("ImageUrls"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ForceMergeExcel) {
		query["ForceMergeExcel"] = request.ForceMergeExcel
	}

	if !dara.IsNil(request.ImageNameExtension) {
		query["ImageNameExtension"] = request.ImageNameExtension
	}

	if !dara.IsNil(request.ImageNamesShrink) {
		query["ImageNames"] = request.ImageNamesShrink
	}

	if !dara.IsNil(request.ImageUrlsShrink) {
		query["ImageUrls"] = request.ImageUrlsShrink
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitConvertImageToExcelJob"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitConvertImageToExcelJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 图片转markdown
//
// @param tmpReq - SubmitConvertImageToMarkdownJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitConvertImageToMarkdownJobResponse
func (client *Client) SubmitConvertImageToMarkdownJobWithContext(ctx context.Context, tmpReq *SubmitConvertImageToMarkdownJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitConvertImageToMarkdownJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SubmitConvertImageToMarkdownJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ImageNames) {
		request.ImageNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImageNames, dara.String("ImageNames"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.ImageUrls) {
		request.ImageUrlsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImageUrls, dara.String("ImageUrls"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageNameExtension) {
		query["ImageNameExtension"] = request.ImageNameExtension
	}

	if !dara.IsNil(request.ImageNamesShrink) {
		query["ImageNames"] = request.ImageNamesShrink
	}

	if !dara.IsNil(request.ImageUrlsShrink) {
		query["ImageUrls"] = request.ImageUrlsShrink
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitConvertImageToMarkdownJob"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitConvertImageToMarkdownJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 图片转pdf
//
// @param tmpReq - SubmitConvertImageToPdfJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitConvertImageToPdfJobResponse
func (client *Client) SubmitConvertImageToPdfJobWithContext(ctx context.Context, tmpReq *SubmitConvertImageToPdfJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitConvertImageToPdfJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SubmitConvertImageToPdfJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ImageNames) {
		request.ImageNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImageNames, dara.String("ImageNames"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.ImageUrls) {
		request.ImageUrlsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImageUrls, dara.String("ImageUrls"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageNameExtension) {
		query["ImageNameExtension"] = request.ImageNameExtension
	}

	if !dara.IsNil(request.ImageNamesShrink) {
		query["ImageNames"] = request.ImageNamesShrink
	}

	if !dara.IsNil(request.ImageUrlsShrink) {
		query["ImageUrls"] = request.ImageUrlsShrink
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitConvertImageToPdfJob"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitConvertImageToPdfJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 图片转word
//
// @param tmpReq - SubmitConvertImageToWordJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitConvertImageToWordJobResponse
func (client *Client) SubmitConvertImageToWordJobWithContext(ctx context.Context, tmpReq *SubmitConvertImageToWordJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitConvertImageToWordJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SubmitConvertImageToWordJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ImageNames) {
		request.ImageNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImageNames, dara.String("ImageNames"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.ImageUrls) {
		request.ImageUrlsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImageUrls, dara.String("ImageUrls"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageNameExtension) {
		query["ImageNameExtension"] = request.ImageNameExtension
	}

	if !dara.IsNil(request.ImageNamesShrink) {
		query["ImageNames"] = request.ImageNamesShrink
	}

	if !dara.IsNil(request.ImageUrlsShrink) {
		query["ImageUrls"] = request.ImageUrlsShrink
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitConvertImageToWordJob"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitConvertImageToWordJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// pdf转excel
//
// @param request - SubmitConvertPdfToExcelJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitConvertPdfToExcelJobResponse
func (client *Client) SubmitConvertPdfToExcelJobWithContext(ctx context.Context, request *SubmitConvertPdfToExcelJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitConvertPdfToExcelJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileUrl) {
		query["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.ForceExportInnerImage) {
		query["ForceExportInnerImage"] = request.ForceExportInnerImage
	}

	if !dara.IsNil(request.ForceMergeExcel) {
		query["ForceMergeExcel"] = request.ForceMergeExcel
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitConvertPdfToExcelJob"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitConvertPdfToExcelJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// pdf转图片
//
// @param request - SubmitConvertPdfToImageJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitConvertPdfToImageJobResponse
func (client *Client) SubmitConvertPdfToImageJobWithContext(ctx context.Context, request *SubmitConvertPdfToImageJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitConvertPdfToImageJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileUrl) {
		query["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitConvertPdfToImageJob"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitConvertPdfToImageJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// pdf转markdown
//
// @param request - SubmitConvertPdfToMarkdownJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitConvertPdfToMarkdownJobResponse
func (client *Client) SubmitConvertPdfToMarkdownJobWithContext(ctx context.Context, request *SubmitConvertPdfToMarkdownJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitConvertPdfToMarkdownJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileUrl) {
		query["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitConvertPdfToMarkdownJob"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitConvertPdfToMarkdownJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// pdf转word
//
// @param request - SubmitConvertPdfToWordJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitConvertPdfToWordJobResponse
func (client *Client) SubmitConvertPdfToWordJobWithContext(ctx context.Context, request *SubmitConvertPdfToWordJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitConvertPdfToWordJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileUrl) {
		query["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.ForceExportInnerImage) {
		query["ForceExportInnerImage"] = request.ForceExportInnerImage
	}

	if !dara.IsNil(request.FormulaEnhancement) {
		query["FormulaEnhancement"] = request.FormulaEnhancement
	}

	if !dara.IsNil(request.Option) {
		query["Option"] = request.Option
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitConvertPdfToWordJob"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitConvertPdfToWordJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 电子解析
//
// @param request - SubmitDigitalDocStructureJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitDigitalDocStructureJobResponse
func (client *Client) SubmitDigitalDocStructureJobWithContext(ctx context.Context, request *SubmitDigitalDocStructureJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitDigitalDocStructureJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileNameExtension) {
		query["FileNameExtension"] = request.FileNameExtension
	}

	if !dara.IsNil(request.FileUrl) {
		query["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.ImageStrategy) {
		query["ImageStrategy"] = request.ImageStrategy
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	if !dara.IsNil(request.RevealMarkdown) {
		query["RevealMarkdown"] = request.RevealMarkdown
	}

	if !dara.IsNil(request.UseUrlResponseBody) {
		query["UseUrlResponseBody"] = request.UseUrlResponseBody
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitDigitalDocStructureJob"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitDigitalDocStructureJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档智能解析流式输出
//
// @param request - SubmitDocParserJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitDocParserJobResponse
func (client *Client) SubmitDocParserJobWithContext(ctx context.Context, request *SubmitDocParserJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitDocParserJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileNameExtension) {
		query["FileNameExtension"] = request.FileNameExtension
	}

	if !dara.IsNil(request.FileUrl) {
		query["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.FormulaEnhancement) {
		query["FormulaEnhancement"] = request.FormulaEnhancement
	}

	if !dara.IsNil(request.LlmEnhancement) {
		query["LlmEnhancement"] = request.LlmEnhancement
	}

	if !dara.IsNil(request.Option) {
		query["Option"] = request.Option
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	if !dara.IsNil(request.OutputHtmlTable) {
		query["OutputHtmlTable"] = request.OutputHtmlTable
	}

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitDocParserJob"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitDocParserJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档智能解析
//
// @param request - SubmitDocStructureJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitDocStructureJobResponse
func (client *Client) SubmitDocStructureJobWithContext(ctx context.Context, request *SubmitDocStructureJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitDocStructureJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllowPptFormat) {
		query["AllowPptFormat"] = request.AllowPptFormat
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileNameExtension) {
		query["FileNameExtension"] = request.FileNameExtension
	}

	if !dara.IsNil(request.FileUrl) {
		query["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.FormulaEnhancement) {
		query["FormulaEnhancement"] = request.FormulaEnhancement
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.StructureType) {
		query["StructureType"] = request.StructureType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitDocStructureJob"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitDocStructureJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档抽取
//
// @param request - SubmitDocumentExtractJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitDocumentExtractJobResponse
func (client *Client) SubmitDocumentExtractJobWithContext(ctx context.Context, request *SubmitDocumentExtractJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitDocumentExtractJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileNameExtension) {
		query["FileNameExtension"] = request.FileNameExtension
	}

	if !dara.IsNil(request.FileUrl) {
		query["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitDocumentExtractJob"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitDocumentExtractJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 表格智能解析
//
// @param request - SubmitTableUnderstandingJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitTableUnderstandingJobResponse
func (client *Client) SubmitTableUnderstandingJobWithContext(ctx context.Context, request *SubmitTableUnderstandingJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitTableUnderstandingJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileNameExtension) {
		query["FileNameExtension"] = request.FileNameExtension
	}

	if !dara.IsNil(request.FileUrl) {
		query["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitTableUnderstandingJob"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitTableUnderstandingJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
