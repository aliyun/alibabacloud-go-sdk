// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCrawlerType interface {
	dara.Model
	String() string
	GoString() string
	SetDisplayName(v string) *CrawlerType
	GetDisplayName() *string
	SetSupportedEntityTypes(v []*CrawlerTypeSupportedEntityTypes) *CrawlerType
	GetSupportedEntityTypes() []*CrawlerTypeSupportedEntityTypes
	SetType(v string) *CrawlerType
	GetType() *string
}

type CrawlerType struct {
	// The display name of the metadata crawler.
	//
	// example:
	//
	// Data Lake Formation
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The supported entity types. The entity types are sorted based on the declaration order.
	SupportedEntityTypes []*CrawlerTypeSupportedEntityTypes `json:"SupportedEntityTypes,omitempty" xml:"SupportedEntityTypes,omitempty" type:"Repeated"`
	// The identifier of the metadata crawler type. Valid values:
	//
	// 	- `maxcompute`
	//
	// 	- `dlf`
	//
	// 	- `hms`: This type of crawler can be used to collect metadata from E-MapReduce (EMR) and CDH Hive clusters.
	//
	// 	- `holo`
	//
	// 	- `mysql`
	//
	// 	- `oracle`
	//
	// 	- `postgresql`
	//
	// 	- `sqlserver`
	//
	// 	- `analyticdb_for_mysql`
	//
	// 	- `ads`
	//
	// 	- `hybriddb_for_postgresql`
	//
	// 	- `ots`
	//
	// 	- `clickhouse`
	//
	// 	- `starrocks`: This type of crawler can be used to query metadata entities only in internal catalogs.
	//
	// example:
	//
	// dlf
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CrawlerType) String() string {
	return dara.Prettify(s)
}

func (s CrawlerType) GoString() string {
	return s.String()
}

func (s *CrawlerType) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CrawlerType) GetSupportedEntityTypes() []*CrawlerTypeSupportedEntityTypes {
	return s.SupportedEntityTypes
}

func (s *CrawlerType) GetType() *string {
	return s.Type
}

func (s *CrawlerType) SetDisplayName(v string) *CrawlerType {
	s.DisplayName = &v
	return s
}

func (s *CrawlerType) SetSupportedEntityTypes(v []*CrawlerTypeSupportedEntityTypes) *CrawlerType {
	s.SupportedEntityTypes = v
	return s
}

func (s *CrawlerType) SetType(v string) *CrawlerType {
	s.Type = &v
	return s
}

func (s *CrawlerType) Validate() error {
	if s.SupportedEntityTypes != nil {
		for _, item := range s.SupportedEntityTypes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CrawlerTypeSupportedEntityTypes struct {
	// Specifies whether the entity type is optional.
	//
	// For example, whether the schema level of the MaxCompute crawler type is optional depends on whether the three-layer model is enabled for a MaxCompute project.
	//
	// example:
	//
	// true
	Optional *bool `json:"Optional,omitempty" xml:"Optional,omitempty"`
	// The subtype of the parent entity. If the subtype does not exist, null is returned.
	//
	// example:
	//
	// database
	ParentSubType *string `json:"ParentSubType,omitempty" xml:"ParentSubType,omitempty"`
	// The identifier of the entity subtype. Valid values: `catalog, database, schema, table, and column`.
	//
	// example:
	//
	// table
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	// The identifier of the entity type. The value of this parameter varies based on the type of the metadata crawler. Configure this parameter in the `${Crawler type}-${Subtype}` format.
	//
	// example:
	//
	// dlf-table
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CrawlerTypeSupportedEntityTypes) String() string {
	return dara.Prettify(s)
}

func (s CrawlerTypeSupportedEntityTypes) GoString() string {
	return s.String()
}

func (s *CrawlerTypeSupportedEntityTypes) GetOptional() *bool {
	return s.Optional
}

func (s *CrawlerTypeSupportedEntityTypes) GetParentSubType() *string {
	return s.ParentSubType
}

func (s *CrawlerTypeSupportedEntityTypes) GetSubType() *string {
	return s.SubType
}

func (s *CrawlerTypeSupportedEntityTypes) GetType() *string {
	return s.Type
}

func (s *CrawlerTypeSupportedEntityTypes) SetOptional(v bool) *CrawlerTypeSupportedEntityTypes {
	s.Optional = &v
	return s
}

func (s *CrawlerTypeSupportedEntityTypes) SetParentSubType(v string) *CrawlerTypeSupportedEntityTypes {
	s.ParentSubType = &v
	return s
}

func (s *CrawlerTypeSupportedEntityTypes) SetSubType(v string) *CrawlerTypeSupportedEntityTypes {
	s.SubType = &v
	return s
}

func (s *CrawlerTypeSupportedEntityTypes) SetType(v string) *CrawlerTypeSupportedEntityTypes {
	s.Type = &v
	return s
}

func (s *CrawlerTypeSupportedEntityTypes) Validate() error {
	return dara.Validate(s)
}
