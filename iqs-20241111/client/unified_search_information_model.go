// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnifiedSearchInformation interface {
	dara.Model
	String() string
	GoString() string
	SetSearchTime(v int64) *UnifiedSearchInformation
	GetSearchTime() *int64
}

type UnifiedSearchInformation struct {
	SearchTime *int64 `json:"searchTime,omitempty" xml:"searchTime,omitempty"`
}

func (s UnifiedSearchInformation) String() string {
	return dara.Prettify(s)
}

func (s UnifiedSearchInformation) GoString() string {
	return s.String()
}

func (s *UnifiedSearchInformation) GetSearchTime() *int64 {
	return s.SearchTime
}

func (s *UnifiedSearchInformation) SetSearchTime(v int64) *UnifiedSearchInformation {
	s.SearchTime = &v
	return s
}

func (s *UnifiedSearchInformation) Validate() error {
	return dara.Validate(s)
}
