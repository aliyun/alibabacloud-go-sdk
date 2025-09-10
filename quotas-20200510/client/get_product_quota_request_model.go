// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProductQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDimensions(v []*GetProductQuotaRequestDimensions) *GetProductQuotaRequest
	GetDimensions() []*GetProductQuotaRequestDimensions
	SetProductCode(v string) *GetProductQuotaRequest
	GetProductCode() *string
	SetQuotaActionCode(v string) *GetProductQuotaRequest
	GetQuotaActionCode() *string
}

type GetProductQuotaRequest struct {
	// The quota dimensions.
	//
	// example:
	//
	// {\\"regionId\\":\\"cn-beijing\\"}
	Dimensions []*GetProductQuotaRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// >  To query the abbreviation of an Alibaba Cloud service name, call the [ListProducts](https://help.aliyun.com/document_detail/440555.html) operation and check the value of `ProductCode` in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// >  To query a quota ID of an Alibaba Cloud service, call the [ListProductQuotas](https://help.aliyun.com/document_detail/440554.html) operation and check the value of `QuotaActionCode` in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// q_security-groups
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
}

func (s GetProductQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProductQuotaRequest) GoString() string {
	return s.String()
}

func (s *GetProductQuotaRequest) GetDimensions() []*GetProductQuotaRequestDimensions {
	return s.Dimensions
}

func (s *GetProductQuotaRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *GetProductQuotaRequest) GetQuotaActionCode() *string {
	return s.QuotaActionCode
}

func (s *GetProductQuotaRequest) SetDimensions(v []*GetProductQuotaRequestDimensions) *GetProductQuotaRequest {
	s.Dimensions = v
	return s
}

func (s *GetProductQuotaRequest) SetProductCode(v string) *GetProductQuotaRequest {
	s.ProductCode = &v
	return s
}

func (s *GetProductQuotaRequest) SetQuotaActionCode(v string) *GetProductQuotaRequest {
	s.QuotaActionCode = &v
	return s
}

func (s *GetProductQuotaRequest) Validate() error {
	return dara.Validate(s)
}

type GetProductQuotaRequestDimensions struct {
	// The key of the dimension.
	//
	// >  This parameter is required for cloud services that support dimensions. You must specify both `Dimensions.N.Key` and `Dimensions.N.Value`. The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service. The following cloud services support dimensions: ECS whose service code is ecs, Enterprise Distributed Application Service (EDAS) whose service code is edas, ECS Quotas by Instance Type whose service code is ecs-spec, and Auto Scaling whose service code is ess.
	//
	// example:
	//
	// regionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the dimension.
	//
	// >  This parameter is required for cloud services that support dimensions. You must specify both `Dimensions.N.Key` and `Dimensions.N.Value`. The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service. The following cloud services support dimensions: ECS whose service code is ecs, EDAS whose service code is edas, ECS Quotas by Instance Type whose service code is ecs-spec, and Auto Scaling whose service code is ess.
	//
	// example:
	//
	// cn-hangzhou
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetProductQuotaRequestDimensions) String() string {
	return dara.Prettify(s)
}

func (s GetProductQuotaRequestDimensions) GoString() string {
	return s.String()
}

func (s *GetProductQuotaRequestDimensions) GetKey() *string {
	return s.Key
}

func (s *GetProductQuotaRequestDimensions) GetValue() *string {
	return s.Value
}

func (s *GetProductQuotaRequestDimensions) SetKey(v string) *GetProductQuotaRequestDimensions {
	s.Key = &v
	return s
}

func (s *GetProductQuotaRequestDimensions) SetValue(v string) *GetProductQuotaRequestDimensions {
	s.Value = &v
	return s
}

func (s *GetProductQuotaRequestDimensions) Validate() error {
	return dara.Validate(s)
}
