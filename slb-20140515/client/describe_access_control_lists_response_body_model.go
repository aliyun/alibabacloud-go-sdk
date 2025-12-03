// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessControlListsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcls(v *DescribeAccessControlListsResponseBodyAcls) *DescribeAccessControlListsResponseBody
	GetAcls() *DescribeAccessControlListsResponseBodyAcls
	SetCount(v int32) *DescribeAccessControlListsResponseBody
	GetCount() *int32
	SetPageNumber(v int32) *DescribeAccessControlListsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAccessControlListsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeAccessControlListsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeAccessControlListsResponseBody
	GetTotalCount() *int32
}

type DescribeAccessControlListsResponseBody struct {
	// A list of ACLs.
	Acls *DescribeAccessControlListsResponseBodyAcls `json:"Acls,omitempty" xml:"Acls,omitempty" type:"Struct"`
	// The number of ACLs on the current page.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The number of the returned page. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3CB646EF-6147-4566-A9D9-CE8FBE86F971
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of ACLs.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAccessControlListsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessControlListsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListsResponseBody) GetAcls() *DescribeAccessControlListsResponseBodyAcls {
	return s.Acls
}

func (s *DescribeAccessControlListsResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeAccessControlListsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAccessControlListsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAccessControlListsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccessControlListsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAccessControlListsResponseBody) SetAcls(v *DescribeAccessControlListsResponseBodyAcls) *DescribeAccessControlListsResponseBody {
	s.Acls = v
	return s
}

func (s *DescribeAccessControlListsResponseBody) SetCount(v int32) *DescribeAccessControlListsResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeAccessControlListsResponseBody) SetPageNumber(v int32) *DescribeAccessControlListsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccessControlListsResponseBody) SetPageSize(v int32) *DescribeAccessControlListsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAccessControlListsResponseBody) SetRequestId(v string) *DescribeAccessControlListsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccessControlListsResponseBody) SetTotalCount(v int32) *DescribeAccessControlListsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAccessControlListsResponseBody) Validate() error {
	if s.Acls != nil {
		if err := s.Acls.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAccessControlListsResponseBodyAcls struct {
	Acl []*DescribeAccessControlListsResponseBodyAclsAcl `json:"Acl,omitempty" xml:"Acl,omitempty" type:"Repeated"`
}

func (s DescribeAccessControlListsResponseBodyAcls) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessControlListsResponseBodyAcls) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListsResponseBodyAcls) GetAcl() []*DescribeAccessControlListsResponseBodyAclsAcl {
	return s.Acl
}

func (s *DescribeAccessControlListsResponseBodyAcls) SetAcl(v []*DescribeAccessControlListsResponseBodyAclsAcl) *DescribeAccessControlListsResponseBodyAcls {
	s.Acl = v
	return s
}

func (s *DescribeAccessControlListsResponseBodyAcls) Validate() error {
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

type DescribeAccessControlListsResponseBodyAclsAcl struct {
	// The ACL ID.
	//
	// example:
	//
	// acl-bp1l0kk4gxce43k*****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The ACL name.
	//
	// example:
	//
	// rule1
	AclName *string `json:"AclName,omitempty" xml:"AclName,omitempty"`
	// The IP version that is used by the CLB instance associated with the ACL.
	//
	// example:
	//
	// ipv4
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	// The time when the CLB instance was created. The time follows the `YYYY-MM-DDThh:mm:ssZ` format.
	//
	// example:
	//
	// 2022-08-31T02:49:05Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-jfenfbp1lhl0****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of tags added to the network ACL. The value of this parameter must be a STRING list in the JSON format.
	Tags *DescribeAccessControlListsResponseBodyAclsAclTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s DescribeAccessControlListsResponseBodyAclsAcl) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessControlListsResponseBodyAclsAcl) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListsResponseBodyAclsAcl) GetAclId() *string {
	return s.AclId
}

func (s *DescribeAccessControlListsResponseBodyAclsAcl) GetAclName() *string {
	return s.AclName
}

func (s *DescribeAccessControlListsResponseBodyAclsAcl) GetAddressIPVersion() *string {
	return s.AddressIPVersion
}

func (s *DescribeAccessControlListsResponseBodyAclsAcl) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeAccessControlListsResponseBodyAclsAcl) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeAccessControlListsResponseBodyAclsAcl) GetTags() *DescribeAccessControlListsResponseBodyAclsAclTags {
	return s.Tags
}

func (s *DescribeAccessControlListsResponseBodyAclsAcl) SetAclId(v string) *DescribeAccessControlListsResponseBodyAclsAcl {
	s.AclId = &v
	return s
}

func (s *DescribeAccessControlListsResponseBodyAclsAcl) SetAclName(v string) *DescribeAccessControlListsResponseBodyAclsAcl {
	s.AclName = &v
	return s
}

func (s *DescribeAccessControlListsResponseBodyAclsAcl) SetAddressIPVersion(v string) *DescribeAccessControlListsResponseBodyAclsAcl {
	s.AddressIPVersion = &v
	return s
}

func (s *DescribeAccessControlListsResponseBodyAclsAcl) SetCreateTime(v string) *DescribeAccessControlListsResponseBodyAclsAcl {
	s.CreateTime = &v
	return s
}

func (s *DescribeAccessControlListsResponseBodyAclsAcl) SetResourceGroupId(v string) *DescribeAccessControlListsResponseBodyAclsAcl {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeAccessControlListsResponseBodyAclsAcl) SetTags(v *DescribeAccessControlListsResponseBodyAclsAclTags) *DescribeAccessControlListsResponseBodyAclsAcl {
	s.Tags = v
	return s
}

func (s *DescribeAccessControlListsResponseBodyAclsAcl) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAccessControlListsResponseBodyAclsAclTags struct {
	Tag []*DescribeAccessControlListsResponseBodyAclsAclTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeAccessControlListsResponseBodyAclsAclTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessControlListsResponseBodyAclsAclTags) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListsResponseBodyAclsAclTags) GetTag() []*DescribeAccessControlListsResponseBodyAclsAclTagsTag {
	return s.Tag
}

func (s *DescribeAccessControlListsResponseBodyAclsAclTags) SetTag(v []*DescribeAccessControlListsResponseBodyAclsAclTagsTag) *DescribeAccessControlListsResponseBodyAclsAclTags {
	s.Tag = v
	return s
}

func (s *DescribeAccessControlListsResponseBodyAclsAclTags) Validate() error {
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

type DescribeAccessControlListsResponseBodyAclsAclTagsTag struct {
	// The tag key.
	//
	// example:
	//
	// test
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// 1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeAccessControlListsResponseBodyAclsAclTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessControlListsResponseBodyAclsAclTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListsResponseBodyAclsAclTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeAccessControlListsResponseBodyAclsAclTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeAccessControlListsResponseBodyAclsAclTagsTag) SetTagKey(v string) *DescribeAccessControlListsResponseBodyAclsAclTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeAccessControlListsResponseBodyAclsAclTagsTag) SetTagValue(v string) *DescribeAccessControlListsResponseBodyAclsAclTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeAccessControlListsResponseBodyAclsAclTagsTag) Validate() error {
	return dara.Validate(s)
}
