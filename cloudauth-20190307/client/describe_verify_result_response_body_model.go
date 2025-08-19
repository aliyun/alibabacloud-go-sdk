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
	// example:
	//
	// 97
	AuthorityComparisionScore *float32 `json:"AuthorityComparisionScore,omitempty" xml:"AuthorityComparisionScore,omitempty"`
	// example:
	//
	// 97
	FaceComparisonScore *float32 `json:"FaceComparisonScore,omitempty" xml:"FaceComparisonScore,omitempty"`
	// example:
	//
	// 97
	IdCardFaceComparisonScore *float32                                  `json:"IdCardFaceComparisonScore,omitempty" xml:"IdCardFaceComparisonScore,omitempty"`
	Material                  *DescribeVerifyResultResponseBodyMaterial `json:"Material,omitempty" xml:"Material,omitempty" type:"Struct"`
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	return dara.Validate(s)
}

type DescribeVerifyResultResponseBodyMaterial struct {
	// example:
	//
	// http://image-demo.img-cn-hangzhou.aliyuncs.com/face-global-example.jpg
	FaceGlobalUrl *string `json:"FaceGlobalUrl,omitempty" xml:"FaceGlobalUrl,omitempty"`
	// example:
	//
	// http://image-demo.img-cn-hangzhou.aliyuncs.com/face-image-example.jpg
	FaceImageUrl *string `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty"`
	// example:
	//
	// false
	FaceMask *bool `json:"FaceMask,omitempty" xml:"FaceMask,omitempty"`
	// example:
	//
	// NORMAL
	FaceQuality *string                                             `json:"FaceQuality,omitempty" xml:"FaceQuality,omitempty"`
	IdCardInfo  *DescribeVerifyResultResponseBodyMaterialIdCardInfo `json:"IdCardInfo,omitempty" xml:"IdCardInfo,omitempty" type:"Struct"`
	IdCardName  *string                                             `json:"IdCardName,omitempty" xml:"IdCardName,omitempty"`
	// example:
	//
	// 02343218901123****
	IdCardNumber *string   `json:"IdCardNumber,omitempty" xml:"IdCardNumber,omitempty"`
	VideoUrls    []*string `json:"VideoUrls,omitempty" xml:"VideoUrls,omitempty" type:"Repeated"`
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
	return dara.Validate(s)
}

type DescribeVerifyResultResponseBodyMaterialIdCardInfo struct {
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
