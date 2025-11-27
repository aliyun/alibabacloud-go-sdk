// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOrgsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribeOrgsResponseBody
	GetNextToken() *string
	SetOrgs(v []*DescribeOrgsResponseBodyOrgs) *DescribeOrgsResponseBody
	GetOrgs() []*DescribeOrgsResponseBodyOrgs
	SetRequestId(v string) *DescribeOrgsResponseBody
	GetRequestId() *string
}

type DescribeOrgsResponseBody struct {
	// The token that determines the start point of the query. The return value is the value of the NextToken response parameter that was returned last time the DescribeOrgs operation was called.
	//
	// example:
	//
	// AAAAAV3MpHK****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The organizations.
	Orgs []*DescribeOrgsResponseBodyOrgs `json:"Orgs,omitempty" xml:"Orgs,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0B4BB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOrgsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOrgsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOrgsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeOrgsResponseBody) GetOrgs() []*DescribeOrgsResponseBodyOrgs {
	return s.Orgs
}

func (s *DescribeOrgsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOrgsResponseBody) SetNextToken(v string) *DescribeOrgsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeOrgsResponseBody) SetOrgs(v []*DescribeOrgsResponseBodyOrgs) *DescribeOrgsResponseBody {
	s.Orgs = v
	return s
}

func (s *DescribeOrgsResponseBody) SetRequestId(v string) *DescribeOrgsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOrgsResponseBody) Validate() error {
	if s.Orgs != nil {
		for _, item := range s.Orgs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeOrgsResponseBodyOrgs struct {
	// The organization ID.
	//
	// example:
	//
	// org-****
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// The name of the organizational unit.
	//
	// example:
	//
	// org****
	OrgName     *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	OrgNamePath *string `json:"OrgNamePath,omitempty" xml:"OrgNamePath,omitempty"`
	// The parent organization ID.
	//
	// example:
	//
	// org-****
	ParentOrgId *string `json:"ParentOrgId,omitempty" xml:"ParentOrgId,omitempty"`
}

func (s DescribeOrgsResponseBodyOrgs) String() string {
	return dara.Prettify(s)
}

func (s DescribeOrgsResponseBodyOrgs) GoString() string {
	return s.String()
}

func (s *DescribeOrgsResponseBodyOrgs) GetOrgId() *string {
	return s.OrgId
}

func (s *DescribeOrgsResponseBodyOrgs) GetOrgName() *string {
	return s.OrgName
}

func (s *DescribeOrgsResponseBodyOrgs) GetOrgNamePath() *string {
	return s.OrgNamePath
}

func (s *DescribeOrgsResponseBodyOrgs) GetParentOrgId() *string {
	return s.ParentOrgId
}

func (s *DescribeOrgsResponseBodyOrgs) SetOrgId(v string) *DescribeOrgsResponseBodyOrgs {
	s.OrgId = &v
	return s
}

func (s *DescribeOrgsResponseBodyOrgs) SetOrgName(v string) *DescribeOrgsResponseBodyOrgs {
	s.OrgName = &v
	return s
}

func (s *DescribeOrgsResponseBodyOrgs) SetOrgNamePath(v string) *DescribeOrgsResponseBodyOrgs {
	s.OrgNamePath = &v
	return s
}

func (s *DescribeOrgsResponseBodyOrgs) SetParentOrgId(v string) *DescribeOrgsResponseBodyOrgs {
	s.ParentOrgId = &v
	return s
}

func (s *DescribeOrgsResponseBodyOrgs) Validate() error {
	return dara.Validate(s)
}
