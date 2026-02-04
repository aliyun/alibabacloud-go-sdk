// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNormalizationSchemaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CreateNormalizationSchemaRequest
	GetLang() *string
	SetNormalizationCategoryId(v string) *CreateNormalizationSchemaRequest
	GetNormalizationCategoryId() *string
	SetNormalizationFields(v []*CreateNormalizationSchemaRequestNormalizationFields) *CreateNormalizationSchemaRequest
	GetNormalizationFields() []*CreateNormalizationSchemaRequestNormalizationFields
	SetNormalizationSchemaDescription(v string) *CreateNormalizationSchemaRequest
	GetNormalizationSchemaDescription() *string
	SetNormalizationSchemaId(v string) *CreateNormalizationSchemaRequest
	GetNormalizationSchemaId() *string
	SetNormalizationSchemaName(v string) *CreateNormalizationSchemaRequest
	GetNormalizationSchemaName() *string
	SetNormalizationSchemaType(v string) *CreateNormalizationSchemaRequest
	GetNormalizationSchemaType() *string
	SetRegionId(v string) *CreateNormalizationSchemaRequest
	GetRegionId() *string
	SetRoleFor(v int64) *CreateNormalizationSchemaRequest
	GetRoleFor() *int64
	SetTargetLogStore(v string) *CreateNormalizationSchemaRequest
	GetTargetLogStore() *string
}

type CreateNormalizationSchemaRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// NETWORK_CATEGORY
	NormalizationCategoryId *string                                                `json:"NormalizationCategoryId,omitempty" xml:"NormalizationCategoryId,omitempty"`
	NormalizationFields     []*CreateNormalizationSchemaRequestNormalizationFields `json:"NormalizationFields,omitempty" xml:"NormalizationFields,omitempty" type:"Repeated"`
	// example:
	//
	// ProcessQuery
	NormalizationSchemaDescription *string `json:"NormalizationSchemaDescription,omitempty" xml:"NormalizationSchemaDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROCESS_QUERY_DNS_ACTIVITY
	NormalizationSchemaId *string `json:"NormalizationSchemaId,omitempty" xml:"NormalizationSchemaId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ProcessQuery
	NormalizationSchemaName *string `json:"NormalizationSchemaName,omitempty" xml:"NormalizationSchemaName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// log
	NormalizationSchemaType *string `json:"NormalizationSchemaType,omitempty" xml:"NormalizationSchemaType,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx-activity
	TargetLogStore *string `json:"TargetLogStore,omitempty" xml:"TargetLogStore,omitempty"`
}

func (s CreateNormalizationSchemaRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNormalizationSchemaRequest) GoString() string {
	return s.String()
}

func (s *CreateNormalizationSchemaRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateNormalizationSchemaRequest) GetNormalizationCategoryId() *string {
	return s.NormalizationCategoryId
}

func (s *CreateNormalizationSchemaRequest) GetNormalizationFields() []*CreateNormalizationSchemaRequestNormalizationFields {
	return s.NormalizationFields
}

func (s *CreateNormalizationSchemaRequest) GetNormalizationSchemaDescription() *string {
	return s.NormalizationSchemaDescription
}

func (s *CreateNormalizationSchemaRequest) GetNormalizationSchemaId() *string {
	return s.NormalizationSchemaId
}

func (s *CreateNormalizationSchemaRequest) GetNormalizationSchemaName() *string {
	return s.NormalizationSchemaName
}

func (s *CreateNormalizationSchemaRequest) GetNormalizationSchemaType() *string {
	return s.NormalizationSchemaType
}

func (s *CreateNormalizationSchemaRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateNormalizationSchemaRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *CreateNormalizationSchemaRequest) GetTargetLogStore() *string {
	return s.TargetLogStore
}

func (s *CreateNormalizationSchemaRequest) SetLang(v string) *CreateNormalizationSchemaRequest {
	s.Lang = &v
	return s
}

func (s *CreateNormalizationSchemaRequest) SetNormalizationCategoryId(v string) *CreateNormalizationSchemaRequest {
	s.NormalizationCategoryId = &v
	return s
}

func (s *CreateNormalizationSchemaRequest) SetNormalizationFields(v []*CreateNormalizationSchemaRequestNormalizationFields) *CreateNormalizationSchemaRequest {
	s.NormalizationFields = v
	return s
}

func (s *CreateNormalizationSchemaRequest) SetNormalizationSchemaDescription(v string) *CreateNormalizationSchemaRequest {
	s.NormalizationSchemaDescription = &v
	return s
}

func (s *CreateNormalizationSchemaRequest) SetNormalizationSchemaId(v string) *CreateNormalizationSchemaRequest {
	s.NormalizationSchemaId = &v
	return s
}

func (s *CreateNormalizationSchemaRequest) SetNormalizationSchemaName(v string) *CreateNormalizationSchemaRequest {
	s.NormalizationSchemaName = &v
	return s
}

func (s *CreateNormalizationSchemaRequest) SetNormalizationSchemaType(v string) *CreateNormalizationSchemaRequest {
	s.NormalizationSchemaType = &v
	return s
}

func (s *CreateNormalizationSchemaRequest) SetRegionId(v string) *CreateNormalizationSchemaRequest {
	s.RegionId = &v
	return s
}

func (s *CreateNormalizationSchemaRequest) SetRoleFor(v int64) *CreateNormalizationSchemaRequest {
	s.RoleFor = &v
	return s
}

func (s *CreateNormalizationSchemaRequest) SetTargetLogStore(v string) *CreateNormalizationSchemaRequest {
	s.TargetLogStore = &v
	return s
}

func (s *CreateNormalizationSchemaRequest) Validate() error {
	if s.NormalizationFields != nil {
		for _, item := range s.NormalizationFields {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateNormalizationSchemaRequestNormalizationFields struct {
	// example:
	//
	// cloud_user
	NormalizationFieldDescription *string `json:"NormalizationFieldDescription,omitempty" xml:"NormalizationFieldDescription,omitempty"`
	// example:
	//
	// 173326*******
	NormalizationFieldExample *string `json:"NormalizationFieldExample,omitempty" xml:"NormalizationFieldExample,omitempty"`
	// example:
	//
	// preset
	NormalizationFieldFrom *string `json:"NormalizationFieldFrom,omitempty" xml:"NormalizationFieldFrom,omitempty"`
	// example:
	//
	// true
	NormalizationFieldJsonIndexAll *bool                                                                            `json:"NormalizationFieldJsonIndexAll,omitempty" xml:"NormalizationFieldJsonIndexAll,omitempty"`
	NormalizationFieldJsonKeys     []*CreateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys `json:"NormalizationFieldJsonKeys,omitempty" xml:"NormalizationFieldJsonKeys,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// cloud_user
	NormalizationFieldName *string `json:"NormalizationFieldName,omitempty" xml:"NormalizationFieldName,omitempty"`
	// example:
	//
	// true
	NormalizationFieldRequired *bool `json:"NormalizationFieldRequired,omitempty" xml:"NormalizationFieldRequired,omitempty"`
	// example:
	//
	// true
	NormalizationFieldRequirement *bool `json:"NormalizationFieldRequirement,omitempty" xml:"NormalizationFieldRequirement,omitempty"`
	// example:
	//
	// true
	NormalizationFieldReserved *bool `json:"NormalizationFieldReserved,omitempty" xml:"NormalizationFieldReserved,omitempty"`
	// example:
	//
	// true
	NormalizationFieldTokenize *bool `json:"NormalizationFieldTokenize,omitempty" xml:"NormalizationFieldTokenize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// varchar
	NormalizationFieldType *string `json:"NormalizationFieldType,omitempty" xml:"NormalizationFieldType,omitempty"`
}

func (s CreateNormalizationSchemaRequestNormalizationFields) String() string {
	return dara.Prettify(s)
}

func (s CreateNormalizationSchemaRequestNormalizationFields) GoString() string {
	return s.String()
}

func (s *CreateNormalizationSchemaRequestNormalizationFields) GetNormalizationFieldDescription() *string {
	return s.NormalizationFieldDescription
}

func (s *CreateNormalizationSchemaRequestNormalizationFields) GetNormalizationFieldExample() *string {
	return s.NormalizationFieldExample
}

func (s *CreateNormalizationSchemaRequestNormalizationFields) GetNormalizationFieldFrom() *string {
	return s.NormalizationFieldFrom
}

func (s *CreateNormalizationSchemaRequestNormalizationFields) GetNormalizationFieldJsonIndexAll() *bool {
	return s.NormalizationFieldJsonIndexAll
}

func (s *CreateNormalizationSchemaRequestNormalizationFields) GetNormalizationFieldJsonKeys() []*CreateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys {
	return s.NormalizationFieldJsonKeys
}

func (s *CreateNormalizationSchemaRequestNormalizationFields) GetNormalizationFieldName() *string {
	return s.NormalizationFieldName
}

func (s *CreateNormalizationSchemaRequestNormalizationFields) GetNormalizationFieldRequired() *bool {
	return s.NormalizationFieldRequired
}

func (s *CreateNormalizationSchemaRequestNormalizationFields) GetNormalizationFieldRequirement() *bool {
	return s.NormalizationFieldRequirement
}

func (s *CreateNormalizationSchemaRequestNormalizationFields) GetNormalizationFieldReserved() *bool {
	return s.NormalizationFieldReserved
}

func (s *CreateNormalizationSchemaRequestNormalizationFields) GetNormalizationFieldTokenize() *bool {
	return s.NormalizationFieldTokenize
}

func (s *CreateNormalizationSchemaRequestNormalizationFields) GetNormalizationFieldType() *string {
	return s.NormalizationFieldType
}

func (s *CreateNormalizationSchemaRequestNormalizationFields) SetNormalizationFieldDescription(v string) *CreateNormalizationSchemaRequestNormalizationFields {
	s.NormalizationFieldDescription = &v
	return s
}

func (s *CreateNormalizationSchemaRequestNormalizationFields) SetNormalizationFieldExample(v string) *CreateNormalizationSchemaRequestNormalizationFields {
	s.NormalizationFieldExample = &v
	return s
}

func (s *CreateNormalizationSchemaRequestNormalizationFields) SetNormalizationFieldFrom(v string) *CreateNormalizationSchemaRequestNormalizationFields {
	s.NormalizationFieldFrom = &v
	return s
}

func (s *CreateNormalizationSchemaRequestNormalizationFields) SetNormalizationFieldJsonIndexAll(v bool) *CreateNormalizationSchemaRequestNormalizationFields {
	s.NormalizationFieldJsonIndexAll = &v
	return s
}

func (s *CreateNormalizationSchemaRequestNormalizationFields) SetNormalizationFieldJsonKeys(v []*CreateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys) *CreateNormalizationSchemaRequestNormalizationFields {
	s.NormalizationFieldJsonKeys = v
	return s
}

func (s *CreateNormalizationSchemaRequestNormalizationFields) SetNormalizationFieldName(v string) *CreateNormalizationSchemaRequestNormalizationFields {
	s.NormalizationFieldName = &v
	return s
}

func (s *CreateNormalizationSchemaRequestNormalizationFields) SetNormalizationFieldRequired(v bool) *CreateNormalizationSchemaRequestNormalizationFields {
	s.NormalizationFieldRequired = &v
	return s
}

func (s *CreateNormalizationSchemaRequestNormalizationFields) SetNormalizationFieldRequirement(v bool) *CreateNormalizationSchemaRequestNormalizationFields {
	s.NormalizationFieldRequirement = &v
	return s
}

func (s *CreateNormalizationSchemaRequestNormalizationFields) SetNormalizationFieldReserved(v bool) *CreateNormalizationSchemaRequestNormalizationFields {
	s.NormalizationFieldReserved = &v
	return s
}

func (s *CreateNormalizationSchemaRequestNormalizationFields) SetNormalizationFieldTokenize(v bool) *CreateNormalizationSchemaRequestNormalizationFields {
	s.NormalizationFieldTokenize = &v
	return s
}

func (s *CreateNormalizationSchemaRequestNormalizationFields) SetNormalizationFieldType(v string) *CreateNormalizationSchemaRequestNormalizationFields {
	s.NormalizationFieldType = &v
	return s
}

func (s *CreateNormalizationSchemaRequestNormalizationFields) Validate() error {
	if s.NormalizationFieldJsonKeys != nil {
		for _, item := range s.NormalizationFieldJsonKeys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys struct {
	NormalizationFieldDescription *string `json:"NormalizationFieldDescription,omitempty" xml:"NormalizationFieldDescription,omitempty"`
	NormalizationFieldExample     *string `json:"NormalizationFieldExample,omitempty" xml:"NormalizationFieldExample,omitempty"`
	// example:
	//
	// preset
	NormalizationFieldFrom *string `json:"NormalizationFieldFrom,omitempty" xml:"NormalizationFieldFrom,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// alert_name
	NormalizationFieldName *string `json:"NormalizationFieldName,omitempty" xml:"NormalizationFieldName,omitempty"`
	// example:
	//
	// true
	NormalizationFieldRequired *bool `json:"NormalizationFieldRequired,omitempty" xml:"NormalizationFieldRequired,omitempty"`
	// example:
	//
	// true
	NormalizationFieldTokenize *bool `json:"NormalizationFieldTokenize,omitempty" xml:"NormalizationFieldTokenize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// text
	NormalizationFieldType *string `json:"NormalizationFieldType,omitempty" xml:"NormalizationFieldType,omitempty"`
}

func (s CreateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys) String() string {
	return dara.Prettify(s)
}

func (s CreateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys) GoString() string {
	return s.String()
}

func (s *CreateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys) GetNormalizationFieldDescription() *string {
	return s.NormalizationFieldDescription
}

func (s *CreateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys) GetNormalizationFieldExample() *string {
	return s.NormalizationFieldExample
}

func (s *CreateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys) GetNormalizationFieldFrom() *string {
	return s.NormalizationFieldFrom
}

func (s *CreateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys) GetNormalizationFieldName() *string {
	return s.NormalizationFieldName
}

func (s *CreateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys) GetNormalizationFieldRequired() *bool {
	return s.NormalizationFieldRequired
}

func (s *CreateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys) GetNormalizationFieldTokenize() *bool {
	return s.NormalizationFieldTokenize
}

func (s *CreateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys) GetNormalizationFieldType() *string {
	return s.NormalizationFieldType
}

func (s *CreateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys) SetNormalizationFieldDescription(v string) *CreateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys {
	s.NormalizationFieldDescription = &v
	return s
}

func (s *CreateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys) SetNormalizationFieldExample(v string) *CreateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys {
	s.NormalizationFieldExample = &v
	return s
}

func (s *CreateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys) SetNormalizationFieldFrom(v string) *CreateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys {
	s.NormalizationFieldFrom = &v
	return s
}

func (s *CreateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys) SetNormalizationFieldName(v string) *CreateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys {
	s.NormalizationFieldName = &v
	return s
}

func (s *CreateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys) SetNormalizationFieldRequired(v bool) *CreateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys {
	s.NormalizationFieldRequired = &v
	return s
}

func (s *CreateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys) SetNormalizationFieldTokenize(v bool) *CreateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys {
	s.NormalizationFieldTokenize = &v
	return s
}

func (s *CreateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys) SetNormalizationFieldType(v string) *CreateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys {
	s.NormalizationFieldType = &v
	return s
}

func (s *CreateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys) Validate() error {
	return dara.Validate(s)
}
