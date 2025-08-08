// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushFilter interface {
	dara.Model
	String() string
	GoString() string
	SetBranch(v string) *PushFilter
	GetBranch() *string
	SetTag(v string) *PushFilter
	GetTag() *string
}

type PushFilter struct {
	// example:
	//
	// master
	Branch *string `json:"branch,omitempty" xml:"branch,omitempty"`
	// example:
	//
	// prod-.*
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s PushFilter) String() string {
	return dara.Prettify(s)
}

func (s PushFilter) GoString() string {
	return s.String()
}

func (s *PushFilter) GetBranch() *string {
	return s.Branch
}

func (s *PushFilter) GetTag() *string {
	return s.Tag
}

func (s *PushFilter) SetBranch(v string) *PushFilter {
	s.Branch = &v
	return s
}

func (s *PushFilter) SetTag(v string) *PushFilter {
	s.Tag = &v
	return s
}

func (s *PushFilter) Validate() error {
	return dara.Validate(s)
}
