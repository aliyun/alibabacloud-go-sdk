// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiagnosisTarget interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DiagnosisTarget
	GetEndTime() *string
	SetExtra(v map[string]*string) *DiagnosisTarget
	GetExtra() map[string]*string
	SetNamespace(v string) *DiagnosisTarget
	GetNamespace() *string
	SetRelatedId(v string) *DiagnosisTarget
	GetRelatedId() *string
	SetRepository(v string) *DiagnosisTarget
	GetRepository() *string
	SetStartTime(v string) *DiagnosisTarget
	GetStartTime() *string
	SetTag(v string) *DiagnosisTarget
	GetTag() *string
}

type DiagnosisTarget struct {
	// The end of the diagnostic time window, in ISO 8601 format. Must be after `StartTime`.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-01-01T00:30:00+08:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// A map of key-value pairs providing additional context for the diagnosis.
	Extra map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// The namespace that contains the repository.
	//
	// example:
	//
	// test_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The ID of a related operation, such as a previous diagnostic task.
	RelatedId *string `json:"RelatedId,omitempty" xml:"RelatedId,omitempty"`
	// The name of the repository.
	//
	// example:
	//
	// test_repo
	Repository *string `json:"Repository,omitempty" xml:"Repository,omitempty"`
	// The start of the diagnostic time window, in ISO 8601 format.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-01-01T00:00:00+08:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The container image tag.
	//
	// example:
	//
	// latest
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s DiagnosisTarget) String() string {
	return dara.Prettify(s)
}

func (s DiagnosisTarget) GoString() string {
	return s.String()
}

func (s *DiagnosisTarget) GetEndTime() *string {
	return s.EndTime
}

func (s *DiagnosisTarget) GetExtra() map[string]*string {
	return s.Extra
}

func (s *DiagnosisTarget) GetNamespace() *string {
	return s.Namespace
}

func (s *DiagnosisTarget) GetRelatedId() *string {
	return s.RelatedId
}

func (s *DiagnosisTarget) GetRepository() *string {
	return s.Repository
}

func (s *DiagnosisTarget) GetStartTime() *string {
	return s.StartTime
}

func (s *DiagnosisTarget) GetTag() *string {
	return s.Tag
}

func (s *DiagnosisTarget) SetEndTime(v string) *DiagnosisTarget {
	s.EndTime = &v
	return s
}

func (s *DiagnosisTarget) SetExtra(v map[string]*string) *DiagnosisTarget {
	s.Extra = v
	return s
}

func (s *DiagnosisTarget) SetNamespace(v string) *DiagnosisTarget {
	s.Namespace = &v
	return s
}

func (s *DiagnosisTarget) SetRelatedId(v string) *DiagnosisTarget {
	s.RelatedId = &v
	return s
}

func (s *DiagnosisTarget) SetRepository(v string) *DiagnosisTarget {
	s.Repository = &v
	return s
}

func (s *DiagnosisTarget) SetStartTime(v string) *DiagnosisTarget {
	s.StartTime = &v
	return s
}

func (s *DiagnosisTarget) SetTag(v string) *DiagnosisTarget {
	s.Tag = &v
	return s
}

func (s *DiagnosisTarget) Validate() error {
	return dara.Validate(s)
}
