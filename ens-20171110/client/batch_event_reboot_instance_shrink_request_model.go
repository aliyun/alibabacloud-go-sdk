// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchEventRebootInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventInfosShrink(v string) *BatchEventRebootInstanceShrinkRequest
	GetEventInfosShrink() *string
}

type BatchEventRebootInstanceShrinkRequest struct {
	EventInfosShrink *string `json:"EventInfos,omitempty" xml:"EventInfos,omitempty"`
}

func (s BatchEventRebootInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchEventRebootInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchEventRebootInstanceShrinkRequest) GetEventInfosShrink() *string {
	return s.EventInfosShrink
}

func (s *BatchEventRebootInstanceShrinkRequest) SetEventInfosShrink(v string) *BatchEventRebootInstanceShrinkRequest {
	s.EventInfosShrink = &v
	return s
}

func (s *BatchEventRebootInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
