// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// # Add AIGC Face Detection Capability
//
// @param request - AIGCFaceVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AIGCFaceVerifyResponse
func (client *Client) AIGCFaceVerifyWithContext(ctx context.Context, request *AIGCFaceVerifyRequest, runtime *dara.RuntimeOptions) (_result *AIGCFaceVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FaceContrastPictureUrl) {
		query["FaceContrastPictureUrl"] = request.FaceContrastPictureUrl
	}

	if !dara.IsNil(request.OssBucketName) {
		query["OssBucketName"] = request.OssBucketName
	}

	if !dara.IsNil(request.OssObjectName) {
		query["OssObjectName"] = request.OssObjectName
	}

	if !dara.IsNil(request.OuterOrderNo) {
		query["OuterOrderNo"] = request.OuterOrderNo
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FaceContrastPicture) {
		body["FaceContrastPicture"] = request.FaceContrastPicture
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AIGCFaceVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AIGCFaceVerifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Bank Card Element Verification Interface
//
// Description:
//
// Bank card verification, including: two elements (name + bank card number), three elements (name + ID number + bank card number), and four elements (name + ID number + mobile phone number + bank card number) consistency verification.
//
// - Service address:
//
//   - Beijing region: cloudauth.cn-beijing.aliyuncs.com (IPv4) or cloudauth-dualstack.cn-beijing.aliyuncs.com (IPv6).
//
//   - Shanghai region: cloudauth.cn-shanghai.aliyuncs.com (IPv4) or cloudauth-dualstack.cn-shanghai.aliyuncs.com (IPv6).
//
// - Request method: POST and GET.
//
// - Transfer protocol: HTTPS.
//
// @param request - BankMetaVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BankMetaVerifyResponse
func (client *Client) BankMetaVerifyWithContext(ctx context.Context, request *BankMetaVerifyRequest, runtime *dara.RuntimeOptions) (_result *BankMetaVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BankCard) {
		query["BankCard"] = request.BankCard
	}

	if !dara.IsNil(request.IdentifyNum) {
		query["IdentifyNum"] = request.IdentifyNum
	}

	if !dara.IsNil(request.IdentityType) {
		query["IdentityType"] = request.IdentityType
	}

	if !dara.IsNil(request.Mobile) {
		query["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.ParamType) {
		query["ParamType"] = request.ParamType
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	if !dara.IsNil(request.VerifyMode) {
		query["VerifyMode"] = request.VerifyMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BankMetaVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BankMetaVerifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Financial-grade Pure Server-Side API for Face Comparison.
//
// Description:
//
// - API Name: CompareFaceVerify.
//
// - Service Address: cloudauth.aliyuncs.com.
//
// - Request Method: HTTPS POST and GET.
//
// - API Description: An interface to achieve real-person authentication through server-side integration.
//
// #### Photo Format Requirements
//
// When performing face comparison, please upload 2 facial photos that meet all the following conditions:
//
// - Recent photo/recent database photo, with a complete, clear, unobstructed face, natural expression, and facing the camera directly.
//
// - Clear photo with normal exposure, no overly dark, overly bright, or halo effects on the face, and no significant angle deviation.
//
// - Resolution not exceeding 1920*1080, at least 640*480, recommended to scale the shorter side to 720 pixels, with a compression ratio greater than 0.9.
//
// - Photo size: <1MB.
//
// - Supports 90, 180, and 270-degree photos; in cases of multiple faces, the largest face will be selected.
//
// @param request - CompareFaceVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompareFaceVerifyResponse
func (client *Client) CompareFaceVerifyWithContext(ctx context.Context, request *CompareFaceVerifyRequest, runtime *dara.RuntimeOptions) (_result *CompareFaceVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Crop) {
		body["Crop"] = request.Crop
	}

	if !dara.IsNil(request.OuterOrderNo) {
		body["OuterOrderNo"] = request.OuterOrderNo
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.SourceCertifyId) {
		body["SourceCertifyId"] = request.SourceCertifyId
	}

	if !dara.IsNil(request.SourceFaceContrastPicture) {
		body["SourceFaceContrastPicture"] = request.SourceFaceContrastPicture
	}

	if !dara.IsNil(request.SourceFaceContrastPictureUrl) {
		body["SourceFaceContrastPictureUrl"] = request.SourceFaceContrastPictureUrl
	}

	if !dara.IsNil(request.SourceOssBucketName) {
		body["SourceOssBucketName"] = request.SourceOssBucketName
	}

	if !dara.IsNil(request.SourceOssObjectName) {
		body["SourceOssObjectName"] = request.SourceOssObjectName
	}

	if !dara.IsNil(request.TargetCertifyId) {
		body["TargetCertifyId"] = request.TargetCertifyId
	}

	if !dara.IsNil(request.TargetFaceContrastPicture) {
		body["TargetFaceContrastPicture"] = request.TargetFaceContrastPicture
	}

	if !dara.IsNil(request.TargetFaceContrastPictureUrl) {
		body["TargetFaceContrastPictureUrl"] = request.TargetFaceContrastPictureUrl
	}

	if !dara.IsNil(request.TargetOssBucketName) {
		body["TargetOssBucketName"] = request.TargetOssBucketName
	}

	if !dara.IsNil(request.TargetOssObjectName) {
		body["TargetOssObjectName"] = request.TargetOssObjectName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CompareFaceVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CompareFaceVerifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke CompareFaces for face comparison.
//
// Description:
//
// Request Method: Only supports sending requests via HTTPS POST.
//
// Interface Description: Compares two face images and outputs the similarity score of the faces in the two images as the result.
//
// - At least one of the specified comparison images should be a face photo (FacePic).
//
// - If an image contains multiple faces, the algorithm will automatically select the largest face in the image.
//
// - If one of the two comparison images does not detect a face, the system will return an error message stating \\"No face detected\\".
//
// When uploading images, you need to provide the HTTP address or base64 encoding of the image.
//
// - HTTP Address: A publicly accessible HTTP address. For example, `http://image-demo.img-cn-hangzhou.aliyuncs.com/example.jpg`.
//
// - Base64 Encoding: An image encoded in base64, formatted as `base64://<base64 string of the image>`.
//
// # Image Restrictions
//
// - Does not support relative or absolute paths for local images.
//
// - Please keep the size of a single image within 2MB to avoid timeout during retrieval by the algorithm.
//
// - The body of a single request has a size limit of 8MB; please calculate the total size of all images and other information in the request to ensure it does not exceed this limit.
//
// - When using base64 to transmit images, the request method must be changed to POST; the header description such as `data:image/png;base64,` should be removed from the base64 string of the image.
//
// @param request - CompareFacesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompareFacesResponse
func (client *Client) CompareFacesWithContext(ctx context.Context, request *CompareFacesRequest, runtime *dara.RuntimeOptions) (_result *CompareFacesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SourceImageType) {
		body["SourceImageType"] = request.SourceImageType
	}

	if !dara.IsNil(request.SourceImageValue) {
		body["SourceImageValue"] = request.SourceImageValue
	}

	if !dara.IsNil(request.TargetImageType) {
		body["TargetImageType"] = request.TargetImageType
	}

	if !dara.IsNil(request.TargetImageValue) {
		body["TargetImageValue"] = request.TargetImageValue
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CompareFaces"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CompareFacesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This interface is used to submit authentication materials for verification and comparison, and it synchronously returns the authentication result.
//
// Description:
//
// - Interface Name: ContrastFaceVerify.
//
// - Service Address: cloudauth.aliyuncs.com.
//
// - Request Method: HTTPS POST and GET.
//
// - Interface Description: An interface for real person authentication through server-side integration.
//
// #### Image Format Requirements
//
// When performing real person authentication, please ensure that the images you upload meet all of the following conditions:
//
// - Recent photo with a clear, unobstructed, and natural expression, facing the camera directly.
//
// - Clear and properly exposed photo, without overly dark, bright, or haloed faces, and with minimal angle deviation.
//
// - Resolution not exceeding 1920*1080, at least 640*480, with the shorter side recommended to be resized to 720 pixels, and a compression ratio greater than 0.9.
//
// - Photo size: <1MB.
//
// - Supports 90, 180, and 270-degree photos; in cases of multiple faces, the largest face will be selected.
//
// @param request - ContrastFaceVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ContrastFaceVerifyResponse
func (client *Client) ContrastFaceVerifyWithContext(ctx context.Context, request *ContrastFaceVerifyRequest, runtime *dara.RuntimeOptions) (_result *ContrastFaceVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Model) {
		query["Model"] = request.Model
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CertName) {
		body["CertName"] = request.CertName
	}

	if !dara.IsNil(request.CertNo) {
		body["CertNo"] = request.CertNo
	}

	if !dara.IsNil(request.CertType) {
		body["CertType"] = request.CertType
	}

	if !dara.IsNil(request.CertifyId) {
		body["CertifyId"] = request.CertifyId
	}

	if !dara.IsNil(request.Crop) {
		body["Crop"] = request.Crop
	}

	if !dara.IsNil(request.DeviceToken) {
		body["DeviceToken"] = request.DeviceToken
	}

	if !dara.IsNil(request.EncryptType) {
		body["EncryptType"] = request.EncryptType
	}

	if !dara.IsNil(request.FaceContrastFile) {
		body["FaceContrastFile"] = request.FaceContrastFile
	}

	if !dara.IsNil(request.FaceContrastPicture) {
		body["FaceContrastPicture"] = request.FaceContrastPicture
	}

	if !dara.IsNil(request.FaceContrastPictureUrl) {
		body["FaceContrastPictureUrl"] = request.FaceContrastPictureUrl
	}

	if !dara.IsNil(request.Ip) {
		body["Ip"] = request.Ip
	}

	if !dara.IsNil(request.Mobile) {
		body["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.OssBucketName) {
		body["OssBucketName"] = request.OssBucketName
	}

	if !dara.IsNil(request.OssObjectName) {
		body["OssObjectName"] = request.OssObjectName
	}

	if !dara.IsNil(request.OuterOrderNo) {
		body["OuterOrderNo"] = request.OuterOrderNo
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ContrastFaceVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ContrastFaceVerifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create a financial-grade authentication scenario
//
// Description:
//
// Request Method: Supports sending requests via HTTPS POST and GET methods.
//
// > The authorization key is valid for 30 minutes and cannot be reused. It is recommended to reacquire it before each activation.
//
// @param request - CreateAntCloudAuthSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAntCloudAuthSceneResponse
func (client *Client) CreateAntCloudAuthSceneWithContext(ctx context.Context, request *CreateAntCloudAuthSceneRequest, runtime *dara.RuntimeOptions) (_result *CreateAntCloudAuthSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BindMiniProgram) {
		query["BindMiniProgram"] = request.BindMiniProgram
	}

	if !dara.IsNil(request.CheckFileBody) {
		query["CheckFileBody"] = request.CheckFileBody
	}

	if !dara.IsNil(request.CheckFileName) {
		query["CheckFileName"] = request.CheckFileName
	}

	if !dara.IsNil(request.MiniProgramName) {
		query["MiniProgramName"] = request.MiniProgramName
	}

	if !dara.IsNil(request.Platform) {
		query["Platform"] = request.Platform
	}

	if !dara.IsNil(request.ReturnPicCount) {
		query["ReturnPicCount"] = request.ReturnPicCount
	}

	if !dara.IsNil(request.ReturnVideoLength) {
		query["ReturnVideoLength"] = request.ReturnVideoLength
	}

	if !dara.IsNil(request.SceneName) {
		query["SceneName"] = request.SceneName
	}

	if !dara.IsNil(request.StoreImage) {
		query["StoreImage"] = request.StoreImage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAntCloudAuthScene"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAntCloudAuthSceneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call CreateAuthKey to get the authorization key, which is used for offline face recognition SDK activation.
//
// Description:
//
// Request Method: Supports sending requests via HTTPS POST and GET methods.
//
// > The authorization key is valid for 30 minutes and cannot be reused. It is recommended to re-obtain it before each activation.
//
// @param request - CreateAuthKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAuthKeyResponse
func (client *Client) CreateAuthKeyWithContext(ctx context.Context, request *CreateAuthKeyRequest, runtime *dara.RuntimeOptions) (_result *CreateAuthKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthYears) {
		query["AuthYears"] = request.AuthYears
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.Test) {
		query["Test"] = request.Test
	}

	if !dara.IsNil(request.UserDeviceId) {
		query["UserDeviceId"] = request.UserDeviceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAuthKey"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAuthKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Cloud Scene
//
// Description:
//
// Request Method: Supports sending requests via HTTPS POST and GET methods.
//
// > The authorization key is valid for 30 minutes and cannot be reused. It is recommended to reacquire it before each activation.
//
// @param request - CreateCloudauthstSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCloudauthstSceneResponse
func (client *Client) CreateCloudauthstSceneWithContext(ctx context.Context, request *CreateCloudauthstSceneRequest, runtime *dara.RuntimeOptions) (_result *CreateCloudauthstSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.SceneName) {
		query["SceneName"] = request.SceneName
	}

	if !dara.IsNil(request.StoreImage) {
		query["StoreImage"] = request.StoreImage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCloudauthstScene"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCloudauthstSceneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Scene Configuration
//
// Description:
//
// Request Method: Supports sending requests via HTTPS POST.
//
// Request Address: cloudauth.aliyuncs.com.
//
// > The authorization key is valid for 30 minutes and cannot be reused. It is recommended to reacquire it before each activation.
//
// @param request - CreateSceneConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSceneConfigResponse
func (client *Client) CreateSceneConfigWithContext(ctx context.Context, request *CreateSceneConfigRequest, runtime *dara.RuntimeOptions) (_result *CreateSceneConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["config"] = request.Config
	}

	if !dara.IsNil(request.SceneId) {
		body["sceneId"] = request.SceneId
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSceneConfig"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSceneConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call CreateVerifySetting to create a verification scenario configuration. This operation is equivalent to creating a new verification scenario on the console.
//
// Description:
//
// Request Method: Only supports sending requests via HTTPS POST.
//
// @param request - CreateVerifySettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVerifySettingResponse
func (client *Client) CreateVerifySettingWithContext(ctx context.Context, request *CreateVerifySettingRequest, runtime *dara.RuntimeOptions) (_result *CreateVerifySettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizName) {
		query["BizName"] = request.BizName
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.GuideStep) {
		query["GuideStep"] = request.GuideStep
	}

	if !dara.IsNil(request.PrivacyStep) {
		query["PrivacyStep"] = request.PrivacyStep
	}

	if !dara.IsNil(request.ResultStep) {
		query["ResultStep"] = request.ResultStep
	}

	if !dara.IsNil(request.Solution) {
		query["Solution"] = request.Solution
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVerifySetting"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVerifySettingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Whitelist
//
// Description:
//
// Request Method: Only supports sending requests via HTTPS POST.
//
// @param request - CreateWhitelistSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWhitelistSettingResponse
func (client *Client) CreateWhitelistSettingWithContext(ctx context.Context, request *CreateWhitelistSettingRequest, runtime *dara.RuntimeOptions) (_result *CreateWhitelistSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertNo) {
		query["CertNo"] = request.CertNo
	}

	if !dara.IsNil(request.CertifyId) {
		query["CertifyId"] = request.CertifyId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.ValidDay) {
		query["ValidDay"] = request.ValidDay
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWhitelistSetting"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWhitelistSettingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Product Credential Verification
//
// Description:
//
// Upload e-commerce product images to perform tampering, quality (clarity), and similar image detection, returning risk labels and scores.
//
// @param request - CredentialProductVerifyV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CredentialProductVerifyV2Response
func (client *Client) CredentialProductVerifyV2WithContext(ctx context.Context, request *CredentialProductVerifyV2Request, runtime *dara.RuntimeOptions) (_result *CredentialProductVerifyV2Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CredName) {
		query["CredName"] = request.CredName
	}

	if !dara.IsNil(request.CredType) {
		query["CredType"] = request.CredType
	}

	if !dara.IsNil(request.ImageUrl) {
		query["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.MerchantId) {
		query["MerchantId"] = request.MerchantId
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageFile) {
		body["ImageFile"] = request.ImageFile
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CredentialProductVerifyV2"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CredentialProductVerifyV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Credential verification
//
// @param tmpReq - CredentialVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CredentialVerifyResponse
func (client *Client) CredentialVerifyWithContext(ctx context.Context, tmpReq *CredentialVerifyRequest, runtime *dara.RuntimeOptions) (_result *CredentialVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CredentialVerifyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MerchantDetail) {
		request.MerchantDetailShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MerchantDetail, dara.String("MerchantDetail"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CertNum) {
		query["CertNum"] = request.CertNum
	}

	if !dara.IsNil(request.CredName) {
		query["CredName"] = request.CredName
	}

	if !dara.IsNil(request.CredType) {
		query["CredType"] = request.CredType
	}

	if !dara.IsNil(request.IdentifyNum) {
		query["IdentifyNum"] = request.IdentifyNum
	}

	if !dara.IsNil(request.ImageUrl) {
		query["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.IsCheck) {
		query["IsCheck"] = request.IsCheck
	}

	if !dara.IsNil(request.IsOCR) {
		query["IsOCR"] = request.IsOCR
	}

	if !dara.IsNil(request.MerchantDetailShrink) {
		query["MerchantDetail"] = request.MerchantDetailShrink
	}

	if !dara.IsNil(request.MerchantId) {
		query["MerchantId"] = request.MerchantId
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.Prompt) {
		query["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.PromptModel) {
		query["PromptModel"] = request.PromptModel
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageContext) {
		body["ImageContext"] = request.ImageContext
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CredentialVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CredentialVerifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Credential Verification
//
// Description:
//
// Input credential image information, perform image tampering and forgery detection, and image semantic understanding. Return the risk detection results.
//
// @param tmpReq - CredentialVerifyV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CredentialVerifyV2Response
func (client *Client) CredentialVerifyV2WithContext(ctx context.Context, tmpReq *CredentialVerifyV2Request, runtime *dara.RuntimeOptions) (_result *CredentialVerifyV2Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CredentialVerifyV2ShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MerchantDetail) {
		request.MerchantDetailShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MerchantDetail, dara.String("MerchantDetail"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CertNum) {
		query["CertNum"] = request.CertNum
	}

	if !dara.IsNil(request.CredName) {
		query["CredName"] = request.CredName
	}

	if !dara.IsNil(request.CredType) {
		query["CredType"] = request.CredType
	}

	if !dara.IsNil(request.IdentifyNum) {
		query["IdentifyNum"] = request.IdentifyNum
	}

	if !dara.IsNil(request.ImageUrl) {
		query["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.IsCheck) {
		query["IsCheck"] = request.IsCheck
	}

	if !dara.IsNil(request.IsOcr) {
		query["IsOcr"] = request.IsOcr
	}

	if !dara.IsNil(request.MerchantDetailShrink) {
		query["MerchantDetail"] = request.MerchantDetailShrink
	}

	if !dara.IsNil(request.MerchantId) {
		query["MerchantId"] = request.MerchantId
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.Prompt) {
		query["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.PromptModel) {
		query["PromptModel"] = request.PromptModel
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageContext) {
		body["ImageContext"] = request.ImageContext
	}

	if !dara.IsNil(request.ImageFile) {
		body["ImageFile"] = request.ImageFile
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CredentialVerifyV2"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CredentialVerifyV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Face Credential Verification Service
//
// Description:
//
// > The Face Deepfake Detection API is currently in the free public beta stage, which will end on August 30, 2024, at 23:59:59. During the public beta, the QPS (Queries Per Second) cannot exceed 3 times/second.
//
// - Service address: cloudauth.aliyuncs.com (IPv4) or cloudauth-dualstack.aliyuncs.com (IPv6).
//
// - Request method: POST and GET.
//
// - Transfer protocol: HTTPS.
//
// @param request - DeepfakeDetectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeepfakeDetectResponse
func (client *Client) DeepfakeDetectWithContext(ctx context.Context, request *DeepfakeDetectRequest, runtime *dara.RuntimeOptions) (_result *DeepfakeDetectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FaceInputType) {
		query["FaceInputType"] = request.FaceInputType
	}

	if !dara.IsNil(request.FaceUrl) {
		query["FaceUrl"] = request.FaceUrl
	}

	if !dara.IsNil(request.OuterOrderNo) {
		query["OuterOrderNo"] = request.OuterOrderNo
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FaceBase64) {
		body["FaceBase64"] = request.FaceBase64
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeepfakeDetect"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeepfakeDetectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete All Custom Flow Control Strategies
//
// Description:
//
// Request Method: Supports sending requests via HTTPS POST and GET methods.
//
// > The authorization key is valid for 30 minutes and cannot be reused. It is recommended to reacquire it before each activation.
//
// @param request - DeleteAllCustomizeFlowStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAllCustomizeFlowStrategyResponse
func (client *Client) DeleteAllCustomizeFlowStrategyWithContext(ctx context.Context, request *DeleteAllCustomizeFlowStrategyRequest, runtime *dara.RuntimeOptions) (_result *DeleteAllCustomizeFlowStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAllCustomizeFlowStrategy"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAllCustomizeFlowStrategyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Watermark Scene
//
// Description:
//
// - Service Address: cloudauth.aliyuncs.com.
//
// - Request Method: HTTPS POST and GET.
//
// @param request - DeleteAntCloudAuthSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAntCloudAuthSceneResponse
func (client *Client) DeleteAntCloudAuthSceneWithContext(ctx context.Context, request *DeleteAntCloudAuthSceneRequest, runtime *dara.RuntimeOptions) (_result *DeleteAntCloudAuthSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAntCloudAuthScene"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAntCloudAuthSceneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Black and White List Policy
//
// Description:
//
// Request Method: Only supports sending requests via HTTPS POST method.
//
// @param request - DeleteBlackListStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBlackListStrategyResponse
func (client *Client) DeleteBlackListStrategyWithContext(ctx context.Context, request *DeleteBlackListStrategyRequest, runtime *dara.RuntimeOptions) (_result *DeleteBlackListStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.ProductName) {
		query["ProductName"] = request.ProductName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBlackListStrategy"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBlackListStrategyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Cloud Scene
//
// Description:
//
// Request Method: Supports sending requests using HTTPS POST and GET methods.
//
// > The authorization key is valid for 30 minutes and cannot be reused. It is recommended to re-obtain it before each activation.
//
// @param request - DeleteCloudauthstSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCloudauthstSceneResponse
func (client *Client) DeleteCloudauthstSceneWithContext(ctx context.Context, request *DeleteCloudauthstSceneRequest, runtime *dara.RuntimeOptions) (_result *DeleteCloudauthstSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCloudauthstScene"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCloudauthstSceneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Security Control Strategy
//
// Description:
//
// Request Method: Supports sending requests via HTTPS POST.
//
// Request URL: cloudauth.aliyuncs.com.
//
// @param request - DeleteControlStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteControlStrategyResponse
func (client *Client) DeleteControlStrategyWithContext(ctx context.Context, request *DeleteControlStrategyRequest, runtime *dara.RuntimeOptions) (_result *DeleteControlStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiName) {
		query["ApiName"] = request.ApiName
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteControlStrategy"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteControlStrategyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Customized Flow Control Strategy
//
// Description:
//
// Request Method: Supports sending requests using HTTPS POST and GET methods.
//
// > The authorization key is valid for 30 minutes and cannot be reused. It is recommended to reacquire it before each activation.
//
// @param request - DeleteCustomizeFlowStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomizeFlowStrategyResponse
func (client *Client) DeleteCustomizeFlowStrategyWithContext(ctx context.Context, request *DeleteCustomizeFlowStrategyRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomizeFlowStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiName) {
		query["ApiName"] = request.ApiName
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomizeFlowStrategy"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomizeFlowStrategyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Financial Level Sensitive Data Deletion Interface
//
// Description:
//
// Deletes all personal information fields in the request, including name, ID number, phone number, IP, images, videos, and device information, etc.
//
// @param request - DeleteFaceVerifyResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFaceVerifyResultResponse
func (client *Client) DeleteFaceVerifyResultWithContext(ctx context.Context, request *DeleteFaceVerifyResultRequest, runtime *dara.RuntimeOptions) (_result *DeleteFaceVerifyResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertifyId) {
		query["CertifyId"] = request.CertifyId
	}

	if !dara.IsNil(request.DeleteAfterQuery) {
		query["DeleteAfterQuery"] = request.DeleteAfterQuery
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFaceVerifyResult"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFaceVerifyResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Scene Configuration
//
// Description:
//
// - Request Method: Supports sending requests via HTTPS POST and GET methods.
//
// - Request URL: cloudauth.aliyuncs.com.
//
// > The authorization key is valid for 30 minutes and cannot be reused. It is recommended to re-obtain it before each activation.
//
// @param request - DeleteSceneConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSceneConfigResponse
func (client *Client) DeleteSceneConfigWithContext(ctx context.Context, request *DeleteSceneConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteSceneConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SceneConfigId) {
		body["sceneConfigId"] = request.SceneConfigId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSceneConfig"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSceneConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Whitelist Configuration
//
// Description:
//
// Request Method: Only supports sending requests via HTTPS POST method.
//
// @param request - DeleteWhitelistSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWhitelistSettingResponse
func (client *Client) DeleteWhitelistSettingWithContext(ctx context.Context, request *DeleteWhitelistSettingRequest, runtime *dara.RuntimeOptions) (_result *DeleteWhitelistSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Ids) {
		query["Ids"] = request.Ids
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWhitelistSetting"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWhitelistSettingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取结果
//
// @param request - DescribeAuthVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAuthVerifyResponse
func (client *Client) DescribeAuthVerifyWithContext(ctx context.Context, request *DescribeAuthVerifyRequest, runtime *dara.RuntimeOptions) (_result *DescribeAuthVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CertifyId) {
		body["CertifyId"] = request.CertifyId
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAuthVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAuthVerifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Obtain Authentication Results from Image Element Verification
//
// Description:
//
// After receiving the callback notification, you can use this interface on the server side to obtain the corresponding authentication status and information.
//
// @param request - DescribeCardVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCardVerifyResponse
func (client *Client) DescribeCardVerifyWithContext(ctx context.Context, request *DescribeCardVerifyRequest, runtime *dara.RuntimeOptions) (_result *DescribeCardVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertifyId) {
		query["CertifyId"] = request.CertifyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCardVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCardVerifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Dashboard Data
//
// Description:
//
// Request Method: Supports sending requests via HTTPS POST and GET methods.
//
// > The authorization key is valid for 30 minutes and cannot be reused. It is recommended to reacquire it before each activation.
//
// @param request - DescribeCloudauthstSceneListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudauthstSceneListResponse
func (client *Client) DescribeCloudauthstSceneListWithContext(ctx context.Context, request *DescribeCloudauthstSceneListRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudauthstSceneListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudauthstSceneList"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudauthstSceneListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call DescribeDeviceInfo to query device-related information, such as the validity period of authorization, business identifiers customized by the access party, and device ID, etc.
//
// Description:
//
// Request Method: Supports sending requests using HTTPS POST and GET methods.
//
// @param request - DescribeDeviceInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeviceInfoResponse
func (client *Client) DescribeDeviceInfoWithContext(ctx context.Context, request *DescribeDeviceInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDeviceInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.ExpiredEndDay) {
		query["ExpiredEndDay"] = request.ExpiredEndDay
	}

	if !dara.IsNil(request.ExpiredStartDay) {
		query["ExpiredStartDay"] = request.ExpiredStartDay
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.UserDeviceId) {
		query["UserDeviceId"] = request.UserDeviceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDeviceInfo"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDeviceInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Financial-Grade Face Guard Service
//
// @param request - DescribeFaceGuardRiskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFaceGuardRiskResponse
func (client *Client) DescribeFaceGuardRiskWithContext(ctx context.Context, request *DescribeFaceGuardRiskRequest, runtime *dara.RuntimeOptions) (_result *DescribeFaceGuardRiskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.DeviceToken) {
		query["DeviceToken"] = request.DeviceToken
	}

	if !dara.IsNil(request.OuterOrderNo) {
		query["OuterOrderNo"] = request.OuterOrderNo
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFaceGuardRisk"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFaceGuardRiskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// After the mobile end of the integrator receives the callback, its server can call this interface to obtain the corresponding authentication status and authentication information.
//
// Description:
//
// - Service Address: cloudauth.aliyuncs.com.
//
// - Request Method: HTTPS POST and GET.
//
// @param request - DescribeFaceVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFaceVerifyResponse
func (client *Client) DescribeFaceVerifyWithContext(ctx context.Context, request *DescribeFaceVerifyRequest, runtime *dara.RuntimeOptions) (_result *DescribeFaceVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertifyId) {
		query["CertifyId"] = request.CertifyId
	}

	if !dara.IsNil(request.PictureReturnType) {
		query["PictureReturnType"] = request.PictureReturnType
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFaceVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFaceVerifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询任务导出记录
//
// @param request - DescribeInfoCheckExportRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInfoCheckExportRecordResponse
func (client *Client) DescribeInfoCheckExportRecordWithContext(ctx context.Context, request *DescribeInfoCheckExportRecordRequest, runtime *dara.RuntimeOptions) (_result *DescribeInfoCheckExportRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInfoCheckExportRecord"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInfoCheckExportRecordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the cloud scenario authentication records of a specific region
//
// Description:
//
// Request Method: Supports sending requests via HTTPS POST and GET methods.
//
// > The authorization key is valid for 30 minutes and cannot be reused. It is recommended to re-obtain it before each activation.
//
// @param request - DescribeListAntCloudAuthScenesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeListAntCloudAuthScenesResponse
func (client *Client) DescribeListAntCloudAuthScenesWithContext(ctx context.Context, request *DescribeListAntCloudAuthScenesRequest, runtime *dara.RuntimeOptions) (_result *DescribeListAntCloudAuthScenesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeListAntCloudAuthScenes"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeListAntCloudAuthScenesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Face Verification Data
//
// Description:
//
// - Service Address: cloudauth.aliyuncs.com.
//
// - Request Method: HTTPS POST and GET.
//
// @param request - DescribeListFaceVerifyDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeListFaceVerifyDataResponse
func (client *Client) DescribeListFaceVerifyDataWithContext(ctx context.Context, request *DescribeListFaceVerifyDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeListFaceVerifyDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GmtEnd) {
		query["GmtEnd"] = request.GmtEnd
	}

	if !dara.IsNil(request.GmtStart) {
		query["GmtStart"] = request.GmtStart
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeListFaceVerifyData"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeListFaceVerifyDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Face Verification Information
//
// Description:
//
// - Service address: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST and GET.
//
// @param request - DescribeListFaceVerifyInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeListFaceVerifyInfosResponse
func (client *Client) DescribeListFaceVerifyInfosWithContext(ctx context.Context, request *DescribeListFaceVerifyInfosRequest, runtime *dara.RuntimeOptions) (_result *DescribeListFaceVerifyInfosResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertifyId) {
		query["CertifyId"] = request.CertifyId
	}

	if !dara.IsNil(request.GmtEnd) {
		query["GmtEnd"] = request.GmtEnd
	}

	if !dara.IsNil(request.GmtStart) {
		query["GmtStart"] = request.GmtStart
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeListFaceVerifyInfos"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeListFaceVerifyInfosResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询页面元数据
//
// @param request - DescribeMetaSearchPageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetaSearchPageListResponse
func (client *Client) DescribeMetaSearchPageListWithContext(ctx context.Context, request *DescribeMetaSearchPageListRequest, runtime *dara.RuntimeOptions) (_result *DescribeMetaSearchPageListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Api) {
		query["Api"] = request.Api
	}

	if !dara.IsNil(request.BankCard) {
		query["BankCard"] = request.BankCard
	}

	if !dara.IsNil(request.BizCode) {
		query["BizCode"] = request.BizCode
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.IdentifyNum) {
		query["IdentifyNum"] = request.IdentifyNum
	}

	if !dara.IsNil(request.IspName) {
		query["IspName"] = request.IspName
	}

	if !dara.IsNil(request.Mobile) {
		query["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ReqId) {
		query["ReqId"] = request.ReqId
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.SubCode) {
		query["SubCode"] = request.SubCode
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	if !dara.IsNil(request.VehicleNum) {
		query["VehicleNum"] = request.VehicleNum
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMetaSearchPageList"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMetaSearchPageListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询认证统计信息
//
// @param request - DescribeMetaStatisticsListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetaStatisticsListResponse
func (client *Client) DescribeMetaStatisticsListWithContext(ctx context.Context, request *DescribeMetaStatisticsListRequest, runtime *dara.RuntimeOptions) (_result *DescribeMetaStatisticsListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Api) {
		query["Api"] = request.Api
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMetaStatisticsList"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMetaStatisticsListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询认证统计页面
//
// @param request - DescribeMetaStatisticsPageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetaStatisticsPageListResponse
func (client *Client) DescribeMetaStatisticsPageListWithContext(ctx context.Context, request *DescribeMetaStatisticsPageListRequest, runtime *dara.RuntimeOptions) (_result *DescribeMetaStatisticsPageListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Api) {
		query["Api"] = request.Api
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMetaStatisticsPageList"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMetaStatisticsPageListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query OSS status
//
// Description:
//
// - Request Method: Supports sending requests via HTTPS POST and GET methods.
//
// - Service Address: cloudauth.aliyuncs.com.
//
// @param request - DescribeOssStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOssStatusResponse
func (client *Client) DescribeOssStatusWithContext(ctx context.Context, request *DescribeOssStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeOssStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOssStatus"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOssStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get OSS Activation Status
//
// Description:
//
// - Request Method: Supports sending requests via HTTPS POST and GET methods.
//
// - Service Address: cloudauth.aliyuncs.com.
//
// @param request - DescribeOssStatusV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOssStatusV2Response
func (client *Client) DescribeOssStatusV2WithContext(ctx context.Context, request *DescribeOssStatusV2Request, runtime *dara.RuntimeOptions) (_result *DescribeOssStatusV2Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOssStatusV2"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOssStatusV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Open API adds financial-grade data statistics API
//
// @param request - DescribePageFaceVerifyDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePageFaceVerifyDataResponse
func (client *Client) DescribePageFaceVerifyDataWithContext(ctx context.Context, request *DescribePageFaceVerifyDataRequest, runtime *dara.RuntimeOptions) (_result *DescribePageFaceVerifyDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePageFaceVerifyData"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePageFaceVerifyDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Enhanced Real Person Authentication Call Statistics Pagination Query Interface
//
// @param request - DescribeSmartStatisticsPageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSmartStatisticsPageListResponse
func (client *Client) DescribeSmartStatisticsPageListWithContext(ctx context.Context, request *DescribeSmartStatisticsPageListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSmartStatisticsPageListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSmartStatisticsPageList"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSmartStatisticsPageListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Verification Device Statistics
//
// Description:
//
// - Service address: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST and GET.
//
// @param request - DescribeVerifyDeviceRiskStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVerifyDeviceRiskStatisticsResponse
func (client *Client) DescribeVerifyDeviceRiskStatisticsWithContext(ctx context.Context, request *DescribeVerifyDeviceRiskStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeVerifyDeviceRiskStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVerifyDeviceRiskStatistics"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVerifyDeviceRiskStatisticsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Overview of authentication request statistics
//
// Description:
//
// - Service address: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST and GET.
//
// @param request - DescribeVerifyFailStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVerifyFailStatisticsResponse
func (client *Client) DescribeVerifyFailStatisticsWithContext(ctx context.Context, request *DescribeVerifyFailStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeVerifyFailStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgeGt) {
		query["AgeGt"] = request.AgeGt
	}

	if !dara.IsNil(request.Api) {
		query["Api"] = request.Api
	}

	if !dara.IsNil(request.DeviceType) {
		query["DeviceType"] = request.DeviceType
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVerifyFailStatistics"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVerifyFailStatisticsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Statistics on Device Face Comparison
//
// Description:
//
// - Service Address: cloudauth.aliyuncs.com.
//
// - Request Method: HTTPS POST and GET.
//
// @param request - DescribeVerifyPersonasDeviceModelStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVerifyPersonasDeviceModelStatisticsResponse
func (client *Client) DescribeVerifyPersonasDeviceModelStatisticsWithContext(ctx context.Context, request *DescribeVerifyPersonasDeviceModelStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeVerifyPersonasDeviceModelStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.TimeRange) {
		query["TimeRange"] = request.TimeRange
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVerifyPersonasDeviceModelStatistics"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVerifyPersonasDeviceModelStatisticsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Authentication Personnel Statistics
//
// Description:
//
// - Service address: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST and GET.
//
// @param request - DescribeVerifyPersonasOsStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVerifyPersonasOsStatisticsResponse
func (client *Client) DescribeVerifyPersonasOsStatisticsWithContext(ctx context.Context, request *DescribeVerifyPersonasOsStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeVerifyPersonasOsStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.TimeRange) {
		query["TimeRange"] = request.TimeRange
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVerifyPersonasOsStatistics"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVerifyPersonasOsStatisticsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Obtain statistical information on the location of authenticated individuals
//
// Description:
//
// - Service Address: cloudauth.aliyuncs.com.
//
// - Request Method: HTTPS POST and GET.
//
// @param request - DescribeVerifyPersonasProvinceStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVerifyPersonasProvinceStatisticsResponse
func (client *Client) DescribeVerifyPersonasProvinceStatisticsWithContext(ctx context.Context, request *DescribeVerifyPersonasProvinceStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeVerifyPersonasProvinceStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.TimeRange) {
		query["TimeRange"] = request.TimeRange
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVerifyPersonasProvinceStatistics"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVerifyPersonasProvinceStatisticsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query gender statistics of authentication
//
// Description:
//
// - Service address: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST and GET.
//
// @param request - DescribeVerifyPersonasSexStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVerifyPersonasSexStatisticsResponse
func (client *Client) DescribeVerifyPersonasSexStatisticsWithContext(ctx context.Context, request *DescribeVerifyPersonasSexStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeVerifyPersonasSexStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.TimeRange) {
		query["TimeRange"] = request.TimeRange
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVerifyPersonasSexStatistics"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVerifyPersonasSexStatisticsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the result of real-person authentication.
//
// Description:
//
// Prerequisites: Before accessing this API, please ensure that you have completed the necessary preparations. For more details, see [Real Person Authentication Server-side Preparation](https://help.aliyun.com/document_detail/127471.html) and [Liveness Face Verification Server-side Preparation](https://help.aliyun.com/document_detail/127717.html).
//
// > Alibaba Cloud Real Person Authentication only stores authentication data for the last 180 days. For any subsequent business use, please call this interface in a timely manner to retrieve and store the data yourself to avoid any impact on usage.
//
// Request Method: HTTPS POST and GET.
//
// Interface Description: After the mobile end of the access party receives the callback, its server can call this interface to obtain the corresponding authentication status and authentication information.
//
// Applicable Scope: This interface is applicable to the authentication solution with SDK + server-side integration.
//
// @param request - DescribeVerifyResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVerifyResultResponse
func (client *Client) DescribeVerifyResultWithContext(ctx context.Context, request *DescribeVerifyResultRequest, runtime *dara.RuntimeOptions) (_result *DescribeVerifyResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVerifyResult"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVerifyResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call DescribeVerifySDK to get the offline SDK download address.
//
// Description:
//
// Request Method: Supports sending requests via HTTPS POST and GET methods.
//
// Interface Description: Obtain the SDK generation result based on the task ID for generating an offline facial recognition SDK.
//
// @param request - DescribeVerifySDKRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVerifySDKResponse
func (client *Client) DescribeVerifySDKWithContext(ctx context.Context, request *DescribeVerifySDKRequest, runtime *dara.RuntimeOptions) (_result *DescribeVerifySDKResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVerifySDK"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVerifySDKResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the list of authentication schemes
//
// Description:
//
// - Service Address: cloudauth.aliyuncs.com.
//
// - Request Method: HTTPS POST and GET.
//
// @param request - DescribeVerifySearchPageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVerifySearchPageListResponse
func (client *Client) DescribeVerifySearchPageListWithContext(ctx context.Context, request *DescribeVerifySearchPageListRequest, runtime *dara.RuntimeOptions) (_result *DescribeVerifySearchPageListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertNo) {
		query["CertNo"] = request.CertNo
	}

	if !dara.IsNil(request.CertifyId) {
		query["CertifyId"] = request.CertifyId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.HasDeviceRisk) {
		query["HasDeviceRisk"] = request.HasDeviceRisk
	}

	if !dara.IsNil(request.Model) {
		query["Model"] = request.Model
	}

	if !dara.IsNil(request.OuterOrderNo) {
		query["OuterOrderNo"] = request.OuterOrderNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Passed) {
		query["Passed"] = request.Passed
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.Root) {
		query["Root"] = request.Root
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.Simulator) {
		query["Simulator"] = request.Simulator
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.SubCode) {
		query["SubCode"] = request.SubCode
	}

	if !dara.IsNil(request.SubCodes) {
		query["SubCodes"] = request.SubCodes
	}

	if !dara.IsNil(request.VirtualVideo) {
		query["VirtualVideo"] = request.VirtualVideo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVerifySearchPageList"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVerifySearchPageListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Authentication Statistics
//
// Description:
//
// - Request Method: Supports sending requests using HTTPS POST and GET methods.
//
// - Service Address: cloudauth.aliyuncs.com.
//
// @param request - DescribeVerifyStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVerifyStatisticsResponse
func (client *Client) DescribeVerifyStatisticsWithContext(ctx context.Context, request *DescribeVerifyStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeVerifyStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgeGt) {
		query["AgeGt"] = request.AgeGt
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVerifyStatistics"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVerifyStatisticsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call DescribeVerifyToken to initiate an authentication request and obtain an authentication Token. This interface is suitable for authentication solutions using SDK + server-side integration.
//
// Description:
//
// Preparation for Access: When integrating this API, please ensure that the corresponding preparations have been completed. For details, see [Overview of Real Person Authentication Solution Integration Process](https://help.aliyun.com/document_detail/127536.html) and [Overview of Live Face Verification Solution (Liveness Detection Solution) Integration Process](https://help.aliyun.com/document_detail/127687.html).
//
// Request Method: HTTPS POST and GET
//
// API Description: Before each authentication, use this interface to obtain an authentication Token (VerifyToken), which is used to link various interfaces in the authentication request.
//
// Applicable Scope: This interface is suitable for wireless SDK integration.
//
// Image Address: Use HTTP or HTTPS addresses that are publicly accessible over the Internet. For example, `http://image-demo.img-cn-hangzhou.aliyuncs.com/example.jpg`.
//
// Image Restrictions:
//
// - Relative or absolute paths of local images are not supported.
//
// - The size of a single image should be controlled within 2 MB to avoid algorithm retrieval timeout.
//
// - The face area in the image must be at least 64*64 pixels (px).
//
// - There is an 8 MB size limit for the Body of a single request. Please calculate the total size of all images and other information in the request to ensure it does not exceed the limit.
//
// @param request - DescribeVerifyTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVerifyTokenResponse
func (client *Client) DescribeVerifyTokenWithContext(ctx context.Context, request *DescribeVerifyTokenRequest, runtime *dara.RuntimeOptions) (_result *DescribeVerifyTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.CallbackSeed) {
		query["CallbackSeed"] = request.CallbackSeed
	}

	if !dara.IsNil(request.CallbackUrl) {
		query["CallbackUrl"] = request.CallbackUrl
	}

	if !dara.IsNil(request.FaceRetainedImageUrl) {
		query["FaceRetainedImageUrl"] = request.FaceRetainedImageUrl
	}

	if !dara.IsNil(request.FailedRedirectUrl) {
		query["FailedRedirectUrl"] = request.FailedRedirectUrl
	}

	if !dara.IsNil(request.IdCardBackImageUrl) {
		query["IdCardBackImageUrl"] = request.IdCardBackImageUrl
	}

	if !dara.IsNil(request.IdCardFrontImageUrl) {
		query["IdCardFrontImageUrl"] = request.IdCardFrontImageUrl
	}

	if !dara.IsNil(request.IdCardNumber) {
		query["IdCardNumber"] = request.IdCardNumber
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PassedRedirectUrl) {
		query["PassedRedirectUrl"] = request.PassedRedirectUrl
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.UserIp) {
		query["UserIp"] = request.UserIp
	}

	if !dara.IsNil(request.UserPhoneNumber) {
		query["UserPhoneNumber"] = request.UserPhoneNumber
	}

	if !dara.IsNil(request.UserRegistTime) {
		query["UserRegistTime"] = request.UserRegistTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVerifyToken"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVerifyTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Whitelist Collection Get Whitelist Collection
//
// Description:
//
// Request Method: Only supports sending requests via HTTPS POST method.
//
// @param request - DescribeWhitelistSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWhitelistSettingResponse
func (client *Client) DescribeWhitelistSettingWithContext(ctx context.Context, request *DescribeWhitelistSettingRequest, runtime *dara.RuntimeOptions) (_result *DescribeWhitelistSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertNo) {
		query["CertNo"] = request.CertNo
	}

	if !dara.IsNil(request.CertifyId) {
		query["CertifyId"] = request.CertifyId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.ValidEndDate) {
		query["ValidEndDate"] = request.ValidEndDate
	}

	if !dara.IsNil(request.ValidStartDate) {
		query["ValidStartDate"] = request.ValidStartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWhitelistSetting"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWhitelistSettingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Detect Validity Attributes in Face Photos
//
// Description:
//
// Request Method: Only supports sending requests via HTTPS POST.
//
// Interface Description: Detects the validity-related attributes of faces in the input photo, which helps the business side to determine whether the photo meets their own business retention or comparison requirements. The currently supported face validity-related attributes include: whether it is a face, whether it is blurry, whether glasses are worn, face pose, whether it is a smile, etc.
//
// Instructions for Uploading Image Addresses: When passing in images, you need to upload their corresponding HTTP, OSS addresses, or Base64 encoding.
//
// - HTTP Address: A publicly accessible HTTP address. For example, `http://image-demo.img-cn-hangzhou.aliyuncs.com/example.jpg`.
//
// - Base64 Encoding: An image encoded through base64, with the format being `base64://<image base64 string>`.
//
// Image Limitations:
//
// - Does not support relative or absolute paths of local images.
//
// - Please keep the size of a single image within 2 MB to avoid algorithm timeout.
//
// - There is an 8 MB size limit for the Body of a single request; please calculate the total size of all images and other information in the request and do not exceed the limit.
//
// - When using Base64 to pass images, the request method needs to be changed to POST; the header description of the image Base64 string, such as `data:image/png;base64`, should be removed.
//
// @param request - DetectFaceAttributesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectFaceAttributesResponse
func (client *Client) DetectFaceAttributesWithContext(ctx context.Context, request *DetectFaceAttributesRequest, runtime *dara.RuntimeOptions) (_result *DetectFaceAttributesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		body["BizType"] = request.BizType
	}

	if !dara.IsNil(request.MaterialValue) {
		body["MaterialValue"] = request.MaterialValue
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetectFaceAttributes"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetectFaceAttributesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Real-person Authentication Record Download
//
// Description:
//
// Obtain the download link for statistical call data files under the product plan based on query conditions.
//
// - Method: HTTPS POST
//
// - Service Address: cloudauth.aliyuncs.com
//
// > Real-person authentication products use CertifyId to count call volumes. For ease of reconciliation, please retain the CertifyId field in your system.
//
// @param request - DownloadVerifyRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DownloadVerifyRecordsResponse
func (client *Client) DownloadVerifyRecordsWithContext(ctx context.Context, request *DownloadVerifyRecordsRequest, runtime *dara.RuntimeOptions) (_result *DownloadVerifyRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizParam) {
		query["BizParam"] = request.BizParam
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DownloadVerifyRecords"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DownloadVerifyRecordsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Two-Factor Validity Verification API
//
// @param request - Id2MetaPeriodVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Id2MetaPeriodVerifyResponse
func (client *Client) Id2MetaPeriodVerifyWithContext(ctx context.Context, request *Id2MetaPeriodVerifyRequest, runtime *dara.RuntimeOptions) (_result *Id2MetaPeriodVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IdentifyNum) {
		body["IdentifyNum"] = request.IdentifyNum
	}

	if !dara.IsNil(request.ParamType) {
		body["ParamType"] = request.ParamType
	}

	if !dara.IsNil(request.UserName) {
		body["UserName"] = request.UserName
	}

	if !dara.IsNil(request.ValidityEndDate) {
		body["ValidityEndDate"] = request.ValidityEndDate
	}

	if !dara.IsNil(request.ValidityStartDate) {
		body["ValidityStartDate"] = request.ValidityStartDate
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Id2MetaPeriodVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &Id2MetaPeriodVerifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Identity Two-Factor Standard
//
// @param request - Id2MetaStandardVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Id2MetaStandardVerifyResponse
func (client *Client) Id2MetaStandardVerifyWithContext(ctx context.Context, request *Id2MetaStandardVerifyRequest, runtime *dara.RuntimeOptions) (_result *Id2MetaStandardVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IdentifyNum) {
		body["IdentifyNum"] = request.IdentifyNum
	}

	if !dara.IsNil(request.ParamType) {
		body["ParamType"] = request.ParamType
	}

	if !dara.IsNil(request.UserName) {
		body["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Id2MetaStandardVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &Id2MetaStandardVerifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Identity Two-Factor Interface
//
// Description:
//
// - Service address: cloudauth.aliyuncs.com (IPv4) or cloudauth-dualstack.aliyuncs.com (IPv6).
//
// - Request method: POST and GET.
//
// - Transfer protocol: HTTPS.
//
// @param request - Id2MetaVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Id2MetaVerifyResponse
func (client *Client) Id2MetaVerifyWithContext(ctx context.Context, request *Id2MetaVerifyRequest, runtime *dara.RuntimeOptions) (_result *Id2MetaVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IdentifyNum) {
		body["IdentifyNum"] = request.IdentifyNum
	}

	if !dara.IsNil(request.ParamType) {
		body["ParamType"] = request.ParamType
	}

	if !dara.IsNil(request.UserName) {
		body["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Id2MetaVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &Id2MetaVerifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ID Two-Factor Image Verification
//
// Description:
//
// Upload both sides of the ID card, and get the verification result of the two factors from an authoritative data source.
//
// @param request - Id2MetaVerifyWithOCRRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Id2MetaVerifyWithOCRResponse
func (client *Client) Id2MetaVerifyWithOCRWithContext(ctx context.Context, request *Id2MetaVerifyWithOCRRequest, runtime *dara.RuntimeOptions) (_result *Id2MetaVerifyWithOCRResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CertFile) {
		body["CertFile"] = request.CertFile
	}

	if !dara.IsNil(request.CertNationalFile) {
		body["CertNationalFile"] = request.CertNationalFile
	}

	if !dara.IsNil(request.CertNationalUrl) {
		body["CertNationalUrl"] = request.CertNationalUrl
	}

	if !dara.IsNil(request.CertUrl) {
		body["CertUrl"] = request.CertUrl
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Id2MetaVerifyWithOCR"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &Id2MetaVerifyWithOCRResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Identity Three Elements Verification
//
// Description:
//
// Input name, ID number, and face photo to verify their authenticity and consistency through authoritative sources.
//
// @param request - Id3MetaVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Id3MetaVerifyResponse
func (client *Client) Id3MetaVerifyWithContext(ctx context.Context, request *Id3MetaVerifyRequest, runtime *dara.RuntimeOptions) (_result *Id3MetaVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Crop) {
		body["Crop"] = request.Crop
	}

	if !dara.IsNil(request.FaceFile) {
		body["FaceFile"] = request.FaceFile
	}

	if !dara.IsNil(request.FaceUrl) {
		body["FaceUrl"] = request.FaceUrl
	}

	if !dara.IsNil(request.IdentifyNum) {
		body["IdentifyNum"] = request.IdentifyNum
	}

	if !dara.IsNil(request.ParamType) {
		body["ParamType"] = request.ParamType
	}

	if !dara.IsNil(request.UserName) {
		body["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Id3MetaVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &Id3MetaVerifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Identity Three Elements Image Verification
//
// Description:
//
// Upload both sides of the ID card to get the verification result of the three identity elements from an authoritative data source.
//
// @param request - Id3MetaVerifyWithOCRRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Id3MetaVerifyWithOCRResponse
func (client *Client) Id3MetaVerifyWithOCRWithContext(ctx context.Context, request *Id3MetaVerifyWithOCRRequest, runtime *dara.RuntimeOptions) (_result *Id3MetaVerifyWithOCRResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CertFile) {
		body["CertFile"] = request.CertFile
	}

	if !dara.IsNil(request.CertNationalFile) {
		body["CertNationalFile"] = request.CertNationalFile
	}

	if !dara.IsNil(request.CertNationalUrl) {
		body["CertNationalUrl"] = request.CertNationalUrl
	}

	if !dara.IsNil(request.CertUrl) {
		body["CertUrl"] = request.CertUrl
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Id3MetaVerifyWithOCR"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &Id3MetaVerifyWithOCRResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务端初始化
//
// @param request - InitAuthVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitAuthVerifyResponse
func (client *Client) InitAuthVerifyWithContext(ctx context.Context, request *InitAuthVerifyRequest, runtime *dara.RuntimeOptions) (_result *InitAuthVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CallbackToken) {
		body["CallbackToken"] = request.CallbackToken
	}

	if !dara.IsNil(request.CallbackUrl) {
		body["CallbackUrl"] = request.CallbackUrl
	}

	if !dara.IsNil(request.CardPageNumber) {
		body["CardPageNumber"] = request.CardPageNumber
	}

	if !dara.IsNil(request.CardType) {
		body["CardType"] = request.CardType
	}

	if !dara.IsNil(request.DocScanMode) {
		body["DocScanMode"] = request.DocScanMode
	}

	if !dara.IsNil(request.IdSpoof) {
		body["IdSpoof"] = request.IdSpoof
	}

	if !dara.IsNil(request.MetaInfo) {
		body["MetaInfo"] = request.MetaInfo
	}

	if !dara.IsNil(request.OuterOrderNo) {
		body["OuterOrderNo"] = request.OuterOrderNo
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InitAuthVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InitAuthVerifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Initiate an authentication request for image verification
//
// Description:
//
// Before each authentication, use this interface to obtain the CertifyId, which is used to link various interfaces in the authentication request.
//
// @param request - InitCardVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitCardVerifyResponse
func (client *Client) InitCardVerifyWithContext(ctx context.Context, request *InitCardVerifyRequest, runtime *dara.RuntimeOptions) (_result *InitCardVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallbackToken) {
		query["CallbackToken"] = request.CallbackToken
	}

	if !dara.IsNil(request.CallbackUrl) {
		query["CallbackUrl"] = request.CallbackUrl
	}

	if !dara.IsNil(request.CardPageNumber) {
		query["CardPageNumber"] = request.CardPageNumber
	}

	if !dara.IsNil(request.CardType) {
		query["CardType"] = request.CardType
	}

	if !dara.IsNil(request.DocScanMode) {
		query["DocScanMode"] = request.DocScanMode
	}

	if !dara.IsNil(request.MerchantBizId) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !dara.IsNil(request.MetaInfo) {
		query["MetaInfo"] = request.MetaInfo
	}

	if !dara.IsNil(request.Model) {
		query["Model"] = request.Model
	}

	if !dara.IsNil(request.PictureSave) {
		query["PictureSave"] = request.PictureSave
	}

	if !dara.IsNil(request.VerifyMeta) {
		query["VerifyMeta"] = request.VerifyMeta
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InitCardVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InitCardVerifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Real-Person Server Initialization Interface
//
// Description:
//
// - Service Address: cloudauth.aliyuncs.com
//
// - Request Method: HTTPS POST and GET.
//
// - This interface uses different parameters for different product solutions. For details, please refer to the [official documentation](https://help.aliyun.com/zh/id-verification/financial-grade-id-verification/product-overview/introduction/?spm=a2c4g.11186623.help-menu-2401581.d_0_0.13f644ecRzFHfm&scm=20140722.H_99169._.OR_help-T_cn~zh-V_1).
//
// #### Image Format Requirements
//
// When performing real-person authentication, please provide images that meet all of the following conditions:
//
// - Recent photo with a clear, unobstructed face, natural expression, and facing the camera directly.
//
// - Clear photo with normal exposure, no overexposure, underexposure, or halo effects, and no significant angle deviation.
//
// - Resolution not exceeding 1920*1080, at least 640*480, recommended short side scaled to 720 pixels, compression ratio greater than 0.9.
//
// - Photo size: <1MB.
//
// - Supports 90, 180, and 270-degree photos; in the case of multiple faces, the largest face will be selected.
//
// @param request - InitFaceVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitFaceVerifyResponse
func (client *Client) InitFaceVerifyWithContext(ctx context.Context, request *InitFaceVerifyRequest, runtime *dara.RuntimeOptions) (_result *InitFaceVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppQualityCheck) {
		query["AppQualityCheck"] = request.AppQualityCheck
	}

	if !dara.IsNil(request.Birthday) {
		query["Birthday"] = request.Birthday
	}

	if !dara.IsNil(request.CallbackToken) {
		query["CallbackToken"] = request.CallbackToken
	}

	if !dara.IsNil(request.CallbackUrl) {
		query["CallbackUrl"] = request.CallbackUrl
	}

	if !dara.IsNil(request.CameraSelection) {
		query["CameraSelection"] = request.CameraSelection
	}

	if !dara.IsNil(request.CertName) {
		query["CertName"] = request.CertName
	}

	if !dara.IsNil(request.CertNo) {
		query["CertNo"] = request.CertNo
	}

	if !dara.IsNil(request.CertType) {
		query["CertType"] = request.CertType
	}

	if !dara.IsNil(request.CertifyId) {
		query["CertifyId"] = request.CertifyId
	}

	if !dara.IsNil(request.CertifyUrlStyle) {
		query["CertifyUrlStyle"] = request.CertifyUrlStyle
	}

	if !dara.IsNil(request.CertifyUrlType) {
		query["CertifyUrlType"] = request.CertifyUrlType
	}

	if !dara.IsNil(request.EncryptType) {
		query["EncryptType"] = request.EncryptType
	}

	if !dara.IsNil(request.FaceContrastPictureUrl) {
		query["FaceContrastPictureUrl"] = request.FaceContrastPictureUrl
	}

	if !dara.IsNil(request.FaceGuardOutput) {
		query["FaceGuardOutput"] = request.FaceGuardOutput
	}

	if !dara.IsNil(request.H5DegradeConfirmBtn) {
		query["H5DegradeConfirmBtn"] = request.H5DegradeConfirmBtn
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.MetaInfo) {
		query["MetaInfo"] = request.MetaInfo
	}

	if !dara.IsNil(request.Mobile) {
		query["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	if !dara.IsNil(request.NeedMultiFaceCheck) {
		query["NeedMultiFaceCheck"] = request.NeedMultiFaceCheck
	}

	if !dara.IsNil(request.OssBucketName) {
		query["OssBucketName"] = request.OssBucketName
	}

	if !dara.IsNil(request.OssObjectName) {
		query["OssObjectName"] = request.OssObjectName
	}

	if !dara.IsNil(request.OuterOrderNo) {
		query["OuterOrderNo"] = request.OuterOrderNo
	}

	if !dara.IsNil(request.ProcedurePriority) {
		query["ProcedurePriority"] = request.ProcedurePriority
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.RarelyCharacters) {
		query["RarelyCharacters"] = request.RarelyCharacters
	}

	if !dara.IsNil(request.ReadImg) {
		query["ReadImg"] = request.ReadImg
	}

	if !dara.IsNil(request.ReturnUrl) {
		query["ReturnUrl"] = request.ReturnUrl
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.SuitableType) {
		query["SuitableType"] = request.SuitableType
	}

	if !dara.IsNil(request.UiCustomUrl) {
		query["UiCustomUrl"] = request.UiCustomUrl
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.ValidityDate) {
		query["ValidityDate"] = request.ValidityDate
	}

	if !dara.IsNil(request.VideoEvidence) {
		query["VideoEvidence"] = request.VideoEvidence
	}

	if !dara.IsNil(request.VoluntaryCustomizedContent) {
		query["VoluntaryCustomizedContent"] = request.VoluntaryCustomizedContent
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AuthId) {
		body["AuthId"] = request.AuthId
	}

	if !dara.IsNil(request.Crop) {
		body["Crop"] = request.Crop
	}

	if !dara.IsNil(request.FaceContrastPicture) {
		body["FaceContrastPicture"] = request.FaceContrastPicture
	}

	if !dara.IsNil(request.Model) {
		body["Model"] = request.Model
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InitFaceVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InitFaceVerifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Add Real Person Whitelist
//
// @param request - InsertWhiteListSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsertWhiteListSettingResponse
func (client *Client) InsertWhiteListSettingWithContext(ctx context.Context, request *InsertWhiteListSettingRequest, runtime *dara.RuntimeOptions) (_result *InsertWhiteListSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertNo) {
		query["CertNo"] = request.CertNo
	}

	if !dara.IsNil(request.CertifyId) {
		query["CertifyId"] = request.CertifyId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.ValidDay) {
		query["ValidDay"] = request.ValidDay
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InsertWhiteListSetting"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InsertWhiteListSettingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Silent Liveness Face (LivenessFaceVerify) refers to a service that performs real face detection by inputting pre-obtained face images through an API. The algorithm primarily identifies whether the face is from a screen recording, printed picture, or other types of liveness attacks. This service is suitable for low-risk business scenarios or in conjunction with offline facial recognition SDKs. If your business has higher requirements for real face security, it is recommended to integrate App or WebSDK mode, or integrate with Deepfake face detection services to assist in identifying more dimensions of face forgery risks.
//
// Description:
//
// Invoke the LivenessFaceVerify interface to perform liveness detection on a face image.
//
// @param request - LivenessFaceVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LivenessFaceVerifyResponse
func (client *Client) LivenessFaceVerifyWithContext(ctx context.Context, request *LivenessFaceVerifyRequest, runtime *dara.RuntimeOptions) (_result *LivenessFaceVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Model) {
		query["Model"] = request.Model
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CertifyId) {
		body["CertifyId"] = request.CertifyId
	}

	if !dara.IsNil(request.Crop) {
		body["Crop"] = request.Crop
	}

	if !dara.IsNil(request.DeviceToken) {
		body["DeviceToken"] = request.DeviceToken
	}

	if !dara.IsNil(request.FaceContrastPicture) {
		body["FaceContrastPicture"] = request.FaceContrastPicture
	}

	if !dara.IsNil(request.FaceContrastPictureUrl) {
		body["FaceContrastPictureUrl"] = request.FaceContrastPictureUrl
	}

	if !dara.IsNil(request.Ip) {
		body["Ip"] = request.Ip
	}

	if !dara.IsNil(request.Mobile) {
		body["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.OssBucketName) {
		body["OssBucketName"] = request.OssBucketName
	}

	if !dara.IsNil(request.OssObjectName) {
		body["OssObjectName"] = request.OssObjectName
	}

	if !dara.IsNil(request.OuterOrderNo) {
		body["OuterOrderNo"] = request.OuterOrderNo
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LivenessFaceVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LivenessFaceVerifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Mobile Two-Factor Verification
//
// Description:
//
// Input the phone number and name, verify their authenticity and consistency through authoritative data sources.
//
// @param request - Mobile2MetaVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Mobile2MetaVerifyResponse
func (client *Client) Mobile2MetaVerifyWithContext(ctx context.Context, request *Mobile2MetaVerifyRequest, runtime *dara.RuntimeOptions) (_result *Mobile2MetaVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Mobile) {
		body["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.ParamType) {
		body["ParamType"] = request.ParamType
	}

	if !dara.IsNil(request.UserName) {
		body["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Mobile2MetaVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &Mobile2MetaVerifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detailed Three-Element Verification for Phone Number_Standard Version
//
// Description:
//
// Input the phone number, name, and ID number to verify their authenticity and consistency through authoritative data sources. If they do not match, the reason for the mismatch is returned.
//
// @param request - Mobile3MetaDetailStandardVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Mobile3MetaDetailStandardVerifyResponse
func (client *Client) Mobile3MetaDetailStandardVerifyWithContext(ctx context.Context, request *Mobile3MetaDetailStandardVerifyRequest, runtime *dara.RuntimeOptions) (_result *Mobile3MetaDetailStandardVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IdentifyNum) {
		body["IdentifyNum"] = request.IdentifyNum
	}

	if !dara.IsNil(request.Mobile) {
		body["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.ParamType) {
		body["ParamType"] = request.ParamType
	}

	if !dara.IsNil(request.UserName) {
		body["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Mobile3MetaDetailStandardVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &Mobile3MetaDetailStandardVerifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Detailed Mobile Three-Element Verification Interface
//
// Description:
//
// - Service address: cloudauth.aliyuncs.com (IPv4) or cloudauth-dualstack.aliyuncs.com (IPv6).
//
// - Request method: POST and GET.
//
// - Transfer protocol: HTTPS.
//
// @param request - Mobile3MetaDetailVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Mobile3MetaDetailVerifyResponse
func (client *Client) Mobile3MetaDetailVerifyWithContext(ctx context.Context, request *Mobile3MetaDetailVerifyRequest, runtime *dara.RuntimeOptions) (_result *Mobile3MetaDetailVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IdentifyNum) {
		body["IdentifyNum"] = request.IdentifyNum
	}

	if !dara.IsNil(request.Mobile) {
		body["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.ParamType) {
		body["ParamType"] = request.ParamType
	}

	if !dara.IsNil(request.UserName) {
		body["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Mobile3MetaDetailVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &Mobile3MetaDetailVerifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Simplified Three-Element Verification for Phone Number_Standard Version
//
// Description:
//
// Input the phone number, name, and ID number to verify their authenticity and consistency through authoritative data sources.
//
// @param request - Mobile3MetaSimpleStandardVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Mobile3MetaSimpleStandardVerifyResponse
func (client *Client) Mobile3MetaSimpleStandardVerifyWithContext(ctx context.Context, request *Mobile3MetaSimpleStandardVerifyRequest, runtime *dara.RuntimeOptions) (_result *Mobile3MetaSimpleStandardVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IdentifyNum) {
		body["IdentifyNum"] = request.IdentifyNum
	}

	if !dara.IsNil(request.Mobile) {
		body["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.ParamType) {
		body["ParamType"] = request.ParamType
	}

	if !dara.IsNil(request.UserName) {
		body["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Mobile3MetaSimpleStandardVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &Mobile3MetaSimpleStandardVerifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Mobile Three Elements Simplified Interface
//
// Description:
//
// - Service address: cloudauth.aliyuncs.com (IPv4) or cloudauth-dualstack.aliyuncs.com (IPv6).
//
// - Request method: POST and GET.
//
// - Transfer protocol: HTTPS.
//
// @param request - Mobile3MetaSimpleVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Mobile3MetaSimpleVerifyResponse
func (client *Client) Mobile3MetaSimpleVerifyWithContext(ctx context.Context, request *Mobile3MetaSimpleVerifyRequest, runtime *dara.RuntimeOptions) (_result *Mobile3MetaSimpleVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IdentifyNum) {
		body["IdentifyNum"] = request.IdentifyNum
	}

	if !dara.IsNil(request.Mobile) {
		body["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.ParamType) {
		body["ParamType"] = request.ParamType
	}

	if !dara.IsNil(request.UserName) {
		body["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Mobile3MetaSimpleVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &Mobile3MetaSimpleVerifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Number Detection
//
// @param request - MobileDetectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MobileDetectResponse
func (client *Client) MobileDetectWithContext(ctx context.Context, request *MobileDetectRequest, runtime *dara.RuntimeOptions) (_result *MobileDetectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Mobiles) {
		body["Mobiles"] = request.Mobiles
	}

	if !dara.IsNil(request.ParamType) {
		body["ParamType"] = request.ParamType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MobileDetect"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MobileDetectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the online status of a mobile number
//
// @param request - MobileOnlineStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MobileOnlineStatusResponse
func (client *Client) MobileOnlineStatusWithContext(ctx context.Context, request *MobileOnlineStatusRequest, runtime *dara.RuntimeOptions) (_result *MobileOnlineStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Mobile) {
		body["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.ParamType) {
		body["ParamType"] = request.ParamType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MobileOnlineStatus"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MobileOnlineStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the online duration of a mobile number
//
// @param request - MobileOnlineTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MobileOnlineTimeResponse
func (client *Client) MobileOnlineTimeWithContext(ctx context.Context, request *MobileOnlineTimeRequest, runtime *dara.RuntimeOptions) (_result *MobileOnlineTimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Mobile) {
		body["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.ParamType) {
		body["ParamType"] = request.ParamType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MobileOnlineTime"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MobileOnlineTimeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify Black and White List Policy
//
// Description:
//
// - Service Address: cloudauth.aliyuncs.com.
//
// - Request Method: HTTPS POST and GET.
//
// - Interface Description: Add or modify blacklist rule.
//
// @param tmpReq - ModifyBlackListStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBlackListStrategyResponse
func (client *Client) ModifyBlackListStrategyWithContext(ctx context.Context, tmpReq *ModifyBlackListStrategyRequest, runtime *dara.RuntimeOptions) (_result *ModifyBlackListStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyBlackListStrategyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BlackListStrategy) {
		request.BlackListStrategyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BlackListStrategy, dara.String("BlackListStrategy"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BlackListStrategyShrink) {
		query["BlackListStrategy"] = request.BlackListStrategyShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBlackListStrategy"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyBlackListStrategyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify Security Control Strategy
//
// Description:
//
// - Request Method: Supports sending requests via HTTPS POST method.
//
// - Request Address: cloudauth.aliyuncs.com.
//
// @param tmpReq - ModifyControlStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyControlStrategyResponse
func (client *Client) ModifyControlStrategyWithContext(ctx context.Context, tmpReq *ModifyControlStrategyRequest, runtime *dara.RuntimeOptions) (_result *ModifyControlStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyControlStrategyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ControlStrategyList) {
		request.ControlStrategyListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ControlStrategyList, dara.String("ControlStrategyList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ControlStrategyListShrink) {
		query["ControlStrategyList"] = request.ControlStrategyListShrink
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyControlStrategy"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyControlStrategyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Information Verification Security Management
//
// Description:
//
// - Request Method: Supports sending requests via HTTPS POST and GET methods.
//
// - Service Address: cloudauth.aliyuncs.com.
//
// @param tmpReq - ModifyCustomizeFlowStrategyListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCustomizeFlowStrategyListResponse
func (client *Client) ModifyCustomizeFlowStrategyListWithContext(ctx context.Context, tmpReq *ModifyCustomizeFlowStrategyListRequest, runtime *dara.RuntimeOptions) (_result *ModifyCustomizeFlowStrategyListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyCustomizeFlowStrategyListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.StrategyObject) {
		request.StrategyObjectShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StrategyObject, dara.String("StrategyObject"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.StrategyObjectShrink) {
		query["StrategyObject"] = request.StrategyObjectShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCustomizeFlowStrategyList"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCustomizeFlowStrategyListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call ModifyDeviceInfo to update device-related information, such as extending the device authorization validity period, updating the business party\\"s custom business identifier, and device ID. Suitable for scenarios like renewing the device validity period.
//
// Description:
//
// Request Method: Supports sending requests using HTTPS POST and GET methods.
//
// @param request - ModifyDeviceInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDeviceInfoResponse
func (client *Client) ModifyDeviceInfoWithContext(ctx context.Context, request *ModifyDeviceInfoRequest, runtime *dara.RuntimeOptions) (_result *ModifyDeviceInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.ExpiredDay) {
		query["ExpiredDay"] = request.ExpiredDay
	}

	if !dara.IsNil(request.UserDeviceId) {
		query["UserDeviceId"] = request.UserDeviceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDeviceInfo"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDeviceInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Paging Query for Real Person Whitelist Configuration
//
// @param request - PageQueryWhiteListSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PageQueryWhiteListSettingResponse
func (client *Client) PageQueryWhiteListSettingWithContext(ctx context.Context, request *PageQueryWhiteListSettingRequest, runtime *dara.RuntimeOptions) (_result *PageQueryWhiteListSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertNo) {
		query["CertNo"] = request.CertNo
	}

	if !dara.IsNil(request.CertifyId) {
		query["CertifyId"] = request.CertifyId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.ValidEndDate) {
		query["ValidEndDate"] = request.ValidEndDate
	}

	if !dara.IsNil(request.ValidStartDate) {
		query["ValidStartDate"] = request.ValidStartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PageQueryWhiteListSetting"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PageQueryWhiteListSettingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Blacklist and Whitelist Policies
//
// Description:
//
// - Request URL: cloudauth.aliyuncs.com
//
// - Request Method: HTTPS POST and GET.
//
// > Supports setting blacklists for IP, ID number, phone number, bank card number, etc. When a blacklist is hit, the system rejects the request and returns a fixed error code.
//
// Supports setting blacklists for IP, ID number, phone number, bank card number, etc. When a blacklist is hit, the system rejects the request and returns a fixed error code.
//
// @param request - QueryBlackListStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryBlackListStrategyResponse
func (client *Client) QueryBlackListStrategyWithContext(ctx context.Context, request *QueryBlackListStrategyRequest, runtime *dara.RuntimeOptions) (_result *QueryBlackListStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryBlackListStrategy"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryBlackListStrategyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Security Control Strategy
//
// Description:
//
// - Request Method: Supports sending requests via HTTPS POST and GET methods.
//
// - Request Address: cloudauth.aliyuncs.com.
//
// @param request - QueryControlStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryControlStrategyResponse
func (client *Client) QueryControlStrategyWithContext(ctx context.Context, request *QueryControlStrategyRequest, runtime *dara.RuntimeOptions) (_result *QueryControlStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryControlStrategy"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryControlStrategyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Custom Flow Control Strategy
//
// Description:
//
// - Service Address: cloudauth.aliyuncs.com
//
// - Request Method: HTTPS POST and GET.
//
// - Security Rules: These are rules to ensure system security, such as monitoring for API abuse, account theft, etc. When a threshold is triggered, the system supports alerting.
//
// @param request - QueryCustomizeFlowStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCustomizeFlowStrategyResponse
func (client *Client) QueryCustomizeFlowStrategyWithContext(ctx context.Context, request *QueryCustomizeFlowStrategyRequest, runtime *dara.RuntimeOptions) (_result *QueryCustomizeFlowStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCustomizeFlowStrategy"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCustomizeFlowStrategyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Scene Configuration
//
// Description:
//
// - Service address: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST and GET.
//
// @param request - QuerySceneConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySceneConfigsResponse
func (client *Client) QuerySceneConfigsWithContext(ctx context.Context, request *QuerySceneConfigsRequest, runtime *dara.RuntimeOptions) (_result *QuerySceneConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySceneConfigs"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySceneConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Real-Person Download Task
//
// Description:
//
// Obtain the download link for statistical call data files under the product plan based on query conditions.
//
// - Method: HTTPS POST
//
// - Service Address: cloudauth.aliyuncs.com
//
// > The real-person authentication product uses CertifyId to count the number of calls. For ease of reconciliation, please retain the CertifyId field in your system.
//
// @param request - QueryVerifyDownloadTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryVerifyDownloadTaskResponse
func (client *Client) QueryVerifyDownloadTaskWithContext(ctx context.Context, request *QueryVerifyDownloadTaskRequest, runtime *dara.RuntimeOptions) (_result *QueryVerifyDownloadTaskResponse, _err error) {
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
		Action:      dara.String("QueryVerifyDownloadTask"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryVerifyDownloadTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Flow Package
//
// Description:
//
// - Service address: cloudauth.aliyuncs.com
//
// - Request method: HTTPS POST and GET.
//
// - This interface uses different parameters for different product solutions. For details, please refer to the [official documentation](https://help.aliyun.com/zh/id-verification/financial-grade-id-verification/product-overview/introduction/?spm=a2c4g.11186623.help-menu-2401581.d_0_0.13f644ecRzFHfm&scm=20140722.H_99169._.OR_help-T_cn~zh-V_1).
//
// @param request - QueryVerifyFlowPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryVerifyFlowPackageResponse
func (client *Client) QueryVerifyFlowPackageWithContext(ctx context.Context, request *QueryVerifyFlowPackageRequest, runtime *dara.RuntimeOptions) (_result *QueryVerifyFlowPackageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryVerifyFlowPackage"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryVerifyFlowPackageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Call Volume Statistics
//
// Description:
//
// - Request URL: cloudauth.aliyuncs.com
//
// - Request Method: HTTPS POST and GET.
//
// > Real-person authentication products use CertifyId to count call volume. For ease of reconciliation, please retain the CertifyId field in your system.
//
// @param request - QueryVerifyInvokeSatisticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryVerifyInvokeSatisticResponse
func (client *Client) QueryVerifyInvokeSatisticWithContext(ctx context.Context, request *QueryVerifyInvokeSatisticRequest, runtime *dara.RuntimeOptions) (_result *QueryVerifyInvokeSatisticResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductProgramList) {
		query["ProductProgramList"] = request.ProductProgramList
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.SceneIdList) {
		query["SceneIdList"] = request.SceneIdList
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.StatisticsType) {
		query["StatisticsType"] = request.StatisticsType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryVerifyInvokeSatistic"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryVerifyInvokeSatisticResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Real Person Whitelist
//
// @param tmpReq - RemoveWhiteListSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveWhiteListSettingResponse
func (client *Client) RemoveWhiteListSettingWithContext(ctx context.Context, tmpReq *RemoveWhiteListSettingRequest, runtime *dara.RuntimeOptions) (_result *RemoveWhiteListSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RemoveWhiteListSettingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.IdsShrink) {
		query["Ids"] = request.IdsShrink
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveWhiteListSetting"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveWhiteListSettingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update Ant Blockchain Transaction Scenario
//
// Description:
//
// Update the information of a financial-level authentication scenario based on the scenario ID.
//
// - Service address: cloudauth.aliyuncs.com.
//
// - Request method: HTTPS POST.
//
// @param request - UpdateAntCloudAuthSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAntCloudAuthSceneResponse
func (client *Client) UpdateAntCloudAuthSceneWithContext(ctx context.Context, request *UpdateAntCloudAuthSceneRequest, runtime *dara.RuntimeOptions) (_result *UpdateAntCloudAuthSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BindMiniProgram) {
		query["BindMiniProgram"] = request.BindMiniProgram
	}

	if !dara.IsNil(request.CheckFileBody) {
		query["CheckFileBody"] = request.CheckFileBody
	}

	if !dara.IsNil(request.CheckFileName) {
		query["CheckFileName"] = request.CheckFileName
	}

	if !dara.IsNil(request.MiniProgramName) {
		query["MiniProgramName"] = request.MiniProgramName
	}

	if !dara.IsNil(request.Platform) {
		query["Platform"] = request.Platform
	}

	if !dara.IsNil(request.ReturnPicCount) {
		query["ReturnPicCount"] = request.ReturnPicCount
	}

	if !dara.IsNil(request.ReturnVideoLength) {
		query["ReturnVideoLength"] = request.ReturnVideoLength
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.SceneName) {
		query["SceneName"] = request.SceneName
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.StoreImage) {
		query["StoreImage"] = request.StoreImage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAntCloudAuthScene"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAntCloudAuthSceneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update Scene Configuration
//
// Description:
//
// - Request Method: Supports sending requests via HTTPS POST.
//
// - Request URL: cloudauth.aliyuncs.com.
//
// @param request - UpdateSceneConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSceneConfigResponse
func (client *Client) UpdateSceneConfigWithContext(ctx context.Context, request *UpdateSceneConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateSceneConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["config"] = request.Config
	}

	if !dara.IsNil(request.Id) {
		body["id"] = request.Id
	}

	if !dara.IsNil(request.SceneId) {
		body["sceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSceneConfig"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSceneConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Five-Item Vehicle Information Recognition
//
// Description:
//
// Query basic vehicle information through the license plate number and vehicle type.
//
// @param request - Vehicle5ItemQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Vehicle5ItemQueryResponse
func (client *Client) Vehicle5ItemQueryWithContext(ctx context.Context, request *Vehicle5ItemQueryRequest, runtime *dara.RuntimeOptions) (_result *Vehicle5ItemQueryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ParamType) {
		query["ParamType"] = request.ParamType
	}

	if !dara.IsNil(request.VehicleNum) {
		query["VehicleNum"] = request.VehicleNum
	}

	if !dara.IsNil(request.VehicleType) {
		query["VehicleType"] = request.VehicleType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Vehicle5ItemQuery"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &Vehicle5ItemQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Vehicle Insurance Date Inquiry
//
// Description:
//
// Query the vehicle insurance date through the license plate number, vehicle type, and vehicle identification number (VIN).
//
// @param request - VehicleInsureQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VehicleInsureQueryResponse
func (client *Client) VehicleInsureQueryWithContext(ctx context.Context, request *VehicleInsureQueryRequest, runtime *dara.RuntimeOptions) (_result *VehicleInsureQueryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ParamType) {
		query["ParamType"] = request.ParamType
	}

	if !dara.IsNil(request.VehicleNum) {
		query["VehicleNum"] = request.VehicleNum
	}

	if !dara.IsNil(request.VehicleType) {
		query["VehicleType"] = request.VehicleType
	}

	if !dara.IsNil(request.Vin) {
		query["Vin"] = request.Vin
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VehicleInsureQuery"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VehicleInsureQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Vehicle Element Verification
//
// Description:
//
// Verifies the consistency of name, ID number, vehicle license plate, and vehicle type.
//
// @param request - VehicleMetaVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VehicleMetaVerifyResponse
func (client *Client) VehicleMetaVerifyWithContext(ctx context.Context, request *VehicleMetaVerifyRequest, runtime *dara.RuntimeOptions) (_result *VehicleMetaVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IdentifyNum) {
		query["IdentifyNum"] = request.IdentifyNum
	}

	if !dara.IsNil(request.ParamType) {
		query["ParamType"] = request.ParamType
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	if !dara.IsNil(request.VehicleNum) {
		query["VehicleNum"] = request.VehicleNum
	}

	if !dara.IsNil(request.VehicleType) {
		query["VehicleType"] = request.VehicleType
	}

	if !dara.IsNil(request.VerifyMetaType) {
		query["VerifyMetaType"] = request.VerifyMetaType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VehicleMetaVerify"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VehicleMetaVerifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Enhanced Vehicle Element Verification
//
// Description:
//
// Verifies the consistency of name, ID number, license plate number, and vehicle type, and supports returning detailed vehicle information.
//
// @param request - VehicleMetaVerifyV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VehicleMetaVerifyV2Response
func (client *Client) VehicleMetaVerifyV2WithContext(ctx context.Context, request *VehicleMetaVerifyV2Request, runtime *dara.RuntimeOptions) (_result *VehicleMetaVerifyV2Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IdentifyNum) {
		query["IdentifyNum"] = request.IdentifyNum
	}

	if !dara.IsNil(request.ParamType) {
		query["ParamType"] = request.ParamType
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	if !dara.IsNil(request.VehicleNum) {
		query["VehicleNum"] = request.VehicleNum
	}

	if !dara.IsNil(request.VehicleType) {
		query["VehicleType"] = request.VehicleType
	}

	if !dara.IsNil(request.VerifyMetaType) {
		query["VerifyMetaType"] = request.VerifyMetaType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VehicleMetaVerifyV2"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VehicleMetaVerifyV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Vehicle Information Recognition
//
// Description:
//
// Query detailed vehicle information through the license plate number and vehicle type.
//
// @param request - VehicleQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VehicleQueryResponse
func (client *Client) VehicleQueryWithContext(ctx context.Context, request *VehicleQueryRequest, runtime *dara.RuntimeOptions) (_result *VehicleQueryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ParamType) {
		query["ParamType"] = request.ParamType
	}

	if !dara.IsNil(request.VehicleNum) {
		query["VehicleNum"] = request.VehicleNum
	}

	if !dara.IsNil(request.VehicleType) {
		query["VehicleType"] = request.VehicleType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VehicleQuery"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VehicleQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call VerifyMaterial, in a pure server-side access authentication solution, input name, ID number, portrait photo, and front and back photos of the ID card (optional) for real-person authentication, and return the authentication result synchronously.
//
// Description:
//
// Preparation for Access: When integrating this API, please ensure that the corresponding preparatory work has been completed. For details, please refer to [Server-side Access Preparation](https://help.aliyun.com/document_detail/127471.html).
//
// Request Method: HTTPS POST and GET.
//
// API Description: The server of the access party submits the authentication materials to the real-person authentication service for verification and comparison, with the results returned synchronously.
//
// Applicable Scope: This interface is only applicable to pure server-side access authentication solutions.
//
// Image Upload Address Explanation:
//
// - HTTP or HTTPS address: Supports publicly accessible HTTP or HTTPS addresses. For example, `http://image-demo.img-cn-hangzhou.aliyuncs.com/example.jpg`.
//
// - OSS address: If the images from the access party are local files, Alibaba Cloud also provides an upload SDK, supporting the business party to upload the images to the specified OSS bucket of the real-person authentication service, and use the obtained OSS address as the image address parameter in the interface. If your business needs to use the upload SDK, please submit a [ticket](https://selfservice.console.aliyun.com/ticket/category/cloudauth/today) to contact us for acquisition.
//
// Image Limitations:
//
// - Does not support relative or absolute paths of local images.
//
// - Please keep the size of a single image within 2 MB to avoid algorithm retrieval timeout.
//
// - The face area in the image should be at least 64*64 pixels.
//
// - There is an 8 MB size limit for the Body of a single request. Please calculate the total size of all images and other information in the request, and do not exceed the limit.
//
// @param request - VerifyMaterialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyMaterialResponse
func (client *Client) VerifyMaterialWithContext(ctx context.Context, request *VerifyMaterialRequest, runtime *dara.RuntimeOptions) (_result *VerifyMaterialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.FaceImageUrl) {
		query["FaceImageUrl"] = request.FaceImageUrl
	}

	if !dara.IsNil(request.IdCardBackImageUrl) {
		query["IdCardBackImageUrl"] = request.IdCardBackImageUrl
	}

	if !dara.IsNil(request.IdCardFrontImageUrl) {
		query["IdCardFrontImageUrl"] = request.IdCardFrontImageUrl
	}

	if !dara.IsNil(request.IdCardNumber) {
		query["IdCardNumber"] = request.IdCardNumber
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyMaterial"),
		Version:     dara.String("2019-03-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyMaterialResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
