// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBranchFilter interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *BranchFilter
	GetName() *string
}

type BranchFilter struct {
	// example:
	//
	// master
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s BranchFilter) String() string {
	return dara.Prettify(s)
}

func (s BranchFilter) GoString() string {
	return s.String()
}

func (s *BranchFilter) GetName() *string {
	return s.Name
}

func (s *BranchFilter) SetName(v string) *BranchFilter {
	s.Name = &v
	return s
}

func (s *BranchFilter) Validate() error {
	return dara.Validate(s)
}
