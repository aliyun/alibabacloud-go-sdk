// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKeyFilterItem interface {
	dara.Model
	String() string
	GoString() string
	SetRegex(v []*string) *KeyFilterItem
	GetRegex() []*string
}

type KeyFilterItem struct {
	Regex []*string `json:"Regex,omitempty" xml:"Regex,omitempty" type:"Repeated"`
}

func (s KeyFilterItem) String() string {
	return dara.Prettify(s)
}

func (s KeyFilterItem) GoString() string {
	return s.String()
}

func (s *KeyFilterItem) GetRegex() []*string {
	return s.Regex
}

func (s *KeyFilterItem) SetRegex(v []*string) *KeyFilterItem {
	s.Regex = v
	return s
}

func (s *KeyFilterItem) Validate() error {
	return dara.Validate(s)
}
