// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveAfterAnswerDelayPlaybackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAfterAnswerDelayPlayback(v int32) *SaveAfterAnswerDelayPlaybackRequest
	GetAfterAnswerDelayPlayback() *int32
	SetEntryId(v string) *SaveAfterAnswerDelayPlaybackRequest
	GetEntryId() *string
	SetStrategyLevel(v int32) *SaveAfterAnswerDelayPlaybackRequest
	GetStrategyLevel() *int32
}

type SaveAfterAnswerDelayPlaybackRequest struct {
	// Playback time for delayed playback. The default value is 0 if this parameter is not specified.
	//
	// example:
	//
	// 1
	AfterAnswerDelayPlayback *int32 `json:"AfterAnswerDelayPlayback,omitempty" xml:"AfterAnswerDelayPlayback,omitempty"`
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 980116ec-2d3d-4747-8059-dc25e7af8501
	EntryId *string `json:"EntryId,omitempty" xml:"EntryId,omitempty"`
	// Policy level (Required).
	//
	// - 2: instance
	//
	// example:
	//
	// 2
	StrategyLevel *int32 `json:"StrategyLevel,omitempty" xml:"StrategyLevel,omitempty"`
}

func (s SaveAfterAnswerDelayPlaybackRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveAfterAnswerDelayPlaybackRequest) GoString() string {
	return s.String()
}

func (s *SaveAfterAnswerDelayPlaybackRequest) GetAfterAnswerDelayPlayback() *int32 {
	return s.AfterAnswerDelayPlayback
}

func (s *SaveAfterAnswerDelayPlaybackRequest) GetEntryId() *string {
	return s.EntryId
}

func (s *SaveAfterAnswerDelayPlaybackRequest) GetStrategyLevel() *int32 {
	return s.StrategyLevel
}

func (s *SaveAfterAnswerDelayPlaybackRequest) SetAfterAnswerDelayPlayback(v int32) *SaveAfterAnswerDelayPlaybackRequest {
	s.AfterAnswerDelayPlayback = &v
	return s
}

func (s *SaveAfterAnswerDelayPlaybackRequest) SetEntryId(v string) *SaveAfterAnswerDelayPlaybackRequest {
	s.EntryId = &v
	return s
}

func (s *SaveAfterAnswerDelayPlaybackRequest) SetStrategyLevel(v int32) *SaveAfterAnswerDelayPlaybackRequest {
	s.StrategyLevel = &v
	return s
}

func (s *SaveAfterAnswerDelayPlaybackRequest) Validate() error {
	return dara.Validate(s)
}
