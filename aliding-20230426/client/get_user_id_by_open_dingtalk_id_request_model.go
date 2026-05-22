// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserIdByOpenDingtalkIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *GetUserIdByOpenDingtalkIdRequestTenantContext) *GetUserIdByOpenDingtalkIdRequest
	GetTenantContext() *GetUserIdByOpenDingtalkIdRequestTenantContext
	SetOpenDingtalkId(v string) *GetUserIdByOpenDingtalkIdRequest
	GetOpenDingtalkId() *string
}

type GetUserIdByOpenDingtalkIdRequest struct {
	TenantContext *GetUserIdByOpenDingtalkIdRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// DTOJdYJ2IQC4HuexhtjsSXXXX
	OpenDingtalkId *string `json:"openDingtalkId,omitempty" xml:"openDingtalkId,omitempty"`
}

func (s GetUserIdByOpenDingtalkIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdByOpenDingtalkIdRequest) GoString() string {
	return s.String()
}

func (s *GetUserIdByOpenDingtalkIdRequest) GetTenantContext() *GetUserIdByOpenDingtalkIdRequestTenantContext {
	return s.TenantContext
}

func (s *GetUserIdByOpenDingtalkIdRequest) GetOpenDingtalkId() *string {
	return s.OpenDingtalkId
}

func (s *GetUserIdByOpenDingtalkIdRequest) SetTenantContext(v *GetUserIdByOpenDingtalkIdRequestTenantContext) *GetUserIdByOpenDingtalkIdRequest {
	s.TenantContext = v
	return s
}

func (s *GetUserIdByOpenDingtalkIdRequest) SetOpenDingtalkId(v string) *GetUserIdByOpenDingtalkIdRequest {
	s.OpenDingtalkId = &v
	return s
}

func (s *GetUserIdByOpenDingtalkIdRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserIdByOpenDingtalkIdRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetUserIdByOpenDingtalkIdRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdByOpenDingtalkIdRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetUserIdByOpenDingtalkIdRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetUserIdByOpenDingtalkIdRequestTenantContext) SetTenantId(v string) *GetUserIdByOpenDingtalkIdRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetUserIdByOpenDingtalkIdRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
