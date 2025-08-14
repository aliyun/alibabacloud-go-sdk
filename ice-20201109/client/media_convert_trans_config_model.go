// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMediaConvertTransConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAdjDarMethod(v string) *MediaConvertTransConfig
	GetAdjDarMethod() *string
	SetIsCheckAudioBitrate(v bool) *MediaConvertTransConfig
	GetIsCheckAudioBitrate() *bool
	SetIsCheckAudioBitrateFail(v bool) *MediaConvertTransConfig
	GetIsCheckAudioBitrateFail() *bool
	SetIsCheckReso(v bool) *MediaConvertTransConfig
	GetIsCheckReso() *bool
	SetIsCheckResoFail(v bool) *MediaConvertTransConfig
	GetIsCheckResoFail() *bool
	SetIsCheckVideoBitrate(v bool) *MediaConvertTransConfig
	GetIsCheckVideoBitrate() *bool
	SetIsCheckVideoBitrateFail(v bool) *MediaConvertTransConfig
	GetIsCheckVideoBitrateFail() *bool
	SetTransMode(v string) *MediaConvertTransConfig
	GetTransMode() *string
}

type MediaConvertTransConfig struct {
	AdjDarMethod            *string `json:"AdjDarMethod,omitempty" xml:"AdjDarMethod,omitempty"`
	IsCheckAudioBitrate     *bool   `json:"IsCheckAudioBitrate,omitempty" xml:"IsCheckAudioBitrate,omitempty"`
	IsCheckAudioBitrateFail *bool   `json:"IsCheckAudioBitrateFail,omitempty" xml:"IsCheckAudioBitrateFail,omitempty"`
	IsCheckReso             *bool   `json:"IsCheckReso,omitempty" xml:"IsCheckReso,omitempty"`
	IsCheckResoFail         *bool   `json:"IsCheckResoFail,omitempty" xml:"IsCheckResoFail,omitempty"`
	IsCheckVideoBitrate     *bool   `json:"IsCheckVideoBitrate,omitempty" xml:"IsCheckVideoBitrate,omitempty"`
	IsCheckVideoBitrateFail *bool   `json:"IsCheckVideoBitrateFail,omitempty" xml:"IsCheckVideoBitrateFail,omitempty"`
	TransMode               *string `json:"TransMode,omitempty" xml:"TransMode,omitempty"`
}

func (s MediaConvertTransConfig) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertTransConfig) GoString() string {
	return s.String()
}

func (s *MediaConvertTransConfig) GetAdjDarMethod() *string {
	return s.AdjDarMethod
}

func (s *MediaConvertTransConfig) GetIsCheckAudioBitrate() *bool {
	return s.IsCheckAudioBitrate
}

func (s *MediaConvertTransConfig) GetIsCheckAudioBitrateFail() *bool {
	return s.IsCheckAudioBitrateFail
}

func (s *MediaConvertTransConfig) GetIsCheckReso() *bool {
	return s.IsCheckReso
}

func (s *MediaConvertTransConfig) GetIsCheckResoFail() *bool {
	return s.IsCheckResoFail
}

func (s *MediaConvertTransConfig) GetIsCheckVideoBitrate() *bool {
	return s.IsCheckVideoBitrate
}

func (s *MediaConvertTransConfig) GetIsCheckVideoBitrateFail() *bool {
	return s.IsCheckVideoBitrateFail
}

func (s *MediaConvertTransConfig) GetTransMode() *string {
	return s.TransMode
}

func (s *MediaConvertTransConfig) SetAdjDarMethod(v string) *MediaConvertTransConfig {
	s.AdjDarMethod = &v
	return s
}

func (s *MediaConvertTransConfig) SetIsCheckAudioBitrate(v bool) *MediaConvertTransConfig {
	s.IsCheckAudioBitrate = &v
	return s
}

func (s *MediaConvertTransConfig) SetIsCheckAudioBitrateFail(v bool) *MediaConvertTransConfig {
	s.IsCheckAudioBitrateFail = &v
	return s
}

func (s *MediaConvertTransConfig) SetIsCheckReso(v bool) *MediaConvertTransConfig {
	s.IsCheckReso = &v
	return s
}

func (s *MediaConvertTransConfig) SetIsCheckResoFail(v bool) *MediaConvertTransConfig {
	s.IsCheckResoFail = &v
	return s
}

func (s *MediaConvertTransConfig) SetIsCheckVideoBitrate(v bool) *MediaConvertTransConfig {
	s.IsCheckVideoBitrate = &v
	return s
}

func (s *MediaConvertTransConfig) SetIsCheckVideoBitrateFail(v bool) *MediaConvertTransConfig {
	s.IsCheckVideoBitrateFail = &v
	return s
}

func (s *MediaConvertTransConfig) SetTransMode(v string) *MediaConvertTransConfig {
	s.TransMode = &v
	return s
}

func (s *MediaConvertTransConfig) Validate() error {
	return dara.Validate(s)
}
