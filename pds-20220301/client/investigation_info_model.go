// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvestigationInfo interface {
	dara.Model
	String() string
	GoString() string
	SetStatus(v int64) *InvestigationInfo
	GetStatus() *int64
	SetSuggestion(v string) *InvestigationInfo
	GetSuggestion() *string
	SetVideoDetail(v *InvestigationInfoVideoDetail) *InvestigationInfo
	GetVideoDetail() *InvestigationInfoVideoDetail
}

type InvestigationInfo struct {
	// The status of the review.
	//
	// Valid values:
	//
	// 	- 0: The review is not performed.
	//
	// 	- 1: The review is not supported.
	//
	// 	- 2: The review fails.
	//
	// 	- 3: The review is in progress.
	//
	// 	- 4: The review is complete.
	//
	// 	- 5: Penalty methods are applied.
	//
	// example:
	//
	// 4
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
	// The recommended operation provided by the review.
	//
	// Valid values:
	//
	// 	- pass: The review is passed..
	//
	// 	- block: The review is not passed. It is recommended to limit the use of the image.
	//
	// example:
	//
	// block
	Suggestion *string `json:"suggestion,omitempty" xml:"suggestion,omitempty"`
	// Video review information
	VideoDetail *InvestigationInfoVideoDetail `json:"video_detail,omitempty" xml:"video_detail,omitempty" type:"Struct"`
}

func (s InvestigationInfo) String() string {
	return dara.Prettify(s)
}

func (s InvestigationInfo) GoString() string {
	return s.String()
}

func (s *InvestigationInfo) GetStatus() *int64 {
	return s.Status
}

func (s *InvestigationInfo) GetSuggestion() *string {
	return s.Suggestion
}

func (s *InvestigationInfo) GetVideoDetail() *InvestigationInfoVideoDetail {
	return s.VideoDetail
}

func (s *InvestigationInfo) SetStatus(v int64) *InvestigationInfo {
	s.Status = &v
	return s
}

func (s *InvestigationInfo) SetSuggestion(v string) *InvestigationInfo {
	s.Suggestion = &v
	return s
}

func (s *InvestigationInfo) SetVideoDetail(v *InvestigationInfoVideoDetail) *InvestigationInfo {
	s.VideoDetail = v
	return s
}

func (s *InvestigationInfo) Validate() error {
	if s.VideoDetail != nil {
		if err := s.VideoDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InvestigationInfoVideoDetail struct {
	// Violation frame information
	BlockFrames []*InvestigationInfoVideoDetailBlockFrames `json:"block_frames,omitempty" xml:"block_frames,omitempty" type:"Repeated"`
}

func (s InvestigationInfoVideoDetail) String() string {
	return dara.Prettify(s)
}

func (s InvestigationInfoVideoDetail) GoString() string {
	return s.String()
}

func (s *InvestigationInfoVideoDetail) GetBlockFrames() []*InvestigationInfoVideoDetailBlockFrames {
	return s.BlockFrames
}

func (s *InvestigationInfoVideoDetail) SetBlockFrames(v []*InvestigationInfoVideoDetailBlockFrames) *InvestigationInfoVideoDetail {
	s.BlockFrames = v
	return s
}

func (s *InvestigationInfoVideoDetail) Validate() error {
	if s.BlockFrames != nil {
		for _, item := range s.BlockFrames {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InvestigationInfoVideoDetailBlockFrames struct {
	// Category of review results
	//
	// example:
	//
	// porn
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// Time (in seconds)
	//
	// example:
	//
	// 3
	Offset *int64 `json:"offset,omitempty" xml:"offset,omitempty"`
	// The confidence level. Valid values: 0 to 100.
	//
	// example:
	//
	// 99.1
	Rate *float64 `json:"rate,omitempty" xml:"rate,omitempty"`
}

func (s InvestigationInfoVideoDetailBlockFrames) String() string {
	return dara.Prettify(s)
}

func (s InvestigationInfoVideoDetailBlockFrames) GoString() string {
	return s.String()
}

func (s *InvestigationInfoVideoDetailBlockFrames) GetLabel() *string {
	return s.Label
}

func (s *InvestigationInfoVideoDetailBlockFrames) GetOffset() *int64 {
	return s.Offset
}

func (s *InvestigationInfoVideoDetailBlockFrames) GetRate() *float64 {
	return s.Rate
}

func (s *InvestigationInfoVideoDetailBlockFrames) SetLabel(v string) *InvestigationInfoVideoDetailBlockFrames {
	s.Label = &v
	return s
}

func (s *InvestigationInfoVideoDetailBlockFrames) SetOffset(v int64) *InvestigationInfoVideoDetailBlockFrames {
	s.Offset = &v
	return s
}

func (s *InvestigationInfoVideoDetailBlockFrames) SetRate(v float64) *InvestigationInfoVideoDetailBlockFrames {
	s.Rate = &v
	return s
}

func (s *InvestigationInfoVideoDetailBlockFrames) Validate() error {
	return dara.Validate(s)
}
