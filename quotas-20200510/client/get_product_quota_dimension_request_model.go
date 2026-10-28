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
	// The quota dimensions on which the target quota dimension depends.
	DependentDimensions []*GetProductQuotaDimensionRequestDependentDimensions `json:"DependentDimensions,omitempty" xml:"DependentDimensions,omitempty" type:"Repeated"`
	// The key of the target quota dimension.
	//
	// > If the target quota dimension has dependent quota dimensions, you must specify the keys and values of the dependent quota dimensions at the same time.
	//
	// example:
	//
	// zoneId
	DimensionKey *string `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	// The abbreviation of the cloud service.
	//
	// > For more information about how to obtain the abbreviation of a cloud service, see the ProductCode parameter in [ListProducts](https://help.aliyun.com/document_detail/440555.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// disk
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

type GetProductQuotaDimensionRequestDependentDimensions struct {
	// The key of the quota dimension on which the target quota dimension depends.
	//
	// > - The value of N depends on the number of quota dimensions on which the target quota dimension depends.
	//
	// > - You must specify the key and value of a dependent quota dimension at the same time.
	//
	// example:
	//
	// regionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the quota dimension on which the target quota dimension depends.
	//
	// > - The value of N depends on the number of quota dimensions on which the target quota dimension depends.
	//
	// > - You must specify the key and value of a dependent quota dimension at the same time.
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
