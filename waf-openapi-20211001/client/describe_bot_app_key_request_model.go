// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBotAppKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeBotAppKeyRequest
	GetInstanceId() *string
	SetKeyVersion(v string) *DescribeBotAppKeyRequest
	GetKeyVersion() *string
	SetRegionId(v string) *DescribeBotAppKeyRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeBotAppKeyRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeBotAppKeyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-5v23u******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	KeyVersion *string `json:"KeyVersion,omitempty" xml:"KeyVersion,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-aek2uo27badl***
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeBotAppKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBotAppKeyRequest) GoString() string {
	return s.String()
}

func (s *DescribeBotAppKeyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeBotAppKeyRequest) GetKeyVersion() *string {
	return s.KeyVersion
}

func (s *DescribeBotAppKeyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBotAppKeyRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeBotAppKeyRequest) SetInstanceId(v string) *DescribeBotAppKeyRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeBotAppKeyRequest) SetKeyVersion(v string) *DescribeBotAppKeyRequest {
	s.KeyVersion = &v
	return s
}

func (s *DescribeBotAppKeyRequest) SetRegionId(v string) *DescribeBotAppKeyRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBotAppKeyRequest) SetResourceManagerResourceGroupId(v string) *DescribeBotAppKeyRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeBotAppKeyRequest) Validate() error {
	return dara.Validate(s)
}
