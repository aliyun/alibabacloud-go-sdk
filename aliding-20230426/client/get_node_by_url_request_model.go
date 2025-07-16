// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeByUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOption(v *GetNodeByUrlRequestOption) *GetNodeByUrlRequest
	GetOption() *GetNodeByUrlRequestOption
	SetTenantContext(v *GetNodeByUrlRequestTenantContext) *GetNodeByUrlRequest
	GetTenantContext() *GetNodeByUrlRequestTenantContext
	SetUrl(v string) *GetNodeByUrlRequest
	GetUrl() *string
}

type GetNodeByUrlRequest struct {
	Option        *GetNodeByUrlRequestOption        `json:"Option,omitempty" xml:"Option,omitempty" type:"Struct"`
	TenantContext *GetNodeByUrlRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// https://alidocs.dingtalk.com/i/nodes/EpGBa2L*********gN7R35y
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetNodeByUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNodeByUrlRequest) GoString() string {
	return s.String()
}

func (s *GetNodeByUrlRequest) GetOption() *GetNodeByUrlRequestOption {
	return s.Option
}

func (s *GetNodeByUrlRequest) GetTenantContext() *GetNodeByUrlRequestTenantContext {
	return s.TenantContext
}

func (s *GetNodeByUrlRequest) GetUrl() *string {
	return s.Url
}

func (s *GetNodeByUrlRequest) SetOption(v *GetNodeByUrlRequestOption) *GetNodeByUrlRequest {
	s.Option = v
	return s
}

func (s *GetNodeByUrlRequest) SetTenantContext(v *GetNodeByUrlRequestTenantContext) *GetNodeByUrlRequest {
	s.TenantContext = v
	return s
}

func (s *GetNodeByUrlRequest) SetUrl(v string) *GetNodeByUrlRequest {
	s.Url = &v
	return s
}

func (s *GetNodeByUrlRequest) Validate() error {
	return dara.Validate(s)
}

type GetNodeByUrlRequestOption struct {
	// example:
	//
	// false
	WithPermissionRole *bool `json:"WithPermissionRole,omitempty" xml:"WithPermissionRole,omitempty"`
	// example:
	//
	// false
	WithStatisticalInfo *bool `json:"WithStatisticalInfo,omitempty" xml:"WithStatisticalInfo,omitempty"`
}

func (s GetNodeByUrlRequestOption) String() string {
	return dara.Prettify(s)
}

func (s GetNodeByUrlRequestOption) GoString() string {
	return s.String()
}

func (s *GetNodeByUrlRequestOption) GetWithPermissionRole() *bool {
	return s.WithPermissionRole
}

func (s *GetNodeByUrlRequestOption) GetWithStatisticalInfo() *bool {
	return s.WithStatisticalInfo
}

func (s *GetNodeByUrlRequestOption) SetWithPermissionRole(v bool) *GetNodeByUrlRequestOption {
	s.WithPermissionRole = &v
	return s
}

func (s *GetNodeByUrlRequestOption) SetWithStatisticalInfo(v bool) *GetNodeByUrlRequestOption {
	s.WithStatisticalInfo = &v
	return s
}

func (s *GetNodeByUrlRequestOption) Validate() error {
	return dara.Validate(s)
}

type GetNodeByUrlRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetNodeByUrlRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetNodeByUrlRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetNodeByUrlRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetNodeByUrlRequestTenantContext) SetTenantId(v string) *GetNodeByUrlRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetNodeByUrlRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
