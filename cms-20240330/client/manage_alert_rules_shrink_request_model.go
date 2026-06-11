// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManageAlertRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *ManageAlertRulesShrinkRequest
	GetBodyShrink() *string
}

type ManageAlertRulesShrinkRequest struct {
	// A unified request body for managing alert rules with the CREATE, UPDATE, PATCH, and BATCH_DELETE actions. The required fields depend on the specified action.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ManageAlertRulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ManageAlertRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ManageAlertRulesShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *ManageAlertRulesShrinkRequest) SetBodyShrink(v string) *ManageAlertRulesShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *ManageAlertRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
