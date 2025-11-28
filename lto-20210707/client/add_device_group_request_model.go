// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDeviceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizedCount(v int32) *AddDeviceGroupRequest
	GetAuthorizedCount() *int32
	SetProductKey(v string) *AddDeviceGroupRequest
	GetProductKey() *string
	SetRegionId(v string) *AddDeviceGroupRequest
	GetRegionId() *string
	SetRemark(v string) *AddDeviceGroupRequest
	GetRemark() *string
}

type AddDeviceGroupRequest struct {
	AuthorizedCount *int32 `json:"AuthorizedCount,omitempty" xml:"AuthorizedCount,omitempty"`
	// This parameter is required.
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s AddDeviceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDeviceGroupRequest) GoString() string {
	return s.String()
}

func (s *AddDeviceGroupRequest) GetAuthorizedCount() *int32 {
	return s.AuthorizedCount
}

func (s *AddDeviceGroupRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *AddDeviceGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddDeviceGroupRequest) GetRemark() *string {
	return s.Remark
}

func (s *AddDeviceGroupRequest) SetAuthorizedCount(v int32) *AddDeviceGroupRequest {
	s.AuthorizedCount = &v
	return s
}

func (s *AddDeviceGroupRequest) SetProductKey(v string) *AddDeviceGroupRequest {
	s.ProductKey = &v
	return s
}

func (s *AddDeviceGroupRequest) SetRegionId(v string) *AddDeviceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *AddDeviceGroupRequest) SetRemark(v string) *AddDeviceGroupRequest {
	s.Remark = &v
	return s
}

func (s *AddDeviceGroupRequest) Validate() error {
	return dara.Validate(s)
}
