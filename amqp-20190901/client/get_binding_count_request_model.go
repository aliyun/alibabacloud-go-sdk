// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBindingCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBindingType(v int32) *GetBindingCountRequest
	GetBindingType() *int32
	SetConsoleSessionId(v string) *GetBindingCountRequest
	GetConsoleSessionId() *string
	SetInstanceId(v string) *GetBindingCountRequest
	GetInstanceId() *string
	SetResourceName(v string) *GetBindingCountRequest
	GetResourceName() *string
	SetUpstream(v bool) *GetBindingCountRequest
	GetUpstream() *bool
	SetVhostName(v string) *GetBindingCountRequest
	GetVhostName() *string
}

type GetBindingCountRequest struct {
	// This parameter is required.
	BindingType      *int32  `json:"BindingType,omitempty" xml:"BindingType,omitempty"`
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// This parameter is required.
	Upstream *bool `json:"Upstream,omitempty" xml:"Upstream,omitempty"`
	// This parameter is required.
	VhostName *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s GetBindingCountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBindingCountRequest) GoString() string {
	return s.String()
}

func (s *GetBindingCountRequest) GetBindingType() *int32 {
	return s.BindingType
}

func (s *GetBindingCountRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *GetBindingCountRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetBindingCountRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *GetBindingCountRequest) GetUpstream() *bool {
	return s.Upstream
}

func (s *GetBindingCountRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *GetBindingCountRequest) SetBindingType(v int32) *GetBindingCountRequest {
	s.BindingType = &v
	return s
}

func (s *GetBindingCountRequest) SetConsoleSessionId(v string) *GetBindingCountRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *GetBindingCountRequest) SetInstanceId(v string) *GetBindingCountRequest {
	s.InstanceId = &v
	return s
}

func (s *GetBindingCountRequest) SetResourceName(v string) *GetBindingCountRequest {
	s.ResourceName = &v
	return s
}

func (s *GetBindingCountRequest) SetUpstream(v bool) *GetBindingCountRequest {
	s.Upstream = &v
	return s
}

func (s *GetBindingCountRequest) SetVhostName(v string) *GetBindingCountRequest {
	s.VhostName = &v
	return s
}

func (s *GetBindingCountRequest) Validate() error {
	return dara.Validate(s)
}
