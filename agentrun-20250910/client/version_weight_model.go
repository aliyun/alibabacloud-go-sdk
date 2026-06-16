// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVersionWeight interface {
	dara.Model
	String() string
	GoString() string
	SetVersion(v string) *VersionWeight
	GetVersion() *string
	SetWeight(v float32) *VersionWeight
	GetWeight() *float32
}

type VersionWeight struct {
	// Agent runtime version number
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// Traffic weight ratio (0.0–1.0)
	Weight *float32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s VersionWeight) String() string {
	return dara.Prettify(s)
}

func (s VersionWeight) GoString() string {
	return s.String()
}

func (s *VersionWeight) GetVersion() *string {
	return s.Version
}

func (s *VersionWeight) GetWeight() *float32 {
	return s.Weight
}

func (s *VersionWeight) SetVersion(v string) *VersionWeight {
	s.Version = &v
	return s
}

func (s *VersionWeight) SetWeight(v float32) *VersionWeight {
	s.Weight = &v
	return s
}

func (s *VersionWeight) Validate() error {
	return dara.Validate(s)
}
