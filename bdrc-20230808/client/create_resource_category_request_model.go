// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceCategoryName(v string) *CreateResourceCategoryRequest
	GetResourceCategoryName() *string
	SetResourceMatcher(v string) *CreateResourceCategoryRequest
	GetResourceMatcher() *string
	SetResourceType(v string) *CreateResourceCategoryRequest
	GetResourceType() *string
}

type CreateResourceCategoryRequest struct {
	// Resource category name.
	//
	// This parameter is required.
	//
	// example:
	//
	// My***ResourceCategory
	ResourceCategoryName *string `json:"ResourceCategoryName,omitempty" xml:"ResourceCategoryName,omitempty"`
	// Resource matcher.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//    "type": "BOOL",
	//
	//    "operator": "AND",
	//
	//    "values": [
	//
	//      {
	//
	//        "type": "TAG",
	//
	//        "operator": "EQUAL",
	//
	//        "key": "key0",
	//
	//        "values": [
	//
	//          "value0"
	//
	//        ]
	//
	//      },
	//
	//      {
	//
	//        "type": "TAG",
	//
	//        "operator": "KEY_EXIST",
	//
	//        "key": "key1"
	//
	//      },
	//
	//      {
	//
	//        "type": "TAG",
	//
	//        "operator": "KEY_EXIST_NOT_EQUAL",
	//
	//        "key": "key2",
	//
	//        "values": [
	//
	//          "value2-wrong"
	//
	//        ]
	//
	//      }
	//
	//    ]
	//
	//  }
	ResourceMatcher *string `json:"ResourceMatcher,omitempty" xml:"ResourceMatcher,omitempty"`
	// Resource type.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s CreateResourceCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceCategoryRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceCategoryRequest) GetResourceCategoryName() *string {
	return s.ResourceCategoryName
}

func (s *CreateResourceCategoryRequest) GetResourceMatcher() *string {
	return s.ResourceMatcher
}

func (s *CreateResourceCategoryRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateResourceCategoryRequest) SetResourceCategoryName(v string) *CreateResourceCategoryRequest {
	s.ResourceCategoryName = &v
	return s
}

func (s *CreateResourceCategoryRequest) SetResourceMatcher(v string) *CreateResourceCategoryRequest {
	s.ResourceMatcher = &v
	return s
}

func (s *CreateResourceCategoryRequest) SetResourceType(v string) *CreateResourceCategoryRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateResourceCategoryRequest) Validate() error {
	return dara.Validate(s)
}
