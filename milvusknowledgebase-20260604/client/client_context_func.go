// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Registers files that are uploaded to the knowledge base storage as knowledge base documents and **automatically triggers parsing*	- (chunking and embedding). Two import types are supported:
//
// - `LOCAL_UPLOAD`: Works with the `GetKnowledgeBasePreSignedUrl` direct upload flow. This operation only registers the file and does not verify whether the file is actually uploaded. Therefore, you must complete the PUT upload before calling this operation.
//
// - `OSS_IMPORT`: Imports files from an external OSS bucket. The operation creates an asynchronous import task and returns a `knowledge_import_task_id`. The system downloads and registers the files in the background.
//
// A maximum of 100 files can be registered in a single request.
//
// @param request - AddDocumentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDocumentsResponse
func (client *Client) AddDocumentsWithContext(ctx context.Context, datasetId *string, request *AddDocumentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddDocumentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Dedup) {
		body["Dedup"] = request.Dedup
	}

	if !dara.IsNil(request.Documents) {
		body["Documents"] = request.Documents
	}

	if !dara.IsNil(request.ImportType) {
		body["ImportType"] = request.ImportType
	}

	if !dara.IsNil(request.KnowledgeBaseId) {
		body["KnowledgeBaseId"] = request.KnowledgeBaseId
	}

	if !dara.IsNil(request.MetaFields) {
		body["MetaFields"] = request.MetaFields
	}

	if !dara.IsNil(request.StrategyId) {
		body["StrategyId"] = request.StrategyId
	}

	if !dara.IsNil(request.DingTalkConfiguration) {
		body["dingTalkConfiguration"] = request.DingTalkConfiguration
	}

	if !dara.IsNil(request.ParentId) {
		body["parentId"] = request.ParentId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDocuments"),
		Version:     dara.String("2026-06-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasets/" + dara.PercentEncode(dara.StringValue(datasetId)) + "/documents/addDocuments"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDocumentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates an **OSS pre-signed PUT URL*	- pointing to the knowledge base dedicated storage for each file in `Documents`. The caller uses the URL to upload file content directly to Object Storage Service (OSS), and then calls `AddDocuments` to register the files. A maximum of 100 files can be processed per request.
//
// @param request - GetKnowledgeBasePreSignedUrlRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetKnowledgeBasePreSignedUrlResponse
func (client *Client) GetKnowledgeBasePreSignedUrlWithContext(ctx context.Context, datasetId *string, request *GetKnowledgeBasePreSignedUrlRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetKnowledgeBasePreSignedUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Documents) {
		body["Documents"] = request.Documents
	}

	if !dara.IsNil(request.ExpiresIn) {
		body["ExpiresIn"] = request.ExpiresIn
	}

	if !dara.IsNil(request.KnowledgeBaseId) {
		body["KnowledgeBaseId"] = request.KnowledgeBaseId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetKnowledgeBasePreSignedUrl"),
		Version:     dara.String("2026-06-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasets/" + dara.PercentEncode(dara.StringValue(datasetId)) + "/getKnowledgeBasePreSignedUrl"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetKnowledgeBasePreSignedUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves documents from a knowledge base.
//
// Description:
//
// Retrieves documents from a specified knowledge base by question or image. Use DRAFT, LATEST_PUBLISHED, or vN display names for the version. Pass tag filter conditions using the actual backend operators.
//
// @param request - SearchKnowledgeBaseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchKnowledgeBaseResponse
func (client *Client) SearchKnowledgeBaseWithContext(ctx context.Context, knowledgeBaseId *string, request *SearchKnowledgeBaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SearchKnowledgeBaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocumentIds) {
		body["documentIds"] = request.DocumentIds
	}

	if !dara.IsNil(request.EnableKnowledgeGraph) {
		body["enableKnowledgeGraph"] = request.EnableKnowledgeGraph
	}

	if !dara.IsNil(request.Image) {
		body["image"] = request.Image
	}

	if !dara.IsNil(request.PageNumber) {
		body["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Query) {
		body["query"] = request.Query
	}

	if !dara.IsNil(request.RerankModelId) {
		body["rerankModelId"] = request.RerankModelId
	}

	if !dara.IsNil(request.RerankModelName) {
		body["rerankModelName"] = request.RerankModelName
	}

	if !dara.IsNil(request.RetrievalConfig) {
		body["retrievalConfig"] = request.RetrievalConfig
	}

	if !dara.IsNil(request.TagFilter) {
		body["tagFilter"] = request.TagFilter
	}

	if !dara.IsNil(request.Version) {
		body["version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchKnowledgeBase"),
		Version:     dara.String("2026-06-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/knowledge-bases/" + dara.PercentEncode(dara.StringValue(knowledgeBaseId)) + "/search"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchKnowledgeBaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
