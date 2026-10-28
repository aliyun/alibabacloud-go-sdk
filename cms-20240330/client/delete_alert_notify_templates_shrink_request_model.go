// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertNotifyTemplatesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertNotifyTemplateIdsShrink(v string) *DeleteAlertNotifyTemplatesShrinkRequest
	GetAlertNotifyTemplateIdsShrink() *string
}

type DeleteAlertNotifyTemplatesShrinkRequest struct {
	// The IDs of the notification templates.
	//
	// This parameter is required.
	AlertNotifyTemplateIdsShrink *string `json:"alertNotifyTemplateIds,omitempty" xml:"alertNotifyTemplateIds,omitempty"`
}

func (s DeleteAlertNotifyTemplatesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertNotifyTemplatesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlertNotifyTemplatesShrinkRequest) GetAlertNotifyTemplateIdsShrink() *string {
	return s.AlertNotifyTemplateIdsShrink
}

func (s *DeleteAlertNotifyTemplatesShrinkRequest) SetAlertNotifyTemplateIdsShrink(v string) *DeleteAlertNotifyTemplatesShrinkRequest {
	s.AlertNotifyTemplateIdsShrink = &v
	return s
}

func (s *DeleteAlertNotifyTemplatesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
