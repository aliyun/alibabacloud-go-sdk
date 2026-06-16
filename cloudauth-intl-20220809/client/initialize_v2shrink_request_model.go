// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitializeV2ShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppQualityCheck(v string) *InitializeV2ShrinkRequest
	GetAppQualityCheck() *string
	SetAuthorize(v string) *InitializeV2ShrinkRequest
	GetAuthorize() *string
	SetAutoRegistration(v string) *InitializeV2ShrinkRequest
	GetAutoRegistration() *string
	SetCallbackToken(v string) *InitializeV2ShrinkRequest
	GetCallbackToken() *string
	SetCallbackUrl(v string) *InitializeV2ShrinkRequest
	GetCallbackUrl() *string
	SetChameleonFrameEnable(v string) *InitializeV2ShrinkRequest
	GetChameleonFrameEnable() *string
	SetCrop(v string) *InitializeV2ShrinkRequest
	GetCrop() *string
	SetDateOfBirth(v string) *InitializeV2ShrinkRequest
	GetDateOfBirth() *string
	SetDateOfExpiry(v string) *InitializeV2ShrinkRequest
	GetDateOfExpiry() *string
	SetDocName(v string) *InitializeV2ShrinkRequest
	GetDocName() *string
	SetDocNo(v string) *InitializeV2ShrinkRequest
	GetDocNo() *string
	SetDocPageConfigShrink(v string) *InitializeV2ShrinkRequest
	GetDocPageConfigShrink() *string
	SetDocScanMode(v string) *InitializeV2ShrinkRequest
	GetDocScanMode() *string
	SetDocType(v string) *InitializeV2ShrinkRequest
	GetDocType() *string
	SetDocVideo(v string) *InitializeV2ShrinkRequest
	GetDocVideo() *string
	SetDocumentNumber(v string) *InitializeV2ShrinkRequest
	GetDocumentNumber() *string
	SetEditOcrResult(v string) *InitializeV2ShrinkRequest
	GetEditOcrResult() *string
	SetEmail(v string) *InitializeV2ShrinkRequest
	GetEmail() *string
	SetExperienceCode(v string) *InitializeV2ShrinkRequest
	GetExperienceCode() *string
	SetFaceGroupCodes(v string) *InitializeV2ShrinkRequest
	GetFaceGroupCodes() *string
	SetFacePictureBase64(v string) *InitializeV2ShrinkRequest
	GetFacePictureBase64() *string
	SetFacePictureFile(v string) *InitializeV2ShrinkRequest
	GetFacePictureFile() *string
	SetFacePictureUrl(v string) *InitializeV2ShrinkRequest
	GetFacePictureUrl() *string
	SetFaceRegisterGroupCode(v string) *InitializeV2ShrinkRequest
	GetFaceRegisterGroupCode() *string
	SetFaceVerifyThreshold(v string) *InitializeV2ShrinkRequest
	GetFaceVerifyThreshold() *string
	SetIdFaceQuality(v string) *InitializeV2ShrinkRequest
	GetIdFaceQuality() *string
	SetIdSpoof(v string) *InitializeV2ShrinkRequest
	GetIdSpoof() *string
	SetIdThreshold(v string) *InitializeV2ShrinkRequest
	GetIdThreshold() *string
	SetLanguageConfig(v string) *InitializeV2ShrinkRequest
	GetLanguageConfig() *string
	SetMRTDInput(v string) *InitializeV2ShrinkRequest
	GetMRTDInput() *string
	SetMerchantBizId(v string) *InitializeV2ShrinkRequest
	GetMerchantBizId() *string
	SetMerchantUserId(v string) *InitializeV2ShrinkRequest
	GetMerchantUserId() *string
	SetMetaInfo(v string) *InitializeV2ShrinkRequest
	GetMetaInfo() *string
	SetMobile(v string) *InitializeV2ShrinkRequest
	GetMobile() *string
	SetModel(v string) *InitializeV2ShrinkRequest
	GetModel() *string
	SetOcr(v string) *InitializeV2ShrinkRequest
	GetOcr() *string
	SetOcrValueStandard(v string) *InitializeV2ShrinkRequest
	GetOcrValueStandard() *string
	SetPages(v string) *InitializeV2ShrinkRequest
	GetPages() *string
	SetProcedurePriority(v string) *InitializeV2ShrinkRequest
	GetProcedurePriority() *string
	SetProductCode(v string) *InitializeV2ShrinkRequest
	GetProductCode() *string
	SetProductFlow(v string) *InitializeV2ShrinkRequest
	GetProductFlow() *string
	SetReturnFaces(v string) *InitializeV2ShrinkRequest
	GetReturnFaces() *string
	SetReturnUrl(v string) *InitializeV2ShrinkRequest
	GetReturnUrl() *string
	SetSaveFacePicture(v string) *InitializeV2ShrinkRequest
	GetSaveFacePicture() *string
	SetSceneCode(v string) *InitializeV2ShrinkRequest
	GetSceneCode() *string
	SetSecurityLevel(v string) *InitializeV2ShrinkRequest
	GetSecurityLevel() *string
	SetShowAlbumIcon(v string) *InitializeV2ShrinkRequest
	GetShowAlbumIcon() *string
	SetShowGuidePage(v string) *InitializeV2ShrinkRequest
	GetShowGuidePage() *string
	SetShowOcrResult(v string) *InitializeV2ShrinkRequest
	GetShowOcrResult() *string
	SetStyleConfig(v string) *InitializeV2ShrinkRequest
	GetStyleConfig() *string
	SetTargetFacePicture(v string) *InitializeV2ShrinkRequest
	GetTargetFacePicture() *string
	SetTargetFacePictureFile(v string) *InitializeV2ShrinkRequest
	GetTargetFacePictureFile() *string
	SetTargetFacePictureUrl(v string) *InitializeV2ShrinkRequest
	GetTargetFacePictureUrl() *string
	SetTemplateConfig(v string) *InitializeV2ShrinkRequest
	GetTemplateConfig() *string
	SetTemplateRanCount(v string) *InitializeV2ShrinkRequest
	GetTemplateRanCount() *string
	SetTemplateType(v string) *InitializeV2ShrinkRequest
	GetTemplateType() *string
	SetUseNFC(v string) *InitializeV2ShrinkRequest
	GetUseNFC() *string
	SetVerifyModel(v string) *InitializeV2ShrinkRequest
	GetVerifyModel() *string
}

type InitializeV2ShrinkRequest struct {
	// example:
	//
	// N
	AppQualityCheck *string `json:"AppQualityCheck,omitempty" xml:"AppQualityCheck,omitempty"`
	// example:
	//
	// T
	Authorize *string `json:"Authorize,omitempty" xml:"Authorize,omitempty"`
	// example:
	//
	// 0
	AutoRegistration *string `json:"AutoRegistration,omitempty" xml:"AutoRegistration,omitempty"`
	// example:
	//
	// 7ca5c68d869344ea8eeb30cdfd544544-6358700
	CallbackToken *string `json:"CallbackToken,omitempty" xml:"CallbackToken,omitempty"`
	// example:
	//
	// https://www.aliyun.com?callbackToken=1000004826&transactionId=shaxxxx&passed=Y&subCode=200
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// example:
	//
	// N
	ChameleonFrameEnable *string `json:"ChameleonFrameEnable,omitempty" xml:"ChameleonFrameEnable,omitempty"`
	// example:
	//
	// N
	Crop *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	// example:
	//
	// -
	DateOfBirth *string `json:"DateOfBirth,omitempty" xml:"DateOfBirth,omitempty"`
	// example:
	//
	// -
	DateOfExpiry *string `json:"DateOfExpiry,omitempty" xml:"DateOfExpiry,omitempty"`
	// example:
	//
	// 张**
	DocName *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
	// example:
	//
	// 410***************
	DocNo *string `json:"DocNo,omitempty" xml:"DocNo,omitempty"`
	// example:
	//
	// OCR_ID_BACK
	DocPageConfigShrink *string `json:"DocPageConfig,omitempty" xml:"DocPageConfig,omitempty"`
	// example:
	//
	// manual
	DocScanMode *string `json:"DocScanMode,omitempty" xml:"DocScanMode,omitempty"`
	// example:
	//
	// 00000001
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// example:
	//
	// N
	DocVideo *string `json:"DocVideo,omitempty" xml:"DocVideo,omitempty"`
	// example:
	//
	// -
	DocumentNumber *string `json:"DocumentNumber,omitempty" xml:"DocumentNumber,omitempty"`
	// example:
	//
	// 0
	EditOcrResult *string `json:"EditOcrResult,omitempty" xml:"EditOcrResult,omitempty"`
	// example:
	//
	// evxxx@imigxxxxx.go.id
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// 9be7b7d0180041219e5ab03ac6dab5fb
	ExperienceCode *string `json:"ExperienceCode,omitempty" xml:"ExperienceCode,omitempty"`
	// example:
	//
	// 0e0c34a77f
	FaceGroupCodes *string `json:"FaceGroupCodes,omitempty" xml:"FaceGroupCodes,omitempty"`
	// example:
	//
	// Base64
	FacePictureBase64 *string `json:"FacePictureBase64,omitempty" xml:"FacePictureBase64,omitempty"`
	// example:
	//
	// InputStream
	FacePictureFile *string `json:"FacePictureFile,omitempty" xml:"FacePictureFile,omitempty"`
	// example:
	//
	// https://digital-face-prod8.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
	FacePictureUrl *string `json:"FacePictureUrl,omitempty" xml:"FacePictureUrl,omitempty"`
	// example:
	//
	// 0e0c34a77f
	FaceRegisterGroupCode *string `json:"FaceRegisterGroupCode,omitempty" xml:"FaceRegisterGroupCode,omitempty"`
	// example:
	//
	// 0.5
	FaceVerifyThreshold *string `json:"FaceVerifyThreshold,omitempty" xml:"FaceVerifyThreshold,omitempty"`
	// example:
	//
	// Y
	IdFaceQuality *string `json:"IdFaceQuality,omitempty" xml:"IdFaceQuality,omitempty"`
	// example:
	//
	// Y
	IdSpoof *string `json:"IdSpoof,omitempty" xml:"IdSpoof,omitempty"`
	// example:
	//
	// 0
	IdThreshold *string `json:"IdThreshold,omitempty" xml:"IdThreshold,omitempty"`
	// example:
	//
	// en
	LanguageConfig *string `json:"LanguageConfig,omitempty" xml:"LanguageConfig,omitempty"`
	// example:
	//
	// 0
	MRTDInput *string `json:"MRTDInput,omitempty" xml:"MRTDInput,omitempty"`
	// example:
	//
	// e0c34a***353888
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// example:
	//
	// 1221****6543
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// example:
	//
	// {\\"bioMetaInfo\\":\\"4.1.0:2916352,0\\",\\"deviceType\\":\\"web\\",\\"ua\\":\\"Mozilla/5.0\\"}
	MetaInfo *string `json:"MetaInfo,omitempty" xml:"MetaInfo,omitempty"`
	// example:
	//
	// +6281293671234
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// example:
	//
	// PHOTINUS_LIVENESS
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// Y
	Ocr *string `json:"Ocr,omitempty" xml:"Ocr,omitempty"`
	// example:
	//
	// 0
	OcrValueStandard *string `json:"OcrValueStandard,omitempty" xml:"OcrValueStandard,omitempty"`
	// example:
	//
	// 01
	Pages *string `json:"Pages,omitempty" xml:"Pages,omitempty"`
	// example:
	//
	// url
	ProcedurePriority *string `json:"ProcedurePriority,omitempty" xml:"ProcedurePriority,omitempty"`
	// example:
	//
	// eKYC
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// DOC_FACE
	ProductFlow *string `json:"ProductFlow,omitempty" xml:"ProductFlow,omitempty"`
	// example:
	//
	// 1
	ReturnFaces *string `json:"ReturnFaces,omitempty" xml:"ReturnFaces,omitempty"`
	// example:
	//
	// http*****
	ReturnUrl *string `json:"ReturnUrl,omitempty" xml:"ReturnUrl,omitempty"`
	// example:
	//
	// 0
	SaveFacePicture *string `json:"SaveFacePicture,omitempty" xml:"SaveFacePicture,omitempty"`
	// example:
	//
	// 1234567890
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// example:
	//
	// 01
	SecurityLevel *string `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	// example:
	//
	// 1
	ShowAlbumIcon *string `json:"ShowAlbumIcon,omitempty" xml:"ShowAlbumIcon,omitempty"`
	// example:
	//
	// 1
	ShowGuidePage *string `json:"ShowGuidePage,omitempty" xml:"ShowGuidePage,omitempty"`
	// example:
	//
	// 1
	ShowOcrResult *string `json:"ShowOcrResult,omitempty" xml:"ShowOcrResult,omitempty"`
	// example:
	//
	// {
	//
	//   "guidepage:": {****},
	//
	//   "ocrPage": {****},
	//
	//   "ocrResultPage": [****],
	//
	//   "facePage": {****},
	//
	// }
	StyleConfig *string `json:"StyleConfig,omitempty" xml:"StyleConfig,omitempty"`
	// example:
	//
	// base64
	TargetFacePicture *string `json:"TargetFacePicture,omitempty" xml:"TargetFacePicture,omitempty"`
	// example:
	//
	// InputStream
	TargetFacePictureFile *string `json:"TargetFacePictureFile,omitempty" xml:"TargetFacePictureFile,omitempty"`
	// example:
	//
	// https://www.xxxxx.com/1.jpg
	TargetFacePictureUrl *string `json:"TargetFacePictureUrl,omitempty" xml:"TargetFacePictureUrl,omitempty"`
	// example:
	//
	// 01,02,07
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// example:
	//
	// 2
	TemplateRanCount *string `json:"TemplateRanCount,omitempty" xml:"TemplateRanCount,omitempty"`
	// example:
	//
	// Seq
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// example:
	//
	// N
	UseNFC *string `json:"UseNFC,omitempty" xml:"UseNFC,omitempty"`
	// example:
	//
	// 0
	VerifyModel *string `json:"VerifyModel,omitempty" xml:"VerifyModel,omitempty"`
}

func (s InitializeV2ShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s InitializeV2ShrinkRequest) GoString() string {
	return s.String()
}

func (s *InitializeV2ShrinkRequest) GetAppQualityCheck() *string {
	return s.AppQualityCheck
}

func (s *InitializeV2ShrinkRequest) GetAuthorize() *string {
	return s.Authorize
}

func (s *InitializeV2ShrinkRequest) GetAutoRegistration() *string {
	return s.AutoRegistration
}

func (s *InitializeV2ShrinkRequest) GetCallbackToken() *string {
	return s.CallbackToken
}

func (s *InitializeV2ShrinkRequest) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *InitializeV2ShrinkRequest) GetChameleonFrameEnable() *string {
	return s.ChameleonFrameEnable
}

func (s *InitializeV2ShrinkRequest) GetCrop() *string {
	return s.Crop
}

func (s *InitializeV2ShrinkRequest) GetDateOfBirth() *string {
	return s.DateOfBirth
}

func (s *InitializeV2ShrinkRequest) GetDateOfExpiry() *string {
	return s.DateOfExpiry
}

func (s *InitializeV2ShrinkRequest) GetDocName() *string {
	return s.DocName
}

func (s *InitializeV2ShrinkRequest) GetDocNo() *string {
	return s.DocNo
}

func (s *InitializeV2ShrinkRequest) GetDocPageConfigShrink() *string {
	return s.DocPageConfigShrink
}

func (s *InitializeV2ShrinkRequest) GetDocScanMode() *string {
	return s.DocScanMode
}

func (s *InitializeV2ShrinkRequest) GetDocType() *string {
	return s.DocType
}

func (s *InitializeV2ShrinkRequest) GetDocVideo() *string {
	return s.DocVideo
}

func (s *InitializeV2ShrinkRequest) GetDocumentNumber() *string {
	return s.DocumentNumber
}

func (s *InitializeV2ShrinkRequest) GetEditOcrResult() *string {
	return s.EditOcrResult
}

func (s *InitializeV2ShrinkRequest) GetEmail() *string {
	return s.Email
}

func (s *InitializeV2ShrinkRequest) GetExperienceCode() *string {
	return s.ExperienceCode
}

func (s *InitializeV2ShrinkRequest) GetFaceGroupCodes() *string {
	return s.FaceGroupCodes
}

func (s *InitializeV2ShrinkRequest) GetFacePictureBase64() *string {
	return s.FacePictureBase64
}

func (s *InitializeV2ShrinkRequest) GetFacePictureFile() *string {
	return s.FacePictureFile
}

func (s *InitializeV2ShrinkRequest) GetFacePictureUrl() *string {
	return s.FacePictureUrl
}

func (s *InitializeV2ShrinkRequest) GetFaceRegisterGroupCode() *string {
	return s.FaceRegisterGroupCode
}

func (s *InitializeV2ShrinkRequest) GetFaceVerifyThreshold() *string {
	return s.FaceVerifyThreshold
}

func (s *InitializeV2ShrinkRequest) GetIdFaceQuality() *string {
	return s.IdFaceQuality
}

func (s *InitializeV2ShrinkRequest) GetIdSpoof() *string {
	return s.IdSpoof
}

func (s *InitializeV2ShrinkRequest) GetIdThreshold() *string {
	return s.IdThreshold
}

func (s *InitializeV2ShrinkRequest) GetLanguageConfig() *string {
	return s.LanguageConfig
}

func (s *InitializeV2ShrinkRequest) GetMRTDInput() *string {
	return s.MRTDInput
}

func (s *InitializeV2ShrinkRequest) GetMerchantBizId() *string {
	return s.MerchantBizId
}

func (s *InitializeV2ShrinkRequest) GetMerchantUserId() *string {
	return s.MerchantUserId
}

func (s *InitializeV2ShrinkRequest) GetMetaInfo() *string {
	return s.MetaInfo
}

func (s *InitializeV2ShrinkRequest) GetMobile() *string {
	return s.Mobile
}

func (s *InitializeV2ShrinkRequest) GetModel() *string {
	return s.Model
}

func (s *InitializeV2ShrinkRequest) GetOcr() *string {
	return s.Ocr
}

func (s *InitializeV2ShrinkRequest) GetOcrValueStandard() *string {
	return s.OcrValueStandard
}

func (s *InitializeV2ShrinkRequest) GetPages() *string {
	return s.Pages
}

func (s *InitializeV2ShrinkRequest) GetProcedurePriority() *string {
	return s.ProcedurePriority
}

func (s *InitializeV2ShrinkRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *InitializeV2ShrinkRequest) GetProductFlow() *string {
	return s.ProductFlow
}

func (s *InitializeV2ShrinkRequest) GetReturnFaces() *string {
	return s.ReturnFaces
}

func (s *InitializeV2ShrinkRequest) GetReturnUrl() *string {
	return s.ReturnUrl
}

func (s *InitializeV2ShrinkRequest) GetSaveFacePicture() *string {
	return s.SaveFacePicture
}

func (s *InitializeV2ShrinkRequest) GetSceneCode() *string {
	return s.SceneCode
}

func (s *InitializeV2ShrinkRequest) GetSecurityLevel() *string {
	return s.SecurityLevel
}

func (s *InitializeV2ShrinkRequest) GetShowAlbumIcon() *string {
	return s.ShowAlbumIcon
}

func (s *InitializeV2ShrinkRequest) GetShowGuidePage() *string {
	return s.ShowGuidePage
}

func (s *InitializeV2ShrinkRequest) GetShowOcrResult() *string {
	return s.ShowOcrResult
}

func (s *InitializeV2ShrinkRequest) GetStyleConfig() *string {
	return s.StyleConfig
}

func (s *InitializeV2ShrinkRequest) GetTargetFacePicture() *string {
	return s.TargetFacePicture
}

func (s *InitializeV2ShrinkRequest) GetTargetFacePictureFile() *string {
	return s.TargetFacePictureFile
}

func (s *InitializeV2ShrinkRequest) GetTargetFacePictureUrl() *string {
	return s.TargetFacePictureUrl
}

func (s *InitializeV2ShrinkRequest) GetTemplateConfig() *string {
	return s.TemplateConfig
}

func (s *InitializeV2ShrinkRequest) GetTemplateRanCount() *string {
	return s.TemplateRanCount
}

func (s *InitializeV2ShrinkRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *InitializeV2ShrinkRequest) GetUseNFC() *string {
	return s.UseNFC
}

func (s *InitializeV2ShrinkRequest) GetVerifyModel() *string {
	return s.VerifyModel
}

func (s *InitializeV2ShrinkRequest) SetAppQualityCheck(v string) *InitializeV2ShrinkRequest {
	s.AppQualityCheck = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetAuthorize(v string) *InitializeV2ShrinkRequest {
	s.Authorize = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetAutoRegistration(v string) *InitializeV2ShrinkRequest {
	s.AutoRegistration = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetCallbackToken(v string) *InitializeV2ShrinkRequest {
	s.CallbackToken = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetCallbackUrl(v string) *InitializeV2ShrinkRequest {
	s.CallbackUrl = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetChameleonFrameEnable(v string) *InitializeV2ShrinkRequest {
	s.ChameleonFrameEnable = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetCrop(v string) *InitializeV2ShrinkRequest {
	s.Crop = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetDateOfBirth(v string) *InitializeV2ShrinkRequest {
	s.DateOfBirth = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetDateOfExpiry(v string) *InitializeV2ShrinkRequest {
	s.DateOfExpiry = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetDocName(v string) *InitializeV2ShrinkRequest {
	s.DocName = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetDocNo(v string) *InitializeV2ShrinkRequest {
	s.DocNo = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetDocPageConfigShrink(v string) *InitializeV2ShrinkRequest {
	s.DocPageConfigShrink = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetDocScanMode(v string) *InitializeV2ShrinkRequest {
	s.DocScanMode = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetDocType(v string) *InitializeV2ShrinkRequest {
	s.DocType = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetDocVideo(v string) *InitializeV2ShrinkRequest {
	s.DocVideo = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetDocumentNumber(v string) *InitializeV2ShrinkRequest {
	s.DocumentNumber = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetEditOcrResult(v string) *InitializeV2ShrinkRequest {
	s.EditOcrResult = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetEmail(v string) *InitializeV2ShrinkRequest {
	s.Email = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetExperienceCode(v string) *InitializeV2ShrinkRequest {
	s.ExperienceCode = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetFaceGroupCodes(v string) *InitializeV2ShrinkRequest {
	s.FaceGroupCodes = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetFacePictureBase64(v string) *InitializeV2ShrinkRequest {
	s.FacePictureBase64 = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetFacePictureFile(v string) *InitializeV2ShrinkRequest {
	s.FacePictureFile = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetFacePictureUrl(v string) *InitializeV2ShrinkRequest {
	s.FacePictureUrl = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetFaceRegisterGroupCode(v string) *InitializeV2ShrinkRequest {
	s.FaceRegisterGroupCode = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetFaceVerifyThreshold(v string) *InitializeV2ShrinkRequest {
	s.FaceVerifyThreshold = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetIdFaceQuality(v string) *InitializeV2ShrinkRequest {
	s.IdFaceQuality = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetIdSpoof(v string) *InitializeV2ShrinkRequest {
	s.IdSpoof = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetIdThreshold(v string) *InitializeV2ShrinkRequest {
	s.IdThreshold = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetLanguageConfig(v string) *InitializeV2ShrinkRequest {
	s.LanguageConfig = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetMRTDInput(v string) *InitializeV2ShrinkRequest {
	s.MRTDInput = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetMerchantBizId(v string) *InitializeV2ShrinkRequest {
	s.MerchantBizId = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetMerchantUserId(v string) *InitializeV2ShrinkRequest {
	s.MerchantUserId = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetMetaInfo(v string) *InitializeV2ShrinkRequest {
	s.MetaInfo = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetMobile(v string) *InitializeV2ShrinkRequest {
	s.Mobile = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetModel(v string) *InitializeV2ShrinkRequest {
	s.Model = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetOcr(v string) *InitializeV2ShrinkRequest {
	s.Ocr = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetOcrValueStandard(v string) *InitializeV2ShrinkRequest {
	s.OcrValueStandard = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetPages(v string) *InitializeV2ShrinkRequest {
	s.Pages = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetProcedurePriority(v string) *InitializeV2ShrinkRequest {
	s.ProcedurePriority = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetProductCode(v string) *InitializeV2ShrinkRequest {
	s.ProductCode = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetProductFlow(v string) *InitializeV2ShrinkRequest {
	s.ProductFlow = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetReturnFaces(v string) *InitializeV2ShrinkRequest {
	s.ReturnFaces = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetReturnUrl(v string) *InitializeV2ShrinkRequest {
	s.ReturnUrl = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetSaveFacePicture(v string) *InitializeV2ShrinkRequest {
	s.SaveFacePicture = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetSceneCode(v string) *InitializeV2ShrinkRequest {
	s.SceneCode = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetSecurityLevel(v string) *InitializeV2ShrinkRequest {
	s.SecurityLevel = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetShowAlbumIcon(v string) *InitializeV2ShrinkRequest {
	s.ShowAlbumIcon = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetShowGuidePage(v string) *InitializeV2ShrinkRequest {
	s.ShowGuidePage = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetShowOcrResult(v string) *InitializeV2ShrinkRequest {
	s.ShowOcrResult = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetStyleConfig(v string) *InitializeV2ShrinkRequest {
	s.StyleConfig = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetTargetFacePicture(v string) *InitializeV2ShrinkRequest {
	s.TargetFacePicture = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetTargetFacePictureFile(v string) *InitializeV2ShrinkRequest {
	s.TargetFacePictureFile = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetTargetFacePictureUrl(v string) *InitializeV2ShrinkRequest {
	s.TargetFacePictureUrl = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetTemplateConfig(v string) *InitializeV2ShrinkRequest {
	s.TemplateConfig = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetTemplateRanCount(v string) *InitializeV2ShrinkRequest {
	s.TemplateRanCount = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetTemplateType(v string) *InitializeV2ShrinkRequest {
	s.TemplateType = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetUseNFC(v string) *InitializeV2ShrinkRequest {
	s.UseNFC = &v
	return s
}

func (s *InitializeV2ShrinkRequest) SetVerifyModel(v string) *InitializeV2ShrinkRequest {
	s.VerifyModel = &v
	return s
}

func (s *InitializeV2ShrinkRequest) Validate() error {
	return dara.Validate(s)
}
