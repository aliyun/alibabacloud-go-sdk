// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExtCodeSignRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExistExtCode(v string) *UpdateExtCodeSignRequest
	GetExistExtCode() *string
	SetNewExtCode(v string) *UpdateExtCodeSignRequest
	GetNewExtCode() *string
	SetOwnerId(v int64) *UpdateExtCodeSignRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpdateExtCodeSignRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateExtCodeSignRequest
	GetResourceOwnerId() *int64
	SetSignName(v string) *UpdateExtCodeSignRequest
	GetSignName() *string
}

type UpdateExtCodeSignRequest struct {
	// 要修改的扩展码A3
	//
	// This parameter is required.
	//
	// example:
	//
	// 01
	ExistExtCode *string `json:"ExistExtCode,omitempty" xml:"ExistExtCode,omitempty"`
	// 修改后的扩展码A3
	//
	// This parameter is required.
	//
	// example:
	//
	// 02
	NewExtCode           *string `json:"NewExtCode,omitempty" xml:"NewExtCode,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 签名
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
}

func (s UpdateExtCodeSignRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateExtCodeSignRequest) GoString() string {
	return s.String()
}

func (s *UpdateExtCodeSignRequest) GetExistExtCode() *string {
	return s.ExistExtCode
}

func (s *UpdateExtCodeSignRequest) GetNewExtCode() *string {
	return s.NewExtCode
}

func (s *UpdateExtCodeSignRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateExtCodeSignRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateExtCodeSignRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateExtCodeSignRequest) GetSignName() *string {
	return s.SignName
}

func (s *UpdateExtCodeSignRequest) SetExistExtCode(v string) *UpdateExtCodeSignRequest {
	s.ExistExtCode = &v
	return s
}

func (s *UpdateExtCodeSignRequest) SetNewExtCode(v string) *UpdateExtCodeSignRequest {
	s.NewExtCode = &v
	return s
}

func (s *UpdateExtCodeSignRequest) SetOwnerId(v int64) *UpdateExtCodeSignRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateExtCodeSignRequest) SetResourceOwnerAccount(v string) *UpdateExtCodeSignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateExtCodeSignRequest) SetResourceOwnerId(v int64) *UpdateExtCodeSignRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateExtCodeSignRequest) SetSignName(v string) *UpdateExtCodeSignRequest {
	s.SignName = &v
	return s
}

func (s *UpdateExtCodeSignRequest) Validate() error {
	return dara.Validate(s)
}
