// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryClusterDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *QueryClusterDetailRequest
	GetAcceptLanguage() *string
	SetAclSwitch(v bool) *QueryClusterDetailRequest
	GetAclSwitch() *bool
	SetInstanceId(v string) *QueryClusterDetailRequest
	GetInstanceId() *string
	SetOrderId(v string) *QueryClusterDetailRequest
	GetOrderId() *string
}

type QueryClusterDetailRequest struct {
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
	// Specifies whether to query access control lists (ACLs).
	//
	// example:
	//
	// false
	AclSwitch *bool `json:"AclSwitch,omitempty" xml:"AclSwitch,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// mse-cn-st21ri2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 20576750143****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s QueryClusterDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryClusterDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryClusterDetailRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *QueryClusterDetailRequest) GetAclSwitch() *bool {
	return s.AclSwitch
}

func (s *QueryClusterDetailRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryClusterDetailRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *QueryClusterDetailRequest) SetAcceptLanguage(v string) *QueryClusterDetailRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *QueryClusterDetailRequest) SetAclSwitch(v bool) *QueryClusterDetailRequest {
	s.AclSwitch = &v
	return s
}

func (s *QueryClusterDetailRequest) SetInstanceId(v string) *QueryClusterDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryClusterDetailRequest) SetOrderId(v string) *QueryClusterDetailRequest {
	s.OrderId = &v
	return s
}

func (s *QueryClusterDetailRequest) Validate() error {
	return dara.Validate(s)
}
