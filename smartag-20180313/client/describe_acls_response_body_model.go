// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeACLsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcls(v *DescribeACLsResponseBodyAcls) *DescribeACLsResponseBody
	GetAcls() *DescribeACLsResponseBodyAcls
	SetPageNumber(v int32) *DescribeACLsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeACLsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeACLsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeACLsResponseBody
	GetTotalCount() *int32
}

type DescribeACLsResponseBody struct {
	Acls *DescribeACLsResponseBodyAcls `json:"Acls,omitempty" xml:"Acls,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 660F303F-C88E-4026-BC6A-FC24B78FD7EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeACLsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeACLsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeACLsResponseBody) GetAcls() *DescribeACLsResponseBodyAcls {
	return s.Acls
}

func (s *DescribeACLsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeACLsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeACLsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeACLsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeACLsResponseBody) SetAcls(v *DescribeACLsResponseBodyAcls) *DescribeACLsResponseBody {
	s.Acls = v
	return s
}

func (s *DescribeACLsResponseBody) SetPageNumber(v int32) *DescribeACLsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeACLsResponseBody) SetPageSize(v int32) *DescribeACLsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeACLsResponseBody) SetRequestId(v string) *DescribeACLsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeACLsResponseBody) SetTotalCount(v int32) *DescribeACLsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeACLsResponseBody) Validate() error {
	if s.Acls != nil {
		if err := s.Acls.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeACLsResponseBodyAcls struct {
	Acl []*DescribeACLsResponseBodyAclsAcl `json:"Acl,omitempty" xml:"Acl,omitempty" type:"Repeated"`
}

func (s DescribeACLsResponseBodyAcls) String() string {
	return dara.Prettify(s)
}

func (s DescribeACLsResponseBodyAcls) GoString() string {
	return s.String()
}

func (s *DescribeACLsResponseBodyAcls) GetAcl() []*DescribeACLsResponseBodyAclsAcl {
	return s.Acl
}

func (s *DescribeACLsResponseBodyAcls) SetAcl(v []*DescribeACLsResponseBodyAclsAcl) *DescribeACLsResponseBodyAcls {
	s.Acl = v
	return s
}

func (s *DescribeACLsResponseBodyAcls) Validate() error {
	if s.Acl != nil {
		for _, item := range s.Acl {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeACLsResponseBodyAclsAcl struct {
	AclId           *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	AclType         *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SagCount        *string `json:"SagCount,omitempty" xml:"SagCount,omitempty"`
}

func (s DescribeACLsResponseBodyAclsAcl) String() string {
	return dara.Prettify(s)
}

func (s DescribeACLsResponseBodyAclsAcl) GoString() string {
	return s.String()
}

func (s *DescribeACLsResponseBodyAclsAcl) GetAclId() *string {
	return s.AclId
}

func (s *DescribeACLsResponseBodyAclsAcl) GetAclType() *string {
	return s.AclType
}

func (s *DescribeACLsResponseBodyAclsAcl) GetName() *string {
	return s.Name
}

func (s *DescribeACLsResponseBodyAclsAcl) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeACLsResponseBodyAclsAcl) GetSagCount() *string {
	return s.SagCount
}

func (s *DescribeACLsResponseBodyAclsAcl) SetAclId(v string) *DescribeACLsResponseBodyAclsAcl {
	s.AclId = &v
	return s
}

func (s *DescribeACLsResponseBodyAclsAcl) SetAclType(v string) *DescribeACLsResponseBodyAclsAcl {
	s.AclType = &v
	return s
}

func (s *DescribeACLsResponseBodyAclsAcl) SetName(v string) *DescribeACLsResponseBodyAclsAcl {
	s.Name = &v
	return s
}

func (s *DescribeACLsResponseBodyAclsAcl) SetResourceGroupId(v string) *DescribeACLsResponseBodyAclsAcl {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeACLsResponseBodyAclsAcl) SetSagCount(v string) *DescribeACLsResponseBodyAclsAcl {
	s.SagCount = &v
	return s
}

func (s *DescribeACLsResponseBodyAclsAcl) Validate() error {
	return dara.Validate(s)
}
