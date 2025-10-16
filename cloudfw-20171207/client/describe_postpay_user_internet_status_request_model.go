// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePostpayUserInternetStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribePostpayUserInternetStatusRequest
	GetInstanceId() *string
	SetLang(v string) *DescribePostpayUserInternetStatusRequest
	GetLang() *string
}

type DescribePostpayUserInternetStatusRequest struct {
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

func (s DescribePostpayUserInternetStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePostpayUserInternetStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribePostpayUserInternetStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribePostpayUserInternetStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribePostpayUserInternetStatusRequest) SetInstanceId(v string) *DescribePostpayUserInternetStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribePostpayUserInternetStatusRequest) SetLang(v string) *DescribePostpayUserInternetStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribePostpayUserInternetStatusRequest) Validate() error {
	return dara.Validate(s)
}
