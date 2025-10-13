// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAclConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAclEntries(v []*AclEntryConfig) *AclConfig
	GetAclEntries() []*AclEntryConfig
}

type AclConfig struct {
	// This parameter is required.
	//
	// if can be null:
	// true
	AclEntries []*AclEntryConfig `json:"aclEntries,omitempty" xml:"aclEntries,omitempty" type:"Repeated"`
}

func (s AclConfig) String() string {
	return dara.Prettify(s)
}

func (s AclConfig) GoString() string {
	return s.String()
}

func (s *AclConfig) GetAclEntries() []*AclEntryConfig {
	return s.AclEntries
}

func (s *AclConfig) SetAclEntries(v []*AclEntryConfig) *AclConfig {
	s.AclEntries = v
	return s
}

func (s *AclConfig) Validate() error {
	if s.AclEntries != nil {
		for _, item := range s.AclEntries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
