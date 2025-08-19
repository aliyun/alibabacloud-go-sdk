// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchWindow interface {
	dara.Model
	String() string
	GoString() string
	SetCountBasedWindow(v int32) *BatchWindow
	GetCountBasedWindow() *int32
	SetTimeBasedWindow(v int32) *BatchWindow
	GetTimeBasedWindow() *int32
}

type BatchWindow struct {
	// example:
	//
	// 100
	CountBasedWindow *int32 `json:"CountBasedWindow,omitempty" xml:"CountBasedWindow,omitempty"`
	// example:
	//
	// 10
	TimeBasedWindow *int32 `json:"TimeBasedWindow,omitempty" xml:"TimeBasedWindow,omitempty"`
}

func (s BatchWindow) String() string {
	return dara.Prettify(s)
}

func (s BatchWindow) GoString() string {
	return s.String()
}

func (s *BatchWindow) GetCountBasedWindow() *int32 {
	return s.CountBasedWindow
}

func (s *BatchWindow) GetTimeBasedWindow() *int32 {
	return s.TimeBasedWindow
}

func (s *BatchWindow) SetCountBasedWindow(v int32) *BatchWindow {
	s.CountBasedWindow = &v
	return s
}

func (s *BatchWindow) SetTimeBasedWindow(v int32) *BatchWindow {
	s.TimeBasedWindow = &v
	return s
}

func (s *BatchWindow) Validate() error {
	return dara.Validate(s)
}
