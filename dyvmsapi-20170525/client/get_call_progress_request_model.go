// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCallProgressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *GetCallProgressRequest
	GetCallId() *string
	SetCalledNum(v string) *GetCallProgressRequest
	GetCalledNum() *string
	SetOwnerId(v int64) *GetCallProgressRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetCallProgressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetCallProgressRequest
	GetResourceOwnerId() *int64
}

type GetCallProgressRequest struct {
	// example:
	//
	// 示例值
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	CalledNum            *string `json:"CalledNum,omitempty" xml:"CalledNum,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetCallProgressRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCallProgressRequest) GoString() string {
	return s.String()
}

func (s *GetCallProgressRequest) GetCallId() *string {
	return s.CallId
}

func (s *GetCallProgressRequest) GetCalledNum() *string {
	return s.CalledNum
}

func (s *GetCallProgressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetCallProgressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetCallProgressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetCallProgressRequest) SetCallId(v string) *GetCallProgressRequest {
	s.CallId = &v
	return s
}

func (s *GetCallProgressRequest) SetCalledNum(v string) *GetCallProgressRequest {
	s.CalledNum = &v
	return s
}

func (s *GetCallProgressRequest) SetOwnerId(v int64) *GetCallProgressRequest {
	s.OwnerId = &v
	return s
}

func (s *GetCallProgressRequest) SetResourceOwnerAccount(v string) *GetCallProgressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetCallProgressRequest) SetResourceOwnerId(v int64) *GetCallProgressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetCallProgressRequest) Validate() error {
	return dara.Validate(s)
}
