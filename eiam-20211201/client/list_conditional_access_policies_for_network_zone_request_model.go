// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConditionalAccessPoliciesForNetworkZoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListConditionalAccessPoliciesForNetworkZoneRequest
	GetInstanceId() *string
	SetNetworkZoneId(v string) *ListConditionalAccessPoliciesForNetworkZoneRequest
	GetNetworkZoneId() *string
}

type ListConditionalAccessPoliciesForNetworkZoneRequest struct {
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Application ID associated with the conditional access policy
	//
	// This parameter is required.
	//
	// example:
	//
	// app_11111
	NetworkZoneId *string `json:"NetworkZoneId,omitempty" xml:"NetworkZoneId,omitempty"`
}

func (s ListConditionalAccessPoliciesForNetworkZoneRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConditionalAccessPoliciesForNetworkZoneRequest) GoString() string {
	return s.String()
}

func (s *ListConditionalAccessPoliciesForNetworkZoneRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListConditionalAccessPoliciesForNetworkZoneRequest) GetNetworkZoneId() *string {
	return s.NetworkZoneId
}

func (s *ListConditionalAccessPoliciesForNetworkZoneRequest) SetInstanceId(v string) *ListConditionalAccessPoliciesForNetworkZoneRequest {
	s.InstanceId = &v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneRequest) SetNetworkZoneId(v string) *ListConditionalAccessPoliciesForNetworkZoneRequest {
	s.NetworkZoneId = &v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneRequest) Validate() error {
	return dara.Validate(s)
}
