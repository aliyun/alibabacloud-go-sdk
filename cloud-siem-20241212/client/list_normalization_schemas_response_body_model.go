// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNormalizationSchemasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListNormalizationSchemasResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListNormalizationSchemasResponseBody
	GetNextToken() *string
	SetNormalizationSchemas(v []*ListNormalizationSchemasResponseBodyNormalizationSchemas) *ListNormalizationSchemasResponseBody
	GetNormalizationSchemas() []*ListNormalizationSchemasResponseBodyNormalizationSchemas
	SetRequestId(v string) *ListNormalizationSchemasResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListNormalizationSchemasResponseBody
	GetTotalCount() *int32
}

type ListNormalizationSchemasResponseBody struct {
	// example:
	//
	// 50。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****。
	NextToken            *string                                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	NormalizationSchemas []*ListNormalizationSchemasResponseBodyNormalizationSchemas `json:"NormalizationSchemas,omitempty" xml:"NormalizationSchemas,omitempty" type:"Repeated"`
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 57。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNormalizationSchemasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNormalizationSchemasResponseBody) GoString() string {
	return s.String()
}

func (s *ListNormalizationSchemasResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNormalizationSchemasResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNormalizationSchemasResponseBody) GetNormalizationSchemas() []*ListNormalizationSchemasResponseBodyNormalizationSchemas {
	return s.NormalizationSchemas
}

func (s *ListNormalizationSchemasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNormalizationSchemasResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListNormalizationSchemasResponseBody) SetMaxResults(v int32) *ListNormalizationSchemasResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListNormalizationSchemasResponseBody) SetNextToken(v string) *ListNormalizationSchemasResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListNormalizationSchemasResponseBody) SetNormalizationSchemas(v []*ListNormalizationSchemasResponseBodyNormalizationSchemas) *ListNormalizationSchemasResponseBody {
	s.NormalizationSchemas = v
	return s
}

func (s *ListNormalizationSchemasResponseBody) SetRequestId(v string) *ListNormalizationSchemasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNormalizationSchemasResponseBody) SetTotalCount(v int32) *ListNormalizationSchemasResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListNormalizationSchemasResponseBody) Validate() error {
	if s.NormalizationSchemas != nil {
		for _, item := range s.NormalizationSchemas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNormalizationSchemasResponseBodyNormalizationSchemas struct {
	// example:
	//
	// 1736386610000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// NETWORK_CATEGORY。
	NormalizationCategoryId *string `json:"NormalizationCategoryId,omitempty" xml:"NormalizationCategoryId,omitempty"`
	// example:
	//
	// 网络五元组日志
	NormalizationSchemaDescription *string `json:"NormalizationSchemaDescription,omitempty" xml:"NormalizationSchemaDescription,omitempty"`
	// example:
	//
	// preset
	NormalizationSchemaFrom *string `json:"NormalizationSchemaFrom,omitempty" xml:"NormalizationSchemaFrom,omitempty"`
	// example:
	//
	// HTTP_ACTIVITY。
	NormalizationSchemaId *string `json:"NormalizationSchemaId,omitempty" xml:"NormalizationSchemaId,omitempty"`
	// example:
	//
	// normalization_rule_Z57np。
	NormalizationSchemaName *string `json:"NormalizationSchemaName,omitempty" xml:"NormalizationSchemaName,omitempty"`
	// example:
	//
	// network-activity。
	NormalizationSchemaTargetLogStore *string `json:"NormalizationSchemaTargetLogStore,omitempty" xml:"NormalizationSchemaTargetLogStore,omitempty"`
	// example:
	//
	// network-activity
	TargetLogStore *string `json:"TargetLogStore,omitempty" xml:"TargetLogStore,omitempty"`
	// example:
	//
	// network-activity
	TargetStoreView *string `json:"TargetStoreView,omitempty" xml:"TargetStoreView,omitempty"`
	// example:
	//
	// 1736386610000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListNormalizationSchemasResponseBodyNormalizationSchemas) String() string {
	return dara.Prettify(s)
}

func (s ListNormalizationSchemasResponseBodyNormalizationSchemas) GoString() string {
	return s.String()
}

func (s *ListNormalizationSchemasResponseBodyNormalizationSchemas) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListNormalizationSchemasResponseBodyNormalizationSchemas) GetNormalizationCategoryId() *string {
	return s.NormalizationCategoryId
}

func (s *ListNormalizationSchemasResponseBodyNormalizationSchemas) GetNormalizationSchemaDescription() *string {
	return s.NormalizationSchemaDescription
}

func (s *ListNormalizationSchemasResponseBodyNormalizationSchemas) GetNormalizationSchemaFrom() *string {
	return s.NormalizationSchemaFrom
}

func (s *ListNormalizationSchemasResponseBodyNormalizationSchemas) GetNormalizationSchemaId() *string {
	return s.NormalizationSchemaId
}

func (s *ListNormalizationSchemasResponseBodyNormalizationSchemas) GetNormalizationSchemaName() *string {
	return s.NormalizationSchemaName
}

func (s *ListNormalizationSchemasResponseBodyNormalizationSchemas) GetNormalizationSchemaTargetLogStore() *string {
	return s.NormalizationSchemaTargetLogStore
}

func (s *ListNormalizationSchemasResponseBodyNormalizationSchemas) GetTargetLogStore() *string {
	return s.TargetLogStore
}

func (s *ListNormalizationSchemasResponseBodyNormalizationSchemas) GetTargetStoreView() *string {
	return s.TargetStoreView
}

func (s *ListNormalizationSchemasResponseBodyNormalizationSchemas) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListNormalizationSchemasResponseBodyNormalizationSchemas) SetCreateTime(v int64) *ListNormalizationSchemasResponseBodyNormalizationSchemas {
	s.CreateTime = &v
	return s
}

func (s *ListNormalizationSchemasResponseBodyNormalizationSchemas) SetNormalizationCategoryId(v string) *ListNormalizationSchemasResponseBodyNormalizationSchemas {
	s.NormalizationCategoryId = &v
	return s
}

func (s *ListNormalizationSchemasResponseBodyNormalizationSchemas) SetNormalizationSchemaDescription(v string) *ListNormalizationSchemasResponseBodyNormalizationSchemas {
	s.NormalizationSchemaDescription = &v
	return s
}

func (s *ListNormalizationSchemasResponseBodyNormalizationSchemas) SetNormalizationSchemaFrom(v string) *ListNormalizationSchemasResponseBodyNormalizationSchemas {
	s.NormalizationSchemaFrom = &v
	return s
}

func (s *ListNormalizationSchemasResponseBodyNormalizationSchemas) SetNormalizationSchemaId(v string) *ListNormalizationSchemasResponseBodyNormalizationSchemas {
	s.NormalizationSchemaId = &v
	return s
}

func (s *ListNormalizationSchemasResponseBodyNormalizationSchemas) SetNormalizationSchemaName(v string) *ListNormalizationSchemasResponseBodyNormalizationSchemas {
	s.NormalizationSchemaName = &v
	return s
}

func (s *ListNormalizationSchemasResponseBodyNormalizationSchemas) SetNormalizationSchemaTargetLogStore(v string) *ListNormalizationSchemasResponseBodyNormalizationSchemas {
	s.NormalizationSchemaTargetLogStore = &v
	return s
}

func (s *ListNormalizationSchemasResponseBodyNormalizationSchemas) SetTargetLogStore(v string) *ListNormalizationSchemasResponseBodyNormalizationSchemas {
	s.TargetLogStore = &v
	return s
}

func (s *ListNormalizationSchemasResponseBodyNormalizationSchemas) SetTargetStoreView(v string) *ListNormalizationSchemasResponseBodyNormalizationSchemas {
	s.TargetStoreView = &v
	return s
}

func (s *ListNormalizationSchemasResponseBodyNormalizationSchemas) SetUpdateTime(v int64) *ListNormalizationSchemasResponseBodyNormalizationSchemas {
	s.UpdateTime = &v
	return s
}

func (s *ListNormalizationSchemasResponseBodyNormalizationSchemas) Validate() error {
	return dara.Validate(s)
}
