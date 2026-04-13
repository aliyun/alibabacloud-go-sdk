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
	// This parameter is required.
	//
	// example:
	//
	// 0123456789ABCDEF0123456789ABCDEF
	Cak *string `json:"Cak,omitempty" xml:"Cak,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// GCM-AES-128
	CipherSuite *string `json:"CipherSuite,omitempty" xml:"CipherSuite,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0123456789ABCDEF0123456789ABCDEF
	Ckn *string `json:"Ckn,omitempty" xml:"Ckn,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-bp1hp0wr072f6****
	PhysicalConnectionId *string `json:"PhysicalConnectionId,omitempty" xml:"PhysicalConnectionId,omitempty"`
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
