// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCrawlerTypeCapabilitiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCrawlerTypes(v []*GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes) *GetCrawlerTypeCapabilitiesResponseBody
	GetCrawlerTypes() []*GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes
	SetRequestId(v string) *GetCrawlerTypeCapabilitiesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCrawlerTypeCapabilitiesResponseBody
	GetSuccess() *bool
}

type GetCrawlerTypeCapabilitiesResponseBody struct {
	CrawlerTypes []*GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes `json:"CrawlerTypes,omitempty" xml:"CrawlerTypes,omitempty" type:"Repeated"`
	// example:
	//
	// 9252F32F-D855-549E-8898-61CF5A733050
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCrawlerTypeCapabilitiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCrawlerTypeCapabilitiesResponseBody) GoString() string {
	return s.String()
}

func (s *GetCrawlerTypeCapabilitiesResponseBody) GetCrawlerTypes() []*GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes {
	return s.CrawlerTypes
}

func (s *GetCrawlerTypeCapabilitiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCrawlerTypeCapabilitiesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCrawlerTypeCapabilitiesResponseBody) SetCrawlerTypes(v []*GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes) *GetCrawlerTypeCapabilitiesResponseBody {
	s.CrawlerTypes = v
	return s
}

func (s *GetCrawlerTypeCapabilitiesResponseBody) SetRequestId(v string) *GetCrawlerTypeCapabilitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCrawlerTypeCapabilitiesResponseBody) SetSuccess(v bool) *GetCrawlerTypeCapabilitiesResponseBody {
	s.Success = &v
	return s
}

func (s *GetCrawlerTypeCapabilitiesResponseBody) Validate() error {
	if s.CrawlerTypes != nil {
		for _, item := range s.CrawlerTypes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes struct {
	// example:
	//
	// DATABASE
	DefaultScopeUnit *string `json:"DefaultScopeUnit,omitempty" xml:"DefaultScopeUnit,omitempty"`
	// example:
	//
	// Hologres
	DisplayName              *string                                                                   `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	RequireResourceGroup     *bool                                                                     `json:"RequireResourceGroup,omitempty" xml:"RequireResourceGroup,omitempty"`
	SupportAiComment         *bool                                                                     `json:"SupportAiComment,omitempty" xml:"SupportAiComment,omitempty"`
	SupportExcludeRegex      *bool                                                                     `json:"SupportExcludeRegex,omitempty" xml:"SupportExcludeRegex,omitempty"`
	SupportSchedule          *bool                                                                     `json:"SupportSchedule,omitempty" xml:"SupportSchedule,omitempty"`
	SupportedDatasourceTypes []*string                                                                 `json:"SupportedDatasourceTypes,omitempty" xml:"SupportedDatasourceTypes,omitempty" type:"Repeated"`
	SupportedEntityTypes     []*GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedEntityTypes `json:"SupportedEntityTypes,omitempty" xml:"SupportedEntityTypes,omitempty" type:"Repeated"`
	SupportedOptionKeys      []*GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedOptionKeys  `json:"SupportedOptionKeys,omitempty" xml:"SupportedOptionKeys,omitempty" type:"Repeated"`
	SupportedScopeUnits      []*string                                                                 `json:"SupportedScopeUnits,omitempty" xml:"SupportedScopeUnits,omitempty" type:"Repeated"`
	// example:
	//
	// holo
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes) String() string {
	return dara.Prettify(s)
}

func (s GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes) GoString() string {
	return s.String()
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes) GetDefaultScopeUnit() *string {
	return s.DefaultScopeUnit
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes) GetRequireResourceGroup() *bool {
	return s.RequireResourceGroup
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes) GetSupportAiComment() *bool {
	return s.SupportAiComment
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes) GetSupportExcludeRegex() *bool {
	return s.SupportExcludeRegex
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes) GetSupportSchedule() *bool {
	return s.SupportSchedule
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes) GetSupportedDatasourceTypes() []*string {
	return s.SupportedDatasourceTypes
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes) GetSupportedEntityTypes() []*GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedEntityTypes {
	return s.SupportedEntityTypes
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes) GetSupportedOptionKeys() []*GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedOptionKeys {
	return s.SupportedOptionKeys
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes) GetSupportedScopeUnits() []*string {
	return s.SupportedScopeUnits
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes) GetType() *string {
	return s.Type
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes) SetDefaultScopeUnit(v string) *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes {
	s.DefaultScopeUnit = &v
	return s
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes) SetDisplayName(v string) *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes {
	s.DisplayName = &v
	return s
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes) SetRequireResourceGroup(v bool) *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes {
	s.RequireResourceGroup = &v
	return s
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes) SetSupportAiComment(v bool) *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes {
	s.SupportAiComment = &v
	return s
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes) SetSupportExcludeRegex(v bool) *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes {
	s.SupportExcludeRegex = &v
	return s
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes) SetSupportSchedule(v bool) *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes {
	s.SupportSchedule = &v
	return s
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes) SetSupportedDatasourceTypes(v []*string) *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes {
	s.SupportedDatasourceTypes = v
	return s
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes) SetSupportedEntityTypes(v []*GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedEntityTypes) *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes {
	s.SupportedEntityTypes = v
	return s
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes) SetSupportedOptionKeys(v []*GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedOptionKeys) *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes {
	s.SupportedOptionKeys = v
	return s
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes) SetSupportedScopeUnits(v []*string) *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes {
	s.SupportedScopeUnits = v
	return s
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes) SetType(v string) *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes {
	s.Type = &v
	return s
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypes) Validate() error {
	if s.SupportedEntityTypes != nil {
		for _, item := range s.SupportedEntityTypes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SupportedOptionKeys != nil {
		for _, item := range s.SupportedOptionKeys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedEntityTypes struct {
	Optional *bool `json:"Optional,omitempty" xml:"Optional,omitempty"`
	// example:
	//
	// instance
	ParentSubType *string `json:"ParentSubType,omitempty" xml:"ParentSubType,omitempty"`
	// example:
	//
	// database
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	// example:
	//
	// holo
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedEntityTypes) String() string {
	return dara.Prettify(s)
}

func (s GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedEntityTypes) GoString() string {
	return s.String()
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedEntityTypes) GetOptional() *bool {
	return s.Optional
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedEntityTypes) GetParentSubType() *string {
	return s.ParentSubType
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedEntityTypes) GetSubType() *string {
	return s.SubType
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedEntityTypes) GetType() *string {
	return s.Type
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedEntityTypes) SetOptional(v bool) *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedEntityTypes {
	s.Optional = &v
	return s
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedEntityTypes) SetParentSubType(v string) *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedEntityTypes {
	s.ParentSubType = &v
	return s
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedEntityTypes) SetSubType(v string) *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedEntityTypes {
	s.SubType = &v
	return s
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedEntityTypes) SetType(v string) *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedEntityTypes {
	s.Type = &v
	return s
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedEntityTypes) Validate() error {
	return dara.Validate(s)
}

type GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedOptionKeys struct {
	AllowedValues []*string `json:"AllowedValues,omitempty" xml:"AllowedValues,omitempty" type:"Repeated"`
	// example:
	//
	// false
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// example:
	//
	// CollectRecyclebin
	Key      *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Required *bool   `json:"Required,omitempty" xml:"Required,omitempty"`
	// example:
	//
	// BOOLEAN
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedOptionKeys) String() string {
	return dara.Prettify(s)
}

func (s GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedOptionKeys) GoString() string {
	return s.String()
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedOptionKeys) GetAllowedValues() []*string {
	return s.AllowedValues
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedOptionKeys) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedOptionKeys) GetKey() *string {
	return s.Key
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedOptionKeys) GetRequired() *bool {
	return s.Required
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedOptionKeys) GetValueType() *string {
	return s.ValueType
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedOptionKeys) SetAllowedValues(v []*string) *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedOptionKeys {
	s.AllowedValues = v
	return s
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedOptionKeys) SetDefaultValue(v string) *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedOptionKeys {
	s.DefaultValue = &v
	return s
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedOptionKeys) SetKey(v string) *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedOptionKeys {
	s.Key = &v
	return s
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedOptionKeys) SetRequired(v bool) *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedOptionKeys {
	s.Required = &v
	return s
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedOptionKeys) SetValueType(v string) *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedOptionKeys {
	s.ValueType = &v
	return s
}

func (s *GetCrawlerTypeCapabilitiesResponseBodyCrawlerTypesSupportedOptionKeys) Validate() error {
	return dara.Validate(s)
}
