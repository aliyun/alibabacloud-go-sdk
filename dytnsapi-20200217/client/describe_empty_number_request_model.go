// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEmptyNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *DescribeEmptyNumberRequest
	GetAuthCode() *string
	SetInputNumber(v string) *DescribeEmptyNumberRequest
	GetInputNumber() *string
	SetMask(v string) *DescribeEmptyNumberRequest
	GetMask() *string
	SetOwnerId(v int64) *DescribeEmptyNumberRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeEmptyNumberRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeEmptyNumberRequest
	GetResourceOwnerId() *int64
}

type DescribeEmptyNumberRequest struct {
	// The authorization code.
	//
	// >  On the **My Applications*	- page in the [Cell Phone Number Service console](https://dytns.console.aliyun.com/analysis/apply), you can obtain the authorization code (also known as authorization ID).
	//
	// This parameter is required.
	//
	// example:
	//
	// Dd1r***4id
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The phone number to be queried.
	//
	// >  You can query only one phone number a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 189****1234
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// The encryption method of the phone number. Valid values:
	//
	// 	- **NORMAL**: The phone number is not encrypted.
	//
	// 	- **MD5**
	//
	// 	- **SHA256**
	//
	// This parameter is required.
	//
	// example:
	//
	// NORMAL
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeEmptyNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEmptyNumberRequest) GoString() string {
	return s.String()
}

func (s *DescribeEmptyNumberRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *DescribeEmptyNumberRequest) GetInputNumber() *string {
	return s.InputNumber
}

func (s *DescribeEmptyNumberRequest) GetMask() *string {
	return s.Mask
}

func (s *DescribeEmptyNumberRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeEmptyNumberRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeEmptyNumberRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeEmptyNumberRequest) SetAuthCode(v string) *DescribeEmptyNumberRequest {
	s.AuthCode = &v
	return s
}

func (s *DescribeEmptyNumberRequest) SetInputNumber(v string) *DescribeEmptyNumberRequest {
	s.InputNumber = &v
	return s
}

func (s *DescribeEmptyNumberRequest) SetMask(v string) *DescribeEmptyNumberRequest {
	s.Mask = &v
	return s
}

func (s *DescribeEmptyNumberRequest) SetOwnerId(v int64) *DescribeEmptyNumberRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeEmptyNumberRequest) SetResourceOwnerAccount(v string) *DescribeEmptyNumberRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeEmptyNumberRequest) SetResourceOwnerId(v int64) *DescribeEmptyNumberRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeEmptyNumberRequest) Validate() error {
	return dara.Validate(s)
}
