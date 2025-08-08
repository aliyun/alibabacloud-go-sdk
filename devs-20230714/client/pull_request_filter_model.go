// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPullRequestFilter interface {
	dara.Model
	String() string
	GoString() string
	SetSourceBranch(v string) *PullRequestFilter
	GetSourceBranch() *string
	SetTargetBranch(v string) *PullRequestFilter
	GetTargetBranch() *string
	SetTypes(v []*string) *PullRequestFilter
	GetTypes() []*string
}

type PullRequestFilter struct {
	// example:
	//
	// feature-.*
	SourceBranch *string `json:"sourceBranch,omitempty" xml:"sourceBranch,omitempty"`
	// example:
	//
	// master
	TargetBranch *string   `json:"targetBranch,omitempty" xml:"targetBranch,omitempty"`
	Types        []*string `json:"types,omitempty" xml:"types,omitempty" type:"Repeated"`
}

func (s PullRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s PullRequestFilter) GoString() string {
	return s.String()
}

func (s *PullRequestFilter) GetSourceBranch() *string {
	return s.SourceBranch
}

func (s *PullRequestFilter) GetTargetBranch() *string {
	return s.TargetBranch
}

func (s *PullRequestFilter) GetTypes() []*string {
	return s.Types
}

func (s *PullRequestFilter) SetSourceBranch(v string) *PullRequestFilter {
	s.SourceBranch = &v
	return s
}

func (s *PullRequestFilter) SetTargetBranch(v string) *PullRequestFilter {
	s.TargetBranch = &v
	return s
}

func (s *PullRequestFilter) SetTypes(v []*string) *PullRequestFilter {
	s.Types = v
	return s
}

func (s *PullRequestFilter) Validate() error {
	return dara.Validate(s)
}
