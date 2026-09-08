// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRLProgressProcessed interface {
	dara.Model
	String() string
	GoString() string
	SetDone(v int32) *RLProgressProcessed
	GetDone() *int32
	SetTotal(v int32) *RLProgressProcessed
	GetTotal() *int32
}

type RLProgressProcessed struct {
	// 已处理条数
	//
	// example:
	//
	// true
	Done *int32 `json:"Done,omitempty" xml:"Done,omitempty"`
	// 总条数
	//
	// example:
	//
	// 3
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s RLProgressProcessed) String() string {
	return dara.Prettify(s)
}

func (s RLProgressProcessed) GoString() string {
	return s.String()
}

func (s *RLProgressProcessed) GetDone() *int32 {
	return s.Done
}

func (s *RLProgressProcessed) GetTotal() *int32 {
	return s.Total
}

func (s *RLProgressProcessed) SetDone(v int32) *RLProgressProcessed {
	s.Done = &v
	return s
}

func (s *RLProgressProcessed) SetTotal(v int32) *RLProgressProcessed {
	s.Total = &v
	return s
}

func (s *RLProgressProcessed) Validate() error {
	return dara.Validate(s)
}
