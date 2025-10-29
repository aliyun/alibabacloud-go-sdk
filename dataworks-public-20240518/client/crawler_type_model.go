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
	// example:
	//
	// Data Lake Formation
	DisplayName          *string                            `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	SupportedEntityTypes []*CrawlerTypeSupportedEntityTypes `json:"SupportedEntityTypes,omitempty" xml:"SupportedEntityTypes,omitempty" type:"Repeated"`
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
	// example:
	//
	// 如对于maxcompute-schema类型，schema层级是否存在可选（是否开启三层模型）
	Optional *bool `json:"Optional,omitempty" xml:"Optional,omitempty"`
	// example:
	//
	// database
	ParentSubType *string `json:"ParentSubType,omitempty" xml:"ParentSubType,omitempty"`
	// example:
	//
	// table
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
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
