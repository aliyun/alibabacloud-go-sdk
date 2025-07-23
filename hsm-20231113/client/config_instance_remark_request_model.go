// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigInstanceRemarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ConfigInstanceRemarkRequest
	GetInstanceId() *string
	SetRemark(v string) *ConfigInstanceRemarkRequest
	GetRemark() *string
}

type ConfigInstanceRemarkRequest struct {
	// The ID of the HSM.
	//
	// This parameter is required.
	//
	// example:
	//
	// hsm-cn-vj30bil8****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The description of the HSM.
	//
	// This parameter is required.
	//
	// example:
	//
	// hsmOnline
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s ConfigInstanceRemarkRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigInstanceRemarkRequest) GoString() string {
	return s.String()
}

func (s *ConfigInstanceRemarkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ConfigInstanceRemarkRequest) GetRemark() *string {
	return s.Remark
}

func (s *ConfigInstanceRemarkRequest) SetInstanceId(v string) *ConfigInstanceRemarkRequest {
	s.InstanceId = &v
	return s
}

func (s *ConfigInstanceRemarkRequest) SetRemark(v string) *ConfigInstanceRemarkRequest {
	s.Remark = &v
	return s
}

func (s *ConfigInstanceRemarkRequest) Validate() error {
	return dara.Validate(s)
}
