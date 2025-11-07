// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyMaterialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorityComparisionScore(v float32) *VerifyMaterialResponseBody
	GetAuthorityComparisionScore() *float32
	SetIdCardFaceComparisonScore(v float32) *VerifyMaterialResponseBody
	GetIdCardFaceComparisonScore() *float32
	SetMaterial(v *VerifyMaterialResponseBodyMaterial) *VerifyMaterialResponseBody
	GetMaterial() *VerifyMaterialResponseBodyMaterial
	SetRequestId(v string) *VerifyMaterialResponseBody
	GetRequestId() *string
	SetVerifyStatus(v int32) *VerifyMaterialResponseBody
	GetVerifyStatus() *int32
	SetVerifyToken(v string) *VerifyMaterialResponseBody
	GetVerifyToken() *string
}

type VerifyMaterialResponseBody struct {
	// Comparison score between the facial photo submitted during the authentication process and authoritative data, with a range of **0**~**100**.
	//
	// Confidence threshold references:
	//
	// - False recognition rate 0.001% corresponds to a threshold of 95.
	//
	// - False recognition rate 0.01% corresponds to a threshold of 90.
	//
	// - False recognition rate 0.1% corresponds to a threshold of 80.
	//
	// - False recognition rate 1% corresponds to a threshold of 60.
	//
	// > This field only indicates the comparison result between the face and authoritative data, serving as a reference score. It is generally not recommended to use this score alone as the pass/fail criterion. For the comprehensive authentication result, please refer to the **VerifyStatus*	- field. The **VerifyStatus*	- result integrates the face-to-authoritative data comparison and other various strategies, enhancing security levels.
	//
	// example:
	//
	// 97
	AuthorityComparisionScore *float32 `json:"AuthorityComparisionScore,omitempty" xml:"AuthorityComparisionScore,omitempty"`
	// Comparison score between the facial photo submitted during the authentication process and the face on the portrait side of the ID card image, with a range of **0**~**100**.
	//
	// Confidence threshold references:
	//
	// - False recognition rate 0.001% corresponds to a threshold of 95.
	//
	// - False recognition rate 0.01% corresponds to a threshold of 90.
	//
	// - False recognition rate 0.1% corresponds to a threshold of 80.
	//
	// - False recognition rate 1% corresponds to a threshold of 60.
	//
	// example:
	//
	// 97
	IdCardFaceComparisonScore *float32 `json:"IdCardFaceComparisonScore,omitempty" xml:"IdCardFaceComparisonScore,omitempty"`
	// Authentication materials.
	Material *VerifyMaterialResponseBodyMaterial `json:"Material,omitempty" xml:"Material,omitempty" type:"Struct"`
	// Request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Authentication status. Values:
	//
	// - **1**: Authentication passed.
	//
	// - **2**~**n**: Authentication failed due to various reasons. For detailed descriptions, see the **Authentication Status Explanation*	- below.
	//
	// example:
	//
	// 1
	VerifyStatus *int32 `json:"VerifyStatus,omitempty" xml:"VerifyStatus,omitempty"`
	// Token for this authentication, used to link various interfaces in the authentication request, valid for 30 minutes.
	//
	// example:
	//
	// c302c0797679457685410ee51a5ba375
	VerifyToken *string `json:"VerifyToken,omitempty" xml:"VerifyToken,omitempty"`
}

func (s VerifyMaterialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyMaterialResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyMaterialResponseBody) GetAuthorityComparisionScore() *float32 {
	return s.AuthorityComparisionScore
}

func (s *VerifyMaterialResponseBody) GetIdCardFaceComparisonScore() *float32 {
	return s.IdCardFaceComparisonScore
}

func (s *VerifyMaterialResponseBody) GetMaterial() *VerifyMaterialResponseBodyMaterial {
	return s.Material
}

func (s *VerifyMaterialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyMaterialResponseBody) GetVerifyStatus() *int32 {
	return s.VerifyStatus
}

func (s *VerifyMaterialResponseBody) GetVerifyToken() *string {
	return s.VerifyToken
}

func (s *VerifyMaterialResponseBody) SetAuthorityComparisionScore(v float32) *VerifyMaterialResponseBody {
	s.AuthorityComparisionScore = &v
	return s
}

func (s *VerifyMaterialResponseBody) SetIdCardFaceComparisonScore(v float32) *VerifyMaterialResponseBody {
	s.IdCardFaceComparisonScore = &v
	return s
}

func (s *VerifyMaterialResponseBody) SetMaterial(v *VerifyMaterialResponseBodyMaterial) *VerifyMaterialResponseBody {
	s.Material = v
	return s
}

func (s *VerifyMaterialResponseBody) SetRequestId(v string) *VerifyMaterialResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyMaterialResponseBody) SetVerifyStatus(v int32) *VerifyMaterialResponseBody {
	s.VerifyStatus = &v
	return s
}

func (s *VerifyMaterialResponseBody) SetVerifyToken(v string) *VerifyMaterialResponseBody {
	s.VerifyToken = &v
	return s
}

func (s *VerifyMaterialResponseBody) Validate() error {
	if s.Material != nil {
		if err := s.Material.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VerifyMaterialResponseBodyMaterial struct {
	// Global camera image captured by the real-person authentication SDK.
	//
	// > This parameter will take effect after configuration. If you need to use this parameter, please submit a [ticket](https://selfservice.console.aliyun.com/ticket/category/cloudauth/today) to contact us.
	//
	// example:
	//
	// http://image-demo.img-cn-hangzhou.aliyuncs.com/face-global-example.jpg
	FaceGlobalUrl *string `json:"FaceGlobalUrl,omitempty" xml:"FaceGlobalUrl,omitempty"`
	// HTTP or HTTPS link to the frontal face image, corresponding to the request parameter **FaceImageUrl**. The link is valid for 5 minutes, and it is recommended to store it in your business to avoid affecting usage.
	//
	// example:
	//
	// http://image-demo.img-cn-hangzhou.aliyuncs.com/example.jpg
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
	FaceMask *string `json:"FaceMask,omitempty" xml:"FaceMask,omitempty"`
	// The quality of the frontal face image. Possible values:
	//
	// - **UNQUALIFIED**: Poor quality
	//
	// - **LOW**: Low
	//
	// - **NORMAL**: Average
	//
	// - **HIGH**: High
	//
	// example:
	//
	// NORMAL
	FaceQuality *string `json:"FaceQuality,omitempty" xml:"FaceQuality,omitempty"`
	// OCR result of the ID card information.
	//
	// > If there is no front or back of the ID card during the verification process, the OCR result of the ID card information will not be returned. Even if the front and back of the ID card are present during the verification process, it does not guarantee that all the information on the ID card will be returned. Due to issues such as poor ID card photography, the OCR may fail to recognize some information, leading to incomplete OCR results. It is recommended that the business side does not heavily rely on the ID card OCR information.
	IdCardInfo *VerifyMaterialResponseBodyMaterialIdCardInfo `json:"IdCardInfo,omitempty" xml:"IdCardInfo,omitempty" type:"Struct"`
	// Name, corresponding to the request parameter **Name**.
	//
	// example:
	//
	// 张三
	IdCardName *string `json:"IdCardName,omitempty" xml:"IdCardName,omitempty"`
	// ID number, corresponding to the request parameter **IdCardNumber**.
	//
	// example:
	//
	// 02343218901123****
	IdCardNumber *string `json:"IdCardNumber,omitempty" xml:"IdCardNumber,omitempty"`
}

func (s VerifyMaterialResponseBodyMaterial) String() string {
	return dara.Prettify(s)
}

func (s VerifyMaterialResponseBodyMaterial) GoString() string {
	return s.String()
}

func (s *VerifyMaterialResponseBodyMaterial) GetFaceGlobalUrl() *string {
	return s.FaceGlobalUrl
}

func (s *VerifyMaterialResponseBodyMaterial) GetFaceImageUrl() *string {
	return s.FaceImageUrl
}

func (s *VerifyMaterialResponseBodyMaterial) GetFaceMask() *string {
	return s.FaceMask
}

func (s *VerifyMaterialResponseBodyMaterial) GetFaceQuality() *string {
	return s.FaceQuality
}

func (s *VerifyMaterialResponseBodyMaterial) GetIdCardInfo() *VerifyMaterialResponseBodyMaterialIdCardInfo {
	return s.IdCardInfo
}

func (s *VerifyMaterialResponseBodyMaterial) GetIdCardName() *string {
	return s.IdCardName
}

func (s *VerifyMaterialResponseBodyMaterial) GetIdCardNumber() *string {
	return s.IdCardNumber
}

func (s *VerifyMaterialResponseBodyMaterial) SetFaceGlobalUrl(v string) *VerifyMaterialResponseBodyMaterial {
	s.FaceGlobalUrl = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterial) SetFaceImageUrl(v string) *VerifyMaterialResponseBodyMaterial {
	s.FaceImageUrl = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterial) SetFaceMask(v string) *VerifyMaterialResponseBodyMaterial {
	s.FaceMask = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterial) SetFaceQuality(v string) *VerifyMaterialResponseBodyMaterial {
	s.FaceQuality = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterial) SetIdCardInfo(v *VerifyMaterialResponseBodyMaterialIdCardInfo) *VerifyMaterialResponseBodyMaterial {
	s.IdCardInfo = v
	return s
}

func (s *VerifyMaterialResponseBodyMaterial) SetIdCardName(v string) *VerifyMaterialResponseBodyMaterial {
	s.IdCardName = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterial) SetIdCardNumber(v string) *VerifyMaterialResponseBodyMaterial {
	s.IdCardNumber = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterial) Validate() error {
	if s.IdCardInfo != nil {
		if err := s.IdCardInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VerifyMaterialResponseBodyMaterialIdCardInfo struct {
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
	// HTTP or HTTPS link to the national emblem side of the ID card. The link is valid for 5 minutes. It is recommended to store it in your business system to avoid any impact on usage.
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
	// End date of the document\\"s validity period. Format: yyyymmdd.
	//
	// example:
	//
	// 20201101
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// HTTP or HTTPS link to the portrait side of the ID card. The link is valid for 5 minutes. It is recommended to store it in your business system to avoid any impact on usage.
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
	// ID number.
	//
	// example:
	//
	// 02343218901123****
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// Start date of the document\\"s validity period. Format: yyyymmdd.
	//
	// example:
	//
	// 20201101
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s VerifyMaterialResponseBodyMaterialIdCardInfo) String() string {
	return dara.Prettify(s)
}

func (s VerifyMaterialResponseBodyMaterialIdCardInfo) GoString() string {
	return s.String()
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) GetAddress() *string {
	return s.Address
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) GetAuthority() *string {
	return s.Authority
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) GetBackImageUrl() *string {
	return s.BackImageUrl
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) GetBirth() *string {
	return s.Birth
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) GetEndDate() *string {
	return s.EndDate
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) GetFrontImageUrl() *string {
	return s.FrontImageUrl
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) GetName() *string {
	return s.Name
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) GetNationality() *string {
	return s.Nationality
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) GetNumber() *string {
	return s.Number
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) GetStartDate() *string {
	return s.StartDate
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) SetAddress(v string) *VerifyMaterialResponseBodyMaterialIdCardInfo {
	s.Address = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) SetAuthority(v string) *VerifyMaterialResponseBodyMaterialIdCardInfo {
	s.Authority = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) SetBackImageUrl(v string) *VerifyMaterialResponseBodyMaterialIdCardInfo {
	s.BackImageUrl = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) SetBirth(v string) *VerifyMaterialResponseBodyMaterialIdCardInfo {
	s.Birth = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) SetEndDate(v string) *VerifyMaterialResponseBodyMaterialIdCardInfo {
	s.EndDate = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) SetFrontImageUrl(v string) *VerifyMaterialResponseBodyMaterialIdCardInfo {
	s.FrontImageUrl = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) SetName(v string) *VerifyMaterialResponseBodyMaterialIdCardInfo {
	s.Name = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) SetNationality(v string) *VerifyMaterialResponseBodyMaterialIdCardInfo {
	s.Nationality = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) SetNumber(v string) *VerifyMaterialResponseBodyMaterialIdCardInfo {
	s.Number = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) SetStartDate(v string) *VerifyMaterialResponseBodyMaterialIdCardInfo {
	s.StartDate = &v
	return s
}

func (s *VerifyMaterialResponseBodyMaterialIdCardInfo) Validate() error {
	return dara.Validate(s)
}
