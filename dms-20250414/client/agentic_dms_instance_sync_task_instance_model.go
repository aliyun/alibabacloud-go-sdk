// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgenticDmsInstanceSyncTaskInstance interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogUuid(v string) *AgenticDmsInstanceSyncTaskInstance
	GetCatalogUuid() *string
	SetCrawlerTaskId(v string) *AgenticDmsInstanceSyncTaskInstance
	GetCrawlerTaskId() *string
	SetDatasourceUuid(v string) *AgenticDmsInstanceSyncTaskInstance
	GetDatasourceUuid() *string
	SetDbType(v string) *AgenticDmsInstanceSyncTaskInstance
	GetDbType() *string
	SetDmsInstanceId(v string) *AgenticDmsInstanceSyncTaskInstance
	GetDmsInstanceId() *string
	SetDmsInstanceSummary(v *AgenticDmsInstanceSyncTaskInstanceDmsInstanceSummary) *AgenticDmsInstanceSyncTaskInstance
	GetDmsInstanceSummary() *AgenticDmsInstanceSyncTaskInstanceDmsInstanceSummary
	SetDmsRegionId(v string) *AgenticDmsInstanceSyncTaskInstance
	GetDmsRegionId() *string
	SetErrorCode(v string) *AgenticDmsInstanceSyncTaskInstance
	GetErrorCode() *string
	SetErrorSummary(v string) *AgenticDmsInstanceSyncTaskInstance
	GetErrorSummary() *string
	SetGmtCreate(v string) *AgenticDmsInstanceSyncTaskInstance
	GetGmtCreate() *string
	SetGmtModified(v string) *AgenticDmsInstanceSyncTaskInstance
	GetGmtModified() *string
	SetPhase(v string) *AgenticDmsInstanceSyncTaskInstance
	GetPhase() *string
	SetStatus(v string) *AgenticDmsInstanceSyncTaskInstance
	GetStatus() *string
}

type AgenticDmsInstanceSyncTaskInstance struct {
	CatalogUuid        *string                                               `json:"CatalogUuid,omitempty" xml:"CatalogUuid,omitempty"`
	CrawlerTaskId      *string                                               `json:"CrawlerTaskId,omitempty" xml:"CrawlerTaskId,omitempty"`
	DatasourceUuid     *string                                               `json:"DatasourceUuid,omitempty" xml:"DatasourceUuid,omitempty"`
	DbType             *string                                               `json:"DbType,omitempty" xml:"DbType,omitempty"`
	DmsInstanceId      *string                                               `json:"DmsInstanceId,omitempty" xml:"DmsInstanceId,omitempty"`
	DmsInstanceSummary *AgenticDmsInstanceSyncTaskInstanceDmsInstanceSummary `json:"DmsInstanceSummary,omitempty" xml:"DmsInstanceSummary,omitempty" type:"Struct"`
	DmsRegionId        *string                                               `json:"DmsRegionId,omitempty" xml:"DmsRegionId,omitempty"`
	ErrorCode          *string                                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorSummary       *string                                               `json:"ErrorSummary,omitempty" xml:"ErrorSummary,omitempty"`
	GmtCreate          *string                                               `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified        *string                                               `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Phase              *string                                               `json:"Phase,omitempty" xml:"Phase,omitempty"`
	Status             *string                                               `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s AgenticDmsInstanceSyncTaskInstance) String() string {
	return dara.Prettify(s)
}

func (s AgenticDmsInstanceSyncTaskInstance) GoString() string {
	return s.String()
}

func (s *AgenticDmsInstanceSyncTaskInstance) GetCatalogUuid() *string {
	return s.CatalogUuid
}

func (s *AgenticDmsInstanceSyncTaskInstance) GetCrawlerTaskId() *string {
	return s.CrawlerTaskId
}

func (s *AgenticDmsInstanceSyncTaskInstance) GetDatasourceUuid() *string {
	return s.DatasourceUuid
}

func (s *AgenticDmsInstanceSyncTaskInstance) GetDbType() *string {
	return s.DbType
}

func (s *AgenticDmsInstanceSyncTaskInstance) GetDmsInstanceId() *string {
	return s.DmsInstanceId
}

func (s *AgenticDmsInstanceSyncTaskInstance) GetDmsInstanceSummary() *AgenticDmsInstanceSyncTaskInstanceDmsInstanceSummary {
	return s.DmsInstanceSummary
}

func (s *AgenticDmsInstanceSyncTaskInstance) GetDmsRegionId() *string {
	return s.DmsRegionId
}

func (s *AgenticDmsInstanceSyncTaskInstance) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AgenticDmsInstanceSyncTaskInstance) GetErrorSummary() *string {
	return s.ErrorSummary
}

func (s *AgenticDmsInstanceSyncTaskInstance) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *AgenticDmsInstanceSyncTaskInstance) GetGmtModified() *string {
	return s.GmtModified
}

func (s *AgenticDmsInstanceSyncTaskInstance) GetPhase() *string {
	return s.Phase
}

func (s *AgenticDmsInstanceSyncTaskInstance) GetStatus() *string {
	return s.Status
}

func (s *AgenticDmsInstanceSyncTaskInstance) SetCatalogUuid(v string) *AgenticDmsInstanceSyncTaskInstance {
	s.CatalogUuid = &v
	return s
}

func (s *AgenticDmsInstanceSyncTaskInstance) SetCrawlerTaskId(v string) *AgenticDmsInstanceSyncTaskInstance {
	s.CrawlerTaskId = &v
	return s
}

func (s *AgenticDmsInstanceSyncTaskInstance) SetDatasourceUuid(v string) *AgenticDmsInstanceSyncTaskInstance {
	s.DatasourceUuid = &v
	return s
}

func (s *AgenticDmsInstanceSyncTaskInstance) SetDbType(v string) *AgenticDmsInstanceSyncTaskInstance {
	s.DbType = &v
	return s
}

func (s *AgenticDmsInstanceSyncTaskInstance) SetDmsInstanceId(v string) *AgenticDmsInstanceSyncTaskInstance {
	s.DmsInstanceId = &v
	return s
}

func (s *AgenticDmsInstanceSyncTaskInstance) SetDmsInstanceSummary(v *AgenticDmsInstanceSyncTaskInstanceDmsInstanceSummary) *AgenticDmsInstanceSyncTaskInstance {
	s.DmsInstanceSummary = v
	return s
}

func (s *AgenticDmsInstanceSyncTaskInstance) SetDmsRegionId(v string) *AgenticDmsInstanceSyncTaskInstance {
	s.DmsRegionId = &v
	return s
}

func (s *AgenticDmsInstanceSyncTaskInstance) SetErrorCode(v string) *AgenticDmsInstanceSyncTaskInstance {
	s.ErrorCode = &v
	return s
}

func (s *AgenticDmsInstanceSyncTaskInstance) SetErrorSummary(v string) *AgenticDmsInstanceSyncTaskInstance {
	s.ErrorSummary = &v
	return s
}

func (s *AgenticDmsInstanceSyncTaskInstance) SetGmtCreate(v string) *AgenticDmsInstanceSyncTaskInstance {
	s.GmtCreate = &v
	return s
}

func (s *AgenticDmsInstanceSyncTaskInstance) SetGmtModified(v string) *AgenticDmsInstanceSyncTaskInstance {
	s.GmtModified = &v
	return s
}

func (s *AgenticDmsInstanceSyncTaskInstance) SetPhase(v string) *AgenticDmsInstanceSyncTaskInstance {
	s.Phase = &v
	return s
}

func (s *AgenticDmsInstanceSyncTaskInstance) SetStatus(v string) *AgenticDmsInstanceSyncTaskInstance {
	s.Status = &v
	return s
}

func (s *AgenticDmsInstanceSyncTaskInstance) Validate() error {
	if s.DmsInstanceSummary != nil {
		if err := s.DmsInstanceSummary.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AgenticDmsInstanceSyncTaskInstanceDmsInstanceSummary struct {
	Alias              *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	DbType             *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	EnvType            *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	Host               *string `json:"Host,omitempty" xml:"Host,omitempty"`
	InstanceResourceId *string `json:"InstanceResourceId,omitempty" xml:"InstanceResourceId,omitempty"`
	InstanceSource     *string `json:"InstanceSource,omitempty" xml:"InstanceSource,omitempty"`
	Port               *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AgenticDmsInstanceSyncTaskInstanceDmsInstanceSummary) String() string {
	return dara.Prettify(s)
}

func (s AgenticDmsInstanceSyncTaskInstanceDmsInstanceSummary) GoString() string {
	return s.String()
}

func (s *AgenticDmsInstanceSyncTaskInstanceDmsInstanceSummary) GetAlias() *string {
	return s.Alias
}

func (s *AgenticDmsInstanceSyncTaskInstanceDmsInstanceSummary) GetDbType() *string {
	return s.DbType
}

func (s *AgenticDmsInstanceSyncTaskInstanceDmsInstanceSummary) GetEnvType() *string {
	return s.EnvType
}

func (s *AgenticDmsInstanceSyncTaskInstanceDmsInstanceSummary) GetHost() *string {
	return s.Host
}

func (s *AgenticDmsInstanceSyncTaskInstanceDmsInstanceSummary) GetInstanceResourceId() *string {
	return s.InstanceResourceId
}

func (s *AgenticDmsInstanceSyncTaskInstanceDmsInstanceSummary) GetInstanceSource() *string {
	return s.InstanceSource
}

func (s *AgenticDmsInstanceSyncTaskInstanceDmsInstanceSummary) GetPort() *int32 {
	return s.Port
}

func (s *AgenticDmsInstanceSyncTaskInstanceDmsInstanceSummary) GetRegionId() *string {
	return s.RegionId
}

func (s *AgenticDmsInstanceSyncTaskInstanceDmsInstanceSummary) SetAlias(v string) *AgenticDmsInstanceSyncTaskInstanceDmsInstanceSummary {
	s.Alias = &v
	return s
}

func (s *AgenticDmsInstanceSyncTaskInstanceDmsInstanceSummary) SetDbType(v string) *AgenticDmsInstanceSyncTaskInstanceDmsInstanceSummary {
	s.DbType = &v
	return s
}

func (s *AgenticDmsInstanceSyncTaskInstanceDmsInstanceSummary) SetEnvType(v string) *AgenticDmsInstanceSyncTaskInstanceDmsInstanceSummary {
	s.EnvType = &v
	return s
}

func (s *AgenticDmsInstanceSyncTaskInstanceDmsInstanceSummary) SetHost(v string) *AgenticDmsInstanceSyncTaskInstanceDmsInstanceSummary {
	s.Host = &v
	return s
}

func (s *AgenticDmsInstanceSyncTaskInstanceDmsInstanceSummary) SetInstanceResourceId(v string) *AgenticDmsInstanceSyncTaskInstanceDmsInstanceSummary {
	s.InstanceResourceId = &v
	return s
}

func (s *AgenticDmsInstanceSyncTaskInstanceDmsInstanceSummary) SetInstanceSource(v string) *AgenticDmsInstanceSyncTaskInstanceDmsInstanceSummary {
	s.InstanceSource = &v
	return s
}

func (s *AgenticDmsInstanceSyncTaskInstanceDmsInstanceSummary) SetPort(v int32) *AgenticDmsInstanceSyncTaskInstanceDmsInstanceSummary {
	s.Port = &v
	return s
}

func (s *AgenticDmsInstanceSyncTaskInstanceDmsInstanceSummary) SetRegionId(v string) *AgenticDmsInstanceSyncTaskInstanceDmsInstanceSummary {
	s.RegionId = &v
	return s
}

func (s *AgenticDmsInstanceSyncTaskInstanceDmsInstanceSummary) Validate() error {
	return dara.Validate(s)
}
