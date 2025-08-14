// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSnapshotUrlsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetSnapshotUrlsRequest
	GetJobId() *string
	SetOrderBy(v string) *GetSnapshotUrlsRequest
	GetOrderBy() *string
	SetPageNumber(v int32) *GetSnapshotUrlsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetSnapshotUrlsRequest
	GetPageSize() *int32
	SetTimeout(v int64) *GetSnapshotUrlsRequest
	GetTimeout() *int64
}

type GetSnapshotUrlsRequest struct {
	// The job ID.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The order that you use to sort the query results. Valid values: Asc and Desc.
	//
	// - Asc
	//
	// - Desc
	//
	// example:
	//
	// Asc
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 30. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The authentication timeout period. Unit: seconds Default value: 3600. Maximum value: 129600 (36 hours).
	//
	// example:
	//
	// 3600
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s GetSnapshotUrlsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSnapshotUrlsRequest) GoString() string {
	return s.String()
}

func (s *GetSnapshotUrlsRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetSnapshotUrlsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *GetSnapshotUrlsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetSnapshotUrlsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSnapshotUrlsRequest) GetTimeout() *int64 {
	return s.Timeout
}

func (s *GetSnapshotUrlsRequest) SetJobId(v string) *GetSnapshotUrlsRequest {
	s.JobId = &v
	return s
}

func (s *GetSnapshotUrlsRequest) SetOrderBy(v string) *GetSnapshotUrlsRequest {
	s.OrderBy = &v
	return s
}

func (s *GetSnapshotUrlsRequest) SetPageNumber(v int32) *GetSnapshotUrlsRequest {
	s.PageNumber = &v
	return s
}

func (s *GetSnapshotUrlsRequest) SetPageSize(v int32) *GetSnapshotUrlsRequest {
	s.PageSize = &v
	return s
}

func (s *GetSnapshotUrlsRequest) SetTimeout(v int64) *GetSnapshotUrlsRequest {
	s.Timeout = &v
	return s
}

func (s *GetSnapshotUrlsRequest) Validate() error {
	return dara.Validate(s)
}
