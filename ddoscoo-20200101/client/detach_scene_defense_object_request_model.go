// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachSceneDefenseObjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetObjectType(v string) *DetachSceneDefenseObjectRequest
	GetObjectType() *string
	SetObjects(v string) *DetachSceneDefenseObjectRequest
	GetObjects() *string
	SetPolicyId(v string) *DetachSceneDefenseObjectRequest
	GetPolicyId() *string
}

type DetachSceneDefenseObjectRequest struct {
	// The type of the object. Set the value to **Domain**, which indicates a domain name.
	//
	// example:
	//
	// Domain
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The protection asset that you want to remove from a policy. Separate multiple protection assets with commas (,).
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

func (s DetachSceneDefenseObjectRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachSceneDefenseObjectRequest) GoString() string {
	return s.String()
}

func (s *DetachSceneDefenseObjectRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *DetachSceneDefenseObjectRequest) GetObjects() *string {
	return s.Objects
}

func (s *DetachSceneDefenseObjectRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DetachSceneDefenseObjectRequest) SetObjectType(v string) *DetachSceneDefenseObjectRequest {
	s.ObjectType = &v
	return s
}

func (s *DetachSceneDefenseObjectRequest) SetObjects(v string) *DetachSceneDefenseObjectRequest {
	s.Objects = &v
	return s
}

func (s *DetachSceneDefenseObjectRequest) SetPolicyId(v string) *DetachSceneDefenseObjectRequest {
	s.PolicyId = &v
	return s
}

func (s *DetachSceneDefenseObjectRequest) Validate() error {
	return dara.Validate(s)
}
