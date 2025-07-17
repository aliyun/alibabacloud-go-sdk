// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIMRobotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRobotId(v int64) *DeleteIMRobotRequest
	GetRobotId() *int64
}

type DeleteIMRobotRequest struct {
	// The ID of the IM chatbot.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	RobotId *int64 `json:"RobotId,omitempty" xml:"RobotId,omitempty"`
}

func (s DeleteIMRobotRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIMRobotRequest) GoString() string {
	return s.String()
}

func (s *DeleteIMRobotRequest) GetRobotId() *int64 {
	return s.RobotId
}

func (s *DeleteIMRobotRequest) SetRobotId(v int64) *DeleteIMRobotRequest {
	s.RobotId = &v
	return s
}

func (s *DeleteIMRobotRequest) Validate() error {
	return dara.Validate(s)
}
