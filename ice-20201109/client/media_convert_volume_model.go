// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMediaConvertVolume interface {
	dara.Model
	String() string
	GoString() string
	SetIntegratedLoudnessTarget(v int32) *MediaConvertVolume
	GetIntegratedLoudnessTarget() *int32
	SetLevel(v int32) *MediaConvertVolume
	GetLevel() *int32
	SetLoudnessRangeTarget(v int32) *MediaConvertVolume
	GetLoudnessRangeTarget() *int32
	SetMethod(v string) *MediaConvertVolume
	GetMethod() *string
	SetTruePeak(v int32) *MediaConvertVolume
	GetTruePeak() *int32
}

type MediaConvertVolume struct {
	IntegratedLoudnessTarget *int32  `json:"IntegratedLoudnessTarget,omitempty" xml:"IntegratedLoudnessTarget,omitempty"`
	Level                    *int32  `json:"Level,omitempty" xml:"Level,omitempty"`
	LoudnessRangeTarget      *int32  `json:"LoudnessRangeTarget,omitempty" xml:"LoudnessRangeTarget,omitempty"`
	Method                   *string `json:"Method,omitempty" xml:"Method,omitempty"`
	TruePeak                 *int32  `json:"TruePeak,omitempty" xml:"TruePeak,omitempty"`
}

func (s MediaConvertVolume) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertVolume) GoString() string {
	return s.String()
}

func (s *MediaConvertVolume) GetIntegratedLoudnessTarget() *int32 {
	return s.IntegratedLoudnessTarget
}

func (s *MediaConvertVolume) GetLevel() *int32 {
	return s.Level
}

func (s *MediaConvertVolume) GetLoudnessRangeTarget() *int32 {
	return s.LoudnessRangeTarget
}

func (s *MediaConvertVolume) GetMethod() *string {
	return s.Method
}

func (s *MediaConvertVolume) GetTruePeak() *int32 {
	return s.TruePeak
}

func (s *MediaConvertVolume) SetIntegratedLoudnessTarget(v int32) *MediaConvertVolume {
	s.IntegratedLoudnessTarget = &v
	return s
}

func (s *MediaConvertVolume) SetLevel(v int32) *MediaConvertVolume {
	s.Level = &v
	return s
}

func (s *MediaConvertVolume) SetLoudnessRangeTarget(v int32) *MediaConvertVolume {
	s.LoudnessRangeTarget = &v
	return s
}

func (s *MediaConvertVolume) SetMethod(v string) *MediaConvertVolume {
	s.Method = &v
	return s
}

func (s *MediaConvertVolume) SetTruePeak(v int32) *MediaConvertVolume {
	s.TruePeak = &v
	return s
}

func (s *MediaConvertVolume) Validate() error {
	return dara.Validate(s)
}
