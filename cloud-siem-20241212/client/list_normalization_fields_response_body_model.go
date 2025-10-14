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
	// example:
	//
	// 50。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****。
	NextToken           *string                                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	NormalizationFields []*ListNormalizationFieldsResponseBodyNormalizationFields `json:"NormalizationFields,omitempty" xml:"NormalizationFields,omitempty" type:"Repeated"`
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 57。
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
	// example:
	//
	// NETWORK_CATEGORY。
	NormalizationCategoryId *string `json:"NormalizationCategoryId,omitempty" xml:"NormalizationCategoryId,omitempty"`
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
	// example:
	//
	// NETWORK_SESSION_ACTIVITY。
	NormalizationSchemaId *string `json:"NormalizationSchemaId,omitempty" xml:"NormalizationSchemaId,omitempty"`
}

func (s ListNormalizationFieldsResponseBodyNormalizationFields) String() string {
	return dara.Prettify(s)
}

func (s ListNormalizationFieldsResponseBodyNormalizationFields) GoString() string {
	return s.String()
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

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) GetNormalizationFieldName() *string {
	return s.NormalizationFieldName
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) GetNormalizationFieldRequirement() *bool {
	return s.NormalizationFieldRequirement
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) GetNormalizationFieldReserved() *bool {
	return s.NormalizationFieldReserved
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) GetNormalizationFieldType() *string {
	return s.NormalizationFieldType
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) GetNormalizationSchemaId() *string {
	return s.NormalizationSchemaId
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

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) SetNormalizationFieldName(v string) *ListNormalizationFieldsResponseBodyNormalizationFields {
	s.NormalizationFieldName = &v
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

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) SetNormalizationFieldType(v string) *ListNormalizationFieldsResponseBodyNormalizationFields {
	s.NormalizationFieldType = &v
	return s
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) SetNormalizationSchemaId(v string) *ListNormalizationFieldsResponseBodyNormalizationFields {
	s.NormalizationSchemaId = &v
	return s
}

func (s *ListNormalizationFieldsResponseBodyNormalizationFields) Validate() error {
	return dara.Validate(s)
}
