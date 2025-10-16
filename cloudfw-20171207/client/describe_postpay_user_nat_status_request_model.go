// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePostpayUserNatStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribePostpayUserNatStatusRequest
	GetInstanceId() *string
	SetLang(v string) *DescribePostpayUserNatStatusRequest
	GetLang() *string
}

type DescribePostpayUserNatStatusRequest struct {
	// The instance ID of Cloud Firewall.
	//
	// example:
	//
	// cfw_elasticity_public_cn-zsk39m******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// 	- **zh*	- (default)
	//
	// 	- **en**
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribePostpayUserNatStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePostpayUserNatStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribePostpayUserNatStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribePostpayUserNatStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribePostpayUserNatStatusRequest) SetInstanceId(v string) *DescribePostpayUserNatStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribePostpayUserNatStatusRequest) SetLang(v string) *DescribePostpayUserNatStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribePostpayUserNatStatusRequest) Validate() error {
	return dara.Validate(s)
}
