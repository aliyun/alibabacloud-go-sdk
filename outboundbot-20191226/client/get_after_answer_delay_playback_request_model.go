// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAfterAnswerDelayPlaybackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntryId(v string) *GetAfterAnswerDelayPlaybackRequest
	GetEntryId() *string
	SetStrategyLevel(v int32) *GetAfterAnswerDelayPlaybackRequest
	GetStrategyLevel() *int32
}

type GetAfterAnswerDelayPlaybackRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c46001bc-3ead-4bfd-9a69-4b5b66a4a3f4
	EntryId *string `json:"EntryId,omitempty" xml:"EntryId,omitempty"`
	// example:
	//
	// 2
	StrategyLevel *int32 `json:"StrategyLevel,omitempty" xml:"StrategyLevel,omitempty"`
}

func (s GetAfterAnswerDelayPlaybackRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAfterAnswerDelayPlaybackRequest) GoString() string {
	return s.String()
}

func (s *GetAfterAnswerDelayPlaybackRequest) GetEntryId() *string {
	return s.EntryId
}

func (s *GetAfterAnswerDelayPlaybackRequest) GetStrategyLevel() *int32 {
	return s.StrategyLevel
}

func (s *GetAfterAnswerDelayPlaybackRequest) SetEntryId(v string) *GetAfterAnswerDelayPlaybackRequest {
	s.EntryId = &v
	return s
}

func (s *GetAfterAnswerDelayPlaybackRequest) SetStrategyLevel(v int32) *GetAfterAnswerDelayPlaybackRequest {
	s.StrategyLevel = &v
	return s
}

func (s *GetAfterAnswerDelayPlaybackRequest) Validate() error {
	return dara.Validate(s)
}
