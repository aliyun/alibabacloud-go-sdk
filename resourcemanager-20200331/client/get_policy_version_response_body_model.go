// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolicyVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyVersion(v *GetPolicyVersionResponseBodyPolicyVersion) *GetPolicyVersionResponseBody
	GetPolicyVersion() *GetPolicyVersionResponseBodyPolicyVersion
	SetRequestId(v string) *GetPolicyVersionResponseBody
	GetRequestId() *string
}

type GetPolicyVersionResponseBody struct {
	// The information about the policy version.
	PolicyVersion *GetPolicyVersionResponseBodyPolicyVersion `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPolicyVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetPolicyVersionResponseBody) GetPolicyVersion() *GetPolicyVersionResponseBodyPolicyVersion {
	return s.PolicyVersion
}

func (s *GetPolicyVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPolicyVersionResponseBody) SetPolicyVersion(v *GetPolicyVersionResponseBodyPolicyVersion) *GetPolicyVersionResponseBody {
	s.PolicyVersion = v
	return s
}

func (s *GetPolicyVersionResponseBody) SetRequestId(v string) *GetPolicyVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPolicyVersionResponseBody) Validate() error {
	if s.PolicyVersion != nil {
		if err := s.PolicyVersion.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPolicyVersionResponseBodyPolicyVersion struct {
	// The time when the policy version was created.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// Indicates whether the policy version is the default version.
	//
	// example:
	//
	// false
	IsDefaultVersion *bool `json:"IsDefaultVersion,omitempty" xml:"IsDefaultVersion,omitempty"`
	// The document of the permission policy.
	//
	// example:
	//
	// { \\"Statement\\": [{ \\"Action\\": [\\"oss:*\\"], \\"Effect\\": \\"Allow\\", \\"Resource\\": [\\"acs:oss:*:*:*\\"]}], \\"Version\\": \\"1\\"}
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// The ID of the policy version.
	//
	// example:
	//
	// v3
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetPolicyVersionResponseBodyPolicyVersion) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyVersionResponseBodyPolicyVersion) GoString() string {
	return s.String()
}

func (s *GetPolicyVersionResponseBodyPolicyVersion) GetCreateDate() *string {
	return s.CreateDate
}

func (s *GetPolicyVersionResponseBodyPolicyVersion) GetIsDefaultVersion() *bool {
	return s.IsDefaultVersion
}

func (s *GetPolicyVersionResponseBodyPolicyVersion) GetPolicyDocument() *string {
	return s.PolicyDocument
}

func (s *GetPolicyVersionResponseBodyPolicyVersion) GetVersionId() *string {
	return s.VersionId
}

func (s *GetPolicyVersionResponseBodyPolicyVersion) SetCreateDate(v string) *GetPolicyVersionResponseBodyPolicyVersion {
	s.CreateDate = &v
	return s
}

func (s *GetPolicyVersionResponseBodyPolicyVersion) SetIsDefaultVersion(v bool) *GetPolicyVersionResponseBodyPolicyVersion {
	s.IsDefaultVersion = &v
	return s
}

func (s *GetPolicyVersionResponseBodyPolicyVersion) SetPolicyDocument(v string) *GetPolicyVersionResponseBodyPolicyVersion {
	s.PolicyDocument = &v
	return s
}

func (s *GetPolicyVersionResponseBodyPolicyVersion) SetVersionId(v string) *GetPolicyVersionResponseBodyPolicyVersion {
	s.VersionId = &v
	return s
}

func (s *GetPolicyVersionResponseBodyPolicyVersion) Validate() error {
	return dara.Validate(s)
}
