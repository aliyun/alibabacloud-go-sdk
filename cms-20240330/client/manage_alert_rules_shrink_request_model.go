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
	SetCallSource(v string) *ManageAlertRulesShrinkRequest
	GetCallSource() *string
}

type ManageAlertRulesShrinkRequest struct {
	// The aggregated request body for managing alert rules. This body is shared by the CREATE, UPDATE, PATCH, and BATCH_DELETE operations. Specify fields based on the action.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// The call source identifier that specifies the internal integration channel of the caller, such as bailian, integrationCenter, or managed_service_for_prometheus. This parameter is used to isolate traffic from different sources. You do not need to specify this parameter for standard OpenAPI calls.
	//
	// example:
	//
	// bailian
	CallSource *string `json:"callSource,omitempty" xml:"callSource,omitempty"`
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

func (s *ManageAlertRulesShrinkRequest) GetCallSource() *string {
	return s.CallSource
}

func (s *ManageAlertRulesShrinkRequest) SetBodyShrink(v string) *ManageAlertRulesShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *ManageAlertRulesShrinkRequest) SetCallSource(v string) *ManageAlertRulesShrinkRequest {
	s.CallSource = &v
	return s
}

func (s *ManageAlertRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
