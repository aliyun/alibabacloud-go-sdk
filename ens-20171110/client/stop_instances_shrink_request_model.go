// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIdsShrink(v string) *StopInstancesShrinkRequest
	GetInstanceIdsShrink() *string
}

type StopInstancesShrinkRequest struct {
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s StopInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StopInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *StopInstancesShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *StopInstancesShrinkRequest) SetInstanceIdsShrink(v string) *StopInstancesShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *StopInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
