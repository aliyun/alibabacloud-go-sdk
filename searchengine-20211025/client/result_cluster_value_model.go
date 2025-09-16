// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResultClusterValue interface {
	dara.Model
	String() string
	GoString() string
	SetBuildParallelNum(v int32) *ResultClusterValue
	GetBuildParallelNum() *int32
	SetMergeParallelNum(v int32) *ResultClusterValue
	GetMergeParallelNum() *int32
}

type ResultClusterValue struct {
	// The maximum number of full indexes that can be concurrently built.
	//
	// example:
	//
	// 2
	BuildParallelNum *int32 `json:"buildParallelNum,omitempty" xml:"buildParallelNum,omitempty"`
	// The maximum number of full indexes that can be concurrently merged.
	//
	// example:
	//
	// 2
	MergeParallelNum *int32 `json:"mergeParallelNum,omitempty" xml:"mergeParallelNum,omitempty"`
}

func (s ResultClusterValue) String() string {
	return dara.Prettify(s)
}

func (s ResultClusterValue) GoString() string {
	return s.String()
}

func (s *ResultClusterValue) GetBuildParallelNum() *int32 {
	return s.BuildParallelNum
}

func (s *ResultClusterValue) GetMergeParallelNum() *int32 {
	return s.MergeParallelNum
}

func (s *ResultClusterValue) SetBuildParallelNum(v int32) *ResultClusterValue {
	s.BuildParallelNum = &v
	return s
}

func (s *ResultClusterValue) SetMergeParallelNum(v int32) *ResultClusterValue {
	s.MergeParallelNum = &v
	return s
}

func (s *ResultClusterValue) Validate() error {
	return dara.Validate(s)
}
