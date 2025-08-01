// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInstancesInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *QueryInstancesInfoRequest
	GetAcceptLanguage() *string
	SetClusterId(v string) *QueryInstancesInfoRequest
	GetClusterId() *string
	SetInstanceId(v string) *QueryInstancesInfoRequest
	GetInstanceId() *string
	SetOrderId(v string) *QueryInstancesInfoRequest
	GetOrderId() *string
	SetRegionId(v string) *QueryInstancesInfoRequest
	GetRegionId() *string
	SetRequestPars(v string) *QueryInstancesInfoRequest
	GetRequestPars() *string
}

type QueryInstancesInfoRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// mse-09k1q11****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// mse_prepaid_public_cn-tl32g1u9k01
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 20574710974****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the region where the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The extended request parameters in the JSON format.
	//
	// example:
	//
	// {}
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
}

func (s QueryInstancesInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryInstancesInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryInstancesInfoRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *QueryInstancesInfoRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *QueryInstancesInfoRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryInstancesInfoRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *QueryInstancesInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryInstancesInfoRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *QueryInstancesInfoRequest) SetAcceptLanguage(v string) *QueryInstancesInfoRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *QueryInstancesInfoRequest) SetClusterId(v string) *QueryInstancesInfoRequest {
	s.ClusterId = &v
	return s
}

func (s *QueryInstancesInfoRequest) SetInstanceId(v string) *QueryInstancesInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryInstancesInfoRequest) SetOrderId(v string) *QueryInstancesInfoRequest {
	s.OrderId = &v
	return s
}

func (s *QueryInstancesInfoRequest) SetRegionId(v string) *QueryInstancesInfoRequest {
	s.RegionId = &v
	return s
}

func (s *QueryInstancesInfoRequest) SetRequestPars(v string) *QueryInstancesInfoRequest {
	s.RequestPars = &v
	return s
}

func (s *QueryInstancesInfoRequest) Validate() error {
	return dara.Validate(s)
}
