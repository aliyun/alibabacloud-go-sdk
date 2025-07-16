// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileDownloadInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDentryId(v string) *GetFileDownloadInfoRequest
	GetDentryId() *string
	SetOption(v *GetFileDownloadInfoRequestOption) *GetFileDownloadInfoRequest
	GetOption() *GetFileDownloadInfoRequestOption
	SetSpaceId(v string) *GetFileDownloadInfoRequest
	GetSpaceId() *string
	SetTenantContext(v *GetFileDownloadInfoRequestTenantContext) *GetFileDownloadInfoRequest
	GetTenantContext() *GetFileDownloadInfoRequestTenantContext
}

type GetFileDownloadInfoRequest struct {
	// example:
	//
	// 798xxxxx
	DentryId *string                           `json:"DentryId,omitempty" xml:"DentryId,omitempty"`
	Option   *GetFileDownloadInfoRequestOption `json:"Option,omitempty" xml:"Option,omitempty" type:"Struct"`
	// example:
	//
	// 854xxxx
	SpaceId       *string                                  `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	TenantContext *GetFileDownloadInfoRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s GetFileDownloadInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFileDownloadInfoRequest) GoString() string {
	return s.String()
}

func (s *GetFileDownloadInfoRequest) GetDentryId() *string {
	return s.DentryId
}

func (s *GetFileDownloadInfoRequest) GetOption() *GetFileDownloadInfoRequestOption {
	return s.Option
}

func (s *GetFileDownloadInfoRequest) GetSpaceId() *string {
	return s.SpaceId
}

func (s *GetFileDownloadInfoRequest) GetTenantContext() *GetFileDownloadInfoRequestTenantContext {
	return s.TenantContext
}

func (s *GetFileDownloadInfoRequest) SetDentryId(v string) *GetFileDownloadInfoRequest {
	s.DentryId = &v
	return s
}

func (s *GetFileDownloadInfoRequest) SetOption(v *GetFileDownloadInfoRequestOption) *GetFileDownloadInfoRequest {
	s.Option = v
	return s
}

func (s *GetFileDownloadInfoRequest) SetSpaceId(v string) *GetFileDownloadInfoRequest {
	s.SpaceId = &v
	return s
}

func (s *GetFileDownloadInfoRequest) SetTenantContext(v *GetFileDownloadInfoRequestTenantContext) *GetFileDownloadInfoRequest {
	s.TenantContext = v
	return s
}

func (s *GetFileDownloadInfoRequest) Validate() error {
	return dara.Validate(s)
}

type GetFileDownloadInfoRequestOption struct {
	// example:
	//
	// true
	PreferIntranet *bool `json:"PreferIntranet,omitempty" xml:"PreferIntranet,omitempty"`
	// example:
	//
	// 1
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetFileDownloadInfoRequestOption) String() string {
	return dara.Prettify(s)
}

func (s GetFileDownloadInfoRequestOption) GoString() string {
	return s.String()
}

func (s *GetFileDownloadInfoRequestOption) GetPreferIntranet() *bool {
	return s.PreferIntranet
}

func (s *GetFileDownloadInfoRequestOption) GetVersion() *int64 {
	return s.Version
}

func (s *GetFileDownloadInfoRequestOption) SetPreferIntranet(v bool) *GetFileDownloadInfoRequestOption {
	s.PreferIntranet = &v
	return s
}

func (s *GetFileDownloadInfoRequestOption) SetVersion(v int64) *GetFileDownloadInfoRequestOption {
	s.Version = &v
	return s
}

func (s *GetFileDownloadInfoRequestOption) Validate() error {
	return dara.Validate(s)
}

type GetFileDownloadInfoRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetFileDownloadInfoRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetFileDownloadInfoRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetFileDownloadInfoRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetFileDownloadInfoRequestTenantContext) SetTenantId(v string) *GetFileDownloadInfoRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetFileDownloadInfoRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
