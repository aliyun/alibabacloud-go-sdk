// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLiveId(v string) *DeleteLiveRequest
	GetLiveId() *string
	SetTenantContext(v *DeleteLiveRequestTenantContext) *DeleteLiveRequest
	GetTenantContext() *DeleteLiveRequestTenantContext
}

type DeleteLiveRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4d38xxxxx
	LiveId        *string                         `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	TenantContext *DeleteLiveRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s DeleteLiveRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveRequest) GetLiveId() *string {
	return s.LiveId
}

func (s *DeleteLiveRequest) GetTenantContext() *DeleteLiveRequestTenantContext {
	return s.TenantContext
}

func (s *DeleteLiveRequest) SetLiveId(v string) *DeleteLiveRequest {
	s.LiveId = &v
	return s
}

func (s *DeleteLiveRequest) SetTenantContext(v *DeleteLiveRequestTenantContext) *DeleteLiveRequest {
	s.TenantContext = v
	return s
}

func (s *DeleteLiveRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteLiveRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s DeleteLiveRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveRequestTenantContext) GoString() string {
	return s.String()
}

func (s *DeleteLiveRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteLiveRequestTenantContext) SetTenantId(v string) *DeleteLiveRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *DeleteLiveRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
