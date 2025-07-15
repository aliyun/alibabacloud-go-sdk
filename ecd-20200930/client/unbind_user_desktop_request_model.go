// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindUserDesktopRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopAgentIds(v []*string) *UnbindUserDesktopRequest
	GetDesktopAgentIds() []*string
	SetDesktopGroupId(v string) *UnbindUserDesktopRequest
	GetDesktopGroupId() *string
	SetDesktopIds(v []*string) *UnbindUserDesktopRequest
	GetDesktopIds() []*string
	SetForce(v bool) *UnbindUserDesktopRequest
	GetForce() *bool
	SetReason(v string) *UnbindUserDesktopRequest
	GetReason() *string
	SetUserDesktopIds(v []*string) *UnbindUserDesktopRequest
	GetUserDesktopIds() []*string
}

type UnbindUserDesktopRequest struct {
	DesktopAgentIds []*string `json:"DesktopAgentIds,omitempty" xml:"DesktopAgentIds,omitempty" type:"Repeated"`
	DesktopGroupId  *string   `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	DesktopIds      []*string `json:"DesktopIds,omitempty" xml:"DesktopIds,omitempty" type:"Repeated"`
	Force           *bool     `json:"Force,omitempty" xml:"Force,omitempty"`
	Reason          *string   `json:"Reason,omitempty" xml:"Reason,omitempty"`
	UserDesktopIds  []*string `json:"UserDesktopIds,omitempty" xml:"UserDesktopIds,omitempty" type:"Repeated"`
}

func (s UnbindUserDesktopRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindUserDesktopRequest) GoString() string {
	return s.String()
}

func (s *UnbindUserDesktopRequest) GetDesktopAgentIds() []*string {
	return s.DesktopAgentIds
}

func (s *UnbindUserDesktopRequest) GetDesktopGroupId() *string {
	return s.DesktopGroupId
}

func (s *UnbindUserDesktopRequest) GetDesktopIds() []*string {
	return s.DesktopIds
}

func (s *UnbindUserDesktopRequest) GetForce() *bool {
	return s.Force
}

func (s *UnbindUserDesktopRequest) GetReason() *string {
	return s.Reason
}

func (s *UnbindUserDesktopRequest) GetUserDesktopIds() []*string {
	return s.UserDesktopIds
}

func (s *UnbindUserDesktopRequest) SetDesktopAgentIds(v []*string) *UnbindUserDesktopRequest {
	s.DesktopAgentIds = v
	return s
}

func (s *UnbindUserDesktopRequest) SetDesktopGroupId(v string) *UnbindUserDesktopRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *UnbindUserDesktopRequest) SetDesktopIds(v []*string) *UnbindUserDesktopRequest {
	s.DesktopIds = v
	return s
}

func (s *UnbindUserDesktopRequest) SetForce(v bool) *UnbindUserDesktopRequest {
	s.Force = &v
	return s
}

func (s *UnbindUserDesktopRequest) SetReason(v string) *UnbindUserDesktopRequest {
	s.Reason = &v
	return s
}

func (s *UnbindUserDesktopRequest) SetUserDesktopIds(v []*string) *UnbindUserDesktopRequest {
	s.UserDesktopIds = v
	return s
}

func (s *UnbindUserDesktopRequest) Validate() error {
	return dara.Validate(s)
}
