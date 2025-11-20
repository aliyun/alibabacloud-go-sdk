// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeDingTalkIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDingTalkId(v string) *ChangeDingTalkIdRequest
	GetDingTalkId() *string
	SetTenantContext(v *ChangeDingTalkIdRequestTenantContext) *ChangeDingTalkIdRequest
	GetTenantContext() *ChangeDingTalkIdRequestTenantContext
}

type ChangeDingTalkIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4uf_iw54grufg9
	DingTalkId    *string                               `json:"DingTalkId,omitempty" xml:"DingTalkId,omitempty"`
	TenantContext *ChangeDingTalkIdRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s ChangeDingTalkIdRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeDingTalkIdRequest) GoString() string {
	return s.String()
}

func (s *ChangeDingTalkIdRequest) GetDingTalkId() *string {
	return s.DingTalkId
}

func (s *ChangeDingTalkIdRequest) GetTenantContext() *ChangeDingTalkIdRequestTenantContext {
	return s.TenantContext
}

func (s *ChangeDingTalkIdRequest) SetDingTalkId(v string) *ChangeDingTalkIdRequest {
	s.DingTalkId = &v
	return s
}

func (s *ChangeDingTalkIdRequest) SetTenantContext(v *ChangeDingTalkIdRequestTenantContext) *ChangeDingTalkIdRequest {
	s.TenantContext = v
	return s
}

func (s *ChangeDingTalkIdRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChangeDingTalkIdRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ChangeDingTalkIdRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s ChangeDingTalkIdRequestTenantContext) GoString() string {
	return s.String()
}

func (s *ChangeDingTalkIdRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *ChangeDingTalkIdRequestTenantContext) SetTenantId(v string) *ChangeDingTalkIdRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *ChangeDingTalkIdRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
