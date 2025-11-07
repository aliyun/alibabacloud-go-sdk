// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifyResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorityComparisionScore(v float32) *DescribeVerifyResultResponseBody
	GetAuthorityComparisionScore() *float32
	SetFaceComparisonScore(v float32) *DescribeVerifyResultResponseBody
	GetFaceComparisonScore() *float32
	SetIdCardFaceComparisonScore(v float32) *DescribeVerifyResultResponseBody
	GetIdCardFaceComparisonScore() *float32
	SetMaterial(v *DescribeVerifyResultResponseBodyMaterial) *DescribeVerifyResultResponseBody
	GetMaterial() *DescribeVerifyResultResponseBodyMaterial
	SetRequestId(v string) *DescribeVerifyResultResponseBody
	GetRequestId() *string
	SetVerifyStatus(v int32) *DescribeVerifyResultResponseBody
	GetVerifyStatus() *int32
}

type DescribeVerifyResultResponseBody struct {
	// The comparison score between the face photo submitted during the authentication process and the authoritative data, with a value range of **0*	- to **100**.
	//
	// Confidence threshold references:
	//
	// - When the false acceptance rate is 0.001%, the corresponding threshold is 95. - When the false acceptance rate is 0.01%, the corresponding threshold is 90. - When the false acceptance rate is 0.1%, the corresponding threshold is 80. - When the false acceptance rate is 1%, the corresponding threshold is 60.
	//
	// > This field only indicates the comparison result between the face and the authoritative data, for your reference only. It is generally not recommended to use this value alone as the standard for whether the authentication passes. For a comprehensive authentication result, please refer to the **VerifyStatus*	- field. The **VerifyStatus*	- result integrates the comparison of the face with the authoritative data and various other strategies, which can enhance security levels.
	//
	// example:
	//
	// 97
	AuthorityComparisionScore *float32 `json:"AuthorityComparisionScore,omitempty" xml:"AuthorityComparisionScore,omitempty"`
	// The comparison score between the face photo submitted during the authentication process and the face in the retained face image. The value range is **0**~**100**.
	//
	// Confidence threshold reference:
	//
	// - When the false recognition rate is 0.001%, the corresponding threshold is 95.
	//
	// - When the false recognition rate is 0.01%, the corresponding threshold is 90.
	//
	// - When the false recognition rate is 0.1%, the corresponding threshold is 80.
	//
	// - When the false recognition rate is 1%, the corresponding threshold is 60.
	//
	// example:
	//
	// 97
	FaceComparisonScore *float32 `json:"FaceComparisonScore,omitempty" xml:"FaceComparisonScore,omitempty"`
	// The comparison score between the face photo submitted during the authentication process and the face on the ID card\\"s face side. The value range is **0**~**100**.
	//
	// Confidence threshold reference:
	//
	// - When the false recognition rate is 0.001%, the corresponding threshold is 95.
	//
	// - When the false recognition rate is 0.01%, the corresponding threshold is 90.
	//
	// - When the false recognition rate is 0.1%, the corresponding threshold is 80.
	//
	// - When the false recognition rate is 1%, the corresponding threshold is 60.
	//
	// example:
	//
	// 97
	IdCardFaceComparisonScore *float32 `json:"IdCardFaceComparisonScore,omitempty" xml:"IdCardFaceComparisonScore,omitempty"`
	// Authentication materials.
	Material *DescribeVerifyResultResponseBodyMaterial `json:"Material,omitempty" xml:"Material,omitempty" type:"Struct"`
	// The ID of this request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Authentication status, values:
	//
	// - **-1**: Not authenticated - **1**: Authentication passed - **2*	- to **n**: Authentication failed for various reasons. For detailed descriptions, see the authentication status explanation.
	//
	// example:
	//
	// 1
	VerifyStatus *int32 `json:"VerifyStatus,omitempty" xml:"VerifyStatus,omitempty"`
}

func (s DescribeVerifyResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVerifyResultResponseBody) GetAuthorityComparisionScore() *float32 {
	return s.AuthorityComparisionScore
}

func (s *DescribeVerifyResultResponseBody) GetFaceComparisonScore() *float32 {
	return s.FaceComparisonScore
}

func (s *DescribeVerifyResultResponseBody) GetIdCardFaceComparisonScore() *float32 {
	return s.IdCardFaceComparisonScore
}

func (s *DescribeVerifyResultResponseBody) GetMaterial() *DescribeVerifyResultResponseBodyMaterial {
	return s.Material
}

func (s *DescribeVerifyResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVerifyResultResponseBody) GetVerifyStatus() *int32 {
	return s.VerifyStatus
}

func (s *DescribeVerifyResultResponseBody) SetAuthorityComparisionScore(v float32) *DescribeVerifyResultResponseBody {
	s.AuthorityComparisionScore = &v
	return s
}

func (s *DescribeVerifyResultResponseBody) SetFaceComparisonScore(v float32) *DescribeVerifyResultResponseBody {
	s.FaceComparisonScore = &v
	return s
}

func (s *DescribeVerifyResultResponseBody) SetIdCardFaceComparisonScore(v float32) *DescribeVerifyResultResponseBody {
	s.IdCardFaceComparisonScore = &v
	return s
}

func (s *DescribeVerifyResultResponseBody) SetMaterial(v *DescribeVerifyResultResponseBodyMaterial) *DescribeVerifyResultResponseBody {
	s.Material = v
	return s
}

func (s *DescribeVerifyResultResponseBody) SetRequestId(v string) *DescribeVerifyResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVerifyResultResponseBody) SetVerifyStatus(v int32) *DescribeVerifyResultResponseBody {
	s.VerifyStatus = &v
	return s
}

func (s *DescribeVerifyResultResponseBody) Validate() error {
	if s.Material != nil {
		if err := s.Material.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVerifyResultResponseBodyMaterial struct {
	// The global camera image captured by the real-person authentication SDK.
	//
	// > This parameter will only take effect after configuration. If you need to use this parameter, please submit a [ticket](https://selfservice.console.aliyun.com/ticket/category/cloudauth/today) to contact us.
	//
	// example:
	//
	// http://image-demo.img-cn-hangzhou.aliyuncs.com/face-global-example.jpg
	FaceGlobalUrl *string `json:"FaceGlobalUrl,omitempty" xml:"FaceGlobalUrl,omitempty"`
	// The HTTP or HTTPS link to the frontal face image. The link is valid for 5 minutes, and it is recommended to store it elsewhere to avoid any impact on usage.
	//
	// > If the HTTP or HTTPS link to the frontal face image expires, you can call [DescribeVerifyResult](https://help.aliyun.com/document_detail/154606.html) again to get the link.
	//
	// example:
	//
	// http://image-demo.img-cn-hangzhou.aliyuncs.com/face-image-example.jpg
	FaceImageUrl *string `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty"`
	// Whether the face is wearing a mask. Values:
	//
	// - **true**: Wearing a mask
	//
	// - **false**: Not wearing a mask
	//
	// example:
	//
	// false
	FaceMask *bool `json:"FaceMask,omitempty" xml:"FaceMask,omitempty"`
	// The quality of the frontal face image. Values:
	//
	// - **UNQUALIFIED**: Poor quality
	//
	// - **LOW**: Low
	//
	// - **NORMAL**: Normal
	//
	// - **HIGH**: High
	//
	// example:
	//
	// NORMAL
	FaceQuality *string `json:"FaceQuality,omitempty" xml:"FaceQuality,omitempty"`
	// OCR results of the ID card information.
	//
	// > If there is no front and back information of the ID card during the authentication process, the real-person authentication service will not return the OCR results of the ID card. Even if there is front and back information of the ID card during the authentication process, the real-person authentication service does not guarantee to return all the information on the ID card. Due to issues with ID card photography, the OCR may fail to recognize some information, resulting in incomplete OCR information. It is recommended that your business does not strongly rely on the ID card OCR information.
	IdCardInfo *DescribeVerifyResultResponseBodyMaterialIdCardInfo `json:"IdCardInfo,omitempty" xml:"IdCardInfo,omitempty" type:"Struct"`
	// Name.
	//
	// example:
	//
	// 张三
	IdCardName *string `json:"IdCardName,omitempty" xml:"IdCardName,omitempty"`
	// ID number.
	//
	// example:
	//
	// 02343218901123****
	IdCardNumber *string `json:"IdCardNumber,omitempty" xml:"IdCardNumber,omitempty"`
	// The URL addresses of the recorded videos returned by the historical RPH5BioOnly solution.
	VideoUrls []*string `json:"VideoUrls,omitempty" xml:"VideoUrls,omitempty" type:"Repeated"`
}

func (s DescribeVerifyResultResponseBodyMaterial) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyResultResponseBodyMaterial) GoString() string {
	return s.String()
}

func (s *DescribeVerifyResultResponseBodyMaterial) GetFaceGlobalUrl() *string {
	return s.FaceGlobalUrl
}

func (s *DescribeVerifyResultResponseBodyMaterial) GetFaceImageUrl() *string {
	return s.FaceImageUrl
}

func (s *DescribeVerifyResultResponseBodyMaterial) GetFaceMask() *bool {
	return s.FaceMask
}

func (s *DescribeVerifyResultResponseBodyMaterial) GetFaceQuality() *string {
	return s.FaceQuality
}

func (s *DescribeVerifyResultResponseBodyMaterial) GetIdCardInfo() *DescribeVerifyResultResponseBodyMaterialIdCardInfo {
	return s.IdCardInfo
}

func (s *DescribeVerifyResultResponseBodyMaterial) GetIdCardName() *string {
	return s.IdCardName
}

func (s *DescribeVerifyResultResponseBodyMaterial) GetIdCardNumber() *string {
	return s.IdCardNumber
}

func (s *DescribeVerifyResultResponseBodyMaterial) GetVideoUrls() []*string {
	return s.VideoUrls
}

func (s *DescribeVerifyResultResponseBodyMaterial) SetFaceGlobalUrl(v string) *DescribeVerifyResultResponseBodyMaterial {
	s.FaceGlobalUrl = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterial) SetFaceImageUrl(v string) *DescribeVerifyResultResponseBodyMaterial {
	s.FaceImageUrl = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterial) SetFaceMask(v bool) *DescribeVerifyResultResponseBodyMaterial {
	s.FaceMask = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterial) SetFaceQuality(v string) *DescribeVerifyResultResponseBodyMaterial {
	s.FaceQuality = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterial) SetIdCardInfo(v *DescribeVerifyResultResponseBodyMaterialIdCardInfo) *DescribeVerifyResultResponseBodyMaterial {
	s.IdCardInfo = v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterial) SetIdCardName(v string) *DescribeVerifyResultResponseBodyMaterial {
	s.IdCardName = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterial) SetIdCardNumber(v string) *DescribeVerifyResultResponseBodyMaterial {
	s.IdCardNumber = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterial) SetVideoUrls(v []*string) *DescribeVerifyResultResponseBodyMaterial {
	s.VideoUrls = v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterial) Validate() error {
	if s.IdCardInfo != nil {
		if err := s.IdCardInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVerifyResultResponseBodyMaterialIdCardInfo struct {
	// Address.
	//
	// example:
	//
	// 浙江省杭州市余杭区文一西路969号
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// Issuing authority.
	//
	// example:
	//
	// 杭州市公安局
	Authority *string `json:"Authority,omitempty" xml:"Authority,omitempty"`
	// HTTP/HTTPS link to the image of the back side (national emblem side) of the ID card. The link is valid for 5 minutes, and it is recommended to store it for business use to avoid any impact.
	//
	// > If the HTTP/HTTPS link to the front-facing portrait image expires, you can call DescribeVerifyResult again to get the link.
	//
	// example:
	//
	// http://image-demo.img-cn-hangzhou.aliyuncs.com/example3.jpg
	BackImageUrl *string `json:"BackImageUrl,omitempty" xml:"BackImageUrl,omitempty"`
	// Date of birth.
	//
	// example:
	//
	// 19900101
	Birth *string `json:"Birth,omitempty" xml:"Birth,omitempty"`
	// The end date of the document\\"s validity period. Format: yyyymmdd.
	//
	// example:
	//
	// 20201101
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// HTTP/HTTPS link to the image of the front side (portrait side) of the ID card. The link is valid for 5 minutes, and it is recommended to store it for business use to avoid any impact.
	//
	// > If the HTTP/HTTPS link to the front-facing portrait image expires, you can call DescribeVerifyResult again to get the link.
	//
	// example:
	//
	// http://image-demo.img-cn-hangzhou.aliyuncs.com/example2.jpg
	FrontImageUrl *string `json:"FrontImageUrl,omitempty" xml:"FrontImageUrl,omitempty"`
	// Name.
	//
	// example:
	//
	// 张三
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Nationality.
	//
	// example:
	//
	// 汉
	Nationality *string `json:"Nationality,omitempty" xml:"Nationality,omitempty"`
	// ID card number.
	//
	// example:
	//
	// 02343218901123****
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// Start date of the document\\"s validity. Format: yyyymmdd.
	//
	// example:
	//
	// 20201101
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeVerifyResultResponseBodyMaterialIdCardInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyResultResponseBodyMaterialIdCardInfo) GoString() string {
	return s.String()
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) GetAddress() *string {
	return s.Address
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) GetAuthority() *string {
	return s.Authority
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) GetBackImageUrl() *string {
	return s.BackImageUrl
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) GetBirth() *string {
	return s.Birth
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) GetFrontImageUrl() *string {
	return s.FrontImageUrl
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) GetName() *string {
	return s.Name
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) GetNationality() *string {
	return s.Nationality
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) GetNumber() *string {
	return s.Number
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) SetAddress(v string) *DescribeVerifyResultResponseBodyMaterialIdCardInfo {
	s.Address = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) SetAuthority(v string) *DescribeVerifyResultResponseBodyMaterialIdCardInfo {
	s.Authority = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) SetBackImageUrl(v string) *DescribeVerifyResultResponseBodyMaterialIdCardInfo {
	s.BackImageUrl = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) SetBirth(v string) *DescribeVerifyResultResponseBodyMaterialIdCardInfo {
	s.Birth = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) SetEndDate(v string) *DescribeVerifyResultResponseBodyMaterialIdCardInfo {
	s.EndDate = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) SetFrontImageUrl(v string) *DescribeVerifyResultResponseBodyMaterialIdCardInfo {
	s.FrontImageUrl = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) SetName(v string) *DescribeVerifyResultResponseBodyMaterialIdCardInfo {
	s.Name = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) SetNationality(v string) *DescribeVerifyResultResponseBodyMaterialIdCardInfo {
	s.Nationality = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) SetNumber(v string) *DescribeVerifyResultResponseBodyMaterialIdCardInfo {
	s.Number = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) SetStartDate(v string) *DescribeVerifyResultResponseBodyMaterialIdCardInfo {
	s.StartDate = &v
	return s
}

func (s *DescribeVerifyResultResponseBodyMaterialIdCardInfo) Validate() error {
	return dara.Validate(s)
}
