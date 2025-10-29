// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchEventMigrateInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventInfosShrink(v string) *BatchEventMigrateInstanceShrinkRequest
	GetEventInfosShrink() *string
}

type BatchEventMigrateInstanceShrinkRequest struct {
	// The details of events.
	EventInfosShrink *string `json:"EventInfos,omitempty" xml:"EventInfos,omitempty"`
}

func (s BatchEventMigrateInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchEventMigrateInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchEventMigrateInstanceShrinkRequest) GetEventInfosShrink() *string {
	return s.EventInfosShrink
}

func (s *BatchEventMigrateInstanceShrinkRequest) SetEventInfosShrink(v string) *BatchEventMigrateInstanceShrinkRequest {
	s.EventInfosShrink = &v
	return s
}

func (s *BatchEventMigrateInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
