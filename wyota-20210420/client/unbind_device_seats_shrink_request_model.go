// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindDeviceSeatsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSerialNoListShrink(v string) *UnbindDeviceSeatsShrinkRequest
	GetSerialNoListShrink() *string
}

type UnbindDeviceSeatsShrinkRequest struct {
	SerialNoListShrink *string `json:"SerialNoList,omitempty" xml:"SerialNoList,omitempty"`
}

func (s UnbindDeviceSeatsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindDeviceSeatsShrinkRequest) GoString() string {
	return s.String()
}

func (s *UnbindDeviceSeatsShrinkRequest) GetSerialNoListShrink() *string {
	return s.SerialNoListShrink
}

func (s *UnbindDeviceSeatsShrinkRequest) SetSerialNoListShrink(v string) *UnbindDeviceSeatsShrinkRequest {
	s.SerialNoListShrink = &v
	return s
}

func (s *UnbindDeviceSeatsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
