// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRegisterLineageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddRegisterLineageCommand(v *AddRegisterLineageRequestAddRegisterLineageCommand) *AddRegisterLineageRequest
	GetAddRegisterLineageCommand() *AddRegisterLineageRequestAddRegisterLineageCommand
	SetOpTenantId(v int64) *AddRegisterLineageRequest
	GetOpTenantId() *int64
}

type AddRegisterLineageRequest struct {
	// This parameter is required.
	AddRegisterLineageCommand *AddRegisterLineageRequestAddRegisterLineageCommand `json:"AddRegisterLineageCommand,omitempty" xml:"AddRegisterLineageCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s AddRegisterLineageRequest) String() string {
	return dara.Prettify(s)
}

func (s AddRegisterLineageRequest) GoString() string {
	return s.String()
}

func (s *AddRegisterLineageRequest) GetAddRegisterLineageCommand() *AddRegisterLineageRequestAddRegisterLineageCommand {
	return s.AddRegisterLineageCommand
}

func (s *AddRegisterLineageRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *AddRegisterLineageRequest) SetAddRegisterLineageCommand(v *AddRegisterLineageRequestAddRegisterLineageCommand) *AddRegisterLineageRequest {
	s.AddRegisterLineageCommand = v
	return s
}

func (s *AddRegisterLineageRequest) SetOpTenantId(v int64) *AddRegisterLineageRequest {
	s.OpTenantId = &v
	return s
}

func (s *AddRegisterLineageRequest) Validate() error {
	if s.AddRegisterLineageCommand != nil {
		if err := s.AddRegisterLineageCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddRegisterLineageRequestAddRegisterLineageCommand struct {
	CheckAssetExist    *bool                                                                 `json:"CheckAssetExist,omitempty" xml:"CheckAssetExist,omitempty"`
	DetailedLineages   []*AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineages `json:"DetailedLineages,omitempty" xml:"DetailedLineages,omitempty" type:"Repeated"`
	RelationProperties map[string]interface{}                                                `json:"RelationProperties,omitempty" xml:"RelationProperties,omitempty"`
	// This parameter is required.
	Source *AddRegisterLineageRequestAddRegisterLineageCommandSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	// This parameter is required.
	Target *AddRegisterLineageRequestAddRegisterLineageCommandTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	// example:
	//
	// 300001234
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// 300004567
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AddRegisterLineageRequestAddRegisterLineageCommand) String() string {
	return dara.Prettify(s)
}

func (s AddRegisterLineageRequestAddRegisterLineageCommand) GoString() string {
	return s.String()
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommand) GetCheckAssetExist() *bool {
	return s.CheckAssetExist
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommand) GetDetailedLineages() []*AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineages {
	return s.DetailedLineages
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommand) GetRelationProperties() map[string]interface{} {
	return s.RelationProperties
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommand) GetSource() *AddRegisterLineageRequestAddRegisterLineageCommandSource {
	return s.Source
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommand) GetTarget() *AddRegisterLineageRequestAddRegisterLineageCommandTarget {
	return s.Target
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommand) GetTenantId() *int64 {
	return s.TenantId
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommand) GetUserId() *string {
	return s.UserId
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommand) SetCheckAssetExist(v bool) *AddRegisterLineageRequestAddRegisterLineageCommand {
	s.CheckAssetExist = &v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommand) SetDetailedLineages(v []*AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineages) *AddRegisterLineageRequestAddRegisterLineageCommand {
	s.DetailedLineages = v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommand) SetRelationProperties(v map[string]interface{}) *AddRegisterLineageRequestAddRegisterLineageCommand {
	s.RelationProperties = v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommand) SetSource(v *AddRegisterLineageRequestAddRegisterLineageCommandSource) *AddRegisterLineageRequestAddRegisterLineageCommand {
	s.Source = v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommand) SetTarget(v *AddRegisterLineageRequestAddRegisterLineageCommandTarget) *AddRegisterLineageRequestAddRegisterLineageCommand {
	s.Target = v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommand) SetTenantId(v int64) *AddRegisterLineageRequestAddRegisterLineageCommand {
	s.TenantId = &v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommand) SetUserId(v string) *AddRegisterLineageRequestAddRegisterLineageCommand {
	s.UserId = &v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommand) Validate() error {
	if s.DetailedLineages != nil {
		for _, item := range s.DetailedLineages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Source != nil {
		if err := s.Source.Validate(); err != nil {
			return err
		}
	}
	if s.Target != nil {
		if err := s.Target.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineages struct {
	IsDirect *bool `json:"IsDirect,omitempty" xml:"IsDirect,omitempty"`
	// This parameter is required.
	Source *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	// This parameter is required.
	Target *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
}

func (s AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineages) String() string {
	return dara.Prettify(s)
}

func (s AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineages) GoString() string {
	return s.String()
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineages) GetIsDirect() *bool {
	return s.IsDirect
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineages) GetSource() *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesSource {
	return s.Source
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineages) GetTarget() *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesTarget {
	return s.Target
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineages) SetIsDirect(v bool) *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineages {
	s.IsDirect = &v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineages) SetSource(v *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesSource) *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineages {
	s.Source = v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineages) SetTarget(v *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesTarget) *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineages {
	s.Target = v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineages) Validate() error {
	if s.Source != nil {
		if err := s.Source.Validate(); err != nil {
			return err
		}
	}
	if s.Target != nil {
		if err := s.Target.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesSource struct {
	// example:
	//
	// dataphin
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// example:
	//
	// DEV, PROD
	Env           *string                `json:"Env,omitempty" xml:"Env,omitempty"`
	ExtProperties map[string]interface{} `json:"ExtProperties,omitempty" xml:"ExtProperties,omitempty"`
	// example:
	//
	// odps.300000001.project1.table1.column1
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
	// example:
	//
	// COLUMN
	MetadataType *string `json:"MetadataType,omitempty" xml:"MetadataType,omitempty"`
	// example:
	//
	// column1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// odps.300000001.project1.table1
	ParentGuid *string `json:"ParentGuid,omitempty" xml:"ParentGuid,omitempty"`
	// example:
	//
	// BY_GUID, BY_PROPERTY
	ReferenceType *string `json:"ReferenceType,omitempty" xml:"ReferenceType,omitempty"`
	// example:
	//
	// project1, bizUnit1
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
}

func (s AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesSource) String() string {
	return dara.Prettify(s)
}

func (s AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesSource) GoString() string {
	return s.String()
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesSource) GetCatalog() *string {
	return s.Catalog
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesSource) GetEnv() *string {
	return s.Env
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesSource) GetExtProperties() map[string]interface{} {
	return s.ExtProperties
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesSource) GetGuid() *string {
	return s.Guid
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesSource) GetMetadataType() *string {
	return s.MetadataType
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesSource) GetName() *string {
	return s.Name
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesSource) GetParentGuid() *string {
	return s.ParentGuid
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesSource) GetReferenceType() *string {
	return s.ReferenceType
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesSource) GetSchema() *string {
	return s.Schema
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesSource) SetCatalog(v string) *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesSource {
	s.Catalog = &v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesSource) SetEnv(v string) *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesSource {
	s.Env = &v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesSource) SetExtProperties(v map[string]interface{}) *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesSource {
	s.ExtProperties = v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesSource) SetGuid(v string) *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesSource {
	s.Guid = &v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesSource) SetMetadataType(v string) *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesSource {
	s.MetadataType = &v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesSource) SetName(v string) *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesSource {
	s.Name = &v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesSource) SetParentGuid(v string) *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesSource {
	s.ParentGuid = &v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesSource) SetReferenceType(v string) *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesSource {
	s.ReferenceType = &v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesSource) SetSchema(v string) *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesSource {
	s.Schema = &v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesSource) Validate() error {
	return dara.Validate(s)
}

type AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesTarget struct {
	// example:
	//
	// dataphin
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// example:
	//
	// DEV, PROD
	Env           *string                `json:"Env,omitempty" xml:"Env,omitempty"`
	ExtProperties map[string]interface{} `json:"ExtProperties,omitempty" xml:"ExtProperties,omitempty"`
	// example:
	//
	// odps.300000001.project1.table1.column1
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
	// example:
	//
	// COLUMN
	MetadataType *string `json:"MetadataType,omitempty" xml:"MetadataType,omitempty"`
	// example:
	//
	// column1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// odps.300000001.project1.table1
	ParentGuid *string `json:"ParentGuid,omitempty" xml:"ParentGuid,omitempty"`
	// example:
	//
	// BY_GUID, BY_PROPERTY
	ReferenceType *string `json:"ReferenceType,omitempty" xml:"ReferenceType,omitempty"`
	// example:
	//
	// project1, bizUnit1
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
}

func (s AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesTarget) String() string {
	return dara.Prettify(s)
}

func (s AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesTarget) GoString() string {
	return s.String()
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesTarget) GetCatalog() *string {
	return s.Catalog
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesTarget) GetEnv() *string {
	return s.Env
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesTarget) GetExtProperties() map[string]interface{} {
	return s.ExtProperties
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesTarget) GetGuid() *string {
	return s.Guid
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesTarget) GetMetadataType() *string {
	return s.MetadataType
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesTarget) GetName() *string {
	return s.Name
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesTarget) GetParentGuid() *string {
	return s.ParentGuid
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesTarget) GetReferenceType() *string {
	return s.ReferenceType
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesTarget) GetSchema() *string {
	return s.Schema
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesTarget) SetCatalog(v string) *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesTarget {
	s.Catalog = &v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesTarget) SetEnv(v string) *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesTarget {
	s.Env = &v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesTarget) SetExtProperties(v map[string]interface{}) *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesTarget {
	s.ExtProperties = v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesTarget) SetGuid(v string) *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesTarget {
	s.Guid = &v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesTarget) SetMetadataType(v string) *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesTarget {
	s.MetadataType = &v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesTarget) SetName(v string) *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesTarget {
	s.Name = &v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesTarget) SetParentGuid(v string) *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesTarget {
	s.ParentGuid = &v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesTarget) SetReferenceType(v string) *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesTarget {
	s.ReferenceType = &v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesTarget) SetSchema(v string) *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesTarget {
	s.Schema = &v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandDetailedLineagesTarget) Validate() error {
	return dara.Validate(s)
}

type AddRegisterLineageRequestAddRegisterLineageCommandSource struct {
	// example:
	//
	// dataphin
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// example:
	//
	// DEV, PROD
	Env           *string                `json:"Env,omitempty" xml:"Env,omitempty"`
	ExtProperties map[string]interface{} `json:"ExtProperties,omitempty" xml:"ExtProperties,omitempty"`
	// example:
	//
	// odps.300000001.project1.table1
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
	// example:
	//
	// PHYSICAL_TABLE, PHYSICAL_VIEW, PHYSICAL_MATERIALIZED_VIEW, DATASOURCE_TABLE, DATASOURCE_VIEW, DATASOURCE_MATERIALIZED_VIEW, DIM_NORMAL, DIM_LEVEL, DIM_ENUM, DIM_VIRTUAL, FACT_EVENT, FACT_PROCESS, FACT_SNAPSHOT, SUM_BIZ_UNIT
	MetadataSubType *string `json:"MetadataSubType,omitempty" xml:"MetadataSubType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TABLE
	MetadataType *string `json:"MetadataType,omitempty" xml:"MetadataType,omitempty"`
	// example:
	//
	// table1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// BY_GUID, BY_PROPERTY
	ReferenceType *string `json:"ReferenceType,omitempty" xml:"ReferenceType,omitempty"`
	// example:
	//
	// project1, bizUnit1
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
}

func (s AddRegisterLineageRequestAddRegisterLineageCommandSource) String() string {
	return dara.Prettify(s)
}

func (s AddRegisterLineageRequestAddRegisterLineageCommandSource) GoString() string {
	return s.String()
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandSource) GetCatalog() *string {
	return s.Catalog
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandSource) GetEnv() *string {
	return s.Env
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandSource) GetExtProperties() map[string]interface{} {
	return s.ExtProperties
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandSource) GetGuid() *string {
	return s.Guid
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandSource) GetMetadataSubType() *string {
	return s.MetadataSubType
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandSource) GetMetadataType() *string {
	return s.MetadataType
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandSource) GetName() *string {
	return s.Name
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandSource) GetReferenceType() *string {
	return s.ReferenceType
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandSource) GetSchema() *string {
	return s.Schema
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandSource) SetCatalog(v string) *AddRegisterLineageRequestAddRegisterLineageCommandSource {
	s.Catalog = &v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandSource) SetEnv(v string) *AddRegisterLineageRequestAddRegisterLineageCommandSource {
	s.Env = &v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandSource) SetExtProperties(v map[string]interface{}) *AddRegisterLineageRequestAddRegisterLineageCommandSource {
	s.ExtProperties = v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandSource) SetGuid(v string) *AddRegisterLineageRequestAddRegisterLineageCommandSource {
	s.Guid = &v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandSource) SetMetadataSubType(v string) *AddRegisterLineageRequestAddRegisterLineageCommandSource {
	s.MetadataSubType = &v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandSource) SetMetadataType(v string) *AddRegisterLineageRequestAddRegisterLineageCommandSource {
	s.MetadataType = &v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandSource) SetName(v string) *AddRegisterLineageRequestAddRegisterLineageCommandSource {
	s.Name = &v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandSource) SetReferenceType(v string) *AddRegisterLineageRequestAddRegisterLineageCommandSource {
	s.ReferenceType = &v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandSource) SetSchema(v string) *AddRegisterLineageRequestAddRegisterLineageCommandSource {
	s.Schema = &v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandSource) Validate() error {
	return dara.Validate(s)
}

type AddRegisterLineageRequestAddRegisterLineageCommandTarget struct {
	// example:
	//
	// dataphin
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// example:
	//
	// DEV, PROD
	Env           *string                `json:"Env,omitempty" xml:"Env,omitempty"`
	ExtProperties map[string]interface{} `json:"ExtProperties,omitempty" xml:"ExtProperties,omitempty"`
	// example:
	//
	// odps.300000001.project1.table1
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
	// example:
	//
	// PHYSICAL_TABLE, PHYSICAL_VIEW, PHYSICAL_MATERIALIZED_VIEW, DATASOURCE_TABLE, DATASOURCE_VIEW, DATASOURCE_MATERIALIZED_VIEW, DIM_NORMAL, DIM_LEVEL, DIM_ENUM, DIM_VIRTUAL, FACT_EVENT, FACT_PROCESS, FACT_SNAPSHOT, SUM_BIZ_UNIT
	MetadataSubType *string `json:"MetadataSubType,omitempty" xml:"MetadataSubType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TABLE
	MetadataType *string `json:"MetadataType,omitempty" xml:"MetadataType,omitempty"`
	// example:
	//
	// table1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// BY_GUID, BY_PROPERTY
	ReferenceType *string `json:"ReferenceType,omitempty" xml:"ReferenceType,omitempty"`
	// example:
	//
	// project1, bizUnit1
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
}

func (s AddRegisterLineageRequestAddRegisterLineageCommandTarget) String() string {
	return dara.Prettify(s)
}

func (s AddRegisterLineageRequestAddRegisterLineageCommandTarget) GoString() string {
	return s.String()
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandTarget) GetCatalog() *string {
	return s.Catalog
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandTarget) GetEnv() *string {
	return s.Env
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandTarget) GetExtProperties() map[string]interface{} {
	return s.ExtProperties
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandTarget) GetGuid() *string {
	return s.Guid
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandTarget) GetMetadataSubType() *string {
	return s.MetadataSubType
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandTarget) GetMetadataType() *string {
	return s.MetadataType
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandTarget) GetName() *string {
	return s.Name
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandTarget) GetReferenceType() *string {
	return s.ReferenceType
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandTarget) GetSchema() *string {
	return s.Schema
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandTarget) SetCatalog(v string) *AddRegisterLineageRequestAddRegisterLineageCommandTarget {
	s.Catalog = &v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandTarget) SetEnv(v string) *AddRegisterLineageRequestAddRegisterLineageCommandTarget {
	s.Env = &v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandTarget) SetExtProperties(v map[string]interface{}) *AddRegisterLineageRequestAddRegisterLineageCommandTarget {
	s.ExtProperties = v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandTarget) SetGuid(v string) *AddRegisterLineageRequestAddRegisterLineageCommandTarget {
	s.Guid = &v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandTarget) SetMetadataSubType(v string) *AddRegisterLineageRequestAddRegisterLineageCommandTarget {
	s.MetadataSubType = &v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandTarget) SetMetadataType(v string) *AddRegisterLineageRequestAddRegisterLineageCommandTarget {
	s.MetadataType = &v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandTarget) SetName(v string) *AddRegisterLineageRequestAddRegisterLineageCommandTarget {
	s.Name = &v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandTarget) SetReferenceType(v string) *AddRegisterLineageRequestAddRegisterLineageCommandTarget {
	s.ReferenceType = &v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandTarget) SetSchema(v string) *AddRegisterLineageRequestAddRegisterLineageCommandTarget {
	s.Schema = &v
	return s
}

func (s *AddRegisterLineageRequestAddRegisterLineageCommandTarget) Validate() error {
	return dara.Validate(s)
}
