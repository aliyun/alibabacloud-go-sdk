// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileUploadInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOption(v *GetFileUploadInfoRequestOption) *GetFileUploadInfoRequest
	GetOption() *GetFileUploadInfoRequestOption
	SetParentDentryUuid(v string) *GetFileUploadInfoRequest
	GetParentDentryUuid() *string
	SetProtocol(v string) *GetFileUploadInfoRequest
	GetProtocol() *string
	SetTenantContext(v *GetFileUploadInfoRequestTenantContext) *GetFileUploadInfoRequest
	GetTenantContext() *GetFileUploadInfoRequestTenantContext
}

type GetFileUploadInfoRequest struct {
	Option *GetFileUploadInfoRequestOption `json:"Option,omitempty" xml:"Option,omitempty" type:"Struct"`
	// example:
	//
	// dentryUuid
	ParentDentryUuid *string `json:"ParentDentryUuid,omitempty" xml:"ParentDentryUuid,omitempty"`
	// example:
	//
	// HEADER_SIGNATURE
	Protocol      *string                                `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	TenantContext *GetFileUploadInfoRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s GetFileUploadInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFileUploadInfoRequest) GoString() string {
	return s.String()
}

func (s *GetFileUploadInfoRequest) GetOption() *GetFileUploadInfoRequestOption {
	return s.Option
}

func (s *GetFileUploadInfoRequest) GetParentDentryUuid() *string {
	return s.ParentDentryUuid
}

func (s *GetFileUploadInfoRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *GetFileUploadInfoRequest) GetTenantContext() *GetFileUploadInfoRequestTenantContext {
	return s.TenantContext
}

func (s *GetFileUploadInfoRequest) SetOption(v *GetFileUploadInfoRequestOption) *GetFileUploadInfoRequest {
	s.Option = v
	return s
}

func (s *GetFileUploadInfoRequest) SetParentDentryUuid(v string) *GetFileUploadInfoRequest {
	s.ParentDentryUuid = &v
	return s
}

func (s *GetFileUploadInfoRequest) SetProtocol(v string) *GetFileUploadInfoRequest {
	s.Protocol = &v
	return s
}

func (s *GetFileUploadInfoRequest) SetTenantContext(v *GetFileUploadInfoRequestTenantContext) *GetFileUploadInfoRequest {
	s.TenantContext = v
	return s
}

func (s *GetFileUploadInfoRequest) Validate() error {
	return dara.Validate(s)
}

type GetFileUploadInfoRequestOption struct {
	PreCheckParam *GetFileUploadInfoRequestOptionPreCheckParam `json:"PreCheckParam,omitempty" xml:"PreCheckParam,omitempty" type:"Struct"`
	// example:
	//
	// true
	PreferIntranet *bool `json:"PreferIntranet,omitempty" xml:"PreferIntranet,omitempty"`
	// example:
	//
	// ZHANGJIAKOU
	PreferRegion *string `json:"PreferRegion,omitempty" xml:"PreferRegion,omitempty"`
	// example:
	//
	// DINGTALK
	StorageDriver *string `json:"StorageDriver,omitempty" xml:"StorageDriver,omitempty"`
}

func (s GetFileUploadInfoRequestOption) String() string {
	return dara.Prettify(s)
}

func (s GetFileUploadInfoRequestOption) GoString() string {
	return s.String()
}

func (s *GetFileUploadInfoRequestOption) GetPreCheckParam() *GetFileUploadInfoRequestOptionPreCheckParam {
	return s.PreCheckParam
}

func (s *GetFileUploadInfoRequestOption) GetPreferIntranet() *bool {
	return s.PreferIntranet
}

func (s *GetFileUploadInfoRequestOption) GetPreferRegion() *string {
	return s.PreferRegion
}

func (s *GetFileUploadInfoRequestOption) GetStorageDriver() *string {
	return s.StorageDriver
}

func (s *GetFileUploadInfoRequestOption) SetPreCheckParam(v *GetFileUploadInfoRequestOptionPreCheckParam) *GetFileUploadInfoRequestOption {
	s.PreCheckParam = v
	return s
}

func (s *GetFileUploadInfoRequestOption) SetPreferIntranet(v bool) *GetFileUploadInfoRequestOption {
	s.PreferIntranet = &v
	return s
}

func (s *GetFileUploadInfoRequestOption) SetPreferRegion(v string) *GetFileUploadInfoRequestOption {
	s.PreferRegion = &v
	return s
}

func (s *GetFileUploadInfoRequestOption) SetStorageDriver(v string) *GetFileUploadInfoRequestOption {
	s.StorageDriver = &v
	return s
}

func (s *GetFileUploadInfoRequestOption) Validate() error {
	return dara.Validate(s)
}

type GetFileUploadInfoRequestOptionPreCheckParam struct {
	// example:
	//
	// None
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// None
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s GetFileUploadInfoRequestOptionPreCheckParam) String() string {
	return dara.Prettify(s)
}

func (s GetFileUploadInfoRequestOptionPreCheckParam) GoString() string {
	return s.String()
}

func (s *GetFileUploadInfoRequestOptionPreCheckParam) GetName() *string {
	return s.Name
}

func (s *GetFileUploadInfoRequestOptionPreCheckParam) GetSize() *int64 {
	return s.Size
}

func (s *GetFileUploadInfoRequestOptionPreCheckParam) SetName(v string) *GetFileUploadInfoRequestOptionPreCheckParam {
	s.Name = &v
	return s
}

func (s *GetFileUploadInfoRequestOptionPreCheckParam) SetSize(v int64) *GetFileUploadInfoRequestOptionPreCheckParam {
	s.Size = &v
	return s
}

func (s *GetFileUploadInfoRequestOptionPreCheckParam) Validate() error {
	return dara.Validate(s)
}

type GetFileUploadInfoRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetFileUploadInfoRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetFileUploadInfoRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetFileUploadInfoRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetFileUploadInfoRequestTenantContext) SetTenantId(v string) *GetFileUploadInfoRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetFileUploadInfoRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
