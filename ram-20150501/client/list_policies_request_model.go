// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMarker(v string) *ListPoliciesRequest
	GetMarker() *string
	SetMaxItems(v int32) *ListPoliciesRequest
	GetMaxItems() *int32
	SetPolicyType(v string) *ListPoliciesRequest
	GetPolicyType() *string
	SetTag(v []*ListPoliciesRequestTag) *ListPoliciesRequest
	GetTag() []*ListPoliciesRequestTag
}

type ListPoliciesRequest struct {
	// The `marker`. If part of a previous response is truncated, you can use this parameter to obtain the truncated part.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The number of entries to return. If a response is truncated because it reaches the value of `MaxItems`, the value of `IsTruncated` will be `true`.
	//
	// Valid values: 1 to 1000. Default value: 100.
	//
	// example:
	//
	// 100
	MaxItems *int32 `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
	// The type of the policies. Valid values: `System` and `Custom`. If you do not specify the parameter, all policies are returned.``
	//
	// example:
	//
	// System
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The tags.
	Tag []*ListPoliciesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListPoliciesRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListPoliciesRequest) GetMaxItems() *int32 {
	return s.MaxItems
}

func (s *ListPoliciesRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListPoliciesRequest) GetTag() []*ListPoliciesRequestTag {
	return s.Tag
}

func (s *ListPoliciesRequest) SetMarker(v string) *ListPoliciesRequest {
	s.Marker = &v
	return s
}

func (s *ListPoliciesRequest) SetMaxItems(v int32) *ListPoliciesRequest {
	s.MaxItems = &v
	return s
}

func (s *ListPoliciesRequest) SetPolicyType(v string) *ListPoliciesRequest {
	s.PolicyType = &v
	return s
}

func (s *ListPoliciesRequest) SetTag(v []*ListPoliciesRequestTag) *ListPoliciesRequest {
	s.Tag = v
	return s
}

func (s *ListPoliciesRequest) Validate() error {
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

type ListPoliciesRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// owner
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// alice
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListPoliciesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesRequestTag) GoString() string {
	return s.String()
}

func (s *ListPoliciesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListPoliciesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListPoliciesRequestTag) SetKey(v string) *ListPoliciesRequestTag {
	s.Key = &v
	return s
}

func (s *ListPoliciesRequestTag) SetValue(v string) *ListPoliciesRequestTag {
	s.Value = &v
	return s
}

func (s *ListPoliciesRequestTag) Validate() error {
	return dara.Validate(s)
}
