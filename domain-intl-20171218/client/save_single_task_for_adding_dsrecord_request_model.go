// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForAddingDSRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v int32) *SaveSingleTaskForAddingDSRecordRequest
	GetAlgorithm() *int32
	SetDigest(v string) *SaveSingleTaskForAddingDSRecordRequest
	GetDigest() *string
	SetDigestType(v int32) *SaveSingleTaskForAddingDSRecordRequest
	GetDigestType() *int32
	SetDomainName(v string) *SaveSingleTaskForAddingDSRecordRequest
	GetDomainName() *string
	SetKeyTag(v int32) *SaveSingleTaskForAddingDSRecordRequest
	GetKeyTag() *int32
	SetLang(v string) *SaveSingleTaskForAddingDSRecordRequest
	GetLang() *string
	SetUserClientIp(v string) *SaveSingleTaskForAddingDSRecordRequest
	GetUserClientIp() *string
}

type SaveSingleTaskForAddingDSRecordRequest struct {
	// This parameter is required.
	Algorithm *int32 `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// This parameter is required.
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// This parameter is required.
	DigestType *int32 `json:"DigestType,omitempty" xml:"DigestType,omitempty"`
	// This parameter is required.
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// This parameter is required.
	KeyTag       *int32  `json:"KeyTag,omitempty" xml:"KeyTag,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveSingleTaskForAddingDSRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForAddingDSRecordRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForAddingDSRecordRequest) GetAlgorithm() *int32 {
	return s.Algorithm
}

func (s *SaveSingleTaskForAddingDSRecordRequest) GetDigest() *string {
	return s.Digest
}

func (s *SaveSingleTaskForAddingDSRecordRequest) GetDigestType() *int32 {
	return s.DigestType
}

func (s *SaveSingleTaskForAddingDSRecordRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SaveSingleTaskForAddingDSRecordRequest) GetKeyTag() *int32 {
	return s.KeyTag
}

func (s *SaveSingleTaskForAddingDSRecordRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveSingleTaskForAddingDSRecordRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveSingleTaskForAddingDSRecordRequest) SetAlgorithm(v int32) *SaveSingleTaskForAddingDSRecordRequest {
	s.Algorithm = &v
	return s
}

func (s *SaveSingleTaskForAddingDSRecordRequest) SetDigest(v string) *SaveSingleTaskForAddingDSRecordRequest {
	s.Digest = &v
	return s
}

func (s *SaveSingleTaskForAddingDSRecordRequest) SetDigestType(v int32) *SaveSingleTaskForAddingDSRecordRequest {
	s.DigestType = &v
	return s
}

func (s *SaveSingleTaskForAddingDSRecordRequest) SetDomainName(v string) *SaveSingleTaskForAddingDSRecordRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForAddingDSRecordRequest) SetKeyTag(v int32) *SaveSingleTaskForAddingDSRecordRequest {
	s.KeyTag = &v
	return s
}

func (s *SaveSingleTaskForAddingDSRecordRequest) SetLang(v string) *SaveSingleTaskForAddingDSRecordRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForAddingDSRecordRequest) SetUserClientIp(v string) *SaveSingleTaskForAddingDSRecordRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveSingleTaskForAddingDSRecordRequest) Validate() error {
	return dara.Validate(s)
}
