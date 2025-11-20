// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommitFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CommitFileRequest
	GetName() *string
	SetOption(v *CommitFileRequestOption) *CommitFileRequest
	GetOption() *CommitFileRequestOption
	SetParentDentryUuid(v string) *CommitFileRequest
	GetParentDentryUuid() *string
	SetTenantContext(v *CommitFileRequestTenantContext) *CommitFileRequest
	GetTenantContext() *CommitFileRequestTenantContext
	SetUploadKey(v string) *CommitFileRequest
	GetUploadKey() *string
}

type CommitFileRequest struct {
	// example:
	//
	// None
	Name   *string                  `json:"Name,omitempty" xml:"Name,omitempty"`
	Option *CommitFileRequestOption `json:"Option,omitempty" xml:"Option,omitempty" type:"Struct"`
	// example:
	//
	// dentryUuid
	ParentDentryUuid *string                         `json:"ParentDentryUuid,omitempty" xml:"ParentDentryUuid,omitempty"`
	TenantContext    *CommitFileRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// example:
	//
	// upload_key
	UploadKey *string `json:"UploadKey,omitempty" xml:"UploadKey,omitempty"`
}

func (s CommitFileRequest) String() string {
	return dara.Prettify(s)
}

func (s CommitFileRequest) GoString() string {
	return s.String()
}

func (s *CommitFileRequest) GetName() *string {
	return s.Name
}

func (s *CommitFileRequest) GetOption() *CommitFileRequestOption {
	return s.Option
}

func (s *CommitFileRequest) GetParentDentryUuid() *string {
	return s.ParentDentryUuid
}

func (s *CommitFileRequest) GetTenantContext() *CommitFileRequestTenantContext {
	return s.TenantContext
}

func (s *CommitFileRequest) GetUploadKey() *string {
	return s.UploadKey
}

func (s *CommitFileRequest) SetName(v string) *CommitFileRequest {
	s.Name = &v
	return s
}

func (s *CommitFileRequest) SetOption(v *CommitFileRequestOption) *CommitFileRequest {
	s.Option = v
	return s
}

func (s *CommitFileRequest) SetParentDentryUuid(v string) *CommitFileRequest {
	s.ParentDentryUuid = &v
	return s
}

func (s *CommitFileRequest) SetTenantContext(v *CommitFileRequestTenantContext) *CommitFileRequest {
	s.TenantContext = v
	return s
}

func (s *CommitFileRequest) SetUploadKey(v string) *CommitFileRequest {
	s.UploadKey = &v
	return s
}

func (s *CommitFileRequest) Validate() error {
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

type CommitFileRequestOption struct {
	AppProperties []*CommitFileRequestOptionAppProperties `json:"AppProperties,omitempty" xml:"AppProperties,omitempty" type:"Repeated"`
	// example:
	//
	// AUTO_RENAME
	ConflictStrategy *string `json:"ConflictStrategy,omitempty" xml:"ConflictStrategy,omitempty"`
	// example:
	//
	// false
	ConvertToOnlineDoc *bool `json:"ConvertToOnlineDoc,omitempty" xml:"ConvertToOnlineDoc,omitempty"`
	// example:
	//
	// DOC
	ConvertToOnlineDocTargetDocumentType *string `json:"ConvertToOnlineDocTargetDocumentType,omitempty" xml:"ConvertToOnlineDocTargetDocumentType,omitempty"`
	// example:
	//
	// 512
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s CommitFileRequestOption) String() string {
	return dara.Prettify(s)
}

func (s CommitFileRequestOption) GoString() string {
	return s.String()
}

func (s *CommitFileRequestOption) GetAppProperties() []*CommitFileRequestOptionAppProperties {
	return s.AppProperties
}

func (s *CommitFileRequestOption) GetConflictStrategy() *string {
	return s.ConflictStrategy
}

func (s *CommitFileRequestOption) GetConvertToOnlineDoc() *bool {
	return s.ConvertToOnlineDoc
}

func (s *CommitFileRequestOption) GetConvertToOnlineDocTargetDocumentType() *string {
	return s.ConvertToOnlineDocTargetDocumentType
}

func (s *CommitFileRequestOption) GetSize() *int64 {
	return s.Size
}

func (s *CommitFileRequestOption) SetAppProperties(v []*CommitFileRequestOptionAppProperties) *CommitFileRequestOption {
	s.AppProperties = v
	return s
}

func (s *CommitFileRequestOption) SetConflictStrategy(v string) *CommitFileRequestOption {
	s.ConflictStrategy = &v
	return s
}

func (s *CommitFileRequestOption) SetConvertToOnlineDoc(v bool) *CommitFileRequestOption {
	s.ConvertToOnlineDoc = &v
	return s
}

func (s *CommitFileRequestOption) SetConvertToOnlineDocTargetDocumentType(v string) *CommitFileRequestOption {
	s.ConvertToOnlineDocTargetDocumentType = &v
	return s
}

func (s *CommitFileRequestOption) SetSize(v int64) *CommitFileRequestOption {
	s.Size = &v
	return s
}

func (s *CommitFileRequestOption) Validate() error {
	if s.AppProperties != nil {
		for _, item := range s.AppProperties {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CommitFileRequestOptionAppProperties struct {
	// example:
	//
	// property_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// property_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// example:
	//
	// PUBLIC
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s CommitFileRequestOptionAppProperties) String() string {
	return dara.Prettify(s)
}

func (s CommitFileRequestOptionAppProperties) GoString() string {
	return s.String()
}

func (s *CommitFileRequestOptionAppProperties) GetName() *string {
	return s.Name
}

func (s *CommitFileRequestOptionAppProperties) GetValue() *string {
	return s.Value
}

func (s *CommitFileRequestOptionAppProperties) GetVisibility() *string {
	return s.Visibility
}

func (s *CommitFileRequestOptionAppProperties) SetName(v string) *CommitFileRequestOptionAppProperties {
	s.Name = &v
	return s
}

func (s *CommitFileRequestOptionAppProperties) SetValue(v string) *CommitFileRequestOptionAppProperties {
	s.Value = &v
	return s
}

func (s *CommitFileRequestOptionAppProperties) SetVisibility(v string) *CommitFileRequestOptionAppProperties {
	s.Visibility = &v
	return s
}

func (s *CommitFileRequestOptionAppProperties) Validate() error {
	return dara.Validate(s)
}

type CommitFileRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CommitFileRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s CommitFileRequestTenantContext) GoString() string {
	return s.String()
}

func (s *CommitFileRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *CommitFileRequestTenantContext) SetTenantId(v string) *CommitFileRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *CommitFileRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
