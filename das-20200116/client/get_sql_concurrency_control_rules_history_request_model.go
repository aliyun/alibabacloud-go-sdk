// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSqlConcurrencyControlRulesHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleContext(v string) *GetSqlConcurrencyControlRulesHistoryRequest
	GetConsoleContext() *string
	SetInstanceId(v string) *GetSqlConcurrencyControlRulesHistoryRequest
	GetInstanceId() *string
	SetPageNo(v int64) *GetSqlConcurrencyControlRulesHistoryRequest
	GetPageNo() *int64
	SetPageSize(v int64) *GetSqlConcurrencyControlRulesHistoryRequest
	GetPageSize() *int64
}

type GetSqlConcurrencyControlRulesHistoryRequest struct {
	// The reserved parameter.
	//
	// example:
	//
	// None
	ConsoleContext *string `json:"ConsoleContext,omitempty" xml:"ConsoleContext,omitempty"`
	// The instance ID.
	//
	// >  Only ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze1jdv45i7l6****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. The value must be an integer that is greater than 0. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. The value must be an integer that is greater than 0. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetSqlConcurrencyControlRulesHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSqlConcurrencyControlRulesHistoryRequest) GoString() string {
	return s.String()
}

func (s *GetSqlConcurrencyControlRulesHistoryRequest) GetConsoleContext() *string {
	return s.ConsoleContext
}

func (s *GetSqlConcurrencyControlRulesHistoryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSqlConcurrencyControlRulesHistoryRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *GetSqlConcurrencyControlRulesHistoryRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetSqlConcurrencyControlRulesHistoryRequest) SetConsoleContext(v string) *GetSqlConcurrencyControlRulesHistoryRequest {
	s.ConsoleContext = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryRequest) SetInstanceId(v string) *GetSqlConcurrencyControlRulesHistoryRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryRequest) SetPageNo(v int64) *GetSqlConcurrencyControlRulesHistoryRequest {
	s.PageNo = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryRequest) SetPageSize(v int64) *GetSqlConcurrencyControlRulesHistoryRequest {
	s.PageSize = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryRequest) Validate() error {
	return dara.Validate(s)
}
