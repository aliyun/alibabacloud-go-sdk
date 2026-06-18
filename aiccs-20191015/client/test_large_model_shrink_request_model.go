// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestLargeModelShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseModelShrink(v string) *TestLargeModelShrinkRequest
	GetBaseModelShrink() *string
	SetModelCode(v string) *TestLargeModelShrinkRequest
	GetModelCode() *string
	SetOwnerId(v int64) *TestLargeModelShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *TestLargeModelShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *TestLargeModelShrinkRequest
	GetResourceOwnerId() *int64
	SetUserDialogContent(v string) *TestLargeModelShrinkRequest
	GetUserDialogContent() *string
}

type TestLargeModelShrinkRequest struct {
	// The base models.
	BaseModelShrink *string `json:"BaseModel,omitempty" xml:"BaseModel,omitempty"`
	// The ID of the test scenario.
	//
	// example:
	//
	// 1232
	ModelCode            *string `json:"ModelCode,omitempty" xml:"ModelCode,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The user dialog content.
	//
	// example:
	//
	// 你好。
	UserDialogContent *string `json:"UserDialogContent,omitempty" xml:"UserDialogContent,omitempty"`
}

func (s TestLargeModelShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s TestLargeModelShrinkRequest) GoString() string {
	return s.String()
}

func (s *TestLargeModelShrinkRequest) GetBaseModelShrink() *string {
	return s.BaseModelShrink
}

func (s *TestLargeModelShrinkRequest) GetModelCode() *string {
	return s.ModelCode
}

func (s *TestLargeModelShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TestLargeModelShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TestLargeModelShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *TestLargeModelShrinkRequest) GetUserDialogContent() *string {
	return s.UserDialogContent
}

func (s *TestLargeModelShrinkRequest) SetBaseModelShrink(v string) *TestLargeModelShrinkRequest {
	s.BaseModelShrink = &v
	return s
}

func (s *TestLargeModelShrinkRequest) SetModelCode(v string) *TestLargeModelShrinkRequest {
	s.ModelCode = &v
	return s
}

func (s *TestLargeModelShrinkRequest) SetOwnerId(v int64) *TestLargeModelShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *TestLargeModelShrinkRequest) SetResourceOwnerAccount(v string) *TestLargeModelShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TestLargeModelShrinkRequest) SetResourceOwnerId(v int64) *TestLargeModelShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TestLargeModelShrinkRequest) SetUserDialogContent(v string) *TestLargeModelShrinkRequest {
	s.UserDialogContent = &v
	return s
}

func (s *TestLargeModelShrinkRequest) Validate() error {
	return dara.Validate(s)
}
