// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSanityCheckResultItem interface {
	dara.Model
	String() string
	GoString() string
	SetCheckNumber(v int32) *SanityCheckResultItem
	GetCheckNumber() *int32
	SetFinishedAt(v string) *SanityCheckResultItem
	GetFinishedAt() *string
	SetMessage(v string) *SanityCheckResultItem
	GetMessage() *string
	SetPhase(v string) *SanityCheckResultItem
	GetPhase() *string
	SetStartedAt(v string) *SanityCheckResultItem
	GetStartedAt() *string
	SetStatus(v string) *SanityCheckResultItem
	GetStatus() *string
}

type SanityCheckResultItem struct {
	// example:
	//
	// 1
	CheckNumber *int32 `json:"CheckNumber,omitempty" xml:"CheckNumber,omitempty"`
	// example:
	//
	// ”2023-11-30T16:47:30.378817+08:00"
	FinishedAt *string `json:"FinishedAt,omitempty" xml:"FinishedAt,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// CheckInit
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// example:
	//
	// ”2023-11-30T16:47:30.343005+08:00“
	StartedAt *string `json:"StartedAt,omitempty" xml:"StartedAt,omitempty"`
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SanityCheckResultItem) String() string {
	return dara.Prettify(s)
}

func (s SanityCheckResultItem) GoString() string {
	return s.String()
}

func (s *SanityCheckResultItem) GetCheckNumber() *int32 {
	return s.CheckNumber
}

func (s *SanityCheckResultItem) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *SanityCheckResultItem) GetMessage() *string {
	return s.Message
}

func (s *SanityCheckResultItem) GetPhase() *string {
	return s.Phase
}

func (s *SanityCheckResultItem) GetStartedAt() *string {
	return s.StartedAt
}

func (s *SanityCheckResultItem) GetStatus() *string {
	return s.Status
}

func (s *SanityCheckResultItem) SetCheckNumber(v int32) *SanityCheckResultItem {
	s.CheckNumber = &v
	return s
}

func (s *SanityCheckResultItem) SetFinishedAt(v string) *SanityCheckResultItem {
	s.FinishedAt = &v
	return s
}

func (s *SanityCheckResultItem) SetMessage(v string) *SanityCheckResultItem {
	s.Message = &v
	return s
}

func (s *SanityCheckResultItem) SetPhase(v string) *SanityCheckResultItem {
	s.Phase = &v
	return s
}

func (s *SanityCheckResultItem) SetStartedAt(v string) *SanityCheckResultItem {
	s.StartedAt = &v
	return s
}

func (s *SanityCheckResultItem) SetStatus(v string) *SanityCheckResultItem {
	s.Status = &v
	return s
}

func (s *SanityCheckResultItem) Validate() error {
	return dara.Validate(s)
}
