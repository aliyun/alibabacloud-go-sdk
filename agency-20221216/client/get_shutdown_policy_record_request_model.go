// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetShutdownPolicyRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerUID(v int64) *GetShutdownPolicyRecordRequest
	GetCustomerUID() *int64
	SetOperationTime(v string) *GetShutdownPolicyRecordRequest
	GetOperationTime() *string
	SetPageNo(v int32) *GetShutdownPolicyRecordRequest
	GetPageNo() *int32
	SetPageSize(v int32) *GetShutdownPolicyRecordRequest
	GetPageSize() *int32
}

type GetShutdownPolicyRecordRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 292828565558721922
	CustomerUID *int64 `json:"CustomerUID,omitempty" xml:"CustomerUID,omitempty"`
	// example:
	//
	// 2025-12-15 10:34:36
	OperationTime *string `json:"OperationTime,omitempty" xml:"OperationTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetShutdownPolicyRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s GetShutdownPolicyRecordRequest) GoString() string {
	return s.String()
}

func (s *GetShutdownPolicyRecordRequest) GetCustomerUID() *int64 {
	return s.CustomerUID
}

func (s *GetShutdownPolicyRecordRequest) GetOperationTime() *string {
	return s.OperationTime
}

func (s *GetShutdownPolicyRecordRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetShutdownPolicyRecordRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetShutdownPolicyRecordRequest) SetCustomerUID(v int64) *GetShutdownPolicyRecordRequest {
	s.CustomerUID = &v
	return s
}

func (s *GetShutdownPolicyRecordRequest) SetOperationTime(v string) *GetShutdownPolicyRecordRequest {
	s.OperationTime = &v
	return s
}

func (s *GetShutdownPolicyRecordRequest) SetPageNo(v int32) *GetShutdownPolicyRecordRequest {
	s.PageNo = &v
	return s
}

func (s *GetShutdownPolicyRecordRequest) SetPageSize(v int32) *GetShutdownPolicyRecordRequest {
	s.PageSize = &v
	return s
}

func (s *GetShutdownPolicyRecordRequest) Validate() error {
	return dara.Validate(s)
}
