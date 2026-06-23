// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateMacSecKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCkn(v string) *DisassociateMacSecKeyRequest
	GetCkn() *string
	SetPhysicalConnectionId(v string) *DisassociateMacSecKeyRequest
	GetPhysicalConnectionId() *string
	SetRegionId(v string) *DisassociateMacSecKeyRequest
	GetRegionId() *string
}

type DisassociateMacSecKeyRequest struct {
	// The key name. Only hexadecimal characters are supported. Lowercase characters are automatically converted to uppercase. When the encryption algorithm type is GCM-AES-128 or GCM-AES-XPN-128, the length must be 32 hexadecimal characters. When the encryption algorithm type is GCM-AES-256 or GCM-AES-XPN-256, the length must be 64 hexadecimal characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0123456789ABCDEF0123456789ABCDEF
	Ckn *string `json:"Ckn,omitempty" xml:"Ckn,omitempty"`
	// The ID of the Express Connect circuit.
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

func (s DisassociateMacSecKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DisassociateMacSecKeyRequest) GoString() string {
	return s.String()
}

func (s *DisassociateMacSecKeyRequest) GetCkn() *string {
	return s.Ckn
}

func (s *DisassociateMacSecKeyRequest) GetPhysicalConnectionId() *string {
	return s.PhysicalConnectionId
}

func (s *DisassociateMacSecKeyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisassociateMacSecKeyRequest) SetCkn(v string) *DisassociateMacSecKeyRequest {
	s.Ckn = &v
	return s
}

func (s *DisassociateMacSecKeyRequest) SetPhysicalConnectionId(v string) *DisassociateMacSecKeyRequest {
	s.PhysicalConnectionId = &v
	return s
}

func (s *DisassociateMacSecKeyRequest) SetRegionId(v string) *DisassociateMacSecKeyRequest {
	s.RegionId = &v
	return s
}

func (s *DisassociateMacSecKeyRequest) Validate() error {
	return dara.Validate(s)
}
