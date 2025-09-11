// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExtCodeSignRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtCode(v string) *DeleteExtCodeSignRequest
	GetExtCode() *string
	SetOwnerId(v int64) *DeleteExtCodeSignRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteExtCodeSignRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteExtCodeSignRequest
	GetResourceOwnerId() *int64
	SetSignName(v string) *DeleteExtCodeSignRequest
	GetSignName() *string
}

type DeleteExtCodeSignRequest struct {
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

func (s DeleteExtCodeSignRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteExtCodeSignRequest) GoString() string {
	return s.String()
}

func (s *DeleteExtCodeSignRequest) GetExtCode() *string {
	return s.ExtCode
}

func (s *DeleteExtCodeSignRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteExtCodeSignRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteExtCodeSignRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteExtCodeSignRequest) GetSignName() *string {
	return s.SignName
}

func (s *DeleteExtCodeSignRequest) SetExtCode(v string) *DeleteExtCodeSignRequest {
	s.ExtCode = &v
	return s
}

func (s *DeleteExtCodeSignRequest) SetOwnerId(v int64) *DeleteExtCodeSignRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteExtCodeSignRequest) SetResourceOwnerAccount(v string) *DeleteExtCodeSignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteExtCodeSignRequest) SetResourceOwnerId(v int64) *DeleteExtCodeSignRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteExtCodeSignRequest) SetSignName(v string) *DeleteExtCodeSignRequest {
	s.SignName = &v
	return s
}

func (s *DeleteExtCodeSignRequest) Validate() error {
	return dara.Validate(s)
}
