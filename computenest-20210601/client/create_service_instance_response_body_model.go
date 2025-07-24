// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMarketInstanceId(v string) *CreateServiceInstanceResponseBody
	GetMarketInstanceId() *string
	SetOrderId(v string) *CreateServiceInstanceResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateServiceInstanceResponseBody
	GetRequestId() *string
	SetServiceInstanceId(v string) *CreateServiceInstanceResponseBody
	GetServiceInstanceId() *string
	SetStatus(v string) *CreateServiceInstanceResponseBody
	GetStatus() *string
}

type CreateServiceInstanceResponseBody struct {
	// The MartketInstance ID.
	//
	// example:
	//
	// 786***45
	MarketInstanceId *string `json:"MarketInstanceId,omitempty" xml:"MarketInstanceId,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 2306175xxxxxxxx
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service instance ID.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The status of the service instance. Valid values:
	//
	// 	- **Created**
	//
	// 	- **Deploying**
	//
	// 	- **DeployedFailed**
	//
	// 	- **Deployed**
	//
	// 	- **Upgrading**
	//
	// 	- **Deleting**
	//
	// 	- **Deleted**
	//
	// 	- **DeletedFailed**
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateServiceInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceResponseBody) GetMarketInstanceId() *string {
	return s.MarketInstanceId
}

func (s *CreateServiceInstanceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateServiceInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceInstanceResponseBody) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *CreateServiceInstanceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateServiceInstanceResponseBody) SetMarketInstanceId(v string) *CreateServiceInstanceResponseBody {
	s.MarketInstanceId = &v
	return s
}

func (s *CreateServiceInstanceResponseBody) SetOrderId(v string) *CreateServiceInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateServiceInstanceResponseBody) SetRequestId(v string) *CreateServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceInstanceResponseBody) SetServiceInstanceId(v string) *CreateServiceInstanceResponseBody {
	s.ServiceInstanceId = &v
	return s
}

func (s *CreateServiceInstanceResponseBody) SetStatus(v string) *CreateServiceInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *CreateServiceInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
