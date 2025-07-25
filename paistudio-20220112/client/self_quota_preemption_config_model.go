// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSelfQuotaPreemptionConfig interface {
	dara.Model
	String() string
	GoString() string
	SetPreemptedPriorities(v []*int32) *SelfQuotaPreemptionConfig
	GetPreemptedPriorities() []*int32
	SetPreemptedProducts(v []*string) *SelfQuotaPreemptionConfig
	GetPreemptedProducts() []*string
	SetPreemptorPriorities(v []*int32) *SelfQuotaPreemptionConfig
	GetPreemptorPriorities() []*int32
}

type SelfQuotaPreemptionConfig struct {
	PreemptedPriorities []*int32  `json:"PreemptedPriorities,omitempty" xml:"PreemptedPriorities,omitempty" type:"Repeated"`
	PreemptedProducts   []*string `json:"PreemptedProducts,omitempty" xml:"PreemptedProducts,omitempty" type:"Repeated"`
	PreemptorPriorities []*int32  `json:"PreemptorPriorities,omitempty" xml:"PreemptorPriorities,omitempty" type:"Repeated"`
}

func (s SelfQuotaPreemptionConfig) String() string {
	return dara.Prettify(s)
}

func (s SelfQuotaPreemptionConfig) GoString() string {
	return s.String()
}

func (s *SelfQuotaPreemptionConfig) GetPreemptedPriorities() []*int32 {
	return s.PreemptedPriorities
}

func (s *SelfQuotaPreemptionConfig) GetPreemptedProducts() []*string {
	return s.PreemptedProducts
}

func (s *SelfQuotaPreemptionConfig) GetPreemptorPriorities() []*int32 {
	return s.PreemptorPriorities
}

func (s *SelfQuotaPreemptionConfig) SetPreemptedPriorities(v []*int32) *SelfQuotaPreemptionConfig {
	s.PreemptedPriorities = v
	return s
}

func (s *SelfQuotaPreemptionConfig) SetPreemptedProducts(v []*string) *SelfQuotaPreemptionConfig {
	s.PreemptedProducts = v
	return s
}

func (s *SelfQuotaPreemptionConfig) SetPreemptorPriorities(v []*int32) *SelfQuotaPreemptionConfig {
	s.PreemptorPriorities = v
	return s
}

func (s *SelfQuotaPreemptionConfig) Validate() error {
	return dara.Validate(s)
}
