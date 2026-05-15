// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecordDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DescribeRecordDataRequest
	GetAccountId() *string
	SetAccountType(v string) *DescribeRecordDataRequest
	GetAccountType() *string
	SetAcid(v string) *DescribeRecordDataRequest
	GetAcid() *string
	SetOwnerId(v int64) *DescribeRecordDataRequest
	GetOwnerId() *int64
	SetProdCode(v string) *DescribeRecordDataRequest
	GetProdCode() *string
	SetResourceOwnerAccount(v string) *DescribeRecordDataRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeRecordDataRequest
	GetResourceOwnerId() *int64
	SetSecLevel(v int32) *DescribeRecordDataRequest
	GetSecLevel() *int32
}

type DescribeRecordDataRequest struct {
	// example:
	//
	// 2235****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// example:
	//
	// BUC_TYPE
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// example:
	//
	// 1004849****
	Acid    *string `json:"Acid,omitempty" xml:"Acid,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// aiccs
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 2
	SecLevel *int32 `json:"SecLevel,omitempty" xml:"SecLevel,omitempty"`
}

func (s DescribeRecordDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecordDataRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DescribeRecordDataRequest) GetAccountType() *string {
	return s.AccountType
}

func (s *DescribeRecordDataRequest) GetAcid() *string {
	return s.Acid
}

func (s *DescribeRecordDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRecordDataRequest) GetProdCode() *string {
	return s.ProdCode
}

func (s *DescribeRecordDataRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeRecordDataRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeRecordDataRequest) GetSecLevel() *int32 {
	return s.SecLevel
}

func (s *DescribeRecordDataRequest) SetAccountId(v string) *DescribeRecordDataRequest {
	s.AccountId = &v
	return s
}

func (s *DescribeRecordDataRequest) SetAccountType(v string) *DescribeRecordDataRequest {
	s.AccountType = &v
	return s
}

func (s *DescribeRecordDataRequest) SetAcid(v string) *DescribeRecordDataRequest {
	s.Acid = &v
	return s
}

func (s *DescribeRecordDataRequest) SetOwnerId(v int64) *DescribeRecordDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRecordDataRequest) SetProdCode(v string) *DescribeRecordDataRequest {
	s.ProdCode = &v
	return s
}

func (s *DescribeRecordDataRequest) SetResourceOwnerAccount(v string) *DescribeRecordDataRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRecordDataRequest) SetResourceOwnerId(v int64) *DescribeRecordDataRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRecordDataRequest) SetSecLevel(v int32) *DescribeRecordDataRequest {
	s.SecLevel = &v
	return s
}

func (s *DescribeRecordDataRequest) Validate() error {
	return dara.Validate(s)
}
