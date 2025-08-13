// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *CreateTagsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateTagsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateTagsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateTagsRequest
	GetResourceOwnerAccount() *string
	SetTagKeyValueParamList(v []*CreateTagsRequestTagKeyValueParamList) *CreateTagsRequest
	GetTagKeyValueParamList() []*CreateTagsRequestTagKeyValueParamList
}

type CreateTagsRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// > Only `cn-hangzhou` is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The information about the tags.
	//
	// This parameter is required.
	TagKeyValueParamList []*CreateTagsRequestTagKeyValueParamList `json:"TagKeyValueParamList,omitempty" xml:"TagKeyValueParamList,omitempty" type:"Repeated"`
}

func (s CreateTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTagsRequest) GoString() string {
	return s.String()
}

func (s *CreateTagsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateTagsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateTagsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTagsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateTagsRequest) GetTagKeyValueParamList() []*CreateTagsRequestTagKeyValueParamList {
	return s.TagKeyValueParamList
}

func (s *CreateTagsRequest) SetOwnerAccount(v string) *CreateTagsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateTagsRequest) SetOwnerId(v int64) *CreateTagsRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTagsRequest) SetRegionId(v string) *CreateTagsRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTagsRequest) SetResourceOwnerAccount(v string) *CreateTagsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTagsRequest) SetTagKeyValueParamList(v []*CreateTagsRequestTagKeyValueParamList) *CreateTagsRequest {
	s.TagKeyValueParamList = v
	return s
}

func (s *CreateTagsRequest) Validate() error {
	return dara.Validate(s)
}

type CreateTagsRequestTagKeyValueParamList struct {
	// The description of the key for tag N.
	//
	// Valid values of N: 1 to 10.
	//
	// example:
	//
	// Business environment
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The value of tag N.
	//
	// The tag value can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. The tag key cannot contain `http://` or `https://`.
	//
	// Valid values of N: 1 to 10.
	//
	// This parameter is required.
	//
	// example:
	//
	// Environment
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The information about the tag values.
	TagValueParamList []*CreateTagsRequestTagKeyValueParamListTagValueParamList `json:"TagValueParamList,omitempty" xml:"TagValueParamList,omitempty" type:"Repeated"`
}

func (s CreateTagsRequestTagKeyValueParamList) String() string {
	return dara.Prettify(s)
}

func (s CreateTagsRequestTagKeyValueParamList) GoString() string {
	return s.String()
}

func (s *CreateTagsRequestTagKeyValueParamList) GetDescription() *string {
	return s.Description
}

func (s *CreateTagsRequestTagKeyValueParamList) GetKey() *string {
	return s.Key
}

func (s *CreateTagsRequestTagKeyValueParamList) GetTagValueParamList() []*CreateTagsRequestTagKeyValueParamListTagValueParamList {
	return s.TagValueParamList
}

func (s *CreateTagsRequestTagKeyValueParamList) SetDescription(v string) *CreateTagsRequestTagKeyValueParamList {
	s.Description = &v
	return s
}

func (s *CreateTagsRequestTagKeyValueParamList) SetKey(v string) *CreateTagsRequestTagKeyValueParamList {
	s.Key = &v
	return s
}

func (s *CreateTagsRequestTagKeyValueParamList) SetTagValueParamList(v []*CreateTagsRequestTagKeyValueParamListTagValueParamList) *CreateTagsRequestTagKeyValueParamList {
	s.TagValueParamList = v
	return s
}

func (s *CreateTagsRequestTagKeyValueParamList) Validate() error {
	return dara.Validate(s)
}

type CreateTagsRequestTagKeyValueParamListTagValueParamList struct {
	// The description of the value for tag N.
	//
	// Valid values of N: 1 to 10.
	//
	// example:
	//
	// Test environment
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The value of tag N.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// Valid values of N: 1 to 10.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTagsRequestTagKeyValueParamListTagValueParamList) String() string {
	return dara.Prettify(s)
}

func (s CreateTagsRequestTagKeyValueParamListTagValueParamList) GoString() string {
	return s.String()
}

func (s *CreateTagsRequestTagKeyValueParamListTagValueParamList) GetDescription() *string {
	return s.Description
}

func (s *CreateTagsRequestTagKeyValueParamListTagValueParamList) GetValue() *string {
	return s.Value
}

func (s *CreateTagsRequestTagKeyValueParamListTagValueParamList) SetDescription(v string) *CreateTagsRequestTagKeyValueParamListTagValueParamList {
	s.Description = &v
	return s
}

func (s *CreateTagsRequestTagKeyValueParamListTagValueParamList) SetValue(v string) *CreateTagsRequestTagKeyValueParamListTagValueParamList {
	s.Value = &v
	return s
}

func (s *CreateTagsRequestTagKeyValueParamListTagValueParamList) Validate() error {
	return dara.Validate(s)
}
