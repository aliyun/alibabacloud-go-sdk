// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ModifyInstanceTypeRequest
	GetClusterId() *string
	SetCoreInstanceType(v string) *ModifyInstanceTypeRequest
	GetCoreInstanceType() *string
	SetMasterInstanceType(v string) *ModifyInstanceTypeRequest
	GetMasterInstanceType() *string
}

type ModifyInstanceTypeRequest struct {
	// The ID of target instance. You can call [DescribeInstances](https://help.aliyun.com/document_detail/144595.html) to obtain target instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// hb-bp1x940uh********
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The new node specifications of the core node. You can invoke [DescribeInstanceType](https://help.aliyun.com/document_detail/145796.html) to obtain the available node specifications.
	//
	// > You must specify either the MasterInstanceType parameter or the CoreInstanceType parameter.
	//
	// example:
	//
	// hbase.sn1.8xlarge
	CoreInstanceType *string `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
	// The new node specifications of the master node. You can invoke [DescribeInstanceType](https://help.aliyun.com/document_detail/145796.html) to obtain the available node specifications.
	//
	// > You must specify either the MasterInstanceType parameter or the CoreInstanceType parameter.
	//
	// example:
	//
	// hbase.sn1.large
	MasterInstanceType *string `json:"MasterInstanceType,omitempty" xml:"MasterInstanceType,omitempty"`
}

func (s ModifyInstanceTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceTypeRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyInstanceTypeRequest) GetCoreInstanceType() *string {
	return s.CoreInstanceType
}

func (s *ModifyInstanceTypeRequest) GetMasterInstanceType() *string {
	return s.MasterInstanceType
}

func (s *ModifyInstanceTypeRequest) SetClusterId(v string) *ModifyInstanceTypeRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyInstanceTypeRequest) SetCoreInstanceType(v string) *ModifyInstanceTypeRequest {
	s.CoreInstanceType = &v
	return s
}

func (s *ModifyInstanceTypeRequest) SetMasterInstanceType(v string) *ModifyInstanceTypeRequest {
	s.MasterInstanceType = &v
	return s
}

func (s *ModifyInstanceTypeRequest) Validate() error {
	return dara.Validate(s)
}
