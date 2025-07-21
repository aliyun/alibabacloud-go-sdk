// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecretRotationPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableAutomaticRotation(v bool) *UpdateSecretRotationPolicyRequest
	GetEnableAutomaticRotation() *bool
	SetRotationInterval(v string) *UpdateSecretRotationPolicyRequest
	GetRotationInterval() *string
	SetSecretName(v string) *UpdateSecretRotationPolicyRequest
	GetSecretName() *string
}

type UpdateSecretRotationPolicyRequest struct {
	// Specifies whether to enable automatic rotation. Valid values:
	//
	// 	- true: enables automatic rotation.
	//
	// 	- false: does not enable automatic rotation. This is the default value.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	EnableAutomaticRotation *bool `json:"EnableAutomaticRotation,omitempty" xml:"EnableAutomaticRotation,omitempty"`
	// The interval for automatic rotation. Valid values: 6 hours to 8,760 hours (365 days).
	//
	// The value is in the `integer[unit]` format.````
	//
	// The unit can be d (day), h (hour), m (minute), or s (second). For example, both 7d and 604800s indicate a seven-day interval.
	//
	// >  This parameter is required if you set the EnableAutomaticRotation parameter to true. This parameter is ignored if you set the EnableAutomaticRotation parameter to false or does not specify the EnableAutomaticRotation parameter.
	//
	// example:
	//
	// 30d
	RotationInterval *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
	// The name of the secret.
	//
	// This parameter is required.
	//
	// example:
	//
	// RdsSecret/Mysql5.4/MyCred
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
}

func (s UpdateSecretRotationPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecretRotationPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecretRotationPolicyRequest) GetEnableAutomaticRotation() *bool {
	return s.EnableAutomaticRotation
}

func (s *UpdateSecretRotationPolicyRequest) GetRotationInterval() *string {
	return s.RotationInterval
}

func (s *UpdateSecretRotationPolicyRequest) GetSecretName() *string {
	return s.SecretName
}

func (s *UpdateSecretRotationPolicyRequest) SetEnableAutomaticRotation(v bool) *UpdateSecretRotationPolicyRequest {
	s.EnableAutomaticRotation = &v
	return s
}

func (s *UpdateSecretRotationPolicyRequest) SetRotationInterval(v string) *UpdateSecretRotationPolicyRequest {
	s.RotationInterval = &v
	return s
}

func (s *UpdateSecretRotationPolicyRequest) SetSecretName(v string) *UpdateSecretRotationPolicyRequest {
	s.SecretName = &v
	return s
}

func (s *UpdateSecretRotationPolicyRequest) Validate() error {
	return dara.Validate(s)
}
