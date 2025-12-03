// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHbaseHaSlbRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBdsId(v string) *DeleteHbaseHaSlbRequest
	GetBdsId() *string
	SetHaId(v string) *DeleteHbaseHaSlbRequest
	GetHaId() *string
	SetHaTypes(v string) *DeleteHbaseHaSlbRequest
	GetHaTypes() *string
}

type DeleteHbaseHaSlbRequest struct {
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
}

func (s DeleteHbaseHaSlbRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHbaseHaSlbRequest) GoString() string {
	return s.String()
}

func (s *DeleteHbaseHaSlbRequest) GetBdsId() *string {
	return s.BdsId
}

func (s *DeleteHbaseHaSlbRequest) GetHaId() *string {
	return s.HaId
}

func (s *DeleteHbaseHaSlbRequest) GetHaTypes() *string {
	return s.HaTypes
}

func (s *DeleteHbaseHaSlbRequest) SetBdsId(v string) *DeleteHbaseHaSlbRequest {
	s.BdsId = &v
	return s
}

func (s *DeleteHbaseHaSlbRequest) SetHaId(v string) *DeleteHbaseHaSlbRequest {
	s.HaId = &v
	return s
}

func (s *DeleteHbaseHaSlbRequest) SetHaTypes(v string) *DeleteHbaseHaSlbRequest {
	s.HaTypes = &v
	return s
}

func (s *DeleteHbaseHaSlbRequest) Validate() error {
	return dara.Validate(s)
}
