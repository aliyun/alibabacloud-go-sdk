// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceRemarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyInstanceRemarkRequest
	GetInstanceId() *string
	SetRemark(v string) *ModifyInstanceRemarkRequest
	GetRemark() *string
	SetSourceIp(v string) *ModifyInstanceRemarkRequest
	GetSourceIp() *string
}

type ModifyInstanceRemarkRequest struct {
	// example:
	//
	// ddoscoo-cn-XXXXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s ModifyInstanceRemarkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceRemarkRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceRemarkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceRemarkRequest) GetRemark() *string {
	return s.Remark
}

func (s *ModifyInstanceRemarkRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ModifyInstanceRemarkRequest) SetInstanceId(v string) *ModifyInstanceRemarkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceRemarkRequest) SetRemark(v string) *ModifyInstanceRemarkRequest {
	s.Remark = &v
	return s
}

func (s *ModifyInstanceRemarkRequest) SetSourceIp(v string) *ModifyInstanceRemarkRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyInstanceRemarkRequest) Validate() error {
	return dara.Validate(s)
}
