// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiagnosisIssue interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DiagnosisIssue
	GetCode() *string
	SetExtra(v map[string]*string) *DiagnosisIssue
	GetExtra() map[string]*string
	SetFirstOccurrence(v string) *DiagnosisIssue
	GetFirstOccurrence() *string
	SetLastOccurrence(v string) *DiagnosisIssue
	GetLastOccurrence() *string
	SetLevel(v string) *DiagnosisIssue
	GetLevel() *string
	SetOccurrenceCount(v int64) *DiagnosisIssue
	GetOccurrenceCount() *int64
	SetSolution(v string) *DiagnosisIssue
	GetSolution() *string
}

type DiagnosisIssue struct {
	// The event code.
	//
	// example:
	//
	// RepoNumOverLimit
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information.
	Extra map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// The first occurrence time of the event.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-01-01T00:00:00+08:00
	FirstOccurrence *string `json:"FirstOccurrence,omitempty" xml:"FirstOccurrence,omitempty"`
	// The last occurrence time of the event.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-01-01T00:30:00+08:00
	LastOccurrence *string `json:"LastOccurrence,omitempty" xml:"LastOccurrence,omitempty"`
	// The event level.
	//
	// example:
	//
	// INFO
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The number of occurrences of the event.
	//
	// example:
	//
	// 10
	OccurrenceCount *int64 `json:"OccurrenceCount,omitempty" xml:"OccurrenceCount,omitempty"`
	// The Solutions.
	//
	// example:
	//
	// Buy repository quota
	Solution *string `json:"Solution,omitempty" xml:"Solution,omitempty"`
}

func (s DiagnosisIssue) String() string {
	return dara.Prettify(s)
}

func (s DiagnosisIssue) GoString() string {
	return s.String()
}

func (s *DiagnosisIssue) GetCode() *string {
	return s.Code
}

func (s *DiagnosisIssue) GetExtra() map[string]*string {
	return s.Extra
}

func (s *DiagnosisIssue) GetFirstOccurrence() *string {
	return s.FirstOccurrence
}

func (s *DiagnosisIssue) GetLastOccurrence() *string {
	return s.LastOccurrence
}

func (s *DiagnosisIssue) GetLevel() *string {
	return s.Level
}

func (s *DiagnosisIssue) GetOccurrenceCount() *int64 {
	return s.OccurrenceCount
}

func (s *DiagnosisIssue) GetSolution() *string {
	return s.Solution
}

func (s *DiagnosisIssue) SetCode(v string) *DiagnosisIssue {
	s.Code = &v
	return s
}

func (s *DiagnosisIssue) SetExtra(v map[string]*string) *DiagnosisIssue {
	s.Extra = v
	return s
}

func (s *DiagnosisIssue) SetFirstOccurrence(v string) *DiagnosisIssue {
	s.FirstOccurrence = &v
	return s
}

func (s *DiagnosisIssue) SetLastOccurrence(v string) *DiagnosisIssue {
	s.LastOccurrence = &v
	return s
}

func (s *DiagnosisIssue) SetLevel(v string) *DiagnosisIssue {
	s.Level = &v
	return s
}

func (s *DiagnosisIssue) SetOccurrenceCount(v int64) *DiagnosisIssue {
	s.OccurrenceCount = &v
	return s
}

func (s *DiagnosisIssue) SetSolution(v string) *DiagnosisIssue {
	s.Solution = &v
	return s
}

func (s *DiagnosisIssue) Validate() error {
	return dara.Validate(s)
}
