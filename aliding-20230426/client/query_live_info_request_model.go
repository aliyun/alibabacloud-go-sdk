// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLiveInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLiveId(v string) *QueryLiveInfoRequest
	GetLiveId() *string
	SetTenantContext(v *QueryLiveInfoRequestTenantContext) *QueryLiveInfoRequest
	GetTenantContext() *QueryLiveInfoRequestTenantContext
}

type QueryLiveInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4d38xxxxx
	LiveId        *string                            `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	TenantContext *QueryLiveInfoRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s QueryLiveInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryLiveInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryLiveInfoRequest) GetLiveId() *string {
	return s.LiveId
}

func (s *QueryLiveInfoRequest) GetTenantContext() *QueryLiveInfoRequestTenantContext {
	return s.TenantContext
}

func (s *QueryLiveInfoRequest) SetLiveId(v string) *QueryLiveInfoRequest {
	s.LiveId = &v
	return s
}

func (s *QueryLiveInfoRequest) SetTenantContext(v *QueryLiveInfoRequestTenantContext) *QueryLiveInfoRequest {
	s.TenantContext = v
	return s
}

func (s *QueryLiveInfoRequest) Validate() error {
	return dara.Validate(s)
}

type QueryLiveInfoRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QueryLiveInfoRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s QueryLiveInfoRequestTenantContext) GoString() string {
	return s.String()
}

func (s *QueryLiveInfoRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryLiveInfoRequestTenantContext) SetTenantId(v string) *QueryLiveInfoRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *QueryLiveInfoRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
