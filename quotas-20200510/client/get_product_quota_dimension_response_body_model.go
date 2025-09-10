// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProductQuotaDimensionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQuotaDimension(v *GetProductQuotaDimensionResponseBodyQuotaDimension) *GetProductQuotaDimensionResponseBody
	GetQuotaDimension() *GetProductQuotaDimensionResponseBodyQuotaDimension
	SetRequestId(v string) *GetProductQuotaDimensionResponseBody
	GetRequestId() *string
}

type GetProductQuotaDimensionResponseBody struct {
	// The details of the quota dimension.
	QuotaDimension *GetProductQuotaDimensionResponseBodyQuotaDimension `json:"QuotaDimension,omitempty" xml:"QuotaDimension,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 1FA5F0E2-368E-4BA4-A8D0-6060FC6BB8F3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetProductQuotaDimensionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProductQuotaDimensionResponseBody) GoString() string {
	return s.String()
}

func (s *GetProductQuotaDimensionResponseBody) GetQuotaDimension() *GetProductQuotaDimensionResponseBodyQuotaDimension {
	return s.QuotaDimension
}

func (s *GetProductQuotaDimensionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProductQuotaDimensionResponseBody) SetQuotaDimension(v *GetProductQuotaDimensionResponseBodyQuotaDimension) *GetProductQuotaDimensionResponseBody {
	s.QuotaDimension = v
	return s
}

func (s *GetProductQuotaDimensionResponseBody) SetRequestId(v string) *GetProductQuotaDimensionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProductQuotaDimensionResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetProductQuotaDimensionResponseBodyQuotaDimension struct {
	// The quota dimensions on which the quota dimension that you want to query is dependent.
	DependentDimensions []*string `json:"DependentDimensions,omitempty" xml:"DependentDimensions,omitempty" type:"Repeated"`
	// The key of the quota dimension. Valid values:
	//
	// 	- regionId: the region ID.
	//
	// 	- zoneId: the zone ID.
	//
	// 	- chargeType: the billing method.
	//
	// 	- networkType: the network type.
	//
	// example:
	//
	// regionId
	DimensionKey *string `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	// The details of the quota dimension value.
	DimensionValueDetail []*GetProductQuotaDimensionResponseBodyQuotaDimensionDimensionValueDetail `json:"DimensionValueDetail,omitempty" xml:"DimensionValueDetail,omitempty" type:"Repeated"`
	// The values of the quota dimension.
	DimensionValues []*string `json:"DimensionValues,omitempty" xml:"DimensionValues,omitempty" type:"Repeated"`
	// The name of the quota dimension.
	//
	// example:
	//
	// region
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetProductQuotaDimensionResponseBodyQuotaDimension) String() string {
	return dara.Prettify(s)
}

func (s GetProductQuotaDimensionResponseBodyQuotaDimension) GoString() string {
	return s.String()
}

func (s *GetProductQuotaDimensionResponseBodyQuotaDimension) GetDependentDimensions() []*string {
	return s.DependentDimensions
}

func (s *GetProductQuotaDimensionResponseBodyQuotaDimension) GetDimensionKey() *string {
	return s.DimensionKey
}

func (s *GetProductQuotaDimensionResponseBodyQuotaDimension) GetDimensionValueDetail() []*GetProductQuotaDimensionResponseBodyQuotaDimensionDimensionValueDetail {
	return s.DimensionValueDetail
}

func (s *GetProductQuotaDimensionResponseBodyQuotaDimension) GetDimensionValues() []*string {
	return s.DimensionValues
}

func (s *GetProductQuotaDimensionResponseBodyQuotaDimension) GetName() *string {
	return s.Name
}

func (s *GetProductQuotaDimensionResponseBodyQuotaDimension) SetDependentDimensions(v []*string) *GetProductQuotaDimensionResponseBodyQuotaDimension {
	s.DependentDimensions = v
	return s
}

func (s *GetProductQuotaDimensionResponseBodyQuotaDimension) SetDimensionKey(v string) *GetProductQuotaDimensionResponseBodyQuotaDimension {
	s.DimensionKey = &v
	return s
}

func (s *GetProductQuotaDimensionResponseBodyQuotaDimension) SetDimensionValueDetail(v []*GetProductQuotaDimensionResponseBodyQuotaDimensionDimensionValueDetail) *GetProductQuotaDimensionResponseBodyQuotaDimension {
	s.DimensionValueDetail = v
	return s
}

func (s *GetProductQuotaDimensionResponseBodyQuotaDimension) SetDimensionValues(v []*string) *GetProductQuotaDimensionResponseBodyQuotaDimension {
	s.DimensionValues = v
	return s
}

func (s *GetProductQuotaDimensionResponseBodyQuotaDimension) SetName(v string) *GetProductQuotaDimensionResponseBodyQuotaDimension {
	s.Name = &v
	return s
}

func (s *GetProductQuotaDimensionResponseBodyQuotaDimension) Validate() error {
	return dara.Validate(s)
}

type GetProductQuotaDimensionResponseBodyQuotaDimensionDimensionValueDetail struct {
	// The name of the quota dimension.
	//
	// example:
	//
	// cn-hangzhou
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the quota dimension.
	//
	// example:
	//
	// cn-hangzhou
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetProductQuotaDimensionResponseBodyQuotaDimensionDimensionValueDetail) String() string {
	return dara.Prettify(s)
}

func (s GetProductQuotaDimensionResponseBodyQuotaDimensionDimensionValueDetail) GoString() string {
	return s.String()
}

func (s *GetProductQuotaDimensionResponseBodyQuotaDimensionDimensionValueDetail) GetName() *string {
	return s.Name
}

func (s *GetProductQuotaDimensionResponseBodyQuotaDimensionDimensionValueDetail) GetValue() *string {
	return s.Value
}

func (s *GetProductQuotaDimensionResponseBodyQuotaDimensionDimensionValueDetail) SetName(v string) *GetProductQuotaDimensionResponseBodyQuotaDimensionDimensionValueDetail {
	s.Name = &v
	return s
}

func (s *GetProductQuotaDimensionResponseBodyQuotaDimensionDimensionValueDetail) SetValue(v string) *GetProductQuotaDimensionResponseBodyQuotaDimensionDimensionValueDetail {
	s.Value = &v
	return s
}

func (s *GetProductQuotaDimensionResponseBodyQuotaDimensionDimensionValueDetail) Validate() error {
	return dara.Validate(s)
}
