// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertActionsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertActionIdsShrink(v string) *DeleteAlertActionsShrinkRequest
	GetAlertActionIdsShrink() *string
}

type DeleteAlertActionsShrinkRequest struct {
	// This parameter is required.
	AlertActionIdsShrink *string `json:"alertActionIds,omitempty" xml:"alertActionIds,omitempty"`
}

func (s DeleteAlertActionsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertActionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlertActionsShrinkRequest) GetAlertActionIdsShrink() *string {
	return s.AlertActionIdsShrink
}

func (s *DeleteAlertActionsShrinkRequest) SetAlertActionIdsShrink(v string) *DeleteAlertActionsShrinkRequest {
	s.AlertActionIdsShrink = &v
	return s
}

func (s *DeleteAlertActionsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
