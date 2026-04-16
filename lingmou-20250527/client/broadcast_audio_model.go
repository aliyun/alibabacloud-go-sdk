// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBroadcastAudio interface {
	dara.Model
	String() string
	GoString() string
	SetAudioLength(v int32) *BroadcastAudio
	GetAudioLength() *int32
	SetCreateTime(v string) *BroadcastAudio
	GetCreateTime() *string
	SetErrorCode(v string) *BroadcastAudio
	GetErrorCode() *string
	SetId(v string) *BroadcastAudio
	GetId() *string
	SetModifiedTime(v string) *BroadcastAudio
	GetModifiedTime() *string
	SetName(v string) *BroadcastAudio
	GetName() *string
	SetStatus(v string) *BroadcastAudio
	GetStatus() *string
}

type BroadcastAudio struct {
	// example:
	//
	// 10
	AudioLength *int32 `json:"audioLength,omitempty" xml:"audioLength,omitempty"`
	// example:
	//
	// 2026-01-22T01:59:03
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	ErrorCode  *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// M1Ju6XhHog_e-lSeb80Slp9g
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 2026-01-22T01:59:03
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s BroadcastAudio) String() string {
	return dara.Prettify(s)
}

func (s BroadcastAudio) GoString() string {
	return s.String()
}

func (s *BroadcastAudio) GetAudioLength() *int32 {
	return s.AudioLength
}

func (s *BroadcastAudio) GetCreateTime() *string {
	return s.CreateTime
}

func (s *BroadcastAudio) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *BroadcastAudio) GetId() *string {
	return s.Id
}

func (s *BroadcastAudio) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *BroadcastAudio) GetName() *string {
	return s.Name
}

func (s *BroadcastAudio) GetStatus() *string {
	return s.Status
}

func (s *BroadcastAudio) SetAudioLength(v int32) *BroadcastAudio {
	s.AudioLength = &v
	return s
}

func (s *BroadcastAudio) SetCreateTime(v string) *BroadcastAudio {
	s.CreateTime = &v
	return s
}

func (s *BroadcastAudio) SetErrorCode(v string) *BroadcastAudio {
	s.ErrorCode = &v
	return s
}

func (s *BroadcastAudio) SetId(v string) *BroadcastAudio {
	s.Id = &v
	return s
}

func (s *BroadcastAudio) SetModifiedTime(v string) *BroadcastAudio {
	s.ModifiedTime = &v
	return s
}

func (s *BroadcastAudio) SetName(v string) *BroadcastAudio {
	s.Name = &v
	return s
}

func (s *BroadcastAudio) SetStatus(v string) *BroadcastAudio {
	s.Status = &v
	return s
}

func (s *BroadcastAudio) Validate() error {
	return dara.Validate(s)
}
