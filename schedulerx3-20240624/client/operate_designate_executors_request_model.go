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
	// This parameter is required.
	AddressList []*string `json:"AddressList,omitempty" xml:"AddressList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	DesignateType *int32 `json:"DesignateType,omitempty" xml:"DesignateType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 74
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
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
