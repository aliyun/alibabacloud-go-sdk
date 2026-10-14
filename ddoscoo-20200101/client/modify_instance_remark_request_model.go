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
}

type ModifyInstanceRemarkRequest struct {
	// The ID of the Anti-DDoS Pro or Anti-DDoS Premium instance to manage.
	//
	// > You can invoke [DescribeInstanceIds](https://help.aliyun.com/document_detail/157459.html) to query the IDs of all Anti-DDoS Pro or Anti-DDoS Premium instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// new-remark
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
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

func (s *ModifyInstanceRemarkRequest) SetInstanceId(v string) *ModifyInstanceRemarkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceRemarkRequest) SetRemark(v string) *ModifyInstanceRemarkRequest {
	s.Remark = &v
	return s
}

func (s *ModifyInstanceRemarkRequest) Validate() error {
	return dara.Validate(s)
}
