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
	// The maximum number of entries to return in the current request.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to retrieve the next page of results. Leave this parameter empty for the first request or when no more results are available. For subsequent requests, set this parameter to the NextToken value returned in the previous API call.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of normalization schemas.
	NormalizationSchemas []*ListNormalizationSchemasResponseBodyNormalizationSchemas `json:"NormalizationSchemas,omitempty" xml:"NormalizationSchemas,omitempty" type:"Repeated"`
	// The ID of the request.
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
	// The time when the normalization schema was created. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1736386610000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the normalization rule category.
	//
	// example:
	//
	// NETWORK_CATEGORY
	NormalizationCategoryId *string `json:"NormalizationCategoryId,omitempty" xml:"NormalizationCategoryId,omitempty"`
	// The source of the field. Valid values:
	//
	// - normalized: Normalized field.
	//
	// - native: Native field.
	//
	// example:
	//
	// normalized
	NormalizationFieldSource *string `json:"NormalizationFieldSource,omitempty" xml:"NormalizationFieldSource,omitempty"`
	// The description of the normalization schema.
	//
	// example:
	//
	// Network flow log
	NormalizationSchemaDescription *string `json:"NormalizationSchemaDescription,omitempty" xml:"NormalizationSchemaDescription,omitempty"`
	// The source of the normalization schema. Valid values: preset and custom.
	//
	// example:
	//
	// preset
	NormalizationSchemaFrom *string `json:"NormalizationSchemaFrom,omitempty" xml:"NormalizationSchemaFrom,omitempty"`
	// The ID of the normalization schema.
	//
	// example:
	//
	// HTTP_ACTIVITY
	NormalizationSchemaId *string `json:"NormalizationSchemaId,omitempty" xml:"NormalizationSchemaId,omitempty"`
	// The name of the normalization schema.
	//
	// example:
	//
	// normalization_rule_Z57np
	NormalizationSchemaName *string `json:"NormalizationSchemaName,omitempty" xml:"NormalizationSchemaName,omitempty"`
	// The Logstore for normalized output.
	//
	// example:
	//
	// network-activity
	NormalizationSchemaTargetLogStore *string `json:"NormalizationSchemaTargetLogStore,omitempty" xml:"NormalizationSchemaTargetLogStore,omitempty"`
	// The type of the normalization schema.
	//
	// example:
	//
	// log
	NormalizationSchemaType *string `json:"NormalizationSchemaType,omitempty" xml:"NormalizationSchemaType,omitempty"`
	// The ID of the security domain.
	//
	// example:
	//
	// NETWORK_AND_WEB_SECURITY
	NormalizationSecurityDomainId *string `json:"NormalizationSecurityDomainId,omitempty" xml:"NormalizationSecurityDomainId,omitempty"`
	// The ID of the product.
	//
	// example:
	//
	// sas
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The list of recommended entities.
	RecommendEntities []*string `json:"RecommendEntities,omitempty" xml:"RecommendEntities,omitempty" type:"Repeated"`
	// The Logstore of Simple Log Service.
	//
	// example:
	//
	// network-activity
	TargetLogStore *string `json:"TargetLogStore,omitempty" xml:"TargetLogStore,omitempty"`
	// The StoreView of Simple Log Service.
	//
	// example:
	//
	// network-activity
	TargetStoreView *string `json:"TargetStoreView,omitempty" xml:"TargetStoreView,omitempty"`
	// The time when the normalization schema was updated. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1736386610000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The ID of the vendor.
	//
	// example:
	//
	// alibaba
	VendorId *string `json:"VendorId,omitempty" xml:"VendorId,omitempty"`
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

func (s *ListNormalizationSchemasResponseBodyNormalizationSchemas) GetNormalizationFieldSource() *string {
	return s.NormalizationFieldSource
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

func (s *ListNormalizationSchemasResponseBodyNormalizationSchemas) GetNormalizationSchemaType() *string {
	return s.NormalizationSchemaType
}

func (s *ListNormalizationSchemasResponseBodyNormalizationSchemas) GetNormalizationSecurityDomainId() *string {
	return s.NormalizationSecurityDomainId
}

func (s *ListNormalizationSchemasResponseBodyNormalizationSchemas) GetProductId() *string {
	return s.ProductId
}

func (s *ListNormalizationSchemasResponseBodyNormalizationSchemas) GetRecommendEntities() []*string {
	return s.RecommendEntities
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

func (s *ListNormalizationSchemasResponseBodyNormalizationSchemas) GetVendorId() *string {
	return s.VendorId
}

func (s *ListNormalizationSchemasResponseBodyNormalizationSchemas) SetCreateTime(v int64) *ListNormalizationSchemasResponseBodyNormalizationSchemas {
	s.CreateTime = &v
	return s
}

func (s *ListNormalizationSchemasResponseBodyNormalizationSchemas) SetNormalizationCategoryId(v string) *ListNormalizationSchemasResponseBodyNormalizationSchemas {
	s.NormalizationCategoryId = &v
	return s
}

func (s *ListNormalizationSchemasResponseBodyNormalizationSchemas) SetNormalizationFieldSource(v string) *ListNormalizationSchemasResponseBodyNormalizationSchemas {
	s.NormalizationFieldSource = &v
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

func (s *ListNormalizationSchemasResponseBodyNormalizationSchemas) SetNormalizationSchemaType(v string) *ListNormalizationSchemasResponseBodyNormalizationSchemas {
	s.NormalizationSchemaType = &v
	return s
}

func (s *ListNormalizationSchemasResponseBodyNormalizationSchemas) SetNormalizationSecurityDomainId(v string) *ListNormalizationSchemasResponseBodyNormalizationSchemas {
	s.NormalizationSecurityDomainId = &v
	return s
}

func (s *ListNormalizationSchemasResponseBodyNormalizationSchemas) SetProductId(v string) *ListNormalizationSchemasResponseBodyNormalizationSchemas {
	s.ProductId = &v
	return s
}

func (s *ListNormalizationSchemasResponseBodyNormalizationSchemas) SetRecommendEntities(v []*string) *ListNormalizationSchemasResponseBodyNormalizationSchemas {
	s.RecommendEntities = v
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

func (s *ListNormalizationSchemasResponseBodyNormalizationSchemas) SetVendorId(v string) *ListNormalizationSchemasResponseBodyNormalizationSchemas {
	s.VendorId = &v
	return s
}

func (s *ListNormalizationSchemasResponseBodyNormalizationSchemas) Validate() error {
	return dara.Validate(s)
}
