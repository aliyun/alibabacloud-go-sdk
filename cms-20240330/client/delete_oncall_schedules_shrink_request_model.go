// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOncallSchedulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOncallScheduleIdsShrink(v string) *DeleteOncallSchedulesShrinkRequest
	GetOncallScheduleIdsShrink() *string
}

type DeleteOncallSchedulesShrinkRequest struct {
	OncallScheduleIdsShrink *string `json:"oncallScheduleIds,omitempty" xml:"oncallScheduleIds,omitempty"`
}

func (s DeleteOncallSchedulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteOncallSchedulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteOncallSchedulesShrinkRequest) GetOncallScheduleIdsShrink() *string {
	return s.OncallScheduleIdsShrink
}

func (s *DeleteOncallSchedulesShrinkRequest) SetOncallScheduleIdsShrink(v string) *DeleteOncallSchedulesShrinkRequest {
	s.OncallScheduleIdsShrink = &v
	return s
}

func (s *DeleteOncallSchedulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
