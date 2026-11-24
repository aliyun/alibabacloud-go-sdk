// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInvoiceForIsvRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTimeEnd(v string) *DescribeInvoiceForIsvRequest
	GetCreateTimeEnd() *string
	SetCreateTimeStart(v string) *DescribeInvoiceForIsvRequest
	GetCreateTimeStart() *string
	SetInvoiceId(v int64) *DescribeInvoiceForIsvRequest
	GetInvoiceId() *int64
	SetMaxResults(v int32) *DescribeInvoiceForIsvRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeInvoiceForIsvRequest
	GetNextToken() *string
	SetPageIndex(v int64) *DescribeInvoiceForIsvRequest
	GetPageIndex() *int64
	SetPageSize(v int64) *DescribeInvoiceForIsvRequest
	GetPageSize() *int64
	SetStatus(v int64) *DescribeInvoiceForIsvRequest
	GetStatus() *int64
	SetType(v int64) *DescribeInvoiceForIsvRequest
	GetType() *int64
	SetUserId(v int64) *DescribeInvoiceForIsvRequest
	GetUserId() *int64
}

type DescribeInvoiceForIsvRequest struct {
	// The end time. By default, 23:59:59 is appended to the date.
	//
	// example:
	//
	// 2025-01-01 00:00:00
	CreateTimeEnd *string `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	// The start time. By default, 00:00:00 is appended to the date.
	//
	// example:
	//
	// 2025-01-31 23:59:59
	CreateTimeStart *string `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	// The ID of the invoice application. This corresponds to the Result.Id field in the response of the DescribeInvoiceForIsv operation.
	//
	// example:
	//
	// 4072040****
	InvoiceId *int64 `json:"InvoiceId,omitempty" xml:"InvoiceId,omitempty"`
	// The maximum number of entries per page in a paging query. Maximum value: 50. Default value: 10. (This parameter is not enabled.)
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token returned by the current call. This parameter is not enabled.
	//
	// example:
	//
	// 3v3mzZN1QdVsTPNiT0OkD9v+MeV8LSqSQDSuTTVFCUMMWShmnj/LXRMScqm242bB
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number. Pages start from 1.
	//
	// example:
	//
	// 1
	PageIndex *int64 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The current invoice status. Valid values:
	//
	// - 0: Applying.
	//
	// - 1: Completed.
	//
	// example:
	//
	// 0
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The invoice type. If this parameter is left empty, all types are queried by default. Valid values:
	//
	// - 2: Digital electronic ordinary invoice.
	//
	// - 3: Digital electronic special invoice.
	//
	// example:
	//
	// 2
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
	// The Alibaba Cloud user ID of the customer.
	//
	// example:
	//
	// 174452687724****
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeInvoiceForIsvRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvoiceForIsvRequest) GoString() string {
	return s.String()
}

func (s *DescribeInvoiceForIsvRequest) GetCreateTimeEnd() *string {
	return s.CreateTimeEnd
}

func (s *DescribeInvoiceForIsvRequest) GetCreateTimeStart() *string {
	return s.CreateTimeStart
}

func (s *DescribeInvoiceForIsvRequest) GetInvoiceId() *int64 {
	return s.InvoiceId
}

func (s *DescribeInvoiceForIsvRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeInvoiceForIsvRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeInvoiceForIsvRequest) GetPageIndex() *int64 {
	return s.PageIndex
}

func (s *DescribeInvoiceForIsvRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeInvoiceForIsvRequest) GetStatus() *int64 {
	return s.Status
}

func (s *DescribeInvoiceForIsvRequest) GetType() *int64 {
	return s.Type
}

func (s *DescribeInvoiceForIsvRequest) GetUserId() *int64 {
	return s.UserId
}

func (s *DescribeInvoiceForIsvRequest) SetCreateTimeEnd(v string) *DescribeInvoiceForIsvRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *DescribeInvoiceForIsvRequest) SetCreateTimeStart(v string) *DescribeInvoiceForIsvRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *DescribeInvoiceForIsvRequest) SetInvoiceId(v int64) *DescribeInvoiceForIsvRequest {
	s.InvoiceId = &v
	return s
}

func (s *DescribeInvoiceForIsvRequest) SetMaxResults(v int32) *DescribeInvoiceForIsvRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeInvoiceForIsvRequest) SetNextToken(v string) *DescribeInvoiceForIsvRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeInvoiceForIsvRequest) SetPageIndex(v int64) *DescribeInvoiceForIsvRequest {
	s.PageIndex = &v
	return s
}

func (s *DescribeInvoiceForIsvRequest) SetPageSize(v int64) *DescribeInvoiceForIsvRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInvoiceForIsvRequest) SetStatus(v int64) *DescribeInvoiceForIsvRequest {
	s.Status = &v
	return s
}

func (s *DescribeInvoiceForIsvRequest) SetType(v int64) *DescribeInvoiceForIsvRequest {
	s.Type = &v
	return s
}

func (s *DescribeInvoiceForIsvRequest) SetUserId(v int64) *DescribeInvoiceForIsvRequest {
	s.UserId = &v
	return s
}

func (s *DescribeInvoiceForIsvRequest) Validate() error {
	return dara.Validate(s)
}
