// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAttackTargetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFirstScanTimeEnd(v string) *ListAttackTargetsRequest
	GetFirstScanTimeEnd() *string
	SetFirstScanTimeStart(v string) *ListAttackTargetsRequest
	GetFirstScanTimeStart() *string
	SetLastScanStatus(v string) *ListAttackTargetsRequest
	GetLastScanStatus() *string
	SetLastScanTimeEnd(v string) *ListAttackTargetsRequest
	GetLastScanTimeEnd() *string
	SetLastScanTimeStart(v string) *ListAttackTargetsRequest
	GetLastScanTimeStart() *string
	SetPageNumber(v int64) *ListAttackTargetsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListAttackTargetsRequest
	GetPageSize() *int64
	SetProvider(v string) *ListAttackTargetsRequest
	GetProvider() *string
	SetRiskLevel(v string) *ListAttackTargetsRequest
	GetRiskLevel() *string
	SetSortField(v string) *ListAttackTargetsRequest
	GetSortField() *string
	SetSortOrder(v string) *ListAttackTargetsRequest
	GetSortOrder() *string
	SetTargetName(v string) *ListAttackTargetsRequest
	GetTargetName() *string
	SetTargetType(v string) *ListAttackTargetsRequest
	GetTargetType() *string
}

type ListAttackTargetsRequest struct {
	// The upper bound (inclusive) of the first scan time range. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1735689600000
	FirstScanTimeEnd *string `json:"FirstScanTimeEnd,omitempty" xml:"FirstScanTimeEnd,omitempty"`
	// The lower bound (inclusive) of the first scan time range. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1735689600000
	FirstScanTimeStart *string `json:"FirstScanTimeStart,omitempty" xml:"FirstScanTimeStart,omitempty"`
	// Filters targets by the status of the most recent scan task.
	//
	// example:
	//
	// completed
	LastScanStatus *string `json:"LastScanStatus,omitempty" xml:"LastScanStatus,omitempty"`
	// The upper bound (inclusive) of the last scan time range. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1735689600000
	LastScanTimeEnd *string `json:"LastScanTimeEnd,omitempty" xml:"LastScanTimeEnd,omitempty"`
	// The lower bound (inclusive) of the last scan time range. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1735689600000
	LastScanTimeStart *string `json:"LastScanTimeStart,omitempty" xml:"LastScanTimeStart,omitempty"`
	// The page number. Pages start from 1. Values less than 1 are normalized to 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100. Values greater than 100 are clamped to 100. Values less than 1 return HTTP status code 400.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Filters targets by the exact business label of the model or agent provider. This parameter is decoupled from ConnectionMethod (technical protocol).
	//
	// example:
	//
	// bailian
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// Filters targets by the risk level derived from the most recent completed scan task. Targets that have never been scanned do not have a risk level and are not matched by any value.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The field used for sorting. Only the following three aggregate fields are supported. Sorting is performed in memory. If this parameter is not specified, no additional sorting is applied.
	//
	// example:
	//
	// lastScanTime
	SortField *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	// The sort order. Targets with null aggregate values are always placed last regardless of the sort order.
	//
	// example:
	//
	// desc
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// Filters targets by name using fuzzy match (substring match). If this parameter is not specified, all targets are returned.
	//
	// example:
	//
	// Bailian
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// Filters targets by the exact scan target type.
	//
	// example:
	//
	// model
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListAttackTargetsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAttackTargetsRequest) GoString() string {
	return s.String()
}

func (s *ListAttackTargetsRequest) GetFirstScanTimeEnd() *string {
	return s.FirstScanTimeEnd
}

func (s *ListAttackTargetsRequest) GetFirstScanTimeStart() *string {
	return s.FirstScanTimeStart
}

func (s *ListAttackTargetsRequest) GetLastScanStatus() *string {
	return s.LastScanStatus
}

func (s *ListAttackTargetsRequest) GetLastScanTimeEnd() *string {
	return s.LastScanTimeEnd
}

func (s *ListAttackTargetsRequest) GetLastScanTimeStart() *string {
	return s.LastScanTimeStart
}

func (s *ListAttackTargetsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListAttackTargetsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAttackTargetsRequest) GetProvider() *string {
	return s.Provider
}

func (s *ListAttackTargetsRequest) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *ListAttackTargetsRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListAttackTargetsRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListAttackTargetsRequest) GetTargetName() *string {
	return s.TargetName
}

func (s *ListAttackTargetsRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *ListAttackTargetsRequest) SetFirstScanTimeEnd(v string) *ListAttackTargetsRequest {
	s.FirstScanTimeEnd = &v
	return s
}

func (s *ListAttackTargetsRequest) SetFirstScanTimeStart(v string) *ListAttackTargetsRequest {
	s.FirstScanTimeStart = &v
	return s
}

func (s *ListAttackTargetsRequest) SetLastScanStatus(v string) *ListAttackTargetsRequest {
	s.LastScanStatus = &v
	return s
}

func (s *ListAttackTargetsRequest) SetLastScanTimeEnd(v string) *ListAttackTargetsRequest {
	s.LastScanTimeEnd = &v
	return s
}

func (s *ListAttackTargetsRequest) SetLastScanTimeStart(v string) *ListAttackTargetsRequest {
	s.LastScanTimeStart = &v
	return s
}

func (s *ListAttackTargetsRequest) SetPageNumber(v int64) *ListAttackTargetsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAttackTargetsRequest) SetPageSize(v int64) *ListAttackTargetsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAttackTargetsRequest) SetProvider(v string) *ListAttackTargetsRequest {
	s.Provider = &v
	return s
}

func (s *ListAttackTargetsRequest) SetRiskLevel(v string) *ListAttackTargetsRequest {
	s.RiskLevel = &v
	return s
}

func (s *ListAttackTargetsRequest) SetSortField(v string) *ListAttackTargetsRequest {
	s.SortField = &v
	return s
}

func (s *ListAttackTargetsRequest) SetSortOrder(v string) *ListAttackTargetsRequest {
	s.SortOrder = &v
	return s
}

func (s *ListAttackTargetsRequest) SetTargetName(v string) *ListAttackTargetsRequest {
	s.TargetName = &v
	return s
}

func (s *ListAttackTargetsRequest) SetTargetType(v string) *ListAttackTargetsRequest {
	s.TargetType = &v
	return s
}

func (s *ListAttackTargetsRequest) Validate() error {
	return dara.Validate(s)
}
