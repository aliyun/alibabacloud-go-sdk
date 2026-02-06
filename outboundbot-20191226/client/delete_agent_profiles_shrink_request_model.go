// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentProfilesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentProfileIdsShrink(v string) *DeleteAgentProfilesShrinkRequest
	GetAgentProfileIdsShrink() *string
	SetAppIp(v string) *DeleteAgentProfilesShrinkRequest
	GetAppIp() *string
}

type DeleteAgentProfilesShrinkRequest struct {
	AgentProfileIdsShrink *string `json:"AgentProfileIds,omitempty" xml:"AgentProfileIds,omitempty"`
	// example:
	//
	// 127.0.0.1
	AppIp *string `json:"AppIp,omitempty" xml:"AppIp,omitempty"`
}

func (s DeleteAgentProfilesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentProfilesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteAgentProfilesShrinkRequest) GetAgentProfileIdsShrink() *string {
	return s.AgentProfileIdsShrink
}

func (s *DeleteAgentProfilesShrinkRequest) GetAppIp() *string {
	return s.AppIp
}

func (s *DeleteAgentProfilesShrinkRequest) SetAgentProfileIdsShrink(v string) *DeleteAgentProfilesShrinkRequest {
	s.AgentProfileIdsShrink = &v
	return s
}

func (s *DeleteAgentProfilesShrinkRequest) SetAppIp(v string) *DeleteAgentProfilesShrinkRequest {
	s.AppIp = &v
	return s
}

func (s *DeleteAgentProfilesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
