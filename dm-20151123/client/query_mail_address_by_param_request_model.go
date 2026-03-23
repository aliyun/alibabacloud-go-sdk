// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMailAddressByParamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyWord(v string) *QueryMailAddressByParamRequest
	GetKeyWord() *string
	SetOwnerId(v int64) *QueryMailAddressByParamRequest
	GetOwnerId() *int64
	SetPageNo(v int32) *QueryMailAddressByParamRequest
	GetPageNo() *int32
	SetPageSize(v int32) *QueryMailAddressByParamRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *QueryMailAddressByParamRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryMailAddressByParamRequest
	GetResourceOwnerId() *int64
	SetSendtype(v string) *QueryMailAddressByParamRequest
	GetSendtype() *string
}

type QueryMailAddressByParamRequest struct {
	// The email address. The length is 1 to 60 characters. It supports digits, letters, periods (.), hyphens (-), and at signs (@).
	//
	// example:
	//
	// 账号+@+域名
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The current page number. The default value is 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. The default value is 10.
	//
	// example:
	//
	// 10
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the sender address. Valid values:
	//
	// - batch: batch emails
	//
	// - trigger: triggered emails
	//
	// example:
	//
	// batch
	Sendtype *string `json:"Sendtype,omitempty" xml:"Sendtype,omitempty"`
}

func (s QueryMailAddressByParamRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMailAddressByParamRequest) GoString() string {
	return s.String()
}

func (s *QueryMailAddressByParamRequest) GetKeyWord() *string {
	return s.KeyWord
}

func (s *QueryMailAddressByParamRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryMailAddressByParamRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *QueryMailAddressByParamRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryMailAddressByParamRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryMailAddressByParamRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryMailAddressByParamRequest) GetSendtype() *string {
	return s.Sendtype
}

func (s *QueryMailAddressByParamRequest) SetKeyWord(v string) *QueryMailAddressByParamRequest {
	s.KeyWord = &v
	return s
}

func (s *QueryMailAddressByParamRequest) SetOwnerId(v int64) *QueryMailAddressByParamRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryMailAddressByParamRequest) SetPageNo(v int32) *QueryMailAddressByParamRequest {
	s.PageNo = &v
	return s
}

func (s *QueryMailAddressByParamRequest) SetPageSize(v int32) *QueryMailAddressByParamRequest {
	s.PageSize = &v
	return s
}

func (s *QueryMailAddressByParamRequest) SetResourceOwnerAccount(v string) *QueryMailAddressByParamRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryMailAddressByParamRequest) SetResourceOwnerId(v int64) *QueryMailAddressByParamRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryMailAddressByParamRequest) SetSendtype(v string) *QueryMailAddressByParamRequest {
	s.Sendtype = &v
	return s
}

func (s *QueryMailAddressByParamRequest) Validate() error {
	return dara.Validate(s)
}
