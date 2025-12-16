// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApp interface {
	dara.Model
	String() string
	GoString() string
	SetAutoSwitch(v bool) *App
	GetAutoSwitch() *bool
	SetCluster(v *AppCluster) *App
	GetCluster() *AppCluster
	SetDataSources(v []*DataSource) *App
	GetDataSources() []*DataSource
	SetDescription(v string) *App
	GetDescription() *string
	SetDomain(v *Domain) *App
	GetDomain() *Domain
	SetFetchFields(v []*string) *App
	GetFetchFields() []*string
	SetFirstRanks(v []*FirstRank) *App
	GetFirstRanks() []*FirstRank
	SetNetworkType(v string) *App
	GetNetworkType() *string
	SetQueryProcessors(v []*QueryProcessor) *App
	GetQueryProcessors() []*QueryProcessor
	SetQuota(v *Quota) *App
	GetQuota() *Quota
	SetRealtimeShared(v bool) *App
	GetRealtimeShared() *bool
	SetSchema(v *Schema) *App
	GetSchema() *Schema
	SetSchemas(v []*Schema) *App
	GetSchemas() []*Schema
	SetSecondRanks(v []*SecondRank) *App
	GetSecondRanks() []*SecondRank
	SetSummaries(v []*Summary) *App
	GetSummaries() []*Summary
	SetType(v string) *App
	GetType() *string
}

type App struct {
	AutoSwitch      *bool             `json:"autoSwitch,omitempty" xml:"autoSwitch,omitempty"`
	Cluster         *AppCluster       `json:"cluster,omitempty" xml:"cluster,omitempty" type:"Struct"`
	DataSources     []*DataSource     `json:"dataSources,omitempty" xml:"dataSources,omitempty" type:"Repeated"`
	Description     *string           `json:"description,omitempty" xml:"description,omitempty"`
	Domain          *Domain           `json:"domain,omitempty" xml:"domain,omitempty"`
	FetchFields     []*string         `json:"fetchFields,omitempty" xml:"fetchFields,omitempty" type:"Repeated"`
	FirstRanks      []*FirstRank      `json:"firstRanks,omitempty" xml:"firstRanks,omitempty" type:"Repeated"`
	NetworkType     *string           `json:"networkType,omitempty" xml:"networkType,omitempty"`
	QueryProcessors []*QueryProcessor `json:"queryProcessors,omitempty" xml:"queryProcessors,omitempty" type:"Repeated"`
	Quota           *Quota            `json:"quota,omitempty" xml:"quota,omitempty"`
	RealtimeShared  *bool             `json:"realtimeShared,omitempty" xml:"realtimeShared,omitempty"`
	Schema          *Schema           `json:"schema,omitempty" xml:"schema,omitempty"`
	Schemas         []*Schema         `json:"schemas,omitempty" xml:"schemas,omitempty" type:"Repeated"`
	SecondRanks     []*SecondRank     `json:"secondRanks,omitempty" xml:"secondRanks,omitempty" type:"Repeated"`
	Summaries       []*Summary        `json:"summaries,omitempty" xml:"summaries,omitempty" type:"Repeated"`
	Type            *string           `json:"type,omitempty" xml:"type,omitempty"`
}

func (s App) String() string {
	return dara.Prettify(s)
}

func (s App) GoString() string {
	return s.String()
}

func (s *App) GetAutoSwitch() *bool {
	return s.AutoSwitch
}

func (s *App) GetCluster() *AppCluster {
	return s.Cluster
}

func (s *App) GetDataSources() []*DataSource {
	return s.DataSources
}

func (s *App) GetDescription() *string {
	return s.Description
}

func (s *App) GetDomain() *Domain {
	return s.Domain
}

func (s *App) GetFetchFields() []*string {
	return s.FetchFields
}

func (s *App) GetFirstRanks() []*FirstRank {
	return s.FirstRanks
}

func (s *App) GetNetworkType() *string {
	return s.NetworkType
}

func (s *App) GetQueryProcessors() []*QueryProcessor {
	return s.QueryProcessors
}

func (s *App) GetQuota() *Quota {
	return s.Quota
}

func (s *App) GetRealtimeShared() *bool {
	return s.RealtimeShared
}

func (s *App) GetSchema() *Schema {
	return s.Schema
}

func (s *App) GetSchemas() []*Schema {
	return s.Schemas
}

func (s *App) GetSecondRanks() []*SecondRank {
	return s.SecondRanks
}

func (s *App) GetSummaries() []*Summary {
	return s.Summaries
}

func (s *App) GetType() *string {
	return s.Type
}

func (s *App) SetAutoSwitch(v bool) *App {
	s.AutoSwitch = &v
	return s
}

func (s *App) SetCluster(v *AppCluster) *App {
	s.Cluster = v
	return s
}

func (s *App) SetDataSources(v []*DataSource) *App {
	s.DataSources = v
	return s
}

func (s *App) SetDescription(v string) *App {
	s.Description = &v
	return s
}

func (s *App) SetDomain(v *Domain) *App {
	s.Domain = v
	return s
}

func (s *App) SetFetchFields(v []*string) *App {
	s.FetchFields = v
	return s
}

func (s *App) SetFirstRanks(v []*FirstRank) *App {
	s.FirstRanks = v
	return s
}

func (s *App) SetNetworkType(v string) *App {
	s.NetworkType = &v
	return s
}

func (s *App) SetQueryProcessors(v []*QueryProcessor) *App {
	s.QueryProcessors = v
	return s
}

func (s *App) SetQuota(v *Quota) *App {
	s.Quota = v
	return s
}

func (s *App) SetRealtimeShared(v bool) *App {
	s.RealtimeShared = &v
	return s
}

func (s *App) SetSchema(v *Schema) *App {
	s.Schema = v
	return s
}

func (s *App) SetSchemas(v []*Schema) *App {
	s.Schemas = v
	return s
}

func (s *App) SetSecondRanks(v []*SecondRank) *App {
	s.SecondRanks = v
	return s
}

func (s *App) SetSummaries(v []*Summary) *App {
	s.Summaries = v
	return s
}

func (s *App) SetType(v string) *App {
	s.Type = &v
	return s
}

func (s *App) Validate() error {
	if s.Cluster != nil {
		if err := s.Cluster.Validate(); err != nil {
			return err
		}
	}
	if s.DataSources != nil {
		for _, item := range s.DataSources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Domain != nil {
		if err := s.Domain.Validate(); err != nil {
			return err
		}
	}
	if s.FirstRanks != nil {
		for _, item := range s.FirstRanks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.QueryProcessors != nil {
		for _, item := range s.QueryProcessors {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	if s.Schema != nil {
		if err := s.Schema.Validate(); err != nil {
			return err
		}
	}
	if s.Schemas != nil {
		for _, item := range s.Schemas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SecondRanks != nil {
		for _, item := range s.SecondRanks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Summaries != nil {
		for _, item := range s.Summaries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AppCluster struct {
	MaxQueryClauseLength *int32 `json:"maxQueryClauseLength,omitempty" xml:"maxQueryClauseLength,omitempty"`
	MaxTimeoutMS         *int32 `json:"maxTimeoutMS,omitempty" xml:"maxTimeoutMS,omitempty"`
}

func (s AppCluster) String() string {
	return dara.Prettify(s)
}

func (s AppCluster) GoString() string {
	return s.String()
}

func (s *AppCluster) GetMaxQueryClauseLength() *int32 {
	return s.MaxQueryClauseLength
}

func (s *AppCluster) GetMaxTimeoutMS() *int32 {
	return s.MaxTimeoutMS
}

func (s *AppCluster) SetMaxQueryClauseLength(v int32) *AppCluster {
	s.MaxQueryClauseLength = &v
	return s
}

func (s *AppCluster) SetMaxTimeoutMS(v int32) *AppCluster {
	s.MaxTimeoutMS = &v
	return s
}

func (s *AppCluster) Validate() error {
	return dara.Validate(s)
}
