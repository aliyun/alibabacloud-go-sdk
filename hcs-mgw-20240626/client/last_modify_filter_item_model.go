// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLastModifyFilterItem interface {
	dara.Model
	String() string
	GoString() string
	SetTimeFilter(v []*TimeFilter) *LastModifyFilterItem
	GetTimeFilter() []*TimeFilter
}

type LastModifyFilterItem struct {
	TimeFilter []*TimeFilter `json:"TimeFilter,omitempty" xml:"TimeFilter,omitempty" type:"Repeated"`
}

func (s LastModifyFilterItem) String() string {
	return dara.Prettify(s)
}

func (s LastModifyFilterItem) GoString() string {
	return s.String()
}

func (s *LastModifyFilterItem) GetTimeFilter() []*TimeFilter {
	return s.TimeFilter
}

func (s *LastModifyFilterItem) SetTimeFilter(v []*TimeFilter) *LastModifyFilterItem {
	s.TimeFilter = v
	return s
}

func (s *LastModifyFilterItem) Validate() error {
	if s.TimeFilter != nil {
		for _, item := range s.TimeFilter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
