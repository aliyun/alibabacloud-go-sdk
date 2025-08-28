// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseVideoFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *PauseVideoFileRequest
	GetCallId() *string
	SetCalledNumber(v string) *PauseVideoFileRequest
	GetCalledNumber() *string
	SetOwnerId(v int64) *PauseVideoFileRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *PauseVideoFileRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *PauseVideoFileRequest
	GetResourceOwnerId() *int64
}

type PauseVideoFileRequest struct {
	// example:
	//
	// 116012854210^10281427****
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	CalledNumber         *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s PauseVideoFileRequest) String() string {
	return dara.Prettify(s)
}

func (s PauseVideoFileRequest) GoString() string {
	return s.String()
}

func (s *PauseVideoFileRequest) GetCallId() *string {
	return s.CallId
}

func (s *PauseVideoFileRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *PauseVideoFileRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *PauseVideoFileRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *PauseVideoFileRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *PauseVideoFileRequest) SetCallId(v string) *PauseVideoFileRequest {
	s.CallId = &v
	return s
}

func (s *PauseVideoFileRequest) SetCalledNumber(v string) *PauseVideoFileRequest {
	s.CalledNumber = &v
	return s
}

func (s *PauseVideoFileRequest) SetOwnerId(v int64) *PauseVideoFileRequest {
	s.OwnerId = &v
	return s
}

func (s *PauseVideoFileRequest) SetResourceOwnerAccount(v string) *PauseVideoFileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PauseVideoFileRequest) SetResourceOwnerId(v int64) *PauseVideoFileRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *PauseVideoFileRequest) Validate() error {
	return dara.Validate(s)
}
