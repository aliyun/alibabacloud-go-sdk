// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePocFunctionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreatePocFunctionRequest
	GetInstanceId() *string
	SetRegionId(v string) *CreatePocFunctionRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *CreatePocFunctionRequest
	GetResourceManagerResourceGroupId() *string
	SetType(v string) *CreatePocFunctionRequest
	GetType() *string
}

type CreatePocFunctionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// botWeb
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreatePocFunctionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePocFunctionRequest) GoString() string {
	return s.String()
}

func (s *CreatePocFunctionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreatePocFunctionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePocFunctionRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *CreatePocFunctionRequest) GetType() *string {
	return s.Type
}

func (s *CreatePocFunctionRequest) SetInstanceId(v string) *CreatePocFunctionRequest {
	s.InstanceId = &v
	return s
}

func (s *CreatePocFunctionRequest) SetRegionId(v string) *CreatePocFunctionRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePocFunctionRequest) SetResourceManagerResourceGroupId(v string) *CreatePocFunctionRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *CreatePocFunctionRequest) SetType(v string) *CreatePocFunctionRequest {
	s.Type = &v
	return s
}

func (s *CreatePocFunctionRequest) Validate() error {
	return dara.Validate(s)
}
