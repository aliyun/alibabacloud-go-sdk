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
	// The properties of the quota item. The combination of dimensions and QuotaActionCode determines a unique quota item.
	//
	// > Dimensions are required for some cloud services. You can call the [ListProductQuotaDimensions](~~ListProductQuotaDimensions~~) operation to query the quota dimensions of a cloud service and check whether a dimension is required. The Requisite response parameter indicates whether the dimension is required.
	//
	// example:
	//
	// {\\"regionId\\":\\"cn-beijing\\"}
	Dimensions []*GetProductQuotaRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The abbreviated name of the cloud service.
	//
	// > For more information about how to obtain the abbreviated name of a cloud service, see the ProductCode parameter in [ListProducts](https://help.aliyun.com/document_detail/440555.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// > For more information about how to obtain the quota ID of a specified cloud service, see the QuotaActionCode parameter in [ListProductQuotas](https://help.aliyun.com/document_detail/440554.html).
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

type GetProductQuotaRequestDimensions struct {
	// The key of the quota dimension.
	//
	// > You must specify both Dimensions.N.Key and Dimensions.N.Value. The value range of N depends on the number of dimensions supported by the cloud service. You can call the [ListProductQuotaDimensions](~~ListProductQuotaDimensions~~) operation to query the quota dimensions supported by the cloud service. The length of the returned array is the maximum value of N.
	//
	// example:
	//
	// regionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the quota dimension.
	//
	// > You must specify both Dimensions.N.Key and Dimensions.N.Value. The value range of N depends on the number of dimensions supported by the cloud service. You can call the [ListProductQuotaDimensions](~~ListProductQuotaDimensions~~) operation to query the quota dimensions supported by the cloud service. The length of the returned array is the maximum value of N.
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
