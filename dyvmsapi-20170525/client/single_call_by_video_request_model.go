// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSingleCallByVideoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalledNumber(v string) *SingleCallByVideoRequest
	GetCalledNumber() *string
	SetCalledShowNumber(v string) *SingleCallByVideoRequest
	GetCalledShowNumber() *string
	SetOutId(v string) *SingleCallByVideoRequest
	GetOutId() *string
	SetOwnerId(v int64) *SingleCallByVideoRequest
	GetOwnerId() *int64
	SetPlayTimes(v int32) *SingleCallByVideoRequest
	GetPlayTimes() *int32
	SetResourceOwnerAccount(v string) *SingleCallByVideoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SingleCallByVideoRequest
	GetResourceOwnerId() *int64
	SetSpeed(v int32) *SingleCallByVideoRequest
	GetSpeed() *int32
	SetVideoCode(v string) *SingleCallByVideoRequest
	GetVideoCode() *string
	SetVoiceCode(v string) *SingleCallByVideoRequest
	GetVoiceCode() *string
	SetVolume(v int32) *SingleCallByVideoRequest
	GetVolume() *int32
}

type SingleCallByVideoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1590****000
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// example:
	//
	// 0571****5678
	CalledShowNumber *string `json:"CalledShowNumber,omitempty" xml:"CalledShowNumber,omitempty"`
	// example:
	//
	// abcdefgh
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 3
	PlayTimes            *int32  `json:"PlayTimes,omitempty" xml:"PlayTimes,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 100
	Speed *int32 `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// This parameter is required.
	VideoCode *string `json:"VideoCode,omitempty" xml:"VideoCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2d4c-4e78-8d2a-afbb06cf****.wav
	VoiceCode *string `json:"VoiceCode,omitempty" xml:"VoiceCode,omitempty"`
	// example:
	//
	// 100
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s SingleCallByVideoRequest) String() string {
	return dara.Prettify(s)
}

func (s SingleCallByVideoRequest) GoString() string {
	return s.String()
}

func (s *SingleCallByVideoRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *SingleCallByVideoRequest) GetCalledShowNumber() *string {
	return s.CalledShowNumber
}

func (s *SingleCallByVideoRequest) GetOutId() *string {
	return s.OutId
}

func (s *SingleCallByVideoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SingleCallByVideoRequest) GetPlayTimes() *int32 {
	return s.PlayTimes
}

func (s *SingleCallByVideoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SingleCallByVideoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SingleCallByVideoRequest) GetSpeed() *int32 {
	return s.Speed
}

func (s *SingleCallByVideoRequest) GetVideoCode() *string {
	return s.VideoCode
}

func (s *SingleCallByVideoRequest) GetVoiceCode() *string {
	return s.VoiceCode
}

func (s *SingleCallByVideoRequest) GetVolume() *int32 {
	return s.Volume
}

func (s *SingleCallByVideoRequest) SetCalledNumber(v string) *SingleCallByVideoRequest {
	s.CalledNumber = &v
	return s
}

func (s *SingleCallByVideoRequest) SetCalledShowNumber(v string) *SingleCallByVideoRequest {
	s.CalledShowNumber = &v
	return s
}

func (s *SingleCallByVideoRequest) SetOutId(v string) *SingleCallByVideoRequest {
	s.OutId = &v
	return s
}

func (s *SingleCallByVideoRequest) SetOwnerId(v int64) *SingleCallByVideoRequest {
	s.OwnerId = &v
	return s
}

func (s *SingleCallByVideoRequest) SetPlayTimes(v int32) *SingleCallByVideoRequest {
	s.PlayTimes = &v
	return s
}

func (s *SingleCallByVideoRequest) SetResourceOwnerAccount(v string) *SingleCallByVideoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SingleCallByVideoRequest) SetResourceOwnerId(v int64) *SingleCallByVideoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SingleCallByVideoRequest) SetSpeed(v int32) *SingleCallByVideoRequest {
	s.Speed = &v
	return s
}

func (s *SingleCallByVideoRequest) SetVideoCode(v string) *SingleCallByVideoRequest {
	s.VideoCode = &v
	return s
}

func (s *SingleCallByVideoRequest) SetVoiceCode(v string) *SingleCallByVideoRequest {
	s.VoiceCode = &v
	return s
}

func (s *SingleCallByVideoRequest) SetVolume(v int32) *SingleCallByVideoRequest {
	s.Volume = &v
	return s
}

func (s *SingleCallByVideoRequest) Validate() error {
	return dara.Validate(s)
}
