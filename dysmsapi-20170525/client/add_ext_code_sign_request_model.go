// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddExtCodeSignRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtCode(v string) *AddExtCodeSignRequest
	GetExtCode() *string
	SetOwnerId(v int64) *AddExtCodeSignRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AddExtCodeSignRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddExtCodeSignRequest
	GetResourceOwnerId() *int64
	SetSignName(v string) *AddExtCodeSignRequest
	GetSignName() *string
}

type AddExtCodeSignRequest struct {
	// 扩展码A3
	//
	// This parameter is required.
	//
	// example:
	//
	// 01
	ExtCode              *string `json:"ExtCode,omitempty" xml:"ExtCode,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 签名
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
}

func (s AddExtCodeSignRequest) String() string {
	return dara.Prettify(s)
}

func (s AddExtCodeSignRequest) GoString() string {
	return s.String()
}

func (s *AddExtCodeSignRequest) GetExtCode() *string {
	return s.ExtCode
}

func (s *AddExtCodeSignRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddExtCodeSignRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddExtCodeSignRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddExtCodeSignRequest) GetSignName() *string {
	return s.SignName
}

func (s *AddExtCodeSignRequest) SetExtCode(v string) *AddExtCodeSignRequest {
	s.ExtCode = &v
	return s
}

func (s *AddExtCodeSignRequest) SetOwnerId(v int64) *AddExtCodeSignRequest {
	s.OwnerId = &v
	return s
}

func (s *AddExtCodeSignRequest) SetResourceOwnerAccount(v string) *AddExtCodeSignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddExtCodeSignRequest) SetResourceOwnerId(v int64) *AddExtCodeSignRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddExtCodeSignRequest) SetSignName(v string) *AddExtCodeSignRequest {
	s.SignName = &v
	return s
}

func (s *AddExtCodeSignRequest) Validate() error {
	return dara.Validate(s)
}
