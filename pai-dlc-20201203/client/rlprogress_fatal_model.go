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
	// The agent_collect_time of the anchor line, represented as a nanosecond string that exceeds the JavaScript safe integer limit.
	//
	// example:
	//
	// 1787474487713456789
	CollectNs *string `json:"CollectNs,omitempty" xml:"CollectNs,omitempty"`
	// The error message, truncated to 500 characters.
	//
	// example:
	//
	// CUDA out of memory. Tried to allocate 2.00 GiB
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The raw log line, truncated to 2000 characters. It is passed as AnchorMessage when calling GetRLLogContext.
	//
	// example:
	//
	// [rank0]: torch.OutOfMemoryError: CUDA out of memory. Tried to allocate 2.00 GiB
	RawMessage *string `json:"RawMessage,omitempty" xml:"RawMessage,omitempty"`
	// The nanosecond offset within the same second, used for sorting logs within the same second.
	//
	// example:
	//
	// 123456789
	SubsecNs *int64 `json:"SubsecNs,omitempty" xml:"SubsecNs,omitempty"`
	// The log time in UNIX seconds.
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
