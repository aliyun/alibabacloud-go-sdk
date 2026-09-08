// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRLProgressSlow interface {
	dara.Model
	String() string
	GoString() string
	SetDetails(v []*RLProgressSlowDetail) *RLProgressSlow
	GetDetails() []*RLProgressSlowDetail
	SetElapsed(v float64) *RLProgressSlow
	GetElapsed() *float64
	SetTime(v int64) *RLProgressSlow
	GetTime() *int64
}

type RLProgressSlow struct {
	// 慢推理明细，最多 20 条
	//
	// example:
	//
	// [{"Elapsed":42.5,"Time":1787474487,"Message":"rollout generation slow","Rank":0,"Pod":"dlc193cpaitk8eny-master-0","WorkerPid":12345,"Ip":"192.168.0.12","Rid":"req-8f3a2c1d","TokenizerPid":"12360","Ipc":"ipc://worker-0","IsPause":"false","StatePresent":"true","OutQueue":"3"}]
	Details []*RLProgressSlowDetail `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	// 最慢一条的已耗时（秒）
	//
	// example:
	//
	// 42.5
	Elapsed *float64 `json:"Elapsed,omitempty" xml:"Elapsed,omitempty"`
	// 最慢一条的日志时间（unix 秒）
	//
	// example:
	//
	// 1787474487
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s RLProgressSlow) String() string {
	return dara.Prettify(s)
}

func (s RLProgressSlow) GoString() string {
	return s.String()
}

func (s *RLProgressSlow) GetDetails() []*RLProgressSlowDetail {
	return s.Details
}

func (s *RLProgressSlow) GetElapsed() *float64 {
	return s.Elapsed
}

func (s *RLProgressSlow) GetTime() *int64 {
	return s.Time
}

func (s *RLProgressSlow) SetDetails(v []*RLProgressSlowDetail) *RLProgressSlow {
	s.Details = v
	return s
}

func (s *RLProgressSlow) SetElapsed(v float64) *RLProgressSlow {
	s.Elapsed = &v
	return s
}

func (s *RLProgressSlow) SetTime(v int64) *RLProgressSlow {
	s.Time = &v
	return s
}

func (s *RLProgressSlow) Validate() error {
	if s.Details != nil {
		for _, item := range s.Details {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
