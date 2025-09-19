// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBinarySecurityPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusters(v string) *ModifyBinarySecurityPolicyRequest
	GetClusters() *string
	SetName(v string) *ModifyBinarySecurityPolicyRequest
	GetName() *string
	SetPolicy(v string) *ModifyBinarySecurityPolicyRequest
	GetPolicy() *string
	SetRemark(v string) *ModifyBinarySecurityPolicyRequest
	GetRemark() *string
	SetResourceOwnerId(v int64) *ModifyBinarySecurityPolicyRequest
	GetResourceOwnerId() *int64
	SetSourceIp(v string) *ModifyBinarySecurityPolicyRequest
	GetSourceIp() *string
	SetStatus(v string) *ModifyBinarySecurityPolicyRequest
	GetStatus() *string
}

type ModifyBinarySecurityPolicyRequest struct {
	// Cluster information.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{\\"Namespaces\\":[\\"default\\"],\\"ClusterId\\":\\"c9f5b93a8da8f4341b77***********\\"},{\\"Namespaces\\":[\\"default\\"],\\"ClusterId\\":\\"cc12429dbb8f644f690b0***********\\"}]
	Clusters *string `json:"Clusters,omitempty" xml:"Clusters,omitempty"`
	// Policy name.
	//
	// example:
	//
	// test-policy-04-11
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Policy content. JSON format, Key values:
	//
	// - **policyMode**: Type of policy, default is requireAttestor.
	//
	// - **requiredAttestors**: Required attestors.
	//
	// This parameter is required.
	//
	// example:
	//
	// {\\"PolicyMode\\":\\"requireAttestor\\",\\"RequiredAttestors\\":[\\"test-xcs-04-11-hhht\\"]}
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// Remark.
	//
	// example:
	//
	// Remark
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The IP address of the access source.
	//
	// example:
	//
	// 1.2.3.4
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// Set to enable or disable the policy. Values:
	//
	// - **enabled**: Enable the protection policy.
	//
	// - **disabled**: Disable the protection policy.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyBinarySecurityPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBinarySecurityPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyBinarySecurityPolicyRequest) GetClusters() *string {
	return s.Clusters
}

func (s *ModifyBinarySecurityPolicyRequest) GetName() *string {
	return s.Name
}

func (s *ModifyBinarySecurityPolicyRequest) GetPolicy() *string {
	return s.Policy
}

func (s *ModifyBinarySecurityPolicyRequest) GetRemark() *string {
	return s.Remark
}

func (s *ModifyBinarySecurityPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyBinarySecurityPolicyRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ModifyBinarySecurityPolicyRequest) GetStatus() *string {
	return s.Status
}

func (s *ModifyBinarySecurityPolicyRequest) SetClusters(v string) *ModifyBinarySecurityPolicyRequest {
	s.Clusters = &v
	return s
}

func (s *ModifyBinarySecurityPolicyRequest) SetName(v string) *ModifyBinarySecurityPolicyRequest {
	s.Name = &v
	return s
}

func (s *ModifyBinarySecurityPolicyRequest) SetPolicy(v string) *ModifyBinarySecurityPolicyRequest {
	s.Policy = &v
	return s
}

func (s *ModifyBinarySecurityPolicyRequest) SetRemark(v string) *ModifyBinarySecurityPolicyRequest {
	s.Remark = &v
	return s
}

func (s *ModifyBinarySecurityPolicyRequest) SetResourceOwnerId(v int64) *ModifyBinarySecurityPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyBinarySecurityPolicyRequest) SetSourceIp(v string) *ModifyBinarySecurityPolicyRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyBinarySecurityPolicyRequest) SetStatus(v string) *ModifyBinarySecurityPolicyRequest {
	s.Status = &v
	return s
}

func (s *ModifyBinarySecurityPolicyRequest) Validate() error {
	return dara.Validate(s)
}
