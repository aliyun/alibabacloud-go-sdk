// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Uploads a file directly to the Bucket/ObjectKey specified in the response, and then uses the object URL as OssFileUrl to create a parsing task.
//
// @param request - AuthorizeFileUploadRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeFileUploadResponse
func (client *Client) AuthorizeFileUploadWithContext(ctx context.Context, request *AuthorizeFileUploadRequest, runtime *dara.RuntimeOptions) (_result *AuthorizeFileUploadResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentName) {
		query["AgentName"] = request.AgentName
	}

	if !dara.IsNil(request.BatchSize) {
		query["BatchSize"] = request.BatchSize
	}

	if !dara.IsNil(request.FileFormat) {
		query["FileFormat"] = request.FileFormat
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuthorizeFileUpload"),
		Version:     dara.String("2026-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AuthorizeFileUploadResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a document parsing task.
//
// Description:
//
// - Region: Only China (Beijing) is supported.
//
// - Fees: The service is free of charge during the public preview period.
//
// @param request - CreateDocParserJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDocParserJobResponse
func (client *Client) CreateDocParserJobWithContext(ctx context.Context, request *CreateDocParserJobRequest, runtime *dara.RuntimeOptions) (_result *CreateDocParserJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentName) {
		query["AgentName"] = request.AgentName
	}

	if !dara.IsNil(request.AsrLanguage) {
		query["AsrLanguage"] = request.AsrLanguage
	}

	if !dara.IsNil(request.AudioClipOutput) {
		query["AudioClipOutput"] = request.AudioClipOutput
	}

	if !dara.IsNil(request.AudioWindowSeconds) {
		query["AudioWindowSeconds"] = request.AudioWindowSeconds
	}

	if !dara.IsNil(request.ChunkSummary) {
		query["ChunkSummary"] = request.ChunkSummary
	}

	if !dara.IsNil(request.FileFormat) {
		query["FileFormat"] = request.FileFormat
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileUrl) {
		query["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.FrameOutput) {
		query["FrameOutput"] = request.FrameOutput
	}

	if !dara.IsNil(request.GlobalSummary) {
		query["GlobalSummary"] = request.GlobalSummary
	}

	if !dara.IsNil(request.ImageMode) {
		query["ImageMode"] = request.ImageMode
	}

	if !dara.IsNil(request.ImageUnderstanding) {
		query["ImageUnderstanding"] = request.ImageUnderstanding
	}

	if !dara.IsNil(request.MediaChunkIntervalSeconds) {
		query["MediaChunkIntervalSeconds"] = request.MediaChunkIntervalSeconds
	}

	if !dara.IsNil(request.MediaChunkStrategy) {
		query["MediaChunkStrategy"] = request.MediaChunkStrategy
	}

	if !dara.IsNil(request.MediaFramesPerMinute) {
		query["MediaFramesPerMinute"] = request.MediaFramesPerMinute
	}

	if !dara.IsNil(request.MediaMaxFrameBudget) {
		query["MediaMaxFrameBudget"] = request.MediaMaxFrameBudget
	}

	if !dara.IsNil(request.MediaMinFrameBudget) {
		query["MediaMinFrameBudget"] = request.MediaMinFrameBudget
	}

	if !dara.IsNil(request.OssFileUrl) {
		query["OssFileUrl"] = request.OssFileUrl
	}

	if !dara.IsNil(request.OutputFormat) {
		query["OutputFormat"] = request.OutputFormat
	}

	if !dara.IsNil(request.ParseScene) {
		query["ParseScene"] = request.ParseScene
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResponseMode) {
		query["ResponseMode"] = request.ResponseMode
	}

	if !dara.IsNil(request.ResultType) {
		query["ResultType"] = request.ResultType
	}

	if !dara.IsNil(request.TableFormat) {
		query["TableFormat"] = request.TableFormat
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDocParserJob"),
		Version:     dara.String("2026-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDocParserJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the result of a document parsing task.
//
// Description:
//
// - Region: Only China (Beijing) is supported.
//
// - Fees: Free of charge during the public preview period.
//
// - Call DescribeDocParserJobResult to retrieve the parsing result of a document parsing task. Call this operation only after DescribeDocParserJobStatus returns a Status of success. Task results are retained for 72 hours and cannot be retrieved after expiration.
//
// @param request - DescribeDocParserJobResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDocParserJobResultResponse
func (client *Client) DescribeDocParserJobResultWithContext(ctx context.Context, request *DescribeDocParserJobResultRequest, runtime *dara.RuntimeOptions) (_result *DescribeDocParserJobResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentName) {
		query["AgentName"] = request.AgentName
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResultType) {
		query["ResultType"] = request.ResultType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDocParserJobResult"),
		Version:     dara.String("2026-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDocParserJobResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of a document parsing task.
//
// Description:
//
// - Region: Only China (Beijing) is supported.
//
// - Fees: The service is free of charge during the public preview period.
//
// @param request - DescribeDocParserJobStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDocParserJobStatusResponse
func (client *Client) DescribeDocParserJobStatusWithContext(ctx context.Context, request *DescribeDocParserJobStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeDocParserJobStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentName) {
		query["AgentName"] = request.AgentName
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDocParserJobStatus"),
		Version:     dara.String("2026-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDocParserJobStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the content of a web page.
//
// Description:
//
// - Region: Only China (Beijing) and Singapore regions are supported.
//
// - Fees: Free of charge during the public preview period.
//
// @param request - WebFetchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WebFetchResponse
func (client *Client) WebFetchWithContext(ctx context.Context, request *WebFetchRequest, runtime *dara.RuntimeOptions) (_result *WebFetchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentName) {
		query["AgentName"] = request.AgentName
	}

	if !dara.IsNil(request.OutputFormat) {
		query["OutputFormat"] = request.OutputFormat
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Url) {
		query["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("WebFetch"),
		Version:     dara.String("2026-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &WebFetchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a web search.
//
// Description:
//
// - Region: Only China (Beijing) and Singapore regions are supported.
//
// - Fees: Free during the public preview period. No fees are charged.
//
// @param request - WebSearchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WebSearchResponse
func (client *Client) WebSearchWithContext(ctx context.Context, request *WebSearchRequest, runtime *dara.RuntimeOptions) (_result *WebSearchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentName) {
		query["AgentName"] = request.AgentName
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UrlScopeDomains) {
		query["UrlScopeDomains"] = request.UrlScopeDomains
	}

	if !dara.IsNil(request.UrlScopeMode) {
		query["UrlScopeMode"] = request.UrlScopeMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("WebSearch"),
		Version:     dara.String("2026-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &WebSearchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
