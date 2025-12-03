// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTLSCipherPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIncludeListener(v bool) *ListTLSCipherPoliciesRequest
	GetIncludeListener() *bool
	SetMaxItems(v int32) *ListTLSCipherPoliciesRequest
	GetMaxItems() *int32
	SetName(v string) *ListTLSCipherPoliciesRequest
	GetName() *string
	SetNextToken(v string) *ListTLSCipherPoliciesRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListTLSCipherPoliciesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListTLSCipherPoliciesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListTLSCipherPoliciesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListTLSCipherPoliciesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListTLSCipherPoliciesRequest
	GetResourceOwnerId() *int64
	SetTLSCipherPolicyId(v string) *ListTLSCipherPoliciesRequest
	GetTLSCipherPolicyId() *string
}

type ListTLSCipherPoliciesRequest struct {
	// Specifies whether to return the information about the associated listeners. Valid values:
	//
	// 	- **true**: returns the information about the associated listeners.
	//
	// 	- **false*	- (default): does not return the information about the associated listeners.
	//
	// example:
	//
	// false
	IncludeListener *bool `json:"IncludeListener,omitempty" xml:"IncludeListener,omitempty"`
	// The maximum number of TLS policies to be queried in this call. Valid values: **1*	- to **100**. If you do not set this parameter, the default value **20*	- is used.
	//
	// example:
	//
	// 20
	MaxItems *int32 `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
	// The name of the TLS policy. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// TLSPolicy-test****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// 	- If this is your first query or no next query is to be sent, ignore this parameter.
	//
	// 	- If a next query is to be sent, set the value to the value of NextToken that is returned from the last call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the Classic Load Balancer (CLB) instance is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the TLS policy.
	//
	// example:
	//
	// tls-bp17elso1h323r****
	TLSCipherPolicyId *string `json:"TLSCipherPolicyId,omitempty" xml:"TLSCipherPolicyId,omitempty"`
}

func (s ListTLSCipherPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTLSCipherPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListTLSCipherPoliciesRequest) GetIncludeListener() *bool {
	return s.IncludeListener
}

func (s *ListTLSCipherPoliciesRequest) GetMaxItems() *int32 {
	return s.MaxItems
}

func (s *ListTLSCipherPoliciesRequest) GetName() *string {
	return s.Name
}

func (s *ListTLSCipherPoliciesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTLSCipherPoliciesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListTLSCipherPoliciesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTLSCipherPoliciesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTLSCipherPoliciesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListTLSCipherPoliciesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListTLSCipherPoliciesRequest) GetTLSCipherPolicyId() *string {
	return s.TLSCipherPolicyId
}

func (s *ListTLSCipherPoliciesRequest) SetIncludeListener(v bool) *ListTLSCipherPoliciesRequest {
	s.IncludeListener = &v
	return s
}

func (s *ListTLSCipherPoliciesRequest) SetMaxItems(v int32) *ListTLSCipherPoliciesRequest {
	s.MaxItems = &v
	return s
}

func (s *ListTLSCipherPoliciesRequest) SetName(v string) *ListTLSCipherPoliciesRequest {
	s.Name = &v
	return s
}

func (s *ListTLSCipherPoliciesRequest) SetNextToken(v string) *ListTLSCipherPoliciesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTLSCipherPoliciesRequest) SetOwnerAccount(v string) *ListTLSCipherPoliciesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTLSCipherPoliciesRequest) SetOwnerId(v int64) *ListTLSCipherPoliciesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTLSCipherPoliciesRequest) SetRegionId(v string) *ListTLSCipherPoliciesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTLSCipherPoliciesRequest) SetResourceOwnerAccount(v string) *ListTLSCipherPoliciesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTLSCipherPoliciesRequest) SetResourceOwnerId(v int64) *ListTLSCipherPoliciesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTLSCipherPoliciesRequest) SetTLSCipherPolicyId(v string) *ListTLSCipherPoliciesRequest {
	s.TLSCipherPolicyId = &v
	return s
}

func (s *ListTLSCipherPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
