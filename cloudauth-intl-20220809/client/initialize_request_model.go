// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitializeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppQualityCheck(v string) *InitializeRequest
	GetAppQualityCheck() *string
	SetAuthorize(v string) *InitializeRequest
	GetAuthorize() *string
	SetAutoRegistration(v string) *InitializeRequest
	GetAutoRegistration() *string
	SetCallbackToken(v string) *InitializeRequest
	GetCallbackToken() *string
	SetCallbackUrl(v string) *InitializeRequest
	GetCallbackUrl() *string
	SetChameleonFrameEnable(v string) *InitializeRequest
	GetChameleonFrameEnable() *string
	SetCrop(v string) *InitializeRequest
	GetCrop() *string
	SetDateOfBirth(v string) *InitializeRequest
	GetDateOfBirth() *string
	SetDateOfExpiry(v string) *InitializeRequest
	GetDateOfExpiry() *string
	SetDocName(v string) *InitializeRequest
	GetDocName() *string
	SetDocNo(v string) *InitializeRequest
	GetDocNo() *string
	SetDocPageConfig(v []*string) *InitializeRequest
	GetDocPageConfig() []*string
	SetDocScanMode(v string) *InitializeRequest
	GetDocScanMode() *string
	SetDocType(v string) *InitializeRequest
	GetDocType() *string
	SetDocVideo(v string) *InitializeRequest
	GetDocVideo() *string
	SetDocumentNumber(v string) *InitializeRequest
	GetDocumentNumber() *string
	SetEditOcrResult(v string) *InitializeRequest
	GetEditOcrResult() *string
	SetExperienceCode(v string) *InitializeRequest
	GetExperienceCode() *string
	SetFaceGroupCodes(v string) *InitializeRequest
	GetFaceGroupCodes() *string
	SetFacePictureBase64(v string) *InitializeRequest
	GetFacePictureBase64() *string
	SetFacePictureUrl(v string) *InitializeRequest
	GetFacePictureUrl() *string
	SetFaceRegisterGroupCode(v string) *InitializeRequest
	GetFaceRegisterGroupCode() *string
	SetFaceVerifyThreshold(v string) *InitializeRequest
	GetFaceVerifyThreshold() *string
	SetIdFaceQuality(v string) *InitializeRequest
	GetIdFaceQuality() *string
	SetIdSpoof(v string) *InitializeRequest
	GetIdSpoof() *string
	SetIdThreshold(v string) *InitializeRequest
	GetIdThreshold() *string
	SetLanguageConfig(v string) *InitializeRequest
	GetLanguageConfig() *string
	SetMRTDInput(v string) *InitializeRequest
	GetMRTDInput() *string
	SetMerchantBizId(v string) *InitializeRequest
	GetMerchantBizId() *string
	SetMerchantUserId(v string) *InitializeRequest
	GetMerchantUserId() *string
	SetMetaInfo(v string) *InitializeRequest
	GetMetaInfo() *string
	SetModel(v string) *InitializeRequest
	GetModel() *string
	SetOcr(v string) *InitializeRequest
	GetOcr() *string
	SetPages(v string) *InitializeRequest
	GetPages() *string
	SetProcedurePriority(v string) *InitializeRequest
	GetProcedurePriority() *string
	SetProductCode(v string) *InitializeRequest
	GetProductCode() *string
	SetProductFlow(v string) *InitializeRequest
	GetProductFlow() *string
	SetReturnFaces(v string) *InitializeRequest
	GetReturnFaces() *string
	SetReturnUrl(v string) *InitializeRequest
	GetReturnUrl() *string
	SetSaveFacePicture(v string) *InitializeRequest
	GetSaveFacePicture() *string
	SetSceneCode(v string) *InitializeRequest
	GetSceneCode() *string
	SetSecurityLevel(v string) *InitializeRequest
	GetSecurityLevel() *string
	SetShowAlbumIcon(v string) *InitializeRequest
	GetShowAlbumIcon() *string
	SetShowGuidePage(v string) *InitializeRequest
	GetShowGuidePage() *string
	SetShowOcrResult(v string) *InitializeRequest
	GetShowOcrResult() *string
	SetStyleConfig(v string) *InitializeRequest
	GetStyleConfig() *string
	SetTargetFacePicture(v string) *InitializeRequest
	GetTargetFacePicture() *string
	SetTargetFacePictureUrl(v string) *InitializeRequest
	GetTargetFacePictureUrl() *string
	SetUseNFC(v string) *InitializeRequest
	GetUseNFC() *string
	SetVerifyModel(v string) *InitializeRequest
	GetVerifyModel() *string
}

type InitializeRequest struct {
	AppQualityCheck *string `json:"AppQualityCheck,omitempty" xml:"AppQualityCheck,omitempty"`
	Authorize       *string `json:"Authorize,omitempty" xml:"Authorize,omitempty"`
	// example:
	//
	// 0
	AutoRegistration     *string `json:"AutoRegistration,omitempty" xml:"AutoRegistration,omitempty"`
	CallbackToken        *string `json:"CallbackToken,omitempty" xml:"CallbackToken,omitempty"`
	CallbackUrl          *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	ChameleonFrameEnable *string `json:"ChameleonFrameEnable,omitempty" xml:"ChameleonFrameEnable,omitempty"`
	// example:
	//
	// *
	Crop          *string   `json:"Crop,omitempty" xml:"Crop,omitempty"`
	DateOfBirth   *string   `json:"DateOfBirth,omitempty" xml:"DateOfBirth,omitempty"`
	DateOfExpiry  *string   `json:"DateOfExpiry,omitempty" xml:"DateOfExpiry,omitempty"`
	DocName       *string   `json:"DocName,omitempty" xml:"DocName,omitempty"`
	DocNo         *string   `json:"DocNo,omitempty" xml:"DocNo,omitempty"`
	DocPageConfig []*string `json:"DocPageConfig,omitempty" xml:"DocPageConfig,omitempty" type:"Repeated"`
	DocScanMode   *string   `json:"DocScanMode,omitempty" xml:"DocScanMode,omitempty"`
	// example:
	//
	// 01000000
	DocType        *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	DocVideo       *string `json:"DocVideo,omitempty" xml:"DocVideo,omitempty"`
	DocumentNumber *string `json:"DocumentNumber,omitempty" xml:"DocumentNumber,omitempty"`
	EditOcrResult  *string `json:"EditOcrResult,omitempty" xml:"EditOcrResult,omitempty"`
	ExperienceCode *string `json:"ExperienceCode,omitempty" xml:"ExperienceCode,omitempty"`
	// example:
	//
	// 0e0c34a77f
	FaceGroupCodes    *string `json:"FaceGroupCodes,omitempty" xml:"FaceGroupCodes,omitempty"`
	FacePictureBase64 *string `json:"FacePictureBase64,omitempty" xml:"FacePictureBase64,omitempty"`
	// example:
	//
	// ***
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
	// *
	IdFaceQuality *string `json:"IdFaceQuality,omitempty" xml:"IdFaceQuality,omitempty"`
	// example:
	//
	// Y
	IdSpoof        *string `json:"IdSpoof,omitempty" xml:"IdSpoof,omitempty"`
	IdThreshold    *string `json:"IdThreshold,omitempty" xml:"IdThreshold,omitempty"`
	LanguageConfig *string `json:"LanguageConfig,omitempty" xml:"LanguageConfig,omitempty"`
	MRTDInput      *string `json:"MRTDInput,omitempty" xml:"MRTDInput,omitempty"`
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
	// {\\"bioMetaInfo\\":\\"4.1.0:2916352,0\\",\\"deviceType\\":\\"web\\",\\"ua\\":\\"Mozilla/5.0 (Macintosh
	MetaInfo *string `json:"MetaInfo,omitempty" xml:"MetaInfo,omitempty"`
	Model    *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// OCR。
	//
	// example:
	//
	// *
	Ocr *string `json:"Ocr,omitempty" xml:"Ocr,omitempty"`
	// example:
	//
	// 1
	Pages             *string `json:"Pages,omitempty" xml:"Pages,omitempty"`
	ProcedurePriority *string `json:"ProcedurePriority,omitempty" xml:"ProcedurePriority,omitempty"`
	// example:
	//
	// eKYC
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
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
	// PAY**
	SceneCode     *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	SecurityLevel *string `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	ShowAlbumIcon *string `json:"ShowAlbumIcon,omitempty" xml:"ShowAlbumIcon,omitempty"`
	ShowGuidePage *string `json:"ShowGuidePage,omitempty" xml:"ShowGuidePage,omitempty"`
	ShowOcrResult *string `json:"ShowOcrResult,omitempty" xml:"ShowOcrResult,omitempty"`
	StyleConfig   *string `json:"StyleConfig,omitempty" xml:"StyleConfig,omitempty"`
	// example:
	//
	// base64
	TargetFacePicture *string `json:"TargetFacePicture,omitempty" xml:"TargetFacePicture,omitempty"`
	// example:
	//
	// https://www.xxxxx.com/1.jpg
	TargetFacePictureUrl *string `json:"TargetFacePictureUrl,omitempty" xml:"TargetFacePictureUrl,omitempty"`
	UseNFC               *string `json:"UseNFC,omitempty" xml:"UseNFC,omitempty"`
	// example:
	//
	// 0
	VerifyModel *string `json:"VerifyModel,omitempty" xml:"VerifyModel,omitempty"`
}

func (s InitializeRequest) String() string {
	return dara.Prettify(s)
}

func (s InitializeRequest) GoString() string {
	return s.String()
}

func (s *InitializeRequest) GetAppQualityCheck() *string {
	return s.AppQualityCheck
}

func (s *InitializeRequest) GetAuthorize() *string {
	return s.Authorize
}

func (s *InitializeRequest) GetAutoRegistration() *string {
	return s.AutoRegistration
}

func (s *InitializeRequest) GetCallbackToken() *string {
	return s.CallbackToken
}

func (s *InitializeRequest) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *InitializeRequest) GetChameleonFrameEnable() *string {
	return s.ChameleonFrameEnable
}

func (s *InitializeRequest) GetCrop() *string {
	return s.Crop
}

func (s *InitializeRequest) GetDateOfBirth() *string {
	return s.DateOfBirth
}

func (s *InitializeRequest) GetDateOfExpiry() *string {
	return s.DateOfExpiry
}

func (s *InitializeRequest) GetDocName() *string {
	return s.DocName
}

func (s *InitializeRequest) GetDocNo() *string {
	return s.DocNo
}

func (s *InitializeRequest) GetDocPageConfig() []*string {
	return s.DocPageConfig
}

func (s *InitializeRequest) GetDocScanMode() *string {
	return s.DocScanMode
}

func (s *InitializeRequest) GetDocType() *string {
	return s.DocType
}

func (s *InitializeRequest) GetDocVideo() *string {
	return s.DocVideo
}

func (s *InitializeRequest) GetDocumentNumber() *string {
	return s.DocumentNumber
}

func (s *InitializeRequest) GetEditOcrResult() *string {
	return s.EditOcrResult
}

func (s *InitializeRequest) GetExperienceCode() *string {
	return s.ExperienceCode
}

func (s *InitializeRequest) GetFaceGroupCodes() *string {
	return s.FaceGroupCodes
}

func (s *InitializeRequest) GetFacePictureBase64() *string {
	return s.FacePictureBase64
}

func (s *InitializeRequest) GetFacePictureUrl() *string {
	return s.FacePictureUrl
}

func (s *InitializeRequest) GetFaceRegisterGroupCode() *string {
	return s.FaceRegisterGroupCode
}

func (s *InitializeRequest) GetFaceVerifyThreshold() *string {
	return s.FaceVerifyThreshold
}

func (s *InitializeRequest) GetIdFaceQuality() *string {
	return s.IdFaceQuality
}

func (s *InitializeRequest) GetIdSpoof() *string {
	return s.IdSpoof
}

func (s *InitializeRequest) GetIdThreshold() *string {
	return s.IdThreshold
}

func (s *InitializeRequest) GetLanguageConfig() *string {
	return s.LanguageConfig
}

func (s *InitializeRequest) GetMRTDInput() *string {
	return s.MRTDInput
}

func (s *InitializeRequest) GetMerchantBizId() *string {
	return s.MerchantBizId
}

func (s *InitializeRequest) GetMerchantUserId() *string {
	return s.MerchantUserId
}

func (s *InitializeRequest) GetMetaInfo() *string {
	return s.MetaInfo
}

func (s *InitializeRequest) GetModel() *string {
	return s.Model
}

func (s *InitializeRequest) GetOcr() *string {
	return s.Ocr
}

func (s *InitializeRequest) GetPages() *string {
	return s.Pages
}

func (s *InitializeRequest) GetProcedurePriority() *string {
	return s.ProcedurePriority
}

func (s *InitializeRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *InitializeRequest) GetProductFlow() *string {
	return s.ProductFlow
}

func (s *InitializeRequest) GetReturnFaces() *string {
	return s.ReturnFaces
}

func (s *InitializeRequest) GetReturnUrl() *string {
	return s.ReturnUrl
}

func (s *InitializeRequest) GetSaveFacePicture() *string {
	return s.SaveFacePicture
}

func (s *InitializeRequest) GetSceneCode() *string {
	return s.SceneCode
}

func (s *InitializeRequest) GetSecurityLevel() *string {
	return s.SecurityLevel
}

func (s *InitializeRequest) GetShowAlbumIcon() *string {
	return s.ShowAlbumIcon
}

func (s *InitializeRequest) GetShowGuidePage() *string {
	return s.ShowGuidePage
}

func (s *InitializeRequest) GetShowOcrResult() *string {
	return s.ShowOcrResult
}

func (s *InitializeRequest) GetStyleConfig() *string {
	return s.StyleConfig
}

func (s *InitializeRequest) GetTargetFacePicture() *string {
	return s.TargetFacePicture
}

func (s *InitializeRequest) GetTargetFacePictureUrl() *string {
	return s.TargetFacePictureUrl
}

func (s *InitializeRequest) GetUseNFC() *string {
	return s.UseNFC
}

func (s *InitializeRequest) GetVerifyModel() *string {
	return s.VerifyModel
}

func (s *InitializeRequest) SetAppQualityCheck(v string) *InitializeRequest {
	s.AppQualityCheck = &v
	return s
}

func (s *InitializeRequest) SetAuthorize(v string) *InitializeRequest {
	s.Authorize = &v
	return s
}

func (s *InitializeRequest) SetAutoRegistration(v string) *InitializeRequest {
	s.AutoRegistration = &v
	return s
}

func (s *InitializeRequest) SetCallbackToken(v string) *InitializeRequest {
	s.CallbackToken = &v
	return s
}

func (s *InitializeRequest) SetCallbackUrl(v string) *InitializeRequest {
	s.CallbackUrl = &v
	return s
}

func (s *InitializeRequest) SetChameleonFrameEnable(v string) *InitializeRequest {
	s.ChameleonFrameEnable = &v
	return s
}

func (s *InitializeRequest) SetCrop(v string) *InitializeRequest {
	s.Crop = &v
	return s
}

func (s *InitializeRequest) SetDateOfBirth(v string) *InitializeRequest {
	s.DateOfBirth = &v
	return s
}

func (s *InitializeRequest) SetDateOfExpiry(v string) *InitializeRequest {
	s.DateOfExpiry = &v
	return s
}

func (s *InitializeRequest) SetDocName(v string) *InitializeRequest {
	s.DocName = &v
	return s
}

func (s *InitializeRequest) SetDocNo(v string) *InitializeRequest {
	s.DocNo = &v
	return s
}

func (s *InitializeRequest) SetDocPageConfig(v []*string) *InitializeRequest {
	s.DocPageConfig = v
	return s
}

func (s *InitializeRequest) SetDocScanMode(v string) *InitializeRequest {
	s.DocScanMode = &v
	return s
}

func (s *InitializeRequest) SetDocType(v string) *InitializeRequest {
	s.DocType = &v
	return s
}

func (s *InitializeRequest) SetDocVideo(v string) *InitializeRequest {
	s.DocVideo = &v
	return s
}

func (s *InitializeRequest) SetDocumentNumber(v string) *InitializeRequest {
	s.DocumentNumber = &v
	return s
}

func (s *InitializeRequest) SetEditOcrResult(v string) *InitializeRequest {
	s.EditOcrResult = &v
	return s
}

func (s *InitializeRequest) SetExperienceCode(v string) *InitializeRequest {
	s.ExperienceCode = &v
	return s
}

func (s *InitializeRequest) SetFaceGroupCodes(v string) *InitializeRequest {
	s.FaceGroupCodes = &v
	return s
}

func (s *InitializeRequest) SetFacePictureBase64(v string) *InitializeRequest {
	s.FacePictureBase64 = &v
	return s
}

func (s *InitializeRequest) SetFacePictureUrl(v string) *InitializeRequest {
	s.FacePictureUrl = &v
	return s
}

func (s *InitializeRequest) SetFaceRegisterGroupCode(v string) *InitializeRequest {
	s.FaceRegisterGroupCode = &v
	return s
}

func (s *InitializeRequest) SetFaceVerifyThreshold(v string) *InitializeRequest {
	s.FaceVerifyThreshold = &v
	return s
}

func (s *InitializeRequest) SetIdFaceQuality(v string) *InitializeRequest {
	s.IdFaceQuality = &v
	return s
}

func (s *InitializeRequest) SetIdSpoof(v string) *InitializeRequest {
	s.IdSpoof = &v
	return s
}

func (s *InitializeRequest) SetIdThreshold(v string) *InitializeRequest {
	s.IdThreshold = &v
	return s
}

func (s *InitializeRequest) SetLanguageConfig(v string) *InitializeRequest {
	s.LanguageConfig = &v
	return s
}

func (s *InitializeRequest) SetMRTDInput(v string) *InitializeRequest {
	s.MRTDInput = &v
	return s
}

func (s *InitializeRequest) SetMerchantBizId(v string) *InitializeRequest {
	s.MerchantBizId = &v
	return s
}

func (s *InitializeRequest) SetMerchantUserId(v string) *InitializeRequest {
	s.MerchantUserId = &v
	return s
}

func (s *InitializeRequest) SetMetaInfo(v string) *InitializeRequest {
	s.MetaInfo = &v
	return s
}

func (s *InitializeRequest) SetModel(v string) *InitializeRequest {
	s.Model = &v
	return s
}

func (s *InitializeRequest) SetOcr(v string) *InitializeRequest {
	s.Ocr = &v
	return s
}

func (s *InitializeRequest) SetPages(v string) *InitializeRequest {
	s.Pages = &v
	return s
}

func (s *InitializeRequest) SetProcedurePriority(v string) *InitializeRequest {
	s.ProcedurePriority = &v
	return s
}

func (s *InitializeRequest) SetProductCode(v string) *InitializeRequest {
	s.ProductCode = &v
	return s
}

func (s *InitializeRequest) SetProductFlow(v string) *InitializeRequest {
	s.ProductFlow = &v
	return s
}

func (s *InitializeRequest) SetReturnFaces(v string) *InitializeRequest {
	s.ReturnFaces = &v
	return s
}

func (s *InitializeRequest) SetReturnUrl(v string) *InitializeRequest {
	s.ReturnUrl = &v
	return s
}

func (s *InitializeRequest) SetSaveFacePicture(v string) *InitializeRequest {
	s.SaveFacePicture = &v
	return s
}

func (s *InitializeRequest) SetSceneCode(v string) *InitializeRequest {
	s.SceneCode = &v
	return s
}

func (s *InitializeRequest) SetSecurityLevel(v string) *InitializeRequest {
	s.SecurityLevel = &v
	return s
}

func (s *InitializeRequest) SetShowAlbumIcon(v string) *InitializeRequest {
	s.ShowAlbumIcon = &v
	return s
}

func (s *InitializeRequest) SetShowGuidePage(v string) *InitializeRequest {
	s.ShowGuidePage = &v
	return s
}

func (s *InitializeRequest) SetShowOcrResult(v string) *InitializeRequest {
	s.ShowOcrResult = &v
	return s
}

func (s *InitializeRequest) SetStyleConfig(v string) *InitializeRequest {
	s.StyleConfig = &v
	return s
}

func (s *InitializeRequest) SetTargetFacePicture(v string) *InitializeRequest {
	s.TargetFacePicture = &v
	return s
}

func (s *InitializeRequest) SetTargetFacePictureUrl(v string) *InitializeRequest {
	s.TargetFacePictureUrl = &v
	return s
}

func (s *InitializeRequest) SetUseNFC(v string) *InitializeRequest {
	s.UseNFC = &v
	return s
}

func (s *InitializeRequest) SetVerifyModel(v string) *InitializeRequest {
	s.VerifyModel = &v
	return s
}

func (s *InitializeRequest) Validate() error {
	return dara.Validate(s)
}
