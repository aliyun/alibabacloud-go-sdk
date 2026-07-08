// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallback(v string) *ModifyGroupRequest
	GetCallback() *string
	SetDescription(v string) *ModifyGroupRequest
	GetDescription() *string
	SetEnabled(v bool) *ModifyGroupRequest
	GetEnabled() *bool
	SetId(v string) *ModifyGroupRequest
	GetId() *string
	SetInProtocol(v string) *ModifyGroupRequest
	GetInProtocol() *string
	SetLazyPull(v bool) *ModifyGroupRequest
	GetLazyPull() *bool
	SetName(v string) *ModifyGroupRequest
	GetName() *string
	SetOutProtocol(v string) *ModifyGroupRequest
	GetOutProtocol() *string
	SetOwnerId(v int64) *ModifyGroupRequest
	GetOwnerId() *int64
	SetPlayDomain(v string) *ModifyGroupRequest
	GetPlayDomain() *string
	SetPushDomain(v string) *ModifyGroupRequest
	GetPushDomain() *string
	SetRegion(v string) *ModifyGroupRequest
	GetRegion() *string
}

type ModifyGroupRequest struct {
	// Callback URL for device or stream status updates in the group.
	//
	// example:
	//
	// http://example.com/callback
	Callback *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	// Description of the space.
	//
	// example:
	//
	// 用于测试使用
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the space is enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The ID of the space.
	//
	// This parameter is required.
	//
	// example:
	//
	// 32388487739092994-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Ingest protocol used by the group. Valid values:
	//
	// - gb28181
	//
	// - rtmp
	//
	// example:
	//
	// gb28181
	InProtocol *string `json:"InProtocol,omitempty" xml:"InProtocol,omitempty"`
	// Whether on-demand stream pulling is enabled. Valid values:
	//
	// - false (default)
	//
	// - true
	//
	// example:
	//
	// false
	LazyPull *bool `json:"LazyPull,omitempty" xml:"LazyPull,omitempty"`
	// Space name.
	//
	// example:
	//
	// myGroup
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Playback protocols used by the group. Separate multiple values with commas. Valid values:
	//
	// - flv
	//
	// - hls
	//
	// - rtmp
	//
	// example:
	//
	// flv,rtmp,hls
	OutProtocol *string `json:"OutProtocol,omitempty" xml:"OutProtocol,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Streaming domain used by the group.
	//
	// example:
	//
	// myplay.com
	PlayDomain *string `json:"PlayDomain,omitempty" xml:"PlayDomain,omitempty"`
	// Ingest domain used by the group. Applies only to groups that use the RTMP ingest protocol.
	//
	// example:
	//
	// mypush.com
	PushDomain *string `json:"PushDomain,omitempty" xml:"PushDomain,omitempty"`
	// The region where the space is located. This region serves as the service center.
	//
	// example:
	//
	// cn-qingdao
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s ModifyGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyGroupRequest) GetCallback() *string {
	return s.Callback
}

func (s *ModifyGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyGroupRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *ModifyGroupRequest) GetId() *string {
	return s.Id
}

func (s *ModifyGroupRequest) GetInProtocol() *string {
	return s.InProtocol
}

func (s *ModifyGroupRequest) GetLazyPull() *bool {
	return s.LazyPull
}

func (s *ModifyGroupRequest) GetName() *string {
	return s.Name
}

func (s *ModifyGroupRequest) GetOutProtocol() *string {
	return s.OutProtocol
}

func (s *ModifyGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyGroupRequest) GetPlayDomain() *string {
	return s.PlayDomain
}

func (s *ModifyGroupRequest) GetPushDomain() *string {
	return s.PushDomain
}

func (s *ModifyGroupRequest) GetRegion() *string {
	return s.Region
}

func (s *ModifyGroupRequest) SetCallback(v string) *ModifyGroupRequest {
	s.Callback = &v
	return s
}

func (s *ModifyGroupRequest) SetDescription(v string) *ModifyGroupRequest {
	s.Description = &v
	return s
}

func (s *ModifyGroupRequest) SetEnabled(v bool) *ModifyGroupRequest {
	s.Enabled = &v
	return s
}

func (s *ModifyGroupRequest) SetId(v string) *ModifyGroupRequest {
	s.Id = &v
	return s
}

func (s *ModifyGroupRequest) SetInProtocol(v string) *ModifyGroupRequest {
	s.InProtocol = &v
	return s
}

func (s *ModifyGroupRequest) SetLazyPull(v bool) *ModifyGroupRequest {
	s.LazyPull = &v
	return s
}

func (s *ModifyGroupRequest) SetName(v string) *ModifyGroupRequest {
	s.Name = &v
	return s
}

func (s *ModifyGroupRequest) SetOutProtocol(v string) *ModifyGroupRequest {
	s.OutProtocol = &v
	return s
}

func (s *ModifyGroupRequest) SetOwnerId(v int64) *ModifyGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyGroupRequest) SetPlayDomain(v string) *ModifyGroupRequest {
	s.PlayDomain = &v
	return s
}

func (s *ModifyGroupRequest) SetPushDomain(v string) *ModifyGroupRequest {
	s.PushDomain = &v
	return s
}

func (s *ModifyGroupRequest) SetRegion(v string) *ModifyGroupRequest {
	s.Region = &v
	return s
}

func (s *ModifyGroupRequest) Validate() error {
	return dara.Validate(s)
}
