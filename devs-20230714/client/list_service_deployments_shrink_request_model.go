// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceDeploymentsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListServiceDeploymentsShrinkRequest
	GetKeyword() *string
	SetLabelSelectorShrink(v string) *ListServiceDeploymentsShrinkRequest
	GetLabelSelectorShrink() *string
	SetPageNumber(v int64) *ListServiceDeploymentsShrinkRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListServiceDeploymentsShrinkRequest
	GetPageSize() *int64
}

type ListServiceDeploymentsShrinkRequest struct {
	// example:
	//
	// demo
	Keyword             *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	LabelSelectorShrink *string `json:"labelSelector,omitempty" xml:"labelSelector,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListServiceDeploymentsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServiceDeploymentsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListServiceDeploymentsShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListServiceDeploymentsShrinkRequest) GetLabelSelectorShrink() *string {
	return s.LabelSelectorShrink
}

func (s *ListServiceDeploymentsShrinkRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListServiceDeploymentsShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListServiceDeploymentsShrinkRequest) SetKeyword(v string) *ListServiceDeploymentsShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListServiceDeploymentsShrinkRequest) SetLabelSelectorShrink(v string) *ListServiceDeploymentsShrinkRequest {
	s.LabelSelectorShrink = &v
	return s
}

func (s *ListServiceDeploymentsShrinkRequest) SetPageNumber(v int64) *ListServiceDeploymentsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListServiceDeploymentsShrinkRequest) SetPageSize(v int64) *ListServiceDeploymentsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListServiceDeploymentsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
