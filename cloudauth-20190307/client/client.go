// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openplatform "github.com/alibabacloud-go/openplatform-20191219/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	oss "github.com/alibabacloud-go/tea-oss-sdk/client"
	ossutil "github.com/alibabacloud-go/tea-oss-utils/service"
	rpcutil "github.com/alibabacloud-go/tea-rpc-utils/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type CompareFaceVerifyRequest struct {
	Crop                         *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	OuterOrderNo                 *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	ProductCode                  *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	SceneId                      *int64  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	SourceCertifyId              *string `json:"SourceCertifyId,omitempty" xml:"SourceCertifyId,omitempty"`
	SourceFaceContrastPicture    *string `json:"SourceFaceContrastPicture,omitempty" xml:"SourceFaceContrastPicture,omitempty"`
	SourceFaceContrastPictureUrl *string `json:"SourceFaceContrastPictureUrl,omitempty" xml:"SourceFaceContrastPictureUrl,omitempty"`
	SourceOssBucketName          *string `json:"SourceOssBucketName,omitempty" xml:"SourceOssBucketName,omitempty"`
	SourceOssObjectName          *string `json:"SourceOssObjectName,omitempty" xml:"SourceOssObjectName,omitempty"`
	TargetCertifyId              *string `json:"TargetCertifyId,omitempty" xml:"TargetCertifyId,omitempty"`
	TargetFaceContrastPicture    *string `json:"TargetFaceContrastPicture,omitempty" xml:"TargetFaceContrastPicture,omitempty"`
	TargetFaceContrastPictureUrl *string `json:"TargetFaceContrastPictureUrl,omitempty" xml:"TargetFaceContrastPictureUrl,omitempty"`
	TargetOssBucketName          *string `json:"TargetOssBucketName,omitempty" xml:"TargetOssBucketName,omitempty"`
	TargetOssObjectName          *string `json:"TargetOssObjectName,omitempty" xml:"TargetOssObjectName,omitempty"`
}

func (s CompareFaceVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s CompareFaceVerifyRequest) GoString() string {
	return s.String()
}

func (s *CompareFaceVerifyRequest) SetCrop(v string) *CompareFaceVerifyRequest {
	s.Crop = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetOuterOrderNo(v string) *CompareFaceVerifyRequest {
	s.OuterOrderNo = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetProductCode(v string) *CompareFaceVerifyRequest {
	s.ProductCode = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetSceneId(v int64) *CompareFaceVerifyRequest {
	s.SceneId = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetSourceCertifyId(v string) *CompareFaceVerifyRequest {
	s.SourceCertifyId = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetSourceFaceContrastPicture(v string) *CompareFaceVerifyRequest {
	s.SourceFaceContrastPicture = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetSourceFaceContrastPictureUrl(v string) *CompareFaceVerifyRequest {
	s.SourceFaceContrastPictureUrl = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetSourceOssBucketName(v string) *CompareFaceVerifyRequest {
	s.SourceOssBucketName = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetSourceOssObjectName(v string) *CompareFaceVerifyRequest {
	s.SourceOssObjectName = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetTargetCertifyId(v string) *CompareFaceVerifyRequest {
	s.TargetCertifyId = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetTargetFaceContrastPicture(v string) *CompareFaceVerifyRequest {
	s.TargetFaceContrastPicture = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetTargetFaceContrastPictureUrl(v string) *CompareFaceVerifyRequest {
	s.TargetFaceContrastPictureUrl = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetTargetOssBucketName(v string) *CompareFaceVerifyRequest {
	s.TargetOssBucketName = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetTargetOssObjectName(v string) *CompareFaceVerifyRequest {
	s.TargetOssObjectName = &v
	return s
}

type CompareFaceVerifyResponse struct {
	Code         *string                                `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message      *string                                `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ResultObject *CompareFaceVerifyResponseResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" require:"true" type:"Struct"`
}

func (s CompareFaceVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s CompareFaceVerifyResponse) GoString() string {
	return s.String()
}

func (s *CompareFaceVerifyResponse) SetCode(v string) *CompareFaceVerifyResponse {
	s.Code = &v
	return s
}

func (s *CompareFaceVerifyResponse) SetMessage(v string) *CompareFaceVerifyResponse {
	s.Message = &v
	return s
}

func (s *CompareFaceVerifyResponse) SetRequestId(v string) *CompareFaceVerifyResponse {
	s.RequestId = &v
	return s
}

func (s *CompareFaceVerifyResponse) SetResultObject(v *CompareFaceVerifyResponseResultObject) *CompareFaceVerifyResponse {
	s.ResultObject = v
	return s
}

type CompareFaceVerifyResponseResultObject struct {
	CertifyId   *string  `json:"CertifyId,omitempty" xml:"CertifyId,omitempty" require:"true"`
	Passed      *string  `json:"Passed,omitempty" xml:"Passed,omitempty" require:"true"`
	VerifyScore *float32 `json:"VerifyScore,omitempty" xml:"VerifyScore,omitempty" require:"true"`
}

func (s CompareFaceVerifyResponseResultObject) String() string {
	return tea.Prettify(s)
}

func (s CompareFaceVerifyResponseResultObject) GoString() string {
	return s.String()
}

func (s *CompareFaceVerifyResponseResultObject) SetCertifyId(v string) *CompareFaceVerifyResponseResultObject {
	s.CertifyId = &v
	return s
}

func (s *CompareFaceVerifyResponseResultObject) SetPassed(v string) *CompareFaceVerifyResponseResultObject {
	s.Passed = &v
	return s
}

func (s *CompareFaceVerifyResponseResultObject) SetVerifyScore(v float32) *CompareFaceVerifyResponseResultObject {
	s.VerifyScore = &v
	return s
}

type CompareFacesRequest struct {
	SourceImageType  *string `json:"SourceImageType,omitempty" xml:"SourceImageType,omitempty"`
	SourceImageValue *string `json:"SourceImageValue,omitempty" xml:"SourceImageValue,omitempty"`
	TargetImageType  *string `json:"TargetImageType,omitempty" xml:"TargetImageType,omitempty"`
	TargetImageValue *string `json:"TargetImageValue,omitempty" xml:"TargetImageValue,omitempty"`
}

func (s CompareFacesRequest) String() string {
	return tea.Prettify(s)
}

func (s CompareFacesRequest) GoString() string {
	return s.String()
}

func (s *CompareFacesRequest) SetSourceImageType(v string) *CompareFacesRequest {
	s.SourceImageType = &v
	return s
}

func (s *CompareFacesRequest) SetSourceImageValue(v string) *CompareFacesRequest {
	s.SourceImageValue = &v
	return s
}

func (s *CompareFacesRequest) SetTargetImageType(v string) *CompareFacesRequest {
	s.TargetImageType = &v
	return s
}

func (s *CompareFacesRequest) SetTargetImageValue(v string) *CompareFacesRequest {
	s.TargetImageValue = &v
	return s
}

type CompareFacesResponse struct {
	Code      *string                   `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                   `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success   *bool                     `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	Data      *CompareFacesResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s CompareFacesResponse) String() string {
	return tea.Prettify(s)
}

func (s CompareFacesResponse) GoString() string {
	return s.String()
}

func (s *CompareFacesResponse) SetCode(v string) *CompareFacesResponse {
	s.Code = &v
	return s
}

func (s *CompareFacesResponse) SetMessage(v string) *CompareFacesResponse {
	s.Message = &v
	return s
}

func (s *CompareFacesResponse) SetRequestId(v string) *CompareFacesResponse {
	s.RequestId = &v
	return s
}

func (s *CompareFacesResponse) SetSuccess(v bool) *CompareFacesResponse {
	s.Success = &v
	return s
}

func (s *CompareFacesResponse) SetData(v *CompareFacesResponseData) *CompareFacesResponse {
	s.Data = v
	return s
}

type CompareFacesResponseData struct {
	ConfidenceThresholds *string  `json:"ConfidenceThresholds,omitempty" xml:"ConfidenceThresholds,omitempty" require:"true"`
	SimilarityScore      *float32 `json:"SimilarityScore,omitempty" xml:"SimilarityScore,omitempty" require:"true"`
}

func (s CompareFacesResponseData) String() string {
	return tea.Prettify(s)
}

func (s CompareFacesResponseData) GoString() string {
	return s.String()
}

func (s *CompareFacesResponseData) SetConfidenceThresholds(v string) *CompareFacesResponseData {
	s.ConfidenceThresholds = &v
	return s
}

func (s *CompareFacesResponseData) SetSimilarityScore(v float32) *CompareFacesResponseData {
	s.SimilarityScore = &v
	return s
}

type ContrastFaceVerifyRequest struct {
	CertName               *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertNo                 *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	CertType               *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	CertifyId              *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	Crop                   *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	DeviceToken            *string `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
	FaceContrastFile       *string `json:"FaceContrastFile,omitempty" xml:"FaceContrastFile,omitempty"`
	FaceContrastPicture    *string `json:"FaceContrastPicture,omitempty" xml:"FaceContrastPicture,omitempty"`
	FaceContrastPictureUrl *string `json:"FaceContrastPictureUrl,omitempty" xml:"FaceContrastPictureUrl,omitempty"`
	Ip                     *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Mobile                 *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Model                  *string `json:"Model,omitempty" xml:"Model,omitempty"`
	OssBucketName          *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	OssObjectName          *string `json:"OssObjectName,omitempty" xml:"OssObjectName,omitempty"`
	OuterOrderNo           *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	ProductCode            *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	SceneId                *int64  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	UserId                 *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ContrastFaceVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s ContrastFaceVerifyRequest) GoString() string {
	return s.String()
}

func (s *ContrastFaceVerifyRequest) SetCertName(v string) *ContrastFaceVerifyRequest {
	s.CertName = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetCertNo(v string) *ContrastFaceVerifyRequest {
	s.CertNo = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetCertType(v string) *ContrastFaceVerifyRequest {
	s.CertType = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetCertifyId(v string) *ContrastFaceVerifyRequest {
	s.CertifyId = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetCrop(v string) *ContrastFaceVerifyRequest {
	s.Crop = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetDeviceToken(v string) *ContrastFaceVerifyRequest {
	s.DeviceToken = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetFaceContrastFile(v string) *ContrastFaceVerifyRequest {
	s.FaceContrastFile = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetFaceContrastPicture(v string) *ContrastFaceVerifyRequest {
	s.FaceContrastPicture = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetFaceContrastPictureUrl(v string) *ContrastFaceVerifyRequest {
	s.FaceContrastPictureUrl = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetIp(v string) *ContrastFaceVerifyRequest {
	s.Ip = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetMobile(v string) *ContrastFaceVerifyRequest {
	s.Mobile = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetModel(v string) *ContrastFaceVerifyRequest {
	s.Model = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetOssBucketName(v string) *ContrastFaceVerifyRequest {
	s.OssBucketName = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetOssObjectName(v string) *ContrastFaceVerifyRequest {
	s.OssObjectName = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetOuterOrderNo(v string) *ContrastFaceVerifyRequest {
	s.OuterOrderNo = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetProductCode(v string) *ContrastFaceVerifyRequest {
	s.ProductCode = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetSceneId(v int64) *ContrastFaceVerifyRequest {
	s.SceneId = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetUserId(v string) *ContrastFaceVerifyRequest {
	s.UserId = &v
	return s
}

type ContrastFaceVerifyResponse struct {
	Code         *string                                 `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message      *string                                 `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ResultObject *ContrastFaceVerifyResponseResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" require:"true" type:"Struct"`
}

func (s ContrastFaceVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s ContrastFaceVerifyResponse) GoString() string {
	return s.String()
}

func (s *ContrastFaceVerifyResponse) SetCode(v string) *ContrastFaceVerifyResponse {
	s.Code = &v
	return s
}

func (s *ContrastFaceVerifyResponse) SetMessage(v string) *ContrastFaceVerifyResponse {
	s.Message = &v
	return s
}

func (s *ContrastFaceVerifyResponse) SetRequestId(v string) *ContrastFaceVerifyResponse {
	s.RequestId = &v
	return s
}

func (s *ContrastFaceVerifyResponse) SetResultObject(v *ContrastFaceVerifyResponseResultObject) *ContrastFaceVerifyResponse {
	s.ResultObject = v
	return s
}

type ContrastFaceVerifyResponseResultObject struct {
	CertifyId    *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty" require:"true"`
	IdentityInfo *string `json:"IdentityInfo,omitempty" xml:"IdentityInfo,omitempty" require:"true"`
	MaterialInfo *string `json:"MaterialInfo,omitempty" xml:"MaterialInfo,omitempty" require:"true"`
	Passed       *string `json:"Passed,omitempty" xml:"Passed,omitempty" require:"true"`
	SubCode      *string `json:"SubCode,omitempty" xml:"SubCode,omitempty" require:"true"`
}

func (s ContrastFaceVerifyResponseResultObject) String() string {
	return tea.Prettify(s)
}

func (s ContrastFaceVerifyResponseResultObject) GoString() string {
	return s.String()
}

func (s *ContrastFaceVerifyResponseResultObject) SetCertifyId(v string) *ContrastFaceVerifyResponseResultObject {
	s.CertifyId = &v
	return s
}

func (s *ContrastFaceVerifyResponseResultObject) SetIdentityInfo(v string) *ContrastFaceVerifyResponseResultObject {
	s.IdentityInfo = &v
	return s
}

func (s *ContrastFaceVerifyResponseResultObject) SetMaterialInfo(v string) *ContrastFaceVerifyResponseResultObject {
	s.MaterialInfo = &v
	return s
}

func (s *ContrastFaceVerifyResponseResultObject) SetPassed(v string) *ContrastFaceVerifyResponseResultObject {
	s.Passed = &v
	return s
}

func (s *ContrastFaceVerifyResponseResultObject) SetSubCode(v string) *ContrastFaceVerifyResponseResultObject {
	s.SubCode = &v
	return s
}

type ContrastFaceVerifyAdvanceRequest struct {
	FaceContrastFileObject io.Reader `json:"FaceContrastFileObject,omitempty" xml:"FaceContrastFileObject,omitempty" require:"true"`
	CertName               *string   `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertNo                 *string   `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	CertType               *string   `json:"CertType,omitempty" xml:"CertType,omitempty"`
	CertifyId              *string   `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	Crop                   *string   `json:"Crop,omitempty" xml:"Crop,omitempty"`
	DeviceToken            *string   `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
	FaceContrastPicture    *string   `json:"FaceContrastPicture,omitempty" xml:"FaceContrastPicture,omitempty"`
	FaceContrastPictureUrl *string   `json:"FaceContrastPictureUrl,omitempty" xml:"FaceContrastPictureUrl,omitempty"`
	Ip                     *string   `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Mobile                 *string   `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Model                  *string   `json:"Model,omitempty" xml:"Model,omitempty"`
	OssBucketName          *string   `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	OssObjectName          *string   `json:"OssObjectName,omitempty" xml:"OssObjectName,omitempty"`
	OuterOrderNo           *string   `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	ProductCode            *string   `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	SceneId                *int64    `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	UserId                 *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ContrastFaceVerifyAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ContrastFaceVerifyAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ContrastFaceVerifyAdvanceRequest) SetFaceContrastFileObject(v io.Reader) *ContrastFaceVerifyAdvanceRequest {
	s.FaceContrastFileObject = v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetCertName(v string) *ContrastFaceVerifyAdvanceRequest {
	s.CertName = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetCertNo(v string) *ContrastFaceVerifyAdvanceRequest {
	s.CertNo = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetCertType(v string) *ContrastFaceVerifyAdvanceRequest {
	s.CertType = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetCertifyId(v string) *ContrastFaceVerifyAdvanceRequest {
	s.CertifyId = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetCrop(v string) *ContrastFaceVerifyAdvanceRequest {
	s.Crop = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetDeviceToken(v string) *ContrastFaceVerifyAdvanceRequest {
	s.DeviceToken = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetFaceContrastPicture(v string) *ContrastFaceVerifyAdvanceRequest {
	s.FaceContrastPicture = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetFaceContrastPictureUrl(v string) *ContrastFaceVerifyAdvanceRequest {
	s.FaceContrastPictureUrl = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetIp(v string) *ContrastFaceVerifyAdvanceRequest {
	s.Ip = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetMobile(v string) *ContrastFaceVerifyAdvanceRequest {
	s.Mobile = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetModel(v string) *ContrastFaceVerifyAdvanceRequest {
	s.Model = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetOssBucketName(v string) *ContrastFaceVerifyAdvanceRequest {
	s.OssBucketName = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetOssObjectName(v string) *ContrastFaceVerifyAdvanceRequest {
	s.OssObjectName = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetOuterOrderNo(v string) *ContrastFaceVerifyAdvanceRequest {
	s.OuterOrderNo = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetProductCode(v string) *ContrastFaceVerifyAdvanceRequest {
	s.ProductCode = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetSceneId(v int64) *ContrastFaceVerifyAdvanceRequest {
	s.SceneId = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetUserId(v string) *ContrastFaceVerifyAdvanceRequest {
	s.UserId = &v
	return s
}

type CreateAuthKeyRequest struct {
	AuthYears    *int    `json:"AuthYears,omitempty" xml:"AuthYears,omitempty"`
	BizType      *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	Test         *bool   `json:"Test,omitempty" xml:"Test,omitempty"`
	UserDeviceId *string `json:"UserDeviceId,omitempty" xml:"UserDeviceId,omitempty"`
}

func (s CreateAuthKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAuthKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateAuthKeyRequest) SetAuthYears(v int) *CreateAuthKeyRequest {
	s.AuthYears = &v
	return s
}

func (s *CreateAuthKeyRequest) SetBizType(v string) *CreateAuthKeyRequest {
	s.BizType = &v
	return s
}

func (s *CreateAuthKeyRequest) SetTest(v bool) *CreateAuthKeyRequest {
	s.Test = &v
	return s
}

func (s *CreateAuthKeyRequest) SetUserDeviceId(v string) *CreateAuthKeyRequest {
	s.UserDeviceId = &v
	return s
}

type CreateAuthKeyResponse struct {
	AuthKey   *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s CreateAuthKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAuthKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateAuthKeyResponse) SetAuthKey(v string) *CreateAuthKeyResponse {
	s.AuthKey = &v
	return s
}

func (s *CreateAuthKeyResponse) SetRequestId(v string) *CreateAuthKeyResponse {
	s.RequestId = &v
	return s
}

type CreateVerifySettingRequest struct {
	BizName     *string `json:"BizName,omitempty" xml:"BizName,omitempty" require:"true"`
	BizType     *string `json:"BizType,omitempty" xml:"BizType,omitempty" require:"true"`
	GuideStep   *bool   `json:"GuideStep,omitempty" xml:"GuideStep,omitempty"`
	PrivacyStep *bool   `json:"PrivacyStep,omitempty" xml:"PrivacyStep,omitempty"`
	ResultStep  *bool   `json:"ResultStep,omitempty" xml:"ResultStep,omitempty"`
	Solution    *string `json:"Solution,omitempty" xml:"Solution,omitempty" require:"true"`
}

func (s CreateVerifySettingRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVerifySettingRequest) GoString() string {
	return s.String()
}

func (s *CreateVerifySettingRequest) SetBizName(v string) *CreateVerifySettingRequest {
	s.BizName = &v
	return s
}

func (s *CreateVerifySettingRequest) SetBizType(v string) *CreateVerifySettingRequest {
	s.BizType = &v
	return s
}

func (s *CreateVerifySettingRequest) SetGuideStep(v bool) *CreateVerifySettingRequest {
	s.GuideStep = &v
	return s
}

func (s *CreateVerifySettingRequest) SetPrivacyStep(v bool) *CreateVerifySettingRequest {
	s.PrivacyStep = &v
	return s
}

func (s *CreateVerifySettingRequest) SetResultStep(v bool) *CreateVerifySettingRequest {
	s.ResultStep = &v
	return s
}

func (s *CreateVerifySettingRequest) SetSolution(v string) *CreateVerifySettingRequest {
	s.Solution = &v
	return s
}

type CreateVerifySettingResponse struct {
	BizName   *string   `json:"BizName,omitempty" xml:"BizName,omitempty" require:"true"`
	BizType   *string   `json:"BizType,omitempty" xml:"BizType,omitempty" require:"true"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Solution  *string   `json:"Solution,omitempty" xml:"Solution,omitempty" require:"true"`
	StepList  []*string `json:"StepList,omitempty" xml:"StepList,omitempty" require:"true" type:"Repeated"`
}

func (s CreateVerifySettingResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVerifySettingResponse) GoString() string {
	return s.String()
}

func (s *CreateVerifySettingResponse) SetBizName(v string) *CreateVerifySettingResponse {
	s.BizName = &v
	return s
}

func (s *CreateVerifySettingResponse) SetBizType(v string) *CreateVerifySettingResponse {
	s.BizType = &v
	return s
}

func (s *CreateVerifySettingResponse) SetRequestId(v string) *CreateVerifySettingResponse {
	s.RequestId = &v
	return s
}

func (s *CreateVerifySettingResponse) SetSolution(v string) *CreateVerifySettingResponse {
	s.Solution = &v
	return s
}

func (s *CreateVerifySettingResponse) SetStepList(v []*string) *CreateVerifySettingResponse {
	s.StepList = v
	return s
}

type DescribeDeviceInfoRequest struct {
	BizType         *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	CurrentPage     *int    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DeviceId        *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	ExpiredEndDay   *string `json:"ExpiredEndDay,omitempty" xml:"ExpiredEndDay,omitempty"`
	ExpiredStartDay *string `json:"ExpiredStartDay,omitempty" xml:"ExpiredStartDay,omitempty"`
	PageSize        *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	UserDeviceId    *string `json:"UserDeviceId,omitempty" xml:"UserDeviceId,omitempty"`
}

func (s DescribeDeviceInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeviceInfoRequest) SetBizType(v string) *DescribeDeviceInfoRequest {
	s.BizType = &v
	return s
}

func (s *DescribeDeviceInfoRequest) SetCurrentPage(v int) *DescribeDeviceInfoRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDeviceInfoRequest) SetDeviceId(v string) *DescribeDeviceInfoRequest {
	s.DeviceId = &v
	return s
}

func (s *DescribeDeviceInfoRequest) SetExpiredEndDay(v string) *DescribeDeviceInfoRequest {
	s.ExpiredEndDay = &v
	return s
}

func (s *DescribeDeviceInfoRequest) SetExpiredStartDay(v string) *DescribeDeviceInfoRequest {
	s.ExpiredStartDay = &v
	return s
}

func (s *DescribeDeviceInfoRequest) SetPageSize(v int) *DescribeDeviceInfoRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDeviceInfoRequest) SetUserDeviceId(v string) *DescribeDeviceInfoRequest {
	s.UserDeviceId = &v
	return s
}

type DescribeDeviceInfoResponse struct {
	CurrentPage    *int                                      `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty" require:"true"`
	PageSize       *int                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	RequestId      *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalCount     *int                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	DeviceInfoList *DescribeDeviceInfoResponseDeviceInfoList `json:"DeviceInfoList,omitempty" xml:"DeviceInfoList,omitempty" require:"true" type:"Struct"`
}

func (s DescribeDeviceInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeviceInfoResponse) SetCurrentPage(v int) *DescribeDeviceInfoResponse {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDeviceInfoResponse) SetPageSize(v int) *DescribeDeviceInfoResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeDeviceInfoResponse) SetRequestId(v string) *DescribeDeviceInfoResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDeviceInfoResponse) SetTotalCount(v int) *DescribeDeviceInfoResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeDeviceInfoResponse) SetDeviceInfoList(v *DescribeDeviceInfoResponseDeviceInfoList) *DescribeDeviceInfoResponse {
	s.DeviceInfoList = v
	return s
}

type DescribeDeviceInfoResponseDeviceInfoList struct {
	DeviceInfo []*DescribeDeviceInfoResponseDeviceInfoListDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDeviceInfoResponseDeviceInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceInfoResponseDeviceInfoList) GoString() string {
	return s.String()
}

func (s *DescribeDeviceInfoResponseDeviceInfoList) SetDeviceInfo(v []*DescribeDeviceInfoResponseDeviceInfoListDeviceInfo) *DescribeDeviceInfoResponseDeviceInfoList {
	s.DeviceInfo = v
	return s
}

type DescribeDeviceInfoResponseDeviceInfoListDeviceInfo struct {
	BeginDay     *string `json:"BeginDay,omitempty" xml:"BeginDay,omitempty" require:"true"`
	BizType      *string `json:"BizType,omitempty" xml:"BizType,omitempty" require:"true"`
	DeviceId     *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty" require:"true"`
	ExpiredDay   *string `json:"ExpiredDay,omitempty" xml:"ExpiredDay,omitempty" require:"true"`
	UserDeviceId *string `json:"UserDeviceId,omitempty" xml:"UserDeviceId,omitempty" require:"true"`
}

func (s DescribeDeviceInfoResponseDeviceInfoListDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceInfoResponseDeviceInfoListDeviceInfo) GoString() string {
	return s.String()
}

func (s *DescribeDeviceInfoResponseDeviceInfoListDeviceInfo) SetBeginDay(v string) *DescribeDeviceInfoResponseDeviceInfoListDeviceInfo {
	s.BeginDay = &v
	return s
}

func (s *DescribeDeviceInfoResponseDeviceInfoListDeviceInfo) SetBizType(v string) *DescribeDeviceInfoResponseDeviceInfoListDeviceInfo {
	s.BizType = &v
	return s
}

func (s *DescribeDeviceInfoResponseDeviceInfoListDeviceInfo) SetDeviceId(v string) *DescribeDeviceInfoResponseDeviceInfoListDeviceInfo {
	s.DeviceId = &v
	return s
}

func (s *DescribeDeviceInfoResponseDeviceInfoListDeviceInfo) SetExpiredDay(v string) *DescribeDeviceInfoResponseDeviceInfoListDeviceInfo {
	s.ExpiredDay = &v
	return s
}

func (s *DescribeDeviceInfoResponseDeviceInfoListDeviceInfo) SetUserDeviceId(v string) *DescribeDeviceInfoResponseDeviceInfoListDeviceInfo {
	s.UserDeviceId = &v
	return s
}

type DescribeFaceVerifyRequest struct {
	CertifyId         *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	PictureReturnType *string `json:"PictureReturnType,omitempty" xml:"PictureReturnType,omitempty"`
	SceneId           *int64  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s DescribeFaceVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaceVerifyRequest) GoString() string {
	return s.String()
}

func (s *DescribeFaceVerifyRequest) SetCertifyId(v string) *DescribeFaceVerifyRequest {
	s.CertifyId = &v
	return s
}

func (s *DescribeFaceVerifyRequest) SetPictureReturnType(v string) *DescribeFaceVerifyRequest {
	s.PictureReturnType = &v
	return s
}

func (s *DescribeFaceVerifyRequest) SetSceneId(v int64) *DescribeFaceVerifyRequest {
	s.SceneId = &v
	return s
}

type DescribeFaceVerifyResponse struct {
	Code         *string                                 `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message      *string                                 `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ResultObject *DescribeFaceVerifyResponseResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" require:"true" type:"Struct"`
}

func (s DescribeFaceVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaceVerifyResponse) GoString() string {
	return s.String()
}

func (s *DescribeFaceVerifyResponse) SetCode(v string) *DescribeFaceVerifyResponse {
	s.Code = &v
	return s
}

func (s *DescribeFaceVerifyResponse) SetMessage(v string) *DescribeFaceVerifyResponse {
	s.Message = &v
	return s
}

func (s *DescribeFaceVerifyResponse) SetRequestId(v string) *DescribeFaceVerifyResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeFaceVerifyResponse) SetResultObject(v *DescribeFaceVerifyResponseResultObject) *DescribeFaceVerifyResponse {
	s.ResultObject = v
	return s
}

type DescribeFaceVerifyResponseResultObject struct {
	DeviceToken  *string `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty" require:"true"`
	IdentityInfo *string `json:"IdentityInfo,omitempty" xml:"IdentityInfo,omitempty" require:"true"`
	MaterialInfo *string `json:"MaterialInfo,omitempty" xml:"MaterialInfo,omitempty" require:"true"`
	Passed       *string `json:"Passed,omitempty" xml:"Passed,omitempty" require:"true"`
	SubCode      *string `json:"SubCode,omitempty" xml:"SubCode,omitempty" require:"true"`
}

func (s DescribeFaceVerifyResponseResultObject) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaceVerifyResponseResultObject) GoString() string {
	return s.String()
}

func (s *DescribeFaceVerifyResponseResultObject) SetDeviceToken(v string) *DescribeFaceVerifyResponseResultObject {
	s.DeviceToken = &v
	return s
}

func (s *DescribeFaceVerifyResponseResultObject) SetIdentityInfo(v string) *DescribeFaceVerifyResponseResultObject {
	s.IdentityInfo = &v
	return s
}

func (s *DescribeFaceVerifyResponseResultObject) SetMaterialInfo(v string) *DescribeFaceVerifyResponseResultObject {
	s.MaterialInfo = &v
	return s
}

func (s *DescribeFaceVerifyResponseResultObject) SetPassed(v string) *DescribeFaceVerifyResponseResultObject {
	s.Passed = &v
	return s
}

func (s *DescribeFaceVerifyResponseResultObject) SetSubCode(v string) *DescribeFaceVerifyResponseResultObject {
	s.SubCode = &v
	return s
}

type DescribeOssUploadTokenRequest struct {
}

func (s DescribeOssUploadTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssUploadTokenRequest) GoString() string {
	return s.String()
}

type DescribeOssUploadTokenResponse struct {
	RequestId      *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	OssUploadToken *DescribeOssUploadTokenResponseOssUploadToken `json:"OssUploadToken,omitempty" xml:"OssUploadToken,omitempty" require:"true" type:"Struct"`
}

func (s DescribeOssUploadTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssUploadTokenResponse) GoString() string {
	return s.String()
}

func (s *DescribeOssUploadTokenResponse) SetRequestId(v string) *DescribeOssUploadTokenResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeOssUploadTokenResponse) SetOssUploadToken(v *DescribeOssUploadTokenResponseOssUploadToken) *DescribeOssUploadTokenResponse {
	s.OssUploadToken = v
	return s
}

type DescribeOssUploadTokenResponseOssUploadToken struct {
	Bucket   *string `json:"Bucket,omitempty" xml:"Bucket,omitempty" require:"true"`
	EndPoint *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty" require:"true"`
	Expired  *int64  `json:"Expired,omitempty" xml:"Expired,omitempty" require:"true"`
	Key      *string `json:"Key,omitempty" xml:"Key,omitempty" require:"true"`
	Path     *string `json:"Path,omitempty" xml:"Path,omitempty" require:"true"`
	Secret   *string `json:"Secret,omitempty" xml:"Secret,omitempty" require:"true"`
	Token    *string `json:"Token,omitempty" xml:"Token,omitempty" require:"true"`
}

func (s DescribeOssUploadTokenResponseOssUploadToken) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssUploadTokenResponseOssUploadToken) GoString() string {
	return s.String()
}

func (s *DescribeOssUploadTokenResponseOssUploadToken) SetBucket(v string) *DescribeOssUploadTokenResponseOssUploadToken {
	s.Bucket = &v
	return s
}

func (s *DescribeOssUploadTokenResponseOssUploadToken) SetEndPoint(v string) *DescribeOssUploadTokenResponseOssUploadToken {
	s.EndPoint = &v
	return s
}

func (s *DescribeOssUploadTokenResponseOssUploadToken) SetExpired(v int64) *DescribeOssUploadTokenResponseOssUploadToken {
	s.Expired = &v
	return s
}

func (s *DescribeOssUploadTokenResponseOssUploadToken) SetKey(v string) *DescribeOssUploadTokenResponseOssUploadToken {
	s.Key = &v
	return s
}

func (s *DescribeOssUploadTokenResponseOssUploadToken) SetPath(v string) *DescribeOssUploadTokenResponseOssUploadToken {
	s.Path = &v
	return s
}

func (s *DescribeOssUploadTokenResponseOssUploadToken) SetSecret(v string) *DescribeOssUploadTokenResponseOssUploadToken {
	s.Secret = &v
	return s
}

func (s *DescribeOssUploadTokenResponseOssUploadToken) SetToken(v string) *DescribeOssUploadTokenResponseOssUploadToken {
	s.Token = &v
	return s
}

type DescribeVerifyResultRequest struct {
	BizId   *string `json:"BizId,omitempty" xml:"BizId,omitempty" require:"true"`
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty" require:"true"`
}

func (s DescribeVerifyResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeVerifyResultRequest) SetBizId(v string) *DescribeVerifyResultRequest {
	s.BizId = &v
	return s
}

func (s *DescribeVerifyResultRequest) SetBizType(v string) *DescribeVerifyResultRequest {
	s.BizType = &v
	return s
}

type DescribeVerifyResultResponse struct {
	AuthorityComparisionScore *float32                              `json:"AuthorityComparisionScore,omitempty" xml:"AuthorityComparisionScore,omitempty" require:"true"`
	FaceComparisonScore       *float32                              `json:"FaceComparisonScore,omitempty" xml:"FaceComparisonScore,omitempty" require:"true"`
	IdCardFaceComparisonScore *float32                              `json:"IdCardFaceComparisonScore,omitempty" xml:"IdCardFaceComparisonScore,omitempty" require:"true"`
	RequestId                 *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	VerifyStatus              *int                                  `json:"VerifyStatus,omitempty" xml:"VerifyStatus,omitempty" require:"true"`
	Material                  *DescribeVerifyResultResponseMaterial `json:"Material,omitempty" xml:"Material,omitempty" require:"true" type:"Struct"`
}

func (s DescribeVerifyResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeVerifyResultResponse) SetAuthorityComparisionScore(v float32) *DescribeVerifyResultResponse {
	s.AuthorityComparisionScore = &v
	return s
}

func (s *DescribeVerifyResultResponse) SetFaceComparisonScore(v float32) *DescribeVerifyResultResponse {
	s.FaceComparisonScore = &v
	return s
}

func (s *DescribeVerifyResultResponse) SetIdCardFaceComparisonScore(v float32) *DescribeVerifyResultResponse {
	s.IdCardFaceComparisonScore = &v
	return s
}

func (s *DescribeVerifyResultResponse) SetRequestId(v string) *DescribeVerifyResultResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeVerifyResultResponse) SetVerifyStatus(v int) *DescribeVerifyResultResponse {
	s.VerifyStatus = &v
	return s
}

func (s *DescribeVerifyResultResponse) SetMaterial(v *DescribeVerifyResultResponseMaterial) *DescribeVerifyResultResponse {
	s.Material = v
	return s
}

type DescribeVerifyResultResponseMaterial struct {
	FaceGlobalUrl *string                                         `json:"FaceGlobalUrl,omitempty" xml:"FaceGlobalUrl,omitempty" require:"true"`
	FaceImageUrl  *string                                         `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty" require:"true"`
	FaceMask      *bool                                           `json:"FaceMask,omitempty" xml:"FaceMask,omitempty" require:"true"`
	FaceQuality   *string                                         `json:"FaceQuality,omitempty" xml:"FaceQuality,omitempty" require:"true"`
	IdCardName    *string                                         `json:"IdCardName,omitempty" xml:"IdCardName,omitempty" require:"true"`
	IdCardNumber  *string                                         `json:"IdCardNumber,omitempty" xml:"IdCardNumber,omitempty" require:"true"`
	IdCardInfo    *DescribeVerifyResultResponseMaterialIdCardInfo `json:"IdCardInfo,omitempty" xml:"IdCardInfo,omitempty" require:"true" type:"Struct"`
	VideoUrls     []*string                                       `json:"VideoUrls,omitempty" xml:"VideoUrls,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeVerifyResultResponseMaterial) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyResultResponseMaterial) GoString() string {
	return s.String()
}

func (s *DescribeVerifyResultResponseMaterial) SetFaceGlobalUrl(v string) *DescribeVerifyResultResponseMaterial {
	s.FaceGlobalUrl = &v
	return s
}

func (s *DescribeVerifyResultResponseMaterial) SetFaceImageUrl(v string) *DescribeVerifyResultResponseMaterial {
	s.FaceImageUrl = &v
	return s
}

func (s *DescribeVerifyResultResponseMaterial) SetFaceMask(v bool) *DescribeVerifyResultResponseMaterial {
	s.FaceMask = &v
	return s
}

func (s *DescribeVerifyResultResponseMaterial) SetFaceQuality(v string) *DescribeVerifyResultResponseMaterial {
	s.FaceQuality = &v
	return s
}

func (s *DescribeVerifyResultResponseMaterial) SetIdCardName(v string) *DescribeVerifyResultResponseMaterial {
	s.IdCardName = &v
	return s
}

func (s *DescribeVerifyResultResponseMaterial) SetIdCardNumber(v string) *DescribeVerifyResultResponseMaterial {
	s.IdCardNumber = &v
	return s
}

func (s *DescribeVerifyResultResponseMaterial) SetIdCardInfo(v *DescribeVerifyResultResponseMaterialIdCardInfo) *DescribeVerifyResultResponseMaterial {
	s.IdCardInfo = v
	return s
}

func (s *DescribeVerifyResultResponseMaterial) SetVideoUrls(v []*string) *DescribeVerifyResultResponseMaterial {
	s.VideoUrls = v
	return s
}

type DescribeVerifyResultResponseMaterialIdCardInfo struct {
	Address       *string `json:"Address,omitempty" xml:"Address,omitempty" require:"true"`
	Authority     *string `json:"Authority,omitempty" xml:"Authority,omitempty" require:"true"`
	BackImageUrl  *string `json:"BackImageUrl,omitempty" xml:"BackImageUrl,omitempty" require:"true"`
	Birth         *string `json:"Birth,omitempty" xml:"Birth,omitempty" require:"true"`
	EndDate       *string `json:"EndDate,omitempty" xml:"EndDate,omitempty" require:"true"`
	FrontImageUrl *string `json:"FrontImageUrl,omitempty" xml:"FrontImageUrl,omitempty" require:"true"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Nationality   *string `json:"Nationality,omitempty" xml:"Nationality,omitempty" require:"true"`
	Number        *string `json:"Number,omitempty" xml:"Number,omitempty" require:"true"`
	StartDate     *string `json:"StartDate,omitempty" xml:"StartDate,omitempty" require:"true"`
}

func (s DescribeVerifyResultResponseMaterialIdCardInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyResultResponseMaterialIdCardInfo) GoString() string {
	return s.String()
}

func (s *DescribeVerifyResultResponseMaterialIdCardInfo) SetAddress(v string) *DescribeVerifyResultResponseMaterialIdCardInfo {
	s.Address = &v
	return s
}

func (s *DescribeVerifyResultResponseMaterialIdCardInfo) SetAuthority(v string) *DescribeVerifyResultResponseMaterialIdCardInfo {
	s.Authority = &v
	return s
}

func (s *DescribeVerifyResultResponseMaterialIdCardInfo) SetBackImageUrl(v string) *DescribeVerifyResultResponseMaterialIdCardInfo {
	s.BackImageUrl = &v
	return s
}

func (s *DescribeVerifyResultResponseMaterialIdCardInfo) SetBirth(v string) *DescribeVerifyResultResponseMaterialIdCardInfo {
	s.Birth = &v
	return s
}

func (s *DescribeVerifyResultResponseMaterialIdCardInfo) SetEndDate(v string) *DescribeVerifyResultResponseMaterialIdCardInfo {
	s.EndDate = &v
	return s
}

func (s *DescribeVerifyResultResponseMaterialIdCardInfo) SetFrontImageUrl(v string) *DescribeVerifyResultResponseMaterialIdCardInfo {
	s.FrontImageUrl = &v
	return s
}

func (s *DescribeVerifyResultResponseMaterialIdCardInfo) SetName(v string) *DescribeVerifyResultResponseMaterialIdCardInfo {
	s.Name = &v
	return s
}

func (s *DescribeVerifyResultResponseMaterialIdCardInfo) SetNationality(v string) *DescribeVerifyResultResponseMaterialIdCardInfo {
	s.Nationality = &v
	return s
}

func (s *DescribeVerifyResultResponseMaterialIdCardInfo) SetNumber(v string) *DescribeVerifyResultResponseMaterialIdCardInfo {
	s.Number = &v
	return s
}

func (s *DescribeVerifyResultResponseMaterialIdCardInfo) SetStartDate(v string) *DescribeVerifyResultResponseMaterialIdCardInfo {
	s.StartDate = &v
	return s
}

type DescribeVerifySDKRequest struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
}

func (s DescribeVerifySDKRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifySDKRequest) GoString() string {
	return s.String()
}

func (s *DescribeVerifySDKRequest) SetTaskId(v string) *DescribeVerifySDKRequest {
	s.TaskId = &v
	return s
}

type DescribeVerifySDKResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SdkUrl    *string `json:"SdkUrl,omitempty" xml:"SdkUrl,omitempty" require:"true"`
}

func (s DescribeVerifySDKResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifySDKResponse) GoString() string {
	return s.String()
}

func (s *DescribeVerifySDKResponse) SetRequestId(v string) *DescribeVerifySDKResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeVerifySDKResponse) SetSdkUrl(v string) *DescribeVerifySDKResponse {
	s.SdkUrl = &v
	return s
}

type DescribeVerifyTokenRequest struct {
	BizId                *string `json:"BizId,omitempty" xml:"BizId,omitempty" require:"true"`
	BizType              *string `json:"BizType,omitempty" xml:"BizType,omitempty" require:"true"`
	CallbackSeed         *string `json:"CallbackSeed,omitempty" xml:"CallbackSeed,omitempty"`
	CallbackUrl          *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	FaceRetainedImageUrl *string `json:"FaceRetainedImageUrl,omitempty" xml:"FaceRetainedImageUrl,omitempty"`
	FailedRedirectUrl    *string `json:"FailedRedirectUrl,omitempty" xml:"FailedRedirectUrl,omitempty"`
	IdCardBackImageUrl   *string `json:"IdCardBackImageUrl,omitempty" xml:"IdCardBackImageUrl,omitempty"`
	IdCardFrontImageUrl  *string `json:"IdCardFrontImageUrl,omitempty" xml:"IdCardFrontImageUrl,omitempty"`
	IdCardNumber         *string `json:"IdCardNumber,omitempty" xml:"IdCardNumber,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PassedRedirectUrl    *string `json:"PassedRedirectUrl,omitempty" xml:"PassedRedirectUrl,omitempty"`
	UserId               *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserIp               *string `json:"UserIp,omitempty" xml:"UserIp,omitempty"`
	UserPhoneNumber      *string `json:"UserPhoneNumber,omitempty" xml:"UserPhoneNumber,omitempty"`
	UserRegistTime       *int64  `json:"UserRegistTime,omitempty" xml:"UserRegistTime,omitempty"`
}

func (s DescribeVerifyTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyTokenRequest) GoString() string {
	return s.String()
}

func (s *DescribeVerifyTokenRequest) SetBizId(v string) *DescribeVerifyTokenRequest {
	s.BizId = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetBizType(v string) *DescribeVerifyTokenRequest {
	s.BizType = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetCallbackSeed(v string) *DescribeVerifyTokenRequest {
	s.CallbackSeed = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetCallbackUrl(v string) *DescribeVerifyTokenRequest {
	s.CallbackUrl = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetFaceRetainedImageUrl(v string) *DescribeVerifyTokenRequest {
	s.FaceRetainedImageUrl = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetFailedRedirectUrl(v string) *DescribeVerifyTokenRequest {
	s.FailedRedirectUrl = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetIdCardBackImageUrl(v string) *DescribeVerifyTokenRequest {
	s.IdCardBackImageUrl = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetIdCardFrontImageUrl(v string) *DescribeVerifyTokenRequest {
	s.IdCardFrontImageUrl = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetIdCardNumber(v string) *DescribeVerifyTokenRequest {
	s.IdCardNumber = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetName(v string) *DescribeVerifyTokenRequest {
	s.Name = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetPassedRedirectUrl(v string) *DescribeVerifyTokenRequest {
	s.PassedRedirectUrl = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetUserId(v string) *DescribeVerifyTokenRequest {
	s.UserId = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetUserIp(v string) *DescribeVerifyTokenRequest {
	s.UserIp = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetUserPhoneNumber(v string) *DescribeVerifyTokenRequest {
	s.UserPhoneNumber = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetUserRegistTime(v int64) *DescribeVerifyTokenRequest {
	s.UserRegistTime = &v
	return s
}

type DescribeVerifyTokenResponse struct {
	RequestId      *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	VerifyPageUrl  *string                                    `json:"VerifyPageUrl,omitempty" xml:"VerifyPageUrl,omitempty" require:"true"`
	VerifyToken    *string                                    `json:"VerifyToken,omitempty" xml:"VerifyToken,omitempty" require:"true"`
	OssUploadToken *DescribeVerifyTokenResponseOssUploadToken `json:"OssUploadToken,omitempty" xml:"OssUploadToken,omitempty" require:"true" type:"Struct"`
}

func (s DescribeVerifyTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyTokenResponse) GoString() string {
	return s.String()
}

func (s *DescribeVerifyTokenResponse) SetRequestId(v string) *DescribeVerifyTokenResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeVerifyTokenResponse) SetVerifyPageUrl(v string) *DescribeVerifyTokenResponse {
	s.VerifyPageUrl = &v
	return s
}

func (s *DescribeVerifyTokenResponse) SetVerifyToken(v string) *DescribeVerifyTokenResponse {
	s.VerifyToken = &v
	return s
}

func (s *DescribeVerifyTokenResponse) SetOssUploadToken(v *DescribeVerifyTokenResponseOssUploadToken) *DescribeVerifyTokenResponse {
	s.OssUploadToken = v
	return s
}

type DescribeVerifyTokenResponseOssUploadToken struct {
	Bucket   *string `json:"Bucket,omitempty" xml:"Bucket,omitempty" require:"true"`
	EndPoint *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty" require:"true"`
	Expired  *int64  `json:"Expired,omitempty" xml:"Expired,omitempty" require:"true"`
	Key      *string `json:"Key,omitempty" xml:"Key,omitempty" require:"true"`
	Path     *string `json:"Path,omitempty" xml:"Path,omitempty" require:"true"`
	Secret   *string `json:"Secret,omitempty" xml:"Secret,omitempty" require:"true"`
	Token    *string `json:"Token,omitempty" xml:"Token,omitempty" require:"true"`
}

func (s DescribeVerifyTokenResponseOssUploadToken) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyTokenResponseOssUploadToken) GoString() string {
	return s.String()
}

func (s *DescribeVerifyTokenResponseOssUploadToken) SetBucket(v string) *DescribeVerifyTokenResponseOssUploadToken {
	s.Bucket = &v
	return s
}

func (s *DescribeVerifyTokenResponseOssUploadToken) SetEndPoint(v string) *DescribeVerifyTokenResponseOssUploadToken {
	s.EndPoint = &v
	return s
}

func (s *DescribeVerifyTokenResponseOssUploadToken) SetExpired(v int64) *DescribeVerifyTokenResponseOssUploadToken {
	s.Expired = &v
	return s
}

func (s *DescribeVerifyTokenResponseOssUploadToken) SetKey(v string) *DescribeVerifyTokenResponseOssUploadToken {
	s.Key = &v
	return s
}

func (s *DescribeVerifyTokenResponseOssUploadToken) SetPath(v string) *DescribeVerifyTokenResponseOssUploadToken {
	s.Path = &v
	return s
}

func (s *DescribeVerifyTokenResponseOssUploadToken) SetSecret(v string) *DescribeVerifyTokenResponseOssUploadToken {
	s.Secret = &v
	return s
}

func (s *DescribeVerifyTokenResponseOssUploadToken) SetToken(v string) *DescribeVerifyTokenResponseOssUploadToken {
	s.Token = &v
	return s
}

type DetectFaceAttributesRequest struct {
	BizType       *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	MaterialValue *string `json:"MaterialValue,omitempty" xml:"MaterialValue,omitempty" require:"true"`
}

func (s DetectFaceAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesRequest) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesRequest) SetBizType(v string) *DetectFaceAttributesRequest {
	s.BizType = &v
	return s
}

func (s *DetectFaceAttributesRequest) SetMaterialValue(v string) *DetectFaceAttributesRequest {
	s.MaterialValue = &v
	return s
}

type DetectFaceAttributesResponse struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	Data      *DetectFaceAttributesResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s DetectFaceAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponse) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponse) SetCode(v string) *DetectFaceAttributesResponse {
	s.Code = &v
	return s
}

func (s *DetectFaceAttributesResponse) SetMessage(v string) *DetectFaceAttributesResponse {
	s.Message = &v
	return s
}

func (s *DetectFaceAttributesResponse) SetRequestId(v string) *DetectFaceAttributesResponse {
	s.RequestId = &v
	return s
}

func (s *DetectFaceAttributesResponse) SetSuccess(v bool) *DetectFaceAttributesResponse {
	s.Success = &v
	return s
}

func (s *DetectFaceAttributesResponse) SetData(v *DetectFaceAttributesResponseData) *DetectFaceAttributesResponse {
	s.Data = v
	return s
}

type DetectFaceAttributesResponseData struct {
	ImgHeight *int                                       `json:"ImgHeight,omitempty" xml:"ImgHeight,omitempty" require:"true"`
	ImgWidth  *int                                       `json:"ImgWidth,omitempty" xml:"ImgWidth,omitempty" require:"true"`
	FaceInfos *DetectFaceAttributesResponseDataFaceInfos `json:"FaceInfos,omitempty" xml:"FaceInfos,omitempty" require:"true" type:"Struct"`
}

func (s DetectFaceAttributesResponseData) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponseData) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseData) SetImgHeight(v int) *DetectFaceAttributesResponseData {
	s.ImgHeight = &v
	return s
}

func (s *DetectFaceAttributesResponseData) SetImgWidth(v int) *DetectFaceAttributesResponseData {
	s.ImgWidth = &v
	return s
}

func (s *DetectFaceAttributesResponseData) SetFaceInfos(v *DetectFaceAttributesResponseDataFaceInfos) *DetectFaceAttributesResponseData {
	s.FaceInfos = v
	return s
}

type DetectFaceAttributesResponseDataFaceInfos struct {
	FaceAttributesDetectInfo []*DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfo `json:"FaceAttributesDetectInfo,omitempty" xml:"FaceAttributesDetectInfo,omitempty" require:"true" type:"Repeated"`
}

func (s DetectFaceAttributesResponseDataFaceInfos) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponseDataFaceInfos) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseDataFaceInfos) SetFaceAttributesDetectInfo(v []*DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfo) *DetectFaceAttributesResponseDataFaceInfos {
	s.FaceAttributesDetectInfo = v
	return s
}

type DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfo struct {
	FaceAttributes *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes `json:"FaceAttributes,omitempty" xml:"FaceAttributes,omitempty" require:"true" type:"Struct"`
	FaceRect       *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceRect       `json:"FaceRect,omitempty" xml:"FaceRect,omitempty" require:"true" type:"Struct"`
}

func (s DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfo) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfo) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfo) SetFaceAttributes(v *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfo {
	s.FaceAttributes = v
	return s
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfo) SetFaceRect(v *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceRect) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfo {
	s.FaceRect = v
	return s
}

type DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes struct {
	Blur       *float32                                                                                 `json:"Blur,omitempty" xml:"Blur,omitempty" require:"true"`
	Facequal   *float32                                                                                 `json:"Facequal,omitempty" xml:"Facequal,omitempty" require:"true"`
	Facetype   *string                                                                                  `json:"Facetype,omitempty" xml:"Facetype,omitempty" require:"true"`
	Glasses    *string                                                                                  `json:"Glasses,omitempty" xml:"Glasses,omitempty" require:"true"`
	Integrity  *int                                                                                     `json:"Integrity,omitempty" xml:"Integrity,omitempty" require:"true"`
	Respirator *string                                                                                  `json:"Respirator,omitempty" xml:"Respirator,omitempty" require:"true"`
	Headpose   *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose `json:"Headpose,omitempty" xml:"Headpose,omitempty" require:"true" type:"Struct"`
	Smiling    *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling  `json:"Smiling,omitempty" xml:"Smiling,omitempty" require:"true" type:"Struct"`
}

func (s DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetBlur(v float32) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Blur = &v
	return s
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetFacequal(v float32) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Facequal = &v
	return s
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetFacetype(v string) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Facetype = &v
	return s
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetGlasses(v string) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Glasses = &v
	return s
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetIntegrity(v int) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Integrity = &v
	return s
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetRespirator(v string) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Respirator = &v
	return s
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetHeadpose(v *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Headpose = v
	return s
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetSmiling(v *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Smiling = v
	return s
}

type DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose struct {
	PitchAngle *float32 `json:"PitchAngle,omitempty" xml:"PitchAngle,omitempty" require:"true"`
	RollAngle  *float32 `json:"RollAngle,omitempty" xml:"RollAngle,omitempty" require:"true"`
	YawAngle   *float32 `json:"YawAngle,omitempty" xml:"YawAngle,omitempty" require:"true"`
}

func (s DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) SetPitchAngle(v float32) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose {
	s.PitchAngle = &v
	return s
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) SetRollAngle(v float32) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose {
	s.RollAngle = &v
	return s
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) SetYawAngle(v float32) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose {
	s.YawAngle = &v
	return s
}

type DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling struct {
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty" require:"true"`
	Value     *float32 `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
}

func (s DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling) SetThreshold(v float32) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling {
	s.Threshold = &v
	return s
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling) SetValue(v float32) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling {
	s.Value = &v
	return s
}

type DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceRect struct {
	Height *int `json:"Height,omitempty" xml:"Height,omitempty" require:"true"`
	Left   *int `json:"Left,omitempty" xml:"Left,omitempty" require:"true"`
	Top    *int `json:"Top,omitempty" xml:"Top,omitempty" require:"true"`
	Width  *int `json:"Width,omitempty" xml:"Width,omitempty" require:"true"`
}

func (s DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceRect) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceRect) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceRect) SetHeight(v int) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceRect {
	s.Height = &v
	return s
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceRect) SetLeft(v int) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceRect {
	s.Left = &v
	return s
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceRect) SetTop(v int) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceRect {
	s.Top = &v
	return s
}

func (s *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceRect) SetWidth(v int) *DetectFaceAttributesResponseDataFaceInfosFaceAttributesDetectInfoFaceRect {
	s.Width = &v
	return s
}

type InitFaceVerifyRequest struct {
	AuthId                 *string `json:"AuthId,omitempty" xml:"AuthId,omitempty"`
	CallbackToken          *string `json:"CallbackToken,omitempty" xml:"CallbackToken,omitempty"`
	CallbackUrl            *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	CertName               *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertNo                 *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	CertType               *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	CertifyId              *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	CertifyUrlType         *string `json:"CertifyUrlType,omitempty" xml:"CertifyUrlType,omitempty"`
	Crop                   *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	FaceContrastPicture    *string `json:"FaceContrastPicture,omitempty" xml:"FaceContrastPicture,omitempty"`
	FaceContrastPictureUrl *string `json:"FaceContrastPictureUrl,omitempty" xml:"FaceContrastPictureUrl,omitempty"`
	Ip                     *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	MetaInfo               *string `json:"MetaInfo,omitempty" xml:"MetaInfo,omitempty"`
	Mobile                 *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Model                  *string `json:"Model,omitempty" xml:"Model,omitempty"`
	OssBucketName          *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	OssObjectName          *string `json:"OssObjectName,omitempty" xml:"OssObjectName,omitempty"`
	OuterOrderNo           *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	ProductCode            *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ReturnUrl              *string `json:"ReturnUrl,omitempty" xml:"ReturnUrl,omitempty"`
	SceneId                *int64  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	UserId                 *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s InitFaceVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s InitFaceVerifyRequest) GoString() string {
	return s.String()
}

func (s *InitFaceVerifyRequest) SetAuthId(v string) *InitFaceVerifyRequest {
	s.AuthId = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCallbackToken(v string) *InitFaceVerifyRequest {
	s.CallbackToken = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCallbackUrl(v string) *InitFaceVerifyRequest {
	s.CallbackUrl = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCertName(v string) *InitFaceVerifyRequest {
	s.CertName = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCertNo(v string) *InitFaceVerifyRequest {
	s.CertNo = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCertType(v string) *InitFaceVerifyRequest {
	s.CertType = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCertifyId(v string) *InitFaceVerifyRequest {
	s.CertifyId = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCertifyUrlType(v string) *InitFaceVerifyRequest {
	s.CertifyUrlType = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCrop(v string) *InitFaceVerifyRequest {
	s.Crop = &v
	return s
}

func (s *InitFaceVerifyRequest) SetFaceContrastPicture(v string) *InitFaceVerifyRequest {
	s.FaceContrastPicture = &v
	return s
}

func (s *InitFaceVerifyRequest) SetFaceContrastPictureUrl(v string) *InitFaceVerifyRequest {
	s.FaceContrastPictureUrl = &v
	return s
}

func (s *InitFaceVerifyRequest) SetIp(v string) *InitFaceVerifyRequest {
	s.Ip = &v
	return s
}

func (s *InitFaceVerifyRequest) SetMetaInfo(v string) *InitFaceVerifyRequest {
	s.MetaInfo = &v
	return s
}

func (s *InitFaceVerifyRequest) SetMobile(v string) *InitFaceVerifyRequest {
	s.Mobile = &v
	return s
}

func (s *InitFaceVerifyRequest) SetModel(v string) *InitFaceVerifyRequest {
	s.Model = &v
	return s
}

func (s *InitFaceVerifyRequest) SetOssBucketName(v string) *InitFaceVerifyRequest {
	s.OssBucketName = &v
	return s
}

func (s *InitFaceVerifyRequest) SetOssObjectName(v string) *InitFaceVerifyRequest {
	s.OssObjectName = &v
	return s
}

func (s *InitFaceVerifyRequest) SetOuterOrderNo(v string) *InitFaceVerifyRequest {
	s.OuterOrderNo = &v
	return s
}

func (s *InitFaceVerifyRequest) SetProductCode(v string) *InitFaceVerifyRequest {
	s.ProductCode = &v
	return s
}

func (s *InitFaceVerifyRequest) SetReturnUrl(v string) *InitFaceVerifyRequest {
	s.ReturnUrl = &v
	return s
}

func (s *InitFaceVerifyRequest) SetSceneId(v int64) *InitFaceVerifyRequest {
	s.SceneId = &v
	return s
}

func (s *InitFaceVerifyRequest) SetUserId(v string) *InitFaceVerifyRequest {
	s.UserId = &v
	return s
}

type InitFaceVerifyResponse struct {
	Code         *string                             `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message      *string                             `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ResultObject *InitFaceVerifyResponseResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" require:"true" type:"Struct"`
}

func (s InitFaceVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s InitFaceVerifyResponse) GoString() string {
	return s.String()
}

func (s *InitFaceVerifyResponse) SetCode(v string) *InitFaceVerifyResponse {
	s.Code = &v
	return s
}

func (s *InitFaceVerifyResponse) SetMessage(v string) *InitFaceVerifyResponse {
	s.Message = &v
	return s
}

func (s *InitFaceVerifyResponse) SetRequestId(v string) *InitFaceVerifyResponse {
	s.RequestId = &v
	return s
}

func (s *InitFaceVerifyResponse) SetResultObject(v *InitFaceVerifyResponseResultObject) *InitFaceVerifyResponse {
	s.ResultObject = v
	return s
}

type InitFaceVerifyResponseResultObject struct {
	CertifyId  *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty" require:"true"`
	CertifyUrl *string `json:"CertifyUrl,omitempty" xml:"CertifyUrl,omitempty" require:"true"`
}

func (s InitFaceVerifyResponseResultObject) String() string {
	return tea.Prettify(s)
}

func (s InitFaceVerifyResponseResultObject) GoString() string {
	return s.String()
}

func (s *InitFaceVerifyResponseResultObject) SetCertifyId(v string) *InitFaceVerifyResponseResultObject {
	s.CertifyId = &v
	return s
}

func (s *InitFaceVerifyResponseResultObject) SetCertifyUrl(v string) *InitFaceVerifyResponseResultObject {
	s.CertifyUrl = &v
	return s
}

type LivenessFaceVerifyRequest struct {
	CertifyId              *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	Crop                   *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	DeviceToken            *string `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
	FaceContrastPicture    *string `json:"FaceContrastPicture,omitempty" xml:"FaceContrastPicture,omitempty"`
	FaceContrastPictureUrl *string `json:"FaceContrastPictureUrl,omitempty" xml:"FaceContrastPictureUrl,omitempty"`
	Ip                     *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Mobile                 *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Model                  *string `json:"Model,omitempty" xml:"Model,omitempty"`
	OssBucketName          *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	OssObjectName          *string `json:"OssObjectName,omitempty" xml:"OssObjectName,omitempty"`
	OuterOrderNo           *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	ProductCode            *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	SceneId                *int64  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	UserId                 *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s LivenessFaceVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s LivenessFaceVerifyRequest) GoString() string {
	return s.String()
}

func (s *LivenessFaceVerifyRequest) SetCertifyId(v string) *LivenessFaceVerifyRequest {
	s.CertifyId = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetCrop(v string) *LivenessFaceVerifyRequest {
	s.Crop = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetDeviceToken(v string) *LivenessFaceVerifyRequest {
	s.DeviceToken = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetFaceContrastPicture(v string) *LivenessFaceVerifyRequest {
	s.FaceContrastPicture = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetFaceContrastPictureUrl(v string) *LivenessFaceVerifyRequest {
	s.FaceContrastPictureUrl = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetIp(v string) *LivenessFaceVerifyRequest {
	s.Ip = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetMobile(v string) *LivenessFaceVerifyRequest {
	s.Mobile = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetModel(v string) *LivenessFaceVerifyRequest {
	s.Model = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetOssBucketName(v string) *LivenessFaceVerifyRequest {
	s.OssBucketName = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetOssObjectName(v string) *LivenessFaceVerifyRequest {
	s.OssObjectName = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetOuterOrderNo(v string) *LivenessFaceVerifyRequest {
	s.OuterOrderNo = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetProductCode(v string) *LivenessFaceVerifyRequest {
	s.ProductCode = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetSceneId(v int64) *LivenessFaceVerifyRequest {
	s.SceneId = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetUserId(v string) *LivenessFaceVerifyRequest {
	s.UserId = &v
	return s
}

type LivenessFaceVerifyResponse struct {
	Code         *string                                 `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message      *string                                 `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ResultObject *LivenessFaceVerifyResponseResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" require:"true" type:"Struct"`
}

func (s LivenessFaceVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s LivenessFaceVerifyResponse) GoString() string {
	return s.String()
}

func (s *LivenessFaceVerifyResponse) SetCode(v string) *LivenessFaceVerifyResponse {
	s.Code = &v
	return s
}

func (s *LivenessFaceVerifyResponse) SetMessage(v string) *LivenessFaceVerifyResponse {
	s.Message = &v
	return s
}

func (s *LivenessFaceVerifyResponse) SetRequestId(v string) *LivenessFaceVerifyResponse {
	s.RequestId = &v
	return s
}

func (s *LivenessFaceVerifyResponse) SetResultObject(v *LivenessFaceVerifyResponseResultObject) *LivenessFaceVerifyResponse {
	s.ResultObject = v
	return s
}

type LivenessFaceVerifyResponseResultObject struct {
	CertifyId    *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty" require:"true"`
	MaterialInfo *string `json:"MaterialInfo,omitempty" xml:"MaterialInfo,omitempty" require:"true"`
	Passed       *string `json:"Passed,omitempty" xml:"Passed,omitempty" require:"true"`
	SubCode      *string `json:"SubCode,omitempty" xml:"SubCode,omitempty" require:"true"`
}

func (s LivenessFaceVerifyResponseResultObject) String() string {
	return tea.Prettify(s)
}

func (s LivenessFaceVerifyResponseResultObject) GoString() string {
	return s.String()
}

func (s *LivenessFaceVerifyResponseResultObject) SetCertifyId(v string) *LivenessFaceVerifyResponseResultObject {
	s.CertifyId = &v
	return s
}

func (s *LivenessFaceVerifyResponseResultObject) SetMaterialInfo(v string) *LivenessFaceVerifyResponseResultObject {
	s.MaterialInfo = &v
	return s
}

func (s *LivenessFaceVerifyResponseResultObject) SetPassed(v string) *LivenessFaceVerifyResponseResultObject {
	s.Passed = &v
	return s
}

func (s *LivenessFaceVerifyResponseResultObject) SetSubCode(v string) *LivenessFaceVerifyResponseResultObject {
	s.SubCode = &v
	return s
}

type ModifyDeviceInfoRequest struct {
	BizType      *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	DeviceId     *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty" require:"true"`
	Duration     *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	ExpiredDay   *string `json:"ExpiredDay,omitempty" xml:"ExpiredDay,omitempty"`
	UserDeviceId *string `json:"UserDeviceId,omitempty" xml:"UserDeviceId,omitempty"`
}

func (s ModifyDeviceInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDeviceInfoRequest) GoString() string {
	return s.String()
}

func (s *ModifyDeviceInfoRequest) SetBizType(v string) *ModifyDeviceInfoRequest {
	s.BizType = &v
	return s
}

func (s *ModifyDeviceInfoRequest) SetDeviceId(v string) *ModifyDeviceInfoRequest {
	s.DeviceId = &v
	return s
}

func (s *ModifyDeviceInfoRequest) SetDuration(v string) *ModifyDeviceInfoRequest {
	s.Duration = &v
	return s
}

func (s *ModifyDeviceInfoRequest) SetExpiredDay(v string) *ModifyDeviceInfoRequest {
	s.ExpiredDay = &v
	return s
}

func (s *ModifyDeviceInfoRequest) SetUserDeviceId(v string) *ModifyDeviceInfoRequest {
	s.UserDeviceId = &v
	return s
}

type ModifyDeviceInfoResponse struct {
	BeginDay     *string `json:"BeginDay,omitempty" xml:"BeginDay,omitempty" require:"true"`
	BizType      *string `json:"BizType,omitempty" xml:"BizType,omitempty" require:"true"`
	DeviceId     *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty" require:"true"`
	ExpiredDay   *string `json:"ExpiredDay,omitempty" xml:"ExpiredDay,omitempty" require:"true"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	UserDeviceId *string `json:"UserDeviceId,omitempty" xml:"UserDeviceId,omitempty" require:"true"`
}

func (s ModifyDeviceInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDeviceInfoResponse) GoString() string {
	return s.String()
}

func (s *ModifyDeviceInfoResponse) SetBeginDay(v string) *ModifyDeviceInfoResponse {
	s.BeginDay = &v
	return s
}

func (s *ModifyDeviceInfoResponse) SetBizType(v string) *ModifyDeviceInfoResponse {
	s.BizType = &v
	return s
}

func (s *ModifyDeviceInfoResponse) SetDeviceId(v string) *ModifyDeviceInfoResponse {
	s.DeviceId = &v
	return s
}

func (s *ModifyDeviceInfoResponse) SetExpiredDay(v string) *ModifyDeviceInfoResponse {
	s.ExpiredDay = &v
	return s
}

func (s *ModifyDeviceInfoResponse) SetRequestId(v string) *ModifyDeviceInfoResponse {
	s.RequestId = &v
	return s
}

func (s *ModifyDeviceInfoResponse) SetUserDeviceId(v string) *ModifyDeviceInfoResponse {
	s.UserDeviceId = &v
	return s
}

type VerifyMaterialRequest struct {
	BizId               *string `json:"BizId,omitempty" xml:"BizId,omitempty" require:"true"`
	BizType             *string `json:"BizType,omitempty" xml:"BizType,omitempty" require:"true"`
	FaceImageUrl        *string `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty" require:"true"`
	IdCardBackImageUrl  *string `json:"IdCardBackImageUrl,omitempty" xml:"IdCardBackImageUrl,omitempty"`
	IdCardFrontImageUrl *string `json:"IdCardFrontImageUrl,omitempty" xml:"IdCardFrontImageUrl,omitempty"`
	IdCardNumber        *string `json:"IdCardNumber,omitempty" xml:"IdCardNumber,omitempty" require:"true"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	UserId              *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s VerifyMaterialRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyMaterialRequest) GoString() string {
	return s.String()
}

func (s *VerifyMaterialRequest) SetBizId(v string) *VerifyMaterialRequest {
	s.BizId = &v
	return s
}

func (s *VerifyMaterialRequest) SetBizType(v string) *VerifyMaterialRequest {
	s.BizType = &v
	return s
}

func (s *VerifyMaterialRequest) SetFaceImageUrl(v string) *VerifyMaterialRequest {
	s.FaceImageUrl = &v
	return s
}

func (s *VerifyMaterialRequest) SetIdCardBackImageUrl(v string) *VerifyMaterialRequest {
	s.IdCardBackImageUrl = &v
	return s
}

func (s *VerifyMaterialRequest) SetIdCardFrontImageUrl(v string) *VerifyMaterialRequest {
	s.IdCardFrontImageUrl = &v
	return s
}

func (s *VerifyMaterialRequest) SetIdCardNumber(v string) *VerifyMaterialRequest {
	s.IdCardNumber = &v
	return s
}

func (s *VerifyMaterialRequest) SetName(v string) *VerifyMaterialRequest {
	s.Name = &v
	return s
}

func (s *VerifyMaterialRequest) SetUserId(v string) *VerifyMaterialRequest {
	s.UserId = &v
	return s
}

type VerifyMaterialResponse struct {
	AuthorityComparisionScore *float32                        `json:"AuthorityComparisionScore,omitempty" xml:"AuthorityComparisionScore,omitempty" require:"true"`
	IdCardFaceComparisonScore *float32                        `json:"IdCardFaceComparisonScore,omitempty" xml:"IdCardFaceComparisonScore,omitempty" require:"true"`
	RequestId                 *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	VerifyStatus              *int                            `json:"VerifyStatus,omitempty" xml:"VerifyStatus,omitempty" require:"true"`
	VerifyToken               *string                         `json:"VerifyToken,omitempty" xml:"VerifyToken,omitempty" require:"true"`
	Material                  *VerifyMaterialResponseMaterial `json:"Material,omitempty" xml:"Material,omitempty" require:"true" type:"Struct"`
}

func (s VerifyMaterialResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyMaterialResponse) GoString() string {
	return s.String()
}

func (s *VerifyMaterialResponse) SetAuthorityComparisionScore(v float32) *VerifyMaterialResponse {
	s.AuthorityComparisionScore = &v
	return s
}

func (s *VerifyMaterialResponse) SetIdCardFaceComparisonScore(v float32) *VerifyMaterialResponse {
	s.IdCardFaceComparisonScore = &v
	return s
}

func (s *VerifyMaterialResponse) SetRequestId(v string) *VerifyMaterialResponse {
	s.RequestId = &v
	return s
}

func (s *VerifyMaterialResponse) SetVerifyStatus(v int) *VerifyMaterialResponse {
	s.VerifyStatus = &v
	return s
}

func (s *VerifyMaterialResponse) SetVerifyToken(v string) *VerifyMaterialResponse {
	s.VerifyToken = &v
	return s
}

func (s *VerifyMaterialResponse) SetMaterial(v *VerifyMaterialResponseMaterial) *VerifyMaterialResponse {
	s.Material = v
	return s
}

type VerifyMaterialResponseMaterial struct {
	FaceGlobalUrl *string                                   `json:"FaceGlobalUrl,omitempty" xml:"FaceGlobalUrl,omitempty" require:"true"`
	FaceImageUrl  *string                                   `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty" require:"true"`
	FaceMask      *string                                   `json:"FaceMask,omitempty" xml:"FaceMask,omitempty" require:"true"`
	FaceQuality   *string                                   `json:"FaceQuality,omitempty" xml:"FaceQuality,omitempty" require:"true"`
	IdCardName    *string                                   `json:"IdCardName,omitempty" xml:"IdCardName,omitempty" require:"true"`
	IdCardNumber  *string                                   `json:"IdCardNumber,omitempty" xml:"IdCardNumber,omitempty" require:"true"`
	IdCardInfo    *VerifyMaterialResponseMaterialIdCardInfo `json:"IdCardInfo,omitempty" xml:"IdCardInfo,omitempty" require:"true" type:"Struct"`
}

func (s VerifyMaterialResponseMaterial) String() string {
	return tea.Prettify(s)
}

func (s VerifyMaterialResponseMaterial) GoString() string {
	return s.String()
}

func (s *VerifyMaterialResponseMaterial) SetFaceGlobalUrl(v string) *VerifyMaterialResponseMaterial {
	s.FaceGlobalUrl = &v
	return s
}

func (s *VerifyMaterialResponseMaterial) SetFaceImageUrl(v string) *VerifyMaterialResponseMaterial {
	s.FaceImageUrl = &v
	return s
}

func (s *VerifyMaterialResponseMaterial) SetFaceMask(v string) *VerifyMaterialResponseMaterial {
	s.FaceMask = &v
	return s
}

func (s *VerifyMaterialResponseMaterial) SetFaceQuality(v string) *VerifyMaterialResponseMaterial {
	s.FaceQuality = &v
	return s
}

func (s *VerifyMaterialResponseMaterial) SetIdCardName(v string) *VerifyMaterialResponseMaterial {
	s.IdCardName = &v
	return s
}

func (s *VerifyMaterialResponseMaterial) SetIdCardNumber(v string) *VerifyMaterialResponseMaterial {
	s.IdCardNumber = &v
	return s
}

func (s *VerifyMaterialResponseMaterial) SetIdCardInfo(v *VerifyMaterialResponseMaterialIdCardInfo) *VerifyMaterialResponseMaterial {
	s.IdCardInfo = v
	return s
}

type VerifyMaterialResponseMaterialIdCardInfo struct {
	Address       *string `json:"Address,omitempty" xml:"Address,omitempty" require:"true"`
	Authority     *string `json:"Authority,omitempty" xml:"Authority,omitempty" require:"true"`
	BackImageUrl  *string `json:"BackImageUrl,omitempty" xml:"BackImageUrl,omitempty" require:"true"`
	Birth         *string `json:"Birth,omitempty" xml:"Birth,omitempty" require:"true"`
	EndDate       *string `json:"EndDate,omitempty" xml:"EndDate,omitempty" require:"true"`
	FrontImageUrl *string `json:"FrontImageUrl,omitempty" xml:"FrontImageUrl,omitempty" require:"true"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Nationality   *string `json:"Nationality,omitempty" xml:"Nationality,omitempty" require:"true"`
	Number        *string `json:"Number,omitempty" xml:"Number,omitempty" require:"true"`
	StartDate     *string `json:"StartDate,omitempty" xml:"StartDate,omitempty" require:"true"`
}

func (s VerifyMaterialResponseMaterialIdCardInfo) String() string {
	return tea.Prettify(s)
}

func (s VerifyMaterialResponseMaterialIdCardInfo) GoString() string {
	return s.String()
}

func (s *VerifyMaterialResponseMaterialIdCardInfo) SetAddress(v string) *VerifyMaterialResponseMaterialIdCardInfo {
	s.Address = &v
	return s
}

func (s *VerifyMaterialResponseMaterialIdCardInfo) SetAuthority(v string) *VerifyMaterialResponseMaterialIdCardInfo {
	s.Authority = &v
	return s
}

func (s *VerifyMaterialResponseMaterialIdCardInfo) SetBackImageUrl(v string) *VerifyMaterialResponseMaterialIdCardInfo {
	s.BackImageUrl = &v
	return s
}

func (s *VerifyMaterialResponseMaterialIdCardInfo) SetBirth(v string) *VerifyMaterialResponseMaterialIdCardInfo {
	s.Birth = &v
	return s
}

func (s *VerifyMaterialResponseMaterialIdCardInfo) SetEndDate(v string) *VerifyMaterialResponseMaterialIdCardInfo {
	s.EndDate = &v
	return s
}

func (s *VerifyMaterialResponseMaterialIdCardInfo) SetFrontImageUrl(v string) *VerifyMaterialResponseMaterialIdCardInfo {
	s.FrontImageUrl = &v
	return s
}

func (s *VerifyMaterialResponseMaterialIdCardInfo) SetName(v string) *VerifyMaterialResponseMaterialIdCardInfo {
	s.Name = &v
	return s
}

func (s *VerifyMaterialResponseMaterialIdCardInfo) SetNationality(v string) *VerifyMaterialResponseMaterialIdCardInfo {
	s.Nationality = &v
	return s
}

func (s *VerifyMaterialResponseMaterialIdCardInfo) SetNumber(v string) *VerifyMaterialResponseMaterialIdCardInfo {
	s.Number = &v
	return s
}

func (s *VerifyMaterialResponseMaterialIdCardInfo) SetStartDate(v string) *VerifyMaterialResponseMaterialIdCardInfo {
	s.StartDate = &v
	return s
}

type Client struct {
	rpc.Client
}

func NewClient(config *rpc.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *rpc.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("cloudauth"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) CompareFaceVerify(request *CompareFaceVerifyRequest, runtime *util.RuntimeOptions) (_result *CompareFaceVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CompareFaceVerifyResponse{}
	_body, _err := client.DoRequest(tea.String("CompareFaceVerify"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CompareFaceVerifySimply(request *CompareFaceVerifyRequest) (_result *CompareFaceVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CompareFaceVerifyResponse{}
	_body, _err := client.CompareFaceVerify(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CompareFaces(request *CompareFacesRequest, runtime *util.RuntimeOptions) (_result *CompareFacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CompareFacesResponse{}
	_body, _err := client.DoRequest(tea.String("CompareFaces"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CompareFacesSimply(request *CompareFacesRequest) (_result *CompareFacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CompareFacesResponse{}
	_body, _err := client.CompareFaces(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ContrastFaceVerify(request *ContrastFaceVerifyRequest, runtime *util.RuntimeOptions) (_result *ContrastFaceVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ContrastFaceVerifyResponse{}
	_body, _err := client.DoRequest(tea.String("ContrastFaceVerify"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ContrastFaceVerifySimply(request *ContrastFaceVerifyRequest) (_result *ContrastFaceVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ContrastFaceVerifyResponse{}
	_body, _err := client.ContrastFaceVerify(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ContrastFaceVerifyAdvance(request *ContrastFaceVerifyAdvanceRequest, runtime *util.RuntimeOptions) (_result *ContrastFaceVerifyResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	openPlatformEndpoint := client.OpenPlatformEndpoint
	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("Cloudauth"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	rpcutil.Convert(runtime, ossRuntime)
	contrastFaceVerifyReq := &ContrastFaceVerifyRequest{}
	rpcutil.Convert(request, contrastFaceVerifyReq)
	if !tea.BoolValue(util.IsUnset(request.FaceContrastFileObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.AccessKeyId
		ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.ObjectKey,
			Content:     request.FaceContrastFileObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.AccessKeyId,
			Policy:              authResponse.EncodedPolicy,
			Signature:           authResponse.Signature,
			Key:                 authResponse.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		contrastFaceVerifyReq.FaceContrastFile = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	}

	contrastFaceVerifyResp, _err := client.ContrastFaceVerify(contrastFaceVerifyReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = contrastFaceVerifyResp
	return _result, _err
}

func (client *Client) CreateAuthKey(request *CreateAuthKeyRequest, runtime *util.RuntimeOptions) (_result *CreateAuthKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateAuthKeyResponse{}
	_body, _err := client.DoRequest(tea.String("CreateAuthKey"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAuthKeySimply(request *CreateAuthKeyRequest) (_result *CreateAuthKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAuthKeyResponse{}
	_body, _err := client.CreateAuthKey(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVerifySetting(request *CreateVerifySettingRequest, runtime *util.RuntimeOptions) (_result *CreateVerifySettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateVerifySettingResponse{}
	_body, _err := client.DoRequest(tea.String("CreateVerifySetting"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVerifySettingSimply(request *CreateVerifySettingRequest) (_result *CreateVerifySettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVerifySettingResponse{}
	_body, _err := client.CreateVerifySetting(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDeviceInfo(request *DescribeDeviceInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeDeviceInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDeviceInfoResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDeviceInfo"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDeviceInfoSimply(request *DescribeDeviceInfoRequest) (_result *DescribeDeviceInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDeviceInfoResponse{}
	_body, _err := client.DescribeDeviceInfo(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFaceVerify(request *DescribeFaceVerifyRequest, runtime *util.RuntimeOptions) (_result *DescribeFaceVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeFaceVerifyResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeFaceVerify"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFaceVerifySimply(request *DescribeFaceVerifyRequest) (_result *DescribeFaceVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFaceVerifyResponse{}
	_body, _err := client.DescribeFaceVerify(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOssUploadToken(request *DescribeOssUploadTokenRequest, runtime *util.RuntimeOptions) (_result *DescribeOssUploadTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeOssUploadTokenResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeOssUploadToken"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOssUploadTokenSimply(request *DescribeOssUploadTokenRequest) (_result *DescribeOssUploadTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOssUploadTokenResponse{}
	_body, _err := client.DescribeOssUploadToken(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVerifyResult(request *DescribeVerifyResultRequest, runtime *util.RuntimeOptions) (_result *DescribeVerifyResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeVerifyResultResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeVerifyResult"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVerifyResultSimply(request *DescribeVerifyResultRequest) (_result *DescribeVerifyResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVerifyResultResponse{}
	_body, _err := client.DescribeVerifyResult(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVerifySDK(request *DescribeVerifySDKRequest, runtime *util.RuntimeOptions) (_result *DescribeVerifySDKResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeVerifySDKResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeVerifySDK"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVerifySDKSimply(request *DescribeVerifySDKRequest) (_result *DescribeVerifySDKResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVerifySDKResponse{}
	_body, _err := client.DescribeVerifySDK(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVerifyToken(request *DescribeVerifyTokenRequest, runtime *util.RuntimeOptions) (_result *DescribeVerifyTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeVerifyTokenResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeVerifyToken"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVerifyTokenSimply(request *DescribeVerifyTokenRequest) (_result *DescribeVerifyTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVerifyTokenResponse{}
	_body, _err := client.DescribeVerifyToken(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectFaceAttributes(request *DetectFaceAttributesRequest, runtime *util.RuntimeOptions) (_result *DetectFaceAttributesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DetectFaceAttributesResponse{}
	_body, _err := client.DoRequest(tea.String("DetectFaceAttributes"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectFaceAttributesSimply(request *DetectFaceAttributesRequest) (_result *DetectFaceAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectFaceAttributesResponse{}
	_body, _err := client.DetectFaceAttributes(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InitFaceVerify(request *InitFaceVerifyRequest, runtime *util.RuntimeOptions) (_result *InitFaceVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &InitFaceVerifyResponse{}
	_body, _err := client.DoRequest(tea.String("InitFaceVerify"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InitFaceVerifySimply(request *InitFaceVerifyRequest) (_result *InitFaceVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InitFaceVerifyResponse{}
	_body, _err := client.InitFaceVerify(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LivenessFaceVerify(request *LivenessFaceVerifyRequest, runtime *util.RuntimeOptions) (_result *LivenessFaceVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &LivenessFaceVerifyResponse{}
	_body, _err := client.DoRequest(tea.String("LivenessFaceVerify"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LivenessFaceVerifySimply(request *LivenessFaceVerifyRequest) (_result *LivenessFaceVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LivenessFaceVerifyResponse{}
	_body, _err := client.LivenessFaceVerify(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDeviceInfo(request *ModifyDeviceInfoRequest, runtime *util.RuntimeOptions) (_result *ModifyDeviceInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyDeviceInfoResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyDeviceInfo"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDeviceInfoSimply(request *ModifyDeviceInfoRequest) (_result *ModifyDeviceInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDeviceInfoResponse{}
	_body, _err := client.ModifyDeviceInfo(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyMaterial(request *VerifyMaterialRequest, runtime *util.RuntimeOptions) (_result *VerifyMaterialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &VerifyMaterialResponse{}
	_body, _err := client.DoRequest(tea.String("VerifyMaterial"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-07"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyMaterialSimply(request *VerifyMaterialRequest) (_result *VerifyMaterialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyMaterialResponse{}
	_body, _err := client.VerifyMaterial(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
