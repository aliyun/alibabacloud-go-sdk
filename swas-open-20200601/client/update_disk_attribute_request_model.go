// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDiskAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateDiskAttributeRequest
	GetClientToken() *string
	SetDiskId(v string) *UpdateDiskAttributeRequest
	GetDiskId() *string
	SetRegionId(v string) *UpdateDiskAttributeRequest
	GetRegionId() *string
	SetRemark(v string) *UpdateDiskAttributeRequest
	GetRemark() *string
}

type UpdateDiskAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The disk ID. You can call the ListDisks operation to query the ID of data disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-bp18kjxg9ebrhsgi****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remarks of the data disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s UpdateDiskAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDiskAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateDiskAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateDiskAttributeRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *UpdateDiskAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateDiskAttributeRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateDiskAttributeRequest) SetClientToken(v string) *UpdateDiskAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateDiskAttributeRequest) SetDiskId(v string) *UpdateDiskAttributeRequest {
	s.DiskId = &v
	return s
}

func (s *UpdateDiskAttributeRequest) SetRegionId(v string) *UpdateDiskAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateDiskAttributeRequest) SetRemark(v string) *UpdateDiskAttributeRequest {
	s.Remark = &v
	return s
}

func (s *UpdateDiskAttributeRequest) Validate() error {
	return dara.Validate(s)
}
