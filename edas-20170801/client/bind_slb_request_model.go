// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindSlbRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *BindSlbRequest
	GetAppId() *string
	SetListenerPort(v int32) *BindSlbRequest
	GetListenerPort() *int32
	SetSlbId(v string) *BindSlbRequest
	GetSlbId() *string
	SetSlbIp(v string) *BindSlbRequest
	GetSlbIp() *string
	SetType(v string) *BindSlbRequest
	GetType() *string
	SetVServerGroupId(v string) *BindSlbRequest
	GetVServerGroupId() *string
}

type BindSlbRequest struct {
	// The ID of the EDAS application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3616cdca-*********
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The listener port for the SLB instance.
	//
	// example:
	//
	// 80
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The ID of the SLB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-wz96ph63r************
	SlbId *string `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
	// The IP address of the SLB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.16*.*.*
	SlbIp *string `json:"SlbIp,omitempty" xml:"SlbIp,omitempty"`
	// The type of the SLB instance. Valid values:
	//
	// 	- internet: Internet-facing SLB instance
	//
	// 	- intranet: internal-facing SLB instance
	//
	// This parameter is required.
	//
	// example:
	//
	// intranet
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the vServer group for the internal-facing SLB instance.
	//
	// example:
	//
	// rsp-cige6******
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s BindSlbRequest) String() string {
	return dara.Prettify(s)
}

func (s BindSlbRequest) GoString() string {
	return s.String()
}

func (s *BindSlbRequest) GetAppId() *string {
	return s.AppId
}

func (s *BindSlbRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *BindSlbRequest) GetSlbId() *string {
	return s.SlbId
}

func (s *BindSlbRequest) GetSlbIp() *string {
	return s.SlbIp
}

func (s *BindSlbRequest) GetType() *string {
	return s.Type
}

func (s *BindSlbRequest) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *BindSlbRequest) SetAppId(v string) *BindSlbRequest {
	s.AppId = &v
	return s
}

func (s *BindSlbRequest) SetListenerPort(v int32) *BindSlbRequest {
	s.ListenerPort = &v
	return s
}

func (s *BindSlbRequest) SetSlbId(v string) *BindSlbRequest {
	s.SlbId = &v
	return s
}

func (s *BindSlbRequest) SetSlbIp(v string) *BindSlbRequest {
	s.SlbIp = &v
	return s
}

func (s *BindSlbRequest) SetType(v string) *BindSlbRequest {
	s.Type = &v
	return s
}

func (s *BindSlbRequest) SetVServerGroupId(v string) *BindSlbRequest {
	s.VServerGroupId = &v
	return s
}

func (s *BindSlbRequest) Validate() error {
	return dara.Validate(s)
}
