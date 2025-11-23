// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSQLReviewOriginSQLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderActionDetail(v *ListSQLReviewOriginSQLRequestOrderActionDetail) *ListSQLReviewOriginSQLRequest
	GetOrderActionDetail() *ListSQLReviewOriginSQLRequestOrderActionDetail
	SetOrderId(v int64) *ListSQLReviewOriginSQLRequest
	GetOrderId() *int64
	SetTid(v int64) *ListSQLReviewOriginSQLRequest
	GetTid() *int64
}

type ListSQLReviewOriginSQLRequest struct {
	// The parameters that are used to filter SQL statements involved in the ticket.
	OrderActionDetail *ListSQLReviewOriginSQLRequestOrderActionDetail `json:"OrderActionDetail,omitempty" xml:"OrderActionDetail,omitempty" type:"Struct"`
	// The ID of the SQL review ticket. You can call the [CreateSQLReviewOrder](https://help.aliyun.com/document_detail/257777.html) operation to query the ticket ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123321
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The tenant ID. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to query the tenant ID.
	//
	// example:
	//
	// 1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListSQLReviewOriginSQLRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSQLReviewOriginSQLRequest) GoString() string {
	return s.String()
}

func (s *ListSQLReviewOriginSQLRequest) GetOrderActionDetail() *ListSQLReviewOriginSQLRequestOrderActionDetail {
	return s.OrderActionDetail
}

func (s *ListSQLReviewOriginSQLRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ListSQLReviewOriginSQLRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListSQLReviewOriginSQLRequest) SetOrderActionDetail(v *ListSQLReviewOriginSQLRequestOrderActionDetail) *ListSQLReviewOriginSQLRequest {
	s.OrderActionDetail = v
	return s
}

func (s *ListSQLReviewOriginSQLRequest) SetOrderId(v int64) *ListSQLReviewOriginSQLRequest {
	s.OrderId = &v
	return s
}

func (s *ListSQLReviewOriginSQLRequest) SetTid(v int64) *ListSQLReviewOriginSQLRequest {
	s.Tid = &v
	return s
}

func (s *ListSQLReviewOriginSQLRequest) Validate() error {
	if s.OrderActionDetail != nil {
		if err := s.OrderActionDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSQLReviewOriginSQLRequestOrderActionDetail struct {
	// The review status of the SQL statement. Valid values:
	//
	// 	- **new**: The SQL statement was waiting to be reviewed.
	//
	// 	- **unknown**: The SQL statement cannot be parsed.
	//
	// 	- **check_not_pass**: The SQL statement failed to pass the review.
	//
	// 	- **check_pass**: The SQL statement passed the review.
	//
	// 	- **force_pass**: The SQL statement passed the manual review.
	//
	// 	- **force_not_pass**: The SQL statement failed to pass the manual review.
	//
	// example:
	//
	// check_not_pass
	CheckStatusResult *string `json:"CheckStatusResult,omitempty" xml:"CheckStatusResult,omitempty"`
	// The ID of the file that contains the SQL statements to be reviewed.
	//
	// example:
	//
	// 123345
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The pagination information.
	Page *ListSQLReviewOriginSQLRequestOrderActionDetailPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// The optimization suggestion for the SQL statement. Valid values:
	//
	// 	- **MUST_IMPROVE**: The SQL statement must be optimized.
	//
	// 	- **POTENTIAL_ISSUE**: The SQL statement contains potential issues.
	//
	// 	- **SUGGEST_IMPROVE**: We recommend that you optimize the SQL statement.
	//
	// 	- **USE_DMS_TOOLKIT**: We recommend that you change schemas without locking tables.
	//
	// 	- **USE_DMS_DML_UNLOCK**: We recommend that you change data without locking tables.
	//
	// 	- **TABLE_INDEX_SUGGEST**: We recommend that you optimize indexes for the SQL statement.
	//
	// example:
	//
	// MUST_IMPROVE
	SQLReviewResult *string `json:"SQLReviewResult,omitempty" xml:"SQLReviewResult,omitempty"`
}

func (s ListSQLReviewOriginSQLRequestOrderActionDetail) String() string {
	return dara.Prettify(s)
}

func (s ListSQLReviewOriginSQLRequestOrderActionDetail) GoString() string {
	return s.String()
}

func (s *ListSQLReviewOriginSQLRequestOrderActionDetail) GetCheckStatusResult() *string {
	return s.CheckStatusResult
}

func (s *ListSQLReviewOriginSQLRequestOrderActionDetail) GetFileId() *int64 {
	return s.FileId
}

func (s *ListSQLReviewOriginSQLRequestOrderActionDetail) GetPage() *ListSQLReviewOriginSQLRequestOrderActionDetailPage {
	return s.Page
}

func (s *ListSQLReviewOriginSQLRequestOrderActionDetail) GetSQLReviewResult() *string {
	return s.SQLReviewResult
}

func (s *ListSQLReviewOriginSQLRequestOrderActionDetail) SetCheckStatusResult(v string) *ListSQLReviewOriginSQLRequestOrderActionDetail {
	s.CheckStatusResult = &v
	return s
}

func (s *ListSQLReviewOriginSQLRequestOrderActionDetail) SetFileId(v int64) *ListSQLReviewOriginSQLRequestOrderActionDetail {
	s.FileId = &v
	return s
}

func (s *ListSQLReviewOriginSQLRequestOrderActionDetail) SetPage(v *ListSQLReviewOriginSQLRequestOrderActionDetailPage) *ListSQLReviewOriginSQLRequestOrderActionDetail {
	s.Page = v
	return s
}

func (s *ListSQLReviewOriginSQLRequestOrderActionDetail) SetSQLReviewResult(v string) *ListSQLReviewOriginSQLRequestOrderActionDetail {
	s.SQLReviewResult = &v
	return s
}

func (s *ListSQLReviewOriginSQLRequestOrderActionDetail) Validate() error {
	if s.Page != nil {
		if err := s.Page.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSQLReviewOriginSQLRequestOrderActionDetailPage struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSQLReviewOriginSQLRequestOrderActionDetailPage) String() string {
	return dara.Prettify(s)
}

func (s ListSQLReviewOriginSQLRequestOrderActionDetailPage) GoString() string {
	return s.String()
}

func (s *ListSQLReviewOriginSQLRequestOrderActionDetailPage) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSQLReviewOriginSQLRequestOrderActionDetailPage) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSQLReviewOriginSQLRequestOrderActionDetailPage) SetPageNumber(v int32) *ListSQLReviewOriginSQLRequestOrderActionDetailPage {
	s.PageNumber = &v
	return s
}

func (s *ListSQLReviewOriginSQLRequestOrderActionDetailPage) SetPageSize(v int32) *ListSQLReviewOriginSQLRequestOrderActionDetailPage {
	s.PageSize = &v
	return s
}

func (s *ListSQLReviewOriginSQLRequestOrderActionDetailPage) Validate() error {
	return dara.Validate(s)
}
