// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAclsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcls(v []*ListAclsResponseBodyAcls) *ListAclsResponseBody
	GetAcls() []*ListAclsResponseBodyAcls
	SetMaxResults(v int32) *ListAclsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListAclsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAclsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAclsResponseBody
	GetTotalCount() *int32
}

type ListAclsResponseBody struct {
	// A list of ACLs.
	Acls []*ListAclsResponseBodyAcls `json:"Acls,omitempty" xml:"Acls,omitempty" type:"Repeated"`
	// The maximum number of network ACLs returned. This parameter is optional. Valid values: **1*	- to **100**. If this parameter is not set, the default value **20*	- is returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// 	- If **NextToken*	- is empty, it indicates that no next query is to be sent.
	//
	// 	- If **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// FFmyTO70t****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 593B0448-D13E-4C56-AC0D-FDF0FDE0E9A3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAclsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAclsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAclsResponseBody) GetAcls() []*ListAclsResponseBodyAcls {
	return s.Acls
}

func (s *ListAclsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAclsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAclsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAclsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAclsResponseBody) SetAcls(v []*ListAclsResponseBodyAcls) *ListAclsResponseBody {
	s.Acls = v
	return s
}

func (s *ListAclsResponseBody) SetMaxResults(v int32) *ListAclsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAclsResponseBody) SetNextToken(v string) *ListAclsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAclsResponseBody) SetRequestId(v string) *ListAclsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAclsResponseBody) SetTotalCount(v int32) *ListAclsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAclsResponseBody) Validate() error {
	if s.Acls != nil {
		for _, item := range s.Acls {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAclsResponseBodyAcls struct {
	// The ACL ID.
	//
	// example:
	//
	// nacl-hp34s2h0xx1ht4nwo****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The name of the ACL.
	//
	// example:
	//
	// test-acl
	AclName *string `json:"AclName,omitempty" xml:"AclName,omitempty"`
	// The status of the ACL. Valid values:
	//
	// 	- **Creating**: The network ACL is being created.
	//
	// 	- **Available**: The network ACL is available.
	//
	// 	- **Configuring**
	//
	// example:
	//
	// Available
	AclStatus *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	// The IP version of the ACL. Only **IPv4*	- may be returned.
	//
	// example:
	//
	// IPv4
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	// Indicates whether configuration management is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	ConfigManagedEnabled *bool `json:"ConfigManagedEnabled,omitempty" xml:"ConfigManagedEnabled,omitempty"`
	// The time when the ACL was created. The follows the `YYYY-MM-DDThh:mm:ssZ` format.
	//
	// example:
	//
	// 2023-02-15T07:37:33Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-atstuj3rtopty****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tags []*ListAclsResponseBodyAclsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListAclsResponseBodyAcls) String() string {
	return dara.Prettify(s)
}

func (s ListAclsResponseBodyAcls) GoString() string {
	return s.String()
}

func (s *ListAclsResponseBodyAcls) GetAclId() *string {
	return s.AclId
}

func (s *ListAclsResponseBodyAcls) GetAclName() *string {
	return s.AclName
}

func (s *ListAclsResponseBodyAcls) GetAclStatus() *string {
	return s.AclStatus
}

func (s *ListAclsResponseBodyAcls) GetAddressIPVersion() *string {
	return s.AddressIPVersion
}

func (s *ListAclsResponseBodyAcls) GetConfigManagedEnabled() *bool {
	return s.ConfigManagedEnabled
}

func (s *ListAclsResponseBodyAcls) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAclsResponseBodyAcls) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListAclsResponseBodyAcls) GetTags() []*ListAclsResponseBodyAclsTags {
	return s.Tags
}

func (s *ListAclsResponseBodyAcls) SetAclId(v string) *ListAclsResponseBodyAcls {
	s.AclId = &v
	return s
}

func (s *ListAclsResponseBodyAcls) SetAclName(v string) *ListAclsResponseBodyAcls {
	s.AclName = &v
	return s
}

func (s *ListAclsResponseBodyAcls) SetAclStatus(v string) *ListAclsResponseBodyAcls {
	s.AclStatus = &v
	return s
}

func (s *ListAclsResponseBodyAcls) SetAddressIPVersion(v string) *ListAclsResponseBodyAcls {
	s.AddressIPVersion = &v
	return s
}

func (s *ListAclsResponseBodyAcls) SetConfigManagedEnabled(v bool) *ListAclsResponseBodyAcls {
	s.ConfigManagedEnabled = &v
	return s
}

func (s *ListAclsResponseBodyAcls) SetCreateTime(v string) *ListAclsResponseBodyAcls {
	s.CreateTime = &v
	return s
}

func (s *ListAclsResponseBodyAcls) SetResourceGroupId(v string) *ListAclsResponseBodyAcls {
	s.ResourceGroupId = &v
	return s
}

func (s *ListAclsResponseBodyAcls) SetTags(v []*ListAclsResponseBodyAclsTags) *ListAclsResponseBodyAcls {
	s.Tags = v
	return s
}

func (s *ListAclsResponseBodyAcls) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAclsResponseBodyAclsTags struct {
	// The tag key. The tag key can be up to 128 characters in length. It cannot start with aliyun or acs: and cannot contain http:// or https://.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 128 characters in length. It cannot start with aliyun or acs: and cannot contain http:// or https://.
	//
	// example:
	//
	// product
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAclsResponseBodyAclsTags) String() string {
	return dara.Prettify(s)
}

func (s ListAclsResponseBodyAclsTags) GoString() string {
	return s.String()
}

func (s *ListAclsResponseBodyAclsTags) GetKey() *string {
	return s.Key
}

func (s *ListAclsResponseBodyAclsTags) GetValue() *string {
	return s.Value
}

func (s *ListAclsResponseBodyAclsTags) SetKey(v string) *ListAclsResponseBodyAclsTags {
	s.Key = &v
	return s
}

func (s *ListAclsResponseBodyAclsTags) SetValue(v string) *ListAclsResponseBodyAclsTags {
	s.Value = &v
	return s
}

func (s *ListAclsResponseBodyAclsTags) Validate() error {
	return dara.Validate(s)
}
