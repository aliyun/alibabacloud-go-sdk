// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCustomHostAlias interface {
	dara.Model
	String() string
	GoString() string
	SetHostAliases(v []*HostAlias) *CustomHostAlias
	GetHostAliases() []*HostAlias
}

type CustomHostAlias struct {
	HostAliases []*HostAlias `json:"hostAliases,omitempty" xml:"hostAliases,omitempty" type:"Repeated"`
}

func (s CustomHostAlias) String() string {
	return dara.Prettify(s)
}

func (s CustomHostAlias) GoString() string {
	return s.String()
}

func (s *CustomHostAlias) GetHostAliases() []*HostAlias {
	return s.HostAliases
}

func (s *CustomHostAlias) SetHostAliases(v []*HostAlias) *CustomHostAlias {
	s.HostAliases = v
	return s
}

func (s *CustomHostAlias) Validate() error {
	if s.HostAliases != nil {
		for _, item := range s.HostAliases {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
