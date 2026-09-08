// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRLProgressBufferDetail interface {
	dara.Model
	String() string
	GoString() string
	SetConsumed(v int32) *RLProgressBufferDetail
	GetConsumed() *int32
	SetFinished(v int32) *RLProgressBufferDetail
	GetFinished() *int32
	SetReady(v int32) *RLProgressBufferDetail
	GetReady() *int32
	SetTag(v int32) *RLProgressBufferDetail
	GetTag() *int32
	SetTotal(v int32) *RLProgressBufferDetail
	GetTotal() *int32
}

type RLProgressBufferDetail struct {
	// 已被 trainer 消费的样本数
	//
	// example:
	//
	// 0
	Consumed *int32 `json:"Consumed,omitempty" xml:"Consumed,omitempty"`
	// 已完成样本数
	//
	// example:
	//
	// 500
	Finished *int32 `json:"Finished,omitempty" xml:"Finished,omitempty"`
	// 已就绪样本数
	//
	// example:
	//
	// 500
	Ready *int32 `json:"Ready,omitempty" xml:"Ready,omitempty"`
	// buffer 标签，即 global batch 序号
	//
	// example:
	//
	// 1
	Tag *int32 `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// 目标样本数
	//
	// example:
	//
	// 3
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s RLProgressBufferDetail) String() string {
	return dara.Prettify(s)
}

func (s RLProgressBufferDetail) GoString() string {
	return s.String()
}

func (s *RLProgressBufferDetail) GetConsumed() *int32 {
	return s.Consumed
}

func (s *RLProgressBufferDetail) GetFinished() *int32 {
	return s.Finished
}

func (s *RLProgressBufferDetail) GetReady() *int32 {
	return s.Ready
}

func (s *RLProgressBufferDetail) GetTag() *int32 {
	return s.Tag
}

func (s *RLProgressBufferDetail) GetTotal() *int32 {
	return s.Total
}

func (s *RLProgressBufferDetail) SetConsumed(v int32) *RLProgressBufferDetail {
	s.Consumed = &v
	return s
}

func (s *RLProgressBufferDetail) SetFinished(v int32) *RLProgressBufferDetail {
	s.Finished = &v
	return s
}

func (s *RLProgressBufferDetail) SetReady(v int32) *RLProgressBufferDetail {
	s.Ready = &v
	return s
}

func (s *RLProgressBufferDetail) SetTag(v int32) *RLProgressBufferDetail {
	s.Tag = &v
	return s
}

func (s *RLProgressBufferDetail) SetTotal(v int32) *RLProgressBufferDetail {
	s.Total = &v
	return s
}

func (s *RLProgressBufferDetail) Validate() error {
	return dara.Validate(s)
}
