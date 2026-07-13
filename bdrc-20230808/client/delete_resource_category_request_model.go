// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceCategoryId(v string) *DeleteResourceCategoryRequest
	GetResourceCategoryId() *string
}

type DeleteResourceCategoryRequest struct {
	// Resource category ID
	//
	// This parameter is required.
	//
	// example:
	//
	// rc-123****7890
	ResourceCategoryId *string `json:"ResourceCategoryId,omitempty" xml:"ResourceCategoryId,omitempty"`
}

func (s DeleteResourceCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceCategoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourceCategoryRequest) GetResourceCategoryId() *string {
	return s.ResourceCategoryId
}

func (s *DeleteResourceCategoryRequest) SetResourceCategoryId(v string) *DeleteResourceCategoryRequest {
	s.ResourceCategoryId = &v
	return s
}

func (s *DeleteResourceCategoryRequest) Validate() error {
	return dara.Validate(s)
}
