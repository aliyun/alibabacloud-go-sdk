// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWafQuotaString interface {
	dara.Model
	String() string
	GoString() string
	SetRegexp(v string) *WafQuotaString
	GetRegexp() *string
}

type WafQuotaString struct {
	Regexp *string `json:"Regexp,omitempty" xml:"Regexp,omitempty"`
}

func (s WafQuotaString) String() string {
	return dara.Prettify(s)
}

func (s WafQuotaString) GoString() string {
	return s.String()
}

func (s *WafQuotaString) GetRegexp() *string {
	return s.Regexp
}

func (s *WafQuotaString) SetRegexp(v string) *WafQuotaString {
	s.Regexp = &v
	return s
}

func (s *WafQuotaString) Validate() error {
	return dara.Validate(s)
}
