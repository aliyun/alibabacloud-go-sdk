// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestLargeModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseModel(v []*string) *TestLargeModelRequest
	GetBaseModel() []*string
	SetModelCode(v string) *TestLargeModelRequest
	GetModelCode() *string
	SetOwnerId(v int64) *TestLargeModelRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *TestLargeModelRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *TestLargeModelRequest
	GetResourceOwnerId() *int64
	SetUserDialogContent(v string) *TestLargeModelRequest
	GetUserDialogContent() *string
}

type TestLargeModelRequest struct {
	// 基础模型
	BaseModel []*string `json:"BaseModel,omitempty" xml:"BaseModel,omitempty" type:"Repeated"`
	// 场景ID
	//
	// example:
	//
	// 1232
	ModelCode            *string `json:"ModelCode,omitempty" xml:"ModelCode,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 用户对话内容
	//
	// example:
	//
	// 示例值示例值示例值
	UserDialogContent *string `json:"UserDialogContent,omitempty" xml:"UserDialogContent,omitempty"`
}

func (s TestLargeModelRequest) String() string {
	return dara.Prettify(s)
}

func (s TestLargeModelRequest) GoString() string {
	return s.String()
}

func (s *TestLargeModelRequest) GetBaseModel() []*string {
	return s.BaseModel
}

func (s *TestLargeModelRequest) GetModelCode() *string {
	return s.ModelCode
}

func (s *TestLargeModelRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TestLargeModelRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TestLargeModelRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *TestLargeModelRequest) GetUserDialogContent() *string {
	return s.UserDialogContent
}

func (s *TestLargeModelRequest) SetBaseModel(v []*string) *TestLargeModelRequest {
	s.BaseModel = v
	return s
}

func (s *TestLargeModelRequest) SetModelCode(v string) *TestLargeModelRequest {
	s.ModelCode = &v
	return s
}

func (s *TestLargeModelRequest) SetOwnerId(v int64) *TestLargeModelRequest {
	s.OwnerId = &v
	return s
}

func (s *TestLargeModelRequest) SetResourceOwnerAccount(v string) *TestLargeModelRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TestLargeModelRequest) SetResourceOwnerId(v int64) *TestLargeModelRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TestLargeModelRequest) SetUserDialogContent(v string) *TestLargeModelRequest {
	s.UserDialogContent = &v
	return s
}

func (s *TestLargeModelRequest) Validate() error {
	return dara.Validate(s)
}
