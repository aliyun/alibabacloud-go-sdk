// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolicyVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyVersions(v *ListPolicyVersionsResponseBodyPolicyVersions) *ListPolicyVersionsResponseBody
	GetPolicyVersions() *ListPolicyVersionsResponseBodyPolicyVersions
	SetRequestId(v string) *ListPolicyVersionsResponseBody
	GetRequestId() *string
}

type ListPolicyVersionsResponseBody struct {
	// The information about the policy version.
	PolicyVersions *ListPolicyVersionsResponseBodyPolicyVersions `json:"PolicyVersions,omitempty" xml:"PolicyVersions,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7B8A4E7D-6CFF-471D-84DF-195A7A241ECB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPolicyVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPolicyVersionsResponseBody) GetPolicyVersions() *ListPolicyVersionsResponseBodyPolicyVersions {
	return s.PolicyVersions
}

func (s *ListPolicyVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPolicyVersionsResponseBody) SetPolicyVersions(v *ListPolicyVersionsResponseBodyPolicyVersions) *ListPolicyVersionsResponseBody {
	s.PolicyVersions = v
	return s
}

func (s *ListPolicyVersionsResponseBody) SetRequestId(v string) *ListPolicyVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPolicyVersionsResponseBody) Validate() error {
	if s.PolicyVersions != nil {
		if err := s.PolicyVersions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPolicyVersionsResponseBodyPolicyVersions struct {
	PolicyVersion []*ListPolicyVersionsResponseBodyPolicyVersionsPolicyVersion `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty" type:"Repeated"`
}

func (s ListPolicyVersionsResponseBodyPolicyVersions) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyVersionsResponseBodyPolicyVersions) GoString() string {
	return s.String()
}

func (s *ListPolicyVersionsResponseBodyPolicyVersions) GetPolicyVersion() []*ListPolicyVersionsResponseBodyPolicyVersionsPolicyVersion {
	return s.PolicyVersion
}

func (s *ListPolicyVersionsResponseBodyPolicyVersions) SetPolicyVersion(v []*ListPolicyVersionsResponseBodyPolicyVersionsPolicyVersion) *ListPolicyVersionsResponseBodyPolicyVersions {
	s.PolicyVersion = v
	return s
}

func (s *ListPolicyVersionsResponseBodyPolicyVersions) Validate() error {
	if s.PolicyVersion != nil {
		for _, item := range s.PolicyVersion {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPolicyVersionsResponseBodyPolicyVersionsPolicyVersion struct {
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
	// The ID of the policy version.
	//
	// example:
	//
	// v3
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s ListPolicyVersionsResponseBodyPolicyVersionsPolicyVersion) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyVersionsResponseBodyPolicyVersionsPolicyVersion) GoString() string {
	return s.String()
}

func (s *ListPolicyVersionsResponseBodyPolicyVersionsPolicyVersion) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ListPolicyVersionsResponseBodyPolicyVersionsPolicyVersion) GetIsDefaultVersion() *bool {
	return s.IsDefaultVersion
}

func (s *ListPolicyVersionsResponseBodyPolicyVersionsPolicyVersion) GetVersionId() *string {
	return s.VersionId
}

func (s *ListPolicyVersionsResponseBodyPolicyVersionsPolicyVersion) SetCreateDate(v string) *ListPolicyVersionsResponseBodyPolicyVersionsPolicyVersion {
	s.CreateDate = &v
	return s
}

func (s *ListPolicyVersionsResponseBodyPolicyVersionsPolicyVersion) SetIsDefaultVersion(v bool) *ListPolicyVersionsResponseBodyPolicyVersionsPolicyVersion {
	s.IsDefaultVersion = &v
	return s
}

func (s *ListPolicyVersionsResponseBodyPolicyVersionsPolicyVersion) SetVersionId(v string) *ListPolicyVersionsResponseBodyPolicyVersionsPolicyVersion {
	s.VersionId = &v
	return s
}

func (s *ListPolicyVersionsResponseBodyPolicyVersionsPolicyVersion) Validate() error {
	return dara.Validate(s)
}
