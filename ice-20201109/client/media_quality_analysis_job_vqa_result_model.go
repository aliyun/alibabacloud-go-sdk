// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMediaQualityAnalysisJobVqaResult interface {
	dara.Model
	String() string
	GoString() string
	SetBlock(v *MediaQualityAnalysisJobVqaScoreDetail) *MediaQualityAnalysisJobVqaResult
	GetBlock() *MediaQualityAnalysisJobVqaScoreDetail
	SetColor(v *MediaQualityAnalysisJobVqaScoreDetail) *MediaQualityAnalysisJobVqaResult
	GetColor() *MediaQualityAnalysisJobVqaScoreDetail
	SetDetail(v *MediaQualityAnalysisJobVqaScoreDetail) *MediaQualityAnalysisJobVqaResult
	GetDetail() *MediaQualityAnalysisJobVqaScoreDetail
	SetNoise(v *MediaQualityAnalysisJobVqaScoreDetail) *MediaQualityAnalysisJobVqaResult
	GetNoise() *MediaQualityAnalysisJobVqaScoreDetail
	SetScoreResult(v *MediaQualityAnalysisJobVqaResultScoreResult) *MediaQualityAnalysisJobVqaResult
	GetScoreResult() *MediaQualityAnalysisJobVqaResultScoreResult
	SetSharp(v *MediaQualityAnalysisJobVqaScoreDetail) *MediaQualityAnalysisJobVqaResult
	GetSharp() *MediaQualityAnalysisJobVqaScoreDetail
	SetState(v string) *MediaQualityAnalysisJobVqaResult
	GetState() *string
}

type MediaQualityAnalysisJobVqaResult struct {
	Block       *MediaQualityAnalysisJobVqaScoreDetail       `json:"Block,omitempty" xml:"Block,omitempty"`
	Color       *MediaQualityAnalysisJobVqaScoreDetail       `json:"Color,omitempty" xml:"Color,omitempty"`
	Detail      *MediaQualityAnalysisJobVqaScoreDetail       `json:"Detail,omitempty" xml:"Detail,omitempty"`
	Noise       *MediaQualityAnalysisJobVqaScoreDetail       `json:"Noise,omitempty" xml:"Noise,omitempty"`
	ScoreResult *MediaQualityAnalysisJobVqaResultScoreResult `json:"ScoreResult,omitempty" xml:"ScoreResult,omitempty" type:"Struct"`
	Sharp       *MediaQualityAnalysisJobVqaScoreDetail       `json:"Sharp,omitempty" xml:"Sharp,omitempty"`
	State       *string                                      `json:"State,omitempty" xml:"State,omitempty"`
}

func (s MediaQualityAnalysisJobVqaResult) String() string {
	return dara.Prettify(s)
}

func (s MediaQualityAnalysisJobVqaResult) GoString() string {
	return s.String()
}

func (s *MediaQualityAnalysisJobVqaResult) GetBlock() *MediaQualityAnalysisJobVqaScoreDetail {
	return s.Block
}

func (s *MediaQualityAnalysisJobVqaResult) GetColor() *MediaQualityAnalysisJobVqaScoreDetail {
	return s.Color
}

func (s *MediaQualityAnalysisJobVqaResult) GetDetail() *MediaQualityAnalysisJobVqaScoreDetail {
	return s.Detail
}

func (s *MediaQualityAnalysisJobVqaResult) GetNoise() *MediaQualityAnalysisJobVqaScoreDetail {
	return s.Noise
}

func (s *MediaQualityAnalysisJobVqaResult) GetScoreResult() *MediaQualityAnalysisJobVqaResultScoreResult {
	return s.ScoreResult
}

func (s *MediaQualityAnalysisJobVqaResult) GetSharp() *MediaQualityAnalysisJobVqaScoreDetail {
	return s.Sharp
}

func (s *MediaQualityAnalysisJobVqaResult) GetState() *string {
	return s.State
}

func (s *MediaQualityAnalysisJobVqaResult) SetBlock(v *MediaQualityAnalysisJobVqaScoreDetail) *MediaQualityAnalysisJobVqaResult {
	s.Block = v
	return s
}

func (s *MediaQualityAnalysisJobVqaResult) SetColor(v *MediaQualityAnalysisJobVqaScoreDetail) *MediaQualityAnalysisJobVqaResult {
	s.Color = v
	return s
}

func (s *MediaQualityAnalysisJobVqaResult) SetDetail(v *MediaQualityAnalysisJobVqaScoreDetail) *MediaQualityAnalysisJobVqaResult {
	s.Detail = v
	return s
}

func (s *MediaQualityAnalysisJobVqaResult) SetNoise(v *MediaQualityAnalysisJobVqaScoreDetail) *MediaQualityAnalysisJobVqaResult {
	s.Noise = v
	return s
}

func (s *MediaQualityAnalysisJobVqaResult) SetScoreResult(v *MediaQualityAnalysisJobVqaResultScoreResult) *MediaQualityAnalysisJobVqaResult {
	s.ScoreResult = v
	return s
}

func (s *MediaQualityAnalysisJobVqaResult) SetSharp(v *MediaQualityAnalysisJobVqaScoreDetail) *MediaQualityAnalysisJobVqaResult {
	s.Sharp = v
	return s
}

func (s *MediaQualityAnalysisJobVqaResult) SetState(v string) *MediaQualityAnalysisJobVqaResult {
	s.State = &v
	return s
}

func (s *MediaQualityAnalysisJobVqaResult) Validate() error {
	if s.Block != nil {
		if err := s.Block.Validate(); err != nil {
			return err
		}
	}
	if s.Color != nil {
		if err := s.Color.Validate(); err != nil {
			return err
		}
	}
	if s.Detail != nil {
		if err := s.Detail.Validate(); err != nil {
			return err
		}
	}
	if s.Noise != nil {
		if err := s.Noise.Validate(); err != nil {
			return err
		}
	}
	if s.ScoreResult != nil {
		if err := s.ScoreResult.Validate(); err != nil {
			return err
		}
	}
	if s.Sharp != nil {
		if err := s.Sharp.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MediaQualityAnalysisJobVqaResultScoreResult struct {
	Block  *MediaQualityAnalysisJobVqaResultScoreResultBlock  `json:"Block,omitempty" xml:"Block,omitempty" type:"Struct"`
	Color  *MediaQualityAnalysisJobVqaResultScoreResultColor  `json:"Color,omitempty" xml:"Color,omitempty" type:"Struct"`
	Detail *MediaQualityAnalysisJobVqaResultScoreResultDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Struct"`
	Noise  *MediaQualityAnalysisJobVqaResultScoreResultNoise  `json:"Noise,omitempty" xml:"Noise,omitempty" type:"Struct"`
	Score  *float64                                           `json:"Score,omitempty" xml:"Score,omitempty"`
	Sharp  *MediaQualityAnalysisJobVqaResultScoreResultSharp  `json:"Sharp,omitempty" xml:"Sharp,omitempty" type:"Struct"`
}

func (s MediaQualityAnalysisJobVqaResultScoreResult) String() string {
	return dara.Prettify(s)
}

func (s MediaQualityAnalysisJobVqaResultScoreResult) GoString() string {
	return s.String()
}

func (s *MediaQualityAnalysisJobVqaResultScoreResult) GetBlock() *MediaQualityAnalysisJobVqaResultScoreResultBlock {
	return s.Block
}

func (s *MediaQualityAnalysisJobVqaResultScoreResult) GetColor() *MediaQualityAnalysisJobVqaResultScoreResultColor {
	return s.Color
}

func (s *MediaQualityAnalysisJobVqaResultScoreResult) GetDetail() *MediaQualityAnalysisJobVqaResultScoreResultDetail {
	return s.Detail
}

func (s *MediaQualityAnalysisJobVqaResultScoreResult) GetNoise() *MediaQualityAnalysisJobVqaResultScoreResultNoise {
	return s.Noise
}

func (s *MediaQualityAnalysisJobVqaResultScoreResult) GetScore() *float64 {
	return s.Score
}

func (s *MediaQualityAnalysisJobVqaResultScoreResult) GetSharp() *MediaQualityAnalysisJobVqaResultScoreResultSharp {
	return s.Sharp
}

func (s *MediaQualityAnalysisJobVqaResultScoreResult) SetBlock(v *MediaQualityAnalysisJobVqaResultScoreResultBlock) *MediaQualityAnalysisJobVqaResultScoreResult {
	s.Block = v
	return s
}

func (s *MediaQualityAnalysisJobVqaResultScoreResult) SetColor(v *MediaQualityAnalysisJobVqaResultScoreResultColor) *MediaQualityAnalysisJobVqaResultScoreResult {
	s.Color = v
	return s
}

func (s *MediaQualityAnalysisJobVqaResultScoreResult) SetDetail(v *MediaQualityAnalysisJobVqaResultScoreResultDetail) *MediaQualityAnalysisJobVqaResultScoreResult {
	s.Detail = v
	return s
}

func (s *MediaQualityAnalysisJobVqaResultScoreResult) SetNoise(v *MediaQualityAnalysisJobVqaResultScoreResultNoise) *MediaQualityAnalysisJobVqaResultScoreResult {
	s.Noise = v
	return s
}

func (s *MediaQualityAnalysisJobVqaResultScoreResult) SetScore(v float64) *MediaQualityAnalysisJobVqaResultScoreResult {
	s.Score = &v
	return s
}

func (s *MediaQualityAnalysisJobVqaResultScoreResult) SetSharp(v *MediaQualityAnalysisJobVqaResultScoreResultSharp) *MediaQualityAnalysisJobVqaResultScoreResult {
	s.Sharp = v
	return s
}

func (s *MediaQualityAnalysisJobVqaResultScoreResult) Validate() error {
	if s.Block != nil {
		if err := s.Block.Validate(); err != nil {
			return err
		}
	}
	if s.Color != nil {
		if err := s.Color.Validate(); err != nil {
			return err
		}
	}
	if s.Detail != nil {
		if err := s.Detail.Validate(); err != nil {
			return err
		}
	}
	if s.Noise != nil {
		if err := s.Noise.Validate(); err != nil {
			return err
		}
	}
	if s.Sharp != nil {
		if err := s.Sharp.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MediaQualityAnalysisJobVqaResultScoreResultBlock struct {
	IntensityValue  *float64 `json:"IntensityValue,omitempty" xml:"IntensityValue,omitempty"`
	PerceptualScore *float64 `json:"PerceptualScore,omitempty" xml:"PerceptualScore,omitempty"`
}

func (s MediaQualityAnalysisJobVqaResultScoreResultBlock) String() string {
	return dara.Prettify(s)
}

func (s MediaQualityAnalysisJobVqaResultScoreResultBlock) GoString() string {
	return s.String()
}

func (s *MediaQualityAnalysisJobVqaResultScoreResultBlock) GetIntensityValue() *float64 {
	return s.IntensityValue
}

func (s *MediaQualityAnalysisJobVqaResultScoreResultBlock) GetPerceptualScore() *float64 {
	return s.PerceptualScore
}

func (s *MediaQualityAnalysisJobVqaResultScoreResultBlock) SetIntensityValue(v float64) *MediaQualityAnalysisJobVqaResultScoreResultBlock {
	s.IntensityValue = &v
	return s
}

func (s *MediaQualityAnalysisJobVqaResultScoreResultBlock) SetPerceptualScore(v float64) *MediaQualityAnalysisJobVqaResultScoreResultBlock {
	s.PerceptualScore = &v
	return s
}

func (s *MediaQualityAnalysisJobVqaResultScoreResultBlock) Validate() error {
	return dara.Validate(s)
}

type MediaQualityAnalysisJobVqaResultScoreResultColor struct {
	IntensityValue  *float64 `json:"IntensityValue,omitempty" xml:"IntensityValue,omitempty"`
	PerceptualScore *float64 `json:"PerceptualScore,omitempty" xml:"PerceptualScore,omitempty"`
}

func (s MediaQualityAnalysisJobVqaResultScoreResultColor) String() string {
	return dara.Prettify(s)
}

func (s MediaQualityAnalysisJobVqaResultScoreResultColor) GoString() string {
	return s.String()
}

func (s *MediaQualityAnalysisJobVqaResultScoreResultColor) GetIntensityValue() *float64 {
	return s.IntensityValue
}

func (s *MediaQualityAnalysisJobVqaResultScoreResultColor) GetPerceptualScore() *float64 {
	return s.PerceptualScore
}

func (s *MediaQualityAnalysisJobVqaResultScoreResultColor) SetIntensityValue(v float64) *MediaQualityAnalysisJobVqaResultScoreResultColor {
	s.IntensityValue = &v
	return s
}

func (s *MediaQualityAnalysisJobVqaResultScoreResultColor) SetPerceptualScore(v float64) *MediaQualityAnalysisJobVqaResultScoreResultColor {
	s.PerceptualScore = &v
	return s
}

func (s *MediaQualityAnalysisJobVqaResultScoreResultColor) Validate() error {
	return dara.Validate(s)
}

type MediaQualityAnalysisJobVqaResultScoreResultDetail struct {
	IntensityValue  *float64 `json:"IntensityValue,omitempty" xml:"IntensityValue,omitempty"`
	PerceptualScore *float64 `json:"PerceptualScore,omitempty" xml:"PerceptualScore,omitempty"`
}

func (s MediaQualityAnalysisJobVqaResultScoreResultDetail) String() string {
	return dara.Prettify(s)
}

func (s MediaQualityAnalysisJobVqaResultScoreResultDetail) GoString() string {
	return s.String()
}

func (s *MediaQualityAnalysisJobVqaResultScoreResultDetail) GetIntensityValue() *float64 {
	return s.IntensityValue
}

func (s *MediaQualityAnalysisJobVqaResultScoreResultDetail) GetPerceptualScore() *float64 {
	return s.PerceptualScore
}

func (s *MediaQualityAnalysisJobVqaResultScoreResultDetail) SetIntensityValue(v float64) *MediaQualityAnalysisJobVqaResultScoreResultDetail {
	s.IntensityValue = &v
	return s
}

func (s *MediaQualityAnalysisJobVqaResultScoreResultDetail) SetPerceptualScore(v float64) *MediaQualityAnalysisJobVqaResultScoreResultDetail {
	s.PerceptualScore = &v
	return s
}

func (s *MediaQualityAnalysisJobVqaResultScoreResultDetail) Validate() error {
	return dara.Validate(s)
}

type MediaQualityAnalysisJobVqaResultScoreResultNoise struct {
	IntensityValue  *float64 `json:"IntensityValue,omitempty" xml:"IntensityValue,omitempty"`
	PerceptualScore *float64 `json:"PerceptualScore,omitempty" xml:"PerceptualScore,omitempty"`
}

func (s MediaQualityAnalysisJobVqaResultScoreResultNoise) String() string {
	return dara.Prettify(s)
}

func (s MediaQualityAnalysisJobVqaResultScoreResultNoise) GoString() string {
	return s.String()
}

func (s *MediaQualityAnalysisJobVqaResultScoreResultNoise) GetIntensityValue() *float64 {
	return s.IntensityValue
}

func (s *MediaQualityAnalysisJobVqaResultScoreResultNoise) GetPerceptualScore() *float64 {
	return s.PerceptualScore
}

func (s *MediaQualityAnalysisJobVqaResultScoreResultNoise) SetIntensityValue(v float64) *MediaQualityAnalysisJobVqaResultScoreResultNoise {
	s.IntensityValue = &v
	return s
}

func (s *MediaQualityAnalysisJobVqaResultScoreResultNoise) SetPerceptualScore(v float64) *MediaQualityAnalysisJobVqaResultScoreResultNoise {
	s.PerceptualScore = &v
	return s
}

func (s *MediaQualityAnalysisJobVqaResultScoreResultNoise) Validate() error {
	return dara.Validate(s)
}

type MediaQualityAnalysisJobVqaResultScoreResultSharp struct {
	IntensityValue  *float64 `json:"IntensityValue,omitempty" xml:"IntensityValue,omitempty"`
	PerceptualScore *float64 `json:"PerceptualScore,omitempty" xml:"PerceptualScore,omitempty"`
}

func (s MediaQualityAnalysisJobVqaResultScoreResultSharp) String() string {
	return dara.Prettify(s)
}

func (s MediaQualityAnalysisJobVqaResultScoreResultSharp) GoString() string {
	return s.String()
}

func (s *MediaQualityAnalysisJobVqaResultScoreResultSharp) GetIntensityValue() *float64 {
	return s.IntensityValue
}

func (s *MediaQualityAnalysisJobVqaResultScoreResultSharp) GetPerceptualScore() *float64 {
	return s.PerceptualScore
}

func (s *MediaQualityAnalysisJobVqaResultScoreResultSharp) SetIntensityValue(v float64) *MediaQualityAnalysisJobVqaResultScoreResultSharp {
	s.IntensityValue = &v
	return s
}

func (s *MediaQualityAnalysisJobVqaResultScoreResultSharp) SetPerceptualScore(v float64) *MediaQualityAnalysisJobVqaResultScoreResultSharp {
	s.PerceptualScore = &v
	return s
}

func (s *MediaQualityAnalysisJobVqaResultScoreResultSharp) Validate() error {
	return dara.Validate(s)
}
