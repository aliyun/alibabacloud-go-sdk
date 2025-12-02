// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyServiceInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ModifyServiceInfoRequest
	GetRegionId() *string
	SetResourceType(v string) *ModifyServiceInfoRequest
	GetResourceType() *string
	SetServiceCode(v string) *ModifyServiceInfoRequest
	GetServiceCode() *string
	SetServiceDesc(v string) *ModifyServiceInfoRequest
	GetServiceDesc() *string
	SetServiceName(v string) *ModifyServiceInfoRequest
	GetServiceName() *string
}

type ModifyServiceInfoRequest struct {
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource type.
	//
	// example:
	//
	// image
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Service code.
	//
	// example:
	//
	// baselineCheck
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// Service description.
	//
	// example:
	//
	// 描述
	ServiceDesc *string `json:"ServiceDesc,omitempty" xml:"ServiceDesc,omitempty"`
	// Service name.
	//
	// example:
	//
	// 通用基线检测
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s ModifyServiceInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyServiceInfoRequest) GoString() string {
	return s.String()
}

func (s *ModifyServiceInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyServiceInfoRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ModifyServiceInfoRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *ModifyServiceInfoRequest) GetServiceDesc() *string {
	return s.ServiceDesc
}

func (s *ModifyServiceInfoRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *ModifyServiceInfoRequest) SetRegionId(v string) *ModifyServiceInfoRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyServiceInfoRequest) SetResourceType(v string) *ModifyServiceInfoRequest {
	s.ResourceType = &v
	return s
}

func (s *ModifyServiceInfoRequest) SetServiceCode(v string) *ModifyServiceInfoRequest {
	s.ServiceCode = &v
	return s
}

func (s *ModifyServiceInfoRequest) SetServiceDesc(v string) *ModifyServiceInfoRequest {
	s.ServiceDesc = &v
	return s
}

func (s *ModifyServiceInfoRequest) SetServiceName(v string) *ModifyServiceInfoRequest {
	s.ServiceName = &v
	return s
}

func (s *ModifyServiceInfoRequest) Validate() error {
	return dara.Validate(s)
}
