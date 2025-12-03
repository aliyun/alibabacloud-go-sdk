// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRestorePlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateRestorePlanRequest
	GetClusterId() *string
	SetRestoreAllTable(v bool) *CreateRestorePlanRequest
	GetRestoreAllTable() *bool
	SetRestoreByCopy(v bool) *CreateRestorePlanRequest
	GetRestoreByCopy() *bool
	SetRestoreToDate(v string) *CreateRestorePlanRequest
	GetRestoreToDate() *string
	SetTables(v string) *CreateRestorePlanRequest
	GetTables() *string
	SetTargetClusterId(v string) *CreateRestorePlanRequest
	GetTargetClusterId() *string
}

type CreateRestorePlanRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-bp150tns0sjxs****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	RestoreAllTable *bool `json:"RestoreAllTable,omitempty" xml:"RestoreAllTable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	RestoreByCopy *bool `json:"RestoreByCopy,omitempty" xml:"RestoreByCopy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2020-11-05T05:49:42Z
	RestoreToDate *string `json:"RestoreToDate,omitempty" xml:"RestoreToDate,omitempty"`
	// example:
	//
	// test_ns:test_table/test_ns:test_table2
	Tables *string `json:"Tables,omitempty" xml:"Tables,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ld-bp169l540vc6c****
	TargetClusterId *string `json:"TargetClusterId,omitempty" xml:"TargetClusterId,omitempty"`
}

func (s CreateRestorePlanRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRestorePlanRequest) GoString() string {
	return s.String()
}

func (s *CreateRestorePlanRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateRestorePlanRequest) GetRestoreAllTable() *bool {
	return s.RestoreAllTable
}

func (s *CreateRestorePlanRequest) GetRestoreByCopy() *bool {
	return s.RestoreByCopy
}

func (s *CreateRestorePlanRequest) GetRestoreToDate() *string {
	return s.RestoreToDate
}

func (s *CreateRestorePlanRequest) GetTables() *string {
	return s.Tables
}

func (s *CreateRestorePlanRequest) GetTargetClusterId() *string {
	return s.TargetClusterId
}

func (s *CreateRestorePlanRequest) SetClusterId(v string) *CreateRestorePlanRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateRestorePlanRequest) SetRestoreAllTable(v bool) *CreateRestorePlanRequest {
	s.RestoreAllTable = &v
	return s
}

func (s *CreateRestorePlanRequest) SetRestoreByCopy(v bool) *CreateRestorePlanRequest {
	s.RestoreByCopy = &v
	return s
}

func (s *CreateRestorePlanRequest) SetRestoreToDate(v string) *CreateRestorePlanRequest {
	s.RestoreToDate = &v
	return s
}

func (s *CreateRestorePlanRequest) SetTables(v string) *CreateRestorePlanRequest {
	s.Tables = &v
	return s
}

func (s *CreateRestorePlanRequest) SetTargetClusterId(v string) *CreateRestorePlanRequest {
	s.TargetClusterId = &v
	return s
}

func (s *CreateRestorePlanRequest) Validate() error {
	return dara.Validate(s)
}
