// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStructSyncJobAnalyzeResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompareType(v string) *GetStructSyncJobAnalyzeResultRequest
	GetCompareType() *string
	SetOrderId(v int64) *GetStructSyncJobAnalyzeResultRequest
	GetOrderId() *int64
	SetPageNumber(v int64) *GetStructSyncJobAnalyzeResultRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *GetStructSyncJobAnalyzeResultRequest
	GetPageSize() *int64
	SetTid(v int64) *GetStructSyncJobAnalyzeResultRequest
	GetTid() *int64
}

type GetStructSyncJobAnalyzeResultRequest struct {
	// The type of the comparison. Valid values:
	//
	// 	- **CREATE_TABLE**: compares the created tables.
	//
	// 	- **ALTER_TABLE**: compares the modified tables.
	//
	// 	- **EQUAL_TABLE**: compares the identical tables.
	//
	// 	- **PASS_TABLE**: compares the tables that are skipped during schema synchronization.
	//
	// 	- **NOT_COMPARE**: does not compare tables.
	//
	// example:
	//
	// CREATE_TABLE
	CompareType *string `json:"CompareType,omitempty" xml:"CompareType,omitempty"`
	// The ID of the ticket.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1342355
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the tenant.
	//
	// > To view the tenant ID, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetStructSyncJobAnalyzeResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStructSyncJobAnalyzeResultRequest) GoString() string {
	return s.String()
}

func (s *GetStructSyncJobAnalyzeResultRequest) GetCompareType() *string {
	return s.CompareType
}

func (s *GetStructSyncJobAnalyzeResultRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetStructSyncJobAnalyzeResultRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetStructSyncJobAnalyzeResultRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetStructSyncJobAnalyzeResultRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetStructSyncJobAnalyzeResultRequest) SetCompareType(v string) *GetStructSyncJobAnalyzeResultRequest {
	s.CompareType = &v
	return s
}

func (s *GetStructSyncJobAnalyzeResultRequest) SetOrderId(v int64) *GetStructSyncJobAnalyzeResultRequest {
	s.OrderId = &v
	return s
}

func (s *GetStructSyncJobAnalyzeResultRequest) SetPageNumber(v int64) *GetStructSyncJobAnalyzeResultRequest {
	s.PageNumber = &v
	return s
}

func (s *GetStructSyncJobAnalyzeResultRequest) SetPageSize(v int64) *GetStructSyncJobAnalyzeResultRequest {
	s.PageSize = &v
	return s
}

func (s *GetStructSyncJobAnalyzeResultRequest) SetTid(v int64) *GetStructSyncJobAnalyzeResultRequest {
	s.Tid = &v
	return s
}

func (s *GetStructSyncJobAnalyzeResultRequest) Validate() error {
	return dara.Validate(s)
}
