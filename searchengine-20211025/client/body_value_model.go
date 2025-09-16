// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBodyValue interface {
	dara.Model
	String() string
	GoString() string
	SetPauseAll(v bool) *BodyValue
	GetPauseAll() *bool
	SetPauseIndex(v bool) *BodyValue
	GetPauseIndex() *bool
	SetPauseIndexBatch(v bool) *BodyValue
	GetPauseIndexBatch() *bool
	SetPauseBiz(v bool) *BodyValue
	GetPauseBiz() *bool
	SetPauseRuntime(v bool) *BodyValue
	GetPauseRuntime() *bool
}

type BodyValue struct {
	// Specifies whether to suspend all pushes.
	//
	// example:
	//
	// true
	PauseAll *bool `json:"pauseAll,omitempty" xml:"pauseAll,omitempty"`
	// Specifies whether to suspend the push for the new full index version.
	//
	// example:
	//
	// true
	PauseIndex *bool `json:"pauseIndex,omitempty" xml:"pauseIndex,omitempty"`
	// Specifies whether to suspend the push for the incremental indexes.
	//
	// example:
	//
	// true
	PauseIndexBatch *bool `json:"pauseIndexBatch,omitempty" xml:"pauseIndexBatch,omitempty"`
	// Specifies whether to suspend the push for the configuration.
	//
	// example:
	//
	// true
	PauseBiz *bool `json:"pauseBiz,omitempty" xml:"pauseBiz,omitempty"`
	// Specifies whether to suspend the push for the real-time incremental indexes.
	//
	// example:
	//
	// true
	PauseRuntime *bool `json:"pauseRuntime,omitempty" xml:"pauseRuntime,omitempty"`
}

func (s BodyValue) String() string {
	return dara.Prettify(s)
}

func (s BodyValue) GoString() string {
	return s.String()
}

func (s *BodyValue) GetPauseAll() *bool {
	return s.PauseAll
}

func (s *BodyValue) GetPauseIndex() *bool {
	return s.PauseIndex
}

func (s *BodyValue) GetPauseIndexBatch() *bool {
	return s.PauseIndexBatch
}

func (s *BodyValue) GetPauseBiz() *bool {
	return s.PauseBiz
}

func (s *BodyValue) GetPauseRuntime() *bool {
	return s.PauseRuntime
}

func (s *BodyValue) SetPauseAll(v bool) *BodyValue {
	s.PauseAll = &v
	return s
}

func (s *BodyValue) SetPauseIndex(v bool) *BodyValue {
	s.PauseIndex = &v
	return s
}

func (s *BodyValue) SetPauseIndexBatch(v bool) *BodyValue {
	s.PauseIndexBatch = &v
	return s
}

func (s *BodyValue) SetPauseBiz(v bool) *BodyValue {
	s.PauseBiz = &v
	return s
}

func (s *BodyValue) SetPauseRuntime(v bool) *BodyValue {
	s.PauseRuntime = &v
	return s
}

func (s *BodyValue) Validate() error {
	return dara.Validate(s)
}
