// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifyTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *DescribeVerifyTokenRequest
	GetBizId() *string
	SetBizType(v string) *DescribeVerifyTokenRequest
	GetBizType() *string
	SetCallbackSeed(v string) *DescribeVerifyTokenRequest
	GetCallbackSeed() *string
	SetCallbackUrl(v string) *DescribeVerifyTokenRequest
	GetCallbackUrl() *string
	SetFaceRetainedImageUrl(v string) *DescribeVerifyTokenRequest
	GetFaceRetainedImageUrl() *string
	SetFailedRedirectUrl(v string) *DescribeVerifyTokenRequest
	GetFailedRedirectUrl() *string
	SetIdCardBackImageUrl(v string) *DescribeVerifyTokenRequest
	GetIdCardBackImageUrl() *string
	SetIdCardFrontImageUrl(v string) *DescribeVerifyTokenRequest
	GetIdCardFrontImageUrl() *string
	SetIdCardNumber(v string) *DescribeVerifyTokenRequest
	GetIdCardNumber() *string
	SetName(v string) *DescribeVerifyTokenRequest
	GetName() *string
	SetPassedRedirectUrl(v string) *DescribeVerifyTokenRequest
	GetPassedRedirectUrl() *string
	SetUserId(v string) *DescribeVerifyTokenRequest
	GetUserId() *string
	SetUserIp(v string) *DescribeVerifyTokenRequest
	GetUserIp() *string
	SetUserPhoneNumber(v string) *DescribeVerifyTokenRequest
	GetUserPhoneNumber() *string
	SetUserRegistTime(v int64) *DescribeVerifyTokenRequest
	GetUserRegistTime() *int64
}

type DescribeVerifyTokenRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 39ecf51e-2f81-4dc5-90ee-ff86125be683
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// RPBasicTest
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// -
	CallbackSeed *string `json:"CallbackSeed,omitempty" xml:"CallbackSeed,omitempty"`
	// example:
	//
	// -
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// example:
	//
	// http%3A%2F%2Fjiangsu.china.com.cn%2Fuploadfile%2F2015%2F0114%2F1421221304095989.jpg
	FaceRetainedImageUrl *string `json:"FaceRetainedImageUrl,omitempty" xml:"FaceRetainedImageUrl,omitempty"`
	// example:
	//
	// -
	FailedRedirectUrl *string `json:"FailedRedirectUrl,omitempty" xml:"FailedRedirectUrl,omitempty"`
	// example:
	//
	// http%3A%2F%2Fjiangsu.china.com.cn%2Fuploadfile%2F2015%2F0114%2F1421221304095989.jpg
	IdCardBackImageUrl *string `json:"IdCardBackImageUrl,omitempty" xml:"IdCardBackImageUrl,omitempty"`
	// example:
	//
	// http%3A%2F%2Fjiangsu.china.com.cn%2Fuploadfile%2F2015%2F0114%2F1421221304095989.jpg
	IdCardFrontImageUrl *string `json:"IdCardFrontImageUrl,omitempty" xml:"IdCardFrontImageUrl,omitempty"`
	// example:
	//
	// 330100xxxxxxxxxxxx
	IdCardNumber *string `json:"IdCardNumber,omitempty" xml:"IdCardNumber,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// -
	PassedRedirectUrl *string `json:"PassedRedirectUrl,omitempty" xml:"PassedRedirectUrl,omitempty"`
	// example:
	//
	// user111
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 192.168.***.***
	UserIp *string `json:"UserIp,omitempty" xml:"UserIp,omitempty"`
	// example:
	//
	// 187********
	UserPhoneNumber *string `json:"UserPhoneNumber,omitempty" xml:"UserPhoneNumber,omitempty"`
	// example:
	//
	// 1577808000000
	UserRegistTime *int64 `json:"UserRegistTime,omitempty" xml:"UserRegistTime,omitempty"`
}

func (s DescribeVerifyTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyTokenRequest) GoString() string {
	return s.String()
}

func (s *DescribeVerifyTokenRequest) GetBizId() *string {
	return s.BizId
}

func (s *DescribeVerifyTokenRequest) GetBizType() *string {
	return s.BizType
}

func (s *DescribeVerifyTokenRequest) GetCallbackSeed() *string {
	return s.CallbackSeed
}

func (s *DescribeVerifyTokenRequest) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *DescribeVerifyTokenRequest) GetFaceRetainedImageUrl() *string {
	return s.FaceRetainedImageUrl
}

func (s *DescribeVerifyTokenRequest) GetFailedRedirectUrl() *string {
	return s.FailedRedirectUrl
}

func (s *DescribeVerifyTokenRequest) GetIdCardBackImageUrl() *string {
	return s.IdCardBackImageUrl
}

func (s *DescribeVerifyTokenRequest) GetIdCardFrontImageUrl() *string {
	return s.IdCardFrontImageUrl
}

func (s *DescribeVerifyTokenRequest) GetIdCardNumber() *string {
	return s.IdCardNumber
}

func (s *DescribeVerifyTokenRequest) GetName() *string {
	return s.Name
}

func (s *DescribeVerifyTokenRequest) GetPassedRedirectUrl() *string {
	return s.PassedRedirectUrl
}

func (s *DescribeVerifyTokenRequest) GetUserId() *string {
	return s.UserId
}

func (s *DescribeVerifyTokenRequest) GetUserIp() *string {
	return s.UserIp
}

func (s *DescribeVerifyTokenRequest) GetUserPhoneNumber() *string {
	return s.UserPhoneNumber
}

func (s *DescribeVerifyTokenRequest) GetUserRegistTime() *int64 {
	return s.UserRegistTime
}

func (s *DescribeVerifyTokenRequest) SetBizId(v string) *DescribeVerifyTokenRequest {
	s.BizId = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetBizType(v string) *DescribeVerifyTokenRequest {
	s.BizType = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetCallbackSeed(v string) *DescribeVerifyTokenRequest {
	s.CallbackSeed = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetCallbackUrl(v string) *DescribeVerifyTokenRequest {
	s.CallbackUrl = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetFaceRetainedImageUrl(v string) *DescribeVerifyTokenRequest {
	s.FaceRetainedImageUrl = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetFailedRedirectUrl(v string) *DescribeVerifyTokenRequest {
	s.FailedRedirectUrl = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetIdCardBackImageUrl(v string) *DescribeVerifyTokenRequest {
	s.IdCardBackImageUrl = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetIdCardFrontImageUrl(v string) *DescribeVerifyTokenRequest {
	s.IdCardFrontImageUrl = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetIdCardNumber(v string) *DescribeVerifyTokenRequest {
	s.IdCardNumber = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetName(v string) *DescribeVerifyTokenRequest {
	s.Name = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetPassedRedirectUrl(v string) *DescribeVerifyTokenRequest {
	s.PassedRedirectUrl = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetUserId(v string) *DescribeVerifyTokenRequest {
	s.UserId = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetUserIp(v string) *DescribeVerifyTokenRequest {
	s.UserIp = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetUserPhoneNumber(v string) *DescribeVerifyTokenRequest {
	s.UserPhoneNumber = &v
	return s
}

func (s *DescribeVerifyTokenRequest) SetUserRegistTime(v int64) *DescribeVerifyTokenRequest {
	s.UserRegistTime = &v
	return s
}

func (s *DescribeVerifyTokenRequest) Validate() error {
	return dara.Validate(s)
}
