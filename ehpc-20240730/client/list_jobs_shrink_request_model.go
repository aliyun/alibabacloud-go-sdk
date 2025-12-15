// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListJobsShrinkRequest
	GetClusterId() *string
	SetJobFilterShrink(v string) *ListJobsShrinkRequest
	GetJobFilterShrink() *string
	SetPageNumber(v string) *ListJobsShrinkRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListJobsShrinkRequest
	GetPageSize() *string
}

type ListJobsShrinkRequest struct {
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-csbua72***
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The job filter information.
	JobFilterShrink *string `json:"JobFilter,omitempty" xml:"JobFilter,omitempty"`
	// The page number of the page to return.
	//
	// 	- Pages start from page 1.
	//
	// 	- Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// 	- Maximum value: 50.
	//
	// 	- Default value: 10
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListJobsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListJobsShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListJobsShrinkRequest) GetJobFilterShrink() *string {
	return s.JobFilterShrink
}

func (s *ListJobsShrinkRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListJobsShrinkRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListJobsShrinkRequest) SetClusterId(v string) *ListJobsShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *ListJobsShrinkRequest) SetJobFilterShrink(v string) *ListJobsShrinkRequest {
	s.JobFilterShrink = &v
	return s
}

func (s *ListJobsShrinkRequest) SetPageNumber(v string) *ListJobsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobsShrinkRequest) SetPageSize(v string) *ListJobsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
