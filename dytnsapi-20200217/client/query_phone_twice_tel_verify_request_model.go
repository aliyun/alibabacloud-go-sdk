// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPhoneTwiceTelVerifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *QueryPhoneTwiceTelVerifyRequest
	GetAuthCode() *string
	SetInputNumber(v string) *QueryPhoneTwiceTelVerifyRequest
	GetInputNumber() *string
	SetMask(v string) *QueryPhoneTwiceTelVerifyRequest
	GetMask() *string
	SetOwnerId(v int64) *QueryPhoneTwiceTelVerifyRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryPhoneTwiceTelVerifyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryPhoneTwiceTelVerifyRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *QueryPhoneTwiceTelVerifyRequest
	GetStartTime() *string
}

type QueryPhoneTwiceTelVerifyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryPhoneTwiceTelVerifyRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPhoneTwiceTelVerifyRequest) GoString() string {
	return s.String()
}

func (s *QueryPhoneTwiceTelVerifyRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *QueryPhoneTwiceTelVerifyRequest) GetInputNumber() *string {
	return s.InputNumber
}

func (s *QueryPhoneTwiceTelVerifyRequest) GetMask() *string {
	return s.Mask
}

func (s *QueryPhoneTwiceTelVerifyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryPhoneTwiceTelVerifyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryPhoneTwiceTelVerifyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryPhoneTwiceTelVerifyRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryPhoneTwiceTelVerifyRequest) SetAuthCode(v string) *QueryPhoneTwiceTelVerifyRequest {
	s.AuthCode = &v
	return s
}

func (s *QueryPhoneTwiceTelVerifyRequest) SetInputNumber(v string) *QueryPhoneTwiceTelVerifyRequest {
	s.InputNumber = &v
	return s
}

func (s *QueryPhoneTwiceTelVerifyRequest) SetMask(v string) *QueryPhoneTwiceTelVerifyRequest {
	s.Mask = &v
	return s
}

func (s *QueryPhoneTwiceTelVerifyRequest) SetOwnerId(v int64) *QueryPhoneTwiceTelVerifyRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryPhoneTwiceTelVerifyRequest) SetResourceOwnerAccount(v string) *QueryPhoneTwiceTelVerifyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryPhoneTwiceTelVerifyRequest) SetResourceOwnerId(v int64) *QueryPhoneTwiceTelVerifyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryPhoneTwiceTelVerifyRequest) SetStartTime(v string) *QueryPhoneTwiceTelVerifyRequest {
	s.StartTime = &v
	return s
}

func (s *QueryPhoneTwiceTelVerifyRequest) Validate() error {
	return dara.Validate(s)
}
