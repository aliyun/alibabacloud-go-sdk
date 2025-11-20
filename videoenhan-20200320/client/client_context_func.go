// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 视频人脸融合模板增加
//
// @param request - AddFaceVideoTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddFaceVideoTemplateResponse
func (client *Client) AddFaceVideoTemplateWithContext(ctx context.Context, request *AddFaceVideoTemplateRequest, runtime *dara.RuntimeOptions) (_result *AddFaceVideoTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.VideoScene) {
		body["VideoScene"] = request.VideoScene
	}

	if !dara.IsNil(request.VideoURL) {
		body["VideoURL"] = request.VideoURL
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddFaceVideoTemplate"),
		Version:     dara.String("2020-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddFaceVideoTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AdjustVideoColorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AdjustVideoColorResponse
func (client *Client) AdjustVideoColorWithContext(ctx context.Context, request *AdjustVideoColorRequest, runtime *dara.RuntimeOptions) (_result *AdjustVideoColorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Mode) {
		body["Mode"] = request.Mode
	}

	if !dara.IsNil(request.VideoBitrate) {
		body["VideoBitrate"] = request.VideoBitrate
	}

	if !dara.IsNil(request.VideoCodec) {
		body["VideoCodec"] = request.VideoCodec
	}

	if !dara.IsNil(request.VideoFormat) {
		body["VideoFormat"] = request.VideoFormat
	}

	if !dara.IsNil(request.VideoUrl) {
		body["VideoUrl"] = request.VideoUrl
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AdjustVideoColor"),
		Version:     dara.String("2020-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AdjustVideoColorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ChangeVideoSizeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeVideoSizeResponse
func (client *Client) ChangeVideoSizeWithContext(ctx context.Context, request *ChangeVideoSizeRequest, runtime *dara.RuntimeOptions) (_result *ChangeVideoSizeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.B) {
		body["B"] = request.B
	}

	if !dara.IsNil(request.CropType) {
		body["CropType"] = request.CropType
	}

	if !dara.IsNil(request.FillType) {
		body["FillType"] = request.FillType
	}

	if !dara.IsNil(request.G) {
		body["G"] = request.G
	}

	if !dara.IsNil(request.Height) {
		body["Height"] = request.Height
	}

	if !dara.IsNil(request.R) {
		body["R"] = request.R
	}

	if !dara.IsNil(request.Tightness) {
		body["Tightness"] = request.Tightness
	}

	if !dara.IsNil(request.VideoUrl) {
		body["VideoUrl"] = request.VideoUrl
	}

	if !dara.IsNil(request.Width) {
		body["Width"] = request.Width
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeVideoSize"),
		Version:     dara.String("2020-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeVideoSizeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视频人脸融合模板删除
//
// @param request - DeleteFaceVideoTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFaceVideoTemplateResponse
func (client *Client) DeleteFaceVideoTemplateWithContext(ctx context.Context, request *DeleteFaceVideoTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteFaceVideoTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFaceVideoTemplate"),
		Version:     dara.String("2020-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFaceVideoTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视频人像增强
//
// @param request - EnhancePortraitVideoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnhancePortraitVideoResponse
func (client *Client) EnhancePortraitVideoWithContext(ctx context.Context, request *EnhancePortraitVideoRequest, runtime *dara.RuntimeOptions) (_result *EnhancePortraitVideoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.VideoUrl) {
		body["VideoUrl"] = request.VideoUrl
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnhancePortraitVideo"),
		Version:     dara.String("2020-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnhancePortraitVideoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - EnhanceVideoQualityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnhanceVideoQualityResponse
func (client *Client) EnhanceVideoQualityWithContext(ctx context.Context, request *EnhanceVideoQualityRequest, runtime *dara.RuntimeOptions) (_result *EnhanceVideoQualityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Bitrate) {
		body["Bitrate"] = request.Bitrate
	}

	if !dara.IsNil(request.FrameRate) {
		body["FrameRate"] = request.FrameRate
	}

	if !dara.IsNil(request.HDRFormat) {
		body["HDRFormat"] = request.HDRFormat
	}

	if !dara.IsNil(request.MaxIlluminance) {
		body["MaxIlluminance"] = request.MaxIlluminance
	}

	if !dara.IsNil(request.OutPutHeight) {
		body["OutPutHeight"] = request.OutPutHeight
	}

	if !dara.IsNil(request.OutPutWidth) {
		body["OutPutWidth"] = request.OutPutWidth
	}

	if !dara.IsNil(request.VideoURL) {
		body["VideoURL"] = request.VideoURL
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnhanceVideoQuality"),
		Version:     dara.String("2020-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnhanceVideoQualityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - EraseVideoLogoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EraseVideoLogoResponse
func (client *Client) EraseVideoLogoWithContext(ctx context.Context, request *EraseVideoLogoRequest, runtime *dara.RuntimeOptions) (_result *EraseVideoLogoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Boxes) {
		body["Boxes"] = request.Boxes
	}

	if !dara.IsNil(request.VideoUrl) {
		body["VideoUrl"] = request.VideoUrl
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EraseVideoLogo"),
		Version:     dara.String("2020-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EraseVideoLogoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - EraseVideoSubtitlesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EraseVideoSubtitlesResponse
func (client *Client) EraseVideoSubtitlesWithContext(ctx context.Context, request *EraseVideoSubtitlesRequest, runtime *dara.RuntimeOptions) (_result *EraseVideoSubtitlesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BH) {
		body["BH"] = request.BH
	}

	if !dara.IsNil(request.BW) {
		body["BW"] = request.BW
	}

	if !dara.IsNil(request.BX) {
		body["BX"] = request.BX
	}

	if !dara.IsNil(request.BY) {
		body["BY"] = request.BY
	}

	if !dara.IsNil(request.VideoUrl) {
		body["VideoUrl"] = request.VideoUrl
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EraseVideoSubtitles"),
		Version:     dara.String("2020-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EraseVideoSubtitlesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视频人像卡通化
//
// @param request - GenerateHumanAnimeStyleVideoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateHumanAnimeStyleVideoResponse
func (client *Client) GenerateHumanAnimeStyleVideoWithContext(ctx context.Context, request *GenerateHumanAnimeStyleVideoRequest, runtime *dara.RuntimeOptions) (_result *GenerateHumanAnimeStyleVideoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CartoonStyle) {
		body["CartoonStyle"] = request.CartoonStyle
	}

	if !dara.IsNil(request.VideoUrl) {
		body["VideoUrl"] = request.VideoUrl
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateHumanAnimeStyleVideo"),
		Version:     dara.String("2020-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateHumanAnimeStyleVideoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GenerateVideoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateVideoResponse
func (client *Client) GenerateVideoWithContext(ctx context.Context, request *GenerateVideoRequest, runtime *dara.RuntimeOptions) (_result *GenerateVideoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Duration) {
		body["Duration"] = request.Duration
	}

	if !dara.IsNil(request.DurationAdaption) {
		body["DurationAdaption"] = request.DurationAdaption
	}

	if !dara.IsNil(request.FileList) {
		body["FileList"] = request.FileList
	}

	if !dara.IsNil(request.Height) {
		body["Height"] = request.Height
	}

	if !dara.IsNil(request.Mute) {
		body["Mute"] = request.Mute
	}

	if !dara.IsNil(request.PuzzleEffect) {
		body["PuzzleEffect"] = request.PuzzleEffect
	}

	if !dara.IsNil(request.Scene) {
		body["Scene"] = request.Scene
	}

	if !dara.IsNil(request.SmartEffect) {
		body["SmartEffect"] = request.SmartEffect
	}

	if !dara.IsNil(request.Style) {
		body["Style"] = request.Style
	}

	if !dara.IsNil(request.TransitionStyle) {
		body["TransitionStyle"] = request.TransitionStyle
	}

	if !dara.IsNil(request.Width) {
		body["Width"] = request.Width
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateVideo"),
		Version:     dara.String("2020-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateVideoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetAsyncJobResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAsyncJobResultResponse
func (client *Client) GetAsyncJobResultWithContext(ctx context.Context, request *GetAsyncJobResultRequest, runtime *dara.RuntimeOptions) (_result *GetAsyncJobResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		body["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAsyncJobResult"),
		Version:     dara.String("2020-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAsyncJobResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - InterpolateVideoFrameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InterpolateVideoFrameResponse
func (client *Client) InterpolateVideoFrameWithContext(ctx context.Context, request *InterpolateVideoFrameRequest, runtime *dara.RuntimeOptions) (_result *InterpolateVideoFrameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Bitrate) {
		body["Bitrate"] = request.Bitrate
	}

	if !dara.IsNil(request.FrameRate) {
		body["FrameRate"] = request.FrameRate
	}

	if !dara.IsNil(request.VideoURL) {
		body["VideoURL"] = request.VideoURL
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InterpolateVideoFrame"),
		Version:     dara.String("2020-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InterpolateVideoFrameResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - MergeVideoFaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MergeVideoFaceResponse
func (client *Client) MergeVideoFaceWithContext(ctx context.Context, request *MergeVideoFaceRequest, runtime *dara.RuntimeOptions) (_result *MergeVideoFaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AddWatermark) {
		body["AddWatermark"] = request.AddWatermark
	}

	if !dara.IsNil(request.Enhance) {
		body["Enhance"] = request.Enhance
	}

	if !dara.IsNil(request.ReferenceURL) {
		body["ReferenceURL"] = request.ReferenceURL
	}

	if !dara.IsNil(request.VideoURL) {
		body["VideoURL"] = request.VideoURL
	}

	if !dara.IsNil(request.WatermarkType) {
		body["WatermarkType"] = request.WatermarkType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MergeVideoFace"),
		Version:     dara.String("2020-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MergeVideoFaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视频模板融合换脸
//
// @param request - MergeVideoModelFaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MergeVideoModelFaceResponse
func (client *Client) MergeVideoModelFaceWithContext(ctx context.Context, request *MergeVideoModelFaceRequest, runtime *dara.RuntimeOptions) (_result *MergeVideoModelFaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AddWatermark) {
		body["AddWatermark"] = request.AddWatermark
	}

	if !dara.IsNil(request.Enhance) {
		body["Enhance"] = request.Enhance
	}

	if !dara.IsNil(request.FaceImageURL) {
		body["FaceImageURL"] = request.FaceImageURL
	}

	if !dara.IsNil(request.MergeInfos) {
		body["MergeInfos"] = request.MergeInfos
	}

	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.WatermarkType) {
		body["WatermarkType"] = request.WatermarkType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MergeVideoModelFace"),
		Version:     dara.String("2020-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MergeVideoModelFaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视频人脸融合模板查询
//
// @param request - QueryFaceVideoTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryFaceVideoTemplateResponse
func (client *Client) QueryFaceVideoTemplateWithContext(ctx context.Context, request *QueryFaceVideoTemplateRequest, runtime *dara.RuntimeOptions) (_result *QueryFaceVideoTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryFaceVideoTemplate"),
		Version:     dara.String("2020-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryFaceVideoTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SuperResolveVideoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SuperResolveVideoResponse
func (client *Client) SuperResolveVideoWithContext(ctx context.Context, request *SuperResolveVideoRequest, runtime *dara.RuntimeOptions) (_result *SuperResolveVideoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BitRate) {
		body["BitRate"] = request.BitRate
	}

	if !dara.IsNil(request.VideoUrl) {
		body["VideoUrl"] = request.VideoUrl
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SuperResolveVideo"),
		Version:     dara.String("2020-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SuperResolveVideoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
