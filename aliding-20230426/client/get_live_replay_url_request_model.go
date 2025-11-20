// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLiveReplayUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLiveId(v string) *GetLiveReplayUrlRequest
	GetLiveId() *string
	SetTenantContext(v *GetLiveReplayUrlRequestTenantContext) *GetLiveReplayUrlRequest
	GetTenantContext() *GetLiveReplayUrlRequestTenantContext
}

type GetLiveReplayUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4d38xxxxx
	LiveId        *string                               `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	TenantContext *GetLiveReplayUrlRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s GetLiveReplayUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLiveReplayUrlRequest) GoString() string {
	return s.String()
}

func (s *GetLiveReplayUrlRequest) GetLiveId() *string {
	return s.LiveId
}

func (s *GetLiveReplayUrlRequest) GetTenantContext() *GetLiveReplayUrlRequestTenantContext {
	return s.TenantContext
}

func (s *GetLiveReplayUrlRequest) SetLiveId(v string) *GetLiveReplayUrlRequest {
	s.LiveId = &v
	return s
}

func (s *GetLiveReplayUrlRequest) SetTenantContext(v *GetLiveReplayUrlRequestTenantContext) *GetLiveReplayUrlRequest {
	s.TenantContext = v
	return s
}

func (s *GetLiveReplayUrlRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLiveReplayUrlRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetLiveReplayUrlRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetLiveReplayUrlRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetLiveReplayUrlRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetLiveReplayUrlRequestTenantContext) SetTenantId(v string) *GetLiveReplayUrlRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetLiveReplayUrlRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
