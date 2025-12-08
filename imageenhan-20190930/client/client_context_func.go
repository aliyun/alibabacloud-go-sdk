// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// @param request - AssessCompositionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssessCompositionResponse
func (client *Client) AssessCompositionWithContext(ctx context.Context, request *AssessCompositionRequest, runtime *dara.RuntimeOptions) (_result *AssessCompositionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageURL) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssessComposition"),
		Version:     dara.String("2019-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssessCompositionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AssessExposureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssessExposureResponse
func (client *Client) AssessExposureWithContext(ctx context.Context, request *AssessExposureRequest, runtime *dara.RuntimeOptions) (_result *AssessExposureResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageURL) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssessExposure"),
		Version:     dara.String("2019-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssessExposureResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AssessSharpnessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssessSharpnessResponse
func (client *Client) AssessSharpnessWithContext(ctx context.Context, request *AssessSharpnessRequest, runtime *dara.RuntimeOptions) (_result *AssessSharpnessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageURL) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssessSharpness"),
		Version:     dara.String("2019-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssessSharpnessResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ChangeImageSizeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeImageSizeResponse
func (client *Client) ChangeImageSizeWithContext(ctx context.Context, request *ChangeImageSizeRequest, runtime *dara.RuntimeOptions) (_result *ChangeImageSizeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Height) {
		body["Height"] = request.Height
	}

	if !dara.IsNil(request.Url) {
		body["Url"] = request.Url
	}

	if !dara.IsNil(request.Width) {
		body["Width"] = request.Width
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeImageSize"),
		Version:     dara.String("2019-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeImageSizeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ColorizeImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ColorizeImageResponse
func (client *Client) ColorizeImageWithContext(ctx context.Context, request *ColorizeImageRequest, runtime *dara.RuntimeOptions) (_result *ColorizeImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageURL) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ColorizeImage"),
		Version:     dara.String("2019-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ColorizeImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - EnhanceImageColorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnhanceImageColorResponse
func (client *Client) EnhanceImageColorWithContext(ctx context.Context, request *EnhanceImageColorRequest, runtime *dara.RuntimeOptions) (_result *EnhanceImageColorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageURL) {
		body["ImageURL"] = request.ImageURL
	}

	if !dara.IsNil(request.Mode) {
		body["Mode"] = request.Mode
	}

	if !dara.IsNil(request.OutputFormat) {
		body["OutputFormat"] = request.OutputFormat
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnhanceImageColor"),
		Version:     dara.String("2019-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnhanceImageColorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ErasePersonRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ErasePersonResponse
func (client *Client) ErasePersonWithContext(ctx context.Context, request *ErasePersonRequest, runtime *dara.RuntimeOptions) (_result *ErasePersonResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageURL) {
		body["ImageURL"] = request.ImageURL
	}

	if !dara.IsNil(request.UserMask) {
		body["UserMask"] = request.UserMask
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ErasePerson"),
		Version:     dara.String("2019-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ErasePersonResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ExtendImageStyleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExtendImageStyleResponse
func (client *Client) ExtendImageStyleWithContext(ctx context.Context, request *ExtendImageStyleRequest, runtime *dara.RuntimeOptions) (_result *ExtendImageStyleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MajorUrl) {
		body["MajorUrl"] = request.MajorUrl
	}

	if !dara.IsNil(request.StyleUrl) {
		body["StyleUrl"] = request.StyleUrl
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExtendImageStyle"),
		Version:     dara.String("2019-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExtendImageStyleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成式图像卡通化
//
// @param request - GenerateCartoonizedImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateCartoonizedImageResponse
func (client *Client) GenerateCartoonizedImageWithContext(ctx context.Context, request *GenerateCartoonizedImageRequest, runtime *dara.RuntimeOptions) (_result *GenerateCartoonizedImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageType) {
		body["ImageType"] = request.ImageType
	}

	if !dara.IsNil(request.ImageUrl) {
		body["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.Index) {
		body["Index"] = request.Index
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateCartoonizedImage"),
		Version:     dara.String("2019-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateCartoonizedImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成式图像超分
//
// @param request - GenerateSuperResolutionImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateSuperResolutionImageResponse
func (client *Client) GenerateSuperResolutionImageWithContext(ctx context.Context, request *GenerateSuperResolutionImageRequest, runtime *dara.RuntimeOptions) (_result *GenerateSuperResolutionImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageUrl) {
		body["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.OutputFormat) {
		body["OutputFormat"] = request.OutputFormat
	}

	if !dara.IsNil(request.OutputQuality) {
		body["OutputQuality"] = request.OutputQuality
	}

	if !dara.IsNil(request.Scale) {
		body["Scale"] = request.Scale
	}

	if !dara.IsNil(request.UserData) {
		body["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateSuperResolutionImage"),
		Version:     dara.String("2019-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateSuperResolutionImageResponse{}
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
		Version:     dara.String("2019-09-30"),
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

// @param request - ImageBlindCharacterWatermarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImageBlindCharacterWatermarkResponse
func (client *Client) ImageBlindCharacterWatermarkWithContext(ctx context.Context, request *ImageBlindCharacterWatermarkRequest, runtime *dara.RuntimeOptions) (_result *ImageBlindCharacterWatermarkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FunctionType) {
		body["FunctionType"] = request.FunctionType
	}

	if !dara.IsNil(request.OriginImageURL) {
		body["OriginImageURL"] = request.OriginImageURL
	}

	if !dara.IsNil(request.OutputFileType) {
		body["OutputFileType"] = request.OutputFileType
	}

	if !dara.IsNil(request.QualityFactor) {
		body["QualityFactor"] = request.QualityFactor
	}

	if !dara.IsNil(request.Text) {
		body["Text"] = request.Text
	}

	if !dara.IsNil(request.WatermarkImageURL) {
		body["WatermarkImageURL"] = request.WatermarkImageURL
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImageBlindCharacterWatermark"),
		Version:     dara.String("2019-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImageBlindCharacterWatermarkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ImageBlindPicWatermarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImageBlindPicWatermarkResponse
func (client *Client) ImageBlindPicWatermarkWithContext(ctx context.Context, request *ImageBlindPicWatermarkRequest, runtime *dara.RuntimeOptions) (_result *ImageBlindPicWatermarkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FunctionType) {
		body["FunctionType"] = request.FunctionType
	}

	if !dara.IsNil(request.LogoURL) {
		body["LogoURL"] = request.LogoURL
	}

	if !dara.IsNil(request.OriginImageURL) {
		body["OriginImageURL"] = request.OriginImageURL
	}

	if !dara.IsNil(request.OutputFileType) {
		body["OutputFileType"] = request.OutputFileType
	}

	if !dara.IsNil(request.QualityFactor) {
		body["QualityFactor"] = request.QualityFactor
	}

	if !dara.IsNil(request.WatermarkImageURL) {
		body["WatermarkImageURL"] = request.WatermarkImageURL
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImageBlindPicWatermark"),
		Version:     dara.String("2019-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImageBlindPicWatermarkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ImitatePhotoStyleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImitatePhotoStyleResponse
func (client *Client) ImitatePhotoStyleWithContext(ctx context.Context, request *ImitatePhotoStyleRequest, runtime *dara.RuntimeOptions) (_result *ImitatePhotoStyleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageURL) {
		body["ImageURL"] = request.ImageURL
	}

	if !dara.IsNil(request.StyleUrl) {
		body["StyleUrl"] = request.StyleUrl
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImitatePhotoStyle"),
		Version:     dara.String("2019-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImitatePhotoStyleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - IntelligentCompositionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IntelligentCompositionResponse
func (client *Client) IntelligentCompositionWithContext(ctx context.Context, request *IntelligentCompositionRequest, runtime *dara.RuntimeOptions) (_result *IntelligentCompositionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageURL) {
		body["ImageURL"] = request.ImageURL
	}

	if !dara.IsNil(request.NumBoxes) {
		body["NumBoxes"] = request.NumBoxes
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IntelligentComposition"),
		Version:     dara.String("2019-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &IntelligentCompositionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - MakeSuperResolutionImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MakeSuperResolutionImageResponse
func (client *Client) MakeSuperResolutionImageWithContext(ctx context.Context, request *MakeSuperResolutionImageRequest, runtime *dara.RuntimeOptions) (_result *MakeSuperResolutionImageResponse, _err error) {
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

	if !dara.IsNil(request.OutputFormat) {
		body["OutputFormat"] = request.OutputFormat
	}

	if !dara.IsNil(request.OutputQuality) {
		body["OutputQuality"] = request.OutputQuality
	}

	if !dara.IsNil(request.UpscaleFactor) {
		body["UpscaleFactor"] = request.UpscaleFactor
	}

	if !dara.IsNil(request.Url) {
		body["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MakeSuperResolutionImage"),
		Version:     dara.String("2019-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MakeSuperResolutionImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RecolorHDImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecolorHDImageResponse
func (client *Client) RecolorHDImageWithContext(ctx context.Context, request *RecolorHDImageRequest, runtime *dara.RuntimeOptions) (_result *RecolorHDImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ColorCount) {
		body["ColorCount"] = request.ColorCount
	}

	if !dara.IsNil(request.ColorTemplate) {
		body["ColorTemplate"] = request.ColorTemplate
	}

	if !dara.IsNil(request.Degree) {
		body["Degree"] = request.Degree
	}

	if !dara.IsNil(request.Mode) {
		body["Mode"] = request.Mode
	}

	if !dara.IsNil(request.RefUrl) {
		body["RefUrl"] = request.RefUrl
	}

	if !dara.IsNil(request.Url) {
		body["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RecolorHDImage"),
		Version:     dara.String("2019-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RecolorHDImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RecolorImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecolorImageResponse
func (client *Client) RecolorImageWithContext(ctx context.Context, request *RecolorImageRequest, runtime *dara.RuntimeOptions) (_result *RecolorImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ColorCount) {
		body["ColorCount"] = request.ColorCount
	}

	if !dara.IsNil(request.ColorTemplate) {
		body["ColorTemplate"] = request.ColorTemplate
	}

	if !dara.IsNil(request.Mode) {
		body["Mode"] = request.Mode
	}

	if !dara.IsNil(request.RefUrl) {
		body["RefUrl"] = request.RefUrl
	}

	if !dara.IsNil(request.Url) {
		body["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RecolorImage"),
		Version:     dara.String("2019-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RecolorImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RemoveImageSubtitlesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveImageSubtitlesResponse
func (client *Client) RemoveImageSubtitlesWithContext(ctx context.Context, request *RemoveImageSubtitlesRequest, runtime *dara.RuntimeOptions) (_result *RemoveImageSubtitlesResponse, _err error) {
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

	if !dara.IsNil(request.ImageURL) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveImageSubtitles"),
		Version:     dara.String("2019-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveImageSubtitlesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RemoveImageWatermarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveImageWatermarkResponse
func (client *Client) RemoveImageWatermarkWithContext(ctx context.Context, request *RemoveImageWatermarkRequest, runtime *dara.RuntimeOptions) (_result *RemoveImageWatermarkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageURL) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveImageWatermark"),
		Version:     dara.String("2019-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveImageWatermarkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
