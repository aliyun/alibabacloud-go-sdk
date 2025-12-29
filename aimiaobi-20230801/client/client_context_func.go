// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 添加审核自定义词库记录
//
// @param tmpReq - AddAuditTermsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddAuditTermsResponse
func (client *Client) AddAuditTermsWithContext(ctx context.Context, tmpReq *AddAuditTermsRequest, runtime *dara.RuntimeOptions) (_result *AddAuditTermsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddAuditTermsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ExceptionWord) {
		request.ExceptionWordShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExceptionWord, dara.String("ExceptionWord"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ExceptionWordShrink) {
		body["ExceptionWord"] = request.ExceptionWordShrink
	}

	if !dara.IsNil(request.Keyword) {
		body["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.SuggestWord) {
		body["SuggestWord"] = request.SuggestWord
	}

	if !dara.IsNil(request.TermsDesc) {
		body["TermsDesc"] = request.TermsDesc
	}

	if !dara.IsNil(request.TermsName) {
		body["TermsName"] = request.TermsName
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddAuditTerms"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddAuditTermsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加文档到数据集
//
// @param tmpReq - AddDatasetDocumentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDatasetDocumentResponse
func (client *Client) AddDatasetDocumentWithContext(ctx context.Context, tmpReq *AddDatasetDocumentRequest, runtime *dara.RuntimeOptions) (_result *AddDatasetDocumentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddDatasetDocumentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Document) {
		request.DocumentShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Document, dara.String("Document"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DatasetId) {
		body["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.DatasetName) {
		body["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.DocumentShrink) {
		body["Document"] = request.DocumentShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDatasetDocument"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDatasetDocumentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成剪辑视频任务
//
// @param tmpReq - AsyncCreateClipsTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AsyncCreateClipsTaskResponse
func (client *Client) AsyncCreateClipsTaskWithContext(ctx context.Context, tmpReq *AsyncCreateClipsTaskRequest, runtime *dara.RuntimeOptions) (_result *AsyncCreateClipsTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AsyncCreateClipsTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ColorWords) {
		request.ColorWordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ColorWords, dara.String("ColorWords"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Stickers) {
		request.StickersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Stickers, dara.String("Stickers"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CloseMusic) {
		body["CloseMusic"] = request.CloseMusic
	}

	if !dara.IsNil(request.CloseSubtitle) {
		body["CloseSubtitle"] = request.CloseSubtitle
	}

	if !dara.IsNil(request.CloseVoice) {
		body["CloseVoice"] = request.CloseVoice
	}

	if !dara.IsNil(request.ColorWordsShrink) {
		body["ColorWords"] = request.ColorWordsShrink
	}

	if !dara.IsNil(request.CustomVoiceUrl) {
		body["CustomVoiceUrl"] = request.CustomVoiceUrl
	}

	if !dara.IsNil(request.CustomVoiceVolume) {
		body["CustomVoiceVolume"] = request.CustomVoiceVolume
	}

	if !dara.IsNil(request.Height) {
		body["Height"] = request.Height
	}

	if !dara.IsNil(request.MusicUrl) {
		body["MusicUrl"] = request.MusicUrl
	}

	if !dara.IsNil(request.MusicVolume) {
		body["MusicVolume"] = request.MusicVolume
	}

	if !dara.IsNil(request.StickersShrink) {
		body["Stickers"] = request.StickersShrink
	}

	if !dara.IsNil(request.SubtitleFontSize) {
		body["SubtitleFontSize"] = request.SubtitleFontSize
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.VoiceStyle) {
		body["VoiceStyle"] = request.VoiceStyle
	}

	if !dara.IsNil(request.VoiceVolume) {
		body["VoiceVolume"] = request.VoiceVolume
	}

	if !dara.IsNil(request.Width) {
		body["Width"] = request.Width
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AsyncCreateClipsTask"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AsyncCreateClipsTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 智能剪辑timeline
//
// @param request - AsyncCreateClipsTimeLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AsyncCreateClipsTimeLineResponse
func (client *Client) AsyncCreateClipsTimeLineWithContext(ctx context.Context, request *AsyncCreateClipsTimeLineRequest, runtime *dara.RuntimeOptions) (_result *AsyncCreateClipsTimeLineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AdditionalContent) {
		body["AdditionalContent"] = request.AdditionalContent
	}

	if !dara.IsNil(request.CustomContent) {
		body["CustomContent"] = request.CustomContent
	}

	if !dara.IsNil(request.NoRefVideo) {
		body["NoRefVideo"] = request.NoRefVideo
	}

	if !dara.IsNil(request.ProcessPrompt) {
		body["ProcessPrompt"] = request.ProcessPrompt
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AsyncCreateClipsTimeLine"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AsyncCreateClipsTimeLineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑剪辑任务的timeline
//
// @param tmpReq - AsyncEditTimelineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AsyncEditTimelineResponse
func (client *Client) AsyncEditTimelineWithContext(ctx context.Context, tmpReq *AsyncEditTimelineRequest, runtime *dara.RuntimeOptions) (_result *AsyncEditTimelineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AsyncEditTimelineShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Timelines) {
		request.TimelinesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Timelines, dara.String("Timelines"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoClips) {
		body["AutoClips"] = request.AutoClips
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TimelinesShrink) {
		body["Timelines"] = request.TimelinesShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AsyncEditTimeline"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AsyncEditTimelineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上传招标书文件
//
// @param request - AsyncUploadTenderDocRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AsyncUploadTenderDocResponse
func (client *Client) AsyncUploadTenderDocWithContext(ctx context.Context, request *AsyncUploadTenderDocRequest, runtime *dara.RuntimeOptions) (_result *AsyncUploadTenderDocResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileKey) {
		body["FileKey"] = request.FileKey
	}

	if !dara.IsNil(request.TenderDocName) {
		body["TenderDocName"] = request.TenderDocName
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AsyncUploadTenderDoc"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AsyncUploadTenderDocResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上传剪辑素材
//
// @param tmpReq - AsyncUploadVideoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AsyncUploadVideoResponse
func (client *Client) AsyncUploadVideoWithContext(ctx context.Context, tmpReq *AsyncUploadVideoRequest, runtime *dara.RuntimeOptions) (_result *AsyncUploadVideoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AsyncUploadVideoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ReferenceVideo) {
		request.ReferenceVideoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReferenceVideo, dara.String("ReferenceVideo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SourceVideos) {
		request.SourceVideosShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceVideos, dara.String("SourceVideos"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.VideoRoles) {
		request.VideoRolesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VideoRoles, dara.String("VideoRoles"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AnlysisPrompt) {
		body["AnlysisPrompt"] = request.AnlysisPrompt
	}

	if !dara.IsNil(request.FaceIdentitySimilarityMinScore) {
		body["FaceIdentitySimilarityMinScore"] = request.FaceIdentitySimilarityMinScore
	}

	if !dara.IsNil(request.ReferenceVideoShrink) {
		body["ReferenceVideo"] = request.ReferenceVideoShrink
	}

	if !dara.IsNil(request.RemoveSubtitle) {
		body["RemoveSubtitle"] = request.RemoveSubtitle
	}

	if !dara.IsNil(request.SourceVideosShrink) {
		body["SourceVideos"] = request.SourceVideosShrink
	}

	if !dara.IsNil(request.SplitInterval) {
		body["SplitInterval"] = request.SplitInterval
	}

	if !dara.IsNil(request.VideoRolesShrink) {
		body["VideoRoles"] = request.VideoRolesShrink
	}

	if !dara.IsNil(request.VideoShotFaceIdentityCount) {
		body["VideoShotFaceIdentityCount"] = request.VideoShotFaceIdentityCount
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AsyncUploadVideo"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AsyncUploadVideoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 标书写作接口
//
// @param request - AsyncWritingBiddingDocRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AsyncWritingBiddingDocResponse
func (client *Client) AsyncWritingBiddingDocWithContext(ctx context.Context, request *AsyncWritingBiddingDocRequest, runtime *dara.RuntimeOptions) (_result *AsyncWritingBiddingDocResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CompanyKeyword) {
		body["CompanyKeyword"] = request.CompanyKeyword
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AsyncWritingBiddingDoc"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AsyncWritingBiddingDocResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 绑定PPT作品信息
//
// @param request - BindPptArtifactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindPptArtifactResponse
func (client *Client) BindPptArtifactWithContext(ctx context.Context, request *BindPptArtifactRequest, runtime *dara.RuntimeOptions) (_result *BindPptArtifactResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ArtifactId) {
		body["ArtifactId"] = request.ArtifactId
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindPptArtifact"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindPptArtifactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消异步任务
//
// @param request - CancelAsyncTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelAsyncTaskResponse
func (client *Client) CancelAsyncTaskWithContext(ctx context.Context, request *CancelAsyncTaskRequest, runtime *dara.RuntimeOptions) (_result *CancelAsyncTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelAsyncTask"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelAsyncTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消审核任务
//
// @param request - CancelAuditTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelAuditTaskResponse
func (client *Client) CancelAuditTaskWithContext(ctx context.Context, request *CancelAuditTaskRequest, runtime *dara.RuntimeOptions) (_result *CancelAuditTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ArticleId) {
		body["ArticleId"] = request.ArticleId
	}

	if !dara.IsNil(request.ContentAuditTaskId) {
		body["ContentAuditTaskId"] = request.ContentAuditTaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelAuditTask"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelAuditTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消深度写作任务
//
// @param request - CancelDeepWriteTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelDeepWriteTaskResponse
func (client *Client) CancelDeepWriteTaskWithContext(ctx context.Context, request *CancelDeepWriteTaskRequest, runtime *dara.RuntimeOptions) (_result *CancelDeepWriteTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelDeepWriteTask"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelDeepWriteTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 清除所有干预内容
//
// @param request - ClearIntervenesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClearIntervenesResponse
func (client *Client) ClearIntervenesWithContext(ctx context.Context, request *ClearIntervenesRequest, runtime *dara.RuntimeOptions) (_result *ClearIntervenesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClearIntervenes"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClearIntervenesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 是否将本次提交自定义规则库得到的解析结果用于审核任务。由于解析结果可能不满足用户需求，因此我们为您提供了该接口用于二次确认。如果对提交的规则库解析满意，则可以直接将本次提交任务的 TaskId 作为入参，系统会对您上传的规则库做后处理，使它可以被用于审核。反之，您可以重新调用 SubmitAuditNote 接口上传修改之后的规则库。
//
// @param request - ConfirmAndPostProcessAuditNoteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfirmAndPostProcessAuditNoteResponse
func (client *Client) ConfirmAndPostProcessAuditNoteWithContext(ctx context.Context, request *ConfirmAndPostProcessAuditNoteRequest, runtime *dara.RuntimeOptions) (_result *ConfirmAndPostProcessAuditNoteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfirmAndPostProcessAuditNote"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfirmAndPostProcessAuditNoteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 数据集管理-创建
//
// @param tmpReq - CreateDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatasetResponse
func (client *Client) CreateDatasetWithContext(ctx context.Context, tmpReq *CreateDatasetRequest, runtime *dara.RuntimeOptions) (_result *CreateDatasetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDatasetShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DatasetConfig) {
		request.DatasetConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DatasetConfig, dara.String("DatasetConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DocumentHandleConfig) {
		request.DocumentHandleConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DocumentHandleConfig, dara.String("DocumentHandleConfig"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DatasetConfigShrink) {
		body["DatasetConfig"] = request.DatasetConfigShrink
	}

	if !dara.IsNil(request.DatasetDescription) {
		body["DatasetDescription"] = request.DatasetDescription
	}

	if !dara.IsNil(request.DatasetName) {
		body["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.DatasetType) {
		body["DatasetType"] = request.DatasetType
	}

	if !dara.IsNil(request.DocumentHandleConfigShrink) {
		body["DocumentHandleConfig"] = request.DocumentHandleConfigShrink
	}

	if !dara.IsNil(request.InvokeType) {
		body["InvokeType"] = request.InvokeType
	}

	if !dara.IsNil(request.SearchDatasetEnable) {
		body["SearchDatasetEnable"] = request.SearchDatasetEnable
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataset"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDatasetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通用配置-创建
//
// @param request - CreateGeneralConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGeneralConfigResponse
func (client *Client) CreateGeneralConfigWithContext(ctx context.Context, request *CreateGeneralConfigRequest, runtime *dara.RuntimeOptions) (_result *CreateGeneralConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigKey) {
		body["ConfigKey"] = request.ConfigKey
	}

	if !dara.IsNil(request.ConfigValue) {
		body["ConfigValue"] = request.ConfigValue
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGeneralConfig"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGeneralConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档管理-创建
//
// @param tmpReq - CreateGeneratedContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGeneratedContentResponse
func (client *Client) CreateGeneratedContentWithContext(ctx context.Context, tmpReq *CreateGeneratedContentRequest, runtime *dara.RuntimeOptions) (_result *CreateGeneratedContentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateGeneratedContentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Keywords) {
		request.KeywordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Keywords, dara.String("Keywords"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.ContentDomain) {
		body["ContentDomain"] = request.ContentDomain
	}

	if !dara.IsNil(request.ContentText) {
		body["ContentText"] = request.ContentText
	}

	if !dara.IsNil(request.KeywordsShrink) {
		body["Keywords"] = request.KeywordsShrink
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	if !dara.IsNil(request.Uuid) {
		body["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGeneratedContent"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGeneratedContentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取授权token
//
// @param request - CreateTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTokenResponse
func (client *Client) CreateTokenWithContext(ctx context.Context, request *CreateTokenRequest, runtime *dara.RuntimeOptions) (_result *CreateTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateToken"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除用户账户下所有可供审核使用的自定义规则库。删除后无法找回，如果您有对规则库存档的需求，请预先使用 DownloadAuditNote 接口保存需要的规则库。
//
// @param request - DeleteAuditNoteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAuditNoteResponse
func (client *Client) DeleteAuditNoteWithContext(ctx context.Context, request *DeleteAuditNoteRequest, runtime *dara.RuntimeOptions) (_result *DeleteAuditNoteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.NoteId) {
		body["NoteId"] = request.NoteId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAuditNote"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAuditNoteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定的词库记录
//
// @param tmpReq - DeleteAuditTermsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAuditTermsResponse
func (client *Client) DeleteAuditTermsWithContext(ctx context.Context, tmpReq *DeleteAuditTermsRequest, runtime *dara.RuntimeOptions) (_result *DeleteAuditTermsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteAuditTermsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.IdList) {
		request.IdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IdList, dara.String("IdList"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.IdListShrink) {
		body["IdList"] = request.IdListShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAuditTerms"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAuditTermsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除自定义文本
//
// @param request - DeleteCustomTextRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomTextResponse
func (client *Client) DeleteCustomTextWithContext(ctx context.Context, request *DeleteCustomTextRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomTextResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CommodityCode) {
		body["CommodityCode"] = request.CommodityCode
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomText"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomTextResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据主题删除自定义主题事件
//
// @param request - DeleteCustomTopicByTopicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomTopicByTopicResponse
func (client *Client) DeleteCustomTopicByTopicWithContext(ctx context.Context, request *DeleteCustomTopicByTopicRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomTopicByTopicResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Topic) {
		body["Topic"] = request.Topic
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomTopicByTopic"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomTopicByTopicResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据自定义观点ID删除自定义观点
//
// @param request - DeleteCustomTopicViewPointByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomTopicViewPointByIdResponse
func (client *Client) DeleteCustomTopicViewPointByIdWithContext(ctx context.Context, request *DeleteCustomTopicViewPointByIdRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomTopicViewPointByIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CustomViewPointId) {
		body["CustomViewPointId"] = request.CustomViewPointId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomTopicViewPointById"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomTopicViewPointByIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 数据集管理-删除
//
// @param request - DeleteDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDatasetResponse
func (client *Client) DeleteDatasetWithContext(ctx context.Context, request *DeleteDatasetRequest, runtime *dara.RuntimeOptions) (_result *DeleteDatasetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DatasetId) {
		body["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataset"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDatasetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除数据集文档
//
// @param request - DeleteDatasetDocumentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDatasetDocumentResponse
func (client *Client) DeleteDatasetDocumentWithContext(ctx context.Context, request *DeleteDatasetDocumentRequest, runtime *dara.RuntimeOptions) (_result *DeleteDatasetDocumentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DatasetId) {
		body["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.DatasetName) {
		body["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.DocId) {
		body["DocId"] = request.DocId
	}

	if !dara.IsNil(request.DocUuid) {
		body["DocUuid"] = request.DocUuid
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDatasetDocument"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDatasetDocumentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 妙读删除多个文档
//
// @param tmpReq - DeleteDocsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDocsResponse
func (client *Client) DeleteDocsWithContext(ctx context.Context, tmpReq *DeleteDocsRequest, runtime *dara.RuntimeOptions) (_result *DeleteDocsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteDocsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DocIds) {
		request.DocIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DocIds, dara.String("DocIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DocIdsShrink) {
		body["DocIds"] = request.DocIdsShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDocs"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDocsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定的用于事实性审核的 URL。
//
// @param request - DeleteFactAuditUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFactAuditUrlResponse
func (client *Client) DeleteFactAuditUrlWithContext(ctx context.Context, request *DeleteFactAuditUrlRequest, runtime *dara.RuntimeOptions) (_result *DeleteFactAuditUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Url) {
		body["Url"] = request.Url
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFactAuditUrl"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFactAuditUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通用配置-删除
//
// @param request - DeleteGeneralConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGeneralConfigResponse
func (client *Client) DeleteGeneralConfigWithContext(ctx context.Context, request *DeleteGeneralConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteGeneralConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigKey) {
		body["ConfigKey"] = request.ConfigKey
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGeneralConfig"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGeneralConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档管理-删除。
//
// @param request - DeleteGeneratedContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGeneratedContentResponse
func (client *Client) DeleteGeneratedContentWithContext(ctx context.Context, request *DeleteGeneratedContentRequest, runtime *dara.RuntimeOptions) (_result *DeleteGeneratedContentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGeneratedContent"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGeneratedContentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除干预规则
//
// @param request - DeleteInterveneRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInterveneRuleResponse
func (client *Client) DeleteInterveneRuleWithContext(ctx context.Context, request *DeleteInterveneRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteInterveneRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RuleId) {
		body["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteInterveneRule"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteInterveneRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据ID删除素材
//
// @param request - DeleteMaterialByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMaterialByIdResponse
func (client *Client) DeleteMaterialByIdWithContext(ctx context.Context, request *DeleteMaterialByIdRequest, runtime *dara.RuntimeOptions) (_result *DeleteMaterialByIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMaterialById"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMaterialByIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定自定义文体
//
// @param request - DeleteStyleLearningResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStyleLearningResultResponse
func (client *Client) DeleteStyleLearningResultWithContext(ctx context.Context, request *DeleteStyleLearningResultRequest, runtime *dara.RuntimeOptions) (_result *DeleteStyleLearningResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteStyleLearningResult"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteStyleLearningResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 从链接中提取文档内容
//
// @param tmpReq - DocumentExtractionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DocumentExtractionResponse
func (client *Client) DocumentExtractionWithContext(ctx context.Context, tmpReq *DocumentExtractionRequest, runtime *dara.RuntimeOptions) (_result *DocumentExtractionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DocumentExtractionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Urls) {
		request.UrlsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Urls, dara.String("Urls"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UrlsShrink) {
		body["Urls"] = request.UrlsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DocumentExtraction"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DocumentExtractionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 您可以通过调用该接口下载结构化后的规则库，供您进行进一步处理。该接口同时拥有两个功能：下载未后处理的结构化规则库，或下载当前可用于审核的结构化规则库。具体使用方法，请参考入参说明。
//
// @param request - DownloadAuditNoteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DownloadAuditNoteResponse
func (client *Client) DownloadAuditNoteWithContext(ctx context.Context, request *DownloadAuditNoteRequest, runtime *dara.RuntimeOptions) (_result *DownloadAuditNoteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.NoteId) {
		body["NoteId"] = request.NoteId
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DownloadAuditNote"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DownloadAuditNoteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 标书下载接口
//
// @param request - DownloadBiddingDocRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DownloadBiddingDocResponse
func (client *Client) DownloadBiddingDocWithContext(ctx context.Context, request *DownloadBiddingDocRequest, runtime *dara.RuntimeOptions) (_result *DownloadBiddingDocResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DownloadBiddingDoc"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DownloadBiddingDocResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑审核自定义词库记录
//
// @param tmpReq - EditAuditTermsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditAuditTermsResponse
func (client *Client) EditAuditTermsWithContext(ctx context.Context, tmpReq *EditAuditTermsRequest, runtime *dara.RuntimeOptions) (_result *EditAuditTermsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &EditAuditTermsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ExceptionWord) {
		request.ExceptionWordShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExceptionWord, dara.String("ExceptionWord"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ExceptionWordShrink) {
		body["ExceptionWord"] = request.ExceptionWordShrink
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Keyword) {
		body["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.SuggestWord) {
		body["SuggestWord"] = request.SuggestWord
	}

	if !dara.IsNil(request.TermsDesc) {
		body["TermsDesc"] = request.TermsDesc
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EditAuditTerms"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EditAuditTermsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑标书内容接口
//
// @param request - EditBiddingDocRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditBiddingDocResponse
func (client *Client) EditBiddingDocWithContext(ctx context.Context, request *EditBiddingDocRequest, runtime *dara.RuntimeOptions) (_result *EditBiddingDocResponse, _err error) {
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

	if !dara.IsNil(request.ContentFormat) {
		body["ContentFormat"] = request.ContentFormat
	}

	if !dara.IsNil(request.ContentType) {
		body["ContentType"] = request.ContentType
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EditBiddingDoc"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EditBiddingDocResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出企业VOC分析任务明细列表
//
// @param tmpReq - ExportAnalysisTagDetailByTaskIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportAnalysisTagDetailByTaskIdResponse
func (client *Client) ExportAnalysisTagDetailByTaskIdWithContext(ctx context.Context, tmpReq *ExportAnalysisTagDetailByTaskIdRequest, runtime *dara.RuntimeOptions) (_result *ExportAnalysisTagDetailByTaskIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ExportAnalysisTagDetailByTaskIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Categories) {
		request.CategoriesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Categories, dara.String("Categories"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoriesShrink) {
		body["Categories"] = request.CategoriesShrink
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportAnalysisTagDetailByTaskId"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportAnalysisTagDetailByTaskIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出智能审核报告
//
// @param request - ExportAuditContentResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportAuditContentResultResponse
func (client *Client) ExportAuditContentResultWithContext(ctx context.Context, request *ExportAuditContentResultRequest, runtime *dara.RuntimeOptions) (_result *ExportAuditContentResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportAuditContentResult"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportAuditContentResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出-自定义数据源-选题视角分析任务结果
//
// @param request - ExportCustomSourceAnalysisTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportCustomSourceAnalysisTaskResponse
func (client *Client) ExportCustomSourceAnalysisTaskWithContext(ctx context.Context, request *ExportCustomSourceAnalysisTaskRequest, runtime *dara.RuntimeOptions) (_result *ExportCustomSourceAnalysisTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportCustomSourceAnalysisTask"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportCustomSourceAnalysisTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档管理-导出。
//
// @param request - ExportGeneratedContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportGeneratedContentResponse
func (client *Client) ExportGeneratedContentWithContext(ctx context.Context, request *ExportGeneratedContentRequest, runtime *dara.RuntimeOptions) (_result *ExportGeneratedContentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportGeneratedContent"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportGeneratedContentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出选题策划文档，响应为一个可公开访问的URL。一小时后失效
//
// @param tmpReq - ExportHotTopicPlanningProposalsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportHotTopicPlanningProposalsResponse
func (client *Client) ExportHotTopicPlanningProposalsWithContext(ctx context.Context, tmpReq *ExportHotTopicPlanningProposalsRequest, runtime *dara.RuntimeOptions) (_result *ExportHotTopicPlanningProposalsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ExportHotTopicPlanningProposalsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CustomViewPointIds) {
		request.CustomViewPointIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomViewPointIds, dara.String("CustomViewPointIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Titles) {
		request.TitlesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Titles, dara.String("Titles"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CustomViewPointIdsShrink) {
		body["CustomViewPointIds"] = request.CustomViewPointIdsShrink
	}

	if !dara.IsNil(request.ExportType) {
		body["ExportType"] = request.ExportType
	}

	if !dara.IsNil(request.TitlesShrink) {
		body["Titles"] = request.TitlesShrink
	}

	if !dara.IsNil(request.Topic) {
		body["Topic"] = request.Topic
	}

	if !dara.IsNil(request.TopicSource) {
		body["TopicSource"] = request.TopicSource
	}

	if !dara.IsNil(request.ViewPointType) {
		body["ViewPointType"] = request.ViewPointType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportHotTopicPlanningProposals"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportHotTopicPlanningProposalsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出所有干预内容
//
// @param request - ExportIntervenesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportIntervenesResponse
func (client *Client) ExportIntervenesWithContext(ctx context.Context, request *ExportIntervenesRequest, runtime *dara.RuntimeOptions) (_result *ExportIntervenesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportIntervenes"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportIntervenesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 反馈某次生成的结果
//
// @param tmpReq - FeedbackDialogueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FeedbackDialogueResponse
func (client *Client) FeedbackDialogueWithContext(ctx context.Context, tmpReq *FeedbackDialogueRequest, runtime *dara.RuntimeOptions) (_result *FeedbackDialogueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &FeedbackDialogueShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RatingTags) {
		request.RatingTagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RatingTags, dara.String("RatingTags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CustomerResponse) {
		body["CustomerResponse"] = request.CustomerResponse
	}

	if !dara.IsNil(request.GoodText) {
		body["GoodText"] = request.GoodText
	}

	if !dara.IsNil(request.ModifiedResponse) {
		body["ModifiedResponse"] = request.ModifiedResponse
	}

	if !dara.IsNil(request.Rating) {
		body["Rating"] = request.Rating
	}

	if !dara.IsNil(request.RatingTagsShrink) {
		body["RatingTags"] = request.RatingTagsShrink
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FeedbackDialogue"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FeedbackDialogueResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取词库导出任务结果
//
// @param request - FetchExportTermsTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FetchExportTermsTaskResponse
func (client *Client) FetchExportTermsTaskWithContext(ctx context.Context, request *FetchExportTermsTaskRequest, runtime *dara.RuntimeOptions) (_result *FetchExportTermsTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FetchExportTermsTask"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FetchExportTermsTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取异步导出文档任务结果
//
// @param request - FetchExportWordTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FetchExportWordTaskResponse
func (client *Client) FetchExportWordTaskWithContext(ctx context.Context, request *FetchExportWordTaskRequest, runtime *dara.RuntimeOptions) (_result *FetchExportWordTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FetchExportWordTask"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FetchExportWordTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取图片任务执行结果
//
// @param tmpReq - FetchImageTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FetchImageTaskResponse
func (client *Client) FetchImageTaskWithContext(ctx context.Context, tmpReq *FetchImageTaskRequest, runtime *dara.RuntimeOptions) (_result *FetchImageTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &FetchImageTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TaskIdList) {
		request.TaskIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskIdList, dara.String("TaskIdList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ArticleTaskId) {
		body["ArticleTaskId"] = request.ArticleTaskId
	}

	if !dara.IsNil(request.TaskIdListShrink) {
		body["TaskIdList"] = request.TaskIdListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FetchImageTask"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FetchImageTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取导入词库任务结果
//
// @param request - FetchImportTermsTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FetchImportTermsTaskResponse
func (client *Client) FetchImportTermsTaskWithContext(ctx context.Context, request *FetchImportTermsTaskRequest, runtime *dara.RuntimeOptions) (_result *FetchImportTermsTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FetchImportTermsTask"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FetchImportTermsTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成内容导出文档任务
//
// @param request - GenerateExportWordTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateExportWordTaskResponse
func (client *Client) GenerateExportWordTaskWithContext(ctx context.Context, request *GenerateExportWordTaskRequest, runtime *dara.RuntimeOptions) (_result *GenerateExportWordTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.GeneratedContentId) {
		body["GeneratedContentId"] = request.GeneratedContentId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateExportWordTask"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateExportWordTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成临时可访问的公开url
//
// @param request - GenerateFileUrlByKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateFileUrlByKeyResponse
func (client *Client) GenerateFileUrlByKeyWithContext(ctx context.Context, request *GenerateFileUrlByKeyRequest, runtime *dara.RuntimeOptions) (_result *GenerateFileUrlByKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FileKey) {
		body["FileKey"] = request.FileKey
	}

	if !dara.IsNil(request.FileName) {
		body["FileName"] = request.FileName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateFileUrlByKey"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateFileUrlByKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 智能配图，图片生成任务
//
// @param tmpReq - GenerateImageTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateImageTaskResponse
func (client *Client) GenerateImageTaskWithContext(ctx context.Context, tmpReq *GenerateImageTaskRequest, runtime *dara.RuntimeOptions) (_result *GenerateImageTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GenerateImageTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ParagraphList) {
		request.ParagraphListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ParagraphList, dara.String("ParagraphList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ArticleTaskId) {
		body["ArticleTaskId"] = request.ArticleTaskId
	}

	if !dara.IsNil(request.ParagraphListShrink) {
		body["ParagraphList"] = request.ParagraphListShrink
	}

	if !dara.IsNil(request.Size) {
		body["Size"] = request.Size
	}

	if !dara.IsNil(request.Style) {
		body["Style"] = request.Style
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateImageTask"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateImageTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成上传配置
//
// @param request - GenerateUploadConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateUploadConfigResponse
func (client *Client) GenerateUploadConfigWithContext(ctx context.Context, request *GenerateUploadConfigRequest, runtime *dara.RuntimeOptions) (_result *GenerateUploadConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		body["FileName"] = request.FileName
	}

	if !dara.IsNil(request.ParentDir) {
		body["ParentDir"] = request.ParentDir
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateUploadConfig"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateUploadConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视角生成
//
// @param tmpReq - GenerateViewPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateViewPointResponse
func (client *Client) GenerateViewPointWithContext(ctx context.Context, tmpReq *GenerateViewPointRequest, runtime *dara.RuntimeOptions) (_result *GenerateViewPointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GenerateViewPointShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ReferenceData) {
		request.ReferenceDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReferenceData, dara.String("ReferenceData"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ReferenceDataShrink) {
		body["ReferenceData"] = request.ReferenceDataShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateViewPoint"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateViewPointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询规则库后处理的进度。与 ConfirmAndPostProcessAuditNote 接口配合使用，供您查询当前后处理任务的状态。
//
// @param request - GetAuditNotePostProcessingStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAuditNotePostProcessingStatusResponse
func (client *Client) GetAuditNotePostProcessingStatusWithContext(ctx context.Context, request *GetAuditNotePostProcessingStatusRequest, runtime *dara.RuntimeOptions) (_result *GetAuditNotePostProcessingStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAuditNotePostProcessingStatus"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAuditNotePostProcessingStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户上传规则库的处理状态。通过该接口，用户可以查询到当前规则库上传任务的状态，并获取到解析后的规则库文件大小、存储路径等信息。
//
// @param request - GetAuditNoteProcessingStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAuditNoteProcessingStatusResponse
func (client *Client) GetAuditNoteProcessingStatusWithContext(ctx context.Context, request *GetAuditNoteProcessingStatusRequest, runtime *dara.RuntimeOptions) (_result *GetAuditNoteProcessingStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAuditNoteProcessingStatus"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAuditNoteProcessingStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获得剪辑任务状态
//
// @param request - GetAutoClipsTaskInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAutoClipsTaskInfoResponse
func (client *Client) GetAutoClipsTaskInfoWithContext(ctx context.Context, request *GetAutoClipsTaskInfoRequest, runtime *dara.RuntimeOptions) (_result *GetAutoClipsTaskInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAutoClipsTaskInfo"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAutoClipsTaskInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户当前可供审核的规则库信息，只能查询到当前可用于审核的规则库。如果您想看到自定义规则库的具体内容，请使用 DownloadAuditNote 接口。
//
// @param request - GetAvailableAuditNotesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAvailableAuditNotesResponse
func (client *Client) GetAvailableAuditNotesWithContext(ctx context.Context, request *GetAvailableAuditNotesRequest, runtime *dara.RuntimeOptions) (_result *GetAvailableAuditNotesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.NoteId) {
		body["NoteId"] = request.NoteId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAvailableAuditNotes"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAvailableAuditNotesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获得标书写作结果接口
//
// @param request - GetBiddingDocInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBiddingDocInfoResponse
func (client *Client) GetBiddingDocInfoWithContext(ctx context.Context, request *GetBiddingDocInfoRequest, runtime *dara.RuntimeOptions) (_result *GetBiddingDocInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBiddingDocInfo"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBiddingDocInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获得标书功能剩余额度
//
// @param request - GetBiddingRemainLimitNumRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBiddingRemainLimitNumResponse
func (client *Client) GetBiddingRemainLimitNumWithContext(ctx context.Context, request *GetBiddingRemainLimitNumRequest, runtime *dara.RuntimeOptions) (_result *GetBiddingRemainLimitNumResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiName) {
		body["ApiName"] = request.ApiName
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBiddingRemainLimitNum"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBiddingRemainLimitNumResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取某次标签挖掘结果分类
//
// @param request - GetCategoriesByTaskIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCategoriesByTaskIdResponse
func (client *Client) GetCategoriesByTaskIdWithContext(ctx context.Context, request *GetCategoriesByTaskIdRequest, runtime *dara.RuntimeOptions) (_result *GetCategoriesByTaskIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCategoriesByTaskId"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCategoriesByTaskIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取自定义播报单任务结果
//
// @param request - GetCustomHotTopicBroadcastJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomHotTopicBroadcastJobResponse
func (client *Client) GetCustomHotTopicBroadcastJobWithContext(ctx context.Context, request *GetCustomHotTopicBroadcastJobRequest, runtime *dara.RuntimeOptions) (_result *GetCustomHotTopicBroadcastJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCustomHotTopicBroadcastJob"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCustomHotTopicBroadcastJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取自定义数据源-选题视角分析任务结果
//
// @param request - GetCustomSourceTopicAnalysisTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomSourceTopicAnalysisTaskResponse
func (client *Client) GetCustomSourceTopicAnalysisTaskWithContext(ctx context.Context, request *GetCustomSourceTopicAnalysisTaskRequest, runtime *dara.RuntimeOptions) (_result *GetCustomSourceTopicAnalysisTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCustomSourceTopicAnalysisTask"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCustomSourceTopicAnalysisTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取自定义文本
//
// @param request - GetCustomTextRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomTextResponse
func (client *Client) GetCustomTextWithContext(ctx context.Context, request *GetCustomTextRequest, runtime *dara.RuntimeOptions) (_result *GetCustomTextResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CommodityCode) {
		body["CommodityCode"] = request.CommodityCode
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCustomText"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCustomTextResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取自定义选题视角分析任务结果
//
// @param request - GetCustomTopicSelectionPerspectiveAnalysisTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomTopicSelectionPerspectiveAnalysisTaskResponse
func (client *Client) GetCustomTopicSelectionPerspectiveAnalysisTaskWithContext(ctx context.Context, request *GetCustomTopicSelectionPerspectiveAnalysisTaskRequest, runtime *dara.RuntimeOptions) (_result *GetCustomTopicSelectionPerspectiveAnalysisTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCustomTopicSelectionPerspectiveAnalysisTask"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCustomTopicSelectionPerspectiveAnalysisTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取系统数据源配置和个人配置
//
// @param request - GetDataSourceOrderConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataSourceOrderConfigResponse
func (client *Client) GetDataSourceOrderConfigWithContext(ctx context.Context, request *GetDataSourceOrderConfigRequest, runtime *dara.RuntimeOptions) (_result *GetDataSourceOrderConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.GenerateTechnology) {
		body["GenerateTechnology"] = request.GenerateTechnology
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataSourceOrderConfig"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataSourceOrderConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 数据集管理-详情
//
// @param request - GetDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatasetResponse
func (client *Client) GetDatasetWithContext(ctx context.Context, request *GetDatasetRequest, runtime *dara.RuntimeOptions) (_result *GetDatasetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DatasetId) {
		body["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.DatasetName) {
		body["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataset"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDatasetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据集文档
//
// @param tmpReq - GetDatasetDocumentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatasetDocumentResponse
func (client *Client) GetDatasetDocumentWithContext(ctx context.Context, tmpReq *GetDatasetDocumentRequest, runtime *dara.RuntimeOptions) (_result *GetDatasetDocumentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetDatasetDocumentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.IncludeFields) {
		request.IncludeFieldsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IncludeFields, dara.String("IncludeFields"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DatasetId) {
		body["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.DatasetName) {
		body["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.DocId) {
		body["DocId"] = request.DocId
	}

	if !dara.IsNil(request.DocUuid) {
		body["DocUuid"] = request.DocUuid
	}

	if !dara.IsNil(request.IncludeFieldsShrink) {
		body["IncludeFields"] = request.IncludeFieldsShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDatasetDocument"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDatasetDocumentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询深度写作任务
//
// @param request - GetDeepWriteTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeepWriteTaskResponse
func (client *Client) GetDeepWriteTaskWithContext(ctx context.Context, request *GetDeepWriteTaskRequest, runtime *dara.RuntimeOptions) (_result *GetDeepWriteTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeepWriteTask"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeepWriteTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询深度写作任务的结果
//
// @param request - GetDeepWriteTaskResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeepWriteTaskResultResponse
func (client *Client) GetDeepWriteTaskResultWithContext(ctx context.Context, request *GetDeepWriteTaskResultRequest, runtime *dara.RuntimeOptions) (_result *GetDeepWriteTaskResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeepWriteTaskResult"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeepWriteTaskResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文档聚合任务结果
//
// @param request - GetDocClusterTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocClusterTaskResponse
func (client *Client) GetDocClusterTaskWithContext(ctx context.Context, request *GetDocClusterTaskRequest, runtime *dara.RuntimeOptions) (_result *GetDocClusterTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDocClusterTask"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDocClusterTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 妙读获取文档信息
//
// @param request - GetDocInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocInfoResponse
func (client *Client) GetDocInfoWithContext(ctx context.Context, request *GetDocInfoRequest, runtime *dara.RuntimeOptions) (_result *GetDocInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		body["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.DocId) {
		body["DocId"] = request.DocId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDocInfo"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDocInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业VOC分析任务结果
//
// @param request - GetEnterpriseVocAnalysisTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEnterpriseVocAnalysisTaskResponse
func (client *Client) GetEnterpriseVocAnalysisTaskWithContext(ctx context.Context, request *GetEnterpriseVocAnalysisTaskRequest, runtime *dara.RuntimeOptions) (_result *GetEnterpriseVocAnalysisTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEnterpriseVocAnalysisTask"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEnterpriseVocAnalysisTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取当前正用于事实性审核的信源 URL。
//
// @param request - GetFactAuditUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFactAuditUrlResponse
func (client *Client) GetFactAuditUrlWithContext(ctx context.Context, request *GetFactAuditUrlRequest, runtime *dara.RuntimeOptions) (_result *GetFactAuditUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFactAuditUrl"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFactAuditUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 妙读获得文档字数
//
// @param request - GetFileContentLengthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFileContentLengthResponse
func (client *Client) GetFileContentLengthWithContext(ctx context.Context, request *GetFileContentLengthRequest, runtime *dara.RuntimeOptions) (_result *GetFileContentLengthResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocName) {
		body["DocName"] = request.DocName
	}

	if !dara.IsNil(request.FileUrl) {
		body["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFileContentLength"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFileContentLengthResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通用配置-查询
//
// @param request - GetGeneralConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGeneralConfigResponse
func (client *Client) GetGeneralConfigWithContext(ctx context.Context, request *GetGeneralConfigRequest, runtime *dara.RuntimeOptions) (_result *GetGeneralConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigKey) {
		body["ConfigKey"] = request.ConfigKey
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGeneralConfig"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGeneralConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档管理-查询详情。
//
// @param request - GetGeneratedContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGeneratedContentResponse
func (client *Client) GetGeneratedContentWithContext(ctx context.Context, request *GetGeneratedContentRequest, runtime *dara.RuntimeOptions) (_result *GetGeneratedContentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGeneratedContent"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGeneratedContentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询新闻播报单
//
// @param tmpReq - GetHotTopicBroadcastRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotTopicBroadcastResponse
func (client *Client) GetHotTopicBroadcastWithContext(ctx context.Context, tmpReq *GetHotTopicBroadcastRequest, runtime *dara.RuntimeOptions) (_result *GetHotTopicBroadcastResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetHotTopicBroadcastShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Locations) {
		request.LocationsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Locations, dara.String("Locations"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.StepForCustomSummaryStyleConfig) {
		request.StepForCustomSummaryStyleConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StepForCustomSummaryStyleConfig, dara.String("StepForCustomSummaryStyleConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.StepForNewsBroadcastContentConfig) {
		request.StepForNewsBroadcastContentConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StepForNewsBroadcastContentConfig, dara.String("StepForNewsBroadcastContentConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Topics) {
		request.TopicsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Topics, dara.String("Topics"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CalcTotalToken) {
		body["CalcTotalToken"] = request.CalcTotalToken
	}

	if !dara.IsNil(request.Category) {
		body["Category"] = request.Category
	}

	if !dara.IsNil(request.Current) {
		body["Current"] = request.Current
	}

	if !dara.IsNil(request.HotTopicVersion) {
		body["HotTopicVersion"] = request.HotTopicVersion
	}

	if !dara.IsNil(request.LocationQuery) {
		body["LocationQuery"] = request.LocationQuery
	}

	if !dara.IsNil(request.LocationsShrink) {
		body["Locations"] = request.LocationsShrink
	}

	if !dara.IsNil(request.Query) {
		body["Query"] = request.Query
	}

	if !dara.IsNil(request.Size) {
		body["Size"] = request.Size
	}

	if !dara.IsNil(request.StepForCustomSummaryStyleConfigShrink) {
		body["StepForCustomSummaryStyleConfig"] = request.StepForCustomSummaryStyleConfigShrink
	}

	if !dara.IsNil(request.StepForNewsBroadcastContentConfigShrink) {
		body["StepForNewsBroadcastContentConfig"] = request.StepForNewsBroadcastContentConfigShrink
	}

	if !dara.IsNil(request.TopicsShrink) {
		body["Topics"] = request.TopicsShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHotTopicBroadcast"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHotTopicBroadcastResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获得干预全局回复
//
// @param request - GetInterveneGlobalReplyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInterveneGlobalReplyResponse
func (client *Client) GetInterveneGlobalReplyWithContext(ctx context.Context, request *GetInterveneGlobalReplyRequest, runtime *dara.RuntimeOptions) (_result *GetInterveneGlobalReplyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInterveneGlobalReply"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInterveneGlobalReplyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获得导入任务信息
//
// @param request - GetInterveneImportTaskInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInterveneImportTaskInfoResponse
func (client *Client) GetInterveneImportTaskInfoWithContext(ctx context.Context, request *GetInterveneImportTaskInfoRequest, runtime *dara.RuntimeOptions) (_result *GetInterveneImportTaskInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInterveneImportTaskInfo"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInterveneImportTaskInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获得干预项规则详情
//
// @param request - GetInterveneRuleDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInterveneRuleDetailResponse
func (client *Client) GetInterveneRuleDetailWithContext(ctx context.Context, request *GetInterveneRuleDetailRequest, runtime *dara.RuntimeOptions) (_result *GetInterveneRuleDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RuleId) {
		body["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInterveneRuleDetail"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInterveneRuleDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获得干预导入模版文件下载地址
//
// @param request - GetInterveneTemplateFileUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInterveneTemplateFileUrlResponse
func (client *Client) GetInterveneTemplateFileUrlWithContext(ctx context.Context, request *GetInterveneTemplateFileUrlRequest, runtime *dara.RuntimeOptions) (_result *GetInterveneTemplateFileUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInterveneTemplateFileUrl"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInterveneTemplateFileUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据ID获取素材内容
//
// @param request - GetMaterialByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMaterialByIdResponse
func (client *Client) GetMaterialByIdWithContext(ctx context.Context, request *GetMaterialByIdRequest, runtime *dara.RuntimeOptions) (_result *GetMaterialByIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMaterialById"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMaterialByIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取PPT组件的配置
//
// @param request - GetPptConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPptConfigResponse
func (client *Client) GetPptConfigWithContext(ctx context.Context, request *GetPptConfigRequest, runtime *dara.RuntimeOptions) (_result *GetPptConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPptConfig"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPptConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取当前用户的配置
//
// @param request - GetPropertiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPropertiesResponse
func (client *Client) GetPropertiesWithContext(ctx context.Context, request *GetPropertiesRequest, runtime *dara.RuntimeOptions) (_result *GetPropertiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProperties"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPropertiesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询智能审核结果
//
// @param request - GetSmartAuditResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSmartAuditResultResponse
func (client *Client) GetSmartAuditResultWithContext(ctx context.Context, request *GetSmartAuditResultRequest, runtime *dara.RuntimeOptions) (_result *GetSmartAuditResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSmartAuditResult"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSmartAuditResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询一键成片剪辑任务
//
// @param request - GetSmartClipTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSmartClipTaskResponse
func (client *Client) GetSmartClipTaskWithContext(ctx context.Context, request *GetSmartClipTaskRequest, runtime *dara.RuntimeOptions) (_result *GetSmartClipTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSmartClipTask"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSmartClipTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文体学习分析结果
//
// @param request - GetStyleLearningResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStyleLearningResultResponse
func (client *Client) GetStyleLearningResultWithContext(ctx context.Context, request *GetStyleLearningResultRequest, runtime *dara.RuntimeOptions) (_result *GetStyleLearningResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStyleLearningResult"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStyleLearningResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据ID获取热点事件信息
//
// @param request - GetTopicByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTopicByIdResponse
func (client *Client) GetTopicByIdWithContext(ctx context.Context, request *GetTopicByIdRequest, runtime *dara.RuntimeOptions) (_result *GetTopicByIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTopicById"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTopicByIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取选题视角分析任务结果
//
// @param request - GetTopicSelectionPerspectiveAnalysisTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTopicSelectionPerspectiveAnalysisTaskResponse
func (client *Client) GetTopicSelectionPerspectiveAnalysisTaskWithContext(ctx context.Context, request *GetTopicSelectionPerspectiveAnalysisTaskRequest, runtime *dara.RuntimeOptions) (_result *GetTopicSelectionPerspectiveAnalysisTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTopicSelectionPerspectiveAnalysisTask"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTopicSelectionPerspectiveAnalysisTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导入干预文件
//
// @param request - ImportInterveneFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportInterveneFileResponse
func (client *Client) ImportInterveneFileWithContext(ctx context.Context, request *ImportInterveneFileRequest, runtime *dara.RuntimeOptions) (_result *ImportInterveneFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DocName) {
		body["DocName"] = request.DocName
	}

	if !dara.IsNil(request.FileKey) {
		body["FileKey"] = request.FileKey
	}

	if !dara.IsNil(request.FileUrl) {
		body["FileUrl"] = request.FileUrl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportInterveneFile"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportInterveneFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 异步导入干预文件
//
// @param request - ImportInterveneFileAsyncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportInterveneFileAsyncResponse
func (client *Client) ImportInterveneFileAsyncWithContext(ctx context.Context, request *ImportInterveneFileAsyncRequest, runtime *dara.RuntimeOptions) (_result *ImportInterveneFileAsyncResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DocName) {
		body["DocName"] = request.DocName
	}

	if !dara.IsNil(request.FileKey) {
		body["FileKey"] = request.FileKey
	}

	if !dara.IsNil(request.FileUrl) {
		body["FileUrl"] = request.FileUrl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportInterveneFileAsync"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportInterveneFileAsyncResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 初始化PPT创建操作
//
// @param request - InitiatePptCreationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitiatePptCreationResponse
func (client *Client) InitiatePptCreationWithContext(ctx context.Context, request *InitiatePptCreationRequest, runtime *dara.RuntimeOptions) (_result *InitiatePptCreationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Outline) {
		body["Outline"] = request.Outline
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InitiatePptCreation"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InitiatePptCreationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置干预全局回复
//
// @param tmpReq - InsertInterveneGlobalReplyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsertInterveneGlobalReplyResponse
func (client *Client) InsertInterveneGlobalReplyWithContext(ctx context.Context, tmpReq *InsertInterveneGlobalReplyRequest, runtime *dara.RuntimeOptions) (_result *InsertInterveneGlobalReplyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &InsertInterveneGlobalReplyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ReplyMessagList) {
		request.ReplyMessagListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReplyMessagList, dara.String("ReplyMessagList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ReplyMessagListShrink) {
		body["ReplyMessagList"] = request.ReplyMessagListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InsertInterveneGlobalReply"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InsertInterveneGlobalReplyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 插入干预规则
//
// @param tmpReq - InsertInterveneRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsertInterveneRuleResponse
func (client *Client) InsertInterveneRuleWithContext(ctx context.Context, tmpReq *InsertInterveneRuleRequest, runtime *dara.RuntimeOptions) (_result *InsertInterveneRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &InsertInterveneRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InterveneRuleConfig) {
		request.InterveneRuleConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InterveneRuleConfig, dara.String("InterveneRuleConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InterveneRuleConfigShrink) {
		body["InterveneRuleConfig"] = request.InterveneRuleConfigShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InsertInterveneRule"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InsertInterveneRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页获取企业VOC分析任务明细列表
//
// @param tmpReq - ListAnalysisTagDetailByTaskIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAnalysisTagDetailByTaskIdResponse
func (client *Client) ListAnalysisTagDetailByTaskIdWithContext(ctx context.Context, tmpReq *ListAnalysisTagDetailByTaskIdRequest, runtime *dara.RuntimeOptions) (_result *ListAnalysisTagDetailByTaskIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListAnalysisTagDetailByTaskIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Categories) {
		request.CategoriesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Categories, dara.String("Categories"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoriesShrink) {
		body["Categories"] = request.CategoriesShrink
	}

	if !dara.IsNil(request.Current) {
		body["Current"] = request.Current
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Size) {
		body["Size"] = request.Size
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAnalysisTagDetailByTaskId"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAnalysisTagDetailByTaskIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询任务列表
//
// @param tmpReq - ListAsyncTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAsyncTasksResponse
func (client *Client) ListAsyncTasksWithContext(ctx context.Context, tmpReq *ListAsyncTasksRequest, runtime *dara.RuntimeOptions) (_result *ListAsyncTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListAsyncTasksShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TaskStatusList) {
		request.TaskStatusListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskStatusList, dara.String("TaskStatusList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TaskTypeList) {
		request.TaskTypeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskTypeList, dara.String("TaskTypeList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateTimeEnd) {
		body["CreateTimeEnd"] = request.CreateTimeEnd
	}

	if !dara.IsNil(request.CreateTimeStart) {
		body["CreateTimeStart"] = request.CreateTimeStart
	}

	if !dara.IsNil(request.Current) {
		body["Current"] = request.Current
	}

	if !dara.IsNil(request.Size) {
		body["Size"] = request.Size
	}

	if !dara.IsNil(request.TaskCode) {
		body["TaskCode"] = request.TaskCode
	}

	if !dara.IsNil(request.TaskName) {
		body["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TaskStatus) {
		body["TaskStatus"] = request.TaskStatus
	}

	if !dara.IsNil(request.TaskStatusListShrink) {
		body["TaskStatusList"] = request.TaskStatusListShrink
	}

	if !dara.IsNil(request.TaskType) {
		body["TaskType"] = request.TaskType
	}

	if !dara.IsNil(request.TaskTypeListShrink) {
		body["TaskTypeList"] = request.TaskTypeListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAsyncTasks"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAsyncTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取审核维度列表
//
// @param request - ListAuditContentErrorTypesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuditContentErrorTypesResponse
func (client *Client) ListAuditContentErrorTypesWithContext(ctx context.Context, request *ListAuditContentErrorTypesRequest, runtime *dara.RuntimeOptions) (_result *ListAuditContentErrorTypesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAuditContentErrorTypes"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAuditContentErrorTypesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取词库列表
//
// @param request - ListAuditTermsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuditTermsResponse
func (client *Client) ListAuditTermsWithContext(ctx context.Context, request *ListAuditTermsRequest, runtime *dara.RuntimeOptions) (_result *ListAuditTermsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TermsName) {
		body["TermsName"] = request.TermsName
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAuditTerms"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAuditTermsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获得标书写作任务列表
//
// @param request - ListBiddingDocRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBiddingDocResponse
func (client *Client) ListBiddingDocWithContext(ctx context.Context, request *ListBiddingDocRequest, runtime *dara.RuntimeOptions) (_result *ListBiddingDocResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateTimeEnd) {
		body["CreateTimeEnd"] = request.CreateTimeEnd
	}

	if !dara.IsNil(request.CreateTimeStart) {
		body["CreateTimeStart"] = request.CreateTimeStart
	}

	if !dara.IsNil(request.Current) {
		body["Current"] = request.Current
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Size) {
		body["Size"] = request.Size
	}

	if !dara.IsNil(request.Skip) {
		body["Skip"] = request.Skip
	}

	if !dara.IsNil(request.TaskName) {
		body["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TaskStatus) {
		body["TaskStatus"] = request.TaskStatus
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBiddingDoc"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBiddingDocResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取系统自定义预设
//
// @param request - ListBuildConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBuildConfigsResponse
func (client *Client) ListBuildConfigsWithContext(ctx context.Context, request *ListBuildConfigsRequest, runtime *dara.RuntimeOptions) (_result *ListBuildConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBuildConfigs"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBuildConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 自定义文本列表
//
// @param request - ListCustomTextRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomTextResponse
func (client *Client) ListCustomTextWithContext(ctx context.Context, request *ListCustomTextRequest, runtime *dara.RuntimeOptions) (_result *ListCustomTextResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CommodityCode) {
		body["CommodityCode"] = request.CommodityCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCustomText"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCustomTextResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 自定义视角列表
//
// @param tmpReq - ListCustomViewPointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomViewPointsResponse
func (client *Client) ListCustomViewPointsWithContext(ctx context.Context, tmpReq *ListCustomViewPointsRequest, runtime *dara.RuntimeOptions) (_result *ListCustomViewPointsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListCustomViewPointsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Attitudes) {
		request.AttitudesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Attitudes, dara.String("Attitudes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CustomViewPointIds) {
		request.CustomViewPointIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomViewPointIds, dara.String("CustomViewPointIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Attitude) {
		body["Attitude"] = request.Attitude
	}

	if !dara.IsNil(request.AttitudesShrink) {
		body["Attitudes"] = request.AttitudesShrink
	}

	if !dara.IsNil(request.CustomViewPointId) {
		body["CustomViewPointId"] = request.CustomViewPointId
	}

	if !dara.IsNil(request.CustomViewPointIdsShrink) {
		body["CustomViewPointIds"] = request.CustomViewPointIdsShrink
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Topic) {
		body["Topic"] = request.Topic
	}

	if !dara.IsNil(request.TopicId) {
		body["TopicId"] = request.TopicId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCustomViewPoints"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCustomViewPointsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询数据集文档列表
//
// @param tmpReq - ListDatasetDocumentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatasetDocumentsResponse
func (client *Client) ListDatasetDocumentsWithContext(ctx context.Context, tmpReq *ListDatasetDocumentsRequest, runtime *dara.RuntimeOptions) (_result *ListDatasetDocumentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDatasetDocumentsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ExcludeFields) {
		request.ExcludeFieldsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExcludeFields, dara.String("ExcludeFields"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.IncludeFields) {
		request.IncludeFieldsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IncludeFields, dara.String("IncludeFields"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DatasetDescription) {
		body["DatasetDescription"] = request.DatasetDescription
	}

	if !dara.IsNil(request.DatasetId) {
		body["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.DatasetName) {
		body["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.DocType) {
		body["DocType"] = request.DocType
	}

	if !dara.IsNil(request.ExcludeFieldsShrink) {
		body["ExcludeFields"] = request.ExcludeFieldsShrink
	}

	if !dara.IsNil(request.IncludeFieldsShrink) {
		body["IncludeFields"] = request.IncludeFieldsShrink
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Query) {
		body["Query"] = request.Query
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDatasetDocuments"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDatasetDocumentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 数据集管理-查询
//
// @param request - ListDatasetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatasetsResponse
func (client *Client) ListDatasetsWithContext(ctx context.Context, request *ListDatasetsRequest, runtime *dara.RuntimeOptions) (_result *ListDatasetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DatasetId) {
		body["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.DatasetName) {
		body["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.DatasetType) {
		body["DatasetType"] = request.DatasetType
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IncludeConfig) {
		body["IncludeConfig"] = request.IncludeConfig
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchDatasetEnable) {
		body["SearchDatasetEnable"] = request.SearchDatasetEnable
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDatasets"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDatasetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成历史列表
//
// @param request - ListDialoguesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDialoguesResponse
func (client *Client) ListDialoguesWithContext(ctx context.Context, request *ListDialoguesRequest, runtime *dara.RuntimeOptions) (_result *ListDialoguesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Current) {
		body["Current"] = request.Current
	}

	if !dara.IsNil(request.DialogueType) {
		body["DialogueType"] = request.DialogueType
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Size) {
		body["Size"] = request.Size
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDialogues"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDialoguesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 妙读获取文档列表
//
// @param tmpReq - ListDocsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDocsResponse
func (client *Client) ListDocsWithContext(ctx context.Context, tmpReq *ListDocsRequest, runtime *dara.RuntimeOptions) (_result *ListDocsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDocsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Statuses) {
		request.StatusesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Statuses, dara.String("Statuses"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		body["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.DocName) {
		body["DocName"] = request.DocName
	}

	if !dara.IsNil(request.DocType) {
		body["DocType"] = request.DocType
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Skip) {
		body["Skip"] = request.Skip
	}

	if !dara.IsNil(request.StatusesShrink) {
		body["Statuses"] = request.StatusesShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDocs"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDocsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 公文检索
//
// @param request - ListDocumentRetrieveRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDocumentRetrieveResponse
func (client *Client) ListDocumentRetrieveWithContext(ctx context.Context, request *ListDocumentRetrieveRequest, runtime *dara.RuntimeOptions) (_result *ListDocumentRetrieveResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ContentType) {
		body["ContentType"] = request.ContentType
	}

	if !dara.IsNil(request.ElementScope) {
		body["ElementScope"] = request.ElementScope
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Office) {
		body["Office"] = request.Office
	}

	if !dara.IsNil(request.Query) {
		body["Query"] = request.Query
	}

	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	if !dara.IsNil(request.Source) {
		body["Source"] = request.Source
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.SubContentType) {
		body["SubContentType"] = request.SubContentType
	}

	if !dara.IsNil(request.SubjectClassify) {
		body["SubjectClassify"] = request.SubjectClassify
	}

	if !dara.IsNil(request.WordSize) {
		body["WordSize"] = request.WordSize
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDocumentRetrieve"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDocumentRetrieveResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新颖视角列表
//
// @param request - ListFreshViewPointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFreshViewPointsResponse
func (client *Client) ListFreshViewPointsWithContext(ctx context.Context, request *ListFreshViewPointsRequest, runtime *dara.RuntimeOptions) (_result *ListFreshViewPointsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Topic) {
		body["Topic"] = request.Topic
	}

	if !dara.IsNil(request.TopicSource) {
		body["TopicSource"] = request.TopicSource
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFreshViewPoints"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFreshViewPointsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通用配置-列表
//
// @param request - ListGeneralConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGeneralConfigsResponse
func (client *Client) ListGeneralConfigsWithContext(ctx context.Context, request *ListGeneralConfigsRequest, runtime *dara.RuntimeOptions) (_result *ListGeneralConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGeneralConfigs"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGeneralConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档管理-列表。
//
// @param request - ListGeneratedContentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGeneratedContentsResponse
func (client *Client) ListGeneratedContentsWithContext(ctx context.Context, request *ListGeneratedContentsRequest, runtime *dara.RuntimeOptions) (_result *ListGeneratedContentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ContentDomain) {
		body["ContentDomain"] = request.ContentDomain
	}

	if !dara.IsNil(request.Current) {
		body["Current"] = request.Current
	}

	if !dara.IsNil(request.DataType) {
		body["DataType"] = request.DataType
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Query) {
		body["Query"] = request.Query
	}

	if !dara.IsNil(request.Size) {
		body["Size"] = request.Size
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGeneratedContents"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGeneratedContentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取分类的热点新闻
//
// @param tmpReq - ListHotNewsWithTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHotNewsWithTypeResponse
func (client *Client) ListHotNewsWithTypeWithContext(ctx context.Context, tmpReq *ListHotNewsWithTypeRequest, runtime *dara.RuntimeOptions) (_result *ListHotNewsWithTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListHotNewsWithTypeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NewsTypes) {
		request.NewsTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NewsTypes, dara.String("NewsTypes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Current) {
		body["Current"] = request.Current
	}

	if !dara.IsNil(request.NewsType) {
		body["NewsType"] = request.NewsType
	}

	if !dara.IsNil(request.NewsTypesShrink) {
		body["NewsTypes"] = request.NewsTypesShrink
	}

	if !dara.IsNil(request.Size) {
		body["Size"] = request.Size
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHotNewsWithType"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHotNewsWithTypeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取所有平台热榜源列表
//
// @param request - ListHotSourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHotSourcesResponse
func (client *Client) ListHotSourcesWithContext(ctx context.Context, request *ListHotSourcesRequest, runtime *dara.RuntimeOptions) (_result *ListHotSourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHotSources"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHotSourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取热点事件列表
//
// @param tmpReq - ListHotTopicsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHotTopicsResponse
func (client *Client) ListHotTopicsWithContext(ctx context.Context, tmpReq *ListHotTopicsRequest, runtime *dara.RuntimeOptions) (_result *ListHotTopicsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListHotTopicsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TopicIds) {
		request.TopicIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TopicIds, dara.String("TopicIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Topics) {
		request.TopicsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Topics, dara.String("Topics"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TopicIdsShrink) {
		body["TopicIds"] = request.TopicIdsShrink
	}

	if !dara.IsNil(request.TopicQuery) {
		body["TopicQuery"] = request.TopicQuery
	}

	if !dara.IsNil(request.TopicSource) {
		body["TopicSource"] = request.TopicSource
	}

	if !dara.IsNil(request.TopicVersion) {
		body["TopicVersion"] = request.TopicVersion
	}

	if !dara.IsNil(request.TopicsShrink) {
		body["Topics"] = request.TopicsShrink
	}

	if !dara.IsNil(request.WithNews) {
		body["WithNews"] = request.WithNews
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHotTopics"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHotTopicsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 热门视角列表
//
// @param request - ListHotViewPointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHotViewPointsResponse
func (client *Client) ListHotViewPointsWithContext(ctx context.Context, request *ListHotViewPointsRequest, runtime *dara.RuntimeOptions) (_result *ListHotViewPointsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Topic) {
		body["Topic"] = request.Topic
	}

	if !dara.IsNil(request.TopicSource) {
		body["TopicSource"] = request.TopicSource
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHotViewPoints"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHotViewPointsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获得干预项目数量列表
//
// @param request - ListInterveneCntRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInterveneCntResponse
func (client *Client) ListInterveneCntWithContext(ctx context.Context, request *ListInterveneCntRequest, runtime *dara.RuntimeOptions) (_result *ListInterveneCntResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PageIndex) {
		body["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInterveneCnt"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInterveneCntResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获得导入任务列表
//
// @param request - ListInterveneImportTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInterveneImportTasksResponse
func (client *Client) ListInterveneImportTasksWithContext(ctx context.Context, request *ListInterveneImportTasksRequest, runtime *dara.RuntimeOptions) (_result *ListInterveneImportTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PageIndex) {
		body["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInterveneImportTasks"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInterveneImportTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获得干预规则列表
//
// @param request - ListInterveneRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInterveneRulesResponse
func (client *Client) ListInterveneRulesWithContext(ctx context.Context, request *ListInterveneRulesRequest, runtime *dara.RuntimeOptions) (_result *ListInterveneRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PageIndex) {
		body["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInterveneRules"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInterveneRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获得干预项列表
//
// @param request - ListIntervenesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntervenesResponse
func (client *Client) ListIntervenesWithContext(ctx context.Context, request *ListIntervenesRequest, runtime *dara.RuntimeOptions) (_result *ListIntervenesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InterveneType) {
		body["InterveneType"] = request.InterveneType
	}

	if !dara.IsNil(request.PageIndex) {
		body["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Query) {
		body["Query"] = request.Query
	}

	if !dara.IsNil(request.RuleId) {
		body["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIntervenes"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIntervenesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询素材列表
//
// @param tmpReq - ListMaterialDocumentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMaterialDocumentsResponse
func (client *Client) ListMaterialDocumentsWithContext(ctx context.Context, tmpReq *ListMaterialDocumentsRequest, runtime *dara.RuntimeOptions) (_result *ListMaterialDocumentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListMaterialDocumentsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DocTypeList) {
		request.DocTypeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DocTypeList, dara.String("DocTypeList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Keywords) {
		request.KeywordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Keywords, dara.String("Keywords"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.CreateTimeEnd) {
		body["CreateTimeEnd"] = request.CreateTimeEnd
	}

	if !dara.IsNil(request.CreateTimeStart) {
		body["CreateTimeStart"] = request.CreateTimeStart
	}

	if !dara.IsNil(request.Current) {
		body["Current"] = request.Current
	}

	if !dara.IsNil(request.DocType) {
		body["DocType"] = request.DocType
	}

	if !dara.IsNil(request.DocTypeListShrink) {
		body["DocTypeList"] = request.DocTypeListShrink
	}

	if !dara.IsNil(request.GeneratePublicUrl) {
		body["GeneratePublicUrl"] = request.GeneratePublicUrl
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.KeywordsShrink) {
		body["Keywords"] = request.KeywordsShrink
	}

	if !dara.IsNil(request.Query) {
		body["Query"] = request.Query
	}

	if !dara.IsNil(request.ShareAttr) {
		body["ShareAttr"] = request.ShareAttr
	}

	if !dara.IsNil(request.Size) {
		body["Size"] = request.Size
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	if !dara.IsNil(request.UpdateTimeEnd) {
		body["UpdateTimeEnd"] = request.UpdateTimeEnd
	}

	if !dara.IsNil(request.UpdateTimeStart) {
		body["UpdateTimeStart"] = request.UpdateTimeStart
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMaterialDocuments"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMaterialDocumentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取选题策划列表
//
// @param tmpReq - ListPlanningProposalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPlanningProposalResponse
func (client *Client) ListPlanningProposalWithContext(ctx context.Context, tmpReq *ListPlanningProposalRequest, runtime *dara.RuntimeOptions) (_result *ListPlanningProposalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListPlanningProposalShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CustomViewPointIds) {
		request.CustomViewPointIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomViewPointIds, dara.String("CustomViewPointIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Titles) {
		request.TitlesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Titles, dara.String("Titles"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CustomViewPointId) {
		body["CustomViewPointId"] = request.CustomViewPointId
	}

	if !dara.IsNil(request.CustomViewPointIdsShrink) {
		body["CustomViewPointIds"] = request.CustomViewPointIdsShrink
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TitlesShrink) {
		body["Titles"] = request.TitlesShrink
	}

	if !dara.IsNil(request.Topic) {
		body["Topic"] = request.Topic
	}

	if !dara.IsNil(request.TopicSource) {
		body["TopicSource"] = request.TopicSource
	}

	if !dara.IsNil(request.TopicVersion) {
		body["TopicVersion"] = request.TopicVersion
	}

	if !dara.IsNil(request.ViewPointType) {
		body["ViewPointType"] = request.ViewPointType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPlanningProposal"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPlanningProposalResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询搜索生成任务对话详情中数据列表
//
// @param request - ListSearchTaskDialogueDatasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSearchTaskDialogueDatasResponse
func (client *Client) ListSearchTaskDialogueDatasWithContext(ctx context.Context, request *ListSearchTaskDialogueDatasRequest, runtime *dara.RuntimeOptions) (_result *ListSearchTaskDialogueDatasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IncludeContent) {
		body["IncludeContent"] = request.IncludeContent
	}

	if !dara.IsNil(request.MultimodalSearchType) {
		body["MultimodalSearchType"] = request.MultimodalSearchType
	}

	if !dara.IsNil(request.OriginalSessionId) {
		body["OriginalSessionId"] = request.OriginalSessionId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Query) {
		body["Query"] = request.Query
	}

	if !dara.IsNil(request.SearchModel) {
		body["SearchModel"] = request.SearchModel
	}

	if !dara.IsNil(request.SearchModelDataValue) {
		body["SearchModelDataValue"] = request.SearchModelDataValue
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSearchTaskDialogueDatas"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSearchTaskDialogueDatasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询妙搜搜索生成任务详情列表
//
// @param request - ListSearchTaskDialoguesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSearchTaskDialoguesResponse
func (client *Client) ListSearchTaskDialoguesWithContext(ctx context.Context, request *ListSearchTaskDialoguesRequest, runtime *dara.RuntimeOptions) (_result *ListSearchTaskDialoguesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSearchTaskDialogues"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSearchTaskDialoguesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询妙搜搜索生成历史任务列表
//
// @param tmpReq - ListSearchTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSearchTasksResponse
func (client *Client) ListSearchTasksWithContext(ctx context.Context, tmpReq *ListSearchTasksRequest, runtime *dara.RuntimeOptions) (_result *ListSearchTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListSearchTasksShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DialogueTypes) {
		request.DialogueTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DialogueTypes, dara.String("DialogueTypes"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DialogueTypesShrink) {
		body["DialogueTypes"] = request.DialogueTypesShrink
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSearchTasks"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSearchTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文体学习分析结果列表
//
// @param request - ListStyleLearningResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListStyleLearningResultResponse
func (client *Client) ListStyleLearningResultWithContext(ctx context.Context, request *ListStyleLearningResultRequest, runtime *dara.RuntimeOptions) (_result *ListStyleLearningResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Current) {
		body["Current"] = request.Current
	}

	if !dara.IsNil(request.Size) {
		body["Size"] = request.Size
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListStyleLearningResult"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListStyleLearningResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 时效性视角列表
//
// @param request - ListTimedViewAttitudeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTimedViewAttitudeResponse
func (client *Client) ListTimedViewAttitudeWithContext(ctx context.Context, request *ListTimedViewAttitudeRequest, runtime *dara.RuntimeOptions) (_result *ListTimedViewAttitudeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Topic) {
		body["Topic"] = request.Topic
	}

	if !dara.IsNil(request.TopicSource) {
		body["TopicSource"] = request.TopicSource
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTimedViewAttitude"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTimedViewAttitudeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取热点推荐事件
//
// @param request - ListTopicRecommendEventListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTopicRecommendEventListResponse
func (client *Client) ListTopicRecommendEventListWithContext(ctx context.Context, request *ListTopicRecommendEventListRequest, runtime *dara.RuntimeOptions) (_result *ListTopicRecommendEventListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTopicRecommendEventList"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTopicRecommendEventListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取主题事件推荐观点列表
//
// @param request - ListTopicViewPointRecommendEventListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTopicViewPointRecommendEventListResponse
func (client *Client) ListTopicViewPointRecommendEventListWithContext(ctx context.Context, request *ListTopicViewPointRecommendEventListRequest, runtime *dara.RuntimeOptions) (_result *ListTopicViewPointRecommendEventListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTopicViewPointRecommendEventList"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTopicViewPointRecommendEventListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取系统所有实例信息
//
// @param request - ListVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVersionsResponse
func (client *Client) ListVersionsWithContext(ctx context.Context, request *ListVersionsRequest, runtime *dara.RuntimeOptions) (_result *ListVersionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVersions"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVersionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 网友视角列表
//
// @param request - ListWebReviewPointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWebReviewPointsResponse
func (client *Client) ListWebReviewPointsWithContext(ctx context.Context, request *ListWebReviewPointsRequest, runtime *dara.RuntimeOptions) (_result *ListWebReviewPointsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Topic) {
		body["Topic"] = request.Topic
	}

	if !dara.IsNil(request.TopicSource) {
		body["TopicSource"] = request.TopicSource
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWebReviewPoints"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWebReviewPointsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文体列表
//
// @param request - ListWritingStylesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWritingStylesResponse
func (client *Client) ListWritingStylesWithContext(ctx context.Context, request *ListWritingStylesRequest, runtime *dara.RuntimeOptions) (_result *ListWritingStylesResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Scene) {
		body["Scene"] = request.Scene
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWritingStyles"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWritingStylesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据taskId查询异步任务状态
//
// @param request - QueryAsyncTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAsyncTaskResponse
func (client *Client) QueryAsyncTaskWithContext(ctx context.Context, request *QueryAsyncTaskRequest, runtime *dara.RuntimeOptions) (_result *QueryAsyncTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAsyncTask"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAsyncTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询审核结果
//
// @param request - QueryAuditTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAuditTaskResponse
func (client *Client) QueryAuditTaskWithContext(ctx context.Context, request *QueryAuditTaskRequest, runtime *dara.RuntimeOptions) (_result *QueryAuditTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ArticleId) {
		body["ArticleId"] = request.ArticleId
	}

	if !dara.IsNil(request.ContentAuditTaskId) {
		body["ContentAuditTaskId"] = request.ContentAuditTaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAuditTask"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAuditTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 内容缩写
//
// @param request - RunAbbreviationContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunAbbreviationContentResponse
func (client *Client) RunAbbreviationContentWithSSECtx(ctx context.Context, request *RunAbbreviationContentRequest, runtime *dara.RuntimeOptions, _yield chan *RunAbbreviationContentResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runAbbreviationContentWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, runtime)
	return
}

// Summary:
//
// 内容缩写
//
// @param request - RunAbbreviationContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunAbbreviationContentResponse
func (client *Client) RunAbbreviationContentWithContext(ctx context.Context, request *RunAbbreviationContentRequest, runtime *dara.RuntimeOptions) (_result *RunAbbreviationContentResponse, _err error) {
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

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunAbbreviationContent"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunAbbreviationContentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 妙笔：AI助手写作
//
// @param tmpReq - RunAiHelperWritingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunAiHelperWritingResponse
func (client *Client) RunAiHelperWritingWithSSECtx(ctx context.Context, tmpReq *RunAiHelperWritingRequest, runtime *dara.RuntimeOptions, _yield chan *RunAiHelperWritingResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runAiHelperWritingWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, tmpReq, runtime)
	return
}

// Summary:
//
// 妙笔：AI助手写作
//
// @param tmpReq - RunAiHelperWritingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunAiHelperWritingResponse
func (client *Client) RunAiHelperWritingWithContext(ctx context.Context, tmpReq *RunAiHelperWritingRequest, runtime *dara.RuntimeOptions) (_result *RunAiHelperWritingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunAiHelperWritingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.WritingParams) {
		request.WritingParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WritingParams, dara.String("WritingParams"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DistributeWriting) {
		body["DistributeWriting"] = request.DistributeWriting
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.PromptMode) {
		body["PromptMode"] = request.PromptMode
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	if !dara.IsNil(request.WritingParamsShrink) {
		body["WritingParams"] = request.WritingParamsShrink
	}

	if !dara.IsNil(request.WritingScene) {
		body["WritingScene"] = request.WritingScene
	}

	if !dara.IsNil(request.WritingStyle) {
		body["WritingStyle"] = request.WritingStyle
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunAiHelperWriting"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunAiHelperWritingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 妙读生成书籍脑图
//
// @param request - RunBookBrainmapRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunBookBrainmapResponse
func (client *Client) RunBookBrainmapWithSSECtx(ctx context.Context, request *RunBookBrainmapRequest, runtime *dara.RuntimeOptions, _yield chan *RunBookBrainmapResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runBookBrainmapWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, runtime)
	return
}

// Summary:
//
// 妙读生成书籍脑图
//
// @param request - RunBookBrainmapRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunBookBrainmapResponse
func (client *Client) RunBookBrainmapWithContext(ctx context.Context, request *RunBookBrainmapRequest, runtime *dara.RuntimeOptions) (_result *RunBookBrainmapResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CleanCache) {
		body["CleanCache"] = request.CleanCache
	}

	if !dara.IsNil(request.DocId) {
		body["DocId"] = request.DocId
	}

	if !dara.IsNil(request.NodeNumber) {
		body["NodeNumber"] = request.NodeNumber
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.WordNumber) {
		body["WordNumber"] = request.WordNumber
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunBookBrainmap"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunBookBrainmapResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 书籍导读接口
//
// @param request - RunBookIntroductionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunBookIntroductionResponse
func (client *Client) RunBookIntroductionWithSSECtx(ctx context.Context, request *RunBookIntroductionRequest, runtime *dara.RuntimeOptions, _yield chan *RunBookIntroductionResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runBookIntroductionWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, runtime)
	return
}

// Summary:
//
// 书籍导读接口
//
// @param request - RunBookIntroductionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunBookIntroductionResponse
func (client *Client) RunBookIntroductionWithContext(ctx context.Context, request *RunBookIntroductionRequest, runtime *dara.RuntimeOptions) (_result *RunBookIntroductionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocId) {
		body["DocId"] = request.DocId
	}

	if !dara.IsNil(request.KeyPointPrompt) {
		body["KeyPointPrompt"] = request.KeyPointPrompt
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.SummaryPrompt) {
		body["SummaryPrompt"] = request.SummaryPrompt
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunBookIntroduction"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunBookIntroductionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 书籍智能卡片接口
//
// @param request - RunBookSmartCardRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunBookSmartCardResponse
func (client *Client) RunBookSmartCardWithSSECtx(ctx context.Context, request *RunBookSmartCardRequest, runtime *dara.RuntimeOptions, _yield chan *RunBookSmartCardResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runBookSmartCardWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, runtime)
	return
}

// Summary:
//
// 书籍智能卡片接口
//
// @param request - RunBookSmartCardRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunBookSmartCardResponse
func (client *Client) RunBookSmartCardWithContext(ctx context.Context, request *RunBookSmartCardRequest, runtime *dara.RuntimeOptions) (_result *RunBookSmartCardResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocId) {
		body["DocId"] = request.DocId
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunBookSmartCard"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunBookSmartCardResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 客户之声预测
//
// @param tmpReq - RunCommentGenerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunCommentGenerationResponse
func (client *Client) RunCommentGenerationWithSSECtx(ctx context.Context, tmpReq *RunCommentGenerationRequest, runtime *dara.RuntimeOptions, _yield chan *RunCommentGenerationResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runCommentGenerationWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, tmpReq, runtime)
	return
}

// Summary:
//
// 客户之声预测
//
// @param tmpReq - RunCommentGenerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunCommentGenerationResponse
func (client *Client) RunCommentGenerationWithContext(ctx context.Context, tmpReq *RunCommentGenerationRequest, runtime *dara.RuntimeOptions) (_result *RunCommentGenerationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunCommentGenerationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LengthRange) {
		request.LengthRangeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LengthRange, dara.String("LengthRange"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Sentiment) {
		request.SentimentShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sentiment, dara.String("Sentiment"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Type) {
		request.TypeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Type, dara.String("Type"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AllowEmoji) {
		body["AllowEmoji"] = request.AllowEmoji
	}

	if !dara.IsNil(request.ExtraInfo) {
		body["ExtraInfo"] = request.ExtraInfo
	}

	if !dara.IsNil(request.Length) {
		body["Length"] = request.Length
	}

	if !dara.IsNil(request.LengthRangeShrink) {
		body["LengthRange"] = request.LengthRangeShrink
	}

	if !dara.IsNil(request.ModelId) {
		body["ModelId"] = request.ModelId
	}

	if !dara.IsNil(request.NumComments) {
		body["NumComments"] = request.NumComments
	}

	if !dara.IsNil(request.SentimentShrink) {
		body["Sentiment"] = request.SentimentShrink
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.SourceMaterial) {
		body["SourceMaterial"] = request.SourceMaterial
	}

	if !dara.IsNil(request.Style) {
		body["Style"] = request.Style
	}

	if !dara.IsNil(request.TypeShrink) {
		body["Type"] = request.TypeShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunCommentGeneration"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunCommentGenerationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 内容续写
//
// @param request - RunContinueContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunContinueContentResponse
func (client *Client) RunContinueContentWithSSECtx(ctx context.Context, request *RunContinueContentRequest, runtime *dara.RuntimeOptions, _yield chan *RunContinueContentResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runContinueContentWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, runtime)
	return
}

// Summary:
//
// 内容续写
//
// @param request - RunContinueContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunContinueContentResponse
func (client *Client) RunContinueContentWithContext(ctx context.Context, request *RunContinueContentRequest, runtime *dara.RuntimeOptions) (_result *RunContinueContentResponse, _err error) {
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

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunContinueContent"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunContinueContentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 自定义热点话题分析
//
// @param request - RunCustomHotTopicAnalysisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunCustomHotTopicAnalysisResponse
func (client *Client) RunCustomHotTopicAnalysisWithSSECtx(ctx context.Context, request *RunCustomHotTopicAnalysisRequest, runtime *dara.RuntimeOptions, _yield chan *RunCustomHotTopicAnalysisResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runCustomHotTopicAnalysisWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, runtime)
	return
}

// Summary:
//
// 自定义热点话题分析
//
// @param request - RunCustomHotTopicAnalysisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunCustomHotTopicAnalysisResponse
func (client *Client) RunCustomHotTopicAnalysisWithContext(ctx context.Context, request *RunCustomHotTopicAnalysisRequest, runtime *dara.RuntimeOptions) (_result *RunCustomHotTopicAnalysisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AskUser) {
		body["AskUser"] = request.AskUser
	}

	if !dara.IsNil(request.ForceAnalysisExistsTopic) {
		body["ForceAnalysisExistsTopic"] = request.ForceAnalysisExistsTopic
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.UserBack) {
		body["UserBack"] = request.UserBack
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunCustomHotTopicAnalysis"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunCustomHotTopicAnalysisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 自定义选题视角分析
//
// @param request - RunCustomHotTopicViewPointAnalysisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunCustomHotTopicViewPointAnalysisResponse
func (client *Client) RunCustomHotTopicViewPointAnalysisWithSSECtx(ctx context.Context, request *RunCustomHotTopicViewPointAnalysisRequest, runtime *dara.RuntimeOptions, _yield chan *RunCustomHotTopicViewPointAnalysisResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runCustomHotTopicViewPointAnalysisWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, runtime)
	return
}

// Summary:
//
// 自定义选题视角分析
//
// @param request - RunCustomHotTopicViewPointAnalysisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunCustomHotTopicViewPointAnalysisResponse
func (client *Client) RunCustomHotTopicViewPointAnalysisWithContext(ctx context.Context, request *RunCustomHotTopicViewPointAnalysisRequest, runtime *dara.RuntimeOptions) (_result *RunCustomHotTopicViewPointAnalysisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AskUser) {
		body["AskUser"] = request.AskUser
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.SearchQuery) {
		body["SearchQuery"] = request.SearchQuery
	}

	if !dara.IsNil(request.SkipAskUser) {
		body["SkipAskUser"] = request.SkipAskUser
	}

	if !dara.IsNil(request.Topic) {
		body["Topic"] = request.Topic
	}

	if !dara.IsNil(request.TopicId) {
		body["TopicId"] = request.TopicId
	}

	if !dara.IsNil(request.TopicSource) {
		body["TopicSource"] = request.TopicSource
	}

	if !dara.IsNil(request.TopicVersion) {
		body["TopicVersion"] = request.TopicVersion
	}

	if !dara.IsNil(request.UserBack) {
		body["UserBack"] = request.UserBack
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunCustomHotTopicViewPointAnalysis"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunCustomHotTopicViewPointAnalysisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 流式输出深度写作事件
//
// @param request - RunDeepWritingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunDeepWritingResponse
func (client *Client) RunDeepWritingWithSSECtx(ctx context.Context, request *RunDeepWritingRequest, runtime *dara.RuntimeOptions, _yield chan *RunDeepWritingResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runDeepWritingWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, runtime)
	return
}

// Summary:
//
// 流式输出深度写作事件
//
// @param request - RunDeepWritingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunDeepWritingResponse
func (client *Client) RunDeepWritingWithContext(ctx context.Context, request *RunDeepWritingRequest, runtime *dara.RuntimeOptions) (_result *RunDeepWritingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Cursor) {
		body["Cursor"] = request.Cursor
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunDeepWriting"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunDeepWritingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 妙读脑图生成接口
//
// @param request - RunDocBrainmapRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunDocBrainmapResponse
func (client *Client) RunDocBrainmapWithSSECtx(ctx context.Context, request *RunDocBrainmapRequest, runtime *dara.RuntimeOptions, _yield chan *RunDocBrainmapResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runDocBrainmapWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, runtime)
	return
}

// Summary:
//
// 妙读脑图生成接口
//
// @param request - RunDocBrainmapRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunDocBrainmapResponse
func (client *Client) RunDocBrainmapWithContext(ctx context.Context, request *RunDocBrainmapRequest, runtime *dara.RuntimeOptions) (_result *RunDocBrainmapResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CleanCache) {
		body["CleanCache"] = request.CleanCache
	}

	if !dara.IsNil(request.DocId) {
		body["DocId"] = request.DocId
	}

	if !dara.IsNil(request.ModelName) {
		body["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.NodeNumber) {
		body["NodeNumber"] = request.NodeNumber
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.WordNumber) {
		body["WordNumber"] = request.WordNumber
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	if !dara.IsNil(request.ReferenceContent) {
		body["referenceContent"] = request.ReferenceContent
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunDocBrainmap"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunDocBrainmapResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 妙读文档导读接口
//
// @param request - RunDocIntroductionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunDocIntroductionResponse
func (client *Client) RunDocIntroductionWithSSECtx(ctx context.Context, request *RunDocIntroductionRequest, runtime *dara.RuntimeOptions, _yield chan *RunDocIntroductionResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runDocIntroductionWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, runtime)
	return
}

// Summary:
//
// 妙读文档导读接口
//
// @param request - RunDocIntroductionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunDocIntroductionResponse
func (client *Client) RunDocIntroductionWithContext(ctx context.Context, request *RunDocIntroductionRequest, runtime *dara.RuntimeOptions) (_result *RunDocIntroductionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CleanCache) {
		body["CleanCache"] = request.CleanCache
	}

	if !dara.IsNil(request.DocId) {
		body["DocId"] = request.DocId
	}

	if !dara.IsNil(request.IntroductionPrompt) {
		body["IntroductionPrompt"] = request.IntroductionPrompt
	}

	if !dara.IsNil(request.KeyPointPrompt) {
		body["KeyPointPrompt"] = request.KeyPointPrompt
	}

	if !dara.IsNil(request.ModelName) {
		body["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.SummaryPrompt) {
		body["SummaryPrompt"] = request.SummaryPrompt
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	if !dara.IsNil(request.ReferenceContent) {
		body["referenceContent"] = request.ReferenceContent
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunDocIntroduction"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunDocIntroductionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 妙读问答接口
//
// @param tmpReq - RunDocQaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunDocQaResponse
func (client *Client) RunDocQaWithSSECtx(ctx context.Context, tmpReq *RunDocQaRequest, runtime *dara.RuntimeOptions, _yield chan *RunDocQaResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runDocQaWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, tmpReq, runtime)
	return
}

// Summary:
//
// 妙读问答接口
//
// @param tmpReq - RunDocQaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunDocQaResponse
func (client *Client) RunDocQaWithContext(ctx context.Context, tmpReq *RunDocQaRequest, runtime *dara.RuntimeOptions) (_result *RunDocQaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunDocQaShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CategoryIds) {
		request.CategoryIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CategoryIds, dara.String("CategoryIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ConversationContexts) {
		request.ConversationContextsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ConversationContexts, dara.String("ConversationContexts"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DocIds) {
		request.DocIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DocIds, dara.String("DocIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoryIdsShrink) {
		body["CategoryIds"] = request.CategoryIdsShrink
	}

	if !dara.IsNil(request.ConversationContextsShrink) {
		body["ConversationContexts"] = request.ConversationContextsShrink
	}

	if !dara.IsNil(request.DocIdsShrink) {
		body["DocIds"] = request.DocIdsShrink
	}

	if !dara.IsNil(request.ModelName) {
		body["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.Query) {
		body["Query"] = request.Query
	}

	if !dara.IsNil(request.ReferenceContent) {
		body["ReferenceContent"] = request.ReferenceContent
	}

	if !dara.IsNil(request.SearchSource) {
		body["SearchSource"] = request.SearchSource
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunDocQa"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunDocQaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档智能卡片接口
//
// @param request - RunDocSmartCardRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunDocSmartCardResponse
func (client *Client) RunDocSmartCardWithSSECtx(ctx context.Context, request *RunDocSmartCardRequest, runtime *dara.RuntimeOptions, _yield chan *RunDocSmartCardResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runDocSmartCardWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, runtime)
	return
}

// Summary:
//
// 文档智能卡片接口
//
// @param request - RunDocSmartCardRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunDocSmartCardResponse
func (client *Client) RunDocSmartCardWithContext(ctx context.Context, request *RunDocSmartCardRequest, runtime *dara.RuntimeOptions) (_result *RunDocSmartCardResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocId) {
		body["DocId"] = request.DocId
	}

	if !dara.IsNil(request.ModelName) {
		body["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunDocSmartCard"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunDocSmartCardResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 妙读文档总结摘要接口
//
// @param request - RunDocSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunDocSummaryResponse
func (client *Client) RunDocSummaryWithSSECtx(ctx context.Context, request *RunDocSummaryRequest, runtime *dara.RuntimeOptions, _yield chan *RunDocSummaryResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runDocSummaryWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, runtime)
	return
}

// Summary:
//
// 妙读文档总结摘要接口
//
// @param request - RunDocSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunDocSummaryResponse
func (client *Client) RunDocSummaryWithContext(ctx context.Context, request *RunDocSummaryRequest, runtime *dara.RuntimeOptions) (_result *RunDocSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CleanCache) {
		body["CleanCache"] = request.CleanCache
	}

	if !dara.IsNil(request.DocId) {
		body["DocId"] = request.DocId
	}

	if !dara.IsNil(request.ModelName) {
		body["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.Query) {
		body["Query"] = request.Query
	}

	if !dara.IsNil(request.RecommendContent) {
		body["RecommendContent"] = request.RecommendContent
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunDocSummary"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunDocSummaryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 妙读文档翻译接口
//
// @param request - RunDocTranslationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunDocTranslationResponse
func (client *Client) RunDocTranslationWithSSECtx(ctx context.Context, request *RunDocTranslationRequest, runtime *dara.RuntimeOptions, _yield chan *RunDocTranslationResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runDocTranslationWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, runtime)
	return
}

// Summary:
//
// 妙读文档翻译接口
//
// @param request - RunDocTranslationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunDocTranslationResponse
func (client *Client) RunDocTranslationWithContext(ctx context.Context, request *RunDocTranslationRequest, runtime *dara.RuntimeOptions) (_result *RunDocTranslationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CleanCache) {
		body["CleanCache"] = request.CleanCache
	}

	if !dara.IsNil(request.DocId) {
		body["DocId"] = request.DocId
	}

	if !dara.IsNil(request.ModelName) {
		body["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.RecommendContent) {
		body["RecommendContent"] = request.RecommendContent
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.TransType) {
		body["TransType"] = request.TransType
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunDocTranslation"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunDocTranslationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档改写
//
// @param request - RunDocWashingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunDocWashingResponse
func (client *Client) RunDocWashingWithSSECtx(ctx context.Context, request *RunDocWashingRequest, runtime *dara.RuntimeOptions, _yield chan *RunDocWashingResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runDocWashingWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, runtime)
	return
}

// Summary:
//
// 文档改写
//
// @param request - RunDocWashingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunDocWashingResponse
func (client *Client) RunDocWashingWithContext(ctx context.Context, request *RunDocWashingRequest, runtime *dara.RuntimeOptions) (_result *RunDocWashingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ModelId) {
		body["ModelId"] = request.ModelId
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.ReferenceContent) {
		body["ReferenceContent"] = request.ReferenceContent
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Topic) {
		body["Topic"] = request.Topic
	}

	if !dara.IsNil(request.WordNumber) {
		body["WordNumber"] = request.WordNumber
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	if !dara.IsNil(request.WritingTypeName) {
		body["WritingTypeName"] = request.WritingTypeName
	}

	if !dara.IsNil(request.WritingTypeRefDoc) {
		body["WritingTypeRefDoc"] = request.WritingTypeRefDoc
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunDocWashing"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunDocWashingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 内容扩写
//
// @param request - RunExpandContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunExpandContentResponse
func (client *Client) RunExpandContentWithSSECtx(ctx context.Context, request *RunExpandContentRequest, runtime *dara.RuntimeOptions, _yield chan *RunExpandContentResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runExpandContentWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, runtime)
	return
}

// Summary:
//
// 内容扩写
//
// @param request - RunExpandContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunExpandContentResponse
func (client *Client) RunExpandContentWithContext(ctx context.Context, request *RunExpandContentRequest, runtime *dara.RuntimeOptions) (_result *RunExpandContentResponse, _err error) {
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

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunExpandContent"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunExpandContentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 妙读猜你想问接口
//
// @param request - RunGenerateQuestionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunGenerateQuestionsResponse
func (client *Client) RunGenerateQuestionsWithSSECtx(ctx context.Context, request *RunGenerateQuestionsRequest, runtime *dara.RuntimeOptions, _yield chan *RunGenerateQuestionsResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runGenerateQuestionsWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, runtime)
	return
}

// Summary:
//
// 妙读猜你想问接口
//
// @param request - RunGenerateQuestionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunGenerateQuestionsResponse
func (client *Client) RunGenerateQuestionsWithContext(ctx context.Context, request *RunGenerateQuestionsRequest, runtime *dara.RuntimeOptions) (_result *RunGenerateQuestionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocId) {
		body["DocId"] = request.DocId
	}

	if !dara.IsNil(request.ModelName) {
		body["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.ReferenceContent) {
		body["ReferenceContent"] = request.ReferenceContent
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunGenerateQuestions"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunGenerateQuestionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 妙读文档关键词抽取接口
//
// @param request - RunHotwordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunHotwordResponse
func (client *Client) RunHotwordWithSSECtx(ctx context.Context, request *RunHotwordRequest, runtime *dara.RuntimeOptions, _yield chan *RunHotwordResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runHotwordWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, runtime)
	return
}

// Summary:
//
// 妙读文档关键词抽取接口
//
// @param request - RunHotwordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunHotwordResponse
func (client *Client) RunHotwordWithContext(ctx context.Context, request *RunHotwordRequest, runtime *dara.RuntimeOptions) (_result *RunHotwordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocId) {
		body["DocId"] = request.DocId
	}

	if !dara.IsNil(request.ModelName) {
		body["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.ReferenceContent) {
		body["ReferenceContent"] = request.ReferenceContent
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunHotword"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunHotwordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # AI妙笔-创作-抽取关键词
//
// @param tmpReq - RunKeywordsExtractionGenerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunKeywordsExtractionGenerationResponse
func (client *Client) RunKeywordsExtractionGenerationWithSSECtx(ctx context.Context, tmpReq *RunKeywordsExtractionGenerationRequest, runtime *dara.RuntimeOptions, _yield chan *RunKeywordsExtractionGenerationResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runKeywordsExtractionGenerationWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, tmpReq, runtime)
	return
}

// Summary:
//
// # AI妙笔-创作-抽取关键词
//
// @param tmpReq - RunKeywordsExtractionGenerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunKeywordsExtractionGenerationResponse
func (client *Client) RunKeywordsExtractionGenerationWithContext(ctx context.Context, tmpReq *RunKeywordsExtractionGenerationRequest, runtime *dara.RuntimeOptions) (_result *RunKeywordsExtractionGenerationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunKeywordsExtractionGenerationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ReferenceData) {
		request.ReferenceDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReferenceData, dara.String("ReferenceData"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.ReferenceDataShrink) {
		body["ReferenceData"] = request.ReferenceDataShrink
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunKeywordsExtractionGeneration"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunKeywordsExtractionGenerationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档批量导读
//
// @param tmpReq - RunMultiDocIntroductionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunMultiDocIntroductionResponse
func (client *Client) RunMultiDocIntroductionWithSSECtx(ctx context.Context, tmpReq *RunMultiDocIntroductionRequest, runtime *dara.RuntimeOptions, _yield chan *RunMultiDocIntroductionResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runMultiDocIntroductionWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, tmpReq, runtime)
	return
}

// Summary:
//
// 文档批量导读
//
// @param tmpReq - RunMultiDocIntroductionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunMultiDocIntroductionResponse
func (client *Client) RunMultiDocIntroductionWithContext(ctx context.Context, tmpReq *RunMultiDocIntroductionRequest, runtime *dara.RuntimeOptions) (_result *RunMultiDocIntroductionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunMultiDocIntroductionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DocIds) {
		request.DocIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DocIds, dara.String("DocIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DocIdsShrink) {
		body["DocIds"] = request.DocIdsShrink
	}

	if !dara.IsNil(request.KeyPointPrompt) {
		body["KeyPointPrompt"] = request.KeyPointPrompt
	}

	if !dara.IsNil(request.ModelName) {
		body["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.SummaryPrompt) {
		body["SummaryPrompt"] = request.SummaryPrompt
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunMultiDocIntroduction"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunMultiDocIntroductionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 流式输出PPT大纲
//
// @param request - RunPptOutlineGenerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunPptOutlineGenerationResponse
func (client *Client) RunPptOutlineGenerationWithSSECtx(ctx context.Context, request *RunPptOutlineGenerationRequest, runtime *dara.RuntimeOptions, _yield chan *RunPptOutlineGenerationResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runPptOutlineGenerationWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, runtime)
	return
}

// Summary:
//
// 流式输出PPT大纲
//
// @param request - RunPptOutlineGenerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunPptOutlineGenerationResponse
func (client *Client) RunPptOutlineGenerationWithContext(ctx context.Context, request *RunPptOutlineGenerationRequest, runtime *dara.RuntimeOptions) (_result *RunPptOutlineGenerationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunPptOutlineGeneration"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunPptOutlineGenerationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 快速写作
//
// @param tmpReq - RunQuickWritingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunQuickWritingResponse
func (client *Client) RunQuickWritingWithSSECtx(ctx context.Context, tmpReq *RunQuickWritingRequest, runtime *dara.RuntimeOptions, _yield chan *RunQuickWritingResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runQuickWritingWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, tmpReq, runtime)
	return
}

// Summary:
//
// 快速写作
//
// @param tmpReq - RunQuickWritingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunQuickWritingResponse
func (client *Client) RunQuickWritingWithContext(ctx context.Context, tmpReq *RunQuickWritingRequest, runtime *dara.RuntimeOptions) (_result *RunQuickWritingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunQuickWritingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Articles) {
		request.ArticlesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Articles, dara.String("Articles"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SearchSources) {
		request.SearchSourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SearchSources, dara.String("SearchSources"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ArticlesShrink) {
		body["Articles"] = request.ArticlesShrink
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.SearchSourcesShrink) {
		body["SearchSources"] = request.SearchSourcesShrink
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunQuickWriting"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunQuickWritingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # AI妙搜-智能搜索生成
//
// @param tmpReq - RunSearchGenerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunSearchGenerationResponse
func (client *Client) RunSearchGenerationWithSSECtx(ctx context.Context, tmpReq *RunSearchGenerationRequest, runtime *dara.RuntimeOptions, _yield chan *RunSearchGenerationResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runSearchGenerationWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, tmpReq, runtime)
	return
}

// Summary:
//
// # AI妙搜-智能搜索生成
//
// @param tmpReq - RunSearchGenerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunSearchGenerationResponse
func (client *Client) RunSearchGenerationWithContext(ctx context.Context, tmpReq *RunSearchGenerationRequest, runtime *dara.RuntimeOptions) (_result *RunSearchGenerationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunSearchGenerationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AgentContext) {
		request.AgentContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AgentContext, dara.String("AgentContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ChatConfig) {
		request.ChatConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChatConfig, dara.String("ChatConfig"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentContextShrink) {
		body["AgentContext"] = request.AgentContextShrink
	}

	if !dara.IsNil(request.ChatConfigShrink) {
		body["ChatConfig"] = request.ChatConfigShrink
	}

	if !dara.IsNil(request.FileUrl) {
		body["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.ModelId) {
		body["ModelId"] = request.ModelId
	}

	if !dara.IsNil(request.OriginalSessionId) {
		body["OriginalSessionId"] = request.OriginalSessionId
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunSearchGeneration"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunSearchGenerationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 妙搜-文搜文
//
// @param tmpReq - RunSearchSimilarArticlesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunSearchSimilarArticlesResponse
func (client *Client) RunSearchSimilarArticlesWithSSECtx(ctx context.Context, tmpReq *RunSearchSimilarArticlesRequest, runtime *dara.RuntimeOptions, _yield chan *RunSearchSimilarArticlesResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runSearchSimilarArticlesWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, tmpReq, runtime)
	return
}

// Summary:
//
// 妙搜-文搜文
//
// @param tmpReq - RunSearchSimilarArticlesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunSearchSimilarArticlesResponse
func (client *Client) RunSearchSimilarArticlesWithContext(ctx context.Context, tmpReq *RunSearchSimilarArticlesRequest, runtime *dara.RuntimeOptions) (_result *RunSearchSimilarArticlesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunSearchSimilarArticlesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ChatConfig) {
		request.ChatConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChatConfig, dara.String("ChatConfig"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ChatConfigShrink) {
		body["ChatConfig"] = request.ChatConfigShrink
	}

	if !dara.IsNil(request.DocType) {
		body["DocType"] = request.DocType
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	if !dara.IsNil(request.Url) {
		body["Url"] = request.Url
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunSearchSimilarArticles"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunSearchSimilarArticlesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创作-分步骤写作
//
// @param tmpReq - RunStepByStepWritingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunStepByStepWritingResponse
func (client *Client) RunStepByStepWritingWithSSECtx(ctx context.Context, tmpReq *RunStepByStepWritingRequest, runtime *dara.RuntimeOptions, _yield chan *RunStepByStepWritingResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runStepByStepWritingWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, tmpReq, runtime)
	return
}

// Summary:
//
// 创作-分步骤写作
//
// @param tmpReq - RunStepByStepWritingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunStepByStepWritingResponse
func (client *Client) RunStepByStepWritingWithContext(ctx context.Context, tmpReq *RunStepByStepWritingRequest, runtime *dara.RuntimeOptions) (_result *RunStepByStepWritingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunStepByStepWritingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ReferenceData) {
		request.ReferenceDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReferenceData, dara.String("ReferenceData"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.WritingConfig) {
		request.WritingConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WritingConfig, dara.String("WritingConfig"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OriginSessionId) {
		body["OriginSessionId"] = request.OriginSessionId
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.ReferenceDataShrink) {
		body["ReferenceData"] = request.ReferenceDataShrink
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	if !dara.IsNil(request.WritingConfigShrink) {
		body["WritingConfig"] = request.WritingConfigShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunStepByStepWriting"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunStepByStepWritingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 内容特点分析
//
// @param tmpReq - RunStyleFeatureAnalysisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunStyleFeatureAnalysisResponse
func (client *Client) RunStyleFeatureAnalysisWithSSECtx(ctx context.Context, tmpReq *RunStyleFeatureAnalysisRequest, runtime *dara.RuntimeOptions, _yield chan *RunStyleFeatureAnalysisResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runStyleFeatureAnalysisWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, tmpReq, runtime)
	return
}

// Summary:
//
// 内容特点分析
//
// @param tmpReq - RunStyleFeatureAnalysisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunStyleFeatureAnalysisResponse
func (client *Client) RunStyleFeatureAnalysisWithContext(ctx context.Context, tmpReq *RunStyleFeatureAnalysisRequest, runtime *dara.RuntimeOptions) (_result *RunStyleFeatureAnalysisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunStyleFeatureAnalysisShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Contents) {
		request.ContentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Contents, dara.String("Contents"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.MaterialIds) {
		request.MaterialIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MaterialIds, dara.String("MaterialIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ContentsShrink) {
		body["Contents"] = request.ContentsShrink
	}

	if !dara.IsNil(request.MaterialIdsShrink) {
		body["MaterialIds"] = request.MaterialIdsShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunStyleFeatureAnalysis"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunStyleFeatureAnalysisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 内容摘要生成
//
// @param request - RunSummaryGenerateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunSummaryGenerateResponse
func (client *Client) RunSummaryGenerateWithSSECtx(ctx context.Context, request *RunSummaryGenerateRequest, runtime *dara.RuntimeOptions, _yield chan *RunSummaryGenerateResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runSummaryGenerateWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, runtime)
	return
}

// Summary:
//
// 内容摘要生成
//
// @param request - RunSummaryGenerateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunSummaryGenerateResponse
func (client *Client) RunSummaryGenerateWithContext(ctx context.Context, request *RunSummaryGenerateRequest, runtime *dara.RuntimeOptions) (_result *RunSummaryGenerateResponse, _err error) {
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

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunSummaryGenerate"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunSummaryGenerateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创作-文本润色
//
// @param request - RunTextPolishingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunTextPolishingResponse
func (client *Client) RunTextPolishingWithSSECtx(ctx context.Context, request *RunTextPolishingRequest, runtime *dara.RuntimeOptions, _yield chan *RunTextPolishingResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runTextPolishingWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, runtime)
	return
}

// Summary:
//
// 创作-文本润色
//
// @param request - RunTextPolishingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunTextPolishingResponse
func (client *Client) RunTextPolishingWithContext(ctx context.Context, request *RunTextPolishingRequest, runtime *dara.RuntimeOptions) (_result *RunTextPolishingResponse, _err error) {
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

	if !dara.IsNil(request.OriginContent) {
		body["OriginContent"] = request.OriginContent
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunTextPolishing"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunTextPolishingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 妙笔：标题生成
//
// @param tmpReq - RunTitleGenerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunTitleGenerationResponse
func (client *Client) RunTitleGenerationWithSSECtx(ctx context.Context, tmpReq *RunTitleGenerationRequest, runtime *dara.RuntimeOptions, _yield chan *RunTitleGenerationResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runTitleGenerationWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, tmpReq, runtime)
	return
}

// Summary:
//
// 妙笔：标题生成
//
// @param tmpReq - RunTitleGenerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunTitleGenerationResponse
func (client *Client) RunTitleGenerationWithContext(ctx context.Context, tmpReq *RunTitleGenerationRequest, runtime *dara.RuntimeOptions) (_result *RunTitleGenerationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunTitleGenerationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeduplicatedTitles) {
		request.DeduplicatedTitlesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeduplicatedTitles, dara.String("DeduplicatedTitles"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ReferenceData) {
		request.ReferenceDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReferenceData, dara.String("ReferenceData"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeduplicatedTitlesShrink) {
		body["DeduplicatedTitles"] = request.DeduplicatedTitlesShrink
	}

	if !dara.IsNil(request.ReferenceDataShrink) {
		body["ReferenceData"] = request.ReferenceDataShrink
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TitleCount) {
		body["TitleCount"] = request.TitleCount
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunTitleGeneration"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunTitleGenerationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 妙策选题策划聚合
//
// @param tmpReq - RunTopicSelectionMergeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunTopicSelectionMergeResponse
func (client *Client) RunTopicSelectionMergeWithSSECtx(ctx context.Context, tmpReq *RunTopicSelectionMergeRequest, runtime *dara.RuntimeOptions, _yield chan *RunTopicSelectionMergeResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runTopicSelectionMergeWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, tmpReq, runtime)
	return
}

// Summary:
//
// 妙策选题策划聚合
//
// @param tmpReq - RunTopicSelectionMergeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunTopicSelectionMergeResponse
func (client *Client) RunTopicSelectionMergeWithContext(ctx context.Context, tmpReq *RunTopicSelectionMergeRequest, runtime *dara.RuntimeOptions) (_result *RunTopicSelectionMergeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunTopicSelectionMergeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Topics) {
		request.TopicsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Topics, dara.String("Topics"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.TopicsShrink) {
		body["Topics"] = request.TopicsShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunTopicSelectionMerge"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunTopicSelectionMergeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # AI妙笔-创作-中英文翻译
//
// @param tmpReq - RunTranslateGenerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunTranslateGenerationResponse
func (client *Client) RunTranslateGenerationWithSSECtx(ctx context.Context, tmpReq *RunTranslateGenerationRequest, runtime *dara.RuntimeOptions, _yield chan *RunTranslateGenerationResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runTranslateGenerationWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, tmpReq, runtime)
	return
}

// Summary:
//
// # AI妙笔-创作-中英文翻译
//
// @param tmpReq - RunTranslateGenerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunTranslateGenerationResponse
func (client *Client) RunTranslateGenerationWithContext(ctx context.Context, tmpReq *RunTranslateGenerationRequest, runtime *dara.RuntimeOptions) (_result *RunTranslateGenerationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunTranslateGenerationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ReferenceData) {
		request.ReferenceDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReferenceData, dara.String("ReferenceData"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.ReferenceDataShrink) {
		body["ReferenceData"] = request.ReferenceDataShrink
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunTranslateGeneration"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunTranslateGenerationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # AI生成视频剪辑脚本
//
// @param request - RunVideoScriptGenerateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunVideoScriptGenerateResponse
func (client *Client) RunVideoScriptGenerateWithSSECtx(ctx context.Context, request *RunVideoScriptGenerateRequest, runtime *dara.RuntimeOptions, _yield chan *RunVideoScriptGenerateResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runVideoScriptGenerateWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, runtime)
	return
}

// Summary:
//
// # AI生成视频剪辑脚本
//
// @param request - RunVideoScriptGenerateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunVideoScriptGenerateResponse
func (client *Client) RunVideoScriptGenerateWithContext(ctx context.Context, request *RunVideoScriptGenerateRequest, runtime *dara.RuntimeOptions) (_result *RunVideoScriptGenerateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.ScriptLength) {
		body["ScriptLength"] = request.ScriptLength
	}

	if !dara.IsNil(request.ScriptNumber) {
		body["ScriptNumber"] = request.ScriptNumber
	}

	if !dara.IsNil(request.UseSearch) {
		body["UseSearch"] = request.UseSearch
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunVideoScriptGenerate"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunVideoScriptGenerateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # AI妙笔-创作-文风改写
//
// @param tmpReq - RunWriteToneGenerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunWriteToneGenerationResponse
func (client *Client) RunWriteToneGenerationWithSSECtx(ctx context.Context, tmpReq *RunWriteToneGenerationRequest, runtime *dara.RuntimeOptions, _yield chan *RunWriteToneGenerationResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runWriteToneGenerationWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, tmpReq, runtime)
	return
}

// Summary:
//
// # AI妙笔-创作-文风改写
//
// @param tmpReq - RunWriteToneGenerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunWriteToneGenerationResponse
func (client *Client) RunWriteToneGenerationWithContext(ctx context.Context, tmpReq *RunWriteToneGenerationRequest, runtime *dara.RuntimeOptions) (_result *RunWriteToneGenerationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunWriteToneGenerationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ReferenceData) {
		request.ReferenceDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReferenceData, dara.String("ReferenceData"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.ReferenceDataShrink) {
		body["ReferenceData"] = request.ReferenceDataShrink
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunWriteToneGeneration"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunWriteToneGenerationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 直接写作
//
// @param tmpReq - RunWritingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunWritingResponse
func (client *Client) RunWritingWithSSECtx(ctx context.Context, tmpReq *RunWritingRequest, runtime *dara.RuntimeOptions, _yield chan *RunWritingResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runWritingWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, tmpReq, runtime)
	return
}

// Summary:
//
// 直接写作
//
// @param tmpReq - RunWritingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunWritingResponse
func (client *Client) RunWritingWithContext(ctx context.Context, tmpReq *RunWritingRequest, runtime *dara.RuntimeOptions) (_result *RunWritingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunWritingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ReferenceData) {
		request.ReferenceDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReferenceData, dara.String("ReferenceData"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.WritingConfig) {
		request.WritingConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WritingConfig, dara.String("WritingConfig"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OriginSessionId) {
		body["OriginSessionId"] = request.OriginSessionId
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.ReferenceDataShrink) {
		body["ReferenceData"] = request.ReferenceDataShrink
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	if !dara.IsNil(request.WritingConfigShrink) {
		body["WritingConfig"] = request.WritingConfigShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunWriting"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunWritingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 直接写作
//
// @param tmpReq - RunWritingV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunWritingV2Response
func (client *Client) RunWritingV2WithSSECtx(ctx context.Context, tmpReq *RunWritingV2Request, runtime *dara.RuntimeOptions, _yield chan *RunWritingV2Response, _yieldErr chan error) {
	defer close(_yield)
	client.runWritingV2WithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, tmpReq, runtime)
	return
}

// Summary:
//
// 直接写作
//
// @param tmpReq - RunWritingV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunWritingV2Response
func (client *Client) RunWritingV2WithContext(ctx context.Context, tmpReq *RunWritingV2Request, runtime *dara.RuntimeOptions) (_result *RunWritingV2Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunWritingV2ShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Articles) {
		request.ArticlesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Articles, dara.String("Articles"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Keywords) {
		request.KeywordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Keywords, dara.String("Keywords"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.MiniDocs) {
		request.MiniDocsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MiniDocs, dara.String("MiniDocs"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OutlineList) {
		request.OutlineListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OutlineList, dara.String("OutlineList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Outlines) {
		request.OutlinesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Outlines, dara.String("Outlines"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SearchSources) {
		request.SearchSourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SearchSources, dara.String("SearchSources"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Summarization) {
		request.SummarizationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Summarization, dara.String("Summarization"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.WritingParams) {
		request.WritingParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WritingParams, dara.String("WritingParams"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ArticlesShrink) {
		body["Articles"] = request.ArticlesShrink
	}

	if !dara.IsNil(request.DistributeWriting) {
		body["DistributeWriting"] = request.DistributeWriting
	}

	if !dara.IsNil(request.GcNumberSize) {
		body["GcNumberSize"] = request.GcNumberSize
	}

	if !dara.IsNil(request.GcNumberSizeTag) {
		body["GcNumberSizeTag"] = request.GcNumberSizeTag
	}

	if !dara.IsNil(request.KeywordsShrink) {
		body["Keywords"] = request.KeywordsShrink
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.MiniDocsShrink) {
		body["MiniDocs"] = request.MiniDocsShrink
	}

	if !dara.IsNil(request.OutlineListShrink) {
		body["OutlineList"] = request.OutlineListShrink
	}

	if !dara.IsNil(request.OutlinesShrink) {
		body["Outlines"] = request.OutlinesShrink
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.PromptMode) {
		body["PromptMode"] = request.PromptMode
	}

	if !dara.IsNil(request.SearchSourcesShrink) {
		body["SearchSources"] = request.SearchSourcesShrink
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.SourceTraceMethod) {
		body["SourceTraceMethod"] = request.SourceTraceMethod
	}

	if !dara.IsNil(request.Step) {
		body["Step"] = request.Step
	}

	if !dara.IsNil(request.SummarizationShrink) {
		body["Summarization"] = request.SummarizationShrink
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.UseSearch) {
		body["UseSearch"] = request.UseSearch
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	if !dara.IsNil(request.WritingParamsShrink) {
		body["WritingParams"] = request.WritingParamsShrink
	}

	if !dara.IsNil(request.WritingScene) {
		body["WritingScene"] = request.WritingScene
	}

	if !dara.IsNil(request.WritingStyle) {
		body["WritingStyle"] = request.WritingStyle
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunWritingV2"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunWritingV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存自定义文本
//
// @param request - SaveCustomTextRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveCustomTextResponse
func (client *Client) SaveCustomTextWithContext(ctx context.Context, request *SaveCustomTextRequest, runtime *dara.RuntimeOptions) (_result *SaveCustomTextResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CommodityCode) {
		body["CommodityCode"] = request.CommodityCode
	}

	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveCustomText"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveCustomTextResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存用户的信源配置
//
// @param tmpReq - SaveDataSourceOrderConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveDataSourceOrderConfigResponse
func (client *Client) SaveDataSourceOrderConfigWithContext(ctx context.Context, tmpReq *SaveDataSourceOrderConfigRequest, runtime *dara.RuntimeOptions) (_result *SaveDataSourceOrderConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SaveDataSourceOrderConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserConfigDataSourceList) {
		request.UserConfigDataSourceListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserConfigDataSourceList, dara.String("UserConfigDataSourceList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.GenerateTechnology) {
		body["GenerateTechnology"] = request.GenerateTechnology
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.UserConfigDataSourceListShrink) {
		body["UserConfigDataSourceList"] = request.UserConfigDataSourceListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveDataSourceOrderConfig"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveDataSourceOrderConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存素材
//
// @param tmpReq - SaveMaterialDocumentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveMaterialDocumentResponse
func (client *Client) SaveMaterialDocumentWithContext(ctx context.Context, tmpReq *SaveMaterialDocumentRequest, runtime *dara.RuntimeOptions) (_result *SaveMaterialDocumentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SaveMaterialDocumentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DocKeywords) {
		request.DocKeywordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DocKeywords, dara.String("DocKeywords"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Author) {
		body["Author"] = request.Author
	}

	if !dara.IsNil(request.BothSavePrivateAndShare) {
		body["BothSavePrivateAndShare"] = request.BothSavePrivateAndShare
	}

	if !dara.IsNil(request.DocKeywordsShrink) {
		body["DocKeywords"] = request.DocKeywordsShrink
	}

	if !dara.IsNil(request.DocType) {
		body["DocType"] = request.DocType
	}

	if !dara.IsNil(request.ExternalUrl) {
		body["ExternalUrl"] = request.ExternalUrl
	}

	if !dara.IsNil(request.HtmlContent) {
		body["HtmlContent"] = request.HtmlContent
	}

	if !dara.IsNil(request.PubTime) {
		body["PubTime"] = request.PubTime
	}

	if !dara.IsNil(request.ShareAttr) {
		body["ShareAttr"] = request.ShareAttr
	}

	if !dara.IsNil(request.SrcFrom) {
		body["SrcFrom"] = request.SrcFrom
	}

	if !dara.IsNil(request.Summary) {
		body["Summary"] = request.Summary
	}

	if !dara.IsNil(request.TextContent) {
		body["TextContent"] = request.TextContent
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	if !dara.IsNil(request.Url) {
		body["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveMaterialDocument"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveMaterialDocumentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存自定义文体
//
// @param tmpReq - SaveStyleLearningResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveStyleLearningResultResponse
func (client *Client) SaveStyleLearningResultWithContext(ctx context.Context, tmpReq *SaveStyleLearningResultRequest, runtime *dara.RuntimeOptions) (_result *SaveStyleLearningResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SaveStyleLearningResultShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CustomTextIdList) {
		request.CustomTextIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomTextIdList, dara.String("CustomTextIdList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.MaterialIdList) {
		request.MaterialIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MaterialIdList, dara.String("MaterialIdList"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		body["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.AigcResult) {
		body["AigcResult"] = request.AigcResult
	}

	if !dara.IsNil(request.CustomTextIdListShrink) {
		body["CustomTextIdList"] = request.CustomTextIdListShrink
	}

	if !dara.IsNil(request.MaterialIdListShrink) {
		body["MaterialIdList"] = request.MaterialIdListShrink
	}

	if !dara.IsNil(request.RewriteResult) {
		body["RewriteResult"] = request.RewriteResult
	}

	if !dara.IsNil(request.StyleName) {
		body["StyleName"] = request.StyleName
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveStyleLearningResult"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveStyleLearningResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 搜索数据集文档
//
// @param request - SearchDatasetDocumentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchDatasetDocumentsResponse
func (client *Client) SearchDatasetDocumentsWithContext(ctx context.Context, request *SearchDatasetDocumentsRequest, runtime *dara.RuntimeOptions) (_result *SearchDatasetDocumentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DatasetId) {
		body["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.DatasetName) {
		body["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.Extend1) {
		body["Extend1"] = request.Extend1
	}

	if !dara.IsNil(request.IncludeContent) {
		body["IncludeContent"] = request.IncludeContent
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Query) {
		body["Query"] = request.Query
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchDatasetDocuments"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchDatasetDocumentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新闻检索
//
// @param tmpReq - SearchNewsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchNewsResponse
func (client *Client) SearchNewsWithContext(ctx context.Context, tmpReq *SearchNewsRequest, runtime *dara.RuntimeOptions) (_result *SearchNewsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SearchNewsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SearchSources) {
		request.SearchSourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SearchSources, dara.String("SearchSources"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FilterNotNull) {
		body["FilterNotNull"] = request.FilterNotNull
	}

	if !dara.IsNil(request.IncludeContent) {
		body["IncludeContent"] = request.IncludeContent
	}

	if !dara.IsNil(request.Page) {
		body["Page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Query) {
		body["Query"] = request.Query
	}

	if !dara.IsNil(request.SearchSourcesShrink) {
		body["SearchSources"] = request.SearchSourcesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchNews"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchNewsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交异步任务
//
// @param request - SubmitAsyncTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitAsyncTaskResponse
func (client *Client) SubmitAsyncTaskWithContext(ctx context.Context, request *SubmitAsyncTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitAsyncTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskCode) {
		body["TaskCode"] = request.TaskCode
	}

	if !dara.IsNil(request.TaskExecuteTime) {
		body["TaskExecuteTime"] = request.TaskExecuteTime
	}

	if !dara.IsNil(request.TaskName) {
		body["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TaskParam) {
		body["TaskParam"] = request.TaskParam
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitAsyncTask"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitAsyncTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 妙笔为您提供了与公有云“智能审校”模块中相同的上传自定义规则库的功能。由于鉴权限制，用户需要开通阿里云 OSS 服务后，将自定义规则库文件上传到 OSS 中，再使用该文件的 fileKey 作为入参才能顺利调用本接口。该接口在被调用后，会对用户的自定义规则库进行结构化处理，并生成一个 xlsx 格式的结构化解析结果。您可以调用 GetAuditNoteProcessingStatus 接口查询结构化处理状态，也可以调用 DownloadAuditNote 接口获取结构化之后的规则库。接口功能正在迭代中，预计会在未来使用可访问的文件 URL 作为入参。
//
// @param request - SubmitAuditNoteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitAuditNoteResponse
func (client *Client) SubmitAuditNoteWithContext(ctx context.Context, request *SubmitAuditNoteRequest, runtime *dara.RuntimeOptions) (_result *SubmitAuditNoteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileKey) {
		body["FileKey"] = request.FileKey
	}

	if !dara.IsNil(request.NoteId) {
		body["NoteId"] = request.NoteId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitAuditNote"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitAuditNoteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交审核任务
//
// @param request - SubmitAuditTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitAuditTaskResponse
func (client *Client) SubmitAuditTaskWithContext(ctx context.Context, request *SubmitAuditTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitAuditTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ArticleId) {
		body["ArticleId"] = request.ArticleId
	}

	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.HtmlContent) {
		body["HtmlContent"] = request.HtmlContent
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitAuditTask"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitAuditTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交自定义播报单任务
//
// @param tmpReq - SubmitCustomHotTopicBroadcastJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitCustomHotTopicBroadcastJobResponse
func (client *Client) SubmitCustomHotTopicBroadcastJobWithContext(ctx context.Context, tmpReq *SubmitCustomHotTopicBroadcastJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitCustomHotTopicBroadcastJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SubmitCustomHotTopicBroadcastJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.HotTopicBroadcastConfig) {
		request.HotTopicBroadcastConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HotTopicBroadcastConfig, dara.String("HotTopicBroadcastConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Topics) {
		request.TopicsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Topics, dara.String("Topics"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.HotTopicBroadcastConfigShrink) {
		body["HotTopicBroadcastConfig"] = request.HotTopicBroadcastConfigShrink
	}

	if !dara.IsNil(request.HotTopicVersion) {
		body["HotTopicVersion"] = request.HotTopicVersion
	}

	if !dara.IsNil(request.TopicsShrink) {
		body["Topics"] = request.TopicsShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitCustomHotTopicBroadcastJob"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitCustomHotTopicBroadcastJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 从自定义数据源提交选题热点分析
//
// @param tmpReq - SubmitCustomSourceTopicAnalysisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitCustomSourceTopicAnalysisResponse
func (client *Client) SubmitCustomSourceTopicAnalysisWithContext(ctx context.Context, tmpReq *SubmitCustomSourceTopicAnalysisRequest, runtime *dara.RuntimeOptions) (_result *SubmitCustomSourceTopicAnalysisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SubmitCustomSourceTopicAnalysisShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AnalysisTypes) {
		request.AnalysisTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AnalysisTypes, dara.String("AnalysisTypes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.News) {
		request.NewsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.News, dara.String("News"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AnalysisTypesShrink) {
		body["AnalysisTypes"] = request.AnalysisTypesShrink
	}

	if !dara.IsNil(request.FileType) {
		body["FileType"] = request.FileType
	}

	if !dara.IsNil(request.FileUrl) {
		body["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.MaxTopicSize) {
		body["MaxTopicSize"] = request.MaxTopicSize
	}

	if !dara.IsNil(request.NewsShrink) {
		body["News"] = request.NewsShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitCustomSourceTopicAnalysis"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitCustomSourceTopicAnalysisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交自定义热点选题视角分析任务
//
// @param tmpReq - SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse
func (client *Client) SubmitCustomTopicSelectionPerspectiveAnalysisTaskWithContext(ctx context.Context, tmpReq *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SubmitCustomTopicSelectionPerspectiveAnalysisTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Documents) {
		request.DocumentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Documents, dara.String("Documents"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DocumentsShrink) {
		body["Documents"] = request.DocumentsShrink
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.Topic) {
		body["Topic"] = request.Topic
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitCustomTopicSelectionPerspectiveAnalysisTask"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交深度写作任务
//
// @param tmpReq - SubmitDeepWriteTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitDeepWriteTaskResponse
func (client *Client) SubmitDeepWriteTaskWithContext(ctx context.Context, tmpReq *SubmitDeepWriteTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitDeepWriteTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SubmitDeepWriteTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AgentOrchestration) {
		request.AgentOrchestrationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AgentOrchestration, dara.String("AgentOrchestration"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Files) {
		request.FilesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Files, dara.String("Files"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentOrchestrationShrink) {
		query["AgentOrchestration"] = request.AgentOrchestrationShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FilesShrink) {
		body["Files"] = request.FilesShrink
	}

	if !dara.IsNil(request.Input) {
		body["Input"] = request.Input
	}

	if !dara.IsNil(request.Instructions) {
		body["Instructions"] = request.Instructions
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitDeepWriteTask"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitDeepWriteTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交文档聚合任务
//
// @param tmpReq - SubmitDocClusterTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitDocClusterTaskResponse
func (client *Client) SubmitDocClusterTaskWithContext(ctx context.Context, tmpReq *SubmitDocClusterTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitDocClusterTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SubmitDocClusterTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Documents) {
		request.DocumentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Documents, dara.String("Documents"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DocumentsShrink) {
		body["Documents"] = request.DocumentsShrink
	}

	if !dara.IsNil(request.SummaryLength) {
		body["SummaryLength"] = request.SummaryLength
	}

	if !dara.IsNil(request.TitleLength) {
		body["TitleLength"] = request.TitleLength
	}

	if !dara.IsNil(request.TopicCount) {
		body["TopicCount"] = request.TopicCount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitDocClusterTask"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitDocClusterTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交VOC异步任务
//
// @param tmpReq - SubmitEnterpriseVocAnalysisTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitEnterpriseVocAnalysisTaskResponse
func (client *Client) SubmitEnterpriseVocAnalysisTaskWithContext(ctx context.Context, tmpReq *SubmitEnterpriseVocAnalysisTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitEnterpriseVocAnalysisTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SubmitEnterpriseVocAnalysisTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ContentTags) {
		request.ContentTagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ContentTags, dara.String("ContentTags"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Contents) {
		request.ContentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Contents, dara.String("Contents"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.FilterTags) {
		request.FilterTagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterTags, dara.String("FilterTags"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		body["ApiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.ContentTagsShrink) {
		body["ContentTags"] = request.ContentTagsShrink
	}

	if !dara.IsNil(request.ContentsShrink) {
		body["Contents"] = request.ContentsShrink
	}

	if !dara.IsNil(request.FileKey) {
		body["FileKey"] = request.FileKey
	}

	if !dara.IsNil(request.FilterTagsShrink) {
		body["FilterTags"] = request.FilterTagsShrink
	}

	if !dara.IsNil(request.MaterialType) {
		body["MaterialType"] = request.MaterialType
	}

	if !dara.IsNil(request.ModelId) {
		body["ModelId"] = request.ModelId
	}

	if !dara.IsNil(request.PositiveSample) {
		body["PositiveSample"] = request.PositiveSample
	}

	if !dara.IsNil(request.PositiveSampleFileKey) {
		body["PositiveSampleFileKey"] = request.PositiveSampleFileKey
	}

	if !dara.IsNil(request.TaskType) {
		body["TaskType"] = request.TaskType
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitEnterpriseVocAnalysisTask"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitEnterpriseVocAnalysisTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出词库任务
//
// @param request - SubmitExportTermsTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitExportTermsTaskResponse
func (client *Client) SubmitExportTermsTaskWithContext(ctx context.Context, request *SubmitExportTermsTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitExportTermsTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TermsName) {
		body["TermsName"] = request.TermsName
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitExportTermsTask"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitExportTermsTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 妙笔为您提供了新的事实性审核能力，在联网搜索并判断正误的前提下，还支持用户自定义配置搜索来源 URL。
//
// @param request - SubmitFactAuditUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitFactAuditUrlResponse
func (client *Client) SubmitFactAuditUrlWithContext(ctx context.Context, request *SubmitFactAuditUrlRequest, runtime *dara.RuntimeOptions) (_result *SubmitFactAuditUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Url) {
		body["Url"] = request.Url
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitFactAuditUrl"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitFactAuditUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交导入自定义词库任务
//
// @param request - SubmitImportTermsTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitImportTermsTaskResponse
func (client *Client) SubmitImportTermsTaskWithContext(ctx context.Context, request *SubmitImportTermsTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitImportTermsTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileKey) {
		body["FileKey"] = request.FileKey
	}

	if !dara.IsNil(request.TermsName) {
		body["TermsName"] = request.TermsName
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitImportTermsTask"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitImportTermsTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交智能审核
//
// @param tmpReq - SubmitSmartAuditRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitSmartAuditResponse
func (client *Client) SubmitSmartAuditWithContext(ctx context.Context, tmpReq *SubmitSmartAuditRequest, runtime *dara.RuntimeOptions) (_result *SubmitSmartAuditResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SubmitSmartAuditShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ImageUrlList) {
		request.ImageUrlListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImageUrlList, dara.String("ImageUrlList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SubCodes) {
		request.SubCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SubCodes, dara.String("SubCodes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ImageUrls) {
		request.ImageUrlsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImageUrls, dara.String("imageUrls"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageUrlListShrink) {
		body["ImageUrlList"] = request.ImageUrlListShrink
	}

	if !dara.IsNil(request.NoteId) {
		body["NoteId"] = request.NoteId
	}

	if !dara.IsNil(request.SubCodesShrink) {
		body["SubCodes"] = request.SubCodesShrink
	}

	if !dara.IsNil(request.TermsName) {
		body["TermsName"] = request.TermsName
	}

	if !dara.IsNil(request.Text) {
		body["Text"] = request.Text
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	if !dara.IsNil(request.ImageUrlsShrink) {
		body["imageUrls"] = request.ImageUrlsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitSmartAudit"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitSmartAuditResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交一键成片剪辑任务
//
// @param tmpReq - SubmitSmartClipTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitSmartClipTaskResponse
func (client *Client) SubmitSmartClipTaskWithContext(ctx context.Context, tmpReq *SubmitSmartClipTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitSmartClipTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SubmitSmartClipTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EditingConfig) {
		request.EditingConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EditingConfig, dara.String("EditingConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.InputConfig) {
		request.InputConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InputConfig, dara.String("InputConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OutputConfig) {
		request.OutputConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OutputConfig, dara.String("OutputConfig"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EditingConfigShrink) {
		body["EditingConfig"] = request.EditingConfigShrink
	}

	if !dara.IsNil(request.ExtendParam) {
		body["ExtendParam"] = request.ExtendParam
	}

	if !dara.IsNil(request.InputConfigShrink) {
		body["InputConfig"] = request.InputConfigShrink
	}

	if !dara.IsNil(request.OutputConfigShrink) {
		body["OutputConfig"] = request.OutputConfigShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitSmartClipTask"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitSmartClipTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交选题热点分析任务
//
// @param tmpReq - SubmitTopicSelectionPerspectiveAnalysisTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitTopicSelectionPerspectiveAnalysisTaskResponse
func (client *Client) SubmitTopicSelectionPerspectiveAnalysisTaskWithContext(ctx context.Context, tmpReq *SubmitTopicSelectionPerspectiveAnalysisTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitTopicSelectionPerspectiveAnalysisTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SubmitTopicSelectionPerspectiveAnalysisTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Documents) {
		request.DocumentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Documents, dara.String("Documents"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PerspectiveTypes) {
		request.PerspectiveTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PerspectiveTypes, dara.String("PerspectiveTypes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DocumentsShrink) {
		body["Documents"] = request.DocumentsShrink
	}

	if !dara.IsNil(request.PerspectiveTypesShrink) {
		body["PerspectiveTypes"] = request.PerspectiveTypesShrink
	}

	if !dara.IsNil(request.Topic) {
		body["Topic"] = request.Topic
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitTopicSelectionPerspectiveAnalysisTask"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitTopicSelectionPerspectiveAnalysisTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新自定义文本
//
// @param request - UpdateCustomTextRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCustomTextResponse
func (client *Client) UpdateCustomTextWithContext(ctx context.Context, request *UpdateCustomTextRequest, runtime *dara.RuntimeOptions) (_result *UpdateCustomTextResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CommodityCode) {
		body["CommodityCode"] = request.CommodityCode
	}

	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCustomText"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCustomTextResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 数据集管理-更新
//
// @param tmpReq - UpdateDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDatasetResponse
func (client *Client) UpdateDatasetWithContext(ctx context.Context, tmpReq *UpdateDatasetRequest, runtime *dara.RuntimeOptions) (_result *UpdateDatasetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDatasetShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DatasetConfig) {
		request.DatasetConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DatasetConfig, dara.String("DatasetConfig"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DatasetConfigShrink) {
		body["DatasetConfig"] = request.DatasetConfigShrink
	}

	if !dara.IsNil(request.DatasetDescription) {
		body["DatasetDescription"] = request.DatasetDescription
	}

	if !dara.IsNil(request.DatasetId) {
		body["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.SearchDatasetEnable) {
		body["SearchDatasetEnable"] = request.SearchDatasetEnable
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataset"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDatasetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改数据集文档
//
// @param tmpReq - UpdateDatasetDocumentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDatasetDocumentResponse
func (client *Client) UpdateDatasetDocumentWithContext(ctx context.Context, tmpReq *UpdateDatasetDocumentRequest, runtime *dara.RuntimeOptions) (_result *UpdateDatasetDocumentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDatasetDocumentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Document) {
		request.DocumentShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Document, dara.String("Document"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DatasetId) {
		body["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.DatasetName) {
		body["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.DocumentShrink) {
		body["Document"] = request.DocumentShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDatasetDocument"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDatasetDocumentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通用配置-更新
//
// @param request - UpdateGeneralConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGeneralConfigResponse
func (client *Client) UpdateGeneralConfigWithContext(ctx context.Context, request *UpdateGeneralConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateGeneralConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigKey) {
		body["ConfigKey"] = request.ConfigKey
	}

	if !dara.IsNil(request.ConfigValue) {
		body["ConfigValue"] = request.ConfigValue
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGeneralConfig"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGeneralConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档管理-更新。
//
// @param tmpReq - UpdateGeneratedContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGeneratedContentResponse
func (client *Client) UpdateGeneratedContentWithContext(ctx context.Context, tmpReq *UpdateGeneratedContentRequest, runtime *dara.RuntimeOptions) (_result *UpdateGeneratedContentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateGeneratedContentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Keywords) {
		request.KeywordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Keywords, dara.String("Keywords"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.ContentText) {
		body["ContentText"] = request.ContentText
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.KeywordsShrink) {
		body["Keywords"] = request.KeywordsShrink
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGeneratedContent"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGeneratedContentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据ID更新素材
//
// @param tmpReq - UpdateMaterialDocumentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMaterialDocumentResponse
func (client *Client) UpdateMaterialDocumentWithContext(ctx context.Context, tmpReq *UpdateMaterialDocumentRequest, runtime *dara.RuntimeOptions) (_result *UpdateMaterialDocumentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateMaterialDocumentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DocKeywords) {
		request.DocKeywordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DocKeywords, dara.String("DocKeywords"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Author) {
		body["Author"] = request.Author
	}

	if !dara.IsNil(request.DocKeywordsShrink) {
		body["DocKeywords"] = request.DocKeywordsShrink
	}

	if !dara.IsNil(request.DocType) {
		body["DocType"] = request.DocType
	}

	if !dara.IsNil(request.ExternalUrl) {
		body["ExternalUrl"] = request.ExternalUrl
	}

	if !dara.IsNil(request.HtmlContent) {
		body["HtmlContent"] = request.HtmlContent
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.PubTime) {
		body["PubTime"] = request.PubTime
	}

	if !dara.IsNil(request.ShareAttr) {
		body["ShareAttr"] = request.ShareAttr
	}

	if !dara.IsNil(request.SrcFrom) {
		body["SrcFrom"] = request.SrcFrom
	}

	if !dara.IsNil(request.Summary) {
		body["Summary"] = request.Summary
	}

	if !dara.IsNil(request.TextContent) {
		body["TextContent"] = request.TextContent
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	if !dara.IsNil(request.Url) {
		body["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMaterialDocument"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMaterialDocumentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 妙读上传书籍
//
// @param tmpReq - UploadBookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadBookResponse
func (client *Client) UploadBookWithContext(ctx context.Context, tmpReq *UploadBookRequest, runtime *dara.RuntimeOptions) (_result *UploadBookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UploadBookShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Docs) {
		request.DocsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Docs, dara.String("Docs"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		body["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.DocsShrink) {
		body["Docs"] = request.DocsShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadBook"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadBookResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 妙读上传文档接口
//
// @param tmpReq - UploadDocRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadDocResponse
func (client *Client) UploadDocWithContext(ctx context.Context, tmpReq *UploadDocRequest, runtime *dara.RuntimeOptions) (_result *UploadDocResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UploadDocShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Docs) {
		request.DocsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Docs, dara.String("Docs"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		body["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.DocsShrink) {
		body["Docs"] = request.DocsShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadDoc"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadDocResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 校验企业VOC上传模板
//
// @param request - ValidateUploadTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateUploadTemplateResponse
func (client *Client) ValidateUploadTemplateWithContext(ctx context.Context, request *ValidateUploadTemplateRequest, runtime *dara.RuntimeOptions) (_result *ValidateUploadTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileKey) {
		body["FileKey"] = request.FileKey
	}

	if !dara.IsNil(request.TaskType) {
		body["TaskType"] = request.TaskType
	}

	if !dara.IsNil(request.TemplateType) {
		body["TemplateType"] = request.TemplateType
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ValidateUploadTemplate"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ValidateUploadTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) runAbbreviationContentWithSSECtx_opYieldFunc(_yield chan *RunAbbreviationContentResponse, _yieldErr chan error, ctx context.Context, request *RunAbbreviationContentRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunAbbreviationContent"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runAiHelperWritingWithSSECtx_opYieldFunc(_yield chan *RunAiHelperWritingResponse, _yieldErr chan error, ctx context.Context, tmpReq *RunAiHelperWritingRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &RunAiHelperWritingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.WritingParams) {
		request.WritingParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WritingParams, dara.String("WritingParams"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DistributeWriting) {
		body["DistributeWriting"] = request.DistributeWriting
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.PromptMode) {
		body["PromptMode"] = request.PromptMode
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	if !dara.IsNil(request.WritingParamsShrink) {
		body["WritingParams"] = request.WritingParamsShrink
	}

	if !dara.IsNil(request.WritingScene) {
		body["WritingScene"] = request.WritingScene
	}

	if !dara.IsNil(request.WritingStyle) {
		body["WritingStyle"] = request.WritingStyle
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunAiHelperWriting"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runBookBrainmapWithSSECtx_opYieldFunc(_yield chan *RunBookBrainmapResponse, _yieldErr chan error, ctx context.Context, request *RunBookBrainmapRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CleanCache) {
		body["CleanCache"] = request.CleanCache
	}

	if !dara.IsNil(request.DocId) {
		body["DocId"] = request.DocId
	}

	if !dara.IsNil(request.NodeNumber) {
		body["NodeNumber"] = request.NodeNumber
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.WordNumber) {
		body["WordNumber"] = request.WordNumber
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunBookBrainmap"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runBookIntroductionWithSSECtx_opYieldFunc(_yield chan *RunBookIntroductionResponse, _yieldErr chan error, ctx context.Context, request *RunBookIntroductionRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocId) {
		body["DocId"] = request.DocId
	}

	if !dara.IsNil(request.KeyPointPrompt) {
		body["KeyPointPrompt"] = request.KeyPointPrompt
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.SummaryPrompt) {
		body["SummaryPrompt"] = request.SummaryPrompt
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunBookIntroduction"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runBookSmartCardWithSSECtx_opYieldFunc(_yield chan *RunBookSmartCardResponse, _yieldErr chan error, ctx context.Context, request *RunBookSmartCardRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocId) {
		body["DocId"] = request.DocId
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunBookSmartCard"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runCommentGenerationWithSSECtx_opYieldFunc(_yield chan *RunCommentGenerationResponse, _yieldErr chan error, ctx context.Context, tmpReq *RunCommentGenerationRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &RunCommentGenerationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LengthRange) {
		request.LengthRangeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LengthRange, dara.String("LengthRange"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Sentiment) {
		request.SentimentShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sentiment, dara.String("Sentiment"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Type) {
		request.TypeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Type, dara.String("Type"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AllowEmoji) {
		body["AllowEmoji"] = request.AllowEmoji
	}

	if !dara.IsNil(request.ExtraInfo) {
		body["ExtraInfo"] = request.ExtraInfo
	}

	if !dara.IsNil(request.Length) {
		body["Length"] = request.Length
	}

	if !dara.IsNil(request.LengthRangeShrink) {
		body["LengthRange"] = request.LengthRangeShrink
	}

	if !dara.IsNil(request.ModelId) {
		body["ModelId"] = request.ModelId
	}

	if !dara.IsNil(request.NumComments) {
		body["NumComments"] = request.NumComments
	}

	if !dara.IsNil(request.SentimentShrink) {
		body["Sentiment"] = request.SentimentShrink
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.SourceMaterial) {
		body["SourceMaterial"] = request.SourceMaterial
	}

	if !dara.IsNil(request.Style) {
		body["Style"] = request.Style
	}

	if !dara.IsNil(request.TypeShrink) {
		body["Type"] = request.TypeShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunCommentGeneration"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runContinueContentWithSSECtx_opYieldFunc(_yield chan *RunContinueContentResponse, _yieldErr chan error, ctx context.Context, request *RunContinueContentRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunContinueContent"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runCustomHotTopicAnalysisWithSSECtx_opYieldFunc(_yield chan *RunCustomHotTopicAnalysisResponse, _yieldErr chan error, ctx context.Context, request *RunCustomHotTopicAnalysisRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AskUser) {
		body["AskUser"] = request.AskUser
	}

	if !dara.IsNil(request.ForceAnalysisExistsTopic) {
		body["ForceAnalysisExistsTopic"] = request.ForceAnalysisExistsTopic
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.UserBack) {
		body["UserBack"] = request.UserBack
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunCustomHotTopicAnalysis"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runCustomHotTopicViewPointAnalysisWithSSECtx_opYieldFunc(_yield chan *RunCustomHotTopicViewPointAnalysisResponse, _yieldErr chan error, ctx context.Context, request *RunCustomHotTopicViewPointAnalysisRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AskUser) {
		body["AskUser"] = request.AskUser
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.SearchQuery) {
		body["SearchQuery"] = request.SearchQuery
	}

	if !dara.IsNil(request.SkipAskUser) {
		body["SkipAskUser"] = request.SkipAskUser
	}

	if !dara.IsNil(request.Topic) {
		body["Topic"] = request.Topic
	}

	if !dara.IsNil(request.TopicId) {
		body["TopicId"] = request.TopicId
	}

	if !dara.IsNil(request.TopicSource) {
		body["TopicSource"] = request.TopicSource
	}

	if !dara.IsNil(request.TopicVersion) {
		body["TopicVersion"] = request.TopicVersion
	}

	if !dara.IsNil(request.UserBack) {
		body["UserBack"] = request.UserBack
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunCustomHotTopicViewPointAnalysis"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runDeepWritingWithSSECtx_opYieldFunc(_yield chan *RunDeepWritingResponse, _yieldErr chan error, ctx context.Context, request *RunDeepWritingRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Cursor) {
		body["Cursor"] = request.Cursor
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunDeepWriting"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runDocBrainmapWithSSECtx_opYieldFunc(_yield chan *RunDocBrainmapResponse, _yieldErr chan error, ctx context.Context, request *RunDocBrainmapRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CleanCache) {
		body["CleanCache"] = request.CleanCache
	}

	if !dara.IsNil(request.DocId) {
		body["DocId"] = request.DocId
	}

	if !dara.IsNil(request.ModelName) {
		body["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.NodeNumber) {
		body["NodeNumber"] = request.NodeNumber
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.WordNumber) {
		body["WordNumber"] = request.WordNumber
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	if !dara.IsNil(request.ReferenceContent) {
		body["referenceContent"] = request.ReferenceContent
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunDocBrainmap"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runDocIntroductionWithSSECtx_opYieldFunc(_yield chan *RunDocIntroductionResponse, _yieldErr chan error, ctx context.Context, request *RunDocIntroductionRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CleanCache) {
		body["CleanCache"] = request.CleanCache
	}

	if !dara.IsNil(request.DocId) {
		body["DocId"] = request.DocId
	}

	if !dara.IsNil(request.IntroductionPrompt) {
		body["IntroductionPrompt"] = request.IntroductionPrompt
	}

	if !dara.IsNil(request.KeyPointPrompt) {
		body["KeyPointPrompt"] = request.KeyPointPrompt
	}

	if !dara.IsNil(request.ModelName) {
		body["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.SummaryPrompt) {
		body["SummaryPrompt"] = request.SummaryPrompt
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	if !dara.IsNil(request.ReferenceContent) {
		body["referenceContent"] = request.ReferenceContent
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunDocIntroduction"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runDocQaWithSSECtx_opYieldFunc(_yield chan *RunDocQaResponse, _yieldErr chan error, ctx context.Context, tmpReq *RunDocQaRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &RunDocQaShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CategoryIds) {
		request.CategoryIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CategoryIds, dara.String("CategoryIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ConversationContexts) {
		request.ConversationContextsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ConversationContexts, dara.String("ConversationContexts"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DocIds) {
		request.DocIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DocIds, dara.String("DocIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoryIdsShrink) {
		body["CategoryIds"] = request.CategoryIdsShrink
	}

	if !dara.IsNil(request.ConversationContextsShrink) {
		body["ConversationContexts"] = request.ConversationContextsShrink
	}

	if !dara.IsNil(request.DocIdsShrink) {
		body["DocIds"] = request.DocIdsShrink
	}

	if !dara.IsNil(request.ModelName) {
		body["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.Query) {
		body["Query"] = request.Query
	}

	if !dara.IsNil(request.ReferenceContent) {
		body["ReferenceContent"] = request.ReferenceContent
	}

	if !dara.IsNil(request.SearchSource) {
		body["SearchSource"] = request.SearchSource
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunDocQa"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runDocSmartCardWithSSECtx_opYieldFunc(_yield chan *RunDocSmartCardResponse, _yieldErr chan error, ctx context.Context, request *RunDocSmartCardRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocId) {
		body["DocId"] = request.DocId
	}

	if !dara.IsNil(request.ModelName) {
		body["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunDocSmartCard"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runDocSummaryWithSSECtx_opYieldFunc(_yield chan *RunDocSummaryResponse, _yieldErr chan error, ctx context.Context, request *RunDocSummaryRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CleanCache) {
		body["CleanCache"] = request.CleanCache
	}

	if !dara.IsNil(request.DocId) {
		body["DocId"] = request.DocId
	}

	if !dara.IsNil(request.ModelName) {
		body["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.Query) {
		body["Query"] = request.Query
	}

	if !dara.IsNil(request.RecommendContent) {
		body["RecommendContent"] = request.RecommendContent
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunDocSummary"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runDocTranslationWithSSECtx_opYieldFunc(_yield chan *RunDocTranslationResponse, _yieldErr chan error, ctx context.Context, request *RunDocTranslationRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CleanCache) {
		body["CleanCache"] = request.CleanCache
	}

	if !dara.IsNil(request.DocId) {
		body["DocId"] = request.DocId
	}

	if !dara.IsNil(request.ModelName) {
		body["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.RecommendContent) {
		body["RecommendContent"] = request.RecommendContent
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.TransType) {
		body["TransType"] = request.TransType
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunDocTranslation"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runDocWashingWithSSECtx_opYieldFunc(_yield chan *RunDocWashingResponse, _yieldErr chan error, ctx context.Context, request *RunDocWashingRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ModelId) {
		body["ModelId"] = request.ModelId
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.ReferenceContent) {
		body["ReferenceContent"] = request.ReferenceContent
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Topic) {
		body["Topic"] = request.Topic
	}

	if !dara.IsNil(request.WordNumber) {
		body["WordNumber"] = request.WordNumber
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	if !dara.IsNil(request.WritingTypeName) {
		body["WritingTypeName"] = request.WritingTypeName
	}

	if !dara.IsNil(request.WritingTypeRefDoc) {
		body["WritingTypeRefDoc"] = request.WritingTypeRefDoc
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunDocWashing"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runExpandContentWithSSECtx_opYieldFunc(_yield chan *RunExpandContentResponse, _yieldErr chan error, ctx context.Context, request *RunExpandContentRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunExpandContent"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runGenerateQuestionsWithSSECtx_opYieldFunc(_yield chan *RunGenerateQuestionsResponse, _yieldErr chan error, ctx context.Context, request *RunGenerateQuestionsRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocId) {
		body["DocId"] = request.DocId
	}

	if !dara.IsNil(request.ModelName) {
		body["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.ReferenceContent) {
		body["ReferenceContent"] = request.ReferenceContent
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunGenerateQuestions"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runHotwordWithSSECtx_opYieldFunc(_yield chan *RunHotwordResponse, _yieldErr chan error, ctx context.Context, request *RunHotwordRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocId) {
		body["DocId"] = request.DocId
	}

	if !dara.IsNil(request.ModelName) {
		body["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.ReferenceContent) {
		body["ReferenceContent"] = request.ReferenceContent
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunHotword"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runKeywordsExtractionGenerationWithSSECtx_opYieldFunc(_yield chan *RunKeywordsExtractionGenerationResponse, _yieldErr chan error, ctx context.Context, tmpReq *RunKeywordsExtractionGenerationRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &RunKeywordsExtractionGenerationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ReferenceData) {
		request.ReferenceDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReferenceData, dara.String("ReferenceData"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.ReferenceDataShrink) {
		body["ReferenceData"] = request.ReferenceDataShrink
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunKeywordsExtractionGeneration"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runMultiDocIntroductionWithSSECtx_opYieldFunc(_yield chan *RunMultiDocIntroductionResponse, _yieldErr chan error, ctx context.Context, tmpReq *RunMultiDocIntroductionRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &RunMultiDocIntroductionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DocIds) {
		request.DocIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DocIds, dara.String("DocIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DocIdsShrink) {
		body["DocIds"] = request.DocIdsShrink
	}

	if !dara.IsNil(request.KeyPointPrompt) {
		body["KeyPointPrompt"] = request.KeyPointPrompt
	}

	if !dara.IsNil(request.ModelName) {
		body["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.SummaryPrompt) {
		body["SummaryPrompt"] = request.SummaryPrompt
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunMultiDocIntroduction"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runPptOutlineGenerationWithSSECtx_opYieldFunc(_yield chan *RunPptOutlineGenerationResponse, _yieldErr chan error, ctx context.Context, request *RunPptOutlineGenerationRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunPptOutlineGeneration"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runQuickWritingWithSSECtx_opYieldFunc(_yield chan *RunQuickWritingResponse, _yieldErr chan error, ctx context.Context, tmpReq *RunQuickWritingRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &RunQuickWritingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Articles) {
		request.ArticlesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Articles, dara.String("Articles"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SearchSources) {
		request.SearchSourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SearchSources, dara.String("SearchSources"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ArticlesShrink) {
		body["Articles"] = request.ArticlesShrink
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.SearchSourcesShrink) {
		body["SearchSources"] = request.SearchSourcesShrink
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunQuickWriting"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runSearchGenerationWithSSECtx_opYieldFunc(_yield chan *RunSearchGenerationResponse, _yieldErr chan error, ctx context.Context, tmpReq *RunSearchGenerationRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &RunSearchGenerationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AgentContext) {
		request.AgentContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AgentContext, dara.String("AgentContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ChatConfig) {
		request.ChatConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChatConfig, dara.String("ChatConfig"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentContextShrink) {
		body["AgentContext"] = request.AgentContextShrink
	}

	if !dara.IsNil(request.ChatConfigShrink) {
		body["ChatConfig"] = request.ChatConfigShrink
	}

	if !dara.IsNil(request.FileUrl) {
		body["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.ModelId) {
		body["ModelId"] = request.ModelId
	}

	if !dara.IsNil(request.OriginalSessionId) {
		body["OriginalSessionId"] = request.OriginalSessionId
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunSearchGeneration"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runSearchSimilarArticlesWithSSECtx_opYieldFunc(_yield chan *RunSearchSimilarArticlesResponse, _yieldErr chan error, ctx context.Context, tmpReq *RunSearchSimilarArticlesRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &RunSearchSimilarArticlesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ChatConfig) {
		request.ChatConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChatConfig, dara.String("ChatConfig"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ChatConfigShrink) {
		body["ChatConfig"] = request.ChatConfigShrink
	}

	if !dara.IsNil(request.DocType) {
		body["DocType"] = request.DocType
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	if !dara.IsNil(request.Url) {
		body["Url"] = request.Url
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunSearchSimilarArticles"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runStepByStepWritingWithSSECtx_opYieldFunc(_yield chan *RunStepByStepWritingResponse, _yieldErr chan error, ctx context.Context, tmpReq *RunStepByStepWritingRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &RunStepByStepWritingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ReferenceData) {
		request.ReferenceDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReferenceData, dara.String("ReferenceData"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.WritingConfig) {
		request.WritingConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WritingConfig, dara.String("WritingConfig"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OriginSessionId) {
		body["OriginSessionId"] = request.OriginSessionId
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.ReferenceDataShrink) {
		body["ReferenceData"] = request.ReferenceDataShrink
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	if !dara.IsNil(request.WritingConfigShrink) {
		body["WritingConfig"] = request.WritingConfigShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunStepByStepWriting"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runStyleFeatureAnalysisWithSSECtx_opYieldFunc(_yield chan *RunStyleFeatureAnalysisResponse, _yieldErr chan error, ctx context.Context, tmpReq *RunStyleFeatureAnalysisRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &RunStyleFeatureAnalysisShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Contents) {
		request.ContentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Contents, dara.String("Contents"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.MaterialIds) {
		request.MaterialIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MaterialIds, dara.String("MaterialIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ContentsShrink) {
		body["Contents"] = request.ContentsShrink
	}

	if !dara.IsNil(request.MaterialIdsShrink) {
		body["MaterialIds"] = request.MaterialIdsShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunStyleFeatureAnalysis"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runSummaryGenerateWithSSECtx_opYieldFunc(_yield chan *RunSummaryGenerateResponse, _yieldErr chan error, ctx context.Context, request *RunSummaryGenerateRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunSummaryGenerate"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runTextPolishingWithSSECtx_opYieldFunc(_yield chan *RunTextPolishingResponse, _yieldErr chan error, ctx context.Context, request *RunTextPolishingRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.OriginContent) {
		body["OriginContent"] = request.OriginContent
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunTextPolishing"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runTitleGenerationWithSSECtx_opYieldFunc(_yield chan *RunTitleGenerationResponse, _yieldErr chan error, ctx context.Context, tmpReq *RunTitleGenerationRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &RunTitleGenerationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeduplicatedTitles) {
		request.DeduplicatedTitlesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeduplicatedTitles, dara.String("DeduplicatedTitles"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ReferenceData) {
		request.ReferenceDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReferenceData, dara.String("ReferenceData"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeduplicatedTitlesShrink) {
		body["DeduplicatedTitles"] = request.DeduplicatedTitlesShrink
	}

	if !dara.IsNil(request.ReferenceDataShrink) {
		body["ReferenceData"] = request.ReferenceDataShrink
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TitleCount) {
		body["TitleCount"] = request.TitleCount
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunTitleGeneration"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runTopicSelectionMergeWithSSECtx_opYieldFunc(_yield chan *RunTopicSelectionMergeResponse, _yieldErr chan error, ctx context.Context, tmpReq *RunTopicSelectionMergeRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &RunTopicSelectionMergeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Topics) {
		request.TopicsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Topics, dara.String("Topics"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.TopicsShrink) {
		body["Topics"] = request.TopicsShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunTopicSelectionMerge"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runTranslateGenerationWithSSECtx_opYieldFunc(_yield chan *RunTranslateGenerationResponse, _yieldErr chan error, ctx context.Context, tmpReq *RunTranslateGenerationRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &RunTranslateGenerationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ReferenceData) {
		request.ReferenceDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReferenceData, dara.String("ReferenceData"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.ReferenceDataShrink) {
		body["ReferenceData"] = request.ReferenceDataShrink
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunTranslateGeneration"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runVideoScriptGenerateWithSSECtx_opYieldFunc(_yield chan *RunVideoScriptGenerateResponse, _yieldErr chan error, ctx context.Context, request *RunVideoScriptGenerateRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.ScriptLength) {
		body["ScriptLength"] = request.ScriptLength
	}

	if !dara.IsNil(request.ScriptNumber) {
		body["ScriptNumber"] = request.ScriptNumber
	}

	if !dara.IsNil(request.UseSearch) {
		body["UseSearch"] = request.UseSearch
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunVideoScriptGenerate"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runWriteToneGenerationWithSSECtx_opYieldFunc(_yield chan *RunWriteToneGenerationResponse, _yieldErr chan error, ctx context.Context, tmpReq *RunWriteToneGenerationRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &RunWriteToneGenerationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ReferenceData) {
		request.ReferenceDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReferenceData, dara.String("ReferenceData"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.ReferenceDataShrink) {
		body["ReferenceData"] = request.ReferenceDataShrink
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunWriteToneGeneration"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runWritingWithSSECtx_opYieldFunc(_yield chan *RunWritingResponse, _yieldErr chan error, ctx context.Context, tmpReq *RunWritingRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &RunWritingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ReferenceData) {
		request.ReferenceDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReferenceData, dara.String("ReferenceData"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.WritingConfig) {
		request.WritingConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WritingConfig, dara.String("WritingConfig"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OriginSessionId) {
		body["OriginSessionId"] = request.OriginSessionId
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.ReferenceDataShrink) {
		body["ReferenceData"] = request.ReferenceDataShrink
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	if !dara.IsNil(request.WritingConfigShrink) {
		body["WritingConfig"] = request.WritingConfigShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunWriting"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runWritingV2WithSSECtx_opYieldFunc(_yield chan *RunWritingV2Response, _yieldErr chan error, ctx context.Context, tmpReq *RunWritingV2Request, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &RunWritingV2ShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Articles) {
		request.ArticlesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Articles, dara.String("Articles"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Keywords) {
		request.KeywordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Keywords, dara.String("Keywords"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.MiniDocs) {
		request.MiniDocsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MiniDocs, dara.String("MiniDocs"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OutlineList) {
		request.OutlineListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OutlineList, dara.String("OutlineList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Outlines) {
		request.OutlinesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Outlines, dara.String("Outlines"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SearchSources) {
		request.SearchSourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SearchSources, dara.String("SearchSources"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Summarization) {
		request.SummarizationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Summarization, dara.String("Summarization"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.WritingParams) {
		request.WritingParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WritingParams, dara.String("WritingParams"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ArticlesShrink) {
		body["Articles"] = request.ArticlesShrink
	}

	if !dara.IsNil(request.DistributeWriting) {
		body["DistributeWriting"] = request.DistributeWriting
	}

	if !dara.IsNil(request.GcNumberSize) {
		body["GcNumberSize"] = request.GcNumberSize
	}

	if !dara.IsNil(request.GcNumberSizeTag) {
		body["GcNumberSizeTag"] = request.GcNumberSizeTag
	}

	if !dara.IsNil(request.KeywordsShrink) {
		body["Keywords"] = request.KeywordsShrink
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.MiniDocsShrink) {
		body["MiniDocs"] = request.MiniDocsShrink
	}

	if !dara.IsNil(request.OutlineListShrink) {
		body["OutlineList"] = request.OutlineListShrink
	}

	if !dara.IsNil(request.OutlinesShrink) {
		body["Outlines"] = request.OutlinesShrink
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.PromptMode) {
		body["PromptMode"] = request.PromptMode
	}

	if !dara.IsNil(request.SearchSourcesShrink) {
		body["SearchSources"] = request.SearchSourcesShrink
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.SourceTraceMethod) {
		body["SourceTraceMethod"] = request.SourceTraceMethod
	}

	if !dara.IsNil(request.Step) {
		body["Step"] = request.Step
	}

	if !dara.IsNil(request.SummarizationShrink) {
		body["Summarization"] = request.SummarizationShrink
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.UseSearch) {
		body["UseSearch"] = request.UseSearch
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	if !dara.IsNil(request.WritingParamsShrink) {
		body["WritingParams"] = request.WritingParamsShrink
	}

	if !dara.IsNil(request.WritingScene) {
		body["WritingScene"] = request.WritingScene
	}

	if !dara.IsNil(request.WritingStyle) {
		body["WritingStyle"] = request.WritingStyle
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunWritingV2"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}
