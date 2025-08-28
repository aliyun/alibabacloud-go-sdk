// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeVideoFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *ResumeVideoFileRequest
	GetCallId() *string
	SetCalledNumber(v string) *ResumeVideoFileRequest
	GetCalledNumber() *string
	SetOwnerId(v int64) *ResumeVideoFileRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ResumeVideoFileRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ResumeVideoFileRequest
	GetResourceOwnerId() *int64
}

type ResumeVideoFileRequest struct {
	// example:
	//
	// 100001616500^1000018****
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	CalledNumber         *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ResumeVideoFileRequest) String() string {
	return dara.Prettify(s)
}

func (s ResumeVideoFileRequest) GoString() string {
	return s.String()
}

func (s *ResumeVideoFileRequest) GetCallId() *string {
	return s.CallId
}

func (s *ResumeVideoFileRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *ResumeVideoFileRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ResumeVideoFileRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ResumeVideoFileRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ResumeVideoFileRequest) SetCallId(v string) *ResumeVideoFileRequest {
	s.CallId = &v
	return s
}

func (s *ResumeVideoFileRequest) SetCalledNumber(v string) *ResumeVideoFileRequest {
	s.CalledNumber = &v
	return s
}

func (s *ResumeVideoFileRequest) SetOwnerId(v int64) *ResumeVideoFileRequest {
	s.OwnerId = &v
	return s
}

func (s *ResumeVideoFileRequest) SetResourceOwnerAccount(v string) *ResumeVideoFileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ResumeVideoFileRequest) SetResourceOwnerId(v int64) *ResumeVideoFileRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ResumeVideoFileRequest) Validate() error {
	return dara.Validate(s)
}
