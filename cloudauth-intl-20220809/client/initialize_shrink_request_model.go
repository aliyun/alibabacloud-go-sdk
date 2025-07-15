// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitializeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppQualityCheck(v string) *InitializeShrinkRequest
	GetAppQualityCheck() *string
	SetAuthorize(v string) *InitializeShrinkRequest
	GetAuthorize() *string
	SetCallbackToken(v string) *InitializeShrinkRequest
	GetCallbackToken() *string
	SetCallbackUrl(v string) *InitializeShrinkRequest
	GetCallbackUrl() *string
	SetChameleonFrameEnable(v string) *InitializeShrinkRequest
	GetChameleonFrameEnable() *string
	SetCrop(v string) *InitializeShrinkRequest
	GetCrop() *string
	SetDateOfBirth(v string) *InitializeShrinkRequest
	GetDateOfBirth() *string
	SetDateOfExpiry(v string) *InitializeShrinkRequest
	GetDateOfExpiry() *string
	SetDocName(v string) *InitializeShrinkRequest
	GetDocName() *string
	SetDocNo(v string) *InitializeShrinkRequest
	GetDocNo() *string
	SetDocPageConfigShrink(v string) *InitializeShrinkRequest
	GetDocPageConfigShrink() *string
	SetDocScanMode(v string) *InitializeShrinkRequest
	GetDocScanMode() *string
	SetDocType(v string) *InitializeShrinkRequest
	GetDocType() *string
	SetDocVideo(v string) *InitializeShrinkRequest
	GetDocVideo() *string
	SetDocumentNumber(v string) *InitializeShrinkRequest
	GetDocumentNumber() *string
	SetExperienceCode(v string) *InitializeShrinkRequest
	GetExperienceCode() *string
	SetFacePictureBase64(v string) *InitializeShrinkRequest
	GetFacePictureBase64() *string
	SetFacePictureUrl(v string) *InitializeShrinkRequest
	GetFacePictureUrl() *string
	SetIdFaceQuality(v string) *InitializeShrinkRequest
	GetIdFaceQuality() *string
	SetIdSpoof(v string) *InitializeShrinkRequest
	GetIdSpoof() *string
	SetIdThreshold(v string) *InitializeShrinkRequest
	GetIdThreshold() *string
	SetLanguageConfig(v string) *InitializeShrinkRequest
	GetLanguageConfig() *string
	SetMRTDInput(v string) *InitializeShrinkRequest
	GetMRTDInput() *string
	SetMerchantBizId(v string) *InitializeShrinkRequest
	GetMerchantBizId() *string
	SetMerchantUserId(v string) *InitializeShrinkRequest
	GetMerchantUserId() *string
	SetMetaInfo(v string) *InitializeShrinkRequest
	GetMetaInfo() *string
	SetModel(v string) *InitializeShrinkRequest
	GetModel() *string
	SetOcr(v string) *InitializeShrinkRequest
	GetOcr() *string
	SetPages(v string) *InitializeShrinkRequest
	GetPages() *string
	SetProcedurePriority(v string) *InitializeShrinkRequest
	GetProcedurePriority() *string
	SetProductCode(v string) *InitializeShrinkRequest
	GetProductCode() *string
	SetProductFlow(v string) *InitializeShrinkRequest
	GetProductFlow() *string
	SetReturnUrl(v string) *InitializeShrinkRequest
	GetReturnUrl() *string
	SetSceneCode(v string) *InitializeShrinkRequest
	GetSceneCode() *string
	SetSecurityLevel(v string) *InitializeShrinkRequest
	GetSecurityLevel() *string
	SetShowAlbumIcon(v string) *InitializeShrinkRequest
	GetShowAlbumIcon() *string
	SetShowGuidePage(v string) *InitializeShrinkRequest
	GetShowGuidePage() *string
	SetShowOcrResult(v string) *InitializeShrinkRequest
	GetShowOcrResult() *string
	SetStyleConfig(v string) *InitializeShrinkRequest
	GetStyleConfig() *string
	SetUseNFC(v string) *InitializeShrinkRequest
	GetUseNFC() *string
}

type InitializeShrinkRequest struct {
	AppQualityCheck      *string `json:"AppQualityCheck,omitempty" xml:"AppQualityCheck,omitempty"`
	Authorize            *string `json:"Authorize,omitempty" xml:"Authorize,omitempty"`
	CallbackToken        *string `json:"CallbackToken,omitempty" xml:"CallbackToken,omitempty"`
	CallbackUrl          *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	ChameleonFrameEnable *string `json:"ChameleonFrameEnable,omitempty" xml:"ChameleonFrameEnable,omitempty"`
	// example:
	//
	// *
	Crop                *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	DateOfBirth         *string `json:"DateOfBirth,omitempty" xml:"DateOfBirth,omitempty"`
	DateOfExpiry        *string `json:"DateOfExpiry,omitempty" xml:"DateOfExpiry,omitempty"`
	DocName             *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
	DocNo               *string `json:"DocNo,omitempty" xml:"DocNo,omitempty"`
	DocPageConfigShrink *string `json:"DocPageConfig,omitempty" xml:"DocPageConfig,omitempty"`
	DocScanMode         *string `json:"DocScanMode,omitempty" xml:"DocScanMode,omitempty"`
	// example:
	//
	// 01000000
	DocType           *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	DocVideo          *string `json:"DocVideo,omitempty" xml:"DocVideo,omitempty"`
	DocumentNumber    *string `json:"DocumentNumber,omitempty" xml:"DocumentNumber,omitempty"`
	ExperienceCode    *string `json:"ExperienceCode,omitempty" xml:"ExperienceCode,omitempty"`
	FacePictureBase64 *string `json:"FacePictureBase64,omitempty" xml:"FacePictureBase64,omitempty"`
	// example:
	//
	// ***
	FacePictureUrl *string `json:"FacePictureUrl,omitempty" xml:"FacePictureUrl,omitempty"`
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
	// http*****
	ReturnUrl *string `json:"ReturnUrl,omitempty" xml:"ReturnUrl,omitempty"`
	// example:
	//
	// PAY**
	SceneCode     *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	SecurityLevel *string `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	ShowAlbumIcon *string `json:"ShowAlbumIcon,omitempty" xml:"ShowAlbumIcon,omitempty"`
	ShowGuidePage *string `json:"ShowGuidePage,omitempty" xml:"ShowGuidePage,omitempty"`
	ShowOcrResult *string `json:"ShowOcrResult,omitempty" xml:"ShowOcrResult,omitempty"`
	StyleConfig   *string `json:"StyleConfig,omitempty" xml:"StyleConfig,omitempty"`
	UseNFC        *string `json:"UseNFC,omitempty" xml:"UseNFC,omitempty"`
}

func (s InitializeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s InitializeShrinkRequest) GoString() string {
	return s.String()
}

func (s *InitializeShrinkRequest) GetAppQualityCheck() *string {
	return s.AppQualityCheck
}

func (s *InitializeShrinkRequest) GetAuthorize() *string {
	return s.Authorize
}

func (s *InitializeShrinkRequest) GetCallbackToken() *string {
	return s.CallbackToken
}

func (s *InitializeShrinkRequest) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *InitializeShrinkRequest) GetChameleonFrameEnable() *string {
	return s.ChameleonFrameEnable
}

func (s *InitializeShrinkRequest) GetCrop() *string {
	return s.Crop
}

func (s *InitializeShrinkRequest) GetDateOfBirth() *string {
	return s.DateOfBirth
}

func (s *InitializeShrinkRequest) GetDateOfExpiry() *string {
	return s.DateOfExpiry
}

func (s *InitializeShrinkRequest) GetDocName() *string {
	return s.DocName
}

func (s *InitializeShrinkRequest) GetDocNo() *string {
	return s.DocNo
}

func (s *InitializeShrinkRequest) GetDocPageConfigShrink() *string {
	return s.DocPageConfigShrink
}

func (s *InitializeShrinkRequest) GetDocScanMode() *string {
	return s.DocScanMode
}

func (s *InitializeShrinkRequest) GetDocType() *string {
	return s.DocType
}

func (s *InitializeShrinkRequest) GetDocVideo() *string {
	return s.DocVideo
}

func (s *InitializeShrinkRequest) GetDocumentNumber() *string {
	return s.DocumentNumber
}

func (s *InitializeShrinkRequest) GetExperienceCode() *string {
	return s.ExperienceCode
}

func (s *InitializeShrinkRequest) GetFacePictureBase64() *string {
	return s.FacePictureBase64
}

func (s *InitializeShrinkRequest) GetFacePictureUrl() *string {
	return s.FacePictureUrl
}

func (s *InitializeShrinkRequest) GetIdFaceQuality() *string {
	return s.IdFaceQuality
}

func (s *InitializeShrinkRequest) GetIdSpoof() *string {
	return s.IdSpoof
}

func (s *InitializeShrinkRequest) GetIdThreshold() *string {
	return s.IdThreshold
}

func (s *InitializeShrinkRequest) GetLanguageConfig() *string {
	return s.LanguageConfig
}

func (s *InitializeShrinkRequest) GetMRTDInput() *string {
	return s.MRTDInput
}

func (s *InitializeShrinkRequest) GetMerchantBizId() *string {
	return s.MerchantBizId
}

func (s *InitializeShrinkRequest) GetMerchantUserId() *string {
	return s.MerchantUserId
}

func (s *InitializeShrinkRequest) GetMetaInfo() *string {
	return s.MetaInfo
}

func (s *InitializeShrinkRequest) GetModel() *string {
	return s.Model
}

func (s *InitializeShrinkRequest) GetOcr() *string {
	return s.Ocr
}

func (s *InitializeShrinkRequest) GetPages() *string {
	return s.Pages
}

func (s *InitializeShrinkRequest) GetProcedurePriority() *string {
	return s.ProcedurePriority
}

func (s *InitializeShrinkRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *InitializeShrinkRequest) GetProductFlow() *string {
	return s.ProductFlow
}

func (s *InitializeShrinkRequest) GetReturnUrl() *string {
	return s.ReturnUrl
}

func (s *InitializeShrinkRequest) GetSceneCode() *string {
	return s.SceneCode
}

func (s *InitializeShrinkRequest) GetSecurityLevel() *string {
	return s.SecurityLevel
}

func (s *InitializeShrinkRequest) GetShowAlbumIcon() *string {
	return s.ShowAlbumIcon
}

func (s *InitializeShrinkRequest) GetShowGuidePage() *string {
	return s.ShowGuidePage
}

func (s *InitializeShrinkRequest) GetShowOcrResult() *string {
	return s.ShowOcrResult
}

func (s *InitializeShrinkRequest) GetStyleConfig() *string {
	return s.StyleConfig
}

func (s *InitializeShrinkRequest) GetUseNFC() *string {
	return s.UseNFC
}

func (s *InitializeShrinkRequest) SetAppQualityCheck(v string) *InitializeShrinkRequest {
	s.AppQualityCheck = &v
	return s
}

func (s *InitializeShrinkRequest) SetAuthorize(v string) *InitializeShrinkRequest {
	s.Authorize = &v
	return s
}

func (s *InitializeShrinkRequest) SetCallbackToken(v string) *InitializeShrinkRequest {
	s.CallbackToken = &v
	return s
}

func (s *InitializeShrinkRequest) SetCallbackUrl(v string) *InitializeShrinkRequest {
	s.CallbackUrl = &v
	return s
}

func (s *InitializeShrinkRequest) SetChameleonFrameEnable(v string) *InitializeShrinkRequest {
	s.ChameleonFrameEnable = &v
	return s
}

func (s *InitializeShrinkRequest) SetCrop(v string) *InitializeShrinkRequest {
	s.Crop = &v
	return s
}

func (s *InitializeShrinkRequest) SetDateOfBirth(v string) *InitializeShrinkRequest {
	s.DateOfBirth = &v
	return s
}

func (s *InitializeShrinkRequest) SetDateOfExpiry(v string) *InitializeShrinkRequest {
	s.DateOfExpiry = &v
	return s
}

func (s *InitializeShrinkRequest) SetDocName(v string) *InitializeShrinkRequest {
	s.DocName = &v
	return s
}

func (s *InitializeShrinkRequest) SetDocNo(v string) *InitializeShrinkRequest {
	s.DocNo = &v
	return s
}

func (s *InitializeShrinkRequest) SetDocPageConfigShrink(v string) *InitializeShrinkRequest {
	s.DocPageConfigShrink = &v
	return s
}

func (s *InitializeShrinkRequest) SetDocScanMode(v string) *InitializeShrinkRequest {
	s.DocScanMode = &v
	return s
}

func (s *InitializeShrinkRequest) SetDocType(v string) *InitializeShrinkRequest {
	s.DocType = &v
	return s
}

func (s *InitializeShrinkRequest) SetDocVideo(v string) *InitializeShrinkRequest {
	s.DocVideo = &v
	return s
}

func (s *InitializeShrinkRequest) SetDocumentNumber(v string) *InitializeShrinkRequest {
	s.DocumentNumber = &v
	return s
}

func (s *InitializeShrinkRequest) SetExperienceCode(v string) *InitializeShrinkRequest {
	s.ExperienceCode = &v
	return s
}

func (s *InitializeShrinkRequest) SetFacePictureBase64(v string) *InitializeShrinkRequest {
	s.FacePictureBase64 = &v
	return s
}

func (s *InitializeShrinkRequest) SetFacePictureUrl(v string) *InitializeShrinkRequest {
	s.FacePictureUrl = &v
	return s
}

func (s *InitializeShrinkRequest) SetIdFaceQuality(v string) *InitializeShrinkRequest {
	s.IdFaceQuality = &v
	return s
}

func (s *InitializeShrinkRequest) SetIdSpoof(v string) *InitializeShrinkRequest {
	s.IdSpoof = &v
	return s
}

func (s *InitializeShrinkRequest) SetIdThreshold(v string) *InitializeShrinkRequest {
	s.IdThreshold = &v
	return s
}

func (s *InitializeShrinkRequest) SetLanguageConfig(v string) *InitializeShrinkRequest {
	s.LanguageConfig = &v
	return s
}

func (s *InitializeShrinkRequest) SetMRTDInput(v string) *InitializeShrinkRequest {
	s.MRTDInput = &v
	return s
}

func (s *InitializeShrinkRequest) SetMerchantBizId(v string) *InitializeShrinkRequest {
	s.MerchantBizId = &v
	return s
}

func (s *InitializeShrinkRequest) SetMerchantUserId(v string) *InitializeShrinkRequest {
	s.MerchantUserId = &v
	return s
}

func (s *InitializeShrinkRequest) SetMetaInfo(v string) *InitializeShrinkRequest {
	s.MetaInfo = &v
	return s
}

func (s *InitializeShrinkRequest) SetModel(v string) *InitializeShrinkRequest {
	s.Model = &v
	return s
}

func (s *InitializeShrinkRequest) SetOcr(v string) *InitializeShrinkRequest {
	s.Ocr = &v
	return s
}

func (s *InitializeShrinkRequest) SetPages(v string) *InitializeShrinkRequest {
	s.Pages = &v
	return s
}

func (s *InitializeShrinkRequest) SetProcedurePriority(v string) *InitializeShrinkRequest {
	s.ProcedurePriority = &v
	return s
}

func (s *InitializeShrinkRequest) SetProductCode(v string) *InitializeShrinkRequest {
	s.ProductCode = &v
	return s
}

func (s *InitializeShrinkRequest) SetProductFlow(v string) *InitializeShrinkRequest {
	s.ProductFlow = &v
	return s
}

func (s *InitializeShrinkRequest) SetReturnUrl(v string) *InitializeShrinkRequest {
	s.ReturnUrl = &v
	return s
}

func (s *InitializeShrinkRequest) SetSceneCode(v string) *InitializeShrinkRequest {
	s.SceneCode = &v
	return s
}

func (s *InitializeShrinkRequest) SetSecurityLevel(v string) *InitializeShrinkRequest {
	s.SecurityLevel = &v
	return s
}

func (s *InitializeShrinkRequest) SetShowAlbumIcon(v string) *InitializeShrinkRequest {
	s.ShowAlbumIcon = &v
	return s
}

func (s *InitializeShrinkRequest) SetShowGuidePage(v string) *InitializeShrinkRequest {
	s.ShowGuidePage = &v
	return s
}

func (s *InitializeShrinkRequest) SetShowOcrResult(v string) *InitializeShrinkRequest {
	s.ShowOcrResult = &v
	return s
}

func (s *InitializeShrinkRequest) SetStyleConfig(v string) *InitializeShrinkRequest {
	s.StyleConfig = &v
	return s
}

func (s *InitializeShrinkRequest) SetUseNFC(v string) *InitializeShrinkRequest {
	s.UseNFC = &v
	return s
}

func (s *InitializeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
