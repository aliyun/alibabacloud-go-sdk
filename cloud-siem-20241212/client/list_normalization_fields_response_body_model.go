// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNormalizationFieldsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListNormalizationFieldsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListNormalizationFieldsResponseBody
	GetNextToken() *string
	SetNormalizationFields(v []*ListNormalizationFieldsResponseBodyNormalizationFields) *ListNormalizationFieldsResponseBody
	GetNormalizationFields() []*ListNormalizationFieldsResponseBodyNormalizationFields
	SetRequestId(v string) *ListNormalizationFieldsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListNormalizationFieldsResponseBody
	GetTotalCount() *int32
}

type ListNormalizationFieldsResponseBody struct {
	// The maximum number of entries to return in this request.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Specifies whether a next query token exists. You do not need to specify this parameter for the first query or if no next query exists. If a next query exists, set this parameter to the NextToken value returned by the previous API call.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of normalization fields.
	NormalizationFields []*ListNormalizationFieldsResponseBodyNormalizationFields `json:"NormalizationFields,omitempty" xml:"NormalizationFields,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 57
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNormalizationFieldsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNormalizationFieldsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNormalizationFieldsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNormalizationFieldsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNormalizationFieldsResponseBody) GetNormalizationFields() []*ListNormalizationFieldsResponseBodyNormalizationFields {
	return s.NormalizationFields
}

func (s *ListNormalizationFieldsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNormalizationFieldsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListNormalizationFieldsResponseBody) SetMaxResults(v int32) *ListNormalizationFieldsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListNormalizationFieldsResponseBody) SetNextToken(v string) *ListNormalizationFieldsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListNormalizationFieldsResponseBody) SetNormalizationFields(v []*ListNormalizationFieldsResponseBodyNormalizationFields) *ListNormalizationFieldsResponseBody {
	s.NormalizationFields = v
	return s
}

func (s *ListNormalizationFieldsResponseBody) SetRequestId(v string) *ListNormalizationFieldsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNormalizationFieldsResponseBody) SetTotalCount(v int32) *ListNormalizationFieldsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListNormalizationFieldsResponseBody) Validate() error {
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

type ListNormalizationFieldsResponseBodyNormalizationFields struct {
	// The creation time.
	//
	// example:
	//
	// 1736386610000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the normalization category.
	//
	// example:
	//
	// NETWORK_CATEGORY
	NormalizationCategoryId *string `json:"NormalizationCategoryId,omitempty" xml:"NormalizationCategoryId,omitempty"`
	// The description of the normalization field.
	//
	// example:
	//
	// cloud_user
	NormalizationFieldDescription *string `json:"NormalizationFieldDescription,omitempty" xml:"NormalizationFieldDescription,omitempty"`
	// The example value of the normalization field.
	//
	// example:
	//
	// 173326*******
	NormalizationFieldExample *string `json:"NormalizationFieldExample,omitempty" xml:"NormalizationFieldExample,omitempty"`
	// The source of the normalization field key of the JSON type.
	//
	// example:
	//
	// preset
	NormalizationFieldFrom *string `json:"NormalizationFieldFrom,omitempty" xml:"NormalizationFieldFrom,omitempty"`
	// Indicates whether indexes are created for all keys of the JSON-type normalization field.
	NormalizationFieldJsonIndexAll *bool `json:"NormalizationFieldJsonIndexAll,omitempty" xml:"NormalizationFieldJsonIndexAll,omitempty"`
	// The list of keys for the JSON-type normalization field.
	NormalizationFieldJsonKeys []*ListNormalizationFieldsResponseBodyNormalizationFieldsNormalizationFieldJsonKeys `json:"NormalizationFieldJsonKeys,omitempty" xml:"NormalizationFieldJsonKeys,omitempty" type:"Repeated"`
	// The name of the normalization field.
	//
	// example:
	//
	// cloud_user
	NormalizationFieldName *string `json:"NormalizationFieldName,omitempty" xml:"NormalizationFieldName,omitempty"`
	// Indicates whether the normalization field key is required.
	NormalizationFieldRequired *bool `json:"NormalizationFieldRequired,omitempty" xml:"NormalizationFieldRequired,omitempty"`
	// Indicates whether the normalization field is required.
	//
	// example:
	//
	// true
	NormalizationFieldRequirement *bool `json:"NormalizationFieldRequirement,omitempty" xml:"NormalizationFieldRequirement,omitempty"`
	// Indicates whether the normalization field is reserved.
	//
	// example:
	//
	// true
	NormalizationFieldReserved *bool `json:"NormalizationFieldReserved,omitempty" xml:"NormalizationFieldReserved,omitempty"`
	// Indicates whether the normalization field is tokenized.
	NormalizationFieldTokenize *bool `json:"NormalizationFieldTokenize,omitempty" xml:"NormalizationFieldTokenize,omitempty"`
	// The type of the normalization field. Valid values:
	//
	// - varchar
	//
	// - bigint
	//
	// - double.
	//
	// example:
	//
	// varchar
	NormalizationFieldType *string `json:"NormalizationFieldType,omitempty" xml:"NormalizationFieldType,omitempty"`
	// The ID of the normalization schema.
	//
	// example:
	//
	// NETWORK_SESSION_ACTIVITY
	NormalizationSchemaId *string `json:"NormalizationSchemaId,omitempty" xml:"NormalizationSchemaId,omitempty"`
	// example:
	//
	// NETWORK_AND_WEB_SECURITY
	NormalizationSecurityDomainId *string `json:"NormalizationSecurityDomainId,omitempty" xml:"NormalizationSecurityDomainId,omitempty"`
	// The update time.
	//
	// example:
	//
	// 1736386610000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListNormalizationFieldsResponseBodyNormalizationFields) String() string {
	return dara.Prettify(s)
}

func (s ListNormalizationFieldsResponseBodyNormalizationFields) GoString() string {
	return s.String()
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) GetNormalizationCategoryId() *string {
	return s.NormalizationCategoryId
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) GetNormalizationFieldDescription() *string {
	return s.NormalizationFieldDescription
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) GetNormalizationFieldExample() *string {
	return s.NormalizationFieldExample
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) GetNormalizationFieldFrom() *string {
	return s.NormalizationFieldFrom
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) GetNormalizationFieldJsonIndexAll() *bool {
	return s.NormalizationFieldJsonIndexAll
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) GetNormalizationFieldJsonKeys() []*ListNormalizationFieldsResponseBodyNormalizationFieldsNormalizationFieldJsonKeys {
	return s.NormalizationFieldJsonKeys
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) GetNormalizationFieldName() *string {
	return s.NormalizationFieldName
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) GetNormalizationFieldRequired() *bool {
	return s.NormalizationFieldRequired
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) GetNormalizationFieldRequirement() *bool {
	return s.NormalizationFieldRequirement
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) GetNormalizationFieldReserved() *bool {
	return s.NormalizationFieldReserved
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) GetNormalizationFieldTokenize() *bool {
	return s.NormalizationFieldTokenize
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) GetNormalizationFieldType() *string {
	return s.NormalizationFieldType
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) GetNormalizationSchemaId() *string {
	return s.NormalizationSchemaId
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) GetNormalizationSecurityDomainId() *string {
	return s.NormalizationSecurityDomainId
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) SetCreateTime(v int64) *ListNormalizationFieldsResponseBodyNormalizationFields {
	s.CreateTime = &v
	return s
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) SetNormalizationCategoryId(v string) *ListNormalizationFieldsResponseBodyNormalizationFields {
	s.NormalizationCategoryId = &v
	return s
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) SetNormalizationFieldDescription(v string) *ListNormalizationFieldsResponseBodyNormalizationFields {
	s.NormalizationFieldDescription = &v
	return s
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) SetNormalizationFieldExample(v string) *ListNormalizationFieldsResponseBodyNormalizationFields {
	s.NormalizationFieldExample = &v
	return s
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) SetNormalizationFieldFrom(v string) *ListNormalizationFieldsResponseBodyNormalizationFields {
	s.NormalizationFieldFrom = &v
	return s
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) SetNormalizationFieldJsonIndexAll(v bool) *ListNormalizationFieldsResponseBodyNormalizationFields {
	s.NormalizationFieldJsonIndexAll = &v
	return s
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) SetNormalizationFieldJsonKeys(v []*ListNormalizationFieldsResponseBodyNormalizationFieldsNormalizationFieldJsonKeys) *ListNormalizationFieldsResponseBodyNormalizationFields {
	s.NormalizationFieldJsonKeys = v
	return s
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) SetNormalizationFieldName(v string) *ListNormalizationFieldsResponseBodyNormalizationFields {
	s.NormalizationFieldName = &v
	return s
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) SetNormalizationFieldRequired(v bool) *ListNormalizationFieldsResponseBodyNormalizationFields {
	s.NormalizationFieldRequired = &v
	return s
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) SetNormalizationFieldRequirement(v bool) *ListNormalizationFieldsResponseBodyNormalizationFields {
	s.NormalizationFieldRequirement = &v
	return s
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) SetNormalizationFieldReserved(v bool) *ListNormalizationFieldsResponseBodyNormalizationFields {
	s.NormalizationFieldReserved = &v
	return s
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) SetNormalizationFieldTokenize(v bool) *ListNormalizationFieldsResponseBodyNormalizationFields {
	s.NormalizationFieldTokenize = &v
	return s
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) SetNormalizationFieldType(v string) *ListNormalizationFieldsResponseBodyNormalizationFields {
	s.NormalizationFieldType = &v
	return s
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) SetNormalizationSchemaId(v string) *ListNormalizationFieldsResponseBodyNormalizationFields {
	s.NormalizationSchemaId = &v
	return s
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) SetNormalizationSecurityDomainId(v string) *ListNormalizationFieldsResponseBodyNormalizationFields {
	s.NormalizationSecurityDomainId = &v
	return s
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) SetUpdateTime(v int64) *ListNormalizationFieldsResponseBodyNormalizationFields {
	s.UpdateTime = &v
	return s
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) Validate() error {
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

type ListNormalizationFieldsResponseBodyNormalizationFieldsNormalizationFieldJsonKeys struct {
	// The creation time.
	//
	// example:
	//
	// 1736386610000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the JSON-type normalization field key.
	//
	// example:
	//
	// alert_level
	NormalizationFieldDescription *string `json:"NormalizationFieldDescription,omitempty" xml:"NormalizationFieldDescription,omitempty"`
	// The example value of the JSON-type normalization field key.
	//
	// example:
	//
	// 1
	NormalizationFieldExample *string `json:"NormalizationFieldExample,omitempty" xml:"NormalizationFieldExample,omitempty"`
	// The source of the normalization field key of the JSON type.
	//
	// example:
	//
	// preset
	NormalizationFieldFrom *string `json:"NormalizationFieldFrom,omitempty" xml:"NormalizationFieldFrom,omitempty"`
	// The name of the JSON-type normalization field key.
	//
	// example:
	//
	// alert_level
	NormalizationFieldName *string `json:"NormalizationFieldName,omitempty" xml:"NormalizationFieldName,omitempty"`
	// Indicates whether the JSON-type normalization field key is required.
	NormalizationFieldRequired *bool `json:"NormalizationFieldRequired,omitempty" xml:"NormalizationFieldRequired,omitempty"`
	// Indicates whether the normalization field key is a built-in system field key.
	NormalizationFieldReserved *bool `json:"NormalizationFieldReserved,omitempty" xml:"NormalizationFieldReserved,omitempty"`
	// Indicates whether the JSON-type normalization field key is tokenized.
	NormalizationFieldTokenize *bool `json:"NormalizationFieldTokenize,omitempty" xml:"NormalizationFieldTokenize,omitempty"`
	// The type of the JSON-type normalization field key.
	//
	// example:
	//
	// text
	NormalizationFieldType *string `json:"NormalizationFieldType,omitempty" xml:"NormalizationFieldType,omitempty"`
	// The update time.
	//
	// example:
	//
	// 1736386610000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListNormalizationFieldsResponseBodyNormalizationFieldsNormalizationFieldJsonKeys) String() string {
	return dara.Prettify(s)
}

func (s ListNormalizationFieldsResponseBodyNormalizationFieldsNormalizationFieldJsonKeys) GoString() string {
	return s.String()
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFieldsNormalizationFieldJsonKeys) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFieldsNormalizationFieldJsonKeys) GetNormalizationFieldDescription() *string {
	return s.NormalizationFieldDescription
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFieldsNormalizationFieldJsonKeys) GetNormalizationFieldExample() *string {
	return s.NormalizationFieldExample
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFieldsNormalizationFieldJsonKeys) GetNormalizationFieldFrom() *string {
	return s.NormalizationFieldFrom
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFieldsNormalizationFieldJsonKeys) GetNormalizationFieldName() *string {
	return s.NormalizationFieldName
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFieldsNormalizationFieldJsonKeys) GetNormalizationFieldRequired() *bool {
	return s.NormalizationFieldRequired
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFieldsNormalizationFieldJsonKeys) GetNormalizationFieldReserved() *bool {
	return s.NormalizationFieldReserved
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFieldsNormalizationFieldJsonKeys) GetNormalizationFieldTokenize() *bool {
	return s.NormalizationFieldTokenize
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFieldsNormalizationFieldJsonKeys) GetNormalizationFieldType() *string {
	return s.NormalizationFieldType
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFieldsNormalizationFieldJsonKeys) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFieldsNormalizationFieldJsonKeys) SetCreateTime(v int64) *ListNormalizationFieldsResponseBodyNormalizationFieldsNormalizationFieldJsonKeys {
	s.CreateTime = &v
	return s
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFieldsNormalizationFieldJsonKeys) SetNormalizationFieldDescription(v string) *ListNormalizationFieldsResponseBodyNormalizationFieldsNormalizationFieldJsonKeys {
	s.NormalizationFieldDescription = &v
	return s
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFieldsNormalizationFieldJsonKeys) SetNormalizationFieldExample(v string) *ListNormalizationFieldsResponseBodyNormalizationFieldsNormalizationFieldJsonKeys {
	s.NormalizationFieldExample = &v
	return s
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFieldsNormalizationFieldJsonKeys) SetNormalizationFieldFrom(v string) *ListNormalizationFieldsResponseBodyNormalizationFieldsNormalizationFieldJsonKeys {
	s.NormalizationFieldFrom = &v
	return s
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFieldsNormalizationFieldJsonKeys) SetNormalizationFieldName(v string) *ListNormalizationFieldsResponseBodyNormalizationFieldsNormalizationFieldJsonKeys {
	s.NormalizationFieldName = &v
	return s
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFieldsNormalizationFieldJsonKeys) SetNormalizationFieldRequired(v bool) *ListNormalizationFieldsResponseBodyNormalizationFieldsNormalizationFieldJsonKeys {
	s.NormalizationFieldRequired = &v
	return s
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFieldsNormalizationFieldJsonKeys) SetNormalizationFieldReserved(v bool) *ListNormalizationFieldsResponseBodyNormalizationFieldsNormalizationFieldJsonKeys {
	s.NormalizationFieldReserved = &v
	return s
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFieldsNormalizationFieldJsonKeys) SetNormalizationFieldTokenize(v bool) *ListNormalizationFieldsResponseBodyNormalizationFieldsNormalizationFieldJsonKeys {
	s.NormalizationFieldTokenize = &v
	return s
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFieldsNormalizationFieldJsonKeys) SetNormalizationFieldType(v string) *ListNormalizationFieldsResponseBodyNormalizationFieldsNormalizationFieldJsonKeys {
	s.NormalizationFieldType = &v
	return s
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFieldsNormalizationFieldJsonKeys) SetUpdateTime(v int64) *ListNormalizationFieldsResponseBodyNormalizationFieldsNormalizationFieldJsonKeys {
	s.UpdateTime = &v
	return s
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFieldsNormalizationFieldJsonKeys) Validate() error {
	return dara.Validate(s)
}
