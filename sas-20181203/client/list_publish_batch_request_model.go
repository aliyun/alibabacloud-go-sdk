// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublishBatchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchName(v string) *ListPublishBatchRequest
	GetBatchName() *string
	SetCurrentPage(v int32) *ListPublishBatchRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *ListPublishBatchRequest
	GetPageSize() *int32
	SetUpgradeVersion(v string) *ListPublishBatchRequest
	GetUpgradeVersion() *string
}

type ListPublishBatchRequest struct {
	// The name of the release batch.
	//
	// example:
	//
	// test
	BatchName *string `json:"BatchName,omitempty" xml:"BatchName,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The destination version of the Security Center agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.0.9
	UpgradeVersion *string `json:"UpgradeVersion,omitempty" xml:"UpgradeVersion,omitempty"`
}

func (s ListPublishBatchRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPublishBatchRequest) GoString() string {
	return s.String()
}

func (s *ListPublishBatchRequest) GetBatchName() *string {
	return s.BatchName
}

func (s *ListPublishBatchRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListPublishBatchRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPublishBatchRequest) GetUpgradeVersion() *string {
	return s.UpgradeVersion
}

func (s *ListPublishBatchRequest) SetBatchName(v string) *ListPublishBatchRequest {
	s.BatchName = &v
	return s
}

func (s *ListPublishBatchRequest) SetCurrentPage(v int32) *ListPublishBatchRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListPublishBatchRequest) SetPageSize(v int32) *ListPublishBatchRequest {
	s.PageSize = &v
	return s
}

func (s *ListPublishBatchRequest) SetUpgradeVersion(v string) *ListPublishBatchRequest {
	s.UpgradeVersion = &v
	return s
}

func (s *ListPublishBatchRequest) Validate() error {
	return dara.Validate(s)
}
