// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListTagResourcesRequest
	GetLang() *string
	SetNextToken(v string) *ListTagResourcesRequest
	GetNextToken() *string
	SetResourceId(v []*string) *ListTagResourcesRequest
	GetResourceId() []*string
	SetResourceType(v string) *ListTagResourcesRequest
	GetResourceType() *string
	SetSize(v int32) *ListTagResourcesRequest
	GetSize() *int32
	SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest
	GetTag() []*ListTagResourcesRequestTag
}

type ListTagResourcesRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// Default value: en.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// 23423****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource IDs, which are zone IDs. You can specify up to 50 zone IDs.
	//
	// example:
	//
	// 97fe9321a476d0861f624d3f738dcc38
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type. Valid value: ZONE.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZONE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The number of entries per page. Maximum value: 200. Default value: 20.
	//
	// example:
	//
	// 20
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The tags added to the resources.
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) GetLang() *string {
	return s.Lang
}

func (s *ListTagResourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagResourcesRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *ListTagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagResourcesRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListTagResourcesRequest) GetTag() []*ListTagResourcesRequestTag {
	return s.Tag
}

func (s *ListTagResourcesRequest) SetLang(v string) *ListTagResourcesRequest {
	s.Lang = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetSize(v int32) *ListTagResourcesRequest {
	s.Size = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

func (s *ListTagResourcesRequest) Validate() error {
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

type ListTagResourcesRequestTag struct {
	// The key of tag N added to the resource.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N added to the resource.
	//
	// example:
	//
	// daily
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListTagResourcesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

func (s *ListTagResourcesRequestTag) Validate() error {
	return dara.Validate(s)
}
