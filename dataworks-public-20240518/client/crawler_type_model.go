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
	// The display name.
	//
	// example:
	//
	// Data Lake Formation
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The list of supported entity types. The entity types have a top-down hierarchical relationship based on their declaration order.
	SupportedEntityTypes []*CrawlerTypeSupportedEntityTypes `json:"SupportedEntityTypes,omitempty" xml:"SupportedEntityTypes,omitempty" type:"Repeated"`
	// The type identifier.
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
	// Indicates whether the entity type is optional.
	//
	// example:
	//
	// For example, for the maxcompute-schema type, whether the schema level is optional (whether the three-layer model is enabled)
	Optional *bool `json:"Optional,omitempty" xml:"Optional,omitempty"`
	// The entity subtype of the parent level. The value is null if no parent level exists.
	//
	// example:
	//
	// database
	ParentSubType *string `json:"ParentSubType,omitempty" xml:"ParentSubType,omitempty"`
	// The entity subtype identifier.
	//
	// example:
	//
	// table
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	// The entity type identifier, which is related to the crawler type. The format is (CrawlerType)-{SubType}.
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
