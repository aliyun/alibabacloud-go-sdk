// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubQuotaPreemptionConfig interface {
	dara.Model
	String() string
	GoString() string
	SetPreemptedPriorityUpperBound(v int64) *SubQuotaPreemptionConfig
	GetPreemptedPriorityUpperBound() *int64
	SetPreemptedProducts(v []*string) *SubQuotaPreemptionConfig
	GetPreemptedProducts() []*string
}

type SubQuotaPreemptionConfig struct {
	// example:
	//
	// 9
	PreemptedPriorityUpperBound *int64    `json:"PreemptedPriorityUpperBound,omitempty" xml:"PreemptedPriorityUpperBound,omitempty"`
	PreemptedProducts           []*string `json:"PreemptedProducts,omitempty" xml:"PreemptedProducts,omitempty" type:"Repeated"`
}

func (s SubQuotaPreemptionConfig) String() string {
	return dara.Prettify(s)
}

func (s SubQuotaPreemptionConfig) GoString() string {
	return s.String()
}

func (s *SubQuotaPreemptionConfig) GetPreemptedPriorityUpperBound() *int64 {
	return s.PreemptedPriorityUpperBound
}

func (s *SubQuotaPreemptionConfig) GetPreemptedProducts() []*string {
	return s.PreemptedProducts
}

func (s *SubQuotaPreemptionConfig) SetPreemptedPriorityUpperBound(v int64) *SubQuotaPreemptionConfig {
	s.PreemptedPriorityUpperBound = &v
	return s
}

func (s *SubQuotaPreemptionConfig) SetPreemptedProducts(v []*string) *SubQuotaPreemptionConfig {
	s.PreemptedProducts = v
	return s
}

func (s *SubQuotaPreemptionConfig) Validate() error {
	return dara.Validate(s)
}
