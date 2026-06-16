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
	SetOcrValueStandard(v string) *InitializeShrinkRequest
	GetOcrValueStandard() *string
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
	SetTemplateConfig(v string) *InitializeShrinkRequest
	GetTemplateConfig() *string
	SetTemplateRanCount(v string) *InitializeShrinkRequest
	GetTemplateRanCount() *string
	SetTemplateType(v string) *InitializeShrinkRequest
	GetTemplateType() *string
	SetUseNFC(v string) *InitializeShrinkRequest
	GetUseNFC() *string
	SetVerifyModel(v string) *InitializeShrinkRequest
	GetVerifyModel() *string
}

type InitializeShrinkRequest struct {
	// <warning>This feature is **not supported by the Web SDK**. To use this feature, integrate the App SDK.</warning>
	//
	// Specifies whether to enable strict face quality detection:
	//
	// - Y: enabled (default)
	//
	// - N: disabled.
	//
	// example:
	//
	// N
	AppQualityCheck *string `json:"AppQualityCheck,omitempty" xml:"AppQualityCheck,omitempty"`
	// Specifies whether to enable authoritative identity verification. Currently, this applies only to second-generation ID cards in the Chinese mainland. This is an input parameter for the IDV product.
	//
	// example:
	//
	// Y
	Authorize *string `json:"Authorize,omitempty" xml:"Authorize,omitempty"`
	// Specifies whether to enable auto-registration.
	//
	// example:
	//
	// 0
	AutoRegistration *string `json:"AutoRegistration,omitempty" xml:"AutoRegistration,omitempty"`
	// The security token used for anti-replay and anti-tampering verification. If this parameter is specified, the CallbackToken field is included in the callback URL.
	//
	// example:
	//
	// 7ca5c68d869344ea8eeb30cdfd544544-6358700
	CallbackToken *string `json:"CallbackToken,omitempty" xml:"CallbackToken,omitempty"`
	// The callback URL for the authentication result. The callback request method is GET by default. The callback URL must start with https. After the authentication is complete, the platform calls back this URL and automatically appends the transactionId, passed, and subcode fields.
	//
	// example:
	//
	// https://www.aliyun.com?callbackToken=1000004826&transactionId=shaxxxx&passed=Y&subCode=200
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// Specifies whether to enable the adaptive color-changing window frame.
	//
	// - **Y**: enabled
	//
	// - **N**: disabled.
	//
	// example:
	//
	// N
	ChameleonFrameEnable *string `json:"ChameleonFrameEnable,omitempty" xml:"ChameleonFrameEnable,omitempty"`
	// Specifies whether to enable cropping. This is an input parameter for the IDV product.
	//
	// example:
	//
	// N
	Crop *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	// The date of birth on the document.
	//
	// Required when **MRTDInput = 2**.
	//
	// example:
	//
	// -
	DateOfBirth *string `json:"DateOfBirth,omitempty" xml:"DateOfBirth,omitempty"`
	// The expiration date on the document.
	//
	// Required when **MRTDInput*	- = 2.
	//
	// example:
	//
	// -
	DateOfExpiry *string `json:"DateOfExpiry,omitempty" xml:"DateOfExpiry,omitempty"`
	// The real name of the user.
	//
	// example:
	//
	// 张三
	DocName *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
	// The document number of the user.
	//
	// example:
	//
	// 411xxxxxxxxxxx0001
	DocNo *string `json:"DocNo,omitempty" xml:"DocNo,omitempty"`
	// The custom configuration for whether to capture additional pages.
	DocPageConfigShrink *string `json:"DocPageConfig,omitempty" xml:"DocPageConfig,omitempty"`
	// The document capture photo mode.
	//
	// - manual: manual photo capture.
	//
	// - auto: automatic photo capture (default).
	//
	// example:
	//
	// manual
	DocScanMode *string `json:"DocScanMode,omitempty" xml:"DocScanMode,omitempty"`
	// The document type.
	//
	// >For the eKYC_PRO and ID_OCR_MAX solutions, see the official documentation: https://www.alibabacloud.com/help/zh/ekyc/latest/certificate-code-table?spm=a2c63.p38356.help-menu-445633.d_2_8_2_0.279147abwKAWbr
	//
	// >For the ID_OCR, eKYC, and eKYC_MIN solutions, see the document type list in the official documentation: https://www.alibabacloud.com/help/zh/ekyc/latest/gnhekqy05ni51m4c?spm=a2c63.p38356.help-menu-445633.d_2_3_1_0_0_0.6243244777KoZ7.
	//
	// example:
	//
	// 00000001
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// Specifies whether to save an evidence video.
	//
	// - N: not required (default).
	//
	// - Y: a face verification video (1 to 2 seconds) is captured during the authentication process and returned through the query API.
	//
	// > Because video files are large, the system discards the video file when the network is unstable to prioritize the transmission of images required for authentication.
	//
	// example:
	//
	// N
	DocVideo *string `json:"DocVideo,omitempty" xml:"DocVideo,omitempty"`
	// The document number.
	//
	// Required when **MRTDInput = 2**.
	//
	// example:
	//
	// -
	DocumentNumber *string `json:"DocumentNumber,omitempty" xml:"DocumentNumber,omitempty"`
	// Specifies whether the recognition result page is editable during the document OCR recognition step:
	//
	// - **0**: not editable
	//
	// - **1*	- (default): editable.
	//
	// example:
	//
	// 0
	EditOcrResult *string `json:"EditOcrResult,omitempty" xml:"EditOcrResult,omitempty"`
	// The Indonesian email address. This field takes effect only when Authorize is set to T.
	//
	// >
	//
	// > - This field is required only when the Indonesian data source is enabled.
	//
	// example:
	//
	// evxxx@imigxxxxx.go.id
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The experience code.
	//
	// example:
	//
	// 9be7b7d0180041219e5ab03ac6dab5fb
	ExperienceCode *string `json:"ExperienceCode,omitempty" xml:"ExperienceCode,omitempty"`
	// The face libraries for comparison.
	//
	// example:
	//
	// 0e0c34a77f
	FaceGroupCodes *string `json:"FaceGroupCodes,omitempty" xml:"FaceGroupCodes,omitempty"`
	// The Base64-encoded face photo. If you use FacePictureBase64 to pass in the face photo, check the photo size and do not pass in an excessively large photo.
	//
	// example:
	//
	// Base64
	FacePictureBase64 *string `json:"FacePictureBase64,omitempty" xml:"FacePictureBase64,omitempty"`
	// The URL of the face photo. The URL must be a publicly accessible HTTP or HTTPS link.
	//
	// example:
	//
	// ***
	FacePictureUrl *string `json:"FacePictureUrl,omitempty" xml:"FacePictureUrl,omitempty"`
	// The face registration library.
	//
	// example:
	//
	// 0e0c34a77f
	FaceRegisterGroupCode *string `json:"FaceRegisterGroupCode,omitempty" xml:"FaceRegisterGroupCode,omitempty"`
	// The face verification threshold.
	//
	// example:
	//
	// 0.5
	FaceVerifyThreshold *string `json:"FaceVerifyThreshold,omitempty" xml:"FaceVerifyThreshold,omitempty"`
	// The face image quality. This is an input parameter for the IDV product.
	//
	// example:
	//
	// Y
	IdFaceQuality *string `json:"IdFaceQuality,omitempty" xml:"IdFaceQuality,omitempty"`
	// Specifies whether to enable document anti-spoofing detection. This is an input parameter for the IDV product.
	//
	// example:
	//
	// Y
	IdSpoof *string `json:"IdSpoof,omitempty" xml:"IdSpoof,omitempty"`
	// The custom OCR quality detection threshold mode:
	//
	// - **0**: standard mode
	//
	// - **1**: strict mode
	//
	// - **2**: loose mode
	//
	// - **3*	- (default): quality detection disabled.
	//
	// example:
	//
	// 0
	IdThreshold *string `json:"IdThreshold,omitempty" xml:"IdThreshold,omitempty"`
	// The language configuration. This is an input parameter for the IDV product.
	//
	// example:
	//
	// en
	LanguageConfig *string `json:"LanguageConfig,omitempty" xml:"LanguageConfig,omitempty"`
	// The input source for MRTD verification parameters. This parameter is required for decrypting information when reading document chip data via NFC.
	//
	// - **0**: user input
	//
	// - **1**: OCR reading
	//
	// - **2**: API input.
	//
	// example:
	//
	// 0
	MRTDInput *string `json:"MRTDInput,omitempty" xml:"MRTDInput,omitempty"`
	// The merchant-defined unique business ID used for subsequent troubleshooting. The value can contain letters and digits with a maximum length of 32 characters. Make sure the value is unique.
	//
	// example:
	//
	// e0c34a***353888
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// Your custom user ID or another identifier that can identify a specific user, such as a phone number or email address. We strongly recommend that you desensitize this field value in advance, such as by hashing the value.
	//
	// example:
	//
	// 1221****6543
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// The MetaInfo environment parameter. Obtain this value by using the client SDK.
	//
	// example:
	//
	// {\\"bioMetaInfo\\":\\"4.1.0:2916352,0\\",\\"deviceType\\":\\"web\\",\\"ua\\":\\"Mozilla/5.0 (Macintosh
	MetaInfo *string `json:"MetaInfo,omitempty" xml:"MetaInfo,omitempty"`
	// The Indonesian phone number. The format must start with +62 followed by 9 to 11 digits. This field takes effect only when Authorize is set to T.
	//
	// >
	//
	// > - This field is required only when the Indonesian data source is enabled.
	//
	// example:
	//
	// +6281293671234
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The type of liveness detection:
	//
	// - **LIVENESS*	- (default): blink action liveness detection.
	//
	// - **PHOTINUS_LIVENESS**: blink action liveness + colorful liveness dual detection.
	//
	// >
	//
	// > - For supported SDK versions, see [SDK release notes](https://www.alibabacloud.com/help/zh/ekyc/latest/sdk-publishing-record?spm=a2c63.p38356.0.i99).
	//
	// > - Colorful liveness dual detection is not supported on PC.
	//
	// example:
	//
	// PHOTINUS_LIVENESS
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// Specifies whether to enable OCR. This is an input parameter for the IDV product.
	//
	// example:
	//
	// Y
	Ocr *string `json:"Ocr,omitempty" xml:"Ocr,omitempty"`
	// Specifies whether to return additional OCR recognition standardized format fields:
	//
	// 0: no (default)
	//
	// 1: yes.
	//
	// example:
	//
	// 0
	OcrValueStandard *string `json:"OcrValueStandard,omitempty" xml:"OcrValueStandard,omitempty"`
	// The capture page configuration. Use commas (,) to connect multiple pages. Valid values:
	//
	// - **01**: portrait side of the document
	//
	// - **01,02**: portrait side and back side of the document
	//
	// > When this value is set to 01,02, only Chinese ID cards and Vietnamese ID cards are supported.
	//
	// example:
	//
	// 01
	Pages *string `json:"Pages,omitempty" xml:"Pages,omitempty"`
	// Specifies whether to allow a degraded processing method when compatibility issues occur during mobile H5 authentication.
	//
	// - **url (default)**: degradation supported. The page displays the authentication URL, and the user can copy the URL or switch browsers to continue authentication.
	//
	// - **keep**: degradation not supported. The error reason is returned directly and the authentication flow ends.
	//
	//
	// >
	//
	// > - This switch is not supported on PC.
	//
	// > - If the business scenario involves completing authentication within an in-app embedded web page, set this parameter to keep to disallow URL degradation.
	//
	// example:
	//
	// url
	ProcedurePriority *string `json:"ProcedurePriority,omitempty" xml:"ProcedurePriority,omitempty"`
	// The product solution to use.
	//
	// >For more information, see the official documentation: https://www.alibabacloud.com/help/zh/ekyc/latest/product-introduction?spm=a2c63.p38356.0.i1.
	//
	// example:
	//
	// eKYC
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The order of document and face capture:
	//
	// - DOC_FACE (default)
	//
	// - FACE_DOC
	//
	// Note: This parameter is required only when ProductCode is set to KYC_GLOBAL.
	//
	// example:
	//
	// DOC_FACE
	ProductFlow *string `json:"ProductFlow,omitempty" xml:"ProductFlow,omitempty"`
	// The number of duplicate faces returned.
	//
	// example:
	//
	// 1
	ReturnFaces *string `json:"ReturnFaces,omitempty" xml:"ReturnFaces,omitempty"`
	// The client-side callback URL.
	//
	// example:
	//
	// http*****
	ReturnUrl *string `json:"ReturnUrl,omitempty" xml:"ReturnUrl,omitempty"`
	// Specifies whether to save the face picture.
	//
	// example:
	//
	// 0
	SaveFacePicture *string `json:"SaveFacePicture,omitempty" xml:"SaveFacePicture,omitempty"`
	// The scene code. This is an input parameter for the IDV product.
	//
	// example:
	//
	// 123****123
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// The pattern that represents different security levels of the authentication flow. Valid values:
	//
	// 01: normal pattern (default).
	//
	// 02: safe mode, a relatively strict pattern that is active for high-risk scenarios. This is an input parameter for the IDV product.
	//
	// example:
	//
	// 01
	SecurityLevel *string `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	// Specifies whether to display the album upload entry during the document OCR recognition step:
	//
	// - **1**: display (default)
	//
	// - **0**: do not display.
	//
	// example:
	//
	// 1
	ShowAlbumIcon *string `json:"ShowAlbumIcon,omitempty" xml:"ShowAlbumIcon,omitempty"`
	// Specifies whether to display the guide page:
	//
	// - **1**: display (default)
	//
	// - **0**: do not display.
	//
	// example:
	//
	// 1
	ShowGuidePage *string `json:"ShowGuidePage,omitempty" xml:"ShowGuidePage,omitempty"`
	// Specifies whether to display the recognition result page during the document OCR recognition step:
	//
	// - **1**: display (default)
	//
	// - **0**: do not display.
	//
	// example:
	//
	// 1
	ShowOcrResult *string `json:"ShowOcrResult,omitempty" xml:"ShowOcrResult,omitempty"`
	// The custom UI configuration. Convert your custom UI configuration to a JSON string based on the configuration template and pass it in through this parameter. For more information, see [IDV UI style customization](https://www.alibabacloud.com/help/zh/ekyc/latest/idv-kyc-custom-skin?spm=a2c63.p38356.0.i60).
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
	// The Base64-encoded portrait photo.
	//
	// example:
	//
	// base64
	TargetFacePicture *string `json:"TargetFacePicture,omitempty" xml:"TargetFacePicture,omitempty"`
	// The URL of the portrait image. The URL must be a publicly accessible HTTP or HTTPS link.
	//
	// example:
	//
	// https://www.xxxxx.com/1.jpg
	TargetFacePictureUrl *string `json:"TargetFacePictureUrl,omitempty" xml:"TargetFacePictureUrl,omitempty"`
	// The custom action pool configuration for liveness detection.
	//
	// This parameter is required when Model is set to TEMPLATE.
	//
	// Configuration rule: separate multiple action codes with commas (,). Best practices: include at least one frontal face action (such as blink) and no more than 3 actions in total.
	//
	// Action lookup table:
	//
	// - Blink: 01
	//
	// - Open Mouth: 02
	//
	// - Shake Head Left: 03
	//
	// - Shake Head Right: 04
	//
	// - Move Farther: 05
	//
	// - Move Closer: 06
	//
	// - Photinus: 07.
	//
	// example:
	//
	// 01,02,07
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// The number of actions randomly selected from TemplateConfig.
	//
	// This parameter takes effect only when TemplateType is set to Ran.
	//
	// - Validation rules:
	//
	// - The value must be greater than 1. The value must be less than or equal to the total number of actions configured in TemplateConfig. If not specified, the default value equals the total number of actions in TemplateConfig.
	//
	// example:
	//
	// 2
	TemplateRanCount *string `json:"TemplateRanCount,omitempty" xml:"TemplateRanCount,omitempty"`
	// The execution order of liveness detection actions in TemplateConfig.
	//
	// This parameter is required when Model is set to TEMPLATE.
	//
	// - Seq: actions are executed in the order configured in TemplateConfig from left to right.
	//
	// - Ran: actions are executed in random order. When this option is selected, TemplateConfig must contain more than one action.
	//
	// example:
	//
	// Seq
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// Specifies whether to enable NFC verification when **DocType*	- is set to 01000000 (global passport).
	//
	// - **Y**: enabled
	//
	// - **N**: disabled.
	//
	// example:
	//
	// N
	UseNFC *string `json:"UseNFC,omitempty" xml:"UseNFC,omitempty"`
	// The verification type.
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

func (s *InitializeShrinkRequest) GetOcrValueStandard() *string {
	return s.OcrValueStandard
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

func (s *InitializeShrinkRequest) GetTemplateConfig() *string {
	return s.TemplateConfig
}

func (s *InitializeShrinkRequest) GetTemplateRanCount() *string {
	return s.TemplateRanCount
}

func (s *InitializeShrinkRequest) GetTemplateType() *string {
	return s.TemplateType
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

func (s *InitializeShrinkRequest) SetOcrValueStandard(v string) *InitializeShrinkRequest {
	s.OcrValueStandard = &v
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

func (s *InitializeShrinkRequest) SetTemplateConfig(v string) *InitializeShrinkRequest {
	s.TemplateConfig = &v
	return s
}

func (s *InitializeShrinkRequest) SetTemplateRanCount(v string) *InitializeShrinkRequest {
	s.TemplateRanCount = &v
	return s
}

func (s *InitializeShrinkRequest) SetTemplateType(v string) *InitializeShrinkRequest {
	s.TemplateType = &v
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
