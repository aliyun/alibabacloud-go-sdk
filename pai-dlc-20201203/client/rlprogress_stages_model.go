// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRLProgressStages interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentIndex(v int32) *RLProgressStages
	GetCurrentIndex() *int32
	SetMode(v string) *RLProgressStages
	GetMode() *string
	SetStages(v []*RLProgressStage) *RLProgressStages
	GetStages() []*RLProgressStage
	SetStepDone(v bool) *RLProgressStages
	GetStepDone() *bool
}

type RLProgressStages struct {
	// 当前所处阶段的下标
	//
	// example:
	//
	// 6
	CurrentIndex *int32 `json:"CurrentIndex,omitempty" xml:"CurrentIndex,omitempty"`
	// disagg / colocate / 空串
	//
	// example:
	//
	// colocate
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// 阶段列表，按流水线顺序
	//
	// example:
	//
	// [{"Key":"generation","Label":"生成","Marker":"start/end generation","Optional":false,"Status":"done","StartTime":1787474487,"EndTime":1787474487,"Duration":0.483}]
	Stages []*RLProgressStage `json:"Stages,omitempty" xml:"Stages,omitempty" type:"Repeated"`
	// 本 step 的阶段流水线是否已走完
	//
	// example:
	//
	// false
	StepDone *bool `json:"StepDone,omitempty" xml:"StepDone,omitempty"`
}

func (s RLProgressStages) String() string {
	return dara.Prettify(s)
}

func (s RLProgressStages) GoString() string {
	return s.String()
}

func (s *RLProgressStages) GetCurrentIndex() *int32 {
	return s.CurrentIndex
}

func (s *RLProgressStages) GetMode() *string {
	return s.Mode
}

func (s *RLProgressStages) GetStages() []*RLProgressStage {
	return s.Stages
}

func (s *RLProgressStages) GetStepDone() *bool {
	return s.StepDone
}

func (s *RLProgressStages) SetCurrentIndex(v int32) *RLProgressStages {
	s.CurrentIndex = &v
	return s
}

func (s *RLProgressStages) SetMode(v string) *RLProgressStages {
	s.Mode = &v
	return s
}

func (s *RLProgressStages) SetStages(v []*RLProgressStage) *RLProgressStages {
	s.Stages = v
	return s
}

func (s *RLProgressStages) SetStepDone(v bool) *RLProgressStages {
	s.StepDone = &v
	return s
}

func (s *RLProgressStages) Validate() error {
	if s.Stages != nil {
		for _, item := range s.Stages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
