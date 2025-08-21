// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySnatEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEipAffinity(v bool) *ModifySnatEntryRequest
	GetEipAffinity() *bool
	SetIspAffinity(v bool) *ModifySnatEntryRequest
	GetIspAffinity() *bool
	SetSnatEntryId(v string) *ModifySnatEntryRequest
	GetSnatEntryId() *string
	SetSnatEntryName(v string) *ModifySnatEntryRequest
	GetSnatEntryName() *string
	SetSnatIp(v string) *ModifySnatEntryRequest
	GetSnatIp() *string
}

type ModifySnatEntryRequest struct {
	// Specifies whether to enable IP affinity. Default value: true. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// >  Description After you enable IP affinity, if multiple EIPs are associated with an SNAT entry, one client uses the same EIP to for communication. If IP affinity is disabled, the client uses a random EIP for communication.
	//
	// example:
	//
	// false
	EipAffinity *bool `json:"EipAffinity,omitempty" xml:"EipAffinity,omitempty"`
	// example:
	//
	// true
	IspAffinity *bool `json:"IspAffinity,omitempty" xml:"IspAffinity,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// snat-5tfjp36fsrb36zs36faj0****
	SnatEntryId *string `json:"SnatEntryId,omitempty" xml:"SnatEntryId,omitempty"`
	// example:
	//
	// test0
	SnatEntryName *string `json:"SnatEntryName,omitempty" xml:"SnatEntryName,omitempty"`
	// Separate multiple EIPs in the SNAT entry with commas (,).
	//
	// example:
	//
	// 120.XXX.XXX.71
	SnatIp *string `json:"SnatIp,omitempty" xml:"SnatIp,omitempty"`
}

func (s ModifySnatEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySnatEntryRequest) GoString() string {
	return s.String()
}

func (s *ModifySnatEntryRequest) GetEipAffinity() *bool {
	return s.EipAffinity
}

func (s *ModifySnatEntryRequest) GetIspAffinity() *bool {
	return s.IspAffinity
}

func (s *ModifySnatEntryRequest) GetSnatEntryId() *string {
	return s.SnatEntryId
}

func (s *ModifySnatEntryRequest) GetSnatEntryName() *string {
	return s.SnatEntryName
}

func (s *ModifySnatEntryRequest) GetSnatIp() *string {
	return s.SnatIp
}

func (s *ModifySnatEntryRequest) SetEipAffinity(v bool) *ModifySnatEntryRequest {
	s.EipAffinity = &v
	return s
}

func (s *ModifySnatEntryRequest) SetIspAffinity(v bool) *ModifySnatEntryRequest {
	s.IspAffinity = &v
	return s
}

func (s *ModifySnatEntryRequest) SetSnatEntryId(v string) *ModifySnatEntryRequest {
	s.SnatEntryId = &v
	return s
}

func (s *ModifySnatEntryRequest) SetSnatEntryName(v string) *ModifySnatEntryRequest {
	s.SnatEntryName = &v
	return s
}

func (s *ModifySnatEntryRequest) SetSnatIp(v string) *ModifySnatEntryRequest {
	s.SnatIp = &v
	return s
}

func (s *ModifySnatEntryRequest) Validate() error {
	return dara.Validate(s)
}
