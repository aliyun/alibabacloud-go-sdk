// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScheduledTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetScheduledId(v string) *DeleteScheduledTaskRequest
	GetScheduledId() *string
}

type DeleteScheduledTaskRequest struct {
	// The ID of the scheduled task.
	//
	// This parameter is required.
	//
	// example:
	//
	// sch-260705-agbx1eev
	ScheduledId *string `json:"ScheduledId,omitempty" xml:"ScheduledId,omitempty"`
}

func (s DeleteScheduledTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteScheduledTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteScheduledTaskRequest) GetScheduledId() *string {
	return s.ScheduledId
}

func (s *DeleteScheduledTaskRequest) SetScheduledId(v string) *DeleteScheduledTaskRequest {
	s.ScheduledId = &v
	return s
}

func (s *DeleteScheduledTaskRequest) Validate() error {
	return dara.Validate(s)
}
