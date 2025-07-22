// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableAllSqlConcurrencyControlRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleContext(v string) *DisableAllSqlConcurrencyControlRulesRequest
	GetConsoleContext() *string
	SetInstanceId(v string) *DisableAllSqlConcurrencyControlRulesRequest
	GetInstanceId() *string
}

type DisableAllSqlConcurrencyControlRulesRequest struct {
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
}

func (s DisableAllSqlConcurrencyControlRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableAllSqlConcurrencyControlRulesRequest) GoString() string {
	return s.String()
}

func (s *DisableAllSqlConcurrencyControlRulesRequest) GetConsoleContext() *string {
	return s.ConsoleContext
}

func (s *DisableAllSqlConcurrencyControlRulesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableAllSqlConcurrencyControlRulesRequest) SetConsoleContext(v string) *DisableAllSqlConcurrencyControlRulesRequest {
	s.ConsoleContext = &v
	return s
}

func (s *DisableAllSqlConcurrencyControlRulesRequest) SetInstanceId(v string) *DisableAllSqlConcurrencyControlRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableAllSqlConcurrencyControlRulesRequest) Validate() error {
	return dara.Validate(s)
}
