// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iArtifactLifecyclePolicy interface {
	dara.Model
	String() string
	GoString() string
	SetCondition(v *ArtifactLifecyclePolicyCondition) *ArtifactLifecyclePolicy
	GetCondition() *ArtifactLifecyclePolicyCondition
	SetFilter(v *ArtifactLifecyclePolicyFilter) *ArtifactLifecyclePolicy
	GetFilter() *ArtifactLifecyclePolicyFilter
	SetType(v string) *ArtifactLifecyclePolicy
	GetType() *string
}

type ArtifactLifecyclePolicy struct {
	// The retention conditions for artifact versions.
	Condition *ArtifactLifecyclePolicyCondition `json:"Condition,omitempty" xml:"Condition,omitempty" type:"Struct"`
	// The filter conditions for artifact versions.
	Filter *ArtifactLifecyclePolicyFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	// The type of the artifact version data retention policy.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ArtifactLifecyclePolicy) String() string {
	return dara.Prettify(s)
}

func (s ArtifactLifecyclePolicy) GoString() string {
	return s.String()
}

func (s *ArtifactLifecyclePolicy) GetCondition() *ArtifactLifecyclePolicyCondition {
	return s.Condition
}

func (s *ArtifactLifecyclePolicy) GetFilter() *ArtifactLifecyclePolicyFilter {
	return s.Filter
}

func (s *ArtifactLifecyclePolicy) GetType() *string {
	return s.Type
}

func (s *ArtifactLifecyclePolicy) SetCondition(v *ArtifactLifecyclePolicyCondition) *ArtifactLifecyclePolicy {
	s.Condition = v
	return s
}

func (s *ArtifactLifecyclePolicy) SetFilter(v *ArtifactLifecyclePolicyFilter) *ArtifactLifecyclePolicy {
	s.Filter = v
	return s
}

func (s *ArtifactLifecyclePolicy) SetType(v string) *ArtifactLifecyclePolicy {
	s.Type = &v
	return s
}

func (s *ArtifactLifecyclePolicy) Validate() error {
	if s.Condition != nil {
		if err := s.Condition.Validate(); err != nil {
			return err
		}
	}
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ArtifactLifecyclePolicyCondition struct {
	// Retains artifact versions whose most recent pull time is within the specified number of days.
	//
	// example:
	//
	// 0
	LastPullOlderThanDays *int32 `json:"LastPullOlderThanDays,omitempty" xml:"LastPullOlderThanDays,omitempty"`
	// Retains artifact versions whose most recent push time is within the specified number of days.
	//
	// example:
	//
	// 0
	LastPushOlderThanDays *int32 `json:"LastPushOlderThanDays,omitempty" xml:"LastPushOlderThanDays,omitempty"`
	// The number of most recently pushed artifact versions to retain.
	//
	// example:
	//
	// 0
	LatestTagCount *int32 `json:"LatestTagCount,omitempty" xml:"LatestTagCount,omitempty"`
}

func (s ArtifactLifecyclePolicyCondition) String() string {
	return dara.Prettify(s)
}

func (s ArtifactLifecyclePolicyCondition) GoString() string {
	return s.String()
}

func (s *ArtifactLifecyclePolicyCondition) GetLastPullOlderThanDays() *int32 {
	return s.LastPullOlderThanDays
}

func (s *ArtifactLifecyclePolicyCondition) GetLastPushOlderThanDays() *int32 {
	return s.LastPushOlderThanDays
}

func (s *ArtifactLifecyclePolicyCondition) GetLatestTagCount() *int32 {
	return s.LatestTagCount
}

func (s *ArtifactLifecyclePolicyCondition) SetLastPullOlderThanDays(v int32) *ArtifactLifecyclePolicyCondition {
	s.LastPullOlderThanDays = &v
	return s
}

func (s *ArtifactLifecyclePolicyCondition) SetLastPushOlderThanDays(v int32) *ArtifactLifecyclePolicyCondition {
	s.LastPushOlderThanDays = &v
	return s
}

func (s *ArtifactLifecyclePolicyCondition) SetLatestTagCount(v int32) *ArtifactLifecyclePolicyCondition {
	s.LatestTagCount = &v
	return s
}

func (s *ArtifactLifecyclePolicyCondition) Validate() error {
	return dara.Validate(s)
}

type ArtifactLifecyclePolicyFilter struct {
	// The matching expression for the names of additionally retained artifact versions.
	TagWildcard *string `json:"TagWildcard,omitempty" xml:"TagWildcard,omitempty"`
}

func (s ArtifactLifecyclePolicyFilter) String() string {
	return dara.Prettify(s)
}

func (s ArtifactLifecyclePolicyFilter) GoString() string {
	return s.String()
}

func (s *ArtifactLifecyclePolicyFilter) GetTagWildcard() *string {
	return s.TagWildcard
}

func (s *ArtifactLifecyclePolicyFilter) SetTagWildcard(v string) *ArtifactLifecyclePolicyFilter {
	s.TagWildcard = &v
	return s
}

func (s *ArtifactLifecyclePolicyFilter) Validate() error {
	return dara.Validate(s)
}
