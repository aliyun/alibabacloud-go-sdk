// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMonitorAlertHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertRuleName(v string) *GetMonitorAlertHistoryRequest
	GetAlertRuleName() *string
	SetEndTime(v string) *GetMonitorAlertHistoryRequest
	GetEndTime() *string
	SetEnv(v string) *GetMonitorAlertHistoryRequest
	GetEnv() *string
	SetOrderBy(v string) *GetMonitorAlertHistoryRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *GetMonitorAlertHistoryRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *GetMonitorAlertHistoryRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetMonitorAlertHistoryRequest
	GetPageSize() *int32
	SetPbcId(v int64) *GetMonitorAlertHistoryRequest
	GetPbcId() *int64
	SetServiceGroupId(v int64) *GetMonitorAlertHistoryRequest
	GetServiceGroupId() *int64
	SetStartTime(v string) *GetMonitorAlertHistoryRequest
	GetStartTime() *string
	SetType(v string) *GetMonitorAlertHistoryRequest
	GetType() *string
}

type GetMonitorAlertHistoryRequest struct {
	AlertRuleName *string `json:"alertRuleName,omitempty" xml:"alertRuleName,omitempty"`
	EndTime       *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// This parameter is required.
	Env            *string `json:"env,omitempty" xml:"env,omitempty"`
	OrderBy        *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	PageNumber     *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize       *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 123
	PbcId *int64 `json:"pbcId,omitempty" xml:"pbcId,omitempty"`
	// example:
	//
	// 123
	ServiceGroupId *int64  `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
	StartTime      *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetMonitorAlertHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMonitorAlertHistoryRequest) GoString() string {
	return s.String()
}

func (s *GetMonitorAlertHistoryRequest) GetAlertRuleName() *string {
	return s.AlertRuleName
}

func (s *GetMonitorAlertHistoryRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetMonitorAlertHistoryRequest) GetEnv() *string {
	return s.Env
}

func (s *GetMonitorAlertHistoryRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *GetMonitorAlertHistoryRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *GetMonitorAlertHistoryRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetMonitorAlertHistoryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetMonitorAlertHistoryRequest) GetPbcId() *int64 {
	return s.PbcId
}

func (s *GetMonitorAlertHistoryRequest) GetServiceGroupId() *int64 {
	return s.ServiceGroupId
}

func (s *GetMonitorAlertHistoryRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetMonitorAlertHistoryRequest) GetType() *string {
	return s.Type
}

func (s *GetMonitorAlertHistoryRequest) SetAlertRuleName(v string) *GetMonitorAlertHistoryRequest {
	s.AlertRuleName = &v
	return s
}

func (s *GetMonitorAlertHistoryRequest) SetEndTime(v string) *GetMonitorAlertHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *GetMonitorAlertHistoryRequest) SetEnv(v string) *GetMonitorAlertHistoryRequest {
	s.Env = &v
	return s
}

func (s *GetMonitorAlertHistoryRequest) SetOrderBy(v string) *GetMonitorAlertHistoryRequest {
	s.OrderBy = &v
	return s
}

func (s *GetMonitorAlertHistoryRequest) SetOrderDirection(v string) *GetMonitorAlertHistoryRequest {
	s.OrderDirection = &v
	return s
}

func (s *GetMonitorAlertHistoryRequest) SetPageNumber(v int32) *GetMonitorAlertHistoryRequest {
	s.PageNumber = &v
	return s
}

func (s *GetMonitorAlertHistoryRequest) SetPageSize(v int32) *GetMonitorAlertHistoryRequest {
	s.PageSize = &v
	return s
}

func (s *GetMonitorAlertHistoryRequest) SetPbcId(v int64) *GetMonitorAlertHistoryRequest {
	s.PbcId = &v
	return s
}

func (s *GetMonitorAlertHistoryRequest) SetServiceGroupId(v int64) *GetMonitorAlertHistoryRequest {
	s.ServiceGroupId = &v
	return s
}

func (s *GetMonitorAlertHistoryRequest) SetStartTime(v string) *GetMonitorAlertHistoryRequest {
	s.StartTime = &v
	return s
}

func (s *GetMonitorAlertHistoryRequest) SetType(v string) *GetMonitorAlertHistoryRequest {
	s.Type = &v
	return s
}

func (s *GetMonitorAlertHistoryRequest) Validate() error {
	return dara.Validate(s)
}
