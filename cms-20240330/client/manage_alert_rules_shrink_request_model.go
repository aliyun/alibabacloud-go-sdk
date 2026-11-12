// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManageAlertRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *ManageAlertRulesShrinkRequest
	GetDryRun() *bool
	SetBodyShrink(v string) *ManageAlertRulesShrinkRequest
	GetBodyShrink() *string
	SetCallSource(v string) *ManageAlertRulesShrinkRequest
	GetCallSource() *string
}

type ManageAlertRulesShrinkRequest struct {
	// Specifies whether to perform a dry run. A dry run only validates the request and returns the validation result in dryRunResult without writing to the database, triggering template synchronization, or causing other side effects. Default value: false.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The aggregated request body for managing alert rules. This body is shared by the CREATE, UPDATE, PATCH, and BATCH_DELETE operations. Specify the fields based on the action.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// The call source identifier. It identifies the internal integration channel to which the caller belongs, such as bailian, integrationCenter, or managed_service_for_prometheus, and is used to isolate traffic from different sources. Leave this parameter empty for regular OpenAPI calls.
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

func (s *ManageAlertRulesShrinkRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ManageAlertRulesShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *ManageAlertRulesShrinkRequest) GetCallSource() *string {
	return s.CallSource
}

func (s *ManageAlertRulesShrinkRequest) SetDryRun(v bool) *ManageAlertRulesShrinkRequest {
	s.DryRun = &v
	return s
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
