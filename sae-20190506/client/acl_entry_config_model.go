// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAclEntryConfig interface {
	dara.Model
	String() string
	GoString() string
	SetEntry(v string) *AclEntryConfig
	GetEntry() *string
}

type AclEntryConfig struct {
	// This parameter is required.
	//
	// example:
	//
	// 1.2.3.4/32
	Entry *string `json:"entry,omitempty" xml:"entry,omitempty"`
}

func (s AclEntryConfig) String() string {
	return dara.Prettify(s)
}

func (s AclEntryConfig) GoString() string {
	return s.String()
}

func (s *AclEntryConfig) GetEntry() *string {
	return s.Entry
}

func (s *AclEntryConfig) SetEntry(v string) *AclEntryConfig {
	s.Entry = &v
	return s
}

func (s *AclEntryConfig) Validate() error {
	return dara.Validate(s)
}
