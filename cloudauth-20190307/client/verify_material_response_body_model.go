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
	// The comparison score between the facial photo submitted during verification and the authoritative data. Value range: **0*	- to **100**.
	//
	// Confidence thresholds for reference:
	//
	// - At a false acceptance rate of 0.001%, the corresponding threshold is 95.
	//
	// - At a false acceptance rate of 0.01%, the corresponding threshold is 90.
	//
	// - At a false acceptance rate of 0.1%, the corresponding threshold is 80.
	//
	// - At a false acceptance rate of 1%, the corresponding threshold is 60.
	//
	// > This field only represents the comparison result between the face and the authoritative data and serves as a reference score. We do not recommend using this score alone as the pass/fail criterion. For the comprehensive verification result, refer to the **VerifyStatus*	- field. The **VerifyStatus*	- result combines the face-to-authoritative-data comparison with multiple other strategies to improve the security level.
	//
	// example:
	//
	// 97
	AuthorityComparisionScore *float32 `json:"AuthorityComparisionScore,omitempty" xml:"AuthorityComparisionScore,omitempty"`
	// The comparison score between the facial photo submitted during verification and the face on the portrait side of the ID card. Value range: **0*	- to **100**.
	//
	// Confidence thresholds for reference:
	//
	// - At a false acceptance rate of 0.001%, the corresponding threshold is 95.
	//
	// - At a false acceptance rate of 0.01%, the corresponding threshold is 90.
	//
	// - At a false acceptance rate of 0.1%, the corresponding threshold is 80.
	//
	// - At a false acceptance rate of 1%, the corresponding threshold is 60.
	//
	// example:
	//
	// 97
	IdCardFaceComparisonScore *float32 `json:"IdCardFaceComparisonScore,omitempty" xml:"IdCardFaceComparisonScore,omitempty"`
	// The verification materials.
	Material *VerifyMaterialResponseBodyMaterial `json:"Material,omitempty" xml:"Material,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The verification status. Valid values:
	//
	// - **1**: Verification passed.
	//
	// - **2*	- to **n**: Verification failed due to various reasons. For detailed descriptions, see **Verification status description*	- below.
	//
	// example:
	//
	// 1
	VerifyStatus *int32 `json:"VerifyStatus,omitempty" xml:"VerifyStatus,omitempty"`
	// The token for this verification, used to correlate the various operations within a verification request. The token is valid for 30 minutes.
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
	// The global camera image captured by the ID Verification SDK.
	//
	// > This parameter takes effect only after configuration. If you need to use this parameter, [submit a ticket](https://selfservice.console.aliyun.com/ticket/category/cloudauth/today) to contact us.
	//
	// example:
	//
	// http://image-demo.img-cn-hangzhou.aliyuncs.com/face-global-example.jpg
	FaceGlobalUrl *string `json:"FaceGlobalUrl,omitempty" xml:"FaceGlobalUrl,omitempty"`
	// The HTTP or HTTPS URL of the front-facing facial photo, corresponding to the request parameter **FaceImageUrl**. The URL is valid for 5 minutes. Save the image to your own storage to avoid access issues.
	//
	// example:
	//
	// http://image-demo.img-cn-hangzhou.aliyuncs.com/example.jpg
	FaceImageUrl *string `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty"`
	// Indicates whether the face is wearing a mask. Valid values:
	//
	// - **true**: A mask is detected.
	//
	// - **false**: No mask is detected.
	//
	// example:
	//
	// false
	FaceMask *string `json:"FaceMask,omitempty" xml:"FaceMask,omitempty"`
	// The quality of the front-facing facial photo. Valid values:
	//
	// - **UNQUALIFIED**: poor quality
	//
	// - **LOW**: low quality
	//
	// - **NORMAL**: moderate quality
	//
	// - **HIGH**: high quality.
	//
	// example:
	//
	// NORMAL
	FaceQuality *string `json:"FaceQuality,omitempty" xml:"FaceQuality,omitempty"`
	// The OCR result of the ID card information.
	//
	// > If the front and back photos of the ID card are not provided during verification, the OCR result of the ID card information is not returned. Even if both photos are provided, not all information on the ID card is guaranteed to be returned. OCR may fail to recognize certain information due to issues such as poor photo quality. We recommend that your business logic does not strictly depend on the ID card OCR information.
	IdCardInfo *VerifyMaterialResponseBodyMaterialIdCardInfo `json:"IdCardInfo,omitempty" xml:"IdCardInfo,omitempty" type:"Struct"`
	// The name, corresponding to the request parameter **Name**.
	//
	// example:
	//
	// 张三
	IdCardName *string `json:"IdCardName,omitempty" xml:"IdCardName,omitempty"`
	// The ID card number, corresponding to the request parameter **IdCardNumber**.
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
	// The address.
	//
	// example:
	//
	// 浙江省杭州市余杭区文一西路969号
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The issuing authority.
	//
	// example:
	//
	// 杭州市公安局
	Authority *string `json:"Authority,omitempty" xml:"Authority,omitempty"`
	// The HTTP or HTTPS URL of the national emblem side of the ID card. The URL is valid for 5 minutes. Save the image to your own storage to avoid access issues.
	//
	// example:
	//
	// http://image-demo.img-cn-hangzhou.aliyuncs.com/example3.jpg
	BackImageUrl *string `json:"BackImageUrl,omitempty" xml:"BackImageUrl,omitempty"`
	// The date of birth.
	//
	// example:
	//
	// 19900101
	Birth *string `json:"Birth,omitempty" xml:"Birth,omitempty"`
	// The expiration date of the ID card. Format: yyyymmdd.
	//
	// example:
	//
	// 20201101
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The HTTP or HTTPS URL of the portrait side of the ID card. The URL is valid for 5 minutes. Save the image to your own storage to avoid access issues.
	//
	// example:
	//
	// http://image-demo.img-cn-hangzhou.aliyuncs.com/example2.jpg
	FrontImageUrl *string `json:"FrontImageUrl,omitempty" xml:"FrontImageUrl,omitempty"`
	// The name.
	//
	// example:
	//
	// 张三
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ethnicity.
	//
	// example:
	//
	// 汉
	Nationality *string `json:"Nationality,omitempty" xml:"Nationality,omitempty"`
	// The ID card number.
	//
	// example:
	//
	// 02343218901123****
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// The start date of the ID card validity period. Format: yyyymmdd.
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
