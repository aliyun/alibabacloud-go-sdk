// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSeekVideoFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *SeekVideoFileRequest
	GetCallId() *string
	SetCalledNumber(v string) *SeekVideoFileRequest
	GetCalledNumber() *string
	SetOwnerId(v int64) *SeekVideoFileRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *SeekVideoFileRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SeekVideoFileRequest
	GetResourceOwnerId() *int64
	SetSeekTimes(v int64) *SeekVideoFileRequest
	GetSeekTimes() *int64
}

type SeekVideoFileRequest struct {
	// 呼叫唯一ID
	//
	// example:
	//
	// 示例值
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// 被叫号码
	//
	// example:
	//
	// 示例值示例值示例值
	CalledNumber         *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 快进或快退值，负数为快退，单位秒
	//
	// example:
	//
	// 92
	SeekTimes *int64 `json:"SeekTimes,omitempty" xml:"SeekTimes,omitempty"`
}

func (s SeekVideoFileRequest) String() string {
	return dara.Prettify(s)
}

func (s SeekVideoFileRequest) GoString() string {
	return s.String()
}

func (s *SeekVideoFileRequest) GetCallId() *string {
	return s.CallId
}

func (s *SeekVideoFileRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *SeekVideoFileRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SeekVideoFileRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SeekVideoFileRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SeekVideoFileRequest) GetSeekTimes() *int64 {
	return s.SeekTimes
}

func (s *SeekVideoFileRequest) SetCallId(v string) *SeekVideoFileRequest {
	s.CallId = &v
	return s
}

func (s *SeekVideoFileRequest) SetCalledNumber(v string) *SeekVideoFileRequest {
	s.CalledNumber = &v
	return s
}

func (s *SeekVideoFileRequest) SetOwnerId(v int64) *SeekVideoFileRequest {
	s.OwnerId = &v
	return s
}

func (s *SeekVideoFileRequest) SetResourceOwnerAccount(v string) *SeekVideoFileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SeekVideoFileRequest) SetResourceOwnerId(v int64) *SeekVideoFileRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SeekVideoFileRequest) SetSeekTimes(v int64) *SeekVideoFileRequest {
	s.SeekTimes = &v
	return s
}

func (s *SeekVideoFileRequest) Validate() error {
	return dara.Validate(s)
}
