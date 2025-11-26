// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplacementStrategy interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceCategories(v []*InstanceCategory) *ReplacementStrategy
	GetInstanceCategories() []*InstanceCategory
}

type ReplacementStrategy struct {
	InstanceCategories []*InstanceCategory `json:"InstanceCategories,omitempty" xml:"InstanceCategories,omitempty" type:"Repeated"`
}

func (s ReplacementStrategy) String() string {
	return dara.Prettify(s)
}

func (s ReplacementStrategy) GoString() string {
	return s.String()
}

func (s *ReplacementStrategy) GetInstanceCategories() []*InstanceCategory {
	return s.InstanceCategories
}

func (s *ReplacementStrategy) SetInstanceCategories(v []*InstanceCategory) *ReplacementStrategy {
	s.InstanceCategories = v
	return s
}

func (s *ReplacementStrategy) Validate() error {
	if s.InstanceCategories != nil {
		for _, item := range s.InstanceCategories {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
