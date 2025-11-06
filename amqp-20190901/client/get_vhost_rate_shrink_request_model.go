// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVhostRateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *GetVhostRateShrinkRequest
	GetConsoleSessionId() *string
	SetInstanceId(v string) *GetVhostRateShrinkRequest
	GetInstanceId() *string
	SetVhostNamesShrink(v string) *GetVhostRateShrinkRequest
	GetVhostNamesShrink() *string
}

type GetVhostRateShrinkRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	VhostNamesShrink *string `json:"VhostNames,omitempty" xml:"VhostNames,omitempty"`
}

func (s GetVhostRateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVhostRateShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetVhostRateShrinkRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *GetVhostRateShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetVhostRateShrinkRequest) GetVhostNamesShrink() *string {
	return s.VhostNamesShrink
}

func (s *GetVhostRateShrinkRequest) SetConsoleSessionId(v string) *GetVhostRateShrinkRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *GetVhostRateShrinkRequest) SetInstanceId(v string) *GetVhostRateShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetVhostRateShrinkRequest) SetVhostNamesShrink(v string) *GetVhostRateShrinkRequest {
	s.VhostNamesShrink = &v
	return s
}

func (s *GetVhostRateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
