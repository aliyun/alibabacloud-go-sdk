// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyColumnarClassRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColumnarClass(v string) *ModifyColumnarClassRequest
	GetColumnarClass() *string
	SetColumnarNodeCount(v string) *ModifyColumnarClassRequest
	GetColumnarNodeCount() *string
	SetInstanceName(v string) *ModifyColumnarClassRequest
	GetInstanceName() *string
	SetRegionId(v string) *ModifyColumnarClassRequest
	GetRegionId() *string
	SetSwitchMode(v string) *ModifyColumnarClassRequest
	GetSwitchMode() *string
}

type ModifyColumnarClassRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// polarx.n8.large.col
	ColumnarClass *string `json:"ColumnarClass,omitempty" xml:"ColumnarClass,omitempty"`
	// example:
	//
	// **
	ColumnarNodeCount *string `json:"ColumnarNodeCount,omitempty" xml:"ColumnarNodeCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzrh51fze****pon-cdc
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 0
	SwitchMode *string `json:"SwitchMode,omitempty" xml:"SwitchMode,omitempty"`
}

func (s ModifyColumnarClassRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyColumnarClassRequest) GoString() string {
	return s.String()
}

func (s *ModifyColumnarClassRequest) GetColumnarClass() *string {
	return s.ColumnarClass
}

func (s *ModifyColumnarClassRequest) GetColumnarNodeCount() *string {
	return s.ColumnarNodeCount
}

func (s *ModifyColumnarClassRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyColumnarClassRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyColumnarClassRequest) GetSwitchMode() *string {
	return s.SwitchMode
}

func (s *ModifyColumnarClassRequest) SetColumnarClass(v string) *ModifyColumnarClassRequest {
	s.ColumnarClass = &v
	return s
}

func (s *ModifyColumnarClassRequest) SetColumnarNodeCount(v string) *ModifyColumnarClassRequest {
	s.ColumnarNodeCount = &v
	return s
}

func (s *ModifyColumnarClassRequest) SetInstanceName(v string) *ModifyColumnarClassRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyColumnarClassRequest) SetRegionId(v string) *ModifyColumnarClassRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyColumnarClassRequest) SetSwitchMode(v string) *ModifyColumnarClassRequest {
	s.SwitchMode = &v
	return s
}

func (s *ModifyColumnarClassRequest) Validate() error {
	return dara.Validate(s)
}
