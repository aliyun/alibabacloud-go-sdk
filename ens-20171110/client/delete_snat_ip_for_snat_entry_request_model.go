// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSnatIpForSnatEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSnatEntryId(v string) *DeleteSnatIpForSnatEntryRequest
	GetSnatEntryId() *string
	SetSnatIp(v string) *DeleteSnatIpForSnatEntryRequest
	GetSnatIp() *string
}

type DeleteSnatIpForSnatEntryRequest struct {
	// The ID of the SNAT entry.
	//
	// This parameter is required.
	//
	// example:
	//
	// snat-5tfi6f8gds82mjmlofeym****
	SnatEntryId *string `json:"SnatEntryId,omitempty" xml:"SnatEntryId,omitempty"`
	// The EIP that you want to delete from the SNAT entry.
	//
	// This parameter is required.
	//
	// example:
	//
	// 120.72.56.71
	SnatIp *string `json:"SnatIp,omitempty" xml:"SnatIp,omitempty"`
}

func (s DeleteSnatIpForSnatEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnatIpForSnatEntryRequest) GoString() string {
	return s.String()
}

func (s *DeleteSnatIpForSnatEntryRequest) GetSnatEntryId() *string {
	return s.SnatEntryId
}

func (s *DeleteSnatIpForSnatEntryRequest) GetSnatIp() *string {
	return s.SnatIp
}

func (s *DeleteSnatIpForSnatEntryRequest) SetSnatEntryId(v string) *DeleteSnatIpForSnatEntryRequest {
	s.SnatEntryId = &v
	return s
}

func (s *DeleteSnatIpForSnatEntryRequest) SetSnatIp(v string) *DeleteSnatIpForSnatEntryRequest {
	s.SnatIp = &v
	return s
}

func (s *DeleteSnatIpForSnatEntryRequest) Validate() error {
	return dara.Validate(s)
}
