// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v int32) *ListCertificateRequest
	GetBusinessType() *int32
	SetEndTime(v int64) *ListCertificateRequest
	GetEndTime() *int64
	SetOwnerId(v int64) *ListCertificateRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *ListCertificateRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCertificateRequest
	GetPageSize() *int32
	SetPhoneNo(v string) *ListCertificateRequest
	GetPhoneNo() *string
	SetResourceOwnerAccount(v string) *ListCertificateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListCertificateRequest
	GetResourceOwnerId() *int64
	SetStartTime(v int64) *ListCertificateRequest
	GetStartTime() *int64
}

type ListCertificateRequest struct {
	BusinessType *int32 `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	EndTime      *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId      *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PhoneNo              *string `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCertificateRequest) GoString() string {
	return s.String()
}

func (s *ListCertificateRequest) GetBusinessType() *int32 {
	return s.BusinessType
}

func (s *ListCertificateRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListCertificateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListCertificateRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCertificateRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCertificateRequest) GetPhoneNo() *string {
	return s.PhoneNo
}

func (s *ListCertificateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListCertificateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListCertificateRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListCertificateRequest) SetBusinessType(v int32) *ListCertificateRequest {
	s.BusinessType = &v
	return s
}

func (s *ListCertificateRequest) SetEndTime(v int64) *ListCertificateRequest {
	s.EndTime = &v
	return s
}

func (s *ListCertificateRequest) SetOwnerId(v int64) *ListCertificateRequest {
	s.OwnerId = &v
	return s
}

func (s *ListCertificateRequest) SetPageNumber(v int32) *ListCertificateRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCertificateRequest) SetPageSize(v int32) *ListCertificateRequest {
	s.PageSize = &v
	return s
}

func (s *ListCertificateRequest) SetPhoneNo(v string) *ListCertificateRequest {
	s.PhoneNo = &v
	return s
}

func (s *ListCertificateRequest) SetResourceOwnerAccount(v string) *ListCertificateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListCertificateRequest) SetResourceOwnerId(v int64) *ListCertificateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListCertificateRequest) SetStartTime(v int64) *ListCertificateRequest {
	s.StartTime = &v
	return s
}

func (s *ListCertificateRequest) Validate() error {
	return dara.Validate(s)
}
