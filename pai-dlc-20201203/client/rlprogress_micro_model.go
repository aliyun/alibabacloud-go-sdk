// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRLProgressMicro interface {
	dara.Model
	String() string
	GoString() string
	SetCurrent(v int32) *RLProgressMicro
	GetCurrent() *int32
	SetTotal(v int32) *RLProgressMicro
	GetTotal() *int32
}

type RLProgressMicro struct {
	// 当前 micro-batch 序号
	//
	// example:
	//
	// 3
	Current *int32 `json:"Current,omitempty" xml:"Current,omitempty"`
	// micro-batch 总数
	//
	// example:
	//
	// 3
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s RLProgressMicro) String() string {
	return dara.Prettify(s)
}

func (s RLProgressMicro) GoString() string {
	return s.String()
}

func (s *RLProgressMicro) GetCurrent() *int32 {
	return s.Current
}

func (s *RLProgressMicro) GetTotal() *int32 {
	return s.Total
}

func (s *RLProgressMicro) SetCurrent(v int32) *RLProgressMicro {
	s.Current = &v
	return s
}

func (s *RLProgressMicro) SetTotal(v int32) *RLProgressMicro {
	s.Total = &v
	return s
}

func (s *RLProgressMicro) Validate() error {
	return dara.Validate(s)
}
