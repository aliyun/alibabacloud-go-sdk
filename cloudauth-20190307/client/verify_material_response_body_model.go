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
	// example:
	//
	// 97
	AuthorityComparisionScore *float32 `json:"AuthorityComparisionScore,omitempty" xml:"AuthorityComparisionScore,omitempty"`
	// example:
	//
	// 97
	IdCardFaceComparisonScore *float32                            `json:"IdCardFaceComparisonScore,omitempty" xml:"IdCardFaceComparisonScore,omitempty"`
	Material                  *VerifyMaterialResponseBodyMaterial `json:"Material,omitempty" xml:"Material,omitempty" type:"Struct"`
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	VerifyStatus *int32 `json:"VerifyStatus,omitempty" xml:"VerifyStatus,omitempty"`
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
	return dara.Validate(s)
}

type VerifyMaterialResponseBodyMaterial struct {
	// example:
	//
	// http://image-demo.img-cn-hangzhou.aliyuncs.com/face-global-example.jpg
	FaceGlobalUrl *string `json:"FaceGlobalUrl,omitempty" xml:"FaceGlobalUrl,omitempty"`
	// example:
	//
	// http://image-demo.img-cn-hangzhou.aliyuncs.com/example.jpg
	FaceImageUrl *string `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty"`
	// example:
	//
	// false
	FaceMask *string `json:"FaceMask,omitempty" xml:"FaceMask,omitempty"`
	// example:
	//
	// NORMAL
	FaceQuality *string                                       `json:"FaceQuality,omitempty" xml:"FaceQuality,omitempty"`
	IdCardInfo  *VerifyMaterialResponseBodyMaterialIdCardInfo `json:"IdCardInfo,omitempty" xml:"IdCardInfo,omitempty" type:"Struct"`
	IdCardName  *string                                       `json:"IdCardName,omitempty" xml:"IdCardName,omitempty"`
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
	return dara.Validate(s)
}

type VerifyMaterialResponseBodyMaterialIdCardInfo struct {
	Address   *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Authority *string `json:"Authority,omitempty" xml:"Authority,omitempty"`
	// example:
	//
	// http://image-demo.img-cn-hangzhou.aliyuncs.com/example3.jpg
	BackImageUrl *string `json:"BackImageUrl,omitempty" xml:"BackImageUrl,omitempty"`
	// example:
	//
	// 19900101
	Birth *string `json:"Birth,omitempty" xml:"Birth,omitempty"`
	// example:
	//
	// 20201101
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// http://image-demo.img-cn-hangzhou.aliyuncs.com/example2.jpg
	FrontImageUrl *string `json:"FrontImageUrl,omitempty" xml:"FrontImageUrl,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Nationality   *string `json:"Nationality,omitempty" xml:"Nationality,omitempty"`
	// example:
	//
	// 02343218901123****
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
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
