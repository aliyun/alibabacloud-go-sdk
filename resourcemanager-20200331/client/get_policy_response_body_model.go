// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicy(v *GetPolicyResponseBodyPolicy) *GetPolicyResponseBody
	GetPolicy() *GetPolicyResponseBodyPolicy
	SetRequestId(v string) *GetPolicyResponseBody
	GetRequestId() *string
}

type GetPolicyResponseBody struct {
	// The information about the permission policy.
	Policy *GetPolicyResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 697852FB-50D7-44D9-9774-530C31EAC572
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetPolicyResponseBody) GetPolicy() *GetPolicyResponseBodyPolicy {
	return s.Policy
}

func (s *GetPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPolicyResponseBody) SetPolicy(v *GetPolicyResponseBodyPolicy) *GetPolicyResponseBody {
	s.Policy = v
	return s
}

func (s *GetPolicyResponseBody) SetRequestId(v string) *GetPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPolicyResponseBody) Validate() error {
	if s.Policy != nil {
		if err := s.Policy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPolicyResponseBodyPolicy struct {
	// The number of references to the permission policy.
	//
	// example:
	//
	// 0
	AttachmentCount *int32 `json:"AttachmentCount,omitempty" xml:"AttachmentCount,omitempty"`
	// The time when the permission policy was created.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The default version of the permission policy.
	//
	// example:
	//
	// v1
	DefaultVersion *string `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	// The description of the permission policy.
	//
	// example:
	//
	// OSS administrator
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The document of the permission policy.
	//
	// example:
	//
	// { "Statement": [{ "Action": ["oss:*"], "Effect": "Allow", "Resource": ["acs:oss:*:*:*"]}], "Version": "1"}
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// The name of the permission policy.
	//
	// example:
	//
	// OSS-Administrator
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the permission policy. Valid values:
	//
	// 	- Custom
	//
	// 	- System
	//
	// example:
	//
	// Custom
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The time when the permission policy was updated.
	//
	// example:
	//
	// 2016-01-23T12:33:18Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s GetPolicyResponseBodyPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyResponseBodyPolicy) GoString() string {
	return s.String()
}

func (s *GetPolicyResponseBodyPolicy) GetAttachmentCount() *int32 {
	return s.AttachmentCount
}

func (s *GetPolicyResponseBodyPolicy) GetCreateDate() *string {
	return s.CreateDate
}

func (s *GetPolicyResponseBodyPolicy) GetDefaultVersion() *string {
	return s.DefaultVersion
}

func (s *GetPolicyResponseBodyPolicy) GetDescription() *string {
	return s.Description
}

func (s *GetPolicyResponseBodyPolicy) GetPolicyDocument() *string {
	return s.PolicyDocument
}

func (s *GetPolicyResponseBodyPolicy) GetPolicyName() *string {
	return s.PolicyName
}

func (s *GetPolicyResponseBodyPolicy) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetPolicyResponseBodyPolicy) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *GetPolicyResponseBodyPolicy) SetAttachmentCount(v int32) *GetPolicyResponseBodyPolicy {
	s.AttachmentCount = &v
	return s
}

func (s *GetPolicyResponseBodyPolicy) SetCreateDate(v string) *GetPolicyResponseBodyPolicy {
	s.CreateDate = &v
	return s
}

func (s *GetPolicyResponseBodyPolicy) SetDefaultVersion(v string) *GetPolicyResponseBodyPolicy {
	s.DefaultVersion = &v
	return s
}

func (s *GetPolicyResponseBodyPolicy) SetDescription(v string) *GetPolicyResponseBodyPolicy {
	s.Description = &v
	return s
}

func (s *GetPolicyResponseBodyPolicy) SetPolicyDocument(v string) *GetPolicyResponseBodyPolicy {
	s.PolicyDocument = &v
	return s
}

func (s *GetPolicyResponseBodyPolicy) SetPolicyName(v string) *GetPolicyResponseBodyPolicy {
	s.PolicyName = &v
	return s
}

func (s *GetPolicyResponseBodyPolicy) SetPolicyType(v string) *GetPolicyResponseBodyPolicy {
	s.PolicyType = &v
	return s
}

func (s *GetPolicyResponseBodyPolicy) SetUpdateDate(v string) *GetPolicyResponseBodyPolicy {
	s.UpdateDate = &v
	return s
}

func (s *GetPolicyResponseBodyPolicy) Validate() error {
	return dara.Validate(s)
}
