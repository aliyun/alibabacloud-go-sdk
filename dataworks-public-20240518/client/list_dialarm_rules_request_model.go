// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDIAlarmRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDIAlarmRuleId(v int64) *ListDIAlarmRulesRequest
	GetDIAlarmRuleId() *int64
	SetJobId(v int64) *ListDIAlarmRulesRequest
	GetJobId() *int64
	SetPageNumber(v int32) *ListDIAlarmRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDIAlarmRulesRequest
	GetPageSize() *int32
}

type ListDIAlarmRulesRequest struct {
	// The ID of the alert rule. If you leave this parameter empty, all alert rules of the task are queried.
	//
	// example:
	//
	// 34988
	DIAlarmRuleId *int64 `json:"DIAlarmRuleId,omitempty" xml:"DIAlarmRuleId,omitempty"`
	// The ID of the task for which alert rules are configured.
	//
	// example:
	//
	// 1000001
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDIAlarmRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDIAlarmRulesRequest) GoString() string {
	return s.String()
}

func (s *ListDIAlarmRulesRequest) GetDIAlarmRuleId() *int64 {
	return s.DIAlarmRuleId
}

func (s *ListDIAlarmRulesRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *ListDIAlarmRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDIAlarmRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDIAlarmRulesRequest) SetDIAlarmRuleId(v int64) *ListDIAlarmRulesRequest {
	s.DIAlarmRuleId = &v
	return s
}

func (s *ListDIAlarmRulesRequest) SetJobId(v int64) *ListDIAlarmRulesRequest {
	s.JobId = &v
	return s
}

func (s *ListDIAlarmRulesRequest) SetPageNumber(v int32) *ListDIAlarmRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDIAlarmRulesRequest) SetPageSize(v int32) *ListDIAlarmRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDIAlarmRulesRequest) Validate() error {
	return dara.Validate(s)
}
