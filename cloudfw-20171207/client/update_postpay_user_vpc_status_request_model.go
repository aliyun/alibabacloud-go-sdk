// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePostpayUserVpcStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdatePostpayUserVpcStatusRequest
	GetInstanceId() *string
	SetLang(v string) *UpdatePostpayUserVpcStatusRequest
	GetLang() *string
	SetOperate(v string) *UpdatePostpayUserVpcStatusRequest
	GetOperate() *string
}

type UpdatePostpayUserVpcStatusRequest struct {
	// The ID of the Cloud Firewall instance.
	//
	// example:
	//
	// cfw_elasticity_public_cn-************
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the request and response. Valid values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The operation type. Currently, only the activation operation is supported. Valid values:
	//
	// - open: activation
	//
	// example:
	//
	// open
	Operate *string `json:"Operate,omitempty" xml:"Operate,omitempty"`
}

func (s UpdatePostpayUserVpcStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePostpayUserVpcStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdatePostpayUserVpcStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdatePostpayUserVpcStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdatePostpayUserVpcStatusRequest) GetOperate() *string {
	return s.Operate
}

func (s *UpdatePostpayUserVpcStatusRequest) SetInstanceId(v string) *UpdatePostpayUserVpcStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdatePostpayUserVpcStatusRequest) SetLang(v string) *UpdatePostpayUserVpcStatusRequest {
	s.Lang = &v
	return s
}

func (s *UpdatePostpayUserVpcStatusRequest) SetOperate(v string) *UpdatePostpayUserVpcStatusRequest {
	s.Operate = &v
	return s
}

func (s *UpdatePostpayUserVpcStatusRequest) Validate() error {
	return dara.Validate(s)
}
