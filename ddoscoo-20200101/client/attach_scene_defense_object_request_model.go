// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachSceneDefenseObjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetObjectType(v string) *AttachSceneDefenseObjectRequest
	GetObjectType() *string
	SetObjects(v string) *AttachSceneDefenseObjectRequest
	GetObjects() *string
	SetPolicyId(v string) *AttachSceneDefenseObjectRequest
	GetPolicyId() *string
}

type AttachSceneDefenseObjectRequest struct {
	// The type of the object. Set the value to **Domain**, which indicates a domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Domain
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The object that you want to add to the policy. Separate multiple objects with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Objects *string `json:"Objects,omitempty" xml:"Objects,omitempty"`
	// The ID of the policy.
	//
	// > You can call the [DescribeSceneDefensePolicies](https://help.aliyun.com/document_detail/159382.html) operation to query the IDs of all policies.
	//
	// This parameter is required.
	//
	// example:
	//
	// 321a-fd31-df51-****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s AttachSceneDefenseObjectRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachSceneDefenseObjectRequest) GoString() string {
	return s.String()
}

func (s *AttachSceneDefenseObjectRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *AttachSceneDefenseObjectRequest) GetObjects() *string {
	return s.Objects
}

func (s *AttachSceneDefenseObjectRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *AttachSceneDefenseObjectRequest) SetObjectType(v string) *AttachSceneDefenseObjectRequest {
	s.ObjectType = &v
	return s
}

func (s *AttachSceneDefenseObjectRequest) SetObjects(v string) *AttachSceneDefenseObjectRequest {
	s.Objects = &v
	return s
}

func (s *AttachSceneDefenseObjectRequest) SetPolicyId(v string) *AttachSceneDefenseObjectRequest {
	s.PolicyId = &v
	return s
}

func (s *AttachSceneDefenseObjectRequest) Validate() error {
	return dara.Validate(s)
}
