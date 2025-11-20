// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConversaionSpaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpenConversationId(v string) *GetConversaionSpaceRequest
	GetOpenConversationId() *string
	SetTenantContext(v *GetConversaionSpaceRequestTenantContext) *GetConversaionSpaceRequest
	GetTenantContext() *GetConversaionSpaceRequestTenantContext
}

type GetConversaionSpaceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cidB8Pzg****FIWPv2PMA==
	OpenConversationId *string                                  `json:"OpenConversationId,omitempty" xml:"OpenConversationId,omitempty"`
	TenantContext      *GetConversaionSpaceRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s GetConversaionSpaceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetConversaionSpaceRequest) GoString() string {
	return s.String()
}

func (s *GetConversaionSpaceRequest) GetOpenConversationId() *string {
	return s.OpenConversationId
}

func (s *GetConversaionSpaceRequest) GetTenantContext() *GetConversaionSpaceRequestTenantContext {
	return s.TenantContext
}

func (s *GetConversaionSpaceRequest) SetOpenConversationId(v string) *GetConversaionSpaceRequest {
	s.OpenConversationId = &v
	return s
}

func (s *GetConversaionSpaceRequest) SetTenantContext(v *GetConversaionSpaceRequestTenantContext) *GetConversaionSpaceRequest {
	s.TenantContext = v
	return s
}

func (s *GetConversaionSpaceRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetConversaionSpaceRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetConversaionSpaceRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetConversaionSpaceRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetConversaionSpaceRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetConversaionSpaceRequestTenantContext) SetTenantId(v string) *GetConversaionSpaceRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetConversaionSpaceRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
