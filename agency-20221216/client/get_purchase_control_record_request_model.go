// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPurchaseControlRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerUID(v int64) *GetPurchaseControlRecordRequest
	GetCustomerUID() *int64
	SetOperationTime(v string) *GetPurchaseControlRecordRequest
	GetOperationTime() *string
	SetPageNo(v int32) *GetPurchaseControlRecordRequest
	GetPageNo() *int32
	SetPageSize(v int32) *GetPurchaseControlRecordRequest
	GetPageSize() *int32
}

type GetPurchaseControlRecordRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 292828565558721922
	CustomerUID *int64 `json:"CustomerUID,omitempty" xml:"CustomerUID,omitempty"`
	// example:
	//
	// 2023-12-15 10:34:36
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

func (s GetPurchaseControlRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPurchaseControlRecordRequest) GoString() string {
	return s.String()
}

func (s *GetPurchaseControlRecordRequest) GetCustomerUID() *int64 {
	return s.CustomerUID
}

func (s *GetPurchaseControlRecordRequest) GetOperationTime() *string {
	return s.OperationTime
}

func (s *GetPurchaseControlRecordRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetPurchaseControlRecordRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetPurchaseControlRecordRequest) SetCustomerUID(v int64) *GetPurchaseControlRecordRequest {
	s.CustomerUID = &v
	return s
}

func (s *GetPurchaseControlRecordRequest) SetOperationTime(v string) *GetPurchaseControlRecordRequest {
	s.OperationTime = &v
	return s
}

func (s *GetPurchaseControlRecordRequest) SetPageNo(v int32) *GetPurchaseControlRecordRequest {
	s.PageNo = &v
	return s
}

func (s *GetPurchaseControlRecordRequest) SetPageSize(v int32) *GetPurchaseControlRecordRequest {
	s.PageSize = &v
	return s
}

func (s *GetPurchaseControlRecordRequest) Validate() error {
	return dara.Validate(s)
}
