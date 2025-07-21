// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainByParamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyWord(v string) *QueryDomainByParamRequest
	GetKeyWord() *string
	SetOwnerId(v int64) *QueryDomainByParamRequest
	GetOwnerId() *int64
	SetPageNo(v int32) *QueryDomainByParamRequest
	GetPageNo() *int32
	SetPageSize(v int32) *QueryDomainByParamRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *QueryDomainByParamRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryDomainByParamRequest
	GetResourceOwnerId() *int64
	SetStatus(v int32) *QueryDomainByParamRequest
	GetStatus() *int32
}

type QueryDomainByParamRequest struct {
	// Domain name, length 1-50, can include numbers, uppercase and lowercase letters, ., -.
	//
	// example:
	//
	// example.com
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Current page number. Default: 1
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// Number of items per page, default: 10
	//
	// example:
	//
	// 10
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// - 0 indicates normal
	//
	// - 1 indicates abnormal
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryDomainByParamRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainByParamRequest) GoString() string {
	return s.String()
}

func (s *QueryDomainByParamRequest) GetKeyWord() *string {
	return s.KeyWord
}

func (s *QueryDomainByParamRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryDomainByParamRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *QueryDomainByParamRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryDomainByParamRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryDomainByParamRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryDomainByParamRequest) GetStatus() *int32 {
	return s.Status
}

func (s *QueryDomainByParamRequest) SetKeyWord(v string) *QueryDomainByParamRequest {
	s.KeyWord = &v
	return s
}

func (s *QueryDomainByParamRequest) SetOwnerId(v int64) *QueryDomainByParamRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryDomainByParamRequest) SetPageNo(v int32) *QueryDomainByParamRequest {
	s.PageNo = &v
	return s
}

func (s *QueryDomainByParamRequest) SetPageSize(v int32) *QueryDomainByParamRequest {
	s.PageSize = &v
	return s
}

func (s *QueryDomainByParamRequest) SetResourceOwnerAccount(v string) *QueryDomainByParamRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryDomainByParamRequest) SetResourceOwnerId(v int64) *QueryDomainByParamRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryDomainByParamRequest) SetStatus(v int32) *QueryDomainByParamRequest {
	s.Status = &v
	return s
}

func (s *QueryDomainByParamRequest) Validate() error {
	return dara.Validate(s)
}
