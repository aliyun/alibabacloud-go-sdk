// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelTemplateResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModelTemplateId(v string) *ListModelTemplateResourceGroupRequest
	GetModelTemplateId() *string
	SetPageNumber(v int32) *ListModelTemplateResourceGroupRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModelTemplateResourceGroupRequest
	GetPageSize() *int32
	SetResourceGroupIds(v []*string) *ListModelTemplateResourceGroupRequest
	GetResourceGroupIds() []*string
}

type ListModelTemplateResourceGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mt-xxxxx
	ModelTemplateId *string `json:"ModelTemplateId,omitempty" xml:"ModelTemplateId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize         *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceGroupIds []*string `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty" type:"Repeated"`
}

func (s ListModelTemplateResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelTemplateResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ListModelTemplateResourceGroupRequest) GetModelTemplateId() *string {
	return s.ModelTemplateId
}

func (s *ListModelTemplateResourceGroupRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModelTemplateResourceGroupRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModelTemplateResourceGroupRequest) GetResourceGroupIds() []*string {
	return s.ResourceGroupIds
}

func (s *ListModelTemplateResourceGroupRequest) SetModelTemplateId(v string) *ListModelTemplateResourceGroupRequest {
	s.ModelTemplateId = &v
	return s
}

func (s *ListModelTemplateResourceGroupRequest) SetPageNumber(v int32) *ListModelTemplateResourceGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModelTemplateResourceGroupRequest) SetPageSize(v int32) *ListModelTemplateResourceGroupRequest {
	s.PageSize = &v
	return s
}

func (s *ListModelTemplateResourceGroupRequest) SetResourceGroupIds(v []*string) *ListModelTemplateResourceGroupRequest {
	s.ResourceGroupIds = v
	return s
}

func (s *ListModelTemplateResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
