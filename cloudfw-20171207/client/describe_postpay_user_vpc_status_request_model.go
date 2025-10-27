// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePostpayUserVpcStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribePostpayUserVpcStatusRequest
	GetInstanceId() *string
	SetLang(v string) *DescribePostpayUserVpcStatusRequest
	GetLang() *string
}

type DescribePostpayUserVpcStatusRequest struct {
	// The instance ID of Cloud Firewall.
	//
	// example:
	//
	// cfw_elasticity_public_cn-************
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

func (s DescribePostpayUserVpcStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePostpayUserVpcStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribePostpayUserVpcStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribePostpayUserVpcStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribePostpayUserVpcStatusRequest) SetInstanceId(v string) *DescribePostpayUserVpcStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribePostpayUserVpcStatusRequest) SetLang(v string) *DescribePostpayUserVpcStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribePostpayUserVpcStatusRequest) Validate() error {
	return dara.Validate(s)
}
