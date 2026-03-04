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
	// The output volume.
	//
	// 	- This parameter takes effect only if Method is set to dynamic.
	//
	// 	- Unit: dB.
	//
	// 	- Valid values: [-70,-5].
	//
	// 	- Default value: -6.
	//
	// example:
	//
	// -6
	IntegratedLoudnessTarget *int32 `json:"IntegratedLoudnessTarget,omitempty" xml:"IntegratedLoudnessTarget,omitempty"`
	// The amount of gain to apply, relative to the input audio.
	//
	// 	- This parameter takes effect only if Method is set to linear.
	//
	// 	- Unit: dB.
	//
	// 	- Valid values: less than or equal to 20.
	//
	// 	- Default value: -20.
	//
	// example:
	//
	// -20
	Level *int32 `json:"Level,omitempty" xml:"Level,omitempty"`
	// The target loudness range.
	//
	// 	- This parameter takes effect only if Method is set to dynamic.
	//
	// 	- Unit: dB.
	//
	// 	- Valid values: [1,20].
	//
	// 	- Default value: 8.
	//
	// example:
	//
	// 8
	LoudnessRangeTarget *int32 `json:"LoudnessRangeTarget,omitempty" xml:"LoudnessRangeTarget,omitempty"`
	// The volume adjustment method. Valid values:
	//
	// 	- auto
	//
	// 	- dynamic
	//
	// 	- linear
	//
	// 	- Default value: dynamic.
	//
	// <!---->
	//
	// *
	//
	// *
	//
	// *
	//
	// *
	//
	// example:
	//
	// linear
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The maximum volume.
	//
	// 	- This parameter takes effect only if Method is set to dynamic.
	//
	// 	- Unit: dB.
	//
	// 	- Valid values: [-9,0].
	//
	// 	- Default value: -1.
	//
	// example:
	//
	// -1
	TruePeak *int32 `json:"TruePeak,omitempty" xml:"TruePeak,omitempty"`
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
