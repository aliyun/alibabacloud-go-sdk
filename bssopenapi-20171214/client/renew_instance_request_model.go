// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RenewInstanceRequest
	GetClientToken() *string
	SetInstanceId(v string) *RenewInstanceRequest
	GetInstanceId() *string
	SetOwnerId(v int64) *RenewInstanceRequest
	GetOwnerId() *int64
	SetProductCode(v string) *RenewInstanceRequest
	GetProductCode() *string
	SetProductType(v string) *RenewInstanceRequest
	GetProductType() *string
	SetRenewPeriod(v int32) *RenewInstanceRequest
	GetRenewPeriod() *int32
}

type RenewInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// ASKJHKLASJHAFSLKJH
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the instance
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-skjdhaskjdh
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The code of the service to which the instance belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// rds
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The type of the service.
	//
	// example:
	//
	// rds
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The duration of the subscription renewal. Unit: months. Valid values:
	//
	// 	- 1 to 9
	//
	// 	- 12
	//
	// 	- 24
	//
	// 	- 36
	//
	// This parameter is required.
	//
	// example:
	//
	// 6
	RenewPeriod *int32 `json:"RenewPeriod,omitempty" xml:"RenewPeriod,omitempty"`
}

func (s RenewInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RenewInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RenewInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RenewInstanceRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *RenewInstanceRequest) GetProductType() *string {
	return s.ProductType
}

func (s *RenewInstanceRequest) GetRenewPeriod() *int32 {
	return s.RenewPeriod
}

func (s *RenewInstanceRequest) SetClientToken(v string) *RenewInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RenewInstanceRequest) SetInstanceId(v string) *RenewInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RenewInstanceRequest) SetOwnerId(v int64) *RenewInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *RenewInstanceRequest) SetProductCode(v string) *RenewInstanceRequest {
	s.ProductCode = &v
	return s
}

func (s *RenewInstanceRequest) SetProductType(v string) *RenewInstanceRequest {
	s.ProductType = &v
	return s
}

func (s *RenewInstanceRequest) SetRenewPeriod(v int32) *RenewInstanceRequest {
	s.RenewPeriod = &v
	return s
}

func (s *RenewInstanceRequest) Validate() error {
	return dara.Validate(s)
}
