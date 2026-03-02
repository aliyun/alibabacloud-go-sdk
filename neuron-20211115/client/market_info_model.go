// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMarketInfo interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *MarketInfo
	GetId() *int64
	SetName(v string) *MarketInfo
	GetName() *string
	SetType(v string) *MarketInfo
	GetType() *string
}

type MarketInfo struct {
	Id   *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s MarketInfo) String() string {
	return dara.Prettify(s)
}

func (s MarketInfo) GoString() string {
	return s.String()
}

func (s *MarketInfo) GetId() *int64 {
	return s.Id
}

func (s *MarketInfo) GetName() *string {
	return s.Name
}

func (s *MarketInfo) GetType() *string {
	return s.Type
}

func (s *MarketInfo) SetId(v int64) *MarketInfo {
	s.Id = &v
	return s
}

func (s *MarketInfo) SetName(v string) *MarketInfo {
	s.Name = &v
	return s
}

func (s *MarketInfo) SetType(v string) *MarketInfo {
	s.Type = &v
	return s
}

func (s *MarketInfo) Validate() error {
	return dara.Validate(s)
}
