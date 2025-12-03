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
	// This parameter is required.
	//
	// example:
	//
	// hb-bp1x940uh********
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// hbase.sn1.8xlarge
	CoreInstanceType *string `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
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
