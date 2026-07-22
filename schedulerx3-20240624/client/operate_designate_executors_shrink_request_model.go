// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateDesignateExecutorsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressListShrink(v string) *OperateDesignateExecutorsShrinkRequest
	GetAddressListShrink() *string
	SetAppGroupId(v int64) *OperateDesignateExecutorsShrinkRequest
	GetAppGroupId() *int64
	SetAppName(v string) *OperateDesignateExecutorsShrinkRequest
	GetAppName() *string
	SetClusterId(v string) *OperateDesignateExecutorsShrinkRequest
	GetClusterId() *string
	SetDesignateType(v int32) *OperateDesignateExecutorsShrinkRequest
	GetDesignateType() *int32
	SetJobId(v int64) *OperateDesignateExecutorsShrinkRequest
	GetJobId() *int64
	SetTransferable(v bool) *OperateDesignateExecutorsShrinkRequest
	GetTransferable() *bool
}

type OperateDesignateExecutorsShrinkRequest struct {
	// The address list.
	//
	// This parameter is required.
	AddressListShrink *string `json:"AddressList,omitempty" xml:"AddressList,omitempty"`
	AppGroupId        *int64  `json:"AppGroupId,omitempty" xml:"AppGroupId,omitempty"`
	// The application name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The type of the designated machine. Valid values:
	//
	// - **1**: designated worker.
	//
	// - **2**: designated label.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DesignateType *int32 `json:"DesignateType,omitempty" xml:"DesignateType,omitempty"`
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 74
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Specifies whether to enable failover.
	//
	// example:
	//
	// true
	Transferable *bool `json:"Transferable,omitempty" xml:"Transferable,omitempty"`
}

func (s OperateDesignateExecutorsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateDesignateExecutorsShrinkRequest) GoString() string {
	return s.String()
}

func (s *OperateDesignateExecutorsShrinkRequest) GetAddressListShrink() *string {
	return s.AddressListShrink
}

func (s *OperateDesignateExecutorsShrinkRequest) GetAppGroupId() *int64 {
	return s.AppGroupId
}

func (s *OperateDesignateExecutorsShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *OperateDesignateExecutorsShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *OperateDesignateExecutorsShrinkRequest) GetDesignateType() *int32 {
	return s.DesignateType
}

func (s *OperateDesignateExecutorsShrinkRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *OperateDesignateExecutorsShrinkRequest) GetTransferable() *bool {
	return s.Transferable
}

func (s *OperateDesignateExecutorsShrinkRequest) SetAddressListShrink(v string) *OperateDesignateExecutorsShrinkRequest {
	s.AddressListShrink = &v
	return s
}

func (s *OperateDesignateExecutorsShrinkRequest) SetAppGroupId(v int64) *OperateDesignateExecutorsShrinkRequest {
	s.AppGroupId = &v
	return s
}

func (s *OperateDesignateExecutorsShrinkRequest) SetAppName(v string) *OperateDesignateExecutorsShrinkRequest {
	s.AppName = &v
	return s
}

func (s *OperateDesignateExecutorsShrinkRequest) SetClusterId(v string) *OperateDesignateExecutorsShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateDesignateExecutorsShrinkRequest) SetDesignateType(v int32) *OperateDesignateExecutorsShrinkRequest {
	s.DesignateType = &v
	return s
}

func (s *OperateDesignateExecutorsShrinkRequest) SetJobId(v int64) *OperateDesignateExecutorsShrinkRequest {
	s.JobId = &v
	return s
}

func (s *OperateDesignateExecutorsShrinkRequest) SetTransferable(v bool) *OperateDesignateExecutorsShrinkRequest {
	s.Transferable = &v
	return s
}

func (s *OperateDesignateExecutorsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
