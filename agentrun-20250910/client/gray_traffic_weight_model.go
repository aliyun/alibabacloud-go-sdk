// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrayTrafficWeight interface {
	dara.Model
	String() string
	GoString() string
	SetVersion(v string) *GrayTrafficWeight
	GetVersion() *string
	SetWeight(v float32) *GrayTrafficWeight
	GetWeight() *float32
}

type GrayTrafficWeight struct {
	// 灰度版本号
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// 流量权重比例（0.0-1.0）
	Weight *float32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s GrayTrafficWeight) String() string {
	return dara.Prettify(s)
}

func (s GrayTrafficWeight) GoString() string {
	return s.String()
}

func (s *GrayTrafficWeight) GetVersion() *string {
	return s.Version
}

func (s *GrayTrafficWeight) GetWeight() *float32 {
	return s.Weight
}

func (s *GrayTrafficWeight) SetVersion(v string) *GrayTrafficWeight {
	s.Version = &v
	return s
}

func (s *GrayTrafficWeight) SetWeight(v float32) *GrayTrafficWeight {
	s.Weight = &v
	return s
}

func (s *GrayTrafficWeight) Validate() error {
	return dara.Validate(s)
}
