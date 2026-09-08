// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRLProgressBuffer interface {
	dara.Model
	String() string
	GoString() string
	SetConsumed(v int32) *RLProgressBuffer
	GetConsumed() *int32
	SetDetail(v []*RLProgressBufferDetail) *RLProgressBuffer
	GetDetail() []*RLProgressBufferDetail
	SetEtaSec(v int64) *RLProgressBuffer
	GetEtaSec() *int64
	SetFillRatePerMin(v float64) *RLProgressBuffer
	GetFillRatePerMin() *float64
	SetFinished(v int32) *RLProgressBuffer
	GetFinished() *int32
	SetPct(v float64) *RLProgressBuffer
	GetPct() *float64
	SetReady(v int32) *RLProgressBuffer
	GetReady() *int32
	SetTarget(v int32) *RLProgressBuffer
	GetTarget() *int32
	SetTrainBatchSize(v int32) *RLProgressBuffer
	GetTrainBatchSize() *int32
	SetTraining(v bool) *RLProgressBuffer
	GetTraining() *bool
}

type RLProgressBuffer struct {
	// The total number of consumed samples in incomplete buffers.
	//
	// example:
	//
	// 0
	Consumed *int32 `json:"Consumed,omitempty" xml:"Consumed,omitempty"`
	// The buffer details split by tag.
	//
	// example:
	//
	// [{"Tag":1,"Ready":500,"Consumed":0,"Finished":500,"Total":512}]
	Detail []*RLProgressBufferDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
	// The estimated number of remaining seconds to fill the buffer.
	//
	// example:
	//
	// 0
	EtaSec *int64 `json:"EtaSec,omitempty" xml:"EtaSec,omitempty"`
	// The fill rate in entries per minute, estimated by using the rollout completion rate as a proxy.
	//
	// example:
	//
	// 120.5
	FillRatePerMin *float64 `json:"FillRatePerMin,omitempty" xml:"FillRatePerMin,omitempty"`
	// The total number of finished samples in incomplete buffers.
	//
	// example:
	//
	// 500
	Finished *int32 `json:"Finished,omitempty" xml:"Finished,omitempty"`
	// The readiness percentage, which is the ratio of Ready to Target.
	//
	// example:
	//
	// 100
	Pct *float64 `json:"Pct,omitempty" xml:"Pct,omitempty"`
	// The total number of ready samples in incomplete buffers.
	//
	// example:
	//
	// 500
	Ready *int32 `json:"Ready,omitempty" xml:"Ready,omitempty"`
	// The total number of target samples in incomplete buffers.
	//
	// example:
	//
	// 512
	Target *int32 `json:"Target,omitempty" xml:"Target,omitempty"`
	// The configured training batch size.
	//
	// example:
	//
	// 512
	TrainBatchSize *int32 `json:"TrainBatchSize,omitempty" xml:"TrainBatchSize,omitempty"`
	// Indicates whether Consumed is greater than 0, which means the batch has been fetched and the trainer is updating.
	//
	// example:
	//
	// true
	Training *bool `json:"Training,omitempty" xml:"Training,omitempty"`
}

func (s RLProgressBuffer) String() string {
	return dara.Prettify(s)
}

func (s RLProgressBuffer) GoString() string {
	return s.String()
}

func (s *RLProgressBuffer) GetConsumed() *int32 {
	return s.Consumed
}

func (s *RLProgressBuffer) GetDetail() []*RLProgressBufferDetail {
	return s.Detail
}

func (s *RLProgressBuffer) GetEtaSec() *int64 {
	return s.EtaSec
}

func (s *RLProgressBuffer) GetFillRatePerMin() *float64 {
	return s.FillRatePerMin
}

func (s *RLProgressBuffer) GetFinished() *int32 {
	return s.Finished
}

func (s *RLProgressBuffer) GetPct() *float64 {
	return s.Pct
}

func (s *RLProgressBuffer) GetReady() *int32 {
	return s.Ready
}

func (s *RLProgressBuffer) GetTarget() *int32 {
	return s.Target
}

func (s *RLProgressBuffer) GetTrainBatchSize() *int32 {
	return s.TrainBatchSize
}

func (s *RLProgressBuffer) GetTraining() *bool {
	return s.Training
}

func (s *RLProgressBuffer) SetConsumed(v int32) *RLProgressBuffer {
	s.Consumed = &v
	return s
}

func (s *RLProgressBuffer) SetDetail(v []*RLProgressBufferDetail) *RLProgressBuffer {
	s.Detail = v
	return s
}

func (s *RLProgressBuffer) SetEtaSec(v int64) *RLProgressBuffer {
	s.EtaSec = &v
	return s
}

func (s *RLProgressBuffer) SetFillRatePerMin(v float64) *RLProgressBuffer {
	s.FillRatePerMin = &v
	return s
}

func (s *RLProgressBuffer) SetFinished(v int32) *RLProgressBuffer {
	s.Finished = &v
	return s
}

func (s *RLProgressBuffer) SetPct(v float64) *RLProgressBuffer {
	s.Pct = &v
	return s
}

func (s *RLProgressBuffer) SetReady(v int32) *RLProgressBuffer {
	s.Ready = &v
	return s
}

func (s *RLProgressBuffer) SetTarget(v int32) *RLProgressBuffer {
	s.Target = &v
	return s
}

func (s *RLProgressBuffer) SetTrainBatchSize(v int32) *RLProgressBuffer {
	s.TrainBatchSize = &v
	return s
}

func (s *RLProgressBuffer) SetTraining(v bool) *RLProgressBuffer {
	s.Training = &v
	return s
}

func (s *RLProgressBuffer) Validate() error {
	if s.Detail != nil {
		for _, item := range s.Detail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
