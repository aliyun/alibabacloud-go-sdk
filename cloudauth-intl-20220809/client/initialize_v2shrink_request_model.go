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
	SetAutoDocPageConfig(v string) *InitializeV2ShrinkRequest
	GetAutoDocPageConfig() *string
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
	SetFaceAttributeCheck(v string) *InitializeV2ShrinkRequest
	GetFaceAttributeCheck() *string
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
	SetUpdateFaceIfUserExists(v string) *InitializeV2ShrinkRequest
	GetUpdateFaceIfUserExists() *string
	SetUseNFC(v string) *InitializeV2ShrinkRequest
	GetUseNFC() *string
	SetVerifyModel(v string) *InitializeV2ShrinkRequest
	GetVerifyModel() *string
}

type InitializeV2ShrinkRequest struct {
	// <warning>This feature is not supported by the **Web SDK**. To use this feature, refer to the App SDK integration.</warning>
	//
	// Specifies whether to enable strict face quality detection. Valid values:
	//
	// - Y: Enabled. This is the default value.
	//
	// - N: Not enabled.
	//
	// example:
	//
	// N
	AppQualityCheck *string `json:"AppQualityCheck,omitempty" xml:"AppQualityCheck,omitempty"`
	// Specifies whether to enable authoritative identity verification. Currently, this applies only to second-generation ID cards in the Chinese mainland. This is an input parameter for the IDV product.
	//
	// example:
	//
	// T
	Authorize         *string `json:"Authorize,omitempty" xml:"Authorize,omitempty"`
	AutoDocPageConfig *string `json:"AutoDocPageConfig,omitempty" xml:"AutoDocPageConfig,omitempty"`
	// Specifies whether to enable automatic registration.
	//
	// example:
	//
	// 0
	AutoRegistration *string `json:"AutoRegistration,omitempty" xml:"AutoRegistration,omitempty"`
	// The security token used for anti-replay and anti-tampering verification. If this parameter is passed in, the CallbackToken field is displayed in the callback URL.
	//
	// example:
	//
	// 7ca5c68d869344ea8eeb30cdfd544544-6358700
	CallbackToken *string `json:"CallbackToken,omitempty" xml:"CallbackToken,omitempty"`
	// The callback URL for the authentication result notification. The callback request method is GET by default, and the callback URL must start with https. After the authentication is complete, the platform calls back this URL and automatically appends the transactionId, passed, and subcode fields.
	//
	// example:
	//
	// https://www.aliyun.com?callbackToken=1000004826&transactionId=shaxxxx&passed=Y&subCode=200
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// Specifies whether to enable the adaptive color-changing window frame.
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
	// example:
	//
	// -
	DateOfBirth *string `json:"DateOfBirth,omitempty" xml:"DateOfBirth,omitempty"`
	// The expiration date on the document.
	//
	// example:
	//
	// -
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
	// The custom configuration for whether to collect additional pages.
	//
	// example:
	//
	// OCR_ID_BACK
	DocPageConfigShrink *string `json:"DocPageConfig,omitempty" xml:"DocPageConfig,omitempty"`
	// The document capture photo mode.
	//
	// example:
	//
	// manual
	DocScanMode *string `json:"DocScanMode,omitempty" xml:"DocScanMode,omitempty"`
	// The document type.
	//
	// >For the eKYC_PRO and ID_OCR_MAX solutions, see the official documentation at https://www.alibabacloud.com/help/zh/ekyc/latest/certificate-code-table?spm=a2c63.p38356.help-menu-445633.d_2_8_2_0.279147abwKAWbr
	//
	// >For the ID_OCR, eKYC, and eKYC_MIN solutions, see the document type list in the official documentation at https://www.alibabacloud.com/help/zh/ekyc/latest/gnhekqy05ni51m4c?spm=a2c63.p38356.help-menu-445633.d_2_3_1_0_0_0.6243244777KoZ7
	//
	// example:
	//
	// 00000001
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// Specifies whether to collect a verification video.
	//
	// - N: No verification video is collected. This is the default value.
	//
	// - Y: A short video (1 to 2 seconds) of the user\\"s face verification process is collected during authentication and returned through the query operation.
	//
	// > Because video files are large, the system discards video files when the network is unstable to prioritize the transmission of images required for authentication.
	//
	// example:
	//
	// N
	DocVideo *string `json:"DocVideo,omitempty" xml:"DocVideo,omitempty"`
	// The document number.
	//
	// example:
	//
	// -
	DocumentNumber *string `json:"DocumentNumber,omitempty" xml:"DocumentNumber,omitempty"`
	// Specifies whether the recognition result page is editable during the document OCR recognition step:
	//
	// example:
	//
	// 0
	EditOcrResult *string `json:"EditOcrResult,omitempty" xml:"EditOcrResult,omitempty"`
	// The Indonesian email address. This field takes effect only when Authorize is set to T.
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
	ExperienceCode     *string `json:"ExperienceCode,omitempty" xml:"ExperienceCode,omitempty"`
	FaceAttributeCheck *string `json:"FaceAttributeCheck,omitempty" xml:"FaceAttributeCheck,omitempty"`
	// The face libraries to compare against.
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
	// The file stream of the face photo.
	//
	// example:
	//
	// InputStream
	FacePictureFile *string `json:"FacePictureFile,omitempty" xml:"FacePictureFile,omitempty"`
	// The URL of the face photo. The URL must be a publicly accessible HTTP or HTTPS link.
	//
	// example:
	//
	// https://digital-face-prod8.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
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
	// The source of the MRTD verification parameter input. This parameter is required to decrypt information when reading document chip data via NFC.
	//
	// - **0**: user input
	//
	// - **1**: OCR reading
	//
	// - **2**: API input
	//
	// example:
	//
	// 0
	MRTDInput *string `json:"MRTDInput,omitempty" xml:"MRTDInput,omitempty"`
	// The merchant-defined unique business ID used for subsequent troubleshooting. The value supports a combination of letters and numbers with a maximum length of 32 characters. Ensure that the value is unique.
	//
	// example:
	//
	// e0c34a***353888
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// The custom user ID, or another identifier that can identify a specific user, such as a phone number or email address. We strongly recommend that you hash or otherwise de-identify this value before passing it in.
	//
	// example:
	//
	// 1221****6543
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// The Metainfo environment parameter, which must be obtained through the client SDK.
	//
	// example:
	//
	// {\\"bioMetaInfo\\":\\"4.1.0:2916352,0\\",\\"deviceType\\":\\"web\\",\\"ua\\":\\"Mozilla/5.0\\"}
	MetaInfo *string `json:"MetaInfo,omitempty" xml:"MetaInfo,omitempty"`
	// The Indonesian phone number. The format must start with +62 followed by 9 to 11 digits. This field takes effect only when Authorize is set to T.
	//
	// example:
	//
	// +6281293671234
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The type of liveness detection to perform.
	//
	// - **LIVENESS*	- (default): Blink-based liveness detection.
	//
	// - **PHOTINUS_LIVENESS**: Dual detection that combines blink-based liveness detection and colorful light liveness detection.
	//
	// >
	//
	// > - For supported SDK versions, see [SDK release notes](https://www.alibabacloud.com/help/zh/ekyc/latest/sdk-publishing-record?spm=a2c63.p38356.0.i99).
	//
	// > - Colorful light dual detection is not supported on PCs.
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
	// Specifies whether to additionally return OCR recognition results in a standardized format:
	//
	// example:
	//
	// 0
	OcrValueStandard *string `json:"OcrValueStandard,omitempty" xml:"OcrValueStandard,omitempty"`
	// The configuration for capture pages. Separate multiple pages with commas (,). Valid values:
	//
	// - **01**: the portrait side of the identity document.
	//
	// - **01,02**: the portrait side and back side of the identity document.
	//
	// > When this parameter is set to 01,02, only China identity cards and Vietnam identity cards are supported.
	//
	// example:
	//
	// 01
	Pages *string `json:"Pages,omitempty" xml:"Pages,omitempty"`
	// Specifies whether to allow a degraded processing method when compatibility issues occur during mobile H5 authentication.
	//
	// - **url (default)**: Degradation is supported. The page displays the authentication URL, and the user can copy the URL to open it or switch browsers to continue authentication.
	//
	// - **keep**: Degradation is not supported. The error reason is directly returned, and the authentication process ends.
	//
	//
	// >
	//
	// > - This parameter is not supported on PC.
	//
	// > - If the business scenario involves completing authentication within a webpage embedded in an app, set this parameter to keep to disallow URL degradation.
	//
	// example:
	//
	// url
	ProcedurePriority *string `json:"ProcedurePriority,omitempty" xml:"ProcedurePriority,omitempty"`
	// The product plan to use.
	//
	// >**Note*	- For more information, see the official documentation: https://www.alibabacloud.com/help/zh/ekyc/latest/product-introduction?spm=a2c63.p38356.0.i1
	//
	// example:
	//
	// eKYC
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// Specifies the order of document and face verification steps. Valid values:
	//
	// - DOC_FACE: Document first, then face. This is the default value.
	//
	// - FACE_DOC: Face first, then document.
	//
	// >**Note**: This parameter is required only when ProductCode is set to KYC_GLOBAL.
	//
	// example:
	//
	// DOC_FACE
	ProductFlow *string `json:"ProductFlow,omitempty" xml:"ProductFlow,omitempty"`
	// The number of duplicate faces to return.
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
	// The scene code. This is an input parameter for the IDV product.
	//
	// example:
	//
	// 1234567890
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// The mode that represents different security levels of the authentication process. Valid values:
	//
	// - 01: Normal mode (default).
	//
	// - 02: Safe mode. A relatively strict mode that can be used for high-risk scenarios (input parameter for IDV products).
	//
	// example:
	//
	// 01
	SecurityLevel *string `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	// Specifies whether to display the album upload entry during the document OCR recognition step:
	//
	// example:
	//
	// 1
	ShowAlbumIcon *string `json:"ShowAlbumIcon,omitempty" xml:"ShowAlbumIcon,omitempty"`
	// Specifies whether to display the guide page:
	//
	// example:
	//
	// 1
	ShowGuidePage *string `json:"ShowGuidePage,omitempty" xml:"ShowGuidePage,omitempty"`
	// Specifies whether to display the recognition result page during the document OCR recognition step:
	//
	// example:
	//
	// 1
	ShowOcrResult *string `json:"ShowOcrResult,omitempty" xml:"ShowOcrResult,omitempty"`
	// The custom UI configuration. Convert your custom UI configuration to a JSON string based on the configuration template, and pass it in through this parameter. For more information, see [IDV UI style customization](https://www.alibabacloud.com/help/zh/ekyc/latest/idv-kyc-custom-skin?spm=a2c63.p38356.0.i60).
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
	// The Base64-encoded face photo.
	//
	// example:
	//
	// base64
	TargetFacePicture *string `json:"TargetFacePicture,omitempty" xml:"TargetFacePicture,omitempty"`
	// The file stream of the reference face photo.
	//
	// example:
	//
	// InputStream
	TargetFacePictureFile *string `json:"TargetFacePictureFile,omitempty" xml:"TargetFacePictureFile,omitempty"`
	// The URL of the face photo. The URL must be a publicly accessible HTTP or HTTPS link.
	//
	// example:
	//
	// https://www.xxxxx.com/1.jpg
	TargetFacePictureUrl *string `json:"TargetFacePictureUrl,omitempty" xml:"TargetFacePictureUrl,omitempty"`
	// The custom action pool configuration for liveness detection.
	//
	// example:
	//
	// 01,02,07
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// The number of actions to randomly select from TemplateConfig.
	//
	// example:
	//
	// 2
	TemplateRanCount *string `json:"TemplateRanCount,omitempty" xml:"TemplateRanCount,omitempty"`
	// The execution order of liveness detection actions in TemplateConfig.
	//
	// example:
	//
	// Seq
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// Specifies whether to overwrite the existing face with the current face when the MerchantUserId already exists during automatic registration. Valid values: Y (overwrite) and N (do not overwrite and return a message indicating that the UserId already exists).
	UpdateFaceIfUserExists *string `json:"UpdateFaceIfUserExists,omitempty" xml:"UpdateFaceIfUserExists,omitempty"`
	// When **DocType*	- is set to 01000000 (global passport), specifies whether to enable NFC verification.
	//
	// - **Y**: Enable NFC verification.
	//
	// - **N**: Disable NFC verification.
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

func (s *InitializeV2ShrinkRequest) GetAutoDocPageConfig() *string {
	return s.AutoDocPageConfig
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

func (s *InitializeV2ShrinkRequest) GetFaceAttributeCheck() *string {
	return s.FaceAttributeCheck
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

func (s *InitializeV2ShrinkRequest) GetUpdateFaceIfUserExists() *string {
	return s.UpdateFaceIfUserExists
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

func (s *InitializeV2ShrinkRequest) SetAutoDocPageConfig(v string) *InitializeV2ShrinkRequest {
	s.AutoDocPageConfig = &v
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

func (s *InitializeV2ShrinkRequest) SetFaceAttributeCheck(v string) *InitializeV2ShrinkRequest {
	s.FaceAttributeCheck = &v
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

func (s *InitializeV2ShrinkRequest) SetUpdateFaceIfUserExists(v string) *InitializeV2ShrinkRequest {
	s.UpdateFaceIfUserExists = &v
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
