// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHBaseHaDBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBdsId(v string) *DeleteHBaseHaDBRequest
	GetBdsId() *string
	SetHaId(v string) *DeleteHBaseHaDBRequest
	GetHaId() *string
}

type DeleteHBaseHaDBRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// bds-bp14112fd7g52s1****
	BdsId *string `json:"BdsId,omitempty" xml:"BdsId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ha-sw2o0l01s4r76****
	HaId *string `json:"HaId,omitempty" xml:"HaId,omitempty"`
}

func (s DeleteHBaseHaDBRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHBaseHaDBRequest) GoString() string {
	return s.String()
}

func (s *DeleteHBaseHaDBRequest) GetBdsId() *string {
	return s.BdsId
}

func (s *DeleteHBaseHaDBRequest) GetHaId() *string {
	return s.HaId
}

func (s *DeleteHBaseHaDBRequest) SetBdsId(v string) *DeleteHBaseHaDBRequest {
	s.BdsId = &v
	return s
}

func (s *DeleteHBaseHaDBRequest) SetHaId(v string) *DeleteHBaseHaDBRequest {
	s.HaId = &v
	return s
}

func (s *DeleteHBaseHaDBRequest) Validate() error {
	return dara.Validate(s)
}
