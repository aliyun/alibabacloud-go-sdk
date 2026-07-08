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
// Add a custom term to the audit dictionary.
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
// Adds a document to a data source.
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
// Generates a video clip.
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

	if !dara.IsNil(tmpReq.HighDefSourceVideos) {
		request.HighDefSourceVideosShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HighDefSourceVideos, dara.String("HighDefSourceVideos"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Stickers) {
		request.StickersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Stickers, dara.String("Stickers"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AdaptMode) {
		body["AdaptMode"] = request.AdaptMode
	}

	if !dara.IsNil(request.Alignment) {
		body["Alignment"] = request.Alignment
	}

	if !dara.IsNil(request.CloseMusic) {
		body["CloseMusic"] = request.CloseMusic
	}

	if !dara.IsNil(request.CloseSubtitle) {
		body["CloseSubtitle"] = request.CloseSubtitle
	}

	if !dara.IsNil(request.CloseVoice) {
		body["CloseVoice"] = request.CloseVoice
	}

	if !dara.IsNil(request.ClosingCreditsUrl) {
		body["ClosingCreditsUrl"] = request.ClosingCreditsUrl
	}

	if !dara.IsNil(request.ColorWordsShrink) {
		body["ColorWords"] = request.ColorWordsShrink
	}

	if !dara.IsNil(request.CosyVoiceAppKey) {
		body["CosyVoiceAppKey"] = request.CosyVoiceAppKey
	}

	if !dara.IsNil(request.CosyVoiceToken) {
		body["CosyVoiceToken"] = request.CosyVoiceToken
	}

	if !dara.IsNil(request.CustomVoiceStyle) {
		body["CustomVoiceStyle"] = request.CustomVoiceStyle
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

	if !dara.IsNil(request.HighDefSourceVideosShrink) {
		body["HighDefSourceVideos"] = request.HighDefSourceVideosShrink
	}

	if !dara.IsNil(request.MusicStyle) {
		body["MusicStyle"] = request.MusicStyle
	}

	if !dara.IsNil(request.MusicUrl) {
		body["MusicUrl"] = request.MusicUrl
	}

	if !dara.IsNil(request.MusicVolume) {
		body["MusicVolume"] = request.MusicVolume
	}

	if !dara.IsNil(request.OpeningCreditsUrl) {
		body["OpeningCreditsUrl"] = request.OpeningCreditsUrl
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

	if !dara.IsNil(request.TextWidth) {
		body["TextWidth"] = request.TextWidth
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
// Creates a timeline for smart video editing.
//
// @param tmpReq - AsyncCreateClipsTimeLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AsyncCreateClipsTimeLineResponse
func (client *Client) AsyncCreateClipsTimeLineWithContext(ctx context.Context, tmpReq *AsyncCreateClipsTimeLineRequest, runtime *dara.RuntimeOptions) (_result *AsyncCreateClipsTimeLineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AsyncCreateClipsTimeLineShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.HighLightConfig) {
		request.HighLightConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HighLightConfig, dara.String("HighLightConfig"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AdditionalContent) {
		body["AdditionalContent"] = request.AdditionalContent
	}

	if !dara.IsNil(request.CustomContent) {
		body["CustomContent"] = request.CustomContent
	}

	if !dara.IsNil(request.HighLightConfigShrink) {
		body["HighLightConfig"] = request.HighLightConfigShrink
	}

	if !dara.IsNil(request.NoRefVideo) {
		body["NoRefVideo"] = request.NoRefVideo
	}

	if !dara.IsNil(request.ProcessPrompt) {
		body["ProcessPrompt"] = request.ProcessPrompt
	}

	if !dara.IsNil(request.RecommendAudio) {
		body["RecommendAudio"] = request.RecommendAudio
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TimelineScene) {
		body["TimelineScene"] = request.TimelineScene
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
// Edits the timeline of a video editing task.
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
// Upload a tender document.
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
// Uploads video assets for editing.
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
	if !dara.IsNil(request.AdaptiveThreshold) {
		body["AdaptiveThreshold"] = request.AdaptiveThreshold
	}

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

	if !dara.IsNil(request.TaskName) {
		body["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TaskType) {
		body["TaskType"] = request.TaskType
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
// An API for writing bidding documents.
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
// Binds a PowerPoint (PPT) artifact.
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

	if !dara.IsNil(request.ExternalUserId) {
		body["ExternalUserId"] = request.ExternalUserId
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
// Cancels pending asynchronous tasks.
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
// Cancel an audit task.
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
// Cancels a deep writing task.
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
// Clears all intervention content.
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
// This interface confirms whether the parsing results from your custom rule library submission are used for audit tasks. Because parsing results may not meet your requirements, use this interface to perform a second confirmation. If you are satisfied with the parsing of your submitted rule library, provide the TaskId from that submission as an input parameter. The system then post-processes your uploaded rule library and makes it available for auditing. Otherwise, invoke the SubmitAuditNote interface again to upload the modified rule library.
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
// Batch add permissions:\\
//
// \\- Dataset permissions\\
//
// @param tmpReq - CreateDataPermissionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataPermissionsResponse
func (client *Client) CreateDataPermissionsWithContext(ctx context.Context, tmpReq *CreateDataPermissionsRequest, runtime *dara.RuntimeOptions) (_result *CreateDataPermissionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataPermissionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PermissionUserInfos) {
		request.PermissionUserInfosShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PermissionUserInfos, dara.String("PermissionUserInfos"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DataId) {
		body["DataId"] = request.DataId
	}

	if !dara.IsNil(request.DataType) {
		body["DataType"] = request.DataType
	}

	if !dara.IsNil(request.PermissionUserInfosShrink) {
		body["PermissionUserInfos"] = request.PermissionUserInfosShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataPermissions"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataPermissionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a dataset.
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
	if !dara.IsNil(request.AccessLevel) {
		body["AccessLevel"] = request.AccessLevel
	}

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
// # General configurations — Create
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
// Saves an article created in Miaobi. This operation supports rich text.
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
// Creates a temporary token for the online inference API.
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
// Deletes all custom rule libraries available for audit under the user account. This operation is irreversible. To archive rule libraries, use the DownloadAuditNote API to save them before deletion.
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
// Deletes specified term records.
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
// Deletes a custom text.
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
// Delete a custom hot spot event by topic name.
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
// Deletes a custom topic viewpoint by its ID.
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
// # Batch delete dataset permissions
//
// @param tmpReq - DeleteDataPermissionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataPermissionsResponse
func (client *Client) DeleteDataPermissionsWithContext(ctx context.Context, tmpReq *DeleteDataPermissionsRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataPermissionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteDataPermissionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.IdsShrink) {
		body["Ids"] = request.IdsShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataPermissions"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataPermissionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a dataset from the data source.
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
// Delete source documents.
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
// Deletes multiple documents.
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
// Deletes the specified URL used for factuality audit.
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
// Deletes general configurations.
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
// Deletes an article created in MiaoBi.
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
// Deletes an intervention rule.
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
// Deletes a material from the material library.
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
// # Delete a PPT work
//
// Description:
//
// - This API operation uses the HTTP Server-Sent Events (SSE) protocol.
//
// - The OpenAPI portal is not compatible with the SSE inference protocol. You cannot directly test this API operation in the portal. For more information about how to call this API operation using the software development kit (SDK) for Java or Python, see [PPT Generation Best practices](https://help.aliyun.com/zh/model-studio/ppt-generation-best-practices).
//
// - To obtain the latest version of the asynchronous Java SDK, see [this link](https://api.aliyun.com/api-tools/sdk/AiMiaoBi?spm=a2c4g.11186623.0.0.4cd3170d7rccDC\\&version=2023-08-01\\&language=java-async-tea\\&tab=primer-doc).
//
// @param request - DeletePptArtifactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePptArtifactResponse
func (client *Client) DeletePptArtifactWithContext(ctx context.Context, request *DeletePptArtifactRequest, runtime *dara.RuntimeOptions) (_result *DeletePptArtifactResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ExternalUserId) {
		body["ExternalUserId"] = request.ExternalUserId
	}

	if !dara.IsNil(request.PptArtifactId) {
		body["PptArtifactId"] = request.PptArtifactId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePptArtifact"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePptArtifactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified custom style.
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
// Extracts the content of documents from URLs.
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
// Call this API to download the structured rule library for further processing. You can use this API to download either the structured rule library before post-processing or the structured rule library currently available for auditing. For specific usage, see the input parameter descriptions.
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
// # API for downloading bidding documents
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
// Edit custom audit term records.
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
// Edits the content of a bidding document.
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
// Exports the tag details for a specified enterprise VOC analysis task.
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
// Exports an automated review report.
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
// Exports custom data source topic perspective analysis task results.
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
	if !dara.IsNil(request.ExportType) {
		body["ExportType"] = request.ExportType
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
// Exports the history of articles created in MiaoBi.
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
// Exports topic planning documents and provides a publicly accessible URL that expires in one hour.
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
// Exports all interventions.
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
// Exports a PPT artifact.
//
// @param request - ExportPptArtifactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportPptArtifactResponse
func (client *Client) ExportPptArtifactWithContext(ctx context.Context, request *ExportPptArtifactRequest, runtime *dara.RuntimeOptions) (_result *ExportPptArtifactResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Edit) {
		body["Edit"] = request.Edit
	}

	if !dara.IsNil(request.ExportFileType) {
		body["ExportFileType"] = request.ExportFileType
	}

	if !dara.IsNil(request.ExternalUserId) {
		body["ExternalUserId"] = request.ExternalUserId
	}

	if !dara.IsNil(request.PptArtifactId) {
		body["PptArtifactId"] = request.PptArtifactId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	if !dara.IsNil(request.Zip) {
		body["Zip"] = request.Zip
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportPptArtifact"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportPptArtifactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Provides feedback on the quality of the content that the model generates.
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
// Retrieves the results of a term library export task.
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
// Fetches the result of an asynchronous document export task.
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
// Retrieve the results of image generation tasks.
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
// Retrieves the result of a term import task.
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
// Retrieve the layout task result.
//
// @param request - FetchParseDocumentLayoutTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FetchParseDocumentLayoutTaskResponse
func (client *Client) FetchParseDocumentLayoutTaskWithContext(ctx context.Context, request *FetchParseDocumentLayoutTaskRequest, runtime *dara.RuntimeOptions) (_result *FetchParseDocumentLayoutTaskResponse, _err error) {
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
		Action:      dara.String("FetchParseDocumentLayoutTask"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FetchParseDocumentLayoutTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Start a task to export content as a Word document.
//
// Description:
//
// The Quanmiao product supports iframe embedding. For details, see [Customer Integration: Quanmiao Public Cloud iframe Customization Guide](https://help.aliyun.com/document_detail/3000990.html).
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
// Generate a temporary public URL.
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
// Asynchronously generates an image from text.
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
// Generates a file upload configuration.
//
// 1\\. Call this API to obtain the upload configuration. The API returns the `PostUrl` (an internal OSS address for AI Writing Assistant), temporary OSS authentication information (`key`, `OSSAccessKeyId`, `Signature`, and `policy`), and the unique file identifier `fileKey`.
//
// 2\\. The client uses the `PostUrl` and the temporary authentication information (`key`, `OSSAccessKeyId`, `Signature`, and `policy`) to upload the file.
//
// 3\\. Use the `fileKey` to call subsequent APIs that require a `fileKey`, such as `GenerateFileUrlByKey`.
//
// Description:
//
// This API returns the address and credentials for file uploads. For more information, see [OSS Form Upload](https://help.aliyun.com/zh/oss/user-guide/form-upload?scm=20140722.H_31849._.OR_help-T_cn~zh-V_1).
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
// Generates viewpoints from article snippets.
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
// Queries the progress of a post-processing task for a rule library. Use this operation together with the ConfirmAndPostProcessAuditNote operation to check the status of the current post-processing task.
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
// Checks the processing status of an uploaded rule library. This operation returns the current status of the upload task, the size of the parsed rule library file, and its storage path.
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
// Retrieves the status of a video editing task.
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
	if !dara.IsNil(request.ShowAnalysisResults) {
		body["ShowAnalysisResults"] = request.ShowAnalysisResults
	}

	if !dara.IsNil(request.ShowResourceInfo) {
		body["ShowResourceInfo"] = request.ShowResourceInfo
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
// Query the rule libraries that are currently available for audit. This operation returns only rule libraries that are active for auditing. To view the contents of a custom rule library, use the DownloadAuditNote API.
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
// Retrieves the generation results of a bidding document.
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
// Retrieve the remaining limit for the bidding feature.
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
// Retrieves the classifications from a tag mining task.
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
// Retrieves the built-in resources for smart clipping.
//
// @param request - GetClipsBuildInResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClipsBuildInResourceResponse
func (client *Client) GetClipsBuildInResourceWithContext(ctx context.Context, request *GetClipsBuildInResourceRequest, runtime *dara.RuntimeOptions) (_result *GetClipsBuildInResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetClipsBuildInResource"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetClipsBuildInResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the result of a custom hot topic broadcast job.
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
// Retrieves the results of a topic analysis task for a custom data source.
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
// Retrieve custom text.
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
// Retrieve the result of a custom topic selection perspective analysis task.
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
// Retrieves configuration information for write data sources and general search data sources.
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
// Data source management details.
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
// Retrieve the data source document.
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
// Queries deep writing tasks. You can use it to check the running status of a specific task.
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
// Queries the result of a deep writing task. If the task is not complete, the operation returns its current status—such as queued, running, failed, or canceled. If the task is complete, the operation returns a URL that points to a compressed package of the task output that you can download.
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
// Retrieves the result of a content aggregation task.
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
// Retrieves information about a document.
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
// Retrieves the result of an enterprise Voice of the Customer (VOC) analysis task.
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
// Retrieves the source URL that is currently used for factuality audit.
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
// MiaoRead calculates the word count for a document.
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
// Queries general configurations.
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
// Get generated content. Queries the history of articles generated in MiaoBi.
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
// Query news broadcast orders.
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
// Retrieve the global intervention reply.
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
// Gets information about an import task.
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
// Retrieves the details of an intervention rule.
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
// Retrieves the download URL for the intervention import template.
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
// Retrieves detailed information about a material from the Material Library.
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
// Queries information about a PPT artifact.
//
// Description:
//
// Usage notes:
//
// - This API uses the HTTP Server-Sent Events (SSE) protocol.
//
// - The OpenAPI portal is not compatible with the SSE protocol and cannot be used for direct debugging. For examples of how to call the API using an SDK for Java or Python, see [PPT Generation Best practices](https://help.aliyun.com/zh/model-studio/ppt-generation-best-practices).
//
// - To obtain the latest version of the asynchronous Java SDK, [download it from the API portal](https://api.aliyun.com/api-tools/sdk/AiMiaoBi?spm=a2c4g.11186623.0.0.4cd3170d7rccDC\\&version=2023-08-01\\&language=java-async-tea\\&tab=primer-doc).
//
// @param request - GetPptArtifactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPptArtifactResponse
func (client *Client) GetPptArtifactWithContext(ctx context.Context, request *GetPptArtifactRequest, runtime *dara.RuntimeOptions) (_result *GetPptArtifactResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ExternalUserId) {
		body["ExternalUserId"] = request.ExternalUserId
	}

	if !dara.IsNil(request.PptArtifactId) {
		body["PptArtifactId"] = request.PptArtifactId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPptArtifact"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPptArtifactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the result of a PPT export task.
//
// @param request - GetPptArtifactExportResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPptArtifactExportResultResponse
func (client *Client) GetPptArtifactExportResultWithContext(ctx context.Context, request *GetPptArtifactExportResultRequest, runtime *dara.RuntimeOptions) (_result *GetPptArtifactExportResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ExportTaskId) {
		body["ExportTaskId"] = request.ExportTaskId
	}

	if !dara.IsNil(request.ExternalUserId) {
		body["ExternalUserId"] = request.ExternalUserId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPptArtifactExportResult"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPptArtifactExportResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the configuration of a PPT component.
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
	if !dara.IsNil(request.ExternalUserId) {
		body["ExternalUserId"] = request.ExternalUserId
	}

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
// Gets information about a PPT task.
//
// @param request - GetPptInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPptInfoResponse
func (client *Client) GetPptInfoWithContext(ctx context.Context, request *GetPptInfoRequest, runtime *dara.RuntimeOptions) (_result *GetPptInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ExternalUserId) {
		body["ExternalUserId"] = request.ExternalUserId
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
		Action:      dara.String("GetPptInfo"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPptInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the filters for PowerPoint (PPT) templates.
//
// @param request - GetPptTemplateSelectorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPptTemplateSelectorResponse
func (client *Client) GetPptTemplateSelectorWithContext(ctx context.Context, request *GetPptTemplateSelectorRequest, runtime *dara.RuntimeOptions) (_result *GetPptTemplateSelectorResponse, _err error) {
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
		Action:      dara.String("GetPptTemplateSelector"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPptTemplateSelectorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves configuration information, such as intelligent configuration styles and inference-related metadata configurations.
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
// Queries the result of an automated review.
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
// Queries a one-click video editing task.
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
// Retrieves the analysis result of a style learning task.
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
// Retrieve hot topic event information by ID.
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
// Retrieves the result of a topic selection perspective analysis task.
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
// Imports an intervention file.
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
// Asynchronously import an intervention file.
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
// Important: This is a billable API operation.
//
// This API performs two operations:
//
// 1\\. Returns the initialization code for the "PPT Generation" frontend component.
//
// 2\\. Performs billing.
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
	if !dara.IsNil(request.ExternalUserId) {
		body["ExternalUserId"] = request.ExternalUserId
	}

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
// Starts a task to create a presentation.
//
// @param request - InitiatePptCreationV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitiatePptCreationV2Response
func (client *Client) InitiatePptCreationV2WithContext(ctx context.Context, request *InitiatePptCreationV2Request, runtime *dara.RuntimeOptions) (_result *InitiatePptCreationV2Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ExternalUserId) {
		body["ExternalUserId"] = request.ExternalUserId
	}

	if !dara.IsNil(request.IsMobile) {
		body["IsMobile"] = request.IsMobile
	}

	if !dara.IsNil(request.Outline) {
		body["Outline"] = request.Outline
	}

	if !dara.IsNil(request.PptTemplateId) {
		body["PptTemplateId"] = request.PptTemplateId
	}

	if !dara.IsNil(request.PptTemplateType) {
		body["PptTemplateType"] = request.PptTemplateType
	}

	if !dara.IsNil(request.PptTitle) {
		body["PptTitle"] = request.PptTitle
	}

	if !dara.IsNil(request.ProcessType) {
		body["ProcessType"] = request.ProcessType
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
		Action:      dara.String("InitiatePptCreationV2"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InitiatePptCreationV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets global intervention replies.
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
// Insert an intervention rule.
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
// Retrieves a paginated list of details for an enterprise Voice of the Customer (VOC) analysis task.
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
// Retrieves a list of asynchronous tasks.
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
// Retrieves a list of audit dimensions.
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
// Retrieve a list of term libraries.
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
// Lists smart video editing tasks.
//
// @param request - ListAutoClipsTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAutoClipsTaskResponse
func (client *Client) ListAutoClipsTaskWithContext(ctx context.Context, request *ListAutoClipsTaskRequest, runtime *dara.RuntimeOptions) (_result *ListAutoClipsTaskResponse, _err error) {
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
		Action:      dara.String("ListAutoClipsTask"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAutoClipsTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of bidding document writing tasks.
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
// Retrieves the system-defined presets for the Generate Content workflow. These presets include options such as writing style, article length, output language, and the number of articles to generate.
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
// Retrieve a list of custom texts.
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
// Lists custom viewpoints.
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
// # Permission to list datasets
//
// @param request - ListDataPermissionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataPermissionsResponse
func (client *Client) ListDataPermissionsWithContext(ctx context.Context, request *ListDataPermissionsRequest, runtime *dara.RuntimeOptions) (_result *ListDataPermissionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DataId) {
		body["DataId"] = request.DataId
	}

	if !dara.IsNil(request.DataType) {
		body["DataType"] = request.DataType
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
		Action:      dara.String("ListDataPermissions"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataPermissionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists data source documents.
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
	if !dara.IsNil(tmpReq.CategoryUuids) {
		request.CategoryUuidsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CategoryUuids, dara.String("CategoryUuids"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DocIds) {
		request.DocIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DocIds, dara.String("DocIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DocUuids) {
		request.DocUuidsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DocUuids, dara.String("DocUuids"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ExcludeFields) {
		request.ExcludeFieldsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExcludeFields, dara.String("ExcludeFields"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.IncludeFields) {
		request.IncludeFieldsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IncludeFields, dara.String("IncludeFields"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoryUuidsShrink) {
		body["CategoryUuids"] = request.CategoryUuidsShrink
	}

	if !dara.IsNil(request.CreateTimeEnd) {
		body["CreateTimeEnd"] = request.CreateTimeEnd
	}

	if !dara.IsNil(request.CreateTimeStart) {
		body["CreateTimeStart"] = request.CreateTimeStart
	}

	if !dara.IsNil(request.DatasetDescription) {
		body["DatasetDescription"] = request.DatasetDescription
	}

	if !dara.IsNil(request.DatasetId) {
		body["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.DatasetName) {
		body["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.DocIdsShrink) {
		body["DocIds"] = request.DocIdsShrink
	}

	if !dara.IsNil(request.DocType) {
		body["DocType"] = request.DocType
	}

	if !dara.IsNil(request.DocUuidsShrink) {
		body["DocUuids"] = request.DocUuidsShrink
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ExcludeFieldsShrink) {
		body["ExcludeFields"] = request.ExcludeFieldsShrink
	}

	if !dara.IsNil(request.Extend1) {
		body["Extend1"] = request.Extend1
	}

	if !dara.IsNil(request.Extend2) {
		body["Extend2"] = request.Extend2
	}

	if !dara.IsNil(request.Extend3) {
		body["Extend3"] = request.Extend3
	}

	if !dara.IsNil(request.IncludeFieldsShrink) {
		body["IncludeFields"] = request.IncludeFieldsShrink
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
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

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.TagsShrink) {
		body["Tags"] = request.TagsShrink
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
// # Data source management - query
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
	if !dara.IsNil(request.DatasetDescription) {
		body["DatasetDescription"] = request.DatasetDescription
	}

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
// History of online inference scenarios.
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
// Miàodú retrieves the list of documents.
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
// Searches government document libraries based on complex conditions.
//
// Description:
//
// The Quanmiao product supports iframe embedding. For more information, see [Customer Integration: Quanmiao Public Cloud iframe Customization Plan](https://help.aliyun.com/document_detail/3000990.html).
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
// Lists the enterprise-specific PPT templates.
//
// Description:
//
// Quanmiao supports iframe integration. For details, see the [Quanmiao Public Cloud iframe Customization Guide](https://help.aliyun.com/document_detail/3000990.html).
//
// @param request - ListEnterprisePptTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEnterprisePptTemplatesResponse
func (client *Client) ListEnterprisePptTemplatesWithContext(ctx context.Context, request *ListEnterprisePptTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListEnterprisePptTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.Skip) {
		query["Skip"] = request.Skip
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEnterprisePptTemplates"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEnterprisePptTemplatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// List of novel perspectives.
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
// Lists the general configurations.
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
// Retrieve a list of documents: Query the history of articles created in MiaoBi.
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
// Retrieves the list of trending topic hotspots.
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
// Retrieve the list of hot ranking sources for all platforms.
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
// Retrieves a list of trending topics.
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
	if !dara.IsNil(request.CreateTimeEnd) {
		body["CreateTimeEnd"] = request.CreateTimeEnd
	}

	if !dara.IsNil(request.CreateTimeStart) {
		body["CreateTimeStart"] = request.CreateTimeStart
	}

	if !dara.IsNil(request.CustomField) {
		body["CustomField"] = request.CustomField
	}

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
// List of popular viewpoints.
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
// Lists the number of intervention projects.
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
// Retrieve a list of import tasks.
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
// Retrieves a list of intervention rules.
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
// Retrieves the list of intervention items.
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
// Retrieve the list of materials from the Material Library.
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
// Retrieves a list of planning proposals.
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
// Queries a list of PPT artifacts.
//
// @param request - ListPptArtifactsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPptArtifactsResponse
func (client *Client) ListPptArtifactsWithContext(ctx context.Context, request *ListPptArtifactsRequest, runtime *dara.RuntimeOptions) (_result *ListPptArtifactsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ExternalUserId) {
		body["ExternalUserId"] = request.ExternalUserId
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Query) {
		body["Query"] = request.Query
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPptArtifacts"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPptArtifactsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of PowerPoint templates.
//
// @param request - ListPptTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPptTemplatesResponse
func (client *Client) ListPptTemplatesWithContext(ctx context.Context, request *ListPptTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListPptTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CareerId) {
		body["CareerId"] = request.CareerId
	}

	if !dara.IsNil(request.ColourId) {
		body["ColourId"] = request.ColourId
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.StyleId) {
		body["StyleId"] = request.StyleId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPptTemplates"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPptTemplatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the dialogue data for a search generation task.
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
// Retrieves the task details for MiaoSou search generation tasks.
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
// Queries the list of historical tasks generated by Miaosou Search.
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
// Retrieves the list of style learning analysis results.
//
// Description:
//
// The Quanmiao product supports iframe embedding. For details, see [Customer Integration: Quanmiao Public Cloud iframe Customization Plan](https://help.aliyun.com/document_detail/3000990.html).
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
// List of timeliness perspectives.
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
// Retrieve hot spot recommendation events.
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
// Retrieves recommended viewpoints for hot spot events.
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
// Retrieve version information for your purchased services.
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
// List of viewpoints from netizens.
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
// Retrieves the list of writing styles.
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
// Queries the details of submitted asynchronous task executions.
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
// Queries the results of an audit task.
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
// Queries video audit results.
//
// Description:
//
// Queries video audit results by task ID. The response includes video information, shot information, and audit results.
//
// @param request - QueryVideoAuditResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryVideoAuditResultResponse
func (client *Client) QueryVideoAuditResultWithContext(ctx context.Context, request *QueryVideoAuditResultRequest, runtime *dara.RuntimeOptions) (_result *QueryVideoAuditResultResponse, _err error) {
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
		Action:      dara.String("QueryVideoAuditResult"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryVideoAuditResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Abbreviates the specified content.
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
// Abbreviates the specified content.
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
// MiaoBi: AI-assisted writing
//
// Description:
//
// QuanMiao products support iframe embedding. For more information, see [QuanMiao Public Cloud iframe Customization for Customer Onboarding](https://help.aliyun.com/document_detail/3000990.html).
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
// MiaoBi: AI-assisted writing
//
// Description:
//
// QuanMiao products support iframe embedding. For more information, see [QuanMiao Public Cloud iframe Customization for Customer Onboarding](https://help.aliyun.com/document_detail/3000990.html).
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
// Miaodu generates mind maps of books.
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
// Miaodu generates mind maps of books.
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

	if !dara.IsNil(request.ResponseFormat) {
		body["ResponseFormat"] = request.ResponseFormat
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
// Extracts a summary, structured selling points, and hotwords from a book.
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
// Extracts a summary, structured selling points, and hotwords from a book.
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
	if !dara.IsNil(request.CleanCache) {
		body["CleanCache"] = request.CleanCache
	}

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
// A smart card interface for books.
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
// A smart card interface for books.
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
// Predicts user comments for a specified article.
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
// Predicts user comments for a specified article.
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
// Continues generating content.
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
// Continues generating content.
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
// Analyzes custom hot topics.
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
// Analyzes custom hot topics.
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
// Perspective analysis of custom topics.
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
// Perspective analysis of custom topics.
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
// Queries deep writing events. The system returns detailed information about the task execution as a stream of Server-Sent Events (SSE).
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
// Queries deep writing events. The system returns detailed information about the task execution as a stream of Server-Sent Events (SSE).
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
// Generate a three-level, multilingual mind map from an article or a book, with control over the number of second-level nodes and the word count of leaf nodes.
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
// Generate a three-level, multilingual mind map from an article or a book, with control over the number of second-level nodes and the word count of leaf nodes.
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

	if !dara.IsNil(request.ResponseFormat) {
		body["ResponseFormat"] = request.ResponseFormat
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
// Generates a summary for an article, video, or URL, including a full-text summary, key points, and a chapter overview (i.e., segmented content with summaries and abstracts for each segment). It also supports multilingual input and output. If the user only requires a full-text summary of an article, they can use the RunDocSummary API. For details, see https://help.aliyun.com/zh/model-studio/api-aimiaobi-2023-08-01-rundocsummary.
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
// Generates a summary for an article, video, or URL, including a full-text summary, key points, and a chapter overview (i.e., segmented content with summaries and abstracts for each segment). It also supports multilingual input and output. If the user only requires a full-text summary of an article, they can use the RunDocSummary API. For details, see https://help.aliyun.com/zh/model-studio/api-aimiaobi-2023-08-01-rundocsummary.
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
// Article Q&A: For a natural language query, provide a textual answer within the specified article scope (accompanied by images if available) and display source attribution information.
//
// Multimodal File Q&A: For a natural language query, provide a textual answer within the specified multimodal file scope, along with relevant images, video segments, or text, and display source attribution information.
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
// Article Q&A: For a natural language query, provide a textual answer within the specified article scope (accompanied by images if available) and display source attribution information.
//
// Multimodal File Q&A: For a natural language query, provide a textual answer within the specified multimodal file scope, along with relevant images, video segments, or text, and display source attribution information.
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
// Automatically adds tags to selected text or a specified chat and generates a smart card note.
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
// Automatically adds tags to selected text or a specified chat and generates a smart card note.
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
// Generates a summary of an article, video, or URL—that is, a concise overview of the entire content. It also supports multilingual input and output.
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
// Generates a summary of an article, video, or URL—that is, a concise overview of the entire content. It also supports multilingual input and output.
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
// An API for document translation between English and Chinese.
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
// An API for document translation between English and Chinese.
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
// Rewrites an article in a specified style.
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
// Rewrites an article in a specified style.
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
// Expands content.
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
// Expands content.
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
// Submits a query and returns several related queries.
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
// Submits a query and returns several related queries.
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
// Extract keywords from a specified document. Keywords are domain-specific professional terms or concepts that represent and identify a particular industry or field. They accurately describe and summarize the core content, key people, major events, or technical terms in that domain.
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
// Extract keywords from a specified document. Keywords are domain-specific professional terms or concepts that represent and identify a particular industry or field. They accurately describe and summarize the core content, key people, major events, or technical terms in that domain.
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
// Extracts and generates keywords using AMB.
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
// Extracts and generates keywords using AMB.
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
// Generate an outline-style summary for multiple documents, videos, or URLs. The summary includes a consolidated overview and key points. This operation supports multiple input and output languages.
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
// Generate an outline-style summary for multiple documents, videos, or URLs. The summary includes a consolidated overview and key points. This operation supports multiple input and output languages.
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
// Generates a PowerPoint outline.
//
// Description:
//
// Instructions:
//
// - This API uses the HTTP Server-Sent Events (SSE) protocol.
//
// - You cannot test this API directly in the OpenAPI Portal because the portal is not compatible with the SSE inference protocol. For examples of how to call the API using the SDK for Java or Python, see [PPT Generation Best practices](https://help.aliyun.com/en/model-studio/ppt-generation-best-practices).
//
// - To obtain the latest version of the asynchronous Java SDK, [click this link](https://api.aliyun.com/api-tools/sdk/AiMiaoBi?spm=a2c4g.11186623.0.0.4cd3170d7rccDC\\&version=2023-08-01\\&language=java-async-tea\\&tab=primer-doc).
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
// Generates a PowerPoint outline.
//
// Description:
//
// Instructions:
//
// - This API uses the HTTP Server-Sent Events (SSE) protocol.
//
// - You cannot test this API directly in the OpenAPI Portal because the portal is not compatible with the SSE inference protocol. For examples of how to call the API using the SDK for Java or Python, see [PPT Generation Best practices](https://help.aliyun.com/en/model-studio/ppt-generation-best-practices).
//
// - To obtain the latest version of the asynchronous Java SDK, [click this link](https://api.aliyun.com/api-tools/sdk/AiMiaoBi?spm=a2c4g.11186623.0.0.4cd3170d7rccDC\\&version=2023-08-01\\&language=java-async-tea\\&tab=primer-doc).
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
	if !dara.IsNil(request.ExternalUserId) {
		body["ExternalUserId"] = request.ExternalUserId
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
// Enter writing instructions to quickly generate content.
//
// Description:
//
// ### Access Instructions:
//
// - The OpenAPI portal is incompatible with the Server-Sent Events (SSE) inference protocol. Therefore, you cannot directly debug this operation. For an example of how to call the API using an SDK, see [Miaobi Best Practices](https://help.aliyun.com/zh/model-studio/best-practices-for-miaobi-api?spm=a2c4g.11186623.help-menu-2400256.d_1_12_6_2_1_0.39892421FntuI2\\&scm=20140722.H_2844289._.OR_help-T_cn~zh-V_1).
//
// - Click this [link](https://api.aliyun.com/api-tools/sdk/AiMiaoBi?version=2023-08-01\\&language=java-async-tea\\&tab=primer-doc) to download the latest version of the Java asynchronous SDK.
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
// Enter writing instructions to quickly generate content.
//
// Description:
//
// ### Access Instructions:
//
// - The OpenAPI portal is incompatible with the Server-Sent Events (SSE) inference protocol. Therefore, you cannot directly debug this operation. For an example of how to call the API using an SDK, see [Miaobi Best Practices](https://help.aliyun.com/zh/model-studio/best-practices-for-miaobi-api?spm=a2c4g.11186623.help-menu-2400256.d_1_12_6_2_1_0.39892421FntuI2\\&scm=20140722.H_2844289._.OR_help-T_cn~zh-V_1).
//
// - Click this [link](https://api.aliyun.com/api-tools/sdk/AiMiaoBi?version=2023-08-01\\&language=java-async-tea\\&tab=primer-doc) to download the latest version of the Java asynchronous SDK.
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
// AI Miaosou – Intelligent Search Generation: This API delivers the search and generation capabilities of the Miaosou homepage. It supports general search and media asset search, along with features such as user query clarification, multimodal knowledge search, and multi-agent generation.
//
// – General Search: Performs semantic retrieval on centralized knowledge data and applies multi-agent post-processing to the results, such as summary generation, abstracting, and timeline summarization.
//
// – Media Asset Search: Conducts an exhaustive full-text search to retrieve highly relevant knowledge and supports multi-agent post-processing, such as clustering and news extraction.
//
// Description:
//
// ### Integration notes:
//
// - This API uses the HTTP Server-Sent Events (SSE) protocol.
//
// - The OpenAPI console does not support SSE inference protocols and cannot be used for direct testing. For SDK-based integration examples (Java and Python), see the [Miaosou Best Practices](https://help.aliyun.com/zh/model-studio/user-guide/best-practices-for-miaosou-api/?spm=a2c4g.11186623.help-menu-2400256.d_1_3_3_2_1_2.42a64a34eIyBhn) documentation.
//
// - To obtain the latest version of the Java asynchronous SDK, click [this link](https://api.aliyun.com/api-tools/sdk/AiMiaoBi?version=2023-08-01\\&language=java-async-tea\\&tab=primer-doc).
//
// ### Data sources for search:
//
// Supports three dataset types. See the [Miaosou Best Practices](https://help.aliyun.com/zh/model-studio/user-guide/best-practices-for-miaosou-api/?spm=a2c4g.11186623.help-menu-2400256.d_1_3_3_2_1_2.42a64a34eIyBhn) documentation for details.
//
// - Built-in “Internet search” dataset: Supports open-domain text, images, and video (video is not yet available) from the Internet.
//
// - Semantic (RAG) dataset: Manages enterprise private knowledge bases and supports text, images, video, and voice (voice is not yet available).
//
// - Third-party API dataset: Integrates directly with your own enterprise search APIs.
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
// AI Miaosou – Intelligent Search Generation: This API delivers the search and generation capabilities of the Miaosou homepage. It supports general search and media asset search, along with features such as user query clarification, multimodal knowledge search, and multi-agent generation.
//
// – General Search: Performs semantic retrieval on centralized knowledge data and applies multi-agent post-processing to the results, such as summary generation, abstracting, and timeline summarization.
//
// – Media Asset Search: Conducts an exhaustive full-text search to retrieve highly relevant knowledge and supports multi-agent post-processing, such as clustering and news extraction.
//
// Description:
//
// ### Integration notes:
//
// - This API uses the HTTP Server-Sent Events (SSE) protocol.
//
// - The OpenAPI console does not support SSE inference protocols and cannot be used for direct testing. For SDK-based integration examples (Java and Python), see the [Miaosou Best Practices](https://help.aliyun.com/zh/model-studio/user-guide/best-practices-for-miaosou-api/?spm=a2c4g.11186623.help-menu-2400256.d_1_3_3_2_1_2.42a64a34eIyBhn) documentation.
//
// - To obtain the latest version of the Java asynchronous SDK, click [this link](https://api.aliyun.com/api-tools/sdk/AiMiaoBi?version=2023-08-01\\&language=java-async-tea\\&tab=primer-doc).
//
// ### Data sources for search:
//
// Supports three dataset types. See the [Miaosou Best Practices](https://help.aliyun.com/zh/model-studio/user-guide/best-practices-for-miaosou-api/?spm=a2c4g.11186623.help-menu-2400256.d_1_3_3_2_1_2.42a64a34eIyBhn) documentation for details.
//
// - Built-in “Internet search” dataset: Supports open-domain text, images, and video (video is not yet available) from the Internet.
//
// - Semantic (RAG) dataset: Manages enterprise private knowledge bases and supports text, images, video, and voice (voice is not yet available).
//
// - Third-party API dataset: Integrates directly with your own enterprise search APIs.
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
// Miao Search enables text-to-text search.
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
// Miao Search enables text-to-text search.
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
// Writes content in a step-by-step pattern using an outline and summaries.
//
// Description:
//
// The Quanmiao product supports iframe embedding. For more information, see [Customer Onboarding: Quanmiao Public Cloud iframe Customization](https://help.aliyun.com/document_detail/3000990.html).
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
// Writes content in a step-by-step pattern using an outline and summaries.
//
// Description:
//
// The Quanmiao product supports iframe embedding. For more information, see [Customer Onboarding: Quanmiao Public Cloud iframe Customization](https://help.aliyun.com/document_detail/3000990.html).
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
// Analyzes the stylistic features of content.
//
// Description:
//
// Quanmiao products support iframe embedding. For more information, see [Customer Integration: Quanmiao Public Cloud iframe Customization Solution](https://help.aliyun.com/document_detail/3000990.html).
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
// Analyzes the stylistic features of content.
//
// Description:
//
// Quanmiao products support iframe embedding. For more information, see [Customer Integration: Quanmiao Public Cloud iframe Customization Solution](https://help.aliyun.com/document_detail/3000990.html).
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
// Generates a summary of content.
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
// Generates a summary of content.
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
// Polishes the specified text.
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
// Polishes the specified text.
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
// Miaobi generates titles.
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
// Miaobi generates titles.
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
// Selects and aggregates topics for MiaoCe.
//
// Description:
//
// All Miao products support iframe embedding. For more information, see [Customer Integration: Miao Public Cloud iFrame Customization Plan](https://help.aliyun.com/document_detail/3000990.html).
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
// Selects and aggregates topics for MiaoCe.
//
// Description:
//
// All Miao products support iframe embedding. For more information, see [Customer Integration: Miao Public Cloud iFrame Customization Plan](https://help.aliyun.com/document_detail/3000990.html).
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
// Translates text for content creation using AMB.
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
// Translates text for content creation using AMB.
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
// Generates video clip scripts using AI.
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
// Generates video clip scripts using AI.
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
// Calls AMB to rewrite text in a new tone.
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
// Calls AMB to rewrite text in a new tone.
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
// Direct writing.
//
// Description:
//
// This API is deprecated. For more information, see [RunWritingV2](https://help.aliyun.com/document_detail/2922606.html).
//
// The Quanmiao product supports iframe embedding. For more information, see [Customer integration: Quanmiao Public Cloud iframe customization](https://help.aliyun.com/document_detail/3000990.html).
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
// Direct writing.
//
// Description:
//
// This API is deprecated. For more information, see [RunWritingV2](https://help.aliyun.com/document_detail/2922606.html).
//
// The Quanmiao product supports iframe embedding. For more information, see [Customer integration: Quanmiao Public Cloud iframe customization](https://help.aliyun.com/document_detail/3000990.html).
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
// # AI writing
//
// Description:
//
// For instructions on embedding Quanmiao products using an iframe, see [Customer integration_Quanmiao public cloud iframe customized solution](https://help.aliyun.com/document_detail/3000990.html).
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
// # AI writing
//
// Description:
//
// For instructions on embedding Quanmiao products using an iframe, see [Customer integration_Quanmiao public cloud iframe customized solution](https://help.aliyun.com/document_detail/3000990.html).
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
// Saves custom text.
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
// Saves the data source configuration for content creation and general search.
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
// Save material: Save material to the Material Library.
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
// Configuration: Cloud storage parameter settings
//
// @param request - SaveOrUpdateOssConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveOrUpdateOssConfigResponse
func (client *Client) SaveOrUpdateOssConfigWithContext(ctx context.Context, request *SaveOrUpdateOssConfigRequest, runtime *dara.RuntimeOptions) (_result *SaveOrUpdateOssConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BucketName) {
		body["BucketName"] = request.BucketName
	}

	if !dara.IsNil(request.EndPoint) {
		body["EndPoint"] = request.EndPoint
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveOrUpdateOssConfig"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveOrUpdateOssConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Saves a custom writing style.
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
// Searches documents in a data source.
//
// @param tmpReq - SearchDatasetDocumentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchDatasetDocumentsResponse
func (client *Client) SearchDatasetDocumentsWithContext(ctx context.Context, tmpReq *SearchDatasetDocumentsRequest, runtime *dara.RuntimeOptions) (_result *SearchDatasetDocumentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SearchDatasetDocumentsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CategoryUuids) {
		request.CategoryUuidsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CategoryUuids, dara.String("CategoryUuids"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DocIds) {
		request.DocIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DocIds, dara.String("DocIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DocTypes) {
		request.DocTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DocTypes, dara.String("DocTypes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DocUuids) {
		request.DocUuidsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DocUuids, dara.String("DocUuids"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoryUuidsShrink) {
		body["CategoryUuids"] = request.CategoryUuidsShrink
	}

	if !dara.IsNil(request.CreateTimeEnd) {
		body["CreateTimeEnd"] = request.CreateTimeEnd
	}

	if !dara.IsNil(request.CreateTimeStart) {
		body["CreateTimeStart"] = request.CreateTimeStart
	}

	if !dara.IsNil(request.DatasetId) {
		body["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.DatasetName) {
		body["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.DocIdsShrink) {
		body["DocIds"] = request.DocIdsShrink
	}

	if !dara.IsNil(request.DocTypesShrink) {
		body["DocTypes"] = request.DocTypesShrink
	}

	if !dara.IsNil(request.DocUuidsShrink) {
		body["DocUuids"] = request.DocUuidsShrink
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Extend1) {
		body["Extend1"] = request.Extend1
	}

	if !dara.IsNil(request.Extend2) {
		body["Extend2"] = request.Extend2
	}

	if !dara.IsNil(request.Extend3) {
		body["Extend3"] = request.Extend3
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

	if !dara.IsNil(request.SearchMode) {
		body["SearchMode"] = request.SearchMode
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TagsShrink) {
		body["Tags"] = request.TagsShrink
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
// Searches for news based on your input. This feature is currently limited to web search.
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
// Executes predefined asynchronous tasks.
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
// Model Studio provides the same feature as the Intelligent Proofreading module in Alibaba Cloud public cloud: uploading a custom rule library. Due to authentication restrictions, you must pass the file key (FileKey) of your custom rule library file as an input parameter to successfully call this API. After you call this API, Model Studio processes your custom rule library and returns a structured result in XLSX format. You can call GetAuditNoteProcessingStatus to check the processing status or call DownloadAuditNote to download the processed rule library. This API is under active development and will eventually accept a publicly accessible file URL instead of a FileKey.
//
// Description:
//
// All Model Studio products support iframe embedding. For details, see [Customer Integration: Model Studio Public Cloud iFrame Customization Guide](https://help.aliyun.com/document_detail/3000990.html).
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
// Submits an audit task.
//
// Description:
//
// The Quanmiao product supports iframe embedding. For more information, see [Customer Integration: Quanmiao Public Cloud iframe Customization Plan](https://alidocs.dingtalk.com/i/nodes/m9bN7RYPWdyrPBREcyM6jDQ2VZd1wyK0?cid=116617178%3A898142682\\&utm_source=im\\&utm_scene=team_space\\&iframeQuery=utm_medium%3Dim_card%26utm_source%3Dim\\&utm_medium=im_card\\&corpId=dingd8e1123006514592).
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
// You can submit a custom broadcast list job.
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
// # Submit topic trend analysis using a custom data source
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

	if !dara.IsNil(tmpReq.Topics) {
		request.TopicsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Topics, dara.String("Topics"), dara.String("json"))
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

	if !dara.IsNil(request.TopicsShrink) {
		body["Topics"] = request.TopicsShrink
	}

	if !dara.IsNil(request.TopicsFileUrl) {
		body["TopicsFileUrl"] = request.TopicsFileUrl
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
// Submits a custom analysis task to analyze hot topic perspectives.
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
// You can submit a deep writing task. You can provide information such as questions, instructions, and attachments, based on the topic you want to research or analyze. The system schedules and executes this task in the background.
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
// Submits a content clustering task.
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
// Submits a Voice of the Customer (VOC) asynchronous task.
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
// # Export task for a thesaurus
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
// MiaoBi provides a factuality audit capability that verifies facts using web search and supports custom configuration of search source URLs.
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
// Submit a custom dictionary import task.
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
// 提交版本任务
//
// @param request - SubmitParseDocumentLayoutTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitParseDocumentLayoutTaskResponse
func (client *Client) SubmitParseDocumentLayoutTaskWithContext(ctx context.Context, request *SubmitParseDocumentLayoutTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitParseDocumentLayoutTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitParseDocumentLayoutTask"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitParseDocumentLayoutTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submit a smart audit request.
//
// Description:
//
// The Quanmiao product supports iframe embedding. For more information, see [Customer integration: Quanmiao public cloud iframe customization guide](https://help.aliyun.com/document_detail/3000990.html).
//
// # Supported audit types
//
// ## Audit category overview
//
// | Audit category                 | Description                                                                                                                                                                                                                                                               |
//
// | ------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
//
// | Content accuracy               | Errors due to phonetic or visual similarity; punctuation errors; misuse of Chinese structural particles (的/地/得); inappropriate word choice or syntax errors; incorrect personal names; incorrect place names; incorrect references; errors in proper nouns or terminology |
//
// | Formatting issues              | Inconsistent capitalization; numeric errors; nonstandard units of measurement; use of traditional Chinese characters                                                                                                                                                      |
//
// | Content structure issues       | Redundant text; repeated segments; logical contradictions; unfilled placeholders                                                                                                                                                                                          |
//
// | Political sensitivity issues   | Sensitive content or directional risks; name or sorting errors; conventional phrasing errors; nonstandard institutional names; misquoting important speeches; references to disgraced officials; mismatched name-title pairings; incorrect title expressions              |
//
// | Security and compliance issues | Violence or terrorism; pornography; prohibited content; insults; disgraced celebrities; personal privacy violations; reporting standard violations                                                                                                                        |
//
// | Legal errors                   | Incorrect citation of laws and regulations; errors in legal provisions                                                                                                                                                                                                    |
//
// | Other domain-specific errors   | Violations of advertising law; financial information errors; scientific or technical term errors                                                                                                                                                                          |
//
// | Factuality check               | Factuality verification: correct or incorrect items                                                                                                                                                                                                                       |
//
// | Image audit                    | Image content moderation                                                                                                                                                                                                                                                  |
//
// | Custom word library            | Custom word library audit                                                                                                                                                                                                                                                 |
//
// | Rule library audit             | Rule library audit                                                                                                                                                                                                                                                        |
//
// | English proofreading           | Terminology standardization; verb tense accuracy; punctuation and quotation marks; spelling and language variants; sentence structure and clarity; numeric and percentage formatting; standardized phrasing                                                               |
//
// ***
//
// ## Sub-audit code values
//
// ### 1. Content accuracy
//
// | Description                                    | Code                 |
//
// | ---------------------------------------------- | -------------------- |
//
// | Phonetic or visual similarity errors           | PhoneticSimilarError |
//
// | Punctuation errors                             | PunctuationError     |
//
// | Misuse of Chinese structural particles (的/地/得) | ParticleUsageError   |
//
// | Inappropriate word choice or syntax errors     | WordError            |
//
// | Incorrect personal names                       | PersonNameError      |
//
// | Incorrect place names                          | LocationError        |
//
// | Incorrect references                           | ReferenceError       |
//
// | Errors in proper nouns or terminology          | NounItemError        |
//
// ### 2. Formatting issues
//
// | Description                           | Code                    |
//
// | ------------------------------------- | ----------------------- |
//
// | Inconsistent capitalization           | CapitalizationError     |
//
// | Numeric errors                        | NumberError             |
//
// | Nonstandard units of measurement      | UnitError               |
//
// | Use of traditional Chinese characters | TraditionalChineseError |
//
// ### 3. Content structure issues
//
// | Description            | Code                 |
//
// | ---------------------- | -------------------- |
//
// | Redundant text         | WordRedundancy       |
//
// | Repeated segments      | DuplicateError       |
//
// | Logical contradictions | LogicContradiction   |
//
// | Unfilled placeholders  | PlaceholderNotFilled |
//
// ### 4. Political sensitivity issues
//
// | Description                            | Code                        |
//
// | -------------------------------------- | --------------------------- |
//
// | Sensitive content or directional risks | SensitiveContentRisk        |
//
// | Name or sorting errors                 | NameOrderError              |
//
// | Conventional phrasing errors           | ConventionalExpressionError |
//
// | Nonstandard institutional names        | DepartmentNameError         |
//
// | Misquoting important speeches          | ImportantSpeechError        |
//
// | References to disgraced officials      | FallenOfficialError         |
//
// | Mismatched name-title pairings         | LeaderTitleMatchError       |
//
// | Incorrect title expressions            | TitleError                  |
//
// ### 5. Security and compliance issues
//
// | Description                   | Code                   |
//
// | ----------------------------- | ---------------------- |
//
// | Violence or terrorism         | ViolenceTerrorismError |
//
// | Pornography                   | PornographyError       |
//
// | Prohibited content            | ProhibitedContentError |
//
// | Insults                       | InsultError            |
//
// | Disgraced celebrities         | DisgracedArtistError   |
//
// | Personal privacy violations   | PersonalPrivacyError   |
//
// | Reporting standard violations | ReportingStandardError |
//
// ### 6. Legal errors
//
// | Description                                | Code                 |
//
// | ------------------------------------------ | -------------------- |
//
// | Incorrect citation of laws and regulations | LegalReferenceError  |
//
// | Errors in legal provisions                 | LegalProvisionsError |
//
// ### 7. Other domain-specific errors
//
// | Description                         | Code                            |
//
// | ----------------------------------- | ------------------------------- |
//
// | Violations of advertising law       | AdvertisingProhibitedWordsError |
//
// | Financial information errors        | FinancialInformationError       |
//
// | Scientific or technical term errors | TechnicalTermError              |
//
// ### 8. Factuality check
//
// | Description                              | Code           |
//
// | ---------------------------------------- | -------------- |
//
// | Factuality verification – correct item   | CorrectFact    |
//
// | Factuality verification – incorrect item | WrongFactError |
//
// ### 9. Image audit
//
// | Description | Code       |
//
// | ----------- | ---------- |
//
// | Image audit | ImageAudit |
//
// ### 10. Custom word library
//
// | Description         | Code        |
//
// | ------------------- | ----------- |
//
// | Custom word library | WordLibrary |
//
// ### 11. Rule library audit
//
// | Description        | Code              |
//
// | ------------------ | ----------------- |
//
// | Rule library audit | WrongQuestionBook |
//
// ### 12. English proofreading
//
// | Description                       | Code                         |
//
// | --------------------------------- | ---------------------------- |
//
// | Terminology standardization       | TerminologyNormalisation     |
//
// | Verb tense accuracy               | VerbTenseAccuracy            |
//
// | Punctuation and quotation marks   | PunctuationAndQuotationMarks |
//
// | Spelling and language variants    | SpellingAndLanguageVariety   |
//
// | Sentence structure and clarity    | SentenceStructureAndClarity  |
//
// | Numeric and percentage formatting | NumericAndPercentageStyle    |
//
// | Other standardized phrasing       | Others                       |
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
// Submits a one-click video editing task.
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
// Submits a hot spot analysis task for topic selection.
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
// Submits a video for review.
//
// Description:
//
// Quanmiao products support iframe embedding. For more information, see [Customer Integration_Quanmiao Public Cloud iframe Customization Solution](https://alidocs.dingtalk.com/i/nodes/m9bN7RYPWdyrPBREcyM6jDQ2VZd1wyK0?cid=116617178%3A898142682\\&utm_source=im\\&utm_scene=team_space\\&iframeQuery=utm_medium%3Dim_card%26utm_source%3Dim\\&utm_medium=im_card\\&corpId=dingd8e1123006514592).
//
// @param request - SubmitVideoAuditRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitVideoAuditResponse
func (client *Client) SubmitVideoAuditWithContext(ctx context.Context, request *SubmitVideoAuditRequest, runtime *dara.RuntimeOptions) (_result *SubmitVideoAuditResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Ext) {
		body["Ext"] = request.Ext
	}

	if !dara.IsNil(request.FileKey) {
		body["FileKey"] = request.FileKey
	}

	if !dara.IsNil(request.SnapshotInterval) {
		body["SnapshotInterval"] = request.SnapshotInterval
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
		Action:      dara.String("SubmitVideoAudit"),
		Version:     dara.String("2023-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitVideoAuditResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates custom text.
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
// This operation updates a data source.
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
	if !dara.IsNil(request.AccessLevel) {
		body["AccessLevel"] = request.AccessLevel
	}

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
// Updates a document in a dataset.
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
// Updates general configurations.
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
// Updates the history of an article created in AiMiaoBi.
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
// Update a material in the Material Library.
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
// Use MiaoDu to upload books.
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
// Upload document API for Miaodu.
//
// Description:
//
// Document upload is implemented through asynchronous invocation. After the invocation, you must use the getDocInfo API to periodically check the document status. Only when the document status becomes 1 can you proceed with subsequent operations such as generating a document summary, creating a full-text mind map, summarizing Q&A content, extracting keywords, or rewriting.
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
// Validates an enterprise VOC upload template.
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

	if !dara.IsNil(request.ResponseFormat) {
		body["ResponseFormat"] = request.ResponseFormat
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

func (client *Client) runBookIntroductionWithSSECtx_opYieldFunc(_yield chan *RunBookIntroductionResponse, _yieldErr chan error, ctx context.Context, request *RunBookIntroductionRequest, runtime *dara.RuntimeOptions) {
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

	if !dara.IsNil(request.ResponseFormat) {
		body["ResponseFormat"] = request.ResponseFormat
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

func (client *Client) runPptOutlineGenerationWithSSECtx_opYieldFunc(_yield chan *RunPptOutlineGenerationResponse, _yieldErr chan error, ctx context.Context, request *RunPptOutlineGenerationRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ExternalUserId) {
		body["ExternalUserId"] = request.ExternalUserId
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
