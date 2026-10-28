// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductQuotaDimensionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListProductQuotaDimensionsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListProductQuotaDimensionsResponseBody
	GetNextToken() *string
	SetQuotaDimensions(v []*ListProductQuotaDimensionsResponseBodyQuotaDimensions) *ListProductQuotaDimensionsResponseBody
	GetQuotaDimensions() []*ListProductQuotaDimensionsResponseBodyQuotaDimensions
	SetRequestId(v string) *ListProductQuotaDimensionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListProductQuotaDimensionsResponseBody
	GetTotalCount() *int32
}

type ListProductQuotaDimensionsResponseBody struct {
	// The maximum number of entries returned for this query.
	//
	// example:
	//
	// 5
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that indicates the position of the returned data. If this parameter is not returned, it indicates that all data has been queried.
	//
	// example:
	//
	// 84
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The quota dimensions.
	QuotaDimensions []*ListProductQuotaDimensionsResponseBodyQuotaDimensions `json:"QuotaDimensions,omitempty" xml:"QuotaDimensions,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 7ED584FB-ECBF-4A2A-969D-F54D25EFABF9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned for this request.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProductQuotaDimensionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProductQuotaDimensionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductQuotaDimensionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProductQuotaDimensionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProductQuotaDimensionsResponseBody) GetQuotaDimensions() []*ListProductQuotaDimensionsResponseBodyQuotaDimensions {
	return s.QuotaDimensions
}

func (s *ListProductQuotaDimensionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProductQuotaDimensionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListProductQuotaDimensionsResponseBody) SetMaxResults(v int32) *ListProductQuotaDimensionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListProductQuotaDimensionsResponseBody) SetNextToken(v string) *ListProductQuotaDimensionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListProductQuotaDimensionsResponseBody) SetQuotaDimensions(v []*ListProductQuotaDimensionsResponseBodyQuotaDimensions) *ListProductQuotaDimensionsResponseBody {
	s.QuotaDimensions = v
	return s
}

func (s *ListProductQuotaDimensionsResponseBody) SetRequestId(v string) *ListProductQuotaDimensionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductQuotaDimensionsResponseBody) SetTotalCount(v int32) *ListProductQuotaDimensionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListProductQuotaDimensionsResponseBody) Validate() error {
	if s.QuotaDimensions != nil {
		for _, item := range s.QuotaDimensions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProductQuotaDimensionsResponseBodyQuotaDimensions struct {
	// The quota dimensions on which the target quota dimension depends.
	DependentDimensions []*string `json:"DependentDimensions,omitempty" xml:"DependentDimensions,omitempty" type:"Repeated"`
	// The key of the quota dimension. Valid values:
	//
	// - regionId: the region ID.
	//
	// - zoneId: the zone ID.
	//
	// - chargeType: the billing method.
	//
	// - networkType: the network type.
	//
	// - resourceType: the resource type.
	//
	// example:
	//
	// zoneId
	DimensionKey *string `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	// The details of the quota dimension value.
	DimensionValueDetail []*ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetail `json:"DimensionValueDetail,omitempty" xml:"DimensionValueDetail,omitempty" type:"Repeated"`
	// The list of quota dimension values.
	DimensionValues []*string `json:"DimensionValues,omitempty" xml:"DimensionValues,omitempty" type:"Repeated"`
	// The name of the quota dimension.
	//
	// example:
	//
	// Zone
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether the quota dimension is required when you query quota dimensions. Valid values:
	//
	// - true: The quota dimension is required.
	//
	// - false: The quota dimension is optional.
	//
	// example:
	//
	// false
	Requisite *bool `json:"Requisite,omitempty" xml:"Requisite,omitempty"`
}

func (s ListProductQuotaDimensionsResponseBodyQuotaDimensions) String() string {
	return dara.Prettify(s)
}

func (s ListProductQuotaDimensionsResponseBodyQuotaDimensions) GoString() string {
	return s.String()
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensions) GetDependentDimensions() []*string {
	return s.DependentDimensions
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensions) GetDimensionKey() *string {
	return s.DimensionKey
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensions) GetDimensionValueDetail() []*ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetail {
	return s.DimensionValueDetail
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensions) GetDimensionValues() []*string {
	return s.DimensionValues
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensions) GetName() *string {
	return s.Name
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensions) GetRequisite() *bool {
	return s.Requisite
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensions) SetDependentDimensions(v []*string) *ListProductQuotaDimensionsResponseBodyQuotaDimensions {
	s.DependentDimensions = v
	return s
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensions) SetDimensionKey(v string) *ListProductQuotaDimensionsResponseBodyQuotaDimensions {
	s.DimensionKey = &v
	return s
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensions) SetDimensionValueDetail(v []*ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetail) *ListProductQuotaDimensionsResponseBodyQuotaDimensions {
	s.DimensionValueDetail = v
	return s
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensions) SetDimensionValues(v []*string) *ListProductQuotaDimensionsResponseBodyQuotaDimensions {
	s.DimensionValues = v
	return s
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensions) SetName(v string) *ListProductQuotaDimensionsResponseBodyQuotaDimensions {
	s.Name = &v
	return s
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensions) SetRequisite(v bool) *ListProductQuotaDimensionsResponseBodyQuotaDimensions {
	s.Requisite = &v
	return s
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensions) Validate() error {
	if s.DimensionValueDetail != nil {
		for _, item := range s.DimensionValueDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetail struct {
	// The quota dimensions on which the target quota dimension depends.
	DependentDimensions []*ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetailDependentDimensions `json:"DependentDimensions,omitempty" xml:"DependentDimensions,omitempty" type:"Repeated"`
	// The name of the quota dimension value.
	//
	// example:
	//
	// cn-hangzhou-b
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the quota dimension.
	//
	// example:
	//
	// cn-hangzhou-b
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetail) String() string {
	return dara.Prettify(s)
}

func (s ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetail) GoString() string {
	return s.String()
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetail) GetDependentDimensions() []*ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetailDependentDimensions {
	return s.DependentDimensions
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetail) GetName() *string {
	return s.Name
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetail) GetValue() *string {
	return s.Value
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetail) SetDependentDimensions(v []*ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetailDependentDimensions) *ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetail {
	s.DependentDimensions = v
	return s
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetail) SetName(v string) *ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetail {
	s.Name = &v
	return s
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetail) SetValue(v string) *ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetail {
	s.Value = &v
	return s
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetail) Validate() error {
	if s.DependentDimensions != nil {
		for _, item := range s.DependentDimensions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetailDependentDimensions struct {
	// The key of the quota dimension on which the target quota dimension depends.
	//
	// example:
	//
	// regionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the quota dimension on which the target quota dimension depends.
	//
	// example:
	//
	// cn-hangzhou
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetailDependentDimensions) String() string {
	return dara.Prettify(s)
}

func (s ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetailDependentDimensions) GoString() string {
	return s.String()
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetailDependentDimensions) GetKey() *string {
	return s.Key
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetailDependentDimensions) GetValue() *string {
	return s.Value
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetailDependentDimensions) SetKey(v string) *ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetailDependentDimensions {
	s.Key = &v
	return s
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetailDependentDimensions) SetValue(v string) *ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetailDependentDimensions {
	s.Value = &v
	return s
}

func (s *ListProductQuotaDimensionsResponseBodyQuotaDimensionsDimensionValueDetailDependentDimensions) Validate() error {
	return dara.Validate(s)
}
