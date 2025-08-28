// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSmartCallOperateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *SmartCallOperateRequest
	GetCallId() *string
	SetCommand(v string) *SmartCallOperateRequest
	GetCommand() *string
	SetOwnerId(v int64) *SmartCallOperateRequest
	GetOwnerId() *int64
	SetParam(v string) *SmartCallOperateRequest
	GetParam() *string
	SetResourceOwnerAccount(v string) *SmartCallOperateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SmartCallOperateRequest
	GetResourceOwnerId() *int64
}

type SmartCallOperateRequest struct {
	// The unique receipt ID of the call. You can call the [SmartCall](https://help.aliyun.com/document_detail/393526.html) operation to obtain the receipt ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 116012854210^1028142****
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// The action that is initiated to the called number of an outbound robocall.
	//
	// > Only the value **parallelBridge*	- is supported. This value indicates that a bridge action is initiated between a called number and an agent of the call center.
	//
	// This parameter is required.
	//
	// example:
	//
	// parallelBridge
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The extended field.
	//
	// example:
	//
	// Param
	Param                *string `json:"Param,omitempty" xml:"Param,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SmartCallOperateRequest) String() string {
	return dara.Prettify(s)
}

func (s SmartCallOperateRequest) GoString() string {
	return s.String()
}

func (s *SmartCallOperateRequest) GetCallId() *string {
	return s.CallId
}

func (s *SmartCallOperateRequest) GetCommand() *string {
	return s.Command
}

func (s *SmartCallOperateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SmartCallOperateRequest) GetParam() *string {
	return s.Param
}

func (s *SmartCallOperateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SmartCallOperateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SmartCallOperateRequest) SetCallId(v string) *SmartCallOperateRequest {
	s.CallId = &v
	return s
}

func (s *SmartCallOperateRequest) SetCommand(v string) *SmartCallOperateRequest {
	s.Command = &v
	return s
}

func (s *SmartCallOperateRequest) SetOwnerId(v int64) *SmartCallOperateRequest {
	s.OwnerId = &v
	return s
}

func (s *SmartCallOperateRequest) SetParam(v string) *SmartCallOperateRequest {
	s.Param = &v
	return s
}

func (s *SmartCallOperateRequest) SetResourceOwnerAccount(v string) *SmartCallOperateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SmartCallOperateRequest) SetResourceOwnerId(v int64) *SmartCallOperateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SmartCallOperateRequest) Validate() error {
	return dara.Validate(s)
}
