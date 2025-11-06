// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForModifyingDSRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v int32) *SaveSingleTaskForModifyingDSRecordRequest
	GetAlgorithm() *int32
	SetDigest(v string) *SaveSingleTaskForModifyingDSRecordRequest
	GetDigest() *string
	SetDigestType(v int32) *SaveSingleTaskForModifyingDSRecordRequest
	GetDigestType() *int32
	SetDomainName(v string) *SaveSingleTaskForModifyingDSRecordRequest
	GetDomainName() *string
	SetKeyTag(v int32) *SaveSingleTaskForModifyingDSRecordRequest
	GetKeyTag() *int32
	SetLang(v string) *SaveSingleTaskForModifyingDSRecordRequest
	GetLang() *string
	SetUserClientIp(v string) *SaveSingleTaskForModifyingDSRecordRequest
	GetUserClientIp() *string
}

type SaveSingleTaskForModifyingDSRecordRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Algorithm *int32 `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// f58fa917424383934c7b0cf1a90f61d692745680fa06f5ecdbe0924e86de9598
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	DigestType *int32 `json:"DigestType,omitempty" xml:"DigestType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	KeyTag *int32 `json:"KeyTag,omitempty" xml:"KeyTag,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveSingleTaskForModifyingDSRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForModifyingDSRecordRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForModifyingDSRecordRequest) GetAlgorithm() *int32 {
	return s.Algorithm
}

func (s *SaveSingleTaskForModifyingDSRecordRequest) GetDigest() *string {
	return s.Digest
}

func (s *SaveSingleTaskForModifyingDSRecordRequest) GetDigestType() *int32 {
	return s.DigestType
}

func (s *SaveSingleTaskForModifyingDSRecordRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SaveSingleTaskForModifyingDSRecordRequest) GetKeyTag() *int32 {
	return s.KeyTag
}

func (s *SaveSingleTaskForModifyingDSRecordRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveSingleTaskForModifyingDSRecordRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveSingleTaskForModifyingDSRecordRequest) SetAlgorithm(v int32) *SaveSingleTaskForModifyingDSRecordRequest {
	s.Algorithm = &v
	return s
}

func (s *SaveSingleTaskForModifyingDSRecordRequest) SetDigest(v string) *SaveSingleTaskForModifyingDSRecordRequest {
	s.Digest = &v
	return s
}

func (s *SaveSingleTaskForModifyingDSRecordRequest) SetDigestType(v int32) *SaveSingleTaskForModifyingDSRecordRequest {
	s.DigestType = &v
	return s
}

func (s *SaveSingleTaskForModifyingDSRecordRequest) SetDomainName(v string) *SaveSingleTaskForModifyingDSRecordRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForModifyingDSRecordRequest) SetKeyTag(v int32) *SaveSingleTaskForModifyingDSRecordRequest {
	s.KeyTag = &v
	return s
}

func (s *SaveSingleTaskForModifyingDSRecordRequest) SetLang(v string) *SaveSingleTaskForModifyingDSRecordRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForModifyingDSRecordRequest) SetUserClientIp(v string) *SaveSingleTaskForModifyingDSRecordRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveSingleTaskForModifyingDSRecordRequest) Validate() error {
	return dara.Validate(s)
}
