// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRLProgressFatal interface {
	dara.Model
	String() string
	GoString() string
	SetCollectNs(v string) *RLProgressFatal
	GetCollectNs() *string
	SetMessage(v string) *RLProgressFatal
	GetMessage() *string
	SetRawMessage(v string) *RLProgressFatal
	GetRawMessage() *string
	SetSubsecNs(v int64) *RLProgressFatal
	GetSubsecNs() *int64
	SetTime(v int64) *RLProgressFatal
	GetTime() *int64
}

type RLProgressFatal struct {
	// 锚点行 agent_collect_time（纳秒字符串，超 JS 安全整数）
	//
	// example:
	//
	// 1787474487713456789
	CollectNs *string `json:"CollectNs,omitempty" xml:"CollectNs,omitempty"`
	// 错误文案（截断至 500 字符）
	//
	// example:
	//
	// CUDA out of memory. Tried to allocate 2.00 GiB
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 原始日志行（截断至 2000 字符）；调用 GetRLLogContext 时作为 AnchorMessage 传入
	//
	// example:
	//
	// [rank0]: torch.OutOfMemoryError: CUDA out of memory. Tried to allocate 2.00 GiB
	RawMessage *string `json:"RawMessage,omitempty" xml:"RawMessage,omitempty"`
	// 同秒内的纳秒偏移，用于同秒日志排序
	//
	// example:
	//
	// 123456789
	SubsecNs *int64 `json:"SubsecNs,omitempty" xml:"SubsecNs,omitempty"`
	// 日志时间（unix 秒）
	//
	// example:
	//
	// 1787474487
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s RLProgressFatal) String() string {
	return dara.Prettify(s)
}

func (s RLProgressFatal) GoString() string {
	return s.String()
}

func (s *RLProgressFatal) GetCollectNs() *string {
	return s.CollectNs
}

func (s *RLProgressFatal) GetMessage() *string {
	return s.Message
}

func (s *RLProgressFatal) GetRawMessage() *string {
	return s.RawMessage
}

func (s *RLProgressFatal) GetSubsecNs() *int64 {
	return s.SubsecNs
}

func (s *RLProgressFatal) GetTime() *int64 {
	return s.Time
}

func (s *RLProgressFatal) SetCollectNs(v string) *RLProgressFatal {
	s.CollectNs = &v
	return s
}

func (s *RLProgressFatal) SetMessage(v string) *RLProgressFatal {
	s.Message = &v
	return s
}

func (s *RLProgressFatal) SetRawMessage(v string) *RLProgressFatal {
	s.RawMessage = &v
	return s
}

func (s *RLProgressFatal) SetSubsecNs(v int64) *RLProgressFatal {
	s.SubsecNs = &v
	return s
}

func (s *RLProgressFatal) SetTime(v int64) *RLProgressFatal {
	s.Time = &v
	return s
}

func (s *RLProgressFatal) Validate() error {
	return dara.Validate(s)
}
