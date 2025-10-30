// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRegisterLineageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteRegisterLineageCommand(v *DeleteRegisterLineageRequestDeleteRegisterLineageCommand) *DeleteRegisterLineageRequest
	GetDeleteRegisterLineageCommand() *DeleteRegisterLineageRequestDeleteRegisterLineageCommand
	SetOpTenantId(v int64) *DeleteRegisterLineageRequest
	GetOpTenantId() *int64
}

type DeleteRegisterLineageRequest struct {
	// This parameter is required.
	DeleteRegisterLineageCommand *DeleteRegisterLineageRequestDeleteRegisterLineageCommand `json:"DeleteRegisterLineageCommand,omitempty" xml:"DeleteRegisterLineageCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s DeleteRegisterLineageRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRegisterLineageRequest) GoString() string {
	return s.String()
}

func (s *DeleteRegisterLineageRequest) GetDeleteRegisterLineageCommand() *DeleteRegisterLineageRequestDeleteRegisterLineageCommand {
	return s.DeleteRegisterLineageCommand
}

func (s *DeleteRegisterLineageRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteRegisterLineageRequest) SetDeleteRegisterLineageCommand(v *DeleteRegisterLineageRequestDeleteRegisterLineageCommand) *DeleteRegisterLineageRequest {
	s.DeleteRegisterLineageCommand = v
	return s
}

func (s *DeleteRegisterLineageRequest) SetOpTenantId(v int64) *DeleteRegisterLineageRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteRegisterLineageRequest) Validate() error {
	if s.DeleteRegisterLineageCommand != nil {
		if err := s.DeleteRegisterLineageCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteRegisterLineageRequestDeleteRegisterLineageCommand struct {
	CascadeDeleteLineage *bool                                                                       `json:"CascadeDeleteLineage,omitempty" xml:"CascadeDeleteLineage,omitempty"`
	DetailedLineages     []*DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineages `json:"DetailedLineages,omitempty" xml:"DetailedLineages,omitempty" type:"Repeated"`
	// This parameter is required.
	Source *DeleteRegisterLineageRequestDeleteRegisterLineageCommandSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	// This parameter is required.
	Target *DeleteRegisterLineageRequestDeleteRegisterLineageCommandTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	// example:
	//
	// 300001234
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// 300004567
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteRegisterLineageRequestDeleteRegisterLineageCommand) String() string {
	return dara.Prettify(s)
}

func (s DeleteRegisterLineageRequestDeleteRegisterLineageCommand) GoString() string {
	return s.String()
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommand) GetCascadeDeleteLineage() *bool {
	return s.CascadeDeleteLineage
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommand) GetDetailedLineages() []*DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineages {
	return s.DetailedLineages
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommand) GetSource() *DeleteRegisterLineageRequestDeleteRegisterLineageCommandSource {
	return s.Source
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommand) GetTarget() *DeleteRegisterLineageRequestDeleteRegisterLineageCommandTarget {
	return s.Target
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommand) GetTenantId() *int64 {
	return s.TenantId
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommand) GetUserId() *string {
	return s.UserId
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommand) SetCascadeDeleteLineage(v bool) *DeleteRegisterLineageRequestDeleteRegisterLineageCommand {
	s.CascadeDeleteLineage = &v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommand) SetDetailedLineages(v []*DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineages) *DeleteRegisterLineageRequestDeleteRegisterLineageCommand {
	s.DetailedLineages = v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommand) SetSource(v *DeleteRegisterLineageRequestDeleteRegisterLineageCommandSource) *DeleteRegisterLineageRequestDeleteRegisterLineageCommand {
	s.Source = v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommand) SetTarget(v *DeleteRegisterLineageRequestDeleteRegisterLineageCommandTarget) *DeleteRegisterLineageRequestDeleteRegisterLineageCommand {
	s.Target = v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommand) SetTenantId(v int64) *DeleteRegisterLineageRequestDeleteRegisterLineageCommand {
	s.TenantId = &v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommand) SetUserId(v string) *DeleteRegisterLineageRequestDeleteRegisterLineageCommand {
	s.UserId = &v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommand) Validate() error {
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

type DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineages struct {
	IsDirect *bool `json:"IsDirect,omitempty" xml:"IsDirect,omitempty"`
	// This parameter is required.
	Source *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	// This parameter is required.
	Target *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
}

func (s DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineages) String() string {
	return dara.Prettify(s)
}

func (s DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineages) GoString() string {
	return s.String()
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineages) GetIsDirect() *bool {
	return s.IsDirect
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineages) GetSource() *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesSource {
	return s.Source
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineages) GetTarget() *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesTarget {
	return s.Target
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineages) SetIsDirect(v bool) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineages {
	s.IsDirect = &v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineages) SetSource(v *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesSource) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineages {
	s.Source = v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineages) SetTarget(v *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesTarget) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineages {
	s.Target = v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineages) Validate() error {
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

type DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesSource struct {
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

func (s DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesSource) String() string {
	return dara.Prettify(s)
}

func (s DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesSource) GoString() string {
	return s.String()
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesSource) GetCatalog() *string {
	return s.Catalog
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesSource) GetEnv() *string {
	return s.Env
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesSource) GetExtProperties() map[string]interface{} {
	return s.ExtProperties
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesSource) GetGuid() *string {
	return s.Guid
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesSource) GetMetadataType() *string {
	return s.MetadataType
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesSource) GetName() *string {
	return s.Name
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesSource) GetParentGuid() *string {
	return s.ParentGuid
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesSource) GetReferenceType() *string {
	return s.ReferenceType
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesSource) GetSchema() *string {
	return s.Schema
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesSource) SetCatalog(v string) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesSource {
	s.Catalog = &v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesSource) SetEnv(v string) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesSource {
	s.Env = &v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesSource) SetExtProperties(v map[string]interface{}) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesSource {
	s.ExtProperties = v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesSource) SetGuid(v string) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesSource {
	s.Guid = &v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesSource) SetMetadataType(v string) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesSource {
	s.MetadataType = &v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesSource) SetName(v string) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesSource {
	s.Name = &v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesSource) SetParentGuid(v string) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesSource {
	s.ParentGuid = &v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesSource) SetReferenceType(v string) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesSource {
	s.ReferenceType = &v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesSource) SetSchema(v string) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesSource {
	s.Schema = &v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesSource) Validate() error {
	return dara.Validate(s)
}

type DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesTarget struct {
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

func (s DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesTarget) String() string {
	return dara.Prettify(s)
}

func (s DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesTarget) GoString() string {
	return s.String()
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesTarget) GetCatalog() *string {
	return s.Catalog
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesTarget) GetEnv() *string {
	return s.Env
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesTarget) GetExtProperties() map[string]interface{} {
	return s.ExtProperties
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesTarget) GetGuid() *string {
	return s.Guid
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesTarget) GetMetadataType() *string {
	return s.MetadataType
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesTarget) GetName() *string {
	return s.Name
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesTarget) GetParentGuid() *string {
	return s.ParentGuid
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesTarget) GetReferenceType() *string {
	return s.ReferenceType
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesTarget) GetSchema() *string {
	return s.Schema
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesTarget) SetCatalog(v string) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesTarget {
	s.Catalog = &v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesTarget) SetEnv(v string) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesTarget {
	s.Env = &v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesTarget) SetExtProperties(v map[string]interface{}) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesTarget {
	s.ExtProperties = v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesTarget) SetGuid(v string) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesTarget {
	s.Guid = &v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesTarget) SetMetadataType(v string) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesTarget {
	s.MetadataType = &v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesTarget) SetName(v string) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesTarget {
	s.Name = &v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesTarget) SetParentGuid(v string) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesTarget {
	s.ParentGuid = &v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesTarget) SetReferenceType(v string) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesTarget {
	s.ReferenceType = &v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesTarget) SetSchema(v string) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesTarget {
	s.Schema = &v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandDetailedLineagesTarget) Validate() error {
	return dara.Validate(s)
}

type DeleteRegisterLineageRequestDeleteRegisterLineageCommandSource struct {
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

func (s DeleteRegisterLineageRequestDeleteRegisterLineageCommandSource) String() string {
	return dara.Prettify(s)
}

func (s DeleteRegisterLineageRequestDeleteRegisterLineageCommandSource) GoString() string {
	return s.String()
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandSource) GetCatalog() *string {
	return s.Catalog
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandSource) GetEnv() *string {
	return s.Env
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandSource) GetExtProperties() map[string]interface{} {
	return s.ExtProperties
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandSource) GetGuid() *string {
	return s.Guid
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandSource) GetMetadataSubType() *string {
	return s.MetadataSubType
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandSource) GetMetadataType() *string {
	return s.MetadataType
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandSource) GetName() *string {
	return s.Name
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandSource) GetReferenceType() *string {
	return s.ReferenceType
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandSource) GetSchema() *string {
	return s.Schema
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandSource) SetCatalog(v string) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandSource {
	s.Catalog = &v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandSource) SetEnv(v string) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandSource {
	s.Env = &v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandSource) SetExtProperties(v map[string]interface{}) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandSource {
	s.ExtProperties = v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandSource) SetGuid(v string) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandSource {
	s.Guid = &v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandSource) SetMetadataSubType(v string) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandSource {
	s.MetadataSubType = &v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandSource) SetMetadataType(v string) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandSource {
	s.MetadataType = &v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandSource) SetName(v string) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandSource {
	s.Name = &v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandSource) SetReferenceType(v string) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandSource {
	s.ReferenceType = &v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandSource) SetSchema(v string) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandSource {
	s.Schema = &v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandSource) Validate() error {
	return dara.Validate(s)
}

type DeleteRegisterLineageRequestDeleteRegisterLineageCommandTarget struct {
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
	// PHYSICAL_TABLE, PHYSICAL_VIEW, PHYSICAL_MATERIALIZED_VIEW, DATASOURCE_TABLE, DATASOURCE_VIEW, DATASOURCE_MATERIALIZED_VIEW, DIM_NORMAL, DIM_LEVEL, DIM_ENUM, DIM_VIRTUAL, FACT_EVENT, FACT_PROCESS, FACT_SNAPSHOT,SUM_BIZ_UNIT
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

func (s DeleteRegisterLineageRequestDeleteRegisterLineageCommandTarget) String() string {
	return dara.Prettify(s)
}

func (s DeleteRegisterLineageRequestDeleteRegisterLineageCommandTarget) GoString() string {
	return s.String()
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandTarget) GetCatalog() *string {
	return s.Catalog
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandTarget) GetEnv() *string {
	return s.Env
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandTarget) GetExtProperties() map[string]interface{} {
	return s.ExtProperties
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandTarget) GetGuid() *string {
	return s.Guid
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandTarget) GetMetadataSubType() *string {
	return s.MetadataSubType
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandTarget) GetMetadataType() *string {
	return s.MetadataType
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandTarget) GetName() *string {
	return s.Name
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandTarget) GetReferenceType() *string {
	return s.ReferenceType
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandTarget) GetSchema() *string {
	return s.Schema
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandTarget) SetCatalog(v string) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandTarget {
	s.Catalog = &v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandTarget) SetEnv(v string) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandTarget {
	s.Env = &v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandTarget) SetExtProperties(v map[string]interface{}) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandTarget {
	s.ExtProperties = v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandTarget) SetGuid(v string) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandTarget {
	s.Guid = &v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandTarget) SetMetadataSubType(v string) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandTarget {
	s.MetadataSubType = &v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandTarget) SetMetadataType(v string) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandTarget {
	s.MetadataType = &v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandTarget) SetName(v string) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandTarget {
	s.Name = &v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandTarget) SetReferenceType(v string) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandTarget {
	s.ReferenceType = &v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandTarget) SetSchema(v string) *DeleteRegisterLineageRequestDeleteRegisterLineageCommandTarget {
	s.Schema = &v
	return s
}

func (s *DeleteRegisterLineageRequestDeleteRegisterLineageCommandTarget) Validate() error {
	return dara.Validate(s)
}
