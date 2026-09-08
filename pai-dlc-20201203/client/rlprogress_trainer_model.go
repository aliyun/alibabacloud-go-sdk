// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRLProgressTrainer interface {
	dara.Model
	String() string
	GoString() string
	SetMicro(v *RLProgressMicro) *RLProgressTrainer
	GetMicro() *RLProgressMicro
	SetMiniIdx(v int32) *RLProgressTrainer
	GetMiniIdx() *int32
	SetNumMinibatches(v int32) *RLProgressTrainer
	GetNumMinibatches() *int32
	SetSync(v *RLProgressSync) *RLProgressTrainer
	GetSync() *RLProgressSync
}

type RLProgressTrainer struct {
	// micro-batch 进度
	//
	// if can be null:
	// true
	Micro *RLProgressMicro `json:"Micro,omitempty" xml:"Micro,omitempty"`
	// 当前 mini batch 序号
	//
	// example:
	//
	// 1
	MiniIdx *int32 `json:"MiniIdx,omitempty" xml:"MiniIdx,omitempty"`
	// mini-batch 总数
	//
	// example:
	//
	// 4
	NumMinibatches *int32 `json:"NumMinibatches,omitempty" xml:"NumMinibatches,omitempty"`
	// 参数同步状态
	//
	// if can be null:
	// true
	Sync *RLProgressSync `json:"Sync,omitempty" xml:"Sync,omitempty"`
}

func (s RLProgressTrainer) String() string {
	return dara.Prettify(s)
}

func (s RLProgressTrainer) GoString() string {
	return s.String()
}

func (s *RLProgressTrainer) GetMicro() *RLProgressMicro {
	return s.Micro
}

func (s *RLProgressTrainer) GetMiniIdx() *int32 {
	return s.MiniIdx
}

func (s *RLProgressTrainer) GetNumMinibatches() *int32 {
	return s.NumMinibatches
}

func (s *RLProgressTrainer) GetSync() *RLProgressSync {
	return s.Sync
}

func (s *RLProgressTrainer) SetMicro(v *RLProgressMicro) *RLProgressTrainer {
	s.Micro = v
	return s
}

func (s *RLProgressTrainer) SetMiniIdx(v int32) *RLProgressTrainer {
	s.MiniIdx = &v
	return s
}

func (s *RLProgressTrainer) SetNumMinibatches(v int32) *RLProgressTrainer {
	s.NumMinibatches = &v
	return s
}

func (s *RLProgressTrainer) SetSync(v *RLProgressSync) *RLProgressTrainer {
	s.Sync = v
	return s
}

func (s *RLProgressTrainer) Validate() error {
	if s.Micro != nil {
		if err := s.Micro.Validate(); err != nil {
			return err
		}
	}
	if s.Sync != nil {
		if err := s.Sync.Validate(); err != nil {
			return err
		}
	}
	return nil
}
