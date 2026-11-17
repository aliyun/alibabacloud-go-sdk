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
	// The name of the key. Only hexadecimal characters are supported, and lowercase letters are automatically converted to uppercase letters. The length must be 32 hexadecimal characters when the encryption algorithm is GCM-AES-128 or GCM-AES-XPN-128, and 64 hexadecimal characters when the encryption algorithm is GCM-AES-256 or GCM-AES-XPN-256. The value must be the name of a key that has been associated with the target Express Connect circuit by calling the AssociateMacSecKey operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0123456789ABCDEF0123456789ABCDEF
	Ckn *string `json:"Ckn,omitempty" xml:"Ckn,omitempty"`
	// The ID of the Express Connect circuit. The Express Connect circuit must be in the **Enabled*	- state and its **BusinessStatus*	- must be **Normal**. You can call the DescribePhysicalConnections operation to query the status.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-bp1hp0wr072f6****
	PhysicalConnectionId *string `json:"PhysicalConnectionId,omitempty" xml:"PhysicalConnectionId,omitempty"`
	// The region ID of the Express Connect circuit. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/448570.html) operation to query region IDs.
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
