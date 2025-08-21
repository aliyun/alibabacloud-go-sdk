// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSnatIpForSnatEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSnatEntryId(v string) *StartSnatIpForSnatEntryRequest
	GetSnatEntryId() *string
	SetSnatIp(v string) *StartSnatIpForSnatEntryRequest
	GetSnatIp() *string
}

type StartSnatIpForSnatEntryRequest struct {
	// The ID of the SNAT entry.
	//
	// This parameter is required.
	//
	// example:
	//
	// snat-5tfi6f8gds82mjmlofeym****
	SnatEntryId *string `json:"SnatEntryId,omitempty" xml:"SnatEntryId,omitempty"`
	// The EIP specified in the SNAT entry.
	//
	// This parameter is required.
	//
	// example:
	//
	// 219.152.82.143
	SnatIp *string `json:"SnatIp,omitempty" xml:"SnatIp,omitempty"`
}

func (s StartSnatIpForSnatEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s StartSnatIpForSnatEntryRequest) GoString() string {
	return s.String()
}

func (s *StartSnatIpForSnatEntryRequest) GetSnatEntryId() *string {
	return s.SnatEntryId
}

func (s *StartSnatIpForSnatEntryRequest) GetSnatIp() *string {
	return s.SnatIp
}

func (s *StartSnatIpForSnatEntryRequest) SetSnatEntryId(v string) *StartSnatIpForSnatEntryRequest {
	s.SnatEntryId = &v
	return s
}

func (s *StartSnatIpForSnatEntryRequest) SetSnatIp(v string) *StartSnatIpForSnatEntryRequest {
	s.SnatIp = &v
	return s
}

func (s *StartSnatIpForSnatEntryRequest) Validate() error {
	return dara.Validate(s)
}
