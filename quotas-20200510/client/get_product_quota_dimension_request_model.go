// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProductQuotaDimensionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDependentDimensions(v []*GetProductQuotaDimensionRequestDependentDimensions) *GetProductQuotaDimensionRequest
	GetDependentDimensions() []*GetProductQuotaDimensionRequestDependentDimensions
	SetDimensionKey(v string) *GetProductQuotaDimensionRequest
	GetDimensionKey() *string
	SetProductCode(v string) *GetProductQuotaDimensionRequest
	GetProductCode() *string
}

type GetProductQuotaDimensionRequest struct {
	// The information about quota dimensions.
	DependentDimensions []*GetProductQuotaDimensionRequestDependentDimensions `json:"DependentDimensions,omitempty" xml:"DependentDimensions,omitempty" type:"Repeated"`
	// The key of the quota dimension.
	//
	// example:
	//
	// regionId
	DimensionKey *string `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// > For more information, see [Alibaba Cloud services that support Quota Center](https://help.aliyun.com/document_detail/182368.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs-spec
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s GetProductQuotaDimensionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProductQuotaDimensionRequest) GoString() string {
	return s.String()
}

func (s *GetProductQuotaDimensionRequest) GetDependentDimensions() []*GetProductQuotaDimensionRequestDependentDimensions {
	return s.DependentDimensions
}

func (s *GetProductQuotaDimensionRequest) GetDimensionKey() *string {
	return s.DimensionKey
}

func (s *GetProductQuotaDimensionRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *GetProductQuotaDimensionRequest) SetDependentDimensions(v []*GetProductQuotaDimensionRequestDependentDimensions) *GetProductQuotaDimensionRequest {
	s.DependentDimensions = v
	return s
}

func (s *GetProductQuotaDimensionRequest) SetDimensionKey(v string) *GetProductQuotaDimensionRequest {
	s.DimensionKey = &v
	return s
}

func (s *GetProductQuotaDimensionRequest) SetProductCode(v string) *GetProductQuotaDimensionRequest {
	s.ProductCode = &v
	return s
}

func (s *GetProductQuotaDimensionRequest) Validate() error {
	return dara.Validate(s)
}

type GetProductQuotaDimensionRequestDependentDimensions struct {
	// The key of the quota dimension on which the quota dimension that you want to query is dependent.
	//
	// > The value range of N varies based on the number of quota dimensions that are supported by the specified Alibaba Cloud service.
	//
	// example:
	//
	// regionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the quota dimension on which the quota dimension that you want to query is dependent.
	//
	// > The value range of N varies based on the number of quota dimensions that are supported by the specified Alibaba Cloud service.
	//
	// example:
	//
	// cn-hangzhou
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetProductQuotaDimensionRequestDependentDimensions) String() string {
	return dara.Prettify(s)
}

func (s GetProductQuotaDimensionRequestDependentDimensions) GoString() string {
	return s.String()
}

func (s *GetProductQuotaDimensionRequestDependentDimensions) GetKey() *string {
	return s.Key
}

func (s *GetProductQuotaDimensionRequestDependentDimensions) GetValue() *string {
	return s.Value
}

func (s *GetProductQuotaDimensionRequestDependentDimensions) SetKey(v string) *GetProductQuotaDimensionRequestDependentDimensions {
	s.Key = &v
	return s
}

func (s *GetProductQuotaDimensionRequestDependentDimensions) SetValue(v string) *GetProductQuotaDimensionRequestDependentDimensions {
	s.Value = &v
	return s
}

func (s *GetProductQuotaDimensionRequestDependentDimensions) Validate() error {
	return dara.Validate(s)
}
