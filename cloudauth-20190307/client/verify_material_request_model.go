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
	// A unique ID that identifies a single authentication task, not exceeding 64 characters. For a single authentication task, the system supports unlimited submissions until the final authentication is passed and the task is completed.
	//
	// > Different BizIds are required for different authentication tasks.
	//
	// This parameter is required.
	//
	// example:
	//
	// 39ecf51e-2f81-4dc5-90ee-ff86125b****
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// Identifier for the business scenario using the real-person authentication service. Please refer to [Business Setup](https://help.aliyun.com/document_detail/127885.html) and complete the creation in the console first.
	//
	// This parameter is required.
	//
	// example:
	//
	// RPMinTest
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// HTTP or HTTPS link to the frontal face image.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://image-demo.img-cn-hangzhou.aliyuncs.com/example.jpg
	FaceImageUrl *string `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty"`
	// HTTP or HTTPS link to the national emblem side of the ID card.
	//
	// example:
	//
	// http://image-demo.img-cn-hangzhou.aliyuncs.com/example3.jpg
	IdCardBackImageUrl *string `json:"IdCardBackImageUrl,omitempty" xml:"IdCardBackImageUrl,omitempty"`
	// HTTP or HTTPS link to the portrait side of the ID card image.
	//
	// example:
	//
	// http://image-demo.img-cn-hangzhou.aliyuncs.com/example2.jpg
	IdCardFrontImageUrl *string `json:"IdCardFrontImageUrl,omitempty" xml:"IdCardFrontImageUrl,omitempty"`
	// ID number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 02343218901123****
	IdCardNumber *string `json:"IdCardNumber,omitempty" xml:"IdCardNumber,omitempty"`
	// Name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 张三
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the end user, such as the account ID of the end user.
	//
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
