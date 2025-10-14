// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNormalizationSchemaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNormalizationSchema(v *GetNormalizationSchemaResponseBodyNormalizationSchema) *GetNormalizationSchemaResponseBody
	GetNormalizationSchema() *GetNormalizationSchemaResponseBodyNormalizationSchema
	SetRequestId(v string) *GetNormalizationSchemaResponseBody
	GetRequestId() *string
}

type GetNormalizationSchemaResponseBody struct {
	NormalizationSchema *GetNormalizationSchemaResponseBodyNormalizationSchema `json:"NormalizationSchema,omitempty" xml:"NormalizationSchema,omitempty" type:"Struct"`
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNormalizationSchemaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNormalizationSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *GetNormalizationSchemaResponseBody) GetNormalizationSchema() *GetNormalizationSchemaResponseBodyNormalizationSchema {
	return s.NormalizationSchema
}

func (s *GetNormalizationSchemaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNormalizationSchemaResponseBody) SetNormalizationSchema(v *GetNormalizationSchemaResponseBodyNormalizationSchema) *GetNormalizationSchemaResponseBody {
	s.NormalizationSchema = v
	return s
}

func (s *GetNormalizationSchemaResponseBody) SetRequestId(v string) *GetNormalizationSchemaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNormalizationSchemaResponseBody) Validate() error {
	if s.NormalizationSchema != nil {
		if err := s.NormalizationSchema.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetNormalizationSchemaResponseBodyNormalizationSchema struct {
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// NETWORK_CATEGORY。
	NormalizationCategoryId *string                                                                     `json:"NormalizationCategoryId,omitempty" xml:"NormalizationCategoryId,omitempty"`
	NormalizationFields     []*GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields `json:"NormalizationFields,omitempty" xml:"NormalizationFields,omitempty" type:"Repeated"`
	// example:
	//
	// HTTP_ACTIVITY。
	NormalizationSchemaId *string `json:"NormalizationSchemaId,omitempty" xml:"NormalizationSchemaId,omitempty"`
	// example:
	//
	// normalization_rule_Z57np。
	NormalizationSchemaName *string `json:"NormalizationSchemaName,omitempty" xml:"NormalizationSchemaName,omitempty"`
	NormalizationSchemaType *string `json:"NormalizationSchemaType,omitempty" xml:"NormalizationSchemaType,omitempty"`
	// example:
	//
	// network-activity。
	TargetLogStore *string `json:"TargetLogStore,omitempty" xml:"TargetLogStore,omitempty"`
	// example:
	//
	// network_activity。
	TargetStoreView *string `json:"TargetStoreView,omitempty" xml:"TargetStoreView,omitempty"`
	UpdateTime      *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetNormalizationSchemaResponseBodyNormalizationSchema) String() string {
	return dara.Prettify(s)
}

func (s GetNormalizationSchemaResponseBodyNormalizationSchema) GoString() string {
	return s.String()
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) GetNormalizationCategoryId() *string {
	return s.NormalizationCategoryId
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) GetNormalizationFields() []*GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields {
	return s.NormalizationFields
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) GetNormalizationSchemaId() *string {
	return s.NormalizationSchemaId
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) GetNormalizationSchemaName() *string {
	return s.NormalizationSchemaName
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) GetNormalizationSchemaType() *string {
	return s.NormalizationSchemaType
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) GetTargetLogStore() *string {
	return s.TargetLogStore
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) GetTargetStoreView() *string {
	return s.TargetStoreView
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) SetCreateTime(v int64) *GetNormalizationSchemaResponseBodyNormalizationSchema {
	s.CreateTime = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) SetNormalizationCategoryId(v string) *GetNormalizationSchemaResponseBodyNormalizationSchema {
	s.NormalizationCategoryId = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) SetNormalizationFields(v []*GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) *GetNormalizationSchemaResponseBodyNormalizationSchema {
	s.NormalizationFields = v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) SetNormalizationSchemaId(v string) *GetNormalizationSchemaResponseBodyNormalizationSchema {
	s.NormalizationSchemaId = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) SetNormalizationSchemaName(v string) *GetNormalizationSchemaResponseBodyNormalizationSchema {
	s.NormalizationSchemaName = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) SetNormalizationSchemaType(v string) *GetNormalizationSchemaResponseBodyNormalizationSchema {
	s.NormalizationSchemaType = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) SetTargetLogStore(v string) *GetNormalizationSchemaResponseBodyNormalizationSchema {
	s.TargetLogStore = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) SetTargetStoreView(v string) *GetNormalizationSchemaResponseBodyNormalizationSchema {
	s.TargetStoreView = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) SetUpdateTime(v int64) *GetNormalizationSchemaResponseBodyNormalizationSchema {
	s.UpdateTime = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) Validate() error {
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

type GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields struct {
	// example:
	//
	// cloud_user。
	NormalizationFieldDescription *string `json:"NormalizationFieldDescription,omitempty" xml:"NormalizationFieldDescription,omitempty"`
	// example:
	//
	// 173326*******。
	NormalizationFieldExample *string `json:"NormalizationFieldExample,omitempty" xml:"NormalizationFieldExample,omitempty"`
	// example:
	//
	// cloud_user。
	NormalizationFieldName *string `json:"NormalizationFieldName,omitempty" xml:"NormalizationFieldName,omitempty"`
	// example:
	//
	// true。
	NormalizationFieldRequirement *bool `json:"NormalizationFieldRequirement,omitempty" xml:"NormalizationFieldRequirement,omitempty"`
	// example:
	//
	// true。
	NormalizationFieldReserved *bool `json:"NormalizationFieldReserved,omitempty" xml:"NormalizationFieldReserved,omitempty"`
	// example:
	//
	// varchar。
	NormalizationFieldType *string `json:"NormalizationFieldType,omitempty" xml:"NormalizationFieldType,omitempty"`
}

func (s GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) String() string {
	return dara.Prettify(s)
}

func (s GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) GoString() string {
	return s.String()
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) GetNormalizationFieldDescription() *string {
	return s.NormalizationFieldDescription
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) GetNormalizationFieldExample() *string {
	return s.NormalizationFieldExample
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) GetNormalizationFieldName() *string {
	return s.NormalizationFieldName
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) GetNormalizationFieldRequirement() *bool {
	return s.NormalizationFieldRequirement
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) GetNormalizationFieldReserved() *bool {
	return s.NormalizationFieldReserved
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) GetNormalizationFieldType() *string {
	return s.NormalizationFieldType
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) SetNormalizationFieldDescription(v string) *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields {
	s.NormalizationFieldDescription = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) SetNormalizationFieldExample(v string) *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields {
	s.NormalizationFieldExample = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) SetNormalizationFieldName(v string) *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields {
	s.NormalizationFieldName = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) SetNormalizationFieldRequirement(v bool) *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields {
	s.NormalizationFieldRequirement = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) SetNormalizationFieldReserved(v bool) *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields {
	s.NormalizationFieldReserved = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) SetNormalizationFieldType(v string) *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields {
	s.NormalizationFieldType = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) Validate() error {
	return dara.Validate(s)
}
