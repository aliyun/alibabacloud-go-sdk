// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableCustomPrivacyPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomPrivacyPolicyId(v string) *DisableCustomPrivacyPolicyRequest
	GetCustomPrivacyPolicyId() *string
	SetInstanceId(v string) *DisableCustomPrivacyPolicyRequest
	GetInstanceId() *string
}

type DisableCustomPrivacyPolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pp_xxxx
	CustomPrivacyPolicyId *string `json:"CustomPrivacyPolicyId,omitempty" xml:"CustomPrivacyPolicyId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DisableCustomPrivacyPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableCustomPrivacyPolicyRequest) GoString() string {
	return s.String()
}

func (s *DisableCustomPrivacyPolicyRequest) GetCustomPrivacyPolicyId() *string {
	return s.CustomPrivacyPolicyId
}

func (s *DisableCustomPrivacyPolicyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableCustomPrivacyPolicyRequest) SetCustomPrivacyPolicyId(v string) *DisableCustomPrivacyPolicyRequest {
	s.CustomPrivacyPolicyId = &v
	return s
}

func (s *DisableCustomPrivacyPolicyRequest) SetInstanceId(v string) *DisableCustomPrivacyPolicyRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableCustomPrivacyPolicyRequest) Validate() error {
	return dara.Validate(s)
}
