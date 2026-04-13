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
