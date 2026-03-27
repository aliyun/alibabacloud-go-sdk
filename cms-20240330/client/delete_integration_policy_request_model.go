// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIntegrationPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForce(v bool) *DeleteIntegrationPolicyRequest
	GetForce() *bool
}

type DeleteIntegrationPolicyRequest struct {
	// Whether to forcibly delete the cloud-native all-in-one machine,
	//
	// default value: `false`.
	//
	// example:
	//
	// false
	Force *bool `json:"force,omitempty" xml:"force,omitempty"`
}

func (s DeleteIntegrationPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIntegrationPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteIntegrationPolicyRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteIntegrationPolicyRequest) SetForce(v bool) *DeleteIntegrationPolicyRequest {
	s.Force = &v
	return s
}

func (s *DeleteIntegrationPolicyRequest) Validate() error {
	return dara.Validate(s)
}
