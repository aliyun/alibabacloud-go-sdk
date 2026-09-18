// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertRobotsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRobotIds(v []*string) *DeleteAlertRobotsRequest
	GetRobotIds() []*string
	SetType(v string) *DeleteAlertRobotsRequest
	GetType() *string
}

type DeleteAlertRobotsRequest struct {
	// The chatbot ID.
	//
	// This parameter is required.
	RobotIds []*string `json:"robotIds,omitempty" xml:"robotIds,omitempty" type:"Repeated"`
	// The chatbot type.
	//
	// example:
	//
	// DING
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DeleteAlertRobotsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertRobotsRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlertRobotsRequest) GetRobotIds() []*string {
	return s.RobotIds
}

func (s *DeleteAlertRobotsRequest) GetType() *string {
	return s.Type
}

func (s *DeleteAlertRobotsRequest) SetRobotIds(v []*string) *DeleteAlertRobotsRequest {
	s.RobotIds = v
	return s
}

func (s *DeleteAlertRobotsRequest) SetType(v string) *DeleteAlertRobotsRequest {
	s.Type = &v
	return s
}

func (s *DeleteAlertRobotsRequest) Validate() error {
	return dara.Validate(s)
}
