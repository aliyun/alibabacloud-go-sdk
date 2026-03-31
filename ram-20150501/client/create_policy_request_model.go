// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreatePolicyRequest
	GetDescription() *string
	SetPolicyDocument(v string) *CreatePolicyRequest
	GetPolicyDocument() *string
	SetPolicyName(v string) *CreatePolicyRequest
	GetPolicyName() *string
	SetTag(v []*CreatePolicyRequestTag) *CreatePolicyRequest
	GetTag() []*CreatePolicyRequestTag
}

type CreatePolicyRequest struct {
	// The description of the policy.
	//
	// The description must be 1 to 1,024 characters in length.
	//
	// example:
	//
	// Query ECS instances in a specific region
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The document of the policy.
	//
	// The document must be 1 to 6,144 characters in length.
	//
	// For more information about policy elements and sample policies, see [Policy elements](https://help.aliyun.com/document_detail/93738.html) and [Overview of sample policies](https://help.aliyun.com/document_detail/210969.html).
	//
	// example:
	//
	// {"Statement": [{"Effect": "Allow","Action": "ecs:Describe*","Resource": "acs:ecs:cn-qingdao:*:instance/*"}],"Version": "1"}
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// The name of the policy.
	//
	// The name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
	//
	// example:
	//
	// View-ECS-instances-in-a-specific-region
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The tags.
	Tag []*CreatePolicyRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreatePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePolicyRequest) GetPolicyDocument() *string {
	return s.PolicyDocument
}

func (s *CreatePolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *CreatePolicyRequest) GetTag() []*CreatePolicyRequestTag {
	return s.Tag
}

func (s *CreatePolicyRequest) SetDescription(v string) *CreatePolicyRequest {
	s.Description = &v
	return s
}

func (s *CreatePolicyRequest) SetPolicyDocument(v string) *CreatePolicyRequest {
	s.PolicyDocument = &v
	return s
}

func (s *CreatePolicyRequest) SetPolicyName(v string) *CreatePolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *CreatePolicyRequest) SetTag(v []*CreatePolicyRequestTag) *CreatePolicyRequest {
	s.Tag = v
	return s
}

func (s *CreatePolicyRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreatePolicyRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// owner
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// alice
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePolicyRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyRequestTag) GoString() string {
	return s.String()
}

func (s *CreatePolicyRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreatePolicyRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreatePolicyRequestTag) SetKey(v string) *CreatePolicyRequestTag {
	s.Key = &v
	return s
}

func (s *CreatePolicyRequestTag) SetValue(v string) *CreatePolicyRequestTag {
	s.Value = &v
	return s
}

func (s *CreatePolicyRequestTag) Validate() error {
	return dara.Validate(s)
}
