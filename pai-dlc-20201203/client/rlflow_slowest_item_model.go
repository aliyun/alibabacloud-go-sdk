// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRLFlowSlowestItem interface {
	dara.Model
	String() string
	GoString() string
	SetPromptUid(v string) *RLFlowSlowestItem
	GetPromptUid() *string
	SetSampleIndex(v string) *RLFlowSlowestItem
	GetSampleIndex() *string
	SetSec(v float64) *RLFlowSlowestItem
	GetSec() *float64
}

type RLFlowSlowestItem struct {
	// The UID of the sample.
	//
	// example:
	//
	// 321fa56f-e1e5-4eb3-8047-db7a230c9a75
	PromptUid *string `json:"PromptUid,omitempty" xml:"PromptUid,omitempty"`
	// The ordinal number of the event trace.
	//
	// example:
	//
	// 2
	SampleIndex *string `json:"SampleIndex,omitempty" xml:"SampleIndex,omitempty"`
	// The execution duration of the stage, in seconds.
	//
	// example:
	//
	// 9.2
	Sec *float64 `json:"Sec,omitempty" xml:"Sec,omitempty"`
}

func (s RLFlowSlowestItem) String() string {
	return dara.Prettify(s)
}

func (s RLFlowSlowestItem) GoString() string {
	return s.String()
}

func (s *RLFlowSlowestItem) GetPromptUid() *string {
	return s.PromptUid
}

func (s *RLFlowSlowestItem) GetSampleIndex() *string {
	return s.SampleIndex
}

func (s *RLFlowSlowestItem) GetSec() *float64 {
	return s.Sec
}

func (s *RLFlowSlowestItem) SetPromptUid(v string) *RLFlowSlowestItem {
	s.PromptUid = &v
	return s
}

func (s *RLFlowSlowestItem) SetSampleIndex(v string) *RLFlowSlowestItem {
	s.SampleIndex = &v
	return s
}

func (s *RLFlowSlowestItem) SetSec(v float64) *RLFlowSlowestItem {
	s.Sec = &v
	return s
}

func (s *RLFlowSlowestItem) Validate() error {
	return dara.Validate(s)
}
