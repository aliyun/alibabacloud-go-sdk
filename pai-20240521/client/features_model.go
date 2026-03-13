// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFeatures interface {
	dara.Model
	String() string
	GoString() string
	SetQuota(v *FeaturesQuota) *Features
	GetQuota() *FeaturesQuota
}

type Features struct {
	Quota *FeaturesQuota `json:"Quota,omitempty" xml:"Quota,omitempty" type:"Struct"`
}

func (s Features) String() string {
	return dara.Prettify(s)
}

func (s Features) GoString() string {
	return s.String()
}

func (s *Features) GetQuota() *FeaturesQuota {
	return s.Quota
}

func (s *Features) SetQuota(v *FeaturesQuota) *Features {
	s.Quota = v
	return s
}

func (s *Features) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FeaturesQuota struct {
	// example:
	//
	// true
	IsEnabled *bool `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
}

func (s FeaturesQuota) String() string {
	return dara.Prettify(s)
}

func (s FeaturesQuota) GoString() string {
	return s.String()
}

func (s *FeaturesQuota) GetIsEnabled() *bool {
	return s.IsEnabled
}

func (s *FeaturesQuota) SetIsEnabled(v bool) *FeaturesQuota {
	s.IsEnabled = &v
	return s
}

func (s *FeaturesQuota) Validate() error {
	return dara.Validate(s)
}
