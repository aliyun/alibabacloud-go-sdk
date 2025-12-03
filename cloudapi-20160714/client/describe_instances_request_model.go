// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableTagAuthorization(v bool) *DescribeInstancesRequest
	GetEnableTagAuthorization() *bool
	SetInstanceId(v string) *DescribeInstancesRequest
	GetInstanceId() *string
	SetInstanceType(v string) *DescribeInstancesRequest
	GetInstanceType() *string
	SetLanguage(v string) *DescribeInstancesRequest
	GetLanguage() *string
	SetSecurityToken(v string) *DescribeInstancesRequest
	GetSecurityToken() *string
	SetTag(v []*DescribeInstancesRequestTag) *DescribeInstancesRequest
	GetTag() []*DescribeInstancesRequestTag
}

type DescribeInstancesRequest struct {
	// Specifies whether tag authorization is enabled.
	//
	// example:
	//
	// false
	EnableTagAuthorization *bool `json:"EnableTagAuthorization,omitempty" xml:"EnableTagAuthorization,omitempty"`
	// The instance ID. If you do not specify this parameter, all instances are returned.
	//
	// example:
	//
	// api-shared-vpc-001
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The language in which you want the description of the system policy to be returned. Valid values:
	//
	// 	- en: English
	//
	// 	- zh: Chinese
	//
	// 	- ja: Japanese
	//
	// example:
	//
	// zh
	Language      *string `json:"Language,omitempty" xml:"Language,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The tag that is bound to the instance.
	Tag []*DescribeInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequest) GetEnableTagAuthorization() *bool {
	return s.EnableTagAuthorization
}

func (s *DescribeInstancesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstancesRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeInstancesRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeInstancesRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeInstancesRequest) GetTag() []*DescribeInstancesRequestTag {
	return s.Tag
}

func (s *DescribeInstancesRequest) SetEnableTagAuthorization(v bool) *DescribeInstancesRequest {
	s.EnableTagAuthorization = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceId(v string) *DescribeInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceType(v string) *DescribeInstancesRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstancesRequest) SetLanguage(v string) *DescribeInstancesRequest {
	s.Language = &v
	return s
}

func (s *DescribeInstancesRequest) SetSecurityToken(v string) *DescribeInstancesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeInstancesRequest) SetTag(v []*DescribeInstancesRequestTag) *DescribeInstancesRequest {
	s.Tag = v
	return s
}

func (s *DescribeInstancesRequest) Validate() error {
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

type DescribeInstancesRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstancesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeInstancesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeInstancesRequestTag) SetKey(v string) *DescribeInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeInstancesRequestTag) SetValue(v string) *DescribeInstancesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeInstancesRequestTag) Validate() error {
	return dara.Validate(s)
}
