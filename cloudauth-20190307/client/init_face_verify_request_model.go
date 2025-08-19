// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitFaceVerifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppQualityCheck(v string) *InitFaceVerifyRequest
	GetAppQualityCheck() *string
	SetAuthId(v string) *InitFaceVerifyRequest
	GetAuthId() *string
	SetBirthday(v string) *InitFaceVerifyRequest
	GetBirthday() *string
	SetCallbackToken(v string) *InitFaceVerifyRequest
	GetCallbackToken() *string
	SetCallbackUrl(v string) *InitFaceVerifyRequest
	GetCallbackUrl() *string
	SetCameraSelection(v string) *InitFaceVerifyRequest
	GetCameraSelection() *string
	SetCertName(v string) *InitFaceVerifyRequest
	GetCertName() *string
	SetCertNo(v string) *InitFaceVerifyRequest
	GetCertNo() *string
	SetCertType(v string) *InitFaceVerifyRequest
	GetCertType() *string
	SetCertifyId(v string) *InitFaceVerifyRequest
	GetCertifyId() *string
	SetCertifyUrlStyle(v string) *InitFaceVerifyRequest
	GetCertifyUrlStyle() *string
	SetCertifyUrlType(v string) *InitFaceVerifyRequest
	GetCertifyUrlType() *string
	SetCrop(v string) *InitFaceVerifyRequest
	GetCrop() *string
	SetEncryptType(v string) *InitFaceVerifyRequest
	GetEncryptType() *string
	SetFaceContrastPicture(v string) *InitFaceVerifyRequest
	GetFaceContrastPicture() *string
	SetFaceContrastPictureUrl(v string) *InitFaceVerifyRequest
	GetFaceContrastPictureUrl() *string
	SetFaceGuardOutput(v string) *InitFaceVerifyRequest
	GetFaceGuardOutput() *string
	SetIp(v string) *InitFaceVerifyRequest
	GetIp() *string
	SetMetaInfo(v string) *InitFaceVerifyRequest
	GetMetaInfo() *string
	SetMobile(v string) *InitFaceVerifyRequest
	GetMobile() *string
	SetMode(v string) *InitFaceVerifyRequest
	GetMode() *string
	SetModel(v string) *InitFaceVerifyRequest
	GetModel() *string
	SetOssBucketName(v string) *InitFaceVerifyRequest
	GetOssBucketName() *string
	SetOssObjectName(v string) *InitFaceVerifyRequest
	GetOssObjectName() *string
	SetOuterOrderNo(v string) *InitFaceVerifyRequest
	GetOuterOrderNo() *string
	SetProcedurePriority(v string) *InitFaceVerifyRequest
	GetProcedurePriority() *string
	SetProductCode(v string) *InitFaceVerifyRequest
	GetProductCode() *string
	SetRarelyCharacters(v string) *InitFaceVerifyRequest
	GetRarelyCharacters() *string
	SetReadImg(v string) *InitFaceVerifyRequest
	GetReadImg() *string
	SetReturnUrl(v string) *InitFaceVerifyRequest
	GetReturnUrl() *string
	SetSceneId(v int64) *InitFaceVerifyRequest
	GetSceneId() *int64
	SetSuitableType(v string) *InitFaceVerifyRequest
	GetSuitableType() *string
	SetUiCustomUrl(v string) *InitFaceVerifyRequest
	GetUiCustomUrl() *string
	SetUserId(v string) *InitFaceVerifyRequest
	GetUserId() *string
	SetValidityDate(v string) *InitFaceVerifyRequest
	GetValidityDate() *string
	SetVideoEvidence(v string) *InitFaceVerifyRequest
	GetVideoEvidence() *string
	SetVoluntaryCustomizedContent(v string) *InitFaceVerifyRequest
	GetVoluntaryCustomizedContent() *string
}

type InitFaceVerifyRequest struct {
	AppQualityCheck *string `json:"AppQualityCheck,omitempty" xml:"AppQualityCheck,omitempty"`
	AuthId          *string `json:"AuthId,omitempty" xml:"AuthId,omitempty"`
	Birthday        *string `json:"Birthday,omitempty" xml:"Birthday,omitempty"`
	CallbackToken   *string `json:"CallbackToken,omitempty" xml:"CallbackToken,omitempty"`
	CallbackUrl     *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	CameraSelection *string `json:"CameraSelection,omitempty" xml:"CameraSelection,omitempty"`
	CertName        *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// example:
	//
	// 330103xxxxxxxxxxxx
	CertNo *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	// example:
	//
	// IDENTITY_CARD
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// example:
	//
	// 0bfa7c493f850e5178b9f8613634c9xx
	CertifyId       *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	CertifyUrlStyle *string `json:"CertifyUrlStyle,omitempty" xml:"CertifyUrlStyle,omitempty"`
	CertifyUrlType  *string `json:"CertifyUrlType,omitempty" xml:"CertifyUrlType,omitempty"`
	Crop            *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	EncryptType     *string `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	// example:
	//
	// /9j/4AAQSkZJRgABAQAASxxxxxxx
	FaceContrastPicture *string `json:"FaceContrastPicture,omitempty" xml:"FaceContrastPicture,omitempty"`
	// example:
	//
	// https://cn-shanghai-aliyun-cloudauth-xxxxxx.oss-cn-shanghai.aliyuncs.com/verify/xxxxx/xxxxx.jpeg
	FaceContrastPictureUrl *string `json:"FaceContrastPictureUrl,omitempty" xml:"FaceContrastPictureUrl,omitempty"`
	FaceGuardOutput        *string `json:"FaceGuardOutput,omitempty" xml:"FaceGuardOutput,omitempty"`
	// example:
	//
	// 114.xxx.xxx.xxx
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// example:
	//
	// {"zimVer":"3.0.0","appVersion": "1","bioMetaInfo": "4.1.0:11501568,0","appName": "com.aliyun.antcloudauth","deviceType": "ios","osVersion": "iOS 10.3.2","apdidToken": "","deviceModel": "iPhone9,1"}
	MetaInfo *string `json:"MetaInfo,omitempty" xml:"MetaInfo,omitempty"`
	// example:
	//
	// 130xxxxxxxx
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Mode   *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Model  *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// cn-shanghai-aliyun-cloudauth-xxxxx
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// example:
	//
	// verify/xxxxx/xxxxxx.jpeg
	OssObjectName *string `json:"OssObjectName,omitempty" xml:"OssObjectName,omitempty"`
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c353888
	OuterOrderNo      *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	ProcedurePriority *string `json:"ProcedurePriority,omitempty" xml:"ProcedurePriority,omitempty"`
	// example:
	//
	// ID_PRO
	ProductCode      *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	RarelyCharacters *string `json:"RarelyCharacters,omitempty" xml:"RarelyCharacters,omitempty"`
	ReadImg          *string `json:"ReadImg,omitempty" xml:"ReadImg,omitempty"`
	// example:
	//
	// www.aliyun.com
	ReturnUrl *string `json:"ReturnUrl,omitempty" xml:"ReturnUrl,omitempty"`
	// example:
	//
	// 1000000006
	SceneId      *int64  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	SuitableType *string `json:"SuitableType,omitempty" xml:"SuitableType,omitempty"`
	UiCustomUrl  *string `json:"UiCustomUrl,omitempty" xml:"UiCustomUrl,omitempty"`
	// example:
	//
	// 123456789
	UserId                     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ValidityDate               *string `json:"ValidityDate,omitempty" xml:"ValidityDate,omitempty"`
	VideoEvidence              *string `json:"VideoEvidence,omitempty" xml:"VideoEvidence,omitempty"`
	VoluntaryCustomizedContent *string `json:"VoluntaryCustomizedContent,omitempty" xml:"VoluntaryCustomizedContent,omitempty"`
}

func (s InitFaceVerifyRequest) String() string {
	return dara.Prettify(s)
}

func (s InitFaceVerifyRequest) GoString() string {
	return s.String()
}

func (s *InitFaceVerifyRequest) GetAppQualityCheck() *string {
	return s.AppQualityCheck
}

func (s *InitFaceVerifyRequest) GetAuthId() *string {
	return s.AuthId
}

func (s *InitFaceVerifyRequest) GetBirthday() *string {
	return s.Birthday
}

func (s *InitFaceVerifyRequest) GetCallbackToken() *string {
	return s.CallbackToken
}

func (s *InitFaceVerifyRequest) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *InitFaceVerifyRequest) GetCameraSelection() *string {
	return s.CameraSelection
}

func (s *InitFaceVerifyRequest) GetCertName() *string {
	return s.CertName
}

func (s *InitFaceVerifyRequest) GetCertNo() *string {
	return s.CertNo
}

func (s *InitFaceVerifyRequest) GetCertType() *string {
	return s.CertType
}

func (s *InitFaceVerifyRequest) GetCertifyId() *string {
	return s.CertifyId
}

func (s *InitFaceVerifyRequest) GetCertifyUrlStyle() *string {
	return s.CertifyUrlStyle
}

func (s *InitFaceVerifyRequest) GetCertifyUrlType() *string {
	return s.CertifyUrlType
}

func (s *InitFaceVerifyRequest) GetCrop() *string {
	return s.Crop
}

func (s *InitFaceVerifyRequest) GetEncryptType() *string {
	return s.EncryptType
}

func (s *InitFaceVerifyRequest) GetFaceContrastPicture() *string {
	return s.FaceContrastPicture
}

func (s *InitFaceVerifyRequest) GetFaceContrastPictureUrl() *string {
	return s.FaceContrastPictureUrl
}

func (s *InitFaceVerifyRequest) GetFaceGuardOutput() *string {
	return s.FaceGuardOutput
}

func (s *InitFaceVerifyRequest) GetIp() *string {
	return s.Ip
}

func (s *InitFaceVerifyRequest) GetMetaInfo() *string {
	return s.MetaInfo
}

func (s *InitFaceVerifyRequest) GetMobile() *string {
	return s.Mobile
}

func (s *InitFaceVerifyRequest) GetMode() *string {
	return s.Mode
}

func (s *InitFaceVerifyRequest) GetModel() *string {
	return s.Model
}

func (s *InitFaceVerifyRequest) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *InitFaceVerifyRequest) GetOssObjectName() *string {
	return s.OssObjectName
}

func (s *InitFaceVerifyRequest) GetOuterOrderNo() *string {
	return s.OuterOrderNo
}

func (s *InitFaceVerifyRequest) GetProcedurePriority() *string {
	return s.ProcedurePriority
}

func (s *InitFaceVerifyRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *InitFaceVerifyRequest) GetRarelyCharacters() *string {
	return s.RarelyCharacters
}

func (s *InitFaceVerifyRequest) GetReadImg() *string {
	return s.ReadImg
}

func (s *InitFaceVerifyRequest) GetReturnUrl() *string {
	return s.ReturnUrl
}

func (s *InitFaceVerifyRequest) GetSceneId() *int64 {
	return s.SceneId
}

func (s *InitFaceVerifyRequest) GetSuitableType() *string {
	return s.SuitableType
}

func (s *InitFaceVerifyRequest) GetUiCustomUrl() *string {
	return s.UiCustomUrl
}

func (s *InitFaceVerifyRequest) GetUserId() *string {
	return s.UserId
}

func (s *InitFaceVerifyRequest) GetValidityDate() *string {
	return s.ValidityDate
}

func (s *InitFaceVerifyRequest) GetVideoEvidence() *string {
	return s.VideoEvidence
}

func (s *InitFaceVerifyRequest) GetVoluntaryCustomizedContent() *string {
	return s.VoluntaryCustomizedContent
}

func (s *InitFaceVerifyRequest) SetAppQualityCheck(v string) *InitFaceVerifyRequest {
	s.AppQualityCheck = &v
	return s
}

func (s *InitFaceVerifyRequest) SetAuthId(v string) *InitFaceVerifyRequest {
	s.AuthId = &v
	return s
}

func (s *InitFaceVerifyRequest) SetBirthday(v string) *InitFaceVerifyRequest {
	s.Birthday = &v
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

func (s *InitFaceVerifyRequest) SetCameraSelection(v string) *InitFaceVerifyRequest {
	s.CameraSelection = &v
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

func (s *InitFaceVerifyRequest) SetCertifyUrlStyle(v string) *InitFaceVerifyRequest {
	s.CertifyUrlStyle = &v
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

func (s *InitFaceVerifyRequest) SetEncryptType(v string) *InitFaceVerifyRequest {
	s.EncryptType = &v
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

func (s *InitFaceVerifyRequest) SetFaceGuardOutput(v string) *InitFaceVerifyRequest {
	s.FaceGuardOutput = &v
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

func (s *InitFaceVerifyRequest) SetMode(v string) *InitFaceVerifyRequest {
	s.Mode = &v
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

func (s *InitFaceVerifyRequest) SetProcedurePriority(v string) *InitFaceVerifyRequest {
	s.ProcedurePriority = &v
	return s
}

func (s *InitFaceVerifyRequest) SetProductCode(v string) *InitFaceVerifyRequest {
	s.ProductCode = &v
	return s
}

func (s *InitFaceVerifyRequest) SetRarelyCharacters(v string) *InitFaceVerifyRequest {
	s.RarelyCharacters = &v
	return s
}

func (s *InitFaceVerifyRequest) SetReadImg(v string) *InitFaceVerifyRequest {
	s.ReadImg = &v
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

func (s *InitFaceVerifyRequest) SetSuitableType(v string) *InitFaceVerifyRequest {
	s.SuitableType = &v
	return s
}

func (s *InitFaceVerifyRequest) SetUiCustomUrl(v string) *InitFaceVerifyRequest {
	s.UiCustomUrl = &v
	return s
}

func (s *InitFaceVerifyRequest) SetUserId(v string) *InitFaceVerifyRequest {
	s.UserId = &v
	return s
}

func (s *InitFaceVerifyRequest) SetValidityDate(v string) *InitFaceVerifyRequest {
	s.ValidityDate = &v
	return s
}

func (s *InitFaceVerifyRequest) SetVideoEvidence(v string) *InitFaceVerifyRequest {
	s.VideoEvidence = &v
	return s
}

func (s *InitFaceVerifyRequest) SetVoluntaryCustomizedContent(v string) *InitFaceVerifyRequest {
	s.VoluntaryCustomizedContent = &v
	return s
}

func (s *InitFaceVerifyRequest) Validate() error {
	return dara.Validate(s)
}
