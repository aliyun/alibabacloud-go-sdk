// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iInitializeV2AdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppQualityCheck(v string) *InitializeV2AdvanceRequest
	GetAppQualityCheck() *string
	SetAuthorize(v string) *InitializeV2AdvanceRequest
	GetAuthorize() *string
	SetAutoRegistration(v string) *InitializeV2AdvanceRequest
	GetAutoRegistration() *string
	SetCallbackToken(v string) *InitializeV2AdvanceRequest
	GetCallbackToken() *string
	SetCallbackUrl(v string) *InitializeV2AdvanceRequest
	GetCallbackUrl() *string
	SetChameleonFrameEnable(v string) *InitializeV2AdvanceRequest
	GetChameleonFrameEnable() *string
	SetCrop(v string) *InitializeV2AdvanceRequest
	GetCrop() *string
	SetDateOfBirth(v string) *InitializeV2AdvanceRequest
	GetDateOfBirth() *string
	SetDateOfExpiry(v string) *InitializeV2AdvanceRequest
	GetDateOfExpiry() *string
	SetDocName(v string) *InitializeV2AdvanceRequest
	GetDocName() *string
	SetDocNo(v string) *InitializeV2AdvanceRequest
	GetDocNo() *string
	SetDocPageConfig(v []*string) *InitializeV2AdvanceRequest
	GetDocPageConfig() []*string
	SetDocScanMode(v string) *InitializeV2AdvanceRequest
	GetDocScanMode() *string
	SetDocType(v string) *InitializeV2AdvanceRequest
	GetDocType() *string
	SetDocVideo(v string) *InitializeV2AdvanceRequest
	GetDocVideo() *string
	SetDocumentNumber(v string) *InitializeV2AdvanceRequest
	GetDocumentNumber() *string
	SetEditOcrResult(v string) *InitializeV2AdvanceRequest
	GetEditOcrResult() *string
	SetEmail(v string) *InitializeV2AdvanceRequest
	GetEmail() *string
	SetExperienceCode(v string) *InitializeV2AdvanceRequest
	GetExperienceCode() *string
	SetFaceGroupCodes(v string) *InitializeV2AdvanceRequest
	GetFaceGroupCodes() *string
	SetFacePictureBase64(v string) *InitializeV2AdvanceRequest
	GetFacePictureBase64() *string
	SetFacePictureFileObject(v io.Reader) *InitializeV2AdvanceRequest
	GetFacePictureFileObject() io.Reader
	SetFacePictureUrl(v string) *InitializeV2AdvanceRequest
	GetFacePictureUrl() *string
	SetFaceRegisterGroupCode(v string) *InitializeV2AdvanceRequest
	GetFaceRegisterGroupCode() *string
	SetFaceVerifyThreshold(v string) *InitializeV2AdvanceRequest
	GetFaceVerifyThreshold() *string
	SetIdFaceQuality(v string) *InitializeV2AdvanceRequest
	GetIdFaceQuality() *string
	SetIdSpoof(v string) *InitializeV2AdvanceRequest
	GetIdSpoof() *string
	SetIdThreshold(v string) *InitializeV2AdvanceRequest
	GetIdThreshold() *string
	SetLanguageConfig(v string) *InitializeV2AdvanceRequest
	GetLanguageConfig() *string
	SetMRTDInput(v string) *InitializeV2AdvanceRequest
	GetMRTDInput() *string
	SetMerchantBizId(v string) *InitializeV2AdvanceRequest
	GetMerchantBizId() *string
	SetMerchantUserId(v string) *InitializeV2AdvanceRequest
	GetMerchantUserId() *string
	SetMetaInfo(v string) *InitializeV2AdvanceRequest
	GetMetaInfo() *string
	SetMobile(v string) *InitializeV2AdvanceRequest
	GetMobile() *string
	SetModel(v string) *InitializeV2AdvanceRequest
	GetModel() *string
	SetOcr(v string) *InitializeV2AdvanceRequest
	GetOcr() *string
	SetOcrValueStandard(v string) *InitializeV2AdvanceRequest
	GetOcrValueStandard() *string
	SetPages(v string) *InitializeV2AdvanceRequest
	GetPages() *string
	SetProcedurePriority(v string) *InitializeV2AdvanceRequest
	GetProcedurePriority() *string
	SetProductCode(v string) *InitializeV2AdvanceRequest
	GetProductCode() *string
	SetProductFlow(v string) *InitializeV2AdvanceRequest
	GetProductFlow() *string
	SetReturnFaces(v string) *InitializeV2AdvanceRequest
	GetReturnFaces() *string
	SetReturnUrl(v string) *InitializeV2AdvanceRequest
	GetReturnUrl() *string
	SetSaveFacePicture(v string) *InitializeV2AdvanceRequest
	GetSaveFacePicture() *string
	SetSceneCode(v string) *InitializeV2AdvanceRequest
	GetSceneCode() *string
	SetSecurityLevel(v string) *InitializeV2AdvanceRequest
	GetSecurityLevel() *string
	SetShowAlbumIcon(v string) *InitializeV2AdvanceRequest
	GetShowAlbumIcon() *string
	SetShowGuidePage(v string) *InitializeV2AdvanceRequest
	GetShowGuidePage() *string
	SetShowOcrResult(v string) *InitializeV2AdvanceRequest
	GetShowOcrResult() *string
	SetStyleConfig(v string) *InitializeV2AdvanceRequest
	GetStyleConfig() *string
	SetTargetFacePicture(v string) *InitializeV2AdvanceRequest
	GetTargetFacePicture() *string
	SetTargetFacePictureFileObject(v io.Reader) *InitializeV2AdvanceRequest
	GetTargetFacePictureFileObject() io.Reader
	SetTargetFacePictureUrl(v string) *InitializeV2AdvanceRequest
	GetTargetFacePictureUrl() *string
	SetTemplateConfig(v string) *InitializeV2AdvanceRequest
	GetTemplateConfig() *string
	SetTemplateRanCount(v string) *InitializeV2AdvanceRequest
	GetTemplateRanCount() *string
	SetTemplateType(v string) *InitializeV2AdvanceRequest
	GetTemplateType() *string
	SetUseNFC(v string) *InitializeV2AdvanceRequest
	GetUseNFC() *string
	SetVerifyModel(v string) *InitializeV2AdvanceRequest
	GetVerifyModel() *string
}

type InitializeV2AdvanceRequest struct {
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
	DocPageConfig []*string `json:"DocPageConfig,omitempty" xml:"DocPageConfig,omitempty" type:"Repeated"`
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
	FacePictureFileObject io.Reader `json:"FacePictureFile,omitempty" xml:"FacePictureFile,omitempty"`
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
	TargetFacePictureFileObject io.Reader `json:"TargetFacePictureFile,omitempty" xml:"TargetFacePictureFile,omitempty"`
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

func (s InitializeV2AdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s InitializeV2AdvanceRequest) GoString() string {
	return s.String()
}

func (s *InitializeV2AdvanceRequest) GetAppQualityCheck() *string {
	return s.AppQualityCheck
}

func (s *InitializeV2AdvanceRequest) GetAuthorize() *string {
	return s.Authorize
}

func (s *InitializeV2AdvanceRequest) GetAutoRegistration() *string {
	return s.AutoRegistration
}

func (s *InitializeV2AdvanceRequest) GetCallbackToken() *string {
	return s.CallbackToken
}

func (s *InitializeV2AdvanceRequest) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *InitializeV2AdvanceRequest) GetChameleonFrameEnable() *string {
	return s.ChameleonFrameEnable
}

func (s *InitializeV2AdvanceRequest) GetCrop() *string {
	return s.Crop
}

func (s *InitializeV2AdvanceRequest) GetDateOfBirth() *string {
	return s.DateOfBirth
}

func (s *InitializeV2AdvanceRequest) GetDateOfExpiry() *string {
	return s.DateOfExpiry
}

func (s *InitializeV2AdvanceRequest) GetDocName() *string {
	return s.DocName
}

func (s *InitializeV2AdvanceRequest) GetDocNo() *string {
	return s.DocNo
}

func (s *InitializeV2AdvanceRequest) GetDocPageConfig() []*string {
	return s.DocPageConfig
}

func (s *InitializeV2AdvanceRequest) GetDocScanMode() *string {
	return s.DocScanMode
}

func (s *InitializeV2AdvanceRequest) GetDocType() *string {
	return s.DocType
}

func (s *InitializeV2AdvanceRequest) GetDocVideo() *string {
	return s.DocVideo
}

func (s *InitializeV2AdvanceRequest) GetDocumentNumber() *string {
	return s.DocumentNumber
}

func (s *InitializeV2AdvanceRequest) GetEditOcrResult() *string {
	return s.EditOcrResult
}

func (s *InitializeV2AdvanceRequest) GetEmail() *string {
	return s.Email
}

func (s *InitializeV2AdvanceRequest) GetExperienceCode() *string {
	return s.ExperienceCode
}

func (s *InitializeV2AdvanceRequest) GetFaceGroupCodes() *string {
	return s.FaceGroupCodes
}

func (s *InitializeV2AdvanceRequest) GetFacePictureBase64() *string {
	return s.FacePictureBase64
}

func (s *InitializeV2AdvanceRequest) GetFacePictureFileObject() io.Reader {
	return s.FacePictureFileObject
}

func (s *InitializeV2AdvanceRequest) GetFacePictureUrl() *string {
	return s.FacePictureUrl
}

func (s *InitializeV2AdvanceRequest) GetFaceRegisterGroupCode() *string {
	return s.FaceRegisterGroupCode
}

func (s *InitializeV2AdvanceRequest) GetFaceVerifyThreshold() *string {
	return s.FaceVerifyThreshold
}

func (s *InitializeV2AdvanceRequest) GetIdFaceQuality() *string {
	return s.IdFaceQuality
}

func (s *InitializeV2AdvanceRequest) GetIdSpoof() *string {
	return s.IdSpoof
}

func (s *InitializeV2AdvanceRequest) GetIdThreshold() *string {
	return s.IdThreshold
}

func (s *InitializeV2AdvanceRequest) GetLanguageConfig() *string {
	return s.LanguageConfig
}

func (s *InitializeV2AdvanceRequest) GetMRTDInput() *string {
	return s.MRTDInput
}

func (s *InitializeV2AdvanceRequest) GetMerchantBizId() *string {
	return s.MerchantBizId
}

func (s *InitializeV2AdvanceRequest) GetMerchantUserId() *string {
	return s.MerchantUserId
}

func (s *InitializeV2AdvanceRequest) GetMetaInfo() *string {
	return s.MetaInfo
}

func (s *InitializeV2AdvanceRequest) GetMobile() *string {
	return s.Mobile
}

func (s *InitializeV2AdvanceRequest) GetModel() *string {
	return s.Model
}

func (s *InitializeV2AdvanceRequest) GetOcr() *string {
	return s.Ocr
}

func (s *InitializeV2AdvanceRequest) GetOcrValueStandard() *string {
	return s.OcrValueStandard
}

func (s *InitializeV2AdvanceRequest) GetPages() *string {
	return s.Pages
}

func (s *InitializeV2AdvanceRequest) GetProcedurePriority() *string {
	return s.ProcedurePriority
}

func (s *InitializeV2AdvanceRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *InitializeV2AdvanceRequest) GetProductFlow() *string {
	return s.ProductFlow
}

func (s *InitializeV2AdvanceRequest) GetReturnFaces() *string {
	return s.ReturnFaces
}

func (s *InitializeV2AdvanceRequest) GetReturnUrl() *string {
	return s.ReturnUrl
}

func (s *InitializeV2AdvanceRequest) GetSaveFacePicture() *string {
	return s.SaveFacePicture
}

func (s *InitializeV2AdvanceRequest) GetSceneCode() *string {
	return s.SceneCode
}

func (s *InitializeV2AdvanceRequest) GetSecurityLevel() *string {
	return s.SecurityLevel
}

func (s *InitializeV2AdvanceRequest) GetShowAlbumIcon() *string {
	return s.ShowAlbumIcon
}

func (s *InitializeV2AdvanceRequest) GetShowGuidePage() *string {
	return s.ShowGuidePage
}

func (s *InitializeV2AdvanceRequest) GetShowOcrResult() *string {
	return s.ShowOcrResult
}

func (s *InitializeV2AdvanceRequest) GetStyleConfig() *string {
	return s.StyleConfig
}

func (s *InitializeV2AdvanceRequest) GetTargetFacePicture() *string {
	return s.TargetFacePicture
}

func (s *InitializeV2AdvanceRequest) GetTargetFacePictureFileObject() io.Reader {
	return s.TargetFacePictureFileObject
}

func (s *InitializeV2AdvanceRequest) GetTargetFacePictureUrl() *string {
	return s.TargetFacePictureUrl
}

func (s *InitializeV2AdvanceRequest) GetTemplateConfig() *string {
	return s.TemplateConfig
}

func (s *InitializeV2AdvanceRequest) GetTemplateRanCount() *string {
	return s.TemplateRanCount
}

func (s *InitializeV2AdvanceRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *InitializeV2AdvanceRequest) GetUseNFC() *string {
	return s.UseNFC
}

func (s *InitializeV2AdvanceRequest) GetVerifyModel() *string {
	return s.VerifyModel
}

func (s *InitializeV2AdvanceRequest) SetAppQualityCheck(v string) *InitializeV2AdvanceRequest {
	s.AppQualityCheck = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetAuthorize(v string) *InitializeV2AdvanceRequest {
	s.Authorize = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetAutoRegistration(v string) *InitializeV2AdvanceRequest {
	s.AutoRegistration = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetCallbackToken(v string) *InitializeV2AdvanceRequest {
	s.CallbackToken = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetCallbackUrl(v string) *InitializeV2AdvanceRequest {
	s.CallbackUrl = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetChameleonFrameEnable(v string) *InitializeV2AdvanceRequest {
	s.ChameleonFrameEnable = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetCrop(v string) *InitializeV2AdvanceRequest {
	s.Crop = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetDateOfBirth(v string) *InitializeV2AdvanceRequest {
	s.DateOfBirth = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetDateOfExpiry(v string) *InitializeV2AdvanceRequest {
	s.DateOfExpiry = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetDocName(v string) *InitializeV2AdvanceRequest {
	s.DocName = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetDocNo(v string) *InitializeV2AdvanceRequest {
	s.DocNo = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetDocPageConfig(v []*string) *InitializeV2AdvanceRequest {
	s.DocPageConfig = v
	return s
}

func (s *InitializeV2AdvanceRequest) SetDocScanMode(v string) *InitializeV2AdvanceRequest {
	s.DocScanMode = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetDocType(v string) *InitializeV2AdvanceRequest {
	s.DocType = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetDocVideo(v string) *InitializeV2AdvanceRequest {
	s.DocVideo = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetDocumentNumber(v string) *InitializeV2AdvanceRequest {
	s.DocumentNumber = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetEditOcrResult(v string) *InitializeV2AdvanceRequest {
	s.EditOcrResult = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetEmail(v string) *InitializeV2AdvanceRequest {
	s.Email = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetExperienceCode(v string) *InitializeV2AdvanceRequest {
	s.ExperienceCode = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetFaceGroupCodes(v string) *InitializeV2AdvanceRequest {
	s.FaceGroupCodes = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetFacePictureBase64(v string) *InitializeV2AdvanceRequest {
	s.FacePictureBase64 = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetFacePictureFileObject(v io.Reader) *InitializeV2AdvanceRequest {
	s.FacePictureFileObject = v
	return s
}

func (s *InitializeV2AdvanceRequest) SetFacePictureUrl(v string) *InitializeV2AdvanceRequest {
	s.FacePictureUrl = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetFaceRegisterGroupCode(v string) *InitializeV2AdvanceRequest {
	s.FaceRegisterGroupCode = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetFaceVerifyThreshold(v string) *InitializeV2AdvanceRequest {
	s.FaceVerifyThreshold = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetIdFaceQuality(v string) *InitializeV2AdvanceRequest {
	s.IdFaceQuality = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetIdSpoof(v string) *InitializeV2AdvanceRequest {
	s.IdSpoof = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetIdThreshold(v string) *InitializeV2AdvanceRequest {
	s.IdThreshold = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetLanguageConfig(v string) *InitializeV2AdvanceRequest {
	s.LanguageConfig = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetMRTDInput(v string) *InitializeV2AdvanceRequest {
	s.MRTDInput = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetMerchantBizId(v string) *InitializeV2AdvanceRequest {
	s.MerchantBizId = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetMerchantUserId(v string) *InitializeV2AdvanceRequest {
	s.MerchantUserId = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetMetaInfo(v string) *InitializeV2AdvanceRequest {
	s.MetaInfo = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetMobile(v string) *InitializeV2AdvanceRequest {
	s.Mobile = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetModel(v string) *InitializeV2AdvanceRequest {
	s.Model = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetOcr(v string) *InitializeV2AdvanceRequest {
	s.Ocr = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetOcrValueStandard(v string) *InitializeV2AdvanceRequest {
	s.OcrValueStandard = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetPages(v string) *InitializeV2AdvanceRequest {
	s.Pages = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetProcedurePriority(v string) *InitializeV2AdvanceRequest {
	s.ProcedurePriority = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetProductCode(v string) *InitializeV2AdvanceRequest {
	s.ProductCode = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetProductFlow(v string) *InitializeV2AdvanceRequest {
	s.ProductFlow = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetReturnFaces(v string) *InitializeV2AdvanceRequest {
	s.ReturnFaces = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetReturnUrl(v string) *InitializeV2AdvanceRequest {
	s.ReturnUrl = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetSaveFacePicture(v string) *InitializeV2AdvanceRequest {
	s.SaveFacePicture = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetSceneCode(v string) *InitializeV2AdvanceRequest {
	s.SceneCode = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetSecurityLevel(v string) *InitializeV2AdvanceRequest {
	s.SecurityLevel = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetShowAlbumIcon(v string) *InitializeV2AdvanceRequest {
	s.ShowAlbumIcon = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetShowGuidePage(v string) *InitializeV2AdvanceRequest {
	s.ShowGuidePage = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetShowOcrResult(v string) *InitializeV2AdvanceRequest {
	s.ShowOcrResult = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetStyleConfig(v string) *InitializeV2AdvanceRequest {
	s.StyleConfig = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetTargetFacePicture(v string) *InitializeV2AdvanceRequest {
	s.TargetFacePicture = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetTargetFacePictureFileObject(v io.Reader) *InitializeV2AdvanceRequest {
	s.TargetFacePictureFileObject = v
	return s
}

func (s *InitializeV2AdvanceRequest) SetTargetFacePictureUrl(v string) *InitializeV2AdvanceRequest {
	s.TargetFacePictureUrl = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetTemplateConfig(v string) *InitializeV2AdvanceRequest {
	s.TemplateConfig = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetTemplateRanCount(v string) *InitializeV2AdvanceRequest {
	s.TemplateRanCount = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetTemplateType(v string) *InitializeV2AdvanceRequest {
	s.TemplateType = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetUseNFC(v string) *InitializeV2AdvanceRequest {
	s.UseNFC = &v
	return s
}

func (s *InitializeV2AdvanceRequest) SetVerifyModel(v string) *InitializeV2AdvanceRequest {
	s.VerifyModel = &v
	return s
}

func (s *InitializeV2AdvanceRequest) Validate() error {
	return dara.Validate(s)
}
