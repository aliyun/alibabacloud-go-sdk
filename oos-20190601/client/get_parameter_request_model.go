// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetParameterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *GetParameterRequest
	GetName() *string
	SetParameterVersion(v int32) *GetParameterRequest
	GetParameterVersion() *int32
	SetRegionId(v string) *GetParameterRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *GetParameterRequest
	GetResourceGroupId() *string
}

type GetParameterRequest struct {
	// The name of the common parameter. The name can be up to 200 characters in length and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// This parameter is required.
	//
	// example:
	//
	// MyParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version number of the common parameter. Valid values: 1 to 100.
	//
	// example:
	//
	// 1
	ParameterVersion *int32 `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxsn4m*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s GetParameterRequest) String() string {
	return dara.Prettify(s)
}

func (s GetParameterRequest) GoString() string {
	return s.String()
}

func (s *GetParameterRequest) GetName() *string {
	return s.Name
}

func (s *GetParameterRequest) GetParameterVersion() *int32 {
	return s.ParameterVersion
}

func (s *GetParameterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetParameterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetParameterRequest) SetName(v string) *GetParameterRequest {
	s.Name = &v
	return s
}

func (s *GetParameterRequest) SetParameterVersion(v int32) *GetParameterRequest {
	s.ParameterVersion = &v
	return s
}

func (s *GetParameterRequest) SetRegionId(v string) *GetParameterRequest {
	s.RegionId = &v
	return s
}

func (s *GetParameterRequest) SetResourceGroupId(v string) *GetParameterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *GetParameterRequest) Validate() error {
	return dara.Validate(s)
}
