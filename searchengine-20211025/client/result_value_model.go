// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResultValue interface {
	dara.Model
	String() string
	GoString() string
	SetPauseAll(v bool) *ResultValue
	GetPauseAll() *bool
	SetPauseIndex(v bool) *ResultValue
	GetPauseIndex() *bool
	SetPauseIndexBatch(v bool) *ResultValue
	GetPauseIndexBatch() *bool
	SetPauseBiz(v bool) *ResultValue
	GetPauseBiz() *bool
	SetPauseRuntime(v bool) *ResultValue
	GetPauseRuntime() *bool
}

type ResultValue struct {
	// Indicates whether all pushes are suspended.
	//
	// example:
	//
	// true
	PauseAll *bool `json:"pauseAll,omitempty" xml:"pauseAll,omitempty"`
	// Indicates whether the push is suspended for the new full index version.
	//
	// example:
	//
	// true
	PauseIndex *bool `json:"pauseIndex,omitempty" xml:"pauseIndex,omitempty"`
	// Indicates whether the push is suspended for the incremental indexes.
	//
	// example:
	//
	// true
	PauseIndexBatch *bool `json:"pauseIndexBatch,omitempty" xml:"pauseIndexBatch,omitempty"`
	// Indicates whether the push is suspended for the configuration.
	//
	// example:
	//
	// true
	PauseBiz *bool `json:"pauseBiz,omitempty" xml:"pauseBiz,omitempty"`
	// Indicates whether the push is suspended for the real-time incremental indexes.
	//
	// example:
	//
	// true
	PauseRuntime *bool `json:"pauseRuntime,omitempty" xml:"pauseRuntime,omitempty"`
}

func (s ResultValue) String() string {
	return dara.Prettify(s)
}

func (s ResultValue) GoString() string {
	return s.String()
}

func (s *ResultValue) GetPauseAll() *bool {
	return s.PauseAll
}

func (s *ResultValue) GetPauseIndex() *bool {
	return s.PauseIndex
}

func (s *ResultValue) GetPauseIndexBatch() *bool {
	return s.PauseIndexBatch
}

func (s *ResultValue) GetPauseBiz() *bool {
	return s.PauseBiz
}

func (s *ResultValue) GetPauseRuntime() *bool {
	return s.PauseRuntime
}

func (s *ResultValue) SetPauseAll(v bool) *ResultValue {
	s.PauseAll = &v
	return s
}

func (s *ResultValue) SetPauseIndex(v bool) *ResultValue {
	s.PauseIndex = &v
	return s
}

func (s *ResultValue) SetPauseIndexBatch(v bool) *ResultValue {
	s.PauseIndexBatch = &v
	return s
}

func (s *ResultValue) SetPauseBiz(v bool) *ResultValue {
	s.PauseBiz = &v
	return s
}

func (s *ResultValue) SetPauseRuntime(v bool) *ResultValue {
	s.PauseRuntime = &v
	return s
}

func (s *ResultValue) Validate() error {
	return dara.Validate(s)
}
