// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrefixFiltersIncludeItem interface {
	dara.Model
	String() string
	GoString() string
	SetPrefix(v []*string) *PrefixFiltersIncludeItem
	GetPrefix() []*string
}

type PrefixFiltersIncludeItem struct {
	// The full object key prefix. A maximum of 10,000 prefixes are supported by default, and each prefix can be up to 1,024 bytes in length.
	//
	// example:
	//
	// source-prefix/dir-a/
	Prefix []*string `json:"Prefix,omitempty" xml:"Prefix,omitempty" type:"Repeated"`
}

func (s PrefixFiltersIncludeItem) String() string {
	return dara.Prettify(s)
}

func (s PrefixFiltersIncludeItem) GoString() string {
	return s.String()
}

func (s *PrefixFiltersIncludeItem) GetPrefix() []*string {
	return s.Prefix
}

func (s *PrefixFiltersIncludeItem) SetPrefix(v []*string) *PrefixFiltersIncludeItem {
	s.Prefix = v
	return s
}

func (s *PrefixFiltersIncludeItem) Validate() error {
	return dara.Validate(s)
}
