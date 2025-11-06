// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVhostShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *DeleteVhostShrinkRequest
	GetConsoleSessionId() *string
	SetInstanceId(v string) *DeleteVhostShrinkRequest
	GetInstanceId() *string
	SetVhostName(v string) *DeleteVhostShrinkRequest
	GetVhostName() *string
	SetVhostNamesShrink(v string) *DeleteVhostShrinkRequest
	GetVhostNamesShrink() *string
}

type DeleteVhostShrinkRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	VhostName        *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
	VhostNamesShrink *string `json:"VhostNames,omitempty" xml:"VhostNames,omitempty"`
}

func (s DeleteVhostShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVhostShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteVhostShrinkRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *DeleteVhostShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteVhostShrinkRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *DeleteVhostShrinkRequest) GetVhostNamesShrink() *string {
	return s.VhostNamesShrink
}

func (s *DeleteVhostShrinkRequest) SetConsoleSessionId(v string) *DeleteVhostShrinkRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *DeleteVhostShrinkRequest) SetInstanceId(v string) *DeleteVhostShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteVhostShrinkRequest) SetVhostName(v string) *DeleteVhostShrinkRequest {
	s.VhostName = &v
	return s
}

func (s *DeleteVhostShrinkRequest) SetVhostNamesShrink(v string) *DeleteVhostShrinkRequest {
	s.VhostNamesShrink = &v
	return s
}

func (s *DeleteVhostShrinkRequest) Validate() error {
	return dara.Validate(s)
}
