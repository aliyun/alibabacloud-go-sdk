// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRiskSceneConfig interface {
	dara.Model
	String() string
	GoString() string
	SetDetectChannel(v []*string) *RiskSceneConfig
	GetDetectChannel() []*string
	SetOfficeChannel(v []*string) *RiskSceneConfig
	GetOfficeChannel() []*string
}

type RiskSceneConfig struct {
	DetectChannel []*string `json:"DetectChannel,omitempty" xml:"DetectChannel,omitempty" type:"Repeated"`
	OfficeChannel []*string `json:"OfficeChannel,omitempty" xml:"OfficeChannel,omitempty" type:"Repeated"`
}

func (s RiskSceneConfig) String() string {
	return dara.Prettify(s)
}

func (s RiskSceneConfig) GoString() string {
	return s.String()
}

func (s *RiskSceneConfig) GetDetectChannel() []*string {
	return s.DetectChannel
}

func (s *RiskSceneConfig) GetOfficeChannel() []*string {
	return s.OfficeChannel
}

func (s *RiskSceneConfig) SetDetectChannel(v []*string) *RiskSceneConfig {
	s.DetectChannel = v
	return s
}

func (s *RiskSceneConfig) SetOfficeChannel(v []*string) *RiskSceneConfig {
	s.OfficeChannel = v
	return s
}

func (s *RiskSceneConfig) Validate() error {
	return dara.Validate(s)
}
