// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveIpControlPolicyItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpControlId(v string) *RemoveIpControlPolicyItemRequest
	GetIpControlId() *string
	SetPolicyItemIds(v string) *RemoveIpControlPolicyItemRequest
	GetPolicyItemIds() *string
	SetSecurityToken(v string) *RemoveIpControlPolicyItemRequest
	GetSecurityToken() *string
}

type RemoveIpControlPolicyItemRequest struct {
	// The ID of the ACL. The ID is unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7ea91319a34d48a09b5c9c871d9768b1
	IpControlId *string `json:"IpControlId,omitempty" xml:"IpControlId,omitempty"`
	// The ID of a policy. Separate multiple IDs with semicolons (;). A maximum of 100 IDs can be entered.
	//
	// This parameter is required.
	//
	// example:
	//
	// P151533572852362;P151533557750260
	PolicyItemIds *string `json:"PolicyItemIds,omitempty" xml:"PolicyItemIds,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s RemoveIpControlPolicyItemRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveIpControlPolicyItemRequest) GoString() string {
	return s.String()
}

func (s *RemoveIpControlPolicyItemRequest) GetIpControlId() *string {
	return s.IpControlId
}

func (s *RemoveIpControlPolicyItemRequest) GetPolicyItemIds() *string {
	return s.PolicyItemIds
}

func (s *RemoveIpControlPolicyItemRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *RemoveIpControlPolicyItemRequest) SetIpControlId(v string) *RemoveIpControlPolicyItemRequest {
	s.IpControlId = &v
	return s
}

func (s *RemoveIpControlPolicyItemRequest) SetPolicyItemIds(v string) *RemoveIpControlPolicyItemRequest {
	s.PolicyItemIds = &v
	return s
}

func (s *RemoveIpControlPolicyItemRequest) SetSecurityToken(v string) *RemoveIpControlPolicyItemRequest {
	s.SecurityToken = &v
	return s
}

func (s *RemoveIpControlPolicyItemRequest) Validate() error {
	return dara.Validate(s)
}
