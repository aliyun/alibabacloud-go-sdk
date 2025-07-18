// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPAL7ConfigReplaceRule interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v string) *PAL7ConfigReplaceRule
	GetFrom() *string
	SetTo(v string) *PAL7ConfigReplaceRule
	GetTo() *string
}

type PAL7ConfigReplaceRule struct {
	// example:
	//
	// aaa
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// bbb
	To *string `json:"To,omitempty" xml:"To,omitempty"`
}

func (s PAL7ConfigReplaceRule) String() string {
	return dara.Prettify(s)
}

func (s PAL7ConfigReplaceRule) GoString() string {
	return s.String()
}

func (s *PAL7ConfigReplaceRule) GetFrom() *string {
	return s.From
}

func (s *PAL7ConfigReplaceRule) GetTo() *string {
	return s.To
}

func (s *PAL7ConfigReplaceRule) SetFrom(v string) *PAL7ConfigReplaceRule {
	s.From = &v
	return s
}

func (s *PAL7ConfigReplaceRule) SetTo(v string) *PAL7ConfigReplaceRule {
	s.To = &v
	return s
}

func (s *PAL7ConfigReplaceRule) Validate() error {
	return dara.Validate(s)
}
