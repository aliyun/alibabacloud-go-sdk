// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserBasicInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMarker(v string) *ListUserBasicInfosRequest
	GetMarker() *string
	SetMaxItems(v int32) *ListUserBasicInfosRequest
	GetMaxItems() *int32
	SetStatus(v string) *ListUserBasicInfosRequest
	GetStatus() *string
	SetTag(v []*ListUserBasicInfosRequestTag) *ListUserBasicInfosRequest
	GetTag() []*ListUserBasicInfosRequestTag
}

type ListUserBasicInfosRequest struct {
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
	// The status of the RAM users that you want to query. Valid values: active, freeze, and active,freeze. If you set the value to active,freeze, RAM users in both active and freeze states are queried. If you leave the parameter empty, the value active is used by default. If the Tag parameter is specified, you cannot specify the Status parameter. In this case, RAM users in both states are queried.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tag []*ListUserBasicInfosRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListUserBasicInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserBasicInfosRequest) GoString() string {
	return s.String()
}

func (s *ListUserBasicInfosRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListUserBasicInfosRequest) GetMaxItems() *int32 {
	return s.MaxItems
}

func (s *ListUserBasicInfosRequest) GetStatus() *string {
	return s.Status
}

func (s *ListUserBasicInfosRequest) GetTag() []*ListUserBasicInfosRequestTag {
	return s.Tag
}

func (s *ListUserBasicInfosRequest) SetMarker(v string) *ListUserBasicInfosRequest {
	s.Marker = &v
	return s
}

func (s *ListUserBasicInfosRequest) SetMaxItems(v int32) *ListUserBasicInfosRequest {
	s.MaxItems = &v
	return s
}

func (s *ListUserBasicInfosRequest) SetStatus(v string) *ListUserBasicInfosRequest {
	s.Status = &v
	return s
}

func (s *ListUserBasicInfosRequest) SetTag(v []*ListUserBasicInfosRequestTag) *ListUserBasicInfosRequest {
	s.Tag = v
	return s
}

func (s *ListUserBasicInfosRequest) Validate() error {
	return dara.Validate(s)
}

type ListUserBasicInfosRequestTag struct {
	// The key of tag N.
	//
	// Valid values of N: 1 to 20. N must be consecutive.
	//
	// example:
	//
	// operator
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.
	//
	// Valid values of N: 1 to 20. N must be consecutive.
	//
	// example:
	//
	// alice
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListUserBasicInfosRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListUserBasicInfosRequestTag) GoString() string {
	return s.String()
}

func (s *ListUserBasicInfosRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListUserBasicInfosRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListUserBasicInfosRequestTag) SetKey(v string) *ListUserBasicInfosRequestTag {
	s.Key = &v
	return s
}

func (s *ListUserBasicInfosRequestTag) SetValue(v string) *ListUserBasicInfosRequestTag {
	s.Value = &v
	return s
}

func (s *ListUserBasicInfosRequestTag) Validate() error {
	return dara.Validate(s)
}
