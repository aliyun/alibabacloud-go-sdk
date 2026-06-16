// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomPrivacyPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomPrivacyPolicyId(v string) *GetCustomPrivacyPolicyRequest
	GetCustomPrivacyPolicyId() *string
	SetInstanceId(v string) *GetCustomPrivacyPolicyRequest
	GetInstanceId() *string
}

type GetCustomPrivacyPolicyRequest struct {
	// The ID of the custom privacy policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// pp_xxxxx
	CustomPrivacyPolicyId *string `json:"CustomPrivacyPolicyId,omitempty" xml:"CustomPrivacyPolicyId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetCustomPrivacyPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCustomPrivacyPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetCustomPrivacyPolicyRequest) GetCustomPrivacyPolicyId() *string {
	return s.CustomPrivacyPolicyId
}

func (s *GetCustomPrivacyPolicyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCustomPrivacyPolicyRequest) SetCustomPrivacyPolicyId(v string) *GetCustomPrivacyPolicyRequest {
	s.CustomPrivacyPolicyId = &v
	return s
}

func (s *GetCustomPrivacyPolicyRequest) SetInstanceId(v string) *GetCustomPrivacyPolicyRequest {
	s.InstanceId = &v
	return s
}

func (s *GetCustomPrivacyPolicyRequest) Validate() error {
	return dara.Validate(s)
}
