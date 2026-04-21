// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHiMarketIcon interface {
	dara.Model
	String() string
	GoString() string
	SetType(v string) *HiMarketIcon
	GetType() *string
	SetValue(v string) *HiMarketIcon
	GetValue() *string
}

type HiMarketIcon struct {
	Type  *string `json:"type,omitempty" xml:"type,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s HiMarketIcon) String() string {
	return dara.Prettify(s)
}

func (s HiMarketIcon) GoString() string {
	return s.String()
}

func (s *HiMarketIcon) GetType() *string {
	return s.Type
}

func (s *HiMarketIcon) GetValue() *string {
	return s.Value
}

func (s *HiMarketIcon) SetType(v string) *HiMarketIcon {
	s.Type = &v
	return s
}

func (s *HiMarketIcon) SetValue(v string) *HiMarketIcon {
	s.Value = &v
	return s
}

func (s *HiMarketIcon) Validate() error {
	return dara.Validate(s)
}
