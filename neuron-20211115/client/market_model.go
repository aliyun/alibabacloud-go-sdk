// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMarket interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *Market
	GetId() *int64
	SetName(v string) *Market
	GetName() *string
	SetType(v string) *Market
	GetType() *string
}

type Market struct {
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 内部市场
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s Market) String() string {
	return dara.Prettify(s)
}

func (s Market) GoString() string {
	return s.String()
}

func (s *Market) GetId() *int64 {
	return s.Id
}

func (s *Market) GetName() *string {
	return s.Name
}

func (s *Market) GetType() *string {
	return s.Type
}

func (s *Market) SetId(v int64) *Market {
	s.Id = &v
	return s
}

func (s *Market) SetName(v string) *Market {
	s.Name = &v
	return s
}

func (s *Market) SetType(v string) *Market {
	s.Type = &v
	return s
}

func (s *Market) Validate() error {
	return dara.Validate(s)
}
