// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewChangeInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RenewChangeInstanceRequest
	GetClientToken() *string
	SetInstanceId(v string) *RenewChangeInstanceRequest
	GetInstanceId() *string
	SetOwnerId(v int64) *RenewChangeInstanceRequest
	GetOwnerId() *int64
	SetParameter(v []*RenewChangeInstanceRequestParameter) *RenewChangeInstanceRequest
	GetParameter() []*RenewChangeInstanceRequestParameter
	SetProductCode(v string) *RenewChangeInstanceRequest
	GetProductCode() *string
	SetProductType(v string) *RenewChangeInstanceRequest
	GetProductType() *string
	SetRenewPeriod(v int64) *RenewChangeInstanceRequest
	GetRenewPeriod() *int64
}

type RenewChangeInstanceRequest struct {
	// Ensures the idempotence of the request.
	//
	// example:
	//
	// JASIOFKVNVIXXXXXX
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-skjdhaskjdh
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Module list
	Parameter []*RenewChangeInstanceRequestParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Repeated"`
	// The product code. To obtain it, call **QueryProductList**.
	//
	// This parameter is required.
	//
	// example:
	//
	// rds
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// Product type
	//
	// example:
	//
	// rds
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The subscription renewal duration, in months. Valid values:
	//
	// - 1-9
	//
	// - 12
	//
	// - 24
	//
	// - 36
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	RenewPeriod *int64 `json:"RenewPeriod,omitempty" xml:"RenewPeriod,omitempty"`
}

func (s RenewChangeInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewChangeInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewChangeInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RenewChangeInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RenewChangeInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RenewChangeInstanceRequest) GetParameter() []*RenewChangeInstanceRequestParameter {
	return s.Parameter
}

func (s *RenewChangeInstanceRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *RenewChangeInstanceRequest) GetProductType() *string {
	return s.ProductType
}

func (s *RenewChangeInstanceRequest) GetRenewPeriod() *int64 {
	return s.RenewPeriod
}

func (s *RenewChangeInstanceRequest) SetClientToken(v string) *RenewChangeInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RenewChangeInstanceRequest) SetInstanceId(v string) *RenewChangeInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RenewChangeInstanceRequest) SetOwnerId(v int64) *RenewChangeInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *RenewChangeInstanceRequest) SetParameter(v []*RenewChangeInstanceRequestParameter) *RenewChangeInstanceRequest {
	s.Parameter = v
	return s
}

func (s *RenewChangeInstanceRequest) SetProductCode(v string) *RenewChangeInstanceRequest {
	s.ProductCode = &v
	return s
}

func (s *RenewChangeInstanceRequest) SetProductType(v string) *RenewChangeInstanceRequest {
	s.ProductType = &v
	return s
}

func (s *RenewChangeInstanceRequest) SetRenewPeriod(v int64) *RenewChangeInstanceRequest {
	s.RenewPeriod = &v
	return s
}

func (s *RenewChangeInstanceRequest) Validate() error {
	if s.Parameter != nil {
		for _, item := range s.Parameter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RenewChangeInstanceRequestParameter struct {
	// The code configuration for the Nth module property. The value of N can be an integer from 1 to 100. If you have multiple module properties, create a separate parameter for each by setting n to 1, 2, 3, and so on.
	//
	// This parameter is required.
	//
	// example:
	//
	// Bandwidth
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The value configuration for the nth module property. The value of n ranges from 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// disk
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s RenewChangeInstanceRequestParameter) String() string {
	return dara.Prettify(s)
}

func (s RenewChangeInstanceRequestParameter) GoString() string {
	return s.String()
}

func (s *RenewChangeInstanceRequestParameter) GetCode() *string {
	return s.Code
}

func (s *RenewChangeInstanceRequestParameter) GetValue() *string {
	return s.Value
}

func (s *RenewChangeInstanceRequestParameter) SetCode(v string) *RenewChangeInstanceRequestParameter {
	s.Code = &v
	return s
}

func (s *RenewChangeInstanceRequestParameter) SetValue(v string) *RenewChangeInstanceRequestParameter {
	s.Value = &v
	return s
}

func (s *RenewChangeInstanceRequestParameter) Validate() error {
	return dara.Validate(s)
}
