// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryExtCodeSignRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtCode(v string) *QueryExtCodeSignRequest
	GetExtCode() *string
	SetOwnerId(v int64) *QueryExtCodeSignRequest
	GetOwnerId() *int64
	SetPageNo(v int64) *QueryExtCodeSignRequest
	GetPageNo() *int64
	SetPageSize(v int64) *QueryExtCodeSignRequest
	GetPageSize() *int64
	SetResourceOwnerAccount(v string) *QueryExtCodeSignRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryExtCodeSignRequest
	GetResourceOwnerId() *int64
	SetSignName(v string) *QueryExtCodeSignRequest
	GetSignName() *string
}

type QueryExtCodeSignRequest struct {
	// 扩展码A3
	//
	// example:
	//
	// 01
	ExtCode *string `json:"ExtCode,omitempty" xml:"ExtCode,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 20
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 签名
	//
	// example:
	//
	// 示例值示例值
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
}

func (s QueryExtCodeSignRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryExtCodeSignRequest) GoString() string {
	return s.String()
}

func (s *QueryExtCodeSignRequest) GetExtCode() *string {
	return s.ExtCode
}

func (s *QueryExtCodeSignRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryExtCodeSignRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *QueryExtCodeSignRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryExtCodeSignRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryExtCodeSignRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryExtCodeSignRequest) GetSignName() *string {
	return s.SignName
}

func (s *QueryExtCodeSignRequest) SetExtCode(v string) *QueryExtCodeSignRequest {
	s.ExtCode = &v
	return s
}

func (s *QueryExtCodeSignRequest) SetOwnerId(v int64) *QueryExtCodeSignRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryExtCodeSignRequest) SetPageNo(v int64) *QueryExtCodeSignRequest {
	s.PageNo = &v
	return s
}

func (s *QueryExtCodeSignRequest) SetPageSize(v int64) *QueryExtCodeSignRequest {
	s.PageSize = &v
	return s
}

func (s *QueryExtCodeSignRequest) SetResourceOwnerAccount(v string) *QueryExtCodeSignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryExtCodeSignRequest) SetResourceOwnerId(v int64) *QueryExtCodeSignRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryExtCodeSignRequest) SetSignName(v string) *QueryExtCodeSignRequest {
	s.SignName = &v
	return s
}

func (s *QueryExtCodeSignRequest) Validate() error {
	return dara.Validate(s)
}
