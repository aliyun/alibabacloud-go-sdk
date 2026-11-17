// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFirewallFeatureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFeatures(v []*GetFirewallFeatureResponseBodyFeatures) *GetFirewallFeatureResponseBody
	GetFeatures() []*GetFirewallFeatureResponseBodyFeatures
	SetRequestId(v string) *GetFirewallFeatureResponseBody
	GetRequestId() *string
}

type GetFirewallFeatureResponseBody struct {
	// The list of firewall feature statuses. This list always returns the statuses of application control and web filtering.
	Features []*GetFirewallFeatureResponseBodyFeatures `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFirewallFeatureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFirewallFeatureResponseBody) GoString() string {
	return s.String()
}

func (s *GetFirewallFeatureResponseBody) GetFeatures() []*GetFirewallFeatureResponseBodyFeatures {
	return s.Features
}

func (s *GetFirewallFeatureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFirewallFeatureResponseBody) SetFeatures(v []*GetFirewallFeatureResponseBodyFeatures) *GetFirewallFeatureResponseBody {
	s.Features = v
	return s
}

func (s *GetFirewallFeatureResponseBody) SetRequestId(v string) *GetFirewallFeatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFirewallFeatureResponseBody) Validate() error {
	if s.Features != nil {
		for _, item := range s.Features {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetFirewallFeatureResponseBodyFeatures struct {
	// The feature name. Valid values:
	//
	// - **ApplicationControl**: Application control.
	//
	// - **WebFilter**: Web filtering.
	//
	// example:
	//
	// ApplicationControl
	FeatureName *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	// The activation status of the feature. Valid values:
	//
	// - **Enable**: Enabled.
	//
	// - **Disable**: Disabled.
	//
	// example:
	//
	// Enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetFirewallFeatureResponseBodyFeatures) String() string {
	return dara.Prettify(s)
}

func (s GetFirewallFeatureResponseBodyFeatures) GoString() string {
	return s.String()
}

func (s *GetFirewallFeatureResponseBodyFeatures) GetFeatureName() *string {
	return s.FeatureName
}

func (s *GetFirewallFeatureResponseBodyFeatures) GetStatus() *string {
	return s.Status
}

func (s *GetFirewallFeatureResponseBodyFeatures) SetFeatureName(v string) *GetFirewallFeatureResponseBodyFeatures {
	s.FeatureName = &v
	return s
}

func (s *GetFirewallFeatureResponseBodyFeatures) SetStatus(v string) *GetFirewallFeatureResponseBodyFeatures {
	s.Status = &v
	return s
}

func (s *GetFirewallFeatureResponseBodyFeatures) Validate() error {
	return dara.Validate(s)
}
