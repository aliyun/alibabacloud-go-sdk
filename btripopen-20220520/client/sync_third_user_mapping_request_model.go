// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncThirdUserMappingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStatus(v int32) *SyncThirdUserMappingRequest
	GetStatus() *int32
	SetThirdChannelType(v string) *SyncThirdUserMappingRequest
	GetThirdChannelType() *string
	SetThirdUserId(v string) *SyncThirdUserMappingRequest
	GetThirdUserId() *string
	SetUserId(v string) *SyncThirdUserMappingRequest
	GetUserId() *string
}

type SyncThirdUserMappingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// weCom
	ThirdChannelType *string `json:"third_channel_type,omitempty" xml:"third_channel_type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	ThirdUserId *string `json:"third_user_id,omitempty" xml:"third_user_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s SyncThirdUserMappingRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncThirdUserMappingRequest) GoString() string {
	return s.String()
}

func (s *SyncThirdUserMappingRequest) GetStatus() *int32 {
	return s.Status
}

func (s *SyncThirdUserMappingRequest) GetThirdChannelType() *string {
	return s.ThirdChannelType
}

func (s *SyncThirdUserMappingRequest) GetThirdUserId() *string {
	return s.ThirdUserId
}

func (s *SyncThirdUserMappingRequest) GetUserId() *string {
	return s.UserId
}

func (s *SyncThirdUserMappingRequest) SetStatus(v int32) *SyncThirdUserMappingRequest {
	s.Status = &v
	return s
}

func (s *SyncThirdUserMappingRequest) SetThirdChannelType(v string) *SyncThirdUserMappingRequest {
	s.ThirdChannelType = &v
	return s
}

func (s *SyncThirdUserMappingRequest) SetThirdUserId(v string) *SyncThirdUserMappingRequest {
	s.ThirdUserId = &v
	return s
}

func (s *SyncThirdUserMappingRequest) SetUserId(v string) *SyncThirdUserMappingRequest {
	s.UserId = &v
	return s
}

func (s *SyncThirdUserMappingRequest) Validate() error {
	return dara.Validate(s)
}
