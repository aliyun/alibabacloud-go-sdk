// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBatchResultDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBatchResultDetails(v *DescribeBatchResultDetailResponseBodyBatchResultDetails) *DescribeBatchResultDetailResponseBody
	GetBatchResultDetails() *DescribeBatchResultDetailResponseBodyBatchResultDetails
	SetPageNumber(v int64) *DescribeBatchResultDetailResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeBatchResultDetailResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeBatchResultDetailResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeBatchResultDetailResponseBody
	GetTotalCount() *int64
}

type DescribeBatchResultDetailResponseBody struct {
	// The detailed results of the batch operation.
	BatchResultDetails *DescribeBatchResultDetailResponseBodyBatchResultDetails `json:"BatchResultDetails,omitempty" xml:"BatchResultDetails,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 75446CC1-FC9A-4595-8D96-089D73D7A63D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBatchResultDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBatchResultDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBatchResultDetailResponseBody) GetBatchResultDetails() *DescribeBatchResultDetailResponseBodyBatchResultDetails {
	return s.BatchResultDetails
}

func (s *DescribeBatchResultDetailResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeBatchResultDetailResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeBatchResultDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBatchResultDetailResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeBatchResultDetailResponseBody) SetBatchResultDetails(v *DescribeBatchResultDetailResponseBodyBatchResultDetails) *DescribeBatchResultDetailResponseBody {
	s.BatchResultDetails = v
	return s
}

func (s *DescribeBatchResultDetailResponseBody) SetPageNumber(v int64) *DescribeBatchResultDetailResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBody) SetPageSize(v int64) *DescribeBatchResultDetailResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBody) SetRequestId(v string) *DescribeBatchResultDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBody) SetTotalCount(v int64) *DescribeBatchResultDetailResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeBatchResultDetailResponseBodyBatchResultDetails struct {
	BatchResultDetail []*DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail `json:"BatchResultDetail,omitempty" xml:"BatchResultDetail,omitempty" type:"Repeated"`
}

func (s DescribeBatchResultDetailResponseBodyBatchResultDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeBatchResultDetailResponseBodyBatchResultDetails) GoString() string {
	return s.String()
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetails) GetBatchResultDetail() []*DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail {
	return s.BatchResultDetail
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetails) SetBatchResultDetail(v []*DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail) *DescribeBatchResultDetailResponseBodyBatchResultDetails {
	s.BatchResultDetail = v
	return s
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetails) Validate() error {
	return dara.Validate(s)
}

type DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail struct {
	// The type of the batch operation.
	//
	// example:
	//
	// DOMAIN_ADD
	BatchType *string `json:"BatchType,omitempty" xml:"BatchType,omitempty"`
	// The domain name.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The line code.
	//
	// example:
	//
	// default
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
	// The new hostname.
	//
	// example:
	//
	// demo-batch-7
	NewRr *string `json:"NewRr,omitempty" xml:"NewRr,omitempty"`
	// The new record value.
	//
	// example:
	//
	// 192.0.2.254
	NewValue *string `json:"NewValue,omitempty" xml:"NewValue,omitempty"`
	// The time when the operation was performed. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-08-22 18:02:58
	OperateDateStr *string `json:"OperateDateStr,omitempty" xml:"OperateDateStr,omitempty"`
	// The priority of the mail exchanger (MX) record.
	//
	// example:
	//
	// 10
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The cause of the execution failure.
	//
	// example:
	//
	// Task lock fail
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The ID of the DNS record.
	//
	// example:
	//
	// 123456789
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The description of the DNS record.
	//
	// example:
	//
	// remark
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The hostname.
	//
	// example:
	//
	// www
	Rr *string `json:"Rr,omitempty" xml:"Rr,omitempty"`
	// The status of the DNS record.
	//
	// example:
	//
	// stop
	RrStatus *string `json:"RrStatus,omitempty" xml:"RrStatus,omitempty"`
	// The execution result of the batch operation. Valid values: **true**: The operation succeeded. **false**: The operation failed.
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time-to-live (TTL) of the DNS record.
	//
	// example:
	//
	// 600
	Ttl *string `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The type of the DNS record.
	//
	// example:
	//
	// A
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of the DNS record.
	//
	// example:
	//
	// 192.0.2.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail) GoString() string {
	return s.String()
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail) GetBatchType() *string {
	return s.BatchType
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail) GetDomain() *string {
	return s.Domain
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail) GetLine() *string {
	return s.Line
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail) GetNewRr() *string {
	return s.NewRr
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail) GetNewValue() *string {
	return s.NewValue
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail) GetOperateDateStr() *string {
	return s.OperateDateStr
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail) GetPriority() *string {
	return s.Priority
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail) GetReason() *string {
	return s.Reason
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail) GetRecordId() *string {
	return s.RecordId
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail) GetRemark() *string {
	return s.Remark
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail) GetRr() *string {
	return s.Rr
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail) GetRrStatus() *string {
	return s.RrStatus
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail) GetStatus() *bool {
	return s.Status
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail) GetTtl() *string {
	return s.Ttl
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail) GetType() *string {
	return s.Type
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail) GetValue() *string {
	return s.Value
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail) SetBatchType(v string) *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail {
	s.BatchType = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail) SetDomain(v string) *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail {
	s.Domain = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail) SetLine(v string) *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail {
	s.Line = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail) SetNewRr(v string) *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail {
	s.NewRr = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail) SetNewValue(v string) *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail {
	s.NewValue = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail) SetOperateDateStr(v string) *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail {
	s.OperateDateStr = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail) SetPriority(v string) *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail {
	s.Priority = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail) SetReason(v string) *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail {
	s.Reason = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail) SetRecordId(v string) *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail {
	s.RecordId = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail) SetRemark(v string) *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail {
	s.Remark = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail) SetRr(v string) *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail {
	s.Rr = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail) SetRrStatus(v string) *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail {
	s.RrStatus = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail) SetStatus(v bool) *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail {
	s.Status = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail) SetTtl(v string) *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail {
	s.Ttl = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail) SetType(v string) *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail {
	s.Type = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail) SetValue(v string) *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail {
	s.Value = &v
	return s
}

func (s *DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail) Validate() error {
	return dara.Validate(s)
}
