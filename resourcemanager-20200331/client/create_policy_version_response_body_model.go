// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyVersion(v *CreatePolicyVersionResponseBodyPolicyVersion) *CreatePolicyVersionResponseBody
	GetPolicyVersion() *CreatePolicyVersionResponseBodyPolicyVersion
	SetRequestId(v string) *CreatePolicyVersionResponseBody
	GetRequestId() *string
}

type CreatePolicyVersionResponseBody struct {
	// The information about the policy version.
	PolicyVersion *CreatePolicyVersionResponseBodyPolicyVersion `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePolicyVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePolicyVersionResponseBody) GetPolicyVersion() *CreatePolicyVersionResponseBodyPolicyVersion {
	return s.PolicyVersion
}

func (s *CreatePolicyVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePolicyVersionResponseBody) SetPolicyVersion(v *CreatePolicyVersionResponseBodyPolicyVersion) *CreatePolicyVersionResponseBody {
	s.PolicyVersion = v
	return s
}

func (s *CreatePolicyVersionResponseBody) SetRequestId(v string) *CreatePolicyVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePolicyVersionResponseBody) Validate() error {
	if s.PolicyVersion != nil {
		if err := s.PolicyVersion.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePolicyVersionResponseBodyPolicyVersion struct {
	// The time when the policy version was created.
	//
	// example:
	//
	// 2015-01-23T12:33:18
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// Indicates whether the policy version is the default version.
	//
	// example:
	//
	// false
	IsDefaultVersion *bool `json:"IsDefaultVersion,omitempty" xml:"IsDefaultVersion,omitempty"`
	// The ID of the policy version.
	//
	// example:
	//
	// v3
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s CreatePolicyVersionResponseBodyPolicyVersion) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyVersionResponseBodyPolicyVersion) GoString() string {
	return s.String()
}

func (s *CreatePolicyVersionResponseBodyPolicyVersion) GetCreateDate() *string {
	return s.CreateDate
}

func (s *CreatePolicyVersionResponseBodyPolicyVersion) GetIsDefaultVersion() *bool {
	return s.IsDefaultVersion
}

func (s *CreatePolicyVersionResponseBodyPolicyVersion) GetVersionId() *string {
	return s.VersionId
}

func (s *CreatePolicyVersionResponseBodyPolicyVersion) SetCreateDate(v string) *CreatePolicyVersionResponseBodyPolicyVersion {
	s.CreateDate = &v
	return s
}

func (s *CreatePolicyVersionResponseBodyPolicyVersion) SetIsDefaultVersion(v bool) *CreatePolicyVersionResponseBodyPolicyVersion {
	s.IsDefaultVersion = &v
	return s
}

func (s *CreatePolicyVersionResponseBodyPolicyVersion) SetVersionId(v string) *CreatePolicyVersionResponseBodyPolicyVersion {
	s.VersionId = &v
	return s
}

func (s *CreatePolicyVersionResponseBodyPolicyVersion) Validate() error {
	return dara.Validate(s)
}
