// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryClusterInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *QueryClusterInfoRequest
	GetAcceptLanguage() *string
	SetAclSwitch(v bool) *QueryClusterInfoRequest
	GetAclSwitch() *bool
	SetClusterId(v string) *QueryClusterInfoRequest
	GetClusterId() *string
	SetInstanceId(v string) *QueryClusterInfoRequest
	GetInstanceId() *string
	SetOrderId(v string) *QueryClusterInfoRequest
	GetOrderId() *string
	SetRegionId(v string) *QueryClusterInfoRequest
	GetRegionId() *string
	SetRequestPars(v string) *QueryClusterInfoRequest
	GetRequestPars() *string
}

type QueryClusterInfoRequest struct {
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
	// Specifies whether to query the configuration of a public IP address whitelist.
	//
	// example:
	//
	// false
	AclSwitch *bool `json:"AclSwitch,omitempty" xml:"AclSwitch,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// mse-09k1q11****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// mse_prepaid_public_cn-7mz2t63ci03
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 20574710974****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the region.
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

func (s QueryClusterInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryClusterInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryClusterInfoRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *QueryClusterInfoRequest) GetAclSwitch() *bool {
	return s.AclSwitch
}

func (s *QueryClusterInfoRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *QueryClusterInfoRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryClusterInfoRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *QueryClusterInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryClusterInfoRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *QueryClusterInfoRequest) SetAcceptLanguage(v string) *QueryClusterInfoRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *QueryClusterInfoRequest) SetAclSwitch(v bool) *QueryClusterInfoRequest {
	s.AclSwitch = &v
	return s
}

func (s *QueryClusterInfoRequest) SetClusterId(v string) *QueryClusterInfoRequest {
	s.ClusterId = &v
	return s
}

func (s *QueryClusterInfoRequest) SetInstanceId(v string) *QueryClusterInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryClusterInfoRequest) SetOrderId(v string) *QueryClusterInfoRequest {
	s.OrderId = &v
	return s
}

func (s *QueryClusterInfoRequest) SetRegionId(v string) *QueryClusterInfoRequest {
	s.RegionId = &v
	return s
}

func (s *QueryClusterInfoRequest) SetRequestPars(v string) *QueryClusterInfoRequest {
	s.RequestPars = &v
	return s
}

func (s *QueryClusterInfoRequest) Validate() error {
	return dara.Validate(s)
}
