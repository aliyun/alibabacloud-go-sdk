// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceCategoryId(v string) *UpdateResourceCategoryRequest
	GetResourceCategoryId() *string
	SetResourceCategoryName(v string) *UpdateResourceCategoryRequest
	GetResourceCategoryName() *string
	SetResourceMatcher(v string) *UpdateResourceCategoryRequest
	GetResourceMatcher() *string
}

type UpdateResourceCategoryRequest struct {
	// Resource category ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rc-123****7890
	ResourceCategoryId *string `json:"ResourceCategoryId,omitempty" xml:"ResourceCategoryId,omitempty"`
	// Resource category name.
	//
	// example:
	//
	// My***ResourceCategory
	ResourceCategoryName *string `json:"ResourceCategoryName,omitempty" xml:"ResourceCategoryName,omitempty"`
	// Resource matcher.
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
}

func (s UpdateResourceCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceCategoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceCategoryRequest) GetResourceCategoryId() *string {
	return s.ResourceCategoryId
}

func (s *UpdateResourceCategoryRequest) GetResourceCategoryName() *string {
	return s.ResourceCategoryName
}

func (s *UpdateResourceCategoryRequest) GetResourceMatcher() *string {
	return s.ResourceMatcher
}

func (s *UpdateResourceCategoryRequest) SetResourceCategoryId(v string) *UpdateResourceCategoryRequest {
	s.ResourceCategoryId = &v
	return s
}

func (s *UpdateResourceCategoryRequest) SetResourceCategoryName(v string) *UpdateResourceCategoryRequest {
	s.ResourceCategoryName = &v
	return s
}

func (s *UpdateResourceCategoryRequest) SetResourceMatcher(v string) *UpdateResourceCategoryRequest {
	s.ResourceMatcher = &v
	return s
}

func (s *UpdateResourceCategoryRequest) Validate() error {
	return dara.Validate(s)
}
