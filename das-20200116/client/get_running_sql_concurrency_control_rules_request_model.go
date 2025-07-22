// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRunningSqlConcurrencyControlRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleContext(v string) *GetRunningSqlConcurrencyControlRulesRequest
	GetConsoleContext() *string
	SetInstanceId(v string) *GetRunningSqlConcurrencyControlRulesRequest
	GetInstanceId() *string
	SetPageNo(v int64) *GetRunningSqlConcurrencyControlRulesRequest
	GetPageNo() *int64
	SetPageSize(v int64) *GetRunningSqlConcurrencyControlRulesRequest
	GetPageSize() *int64
}

type GetRunningSqlConcurrencyControlRulesRequest struct {
	// The reserved parameter.
	//
	// example:
	//
	// None
	ConsoleContext *string `json:"ConsoleContext,omitempty" xml:"ConsoleContext,omitempty"`
	// The instance ID.
	//
	// >  You must specify this parameter only if your database instance is an ApsaraDB RDS for MySQL instance or a PolarDB for MySQL cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze1jdv45i7l6****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. The value must be a positive integer. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. The value must be a positive integer. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetRunningSqlConcurrencyControlRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRunningSqlConcurrencyControlRulesRequest) GoString() string {
	return s.String()
}

func (s *GetRunningSqlConcurrencyControlRulesRequest) GetConsoleContext() *string {
	return s.ConsoleContext
}

func (s *GetRunningSqlConcurrencyControlRulesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRunningSqlConcurrencyControlRulesRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *GetRunningSqlConcurrencyControlRulesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetRunningSqlConcurrencyControlRulesRequest) SetConsoleContext(v string) *GetRunningSqlConcurrencyControlRulesRequest {
	s.ConsoleContext = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesRequest) SetInstanceId(v string) *GetRunningSqlConcurrencyControlRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesRequest) SetPageNo(v int64) *GetRunningSqlConcurrencyControlRulesRequest {
	s.PageNo = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesRequest) SetPageSize(v int64) *GetRunningSqlConcurrencyControlRulesRequest {
	s.PageSize = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesRequest) Validate() error {
	return dara.Validate(s)
}
