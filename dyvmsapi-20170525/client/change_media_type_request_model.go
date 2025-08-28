// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeMediaTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *ChangeMediaTypeRequest
	GetCallId() *string
	SetCalledNum(v string) *ChangeMediaTypeRequest
	GetCalledNum() *string
	SetMediaType(v string) *ChangeMediaTypeRequest
	GetMediaType() *string
	SetOutId(v string) *ChangeMediaTypeRequest
	GetOutId() *string
	SetOwnerId(v int64) *ChangeMediaTypeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ChangeMediaTypeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ChangeMediaTypeRequest
	GetResourceOwnerId() *int64
}

type ChangeMediaTypeRequest struct {
	// example:
	//
	// 示例值示例值示例值
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	CalledNum *string `json:"CalledNum,omitempty" xml:"CalledNum,omitempty"`
	// example:
	//
	// 示例值示例值
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// example:
	//
	// 示例值示例值
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ChangeMediaTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeMediaTypeRequest) GoString() string {
	return s.String()
}

func (s *ChangeMediaTypeRequest) GetCallId() *string {
	return s.CallId
}

func (s *ChangeMediaTypeRequest) GetCalledNum() *string {
	return s.CalledNum
}

func (s *ChangeMediaTypeRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *ChangeMediaTypeRequest) GetOutId() *string {
	return s.OutId
}

func (s *ChangeMediaTypeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ChangeMediaTypeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ChangeMediaTypeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ChangeMediaTypeRequest) SetCallId(v string) *ChangeMediaTypeRequest {
	s.CallId = &v
	return s
}

func (s *ChangeMediaTypeRequest) SetCalledNum(v string) *ChangeMediaTypeRequest {
	s.CalledNum = &v
	return s
}

func (s *ChangeMediaTypeRequest) SetMediaType(v string) *ChangeMediaTypeRequest {
	s.MediaType = &v
	return s
}

func (s *ChangeMediaTypeRequest) SetOutId(v string) *ChangeMediaTypeRequest {
	s.OutId = &v
	return s
}

func (s *ChangeMediaTypeRequest) SetOwnerId(v int64) *ChangeMediaTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *ChangeMediaTypeRequest) SetResourceOwnerAccount(v string) *ChangeMediaTypeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ChangeMediaTypeRequest) SetResourceOwnerId(v int64) *ChangeMediaTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ChangeMediaTypeRequest) Validate() error {
	return dara.Validate(s)
}
