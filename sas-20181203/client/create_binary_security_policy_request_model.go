// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBinarySecurityPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusters(v string) *CreateBinarySecurityPolicyRequest
	GetClusters() *string
	SetName(v string) *CreateBinarySecurityPolicyRequest
	GetName() *string
	SetPolicy(v string) *CreateBinarySecurityPolicyRequest
	GetPolicy() *string
	SetRemark(v string) *CreateBinarySecurityPolicyRequest
	GetRemark() *string
	SetResourceOwnerId(v int64) *CreateBinarySecurityPolicyRequest
	GetResourceOwnerId() *int64
	SetSourceIp(v string) *CreateBinarySecurityPolicyRequest
	GetSourceIp() *string
	SetStatus(v string) *CreateBinarySecurityPolicyRequest
	GetStatus() *string
}

type CreateBinarySecurityPolicyRequest struct {
	// The information about the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{\\"ClusterId\\":\\"cc12429dbb8f644f690b0623fb52b4737\\",\\"Namespaces\\":[\\"default\\"]},{\\"ClusterId\\":\\"c9f5b93a8da8f4341b774d79fdbcedb3c\\",\\"Namespaces\\":[\\"default\\"]}]
	Clusters *string `json:"Clusters,omitempty" xml:"Clusters,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// mv-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The content of the policy. Specify a value in the JSON format. You can specify the following keys:
	//
	// 	- **policyMode**: the type of the policy. Default value: requireAttestor.
	//
	// 	- **requiredAttestors**: the required witnesses.
	//
	// This parameter is required.
	//
	// example:
	//
	// {\\"PolicyMode\\":\\"requireAttestor\\",\\"RequiredAttestors\\":[\\"test-xcs-04-12-heyuan\\"]}
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The description.
	//
	// example:
	//
	// remark test
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 59.82.XXX.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The status of the policy. Valid values:
	//
	// 	- **enable**
	//
	// 	- **disable**
	//
	// example:
	//
	// enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateBinarySecurityPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBinarySecurityPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateBinarySecurityPolicyRequest) GetClusters() *string {
	return s.Clusters
}

func (s *CreateBinarySecurityPolicyRequest) GetName() *string {
	return s.Name
}

func (s *CreateBinarySecurityPolicyRequest) GetPolicy() *string {
	return s.Policy
}

func (s *CreateBinarySecurityPolicyRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateBinarySecurityPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateBinarySecurityPolicyRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *CreateBinarySecurityPolicyRequest) GetStatus() *string {
	return s.Status
}

func (s *CreateBinarySecurityPolicyRequest) SetClusters(v string) *CreateBinarySecurityPolicyRequest {
	s.Clusters = &v
	return s
}

func (s *CreateBinarySecurityPolicyRequest) SetName(v string) *CreateBinarySecurityPolicyRequest {
	s.Name = &v
	return s
}

func (s *CreateBinarySecurityPolicyRequest) SetPolicy(v string) *CreateBinarySecurityPolicyRequest {
	s.Policy = &v
	return s
}

func (s *CreateBinarySecurityPolicyRequest) SetRemark(v string) *CreateBinarySecurityPolicyRequest {
	s.Remark = &v
	return s
}

func (s *CreateBinarySecurityPolicyRequest) SetResourceOwnerId(v int64) *CreateBinarySecurityPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateBinarySecurityPolicyRequest) SetSourceIp(v string) *CreateBinarySecurityPolicyRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateBinarySecurityPolicyRequest) SetStatus(v string) *CreateBinarySecurityPolicyRequest {
	s.Status = &v
	return s
}

func (s *CreateBinarySecurityPolicyRequest) Validate() error {
	return dara.Validate(s)
}
