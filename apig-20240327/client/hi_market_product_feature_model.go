// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHiMarketProductFeature interface {
	dara.Model
	String() string
	GoString() string
	SetModelFeature(v *HiMarketModelFeature) *HiMarketProductFeature
	GetModelFeature() *HiMarketModelFeature
}

type HiMarketProductFeature struct {
	ModelFeature *HiMarketModelFeature `json:"modelFeature,omitempty" xml:"modelFeature,omitempty"`
}

func (s HiMarketProductFeature) String() string {
	return dara.Prettify(s)
}

func (s HiMarketProductFeature) GoString() string {
	return s.String()
}

func (s *HiMarketProductFeature) GetModelFeature() *HiMarketModelFeature {
	return s.ModelFeature
}

func (s *HiMarketProductFeature) SetModelFeature(v *HiMarketModelFeature) *HiMarketProductFeature {
	s.ModelFeature = v
	return s
}

func (s *HiMarketProductFeature) Validate() error {
	if s.ModelFeature != nil {
		if err := s.ModelFeature.Validate(); err != nil {
			return err
		}
	}
	return nil
}
