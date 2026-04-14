// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScenegroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpenConversationId(v string) *GetScenegroupRequest
	GetOpenConversationId() *string
	SetTenantContext(v *GetScenegroupRequestTenantContext) *GetScenegroupRequest
	GetTenantContext() *GetScenegroupRequestTenantContext
}

type GetScenegroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cidt*****Xa4K10w==
	OpenConversationId *string                            `json:"OpenConversationId,omitempty" xml:"OpenConversationId,omitempty"`
	TenantContext      *GetScenegroupRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s GetScenegroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScenegroupRequest) GoString() string {
	return s.String()
}

func (s *GetScenegroupRequest) GetOpenConversationId() *string {
	return s.OpenConversationId
}

func (s *GetScenegroupRequest) GetTenantContext() *GetScenegroupRequestTenantContext {
	return s.TenantContext
}

func (s *GetScenegroupRequest) SetOpenConversationId(v string) *GetScenegroupRequest {
	s.OpenConversationId = &v
	return s
}

func (s *GetScenegroupRequest) SetTenantContext(v *GetScenegroupRequestTenantContext) *GetScenegroupRequest {
	s.TenantContext = v
	return s
}

func (s *GetScenegroupRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetScenegroupRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetScenegroupRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetScenegroupRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetScenegroupRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetScenegroupRequestTenantContext) SetTenantId(v string) *GetScenegroupRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetScenegroupRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
