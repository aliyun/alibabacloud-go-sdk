// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateDesignateExecutorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressList(v []*string) *OperateDesignateExecutorsRequest
	GetAddressList() []*string
	SetAppGroupId(v int64) *OperateDesignateExecutorsRequest
	GetAppGroupId() *int64
	SetAppName(v string) *OperateDesignateExecutorsRequest
	GetAppName() *string
	SetClusterId(v string) *OperateDesignateExecutorsRequest
	GetClusterId() *string
	SetDesignateType(v int32) *OperateDesignateExecutorsRequest
	GetDesignateType() *int32
	SetJobId(v int64) *OperateDesignateExecutorsRequest
	GetJobId() *int64
	SetTransferable(v bool) *OperateDesignateExecutorsRequest
	GetTransferable() *bool
}

type OperateDesignateExecutorsRequest struct {
	// The address list.
	//
	// This parameter is required.
	AddressList []*string `json:"AddressList,omitempty" xml:"AddressList,omitempty" type:"Repeated"`
	AppGroupId  *int64    `json:"AppGroupId,omitempty" xml:"AppGroupId,omitempty"`
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

func (s OperateDesignateExecutorsRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateDesignateExecutorsRequest) GoString() string {
	return s.String()
}

func (s *OperateDesignateExecutorsRequest) GetAddressList() []*string {
	return s.AddressList
}

func (s *OperateDesignateExecutorsRequest) GetAppGroupId() *int64 {
	return s.AppGroupId
}

func (s *OperateDesignateExecutorsRequest) GetAppName() *string {
	return s.AppName
}

func (s *OperateDesignateExecutorsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *OperateDesignateExecutorsRequest) GetDesignateType() *int32 {
	return s.DesignateType
}

func (s *OperateDesignateExecutorsRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *OperateDesignateExecutorsRequest) GetTransferable() *bool {
	return s.Transferable
}

func (s *OperateDesignateExecutorsRequest) SetAddressList(v []*string) *OperateDesignateExecutorsRequest {
	s.AddressList = v
	return s
}

func (s *OperateDesignateExecutorsRequest) SetAppGroupId(v int64) *OperateDesignateExecutorsRequest {
	s.AppGroupId = &v
	return s
}

func (s *OperateDesignateExecutorsRequest) SetAppName(v string) *OperateDesignateExecutorsRequest {
	s.AppName = &v
	return s
}

func (s *OperateDesignateExecutorsRequest) SetClusterId(v string) *OperateDesignateExecutorsRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateDesignateExecutorsRequest) SetDesignateType(v int32) *OperateDesignateExecutorsRequest {
	s.DesignateType = &v
	return s
}

func (s *OperateDesignateExecutorsRequest) SetJobId(v int64) *OperateDesignateExecutorsRequest {
	s.JobId = &v
	return s
}

func (s *OperateDesignateExecutorsRequest) SetTransferable(v bool) *OperateDesignateExecutorsRequest {
	s.Transferable = &v
	return s
}

func (s *OperateDesignateExecutorsRequest) Validate() error {
	return dara.Validate(s)
}
