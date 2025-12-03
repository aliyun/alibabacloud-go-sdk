// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrganizationId(v string) *DeleteAppMemberRequest
	GetOrganizationId() *string
	SetSubjectId(v string) *DeleteAppMemberRequest
	GetSubjectId() *string
	SetSubjectType(v string) *DeleteAppMemberRequest
	GetSubjectType() *string
}

type DeleteAppMemberRequest struct {
	// example:
	//
	// 66c0c9fffeb86b450c199fcd
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// example:
	//
	// 1332695887xxxxxx
	SubjectId *string `json:"subjectId,omitempty" xml:"subjectId,omitempty"`
	// example:
	//
	// User
	SubjectType *string `json:"subjectType,omitempty" xml:"subjectType,omitempty"`
}

func (s DeleteAppMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppMemberRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppMemberRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *DeleteAppMemberRequest) GetSubjectId() *string {
	return s.SubjectId
}

func (s *DeleteAppMemberRequest) GetSubjectType() *string {
	return s.SubjectType
}

func (s *DeleteAppMemberRequest) SetOrganizationId(v string) *DeleteAppMemberRequest {
	s.OrganizationId = &v
	return s
}

func (s *DeleteAppMemberRequest) SetSubjectId(v string) *DeleteAppMemberRequest {
	s.SubjectId = &v
	return s
}

func (s *DeleteAppMemberRequest) SetSubjectType(v string) *DeleteAppMemberRequest {
	s.SubjectType = &v
	return s
}

func (s *DeleteAppMemberRequest) Validate() error {
	return dara.Validate(s)
}
