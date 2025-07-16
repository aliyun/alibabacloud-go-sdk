// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLiveWatchDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLiveId(v string) *QueryLiveWatchDetailRequest
	GetLiveId() *string
	SetTenantContext(v *QueryLiveWatchDetailRequestTenantContext) *QueryLiveWatchDetailRequest
	GetTenantContext() *QueryLiveWatchDetailRequestTenantContext
}

type QueryLiveWatchDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4d38xxxxx
	LiveId        *string                                   `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	TenantContext *QueryLiveWatchDetailRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s QueryLiveWatchDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryLiveWatchDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryLiveWatchDetailRequest) GetLiveId() *string {
	return s.LiveId
}

func (s *QueryLiveWatchDetailRequest) GetTenantContext() *QueryLiveWatchDetailRequestTenantContext {
	return s.TenantContext
}

func (s *QueryLiveWatchDetailRequest) SetLiveId(v string) *QueryLiveWatchDetailRequest {
	s.LiveId = &v
	return s
}

func (s *QueryLiveWatchDetailRequest) SetTenantContext(v *QueryLiveWatchDetailRequestTenantContext) *QueryLiveWatchDetailRequest {
	s.TenantContext = v
	return s
}

func (s *QueryLiveWatchDetailRequest) Validate() error {
	return dara.Validate(s)
}

type QueryLiveWatchDetailRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QueryLiveWatchDetailRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s QueryLiveWatchDetailRequestTenantContext) GoString() string {
	return s.String()
}

func (s *QueryLiveWatchDetailRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryLiveWatchDetailRequestTenantContext) SetTenantId(v string) *QueryLiveWatchDetailRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *QueryLiveWatchDetailRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
