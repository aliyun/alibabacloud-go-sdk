// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchHbaseHaSlbRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBdsId(v string) *SwitchHbaseHaSlbRequest
	GetBdsId() *string
	SetHaId(v string) *SwitchHbaseHaSlbRequest
	GetHaId() *string
	SetHaTypes(v string) *SwitchHbaseHaSlbRequest
	GetHaTypes() *string
	SetHbaseType(v string) *SwitchHbaseHaSlbRequest
	GetHbaseType() *string
}

type SwitchHbaseHaSlbRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// bds-t4n3496whj23ia4k
	BdsId *string `json:"BdsId,omitempty" xml:"BdsId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ha-v21tmnxjwh2yuy1il
	HaId *string `json:"HaId,omitempty" xml:"HaId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// thrift
	HaTypes *string `json:"HaTypes,omitempty" xml:"HaTypes,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Active
	HbaseType *string `json:"HbaseType,omitempty" xml:"HbaseType,omitempty"`
}

func (s SwitchHbaseHaSlbRequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchHbaseHaSlbRequest) GoString() string {
	return s.String()
}

func (s *SwitchHbaseHaSlbRequest) GetBdsId() *string {
	return s.BdsId
}

func (s *SwitchHbaseHaSlbRequest) GetHaId() *string {
	return s.HaId
}

func (s *SwitchHbaseHaSlbRequest) GetHaTypes() *string {
	return s.HaTypes
}

func (s *SwitchHbaseHaSlbRequest) GetHbaseType() *string {
	return s.HbaseType
}

func (s *SwitchHbaseHaSlbRequest) SetBdsId(v string) *SwitchHbaseHaSlbRequest {
	s.BdsId = &v
	return s
}

func (s *SwitchHbaseHaSlbRequest) SetHaId(v string) *SwitchHbaseHaSlbRequest {
	s.HaId = &v
	return s
}

func (s *SwitchHbaseHaSlbRequest) SetHaTypes(v string) *SwitchHbaseHaSlbRequest {
	s.HaTypes = &v
	return s
}

func (s *SwitchHbaseHaSlbRequest) SetHbaseType(v string) *SwitchHbaseHaSlbRequest {
	s.HbaseType = &v
	return s
}

func (s *SwitchHbaseHaSlbRequest) Validate() error {
	return dara.Validate(s)
}
