// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchEventRedeployInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventInfosShrink(v string) *BatchEventRedeployInstanceShrinkRequest
	GetEventInfosShrink() *string
}

type BatchEventRedeployInstanceShrinkRequest struct {
	// List of events.
	EventInfosShrink *string `json:"EventInfos,omitempty" xml:"EventInfos,omitempty"`
}

func (s BatchEventRedeployInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchEventRedeployInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchEventRedeployInstanceShrinkRequest) GetEventInfosShrink() *string {
	return s.EventInfosShrink
}

func (s *BatchEventRedeployInstanceShrinkRequest) SetEventInfosShrink(v string) *BatchEventRedeployInstanceShrinkRequest {
	s.EventInfosShrink = &v
	return s
}

func (s *BatchEventRedeployInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
