// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// @param request - AddFaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddFaceResponse
func (client *Client) AddFaceWithContext(ctx context.Context, request *AddFaceRequest, runtime *dara.RuntimeOptions) (_result *AddFaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		body["DbName"] = request.DbName
	}

	if !dara.IsNil(request.EntityId) {
		body["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.ExtraData) {
		body["ExtraData"] = request.ExtraData
	}

	if !dara.IsNil(request.ImageUrl) {
		body["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.QualityScoreThreshold) {
		body["QualityScoreThreshold"] = request.QualityScoreThreshold
	}

	if !dara.IsNil(request.SimilarityScoreThresholdBetweenEntity) {
		body["SimilarityScoreThresholdBetweenEntity"] = request.SimilarityScoreThresholdBetweenEntity
	}

	if !dara.IsNil(request.SimilarityScoreThresholdInEntity) {
		body["SimilarityScoreThresholdInEntity"] = request.SimilarityScoreThresholdInEntity
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddFace"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddFaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddFaceEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddFaceEntityResponse
func (client *Client) AddFaceEntityWithContext(ctx context.Context, request *AddFaceEntityRequest, runtime *dara.RuntimeOptions) (_result *AddFaceEntityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		body["DbName"] = request.DbName
	}

	if !dara.IsNil(request.EntityId) {
		body["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.Labels) {
		body["Labels"] = request.Labels
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddFaceEntity"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddFaceEntityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 图像人脸融合模板增加
//
// @param request - AddFaceImageTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddFaceImageTemplateResponse
func (client *Client) AddFaceImageTemplateWithContext(ctx context.Context, request *AddFaceImageTemplateRequest, runtime *dara.RuntimeOptions) (_result *AddFaceImageTemplateResponse, _err error) {
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
		Action:      dara.String("AddFaceImageTemplate"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddFaceImageTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量添加人脸数据
//
// @param tmpReq - BatchAddFacesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchAddFacesResponse
func (client *Client) BatchAddFacesWithContext(ctx context.Context, tmpReq *BatchAddFacesRequest, runtime *dara.RuntimeOptions) (_result *BatchAddFacesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchAddFacesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Faces) {
		request.FacesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Faces, dara.String("Faces"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		body["DbName"] = request.DbName
	}

	if !dara.IsNil(request.EntityId) {
		body["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.FacesShrink) {
		body["Faces"] = request.FacesShrink
	}

	if !dara.IsNil(request.QualityScoreThreshold) {
		body["QualityScoreThreshold"] = request.QualityScoreThreshold
	}

	if !dara.IsNil(request.SimilarityScoreThresholdBetweenEntity) {
		body["SimilarityScoreThresholdBetweenEntity"] = request.SimilarityScoreThresholdBetweenEntity
	}

	if !dara.IsNil(request.SimilarityScoreThresholdInEntity) {
		body["SimilarityScoreThresholdInEntity"] = request.SimilarityScoreThresholdInEntity
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchAddFaces"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchAddFacesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - BlurFaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BlurFaceResponse
func (client *Client) BlurFaceWithContext(ctx context.Context, request *BlurFaceRequest, runtime *dara.RuntimeOptions) (_result *BlurFaceResponse, _err error) {
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
		Action:      dara.String("BlurFace"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BlurFaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - BodyPostureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BodyPostureResponse
func (client *Client) BodyPostureWithContext(ctx context.Context, request *BodyPostureRequest, runtime *dara.RuntimeOptions) (_result *BodyPostureResponse, _err error) {
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
		Action:      dara.String("BodyPosture"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BodyPostureResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 人脸比对(1:1)
//
// @param request - CompareFaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompareFaceResponse
func (client *Client) CompareFaceWithContext(ctx context.Context, request *CompareFaceRequest, runtime *dara.RuntimeOptions) (_result *CompareFaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageDataA) {
		body["ImageDataA"] = request.ImageDataA
	}

	if !dara.IsNil(request.ImageDataB) {
		body["ImageDataB"] = request.ImageDataB
	}

	if !dara.IsNil(request.ImageURLA) {
		body["ImageURLA"] = request.ImageURLA
	}

	if !dara.IsNil(request.ImageURLB) {
		body["ImageURLB"] = request.ImageURLB
	}

	if !dara.IsNil(request.QualityScoreThreshold) {
		body["QualityScoreThreshold"] = request.QualityScoreThreshold
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CompareFace"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CompareFaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 口罩人脸比对1:1
//
// @param request - CompareFaceWithMaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompareFaceWithMaskResponse
func (client *Client) CompareFaceWithMaskWithContext(ctx context.Context, request *CompareFaceWithMaskRequest, runtime *dara.RuntimeOptions) (_result *CompareFaceWithMaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageURLA) {
		body["ImageURLA"] = request.ImageURLA
	}

	if !dara.IsNil(request.ImageURLB) {
		body["ImageURLB"] = request.ImageURLB
	}

	if !dara.IsNil(request.QualityScoreThreshold) {
		body["QualityScoreThreshold"] = request.QualityScoreThreshold
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CompareFaceWithMask"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CompareFaceWithMaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateFaceDbRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFaceDbResponse
func (client *Client) CreateFaceDbWithContext(ctx context.Context, request *CreateFaceDbRequest, runtime *dara.RuntimeOptions) (_result *CreateFaceDbResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFaceDb"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFaceDbResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 换脸鉴别
//
// @param request - DeepfakeFaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeepfakeFaceResponse
func (client *Client) DeepfakeFaceWithContext(ctx context.Context, request *DeepfakeFaceRequest, runtime *dara.RuntimeOptions) (_result *DeepfakeFaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Tasks) {
		body["Tasks"] = request.Tasks
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeepfakeFace"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeepfakeFaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteFaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFaceResponse
func (client *Client) DeleteFaceWithContext(ctx context.Context, request *DeleteFaceRequest, runtime *dara.RuntimeOptions) (_result *DeleteFaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		body["DbName"] = request.DbName
	}

	if !dara.IsNil(request.FaceId) {
		body["FaceId"] = request.FaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFace"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteFaceDbRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFaceDbResponse
func (client *Client) DeleteFaceDbWithContext(ctx context.Context, request *DeleteFaceDbRequest, runtime *dara.RuntimeOptions) (_result *DeleteFaceDbResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFaceDb"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFaceDbResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteFaceEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFaceEntityResponse
func (client *Client) DeleteFaceEntityWithContext(ctx context.Context, request *DeleteFaceEntityRequest, runtime *dara.RuntimeOptions) (_result *DeleteFaceEntityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		body["DbName"] = request.DbName
	}

	if !dara.IsNil(request.EntityId) {
		body["EntityId"] = request.EntityId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFaceEntity"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFaceEntityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 图像人脸融合模板删除
//
// @param request - DeleteFaceImageTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFaceImageTemplateResponse
func (client *Client) DeleteFaceImageTemplateWithContext(ctx context.Context, request *DeleteFaceImageTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteFaceImageTemplateResponse, _err error) {
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
		Action:      dara.String("DeleteFaceImageTemplate"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFaceImageTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DetectBodyCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectBodyCountResponse
func (client *Client) DetectBodyCountWithContext(ctx context.Context, request *DetectBodyCountRequest, runtime *dara.RuntimeOptions) (_result *DetectBodyCountResponse, _err error) {
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
		Action:      dara.String("DetectBodyCount"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetectBodyCountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DetectCelebrityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectCelebrityResponse
func (client *Client) DetectCelebrityWithContext(ctx context.Context, request *DetectCelebrityRequest, runtime *dara.RuntimeOptions) (_result *DetectCelebrityResponse, _err error) {
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
		Action:      dara.String("DetectCelebrity"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetectCelebrityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DetectFaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectFaceResponse
func (client *Client) DetectFaceWithContext(ctx context.Context, request *DetectFaceRequest, runtime *dara.RuntimeOptions) (_result *DetectFaceResponse, _err error) {
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

	if !dara.IsNil(request.Landmark) {
		body["Landmark"] = request.Landmark
	}

	if !dara.IsNil(request.MaxFaceNumber) {
		body["MaxFaceNumber"] = request.MaxFaceNumber
	}

	if !dara.IsNil(request.Pose) {
		body["Pose"] = request.Pose
	}

	if !dara.IsNil(request.Quality) {
		body["Quality"] = request.Quality
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetectFace"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetectFaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 红外人脸活体检测
//
// @param request - DetectInfraredLivingFaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectInfraredLivingFaceResponse
func (client *Client) DetectInfraredLivingFaceWithContext(ctx context.Context, request *DetectInfraredLivingFaceRequest, runtime *dara.RuntimeOptions) (_result *DetectInfraredLivingFaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Tasks) {
		body["Tasks"] = request.Tasks
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetectInfraredLivingFace"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetectInfraredLivingFaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 人脸活体检测(DetectLivingFace)
//
// @param request - DetectLivingFaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectLivingFaceResponse
func (client *Client) DetectLivingFaceWithContext(ctx context.Context, request *DetectLivingFaceRequest, runtime *dara.RuntimeOptions) (_result *DetectLivingFaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Tasks) {
		body["Tasks"] = request.Tasks
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetectLivingFace"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetectLivingFaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DetectPedestrianRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectPedestrianResponse
func (client *Client) DetectPedestrianWithContext(ctx context.Context, request *DetectPedestrianRequest, runtime *dara.RuntimeOptions) (_result *DetectPedestrianResponse, _err error) {
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
		Action:      dara.String("DetectPedestrian"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetectPedestrianResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DetectVideoLivingFaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectVideoLivingFaceResponse
func (client *Client) DetectVideoLivingFaceWithContext(ctx context.Context, request *DetectVideoLivingFaceRequest, runtime *dara.RuntimeOptions) (_result *DetectVideoLivingFaceResponse, _err error) {
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
		Action:      dara.String("DetectVideoLivingFace"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetectVideoLivingFaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - EnhanceFaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnhanceFaceResponse
func (client *Client) EnhanceFaceWithContext(ctx context.Context, request *EnhanceFaceRequest, runtime *dara.RuntimeOptions) (_result *EnhanceFaceResponse, _err error) {
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
		Action:      dara.String("EnhanceFace"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnhanceFaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 指纹提取
//
// @param request - ExtractFingerPrintRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExtractFingerPrintResponse
func (client *Client) ExtractFingerPrintWithContext(ctx context.Context, request *ExtractFingerPrintRequest, runtime *dara.RuntimeOptions) (_result *ExtractFingerPrintResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageData) {
		body["ImageData"] = request.ImageData
	}

	if !dara.IsNil(request.ImageURL) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExtractFingerPrint"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExtractFingerPrintResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - FaceBeautyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FaceBeautyResponse
func (client *Client) FaceBeautyWithContext(ctx context.Context, request *FaceBeautyRequest, runtime *dara.RuntimeOptions) (_result *FaceBeautyResponse, _err error) {
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

	if !dara.IsNil(request.Sharp) {
		body["Sharp"] = request.Sharp
	}

	if !dara.IsNil(request.Smooth) {
		body["Smooth"] = request.Smooth
	}

	if !dara.IsNil(request.White) {
		body["White"] = request.White
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FaceBeauty"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FaceBeautyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GenRealPersonVerificationTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenRealPersonVerificationTokenResponse
func (client *Client) GenRealPersonVerificationTokenWithContext(ctx context.Context, request *GenRealPersonVerificationTokenRequest, runtime *dara.RuntimeOptions) (_result *GenRealPersonVerificationTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CertificateName) {
		body["CertificateName"] = request.CertificateName
	}

	if !dara.IsNil(request.CertificateNumber) {
		body["CertificateNumber"] = request.CertificateNumber
	}

	if !dara.IsNil(request.MetaInfo) {
		body["MetaInfo"] = request.MetaInfo
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenRealPersonVerificationToken"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenRealPersonVerificationTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GenerateHumanAnimeStyleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateHumanAnimeStyleResponse
func (client *Client) GenerateHumanAnimeStyleWithContext(ctx context.Context, request *GenerateHumanAnimeStyleRequest, runtime *dara.RuntimeOptions) (_result *GenerateHumanAnimeStyleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlgoType) {
		query["AlgoType"] = request.AlgoType
	}

	if !dara.IsNil(request.ImageURL) {
		query["ImageURL"] = request.ImageURL
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateHumanAnimeStyle"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateHumanAnimeStyleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 人像素描风格化
//
// @param request - GenerateHumanSketchStyleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateHumanSketchStyleResponse
func (client *Client) GenerateHumanSketchStyleWithContext(ctx context.Context, request *GenerateHumanSketchStyleRequest, runtime *dara.RuntimeOptions) (_result *GenerateHumanSketchStyleResponse, _err error) {
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

	if !dara.IsNil(request.ReturnType) {
		body["ReturnType"] = request.ReturnType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateHumanSketchStyle"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateHumanSketchStyleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetFaceEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFaceEntityResponse
func (client *Client) GetFaceEntityWithContext(ctx context.Context, request *GetFaceEntityRequest, runtime *dara.RuntimeOptions) (_result *GetFaceEntityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		body["DbName"] = request.DbName
	}

	if !dara.IsNil(request.EntityId) {
		body["EntityId"] = request.EntityId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFaceEntity"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFaceEntityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetRealPersonVerificationResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRealPersonVerificationResultResponse
func (client *Client) GetRealPersonVerificationResultWithContext(ctx context.Context, request *GetRealPersonVerificationResultRequest, runtime *dara.RuntimeOptions) (_result *GetRealPersonVerificationResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.VerificationToken) {
		body["VerificationToken"] = request.VerificationToken
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRealPersonVerificationResult"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRealPersonVerificationResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 智能瘦脸
//
// @param request - LiquifyFaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LiquifyFaceResponse
func (client *Client) LiquifyFaceWithContext(ctx context.Context, request *LiquifyFaceRequest, runtime *dara.RuntimeOptions) (_result *LiquifyFaceResponse, _err error) {
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

	if !dara.IsNil(request.SlimDegree) {
		body["SlimDegree"] = request.SlimDegree
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LiquifyFace"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LiquifyFaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListFaceDbsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFaceDbsResponse
func (client *Client) ListFaceDbsWithContext(ctx context.Context, request *ListFaceDbsRequest, runtime *dara.RuntimeOptions) (_result *ListFaceDbsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Limit) {
		body["Limit"] = request.Limit
	}

	if !dara.IsNil(request.Offset) {
		body["Offset"] = request.Offset
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFaceDbs"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFaceDbsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListFaceEntitiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFaceEntitiesResponse
func (client *Client) ListFaceEntitiesWithContext(ctx context.Context, request *ListFaceEntitiesRequest, runtime *dara.RuntimeOptions) (_result *ListFaceEntitiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		body["DbName"] = request.DbName
	}

	if !dara.IsNil(request.EntityIdPrefix) {
		body["EntityIdPrefix"] = request.EntityIdPrefix
	}

	if !dara.IsNil(request.Labels) {
		body["Labels"] = request.Labels
	}

	if !dara.IsNil(request.Limit) {
		body["Limit"] = request.Limit
	}

	if !dara.IsNil(request.Offset) {
		body["Offset"] = request.Offset
	}

	if !dara.IsNil(request.Order) {
		body["Order"] = request.Order
	}

	if !dara.IsNil(request.Token) {
		body["Token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFaceEntities"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFaceEntitiesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 图像人脸融合
//
// @param request - MergeImageFaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MergeImageFaceResponse
func (client *Client) MergeImageFaceWithContext(ctx context.Context, request *MergeImageFaceRequest, runtime *dara.RuntimeOptions) (_result *MergeImageFaceResponse, _err error) {
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

	if !dara.IsNil(request.ImageURL) {
		body["ImageURL"] = request.ImageURL
	}

	if !dara.IsNil(request.MergeInfos) {
		body["MergeInfos"] = request.MergeInfos
	}

	if !dara.IsNil(request.ModelVersion) {
		body["ModelVersion"] = request.ModelVersion
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
		Action:      dara.String("MergeImageFace"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MergeImageFaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 线上监考
//
// @param request - MonitorExaminationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MonitorExaminationResponse
func (client *Client) MonitorExaminationWithContext(ctx context.Context, request *MonitorExaminationRequest, runtime *dara.RuntimeOptions) (_result *MonitorExaminationResponse, _err error) {
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

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MonitorExamination"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MonitorExaminationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - PedestrianDetectAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PedestrianDetectAttributeResponse
func (client *Client) PedestrianDetectAttributeWithContext(ctx context.Context, request *PedestrianDetectAttributeRequest, runtime *dara.RuntimeOptions) (_result *PedestrianDetectAttributeResponse, _err error) {
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
		Action:      dara.String("PedestrianDetectAttribute"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PedestrianDetectAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 图像人脸融合模板查询
//
// @param request - QueryFaceImageTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryFaceImageTemplateResponse
func (client *Client) QueryFaceImageTemplateWithContext(ctx context.Context, request *QueryFaceImageTemplateRequest, runtime *dara.RuntimeOptions) (_result *QueryFaceImageTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryFaceImageTemplate"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryFaceImageTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 动作行为识别(RecognizeAction)
//
// @param request - RecognizeActionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeActionResponse
func (client *Client) RecognizeActionWithContext(ctx context.Context, request *RecognizeActionRequest, runtime *dara.RuntimeOptions) (_result *RecognizeActionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.URLList) {
		body["URLList"] = request.URLList
	}

	if !dara.IsNil(request.VideoData) {
		body["VideoData"] = request.VideoData
	}

	if !dara.IsNil(request.VideoUrl) {
		body["VideoUrl"] = request.VideoUrl
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RecognizeAction"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RecognizeActionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RecognizeExpressionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeExpressionResponse
func (client *Client) RecognizeExpressionWithContext(ctx context.Context, request *RecognizeExpressionRequest, runtime *dara.RuntimeOptions) (_result *RecognizeExpressionResponse, _err error) {
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
		Action:      dara.String("RecognizeExpression"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RecognizeExpressionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RecognizeFaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeFaceResponse
func (client *Client) RecognizeFaceWithContext(ctx context.Context, request *RecognizeFaceRequest, runtime *dara.RuntimeOptions) (_result *RecognizeFaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Age) {
		body["Age"] = request.Age
	}

	if !dara.IsNil(request.Beauty) {
		body["Beauty"] = request.Beauty
	}

	if !dara.IsNil(request.Expression) {
		body["Expression"] = request.Expression
	}

	if !dara.IsNil(request.Gender) {
		body["Gender"] = request.Gender
	}

	if !dara.IsNil(request.Glass) {
		body["Glass"] = request.Glass
	}

	if !dara.IsNil(request.Hat) {
		body["Hat"] = request.Hat
	}

	if !dara.IsNil(request.ImageURL) {
		body["ImageURL"] = request.ImageURL
	}

	if !dara.IsNil(request.Mask) {
		body["Mask"] = request.Mask
	}

	if !dara.IsNil(request.MaxFaceNumber) {
		body["MaxFaceNumber"] = request.MaxFaceNumber
	}

	if !dara.IsNil(request.Quality) {
		body["Quality"] = request.Quality
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RecognizeFace"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RecognizeFaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 公众人脸识别(RecognizePublicFace)
//
// @param request - RecognizePublicFaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizePublicFaceResponse
func (client *Client) RecognizePublicFaceWithContext(ctx context.Context, request *RecognizePublicFaceRequest, runtime *dara.RuntimeOptions) (_result *RecognizePublicFaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Task) {
		body["Task"] = request.Task
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RecognizePublicFace"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RecognizePublicFaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 美肤
//
// @param request - RetouchSkinRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetouchSkinResponse
func (client *Client) RetouchSkinWithContext(ctx context.Context, request *RetouchSkinRequest, runtime *dara.RuntimeOptions) (_result *RetouchSkinResponse, _err error) {
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

	if !dara.IsNil(request.RetouchDegree) {
		body["RetouchDegree"] = request.RetouchDegree
	}

	if !dara.IsNil(request.WhiteningDegree) {
		body["WhiteningDegree"] = request.WhiteningDegree
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RetouchSkin"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RetouchSkinResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// summary
//
// @param request - SearchFaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchFaceResponse
func (client *Client) SearchFaceWithContext(ctx context.Context, request *SearchFaceRequest, runtime *dara.RuntimeOptions) (_result *SearchFaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		body["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DbNames) {
		body["DbNames"] = request.DbNames
	}

	if !dara.IsNil(request.ImageUrl) {
		body["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.Limit) {
		body["Limit"] = request.Limit
	}

	if !dara.IsNil(request.MaxFaceNum) {
		body["MaxFaceNum"] = request.MaxFaceNum
	}

	if !dara.IsNil(request.QualityScoreThreshold) {
		body["QualityScoreThreshold"] = request.QualityScoreThreshold
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchFace"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchFaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateFaceEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFaceEntityResponse
func (client *Client) UpdateFaceEntityWithContext(ctx context.Context, request *UpdateFaceEntityRequest, runtime *dara.RuntimeOptions) (_result *UpdateFaceEntityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		body["DbName"] = request.DbName
	}

	if !dara.IsNil(request.EntityId) {
		body["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.Labels) {
		body["Labels"] = request.Labels
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFaceEntity"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFaceEntityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
