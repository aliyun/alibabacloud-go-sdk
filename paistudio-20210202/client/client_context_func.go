// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 复制实验
//
// @param request - CopyExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyExperimentResponse
func (client *Client) CopyExperimentWithContext(ctx context.Context, ExperimentId *string, request *CopyExperimentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CopyExperimentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Accessibility) {
		body["Accessibility"] = request.Accessibility
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.FolderId) {
		body["FolderId"] = request.FolderId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Source) {
		body["Source"] = request.Source
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CopyExperiment"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments/" + dara.PercentEncode(dara.StringValue(ExperimentId)) + "/copy"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CopyExperimentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建实验，或根据实验模版创建实验
//
// @param request - CreateExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExperimentResponse
func (client *Client) CreateExperimentWithContext(ctx context.Context, request *CreateExperimentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateExperimentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Accessibility) {
		body["Accessibility"] = request.Accessibility
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.FolderId) {
		body["FolderId"] = request.FolderId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Options) {
		body["Options"] = request.Options
	}

	if !dara.IsNil(request.Source) {
		body["Source"] = request.Source
	}

	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateExperiment"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateExperimentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建算法文件夹
//
// @param request - CreateExperimentFolderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExperimentFolderResponse
func (client *Client) CreateExperimentFolderWithContext(ctx context.Context, request *CreateExperimentFolderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateExperimentFolderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Accessibility) {
		body["Accessibility"] = request.Accessibility
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ParentFolderId) {
		body["ParentFolderId"] = request.ParentFolderId
	}

	if !dara.IsNil(request.Source) {
		body["Source"] = request.Source
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateExperimentFolder"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experimentfolders"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateExperimentFolderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 校验实验是否能迁移
//
// @param request - CreateExperimentMigrateValidationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExperimentMigrateValidationResponse
func (client *Client) CreateExperimentMigrateValidationWithContext(ctx context.Context, request *CreateExperimentMigrateValidationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateExperimentMigrateValidationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SourceExpId) {
		query["SourceExpId"] = request.SourceExpId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateExperimentMigrateValidation"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/migrate/experimentvalidation"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateExperimentMigrateValidationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建一个工作流的作业
//
// @param request - CreateJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateJobResponse
func (client *Client) CreateJobWithContext(ctx context.Context, request *CreateJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ExecuteType) {
		body["ExecuteType"] = request.ExecuteType
	}

	if !dara.IsNil(request.ExperimentId) {
		body["ExperimentId"] = request.ExperimentId
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.Options) {
		body["Options"] = request.Options
	}

	if !dara.IsNil(request.PipelineDraftId) {
		body["PipelineDraftId"] = request.PipelineDraftId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateJob"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/jobs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除实验
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExperimentResponse
func (client *Client) DeleteExperimentWithContext(ctx context.Context, ExperimentId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteExperimentResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteExperiment"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments/" + dara.PercentEncode(dara.StringValue(ExperimentId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteExperimentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除算法文件夹
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExperimentFolderResponse
func (client *Client) DeleteExperimentFolderWithContext(ctx context.Context, FolderId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteExperimentFolderResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteExperimentFolder"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experimentfolders/" + dara.PercentEncode(dara.StringValue(FolderId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteExperimentFolderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取算法树
//
// @param request - GetAlgoTreeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAlgoTreeResponse
func (client *Client) GetAlgoTreeWithContext(ctx context.Context, request *GetAlgoTreeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAlgoTreeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAlgoTree"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/algo/tree"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAlgoTreeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取算法定义
//
// @param request - GetAlgorithmDefRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAlgorithmDefResponse
func (client *Client) GetAlgorithmDefWithContext(ctx context.Context, request *GetAlgorithmDefRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAlgorithmDefResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlgoVersion) {
		query["AlgoVersion"] = request.AlgoVersion
	}

	if !dara.IsNil(request.Identifier) {
		query["Identifier"] = request.Identifier
	}

	if !dara.IsNil(request.Provider) {
		query["Provider"] = request.Provider
	}

	if !dara.IsNil(request.Signature) {
		query["Signature"] = request.Signature
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAlgorithmDef"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/algorithm/def"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAlgorithmDefResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量获取算法定义
//
// @param request - GetAlgorithmDefsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAlgorithmDefsResponse
func (client *Client) GetAlgorithmDefsWithContext(ctx context.Context, request *GetAlgorithmDefsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAlgorithmDefsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LatestTimestamp) {
		query["LatestTimestamp"] = request.LatestTimestamp
	}

	if !dara.IsNil(request.RangeEnd) {
		query["RangeEnd"] = request.RangeEnd
	}

	if !dara.IsNil(request.RangeStart) {
		query["RangeStart"] = request.RangeStart
	}

	if !dara.IsNil(request.Timestamp) {
		query["Timestamp"] = request.Timestamp
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAlgorithmDefs"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/algorithm/defs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAlgorithmDefsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取算法树
//
// @param request - GetAlgorithmTreeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAlgorithmTreeResponse
func (client *Client) GetAlgorithmTreeWithContext(ctx context.Context, request *GetAlgorithmTreeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAlgorithmTreeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAlgorithmTree"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/algorithm/tree"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAlgorithmTreeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实验信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExperimentResponse
func (client *Client) GetExperimentWithContext(ctx context.Context, ExperimentId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetExperimentResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetExperiment"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments/" + dara.PercentEncode(dara.StringValue(ExperimentId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetExperimentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取算法文件夹下的内容
//
// @param request - GetExperimentFolderChildrenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExperimentFolderChildrenResponse
func (client *Client) GetExperimentFolderChildrenWithContext(ctx context.Context, FolderId *string, request *GetExperimentFolderChildrenRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetExperimentFolderChildrenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Accessibility) {
		query["Accessibility"] = request.Accessibility
	}

	if !dara.IsNil(request.OnlyFolder) {
		query["OnlyFolder"] = request.OnlyFolder
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetExperimentFolderChildren"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experimentfolders/" + dara.PercentEncode(dara.StringValue(FolderId)) + "/children"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetExperimentFolderChildrenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实验的元信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExperimentMetaResponse
func (client *Client) GetExperimentMetaWithContext(ctx context.Context, ExperimentId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetExperimentMetaResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetExperimentMeta"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments/" + dara.PercentEncode(dara.StringValue(ExperimentId)) + "/meta"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetExperimentMetaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实验以及实验节点的状态
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExperimentStatusResponse
func (client *Client) GetExperimentStatusWithContext(ctx context.Context, ExperimentId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetExperimentStatusResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetExperimentStatus"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments/" + dara.PercentEncode(dara.StringValue(ExperimentId)) + "/status"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetExperimentStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询实验的可视化meta
//
// @param request - GetExperimentVisualizationMetaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExperimentVisualizationMetaResponse
func (client *Client) GetExperimentVisualizationMetaWithContext(ctx context.Context, ExperimentId *string, request *GetExperimentVisualizationMetaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetExperimentVisualizationMetaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NodeIds) {
		query["NodeIds"] = request.NodeIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetExperimentVisualizationMeta"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments/" + dara.PercentEncode(dara.StringValue(ExperimentId)) + "/visualizationMeta"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetExperimentVisualizationMetaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实验的统计信息
//
// @param request - GetExperimentsStatisticsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExperimentsStatisticsResponse
func (client *Client) GetExperimentsStatisticsWithContext(ctx context.Context, request *GetExperimentsStatisticsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetExperimentsStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.WorkspaceIds) {
		query["WorkspaceIds"] = request.WorkspaceIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetExperimentsStatistics"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/statistics/experiments"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetExperimentsStatisticsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实验或文件夹所有者列表
//
// @param request - GetExperimentsUsersStatisticsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExperimentsUsersStatisticsResponse
func (client *Client) GetExperimentsUsersStatisticsWithContext(ctx context.Context, request *GetExperimentsUsersStatisticsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetExperimentsUsersStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetExperimentsUsersStatistics"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/statistics/experimentsusers"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetExperimentsUsersStatisticsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取一个PAI Studio作业详情
//
// @param request - GetJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobResponse
func (client *Client) GetJobWithContext(ctx context.Context, JobId *string, request *GetJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Verbose) {
		query["Verbose"] = request.Verbose
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetJob"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/jobs/" + dara.PercentEncode(dara.StringValue(JobId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取MaxCompute表schema
//
// @param request - GetMCTableSchemaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMCTableSchemaResponse
func (client *Client) GetMCTableSchemaWithContext(ctx context.Context, TableName *string, request *GetMCTableSchemaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMCTableSchemaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMCTableSchema"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasources/maxcompute/tables/" + dara.PercentEncode(dara.StringValue(TableName)) + "/schema"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMCTableSchemaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实验节点输入桩的输入表的格式
//
// @param request - GetNodeInputSchemaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNodeInputSchemaResponse
func (client *Client) GetNodeInputSchemaWithContext(ctx context.Context, ExperimentId *string, NodeId *string, request *GetNodeInputSchemaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetNodeInputSchemaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InputId) {
		query["InputId"] = request.InputId
	}

	if !dara.IsNil(request.InputIndex) {
		query["InputIndex"] = request.InputIndex
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNodeInputSchema"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments/" + dara.PercentEncode(dara.StringValue(ExperimentId)) + "/nodes/" + dara.PercentEncode(dara.StringValue(NodeId)) + "/schema"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNodeInputSchemaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取某个节点的输出模型信息
//
// @param request - GetNodeOutputRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNodeOutputResponse
func (client *Client) GetNodeOutputWithContext(ctx context.Context, ExperimentId *string, NodeId *string, OutputId *string, request *GetNodeOutputRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetNodeOutputResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OutputIndex) {
		query["OutputIndex"] = request.OutputIndex
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNodeOutput"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments/" + dara.PercentEncode(dara.StringValue(ExperimentId)) + "/nodes/" + dara.PercentEncode(dara.StringValue(NodeId)) + "/outputs/" + dara.PercentEncode(dara.StringValue(OutputId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNodeOutputResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取PAI Studio中指定模板
//
// @param request - GetTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTemplateResponse
func (client *Client) GetTemplateWithContext(ctx context.Context, TemplateId *string, request *GetTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Verbose) {
		query["Verbose"] = request.Verbose
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTemplate"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/templates/" + dara.PercentEncode(dara.StringValue(TemplateId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取授权角色列表
//
// @param request - ListAuthRolesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuthRolesResponse
func (client *Client) ListAuthRolesWithContext(ctx context.Context, request *ListAuthRolesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAuthRolesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsGenerateToken) {
		query["IsGenerateToken"] = request.IsGenerateToken
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAuthRoles"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/authorization/roles"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAuthRolesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实验列表
//
// @param request - ListExperimentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExperimentsResponse
func (client *Client) ListExperimentsWithContext(ctx context.Context, request *ListExperimentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListExperimentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Creator) {
		query["Creator"] = request.Creator
	}

	if !dara.IsNil(request.ExperimentId) {
		query["ExperimentId"] = request.ExperimentId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListExperiments"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListExperimentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举标签
//
// @param request - ListImageLabelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListImageLabelsResponse
func (client *Client) ListImageLabelsWithContext(ctx context.Context, request *ListImageLabelsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListImageLabelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.LabelFilter) {
		query["LabelFilter"] = request.LabelFilter
	}

	if !dara.IsNil(request.LabelKeys) {
		query["LabelKeys"] = request.LabelKeys
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListImageLabels"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/image/labels"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListImageLabelsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举已注册镜像
//
// @param request - ListImagesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListImagesResponse
func (client *Client) ListImagesWithContext(ctx context.Context, request *ListImagesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListImagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Labels) {
		query["Labels"] = request.Labels
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Verbose) {
		query["Verbose"] = request.Verbose
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListImages"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/images"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListImagesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取作业详情
//
// @param request - ListJobsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobsResponse
func (client *Client) ListJobsWithContext(ctx context.Context, request *ListJobsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Creator) {
		query["Creator"] = request.Creator
	}

	if !dara.IsNil(request.ExperimentId) {
		query["ExperimentId"] = request.ExperimentId
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListJobs"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/jobs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取某个节点的输出模型列表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodeOutputsResponse
func (client *Client) ListNodeOutputsWithContext(ctx context.Context, ExperimentId *string, NodeId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListNodeOutputsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNodeOutputs"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments/" + dara.PercentEncode(dara.StringValue(ExperimentId)) + "/nodes/" + dara.PercentEncode(dara.StringValue(NodeId)) + "/outputs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNodeOutputsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取最近的实验
//
// @param request - ListRecentExperimentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecentExperimentsResponse
func (client *Client) ListRecentExperimentsWithContext(ctx context.Context, request *ListRecentExperimentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRecentExperimentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRecentExperiments"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/recentexperiments"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRecentExperimentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取PAI Studio中指定模板列表
//
// @param request - ListTemplatesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTemplatesResponse
func (client *Client) ListTemplatesWithContext(ctx context.Context, request *ListTemplatesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Label) {
		query["Label"] = request.Label
	}

	if !dara.IsNil(request.List) {
		query["List"] = request.List
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.TagId) {
		query["TagId"] = request.TagId
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	if !dara.IsNil(request.TypeId) {
		query["TypeId"] = request.TypeId
	}

	if !dara.IsNil(request.Verbose) {
		query["Verbose"] = request.Verbose
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTemplates"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/templates"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTemplatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 迁移PAI Studio 1.0的实验目录
//
// @param request - MigrateExperimentFoldersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MigrateExperimentFoldersResponse
func (client *Client) MigrateExperimentFoldersWithContext(ctx context.Context, request *MigrateExperimentFoldersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *MigrateExperimentFoldersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Accessibility) {
		query["Accessibility"] = request.Accessibility
	}

	if !dara.IsNil(request.IsOwner) {
		query["IsOwner"] = request.IsOwner
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MigrateExperimentFolders"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/migrate/folders"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &MigrateExperimentFoldersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 迁移PAI Studio 1.0的实验
//
// @param request - MigrateExperimentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MigrateExperimentsResponse
func (client *Client) MigrateExperimentsWithContext(ctx context.Context, request *MigrateExperimentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *MigrateExperimentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Accessibility) {
		query["Accessibility"] = request.Accessibility
	}

	if !dara.IsNil(request.DestFolderId) {
		query["DestFolderId"] = request.DestFolderId
	}

	if !dara.IsNil(request.IsOwner) {
		query["IsOwner"] = request.IsOwner
	}

	if !dara.IsNil(request.SourceExpId) {
		query["SourceExpId"] = request.SourceExpId
	}

	if !dara.IsNil(request.UpdateIfExists) {
		query["UpdateIfExists"] = request.UpdateIfExists
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MigrateExperiments"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/migrate/experiments"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &MigrateExperimentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预览Maxcompute表数据
//
// @param request - PreviewMCTableRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PreviewMCTableResponse
func (client *Client) PreviewMCTableWithContext(ctx context.Context, TableName *string, request *PreviewMCTableRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PreviewMCTableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Endpoint) {
		query["Endpoint"] = request.Endpoint
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.Partition) {
		query["Partition"] = request.Partition
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PreviewMCTable"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasources/maxcompute/tables/" + dara.PercentEncode(dara.StringValue(TableName)) + "/preview"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PreviewMCTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发布实验
//
// @param request - PublishExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishExperimentResponse
func (client *Client) PublishExperimentWithContext(ctx context.Context, ExperimentId *string, request *PublishExperimentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PublishExperimentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FolderId) {
		body["FolderId"] = request.FolderId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishExperiment"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments/" + dara.PercentEncode(dara.StringValue(ExperimentId)) + "/publish"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishExperimentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询实验的可视化数据
//
// @param request - QueryExperimentVisualizationDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryExperimentVisualizationDataResponse
func (client *Client) QueryExperimentVisualizationDataWithContext(ctx context.Context, ExperimentId *string, request *QueryExperimentVisualizationDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryExperimentVisualizationDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ToArray(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryExperimentVisualizationData"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments/" + dara.PercentEncode(dara.StringValue(ExperimentId)) + "/visualizationDataQuery"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryExperimentVisualizationDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 搜索MaxCompute表
//
// @param request - SearchMCTablesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchMCTablesResponse
func (client *Client) SearchMCTablesWithContext(ctx context.Context, request *SearchMCTablesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SearchMCTablesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchMCTables"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasources/maxcompute/tables"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchMCTablesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止实验所有运行中的作业
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopExperimentResponse
func (client *Client) StopExperimentWithContext(ctx context.Context, ExperimentId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopExperimentResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopExperiment"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments/" + dara.PercentEncode(dara.StringValue(ExperimentId)) + "/stop"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopExperimentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止一个实验的作业
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopJobResponse
func (client *Client) StopJobWithContext(ctx context.Context, JobId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopJobResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopJob"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/jobs/" + dara.PercentEncode(dara.StringValue(JobId)) + "/stop"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新实验内容
//
// @param request - UpdateExperimentContentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateExperimentContentResponse
func (client *Client) UpdateExperimentContentWithContext(ctx context.Context, ExperimentId *string, request *UpdateExperimentContentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateExperimentContentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.Version) {
		body["Version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateExperimentContent"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments/" + dara.PercentEncode(dara.StringValue(ExperimentId)) + "/content"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateExperimentContentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新算法文件夹
//
// @param request - UpdateExperimentFolderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateExperimentFolderResponse
func (client *Client) UpdateExperimentFolderWithContext(ctx context.Context, FolderId *string, request *UpdateExperimentFolderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateExperimentFolderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ParentFolderId) {
		body["ParentFolderId"] = request.ParentFolderId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateExperimentFolder"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experimentfolders/" + dara.PercentEncode(dara.StringValue(FolderId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateExperimentFolderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新实验的Meta信息
//
// @param request - UpdateExperimentMetaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateExperimentMetaResponse
func (client *Client) UpdateExperimentMetaWithContext(ctx context.Context, ExperimentId *string, request *UpdateExperimentMetaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateExperimentMetaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Accessibility) {
		body["Accessibility"] = request.Accessibility
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.FolderId) {
		body["FolderId"] = request.FolderId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Options) {
		body["Options"] = request.Options
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateExperimentMeta"),
		Version:     dara.String("2021-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments/" + dara.PercentEncode(dara.StringValue(ExperimentId)) + "/meta"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateExperimentMetaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
