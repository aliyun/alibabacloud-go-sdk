// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotaApplicationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDimensions(v []*ListQuotaApplicationsRequestDimensions) *ListQuotaApplicationsRequest
	GetDimensions() []*ListQuotaApplicationsRequestDimensions
	SetKeyWord(v string) *ListQuotaApplicationsRequest
	GetKeyWord() *string
	SetMaxResults(v int32) *ListQuotaApplicationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListQuotaApplicationsRequest
	GetNextToken() *string
	SetProductCode(v string) *ListQuotaApplicationsRequest
	GetProductCode() *string
	SetQuotaActionCode(v string) *ListQuotaApplicationsRequest
	GetQuotaActionCode() *string
	SetQuotaCategory(v string) *ListQuotaApplicationsRequest
	GetQuotaCategory() *string
	SetStatus(v string) *ListQuotaApplicationsRequest
	GetStatus() *string
}

type ListQuotaApplicationsRequest struct {
	// The quota dimensions.
	//
	// The properties of the quota item. The combination of dimensions and the quota ID (QuotaActionCode) uniquely identifies a quota item.
	//
	// > The dimensions are required for some cloud services. You can call the [ListProductQuotaDimensions](~~ListProductQuotaDimensions~~) operation to query the quota dimensions supported by a cloud service. The Requisite response parameter indicates whether a dimension is required.
	Dimensions []*ListQuotaApplicationsRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The search keyword for quota increase applications.
	//
	// example:
	//
	// Cluster
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	// The maximum number of entries to return for the current query.
	//
	// Valid values: 1 to 200. Default value: 30.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that specifies the start position of the query. An empty value indicates that the query starts from the beginning.
	//
	// example:
	//
	// 1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The abbreviation of the cloud service.
	//
	// example:
	//
	// csk
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// example:
	//
	// q_i5uzm3
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The quota type. Valid values:
	//
	// - CommonQuota: General quota.
	//
	// - FlowControl: API rate limit.
	//
	// - WhiteListLabel: Privilege.
	//
	// example:
	//
	// CommonQuota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
	// The approval status of the quota increase application. Valid values:
	//
	// - Disagree: Rejected.
	//
	// - Agree: Approved.
	//
	// - Process: Under review.
	//
	// - Cancel: Closed.
	//
	// - Withdrawn: Withdrawn.
	//
	// example:
	//
	// Agree
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListQuotaApplicationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaApplicationsRequest) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationsRequest) GetDimensions() []*ListQuotaApplicationsRequestDimensions {
	return s.Dimensions
}

func (s *ListQuotaApplicationsRequest) GetKeyWord() *string {
	return s.KeyWord
}

func (s *ListQuotaApplicationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListQuotaApplicationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListQuotaApplicationsRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListQuotaApplicationsRequest) GetQuotaActionCode() *string {
	return s.QuotaActionCode
}

func (s *ListQuotaApplicationsRequest) GetQuotaCategory() *string {
	return s.QuotaCategory
}

func (s *ListQuotaApplicationsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListQuotaApplicationsRequest) SetDimensions(v []*ListQuotaApplicationsRequestDimensions) *ListQuotaApplicationsRequest {
	s.Dimensions = v
	return s
}

func (s *ListQuotaApplicationsRequest) SetKeyWord(v string) *ListQuotaApplicationsRequest {
	s.KeyWord = &v
	return s
}

func (s *ListQuotaApplicationsRequest) SetMaxResults(v int32) *ListQuotaApplicationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListQuotaApplicationsRequest) SetNextToken(v string) *ListQuotaApplicationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListQuotaApplicationsRequest) SetProductCode(v string) *ListQuotaApplicationsRequest {
	s.ProductCode = &v
	return s
}

func (s *ListQuotaApplicationsRequest) SetQuotaActionCode(v string) *ListQuotaApplicationsRequest {
	s.QuotaActionCode = &v
	return s
}

func (s *ListQuotaApplicationsRequest) SetQuotaCategory(v string) *ListQuotaApplicationsRequest {
	s.QuotaCategory = &v
	return s
}

func (s *ListQuotaApplicationsRequest) SetStatus(v string) *ListQuotaApplicationsRequest {
	s.Status = &v
	return s
}

func (s *ListQuotaApplicationsRequest) Validate() error {
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

type ListQuotaApplicationsRequestDimensions struct {
	// The key of the quota dimension.
	//
	// > Dimensions.N.Key and Dimensions.N.Value must be specified at the same time. The value range of N depends on the number of dimensions supported by the corresponding cloud service. You can call the [ListProductQuotaDimensions](~~ListProductQuotaDimensions~~) operation to query the quota dimensions supported by the cloud service. The length of the returned array is the value of N.
	//
	// example:
	//
	// regionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the quota dimension.
	//
	// > Dimensions.N.Key and Dimensions.N.Value must be specified at the same time. The value range of N depends on the number of dimensions supported by the corresponding cloud service. You can call the [ListProductQuotaDimensions](~~ListProductQuotaDimensions~~) operation to query the quota dimensions supported by the cloud service. The length of the returned array is the value of N.
	//
	// example:
	//
	// cn-hangzhou
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListQuotaApplicationsRequestDimensions) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaApplicationsRequestDimensions) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationsRequestDimensions) GetKey() *string {
	return s.Key
}

func (s *ListQuotaApplicationsRequestDimensions) GetValue() *string {
	return s.Value
}

func (s *ListQuotaApplicationsRequestDimensions) SetKey(v string) *ListQuotaApplicationsRequestDimensions {
	s.Key = &v
	return s
}

func (s *ListQuotaApplicationsRequestDimensions) SetValue(v string) *ListQuotaApplicationsRequestDimensions {
	s.Value = &v
	return s
}

func (s *ListQuotaApplicationsRequestDimensions) Validate() error {
	return dara.Validate(s)
}
