// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMobileOperatorAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *DescribeMobileOperatorAttributeRequest
	GetAuthCode() *string
	SetInputNumber(v string) *DescribeMobileOperatorAttributeRequest
	GetInputNumber() *string
	SetMask(v string) *DescribeMobileOperatorAttributeRequest
	GetMask() *string
	SetOwnerId(v int64) *DescribeMobileOperatorAttributeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeMobileOperatorAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeMobileOperatorAttributeRequest
	GetResourceOwnerId() *int64
}

type DescribeMobileOperatorAttributeRequest struct {
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
	// 示例值
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeMobileOperatorAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMobileOperatorAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeMobileOperatorAttributeRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *DescribeMobileOperatorAttributeRequest) GetInputNumber() *string {
	return s.InputNumber
}

func (s *DescribeMobileOperatorAttributeRequest) GetMask() *string {
	return s.Mask
}

func (s *DescribeMobileOperatorAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeMobileOperatorAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeMobileOperatorAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeMobileOperatorAttributeRequest) SetAuthCode(v string) *DescribeMobileOperatorAttributeRequest {
	s.AuthCode = &v
	return s
}

func (s *DescribeMobileOperatorAttributeRequest) SetInputNumber(v string) *DescribeMobileOperatorAttributeRequest {
	s.InputNumber = &v
	return s
}

func (s *DescribeMobileOperatorAttributeRequest) SetMask(v string) *DescribeMobileOperatorAttributeRequest {
	s.Mask = &v
	return s
}

func (s *DescribeMobileOperatorAttributeRequest) SetOwnerId(v int64) *DescribeMobileOperatorAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMobileOperatorAttributeRequest) SetResourceOwnerAccount(v string) *DescribeMobileOperatorAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeMobileOperatorAttributeRequest) SetResourceOwnerId(v int64) *DescribeMobileOperatorAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeMobileOperatorAttributeRequest) Validate() error {
	return dara.Validate(s)
}
