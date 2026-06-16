// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitializeV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetAppQualityCheck(v string) *InitializeV2Request
	GetAppQualityCheck() *string
	SetAuthorize(v string) *InitializeV2Request
	GetAuthorize() *string
	SetAutoRegistration(v string) *InitializeV2Request
	GetAutoRegistration() *string
	SetCallbackToken(v string) *InitializeV2Request
	GetCallbackToken() *string
	SetCallbackUrl(v string) *InitializeV2Request
	GetCallbackUrl() *string
	SetChameleonFrameEnable(v string) *InitializeV2Request
	GetChameleonFrameEnable() *string
	SetCrop(v string) *InitializeV2Request
	GetCrop() *string
	SetDateOfBirth(v string) *InitializeV2Request
	GetDateOfBirth() *string
	SetDateOfExpiry(v string) *InitializeV2Request
	GetDateOfExpiry() *string
	SetDocName(v string) *InitializeV2Request
	GetDocName() *string
	SetDocNo(v string) *InitializeV2Request
	GetDocNo() *string
	SetDocPageConfig(v []*string) *InitializeV2Request
	GetDocPageConfig() []*string
	SetDocScanMode(v string) *InitializeV2Request
	GetDocScanMode() *string
	SetDocType(v string) *InitializeV2Request
	GetDocType() *string
	SetDocVideo(v string) *InitializeV2Request
	GetDocVideo() *string
	SetDocumentNumber(v string) *InitializeV2Request
	GetDocumentNumber() *string
	SetEditOcrResult(v string) *InitializeV2Request
	GetEditOcrResult() *string
	SetEmail(v string) *InitializeV2Request
	GetEmail() *string
	SetExperienceCode(v string) *InitializeV2Request
	GetExperienceCode() *string
	SetFaceGroupCodes(v string) *InitializeV2Request
	GetFaceGroupCodes() *string
	SetFacePictureBase64(v string) *InitializeV2Request
	GetFacePictureBase64() *string
	SetFacePictureFile(v string) *InitializeV2Request
	GetFacePictureFile() *string
	SetFacePictureUrl(v string) *InitializeV2Request
	GetFacePictureUrl() *string
	SetFaceRegisterGroupCode(v string) *InitializeV2Request
	GetFaceRegisterGroupCode() *string
	SetFaceVerifyThreshold(v string) *InitializeV2Request
	GetFaceVerifyThreshold() *string
	SetIdFaceQuality(v string) *InitializeV2Request
	GetIdFaceQuality() *string
	SetIdSpoof(v string) *InitializeV2Request
	GetIdSpoof() *string
	SetIdThreshold(v string) *InitializeV2Request
	GetIdThreshold() *string
	SetLanguageConfig(v string) *InitializeV2Request
	GetLanguageConfig() *string
	SetMRTDInput(v string) *InitializeV2Request
	GetMRTDInput() *string
	SetMerchantBizId(v string) *InitializeV2Request
	GetMerchantBizId() *string
	SetMerchantUserId(v string) *InitializeV2Request
	GetMerchantUserId() *string
	SetMetaInfo(v string) *InitializeV2Request
	GetMetaInfo() *string
	SetMobile(v string) *InitializeV2Request
	GetMobile() *string
	SetModel(v string) *InitializeV2Request
	GetModel() *string
	SetOcr(v string) *InitializeV2Request
	GetOcr() *string
	SetOcrValueStandard(v string) *InitializeV2Request
	GetOcrValueStandard() *string
	SetPages(v string) *InitializeV2Request
	GetPages() *string
	SetProcedurePriority(v string) *InitializeV2Request
	GetProcedurePriority() *string
	SetProductCode(v string) *InitializeV2Request
	GetProductCode() *string
	SetProductFlow(v string) *InitializeV2Request
	GetProductFlow() *string
	SetReturnFaces(v string) *InitializeV2Request
	GetReturnFaces() *string
	SetReturnUrl(v string) *InitializeV2Request
	GetReturnUrl() *string
	SetSaveFacePicture(v string) *InitializeV2Request
	GetSaveFacePicture() *string
	SetSceneCode(v string) *InitializeV2Request
	GetSceneCode() *string
	SetSecurityLevel(v string) *InitializeV2Request
	GetSecurityLevel() *string
	SetShowAlbumIcon(v string) *InitializeV2Request
	GetShowAlbumIcon() *string
	SetShowGuidePage(v string) *InitializeV2Request
	GetShowGuidePage() *string
	SetShowOcrResult(v string) *InitializeV2Request
	GetShowOcrResult() *string
	SetStyleConfig(v string) *InitializeV2Request
	GetStyleConfig() *string
	SetTargetFacePicture(v string) *InitializeV2Request
	GetTargetFacePicture() *string
	SetTargetFacePictureFile(v string) *InitializeV2Request
	GetTargetFacePictureFile() *string
	SetTargetFacePictureUrl(v string) *InitializeV2Request
	GetTargetFacePictureUrl() *string
	SetTemplateConfig(v string) *InitializeV2Request
	GetTemplateConfig() *string
	SetTemplateRanCount(v string) *InitializeV2Request
	GetTemplateRanCount() *string
	SetTemplateType(v string) *InitializeV2Request
	GetTemplateType() *string
	SetUseNFC(v string) *InitializeV2Request
	GetUseNFC() *string
	SetVerifyModel(v string) *InitializeV2Request
	GetVerifyModel() *string
}

type InitializeV2Request struct {
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

func (s InitializeV2Request) String() string {
	return dara.Prettify(s)
}

func (s InitializeV2Request) GoString() string {
	return s.String()
}

func (s *InitializeV2Request) GetAppQualityCheck() *string {
	return s.AppQualityCheck
}

func (s *InitializeV2Request) GetAuthorize() *string {
	return s.Authorize
}

func (s *InitializeV2Request) GetAutoRegistration() *string {
	return s.AutoRegistration
}

func (s *InitializeV2Request) GetCallbackToken() *string {
	return s.CallbackToken
}

func (s *InitializeV2Request) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *InitializeV2Request) GetChameleonFrameEnable() *string {
	return s.ChameleonFrameEnable
}

func (s *InitializeV2Request) GetCrop() *string {
	return s.Crop
}

func (s *InitializeV2Request) GetDateOfBirth() *string {
	return s.DateOfBirth
}

func (s *InitializeV2Request) GetDateOfExpiry() *string {
	return s.DateOfExpiry
}

func (s *InitializeV2Request) GetDocName() *string {
	return s.DocName
}

func (s *InitializeV2Request) GetDocNo() *string {
	return s.DocNo
}

func (s *InitializeV2Request) GetDocPageConfig() []*string {
	return s.DocPageConfig
}

func (s *InitializeV2Request) GetDocScanMode() *string {
	return s.DocScanMode
}

func (s *InitializeV2Request) GetDocType() *string {
	return s.DocType
}

func (s *InitializeV2Request) GetDocVideo() *string {
	return s.DocVideo
}

func (s *InitializeV2Request) GetDocumentNumber() *string {
	return s.DocumentNumber
}

func (s *InitializeV2Request) GetEditOcrResult() *string {
	return s.EditOcrResult
}

func (s *InitializeV2Request) GetEmail() *string {
	return s.Email
}

func (s *InitializeV2Request) GetExperienceCode() *string {
	return s.ExperienceCode
}

func (s *InitializeV2Request) GetFaceGroupCodes() *string {
	return s.FaceGroupCodes
}

func (s *InitializeV2Request) GetFacePictureBase64() *string {
	return s.FacePictureBase64
}

func (s *InitializeV2Request) GetFacePictureFile() *string {
	return s.FacePictureFile
}

func (s *InitializeV2Request) GetFacePictureUrl() *string {
	return s.FacePictureUrl
}

func (s *InitializeV2Request) GetFaceRegisterGroupCode() *string {
	return s.FaceRegisterGroupCode
}

func (s *InitializeV2Request) GetFaceVerifyThreshold() *string {
	return s.FaceVerifyThreshold
}

func (s *InitializeV2Request) GetIdFaceQuality() *string {
	return s.IdFaceQuality
}

func (s *InitializeV2Request) GetIdSpoof() *string {
	return s.IdSpoof
}

func (s *InitializeV2Request) GetIdThreshold() *string {
	return s.IdThreshold
}

func (s *InitializeV2Request) GetLanguageConfig() *string {
	return s.LanguageConfig
}

func (s *InitializeV2Request) GetMRTDInput() *string {
	return s.MRTDInput
}

func (s *InitializeV2Request) GetMerchantBizId() *string {
	return s.MerchantBizId
}

func (s *InitializeV2Request) GetMerchantUserId() *string {
	return s.MerchantUserId
}

func (s *InitializeV2Request) GetMetaInfo() *string {
	return s.MetaInfo
}

func (s *InitializeV2Request) GetMobile() *string {
	return s.Mobile
}

func (s *InitializeV2Request) GetModel() *string {
	return s.Model
}

func (s *InitializeV2Request) GetOcr() *string {
	return s.Ocr
}

func (s *InitializeV2Request) GetOcrValueStandard() *string {
	return s.OcrValueStandard
}

func (s *InitializeV2Request) GetPages() *string {
	return s.Pages
}

func (s *InitializeV2Request) GetProcedurePriority() *string {
	return s.ProcedurePriority
}

func (s *InitializeV2Request) GetProductCode() *string {
	return s.ProductCode
}

func (s *InitializeV2Request) GetProductFlow() *string {
	return s.ProductFlow
}

func (s *InitializeV2Request) GetReturnFaces() *string {
	return s.ReturnFaces
}

func (s *InitializeV2Request) GetReturnUrl() *string {
	return s.ReturnUrl
}

func (s *InitializeV2Request) GetSaveFacePicture() *string {
	return s.SaveFacePicture
}

func (s *InitializeV2Request) GetSceneCode() *string {
	return s.SceneCode
}

func (s *InitializeV2Request) GetSecurityLevel() *string {
	return s.SecurityLevel
}

func (s *InitializeV2Request) GetShowAlbumIcon() *string {
	return s.ShowAlbumIcon
}

func (s *InitializeV2Request) GetShowGuidePage() *string {
	return s.ShowGuidePage
}

func (s *InitializeV2Request) GetShowOcrResult() *string {
	return s.ShowOcrResult
}

func (s *InitializeV2Request) GetStyleConfig() *string {
	return s.StyleConfig
}

func (s *InitializeV2Request) GetTargetFacePicture() *string {
	return s.TargetFacePicture
}

func (s *InitializeV2Request) GetTargetFacePictureFile() *string {
	return s.TargetFacePictureFile
}

func (s *InitializeV2Request) GetTargetFacePictureUrl() *string {
	return s.TargetFacePictureUrl
}

func (s *InitializeV2Request) GetTemplateConfig() *string {
	return s.TemplateConfig
}

func (s *InitializeV2Request) GetTemplateRanCount() *string {
	return s.TemplateRanCount
}

func (s *InitializeV2Request) GetTemplateType() *string {
	return s.TemplateType
}

func (s *InitializeV2Request) GetUseNFC() *string {
	return s.UseNFC
}

func (s *InitializeV2Request) GetVerifyModel() *string {
	return s.VerifyModel
}

func (s *InitializeV2Request) SetAppQualityCheck(v string) *InitializeV2Request {
	s.AppQualityCheck = &v
	return s
}

func (s *InitializeV2Request) SetAuthorize(v string) *InitializeV2Request {
	s.Authorize = &v
	return s
}

func (s *InitializeV2Request) SetAutoRegistration(v string) *InitializeV2Request {
	s.AutoRegistration = &v
	return s
}

func (s *InitializeV2Request) SetCallbackToken(v string) *InitializeV2Request {
	s.CallbackToken = &v
	return s
}

func (s *InitializeV2Request) SetCallbackUrl(v string) *InitializeV2Request {
	s.CallbackUrl = &v
	return s
}

func (s *InitializeV2Request) SetChameleonFrameEnable(v string) *InitializeV2Request {
	s.ChameleonFrameEnable = &v
	return s
}

func (s *InitializeV2Request) SetCrop(v string) *InitializeV2Request {
	s.Crop = &v
	return s
}

func (s *InitializeV2Request) SetDateOfBirth(v string) *InitializeV2Request {
	s.DateOfBirth = &v
	return s
}

func (s *InitializeV2Request) SetDateOfExpiry(v string) *InitializeV2Request {
	s.DateOfExpiry = &v
	return s
}

func (s *InitializeV2Request) SetDocName(v string) *InitializeV2Request {
	s.DocName = &v
	return s
}

func (s *InitializeV2Request) SetDocNo(v string) *InitializeV2Request {
	s.DocNo = &v
	return s
}

func (s *InitializeV2Request) SetDocPageConfig(v []*string) *InitializeV2Request {
	s.DocPageConfig = v
	return s
}

func (s *InitializeV2Request) SetDocScanMode(v string) *InitializeV2Request {
	s.DocScanMode = &v
	return s
}

func (s *InitializeV2Request) SetDocType(v string) *InitializeV2Request {
	s.DocType = &v
	return s
}

func (s *InitializeV2Request) SetDocVideo(v string) *InitializeV2Request {
	s.DocVideo = &v
	return s
}

func (s *InitializeV2Request) SetDocumentNumber(v string) *InitializeV2Request {
	s.DocumentNumber = &v
	return s
}

func (s *InitializeV2Request) SetEditOcrResult(v string) *InitializeV2Request {
	s.EditOcrResult = &v
	return s
}

func (s *InitializeV2Request) SetEmail(v string) *InitializeV2Request {
	s.Email = &v
	return s
}

func (s *InitializeV2Request) SetExperienceCode(v string) *InitializeV2Request {
	s.ExperienceCode = &v
	return s
}

func (s *InitializeV2Request) SetFaceGroupCodes(v string) *InitializeV2Request {
	s.FaceGroupCodes = &v
	return s
}

func (s *InitializeV2Request) SetFacePictureBase64(v string) *InitializeV2Request {
	s.FacePictureBase64 = &v
	return s
}

func (s *InitializeV2Request) SetFacePictureFile(v string) *InitializeV2Request {
	s.FacePictureFile = &v
	return s
}

func (s *InitializeV2Request) SetFacePictureUrl(v string) *InitializeV2Request {
	s.FacePictureUrl = &v
	return s
}

func (s *InitializeV2Request) SetFaceRegisterGroupCode(v string) *InitializeV2Request {
	s.FaceRegisterGroupCode = &v
	return s
}

func (s *InitializeV2Request) SetFaceVerifyThreshold(v string) *InitializeV2Request {
	s.FaceVerifyThreshold = &v
	return s
}

func (s *InitializeV2Request) SetIdFaceQuality(v string) *InitializeV2Request {
	s.IdFaceQuality = &v
	return s
}

func (s *InitializeV2Request) SetIdSpoof(v string) *InitializeV2Request {
	s.IdSpoof = &v
	return s
}

func (s *InitializeV2Request) SetIdThreshold(v string) *InitializeV2Request {
	s.IdThreshold = &v
	return s
}

func (s *InitializeV2Request) SetLanguageConfig(v string) *InitializeV2Request {
	s.LanguageConfig = &v
	return s
}

func (s *InitializeV2Request) SetMRTDInput(v string) *InitializeV2Request {
	s.MRTDInput = &v
	return s
}

func (s *InitializeV2Request) SetMerchantBizId(v string) *InitializeV2Request {
	s.MerchantBizId = &v
	return s
}

func (s *InitializeV2Request) SetMerchantUserId(v string) *InitializeV2Request {
	s.MerchantUserId = &v
	return s
}

func (s *InitializeV2Request) SetMetaInfo(v string) *InitializeV2Request {
	s.MetaInfo = &v
	return s
}

func (s *InitializeV2Request) SetMobile(v string) *InitializeV2Request {
	s.Mobile = &v
	return s
}

func (s *InitializeV2Request) SetModel(v string) *InitializeV2Request {
	s.Model = &v
	return s
}

func (s *InitializeV2Request) SetOcr(v string) *InitializeV2Request {
	s.Ocr = &v
	return s
}

func (s *InitializeV2Request) SetOcrValueStandard(v string) *InitializeV2Request {
	s.OcrValueStandard = &v
	return s
}

func (s *InitializeV2Request) SetPages(v string) *InitializeV2Request {
	s.Pages = &v
	return s
}

func (s *InitializeV2Request) SetProcedurePriority(v string) *InitializeV2Request {
	s.ProcedurePriority = &v
	return s
}

func (s *InitializeV2Request) SetProductCode(v string) *InitializeV2Request {
	s.ProductCode = &v
	return s
}

func (s *InitializeV2Request) SetProductFlow(v string) *InitializeV2Request {
	s.ProductFlow = &v
	return s
}

func (s *InitializeV2Request) SetReturnFaces(v string) *InitializeV2Request {
	s.ReturnFaces = &v
	return s
}

func (s *InitializeV2Request) SetReturnUrl(v string) *InitializeV2Request {
	s.ReturnUrl = &v
	return s
}

func (s *InitializeV2Request) SetSaveFacePicture(v string) *InitializeV2Request {
	s.SaveFacePicture = &v
	return s
}

func (s *InitializeV2Request) SetSceneCode(v string) *InitializeV2Request {
	s.SceneCode = &v
	return s
}

func (s *InitializeV2Request) SetSecurityLevel(v string) *InitializeV2Request {
	s.SecurityLevel = &v
	return s
}

func (s *InitializeV2Request) SetShowAlbumIcon(v string) *InitializeV2Request {
	s.ShowAlbumIcon = &v
	return s
}

func (s *InitializeV2Request) SetShowGuidePage(v string) *InitializeV2Request {
	s.ShowGuidePage = &v
	return s
}

func (s *InitializeV2Request) SetShowOcrResult(v string) *InitializeV2Request {
	s.ShowOcrResult = &v
	return s
}

func (s *InitializeV2Request) SetStyleConfig(v string) *InitializeV2Request {
	s.StyleConfig = &v
	return s
}

func (s *InitializeV2Request) SetTargetFacePicture(v string) *InitializeV2Request {
	s.TargetFacePicture = &v
	return s
}

func (s *InitializeV2Request) SetTargetFacePictureFile(v string) *InitializeV2Request {
	s.TargetFacePictureFile = &v
	return s
}

func (s *InitializeV2Request) SetTargetFacePictureUrl(v string) *InitializeV2Request {
	s.TargetFacePictureUrl = &v
	return s
}

func (s *InitializeV2Request) SetTemplateConfig(v string) *InitializeV2Request {
	s.TemplateConfig = &v
	return s
}

func (s *InitializeV2Request) SetTemplateRanCount(v string) *InitializeV2Request {
	s.TemplateRanCount = &v
	return s
}

func (s *InitializeV2Request) SetTemplateType(v string) *InitializeV2Request {
	s.TemplateType = &v
	return s
}

func (s *InitializeV2Request) SetUseNFC(v string) *InitializeV2Request {
	s.UseNFC = &v
	return s
}

func (s *InitializeV2Request) SetVerifyModel(v string) *InitializeV2Request {
	s.VerifyModel = &v
	return s
}

func (s *InitializeV2Request) Validate() error {
	return dara.Validate(s)
}
