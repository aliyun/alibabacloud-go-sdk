// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebAclConfig interface {
	dara.Model
	String() string
	GoString() string
	SetWebAclEntries(v []*WebAclEntryConfig) *WebAclConfig
	GetWebAclEntries() []*WebAclEntryConfig
}

type WebAclConfig struct {
	// This parameter is required.
	//
	// if can be null:
	// true
	WebAclEntries []*WebAclEntryConfig `json:"WebAclEntries,omitempty" xml:"WebAclEntries,omitempty" type:"Repeated"`
}

func (s WebAclConfig) String() string {
	return dara.Prettify(s)
}

func (s WebAclConfig) GoString() string {
	return s.String()
}

func (s *WebAclConfig) GetWebAclEntries() []*WebAclEntryConfig {
	return s.WebAclEntries
}

func (s *WebAclConfig) SetWebAclEntries(v []*WebAclEntryConfig) *WebAclConfig {
	s.WebAclEntries = v
	return s
}

func (s *WebAclConfig) Validate() error {
	if s.WebAclEntries != nil {
		for _, item := range s.WebAclEntries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
