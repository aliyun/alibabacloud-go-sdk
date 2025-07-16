// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReportUnReadCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRequest(v map[string]interface{}) *GetReportUnReadCountRequest
	GetRequest() map[string]interface{}
	SetTenantContext(v *GetReportUnReadCountRequestTenantContext) *GetReportUnReadCountRequest
	GetTenantContext() *GetReportUnReadCountRequestTenantContext
}

type GetReportUnReadCountRequest struct {
	// example:
	//
	// null
	Request       map[string]interface{}                    `json:"Request,omitempty" xml:"Request,omitempty"`
	TenantContext *GetReportUnReadCountRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s GetReportUnReadCountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetReportUnReadCountRequest) GoString() string {
	return s.String()
}

func (s *GetReportUnReadCountRequest) GetRequest() map[string]interface{} {
	return s.Request
}

func (s *GetReportUnReadCountRequest) GetTenantContext() *GetReportUnReadCountRequestTenantContext {
	return s.TenantContext
}

func (s *GetReportUnReadCountRequest) SetRequest(v map[string]interface{}) *GetReportUnReadCountRequest {
	s.Request = v
	return s
}

func (s *GetReportUnReadCountRequest) SetTenantContext(v *GetReportUnReadCountRequestTenantContext) *GetReportUnReadCountRequest {
	s.TenantContext = v
	return s
}

func (s *GetReportUnReadCountRequest) Validate() error {
	return dara.Validate(s)
}

type GetReportUnReadCountRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetReportUnReadCountRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetReportUnReadCountRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetReportUnReadCountRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetReportUnReadCountRequestTenantContext) SetTenantId(v string) *GetReportUnReadCountRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetReportUnReadCountRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
