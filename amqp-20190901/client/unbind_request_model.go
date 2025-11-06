// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBindingKey(v string) *UnbindRequest
	GetBindingKey() *string
	SetBindingType(v int32) *UnbindRequest
	GetBindingType() *int32
	SetConsoleSessionId(v string) *UnbindRequest
	GetConsoleSessionId() *string
	SetDstName(v string) *UnbindRequest
	GetDstName() *string
	SetInstanceId(v string) *UnbindRequest
	GetInstanceId() *string
	SetSrcName(v string) *UnbindRequest
	GetSrcName() *string
	SetVhostName(v string) *UnbindRequest
	GetVhostName() *string
}

type UnbindRequest struct {
	BindingKey *string `json:"BindingKey,omitempty" xml:"BindingKey,omitempty"`
	// This parameter is required.
	BindingType      *int32  `json:"BindingType,omitempty" xml:"BindingType,omitempty"`
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	// This parameter is required.
	DstName    *string `json:"DstName,omitempty" xml:"DstName,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	SrcName *string `json:"SrcName,omitempty" xml:"SrcName,omitempty"`
	// This parameter is required.
	VhostName *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s UnbindRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindRequest) GoString() string {
	return s.String()
}

func (s *UnbindRequest) GetBindingKey() *string {
	return s.BindingKey
}

func (s *UnbindRequest) GetBindingType() *int32 {
	return s.BindingType
}

func (s *UnbindRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *UnbindRequest) GetDstName() *string {
	return s.DstName
}

func (s *UnbindRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UnbindRequest) GetSrcName() *string {
	return s.SrcName
}

func (s *UnbindRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *UnbindRequest) SetBindingKey(v string) *UnbindRequest {
	s.BindingKey = &v
	return s
}

func (s *UnbindRequest) SetBindingType(v int32) *UnbindRequest {
	s.BindingType = &v
	return s
}

func (s *UnbindRequest) SetConsoleSessionId(v string) *UnbindRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *UnbindRequest) SetDstName(v string) *UnbindRequest {
	s.DstName = &v
	return s
}

func (s *UnbindRequest) SetInstanceId(v string) *UnbindRequest {
	s.InstanceId = &v
	return s
}

func (s *UnbindRequest) SetSrcName(v string) *UnbindRequest {
	s.SrcName = &v
	return s
}

func (s *UnbindRequest) SetVhostName(v string) *UnbindRequest {
	s.VhostName = &v
	return s
}

func (s *UnbindRequest) Validate() error {
	return dara.Validate(s)
}
