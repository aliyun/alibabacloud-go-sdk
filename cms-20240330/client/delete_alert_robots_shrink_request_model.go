// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertRobotsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRobotIdsShrink(v string) *DeleteAlertRobotsShrinkRequest
	GetRobotIdsShrink() *string
	SetType(v string) *DeleteAlertRobotsShrinkRequest
	GetType() *string
}

type DeleteAlertRobotsShrinkRequest struct {
	// The chatbot ID.
	//
	// This parameter is required.
	RobotIdsShrink *string `json:"robotIds,omitempty" xml:"robotIds,omitempty"`
	// The chatbot type.
	//
	// example:
	//
	// DING
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DeleteAlertRobotsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertRobotsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlertRobotsShrinkRequest) GetRobotIdsShrink() *string {
	return s.RobotIdsShrink
}

func (s *DeleteAlertRobotsShrinkRequest) GetType() *string {
	return s.Type
}

func (s *DeleteAlertRobotsShrinkRequest) SetRobotIdsShrink(v string) *DeleteAlertRobotsShrinkRequest {
	s.RobotIdsShrink = &v
	return s
}

func (s *DeleteAlertRobotsShrinkRequest) SetType(v string) *DeleteAlertRobotsShrinkRequest {
	s.Type = &v
	return s
}

func (s *DeleteAlertRobotsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
