// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDIAlarmRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDIJobId(v int64) *ListDIAlarmRulesRequest
	GetDIJobId() *int64
	SetPageNumber(v int64) *ListDIAlarmRulesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListDIAlarmRulesRequest
	GetPageSize() *int64
}

type ListDIAlarmRulesRequest struct {
	// The ID of the task with which the alert rules are associated.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11260
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDIAlarmRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDIAlarmRulesRequest) GoString() string {
	return s.String()
}

func (s *ListDIAlarmRulesRequest) GetDIJobId() *int64 {
	return s.DIJobId
}

func (s *ListDIAlarmRulesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListDIAlarmRulesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListDIAlarmRulesRequest) SetDIJobId(v int64) *ListDIAlarmRulesRequest {
	s.DIJobId = &v
	return s
}

func (s *ListDIAlarmRulesRequest) SetPageNumber(v int64) *ListDIAlarmRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDIAlarmRulesRequest) SetPageSize(v int64) *ListDIAlarmRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDIAlarmRulesRequest) Validate() error {
	return dara.Validate(s)
}
