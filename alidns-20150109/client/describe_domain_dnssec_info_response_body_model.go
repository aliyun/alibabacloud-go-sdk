// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainDnssecInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *DescribeDomainDnssecInfoResponseBody
	GetAlgorithm() *string
	SetDigest(v string) *DescribeDomainDnssecInfoResponseBody
	GetDigest() *string
	SetDigestType(v string) *DescribeDomainDnssecInfoResponseBody
	GetDigestType() *string
	SetDomainName(v string) *DescribeDomainDnssecInfoResponseBody
	GetDomainName() *string
	SetDsRecord(v string) *DescribeDomainDnssecInfoResponseBody
	GetDsRecord() *string
	SetFlags(v string) *DescribeDomainDnssecInfoResponseBody
	GetFlags() *string
	SetKeyTag(v string) *DescribeDomainDnssecInfoResponseBody
	GetKeyTag() *string
	SetPublicKey(v string) *DescribeDomainDnssecInfoResponseBody
	GetPublicKey() *string
	SetRequestId(v string) *DescribeDomainDnssecInfoResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeDomainDnssecInfoResponseBody
	GetStatus() *string
}

type DescribeDomainDnssecInfoResponseBody struct {
	// The algorithm type. This parameter is returned if DNSSEC is enabled.
	//
	// example:
	//
	// 13
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The digest. This parameter is returned if DNSSEC is enabled.
	//
	// example:
	//
	// C1A0424B97A049F1F9B2EA139CC298533219668164E343BD21203ABC4608C02A
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The digest type. This parameter is returned if DNSSEC is enabled.
	//
	// example:
	//
	// SHA256
	DigestType *string `json:"DigestType,omitempty" xml:"DigestType,omitempty"`
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The delegation signer (DS) record. This parameter is returned if DNSSEC is enabled.
	//
	// example:
	//
	// example.com. 3600 IN DS 2371 13 2 C1A0424B97A049F1F9B2EA139CC298533219668164E343BD21203ABC4608C02A
	DsRecord *string `json:"DsRecord,omitempty" xml:"DsRecord,omitempty"`
	// The flag. This parameter is returned if DNSSEC is enabled.
	//
	// example:
	//
	// 257 (KSK)
	Flags *string `json:"Flags,omitempty" xml:"Flags,omitempty"`
	// The key tag. This parameter is returned if DNSSEC is enabled.
	//
	// example:
	//
	// 54931
	KeyTag *string `json:"KeyTag,omitempty" xml:"KeyTag,omitempty"`
	// The public key. This parameter is returned if DNSSEC is enabled.
	//
	// example:
	//
	// mdsswUyr3DPW132mOi8V9xESWE8jTo0dxCjjnopKl+GqJxpVXckHAeF+KkxLbxILfDLUT0rAK9iUzy1L53eKGQ==
	PublicKey *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The state of the DNSSEC. Valid values:
	//
	// 	- ON
	//
	// 	- OFF
	//
	// example:
	//
	// ON
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDomainDnssecInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainDnssecInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainDnssecInfoResponseBody) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *DescribeDomainDnssecInfoResponseBody) GetDigest() *string {
	return s.Digest
}

func (s *DescribeDomainDnssecInfoResponseBody) GetDigestType() *string {
	return s.DigestType
}

func (s *DescribeDomainDnssecInfoResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainDnssecInfoResponseBody) GetDsRecord() *string {
	return s.DsRecord
}

func (s *DescribeDomainDnssecInfoResponseBody) GetFlags() *string {
	return s.Flags
}

func (s *DescribeDomainDnssecInfoResponseBody) GetKeyTag() *string {
	return s.KeyTag
}

func (s *DescribeDomainDnssecInfoResponseBody) GetPublicKey() *string {
	return s.PublicKey
}

func (s *DescribeDomainDnssecInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainDnssecInfoResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeDomainDnssecInfoResponseBody) SetAlgorithm(v string) *DescribeDomainDnssecInfoResponseBody {
	s.Algorithm = &v
	return s
}

func (s *DescribeDomainDnssecInfoResponseBody) SetDigest(v string) *DescribeDomainDnssecInfoResponseBody {
	s.Digest = &v
	return s
}

func (s *DescribeDomainDnssecInfoResponseBody) SetDigestType(v string) *DescribeDomainDnssecInfoResponseBody {
	s.DigestType = &v
	return s
}

func (s *DescribeDomainDnssecInfoResponseBody) SetDomainName(v string) *DescribeDomainDnssecInfoResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainDnssecInfoResponseBody) SetDsRecord(v string) *DescribeDomainDnssecInfoResponseBody {
	s.DsRecord = &v
	return s
}

func (s *DescribeDomainDnssecInfoResponseBody) SetFlags(v string) *DescribeDomainDnssecInfoResponseBody {
	s.Flags = &v
	return s
}

func (s *DescribeDomainDnssecInfoResponseBody) SetKeyTag(v string) *DescribeDomainDnssecInfoResponseBody {
	s.KeyTag = &v
	return s
}

func (s *DescribeDomainDnssecInfoResponseBody) SetPublicKey(v string) *DescribeDomainDnssecInfoResponseBody {
	s.PublicKey = &v
	return s
}

func (s *DescribeDomainDnssecInfoResponseBody) SetRequestId(v string) *DescribeDomainDnssecInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainDnssecInfoResponseBody) SetStatus(v string) *DescribeDomainDnssecInfoResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDomainDnssecInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
