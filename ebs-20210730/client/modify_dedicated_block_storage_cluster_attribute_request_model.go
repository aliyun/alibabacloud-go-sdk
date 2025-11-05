// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDedicatedBlockStorageClusterAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyDedicatedBlockStorageClusterAttributeRequest
	GetClientToken() *string
	SetDbscId(v string) *ModifyDedicatedBlockStorageClusterAttributeRequest
	GetDbscId() *string
	SetDbscName(v string) *ModifyDedicatedBlockStorageClusterAttributeRequest
	GetDbscName() *string
	SetDescription(v string) *ModifyDedicatedBlockStorageClusterAttributeRequest
	GetDescription() *string
	SetRegionId(v string) *ModifyDedicatedBlockStorageClusterAttributeRequest
	GetRegionId() *string
}

type ModifyDedicatedBlockStorageClusterAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests.
	//
	// The ClientToken value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How do I ensure idempotence ](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the dedicated block storage cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbsc-cn-od43bf****
	DbscId *string `json:"DbscId,omitempty" xml:"DbscId,omitempty"`
	// The new name of the dedicated block storage cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-test-dbsc
	DbscName *string `json:"DbscName,omitempty" xml:"DbscName,omitempty"`
	// The new description of dedicated block storage cluster.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The region ID of the dedicated block storage cluster. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDedicatedBlockStorageClusterAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDedicatedBlockStorageClusterAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedBlockStorageClusterAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDedicatedBlockStorageClusterAttributeRequest) GetDbscId() *string {
	return s.DbscId
}

func (s *ModifyDedicatedBlockStorageClusterAttributeRequest) GetDbscName() *string {
	return s.DbscName
}

func (s *ModifyDedicatedBlockStorageClusterAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyDedicatedBlockStorageClusterAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDedicatedBlockStorageClusterAttributeRequest) SetClientToken(v string) *ModifyDedicatedBlockStorageClusterAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDedicatedBlockStorageClusterAttributeRequest) SetDbscId(v string) *ModifyDedicatedBlockStorageClusterAttributeRequest {
	s.DbscId = &v
	return s
}

func (s *ModifyDedicatedBlockStorageClusterAttributeRequest) SetDbscName(v string) *ModifyDedicatedBlockStorageClusterAttributeRequest {
	s.DbscName = &v
	return s
}

func (s *ModifyDedicatedBlockStorageClusterAttributeRequest) SetDescription(v string) *ModifyDedicatedBlockStorageClusterAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyDedicatedBlockStorageClusterAttributeRequest) SetRegionId(v string) *ModifyDedicatedBlockStorageClusterAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDedicatedBlockStorageClusterAttributeRequest) Validate() error {
	return dara.Validate(s)
}
