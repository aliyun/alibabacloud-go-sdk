// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGroupLiveInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnchorUnionId(v string) *QueryGroupLiveInfoRequest
	GetAnchorUnionId() *string
	SetLiveUuid(v string) *QueryGroupLiveInfoRequest
	GetLiveUuid() *string
	SetTenantContext(v *QueryGroupLiveInfoRequestTenantContext) *QueryGroupLiveInfoRequest
	GetTenantContext() *QueryGroupLiveInfoRequestTenantContext
}

type QueryGroupLiveInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Eijxxx
	AnchorUnionId *string `json:"AnchorUnionId,omitempty" xml:"AnchorUnionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	LiveUuid      *string                                 `json:"LiveUuid,omitempty" xml:"LiveUuid,omitempty"`
	TenantContext *QueryGroupLiveInfoRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s QueryGroupLiveInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryGroupLiveInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryGroupLiveInfoRequest) GetAnchorUnionId() *string {
	return s.AnchorUnionId
}

func (s *QueryGroupLiveInfoRequest) GetLiveUuid() *string {
	return s.LiveUuid
}

func (s *QueryGroupLiveInfoRequest) GetTenantContext() *QueryGroupLiveInfoRequestTenantContext {
	return s.TenantContext
}

func (s *QueryGroupLiveInfoRequest) SetAnchorUnionId(v string) *QueryGroupLiveInfoRequest {
	s.AnchorUnionId = &v
	return s
}

func (s *QueryGroupLiveInfoRequest) SetLiveUuid(v string) *QueryGroupLiveInfoRequest {
	s.LiveUuid = &v
	return s
}

func (s *QueryGroupLiveInfoRequest) SetTenantContext(v *QueryGroupLiveInfoRequestTenantContext) *QueryGroupLiveInfoRequest {
	s.TenantContext = v
	return s
}

func (s *QueryGroupLiveInfoRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryGroupLiveInfoRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QueryGroupLiveInfoRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s QueryGroupLiveInfoRequestTenantContext) GoString() string {
	return s.String()
}

func (s *QueryGroupLiveInfoRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryGroupLiveInfoRequestTenantContext) SetTenantId(v string) *QueryGroupLiveInfoRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *QueryGroupLiveInfoRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
