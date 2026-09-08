// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRLProgressSlowDetail interface {
	dara.Model
	String() string
	GoString() string
	SetElapsed(v float64) *RLProgressSlowDetail
	GetElapsed() *float64
	SetIp(v string) *RLProgressSlowDetail
	GetIp() *string
	SetIpc(v string) *RLProgressSlowDetail
	GetIpc() *string
	SetIsPause(v string) *RLProgressSlowDetail
	GetIsPause() *string
	SetMessage(v string) *RLProgressSlowDetail
	GetMessage() *string
	SetOutQueue(v string) *RLProgressSlowDetail
	GetOutQueue() *string
	SetPod(v string) *RLProgressSlowDetail
	GetPod() *string
	SetRank(v int32) *RLProgressSlowDetail
	GetRank() *int32
	SetRid(v string) *RLProgressSlowDetail
	GetRid() *string
	SetStatePresent(v string) *RLProgressSlowDetail
	GetStatePresent() *string
	SetTime(v int64) *RLProgressSlowDetail
	GetTime() *int64
	SetTokenizerPid(v string) *RLProgressSlowDetail
	GetTokenizerPid() *string
	SetWorkerPid(v int32) *RLProgressSlowDetail
	GetWorkerPid() *int32
}

type RLProgressSlowDetail struct {
	// The elapsed time of the request, in seconds.
	//
	// example:
	//
	// 42.5
	Elapsed *float64 `json:"Elapsed,omitempty" xml:"Elapsed,omitempty"`
	// worker IP
	//
	// example:
	//
	// 192.168.0.12
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The IPC channel identifier, which corresponds to the ipc field in the log.
	//
	// example:
	//
	// ipc://worker-0
	Ipc *string `json:"Ipc,omitempty" xml:"Ipc,omitempty"`
	// Indicates whether the request is paused. This is the raw value of the is_pause field in the log.
	//
	// example:
	//
	// false
	IsPause *string `json:"IsPause,omitempty" xml:"IsPause,omitempty"`
	// The log message, truncated to 700 characters.
	//
	// example:
	//
	// CUDA out of memory. Tried to allocate 2.00 GiB
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The output queue length. This is the raw value of the out_queue field in the log.
	//
	// example:
	//
	// 3
	OutQueue *string `json:"OutQueue,omitempty" xml:"OutQueue,omitempty"`
	// The name of the pod.
	//
	// example:
	//
	// dlc193cpaitk8eny-master-0
	Pod *string `json:"Pod,omitempty" xml:"Pod,omitempty"`
	// The training rank.
	//
	// example:
	//
	// 0
	Rank *int32 `json:"Rank,omitempty" xml:"Rank,omitempty"`
	// The inference request ID, which corresponds to the rid field in the log.
	//
	// example:
	//
	// req-8f3a2c1d
	Rid *string `json:"Rid,omitempty" xml:"Rid,omitempty"`
	// Indicates whether the state is present. This is the raw value of the state_present field in the log.
	//
	// example:
	//
	// true
	StatePresent *string `json:"StatePresent,omitempty" xml:"StatePresent,omitempty"`
	// The log time, in UNIX seconds.
	//
	// example:
	//
	// 1787474487
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// The tokenizer process ID, which corresponds to the pid field in the log.
	//
	// example:
	//
	// 12360
	TokenizerPid *string `json:"TokenizerPid,omitempty" xml:"TokenizerPid,omitempty"`
	// The worker process ID.
	//
	// example:
	//
	// 12345
	WorkerPid *int32 `json:"WorkerPid,omitempty" xml:"WorkerPid,omitempty"`
}

func (s RLProgressSlowDetail) String() string {
	return dara.Prettify(s)
}

func (s RLProgressSlowDetail) GoString() string {
	return s.String()
}

func (s *RLProgressSlowDetail) GetElapsed() *float64 {
	return s.Elapsed
}

func (s *RLProgressSlowDetail) GetIp() *string {
	return s.Ip
}

func (s *RLProgressSlowDetail) GetIpc() *string {
	return s.Ipc
}

func (s *RLProgressSlowDetail) GetIsPause() *string {
	return s.IsPause
}

func (s *RLProgressSlowDetail) GetMessage() *string {
	return s.Message
}

func (s *RLProgressSlowDetail) GetOutQueue() *string {
	return s.OutQueue
}

func (s *RLProgressSlowDetail) GetPod() *string {
	return s.Pod
}

func (s *RLProgressSlowDetail) GetRank() *int32 {
	return s.Rank
}

func (s *RLProgressSlowDetail) GetRid() *string {
	return s.Rid
}

func (s *RLProgressSlowDetail) GetStatePresent() *string {
	return s.StatePresent
}

func (s *RLProgressSlowDetail) GetTime() *int64 {
	return s.Time
}

func (s *RLProgressSlowDetail) GetTokenizerPid() *string {
	return s.TokenizerPid
}

func (s *RLProgressSlowDetail) GetWorkerPid() *int32 {
	return s.WorkerPid
}

func (s *RLProgressSlowDetail) SetElapsed(v float64) *RLProgressSlowDetail {
	s.Elapsed = &v
	return s
}

func (s *RLProgressSlowDetail) SetIp(v string) *RLProgressSlowDetail {
	s.Ip = &v
	return s
}

func (s *RLProgressSlowDetail) SetIpc(v string) *RLProgressSlowDetail {
	s.Ipc = &v
	return s
}

func (s *RLProgressSlowDetail) SetIsPause(v string) *RLProgressSlowDetail {
	s.IsPause = &v
	return s
}

func (s *RLProgressSlowDetail) SetMessage(v string) *RLProgressSlowDetail {
	s.Message = &v
	return s
}

func (s *RLProgressSlowDetail) SetOutQueue(v string) *RLProgressSlowDetail {
	s.OutQueue = &v
	return s
}

func (s *RLProgressSlowDetail) SetPod(v string) *RLProgressSlowDetail {
	s.Pod = &v
	return s
}

func (s *RLProgressSlowDetail) SetRank(v int32) *RLProgressSlowDetail {
	s.Rank = &v
	return s
}

func (s *RLProgressSlowDetail) SetRid(v string) *RLProgressSlowDetail {
	s.Rid = &v
	return s
}

func (s *RLProgressSlowDetail) SetStatePresent(v string) *RLProgressSlowDetail {
	s.StatePresent = &v
	return s
}

func (s *RLProgressSlowDetail) SetTime(v int64) *RLProgressSlowDetail {
	s.Time = &v
	return s
}

func (s *RLProgressSlowDetail) SetTokenizerPid(v string) *RLProgressSlowDetail {
	s.TokenizerPid = &v
	return s
}

func (s *RLProgressSlowDetail) SetWorkerPid(v int32) *RLProgressSlowDetail {
	s.WorkerPid = &v
	return s
}

func (s *RLProgressSlowDetail) Validate() error {
	return dara.Validate(s)
}
