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
	// The encryption algorithm number. For more information, see [Domain Name System Security (DNSSEC) Algorithm Numbers](https://www.iana.org/assignments/dns-sec-alg-numbers/dns-sec-alg-numbers.xhtml). Valid values:
	//
	// - **1**: RSA/MD5.
	//
	// - **2**: Diffie-Hellman.
	//
	// - **3**: DSA/SHA-1.
	//
	// - **5**: RSA/SHA-1.
	//
	// - **6**: DSA-NSEC3-SHA1.
	//
	// - **7**: RSASHA1-NSEC3-SHA1.
	//
	// - **8**: RSA/SHA-256.
	//
	// - **10**: RSA/SHA-512.
	//
	// - **12**: GOST R 34.10-2001.
	//
	// - **13**: ECDSA Curve P-256 with SHA-256.
	//
	// - **14**: ECDSA Curve P-384 with SHA-384.
	//
	// - **15**: Ed25519 and Ed448.
	//
	// - **252**: Reserved for Indirect Keys.
	//
	// - **253**: private algorithm.
	//
	// - **254**: private algorithm OID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Algorithm *int32 `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// Summary.
	//
	// This parameter is required.
	//
	// example:
	//
	// f58fa917424383934c7b0cf1a90f61d692745680fa06f5ecdbe0924e86de9598
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// Summary algorithm type. For more information, see [Delegation Signer (DS) Resource Record (RR) Type Digest Algorithms](https://www.iana.org/assignments/ds-rr-types/ds-rr-types.xhtml). Valid values:
	//
	// - **1**: SHA-1;
	//
	// - **2**: SHA-256;
	//
	// - **3**: GOST R 34.11-94;
	//
	// - **4**: SHA-384.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	DigestType *int32 `json:"DigestType,omitempty" xml:"DigestType,omitempty"`
	// Domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Key tag used to identify DNSSEC records. It is an integer less than 65536.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	KeyTag *int32 `json:"KeyTag,omitempty" xml:"KeyTag,omitempty"`
	// Language of error messages returned by the API. Valid values:
	//
	// - **zh**: Chinese;
	//
	// - **en**: English.
	//
	// Default value: **en**.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// User IP address, which can be set to **127.0.0.1**.
	//
	// example:
	//
	// 127.0.0.1
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
