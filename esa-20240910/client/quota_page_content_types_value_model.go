// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuotaPageContentTypesValue interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v bool) *QuotaPageContentTypesValue
	GetEnable() *bool
	SetContentLength(v *WafQuotaInteger) *QuotaPageContentTypesValue
	GetContentLength() *WafQuotaInteger
}

type QuotaPageContentTypesValue struct {
	// The switch for the Content-Type type in custom response pages.
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The content length quota for the Content-Type in custom response pages.
	ContentLength *WafQuotaInteger `json:"ContentLength,omitempty" xml:"ContentLength,omitempty"`
}

func (s QuotaPageContentTypesValue) String() string {
	return dara.Prettify(s)
}

func (s QuotaPageContentTypesValue) GoString() string {
	return s.String()
}

func (s *QuotaPageContentTypesValue) GetEnable() *bool {
	return s.Enable
}

func (s *QuotaPageContentTypesValue) GetContentLength() *WafQuotaInteger {
	return s.ContentLength
}

func (s *QuotaPageContentTypesValue) SetEnable(v bool) *QuotaPageContentTypesValue {
	s.Enable = &v
	return s
}

func (s *QuotaPageContentTypesValue) SetContentLength(v *WafQuotaInteger) *QuotaPageContentTypesValue {
	s.ContentLength = v
	return s
}

func (s *QuotaPageContentTypesValue) Validate() error {
	if s.ContentLength != nil {
		if err := s.ContentLength.Validate(); err != nil {
			return err
		}
	}
	return nil
}
