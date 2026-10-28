// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductQuotasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDimensions(v []*ListProductQuotasRequestDimensions) *ListProductQuotasRequest
	GetDimensions() []*ListProductQuotasRequestDimensions
	SetGroupCode(v string) *ListProductQuotasRequest
	GetGroupCode() *string
	SetKeyWord(v string) *ListProductQuotasRequest
	GetKeyWord() *string
	SetMaxResults(v int32) *ListProductQuotasRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListProductQuotasRequest
	GetNextToken() *string
	SetProductCode(v string) *ListProductQuotasRequest
	GetProductCode() *string
	SetQuotaActionCode(v string) *ListProductQuotasRequest
	GetQuotaActionCode() *string
	SetQuotaCategory(v string) *ListProductQuotasRequest
	GetQuotaCategory() *string
	SetSortField(v string) *ListProductQuotasRequest
	GetSortField() *string
	SetSortOrder(v string) *ListProductQuotasRequest
	GetSortOrder() *string
}

type ListProductQuotasRequest struct {
	// The quota dimensions.
	Dimensions []*ListProductQuotasRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The code of the dimension group.
	//
	// example:
	//
	// entconsole_w1j3ms****
	GroupCode *string `json:"GroupCode,omitempty" xml:"GroupCode,omitempty"`
	// The search keyword for the quota.
	//
	// > This parameter applies only to Elastic Computing Service server instance type quotas (ecs-spec). The search scope includes `QuotaName` and `QuotaActionCode`.
	//
	// example:
	//
	// security
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	// The maximum number of entries to return in a single request.
	//
	// Valid values: 1 to 100. Default value: 30.
	//
	// example:
	//
	// 30
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to retrieve the next page of results.
	//
	// - Leave this parameter empty for the first request or when no more results are available.
	//
	// - If more results are available, set this parameter to the NextToken value returned when you invoke the previous API call.
	//
	// example:
	//
	// AAAAAd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The abbreviation of the cloud service.
	//
	// > For more information about how to obtain the abbreviation of a cloud service, see `ProductCode` in [ListProducts](https://help.aliyun.com/document_detail/440555.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota.
	//
	// example:
	//
	// q_security-groups
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The quota category.
	//
	// Default value: CommonQuota.
	//
	// example:
	//
	// FlowControl
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
	// The sorting method supported by the quota.
	//
	// > This parameter applies only to Elastic Computing Service server instance type quotas (ecs-spec). If no Settings are specified, it is left empty.
	//
	// example:
	//
	// TIME
	SortField *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	// The sorting method supported by the quota.
	//
	// > This parameter applies only to Elastic Computing Service server instance type quotas (ecs-spec). If no Settings are specified, it is left empty.
	//
	// example:
	//
	// Ascending
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s ListProductQuotasRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProductQuotasRequest) GoString() string {
	return s.String()
}

func (s *ListProductQuotasRequest) GetDimensions() []*ListProductQuotasRequestDimensions {
	return s.Dimensions
}

func (s *ListProductQuotasRequest) GetGroupCode() *string {
	return s.GroupCode
}

func (s *ListProductQuotasRequest) GetKeyWord() *string {
	return s.KeyWord
}

func (s *ListProductQuotasRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProductQuotasRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProductQuotasRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListProductQuotasRequest) GetQuotaActionCode() *string {
	return s.QuotaActionCode
}

func (s *ListProductQuotasRequest) GetQuotaCategory() *string {
	return s.QuotaCategory
}

func (s *ListProductQuotasRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListProductQuotasRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListProductQuotasRequest) SetDimensions(v []*ListProductQuotasRequestDimensions) *ListProductQuotasRequest {
	s.Dimensions = v
	return s
}

func (s *ListProductQuotasRequest) SetGroupCode(v string) *ListProductQuotasRequest {
	s.GroupCode = &v
	return s
}

func (s *ListProductQuotasRequest) SetKeyWord(v string) *ListProductQuotasRequest {
	s.KeyWord = &v
	return s
}

func (s *ListProductQuotasRequest) SetMaxResults(v int32) *ListProductQuotasRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProductQuotasRequest) SetNextToken(v string) *ListProductQuotasRequest {
	s.NextToken = &v
	return s
}

func (s *ListProductQuotasRequest) SetProductCode(v string) *ListProductQuotasRequest {
	s.ProductCode = &v
	return s
}

func (s *ListProductQuotasRequest) SetQuotaActionCode(v string) *ListProductQuotasRequest {
	s.QuotaActionCode = &v
	return s
}

func (s *ListProductQuotasRequest) SetQuotaCategory(v string) *ListProductQuotasRequest {
	s.QuotaCategory = &v
	return s
}

func (s *ListProductQuotasRequest) SetSortField(v string) *ListProductQuotasRequest {
	s.SortField = &v
	return s
}

func (s *ListProductQuotasRequest) SetSortOrder(v string) *ListProductQuotasRequest {
	s.SortOrder = &v
	return s
}

func (s *ListProductQuotasRequest) Validate() error {
	if s.Dimensions != nil {
		for _, item := range s.Dimensions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProductQuotasRequestDimensions struct {
	// The key of the quota dimension.
	//
	// > The value range of N depends on the number of dimensions supported by the corresponding cloud service.
	//
	// example:
	//
	// regionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the quota dimension.
	//
	// > The value range of N depends on the number of dimensions supported by the corresponding cloud service.
	//
	// example:
	//
	// cn-hangzhou
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListProductQuotasRequestDimensions) String() string {
	return dara.Prettify(s)
}

func (s ListProductQuotasRequestDimensions) GoString() string {
	return s.String()
}

func (s *ListProductQuotasRequestDimensions) GetKey() *string {
	return s.Key
}

func (s *ListProductQuotasRequestDimensions) GetValue() *string {
	return s.Value
}

func (s *ListProductQuotasRequestDimensions) SetKey(v string) *ListProductQuotasRequestDimensions {
	s.Key = &v
	return s
}

func (s *ListProductQuotasRequestDimensions) SetValue(v string) *ListProductQuotasRequestDimensions {
	s.Value = &v
	return s
}

func (s *ListProductQuotasRequestDimensions) Validate() error {
	return dara.Validate(s)
}
