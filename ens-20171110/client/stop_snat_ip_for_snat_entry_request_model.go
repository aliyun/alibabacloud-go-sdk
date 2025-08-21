// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopSnatIpForSnatEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSnatEntryId(v string) *StopSnatIpForSnatEntryRequest
	GetSnatEntryId() *string
	SetSnatIp(v string) *StopSnatIpForSnatEntryRequest
	GetSnatIp() *string
}

type StopSnatIpForSnatEntryRequest struct {
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
	// 221.178.103.143
	SnatIp *string `json:"SnatIp,omitempty" xml:"SnatIp,omitempty"`
}

func (s StopSnatIpForSnatEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s StopSnatIpForSnatEntryRequest) GoString() string {
	return s.String()
}

func (s *StopSnatIpForSnatEntryRequest) GetSnatEntryId() *string {
	return s.SnatEntryId
}

func (s *StopSnatIpForSnatEntryRequest) GetSnatIp() *string {
	return s.SnatIp
}

func (s *StopSnatIpForSnatEntryRequest) SetSnatEntryId(v string) *StopSnatIpForSnatEntryRequest {
	s.SnatEntryId = &v
	return s
}

func (s *StopSnatIpForSnatEntryRequest) SetSnatIp(v string) *StopSnatIpForSnatEntryRequest {
	s.SnatIp = &v
	return s
}

func (s *StopSnatIpForSnatEntryRequest) Validate() error {
	return dara.Validate(s)
}
