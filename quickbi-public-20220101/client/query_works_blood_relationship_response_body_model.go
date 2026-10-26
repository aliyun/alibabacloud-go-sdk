// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryWorksBloodRelationshipResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryWorksBloodRelationshipResponseBody
	GetRequestId() *string
	SetResult(v []*QueryWorksBloodRelationshipResponseBodyResult) *QueryWorksBloodRelationshipResponseBody
	GetResult() []*QueryWorksBloodRelationshipResponseBodyResult
	SetSuccess(v bool) *QueryWorksBloodRelationshipResponseBody
	GetSuccess() *bool
}

type QueryWorksBloodRelationshipResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// List of work blood information.
	Result []*QueryWorksBloodRelationshipResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryWorksBloodRelationshipResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryWorksBloodRelationshipResponseBody) GoString() string {
	return s.String()
}

func (s *QueryWorksBloodRelationshipResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryWorksBloodRelationshipResponseBody) GetResult() []*QueryWorksBloodRelationshipResponseBodyResult {
	return s.Result
}

func (s *QueryWorksBloodRelationshipResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryWorksBloodRelationshipResponseBody) SetRequestId(v string) *QueryWorksBloodRelationshipResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryWorksBloodRelationshipResponseBody) SetResult(v []*QueryWorksBloodRelationshipResponseBodyResult) *QueryWorksBloodRelationshipResponseBody {
	s.Result = v
	return s
}

func (s *QueryWorksBloodRelationshipResponseBody) SetSuccess(v bool) *QueryWorksBloodRelationshipResponseBody {
	s.Success = &v
	return s
}

func (s *QueryWorksBloodRelationshipResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryWorksBloodRelationshipResponseBodyResult struct {
	// The ID of the component that you want to modify.
	//
	// example:
	//
	// 0696083a-ca72-4d89-8e7a-c017910e0***
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	// The name of the component.
	//
	// example:
	//
	// Line
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// The type of the image component.
	//
	// example:
	//
	// 3
	ComponentType *int32 `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	// Chinese name of the component type
	//
	// example:
	//
	// ddd
	ComponentTypeCnName *string `json:"ComponentTypeCnName,omitempty" xml:"ComponentTypeCnName,omitempty"`
	// The name of the component type.
	//
	// example:
	//
	// LINE
	ComponentTypeName *string `json:"ComponentTypeName,omitempty" xml:"ComponentTypeName,omitempty"`
	// The ID of the training dataset that you want to remove from the specified custom linguistic model.
	//
	// example:
	//
	// dc78a4ed-880d-452e-b017-90cfc10c83e5_company_sales_record
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// A list of query parameter reference columns.
	QueryParams []*QueryWorksBloodRelationshipResponseBodyResultQueryParams `json:"QueryParams,omitempty" xml:"QueryParams,omitempty" type:"Repeated"`
}

func (s QueryWorksBloodRelationshipResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryWorksBloodRelationshipResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryWorksBloodRelationshipResponseBodyResult) GetComponentId() *string {
	return s.ComponentId
}

func (s *QueryWorksBloodRelationshipResponseBodyResult) GetComponentName() *string {
	return s.ComponentName
}

func (s *QueryWorksBloodRelationshipResponseBodyResult) GetComponentType() *int32 {
	return s.ComponentType
}

func (s *QueryWorksBloodRelationshipResponseBodyResult) GetComponentTypeCnName() *string {
	return s.ComponentTypeCnName
}

func (s *QueryWorksBloodRelationshipResponseBodyResult) GetComponentTypeName() *string {
	return s.ComponentTypeName
}

func (s *QueryWorksBloodRelationshipResponseBodyResult) GetDatasetId() *string {
	return s.DatasetId
}

func (s *QueryWorksBloodRelationshipResponseBodyResult) GetQueryParams() []*QueryWorksBloodRelationshipResponseBodyResultQueryParams {
	return s.QueryParams
}

func (s *QueryWorksBloodRelationshipResponseBodyResult) SetComponentId(v string) *QueryWorksBloodRelationshipResponseBodyResult {
	s.ComponentId = &v
	return s
}

func (s *QueryWorksBloodRelationshipResponseBodyResult) SetComponentName(v string) *QueryWorksBloodRelationshipResponseBodyResult {
	s.ComponentName = &v
	return s
}

func (s *QueryWorksBloodRelationshipResponseBodyResult) SetComponentType(v int32) *QueryWorksBloodRelationshipResponseBodyResult {
	s.ComponentType = &v
	return s
}

func (s *QueryWorksBloodRelationshipResponseBodyResult) SetComponentTypeCnName(v string) *QueryWorksBloodRelationshipResponseBodyResult {
	s.ComponentTypeCnName = &v
	return s
}

func (s *QueryWorksBloodRelationshipResponseBodyResult) SetComponentTypeName(v string) *QueryWorksBloodRelationshipResponseBodyResult {
	s.ComponentTypeName = &v
	return s
}

func (s *QueryWorksBloodRelationshipResponseBodyResult) SetDatasetId(v string) *QueryWorksBloodRelationshipResponseBodyResult {
	s.DatasetId = &v
	return s
}

func (s *QueryWorksBloodRelationshipResponseBodyResult) SetQueryParams(v []*QueryWorksBloodRelationshipResponseBodyResultQueryParams) *QueryWorksBloodRelationshipResponseBodyResult {
	s.QueryParams = v
	return s
}

func (s *QueryWorksBloodRelationshipResponseBodyResult) Validate() error {
	if s.QueryParams != nil {
		for _, item := range s.QueryParams {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryWorksBloodRelationshipResponseBodyResultQueryParams struct {
	// The ID of the owning location.
	//
	// example:
	//
	// area_column
	AreaId *string `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	// The name of the owning location.
	//
	// example:
	//
	// Column (Measure)
	AreaName *string `json:"AreaName,omitempty" xml:"AreaName,omitempty"`
	// The display name of the field.
	//
	// example:
	//
	// order_number
	Caption *string `json:"Caption,omitempty" xml:"Caption,omitempty"`
	// The type of the field. Valid values:
	//
	// 	- string: string type
	//
	// 	- date: a date type that contains only the year, month, and day parts
	//
	// 	- datetime: a common date type
	//
	// 	- time: a date type that contains only hours, minutes, and seconds.
	//
	// 	- number: numeric
	//
	// 	- boolean: Boolean type
	//
	// 	- geographical: geographical location
	//
	// 	- url: string type
	//
	// 	- imageUrl: the type of the image link.
	//
	// 	- multivalue: a multi-value column
	//
	// example:
	//
	// number
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// Calculate field expression.
	//
	// example:
	//
	// BI_DATEADD([date], 100, \\"day\\")
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// Indices whether the metric. Valid values:
	//
	// true false
	//
	// example:
	//
	// true
	IsMeasure *bool `json:"IsMeasure,omitempty" xml:"IsMeasure,omitempty"`
	// The globally unique PathId.
	//
	// example:
	//
	// schema7d1944eb-443e-48c6-8123-bf45a99e7e74.dc78a4ed-880d-452e-b017-90cfc10c83e5_company_sales_record.[Ndc78a4_order_level].[Ndc78a4_order_level].[Ndc78a4_order_level]
	PathId *string `json:"PathId,omitempty" xml:"PathId,omitempty"`
	// The unique ID of the field.
	//
	// example:
	//
	// Ndc78a4_order_number
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s QueryWorksBloodRelationshipResponseBodyResultQueryParams) String() string {
	return dara.Prettify(s)
}

func (s QueryWorksBloodRelationshipResponseBodyResultQueryParams) GoString() string {
	return s.String()
}

func (s *QueryWorksBloodRelationshipResponseBodyResultQueryParams) GetAreaId() *string {
	return s.AreaId
}

func (s *QueryWorksBloodRelationshipResponseBodyResultQueryParams) GetAreaName() *string {
	return s.AreaName
}

func (s *QueryWorksBloodRelationshipResponseBodyResultQueryParams) GetCaption() *string {
	return s.Caption
}

func (s *QueryWorksBloodRelationshipResponseBodyResultQueryParams) GetDataType() *string {
	return s.DataType
}

func (s *QueryWorksBloodRelationshipResponseBodyResultQueryParams) GetExpression() *string {
	return s.Expression
}

func (s *QueryWorksBloodRelationshipResponseBodyResultQueryParams) GetIsMeasure() *bool {
	return s.IsMeasure
}

func (s *QueryWorksBloodRelationshipResponseBodyResultQueryParams) GetPathId() *string {
	return s.PathId
}

func (s *QueryWorksBloodRelationshipResponseBodyResultQueryParams) GetUid() *string {
	return s.Uid
}

func (s *QueryWorksBloodRelationshipResponseBodyResultQueryParams) SetAreaId(v string) *QueryWorksBloodRelationshipResponseBodyResultQueryParams {
	s.AreaId = &v
	return s
}

func (s *QueryWorksBloodRelationshipResponseBodyResultQueryParams) SetAreaName(v string) *QueryWorksBloodRelationshipResponseBodyResultQueryParams {
	s.AreaName = &v
	return s
}

func (s *QueryWorksBloodRelationshipResponseBodyResultQueryParams) SetCaption(v string) *QueryWorksBloodRelationshipResponseBodyResultQueryParams {
	s.Caption = &v
	return s
}

func (s *QueryWorksBloodRelationshipResponseBodyResultQueryParams) SetDataType(v string) *QueryWorksBloodRelationshipResponseBodyResultQueryParams {
	s.DataType = &v
	return s
}

func (s *QueryWorksBloodRelationshipResponseBodyResultQueryParams) SetExpression(v string) *QueryWorksBloodRelationshipResponseBodyResultQueryParams {
	s.Expression = &v
	return s
}

func (s *QueryWorksBloodRelationshipResponseBodyResultQueryParams) SetIsMeasure(v bool) *QueryWorksBloodRelationshipResponseBodyResultQueryParams {
	s.IsMeasure = &v
	return s
}

func (s *QueryWorksBloodRelationshipResponseBodyResultQueryParams) SetPathId(v string) *QueryWorksBloodRelationshipResponseBodyResultQueryParams {
	s.PathId = &v
	return s
}

func (s *QueryWorksBloodRelationshipResponseBodyResultQueryParams) SetUid(v string) *QueryWorksBloodRelationshipResponseBodyResultQueryParams {
	s.Uid = &v
	return s
}

func (s *QueryWorksBloodRelationshipResponseBodyResultQueryParams) Validate() error {
	return dara.Validate(s)
}
