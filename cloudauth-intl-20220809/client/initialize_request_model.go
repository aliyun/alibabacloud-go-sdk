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
	SetAutoDocPageConfig(v string) *InitializeRequest
	GetAutoDocPageConfig() *string
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
	SetFaceAttributeCheck(v string) *InitializeRequest
	GetFaceAttributeCheck() *string
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
	SetOcrValueStandard(v string) *InitializeRequest
	GetOcrValueStandard() *string
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
	SetTemplateConfig(v string) *InitializeRequest
	GetTemplateConfig() *string
	SetTemplateRanCount(v string) *InitializeRequest
	GetTemplateRanCount() *string
	SetTemplateType(v string) *InitializeRequest
	GetTemplateType() *string
	SetUpdateFaceIfUserExists(v string) *InitializeRequest
	GetUpdateFaceIfUserExists() *string
	SetUseNFC(v string) *InitializeRequest
	GetUseNFC() *string
	SetVerifyModel(v string) *InitializeRequest
	GetVerifyModel() *string
}

type InitializeRequest struct {
	// <warning>This feature is not supported by the **Web SDK**. To use this feature, use the App SDK.</warning>
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
	// Specifies whether to enable authoritative identity verification. Currently, this parameter applies only to second-generation ID cards in the Chinese mainland. This is an input parameter for the IDV product.
	//
	// example:
	//
	// Y
	Authorize *string `json:"Authorize,omitempty" xml:"Authorize,omitempty"`
	// The automatic document classification configuration.
	AutoDocPageConfig *string `json:"AutoDocPageConfig,omitempty" xml:"AutoDocPageConfig,omitempty"`
	// Specifies whether to enable automatic registration.
	//
	// example:
	//
	// 0
	AutoRegistration *string `json:"AutoRegistration,omitempty" xml:"AutoRegistration,omitempty"`
	// The security token used for anti-replay and anti-tampering verification. If this parameter is specified, the CallbackToken field is displayed in the callback URL.
	//
	// example:
	//
	// 7ca5c68d869344ea8eeb30cdfd544544-6358700
	CallbackToken *string `json:"CallbackToken,omitempty" xml:"CallbackToken,omitempty"`
	// The callback URL for the authentication result. The callback request method is GET by default, and the callback URL must start with https. After the authentication is complete, the platform calls back this URL and automatically adds the transactionId, passed, and subcode fields.
	//
	// example:
	//
	// https://www.aliyun.com?callbackToken=1000004826&transactionId=shaxxxx&passed=Y&subCode=200
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// Specifies whether to enable the adaptive color-changing window border.
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
	// John Smith
	DocName *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
	// The document number of the user.
	//
	// example:
	//
	// 411xxxxxxxxxxx0001
	DocNo *string `json:"DocNo,omitempty" xml:"DocNo,omitempty"`
	// The custom configuration for whether to collect additional pages.
	DocPageConfig []*string `json:"DocPageConfig,omitempty" xml:"DocPageConfig,omitempty" type:"Repeated"`
	// The document capture and photo mode.
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
	// - N: No (default).
	//
	// - Y: A short video (1 to 2 seconds) of the user\\"s face verification process is collected and returned through the query operation.
	//
	// > Because video files are large, the system discards video files when the network is unstable to prioritize the transmission of images required for verification.
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
	ExperienceCode *string `json:"ExperienceCode,omitempty" xml:"ExperienceCode,omitempty"`
	// Specifies whether to enable face attribute check.
	FaceAttributeCheck *string `json:"FaceAttributeCheck,omitempty" xml:"FaceAttributeCheck,omitempty"`
	// The face libraries to compare against.
	//
	// example:
	//
	// 0e0c34a77f
	FaceGroupCodes *string `json:"FaceGroupCodes,omitempty" xml:"FaceGroupCodes,omitempty"`
	// The Base64-encoded face photo. If you use FacePictureBase64 to pass in a face photo, check the photo size and do not pass in an excessively large photo.
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
	// Specifies whether to enable document anti-forgery detection. This is an input parameter for the IDV product.
	//
	// example:
	//
	// Y
	IdSpoof *string `json:"IdSpoof,omitempty" xml:"IdSpoof,omitempty"`
	// The custom OCR quality detection threshold mode. Valid values:
	//
	// - **0**: Standard mode.
	//
	// - **1**: Strict mode.
	//
	// - **2**: Loose mode.
	//
	// - **3*	- (default): Quality detection disabled.
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
	// The input source of MRTD verification parameters. This parameter is required to decrypt information when reading document chip data via NFC.
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
	// The merchant-defined unique business ID for subsequent troubleshooting. The value can contain letters and digits with a maximum length of 32 characters. Ensure that the value is unique.
	//
	// example:
	//
	// e0c34a***353888
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// Your custom user ID, or another identifier that can identify a specific user, such as a phone number or email address. We strongly recommend that you desensitize this field value in advance, for example, by hashing the value.
	//
	// example:
	//
	// 1221****6543
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// The Metainfo environment parameter, which must be obtained from the client SDK.
	//
	// example:
	//
	// {\\"bioMetaInfo\\":\\"4.1.0:2916352,0\\",\\"deviceType\\":\\"web\\",\\"ua\\":\\"Mozilla/5.0 (Macintosh
	MetaInfo *string `json:"MetaInfo,omitempty" xml:"MetaInfo,omitempty"`
	// The Indonesian phone number. The format must be verified (starting with +62, followed by 9 to 11 digits). This field takes effect only when Authorize is set to T.
	//
	// example:
	//
	// +6281293671234
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The type of liveness detection to perform.
	//
	// - **LIVENESS*	- (default): Blink-based liveness detection.
	//
	// - **PHOTINUS_LIVENESS**: Dual detection combining blink-based liveness detection and flash-based liveness detection.
	//
	// >
	//
	// > - For supported SDK versions, see [SDK release notes](https://www.alibabacloud.com/help/zh/ekyc/latest/sdk-publishing-record?spm=a2c63.p38356.0.i99).
	//
	// > - Flash-based dual liveness detection is not supported on PCs.
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
	// Specifies whether to additionally return OCR recognition results in standardized format fields:
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
	// example:
	//
	// url
	ProcedurePriority *string `json:"ProcedurePriority,omitempty" xml:"ProcedurePriority,omitempty"`
	// The product plan to use.
	//
	// >For more information, see the official documentation: https://www.alibabacloud.com/help/zh/ekyc/latest/product-introduction?spm=a2c63.p38356.0.i1
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
	// >**Note:*	- This parameter is required only when ProductCode is set to KYC_GLOBAL.
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
	// The mode that represents different security levels of the authentication process. Valid values:
	//
	// example:
	//
	// 01
	SecurityLevel *string `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	// Specifies whether to display the album upload entry during the document OCR recognition step. Valid values:
	//
	// - **1**: Display. This is the default value.
	//
	// - **0**: Do not display.
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
	// Configuration rule: Separate multiple action codes with commas. Best practice: Include at least one frontal face action (such as blink), and use no more than 3 actions in total.
	//
	// Action code table:
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
	// Specifies whether to overwrite the existing face with the current face when MerchantUserId already exists during automatic registration. Y: overwrite. N: do not overwrite and return a message indicating that the UserId already exists.
	UpdateFaceIfUserExists *string `json:"UpdateFaceIfUserExists,omitempty" xml:"UpdateFaceIfUserExists,omitempty"`
	// Specifies whether to enable NFC verification when **DocType*	- is set to 01000000 (global passport). Valid values:
	//
	// - **Y**: Enabled.
	//
	// - **N**: Not enabled.
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

func (s *InitializeRequest) GetAutoDocPageConfig() *string {
	return s.AutoDocPageConfig
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

func (s *InitializeRequest) GetFaceAttributeCheck() *string {
	return s.FaceAttributeCheck
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

func (s *InitializeRequest) GetOcrValueStandard() *string {
	return s.OcrValueStandard
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

func (s *InitializeRequest) GetTemplateConfig() *string {
	return s.TemplateConfig
}

func (s *InitializeRequest) GetTemplateRanCount() *string {
	return s.TemplateRanCount
}

func (s *InitializeRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *InitializeRequest) GetUpdateFaceIfUserExists() *string {
	return s.UpdateFaceIfUserExists
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

func (s *InitializeRequest) SetAutoDocPageConfig(v string) *InitializeRequest {
	s.AutoDocPageConfig = &v
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

func (s *InitializeRequest) SetFaceAttributeCheck(v string) *InitializeRequest {
	s.FaceAttributeCheck = &v
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

func (s *InitializeRequest) SetOcrValueStandard(v string) *InitializeRequest {
	s.OcrValueStandard = &v
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

func (s *InitializeRequest) SetTemplateConfig(v string) *InitializeRequest {
	s.TemplateConfig = &v
	return s
}

func (s *InitializeRequest) SetTemplateRanCount(v string) *InitializeRequest {
	s.TemplateRanCount = &v
	return s
}

func (s *InitializeRequest) SetTemplateType(v string) *InitializeRequest {
	s.TemplateType = &v
	return s
}

func (s *InitializeRequest) SetUpdateFaceIfUserExists(v string) *InitializeRequest {
	s.UpdateFaceIfUserExists = &v
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
