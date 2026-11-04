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
	SetEnableAlbumUpload(v string) *InitFaceVerifyRequest
	GetEnableAlbumUpload() *string
	SetEnableBeauty(v string) *InitFaceVerifyRequest
	GetEnableBeauty() *string
	SetEnableTts(v string) *InitFaceVerifyRequest
	GetEnableTts() *string
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
	// Specifies whether to enable strict face quality detection in the SDK. Valid values:
	//
	// - **Y**: Enables strict face quality detection.
	//
	// - **N**: Does not enable strict face quality detection. This is the default value.
	//
	//
	// >
	//
	// > - If you enable this parameter, the SDK must integrate the [strict face quality detection module](https://www.alibabacloud.com/help/en/id-verification/financial-grade-id-verification/description-of-sdk-package-clipping). Affected by strict quality detection, the success rate of face scanning may decrease.
	//
	// > - Only Android SDK 2.3.24 and later are supported.
	//
	// example:
	//
	// N
	AppQualityCheck *string `json:"AppQualityCheck,omitempty" xml:"AppQualityCheck,omitempty"`
	// The user authorization ID. The maximum length is 64 characters.
	//
	// example:
	//
	// 92d46b9e9e2d703f2897f350d5bd4149
	AuthId *string `json:"AuthId,omitempty" xml:"AuthId,omitempty"`
	// The date of birth on the ID document.
	//
	// This parameter is required when the **CertType*	- parameter is set to **PASSPORT*	- and the **Mode*	- parameter is set to **3**.
	//
	// example:
	//
	// 1993-10-10
	Birthday *string `json:"Birthday,omitempty" xml:"Birthday,omitempty"`
	// The security token that you generate to prevent duplicate and tampered requests.
	//
	// If you specify this parameter, the **CallbackToken*	- field is displayed in the callback URL.
	//
	// example:
	//
	// NMjvQanQgplBSaEI0sL86WnQplB
	CallbackToken *string `json:"CallbackToken,omitempty" xml:"CallbackToken,omitempty"`
	// The callback URL for authentication result notifications. The default callback request method is GET. The callback URL must start with `https`. After the authentication is completed, the platform invokes this URL and automatically appends the `certifyId` and `passed` fields. The value of the `passed` field is the subcode. Example: `https://www.aliyun.com?callbackToken=1000004826&certifyId=shaxxxx&passed=200`.
	//
	// <notice>
	//
	// - The callback is triggered only when the authentication is completed, including both successful and failed authentications. If the authentication is abandoned, an abnormal break occurs, or the authentication is not performed, no notification is sent. We recommend that you invoke the query API operation to obtain the authentication details after you receive the callback notification if needed.
	//
	// - The accessibility of the URL is verified before the API operation is invoked. If the URL does not support public network access, a 401 error is returned.
	//
	// - The API operation must return HTTP status code 200 after the callback. Otherwise, a retry is triggered, and two callbacks are sent within 3 seconds.
	//
	// </notice>
	//
	// example:
	//
	// https://www.aliyun.com
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// Specifies whether to enable the camera selection feature. Valid values:
	//
	// - **Y**: Enables the feature.
	//
	// - **N**: Does not enable the feature. This is the default value.
	//
	// > This feature takes effect only in the PC integration mode. After this feature is enabled, users can select a camera for authentication.
	//
	// example:
	//
	// N
	CameraSelection *string `json:"CameraSelection,omitempty" xml:"CameraSelection,omitempty"`
	// The real name.
	//
	// example:
	//
	// Zhang San
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The ID document number.
	//
	// example:
	//
	// 330103xxxxxxxxxxxx
	CertNo *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	// The type of the ID document.
	//
	// Only ID cards are supported. Set the value to IDENTITY_CARD.
	//
	// example:
	//
	// IDENTITY_CARD
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// **[Deprecated]**
	//
	// The CertifyId of a previous successful ID Verification. The photo used in the previous authentication is used as the comparison photo.
	//
	// > You can specify only one of the following parameters to pass the image: FaceContrastPicture, FaceContrastPictureUrl, CertifyId, and OSS.
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
	// The device type of the Web SDK. Valid values: **WEB*	- and **H5**.
	//
	// > Only Web SDK device types are supported.
	//
	// example:
	//
	// WEB
	CertifyUrlType *string `json:"CertifyUrlType,omitempty" xml:"CertifyUrlType,omitempty"`
	// Specifies whether to allow face image cropping. By default, cropping is not allowed. Valid values:
	//
	// - T: Cropping is allowed.
	//
	// - F: Cropping is not allowed.
	//
	// > If the requested image is not collected by the standard liveness detection SDK, we recommend that you allow face image cropping. After this feature is enabled, the requested image is cropped and corrected before the request is sent to the service.
	//
	// example:
	//
	// T
	Crop *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	// Specifies whether to allow album upload. Valid values: Y and N. Default value: N.
	//
	// example:
	//
	// Y
	EnableAlbumUpload *string `json:"EnableAlbumUpload,omitempty" xml:"EnableAlbumUpload,omitempty"`
	// Specifies whether to enable the beauty filter. Valid values: Y and N.
	//
	// example:
	//
	// Y
	EnableBeauty *string `json:"EnableBeauty,omitempty" xml:"EnableBeauty,omitempty"`
	// Specifies whether to enable speech synthesis for H5 willingness authentication. Valid values: Y and N. Default value: Y.
	//
	// example:
	//
	// Y
	EnableTts *string `json:"EnableTts,omitempty" xml:"EnableTts,omitempty"`
	// The encryption algorithm. Only the SM2 algorithm is supported.
	//
	// After encrypted transmission is enabled, you must pass the encrypted CertName and CertNo. For more information about encryption, see [Parameter encryption](https://www.alibabacloud.com/help/en/id-verification/financial-grade-id-verification/description-of-parameter-encryption#task-2229332).
	//
	// example:
	//
	// SM2
	EncryptType *string `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	// The Base64-encoded image.
	//
	// > You can specify only one of the following parameters to pass the image: FaceContrastPicture, FaceContrastPictureUrl, CertifyId, and OSS.
	//
	// example:
	//
	// /9j/4AAQSkZJRgABAQAASxxxxxxx
	FaceContrastPicture *string `json:"FaceContrastPicture,omitempty" xml:"FaceContrastPicture,omitempty"`
	// The URL of the image in OSS. Only authorized OSS image URLs are supported.
	//
	// > You can specify only one of the following parameters to pass the image: FaceContrastPicture, FaceContrastPictureUrl, CertifyId, and OSS.
	//
	// example:
	//
	// https://cn-shanghai-aliyun-cloudauth-xxxxxx.oss-cn-shanghai.aliyuncs.com/verify/xxxxx/xxxxx.jpeg
	FaceContrastPictureUrl *string `json:"FaceContrastPictureUrl,omitempty" xml:"FaceContrastPictureUrl,omitempty"`
	// The type of the Device Assistant tag. Set the value to **DeviceRisk**.
	//
	// >
	//
	// > - Outputting Device Assistant tags incurs additional fees. For more information, see [Value-added services](https://www.alibabacloud.com/help/en/id-verification/financial-grade-id-verification/face-guard).
	//
	// > - If you do not need to output Device Assistant tags, you can leave this parameter empty or do not pass it.
	//
	// example:
	//
	// DeviceRisk
	FaceGuardOutput *string `json:"FaceGuardOutput,omitempty" xml:"FaceGuardOutput,omitempty"`
	// Specifies whether to display the I Have Completed Authentication button on the H5 downgrade page after authentication is completed. Valid values:
	//
	// - **Y**: Displays the button.
	//
	// - **N*	- (default): Hides the button.
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
	// The Metainfo environment parameter. You must obtain this parameter by using the client SDK.
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
	// The method to obtain the passport NFC verification elements. Valid values:
	//
	// - **1**: User input. The end user manually enters the ID document element information by using the UI provided by the Alibaba Cloud SDK.
	//
	// - **3**: External parameter input. The ID document element information is passed from external sources.
	//
	// > When the NFC decoder decrypts the encrypted information of the passport chip, it needs to obtain three passport elements: name, date of birth, and expiration date.
	//
	// example:
	//
	// 1
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The liveness detection type. Valid values:
	//
	// > Only the following liveness detection types are supported. Custom actions or combinations are not supported.
	//
	// Description
	//
	// Only the following liveness detection types are supported. Custom actions or combinations are not supported.
	//
	// - **LIVENESS*	- (default): Blinking.
	//
	// - **PHOTINUS_LIVENESS**: Blinking and colorful light reflection.
	//
	// - **MULTI_ACTION**: Blinking and head turning (the order is random).
	//
	// - **MOVE_ACTION*	- (recommended): Moving closer or farther and blinking.
	//
	// - **MOVE_PHOTINUS**: Moving closer or farther and colorful light reflection.
	//
	// >
	//
	// >- **The default liveness detection type*	- is supported in the following SDK versions:
	//
	// >   - Android SDK 1.2.6 and later
	//
	// >   - iOS SDK 1.2.4 and later
	//
	// >   - Harmony SDK 1.0.0 and later
	//
	// >- Other types are supported in the latest SDK versions for Android, iOS, and Harmony. We recommend that you integrate the latest SDK version.
	//
	// example:
	//
	// MOVE_ACTION
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// Specifies whether to intercept the request when multiple faces are detected on the terminal. Valid values:
	//
	// - **Y**: Intercepts the request. The client prompts the user to scan the face again.
	//
	// - **N*	- (default): Does not intercept the request. The largest face in the face scanning screen is sent to the server for security detection.
	//
	// example:
	//
	// Y
	NeedMultiFaceCheck *string `json:"NeedMultiFaceCheck,omitempty" xml:"NeedMultiFaceCheck,omitempty"`
	// The name of the authorized OSS bucket.
	//
	// > You can specify only one of the following parameters to pass the image: FaceContrastPicture, FaceContrastPictureUrl, CertifyId, and OSS.
	//
	// example:
	//
	// cn-shanghai-aliyun-cloudauth-xxxxx
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// The name of the object in the authorized OSS bucket.
	//
	// > You can specify only one of the following parameters to pass the image: FaceContrastPicture, FaceContrastPictureUrl, CertifyId, and OSS.
	//
	// example:
	//
	// verify/xxxxx/xxxxxx.jpeg
	OssObjectName *string `json:"OssObjectName,omitempty" xml:"OssObjectName,omitempty"`
	// The unique identifier of the merchant request.
	//
	// The value is a 32-character alphanumeric string. The first few characters are a custom abbreviation of the merchant, the middle part can be a period of time, and the last part can be a random or incremental sequence.
	//
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c353888
	OuterOrderNo *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	// The downgrade configuration when WebRTC or WebAssembly is incompatible during mobile H5 authentication. Valid values:
	//
	// - **keep**: Downgrade is not supported. The result is directly returned.
	//
	// - **url*	- (default): Downgrade is supported. The authentication URL is returned. The user opens or switches to a browser by using this URL for authentication.
	//
	// - **video**: Downgrade is supported. The system camera is used to record a 3- to 5-second blinking video for authentication.
	//
	//
	// >
	//
	// > When the downgrade mode is set to video, the following features become invalid and product security is reduced. We recommend that you configure this mode only for security scenarios.
	//
	// > - The liveness detection type setting does not take effect.
	//
	// > - The VideoEvidence feature is not supported.
	//
	// example:
	//
	// url
	ProcedurePriority *string `json:"ProcedurePriority,omitempty" xml:"ProcedurePriority,omitempty"`
	// The fixed value. The value of this parameter varies based on the product solution:
	//
	// - APP authentication solution: Set the value to ID_PRO.
	//
	// - Face liveness verification solution: Set the value to PV_FV.
	//
	// - Liveness detection solution: Set the value to LR_FR.
	//
	// example:
	//
	// ID_PRO
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// Specifies whether to enable the rare character mode. Valid values:
	//
	// - **Y**: Enables the mode. Before authentication, an information input box appears. The user must enter the name with rare characters and the ID card number, and agree to the agreement before the authentication process starts.
	//
	// - **N**: Does not enable the mode. This is the default value.
	//
	// example:
	//
	// Y
	RarelyCharacters *string `json:"RarelyCharacters,omitempty" xml:"RarelyCharacters,omitempty"`
	// Specifies whether to read the ID document photo. Valid values:
	//
	// - **Y**: Reads the photo.
	//
	// - **N**: Does not read the photo.
	//
	// > If you need to use the face photo from the ID document in subsequent authentication steps, we recommend that you set this parameter to Y.
	//
	// example:
	//
	// Y
	ReadImg *string `json:"ReadImg,omitempty" xml:"ReadImg,omitempty"`
	// The redirect URL for the merchant business page.
	//
	// example:
	//
	// www.aliyun.com
	ReturnUrl *string `json:"ReturnUrl,omitempty" xml:"ReturnUrl,omitempty"`
	// The ID of the authentication scenario.
	//
	// example:
	//
	// 1000000006
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The age-friendly configuration parameter. This parameter takes effect for each authentication request. You can select different parameters for each authentication request based on the business property, customer distribution, and operation attribute of your app. The following options are available. Default value: 0.
	//
	// - **0**: Disables the age-friendly pattern for the current authentication request.
	//
	// - **1**: Enables the age-friendly pattern for the current authentication request.
	//
	// - **2**: Allows the user to select the pattern.
	//
	//
	// The end user can select the authentication pattern. The product guide page provides two authentication entries: Start Authentication and Senior Authentication Pattern. When the user selects Senior Authentication Pattern, the system enters the age-friendly pattern.
	//
	// >
	//
	// > - The age-friendly parameter takes effect only when the liveness detection type **Model*	- is set to **LIVENESS*	- or **MULTI_ACTION**.
	//
	// example:
	//
	// 0
	SuitableType *string `json:"SuitableType,omitempty" xml:"SuitableType,omitempty"`
	// The URL of the UI configuration file.
	//
	// For the complete configuration, see [Web SDK UI custom configuration](https://www.alibabacloud.com/help/en/id-verification/financial-grade-id-verification/web-sdk-ui-custom-configuration-description).
	//
	// example:
	//
	// www.aliyundoc.com
	UiCustomUrl *string `json:"UiCustomUrl,omitempty" xml:"UiCustomUrl,omitempty"`
	// The custom user ID in the customer business. Make sure that the ID is unique.
	//
	// example:
	//
	// 123456789
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The expiration date of the ID document.
	//
	// This parameter is required when the **CertType*	- parameter is set to **PASSPORT*	- and the **Mode*	- parameter is set to **3**.
	//
	// example:
	//
	// 2039-06-10
	ValidityDate *string `json:"ValidityDate,omitempty" xml:"ValidityDate,omitempty"`
	// Specifies whether to enable video evidence. Valid values:
	//
	// - **true**: Enables video evidence.
	//
	// - **false**: Does not enable video evidence. This is the default value.
	//
	// > Because video files are large, the system discards video files to prioritize the transmission of necessary authentication images when the network is unstable. We recommend that you set your business to weakly depend on videos.
	//
	// example:
	//
	// false
	VideoEvidence *string `json:"VideoEvidence,omitempty" xml:"VideoEvidence,omitempty"`
	// The custom willingness content. This parameter is required when personalized settings are enabled. The value is a JSON string of a string list.
	//
	// - In the read-along scenario: The content cannot exceed 60 Chinese characters (excluding punctuation marks), and the list contains only one element.
	//
	// - In the Q&A scenario: You can configure up to three questions. Each question cannot exceed 30 Chinese characters, and each question is an element in the list.
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

func (s *InitFaceVerifyRequest) GetEnableAlbumUpload() *string {
	return s.EnableAlbumUpload
}

func (s *InitFaceVerifyRequest) GetEnableBeauty() *string {
	return s.EnableBeauty
}

func (s *InitFaceVerifyRequest) GetEnableTts() *string {
	return s.EnableTts
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

func (s *InitFaceVerifyRequest) SetEnableAlbumUpload(v string) *InitFaceVerifyRequest {
	s.EnableAlbumUpload = &v
	return s
}

func (s *InitFaceVerifyRequest) SetEnableBeauty(v string) *InitFaceVerifyRequest {
	s.EnableBeauty = &v
	return s
}

func (s *InitFaceVerifyRequest) SetEnableTts(v string) *InitFaceVerifyRequest {
	s.EnableTts = &v
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
