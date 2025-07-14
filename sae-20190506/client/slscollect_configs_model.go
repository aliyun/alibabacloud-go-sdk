// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSLSCollectConfigs interface {
	dara.Model
	String() string
	GoString() string
	SetCollectConfigs(v []*SLSCollectConfig) *SLSCollectConfigs
	GetCollectConfigs() []*SLSCollectConfig
}

type SLSCollectConfigs struct {
	CollectConfigs []*SLSCollectConfig `json:"CollectConfigs,omitempty" xml:"CollectConfigs,omitempty" type:"Repeated"`
}

func (s SLSCollectConfigs) String() string {
	return dara.Prettify(s)
}

func (s SLSCollectConfigs) GoString() string {
	return s.String()
}

func (s *SLSCollectConfigs) GetCollectConfigs() []*SLSCollectConfig {
	return s.CollectConfigs
}

func (s *SLSCollectConfigs) SetCollectConfigs(v []*SLSCollectConfig) *SLSCollectConfigs {
	s.CollectConfigs = v
	return s
}

func (s *SLSCollectConfigs) Validate() error {
	return dara.Validate(s)
}
