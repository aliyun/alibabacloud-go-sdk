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
	client.EndpointMap = map[string]*string{
		"cn-beijing": dara.String("aimiaobi.cn-beijing.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("aimiaobi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Add a custom term to the audit dictionary.
//
// @param tmpReq - AddAuditTermsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddAuditTermsResponse
func (client *Client) AddAuditTermsWithOptions(tmpReq *AddAuditTermsRequest, runtime *dara.RuntimeOptions) (_result *AddAuditTermsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Add a custom term to the audit dictionary.
//
// @param request - AddAuditTermsRequest
//
// @return AddAuditTermsResponse
func (client *Client) AddAuditTerms(request *AddAuditTermsRequest) (_result *AddAuditTermsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddAuditTermsResponse{}
	_body, _err := client.AddAuditTermsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AddDatasetDocumentWithOptions(tmpReq *AddDatasetDocumentRequest, runtime *dara.RuntimeOptions) (_result *AddDatasetDocumentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - AddDatasetDocumentRequest
//
// @return AddDatasetDocumentResponse
func (client *Client) AddDatasetDocument(request *AddDatasetDocumentRequest) (_result *AddDatasetDocumentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddDatasetDocumentResponse{}
	_body, _err := client.AddDatasetDocumentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AsyncCreateClipsTaskWithOptions(tmpReq *AsyncCreateClipsTaskRequest, runtime *dara.RuntimeOptions) (_result *AsyncCreateClipsTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - AsyncCreateClipsTaskRequest
//
// @return AsyncCreateClipsTaskResponse
func (client *Client) AsyncCreateClipsTask(request *AsyncCreateClipsTaskRequest) (_result *AsyncCreateClipsTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AsyncCreateClipsTaskResponse{}
	_body, _err := client.AsyncCreateClipsTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AsyncCreateClipsTimeLineWithOptions(tmpReq *AsyncCreateClipsTimeLineRequest, runtime *dara.RuntimeOptions) (_result *AsyncCreateClipsTimeLineResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - AsyncCreateClipsTimeLineRequest
//
// @return AsyncCreateClipsTimeLineResponse
func (client *Client) AsyncCreateClipsTimeLine(request *AsyncCreateClipsTimeLineRequest) (_result *AsyncCreateClipsTimeLineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AsyncCreateClipsTimeLineResponse{}
	_body, _err := client.AsyncCreateClipsTimeLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AsyncEditTimelineWithOptions(tmpReq *AsyncEditTimelineRequest, runtime *dara.RuntimeOptions) (_result *AsyncEditTimelineResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - AsyncEditTimelineRequest
//
// @return AsyncEditTimelineResponse
func (client *Client) AsyncEditTimeline(request *AsyncEditTimelineRequest) (_result *AsyncEditTimelineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AsyncEditTimelineResponse{}
	_body, _err := client.AsyncEditTimelineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AsyncUploadTenderDocWithOptions(request *AsyncUploadTenderDocRequest, runtime *dara.RuntimeOptions) (_result *AsyncUploadTenderDocResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return AsyncUploadTenderDocResponse
func (client *Client) AsyncUploadTenderDoc(request *AsyncUploadTenderDocRequest) (_result *AsyncUploadTenderDocResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AsyncUploadTenderDocResponse{}
	_body, _err := client.AsyncUploadTenderDocWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AsyncUploadVideoWithOptions(tmpReq *AsyncUploadVideoRequest, runtime *dara.RuntimeOptions) (_result *AsyncUploadVideoResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - AsyncUploadVideoRequest
//
// @return AsyncUploadVideoResponse
func (client *Client) AsyncUploadVideo(request *AsyncUploadVideoRequest) (_result *AsyncUploadVideoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AsyncUploadVideoResponse{}
	_body, _err := client.AsyncUploadVideoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AsyncWritingBiddingDocWithOptions(request *AsyncWritingBiddingDocRequest, runtime *dara.RuntimeOptions) (_result *AsyncWritingBiddingDocResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return AsyncWritingBiddingDocResponse
func (client *Client) AsyncWritingBiddingDoc(request *AsyncWritingBiddingDocRequest) (_result *AsyncWritingBiddingDocResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AsyncWritingBiddingDocResponse{}
	_body, _err := client.AsyncWritingBiddingDocWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) BindPptArtifactWithOptions(request *BindPptArtifactRequest, runtime *dara.RuntimeOptions) (_result *BindPptArtifactResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return BindPptArtifactResponse
func (client *Client) BindPptArtifact(request *BindPptArtifactRequest) (_result *BindPptArtifactResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindPptArtifactResponse{}
	_body, _err := client.BindPptArtifactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CancelAsyncTaskWithOptions(request *CancelAsyncTaskRequest, runtime *dara.RuntimeOptions) (_result *CancelAsyncTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CancelAsyncTaskResponse
func (client *Client) CancelAsyncTask(request *CancelAsyncTaskRequest) (_result *CancelAsyncTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelAsyncTaskResponse{}
	_body, _err := client.CancelAsyncTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CancelAuditTaskWithOptions(request *CancelAuditTaskRequest, runtime *dara.RuntimeOptions) (_result *CancelAuditTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CancelAuditTaskResponse
func (client *Client) CancelAuditTask(request *CancelAuditTaskRequest) (_result *CancelAuditTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelAuditTaskResponse{}
	_body, _err := client.CancelAuditTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CancelDeepWriteTaskWithOptions(request *CancelDeepWriteTaskRequest, runtime *dara.RuntimeOptions) (_result *CancelDeepWriteTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CancelDeepWriteTaskResponse
func (client *Client) CancelDeepWriteTask(request *CancelDeepWriteTaskRequest) (_result *CancelDeepWriteTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelDeepWriteTaskResponse{}
	_body, _err := client.CancelDeepWriteTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ClearIntervenesWithOptions(request *ClearIntervenesRequest, runtime *dara.RuntimeOptions) (_result *ClearIntervenesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ClearIntervenesResponse
func (client *Client) ClearIntervenes(request *ClearIntervenesRequest) (_result *ClearIntervenesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClearIntervenesResponse{}
	_body, _err := client.ClearIntervenesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ConfirmAndPostProcessAuditNoteWithOptions(request *ConfirmAndPostProcessAuditNoteRequest, runtime *dara.RuntimeOptions) (_result *ConfirmAndPostProcessAuditNoteResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ConfirmAndPostProcessAuditNoteResponse
func (client *Client) ConfirmAndPostProcessAuditNote(request *ConfirmAndPostProcessAuditNoteRequest) (_result *ConfirmAndPostProcessAuditNoteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfirmAndPostProcessAuditNoteResponse{}
	_body, _err := client.ConfirmAndPostProcessAuditNoteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateDataPermissionsWithOptions(tmpReq *CreateDataPermissionsRequest, runtime *dara.RuntimeOptions) (_result *CreateDataPermissionsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateDataPermissionsRequest
//
// @return CreateDataPermissionsResponse
func (client *Client) CreateDataPermissions(request *CreateDataPermissionsRequest) (_result *CreateDataPermissionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataPermissionsResponse{}
	_body, _err := client.CreateDataPermissionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateDatasetWithOptions(tmpReq *CreateDatasetRequest, runtime *dara.RuntimeOptions) (_result *CreateDatasetResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateDatasetRequest
//
// @return CreateDatasetResponse
func (client *Client) CreateDataset(request *CreateDatasetRequest) (_result *CreateDatasetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDatasetResponse{}
	_body, _err := client.CreateDatasetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateGeneralConfigWithOptions(request *CreateGeneralConfigRequest, runtime *dara.RuntimeOptions) (_result *CreateGeneralConfigResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateGeneralConfigResponse
func (client *Client) CreateGeneralConfig(request *CreateGeneralConfigRequest) (_result *CreateGeneralConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateGeneralConfigResponse{}
	_body, _err := client.CreateGeneralConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateGeneratedContentWithOptions(tmpReq *CreateGeneratedContentRequest, runtime *dara.RuntimeOptions) (_result *CreateGeneratedContentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateGeneratedContentRequest
//
// @return CreateGeneratedContentResponse
func (client *Client) CreateGeneratedContent(request *CreateGeneratedContentRequest) (_result *CreateGeneratedContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateGeneratedContentResponse{}
	_body, _err := client.CreateGeneratedContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateTokenWithOptions(request *CreateTokenRequest, runtime *dara.RuntimeOptions) (_result *CreateTokenResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateTokenResponse
func (client *Client) CreateToken(request *CreateTokenRequest) (_result *CreateTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTokenResponse{}
	_body, _err := client.CreateTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteAuditNoteWithOptions(request *DeleteAuditNoteRequest, runtime *dara.RuntimeOptions) (_result *DeleteAuditNoteResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteAuditNoteResponse
func (client *Client) DeleteAuditNote(request *DeleteAuditNoteRequest) (_result *DeleteAuditNoteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAuditNoteResponse{}
	_body, _err := client.DeleteAuditNoteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteAuditTermsWithOptions(tmpReq *DeleteAuditTermsRequest, runtime *dara.RuntimeOptions) (_result *DeleteAuditTermsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - DeleteAuditTermsRequest
//
// @return DeleteAuditTermsResponse
func (client *Client) DeleteAuditTerms(request *DeleteAuditTermsRequest) (_result *DeleteAuditTermsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAuditTermsResponse{}
	_body, _err := client.DeleteAuditTermsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteCustomTextWithOptions(request *DeleteCustomTextRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomTextResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteCustomTextResponse
func (client *Client) DeleteCustomText(request *DeleteCustomTextRequest) (_result *DeleteCustomTextResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCustomTextResponse{}
	_body, _err := client.DeleteCustomTextWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteCustomTopicByTopicWithOptions(request *DeleteCustomTopicByTopicRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomTopicByTopicResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteCustomTopicByTopicResponse
func (client *Client) DeleteCustomTopicByTopic(request *DeleteCustomTopicByTopicRequest) (_result *DeleteCustomTopicByTopicResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCustomTopicByTopicResponse{}
	_body, _err := client.DeleteCustomTopicByTopicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteCustomTopicViewPointByIdWithOptions(request *DeleteCustomTopicViewPointByIdRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomTopicViewPointByIdResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteCustomTopicViewPointByIdResponse
func (client *Client) DeleteCustomTopicViewPointById(request *DeleteCustomTopicViewPointByIdRequest) (_result *DeleteCustomTopicViewPointByIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCustomTopicViewPointByIdResponse{}
	_body, _err := client.DeleteCustomTopicViewPointByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteDataPermissionsWithOptions(tmpReq *DeleteDataPermissionsRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataPermissionsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - DeleteDataPermissionsRequest
//
// @return DeleteDataPermissionsResponse
func (client *Client) DeleteDataPermissions(request *DeleteDataPermissionsRequest) (_result *DeleteDataPermissionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataPermissionsResponse{}
	_body, _err := client.DeleteDataPermissionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteDatasetWithOptions(request *DeleteDatasetRequest, runtime *dara.RuntimeOptions) (_result *DeleteDatasetResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteDatasetResponse
func (client *Client) DeleteDataset(request *DeleteDatasetRequest) (_result *DeleteDatasetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDatasetResponse{}
	_body, _err := client.DeleteDatasetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteDatasetDocumentWithOptions(request *DeleteDatasetDocumentRequest, runtime *dara.RuntimeOptions) (_result *DeleteDatasetDocumentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteDatasetDocumentResponse
func (client *Client) DeleteDatasetDocument(request *DeleteDatasetDocumentRequest) (_result *DeleteDatasetDocumentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDatasetDocumentResponse{}
	_body, _err := client.DeleteDatasetDocumentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteDocsWithOptions(tmpReq *DeleteDocsRequest, runtime *dara.RuntimeOptions) (_result *DeleteDocsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - DeleteDocsRequest
//
// @return DeleteDocsResponse
func (client *Client) DeleteDocs(request *DeleteDocsRequest) (_result *DeleteDocsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDocsResponse{}
	_body, _err := client.DeleteDocsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteFactAuditUrlWithOptions(request *DeleteFactAuditUrlRequest, runtime *dara.RuntimeOptions) (_result *DeleteFactAuditUrlResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteFactAuditUrlResponse
func (client *Client) DeleteFactAuditUrl(request *DeleteFactAuditUrlRequest) (_result *DeleteFactAuditUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFactAuditUrlResponse{}
	_body, _err := client.DeleteFactAuditUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteGeneralConfigWithOptions(request *DeleteGeneralConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteGeneralConfigResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteGeneralConfigResponse
func (client *Client) DeleteGeneralConfig(request *DeleteGeneralConfigRequest) (_result *DeleteGeneralConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteGeneralConfigResponse{}
	_body, _err := client.DeleteGeneralConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteGeneratedContentWithOptions(request *DeleteGeneratedContentRequest, runtime *dara.RuntimeOptions) (_result *DeleteGeneratedContentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteGeneratedContentResponse
func (client *Client) DeleteGeneratedContent(request *DeleteGeneratedContentRequest) (_result *DeleteGeneratedContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteGeneratedContentResponse{}
	_body, _err := client.DeleteGeneratedContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteInterveneRuleWithOptions(request *DeleteInterveneRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteInterveneRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteInterveneRuleResponse
func (client *Client) DeleteInterveneRule(request *DeleteInterveneRuleRequest) (_result *DeleteInterveneRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteInterveneRuleResponse{}
	_body, _err := client.DeleteInterveneRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteMaterialByIdWithOptions(request *DeleteMaterialByIdRequest, runtime *dara.RuntimeOptions) (_result *DeleteMaterialByIdResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteMaterialByIdResponse
func (client *Client) DeleteMaterialById(request *DeleteMaterialByIdRequest) (_result *DeleteMaterialByIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMaterialByIdResponse{}
	_body, _err := client.DeleteMaterialByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeletePptArtifactWithOptions(request *DeletePptArtifactRequest, runtime *dara.RuntimeOptions) (_result *DeletePptArtifactResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeletePptArtifactResponse
func (client *Client) DeletePptArtifact(request *DeletePptArtifactRequest) (_result *DeletePptArtifactResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePptArtifactResponse{}
	_body, _err := client.DeletePptArtifactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteStyleLearningResultWithOptions(request *DeleteStyleLearningResultRequest, runtime *dara.RuntimeOptions) (_result *DeleteStyleLearningResultResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteStyleLearningResultResponse
func (client *Client) DeleteStyleLearningResult(request *DeleteStyleLearningResultRequest) (_result *DeleteStyleLearningResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteStyleLearningResultResponse{}
	_body, _err := client.DeleteStyleLearningResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DocumentExtractionWithOptions(tmpReq *DocumentExtractionRequest, runtime *dara.RuntimeOptions) (_result *DocumentExtractionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - DocumentExtractionRequest
//
// @return DocumentExtractionResponse
func (client *Client) DocumentExtraction(request *DocumentExtractionRequest) (_result *DocumentExtractionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DocumentExtractionResponse{}
	_body, _err := client.DocumentExtractionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DownloadAuditNoteWithOptions(request *DownloadAuditNoteRequest, runtime *dara.RuntimeOptions) (_result *DownloadAuditNoteResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DownloadAuditNoteResponse
func (client *Client) DownloadAuditNote(request *DownloadAuditNoteRequest) (_result *DownloadAuditNoteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DownloadAuditNoteResponse{}
	_body, _err := client.DownloadAuditNoteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DownloadBiddingDocWithOptions(request *DownloadBiddingDocRequest, runtime *dara.RuntimeOptions) (_result *DownloadBiddingDocResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DownloadBiddingDocResponse
func (client *Client) DownloadBiddingDoc(request *DownloadBiddingDocRequest) (_result *DownloadBiddingDocResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DownloadBiddingDocResponse{}
	_body, _err := client.DownloadBiddingDocWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) EditAuditTermsWithOptions(tmpReq *EditAuditTermsRequest, runtime *dara.RuntimeOptions) (_result *EditAuditTermsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - EditAuditTermsRequest
//
// @return EditAuditTermsResponse
func (client *Client) EditAuditTerms(request *EditAuditTermsRequest) (_result *EditAuditTermsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EditAuditTermsResponse{}
	_body, _err := client.EditAuditTermsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) EditBiddingDocWithOptions(request *EditBiddingDocRequest, runtime *dara.RuntimeOptions) (_result *EditBiddingDocResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return EditBiddingDocResponse
func (client *Client) EditBiddingDoc(request *EditBiddingDocRequest) (_result *EditBiddingDocResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EditBiddingDocResponse{}
	_body, _err := client.EditBiddingDocWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ExportAnalysisTagDetailByTaskIdWithOptions(tmpReq *ExportAnalysisTagDetailByTaskIdRequest, runtime *dara.RuntimeOptions) (_result *ExportAnalysisTagDetailByTaskIdResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ExportAnalysisTagDetailByTaskIdRequest
//
// @return ExportAnalysisTagDetailByTaskIdResponse
func (client *Client) ExportAnalysisTagDetailByTaskId(request *ExportAnalysisTagDetailByTaskIdRequest) (_result *ExportAnalysisTagDetailByTaskIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportAnalysisTagDetailByTaskIdResponse{}
	_body, _err := client.ExportAnalysisTagDetailByTaskIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ExportAuditContentResultWithOptions(request *ExportAuditContentResultRequest, runtime *dara.RuntimeOptions) (_result *ExportAuditContentResultResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ExportAuditContentResultResponse
func (client *Client) ExportAuditContentResult(request *ExportAuditContentResultRequest) (_result *ExportAuditContentResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportAuditContentResultResponse{}
	_body, _err := client.ExportAuditContentResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ExportCustomSourceAnalysisTaskWithOptions(request *ExportCustomSourceAnalysisTaskRequest, runtime *dara.RuntimeOptions) (_result *ExportCustomSourceAnalysisTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ExportCustomSourceAnalysisTaskResponse
func (client *Client) ExportCustomSourceAnalysisTask(request *ExportCustomSourceAnalysisTaskRequest) (_result *ExportCustomSourceAnalysisTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportCustomSourceAnalysisTaskResponse{}
	_body, _err := client.ExportCustomSourceAnalysisTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ExportGeneratedContentWithOptions(request *ExportGeneratedContentRequest, runtime *dara.RuntimeOptions) (_result *ExportGeneratedContentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ExportGeneratedContentResponse
func (client *Client) ExportGeneratedContent(request *ExportGeneratedContentRequest) (_result *ExportGeneratedContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportGeneratedContentResponse{}
	_body, _err := client.ExportGeneratedContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ExportHotTopicPlanningProposalsWithOptions(tmpReq *ExportHotTopicPlanningProposalsRequest, runtime *dara.RuntimeOptions) (_result *ExportHotTopicPlanningProposalsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ExportHotTopicPlanningProposalsRequest
//
// @return ExportHotTopicPlanningProposalsResponse
func (client *Client) ExportHotTopicPlanningProposals(request *ExportHotTopicPlanningProposalsRequest) (_result *ExportHotTopicPlanningProposalsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportHotTopicPlanningProposalsResponse{}
	_body, _err := client.ExportHotTopicPlanningProposalsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ExportIntervenesWithOptions(request *ExportIntervenesRequest, runtime *dara.RuntimeOptions) (_result *ExportIntervenesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ExportIntervenesResponse
func (client *Client) ExportIntervenes(request *ExportIntervenesRequest) (_result *ExportIntervenesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportIntervenesResponse{}
	_body, _err := client.ExportIntervenesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ExportPptArtifactWithOptions(request *ExportPptArtifactRequest, runtime *dara.RuntimeOptions) (_result *ExportPptArtifactResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ExportPptArtifactResponse
func (client *Client) ExportPptArtifact(request *ExportPptArtifactRequest) (_result *ExportPptArtifactResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportPptArtifactResponse{}
	_body, _err := client.ExportPptArtifactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) FeedbackDialogueWithOptions(tmpReq *FeedbackDialogueRequest, runtime *dara.RuntimeOptions) (_result *FeedbackDialogueResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - FeedbackDialogueRequest
//
// @return FeedbackDialogueResponse
func (client *Client) FeedbackDialogue(request *FeedbackDialogueRequest) (_result *FeedbackDialogueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FeedbackDialogueResponse{}
	_body, _err := client.FeedbackDialogueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) FetchExportTermsTaskWithOptions(request *FetchExportTermsTaskRequest, runtime *dara.RuntimeOptions) (_result *FetchExportTermsTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return FetchExportTermsTaskResponse
func (client *Client) FetchExportTermsTask(request *FetchExportTermsTaskRequest) (_result *FetchExportTermsTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FetchExportTermsTaskResponse{}
	_body, _err := client.FetchExportTermsTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) FetchExportWordTaskWithOptions(request *FetchExportWordTaskRequest, runtime *dara.RuntimeOptions) (_result *FetchExportWordTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return FetchExportWordTaskResponse
func (client *Client) FetchExportWordTask(request *FetchExportWordTaskRequest) (_result *FetchExportWordTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FetchExportWordTaskResponse{}
	_body, _err := client.FetchExportWordTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) FetchImageTaskWithOptions(tmpReq *FetchImageTaskRequest, runtime *dara.RuntimeOptions) (_result *FetchImageTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - FetchImageTaskRequest
//
// @return FetchImageTaskResponse
func (client *Client) FetchImageTask(request *FetchImageTaskRequest) (_result *FetchImageTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FetchImageTaskResponse{}
	_body, _err := client.FetchImageTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) FetchImportTermsTaskWithOptions(request *FetchImportTermsTaskRequest, runtime *dara.RuntimeOptions) (_result *FetchImportTermsTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return FetchImportTermsTaskResponse
func (client *Client) FetchImportTermsTask(request *FetchImportTermsTaskRequest) (_result *FetchImportTermsTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FetchImportTermsTaskResponse{}
	_body, _err := client.FetchImportTermsTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) FetchParseDocumentLayoutTaskWithOptions(request *FetchParseDocumentLayoutTaskRequest, runtime *dara.RuntimeOptions) (_result *FetchParseDocumentLayoutTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return FetchParseDocumentLayoutTaskResponse
func (client *Client) FetchParseDocumentLayoutTask(request *FetchParseDocumentLayoutTaskRequest) (_result *FetchParseDocumentLayoutTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FetchParseDocumentLayoutTaskResponse{}
	_body, _err := client.FetchParseDocumentLayoutTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GenerateExportWordTaskWithOptions(request *GenerateExportWordTaskRequest, runtime *dara.RuntimeOptions) (_result *GenerateExportWordTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GenerateExportWordTaskResponse
func (client *Client) GenerateExportWordTask(request *GenerateExportWordTaskRequest) (_result *GenerateExportWordTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateExportWordTaskResponse{}
	_body, _err := client.GenerateExportWordTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GenerateFileUrlByKeyWithOptions(request *GenerateFileUrlByKeyRequest, runtime *dara.RuntimeOptions) (_result *GenerateFileUrlByKeyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GenerateFileUrlByKeyResponse
func (client *Client) GenerateFileUrlByKey(request *GenerateFileUrlByKeyRequest) (_result *GenerateFileUrlByKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateFileUrlByKeyResponse{}
	_body, _err := client.GenerateFileUrlByKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GenerateImageTaskWithOptions(tmpReq *GenerateImageTaskRequest, runtime *dara.RuntimeOptions) (_result *GenerateImageTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GenerateImageTaskRequest
//
// @return GenerateImageTaskResponse
func (client *Client) GenerateImageTask(request *GenerateImageTaskRequest) (_result *GenerateImageTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateImageTaskResponse{}
	_body, _err := client.GenerateImageTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GenerateUploadConfigWithOptions(request *GenerateUploadConfigRequest, runtime *dara.RuntimeOptions) (_result *GenerateUploadConfigResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GenerateUploadConfigResponse
func (client *Client) GenerateUploadConfig(request *GenerateUploadConfigRequest) (_result *GenerateUploadConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateUploadConfigResponse{}
	_body, _err := client.GenerateUploadConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GenerateViewPointWithOptions(tmpReq *GenerateViewPointRequest, runtime *dara.RuntimeOptions) (_result *GenerateViewPointResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GenerateViewPointRequest
//
// @return GenerateViewPointResponse
func (client *Client) GenerateViewPoint(request *GenerateViewPointRequest) (_result *GenerateViewPointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateViewPointResponse{}
	_body, _err := client.GenerateViewPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetAuditNotePostProcessingStatusWithOptions(request *GetAuditNotePostProcessingStatusRequest, runtime *dara.RuntimeOptions) (_result *GetAuditNotePostProcessingStatusResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetAuditNotePostProcessingStatusResponse
func (client *Client) GetAuditNotePostProcessingStatus(request *GetAuditNotePostProcessingStatusRequest) (_result *GetAuditNotePostProcessingStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAuditNotePostProcessingStatusResponse{}
	_body, _err := client.GetAuditNotePostProcessingStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetAuditNoteProcessingStatusWithOptions(request *GetAuditNoteProcessingStatusRequest, runtime *dara.RuntimeOptions) (_result *GetAuditNoteProcessingStatusResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetAuditNoteProcessingStatusResponse
func (client *Client) GetAuditNoteProcessingStatus(request *GetAuditNoteProcessingStatusRequest) (_result *GetAuditNoteProcessingStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAuditNoteProcessingStatusResponse{}
	_body, _err := client.GetAuditNoteProcessingStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetAutoClipsTaskInfoWithOptions(request *GetAutoClipsTaskInfoRequest, runtime *dara.RuntimeOptions) (_result *GetAutoClipsTaskInfoResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetAutoClipsTaskInfoResponse
func (client *Client) GetAutoClipsTaskInfo(request *GetAutoClipsTaskInfoRequest) (_result *GetAutoClipsTaskInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAutoClipsTaskInfoResponse{}
	_body, _err := client.GetAutoClipsTaskInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetAvailableAuditNotesWithOptions(request *GetAvailableAuditNotesRequest, runtime *dara.RuntimeOptions) (_result *GetAvailableAuditNotesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetAvailableAuditNotesResponse
func (client *Client) GetAvailableAuditNotes(request *GetAvailableAuditNotesRequest) (_result *GetAvailableAuditNotesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAvailableAuditNotesResponse{}
	_body, _err := client.GetAvailableAuditNotesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetBiddingDocInfoWithOptions(request *GetBiddingDocInfoRequest, runtime *dara.RuntimeOptions) (_result *GetBiddingDocInfoResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetBiddingDocInfoResponse
func (client *Client) GetBiddingDocInfo(request *GetBiddingDocInfoRequest) (_result *GetBiddingDocInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBiddingDocInfoResponse{}
	_body, _err := client.GetBiddingDocInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetBiddingRemainLimitNumWithOptions(request *GetBiddingRemainLimitNumRequest, runtime *dara.RuntimeOptions) (_result *GetBiddingRemainLimitNumResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetBiddingRemainLimitNumResponse
func (client *Client) GetBiddingRemainLimitNum(request *GetBiddingRemainLimitNumRequest) (_result *GetBiddingRemainLimitNumResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBiddingRemainLimitNumResponse{}
	_body, _err := client.GetBiddingRemainLimitNumWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetCategoriesByTaskIdWithOptions(request *GetCategoriesByTaskIdRequest, runtime *dara.RuntimeOptions) (_result *GetCategoriesByTaskIdResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetCategoriesByTaskIdResponse
func (client *Client) GetCategoriesByTaskId(request *GetCategoriesByTaskIdRequest) (_result *GetCategoriesByTaskIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCategoriesByTaskIdResponse{}
	_body, _err := client.GetCategoriesByTaskIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetClipsBuildInResourceWithOptions(request *GetClipsBuildInResourceRequest, runtime *dara.RuntimeOptions) (_result *GetClipsBuildInResourceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetClipsBuildInResourceResponse
func (client *Client) GetClipsBuildInResource(request *GetClipsBuildInResourceRequest) (_result *GetClipsBuildInResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetClipsBuildInResourceResponse{}
	_body, _err := client.GetClipsBuildInResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetCustomHotTopicBroadcastJobWithOptions(request *GetCustomHotTopicBroadcastJobRequest, runtime *dara.RuntimeOptions) (_result *GetCustomHotTopicBroadcastJobResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetCustomHotTopicBroadcastJobResponse
func (client *Client) GetCustomHotTopicBroadcastJob(request *GetCustomHotTopicBroadcastJobRequest) (_result *GetCustomHotTopicBroadcastJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCustomHotTopicBroadcastJobResponse{}
	_body, _err := client.GetCustomHotTopicBroadcastJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetCustomSourceTopicAnalysisTaskWithOptions(request *GetCustomSourceTopicAnalysisTaskRequest, runtime *dara.RuntimeOptions) (_result *GetCustomSourceTopicAnalysisTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetCustomSourceTopicAnalysisTaskResponse
func (client *Client) GetCustomSourceTopicAnalysisTask(request *GetCustomSourceTopicAnalysisTaskRequest) (_result *GetCustomSourceTopicAnalysisTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCustomSourceTopicAnalysisTaskResponse{}
	_body, _err := client.GetCustomSourceTopicAnalysisTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetCustomTextWithOptions(request *GetCustomTextRequest, runtime *dara.RuntimeOptions) (_result *GetCustomTextResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetCustomTextResponse
func (client *Client) GetCustomText(request *GetCustomTextRequest) (_result *GetCustomTextResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCustomTextResponse{}
	_body, _err := client.GetCustomTextWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetCustomTopicSelectionPerspectiveAnalysisTaskWithOptions(request *GetCustomTopicSelectionPerspectiveAnalysisTaskRequest, runtime *dara.RuntimeOptions) (_result *GetCustomTopicSelectionPerspectiveAnalysisTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetCustomTopicSelectionPerspectiveAnalysisTaskResponse
func (client *Client) GetCustomTopicSelectionPerspectiveAnalysisTask(request *GetCustomTopicSelectionPerspectiveAnalysisTaskRequest) (_result *GetCustomTopicSelectionPerspectiveAnalysisTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCustomTopicSelectionPerspectiveAnalysisTaskResponse{}
	_body, _err := client.GetCustomTopicSelectionPerspectiveAnalysisTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetDataSourceOrderConfigWithOptions(request *GetDataSourceOrderConfigRequest, runtime *dara.RuntimeOptions) (_result *GetDataSourceOrderConfigResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetDataSourceOrderConfigResponse
func (client *Client) GetDataSourceOrderConfig(request *GetDataSourceOrderConfigRequest) (_result *GetDataSourceOrderConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataSourceOrderConfigResponse{}
	_body, _err := client.GetDataSourceOrderConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetDatasetWithOptions(request *GetDatasetRequest, runtime *dara.RuntimeOptions) (_result *GetDatasetResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetDatasetResponse
func (client *Client) GetDataset(request *GetDatasetRequest) (_result *GetDatasetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDatasetResponse{}
	_body, _err := client.GetDatasetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetDatasetDocumentWithOptions(tmpReq *GetDatasetDocumentRequest, runtime *dara.RuntimeOptions) (_result *GetDatasetDocumentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetDatasetDocumentRequest
//
// @return GetDatasetDocumentResponse
func (client *Client) GetDatasetDocument(request *GetDatasetDocumentRequest) (_result *GetDatasetDocumentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDatasetDocumentResponse{}
	_body, _err := client.GetDatasetDocumentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetDeepWriteTaskWithOptions(request *GetDeepWriteTaskRequest, runtime *dara.RuntimeOptions) (_result *GetDeepWriteTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetDeepWriteTaskResponse
func (client *Client) GetDeepWriteTask(request *GetDeepWriteTaskRequest) (_result *GetDeepWriteTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDeepWriteTaskResponse{}
	_body, _err := client.GetDeepWriteTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetDeepWriteTaskResultWithOptions(request *GetDeepWriteTaskResultRequest, runtime *dara.RuntimeOptions) (_result *GetDeepWriteTaskResultResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetDeepWriteTaskResultResponse
func (client *Client) GetDeepWriteTaskResult(request *GetDeepWriteTaskResultRequest) (_result *GetDeepWriteTaskResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDeepWriteTaskResultResponse{}
	_body, _err := client.GetDeepWriteTaskResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetDocClusterTaskWithOptions(request *GetDocClusterTaskRequest, runtime *dara.RuntimeOptions) (_result *GetDocClusterTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetDocClusterTaskResponse
func (client *Client) GetDocClusterTask(request *GetDocClusterTaskRequest) (_result *GetDocClusterTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDocClusterTaskResponse{}
	_body, _err := client.GetDocClusterTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetDocInfoWithOptions(request *GetDocInfoRequest, runtime *dara.RuntimeOptions) (_result *GetDocInfoResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetDocInfoResponse
func (client *Client) GetDocInfo(request *GetDocInfoRequest) (_result *GetDocInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDocInfoResponse{}
	_body, _err := client.GetDocInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetEnterpriseVocAnalysisTaskWithOptions(request *GetEnterpriseVocAnalysisTaskRequest, runtime *dara.RuntimeOptions) (_result *GetEnterpriseVocAnalysisTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetEnterpriseVocAnalysisTaskResponse
func (client *Client) GetEnterpriseVocAnalysisTask(request *GetEnterpriseVocAnalysisTaskRequest) (_result *GetEnterpriseVocAnalysisTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetEnterpriseVocAnalysisTaskResponse{}
	_body, _err := client.GetEnterpriseVocAnalysisTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetFactAuditUrlWithOptions(request *GetFactAuditUrlRequest, runtime *dara.RuntimeOptions) (_result *GetFactAuditUrlResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetFactAuditUrlResponse
func (client *Client) GetFactAuditUrl(request *GetFactAuditUrlRequest) (_result *GetFactAuditUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetFactAuditUrlResponse{}
	_body, _err := client.GetFactAuditUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetFileContentLengthWithOptions(request *GetFileContentLengthRequest, runtime *dara.RuntimeOptions) (_result *GetFileContentLengthResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetFileContentLengthResponse
func (client *Client) GetFileContentLength(request *GetFileContentLengthRequest) (_result *GetFileContentLengthResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetFileContentLengthResponse{}
	_body, _err := client.GetFileContentLengthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetGeneralConfigWithOptions(request *GetGeneralConfigRequest, runtime *dara.RuntimeOptions) (_result *GetGeneralConfigResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetGeneralConfigResponse
func (client *Client) GetGeneralConfig(request *GetGeneralConfigRequest) (_result *GetGeneralConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetGeneralConfigResponse{}
	_body, _err := client.GetGeneralConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetGeneratedContentWithOptions(request *GetGeneratedContentRequest, runtime *dara.RuntimeOptions) (_result *GetGeneratedContentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetGeneratedContentResponse
func (client *Client) GetGeneratedContent(request *GetGeneratedContentRequest) (_result *GetGeneratedContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetGeneratedContentResponse{}
	_body, _err := client.GetGeneratedContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetHotTopicBroadcastWithOptions(tmpReq *GetHotTopicBroadcastRequest, runtime *dara.RuntimeOptions) (_result *GetHotTopicBroadcastResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetHotTopicBroadcastRequest
//
// @return GetHotTopicBroadcastResponse
func (client *Client) GetHotTopicBroadcast(request *GetHotTopicBroadcastRequest) (_result *GetHotTopicBroadcastResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetHotTopicBroadcastResponse{}
	_body, _err := client.GetHotTopicBroadcastWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetInterveneGlobalReplyWithOptions(request *GetInterveneGlobalReplyRequest, runtime *dara.RuntimeOptions) (_result *GetInterveneGlobalReplyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetInterveneGlobalReplyResponse
func (client *Client) GetInterveneGlobalReply(request *GetInterveneGlobalReplyRequest) (_result *GetInterveneGlobalReplyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInterveneGlobalReplyResponse{}
	_body, _err := client.GetInterveneGlobalReplyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetInterveneImportTaskInfoWithOptions(request *GetInterveneImportTaskInfoRequest, runtime *dara.RuntimeOptions) (_result *GetInterveneImportTaskInfoResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetInterveneImportTaskInfoResponse
func (client *Client) GetInterveneImportTaskInfo(request *GetInterveneImportTaskInfoRequest) (_result *GetInterveneImportTaskInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInterveneImportTaskInfoResponse{}
	_body, _err := client.GetInterveneImportTaskInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetInterveneRuleDetailWithOptions(request *GetInterveneRuleDetailRequest, runtime *dara.RuntimeOptions) (_result *GetInterveneRuleDetailResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetInterveneRuleDetailResponse
func (client *Client) GetInterveneRuleDetail(request *GetInterveneRuleDetailRequest) (_result *GetInterveneRuleDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInterveneRuleDetailResponse{}
	_body, _err := client.GetInterveneRuleDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetInterveneTemplateFileUrlWithOptions(request *GetInterveneTemplateFileUrlRequest, runtime *dara.RuntimeOptions) (_result *GetInterveneTemplateFileUrlResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetInterveneTemplateFileUrlResponse
func (client *Client) GetInterveneTemplateFileUrl(request *GetInterveneTemplateFileUrlRequest) (_result *GetInterveneTemplateFileUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInterveneTemplateFileUrlResponse{}
	_body, _err := client.GetInterveneTemplateFileUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetMaterialByIdWithOptions(request *GetMaterialByIdRequest, runtime *dara.RuntimeOptions) (_result *GetMaterialByIdResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetMaterialByIdResponse
func (client *Client) GetMaterialById(request *GetMaterialByIdRequest) (_result *GetMaterialByIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMaterialByIdResponse{}
	_body, _err := client.GetMaterialByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetPptArtifactWithOptions(request *GetPptArtifactRequest, runtime *dara.RuntimeOptions) (_result *GetPptArtifactResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetPptArtifactResponse
func (client *Client) GetPptArtifact(request *GetPptArtifactRequest) (_result *GetPptArtifactResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPptArtifactResponse{}
	_body, _err := client.GetPptArtifactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetPptArtifactExportResultWithOptions(request *GetPptArtifactExportResultRequest, runtime *dara.RuntimeOptions) (_result *GetPptArtifactExportResultResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetPptArtifactExportResultResponse
func (client *Client) GetPptArtifactExportResult(request *GetPptArtifactExportResultRequest) (_result *GetPptArtifactExportResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPptArtifactExportResultResponse{}
	_body, _err := client.GetPptArtifactExportResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetPptConfigWithOptions(request *GetPptConfigRequest, runtime *dara.RuntimeOptions) (_result *GetPptConfigResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetPptConfigResponse
func (client *Client) GetPptConfig(request *GetPptConfigRequest) (_result *GetPptConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPptConfigResponse{}
	_body, _err := client.GetPptConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetPptInfoWithOptions(request *GetPptInfoRequest, runtime *dara.RuntimeOptions) (_result *GetPptInfoResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetPptInfoResponse
func (client *Client) GetPptInfo(request *GetPptInfoRequest) (_result *GetPptInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPptInfoResponse{}
	_body, _err := client.GetPptInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetPptTemplateSelectorWithOptions(request *GetPptTemplateSelectorRequest, runtime *dara.RuntimeOptions) (_result *GetPptTemplateSelectorResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetPptTemplateSelectorResponse
func (client *Client) GetPptTemplateSelector(request *GetPptTemplateSelectorRequest) (_result *GetPptTemplateSelectorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPptTemplateSelectorResponse{}
	_body, _err := client.GetPptTemplateSelectorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetPropertiesWithOptions(request *GetPropertiesRequest, runtime *dara.RuntimeOptions) (_result *GetPropertiesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetPropertiesResponse
func (client *Client) GetProperties(request *GetPropertiesRequest) (_result *GetPropertiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPropertiesResponse{}
	_body, _err := client.GetPropertiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetSmartAuditResultWithOptions(request *GetSmartAuditResultRequest, runtime *dara.RuntimeOptions) (_result *GetSmartAuditResultResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetSmartAuditResultResponse
func (client *Client) GetSmartAuditResult(request *GetSmartAuditResultRequest) (_result *GetSmartAuditResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSmartAuditResultResponse{}
	_body, _err := client.GetSmartAuditResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetSmartClipTaskWithOptions(request *GetSmartClipTaskRequest, runtime *dara.RuntimeOptions) (_result *GetSmartClipTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetSmartClipTaskResponse
func (client *Client) GetSmartClipTask(request *GetSmartClipTaskRequest) (_result *GetSmartClipTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSmartClipTaskResponse{}
	_body, _err := client.GetSmartClipTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetStyleLearningResultWithOptions(request *GetStyleLearningResultRequest, runtime *dara.RuntimeOptions) (_result *GetStyleLearningResultResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetStyleLearningResultResponse
func (client *Client) GetStyleLearningResult(request *GetStyleLearningResultRequest) (_result *GetStyleLearningResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetStyleLearningResultResponse{}
	_body, _err := client.GetStyleLearningResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetTopicByIdWithOptions(request *GetTopicByIdRequest, runtime *dara.RuntimeOptions) (_result *GetTopicByIdResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetTopicByIdResponse
func (client *Client) GetTopicById(request *GetTopicByIdRequest) (_result *GetTopicByIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTopicByIdResponse{}
	_body, _err := client.GetTopicByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetTopicSelectionPerspectiveAnalysisTaskWithOptions(request *GetTopicSelectionPerspectiveAnalysisTaskRequest, runtime *dara.RuntimeOptions) (_result *GetTopicSelectionPerspectiveAnalysisTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetTopicSelectionPerspectiveAnalysisTaskResponse
func (client *Client) GetTopicSelectionPerspectiveAnalysisTask(request *GetTopicSelectionPerspectiveAnalysisTaskRequest) (_result *GetTopicSelectionPerspectiveAnalysisTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTopicSelectionPerspectiveAnalysisTaskResponse{}
	_body, _err := client.GetTopicSelectionPerspectiveAnalysisTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ImportInterveneFileWithOptions(request *ImportInterveneFileRequest, runtime *dara.RuntimeOptions) (_result *ImportInterveneFileResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ImportInterveneFileResponse
func (client *Client) ImportInterveneFile(request *ImportInterveneFileRequest) (_result *ImportInterveneFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ImportInterveneFileResponse{}
	_body, _err := client.ImportInterveneFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ImportInterveneFileAsyncWithOptions(request *ImportInterveneFileAsyncRequest, runtime *dara.RuntimeOptions) (_result *ImportInterveneFileAsyncResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ImportInterveneFileAsyncResponse
func (client *Client) ImportInterveneFileAsync(request *ImportInterveneFileAsyncRequest) (_result *ImportInterveneFileAsyncResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ImportInterveneFileAsyncResponse{}
	_body, _err := client.ImportInterveneFileAsyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) InitiatePptCreationWithOptions(request *InitiatePptCreationRequest, runtime *dara.RuntimeOptions) (_result *InitiatePptCreationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return InitiatePptCreationResponse
func (client *Client) InitiatePptCreation(request *InitiatePptCreationRequest) (_result *InitiatePptCreationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InitiatePptCreationResponse{}
	_body, _err := client.InitiatePptCreationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) InitiatePptCreationV2WithOptions(request *InitiatePptCreationV2Request, runtime *dara.RuntimeOptions) (_result *InitiatePptCreationV2Response, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return InitiatePptCreationV2Response
func (client *Client) InitiatePptCreationV2(request *InitiatePptCreationV2Request) (_result *InitiatePptCreationV2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InitiatePptCreationV2Response{}
	_body, _err := client.InitiatePptCreationV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) InsertInterveneGlobalReplyWithOptions(tmpReq *InsertInterveneGlobalReplyRequest, runtime *dara.RuntimeOptions) (_result *InsertInterveneGlobalReplyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - InsertInterveneGlobalReplyRequest
//
// @return InsertInterveneGlobalReplyResponse
func (client *Client) InsertInterveneGlobalReply(request *InsertInterveneGlobalReplyRequest) (_result *InsertInterveneGlobalReplyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InsertInterveneGlobalReplyResponse{}
	_body, _err := client.InsertInterveneGlobalReplyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) InsertInterveneRuleWithOptions(tmpReq *InsertInterveneRuleRequest, runtime *dara.RuntimeOptions) (_result *InsertInterveneRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - InsertInterveneRuleRequest
//
// @return InsertInterveneRuleResponse
func (client *Client) InsertInterveneRule(request *InsertInterveneRuleRequest) (_result *InsertInterveneRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InsertInterveneRuleResponse{}
	_body, _err := client.InsertInterveneRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListAnalysisTagDetailByTaskIdWithOptions(tmpReq *ListAnalysisTagDetailByTaskIdRequest, runtime *dara.RuntimeOptions) (_result *ListAnalysisTagDetailByTaskIdResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListAnalysisTagDetailByTaskIdRequest
//
// @return ListAnalysisTagDetailByTaskIdResponse
func (client *Client) ListAnalysisTagDetailByTaskId(request *ListAnalysisTagDetailByTaskIdRequest) (_result *ListAnalysisTagDetailByTaskIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAnalysisTagDetailByTaskIdResponse{}
	_body, _err := client.ListAnalysisTagDetailByTaskIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListAsyncTasksWithOptions(tmpReq *ListAsyncTasksRequest, runtime *dara.RuntimeOptions) (_result *ListAsyncTasksResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListAsyncTasksRequest
//
// @return ListAsyncTasksResponse
func (client *Client) ListAsyncTasks(request *ListAsyncTasksRequest) (_result *ListAsyncTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAsyncTasksResponse{}
	_body, _err := client.ListAsyncTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListAuditContentErrorTypesWithOptions(request *ListAuditContentErrorTypesRequest, runtime *dara.RuntimeOptions) (_result *ListAuditContentErrorTypesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListAuditContentErrorTypesResponse
func (client *Client) ListAuditContentErrorTypes(request *ListAuditContentErrorTypesRequest) (_result *ListAuditContentErrorTypesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAuditContentErrorTypesResponse{}
	_body, _err := client.ListAuditContentErrorTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListAuditTermsWithOptions(request *ListAuditTermsRequest, runtime *dara.RuntimeOptions) (_result *ListAuditTermsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListAuditTermsResponse
func (client *Client) ListAuditTerms(request *ListAuditTermsRequest) (_result *ListAuditTermsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAuditTermsResponse{}
	_body, _err := client.ListAuditTermsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListAutoClipsTaskWithOptions(request *ListAutoClipsTaskRequest, runtime *dara.RuntimeOptions) (_result *ListAutoClipsTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListAutoClipsTaskResponse
func (client *Client) ListAutoClipsTask(request *ListAutoClipsTaskRequest) (_result *ListAutoClipsTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAutoClipsTaskResponse{}
	_body, _err := client.ListAutoClipsTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListBiddingDocWithOptions(request *ListBiddingDocRequest, runtime *dara.RuntimeOptions) (_result *ListBiddingDocResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListBiddingDocResponse
func (client *Client) ListBiddingDoc(request *ListBiddingDocRequest) (_result *ListBiddingDocResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBiddingDocResponse{}
	_body, _err := client.ListBiddingDocWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListBuildConfigsWithOptions(request *ListBuildConfigsRequest, runtime *dara.RuntimeOptions) (_result *ListBuildConfigsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListBuildConfigsResponse
func (client *Client) ListBuildConfigs(request *ListBuildConfigsRequest) (_result *ListBuildConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBuildConfigsResponse{}
	_body, _err := client.ListBuildConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListCustomTextWithOptions(request *ListCustomTextRequest, runtime *dara.RuntimeOptions) (_result *ListCustomTextResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListCustomTextResponse
func (client *Client) ListCustomText(request *ListCustomTextRequest) (_result *ListCustomTextResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCustomTextResponse{}
	_body, _err := client.ListCustomTextWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListCustomViewPointsWithOptions(tmpReq *ListCustomViewPointsRequest, runtime *dara.RuntimeOptions) (_result *ListCustomViewPointsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListCustomViewPointsRequest
//
// @return ListCustomViewPointsResponse
func (client *Client) ListCustomViewPoints(request *ListCustomViewPointsRequest) (_result *ListCustomViewPointsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCustomViewPointsResponse{}
	_body, _err := client.ListCustomViewPointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListDataPermissionsWithOptions(request *ListDataPermissionsRequest, runtime *dara.RuntimeOptions) (_result *ListDataPermissionsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListDataPermissionsResponse
func (client *Client) ListDataPermissions(request *ListDataPermissionsRequest) (_result *ListDataPermissionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataPermissionsResponse{}
	_body, _err := client.ListDataPermissionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListDatasetDocumentsWithOptions(tmpReq *ListDatasetDocumentsRequest, runtime *dara.RuntimeOptions) (_result *ListDatasetDocumentsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListDatasetDocumentsRequest
//
// @return ListDatasetDocumentsResponse
func (client *Client) ListDatasetDocuments(request *ListDatasetDocumentsRequest) (_result *ListDatasetDocumentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDatasetDocumentsResponse{}
	_body, _err := client.ListDatasetDocumentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListDatasetsWithOptions(request *ListDatasetsRequest, runtime *dara.RuntimeOptions) (_result *ListDatasetsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListDatasetsResponse
func (client *Client) ListDatasets(request *ListDatasetsRequest) (_result *ListDatasetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDatasetsResponse{}
	_body, _err := client.ListDatasetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListDialoguesWithOptions(request *ListDialoguesRequest, runtime *dara.RuntimeOptions) (_result *ListDialoguesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListDialoguesResponse
func (client *Client) ListDialogues(request *ListDialoguesRequest) (_result *ListDialoguesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDialoguesResponse{}
	_body, _err := client.ListDialoguesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListDocsWithOptions(tmpReq *ListDocsRequest, runtime *dara.RuntimeOptions) (_result *ListDocsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListDocsRequest
//
// @return ListDocsResponse
func (client *Client) ListDocs(request *ListDocsRequest) (_result *ListDocsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDocsResponse{}
	_body, _err := client.ListDocsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListDocumentRetrieveWithOptions(request *ListDocumentRetrieveRequest, runtime *dara.RuntimeOptions) (_result *ListDocumentRetrieveResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListDocumentRetrieveResponse
func (client *Client) ListDocumentRetrieve(request *ListDocumentRetrieveRequest) (_result *ListDocumentRetrieveResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDocumentRetrieveResponse{}
	_body, _err := client.ListDocumentRetrieveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListEnterprisePptTemplatesWithOptions(request *ListEnterprisePptTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListEnterprisePptTemplatesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListEnterprisePptTemplatesResponse
func (client *Client) ListEnterprisePptTemplates(request *ListEnterprisePptTemplatesRequest) (_result *ListEnterprisePptTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEnterprisePptTemplatesResponse{}
	_body, _err := client.ListEnterprisePptTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListFreshViewPointsWithOptions(request *ListFreshViewPointsRequest, runtime *dara.RuntimeOptions) (_result *ListFreshViewPointsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListFreshViewPointsResponse
func (client *Client) ListFreshViewPoints(request *ListFreshViewPointsRequest) (_result *ListFreshViewPointsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListFreshViewPointsResponse{}
	_body, _err := client.ListFreshViewPointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListGeneralConfigsWithOptions(request *ListGeneralConfigsRequest, runtime *dara.RuntimeOptions) (_result *ListGeneralConfigsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListGeneralConfigsResponse
func (client *Client) ListGeneralConfigs(request *ListGeneralConfigsRequest) (_result *ListGeneralConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListGeneralConfigsResponse{}
	_body, _err := client.ListGeneralConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListGeneratedContentsWithOptions(request *ListGeneratedContentsRequest, runtime *dara.RuntimeOptions) (_result *ListGeneratedContentsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListGeneratedContentsResponse
func (client *Client) ListGeneratedContents(request *ListGeneratedContentsRequest) (_result *ListGeneratedContentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListGeneratedContentsResponse{}
	_body, _err := client.ListGeneratedContentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListHotNewsWithTypeWithOptions(tmpReq *ListHotNewsWithTypeRequest, runtime *dara.RuntimeOptions) (_result *ListHotNewsWithTypeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListHotNewsWithTypeRequest
//
// @return ListHotNewsWithTypeResponse
func (client *Client) ListHotNewsWithType(request *ListHotNewsWithTypeRequest) (_result *ListHotNewsWithTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHotNewsWithTypeResponse{}
	_body, _err := client.ListHotNewsWithTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListHotSourcesWithOptions(request *ListHotSourcesRequest, runtime *dara.RuntimeOptions) (_result *ListHotSourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListHotSourcesResponse
func (client *Client) ListHotSources(request *ListHotSourcesRequest) (_result *ListHotSourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHotSourcesResponse{}
	_body, _err := client.ListHotSourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListHotTopicsWithOptions(tmpReq *ListHotTopicsRequest, runtime *dara.RuntimeOptions) (_result *ListHotTopicsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListHotTopicsRequest
//
// @return ListHotTopicsResponse
func (client *Client) ListHotTopics(request *ListHotTopicsRequest) (_result *ListHotTopicsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHotTopicsResponse{}
	_body, _err := client.ListHotTopicsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListHotViewPointsWithOptions(request *ListHotViewPointsRequest, runtime *dara.RuntimeOptions) (_result *ListHotViewPointsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListHotViewPointsResponse
func (client *Client) ListHotViewPoints(request *ListHotViewPointsRequest) (_result *ListHotViewPointsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHotViewPointsResponse{}
	_body, _err := client.ListHotViewPointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListInterveneCntWithOptions(request *ListInterveneCntRequest, runtime *dara.RuntimeOptions) (_result *ListInterveneCntResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListInterveneCntResponse
func (client *Client) ListInterveneCnt(request *ListInterveneCntRequest) (_result *ListInterveneCntResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInterveneCntResponse{}
	_body, _err := client.ListInterveneCntWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListInterveneImportTasksWithOptions(request *ListInterveneImportTasksRequest, runtime *dara.RuntimeOptions) (_result *ListInterveneImportTasksResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListInterveneImportTasksResponse
func (client *Client) ListInterveneImportTasks(request *ListInterveneImportTasksRequest) (_result *ListInterveneImportTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInterveneImportTasksResponse{}
	_body, _err := client.ListInterveneImportTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListInterveneRulesWithOptions(request *ListInterveneRulesRequest, runtime *dara.RuntimeOptions) (_result *ListInterveneRulesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListInterveneRulesResponse
func (client *Client) ListInterveneRules(request *ListInterveneRulesRequest) (_result *ListInterveneRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInterveneRulesResponse{}
	_body, _err := client.ListInterveneRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListIntervenesWithOptions(request *ListIntervenesRequest, runtime *dara.RuntimeOptions) (_result *ListIntervenesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListIntervenesResponse
func (client *Client) ListIntervenes(request *ListIntervenesRequest) (_result *ListIntervenesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListIntervenesResponse{}
	_body, _err := client.ListIntervenesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListMaterialDocumentsWithOptions(tmpReq *ListMaterialDocumentsRequest, runtime *dara.RuntimeOptions) (_result *ListMaterialDocumentsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListMaterialDocumentsRequest
//
// @return ListMaterialDocumentsResponse
func (client *Client) ListMaterialDocuments(request *ListMaterialDocumentsRequest) (_result *ListMaterialDocumentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMaterialDocumentsResponse{}
	_body, _err := client.ListMaterialDocumentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListPlanningProposalWithOptions(tmpReq *ListPlanningProposalRequest, runtime *dara.RuntimeOptions) (_result *ListPlanningProposalResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListPlanningProposalRequest
//
// @return ListPlanningProposalResponse
func (client *Client) ListPlanningProposal(request *ListPlanningProposalRequest) (_result *ListPlanningProposalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPlanningProposalResponse{}
	_body, _err := client.ListPlanningProposalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListPptArtifactsWithOptions(request *ListPptArtifactsRequest, runtime *dara.RuntimeOptions) (_result *ListPptArtifactsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListPptArtifactsResponse
func (client *Client) ListPptArtifacts(request *ListPptArtifactsRequest) (_result *ListPptArtifactsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPptArtifactsResponse{}
	_body, _err := client.ListPptArtifactsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListPptTemplatesWithOptions(request *ListPptTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListPptTemplatesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListPptTemplatesResponse
func (client *Client) ListPptTemplates(request *ListPptTemplatesRequest) (_result *ListPptTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPptTemplatesResponse{}
	_body, _err := client.ListPptTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListSearchTaskDialogueDatasWithOptions(request *ListSearchTaskDialogueDatasRequest, runtime *dara.RuntimeOptions) (_result *ListSearchTaskDialogueDatasResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListSearchTaskDialogueDatasResponse
func (client *Client) ListSearchTaskDialogueDatas(request *ListSearchTaskDialogueDatasRequest) (_result *ListSearchTaskDialogueDatasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSearchTaskDialogueDatasResponse{}
	_body, _err := client.ListSearchTaskDialogueDatasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListSearchTaskDialoguesWithOptions(request *ListSearchTaskDialoguesRequest, runtime *dara.RuntimeOptions) (_result *ListSearchTaskDialoguesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListSearchTaskDialoguesResponse
func (client *Client) ListSearchTaskDialogues(request *ListSearchTaskDialoguesRequest) (_result *ListSearchTaskDialoguesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSearchTaskDialoguesResponse{}
	_body, _err := client.ListSearchTaskDialoguesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListSearchTasksWithOptions(tmpReq *ListSearchTasksRequest, runtime *dara.RuntimeOptions) (_result *ListSearchTasksResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListSearchTasksRequest
//
// @return ListSearchTasksResponse
func (client *Client) ListSearchTasks(request *ListSearchTasksRequest) (_result *ListSearchTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSearchTasksResponse{}
	_body, _err := client.ListSearchTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListStyleLearningResultWithOptions(request *ListStyleLearningResultRequest, runtime *dara.RuntimeOptions) (_result *ListStyleLearningResultResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListStyleLearningResultResponse
func (client *Client) ListStyleLearningResult(request *ListStyleLearningResultRequest) (_result *ListStyleLearningResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListStyleLearningResultResponse{}
	_body, _err := client.ListStyleLearningResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListTimedViewAttitudeWithOptions(request *ListTimedViewAttitudeRequest, runtime *dara.RuntimeOptions) (_result *ListTimedViewAttitudeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListTimedViewAttitudeResponse
func (client *Client) ListTimedViewAttitude(request *ListTimedViewAttitudeRequest) (_result *ListTimedViewAttitudeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTimedViewAttitudeResponse{}
	_body, _err := client.ListTimedViewAttitudeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListTopicRecommendEventListWithOptions(request *ListTopicRecommendEventListRequest, runtime *dara.RuntimeOptions) (_result *ListTopicRecommendEventListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListTopicRecommendEventListResponse
func (client *Client) ListTopicRecommendEventList(request *ListTopicRecommendEventListRequest) (_result *ListTopicRecommendEventListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTopicRecommendEventListResponse{}
	_body, _err := client.ListTopicRecommendEventListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListTopicViewPointRecommendEventListWithOptions(request *ListTopicViewPointRecommendEventListRequest, runtime *dara.RuntimeOptions) (_result *ListTopicViewPointRecommendEventListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListTopicViewPointRecommendEventListResponse
func (client *Client) ListTopicViewPointRecommendEventList(request *ListTopicViewPointRecommendEventListRequest) (_result *ListTopicViewPointRecommendEventListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTopicViewPointRecommendEventListResponse{}
	_body, _err := client.ListTopicViewPointRecommendEventListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListVersionsWithOptions(request *ListVersionsRequest, runtime *dara.RuntimeOptions) (_result *ListVersionsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListVersionsResponse
func (client *Client) ListVersions(request *ListVersionsRequest) (_result *ListVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListVersionsResponse{}
	_body, _err := client.ListVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListWebReviewPointsWithOptions(request *ListWebReviewPointsRequest, runtime *dara.RuntimeOptions) (_result *ListWebReviewPointsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListWebReviewPointsResponse
func (client *Client) ListWebReviewPoints(request *ListWebReviewPointsRequest) (_result *ListWebReviewPointsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListWebReviewPointsResponse{}
	_body, _err := client.ListWebReviewPointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListWritingStylesWithOptions(request *ListWritingStylesRequest, runtime *dara.RuntimeOptions) (_result *ListWritingStylesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListWritingStylesResponse
func (client *Client) ListWritingStyles(request *ListWritingStylesRequest) (_result *ListWritingStylesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListWritingStylesResponse{}
	_body, _err := client.ListWritingStylesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryAsyncTaskWithOptions(request *QueryAsyncTaskRequest, runtime *dara.RuntimeOptions) (_result *QueryAsyncTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryAsyncTaskResponse
func (client *Client) QueryAsyncTask(request *QueryAsyncTaskRequest) (_result *QueryAsyncTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryAsyncTaskResponse{}
	_body, _err := client.QueryAsyncTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryAuditTaskWithOptions(request *QueryAuditTaskRequest, runtime *dara.RuntimeOptions) (_result *QueryAuditTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryAuditTaskResponse
func (client *Client) QueryAuditTask(request *QueryAuditTaskRequest) (_result *QueryAuditTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryAuditTaskResponse{}
	_body, _err := client.QueryAuditTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryVideoAuditResultWithOptions(request *QueryVideoAuditResultRequest, runtime *dara.RuntimeOptions) (_result *QueryVideoAuditResultResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryVideoAuditResultResponse
func (client *Client) QueryVideoAuditResult(request *QueryVideoAuditResultRequest) (_result *QueryVideoAuditResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryVideoAuditResultResponse{}
	_body, _err := client.QueryVideoAuditResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunAbbreviationContentWithSSE(request *RunAbbreviationContentRequest, runtime *dara.RuntimeOptions, _yield chan *RunAbbreviationContentResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runAbbreviationContentWithSSE_opYieldFunc(_yield, _yieldErr, request, runtime)
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
func (client *Client) RunAbbreviationContentWithOptions(request *RunAbbreviationContentRequest, runtime *dara.RuntimeOptions) (_result *RunAbbreviationContentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RunAbbreviationContentResponse
func (client *Client) RunAbbreviationContent(request *RunAbbreviationContentRequest) (_result *RunAbbreviationContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunAbbreviationContentResponse{}
	_body, _err := client.RunAbbreviationContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunAiHelperWritingWithSSE(tmpReq *RunAiHelperWritingRequest, runtime *dara.RuntimeOptions, _yield chan *RunAiHelperWritingResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runAiHelperWritingWithSSE_opYieldFunc(_yield, _yieldErr, tmpReq, runtime)
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
func (client *Client) RunAiHelperWritingWithOptions(tmpReq *RunAiHelperWritingRequest, runtime *dara.RuntimeOptions) (_result *RunAiHelperWritingResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - RunAiHelperWritingRequest
//
// @return RunAiHelperWritingResponse
func (client *Client) RunAiHelperWriting(request *RunAiHelperWritingRequest) (_result *RunAiHelperWritingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunAiHelperWritingResponse{}
	_body, _err := client.RunAiHelperWritingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunBookBrainmapWithSSE(request *RunBookBrainmapRequest, runtime *dara.RuntimeOptions, _yield chan *RunBookBrainmapResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runBookBrainmapWithSSE_opYieldFunc(_yield, _yieldErr, request, runtime)
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
func (client *Client) RunBookBrainmapWithOptions(request *RunBookBrainmapRequest, runtime *dara.RuntimeOptions) (_result *RunBookBrainmapResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RunBookBrainmapResponse
func (client *Client) RunBookBrainmap(request *RunBookBrainmapRequest) (_result *RunBookBrainmapResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunBookBrainmapResponse{}
	_body, _err := client.RunBookBrainmapWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunBookIntroductionWithSSE(request *RunBookIntroductionRequest, runtime *dara.RuntimeOptions, _yield chan *RunBookIntroductionResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runBookIntroductionWithSSE_opYieldFunc(_yield, _yieldErr, request, runtime)
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
func (client *Client) RunBookIntroductionWithOptions(request *RunBookIntroductionRequest, runtime *dara.RuntimeOptions) (_result *RunBookIntroductionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RunBookIntroductionResponse
func (client *Client) RunBookIntroduction(request *RunBookIntroductionRequest) (_result *RunBookIntroductionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunBookIntroductionResponse{}
	_body, _err := client.RunBookIntroductionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunBookSmartCardWithSSE(request *RunBookSmartCardRequest, runtime *dara.RuntimeOptions, _yield chan *RunBookSmartCardResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runBookSmartCardWithSSE_opYieldFunc(_yield, _yieldErr, request, runtime)
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
func (client *Client) RunBookSmartCardWithOptions(request *RunBookSmartCardRequest, runtime *dara.RuntimeOptions) (_result *RunBookSmartCardResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RunBookSmartCardResponse
func (client *Client) RunBookSmartCard(request *RunBookSmartCardRequest) (_result *RunBookSmartCardResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunBookSmartCardResponse{}
	_body, _err := client.RunBookSmartCardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunCommentGenerationWithSSE(tmpReq *RunCommentGenerationRequest, runtime *dara.RuntimeOptions, _yield chan *RunCommentGenerationResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runCommentGenerationWithSSE_opYieldFunc(_yield, _yieldErr, tmpReq, runtime)
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
func (client *Client) RunCommentGenerationWithOptions(tmpReq *RunCommentGenerationRequest, runtime *dara.RuntimeOptions) (_result *RunCommentGenerationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - RunCommentGenerationRequest
//
// @return RunCommentGenerationResponse
func (client *Client) RunCommentGeneration(request *RunCommentGenerationRequest) (_result *RunCommentGenerationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunCommentGenerationResponse{}
	_body, _err := client.RunCommentGenerationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunContinueContentWithSSE(request *RunContinueContentRequest, runtime *dara.RuntimeOptions, _yield chan *RunContinueContentResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runContinueContentWithSSE_opYieldFunc(_yield, _yieldErr, request, runtime)
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
func (client *Client) RunContinueContentWithOptions(request *RunContinueContentRequest, runtime *dara.RuntimeOptions) (_result *RunContinueContentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RunContinueContentResponse
func (client *Client) RunContinueContent(request *RunContinueContentRequest) (_result *RunContinueContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunContinueContentResponse{}
	_body, _err := client.RunContinueContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunCustomHotTopicAnalysisWithSSE(request *RunCustomHotTopicAnalysisRequest, runtime *dara.RuntimeOptions, _yield chan *RunCustomHotTopicAnalysisResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runCustomHotTopicAnalysisWithSSE_opYieldFunc(_yield, _yieldErr, request, runtime)
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
func (client *Client) RunCustomHotTopicAnalysisWithOptions(request *RunCustomHotTopicAnalysisRequest, runtime *dara.RuntimeOptions) (_result *RunCustomHotTopicAnalysisResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RunCustomHotTopicAnalysisResponse
func (client *Client) RunCustomHotTopicAnalysis(request *RunCustomHotTopicAnalysisRequest) (_result *RunCustomHotTopicAnalysisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunCustomHotTopicAnalysisResponse{}
	_body, _err := client.RunCustomHotTopicAnalysisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunCustomHotTopicViewPointAnalysisWithSSE(request *RunCustomHotTopicViewPointAnalysisRequest, runtime *dara.RuntimeOptions, _yield chan *RunCustomHotTopicViewPointAnalysisResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runCustomHotTopicViewPointAnalysisWithSSE_opYieldFunc(_yield, _yieldErr, request, runtime)
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
func (client *Client) RunCustomHotTopicViewPointAnalysisWithOptions(request *RunCustomHotTopicViewPointAnalysisRequest, runtime *dara.RuntimeOptions) (_result *RunCustomHotTopicViewPointAnalysisResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RunCustomHotTopicViewPointAnalysisResponse
func (client *Client) RunCustomHotTopicViewPointAnalysis(request *RunCustomHotTopicViewPointAnalysisRequest) (_result *RunCustomHotTopicViewPointAnalysisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunCustomHotTopicViewPointAnalysisResponse{}
	_body, _err := client.RunCustomHotTopicViewPointAnalysisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunDeepWritingWithSSE(request *RunDeepWritingRequest, runtime *dara.RuntimeOptions, _yield chan *RunDeepWritingResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runDeepWritingWithSSE_opYieldFunc(_yield, _yieldErr, request, runtime)
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
func (client *Client) RunDeepWritingWithOptions(request *RunDeepWritingRequest, runtime *dara.RuntimeOptions) (_result *RunDeepWritingResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RunDeepWritingResponse
func (client *Client) RunDeepWriting(request *RunDeepWritingRequest) (_result *RunDeepWritingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunDeepWritingResponse{}
	_body, _err := client.RunDeepWritingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunDocBrainmapWithSSE(request *RunDocBrainmapRequest, runtime *dara.RuntimeOptions, _yield chan *RunDocBrainmapResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runDocBrainmapWithSSE_opYieldFunc(_yield, _yieldErr, request, runtime)
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
func (client *Client) RunDocBrainmapWithOptions(request *RunDocBrainmapRequest, runtime *dara.RuntimeOptions) (_result *RunDocBrainmapResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RunDocBrainmapResponse
func (client *Client) RunDocBrainmap(request *RunDocBrainmapRequest) (_result *RunDocBrainmapResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunDocBrainmapResponse{}
	_body, _err := client.RunDocBrainmapWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunDocIntroductionWithSSE(request *RunDocIntroductionRequest, runtime *dara.RuntimeOptions, _yield chan *RunDocIntroductionResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runDocIntroductionWithSSE_opYieldFunc(_yield, _yieldErr, request, runtime)
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
func (client *Client) RunDocIntroductionWithOptions(request *RunDocIntroductionRequest, runtime *dara.RuntimeOptions) (_result *RunDocIntroductionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RunDocIntroductionResponse
func (client *Client) RunDocIntroduction(request *RunDocIntroductionRequest) (_result *RunDocIntroductionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunDocIntroductionResponse{}
	_body, _err := client.RunDocIntroductionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunDocQaWithSSE(tmpReq *RunDocQaRequest, runtime *dara.RuntimeOptions, _yield chan *RunDocQaResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runDocQaWithSSE_opYieldFunc(_yield, _yieldErr, tmpReq, runtime)
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
func (client *Client) RunDocQaWithOptions(tmpReq *RunDocQaRequest, runtime *dara.RuntimeOptions) (_result *RunDocQaResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - RunDocQaRequest
//
// @return RunDocQaResponse
func (client *Client) RunDocQa(request *RunDocQaRequest) (_result *RunDocQaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunDocQaResponse{}
	_body, _err := client.RunDocQaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunDocSmartCardWithSSE(request *RunDocSmartCardRequest, runtime *dara.RuntimeOptions, _yield chan *RunDocSmartCardResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runDocSmartCardWithSSE_opYieldFunc(_yield, _yieldErr, request, runtime)
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
func (client *Client) RunDocSmartCardWithOptions(request *RunDocSmartCardRequest, runtime *dara.RuntimeOptions) (_result *RunDocSmartCardResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RunDocSmartCardResponse
func (client *Client) RunDocSmartCard(request *RunDocSmartCardRequest) (_result *RunDocSmartCardResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunDocSmartCardResponse{}
	_body, _err := client.RunDocSmartCardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunDocSummaryWithSSE(request *RunDocSummaryRequest, runtime *dara.RuntimeOptions, _yield chan *RunDocSummaryResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runDocSummaryWithSSE_opYieldFunc(_yield, _yieldErr, request, runtime)
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
func (client *Client) RunDocSummaryWithOptions(request *RunDocSummaryRequest, runtime *dara.RuntimeOptions) (_result *RunDocSummaryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RunDocSummaryResponse
func (client *Client) RunDocSummary(request *RunDocSummaryRequest) (_result *RunDocSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunDocSummaryResponse{}
	_body, _err := client.RunDocSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunDocTranslationWithSSE(request *RunDocTranslationRequest, runtime *dara.RuntimeOptions, _yield chan *RunDocTranslationResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runDocTranslationWithSSE_opYieldFunc(_yield, _yieldErr, request, runtime)
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
func (client *Client) RunDocTranslationWithOptions(request *RunDocTranslationRequest, runtime *dara.RuntimeOptions) (_result *RunDocTranslationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RunDocTranslationResponse
func (client *Client) RunDocTranslation(request *RunDocTranslationRequest) (_result *RunDocTranslationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunDocTranslationResponse{}
	_body, _err := client.RunDocTranslationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunDocWashingWithSSE(request *RunDocWashingRequest, runtime *dara.RuntimeOptions, _yield chan *RunDocWashingResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runDocWashingWithSSE_opYieldFunc(_yield, _yieldErr, request, runtime)
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
func (client *Client) RunDocWashingWithOptions(request *RunDocWashingRequest, runtime *dara.RuntimeOptions) (_result *RunDocWashingResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RunDocWashingResponse
func (client *Client) RunDocWashing(request *RunDocWashingRequest) (_result *RunDocWashingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunDocWashingResponse{}
	_body, _err := client.RunDocWashingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunExpandContentWithSSE(request *RunExpandContentRequest, runtime *dara.RuntimeOptions, _yield chan *RunExpandContentResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runExpandContentWithSSE_opYieldFunc(_yield, _yieldErr, request, runtime)
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
func (client *Client) RunExpandContentWithOptions(request *RunExpandContentRequest, runtime *dara.RuntimeOptions) (_result *RunExpandContentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RunExpandContentResponse
func (client *Client) RunExpandContent(request *RunExpandContentRequest) (_result *RunExpandContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunExpandContentResponse{}
	_body, _err := client.RunExpandContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunGenerateQuestionsWithSSE(request *RunGenerateQuestionsRequest, runtime *dara.RuntimeOptions, _yield chan *RunGenerateQuestionsResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runGenerateQuestionsWithSSE_opYieldFunc(_yield, _yieldErr, request, runtime)
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
func (client *Client) RunGenerateQuestionsWithOptions(request *RunGenerateQuestionsRequest, runtime *dara.RuntimeOptions) (_result *RunGenerateQuestionsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RunGenerateQuestionsResponse
func (client *Client) RunGenerateQuestions(request *RunGenerateQuestionsRequest) (_result *RunGenerateQuestionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunGenerateQuestionsResponse{}
	_body, _err := client.RunGenerateQuestionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunHotwordWithSSE(request *RunHotwordRequest, runtime *dara.RuntimeOptions, _yield chan *RunHotwordResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runHotwordWithSSE_opYieldFunc(_yield, _yieldErr, request, runtime)
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
func (client *Client) RunHotwordWithOptions(request *RunHotwordRequest, runtime *dara.RuntimeOptions) (_result *RunHotwordResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RunHotwordResponse
func (client *Client) RunHotword(request *RunHotwordRequest) (_result *RunHotwordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunHotwordResponse{}
	_body, _err := client.RunHotwordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunKeywordsExtractionGenerationWithSSE(tmpReq *RunKeywordsExtractionGenerationRequest, runtime *dara.RuntimeOptions, _yield chan *RunKeywordsExtractionGenerationResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runKeywordsExtractionGenerationWithSSE_opYieldFunc(_yield, _yieldErr, tmpReq, runtime)
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
func (client *Client) RunKeywordsExtractionGenerationWithOptions(tmpReq *RunKeywordsExtractionGenerationRequest, runtime *dara.RuntimeOptions) (_result *RunKeywordsExtractionGenerationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - RunKeywordsExtractionGenerationRequest
//
// @return RunKeywordsExtractionGenerationResponse
func (client *Client) RunKeywordsExtractionGeneration(request *RunKeywordsExtractionGenerationRequest) (_result *RunKeywordsExtractionGenerationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunKeywordsExtractionGenerationResponse{}
	_body, _err := client.RunKeywordsExtractionGenerationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunMultiDocIntroductionWithSSE(tmpReq *RunMultiDocIntroductionRequest, runtime *dara.RuntimeOptions, _yield chan *RunMultiDocIntroductionResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runMultiDocIntroductionWithSSE_opYieldFunc(_yield, _yieldErr, tmpReq, runtime)
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
func (client *Client) RunMultiDocIntroductionWithOptions(tmpReq *RunMultiDocIntroductionRequest, runtime *dara.RuntimeOptions) (_result *RunMultiDocIntroductionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - RunMultiDocIntroductionRequest
//
// @return RunMultiDocIntroductionResponse
func (client *Client) RunMultiDocIntroduction(request *RunMultiDocIntroductionRequest) (_result *RunMultiDocIntroductionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunMultiDocIntroductionResponse{}
	_body, _err := client.RunMultiDocIntroductionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunPptOutlineGenerationWithSSE(request *RunPptOutlineGenerationRequest, runtime *dara.RuntimeOptions, _yield chan *RunPptOutlineGenerationResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runPptOutlineGenerationWithSSE_opYieldFunc(_yield, _yieldErr, request, runtime)
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
func (client *Client) RunPptOutlineGenerationWithOptions(request *RunPptOutlineGenerationRequest, runtime *dara.RuntimeOptions) (_result *RunPptOutlineGenerationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RunPptOutlineGenerationResponse
func (client *Client) RunPptOutlineGeneration(request *RunPptOutlineGenerationRequest) (_result *RunPptOutlineGenerationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunPptOutlineGenerationResponse{}
	_body, _err := client.RunPptOutlineGenerationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunQuickWritingWithSSE(tmpReq *RunQuickWritingRequest, runtime *dara.RuntimeOptions, _yield chan *RunQuickWritingResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runQuickWritingWithSSE_opYieldFunc(_yield, _yieldErr, tmpReq, runtime)
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
func (client *Client) RunQuickWritingWithOptions(tmpReq *RunQuickWritingRequest, runtime *dara.RuntimeOptions) (_result *RunQuickWritingResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - RunQuickWritingRequest
//
// @return RunQuickWritingResponse
func (client *Client) RunQuickWriting(request *RunQuickWritingRequest) (_result *RunQuickWritingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunQuickWritingResponse{}
	_body, _err := client.RunQuickWritingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunSearchGenerationWithSSE(tmpReq *RunSearchGenerationRequest, runtime *dara.RuntimeOptions, _yield chan *RunSearchGenerationResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runSearchGenerationWithSSE_opYieldFunc(_yield, _yieldErr, tmpReq, runtime)
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
func (client *Client) RunSearchGenerationWithOptions(tmpReq *RunSearchGenerationRequest, runtime *dara.RuntimeOptions) (_result *RunSearchGenerationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - RunSearchGenerationRequest
//
// @return RunSearchGenerationResponse
func (client *Client) RunSearchGeneration(request *RunSearchGenerationRequest) (_result *RunSearchGenerationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunSearchGenerationResponse{}
	_body, _err := client.RunSearchGenerationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunSearchSimilarArticlesWithSSE(tmpReq *RunSearchSimilarArticlesRequest, runtime *dara.RuntimeOptions, _yield chan *RunSearchSimilarArticlesResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runSearchSimilarArticlesWithSSE_opYieldFunc(_yield, _yieldErr, tmpReq, runtime)
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
func (client *Client) RunSearchSimilarArticlesWithOptions(tmpReq *RunSearchSimilarArticlesRequest, runtime *dara.RuntimeOptions) (_result *RunSearchSimilarArticlesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - RunSearchSimilarArticlesRequest
//
// @return RunSearchSimilarArticlesResponse
func (client *Client) RunSearchSimilarArticles(request *RunSearchSimilarArticlesRequest) (_result *RunSearchSimilarArticlesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunSearchSimilarArticlesResponse{}
	_body, _err := client.RunSearchSimilarArticlesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunStepByStepWritingWithSSE(tmpReq *RunStepByStepWritingRequest, runtime *dara.RuntimeOptions, _yield chan *RunStepByStepWritingResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runStepByStepWritingWithSSE_opYieldFunc(_yield, _yieldErr, tmpReq, runtime)
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
func (client *Client) RunStepByStepWritingWithOptions(tmpReq *RunStepByStepWritingRequest, runtime *dara.RuntimeOptions) (_result *RunStepByStepWritingResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - RunStepByStepWritingRequest
//
// @return RunStepByStepWritingResponse
func (client *Client) RunStepByStepWriting(request *RunStepByStepWritingRequest) (_result *RunStepByStepWritingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunStepByStepWritingResponse{}
	_body, _err := client.RunStepByStepWritingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunStyleFeatureAnalysisWithSSE(tmpReq *RunStyleFeatureAnalysisRequest, runtime *dara.RuntimeOptions, _yield chan *RunStyleFeatureAnalysisResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runStyleFeatureAnalysisWithSSE_opYieldFunc(_yield, _yieldErr, tmpReq, runtime)
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
func (client *Client) RunStyleFeatureAnalysisWithOptions(tmpReq *RunStyleFeatureAnalysisRequest, runtime *dara.RuntimeOptions) (_result *RunStyleFeatureAnalysisResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - RunStyleFeatureAnalysisRequest
//
// @return RunStyleFeatureAnalysisResponse
func (client *Client) RunStyleFeatureAnalysis(request *RunStyleFeatureAnalysisRequest) (_result *RunStyleFeatureAnalysisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunStyleFeatureAnalysisResponse{}
	_body, _err := client.RunStyleFeatureAnalysisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunSummaryGenerateWithSSE(request *RunSummaryGenerateRequest, runtime *dara.RuntimeOptions, _yield chan *RunSummaryGenerateResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runSummaryGenerateWithSSE_opYieldFunc(_yield, _yieldErr, request, runtime)
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
func (client *Client) RunSummaryGenerateWithOptions(request *RunSummaryGenerateRequest, runtime *dara.RuntimeOptions) (_result *RunSummaryGenerateResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RunSummaryGenerateResponse
func (client *Client) RunSummaryGenerate(request *RunSummaryGenerateRequest) (_result *RunSummaryGenerateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunSummaryGenerateResponse{}
	_body, _err := client.RunSummaryGenerateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunTextPolishingWithSSE(request *RunTextPolishingRequest, runtime *dara.RuntimeOptions, _yield chan *RunTextPolishingResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runTextPolishingWithSSE_opYieldFunc(_yield, _yieldErr, request, runtime)
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
func (client *Client) RunTextPolishingWithOptions(request *RunTextPolishingRequest, runtime *dara.RuntimeOptions) (_result *RunTextPolishingResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RunTextPolishingResponse
func (client *Client) RunTextPolishing(request *RunTextPolishingRequest) (_result *RunTextPolishingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunTextPolishingResponse{}
	_body, _err := client.RunTextPolishingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunTitleGenerationWithSSE(tmpReq *RunTitleGenerationRequest, runtime *dara.RuntimeOptions, _yield chan *RunTitleGenerationResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runTitleGenerationWithSSE_opYieldFunc(_yield, _yieldErr, tmpReq, runtime)
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
func (client *Client) RunTitleGenerationWithOptions(tmpReq *RunTitleGenerationRequest, runtime *dara.RuntimeOptions) (_result *RunTitleGenerationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - RunTitleGenerationRequest
//
// @return RunTitleGenerationResponse
func (client *Client) RunTitleGeneration(request *RunTitleGenerationRequest) (_result *RunTitleGenerationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunTitleGenerationResponse{}
	_body, _err := client.RunTitleGenerationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunTopicSelectionMergeWithSSE(tmpReq *RunTopicSelectionMergeRequest, runtime *dara.RuntimeOptions, _yield chan *RunTopicSelectionMergeResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runTopicSelectionMergeWithSSE_opYieldFunc(_yield, _yieldErr, tmpReq, runtime)
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
func (client *Client) RunTopicSelectionMergeWithOptions(tmpReq *RunTopicSelectionMergeRequest, runtime *dara.RuntimeOptions) (_result *RunTopicSelectionMergeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - RunTopicSelectionMergeRequest
//
// @return RunTopicSelectionMergeResponse
func (client *Client) RunTopicSelectionMerge(request *RunTopicSelectionMergeRequest) (_result *RunTopicSelectionMergeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunTopicSelectionMergeResponse{}
	_body, _err := client.RunTopicSelectionMergeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunTranslateGenerationWithSSE(tmpReq *RunTranslateGenerationRequest, runtime *dara.RuntimeOptions, _yield chan *RunTranslateGenerationResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runTranslateGenerationWithSSE_opYieldFunc(_yield, _yieldErr, tmpReq, runtime)
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
func (client *Client) RunTranslateGenerationWithOptions(tmpReq *RunTranslateGenerationRequest, runtime *dara.RuntimeOptions) (_result *RunTranslateGenerationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - RunTranslateGenerationRequest
//
// @return RunTranslateGenerationResponse
func (client *Client) RunTranslateGeneration(request *RunTranslateGenerationRequest) (_result *RunTranslateGenerationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunTranslateGenerationResponse{}
	_body, _err := client.RunTranslateGenerationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunVideoScriptGenerateWithSSE(request *RunVideoScriptGenerateRequest, runtime *dara.RuntimeOptions, _yield chan *RunVideoScriptGenerateResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runVideoScriptGenerateWithSSE_opYieldFunc(_yield, _yieldErr, request, runtime)
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
func (client *Client) RunVideoScriptGenerateWithOptions(request *RunVideoScriptGenerateRequest, runtime *dara.RuntimeOptions) (_result *RunVideoScriptGenerateResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RunVideoScriptGenerateResponse
func (client *Client) RunVideoScriptGenerate(request *RunVideoScriptGenerateRequest) (_result *RunVideoScriptGenerateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunVideoScriptGenerateResponse{}
	_body, _err := client.RunVideoScriptGenerateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunWriteToneGenerationWithSSE(tmpReq *RunWriteToneGenerationRequest, runtime *dara.RuntimeOptions, _yield chan *RunWriteToneGenerationResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runWriteToneGenerationWithSSE_opYieldFunc(_yield, _yieldErr, tmpReq, runtime)
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
func (client *Client) RunWriteToneGenerationWithOptions(tmpReq *RunWriteToneGenerationRequest, runtime *dara.RuntimeOptions) (_result *RunWriteToneGenerationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - RunWriteToneGenerationRequest
//
// @return RunWriteToneGenerationResponse
func (client *Client) RunWriteToneGeneration(request *RunWriteToneGenerationRequest) (_result *RunWriteToneGenerationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunWriteToneGenerationResponse{}
	_body, _err := client.RunWriteToneGenerationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunWritingWithSSE(tmpReq *RunWritingRequest, runtime *dara.RuntimeOptions, _yield chan *RunWritingResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runWritingWithSSE_opYieldFunc(_yield, _yieldErr, tmpReq, runtime)
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
func (client *Client) RunWritingWithOptions(tmpReq *RunWritingRequest, runtime *dara.RuntimeOptions) (_result *RunWritingResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - RunWritingRequest
//
// @return RunWritingResponse
func (client *Client) RunWriting(request *RunWritingRequest) (_result *RunWritingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunWritingResponse{}
	_body, _err := client.RunWritingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunWritingV2WithSSE(tmpReq *RunWritingV2Request, runtime *dara.RuntimeOptions, _yield chan *RunWritingV2Response, _yieldErr chan error) {
	defer close(_yield)
	client.runWritingV2WithSSE_opYieldFunc(_yield, _yieldErr, tmpReq, runtime)
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
func (client *Client) RunWritingV2WithOptions(tmpReq *RunWritingV2Request, runtime *dara.RuntimeOptions) (_result *RunWritingV2Response, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - RunWritingV2Request
//
// @return RunWritingV2Response
func (client *Client) RunWritingV2(request *RunWritingV2Request) (_result *RunWritingV2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunWritingV2Response{}
	_body, _err := client.RunWritingV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SaveCustomTextWithOptions(request *SaveCustomTextRequest, runtime *dara.RuntimeOptions) (_result *SaveCustomTextResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SaveCustomTextResponse
func (client *Client) SaveCustomText(request *SaveCustomTextRequest) (_result *SaveCustomTextResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveCustomTextResponse{}
	_body, _err := client.SaveCustomTextWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SaveDataSourceOrderConfigWithOptions(tmpReq *SaveDataSourceOrderConfigRequest, runtime *dara.RuntimeOptions) (_result *SaveDataSourceOrderConfigResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - SaveDataSourceOrderConfigRequest
//
// @return SaveDataSourceOrderConfigResponse
func (client *Client) SaveDataSourceOrderConfig(request *SaveDataSourceOrderConfigRequest) (_result *SaveDataSourceOrderConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveDataSourceOrderConfigResponse{}
	_body, _err := client.SaveDataSourceOrderConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SaveMaterialDocumentWithOptions(tmpReq *SaveMaterialDocumentRequest, runtime *dara.RuntimeOptions) (_result *SaveMaterialDocumentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - SaveMaterialDocumentRequest
//
// @return SaveMaterialDocumentResponse
func (client *Client) SaveMaterialDocument(request *SaveMaterialDocumentRequest) (_result *SaveMaterialDocumentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveMaterialDocumentResponse{}
	_body, _err := client.SaveMaterialDocumentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SaveOrUpdateOssConfigWithOptions(request *SaveOrUpdateOssConfigRequest, runtime *dara.RuntimeOptions) (_result *SaveOrUpdateOssConfigResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SaveOrUpdateOssConfigResponse
func (client *Client) SaveOrUpdateOssConfig(request *SaveOrUpdateOssConfigRequest) (_result *SaveOrUpdateOssConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveOrUpdateOssConfigResponse{}
	_body, _err := client.SaveOrUpdateOssConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SaveStyleLearningResultWithOptions(tmpReq *SaveStyleLearningResultRequest, runtime *dara.RuntimeOptions) (_result *SaveStyleLearningResultResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - SaveStyleLearningResultRequest
//
// @return SaveStyleLearningResultResponse
func (client *Client) SaveStyleLearningResult(request *SaveStyleLearningResultRequest) (_result *SaveStyleLearningResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveStyleLearningResultResponse{}
	_body, _err := client.SaveStyleLearningResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SearchDatasetDocumentsWithOptions(tmpReq *SearchDatasetDocumentsRequest, runtime *dara.RuntimeOptions) (_result *SearchDatasetDocumentsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - SearchDatasetDocumentsRequest
//
// @return SearchDatasetDocumentsResponse
func (client *Client) SearchDatasetDocuments(request *SearchDatasetDocumentsRequest) (_result *SearchDatasetDocumentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchDatasetDocumentsResponse{}
	_body, _err := client.SearchDatasetDocumentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SearchNewsWithOptions(tmpReq *SearchNewsRequest, runtime *dara.RuntimeOptions) (_result *SearchNewsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - SearchNewsRequest
//
// @return SearchNewsResponse
func (client *Client) SearchNews(request *SearchNewsRequest) (_result *SearchNewsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchNewsResponse{}
	_body, _err := client.SearchNewsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SubmitAsyncTaskWithOptions(request *SubmitAsyncTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitAsyncTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SubmitAsyncTaskResponse
func (client *Client) SubmitAsyncTask(request *SubmitAsyncTaskRequest) (_result *SubmitAsyncTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitAsyncTaskResponse{}
	_body, _err := client.SubmitAsyncTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SubmitAuditNoteWithOptions(request *SubmitAuditNoteRequest, runtime *dara.RuntimeOptions) (_result *SubmitAuditNoteResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SubmitAuditNoteResponse
func (client *Client) SubmitAuditNote(request *SubmitAuditNoteRequest) (_result *SubmitAuditNoteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitAuditNoteResponse{}
	_body, _err := client.SubmitAuditNoteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SubmitAuditTaskWithOptions(request *SubmitAuditTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitAuditTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SubmitAuditTaskResponse
func (client *Client) SubmitAuditTask(request *SubmitAuditTaskRequest) (_result *SubmitAuditTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitAuditTaskResponse{}
	_body, _err := client.SubmitAuditTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SubmitCustomHotTopicBroadcastJobWithOptions(tmpReq *SubmitCustomHotTopicBroadcastJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitCustomHotTopicBroadcastJobResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - SubmitCustomHotTopicBroadcastJobRequest
//
// @return SubmitCustomHotTopicBroadcastJobResponse
func (client *Client) SubmitCustomHotTopicBroadcastJob(request *SubmitCustomHotTopicBroadcastJobRequest) (_result *SubmitCustomHotTopicBroadcastJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitCustomHotTopicBroadcastJobResponse{}
	_body, _err := client.SubmitCustomHotTopicBroadcastJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SubmitCustomSourceTopicAnalysisWithOptions(tmpReq *SubmitCustomSourceTopicAnalysisRequest, runtime *dara.RuntimeOptions) (_result *SubmitCustomSourceTopicAnalysisResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - SubmitCustomSourceTopicAnalysisRequest
//
// @return SubmitCustomSourceTopicAnalysisResponse
func (client *Client) SubmitCustomSourceTopicAnalysis(request *SubmitCustomSourceTopicAnalysisRequest) (_result *SubmitCustomSourceTopicAnalysisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitCustomSourceTopicAnalysisResponse{}
	_body, _err := client.SubmitCustomSourceTopicAnalysisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SubmitCustomTopicSelectionPerspectiveAnalysisTaskWithOptions(tmpReq *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest
//
// @return SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse
func (client *Client) SubmitCustomTopicSelectionPerspectiveAnalysisTask(request *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest) (_result *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse{}
	_body, _err := client.SubmitCustomTopicSelectionPerspectiveAnalysisTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SubmitDeepWriteTaskWithOptions(tmpReq *SubmitDeepWriteTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitDeepWriteTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - SubmitDeepWriteTaskRequest
//
// @return SubmitDeepWriteTaskResponse
func (client *Client) SubmitDeepWriteTask(request *SubmitDeepWriteTaskRequest) (_result *SubmitDeepWriteTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitDeepWriteTaskResponse{}
	_body, _err := client.SubmitDeepWriteTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SubmitDocClusterTaskWithOptions(tmpReq *SubmitDocClusterTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitDocClusterTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - SubmitDocClusterTaskRequest
//
// @return SubmitDocClusterTaskResponse
func (client *Client) SubmitDocClusterTask(request *SubmitDocClusterTaskRequest) (_result *SubmitDocClusterTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitDocClusterTaskResponse{}
	_body, _err := client.SubmitDocClusterTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SubmitEnterpriseVocAnalysisTaskWithOptions(tmpReq *SubmitEnterpriseVocAnalysisTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitEnterpriseVocAnalysisTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - SubmitEnterpriseVocAnalysisTaskRequest
//
// @return SubmitEnterpriseVocAnalysisTaskResponse
func (client *Client) SubmitEnterpriseVocAnalysisTask(request *SubmitEnterpriseVocAnalysisTaskRequest) (_result *SubmitEnterpriseVocAnalysisTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitEnterpriseVocAnalysisTaskResponse{}
	_body, _err := client.SubmitEnterpriseVocAnalysisTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SubmitExportTermsTaskWithOptions(request *SubmitExportTermsTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitExportTermsTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SubmitExportTermsTaskResponse
func (client *Client) SubmitExportTermsTask(request *SubmitExportTermsTaskRequest) (_result *SubmitExportTermsTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitExportTermsTaskResponse{}
	_body, _err := client.SubmitExportTermsTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SubmitFactAuditUrlWithOptions(request *SubmitFactAuditUrlRequest, runtime *dara.RuntimeOptions) (_result *SubmitFactAuditUrlResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SubmitFactAuditUrlResponse
func (client *Client) SubmitFactAuditUrl(request *SubmitFactAuditUrlRequest) (_result *SubmitFactAuditUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitFactAuditUrlResponse{}
	_body, _err := client.SubmitFactAuditUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SubmitImportTermsTaskWithOptions(request *SubmitImportTermsTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitImportTermsTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SubmitImportTermsTaskResponse
func (client *Client) SubmitImportTermsTask(request *SubmitImportTermsTaskRequest) (_result *SubmitImportTermsTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitImportTermsTaskResponse{}
	_body, _err := client.SubmitImportTermsTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SubmitParseDocumentLayoutTaskWithOptions(request *SubmitParseDocumentLayoutTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitParseDocumentLayoutTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SubmitParseDocumentLayoutTaskResponse
func (client *Client) SubmitParseDocumentLayoutTask(request *SubmitParseDocumentLayoutTaskRequest) (_result *SubmitParseDocumentLayoutTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitParseDocumentLayoutTaskResponse{}
	_body, _err := client.SubmitParseDocumentLayoutTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SubmitSmartAuditWithOptions(tmpReq *SubmitSmartAuditRequest, runtime *dara.RuntimeOptions) (_result *SubmitSmartAuditResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - SubmitSmartAuditRequest
//
// @return SubmitSmartAuditResponse
func (client *Client) SubmitSmartAudit(request *SubmitSmartAuditRequest) (_result *SubmitSmartAuditResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitSmartAuditResponse{}
	_body, _err := client.SubmitSmartAuditWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SubmitSmartClipTaskWithOptions(tmpReq *SubmitSmartClipTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitSmartClipTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - SubmitSmartClipTaskRequest
//
// @return SubmitSmartClipTaskResponse
func (client *Client) SubmitSmartClipTask(request *SubmitSmartClipTaskRequest) (_result *SubmitSmartClipTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitSmartClipTaskResponse{}
	_body, _err := client.SubmitSmartClipTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SubmitTopicSelectionPerspectiveAnalysisTaskWithOptions(tmpReq *SubmitTopicSelectionPerspectiveAnalysisTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitTopicSelectionPerspectiveAnalysisTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - SubmitTopicSelectionPerspectiveAnalysisTaskRequest
//
// @return SubmitTopicSelectionPerspectiveAnalysisTaskResponse
func (client *Client) SubmitTopicSelectionPerspectiveAnalysisTask(request *SubmitTopicSelectionPerspectiveAnalysisTaskRequest) (_result *SubmitTopicSelectionPerspectiveAnalysisTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitTopicSelectionPerspectiveAnalysisTaskResponse{}
	_body, _err := client.SubmitTopicSelectionPerspectiveAnalysisTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SubmitVideoAuditWithOptions(request *SubmitVideoAuditRequest, runtime *dara.RuntimeOptions) (_result *SubmitVideoAuditResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SubmitVideoAuditResponse
func (client *Client) SubmitVideoAudit(request *SubmitVideoAuditRequest) (_result *SubmitVideoAuditResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitVideoAuditResponse{}
	_body, _err := client.SubmitVideoAuditWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateCustomTextWithOptions(request *UpdateCustomTextRequest, runtime *dara.RuntimeOptions) (_result *UpdateCustomTextResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateCustomTextResponse
func (client *Client) UpdateCustomText(request *UpdateCustomTextRequest) (_result *UpdateCustomTextResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCustomTextResponse{}
	_body, _err := client.UpdateCustomTextWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateDatasetWithOptions(tmpReq *UpdateDatasetRequest, runtime *dara.RuntimeOptions) (_result *UpdateDatasetResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UpdateDatasetRequest
//
// @return UpdateDatasetResponse
func (client *Client) UpdateDataset(request *UpdateDatasetRequest) (_result *UpdateDatasetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDatasetResponse{}
	_body, _err := client.UpdateDatasetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateDatasetDocumentWithOptions(tmpReq *UpdateDatasetDocumentRequest, runtime *dara.RuntimeOptions) (_result *UpdateDatasetDocumentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UpdateDatasetDocumentRequest
//
// @return UpdateDatasetDocumentResponse
func (client *Client) UpdateDatasetDocument(request *UpdateDatasetDocumentRequest) (_result *UpdateDatasetDocumentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDatasetDocumentResponse{}
	_body, _err := client.UpdateDatasetDocumentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateGeneralConfigWithOptions(request *UpdateGeneralConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateGeneralConfigResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateGeneralConfigResponse
func (client *Client) UpdateGeneralConfig(request *UpdateGeneralConfigRequest) (_result *UpdateGeneralConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGeneralConfigResponse{}
	_body, _err := client.UpdateGeneralConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateGeneratedContentWithOptions(tmpReq *UpdateGeneratedContentRequest, runtime *dara.RuntimeOptions) (_result *UpdateGeneratedContentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UpdateGeneratedContentRequest
//
// @return UpdateGeneratedContentResponse
func (client *Client) UpdateGeneratedContent(request *UpdateGeneratedContentRequest) (_result *UpdateGeneratedContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGeneratedContentResponse{}
	_body, _err := client.UpdateGeneratedContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateMaterialDocumentWithOptions(tmpReq *UpdateMaterialDocumentRequest, runtime *dara.RuntimeOptions) (_result *UpdateMaterialDocumentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UpdateMaterialDocumentRequest
//
// @return UpdateMaterialDocumentResponse
func (client *Client) UpdateMaterialDocument(request *UpdateMaterialDocumentRequest) (_result *UpdateMaterialDocumentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateMaterialDocumentResponse{}
	_body, _err := client.UpdateMaterialDocumentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UploadBookWithOptions(tmpReq *UploadBookRequest, runtime *dara.RuntimeOptions) (_result *UploadBookResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UploadBookRequest
//
// @return UploadBookResponse
func (client *Client) UploadBook(request *UploadBookRequest) (_result *UploadBookResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadBookResponse{}
	_body, _err := client.UploadBookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UploadDocWithOptions(tmpReq *UploadDocRequest, runtime *dara.RuntimeOptions) (_result *UploadDocResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UploadDocRequest
//
// @return UploadDocResponse
func (client *Client) UploadDoc(request *UploadDocRequest) (_result *UploadDocResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadDocResponse{}
	_body, _err := client.UploadDocWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ValidateUploadTemplateWithOptions(request *ValidateUploadTemplateRequest, runtime *dara.RuntimeOptions) (_result *ValidateUploadTemplateResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ValidateUploadTemplateResponse
func (client *Client) ValidateUploadTemplate(request *ValidateUploadTemplateRequest) (_result *ValidateUploadTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ValidateUploadTemplateResponse{}
	_body, _err := client.ValidateUploadTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) runAbbreviationContentWithSSE_opYieldFunc(_yield chan *RunAbbreviationContentResponse, _yieldErr chan error, request *RunAbbreviationContentRequest, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runAiHelperWritingWithSSE_opYieldFunc(_yield chan *RunAiHelperWritingResponse, _yieldErr chan error, tmpReq *RunAiHelperWritingRequest, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runBookBrainmapWithSSE_opYieldFunc(_yield chan *RunBookBrainmapResponse, _yieldErr chan error, request *RunBookBrainmapRequest, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runBookIntroductionWithSSE_opYieldFunc(_yield chan *RunBookIntroductionResponse, _yieldErr chan error, request *RunBookIntroductionRequest, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runBookSmartCardWithSSE_opYieldFunc(_yield chan *RunBookSmartCardResponse, _yieldErr chan error, request *RunBookSmartCardRequest, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runCommentGenerationWithSSE_opYieldFunc(_yield chan *RunCommentGenerationResponse, _yieldErr chan error, tmpReq *RunCommentGenerationRequest, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runContinueContentWithSSE_opYieldFunc(_yield chan *RunContinueContentResponse, _yieldErr chan error, request *RunContinueContentRequest, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runCustomHotTopicAnalysisWithSSE_opYieldFunc(_yield chan *RunCustomHotTopicAnalysisResponse, _yieldErr chan error, request *RunCustomHotTopicAnalysisRequest, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runCustomHotTopicViewPointAnalysisWithSSE_opYieldFunc(_yield chan *RunCustomHotTopicViewPointAnalysisResponse, _yieldErr chan error, request *RunCustomHotTopicViewPointAnalysisRequest, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runDeepWritingWithSSE_opYieldFunc(_yield chan *RunDeepWritingResponse, _yieldErr chan error, request *RunDeepWritingRequest, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runDocBrainmapWithSSE_opYieldFunc(_yield chan *RunDocBrainmapResponse, _yieldErr chan error, request *RunDocBrainmapRequest, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runDocIntroductionWithSSE_opYieldFunc(_yield chan *RunDocIntroductionResponse, _yieldErr chan error, request *RunDocIntroductionRequest, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runDocQaWithSSE_opYieldFunc(_yield chan *RunDocQaResponse, _yieldErr chan error, tmpReq *RunDocQaRequest, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runDocSmartCardWithSSE_opYieldFunc(_yield chan *RunDocSmartCardResponse, _yieldErr chan error, request *RunDocSmartCardRequest, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runDocSummaryWithSSE_opYieldFunc(_yield chan *RunDocSummaryResponse, _yieldErr chan error, request *RunDocSummaryRequest, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runDocTranslationWithSSE_opYieldFunc(_yield chan *RunDocTranslationResponse, _yieldErr chan error, request *RunDocTranslationRequest, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runDocWashingWithSSE_opYieldFunc(_yield chan *RunDocWashingResponse, _yieldErr chan error, request *RunDocWashingRequest, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runExpandContentWithSSE_opYieldFunc(_yield chan *RunExpandContentResponse, _yieldErr chan error, request *RunExpandContentRequest, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runGenerateQuestionsWithSSE_opYieldFunc(_yield chan *RunGenerateQuestionsResponse, _yieldErr chan error, request *RunGenerateQuestionsRequest, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runHotwordWithSSE_opYieldFunc(_yield chan *RunHotwordResponse, _yieldErr chan error, request *RunHotwordRequest, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runKeywordsExtractionGenerationWithSSE_opYieldFunc(_yield chan *RunKeywordsExtractionGenerationResponse, _yieldErr chan error, tmpReq *RunKeywordsExtractionGenerationRequest, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runMultiDocIntroductionWithSSE_opYieldFunc(_yield chan *RunMultiDocIntroductionResponse, _yieldErr chan error, tmpReq *RunMultiDocIntroductionRequest, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runPptOutlineGenerationWithSSE_opYieldFunc(_yield chan *RunPptOutlineGenerationResponse, _yieldErr chan error, request *RunPptOutlineGenerationRequest, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runQuickWritingWithSSE_opYieldFunc(_yield chan *RunQuickWritingResponse, _yieldErr chan error, tmpReq *RunQuickWritingRequest, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runSearchGenerationWithSSE_opYieldFunc(_yield chan *RunSearchGenerationResponse, _yieldErr chan error, tmpReq *RunSearchGenerationRequest, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runSearchSimilarArticlesWithSSE_opYieldFunc(_yield chan *RunSearchSimilarArticlesResponse, _yieldErr chan error, tmpReq *RunSearchSimilarArticlesRequest, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runStepByStepWritingWithSSE_opYieldFunc(_yield chan *RunStepByStepWritingResponse, _yieldErr chan error, tmpReq *RunStepByStepWritingRequest, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runStyleFeatureAnalysisWithSSE_opYieldFunc(_yield chan *RunStyleFeatureAnalysisResponse, _yieldErr chan error, tmpReq *RunStyleFeatureAnalysisRequest, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runSummaryGenerateWithSSE_opYieldFunc(_yield chan *RunSummaryGenerateResponse, _yieldErr chan error, request *RunSummaryGenerateRequest, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runTextPolishingWithSSE_opYieldFunc(_yield chan *RunTextPolishingResponse, _yieldErr chan error, request *RunTextPolishingRequest, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runTitleGenerationWithSSE_opYieldFunc(_yield chan *RunTitleGenerationResponse, _yieldErr chan error, tmpReq *RunTitleGenerationRequest, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runTopicSelectionMergeWithSSE_opYieldFunc(_yield chan *RunTopicSelectionMergeResponse, _yieldErr chan error, tmpReq *RunTopicSelectionMergeRequest, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runTranslateGenerationWithSSE_opYieldFunc(_yield chan *RunTranslateGenerationResponse, _yieldErr chan error, tmpReq *RunTranslateGenerationRequest, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runVideoScriptGenerateWithSSE_opYieldFunc(_yield chan *RunVideoScriptGenerateResponse, _yieldErr chan error, request *RunVideoScriptGenerateRequest, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runWriteToneGenerationWithSSE_opYieldFunc(_yield chan *RunWriteToneGenerationResponse, _yieldErr chan error, tmpReq *RunWriteToneGenerationRequest, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runWritingWithSSE_opYieldFunc(_yield chan *RunWritingResponse, _yieldErr chan error, tmpReq *RunWritingRequest, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runWritingV2WithSSE_opYieldFunc(_yield chan *RunWritingV2Response, _yieldErr chan error, tmpReq *RunWritingV2Request, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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
