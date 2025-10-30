// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomPrivacyPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomPrivacyPolicyId(v string) *DeleteCustomPrivacyPolicyRequest
	GetCustomPrivacyPolicyId() *string
	SetInstanceId(v string) *DeleteCustomPrivacyPolicyRequest
	GetInstanceId() *string
}

type DeleteCustomPrivacyPolicyRequest struct {
	// 自定义条款ID
	//
	// This parameter is required.
	//
	// example:
	//
	// pp_xxxxx
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

func (s DeleteCustomPrivacyPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomPrivacyPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomPrivacyPolicyRequest) GetCustomPrivacyPolicyId() *string {
	return s.CustomPrivacyPolicyId
}

func (s *DeleteCustomPrivacyPolicyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteCustomPrivacyPolicyRequest) SetCustomPrivacyPolicyId(v string) *DeleteCustomPrivacyPolicyRequest {
	s.CustomPrivacyPolicyId = &v
	return s
}

func (s *DeleteCustomPrivacyPolicyRequest) SetInstanceId(v string) *DeleteCustomPrivacyPolicyRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteCustomPrivacyPolicyRequest) Validate() error {
	return dara.Validate(s)
}
