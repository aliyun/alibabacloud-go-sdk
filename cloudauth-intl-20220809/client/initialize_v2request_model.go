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
	SetAutoDocPageConfig(v string) *InitializeV2Request
	GetAutoDocPageConfig() *string
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
	SetFaceAttributeCheck(v string) *InitializeV2Request
	GetFaceAttributeCheck() *string
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
	SetUpdateFaceIfUserExists(v string) *InitializeV2Request
	GetUpdateFaceIfUserExists() *string
	SetUseNFC(v string) *InitializeV2Request
	GetUseNFC() *string
	SetVerifyModel(v string) *InitializeV2Request
	GetVerifyModel() *string
}

type InitializeV2Request struct {
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
	DocPageConfig []*string `json:"DocPageConfig,omitempty" xml:"DocPageConfig,omitempty" type:"Repeated"`
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

func (s *InitializeV2Request) GetAutoDocPageConfig() *string {
	return s.AutoDocPageConfig
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

func (s *InitializeV2Request) GetFaceAttributeCheck() *string {
	return s.FaceAttributeCheck
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

func (s *InitializeV2Request) GetUpdateFaceIfUserExists() *string {
	return s.UpdateFaceIfUserExists
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

func (s *InitializeV2Request) SetAutoDocPageConfig(v string) *InitializeV2Request {
	s.AutoDocPageConfig = &v
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

func (s *InitializeV2Request) SetFaceAttributeCheck(v string) *InitializeV2Request {
	s.FaceAttributeCheck = &v
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

func (s *InitializeV2Request) SetUpdateFaceIfUserExists(v string) *InitializeV2Request {
	s.UpdateFaceIfUserExists = &v
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
