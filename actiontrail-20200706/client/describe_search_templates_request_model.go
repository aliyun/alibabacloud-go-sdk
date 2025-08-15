// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSearchTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeSearchTemplatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSearchTemplatesRequest
	GetPageSize() *int32
	SetSceneId(v string) *DescribeSearchTemplatesRequest
	GetSceneId() *string
}

type DescribeSearchTemplatesRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sc-lpYrjKouRfy3MK-wteJW_Q
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s DescribeSearchTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSearchTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSearchTemplatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSearchTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSearchTemplatesRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *DescribeSearchTemplatesRequest) SetPageNumber(v int32) *DescribeSearchTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSearchTemplatesRequest) SetPageSize(v int32) *DescribeSearchTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSearchTemplatesRequest) SetSceneId(v string) *DescribeSearchTemplatesRequest {
	s.SceneId = &v
	return s
}

func (s *DescribeSearchTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
