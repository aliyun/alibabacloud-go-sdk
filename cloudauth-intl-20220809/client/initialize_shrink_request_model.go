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
	SetAutoRegistration(v string) *InitializeShrinkRequest
	GetAutoRegistration() *string
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
	SetEditOcrResult(v string) *InitializeShrinkRequest
	GetEditOcrResult() *string
	SetEmail(v string) *InitializeShrinkRequest
	GetEmail() *string
	SetExperienceCode(v string) *InitializeShrinkRequest
	GetExperienceCode() *string
	SetFaceGroupCodes(v string) *InitializeShrinkRequest
	GetFaceGroupCodes() *string
	SetFacePictureBase64(v string) *InitializeShrinkRequest
	GetFacePictureBase64() *string
	SetFacePictureUrl(v string) *InitializeShrinkRequest
	GetFacePictureUrl() *string
	SetFaceRegisterGroupCode(v string) *InitializeShrinkRequest
	GetFaceRegisterGroupCode() *string
	SetFaceVerifyThreshold(v string) *InitializeShrinkRequest
	GetFaceVerifyThreshold() *string
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
	SetMobile(v string) *InitializeShrinkRequest
	GetMobile() *string
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
	SetReturnFaces(v string) *InitializeShrinkRequest
	GetReturnFaces() *string
	SetReturnUrl(v string) *InitializeShrinkRequest
	GetReturnUrl() *string
	SetSaveFacePicture(v string) *InitializeShrinkRequest
	GetSaveFacePicture() *string
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
	SetTargetFacePicture(v string) *InitializeShrinkRequest
	GetTargetFacePicture() *string
	SetTargetFacePictureUrl(v string) *InitializeShrinkRequest
	GetTargetFacePictureUrl() *string
	SetUseNFC(v string) *InitializeShrinkRequest
	GetUseNFC() *string
	SetVerifyModel(v string) *InitializeShrinkRequest
	GetVerifyModel() *string
}

type InitializeShrinkRequest struct {
	// <warning>This feature is currently not supported by **Web SDK**. Please refer to the App SDK integration if needed.</warning>
	//
	// Whether to enable strict face quality detection:
	//
	// - Y: Enable (default)
	//
	// - N: Disable
	//
	// example:
	//
	// N
	AppQualityCheck *string `json:"AppQualityCheck,omitempty" xml:"AppQualityCheck,omitempty"`
	// Whether to enable authoritative identity verification, currently applicable only to the second-generation ID card in mainland China. (IDV product input parameter)
	//
	// example:
	//
	// Y
	Authorize *string `json:"Authorize,omitempty" xml:"Authorize,omitempty"`
	// Whether to enable automatic registration
	//
	// example:
	//
	// 0
	AutoRegistration *string `json:"AutoRegistration,omitempty" xml:"AutoRegistration,omitempty"`
	// Security Token, used for preventing duplication and tampering. If this parameter is passed, the CallbackToken field will be displayed in the callback address.
	//
	// example:
	//
	// 7ca5c68d869344ea8eeb30cdfd544544-6358700
	CallbackToken *string `json:"CallbackToken,omitempty" xml:"CallbackToken,omitempty"`
	// Callback notification address for authentication results. The default callback request method is GET, and the callback address must start with https. After completing the authentication, the platform will call back this address and automatically add the transactionId, passed, and subcode fields.
	//
	// example:
	//
	// https://www.aliyun.com?callbackToken=1000004826&transactionId=shaxxxx&passed=Y&subCode=200
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// Whether to enable adaptive color-changing window border
	//
	// - **Y**: Enable
	//
	// - **N**: Disable
	//
	// example:
	//
	// N
	ChameleonFrameEnable *string `json:"ChameleonFrameEnable,omitempty" xml:"ChameleonFrameEnable,omitempty"`
	// Whether to crop. (IDV product input parameter)
	//
	// example:
	//
	// N
	Crop *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	// Date of birth on the document
	//
	// **MRTDInput = 2*	- is required.
	//
	// example:
	//
	// -
	DateOfBirth *string `json:"DateOfBirth,omitempty" xml:"DateOfBirth,omitempty"`
	// Expiration date on the document
	//
	// **MRTDInput = 2*	- is required.
	//
	// example:
	//
	// -
	DateOfExpiry *string `json:"DateOfExpiry,omitempty" xml:"DateOfExpiry,omitempty"`
	// User\\"s real name.
	//
	// example:
	//
	// 张三
	DocName *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
	// User\\"s document number.
	//
	// example:
	//
	// 411xxxxxxxxxxx0001
	DocNo *string `json:"DocNo,omitempty" xml:"DocNo,omitempty"`
	// Customer-defined input to specify whether to collect more pages
	DocPageConfigShrink *string `json:"DocPageConfig,omitempty" xml:"DocPageConfig,omitempty"`
	// Document capture mode.
	//
	// - manual: Manual capture.
	//
	// - auto: Automatic capture (default)
	//
	// example:
	//
	// manual
	DocScanMode *string `json:"DocScanMode,omitempty" xml:"DocScanMode,omitempty"`
	// Document type, uniquely identified by an 8-digit combination.
	//
	// Note: This parameter is required only when ProductCode is KYC_GLOBAL, OCR_GLOBAL, or IDR_GLOBAL.
	//
	// example:
	//
	// ​01560000
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// Whether to require a video for evidence.
	//
	// - N: Not required (default).
	//
	// - Y: During the authentication process, a 1~2 second video of the user\\"s face will be captured and returned via the query interface.
	//
	// > Due to the large size of the video file, the system may discard it when the network is unstable, prioritizing the transmission of necessary images for authentication.
	//
	// example:
	//
	// N
	DocVideo *string `json:"DocVideo,omitempty" xml:"DocVideo,omitempty"`
	// Document number
	//
	// **MRTDInput = 2*	- is required.
	//
	// example:
	//
	// -
	DocumentNumber *string `json:"DocumentNumber,omitempty" xml:"DocumentNumber,omitempty"`
	// In the document OCR recognition step, whether the recognition result page is editable:
	//
	// - **0**: Not editable
	//
	// - **1*	- (default): Editable
	//
	// example:
	//
	// 0
	EditOcrResult *string `json:"EditOcrResult,omitempty" xml:"EditOcrResult,omitempty"`
	Email         *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// Experience code
	//
	// example:
	//
	// 9be7b7d0180041219e5ab03ac6dab5fb
	ExperienceCode *string `json:"ExperienceCode,omitempty" xml:"ExperienceCode,omitempty"`
	// Face library to be compared
	//
	// example:
	//
	// 0e0c34a77f
	FaceGroupCodes *string `json:"FaceGroupCodes,omitempty" xml:"FaceGroupCodes,omitempty"`
	// Base64 encoded face image. If you choose to pass the face image via FacePictureBase64, please check the image size and do not upload images larger than 1 MB.
	//
	// When productCode is FV_GLOBAL, choose one of the parameters between FacePictureBase64 and FacePictureUrl to pass in.
	//
	// example:
	//
	// Base64
	FacePictureBase64 *string `json:"FacePictureBase64,omitempty" xml:"FacePictureBase64,omitempty"`
	// Face image URL. A publicly accessible HTTP or HTTPS link. When productCode is FV_GLOBAL, choose one of the parameters between FacePictureUrl and FacePictureBase to pass in.
	//
	// example:
	//
	// ***
	FacePictureUrl *string `json:"FacePictureUrl,omitempty" xml:"FacePictureUrl,omitempty"`
	// Face library for registration.
	//
	// example:
	//
	// 0e0c34a77f
	FaceRegisterGroupCode *string `json:"FaceRegisterGroupCode,omitempty" xml:"FaceRegisterGroupCode,omitempty"`
	// Face verification threshold
	//
	// example:
	//
	// 0.5
	FaceVerifyThreshold *string `json:"FaceVerifyThreshold,omitempty" xml:"FaceVerifyThreshold,omitempty"`
	// Face image quality. (IDV product input parameter)
	//
	// example:
	//
	// Y
	IdFaceQuality *string `json:"IdFaceQuality,omitempty" xml:"IdFaceQuality,omitempty"`
	// Whether to enable certificate anti-counterfeiting detection. (IDV product input parameter)
	//
	// example:
	//
	// Y
	IdSpoof *string `json:"IdSpoof,omitempty" xml:"IdSpoof,omitempty"`
	// Custom OCR quality detection threshold mode:
	//
	// - **0**: Standard mode
	//
	// - **1**: Strict mode
	//
	// - **2**: Lenient mode
	//
	// - **3*	- (default): Disable quality detection
	//
	// example:
	//
	// 0
	IdThreshold *string `json:"IdThreshold,omitempty" xml:"IdThreshold,omitempty"`
	// Language configuration. (IDV product input parameter)
	//
	// example:
	//
	// en
	LanguageConfig *string `json:"LanguageConfig,omitempty" xml:"LanguageConfig,omitempty"`
	// Source of MRTD verification parameters. This parameter is required to decrypt information when reading the document chip via NFC.
	//
	// - **0**: User input
	//
	// - **1**: OCR read
	//
	// - **2**: Passed through the API
	//
	// example:
	//
	// 0
	MRTDInput *string `json:"MRTDInput,omitempty" xml:"MRTDInput,omitempty"`
	// A unique business identifier defined by the merchant, used for subsequent troubleshooting. It supports a combination of letters and numbers, with a maximum length of 32 characters. Please ensure its uniqueness.
	//
	// example:
	//
	// e0c34a***353888
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// Your custom user ID or other identifiers that can recognize specific users, such as phone numbers or email addresses. It is strongly recommended to pre-desensitize the value of this field, for example, by hashing it.
	//
	// example:
	//
	// 1221****6543
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// Metainfo environment parameter, which needs to be obtained through the client SDK.
	//
	// example:
	//
	// {\\"bioMetaInfo\\":\\"4.1.0:2916352,0\\",\\"deviceType\\":\\"web\\",\\"ua\\":\\"Mozilla/5.0 (Macintosh
	MetaInfo *string `json:"MetaInfo,omitempty" xml:"MetaInfo,omitempty"`
	Mobile   *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The type of liveness detection to be performed:
	//
	// - **LIVENESS*	- (default): Blinking action liveness detection.
	//
	// - **PHOTINUS_LIVENESS**: Blinking action liveness + photinus liveness dual detection.
	//
	// >
	//
	// > - For supported SDK versions, see [SDK Publishing Record](https://www.alibabacloud.com/help/zh/ekyc/latest/sdk-publishing-record?spm=a2c63.p38356.0.i99).
	//
	// > - PC does not support photinus liveness dual detection.
	//
	// example:
	//
	// PHOTINUS_LIVENESS
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// Whether to enable OCR. (IDV product input parameter)
	//
	// example:
	//
	// Y
	Ocr *string `json:"Ocr,omitempty" xml:"Ocr,omitempty"`
	// Page configuration for collection, multiple pages are connected using commas. The value range is as follows:
	//
	// - **01**: Front side of the document
	//
	// - **01,02**: Front and back sides of the document
	//
	// > When this value is 01,02, currently only Chinese and Vietnamese IDs are supported.
	//
	// example:
	//
	// 01
	Pages *string `json:"Pages,omitempty" xml:"Pages,omitempty"`
	// When compatibility issues occur with H5-based mobile authentication, whether to allow a fallback handling method.
	//
	// - **url*	- (default): Support fallback. The page displays the authentication URL, which users can copy and open in another browser to continue the authentication process.
	//
	// - **keep**: Do not support fallback. Directly return the error reason and end the authentication process.
	//
	//
	// >
	//
	// > - This switch is not supported on PC.
	//
	// > - If the business scenario involves completing authentication through an embedded web page in an app, it is recommended to set this parameter to `keep` to disallow URL fallback.
	//
	// example:
	//
	// url
	ProcedurePriority *string `json:"ProcedurePriority,omitempty" xml:"ProcedurePriority,omitempty"`
	// The product solution to be integrated. The values are as follows:
	//
	// - KYC_GLOBAL (eKYC product solution)
	//
	// - FV_GLOBAL (Live Face Verification)
	//
	// - FL_GLOBAL (Liveness Detection)
	//
	// - IDR_GLOBAL (Single Document Verification)
	//
	// - OCR_GLOBAL (Single Document OCR)
	//
	// example:
	//
	// KYC_GLOBAL
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// Supports card and face sequential arrangement:
	//
	// - DOC_FACE (default)
	//
	// - FACE_DOC
	//
	// Note: This parameter is required only when ProductCode is KYC_GLOBAL.
	//
	// example:
	//
	// DOC_FACE
	ProductFlow *string `json:"ProductFlow,omitempty" xml:"ProductFlow,omitempty"`
	// Number of duplicate faces returned
	//
	// example:
	//
	// 1
	ReturnFaces *string `json:"ReturnFaces,omitempty" xml:"ReturnFaces,omitempty"`
	// Client-side callback address.
	//
	// example:
	//
	// http*****
	ReturnUrl *string `json:"ReturnUrl,omitempty" xml:"ReturnUrl,omitempty"`
	// Whether to save the face image
	//
	// example:
	//
	// 0
	SaveFacePicture *string `json:"SaveFacePicture,omitempty" xml:"SaveFacePicture,omitempty"`
	// Scene code. (IDV product input parameter)
	//
	// example:
	//
	// 123****123
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// Represents different security levels in the authentication process. The available values are as follows:
	//
	// 01: Normal mode (default).
	//
	// 02: Secure mode, a relatively strict mode, suitable for high-risk scenarios. (IDV product input parameter)
	//
	// example:
	//
	// 01
	SecurityLevel *string `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	// In the document OCR recognition step, whether to display the album upload entry:
	//
	// - **1**: Display (default)
	//
	// - **0**: Do not display
	//
	// example:
	//
	// 1
	ShowAlbumIcon *string `json:"ShowAlbumIcon,omitempty" xml:"ShowAlbumIcon,omitempty"`
	// Switch to control whether to display the guide page:
	//
	// - **1**: Display (default)
	//
	// - **0**: Do not display
	//
	// example:
	//
	// 1
	ShowGuidePage *string `json:"ShowGuidePage,omitempty" xml:"ShowGuidePage,omitempty"`
	// In the document OCR recognition step, whether to display the recognition result page:
	//
	// - **1**: Display (default)
	//
	// - **0**: Do not display
	//
	// example:
	//
	// 1
	ShowOcrResult *string `json:"ShowOcrResult,omitempty" xml:"ShowOcrResult,omitempty"`
	// Custom UI configuration. Based on the configuration template, convert your custom UI configuration into a JSON string and pass it through this interface. For more information, see [IDV UI Customization](https://www.alibabacloud.com/help/zh/ekyc/latest/idv-kyc-custom-skin?spm=a2c63.p38356.0.i60).
	//
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
	// Base64 encoding of the portrait photo.
	//
	// example:
	//
	// base64
	TargetFacePicture *string `json:"TargetFacePicture,omitempty" xml:"TargetFacePicture,omitempty"`
	// Portrait image URL, accessible via public HTTP or HTTPS link.
	//
	// example:
	//
	// https://www.xxxxx.com/1.jpg
	TargetFacePictureUrl *string `json:"TargetFacePictureUrl,omitempty" xml:"TargetFacePictureUrl,omitempty"`
	// When **DocType**=01000000 (global passport), you can choose whether to enable NFC verification.
	//
	// - **Y*	- (enable)
	//
	// - **N*	- (disable)
	//
	// example:
	//
	// N
	UseNFC *string `json:"UseNFC,omitempty" xml:"UseNFC,omitempty"`
	// Type of verification
	//
	// example:
	//
	// 0
	VerifyModel *string `json:"VerifyModel,omitempty" xml:"VerifyModel,omitempty"`
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

func (s *InitializeShrinkRequest) GetAutoRegistration() *string {
	return s.AutoRegistration
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

func (s *InitializeShrinkRequest) GetEditOcrResult() *string {
	return s.EditOcrResult
}

func (s *InitializeShrinkRequest) GetEmail() *string {
	return s.Email
}

func (s *InitializeShrinkRequest) GetExperienceCode() *string {
	return s.ExperienceCode
}

func (s *InitializeShrinkRequest) GetFaceGroupCodes() *string {
	return s.FaceGroupCodes
}

func (s *InitializeShrinkRequest) GetFacePictureBase64() *string {
	return s.FacePictureBase64
}

func (s *InitializeShrinkRequest) GetFacePictureUrl() *string {
	return s.FacePictureUrl
}

func (s *InitializeShrinkRequest) GetFaceRegisterGroupCode() *string {
	return s.FaceRegisterGroupCode
}

func (s *InitializeShrinkRequest) GetFaceVerifyThreshold() *string {
	return s.FaceVerifyThreshold
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

func (s *InitializeShrinkRequest) GetMobile() *string {
	return s.Mobile
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

func (s *InitializeShrinkRequest) GetReturnFaces() *string {
	return s.ReturnFaces
}

func (s *InitializeShrinkRequest) GetReturnUrl() *string {
	return s.ReturnUrl
}

func (s *InitializeShrinkRequest) GetSaveFacePicture() *string {
	return s.SaveFacePicture
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

func (s *InitializeShrinkRequest) GetTargetFacePicture() *string {
	return s.TargetFacePicture
}

func (s *InitializeShrinkRequest) GetTargetFacePictureUrl() *string {
	return s.TargetFacePictureUrl
}

func (s *InitializeShrinkRequest) GetUseNFC() *string {
	return s.UseNFC
}

func (s *InitializeShrinkRequest) GetVerifyModel() *string {
	return s.VerifyModel
}

func (s *InitializeShrinkRequest) SetAppQualityCheck(v string) *InitializeShrinkRequest {
	s.AppQualityCheck = &v
	return s
}

func (s *InitializeShrinkRequest) SetAuthorize(v string) *InitializeShrinkRequest {
	s.Authorize = &v
	return s
}

func (s *InitializeShrinkRequest) SetAutoRegistration(v string) *InitializeShrinkRequest {
	s.AutoRegistration = &v
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

func (s *InitializeShrinkRequest) SetEditOcrResult(v string) *InitializeShrinkRequest {
	s.EditOcrResult = &v
	return s
}

func (s *InitializeShrinkRequest) SetEmail(v string) *InitializeShrinkRequest {
	s.Email = &v
	return s
}

func (s *InitializeShrinkRequest) SetExperienceCode(v string) *InitializeShrinkRequest {
	s.ExperienceCode = &v
	return s
}

func (s *InitializeShrinkRequest) SetFaceGroupCodes(v string) *InitializeShrinkRequest {
	s.FaceGroupCodes = &v
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

func (s *InitializeShrinkRequest) SetFaceRegisterGroupCode(v string) *InitializeShrinkRequest {
	s.FaceRegisterGroupCode = &v
	return s
}

func (s *InitializeShrinkRequest) SetFaceVerifyThreshold(v string) *InitializeShrinkRequest {
	s.FaceVerifyThreshold = &v
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

func (s *InitializeShrinkRequest) SetMobile(v string) *InitializeShrinkRequest {
	s.Mobile = &v
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

func (s *InitializeShrinkRequest) SetReturnFaces(v string) *InitializeShrinkRequest {
	s.ReturnFaces = &v
	return s
}

func (s *InitializeShrinkRequest) SetReturnUrl(v string) *InitializeShrinkRequest {
	s.ReturnUrl = &v
	return s
}

func (s *InitializeShrinkRequest) SetSaveFacePicture(v string) *InitializeShrinkRequest {
	s.SaveFacePicture = &v
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

func (s *InitializeShrinkRequest) SetTargetFacePicture(v string) *InitializeShrinkRequest {
	s.TargetFacePicture = &v
	return s
}

func (s *InitializeShrinkRequest) SetTargetFacePictureUrl(v string) *InitializeShrinkRequest {
	s.TargetFacePictureUrl = &v
	return s
}

func (s *InitializeShrinkRequest) SetUseNFC(v string) *InitializeShrinkRequest {
	s.UseNFC = &v
	return s
}

func (s *InitializeShrinkRequest) SetVerifyModel(v string) *InitializeShrinkRequest {
	s.VerifyModel = &v
	return s
}

func (s *InitializeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
