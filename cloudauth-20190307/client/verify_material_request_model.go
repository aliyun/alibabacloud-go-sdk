// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyMaterialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *VerifyMaterialRequest
	GetBizId() *string
	SetBizType(v string) *VerifyMaterialRequest
	GetBizType() *string
	SetFaceImageUrl(v string) *VerifyMaterialRequest
	GetFaceImageUrl() *string
	SetIdCardBackImageUrl(v string) *VerifyMaterialRequest
	GetIdCardBackImageUrl() *string
	SetIdCardFrontImageUrl(v string) *VerifyMaterialRequest
	GetIdCardFrontImageUrl() *string
	SetIdCardNumber(v string) *VerifyMaterialRequest
	GetIdCardNumber() *string
	SetName(v string) *VerifyMaterialRequest
	GetName() *string
	SetUserId(v string) *VerifyMaterialRequest
	GetUserId() *string
}

type VerifyMaterialRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 39ecf51e-2f81-4dc5-90ee-ff86125b****
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// RPMinTest
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://image-demo.img-cn-hangzhou.aliyuncs.com/example.jpg
	FaceImageUrl *string `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty"`
	// example:
	//
	// http://image-demo.img-cn-hangzhou.aliyuncs.com/example3.jpg
	IdCardBackImageUrl *string `json:"IdCardBackImageUrl,omitempty" xml:"IdCardBackImageUrl,omitempty"`
	// example:
	//
	// http://image-demo.img-cn-hangzhou.aliyuncs.com/example2.jpg
	IdCardFrontImageUrl *string `json:"IdCardFrontImageUrl,omitempty" xml:"IdCardFrontImageUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 02343218901123****
	IdCardNumber *string `json:"IdCardNumber,omitempty" xml:"IdCardNumber,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 54sdj
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s VerifyMaterialRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyMaterialRequest) GoString() string {
	return s.String()
}

func (s *VerifyMaterialRequest) GetBizId() *string {
	return s.BizId
}

func (s *VerifyMaterialRequest) GetBizType() *string {
	return s.BizType
}

func (s *VerifyMaterialRequest) GetFaceImageUrl() *string {
	return s.FaceImageUrl
}

func (s *VerifyMaterialRequest) GetIdCardBackImageUrl() *string {
	return s.IdCardBackImageUrl
}

func (s *VerifyMaterialRequest) GetIdCardFrontImageUrl() *string {
	return s.IdCardFrontImageUrl
}

func (s *VerifyMaterialRequest) GetIdCardNumber() *string {
	return s.IdCardNumber
}

func (s *VerifyMaterialRequest) GetName() *string {
	return s.Name
}

func (s *VerifyMaterialRequest) GetUserId() *string {
	return s.UserId
}

func (s *VerifyMaterialRequest) SetBizId(v string) *VerifyMaterialRequest {
	s.BizId = &v
	return s
}

func (s *VerifyMaterialRequest) SetBizType(v string) *VerifyMaterialRequest {
	s.BizType = &v
	return s
}

func (s *VerifyMaterialRequest) SetFaceImageUrl(v string) *VerifyMaterialRequest {
	s.FaceImageUrl = &v
	return s
}

func (s *VerifyMaterialRequest) SetIdCardBackImageUrl(v string) *VerifyMaterialRequest {
	s.IdCardBackImageUrl = &v
	return s
}

func (s *VerifyMaterialRequest) SetIdCardFrontImageUrl(v string) *VerifyMaterialRequest {
	s.IdCardFrontImageUrl = &v
	return s
}

func (s *VerifyMaterialRequest) SetIdCardNumber(v string) *VerifyMaterialRequest {
	s.IdCardNumber = &v
	return s
}

func (s *VerifyMaterialRequest) SetName(v string) *VerifyMaterialRequest {
	s.Name = &v
	return s
}

func (s *VerifyMaterialRequest) SetUserId(v string) *VerifyMaterialRequest {
	s.UserId = &v
	return s
}

func (s *VerifyMaterialRequest) Validate() error {
	return dara.Validate(s)
}
