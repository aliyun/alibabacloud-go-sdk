// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryReportDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReportId(v string) *QueryReportDetailRequest
	GetReportId() *string
	SetTenantContext(v *QueryReportDetailRequestTenantContext) *QueryReportDetailRequest
	GetTenantContext() *QueryReportDetailRequestTenantContext
}

type QueryReportDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 18054XXX
	ReportId      *string                                `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	TenantContext *QueryReportDetailRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s QueryReportDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryReportDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryReportDetailRequest) GetReportId() *string {
	return s.ReportId
}

func (s *QueryReportDetailRequest) GetTenantContext() *QueryReportDetailRequestTenantContext {
	return s.TenantContext
}

func (s *QueryReportDetailRequest) SetReportId(v string) *QueryReportDetailRequest {
	s.ReportId = &v
	return s
}

func (s *QueryReportDetailRequest) SetTenantContext(v *QueryReportDetailRequestTenantContext) *QueryReportDetailRequest {
	s.TenantContext = v
	return s
}

func (s *QueryReportDetailRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryReportDetailRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QueryReportDetailRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s QueryReportDetailRequestTenantContext) GoString() string {
	return s.String()
}

func (s *QueryReportDetailRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryReportDetailRequestTenantContext) SetTenantId(v string) *QueryReportDetailRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *QueryReportDetailRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
