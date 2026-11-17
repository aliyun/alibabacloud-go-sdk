// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFirewallFeatureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureName(v string) *UpdateFirewallFeatureRequest
	GetFeatureName() *string
	SetStatus(v string) *UpdateFirewallFeatureRequest
	GetStatus() *string
}

type UpdateFirewallFeatureRequest struct {
	// The name of the feature to enable. Valid values:
	//
	// - **ApplicationControl**: application control
	//
	// - **WebFilter**: web filtering
	//
	// This parameter is required.
	//
	// example:
	//
	// ApplicationControl
	FeatureName *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	// The status of the feature. Only **Enable*	- is supported, which enables the feature. Specifying Disable or other values returns an error. This operation is idempotent if the feature is already enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// Enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateFirewallFeatureRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFirewallFeatureRequest) GoString() string {
	return s.String()
}

func (s *UpdateFirewallFeatureRequest) GetFeatureName() *string {
	return s.FeatureName
}

func (s *UpdateFirewallFeatureRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateFirewallFeatureRequest) SetFeatureName(v string) *UpdateFirewallFeatureRequest {
	s.FeatureName = &v
	return s
}

func (s *UpdateFirewallFeatureRequest) SetStatus(v string) *UpdateFirewallFeatureRequest {
	s.Status = &v
	return s
}

func (s *UpdateFirewallFeatureRequest) Validate() error {
	return dara.Validate(s)
}
