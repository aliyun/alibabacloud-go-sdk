// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryForOperationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplyTimeEnd(v int64) *QueryForOperationRequest
	GetApplyTimeEnd() *int64
	SetApplyTimeStart(v int64) *QueryForOperationRequest
	GetApplyTimeStart() *int64
	SetAssetType(v string) *QueryForOperationRequest
	GetAssetType() *string
	SetAuditTimeEnd(v int64) *QueryForOperationRequest
	GetAuditTimeEnd() *int64
	SetAuditTimeStart(v int64) *QueryForOperationRequest
	GetAuditTimeStart() *int64
	SetCurrentPage(v int32) *QueryForOperationRequest
	GetCurrentPage() *int32
	SetCustomerUid(v string) *QueryForOperationRequest
	GetCustomerUid() *string
	SetIp(v string) *QueryForOperationRequest
	GetIp() *string
	SetPageSize(v int32) *QueryForOperationRequest
	GetPageSize() *int32
	SetStatus(v string) *QueryForOperationRequest
	GetStatus() *string
}

type QueryForOperationRequest struct {
	// example:
	//
	// 1791937232000
	ApplyTimeEnd *int64 `json:"ApplyTimeEnd,omitempty" xml:"ApplyTimeEnd,omitempty"`
	// example:
	//
	// 1791937132000
	ApplyTimeStart *int64 `json:"ApplyTimeStart,omitempty" xml:"ApplyTimeStart,omitempty"`
	// example:
	//
	// EIP
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// example:
	//
	// 1791937232000
	AuditTimeEnd *int64 `json:"AuditTimeEnd,omitempty" xml:"AuditTimeEnd,omitempty"`
	// example:
	//
	// 1791937132000
	AuditTimeStart *int64 `json:"AuditTimeStart,omitempty" xml:"AuditTimeStart,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 1647796581073291
	CustomerUid *string `json:"CustomerUid,omitempty" xml:"CustomerUid,omitempty"`
	// example:
	//
	// 12.3*.22.11
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// approved
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryForOperationRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryForOperationRequest) GoString() string {
	return s.String()
}

func (s *QueryForOperationRequest) GetApplyTimeEnd() *int64 {
	return s.ApplyTimeEnd
}

func (s *QueryForOperationRequest) GetApplyTimeStart() *int64 {
	return s.ApplyTimeStart
}

func (s *QueryForOperationRequest) GetAssetType() *string {
	return s.AssetType
}

func (s *QueryForOperationRequest) GetAuditTimeEnd() *int64 {
	return s.AuditTimeEnd
}

func (s *QueryForOperationRequest) GetAuditTimeStart() *int64 {
	return s.AuditTimeStart
}

func (s *QueryForOperationRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryForOperationRequest) GetCustomerUid() *string {
	return s.CustomerUid
}

func (s *QueryForOperationRequest) GetIp() *string {
	return s.Ip
}

func (s *QueryForOperationRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryForOperationRequest) GetStatus() *string {
	return s.Status
}

func (s *QueryForOperationRequest) SetApplyTimeEnd(v int64) *QueryForOperationRequest {
	s.ApplyTimeEnd = &v
	return s
}

func (s *QueryForOperationRequest) SetApplyTimeStart(v int64) *QueryForOperationRequest {
	s.ApplyTimeStart = &v
	return s
}

func (s *QueryForOperationRequest) SetAssetType(v string) *QueryForOperationRequest {
	s.AssetType = &v
	return s
}

func (s *QueryForOperationRequest) SetAuditTimeEnd(v int64) *QueryForOperationRequest {
	s.AuditTimeEnd = &v
	return s
}

func (s *QueryForOperationRequest) SetAuditTimeStart(v int64) *QueryForOperationRequest {
	s.AuditTimeStart = &v
	return s
}

func (s *QueryForOperationRequest) SetCurrentPage(v int32) *QueryForOperationRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryForOperationRequest) SetCustomerUid(v string) *QueryForOperationRequest {
	s.CustomerUid = &v
	return s
}

func (s *QueryForOperationRequest) SetIp(v string) *QueryForOperationRequest {
	s.Ip = &v
	return s
}

func (s *QueryForOperationRequest) SetPageSize(v int32) *QueryForOperationRequest {
	s.PageSize = &v
	return s
}

func (s *QueryForOperationRequest) SetStatus(v string) *QueryForOperationRequest {
	s.Status = &v
	return s
}

func (s *QueryForOperationRequest) Validate() error {
	return dara.Validate(s)
}
