// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSandboxHealthCheckOut interface {
	dara.Model
	String() string
	GoString() string
	SetStatus(v string) *SandboxHealthCheckOut
	GetStatus() *string
}

type SandboxHealthCheckOut struct {
	// 健康状态，OK表示健康
	//
	// This parameter is required.
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s SandboxHealthCheckOut) String() string {
	return dara.Prettify(s)
}

func (s SandboxHealthCheckOut) GoString() string {
	return s.String()
}

func (s *SandboxHealthCheckOut) GetStatus() *string {
	return s.Status
}

func (s *SandboxHealthCheckOut) SetStatus(v string) *SandboxHealthCheckOut {
	s.Status = &v
	return s
}

func (s *SandboxHealthCheckOut) Validate() error {
	return dara.Validate(s)
}
