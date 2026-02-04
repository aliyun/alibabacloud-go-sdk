// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNormalizationSchemaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *UpdateNormalizationSchemaRequest
	GetLang() *string
	SetNormalizationFields(v []*UpdateNormalizationSchemaRequestNormalizationFields) *UpdateNormalizationSchemaRequest
	GetNormalizationFields() []*UpdateNormalizationSchemaRequestNormalizationFields
	SetNormalizationSchemaDescription(v string) *UpdateNormalizationSchemaRequest
	GetNormalizationSchemaDescription() *string
	SetNormalizationSchemaId(v string) *UpdateNormalizationSchemaRequest
	GetNormalizationSchemaId() *string
	SetNormalizationSchemaName(v string) *UpdateNormalizationSchemaRequest
	GetNormalizationSchemaName() *string
	SetNormalizationSchemaType(v string) *UpdateNormalizationSchemaRequest
	GetNormalizationSchemaType() *string
	SetRegionId(v string) *UpdateNormalizationSchemaRequest
	GetRegionId() *string
	SetRoleFor(v int64) *UpdateNormalizationSchemaRequest
	GetRoleFor() *int64
}

type UpdateNormalizationSchemaRequest struct {
	// example:
	//
	// zh
	Lang                *string                                                `json:"Lang,omitempty" xml:"Lang,omitempty"`
	NormalizationFields []*UpdateNormalizationSchemaRequestNormalizationFields `json:"NormalizationFields,omitempty" xml:"NormalizationFields,omitempty" type:"Repeated"`
	// example:
	//
	// ProcessQuery
	NormalizationSchemaDescription *string `json:"NormalizationSchemaDescription,omitempty" xml:"NormalizationSchemaDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// WAF_ALERT_ACTIVITY
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
}

func (s UpdateNormalizationSchemaRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNormalizationSchemaRequest) GoString() string {
	return s.String()
}

func (s *UpdateNormalizationSchemaRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateNormalizationSchemaRequest) GetNormalizationFields() []*UpdateNormalizationSchemaRequestNormalizationFields {
	return s.NormalizationFields
}

func (s *UpdateNormalizationSchemaRequest) GetNormalizationSchemaDescription() *string {
	return s.NormalizationSchemaDescription
}

func (s *UpdateNormalizationSchemaRequest) GetNormalizationSchemaId() *string {
	return s.NormalizationSchemaId
}

func (s *UpdateNormalizationSchemaRequest) GetNormalizationSchemaName() *string {
	return s.NormalizationSchemaName
}

func (s *UpdateNormalizationSchemaRequest) GetNormalizationSchemaType() *string {
	return s.NormalizationSchemaType
}

func (s *UpdateNormalizationSchemaRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateNormalizationSchemaRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *UpdateNormalizationSchemaRequest) SetLang(v string) *UpdateNormalizationSchemaRequest {
	s.Lang = &v
	return s
}

func (s *UpdateNormalizationSchemaRequest) SetNormalizationFields(v []*UpdateNormalizationSchemaRequestNormalizationFields) *UpdateNormalizationSchemaRequest {
	s.NormalizationFields = v
	return s
}

func (s *UpdateNormalizationSchemaRequest) SetNormalizationSchemaDescription(v string) *UpdateNormalizationSchemaRequest {
	s.NormalizationSchemaDescription = &v
	return s
}

func (s *UpdateNormalizationSchemaRequest) SetNormalizationSchemaId(v string) *UpdateNormalizationSchemaRequest {
	s.NormalizationSchemaId = &v
	return s
}

func (s *UpdateNormalizationSchemaRequest) SetNormalizationSchemaName(v string) *UpdateNormalizationSchemaRequest {
	s.NormalizationSchemaName = &v
	return s
}

func (s *UpdateNormalizationSchemaRequest) SetNormalizationSchemaType(v string) *UpdateNormalizationSchemaRequest {
	s.NormalizationSchemaType = &v
	return s
}

func (s *UpdateNormalizationSchemaRequest) SetRegionId(v string) *UpdateNormalizationSchemaRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateNormalizationSchemaRequest) SetRoleFor(v int64) *UpdateNormalizationSchemaRequest {
	s.RoleFor = &v
	return s
}

func (s *UpdateNormalizationSchemaRequest) Validate() error {
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

type UpdateNormalizationSchemaRequestNormalizationFields struct {
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
	NormalizationFieldFrom         *string                                                                          `json:"NormalizationFieldFrom,omitempty" xml:"NormalizationFieldFrom,omitempty"`
	NormalizationFieldJsonIndexAll *bool                                                                            `json:"NormalizationFieldJsonIndexAll,omitempty" xml:"NormalizationFieldJsonIndexAll,omitempty"`
	NormalizationFieldJsonKeys     []*UpdateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys `json:"NormalizationFieldJsonKeys,omitempty" xml:"NormalizationFieldJsonKeys,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// cloud_user
	NormalizationFieldName     *string `json:"NormalizationFieldName,omitempty" xml:"NormalizationFieldName,omitempty"`
	NormalizationFieldRequired *bool   `json:"NormalizationFieldRequired,omitempty" xml:"NormalizationFieldRequired,omitempty"`
	NormalizationFieldTokenize *bool   `json:"NormalizationFieldTokenize,omitempty" xml:"NormalizationFieldTokenize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// varchar
	NormalizationFieldType *string `json:"NormalizationFieldType,omitempty" xml:"NormalizationFieldType,omitempty"`
}

func (s UpdateNormalizationSchemaRequestNormalizationFields) String() string {
	return dara.Prettify(s)
}

func (s UpdateNormalizationSchemaRequestNormalizationFields) GoString() string {
	return s.String()
}

func (s *UpdateNormalizationSchemaRequestNormalizationFields) GetNormalizationFieldDescription() *string {
	return s.NormalizationFieldDescription
}

func (s *UpdateNormalizationSchemaRequestNormalizationFields) GetNormalizationFieldExample() *string {
	return s.NormalizationFieldExample
}

func (s *UpdateNormalizationSchemaRequestNormalizationFields) GetNormalizationFieldFrom() *string {
	return s.NormalizationFieldFrom
}

func (s *UpdateNormalizationSchemaRequestNormalizationFields) GetNormalizationFieldJsonIndexAll() *bool {
	return s.NormalizationFieldJsonIndexAll
}

func (s *UpdateNormalizationSchemaRequestNormalizationFields) GetNormalizationFieldJsonKeys() []*UpdateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys {
	return s.NormalizationFieldJsonKeys
}

func (s *UpdateNormalizationSchemaRequestNormalizationFields) GetNormalizationFieldName() *string {
	return s.NormalizationFieldName
}

func (s *UpdateNormalizationSchemaRequestNormalizationFields) GetNormalizationFieldRequired() *bool {
	return s.NormalizationFieldRequired
}

func (s *UpdateNormalizationSchemaRequestNormalizationFields) GetNormalizationFieldTokenize() *bool {
	return s.NormalizationFieldTokenize
}

func (s *UpdateNormalizationSchemaRequestNormalizationFields) GetNormalizationFieldType() *string {
	return s.NormalizationFieldType
}

func (s *UpdateNormalizationSchemaRequestNormalizationFields) SetNormalizationFieldDescription(v string) *UpdateNormalizationSchemaRequestNormalizationFields {
	s.NormalizationFieldDescription = &v
	return s
}

func (s *UpdateNormalizationSchemaRequestNormalizationFields) SetNormalizationFieldExample(v string) *UpdateNormalizationSchemaRequestNormalizationFields {
	s.NormalizationFieldExample = &v
	return s
}

func (s *UpdateNormalizationSchemaRequestNormalizationFields) SetNormalizationFieldFrom(v string) *UpdateNormalizationSchemaRequestNormalizationFields {
	s.NormalizationFieldFrom = &v
	return s
}

func (s *UpdateNormalizationSchemaRequestNormalizationFields) SetNormalizationFieldJsonIndexAll(v bool) *UpdateNormalizationSchemaRequestNormalizationFields {
	s.NormalizationFieldJsonIndexAll = &v
	return s
}

func (s *UpdateNormalizationSchemaRequestNormalizationFields) SetNormalizationFieldJsonKeys(v []*UpdateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys) *UpdateNormalizationSchemaRequestNormalizationFields {
	s.NormalizationFieldJsonKeys = v
	return s
}

func (s *UpdateNormalizationSchemaRequestNormalizationFields) SetNormalizationFieldName(v string) *UpdateNormalizationSchemaRequestNormalizationFields {
	s.NormalizationFieldName = &v
	return s
}

func (s *UpdateNormalizationSchemaRequestNormalizationFields) SetNormalizationFieldRequired(v bool) *UpdateNormalizationSchemaRequestNormalizationFields {
	s.NormalizationFieldRequired = &v
	return s
}

func (s *UpdateNormalizationSchemaRequestNormalizationFields) SetNormalizationFieldTokenize(v bool) *UpdateNormalizationSchemaRequestNormalizationFields {
	s.NormalizationFieldTokenize = &v
	return s
}

func (s *UpdateNormalizationSchemaRequestNormalizationFields) SetNormalizationFieldType(v string) *UpdateNormalizationSchemaRequestNormalizationFields {
	s.NormalizationFieldType = &v
	return s
}

func (s *UpdateNormalizationSchemaRequestNormalizationFields) Validate() error {
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

type UpdateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys struct {
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
	NormalizationFieldName     *string `json:"NormalizationFieldName,omitempty" xml:"NormalizationFieldName,omitempty"`
	NormalizationFieldRequired *bool   `json:"NormalizationFieldRequired,omitempty" xml:"NormalizationFieldRequired,omitempty"`
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

func (s UpdateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys) String() string {
	return dara.Prettify(s)
}

func (s UpdateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys) GoString() string {
	return s.String()
}

func (s *UpdateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys) GetNormalizationFieldDescription() *string {
	return s.NormalizationFieldDescription
}

func (s *UpdateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys) GetNormalizationFieldExample() *string {
	return s.NormalizationFieldExample
}

func (s *UpdateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys) GetNormalizationFieldFrom() *string {
	return s.NormalizationFieldFrom
}

func (s *UpdateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys) GetNormalizationFieldName() *string {
	return s.NormalizationFieldName
}

func (s *UpdateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys) GetNormalizationFieldRequired() *bool {
	return s.NormalizationFieldRequired
}

func (s *UpdateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys) GetNormalizationFieldTokenize() *bool {
	return s.NormalizationFieldTokenize
}

func (s *UpdateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys) GetNormalizationFieldType() *string {
	return s.NormalizationFieldType
}

func (s *UpdateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys) SetNormalizationFieldDescription(v string) *UpdateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys {
	s.NormalizationFieldDescription = &v
	return s
}

func (s *UpdateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys) SetNormalizationFieldExample(v string) *UpdateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys {
	s.NormalizationFieldExample = &v
	return s
}

func (s *UpdateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys) SetNormalizationFieldFrom(v string) *UpdateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys {
	s.NormalizationFieldFrom = &v
	return s
}

func (s *UpdateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys) SetNormalizationFieldName(v string) *UpdateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys {
	s.NormalizationFieldName = &v
	return s
}

func (s *UpdateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys) SetNormalizationFieldRequired(v bool) *UpdateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys {
	s.NormalizationFieldRequired = &v
	return s
}

func (s *UpdateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys) SetNormalizationFieldTokenize(v bool) *UpdateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys {
	s.NormalizationFieldTokenize = &v
	return s
}

func (s *UpdateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys) SetNormalizationFieldType(v string) *UpdateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys {
	s.NormalizationFieldType = &v
	return s
}

func (s *UpdateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys) Validate() error {
	return dara.Validate(s)
}
