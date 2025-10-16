// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePostpayUserInternetStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdatePostpayUserInternetStatusRequest
	GetInstanceId() *string
	SetLang(v string) *UpdatePostpayUserInternetStatusRequest
	GetLang() *string
	SetOperate(v string) *UpdatePostpayUserInternetStatusRequest
	GetOperate() *string
}

type UpdatePostpayUserInternetStatusRequest struct {
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
	// The operation type.
	//
	// 	- Set the value to open.
	//
	// example:
	//
	// open
	Operate *string `json:"Operate,omitempty" xml:"Operate,omitempty"`
}

func (s UpdatePostpayUserInternetStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePostpayUserInternetStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdatePostpayUserInternetStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdatePostpayUserInternetStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdatePostpayUserInternetStatusRequest) GetOperate() *string {
	return s.Operate
}

func (s *UpdatePostpayUserInternetStatusRequest) SetInstanceId(v string) *UpdatePostpayUserInternetStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdatePostpayUserInternetStatusRequest) SetLang(v string) *UpdatePostpayUserInternetStatusRequest {
	s.Lang = &v
	return s
}

func (s *UpdatePostpayUserInternetStatusRequest) SetOperate(v string) *UpdatePostpayUserInternetStatusRequest {
	s.Operate = &v
	return s
}

func (s *UpdatePostpayUserInternetStatusRequest) Validate() error {
	return dara.Validate(s)
}
