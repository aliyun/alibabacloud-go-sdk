// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateMacSecKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCak(v string) *AssociateMacSecKeyRequest
	GetCak() *string
	SetCipherSuite(v string) *AssociateMacSecKeyRequest
	GetCipherSuite() *string
	SetCkn(v string) *AssociateMacSecKeyRequest
	GetCkn() *string
	SetPhysicalConnectionId(v string) *AssociateMacSecKeyRequest
	GetPhysicalConnectionId() *string
	SetRegionId(v string) *AssociateMacSecKeyRequest
	GetRegionId() *string
}

type AssociateMacSecKeyRequest struct {
	// The passphrase. Only hexadecimal characters are supported. Lowercase characters are automatically transformed to uppercase. When the encryption algorithm type is GCM-AES-128 or GCM-AES-XPN-128, the length must be 32 hexadecimal characters. When the encryption algorithm type is GCM-AES-256 or GCM-AES-XPN-256, the length must be 64 hexadecimal characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0123456789ABCDEF0123456789ABCDEF
	Cak *string `json:"Cak,omitempty" xml:"Cak,omitempty"`
	// The encryption algorithm type. Valid values:
	//
	// - GCM-AES-128
	//
	// - GCM-AES-XPN-128
	//
	// - GCM-AES-256
	//
	// - GCM-AES-XPN-256
	//
	// This parameter is required.
	//
	// example:
	//
	// GCM-AES-128
	CipherSuite *string `json:"CipherSuite,omitempty" xml:"CipherSuite,omitempty"`
	// The key name. Only hexadecimal characters are supported. Lowercase characters are automatically transformed to uppercase. When the encryption algorithm type is GCM-AES-128 or GCM-AES-XPN-128, the length must be 32 hexadecimal characters. When the encryption algorithm type is GCM-AES-256 or GCM-AES-XPN-256, the length must be 64 hexadecimal characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0123456789ABCDEF0123456789ABCDEF
	Ckn *string `json:"Ckn,omitempty" xml:"Ckn,omitempty"`
	// The ID of the Express Connect circuit. The Express Connect circuit must be a dedicated circuit that has completed payment, supports MACsec, and is in the **Enabled*	- state. You can call DescribePhysicalConnections to query the current status and device capabilities.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-bp1hp0wr072f6****
	PhysicalConnectionId *string `json:"PhysicalConnectionId,omitempty" xml:"PhysicalConnectionId,omitempty"`
	// The region ID of the Express Connect circuit.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/448570.html) operation to query region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AssociateMacSecKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateMacSecKeyRequest) GoString() string {
	return s.String()
}

func (s *AssociateMacSecKeyRequest) GetCak() *string {
	return s.Cak
}

func (s *AssociateMacSecKeyRequest) GetCipherSuite() *string {
	return s.CipherSuite
}

func (s *AssociateMacSecKeyRequest) GetCkn() *string {
	return s.Ckn
}

func (s *AssociateMacSecKeyRequest) GetPhysicalConnectionId() *string {
	return s.PhysicalConnectionId
}

func (s *AssociateMacSecKeyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AssociateMacSecKeyRequest) SetCak(v string) *AssociateMacSecKeyRequest {
	s.Cak = &v
	return s
}

func (s *AssociateMacSecKeyRequest) SetCipherSuite(v string) *AssociateMacSecKeyRequest {
	s.CipherSuite = &v
	return s
}

func (s *AssociateMacSecKeyRequest) SetCkn(v string) *AssociateMacSecKeyRequest {
	s.Ckn = &v
	return s
}

func (s *AssociateMacSecKeyRequest) SetPhysicalConnectionId(v string) *AssociateMacSecKeyRequest {
	s.PhysicalConnectionId = &v
	return s
}

func (s *AssociateMacSecKeyRequest) SetRegionId(v string) *AssociateMacSecKeyRequest {
	s.RegionId = &v
	return s
}

func (s *AssociateMacSecKeyRequest) Validate() error {
	return dara.Validate(s)
}
