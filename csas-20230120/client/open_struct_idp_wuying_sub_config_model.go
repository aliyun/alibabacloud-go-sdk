// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenStructIdpWuyingSubConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAliuids(v []*string) *OpenStructIdpWuyingSubConfig
	GetAliuids() []*string
}

type OpenStructIdpWuyingSubConfig struct {
	Aliuids []*string `json:"Aliuids,omitempty" xml:"Aliuids,omitempty" type:"Repeated"`
}

func (s OpenStructIdpWuyingSubConfig) String() string {
	return dara.Prettify(s)
}

func (s OpenStructIdpWuyingSubConfig) GoString() string {
	return s.String()
}

func (s *OpenStructIdpWuyingSubConfig) GetAliuids() []*string {
	return s.Aliuids
}

func (s *OpenStructIdpWuyingSubConfig) SetAliuids(v []*string) *OpenStructIdpWuyingSubConfig {
	s.Aliuids = v
	return s
}

func (s *OpenStructIdpWuyingSubConfig) Validate() error {
	return dara.Validate(s)
}
