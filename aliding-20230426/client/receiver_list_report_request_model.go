// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReceiverListReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOffset(v int64) *ReceiverListReportRequest
	GetOffset() *int64
	SetReportId(v string) *ReceiverListReportRequest
	GetReportId() *string
	SetSize(v int64) *ReceiverListReportRequest
	GetSize() *int64
	SetTenantContext(v *ReceiverListReportRequestTenantContext) *ReceiverListReportRequest
	GetTenantContext() *ReceiverListReportRequestTenantContext
}

type ReceiverListReportRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 10
	Offset *int64 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 173xxxx
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	Size          *int64                                  `json:"Size,omitempty" xml:"Size,omitempty"`
	TenantContext *ReceiverListReportRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s ReceiverListReportRequest) String() string {
	return dara.Prettify(s)
}

func (s ReceiverListReportRequest) GoString() string {
	return s.String()
}

func (s *ReceiverListReportRequest) GetOffset() *int64 {
	return s.Offset
}

func (s *ReceiverListReportRequest) GetReportId() *string {
	return s.ReportId
}

func (s *ReceiverListReportRequest) GetSize() *int64 {
	return s.Size
}

func (s *ReceiverListReportRequest) GetTenantContext() *ReceiverListReportRequestTenantContext {
	return s.TenantContext
}

func (s *ReceiverListReportRequest) SetOffset(v int64) *ReceiverListReportRequest {
	s.Offset = &v
	return s
}

func (s *ReceiverListReportRequest) SetReportId(v string) *ReceiverListReportRequest {
	s.ReportId = &v
	return s
}

func (s *ReceiverListReportRequest) SetSize(v int64) *ReceiverListReportRequest {
	s.Size = &v
	return s
}

func (s *ReceiverListReportRequest) SetTenantContext(v *ReceiverListReportRequestTenantContext) *ReceiverListReportRequest {
	s.TenantContext = v
	return s
}

func (s *ReceiverListReportRequest) Validate() error {
	return dara.Validate(s)
}

type ReceiverListReportRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ReceiverListReportRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s ReceiverListReportRequestTenantContext) GoString() string {
	return s.String()
}

func (s *ReceiverListReportRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *ReceiverListReportRequestTenantContext) SetTenantId(v string) *ReceiverListReportRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *ReceiverListReportRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
