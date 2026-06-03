// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGateVerifyRecordListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthenticationType(v int32) *QueryGateVerifyRecordListRequest
	GetAuthenticationType() *int32
	SetOsType(v string) *QueryGateVerifyRecordListRequest
	GetOsType() *string
	SetOwnerId(v int64) *QueryGateVerifyRecordListRequest
	GetOwnerId() *int64
	SetPhoneNum(v string) *QueryGateVerifyRecordListRequest
	GetPhoneNum() *string
	SetProdCode(v string) *QueryGateVerifyRecordListRequest
	GetProdCode() *string
	SetQueryDate(v string) *QueryGateVerifyRecordListRequest
	GetQueryDate() *string
	SetResourceOwnerAccount(v string) *QueryGateVerifyRecordListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryGateVerifyRecordListRequest
	GetResourceOwnerId() *int64
}

type QueryGateVerifyRecordListRequest struct {
	// This parameter is required.
	AuthenticationType *int32 `json:"AuthenticationType,omitempty" xml:"AuthenticationType,omitempty"`
	// This parameter is required.
	OsType  *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	PhoneNum *string `json:"PhoneNum,omitempty" xml:"PhoneNum,omitempty"`
	ProdCode *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	// This parameter is required.
	QueryDate            *string `json:"QueryDate,omitempty" xml:"QueryDate,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryGateVerifyRecordListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryGateVerifyRecordListRequest) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyRecordListRequest) GetAuthenticationType() *int32 {
	return s.AuthenticationType
}

func (s *QueryGateVerifyRecordListRequest) GetOsType() *string {
	return s.OsType
}

func (s *QueryGateVerifyRecordListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryGateVerifyRecordListRequest) GetPhoneNum() *string {
	return s.PhoneNum
}

func (s *QueryGateVerifyRecordListRequest) GetProdCode() *string {
	return s.ProdCode
}

func (s *QueryGateVerifyRecordListRequest) GetQueryDate() *string {
	return s.QueryDate
}

func (s *QueryGateVerifyRecordListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryGateVerifyRecordListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryGateVerifyRecordListRequest) SetAuthenticationType(v int32) *QueryGateVerifyRecordListRequest {
	s.AuthenticationType = &v
	return s
}

func (s *QueryGateVerifyRecordListRequest) SetOsType(v string) *QueryGateVerifyRecordListRequest {
	s.OsType = &v
	return s
}

func (s *QueryGateVerifyRecordListRequest) SetOwnerId(v int64) *QueryGateVerifyRecordListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryGateVerifyRecordListRequest) SetPhoneNum(v string) *QueryGateVerifyRecordListRequest {
	s.PhoneNum = &v
	return s
}

func (s *QueryGateVerifyRecordListRequest) SetProdCode(v string) *QueryGateVerifyRecordListRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryGateVerifyRecordListRequest) SetQueryDate(v string) *QueryGateVerifyRecordListRequest {
	s.QueryDate = &v
	return s
}

func (s *QueryGateVerifyRecordListRequest) SetResourceOwnerAccount(v string) *QueryGateVerifyRecordListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryGateVerifyRecordListRequest) SetResourceOwnerId(v int64) *QueryGateVerifyRecordListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryGateVerifyRecordListRequest) Validate() error {
	return dara.Validate(s)
}
