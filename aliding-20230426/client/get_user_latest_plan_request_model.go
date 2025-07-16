// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserLatestPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *GetUserLatestPlanRequestTenantContext) *GetUserLatestPlanRequest
	GetTenantContext() *GetUserLatestPlanRequestTenantContext
}

type GetUserLatestPlanRequest struct {
	TenantContext *GetUserLatestPlanRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s GetUserLatestPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserLatestPlanRequest) GoString() string {
	return s.String()
}

func (s *GetUserLatestPlanRequest) GetTenantContext() *GetUserLatestPlanRequestTenantContext {
	return s.TenantContext
}

func (s *GetUserLatestPlanRequest) SetTenantContext(v *GetUserLatestPlanRequestTenantContext) *GetUserLatestPlanRequest {
	s.TenantContext = v
	return s
}

func (s *GetUserLatestPlanRequest) Validate() error {
	return dara.Validate(s)
}

type GetUserLatestPlanRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetUserLatestPlanRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetUserLatestPlanRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetUserLatestPlanRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetUserLatestPlanRequestTenantContext) SetTenantId(v string) *GetUserLatestPlanRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetUserLatestPlanRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
