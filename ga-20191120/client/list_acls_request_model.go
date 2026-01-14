// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAclsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclIds(v []*string) *ListAclsRequest
	GetAclIds() []*string
	SetAclName(v string) *ListAclsRequest
	GetAclName() *string
	SetClientToken(v string) *ListAclsRequest
	GetClientToken() *string
	SetMaxResults(v int32) *ListAclsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAclsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListAclsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListAclsRequest
	GetResourceGroupId() *string
	SetTag(v []*ListAclsRequestTag) *ListAclsRequest
	GetTag() []*ListAclsRequestTag
}

type ListAclsRequest struct {
	// The ACL IDs. You can specify a maximum of 20 ACL IDs in each request.
	AclIds []*string `json:"AclIds,omitempty" xml:"AclIds,omitempty" type:"Repeated"`
	// The ACL name. You can specify a maximum of ACL names in each request.
	//
	// The name must be 1 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// test-acl
	AclName *string `json:"AclName,omitempty" xml:"AclName,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The number of entries per page. Valid values: **1*	- to **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// 	- If this is your first and only query, ignore this parameter.
	//
	// 	- If a subsequent query is to be performed, set the parameter to the value of NextToken that is returned from the last call.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region where the Global Accelerator (GA) instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek2lgw4evw****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tag of the ACL.
	Tag []*ListAclsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListAclsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAclsRequest) GoString() string {
	return s.String()
}

func (s *ListAclsRequest) GetAclIds() []*string {
	return s.AclIds
}

func (s *ListAclsRequest) GetAclName() *string {
	return s.AclName
}

func (s *ListAclsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListAclsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAclsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAclsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAclsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListAclsRequest) GetTag() []*ListAclsRequestTag {
	return s.Tag
}

func (s *ListAclsRequest) SetAclIds(v []*string) *ListAclsRequest {
	s.AclIds = v
	return s
}

func (s *ListAclsRequest) SetAclName(v string) *ListAclsRequest {
	s.AclName = &v
	return s
}

func (s *ListAclsRequest) SetClientToken(v string) *ListAclsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListAclsRequest) SetMaxResults(v int32) *ListAclsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAclsRequest) SetNextToken(v string) *ListAclsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAclsRequest) SetRegionId(v string) *ListAclsRequest {
	s.RegionId = &v
	return s
}

func (s *ListAclsRequest) SetResourceGroupId(v string) *ListAclsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListAclsRequest) SetTag(v []*ListAclsRequestTag) *ListAclsRequest {
	s.Tag = v
	return s
}

func (s *ListAclsRequest) Validate() error {
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

type ListAclsRequestTag struct {
	// The tag key of the ACL. The value of this parameter cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length, and cannot contain `http://` or `https://`. The tag key cannot start with `aliyun` or `acs:`.
	//
	// You can specify a maximum of 20 tag keys.
	//
	// example:
	//
	// tag-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the ACL. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length, and cannot contain `http://` or `https://`. The tag value cannot start with `aliyun` or `acs:`.
	//
	// You can specify a maximum of 20 tag values.
	//
	// example:
	//
	// tag-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAclsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListAclsRequestTag) GoString() string {
	return s.String()
}

func (s *ListAclsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListAclsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListAclsRequestTag) SetKey(v string) *ListAclsRequestTag {
	s.Key = &v
	return s
}

func (s *ListAclsRequestTag) SetValue(v string) *ListAclsRequestTag {
	s.Value = &v
	return s
}

func (s *ListAclsRequestTag) Validate() error {
	return dara.Validate(s)
}
