// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployInstanceInfo interface {
	dara.Model
	String() string
	GoString() string
	SetIp(v string) *DeployInstanceInfo
	GetIp() *string
	SetStatus(v string) *DeployInstanceInfo
	GetStatus() *string
}

type DeployInstanceInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 192.1.1.1
	Ip *string `json:"ip,omitempty" xml:"ip,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// EXECUTING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DeployInstanceInfo) String() string {
	return dara.Prettify(s)
}

func (s DeployInstanceInfo) GoString() string {
	return s.String()
}

func (s *DeployInstanceInfo) GetIp() *string {
	return s.Ip
}

func (s *DeployInstanceInfo) GetStatus() *string {
	return s.Status
}

func (s *DeployInstanceInfo) SetIp(v string) *DeployInstanceInfo {
	s.Ip = &v
	return s
}

func (s *DeployInstanceInfo) SetStatus(v string) *DeployInstanceInfo {
	s.Status = &v
	return s
}

func (s *DeployInstanceInfo) Validate() error {
	return dara.Validate(s)
}
