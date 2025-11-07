// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGitOrganizationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *ListGitOrganizationsResponseBody
	GetCount() *int32
	SetGitOrgs(v []*ListGitOrganizationsResponseBodyGitOrgs) *ListGitOrganizationsResponseBody
	GetGitOrgs() []*ListGitOrganizationsResponseBodyGitOrgs
	SetRequestId(v string) *ListGitOrganizationsResponseBody
	GetRequestId() *string
}

type ListGitOrganizationsResponseBody struct {
	// example:
	//
	// 1
	Count   *int32                                     `json:"Count,omitempty" xml:"Count,omitempty"`
	GitOrgs []*ListGitOrganizationsResponseBodyGitOrgs `json:"GitOrgs,omitempty" xml:"GitOrgs,omitempty" type:"Repeated"`
	// example:
	//
	// 9E011F98-4EE5-5E3A-8FA3-D38F2C531C1F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListGitOrganizationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGitOrganizationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGitOrganizationsResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListGitOrganizationsResponseBody) GetGitOrgs() []*ListGitOrganizationsResponseBodyGitOrgs {
	return s.GitOrgs
}

func (s *ListGitOrganizationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGitOrganizationsResponseBody) SetCount(v int32) *ListGitOrganizationsResponseBody {
	s.Count = &v
	return s
}

func (s *ListGitOrganizationsResponseBody) SetGitOrgs(v []*ListGitOrganizationsResponseBodyGitOrgs) *ListGitOrganizationsResponseBody {
	s.GitOrgs = v
	return s
}

func (s *ListGitOrganizationsResponseBody) SetRequestId(v string) *ListGitOrganizationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGitOrganizationsResponseBody) Validate() error {
	if s.GitOrgs != nil {
		for _, item := range s.GitOrgs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGitOrganizationsResponseBodyGitOrgs struct {
	// example:
	//
	// dfj3535
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// example:
	//
	// aliyun-computenest
	OrgName *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
}

func (s ListGitOrganizationsResponseBodyGitOrgs) String() string {
	return dara.Prettify(s)
}

func (s ListGitOrganizationsResponseBodyGitOrgs) GoString() string {
	return s.String()
}

func (s *ListGitOrganizationsResponseBodyGitOrgs) GetOrgId() *string {
	return s.OrgId
}

func (s *ListGitOrganizationsResponseBodyGitOrgs) GetOrgName() *string {
	return s.OrgName
}

func (s *ListGitOrganizationsResponseBodyGitOrgs) SetOrgId(v string) *ListGitOrganizationsResponseBodyGitOrgs {
	s.OrgId = &v
	return s
}

func (s *ListGitOrganizationsResponseBodyGitOrgs) SetOrgName(v string) *ListGitOrganizationsResponseBodyGitOrgs {
	s.OrgName = &v
	return s
}

func (s *ListGitOrganizationsResponseBodyGitOrgs) Validate() error {
	return dara.Validate(s)
}
