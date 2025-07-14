// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebAclEntryConfig interface {
	dara.Model
	String() string
	GoString() string
	SetEntry(v string) *WebAclEntryConfig
	GetEntry() *string
}

type WebAclEntryConfig struct {
	// This parameter is required.
	//
	// example:
	//
	// 1.2.3.4/32
	Entry *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
}

func (s WebAclEntryConfig) String() string {
	return dara.Prettify(s)
}

func (s WebAclEntryConfig) GoString() string {
	return s.String()
}

func (s *WebAclEntryConfig) GetEntry() *string {
	return s.Entry
}

func (s *WebAclEntryConfig) SetEntry(v string) *WebAclEntryConfig {
	s.Entry = &v
	return s
}

func (s *WebAclEntryConfig) Validate() error {
	return dara.Validate(s)
}
