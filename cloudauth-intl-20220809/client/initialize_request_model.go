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
	SetEmail(v string) *InitializeRequest
	GetEmail() *string
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
	SetMobile(v string) *InitializeRequest
	GetMobile() *string
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
	DocPageConfig []*string `json:"DocPageConfig,omitempty" xml:"DocPageConfig,omitempty" type:"Repeated"`
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

func (s *InitializeRequest) GetEmail() *string {
	return s.Email
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

func (s *InitializeRequest) GetMobile() *string {
	return s.Mobile
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

func (s *InitializeRequest) SetEmail(v string) *InitializeRequest {
	s.Email = &v
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

func (s *InitializeRequest) SetMobile(v string) *InitializeRequest {
	s.Mobile = &v
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
