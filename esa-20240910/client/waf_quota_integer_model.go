// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWafQuotaInteger interface {
	dara.Model
	String() string
	GoString() string
	SetEqual(v int32) *WafQuotaInteger
	GetEqual() *int32
	SetGreaterThan(v int32) *WafQuotaInteger
	GetGreaterThan() *int32
	SetGreaterThanOrEqual(v int32) *WafQuotaInteger
	GetGreaterThanOrEqual() *int32
	SetLessThan(v int32) *WafQuotaInteger
	GetLessThan() *int32
	SetLessThanOrEqual(v int32) *WafQuotaInteger
	GetLessThanOrEqual() *int32
}

type WafQuotaInteger struct {
	Equal              *int32 `json:"Equal,omitempty" xml:"Equal,omitempty"`
	GreaterThan        *int32 `json:"GreaterThan,omitempty" xml:"GreaterThan,omitempty"`
	GreaterThanOrEqual *int32 `json:"GreaterThanOrEqual,omitempty" xml:"GreaterThanOrEqual,omitempty"`
	LessThan           *int32 `json:"LessThan,omitempty" xml:"LessThan,omitempty"`
	LessThanOrEqual    *int32 `json:"LessThanOrEqual,omitempty" xml:"LessThanOrEqual,omitempty"`
}

func (s WafQuotaInteger) String() string {
	return dara.Prettify(s)
}

func (s WafQuotaInteger) GoString() string {
	return s.String()
}

func (s *WafQuotaInteger) GetEqual() *int32 {
	return s.Equal
}

func (s *WafQuotaInteger) GetGreaterThan() *int32 {
	return s.GreaterThan
}

func (s *WafQuotaInteger) GetGreaterThanOrEqual() *int32 {
	return s.GreaterThanOrEqual
}

func (s *WafQuotaInteger) GetLessThan() *int32 {
	return s.LessThan
}

func (s *WafQuotaInteger) GetLessThanOrEqual() *int32 {
	return s.LessThanOrEqual
}

func (s *WafQuotaInteger) SetEqual(v int32) *WafQuotaInteger {
	s.Equal = &v
	return s
}

func (s *WafQuotaInteger) SetGreaterThan(v int32) *WafQuotaInteger {
	s.GreaterThan = &v
	return s
}

func (s *WafQuotaInteger) SetGreaterThanOrEqual(v int32) *WafQuotaInteger {
	s.GreaterThanOrEqual = &v
	return s
}

func (s *WafQuotaInteger) SetLessThan(v int32) *WafQuotaInteger {
	s.LessThan = &v
	return s
}

func (s *WafQuotaInteger) SetLessThanOrEqual(v int32) *WafQuotaInteger {
	s.LessThanOrEqual = &v
	return s
}

func (s *WafQuotaInteger) Validate() error {
	return dara.Validate(s)
}
