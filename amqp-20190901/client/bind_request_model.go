// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArgument(v string) *BindRequest
	GetArgument() *string
	SetBindingKey(v string) *BindRequest
	GetBindingKey() *string
	SetBindingType(v int32) *BindRequest
	GetBindingType() *int32
	SetConsoleSessionId(v string) *BindRequest
	GetConsoleSessionId() *string
	SetDstName(v string) *BindRequest
	GetDstName() *string
	SetInstanceId(v string) *BindRequest
	GetInstanceId() *string
	SetSrcName(v string) *BindRequest
	GetSrcName() *string
	SetVhostName(v string) *BindRequest
	GetVhostName() *string
}

type BindRequest struct {
	Argument   *string `json:"Argument,omitempty" xml:"Argument,omitempty"`
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

func (s BindRequest) String() string {
	return dara.Prettify(s)
}

func (s BindRequest) GoString() string {
	return s.String()
}

func (s *BindRequest) GetArgument() *string {
	return s.Argument
}

func (s *BindRequest) GetBindingKey() *string {
	return s.BindingKey
}

func (s *BindRequest) GetBindingType() *int32 {
	return s.BindingType
}

func (s *BindRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *BindRequest) GetDstName() *string {
	return s.DstName
}

func (s *BindRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BindRequest) GetSrcName() *string {
	return s.SrcName
}

func (s *BindRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *BindRequest) SetArgument(v string) *BindRequest {
	s.Argument = &v
	return s
}

func (s *BindRequest) SetBindingKey(v string) *BindRequest {
	s.BindingKey = &v
	return s
}

func (s *BindRequest) SetBindingType(v int32) *BindRequest {
	s.BindingType = &v
	return s
}

func (s *BindRequest) SetConsoleSessionId(v string) *BindRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *BindRequest) SetDstName(v string) *BindRequest {
	s.DstName = &v
	return s
}

func (s *BindRequest) SetInstanceId(v string) *BindRequest {
	s.InstanceId = &v
	return s
}

func (s *BindRequest) SetSrcName(v string) *BindRequest {
	s.SrcName = &v
	return s
}

func (s *BindRequest) SetVhostName(v string) *BindRequest {
	s.VhostName = &v
	return s
}

func (s *BindRequest) Validate() error {
	return dara.Validate(s)
}
