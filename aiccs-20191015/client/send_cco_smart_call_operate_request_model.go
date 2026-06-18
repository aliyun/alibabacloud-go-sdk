// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendCcoSmartCallOperateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *SendCcoSmartCallOperateRequest
	GetCallId() *string
	SetCommand(v string) *SendCcoSmartCallOperateRequest
	GetCommand() *string
	SetOwnerId(v int64) *SendCcoSmartCallOperateRequest
	GetOwnerId() *int64
	SetParam(v string) *SendCcoSmartCallOperateRequest
	GetParam() *string
	SetProdCode(v string) *SendCcoSmartCallOperateRequest
	GetProdCode() *string
	SetResourceOwnerAccount(v string) *SendCcoSmartCallOperateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SendCcoSmartCallOperateRequest
	GetResourceOwnerId() *int64
}

type SendCcoSmartCallOperateRequest struct {
	// Unique receipt ID of the call. This can be obtained from the response of the [SendCcoSmartCall](https://help.aliyun.com/document_detail/311247.html) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 116012854210^102814279****
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// Specifies the action to be performed for the called number during an Intelligent Outbound Call.
	//
	// > Currently, only the **parallelBridge*	- parameter is supported, which indicates bridging the called number with a call center agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// parallelBridge
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Extension field.
	//
	// example:
	//
	// Param
	Param *string `json:"Param,omitempty" xml:"Param,omitempty"`
	// Product name. Default value: **aiccs**.
	//
	// example:
	//
	// aiccs
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SendCcoSmartCallOperateRequest) String() string {
	return dara.Prettify(s)
}

func (s SendCcoSmartCallOperateRequest) GoString() string {
	return s.String()
}

func (s *SendCcoSmartCallOperateRequest) GetCallId() *string {
	return s.CallId
}

func (s *SendCcoSmartCallOperateRequest) GetCommand() *string {
	return s.Command
}

func (s *SendCcoSmartCallOperateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SendCcoSmartCallOperateRequest) GetParam() *string {
	return s.Param
}

func (s *SendCcoSmartCallOperateRequest) GetProdCode() *string {
	return s.ProdCode
}

func (s *SendCcoSmartCallOperateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SendCcoSmartCallOperateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SendCcoSmartCallOperateRequest) SetCallId(v string) *SendCcoSmartCallOperateRequest {
	s.CallId = &v
	return s
}

func (s *SendCcoSmartCallOperateRequest) SetCommand(v string) *SendCcoSmartCallOperateRequest {
	s.Command = &v
	return s
}

func (s *SendCcoSmartCallOperateRequest) SetOwnerId(v int64) *SendCcoSmartCallOperateRequest {
	s.OwnerId = &v
	return s
}

func (s *SendCcoSmartCallOperateRequest) SetParam(v string) *SendCcoSmartCallOperateRequest {
	s.Param = &v
	return s
}

func (s *SendCcoSmartCallOperateRequest) SetProdCode(v string) *SendCcoSmartCallOperateRequest {
	s.ProdCode = &v
	return s
}

func (s *SendCcoSmartCallOperateRequest) SetResourceOwnerAccount(v string) *SendCcoSmartCallOperateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SendCcoSmartCallOperateRequest) SetResourceOwnerId(v int64) *SendCcoSmartCallOperateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SendCcoSmartCallOperateRequest) Validate() error {
	return dara.Validate(s)
}
