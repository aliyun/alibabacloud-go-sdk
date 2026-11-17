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
	SetAutoDocPageConfig(v string) *InitializeV2AdvanceRequest
	GetAutoDocPageConfig() *string
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
	SetFaceAttributeCheck(v string) *InitializeV2AdvanceRequest
	GetFaceAttributeCheck() *string
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
	SetIsForceFaceCentered(v string) *InitializeV2AdvanceRequest
	GetIsForceFaceCentered() *string
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
	SetNeedMultiFaceCheck(v string) *InitializeV2AdvanceRequest
	GetNeedMultiFaceCheck() *string
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
	SetUpdateFaceIfUserExists(v string) *InitializeV2AdvanceRequest
	GetUpdateFaceIfUserExists() *string
	SetUseNFC(v string) *InitializeV2AdvanceRequest
	GetUseNFC() *string
	SetVerifyModel(v string) *InitializeV2AdvanceRequest
	GetVerifyModel() *string
}

type InitializeV2AdvanceRequest struct {
	// <warning>This feature is not supported by the **Web SDK**. To use this feature, refer to the App SDK integration guide.</warning>
	//
	// Specifies whether to enable strict face quality detection. Valid values:
	//
	// - Y: Enable. This is the default value.
	//
	// - N: Disable.
	//
	// example:
	//
	// N
	AppQualityCheck *string `json:"AppQualityCheck,omitempty" xml:"AppQualityCheck,omitempty"`
	// Specifies whether to enable authoritative identity verification. This feature currently applies only to second-generation ID cards in the Chinese mainland. This parameter is applicable to IDV products.
	//
	// example:
	//
	// T
	Authorize *string `json:"Authorize,omitempty" xml:"Authorize,omitempty"`
	// The automatic document page classification configuration. This parameter takes effect only when DocType is set to a country or region code. You can specify the pages to collect based on the specific document type code. For documents that are not configured, only the front side is collected.
	//
	// example:
	//
	// {"BRA01001":"01","BRA01002":"01,02"}
	AutoDocPageConfig *string `json:"AutoDocPageConfig,omitempty" xml:"AutoDocPageConfig,omitempty"`
	// Specifies whether to enable automatic registration.
	//
	// example:
	//
	// 0
	AutoRegistration *string `json:"AutoRegistration,omitempty" xml:"AutoRegistration,omitempty"`
	// The security token used for anti-replay and anti-tampering verification. If you specify this parameter, the CallbackToken field is displayed in the callback URL.
	//
	// example:
	//
	// 7ca5c68d869344ea8eeb30cdfd544544-6358700
	CallbackToken *string `json:"CallbackToken,omitempty" xml:"CallbackToken,omitempty"`
	// The callback URL for authentication result notifications. The default callback request method is GET, and the callback URL must start with https. After the authentication is completed, the platform calls this URL and automatically appends the transactionId, passed, and subcode fields.
	//
	// example:
	//
	// https://www.aliyun.com?callbackToken=1000004826&transactionId=shaxxxx&passed=Y&subCode=200
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// Specifies whether to enable the adaptive color-changing window border. Valid values:
	//
	// - **Y**: Enable.
	//
	// - **N**: Disable.
	//
	// example:
	//
	// N
	ChameleonFrameEnable *string `json:"ChameleonFrameEnable,omitempty" xml:"ChameleonFrameEnable,omitempty"`
	// Specifies whether to crop the image. This parameter is applicable to IDV products.
	//
	// example:
	//
	// N
	Crop *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	// The date of birth on the document, in the YYMMDD format (6 digits). This parameter is required when **MRTDInput*	- is set to **2**.
	//
	// example:
	//
	// 900101
	DateOfBirth *string `json:"DateOfBirth,omitempty" xml:"DateOfBirth,omitempty"`
	// The expiration date on the document, in the YYMMDD format (6 digits). This parameter is required when **MRTDInput*	- is set to **2**.
	//
	// example:
	//
	// 301231
	DateOfExpiry *string `json:"DateOfExpiry,omitempty" xml:"DateOfExpiry,omitempty"`
	// The real name of the user.
	//
	// example:
	//
	// Zhang**
	DocName *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
	// The document number of the user.
	//
	// example:
	//
	// 410***************
	DocNo *string `json:"DocNo,omitempty" xml:"DocNo,omitempty"`
	// The custom configuration passed by the customer to specify whether to collect more pages.
	DocPageConfig []*string `json:"DocPageConfig,omitempty" xml:"DocPageConfig,omitempty" type:"Repeated"`
	// The photo capture mode for document collection. Valid values:
	//
	// - manual: Manual capture.
	//
	// - auto: Automatic capture. This is the default value.
	//
	// example:
	//
	// manual
	DocScanMode *string `json:"DocScanMode,omitempty" xml:"DocScanMode,omitempty"`
	// The document type.
	//
	// > For the eKYC_PRO and ID_OCR_MAX editions, see the official documentation: https://www.alibabacloud.com/help/zh/ekyc/latest/certificate-code-table?spm=a2c63.p38356.help-menu-445633.d_2_8_2_0.279147abwKAWbr
	//
	// > For the ID_OCR, eKYC, and eYKC_MIN editions, see the document type list in the official documentation: https://www.alibabacloud.com/help/zh/ekyc/latest/gnhekqy05ni51m4c?spm=a2c63.p38356.help-menu-445633.d_2_3_1_0_0_0.6243244777KoZ7
	//
	// example:
	//
	// 00000001
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// Specifies whether to store a video as evidence. Valid values:
	//
	// - N: No. This is the default value.
	//
	// - Y: Yes. During the authentication process, a video of the user\\"s face scanning process (a 1- to 2-second video file) is collected and returned through the query operation.
	//
	// > Because video files are large, the system discards the video file when the network is unstable to prioritize the transmission of images required for authentication.
	//
	// example:
	//
	// N
	DocVideo *string `json:"DocVideo,omitempty" xml:"DocVideo,omitempty"`
	// The document number. This parameter is required when **MRTDInput*	- is set to **2**.
	//
	// example:
	//
	// -
	DocumentNumber *string `json:"DocumentNumber,omitempty" xml:"DocumentNumber,omitempty"`
	// Specifies whether the recognition result page is editable during the document OCR stage. Valid values:
	//
	// - **0**: Not editable.
	//
	// - **1**: Editable. This is the default value.
	//
	// example:
	//
	// 0
	EditOcrResult *string `json:"EditOcrResult,omitempty" xml:"EditOcrResult,omitempty"`
	// The Indonesian email address. This field takes effect only when Authorize is set to T.
	//
	// > Note
	//
	// > This field is required only when the Indonesian data source is enabled.
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
	// Specifies whether to return additional face attribute information. Valid values: N: Do not return. This is the default value. Y: Return. The result is located in Result.ExtFaceInfo.FaceAttributeInfo.
	//
	// example:
	//
	// N
	FaceAttributeCheck *string `json:"FaceAttributeCheck,omitempty" xml:"FaceAttributeCheck,omitempty"`
	// The face databases for comparison.
	//
	// example:
	//
	// 0e0c34a77f
	FaceGroupCodes *string `json:"FaceGroupCodes,omitempty" xml:"FaceGroupCodes,omitempty"`
	// The Base64-encoded string of the photo. If you choose to pass the face photo by using FacePictureBase64, check the photo size and do not pass an excessively large photo.
	//
	// example:
	//
	// Base64
	FacePictureBase64 *string `json:"FacePictureBase64,omitempty" xml:"FacePictureBase64,omitempty"`
	// The file stream of the face photo.
	//
	// example:
	//
	// InputStream
	FacePictureFileObject io.Reader `json:"FacePictureFile,omitempty" xml:"FacePictureFile,omitempty"`
	// The URL of the face photo. The URL must be a publicly accessible HTTP or HTTPS link.
	//
	// example:
	//
	// https://digital-face-prod8.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
	FacePictureUrl *string `json:"FacePictureUrl,omitempty" xml:"FacePictureUrl,omitempty"`
	// The face database for registration.
	//
	// example:
	//
	// 0e0c34a77f
	FaceRegisterGroupCode *string `json:"FaceRegisterGroupCode,omitempty" xml:"FaceRegisterGroupCode,omitempty"`
	// The threshold for face verification.
	//
	// example:
	//
	// 0.5
	FaceVerifyThreshold *string `json:"FaceVerifyThreshold,omitempty" xml:"FaceVerifyThreshold,omitempty"`
	// The quality of the face image. This parameter is applicable to IDV products.
	//
	// example:
	//
	// Y
	IdFaceQuality *string `json:"IdFaceQuality,omitempty" xml:"IdFaceQuality,omitempty"`
	// Specifies whether to enable document anti-spoofing detection. This parameter is applicable to IDV products.
	//
	// example:
	//
	// Y
	IdSpoof *string `json:"IdSpoof,omitempty" xml:"IdSpoof,omitempty"`
	// The custom threshold mode for OCR quality detection. Valid values:
	//
	// - 0: System default.
	//
	// - 1: Strict mode.
	//
	// - 2: Loose mode.
	//
	// - 3: Disable quality detection. This is the default value.
	//
	// example:
	//
	// 0
	IdThreshold *string `json:"IdThreshold,omitempty" xml:"IdThreshold,omitempty"`
	// The forced centering configuration. Specifies whether to intercept the request when an off-center face is detected on the terminal. Valid values:
	//
	// - Y: Intercept. The client prompts the user to scan their face again.
	//
	// - N: Do not intercept. This is the default value. The system extracts the largest face from the face scanning screen and sends it to the server for security detection.
	//
	// example:
	//
	// Y
	IsForceFaceCentered *string `json:"IsForceFaceCentered,omitempty" xml:"IsForceFaceCentered,omitempty"`
	// The language configuration. This parameter is applicable to IDV products.
	//
	// example:
	//
	// en
	LanguageConfig *string `json:"LanguageConfig,omitempty" xml:"LanguageConfig,omitempty"`
	// The input source of the MRTD verification parameter. This parameter is required to decrypt information when NFC reads the document chip information. Valid values:
	//
	// - **0**: User input.
	//
	// - **1**: OCR reading.
	//
	// - **2**: API input.
	//
	// example:
	//
	// 0
	MRTDInput *string `json:"MRTDInput,omitempty" xml:"MRTDInput,omitempty"`
	// The custom unique business identifier on the merchant side, which is used for subsequent troubleshooting. It supports a combination of letters and digits, with a maximum length of 32 characters. Make sure that the value is unique.
	//
	// example:
	//
	// e0c34a***353888
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// Your custom user ID or another identifier that can identify a specific user, such as a mobile phone number or an email address. We strongly recommend that you desensitize the value of this field in advance, for example, by hashing the value.
	//
	// example:
	//
	// 1221****6543
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// The Metainfo environment parameter, which must be obtained by using the client SDK.
	//
	// example:
	//
	// {\\"bioMetaInfo\\":\\"4.1.0:2916352,0\\",\\"deviceType\\":\\"web\\",\\"ua\\":\\"Mozilla/5.0\\"}
	MetaInfo *string `json:"MetaInfo,omitempty" xml:"MetaInfo,omitempty"`
	// The Indonesian phone number. The format must be verified. It must start with +62 followed by 9 to 11 digits. This field takes effect only when Authorize is set to T.
	//
	// >
	//
	// > - This field is required only when the Indonesian data source is enabled.
	//
	// example:
	//
	// +6281293671234
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The type of liveness detection to perform. Valid values:
	//
	// - **LIVENESS**: Liveness detection based on blinking. This is the default value.
	//
	// - **PHOTINUS_LIVENESS**: Dual liveness detection based on blinking and photinus.
	//
	// >
	//
	// > - For supported SDK versions, see [SDK release notes](https://www.alibabacloud.com/help/zh/ekyc/latest/sdk-publishing-record?spm=a2c63.p38356.0.i99).
	//
	// > - Dual liveness detection based on photinus is not supported on PCs.
	//
	// example:
	//
	// PHOTINUS_LIVENESS
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// Specifies whether to intercept the request when multiple faces are detected on the terminal. Valid values:
	//
	// - Y: Intercept. The client prompts the user to scan their face again.
	//
	// - N: Do not intercept. This is the default value. The system extracts the largest face from the face scanning screen and sends it to the server for security detection.
	//
	// example:
	//
	// Y
	NeedMultiFaceCheck *string `json:"NeedMultiFaceCheck,omitempty" xml:"NeedMultiFaceCheck,omitempty"`
	// Specifies whether to enable OCR. This parameter is applicable to IDV products.
	//
	// example:
	//
	// Y
	Ocr *string `json:"Ocr,omitempty" xml:"Ocr,omitempty"`
	// Specifies whether to return additional standardized OCR recognition fields. Valid values:
	//
	// - 0: No. This is the default value.
	//
	// - 1: Yes.
	//
	// example:
	//
	// 0
	OcrValueStandard *string `json:"OcrValueStandard,omitempty" xml:"OcrValueStandard,omitempty"`
	// The configuration of the pages to collect. Separate multiple pages with commas (,). Valid values:
	//
	// - **01**: The portrait side of the document.
	//
	// - **01,02**: The portrait and back sides of the document.
	//
	// > If you set this parameter to 01,02, only Chinese ID cards and Vietnamese ID cards are supported.
	//
	// example:
	//
	// 01
	Pages *string `json:"Pages,omitempty" xml:"Pages,omitempty"`
	// Specifies whether to allow degraded processing when compatibility issues occur during mobile H5 authentication. Valid values:
	//
	// - **url**: Allow degradation. This is the default value. The page displays the authentication URL, and the user can copy the URL to open it or switch browsers to continue the authentication.
	//
	// - **keep**: Do not allow degradation. The system directly returns the error cause and ends the authentication process.
	//
	//
	// >
	//
	// > - This switch is not supported on PCs.
	//
	// > - If the business scenario involves completing authentication in an embedded web page within an app, we recommend that you set this parameter to keep to disable URL degradation.
	//
	// example:
	//
	// url
	ProcedurePriority *string `json:"ProcedurePriority,omitempty" xml:"ProcedurePriority,omitempty"`
	// The product edition to integrate.
	//
	// > Note: For more information, see the official documentation: https://www.alibabacloud.com/help/zh/ekyc/latest/product-introduction?spm=a2c63.p38356.0.i1
	//
	// example:
	//
	// eKYC
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The sequence of document and face verification. Valid values:
	//
	// - DOC_FACE: Document verification followed by face verification. This is the default value.
	//
	// - FACE_DOC: Face verification followed by document verification.
	//
	// > Note: This parameter is required only when ProductCode is set to KYC_GLOBAL.
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
	// Specifies whether to save the face photo.
	//
	// example:
	//
	// 0
	SaveFacePicture *string `json:"SaveFacePicture,omitempty" xml:"SaveFacePicture,omitempty"`
	// The scene code. This parameter is applicable to IDV products.
	//
	// example:
	//
	// 1234567890
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// The pattern that represents different security levels of the authentication flow. Valid values:
	//
	// - 01: Normal pattern. This is the default value.
	//
	// - 02: safe mode. This is a relatively strict pattern that is active in high-risk scenarios. This parameter is applicable to IDV products.
	//
	// example:
	//
	// 01
	SecurityLevel *string `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	// Specifies whether to display the album upload entry during the document OCR stage. Valid values:
	//
	// - **1**: Display. This is the default value.
	//
	// - **0**: Do not display.
	//
	// example:
	//
	// 1
	ShowAlbumIcon *string `json:"ShowAlbumIcon,omitempty" xml:"ShowAlbumIcon,omitempty"`
	// Specifies whether to display the guide page. Valid values:
	//
	// - **1**: Display. This is the default value.
	//
	// - **0**: Do not display.
	//
	// example:
	//
	// 1
	ShowGuidePage *string `json:"ShowGuidePage,omitempty" xml:"ShowGuidePage,omitempty"`
	// Specifies whether to display the recognition result page during the document OCR stage. Valid values:
	//
	// - **1**: Display. This is the default value.
	//
	// - **0**: Do not display.
	//
	// example:
	//
	// 1
	ShowOcrResult *string `json:"ShowOcrResult,omitempty" xml:"ShowOcrResult,omitempty"`
	// The custom UI configuration. Convert your custom UI configuration into a JSON string based on the configuration template, and pass it through this operation. For more information, see [Customize IDV UI styles](https://www.alibabacloud.com/help/zh/ekyc/latest/idv-kyc-custom-skin?spm=a2c63.p38356.0.i60).
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
	// The Base64-encoded string of the portrait photo.
	//
	// example:
	//
	// base64
	TargetFacePicture *string `json:"TargetFacePicture,omitempty" xml:"TargetFacePicture,omitempty"`
	// The file stream of the base face image.
	//
	// example:
	//
	// InputStream
	TargetFacePictureFileObject io.Reader `json:"TargetFacePictureFile,omitempty" xml:"TargetFacePictureFile,omitempty"`
	// The URL of the portrait photo. The URL must be a publicly accessible HTTP or HTTPS link.
	//
	// example:
	//
	// https://www.xxxxx.com/1.jpg
	TargetFacePictureUrl *string `json:"TargetFacePictureUrl,omitempty" xml:"TargetFacePictureUrl,omitempty"`
	// The custom action pool configuration for liveness detection. This parameter is required when Model is set to TEMPLATE.
	//
	// Configuration rules: Separate multiple action codes with commas (,). Best Practices: We recommend that you include at least one frontal face action (such as blinking) and keep the total number of actions to a maximum of 3.
	//
	// Action code lookup table:
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
	// - Photinus: 07
	//
	// example:
	//
	// 01,02,07
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// **[Deprecated]*	- The random mode is no longer recommended. This field is deprecated.
	//
	// example:
	//
	// 2
	TemplateRanCount *string `json:"TemplateRanCount,omitempty" xml:"TemplateRanCount,omitempty"`
	// The execution sequence of liveness detection actions in TemplateConfig. This parameter is required when Model is set to TEMPLATE. Valid values:
	//
	// - Seq: Execute from left to right in the sequence configured in TemplateConfig.
	//
	// example:
	//
	// Seq
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// Specifies whether to overwrite the existing face with the current face if MerchantUserId already exists during automatic registration. Valid values: Y: Overwrite. N: Do not overwrite and return a message indicating that the UserId already exists.
	//
	// example:
	//
	// N
	UpdateFaceIfUserExists *string `json:"UpdateFaceIfUserExists,omitempty" xml:"UpdateFaceIfUserExists,omitempty"`
	// Specifies whether to enable NFC verification when **DocType*	- is set to 01000000 (global passport). Valid values:
	//
	// - **Y**: Enable.
	//
	// - **N**: Disable.
	//
	// example:
	//
	// N
	UseNFC *string `json:"UseNFC,omitempty" xml:"UseNFC,omitempty"`
	// The type of verification.
	//
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

func (s *InitializeV2AdvanceRequest) GetAutoDocPageConfig() *string {
	return s.AutoDocPageConfig
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

func (s *InitializeV2AdvanceRequest) GetFaceAttributeCheck() *string {
	return s.FaceAttributeCheck
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

func (s *InitializeV2AdvanceRequest) GetIsForceFaceCentered() *string {
	return s.IsForceFaceCentered
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

func (s *InitializeV2AdvanceRequest) GetNeedMultiFaceCheck() *string {
	return s.NeedMultiFaceCheck
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

func (s *InitializeV2AdvanceRequest) GetUpdateFaceIfUserExists() *string {
	return s.UpdateFaceIfUserExists
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

func (s *InitializeV2AdvanceRequest) SetAutoDocPageConfig(v string) *InitializeV2AdvanceRequest {
	s.AutoDocPageConfig = &v
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

func (s *InitializeV2AdvanceRequest) SetFaceAttributeCheck(v string) *InitializeV2AdvanceRequest {
	s.FaceAttributeCheck = &v
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

func (s *InitializeV2AdvanceRequest) SetIsForceFaceCentered(v string) *InitializeV2AdvanceRequest {
	s.IsForceFaceCentered = &v
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

func (s *InitializeV2AdvanceRequest) SetNeedMultiFaceCheck(v string) *InitializeV2AdvanceRequest {
	s.NeedMultiFaceCheck = &v
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

func (s *InitializeV2AdvanceRequest) SetUpdateFaceIfUserExists(v string) *InitializeV2AdvanceRequest {
	s.UpdateFaceIfUserExists = &v
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
