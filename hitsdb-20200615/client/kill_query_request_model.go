// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKillQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *KillQueryRequest
	GetInstanceId() *string
	SetProcessId(v string) *KillQueryRequest
	GetProcessId() *string
	SetRegionId(v string) *KillQueryRequest
	GetRegionId() *string
	SetSecurityToken(v string) *KillQueryRequest
	GetSecurityToken() *string
}

type KillQueryRequest struct {
	// This parameter is required.
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ProcessId     *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s KillQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s KillQueryRequest) GoString() string {
	return s.String()
}

func (s *KillQueryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *KillQueryRequest) GetProcessId() *string {
	return s.ProcessId
}

func (s *KillQueryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *KillQueryRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *KillQueryRequest) SetInstanceId(v string) *KillQueryRequest {
	s.InstanceId = &v
	return s
}

func (s *KillQueryRequest) SetProcessId(v string) *KillQueryRequest {
	s.ProcessId = &v
	return s
}

func (s *KillQueryRequest) SetRegionId(v string) *KillQueryRequest {
	s.RegionId = &v
	return s
}

func (s *KillQueryRequest) SetSecurityToken(v string) *KillQueryRequest {
	s.SecurityToken = &v
	return s
}

func (s *KillQueryRequest) Validate() error {
	return dara.Validate(s)
}
