// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReduceCreditSeatsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ReduceCreditSeatsRequest
	GetClientToken() *string
	SetInstanceId(v string) *ReduceCreditSeatsRequest
	GetInstanceId() *string
	SetProductCode(v string) *ReduceCreditSeatsRequest
	GetProductCode() *string
	SetProductType(v string) *ReduceCreditSeatsRequest
	GetProductType() *string
	SetSubscriptionType(v string) *ReduceCreditSeatsRequest
	GetSubscriptionType() *string
}

type ReduceCreditSeatsRequest struct {
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ProductCode      *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductType      *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
}

func (s ReduceCreditSeatsRequest) String() string {
	return dara.Prettify(s)
}

func (s ReduceCreditSeatsRequest) GoString() string {
	return s.String()
}

func (s *ReduceCreditSeatsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ReduceCreditSeatsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReduceCreditSeatsRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ReduceCreditSeatsRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ReduceCreditSeatsRequest) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *ReduceCreditSeatsRequest) SetClientToken(v string) *ReduceCreditSeatsRequest {
	s.ClientToken = &v
	return s
}

func (s *ReduceCreditSeatsRequest) SetInstanceId(v string) *ReduceCreditSeatsRequest {
	s.InstanceId = &v
	return s
}

func (s *ReduceCreditSeatsRequest) SetProductCode(v string) *ReduceCreditSeatsRequest {
	s.ProductCode = &v
	return s
}

func (s *ReduceCreditSeatsRequest) SetProductType(v string) *ReduceCreditSeatsRequest {
	s.ProductType = &v
	return s
}

func (s *ReduceCreditSeatsRequest) SetSubscriptionType(v string) *ReduceCreditSeatsRequest {
	s.SubscriptionType = &v
	return s
}

func (s *ReduceCreditSeatsRequest) Validate() error {
	return dara.Validate(s)
}
