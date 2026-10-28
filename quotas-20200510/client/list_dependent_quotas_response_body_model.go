// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDependentQuotasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQuotas(v []*ListDependentQuotasResponseBodyQuotas) *ListDependentQuotasResponseBody
	GetQuotas() []*ListDependentQuotasResponseBodyQuotas
	SetRequestId(v string) *ListDependentQuotasResponseBody
	GetRequestId() *string
}

type ListDependentQuotasResponseBody struct {
	// The list of quotas on which the target quota depends.
	Quotas []*ListDependentQuotasResponseBodyQuotas `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 920D8A47-26BB-49FA-A09F-F98D7DAA55F3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDependentQuotasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDependentQuotasResponseBody) GoString() string {
	return s.String()
}

func (s *ListDependentQuotasResponseBody) GetQuotas() []*ListDependentQuotasResponseBodyQuotas {
	return s.Quotas
}

func (s *ListDependentQuotasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDependentQuotasResponseBody) SetQuotas(v []*ListDependentQuotasResponseBodyQuotas) *ListDependentQuotasResponseBody {
	s.Quotas = v
	return s
}

func (s *ListDependentQuotasResponseBody) SetRequestId(v string) *ListDependentQuotasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDependentQuotasResponseBody) Validate() error {
	if s.Quotas != nil {
		for _, item := range s.Quotas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDependentQuotasResponseBodyQuotas struct {
	// The quota dimensions on which the target quota depends.
	Dimensions []*ListDependentQuotasResponseBodyQuotasDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The abbreviated name of the cloud service.
	//
	// example:
	//
	// ram
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota on which the target quota depends.
	//
	// example:
	//
	// q_laxtn4
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The ratio between the target quota and the dependent quota.
	//
	// example:
	//
	// 100
	Scale *float32 `json:"Scale,omitempty" xml:"Scale,omitempty"`
}

func (s ListDependentQuotasResponseBodyQuotas) String() string {
	return dara.Prettify(s)
}

func (s ListDependentQuotasResponseBodyQuotas) GoString() string {
	return s.String()
}

func (s *ListDependentQuotasResponseBodyQuotas) GetDimensions() []*ListDependentQuotasResponseBodyQuotasDimensions {
	return s.Dimensions
}

func (s *ListDependentQuotasResponseBodyQuotas) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListDependentQuotasResponseBodyQuotas) GetQuotaActionCode() *string {
	return s.QuotaActionCode
}

func (s *ListDependentQuotasResponseBodyQuotas) GetScale() *float32 {
	return s.Scale
}

func (s *ListDependentQuotasResponseBodyQuotas) SetDimensions(v []*ListDependentQuotasResponseBodyQuotasDimensions) *ListDependentQuotasResponseBodyQuotas {
	s.Dimensions = v
	return s
}

func (s *ListDependentQuotasResponseBodyQuotas) SetProductCode(v string) *ListDependentQuotasResponseBodyQuotas {
	s.ProductCode = &v
	return s
}

func (s *ListDependentQuotasResponseBodyQuotas) SetQuotaActionCode(v string) *ListDependentQuotasResponseBodyQuotas {
	s.QuotaActionCode = &v
	return s
}

func (s *ListDependentQuotasResponseBodyQuotas) SetScale(v float32) *ListDependentQuotasResponseBodyQuotas {
	s.Scale = &v
	return s
}

func (s *ListDependentQuotasResponseBodyQuotas) Validate() error {
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

type ListDependentQuotasResponseBodyQuotasDimensions struct {
	// The quota dimension on which the dependent dimension of the target quota depends.
	DependentDimension []*string `json:"DependentDimension,omitempty" xml:"DependentDimension,omitempty" type:"Repeated"`
	// The key of the quota dimension on which the target quota depends.
	//
	// example:
	//
	// []
	DimensionKey *string `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	// The value of the quota dimension on which the dependent dimension of the target quota depends.
	DimensionValues []*string `json:"DimensionValues,omitempty" xml:"DimensionValues,omitempty" type:"Repeated"`
}

func (s ListDependentQuotasResponseBodyQuotasDimensions) String() string {
	return dara.Prettify(s)
}

func (s ListDependentQuotasResponseBodyQuotasDimensions) GoString() string {
	return s.String()
}

func (s *ListDependentQuotasResponseBodyQuotasDimensions) GetDependentDimension() []*string {
	return s.DependentDimension
}

func (s *ListDependentQuotasResponseBodyQuotasDimensions) GetDimensionKey() *string {
	return s.DimensionKey
}

func (s *ListDependentQuotasResponseBodyQuotasDimensions) GetDimensionValues() []*string {
	return s.DimensionValues
}

func (s *ListDependentQuotasResponseBodyQuotasDimensions) SetDependentDimension(v []*string) *ListDependentQuotasResponseBodyQuotasDimensions {
	s.DependentDimension = v
	return s
}

func (s *ListDependentQuotasResponseBodyQuotasDimensions) SetDimensionKey(v string) *ListDependentQuotasResponseBodyQuotasDimensions {
	s.DimensionKey = &v
	return s
}

func (s *ListDependentQuotasResponseBodyQuotasDimensions) SetDimensionValues(v []*string) *ListDependentQuotasResponseBodyQuotasDimensions {
	s.DimensionValues = v
	return s
}

func (s *ListDependentQuotasResponseBodyQuotasDimensions) Validate() error {
	return dara.Validate(s)
}
