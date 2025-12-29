// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhoneNumberRiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *DescribePhoneNumberRiskRequest
	GetAuthCode() *string
	SetInputNumber(v string) *DescribePhoneNumberRiskRequest
	GetInputNumber() *string
	SetMask(v string) *DescribePhoneNumberRiskRequest
	GetMask() *string
	SetOwnerId(v int64) *DescribePhoneNumberRiskRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribePhoneNumberRiskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribePhoneNumberRiskRequest
	GetResourceOwnerId() *int64
}

type DescribePhoneNumberRiskRequest struct {
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
	// 示例值示例值
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
}

func (s DescribePhoneNumberRiskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberRiskRequest) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberRiskRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *DescribePhoneNumberRiskRequest) GetInputNumber() *string {
	return s.InputNumber
}

func (s *DescribePhoneNumberRiskRequest) GetMask() *string {
	return s.Mask
}

func (s *DescribePhoneNumberRiskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePhoneNumberRiskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribePhoneNumberRiskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribePhoneNumberRiskRequest) SetAuthCode(v string) *DescribePhoneNumberRiskRequest {
	s.AuthCode = &v
	return s
}

func (s *DescribePhoneNumberRiskRequest) SetInputNumber(v string) *DescribePhoneNumberRiskRequest {
	s.InputNumber = &v
	return s
}

func (s *DescribePhoneNumberRiskRequest) SetMask(v string) *DescribePhoneNumberRiskRequest {
	s.Mask = &v
	return s
}

func (s *DescribePhoneNumberRiskRequest) SetOwnerId(v int64) *DescribePhoneNumberRiskRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePhoneNumberRiskRequest) SetResourceOwnerAccount(v string) *DescribePhoneNumberRiskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePhoneNumberRiskRequest) SetResourceOwnerId(v int64) *DescribePhoneNumberRiskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePhoneNumberRiskRequest) Validate() error {
	return dara.Validate(s)
}
