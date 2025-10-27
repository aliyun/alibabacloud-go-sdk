// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePostpayUserNatStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdatePostpayUserNatStatusRequest
	GetInstanceId() *string
	SetLang(v string) *UpdatePostpayUserNatStatusRequest
	GetLang() *string
	SetOperate(v string) *UpdatePostpayUserNatStatusRequest
	GetOperate() *string
}

type UpdatePostpayUserNatStatusRequest struct {
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
	// The operation type.
	//
	// 	- Set the value to open.
	//
	// example:
	//
	// open
	Operate *string `json:"Operate,omitempty" xml:"Operate,omitempty"`
}

func (s UpdatePostpayUserNatStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePostpayUserNatStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdatePostpayUserNatStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdatePostpayUserNatStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdatePostpayUserNatStatusRequest) GetOperate() *string {
	return s.Operate
}

func (s *UpdatePostpayUserNatStatusRequest) SetInstanceId(v string) *UpdatePostpayUserNatStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdatePostpayUserNatStatusRequest) SetLang(v string) *UpdatePostpayUserNatStatusRequest {
	s.Lang = &v
	return s
}

func (s *UpdatePostpayUserNatStatusRequest) SetOperate(v string) *UpdatePostpayUserNatStatusRequest {
	s.Operate = &v
	return s
}

func (s *UpdatePostpayUserNatStatusRequest) Validate() error {
	return dara.Validate(s)
}
