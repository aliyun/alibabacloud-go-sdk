// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiccsSmartCallOperateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *AiccsSmartCallOperateRequest
	GetCallId() *string
	SetCommand(v string) *AiccsSmartCallOperateRequest
	GetCommand() *string
	SetOwnerId(v int64) *AiccsSmartCallOperateRequest
	GetOwnerId() *int64
	SetParam(v string) *AiccsSmartCallOperateRequest
	GetParam() *string
	SetProdCode(v string) *AiccsSmartCallOperateRequest
	GetProdCode() *string
	SetResourceOwnerAccount(v string) *AiccsSmartCallOperateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AiccsSmartCallOperateRequest
	GetResourceOwnerId() *int64
}

type AiccsSmartCallOperateRequest struct {
	// example:
	//
	// 116012854210^102814279****
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// example:
	//
	// parallelBridge
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// Param
	Param *string `json:"Param,omitempty" xml:"Param,omitempty"`
	// example:
	//
	// aiccs
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AiccsSmartCallOperateRequest) String() string {
	return dara.Prettify(s)
}

func (s AiccsSmartCallOperateRequest) GoString() string {
	return s.String()
}

func (s *AiccsSmartCallOperateRequest) GetCallId() *string {
	return s.CallId
}

func (s *AiccsSmartCallOperateRequest) GetCommand() *string {
	return s.Command
}

func (s *AiccsSmartCallOperateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AiccsSmartCallOperateRequest) GetParam() *string {
	return s.Param
}

func (s *AiccsSmartCallOperateRequest) GetProdCode() *string {
	return s.ProdCode
}

func (s *AiccsSmartCallOperateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AiccsSmartCallOperateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AiccsSmartCallOperateRequest) SetCallId(v string) *AiccsSmartCallOperateRequest {
	s.CallId = &v
	return s
}

func (s *AiccsSmartCallOperateRequest) SetCommand(v string) *AiccsSmartCallOperateRequest {
	s.Command = &v
	return s
}

func (s *AiccsSmartCallOperateRequest) SetOwnerId(v int64) *AiccsSmartCallOperateRequest {
	s.OwnerId = &v
	return s
}

func (s *AiccsSmartCallOperateRequest) SetParam(v string) *AiccsSmartCallOperateRequest {
	s.Param = &v
	return s
}

func (s *AiccsSmartCallOperateRequest) SetProdCode(v string) *AiccsSmartCallOperateRequest {
	s.ProdCode = &v
	return s
}

func (s *AiccsSmartCallOperateRequest) SetResourceOwnerAccount(v string) *AiccsSmartCallOperateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AiccsSmartCallOperateRequest) SetResourceOwnerId(v int64) *AiccsSmartCallOperateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AiccsSmartCallOperateRequest) Validate() error {
	return dara.Validate(s)
}
