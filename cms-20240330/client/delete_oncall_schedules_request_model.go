// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOncallSchedulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOncallScheduleIds(v []*string) *DeleteOncallSchedulesRequest
	GetOncallScheduleIds() []*string
}

type DeleteOncallSchedulesRequest struct {
	// The IDs of the on-call schedules.
	OncallScheduleIds []*string `json:"oncallScheduleIds,omitempty" xml:"oncallScheduleIds,omitempty" type:"Repeated"`
}

func (s DeleteOncallSchedulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteOncallSchedulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteOncallSchedulesRequest) GetOncallScheduleIds() []*string {
	return s.OncallScheduleIds
}

func (s *DeleteOncallSchedulesRequest) SetOncallScheduleIds(v []*string) *DeleteOncallSchedulesRequest {
	s.OncallScheduleIds = v
	return s
}

func (s *DeleteOncallSchedulesRequest) Validate() error {
	return dara.Validate(s)
}
