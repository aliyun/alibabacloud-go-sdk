// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSnatIpForSnatEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSnatEntryId(v string) *AddSnatIpForSnatEntryRequest
	GetSnatEntryId() *string
	SetSnatIp(v string) *AddSnatIpForSnatEntryRequest
	GetSnatIp() *string
}

type AddSnatIpForSnatEntryRequest struct {
	// The ID of the SNAT entry.
	//
	// This parameter is required.
	//
	// example:
	//
	// snat-5xkzf89dndkzh8yg9stzqz9m4
	SnatEntryId *string `json:"SnatEntryId,omitempty" xml:"SnatEntryId,omitempty"`
	// The EIP specified in the SNAT entry.
	//
	// This parameter is required.
	//
	// example:
	//
	// 219.152.82.144
	SnatIp *string `json:"SnatIp,omitempty" xml:"SnatIp,omitempty"`
}

func (s AddSnatIpForSnatEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s AddSnatIpForSnatEntryRequest) GoString() string {
	return s.String()
}

func (s *AddSnatIpForSnatEntryRequest) GetSnatEntryId() *string {
	return s.SnatEntryId
}

func (s *AddSnatIpForSnatEntryRequest) GetSnatIp() *string {
	return s.SnatIp
}

func (s *AddSnatIpForSnatEntryRequest) SetSnatEntryId(v string) *AddSnatIpForSnatEntryRequest {
	s.SnatEntryId = &v
	return s
}

func (s *AddSnatIpForSnatEntryRequest) SetSnatIp(v string) *AddSnatIpForSnatEntryRequest {
	s.SnatIp = &v
	return s
}

func (s *AddSnatIpForSnatEntryRequest) Validate() error {
	return dara.Validate(s)
}
