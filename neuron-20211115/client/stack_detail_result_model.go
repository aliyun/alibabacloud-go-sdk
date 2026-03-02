// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStackDetailResult interface {
	dara.Model
	String() string
	GoString() string
	SetStacks(v []*StackInfo) *StackDetailResult
	GetStacks() []*StackInfo
}

type StackDetailResult struct {
	Stacks []*StackInfo `json:"stacks,omitempty" xml:"stacks,omitempty" type:"Repeated"`
}

func (s StackDetailResult) String() string {
	return dara.Prettify(s)
}

func (s StackDetailResult) GoString() string {
	return s.String()
}

func (s *StackDetailResult) GetStacks() []*StackInfo {
	return s.Stacks
}

func (s *StackDetailResult) SetStacks(v []*StackInfo) *StackDetailResult {
	s.Stacks = v
	return s
}

func (s *StackDetailResult) Validate() error {
	if s.Stacks != nil {
		for _, item := range s.Stacks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
