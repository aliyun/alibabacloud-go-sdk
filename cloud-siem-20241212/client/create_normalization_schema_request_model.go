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
	SetNormalizationFieldSource(v string) *CreateNormalizationSchemaRequest
	GetNormalizationFieldSource() *string
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
	SetNormalizationSecurityDomainId(v string) *CreateNormalizationSchemaRequest
	GetNormalizationSecurityDomainId() *string
	SetProductId(v string) *CreateNormalizationSchemaRequest
	GetProductId() *string
	SetRegionId(v string) *CreateNormalizationSchemaRequest
	GetRegionId() *string
	SetRoleFor(v int64) *CreateNormalizationSchemaRequest
	GetRoleFor() *int64
	SetTargetLogStore(v string) *CreateNormalizationSchemaRequest
	GetTargetLogStore() *string
	SetVendorId(v string) *CreateNormalizationSchemaRequest
	GetVendorId() *string
}

type CreateNormalizationSchemaRequest struct {
	// The language of the response messages. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// Note: This parameter is required in actual API calls. If it is not provided, the service returns a 400 error: "The specified Lang can not be empty." Pass the language code explicitly.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The normalization category ID.
	//
	// example:
	//
	// NETWORK_CATEGORY
	NormalizationCategoryId *string `json:"NormalizationCategoryId,omitempty" xml:"NormalizationCategoryId,omitempty"`
	// The field source.
	//
	// example:
	//
	// native
	NormalizationFieldSource *string `json:"NormalizationFieldSource,omitempty" xml:"NormalizationFieldSource,omitempty"`
	// The normalization fields.
	NormalizationFields []*CreateNormalizationSchemaRequestNormalizationFields `json:"NormalizationFields,omitempty" xml:"NormalizationFields,omitempty" type:"Repeated"`
	// The normalization schema description.
	//
	// example:
	//
	// ProcessQuery
	NormalizationSchemaDescription *string `json:"NormalizationSchemaDescription,omitempty" xml:"NormalizationSchemaDescription,omitempty"`
	// The normalization schema ID. When creating a custom schema, the ID must start with custom- and can contain only lowercase letters, digits, and hyphens. Example: custom-my-log-schema.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROCESS_QUERY_DNS_ACTIVITY
	NormalizationSchemaId *string `json:"NormalizationSchemaId,omitempty" xml:"NormalizationSchemaId,omitempty"`
	// The normalization schema name.
	//
	// This parameter is required.
	//
	// example:
	//
	// ProcessQuery
	NormalizationSchemaName *string `json:"NormalizationSchemaName,omitempty" xml:"NormalizationSchemaName,omitempty"`
	// The normalization schema type. Valid values:
	//
	// - log: log.
	//
	// - entity: entity.
	//
	// This parameter is required.
	//
	// example:
	//
	// log
	NormalizationSchemaType *string `json:"NormalizationSchemaType,omitempty" xml:"NormalizationSchemaType,omitempty"`
	// The security domain ID.
	//
	// example:
	//
	// NETWORK_AND_WEB_SECURITY
	NormalizationSecurityDomainId *string `json:"NormalizationSecurityDomainId,omitempty" xml:"NormalizationSecurityDomainId,omitempty"`
	// The product ID.
	//
	// example:
	//
	// sas
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The region where the data management center of Cloud Threat Detection and Response (CTDR) is located. Select the management center region based on the region where your assets reside. Valid values:
	//
	// - cn-hangzhou: assets in the Chinese mainland or Hong Kong (China).
	//
	// - ap-southeast-1: assets in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID that the administrator uses to switch to the view of another member account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The Simple Log Service LogStore.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx-activity
	TargetLogStore *string `json:"TargetLogStore,omitempty" xml:"TargetLogStore,omitempty"`
	// The vendor ID.
	//
	// example:
	//
	// alibaba
	VendorId *string `json:"VendorId,omitempty" xml:"VendorId,omitempty"`
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

func (s *CreateNormalizationSchemaRequest) GetNormalizationFieldSource() *string {
	return s.NormalizationFieldSource
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

func (s *CreateNormalizationSchemaRequest) GetNormalizationSecurityDomainId() *string {
	return s.NormalizationSecurityDomainId
}

func (s *CreateNormalizationSchemaRequest) GetProductId() *string {
	return s.ProductId
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

func (s *CreateNormalizationSchemaRequest) GetVendorId() *string {
	return s.VendorId
}

func (s *CreateNormalizationSchemaRequest) SetLang(v string) *CreateNormalizationSchemaRequest {
	s.Lang = &v
	return s
}

func (s *CreateNormalizationSchemaRequest) SetNormalizationCategoryId(v string) *CreateNormalizationSchemaRequest {
	s.NormalizationCategoryId = &v
	return s
}

func (s *CreateNormalizationSchemaRequest) SetNormalizationFieldSource(v string) *CreateNormalizationSchemaRequest {
	s.NormalizationFieldSource = &v
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

func (s *CreateNormalizationSchemaRequest) SetNormalizationSecurityDomainId(v string) *CreateNormalizationSchemaRequest {
	s.NormalizationSecurityDomainId = &v
	return s
}

func (s *CreateNormalizationSchemaRequest) SetProductId(v string) *CreateNormalizationSchemaRequest {
	s.ProductId = &v
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

func (s *CreateNormalizationSchemaRequest) SetVendorId(v string) *CreateNormalizationSchemaRequest {
	s.VendorId = &v
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
	// The standard field description.
	//
	// example:
	//
	// cloud_user
	NormalizationFieldDescription *string `json:"NormalizationFieldDescription,omitempty" xml:"NormalizationFieldDescription,omitempty"`
	// The standard field example.
	//
	// example:
	//
	// 173326*******
	NormalizationFieldExample *string `json:"NormalizationFieldExample,omitempty" xml:"NormalizationFieldExample,omitempty"`
	// The key source for a standard field of the JSON type.
	//
	// example:
	//
	// preset
	NormalizationFieldFrom *string `json:"NormalizationFieldFrom,omitempty" xml:"NormalizationFieldFrom,omitempty"`
	// Specifies whether to create indexes for all keys in a standard field of the JSON type.
	//
	// example:
	//
	// true
	NormalizationFieldJsonIndexAll *bool `json:"NormalizationFieldJsonIndexAll,omitempty" xml:"NormalizationFieldJsonIndexAll,omitempty"`
	// The key list for a standard field of the JSON type.
	NormalizationFieldJsonKeys []*CreateNormalizationSchemaRequestNormalizationFieldsNormalizationFieldJsonKeys `json:"NormalizationFieldJsonKeys,omitempty" xml:"NormalizationFieldJsonKeys,omitempty" type:"Repeated"`
	// The normalization field name.
	//
	// This parameter is required.
	//
	// example:
	//
	// cloud_user
	NormalizationFieldName *string `json:"NormalizationFieldName,omitempty" xml:"NormalizationFieldName,omitempty"`
	// Specifies whether the standard field is required.
	//
	// example:
	//
	// true
	NormalizationFieldRequired *bool `json:"NormalizationFieldRequired,omitempty" xml:"NormalizationFieldRequired,omitempty"`
	// Specifies whether the normalization field is required.
	//
	// example:
	//
	// true
	NormalizationFieldRequirement *bool `json:"NormalizationFieldRequirement,omitempty" xml:"NormalizationFieldRequirement,omitempty"`
	// Specifies whether the normalization field is reserved.
	//
	// example:
	//
	// true
	NormalizationFieldReserved *bool `json:"NormalizationFieldReserved,omitempty" xml:"NormalizationFieldReserved,omitempty"`
	// Specifies whether the standard field is tokenized.
	//
	// example:
	//
	// true
	NormalizationFieldTokenize *bool `json:"NormalizationFieldTokenize,omitempty" xml:"NormalizationFieldTokenize,omitempty"`
	// The standard field type. Valid values: text, long, double, and json.
	//
	// This parameter is required.
	//
	// example:
	//
	// text
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
	// The key description for a standard field of the JSON type.
	//
	// example:
	//
	// The alert severity levels are represented by the values 1, 2, 3, and 4.
	NormalizationFieldDescription *string `json:"NormalizationFieldDescription,omitempty" xml:"NormalizationFieldDescription,omitempty"`
	// The key example for a standard field of the JSON type.
	//
	// example:
	//
	// 1, 2, 3, 4
	NormalizationFieldExample *string `json:"NormalizationFieldExample,omitempty" xml:"NormalizationFieldExample,omitempty"`
	// The key source for a standard field of the JSON type.
	//
	// example:
	//
	// preset
	NormalizationFieldFrom *string `json:"NormalizationFieldFrom,omitempty" xml:"NormalizationFieldFrom,omitempty"`
	// The key name for a standard field of the JSON type.
	//
	// This parameter is required.
	//
	// example:
	//
	// alert_level
	NormalizationFieldName *string `json:"NormalizationFieldName,omitempty" xml:"NormalizationFieldName,omitempty"`
	// Specifies whether the key for a standard field of the JSON type is required.
	//
	// example:
	//
	// true
	NormalizationFieldRequired *bool `json:"NormalizationFieldRequired,omitempty" xml:"NormalizationFieldRequired,omitempty"`
	// Specifies whether the key for a standard field of the JSON type is tokenized.
	//
	// example:
	//
	// true
	NormalizationFieldTokenize *bool `json:"NormalizationFieldTokenize,omitempty" xml:"NormalizationFieldTokenize,omitempty"`
	// The key type for a standard field of the JSON type. Valid values: text, long, double, and json.
	//
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
