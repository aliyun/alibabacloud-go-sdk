// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTimeMs(v int64) *ListInstancesShrinkRequest
	GetEndTimeMs() *int64
	SetInstanceIdsShrink(v string) *ListInstancesShrinkRequest
	GetInstanceIdsShrink() *string
	SetInstanceStatusShrink(v string) *ListInstancesShrinkRequest
	GetInstanceStatusShrink() *string
	SetLimit(v string) *ListInstancesShrinkRequest
	GetLimit() *string
	SetQualifier(v string) *ListInstancesShrinkRequest
	GetQualifier() *string
	SetStartKey(v string) *ListInstancesShrinkRequest
	GetStartKey() *string
	SetStartTimeMs(v int64) *ListInstancesShrinkRequest
	GetStartTimeMs() *int64
	SetWithAllActive(v bool) *ListInstancesShrinkRequest
	GetWithAllActive() *bool
}

type ListInstancesShrinkRequest struct {
	EndTimeMs            *int64  `json:"endTimeMs,omitempty" xml:"endTimeMs,omitempty"`
	InstanceIdsShrink    *string `json:"instanceIds,omitempty" xml:"instanceIds,omitempty"`
	InstanceStatusShrink *string `json:"instanceStatus,omitempty" xml:"instanceStatus,omitempty"`
	Limit                *string `json:"limit,omitempty" xml:"limit,omitempty"`
	// The function version or alias.
	//
	// example:
	//
	// LATEST
	Qualifier   *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	StartKey    *string `json:"startKey,omitempty" xml:"startKey,omitempty"`
	StartTimeMs *int64  `json:"startTimeMs,omitempty" xml:"startTimeMs,omitempty"`
	// Specifies whether to list all instances. Valid values: true and false.
	//
	// example:
	//
	// true
	WithAllActive *bool `json:"withAllActive,omitempty" xml:"withAllActive,omitempty"`
}

func (s ListInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesShrinkRequest) GetEndTimeMs() *int64 {
	return s.EndTimeMs
}

func (s *ListInstancesShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *ListInstancesShrinkRequest) GetInstanceStatusShrink() *string {
	return s.InstanceStatusShrink
}

func (s *ListInstancesShrinkRequest) GetLimit() *string {
	return s.Limit
}

func (s *ListInstancesShrinkRequest) GetQualifier() *string {
	return s.Qualifier
}

func (s *ListInstancesShrinkRequest) GetStartKey() *string {
	return s.StartKey
}

func (s *ListInstancesShrinkRequest) GetStartTimeMs() *int64 {
	return s.StartTimeMs
}

func (s *ListInstancesShrinkRequest) GetWithAllActive() *bool {
	return s.WithAllActive
}

func (s *ListInstancesShrinkRequest) SetEndTimeMs(v int64) *ListInstancesShrinkRequest {
	s.EndTimeMs = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetInstanceIdsShrink(v string) *ListInstancesShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetInstanceStatusShrink(v string) *ListInstancesShrinkRequest {
	s.InstanceStatusShrink = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetLimit(v string) *ListInstancesShrinkRequest {
	s.Limit = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetQualifier(v string) *ListInstancesShrinkRequest {
	s.Qualifier = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetStartKey(v string) *ListInstancesShrinkRequest {
	s.StartKey = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetStartTimeMs(v int64) *ListInstancesShrinkRequest {
	s.StartTimeMs = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetWithAllActive(v bool) *ListInstancesShrinkRequest {
	s.WithAllActive = &v
	return s
}

func (s *ListInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
