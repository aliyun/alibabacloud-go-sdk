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
	if s.BatchResultDetails != nil {
		if err := s.BatchResultDetails.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.BatchResultDetail != nil {
		for _, item := range s.BatchResultDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBatchResultDetailResponseBodyBatchResultDetailsBatchResultDetail struct {
	BatchType      *string `json:"BatchType,omitempty" xml:"BatchType,omitempty"`
	Domain         *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Line           *string `json:"Line,omitempty" xml:"Line,omitempty"`
	NewRr          *string `json:"NewRr,omitempty" xml:"NewRr,omitempty"`
	NewValue       *string `json:"NewValue,omitempty" xml:"NewValue,omitempty"`
	OperateDateStr *string `json:"OperateDateStr,omitempty" xml:"OperateDateStr,omitempty"`
	Priority       *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Reason         *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	RecordId       *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	Remark         *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Rr             *string `json:"Rr,omitempty" xml:"Rr,omitempty"`
	RrStatus       *string `json:"RrStatus,omitempty" xml:"RrStatus,omitempty"`
	Status         *bool   `json:"Status,omitempty" xml:"Status,omitempty"`
	Ttl            *string `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value          *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
