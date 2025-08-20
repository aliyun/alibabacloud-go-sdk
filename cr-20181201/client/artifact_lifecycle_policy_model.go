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
	Condition *ArtifactLifecyclePolicyCondition `json:"Condition,omitempty" xml:"Condition,omitempty" type:"Struct"`
	Filter    *ArtifactLifecyclePolicyFilter    `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	Type      *string                           `json:"Type,omitempty" xml:"Type,omitempty"`
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
	return dara.Validate(s)
}

type ArtifactLifecyclePolicyCondition struct {
	LastPullOlderThanDays *int32 `json:"LastPullOlderThanDays,omitempty" xml:"LastPullOlderThanDays,omitempty"`
	LastPushOlderThanDays *int32 `json:"LastPushOlderThanDays,omitempty" xml:"LastPushOlderThanDays,omitempty"`
	LatestTagCount        *int32 `json:"LatestTagCount,omitempty" xml:"LatestTagCount,omitempty"`
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
