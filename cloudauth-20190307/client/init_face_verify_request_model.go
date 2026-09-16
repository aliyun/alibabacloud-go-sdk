// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitFaceVerifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppQualityCheck(v string) *InitFaceVerifyRequest
	GetAppQualityCheck() *string
	SetAuthId(v string) *InitFaceVerifyRequest
	GetAuthId() *string
	SetBirthday(v string) *InitFaceVerifyRequest
	GetBirthday() *string
	SetCallbackToken(v string) *InitFaceVerifyRequest
	GetCallbackToken() *string
	SetCallbackUrl(v string) *InitFaceVerifyRequest
	GetCallbackUrl() *string
	SetCameraSelection(v string) *InitFaceVerifyRequest
	GetCameraSelection() *string
	SetCertName(v string) *InitFaceVerifyRequest
	GetCertName() *string
	SetCertNo(v string) *InitFaceVerifyRequest
	GetCertNo() *string
	SetCertType(v string) *InitFaceVerifyRequest
	GetCertType() *string
	SetCertifyId(v string) *InitFaceVerifyRequest
	GetCertifyId() *string
	SetCertifyUrlStyle(v string) *InitFaceVerifyRequest
	GetCertifyUrlStyle() *string
	SetCertifyUrlType(v string) *InitFaceVerifyRequest
	GetCertifyUrlType() *string
	SetCrop(v string) *InitFaceVerifyRequest
	GetCrop() *string
	SetEnableBeauty(v string) *InitFaceVerifyRequest
	GetEnableBeauty() *string
	SetEncryptType(v string) *InitFaceVerifyRequest
	GetEncryptType() *string
	SetFaceContrastPicture(v string) *InitFaceVerifyRequest
	GetFaceContrastPicture() *string
	SetFaceContrastPictureUrl(v string) *InitFaceVerifyRequest
	GetFaceContrastPictureUrl() *string
	SetFaceGuardOutput(v string) *InitFaceVerifyRequest
	GetFaceGuardOutput() *string
	SetH5DegradeConfirmBtn(v string) *InitFaceVerifyRequest
	GetH5DegradeConfirmBtn() *string
	SetIp(v string) *InitFaceVerifyRequest
	GetIp() *string
	SetMetaInfo(v string) *InitFaceVerifyRequest
	GetMetaInfo() *string
	SetMobile(v string) *InitFaceVerifyRequest
	GetMobile() *string
	SetMode(v string) *InitFaceVerifyRequest
	GetMode() *string
	SetModel(v string) *InitFaceVerifyRequest
	GetModel() *string
	SetNeedMultiFaceCheck(v string) *InitFaceVerifyRequest
	GetNeedMultiFaceCheck() *string
	SetOssBucketName(v string) *InitFaceVerifyRequest
	GetOssBucketName() *string
	SetOssObjectName(v string) *InitFaceVerifyRequest
	GetOssObjectName() *string
	SetOuterOrderNo(v string) *InitFaceVerifyRequest
	GetOuterOrderNo() *string
	SetProcedurePriority(v string) *InitFaceVerifyRequest
	GetProcedurePriority() *string
	SetProductCode(v string) *InitFaceVerifyRequest
	GetProductCode() *string
	SetRarelyCharacters(v string) *InitFaceVerifyRequest
	GetRarelyCharacters() *string
	SetReadImg(v string) *InitFaceVerifyRequest
	GetReadImg() *string
	SetReturnUrl(v string) *InitFaceVerifyRequest
	GetReturnUrl() *string
	SetSceneId(v int64) *InitFaceVerifyRequest
	GetSceneId() *int64
	SetSuitableType(v string) *InitFaceVerifyRequest
	GetSuitableType() *string
	SetUiCustomUrl(v string) *InitFaceVerifyRequest
	GetUiCustomUrl() *string
	SetUserId(v string) *InitFaceVerifyRequest
	GetUserId() *string
	SetValidityDate(v string) *InitFaceVerifyRequest
	GetValidityDate() *string
	SetVideoEvidence(v string) *InitFaceVerifyRequest
	GetVideoEvidence() *string
	SetVoluntaryCustomizedContent(v string) *InitFaceVerifyRequest
	GetVoluntaryCustomizedContent() *string
}

type InitFaceVerifyRequest struct {
	// Specifies whether the SDK enables strict face quality detection:
	//
	// - **Y**: Enabled.
	//
	// - **N**: Disabled (default).
	//
	//
	// >
	//
	// > - If this parameter is enabled, the SDK must integrate the [strict face quality detection module](https://www.alibabacloud.com/help/en/id-verification/financial-grade-id-verification/description-of-sdk-package-clipping). Strict quality detection may reduce the face recognition success rate.
	//
	// > - Only Android SDK 2.3.24 and later versions are supported.
	//
	// example:
	//
	// N
	AppQualityCheck *string `json:"AppQualityCheck,omitempty" xml:"AppQualityCheck,omitempty"`
	// The user authorization ID. Maximum length: 64 characters.
	//
	// example:
	//
	// 92d46b9e9e2d703f2897f350d5bd4149
	AuthId *string `json:"AuthId,omitempty" xml:"AuthId,omitempty"`
	// The date of birth on the certificate.
	//
	// This field is required when the certificate type **CertType*	- is set to **PASSPORT*	- and **Mode*	- is set to **3**.
	//
	// example:
	//
	// 1993-10-10
	Birthday *string `json:"Birthday,omitempty" xml:"Birthday,omitempty"`
	// The security token that you generate to prevent duplication and tampering.
	//
	// If this value is set, the **CallbackToken*	- field is displayed in the callback URL.
	//
	// example:
	//
	// NMjvQanQgplBSaEI0sL86WnQplB
	CallbackToken *string `json:"CallbackToken,omitempty" xml:"CallbackToken,omitempty"`
	// The callback notification URL for the authentication result. The default callback request method is GET, and the callback URL must start with `https`. After authentication is complete, the platform calls back this URL and automatically appends the `certifyId` and `passed` fields. The `passed` field returns the subcode value. Example: `https://www.aliyun.com?callbackToken=1000004826&certifyId=shaxxxx&passed=200.`
	//
	// <notice>
	//
	// - The callback is triggered only when authentication is complete (including both passed and failed). If authentication is abandoned, abnormally breaks, or is not performed, no notification is sent. After receiving the callback notification, you can use the query operation to obtain authentication details if needed.
	//
	// - The accessibility of the provided URL is verified before the operation is invoked. If the URL cannot be accessed through public network access, error 401 is returned.
	//
	// - After receiving the callback, your service must return HTTP status code 200. Otherwise, a retry is triggered with two callbacks within 3 seconds.
	//
	// </notice>
	//
	// example:
	//
	// https://www.aliyun.com
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// Specifies whether to enable the camera selection feature:
	//
	// - **Y**: Enabled.
	//
	// - **N**: Disabled (default).
	//
	// > This feature takes effect only for PC integration mode. After it is enabled, users can select a camera for authentication.
	//
	// example:
	//
	// N
	CameraSelection *string `json:"CameraSelection,omitempty" xml:"CameraSelection,omitempty"`
	// The real name.
	//
	// example:
	//
	// Wang Shanshan
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The certificate number.
	//
	// example:
	//
	// 330103xxxxxxxxxxxx
	CertNo *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	// The certificate type.
	//
	// Currently, only ID cards are supported. Set this parameter to IDENTITY_CARD.
	//
	// example:
	//
	// IDENTITY_CARD
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// 	Warning: This parameter will be deprecated.</warning>
	//
	// The CertifyId from a previous successful ID Verification session. The photo from that authentication is used as the comparison photo.
	//
	// > Among the four image input methods (FaceContrastPicture, FaceContrastPictureUrl, CertifyId, and OSS), select only one.
	//
	// example:
	//
	// 0bfa7c493f850e5178b9f8613634c9xx
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	// The type of the returned **CertifyUrl**. Valid values:
	//
	// - **L**: Original long URL.
	//
	// - **S*	- (default): Short URL.
	//
	// example:
	//
	// L
	CertifyUrlStyle *string `json:"CertifyUrlStyle,omitempty" xml:"CertifyUrlStyle,omitempty"`
	// The Web SDK device type. Valid values: **WEB*	- or **H5**.
	//
	// > Only Web SDK device types are supported.
	//
	// example:
	//
	// WEB
	CertifyUrlType *string `json:"CertifyUrlType,omitempty" xml:"CertifyUrlType,omitempty"`
	// Specifies whether to allow cropping of face images. Cropping is not allowed by default.
	//
	// - T: Cropping is allowed.
	//
	// - F: Cropping is not allowed.
	//
	// > If the requested image is not captured by a standard liveness detection SDK, allow cropping of face images. After this feature is enabled, the requested image is first cropped and corrected before the request is sent to the service.
	//
	// example:
	//
	// T
	Crop *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	// Specifies whether to enable beauty mode: Y/N.
	//
	// example:
	//
	// Y
	EnableBeauty *string `json:"EnableBeauty,omitempty" xml:"EnableBeauty,omitempty"`
	// The encryption algorithm. Currently, only the SM2 national cryptographic algorithm is supported.
	//
	// After encrypted transmission is enabled, pass in the encrypted CertName and CertNo. For encryption instructions, refer to [Parameter encryption description](https://www.alibabacloud.com/help/en/id-verification/financial-grade-id-verification/description-of-parameter-encryption#task-2229332).
	//
	// example:
	//
	// SM2
	EncryptType *string `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	// The Base64-encoded photo.
	//
	// > Among the four image input methods (FaceContrastPicture, FaceContrastPictureUrl, CertifyId, and OSS), select only one.
	//
	// example:
	//
	// /9j/4AAQSkZJRgABAQAASxxxxxxx
	FaceContrastPicture *string `json:"FaceContrastPicture,omitempty" xml:"FaceContrastPicture,omitempty"`
	// The OSS photo URL. Currently, only authorized OSS photo URLs are supported.
	//
	// > Among the four image input methods (FaceContrastPicture, FaceContrastPictureUrl, CertifyId, and OSS), select only one.
	//
	// example:
	//
	// https://cn-shanghai-aliyun-cloudauth-xxxxxx.oss-cn-shanghai.aliyuncs.com/verify/xxxxx/xxxxx.jpeg
	FaceContrastPictureUrl *string `json:"FaceContrastPictureUrl,omitempty" xml:"FaceContrastPictureUrl,omitempty"`
	// The device assistant tag type. Valid values: **DeviceRisk**.
	//
	// >
	//
	// > - Selecting device assistant output incurs additional fees. For more information, refer to [Paid value-added services](https://www.alibabacloud.com/help/en/id-verification/financial-grade-id-verification/face-guard).
	//
	// > - If you do not need device assistant tag output, do not pass this parameter or pass an empty value.
	//
	// example:
	//
	// DeviceRisk
	FaceGuardOutput *string `json:"FaceGuardOutput,omitempty" xml:"FaceGuardOutput,omitempty"`
	// Specifies whether to display the "I have completed authentication" button on the H5 degradation page after authentication is complete:
	//
	// - **Y**: Enabled.
	//
	// - **N*	- (default): Disabled.
	//
	// example:
	//
	// Y
	H5DegradeConfirmBtn *string `json:"H5DegradeConfirmBtn,omitempty" xml:"H5DegradeConfirmBtn,omitempty"`
	// The IP address of the user.
	//
	// example:
	//
	// 114.xxx.xxx.xxx
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The Metainfo environment parameter, which must be obtained through the client SDK.
	//
	// example:
	//
	// {"zimVer":"3.0.0","appVersion": "1","bioMetaInfo": "4.1.0:11501568,0","appName": "com.aliyun.antcloudauth","deviceType": "ios","osVersion": "iOS 10.3.2","apdidToken": "","deviceModel": "iPhone9,1"}
	MetaInfo *string `json:"MetaInfo,omitempty" xml:"MetaInfo,omitempty"`
	// The mobile phone number of the user.
	//
	// example:
	//
	// 130xxxxxxxx
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The method for obtaining passport NFC verification elements:
	//
	// - **1**: User input. The end user manually enters certificate element information using the UI provided by the Alibaba Cloud SDK.
	//
	// - **3**: External parameter input. Certificate element information is passed in externally.
	//
	// > NFC decoding of passport chip encrypted information requires three passport elements: name, date of birth, and certificate expiration date.
	//
	// example:
	//
	// 1
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The liveness detection type. Valid values:
	//
	// > The liveness detection type supports only the following values. Custom actions or combinations are not supported.
	//
	// Note:
	//
	// The liveness detection type supports only the following values. Custom actions or combinations are not supported.
	//
	// - **LIVENESS*	- (default): Blink.
	//
	// - **PHOTINUS_LIVENESS**: Blink + colorful light.
	//
	// - **MULTI_ACTION**: Blink + head shake (the order of blink and head shake is random).
	//
	// - **MOVE_ACTION*	- (recommended): Move closer/farther + blink.
	//
	// - **MOVE_PHOTINUS**: Move closer/farther + colorful light.
	//
	// >
	//
	// >- **The default liveness detection type*	- is supported in the following versions:
	//
	// >   - Android SDK 1.2.6 and later
	//
	// >   - iOS SDK 1.2.4 and later
	//
	// >   - Harmony SDK 1.0.0 and later
	//
	// >- Other types are supported in the latest Android/iOS/Harmony SDK versions. Integrate the latest version.
	//
	// example:
	//
	// MOVE_ACTION
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// Specifies whether to block authentication when multiple faces are detected on the device. Valid values:
	//
	// - **Y**: Block. The client prompts the user to redo face recognition.
	//
	// - **N*	- (default): Do not block. The largest face in the frame is sent to the server for security detection.
	//
	// example:
	//
	// Y
	NeedMultiFaceCheck *string `json:"NeedMultiFaceCheck,omitempty" xml:"NeedMultiFaceCheck,omitempty"`
	// The bucket name of the authorized OSS space.
	//
	// > Among the four image input methods (FaceContrastPicture, FaceContrastPictureUrl, CertifyId, and OSS), select only one.
	//
	// example:
	//
	// cn-shanghai-aliyun-cloudauth-xxxxx
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// The file name in the authorized OSS space.
	//
	// > Among the four image input methods (FaceContrastPicture, FaceContrastPictureUrl, CertifyId, and OSS), select only one.
	//
	// example:
	//
	// verify/xxxxx/xxxxxx.jpeg
	OssObjectName *string `json:"OssObjectName,omitempty" xml:"OssObjectName,omitempty"`
	// The unique identifier of the merchant request.
	//
	// The value is a 32-character alphanumeric string. The first few characters are a custom abbreviation defined by the merchant, the middle part can use a time segment, and the last part can use a random or incremental sequence.
	//
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c353888
	OuterOrderNo *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	// The degradation configuration when WebRTC or WebAssembly incompatibility occurs during mobile H5 authentication.
	//
	// - **keep**: Degradation is not supported. The system returns directly.
	//
	// - **url*	- (default): Degradation is supported. An authentication URL is returned. The user opens or switches to a browser to authenticate using this URL.
	//
	// - **video**: Degradation is supported. The system camera is used to record a 3-5 second blink video for authentication.
	//
	//
	// >
	//
	// > When the degradation mode is Video, the following features become ineffective and product security is reduced. Configure this mode only for security scenarios.
	//
	// > - The liveness detection type setting does not take effect.
	//
	// > - The VideoEvidence feature is not supported.
	//
	// example:
	//
	// url
	ProcedurePriority *string `json:"ProcedurePriority,omitempty" xml:"ProcedurePriority,omitempty"`
	// A fixed value. This parameter varies depending on the product plan:
	//
	// - APP authentication plan: The fixed value is ID_PRO.
	//
	// - Face liveness verification plan: The fixed value is PV_FV.
	//
	// - Liveness detection plan: The fixed value is LR_FR.
	//
	// example:
	//
	// ID_PRO
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// Specifies whether to enable the rare character mode:
	//
	// - **Y**: Enabled. An information input box pops up before authentication, requiring the user to enter the rare character name and ID card number and agree to the agreement before starting the authentication process.
	//
	// - **N**: Disabled (default).
	//
	// example:
	//
	// Y
	RarelyCharacters *string `json:"RarelyCharacters,omitempty" xml:"RarelyCharacters,omitempty"`
	// Specifies whether to read the certificate photo:
	//
	// - **Y**: Read.
	//
	// - **N**: Do not read.
	//
	// > If the certificate face photo is needed in subsequent authentication steps, set this parameter to Y.
	//
	// example:
	//
	// Y
	ReadImg *string `json:"ReadImg,omitempty" xml:"ReadImg,omitempty"`
	// The target URL to which the merchant business page redirects.
	//
	// example:
	//
	// www.aliyun.com
	ReturnUrl *string `json:"ReturnUrl,omitempty" xml:"ReturnUrl,omitempty"`
	// The authentication scene ID.
	//
	// example:
	//
	// 1000000006
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The elderly-friendly configuration parameter. This parameter takes effect for each authentication request. You can select different parameters for each authentication request based on the business attributes, customer distribution, and operational characteristics of your app. Valid values:
	//
	// - **0*	- (default): Disabled. The current authentication request does not enable elderly-friendly mode.
	//
	// - **1**: Enabled. The current authentication request enables elderly-friendly mode.
	//
	// - **2**: User choice.
	//
	//
	// Allows end users to select the authentication mode. The product guide page provides two authentication entries: "Start Authentication" and "Senior Authentication Mode". When the user selects "Senior Authentication Mode", the system enters elderly-friendly mode.
	//
	// >
	//
	// > - The elderly-friendly parameter takes effect only when the liveness detection type **Model*	- is set to **LIVENESS*	- or **MULTI_ACTION**.
	//
	// example:
	//
	// 0
	SuitableType *string `json:"SuitableType,omitempty" xml:"SuitableType,omitempty"`
	// The UI configuration file URL.
	//
	// You can view the complete configuration in [Web SDK UI custom configuration](https://www.alibabacloud.com/help/en/id-verification/financial-grade-id-verification/web-sdk-ui-custom-configuration-description).
	//
	// example:
	//
	// www.aliyundoc.com
	UiCustomUrl *string `json:"UiCustomUrl,omitempty" xml:"UiCustomUrl,omitempty"`
	// The custom user ID defined by the business. Keep this value unique.
	//
	// example:
	//
	// 123456789
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The certificate expiration date.
	//
	// This field is required when the certificate type **CertType*	- is set to **PASSPORT*	- and **Mode*	- is set to **3**.
	//
	// example:
	//
	// 2039-06-10
	ValidityDate *string `json:"ValidityDate,omitempty" xml:"ValidityDate,omitempty"`
	// Specifies whether to enable video evidence:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled (default).
	//
	// > Because video files are large, the system discards video files to prioritize the transmission of essential authentication images when the network is unstable. Set video as a weak dependency in your business logic.
	//
	// example:
	//
	// false
	VideoEvidence *string `json:"VideoEvidence,omitempty" xml:"VideoEvidence,omitempty"`
	// The custom voluntary content. This parameter is required when personalized settings are enabled. The format is a JSON string of a String List.
	//
	// - For read-aloud scenarios: The content cannot exceed 60 Chinese characters (excluding punctuation), and the List contains only 1 element.
	//
	// - For Q&A scenarios: A maximum of 3 questions can be set. Each question cannot exceed 30 Chinese characters. Each question is a separate element in the List.
	//
	// example:
	//
	// ["I, Mr. Wang, agree to the **	- agreement."]
	VoluntaryCustomizedContent *string `json:"VoluntaryCustomizedContent,omitempty" xml:"VoluntaryCustomizedContent,omitempty"`
}

func (s InitFaceVerifyRequest) String() string {
	return dara.Prettify(s)
}

func (s InitFaceVerifyRequest) GoString() string {
	return s.String()
}

func (s *InitFaceVerifyRequest) GetAppQualityCheck() *string {
	return s.AppQualityCheck
}

func (s *InitFaceVerifyRequest) GetAuthId() *string {
	return s.AuthId
}

func (s *InitFaceVerifyRequest) GetBirthday() *string {
	return s.Birthday
}

func (s *InitFaceVerifyRequest) GetCallbackToken() *string {
	return s.CallbackToken
}

func (s *InitFaceVerifyRequest) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *InitFaceVerifyRequest) GetCameraSelection() *string {
	return s.CameraSelection
}

func (s *InitFaceVerifyRequest) GetCertName() *string {
	return s.CertName
}

func (s *InitFaceVerifyRequest) GetCertNo() *string {
	return s.CertNo
}

func (s *InitFaceVerifyRequest) GetCertType() *string {
	return s.CertType
}

func (s *InitFaceVerifyRequest) GetCertifyId() *string {
	return s.CertifyId
}

func (s *InitFaceVerifyRequest) GetCertifyUrlStyle() *string {
	return s.CertifyUrlStyle
}

func (s *InitFaceVerifyRequest) GetCertifyUrlType() *string {
	return s.CertifyUrlType
}

func (s *InitFaceVerifyRequest) GetCrop() *string {
	return s.Crop
}

func (s *InitFaceVerifyRequest) GetEnableBeauty() *string {
	return s.EnableBeauty
}

func (s *InitFaceVerifyRequest) GetEncryptType() *string {
	return s.EncryptType
}

func (s *InitFaceVerifyRequest) GetFaceContrastPicture() *string {
	return s.FaceContrastPicture
}

func (s *InitFaceVerifyRequest) GetFaceContrastPictureUrl() *string {
	return s.FaceContrastPictureUrl
}

func (s *InitFaceVerifyRequest) GetFaceGuardOutput() *string {
	return s.FaceGuardOutput
}

func (s *InitFaceVerifyRequest) GetH5DegradeConfirmBtn() *string {
	return s.H5DegradeConfirmBtn
}

func (s *InitFaceVerifyRequest) GetIp() *string {
	return s.Ip
}

func (s *InitFaceVerifyRequest) GetMetaInfo() *string {
	return s.MetaInfo
}

func (s *InitFaceVerifyRequest) GetMobile() *string {
	return s.Mobile
}

func (s *InitFaceVerifyRequest) GetMode() *string {
	return s.Mode
}

func (s *InitFaceVerifyRequest) GetModel() *string {
	return s.Model
}

func (s *InitFaceVerifyRequest) GetNeedMultiFaceCheck() *string {
	return s.NeedMultiFaceCheck
}

func (s *InitFaceVerifyRequest) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *InitFaceVerifyRequest) GetOssObjectName() *string {
	return s.OssObjectName
}

func (s *InitFaceVerifyRequest) GetOuterOrderNo() *string {
	return s.OuterOrderNo
}

func (s *InitFaceVerifyRequest) GetProcedurePriority() *string {
	return s.ProcedurePriority
}

func (s *InitFaceVerifyRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *InitFaceVerifyRequest) GetRarelyCharacters() *string {
	return s.RarelyCharacters
}

func (s *InitFaceVerifyRequest) GetReadImg() *string {
	return s.ReadImg
}

func (s *InitFaceVerifyRequest) GetReturnUrl() *string {
	return s.ReturnUrl
}

func (s *InitFaceVerifyRequest) GetSceneId() *int64 {
	return s.SceneId
}

func (s *InitFaceVerifyRequest) GetSuitableType() *string {
	return s.SuitableType
}

func (s *InitFaceVerifyRequest) GetUiCustomUrl() *string {
	return s.UiCustomUrl
}

func (s *InitFaceVerifyRequest) GetUserId() *string {
	return s.UserId
}

func (s *InitFaceVerifyRequest) GetValidityDate() *string {
	return s.ValidityDate
}

func (s *InitFaceVerifyRequest) GetVideoEvidence() *string {
	return s.VideoEvidence
}

func (s *InitFaceVerifyRequest) GetVoluntaryCustomizedContent() *string {
	return s.VoluntaryCustomizedContent
}

func (s *InitFaceVerifyRequest) SetAppQualityCheck(v string) *InitFaceVerifyRequest {
	s.AppQualityCheck = &v
	return s
}

func (s *InitFaceVerifyRequest) SetAuthId(v string) *InitFaceVerifyRequest {
	s.AuthId = &v
	return s
}

func (s *InitFaceVerifyRequest) SetBirthday(v string) *InitFaceVerifyRequest {
	s.Birthday = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCallbackToken(v string) *InitFaceVerifyRequest {
	s.CallbackToken = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCallbackUrl(v string) *InitFaceVerifyRequest {
	s.CallbackUrl = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCameraSelection(v string) *InitFaceVerifyRequest {
	s.CameraSelection = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCertName(v string) *InitFaceVerifyRequest {
	s.CertName = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCertNo(v string) *InitFaceVerifyRequest {
	s.CertNo = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCertType(v string) *InitFaceVerifyRequest {
	s.CertType = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCertifyId(v string) *InitFaceVerifyRequest {
	s.CertifyId = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCertifyUrlStyle(v string) *InitFaceVerifyRequest {
	s.CertifyUrlStyle = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCertifyUrlType(v string) *InitFaceVerifyRequest {
	s.CertifyUrlType = &v
	return s
}

func (s *InitFaceVerifyRequest) SetCrop(v string) *InitFaceVerifyRequest {
	s.Crop = &v
	return s
}

func (s *InitFaceVerifyRequest) SetEnableBeauty(v string) *InitFaceVerifyRequest {
	s.EnableBeauty = &v
	return s
}

func (s *InitFaceVerifyRequest) SetEncryptType(v string) *InitFaceVerifyRequest {
	s.EncryptType = &v
	return s
}

func (s *InitFaceVerifyRequest) SetFaceContrastPicture(v string) *InitFaceVerifyRequest {
	s.FaceContrastPicture = &v
	return s
}

func (s *InitFaceVerifyRequest) SetFaceContrastPictureUrl(v string) *InitFaceVerifyRequest {
	s.FaceContrastPictureUrl = &v
	return s
}

func (s *InitFaceVerifyRequest) SetFaceGuardOutput(v string) *InitFaceVerifyRequest {
	s.FaceGuardOutput = &v
	return s
}

func (s *InitFaceVerifyRequest) SetH5DegradeConfirmBtn(v string) *InitFaceVerifyRequest {
	s.H5DegradeConfirmBtn = &v
	return s
}

func (s *InitFaceVerifyRequest) SetIp(v string) *InitFaceVerifyRequest {
	s.Ip = &v
	return s
}

func (s *InitFaceVerifyRequest) SetMetaInfo(v string) *InitFaceVerifyRequest {
	s.MetaInfo = &v
	return s
}

func (s *InitFaceVerifyRequest) SetMobile(v string) *InitFaceVerifyRequest {
	s.Mobile = &v
	return s
}

func (s *InitFaceVerifyRequest) SetMode(v string) *InitFaceVerifyRequest {
	s.Mode = &v
	return s
}

func (s *InitFaceVerifyRequest) SetModel(v string) *InitFaceVerifyRequest {
	s.Model = &v
	return s
}

func (s *InitFaceVerifyRequest) SetNeedMultiFaceCheck(v string) *InitFaceVerifyRequest {
	s.NeedMultiFaceCheck = &v
	return s
}

func (s *InitFaceVerifyRequest) SetOssBucketName(v string) *InitFaceVerifyRequest {
	s.OssBucketName = &v
	return s
}

func (s *InitFaceVerifyRequest) SetOssObjectName(v string) *InitFaceVerifyRequest {
	s.OssObjectName = &v
	return s
}

func (s *InitFaceVerifyRequest) SetOuterOrderNo(v string) *InitFaceVerifyRequest {
	s.OuterOrderNo = &v
	return s
}

func (s *InitFaceVerifyRequest) SetProcedurePriority(v string) *InitFaceVerifyRequest {
	s.ProcedurePriority = &v
	return s
}

func (s *InitFaceVerifyRequest) SetProductCode(v string) *InitFaceVerifyRequest {
	s.ProductCode = &v
	return s
}

func (s *InitFaceVerifyRequest) SetRarelyCharacters(v string) *InitFaceVerifyRequest {
	s.RarelyCharacters = &v
	return s
}

func (s *InitFaceVerifyRequest) SetReadImg(v string) *InitFaceVerifyRequest {
	s.ReadImg = &v
	return s
}

func (s *InitFaceVerifyRequest) SetReturnUrl(v string) *InitFaceVerifyRequest {
	s.ReturnUrl = &v
	return s
}

func (s *InitFaceVerifyRequest) SetSceneId(v int64) *InitFaceVerifyRequest {
	s.SceneId = &v
	return s
}

func (s *InitFaceVerifyRequest) SetSuitableType(v string) *InitFaceVerifyRequest {
	s.SuitableType = &v
	return s
}

func (s *InitFaceVerifyRequest) SetUiCustomUrl(v string) *InitFaceVerifyRequest {
	s.UiCustomUrl = &v
	return s
}

func (s *InitFaceVerifyRequest) SetUserId(v string) *InitFaceVerifyRequest {
	s.UserId = &v
	return s
}

func (s *InitFaceVerifyRequest) SetValidityDate(v string) *InitFaceVerifyRequest {
	s.ValidityDate = &v
	return s
}

func (s *InitFaceVerifyRequest) SetVideoEvidence(v string) *InitFaceVerifyRequest {
	s.VideoEvidence = &v
	return s
}

func (s *InitFaceVerifyRequest) SetVoluntaryCustomizedContent(v string) *InitFaceVerifyRequest {
	s.VoluntaryCustomizedContent = &v
	return s
}

func (s *InitFaceVerifyRequest) Validate() error {
	return dara.Validate(s)
}
