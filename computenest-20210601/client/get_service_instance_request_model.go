// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMarketInstanceId(v string) *GetServiceInstanceRequest
	GetMarketInstanceId() *string
	SetRegionId(v string) *GetServiceInstanceRequest
	GetRegionId() *string
	SetServiceInstanceId(v string) *GetServiceInstanceRequest
	GetServiceInstanceId() *string
}

type GetServiceInstanceRequest struct {
	// The MarketInstance ID.
	//
	// example:
	//
	// 704***59
	MarketInstanceId *string `json:"MarketInstanceId,omitempty" xml:"MarketInstanceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service instance ID.
	//
	// >  You must specify either `ServiceInstanceId` or `MarketInstanceId`. Otherwise, the operation fails.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s GetServiceInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceRequest) GetMarketInstanceId() *string {
	return s.MarketInstanceId
}

func (s *GetServiceInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetServiceInstanceRequest) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *GetServiceInstanceRequest) SetMarketInstanceId(v string) *GetServiceInstanceRequest {
	s.MarketInstanceId = &v
	return s
}

func (s *GetServiceInstanceRequest) SetRegionId(v string) *GetServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceInstanceRequest) SetServiceInstanceId(v string) *GetServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *GetServiceInstanceRequest) Validate() error {
	return dara.Validate(s)
}
