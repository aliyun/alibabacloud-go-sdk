// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuotaListItemsValue interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v bool) *QuotaListItemsValue
	GetEnable() *bool
	SetValue(v *WafQuotaString) *QuotaListItemsValue
	GetValue() *WafQuotaString
}

type QuotaListItemsValue struct {
	// The switch for the type of item in the custom list.
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// Format restrictions for the type of item in the custom list.
	Value *WafQuotaString `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QuotaListItemsValue) String() string {
	return dara.Prettify(s)
}

func (s QuotaListItemsValue) GoString() string {
	return s.String()
}

func (s *QuotaListItemsValue) GetEnable() *bool {
	return s.Enable
}

func (s *QuotaListItemsValue) GetValue() *WafQuotaString {
	return s.Value
}

func (s *QuotaListItemsValue) SetEnable(v bool) *QuotaListItemsValue {
	s.Enable = &v
	return s
}

func (s *QuotaListItemsValue) SetValue(v *WafQuotaString) *QuotaListItemsValue {
	s.Value = v
	return s
}

func (s *QuotaListItemsValue) Validate() error {
	if s.Value != nil {
		if err := s.Value.Validate(); err != nil {
			return err
		}
	}
	return nil
}
