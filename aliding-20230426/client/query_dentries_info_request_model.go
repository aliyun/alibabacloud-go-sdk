// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDentriesInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppIdsForAppProperties(v []*string) *QueryDentriesInfoRequest
	GetAppIdsForAppProperties() []*string
	SetDentryId(v string) *QueryDentriesInfoRequest
	GetDentryId() *string
	SetSpaceId(v string) *QueryDentriesInfoRequest
	GetSpaceId() *string
	SetTenantContext(v *QueryDentriesInfoRequestTenantContext) *QueryDentriesInfoRequest
	GetTenantContext() *QueryDentriesInfoRequestTenantContext
	SetWithThumbnail(v bool) *QueryDentriesInfoRequest
	GetWithThumbnail() *bool
}

type QueryDentriesInfoRequest struct {
	AppIdsForAppProperties []*string `json:"AppIdsForAppProperties,omitempty" xml:"AppIdsForAppProperties,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 140901622636
	DentryId *string `json:"DentryId,omitempty" xml:"DentryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 22443475065
	SpaceId       *string                                `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	TenantContext *QueryDentriesInfoRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// example:
	//
	// false
	WithThumbnail *bool `json:"WithThumbnail,omitempty" xml:"WithThumbnail,omitempty"`
}

func (s QueryDentriesInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDentriesInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryDentriesInfoRequest) GetAppIdsForAppProperties() []*string {
	return s.AppIdsForAppProperties
}

func (s *QueryDentriesInfoRequest) GetDentryId() *string {
	return s.DentryId
}

func (s *QueryDentriesInfoRequest) GetSpaceId() *string {
	return s.SpaceId
}

func (s *QueryDentriesInfoRequest) GetTenantContext() *QueryDentriesInfoRequestTenantContext {
	return s.TenantContext
}

func (s *QueryDentriesInfoRequest) GetWithThumbnail() *bool {
	return s.WithThumbnail
}

func (s *QueryDentriesInfoRequest) SetAppIdsForAppProperties(v []*string) *QueryDentriesInfoRequest {
	s.AppIdsForAppProperties = v
	return s
}

func (s *QueryDentriesInfoRequest) SetDentryId(v string) *QueryDentriesInfoRequest {
	s.DentryId = &v
	return s
}

func (s *QueryDentriesInfoRequest) SetSpaceId(v string) *QueryDentriesInfoRequest {
	s.SpaceId = &v
	return s
}

func (s *QueryDentriesInfoRequest) SetTenantContext(v *QueryDentriesInfoRequestTenantContext) *QueryDentriesInfoRequest {
	s.TenantContext = v
	return s
}

func (s *QueryDentriesInfoRequest) SetWithThumbnail(v bool) *QueryDentriesInfoRequest {
	s.WithThumbnail = &v
	return s
}

func (s *QueryDentriesInfoRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDentriesInfoRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QueryDentriesInfoRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s QueryDentriesInfoRequestTenantContext) GoString() string {
	return s.String()
}

func (s *QueryDentriesInfoRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryDentriesInfoRequestTenantContext) SetTenantId(v string) *QueryDentriesInfoRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *QueryDentriesInfoRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
