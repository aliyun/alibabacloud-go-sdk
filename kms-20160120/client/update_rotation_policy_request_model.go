// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRotationPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableAutomaticRotation(v bool) *UpdateRotationPolicyRequest
	GetEnableAutomaticRotation() *bool
	SetKeyId(v string) *UpdateRotationPolicyRequest
	GetKeyId() *string
	SetRotationInterval(v string) *UpdateRotationPolicyRequest
	GetRotationInterval() *string
}

type UpdateRotationPolicyRequest struct {
	// Specifies whether to enable automatic key rotation. Valid values:
	//
	// 	- true: enables automatic key rotation.
	//
	// 	- false: disables automatic key rotation.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	EnableAutomaticRotation *bool `json:"EnableAutomaticRotation,omitempty" xml:"EnableAutomaticRotation,omitempty"`
	// The ID of the customer master key (CMK). The ID must be globally unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-12345678****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The period of automatic key rotation. Specify the value in the integer[unit] format. The following units are supported: d (day), h (hour), m (minute), and s (second). For example, you can use either 7d or 604800s to specify a seven-day period. The period can range from 7 days to 730 days.
	//
	// >  If you set the EnableAutomaticRotation parameter to true, you must also specify this parameter. If you set the EnableAutomaticRotation parameter to false, you can leave this parameter unspecified.
	//
	// example:
	//
	// 30d
	RotationInterval *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
}

func (s UpdateRotationPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRotationPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateRotationPolicyRequest) GetEnableAutomaticRotation() *bool {
	return s.EnableAutomaticRotation
}

func (s *UpdateRotationPolicyRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *UpdateRotationPolicyRequest) GetRotationInterval() *string {
	return s.RotationInterval
}

func (s *UpdateRotationPolicyRequest) SetEnableAutomaticRotation(v bool) *UpdateRotationPolicyRequest {
	s.EnableAutomaticRotation = &v
	return s
}

func (s *UpdateRotationPolicyRequest) SetKeyId(v string) *UpdateRotationPolicyRequest {
	s.KeyId = &v
	return s
}

func (s *UpdateRotationPolicyRequest) SetRotationInterval(v string) *UpdateRotationPolicyRequest {
	s.RotationInterval = &v
	return s
}

func (s *UpdateRotationPolicyRequest) Validate() error {
	return dara.Validate(s)
}
