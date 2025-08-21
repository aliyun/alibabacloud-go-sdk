// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApp(v string) *CreateGroupRequest
	GetApp() *string
	SetCallback(v string) *CreateGroupRequest
	GetCallback() *string
	SetDescription(v string) *CreateGroupRequest
	GetDescription() *string
	SetInProtocol(v string) *CreateGroupRequest
	GetInProtocol() *string
	SetLazyPull(v bool) *CreateGroupRequest
	GetLazyPull() *bool
	SetName(v string) *CreateGroupRequest
	GetName() *string
	SetOutProtocol(v string) *CreateGroupRequest
	GetOutProtocol() *string
	SetOwnerId(v int64) *CreateGroupRequest
	GetOwnerId() *int64
	SetPlayDomain(v string) *CreateGroupRequest
	GetPlayDomain() *string
	SetPushDomain(v string) *CreateGroupRequest
	GetPushDomain() *string
	SetRegion(v string) *CreateGroupRequest
	GetRegion() *string
}

type CreateGroupRequest struct {
	// example:
	//
	// live
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// example:
	//
	// http://example.com/callback
	Callback    *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rtmp
	InProtocol *string `json:"InProtocol,omitempty" xml:"InProtocol,omitempty"`
	// example:
	//
	// false
	LazyPull *bool `json:"LazyPull,omitempty" xml:"LazyPull,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// flv,rtmp
	OutProtocol *string `json:"OutProtocol,omitempty" xml:"OutProtocol,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// demo.aliyundoc.com
	PlayDomain *string `json:"PlayDomain,omitempty" xml:"PlayDomain,omitempty"`
	// example:
	//
	// example.aliyundoc.com
	PushDomain *string `json:"PushDomain,omitempty" xml:"PushDomain,omitempty"`
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s CreateGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupRequest) GetApp() *string {
	return s.App
}

func (s *CreateGroupRequest) GetCallback() *string {
	return s.Callback
}

func (s *CreateGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGroupRequest) GetInProtocol() *string {
	return s.InProtocol
}

func (s *CreateGroupRequest) GetLazyPull() *bool {
	return s.LazyPull
}

func (s *CreateGroupRequest) GetName() *string {
	return s.Name
}

func (s *CreateGroupRequest) GetOutProtocol() *string {
	return s.OutProtocol
}

func (s *CreateGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateGroupRequest) GetPlayDomain() *string {
	return s.PlayDomain
}

func (s *CreateGroupRequest) GetPushDomain() *string {
	return s.PushDomain
}

func (s *CreateGroupRequest) GetRegion() *string {
	return s.Region
}

func (s *CreateGroupRequest) SetApp(v string) *CreateGroupRequest {
	s.App = &v
	return s
}

func (s *CreateGroupRequest) SetCallback(v string) *CreateGroupRequest {
	s.Callback = &v
	return s
}

func (s *CreateGroupRequest) SetDescription(v string) *CreateGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateGroupRequest) SetInProtocol(v string) *CreateGroupRequest {
	s.InProtocol = &v
	return s
}

func (s *CreateGroupRequest) SetLazyPull(v bool) *CreateGroupRequest {
	s.LazyPull = &v
	return s
}

func (s *CreateGroupRequest) SetName(v string) *CreateGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateGroupRequest) SetOutProtocol(v string) *CreateGroupRequest {
	s.OutProtocol = &v
	return s
}

func (s *CreateGroupRequest) SetOwnerId(v int64) *CreateGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateGroupRequest) SetPlayDomain(v string) *CreateGroupRequest {
	s.PlayDomain = &v
	return s
}

func (s *CreateGroupRequest) SetPushDomain(v string) *CreateGroupRequest {
	s.PushDomain = &v
	return s
}

func (s *CreateGroupRequest) SetRegion(v string) *CreateGroupRequest {
	s.Region = &v
	return s
}

func (s *CreateGroupRequest) Validate() error {
	return dara.Validate(s)
}
