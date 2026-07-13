// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceCategoryId(v string) *GetResourceCategoryRequest
	GetResourceCategoryId() *string
}

type GetResourceCategoryRequest struct {
	// Resource category ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rc-123****7890
	ResourceCategoryId *string `json:"ResourceCategoryId,omitempty" xml:"ResourceCategoryId,omitempty"`
}

func (s GetResourceCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceCategoryRequest) GoString() string {
	return s.String()
}

func (s *GetResourceCategoryRequest) GetResourceCategoryId() *string {
	return s.ResourceCategoryId
}

func (s *GetResourceCategoryRequest) SetResourceCategoryId(v string) *GetResourceCategoryRequest {
	s.ResourceCategoryId = &v
	return s
}

func (s *GetResourceCategoryRequest) Validate() error {
	return dara.Validate(s)
}
