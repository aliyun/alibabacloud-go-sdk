// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitMultipartFileUploadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOption(v *InitMultipartFileUploadRequestOption) *InitMultipartFileUploadRequest
	GetOption() *InitMultipartFileUploadRequestOption
	SetParentDentryUuid(v string) *InitMultipartFileUploadRequest
	GetParentDentryUuid() *string
	SetTenantContext(v *InitMultipartFileUploadRequestTenantContext) *InitMultipartFileUploadRequest
	GetTenantContext() *InitMultipartFileUploadRequestTenantContext
}

type InitMultipartFileUploadRequest struct {
	Option *InitMultipartFileUploadRequestOption `json:"Option,omitempty" xml:"Option,omitempty" type:"Struct"`
	// example:
	//
	// dentryUuid
	ParentDentryUuid *string                                      `json:"ParentDentryUuid,omitempty" xml:"ParentDentryUuid,omitempty"`
	TenantContext    *InitMultipartFileUploadRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s InitMultipartFileUploadRequest) String() string {
	return dara.Prettify(s)
}

func (s InitMultipartFileUploadRequest) GoString() string {
	return s.String()
}

func (s *InitMultipartFileUploadRequest) GetOption() *InitMultipartFileUploadRequestOption {
	return s.Option
}

func (s *InitMultipartFileUploadRequest) GetParentDentryUuid() *string {
	return s.ParentDentryUuid
}

func (s *InitMultipartFileUploadRequest) GetTenantContext() *InitMultipartFileUploadRequestTenantContext {
	return s.TenantContext
}

func (s *InitMultipartFileUploadRequest) SetOption(v *InitMultipartFileUploadRequestOption) *InitMultipartFileUploadRequest {
	s.Option = v
	return s
}

func (s *InitMultipartFileUploadRequest) SetParentDentryUuid(v string) *InitMultipartFileUploadRequest {
	s.ParentDentryUuid = &v
	return s
}

func (s *InitMultipartFileUploadRequest) SetTenantContext(v *InitMultipartFileUploadRequestTenantContext) *InitMultipartFileUploadRequest {
	s.TenantContext = v
	return s
}

func (s *InitMultipartFileUploadRequest) Validate() error {
	if s.Option != nil {
		if err := s.Option.Validate(); err != nil {
			return err
		}
	}
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InitMultipartFileUploadRequestOption struct {
	PreCheckParam *InitMultipartFileUploadRequestOptionPreCheckParam `json:"PreCheckParam,omitempty" xml:"PreCheckParam,omitempty" type:"Struct"`
	// example:
	//
	// ZHANGJIAKOU
	PreferRegion *string `json:"PreferRegion,omitempty" xml:"PreferRegion,omitempty"`
	// example:
	//
	// DINGTALK
	StorageDriver *string `json:"StorageDriver,omitempty" xml:"StorageDriver,omitempty"`
}

func (s InitMultipartFileUploadRequestOption) String() string {
	return dara.Prettify(s)
}

func (s InitMultipartFileUploadRequestOption) GoString() string {
	return s.String()
}

func (s *InitMultipartFileUploadRequestOption) GetPreCheckParam() *InitMultipartFileUploadRequestOptionPreCheckParam {
	return s.PreCheckParam
}

func (s *InitMultipartFileUploadRequestOption) GetPreferRegion() *string {
	return s.PreferRegion
}

func (s *InitMultipartFileUploadRequestOption) GetStorageDriver() *string {
	return s.StorageDriver
}

func (s *InitMultipartFileUploadRequestOption) SetPreCheckParam(v *InitMultipartFileUploadRequestOptionPreCheckParam) *InitMultipartFileUploadRequestOption {
	s.PreCheckParam = v
	return s
}

func (s *InitMultipartFileUploadRequestOption) SetPreferRegion(v string) *InitMultipartFileUploadRequestOption {
	s.PreferRegion = &v
	return s
}

func (s *InitMultipartFileUploadRequestOption) SetStorageDriver(v string) *InitMultipartFileUploadRequestOption {
	s.StorageDriver = &v
	return s
}

func (s *InitMultipartFileUploadRequestOption) Validate() error {
	if s.PreCheckParam != nil {
		if err := s.PreCheckParam.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InitMultipartFileUploadRequestOptionPreCheckParam struct {
	// example:
	//
	// md5
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// example:
	//
	// 100
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s InitMultipartFileUploadRequestOptionPreCheckParam) String() string {
	return dara.Prettify(s)
}

func (s InitMultipartFileUploadRequestOptionPreCheckParam) GoString() string {
	return s.String()
}

func (s *InitMultipartFileUploadRequestOptionPreCheckParam) GetMd5() *string {
	return s.Md5
}

func (s *InitMultipartFileUploadRequestOptionPreCheckParam) GetName() *string {
	return s.Name
}

func (s *InitMultipartFileUploadRequestOptionPreCheckParam) GetParentId() *string {
	return s.ParentId
}

func (s *InitMultipartFileUploadRequestOptionPreCheckParam) GetSize() *int64 {
	return s.Size
}

func (s *InitMultipartFileUploadRequestOptionPreCheckParam) SetMd5(v string) *InitMultipartFileUploadRequestOptionPreCheckParam {
	s.Md5 = &v
	return s
}

func (s *InitMultipartFileUploadRequestOptionPreCheckParam) SetName(v string) *InitMultipartFileUploadRequestOptionPreCheckParam {
	s.Name = &v
	return s
}

func (s *InitMultipartFileUploadRequestOptionPreCheckParam) SetParentId(v string) *InitMultipartFileUploadRequestOptionPreCheckParam {
	s.ParentId = &v
	return s
}

func (s *InitMultipartFileUploadRequestOptionPreCheckParam) SetSize(v int64) *InitMultipartFileUploadRequestOptionPreCheckParam {
	s.Size = &v
	return s
}

func (s *InitMultipartFileUploadRequestOptionPreCheckParam) Validate() error {
	return dara.Validate(s)
}

type InitMultipartFileUploadRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s InitMultipartFileUploadRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s InitMultipartFileUploadRequestTenantContext) GoString() string {
	return s.String()
}

func (s *InitMultipartFileUploadRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *InitMultipartFileUploadRequestTenantContext) SetTenantId(v string) *InitMultipartFileUploadRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *InitMultipartFileUploadRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
