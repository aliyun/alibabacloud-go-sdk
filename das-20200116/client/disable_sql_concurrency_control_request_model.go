// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSqlConcurrencyControlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleContext(v string) *DisableSqlConcurrencyControlRequest
	GetConsoleContext() *string
	SetInstanceId(v string) *DisableSqlConcurrencyControlRequest
	GetInstanceId() *string
	SetItemId(v int64) *DisableSqlConcurrencyControlRequest
	GetItemId() *int64
}

type DisableSqlConcurrencyControlRequest struct {
	// The reserved parameter.
	//
	// example:
	//
	// None
	ConsoleContext *string `json:"ConsoleContext,omitempty" xml:"ConsoleContext,omitempty"`
	// The instance ID.
	//
	// >  The database instance must be an ApsaraDB RDS for MySQL instance or a PolarDB for MySQL cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze1jdv45i7l6****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the throttling rule that is applied to the instance. You can call the [GetRunningSqlConcurrencyControlRules](https://help.aliyun.com/document_detail/223538.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ItemId *int64 `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
}

func (s DisableSqlConcurrencyControlRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableSqlConcurrencyControlRequest) GoString() string {
	return s.String()
}

func (s *DisableSqlConcurrencyControlRequest) GetConsoleContext() *string {
	return s.ConsoleContext
}

func (s *DisableSqlConcurrencyControlRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableSqlConcurrencyControlRequest) GetItemId() *int64 {
	return s.ItemId
}

func (s *DisableSqlConcurrencyControlRequest) SetConsoleContext(v string) *DisableSqlConcurrencyControlRequest {
	s.ConsoleContext = &v
	return s
}

func (s *DisableSqlConcurrencyControlRequest) SetInstanceId(v string) *DisableSqlConcurrencyControlRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableSqlConcurrencyControlRequest) SetItemId(v int64) *DisableSqlConcurrencyControlRequest {
	s.ItemId = &v
	return s
}

func (s *DisableSqlConcurrencyControlRequest) Validate() error {
	return dara.Validate(s)
}
