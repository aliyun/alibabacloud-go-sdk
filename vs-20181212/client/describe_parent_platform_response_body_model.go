// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParentPlatformResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoStart(v bool) *DescribeParentPlatformResponseBody
	GetAutoStart() *bool
	SetClientAuth(v bool) *DescribeParentPlatformResponseBody
	GetClientAuth() *bool
	SetClientGbId(v string) *DescribeParentPlatformResponseBody
	GetClientGbId() *string
	SetClientIp(v string) *DescribeParentPlatformResponseBody
	GetClientIp() *string
	SetClientPassword(v string) *DescribeParentPlatformResponseBody
	GetClientPassword() *string
	SetClientPort(v int64) *DescribeParentPlatformResponseBody
	GetClientPort() *int64
	SetClientUsername(v string) *DescribeParentPlatformResponseBody
	GetClientUsername() *string
	SetCreatedTime(v string) *DescribeParentPlatformResponseBody
	GetCreatedTime() *string
	SetDescription(v string) *DescribeParentPlatformResponseBody
	GetDescription() *string
	SetGbId(v string) *DescribeParentPlatformResponseBody
	GetGbId() *string
	SetId(v string) *DescribeParentPlatformResponseBody
	GetId() *string
	SetIp(v string) *DescribeParentPlatformResponseBody
	GetIp() *string
	SetName(v string) *DescribeParentPlatformResponseBody
	GetName() *string
	SetPort(v int64) *DescribeParentPlatformResponseBody
	GetPort() *int64
	SetProtocol(v string) *DescribeParentPlatformResponseBody
	GetProtocol() *string
	SetRequestId(v string) *DescribeParentPlatformResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeParentPlatformResponseBody
	GetStatus() *string
}

type DescribeParentPlatformResponseBody struct {
	// Specifies whether to enable the parent platform automatically. Valid values:
	//
	// - false (default)
	//
	// - true
	//
	// example:
	//
	// false
	AutoStart *bool `json:"AutoStart,omitempty" xml:"AutoStart,omitempty"`
	// Specifies whether to enable local authentication. Valid values:
	//
	// - true (default)
	//
	// - false
	//
	// example:
	//
	// true
	ClientAuth *bool `json:"ClientAuth,omitempty" xml:"ClientAuth,omitempty"`
	// The local GB ID.
	//
	// example:
	//
	// 31010*****317542006
	ClientGbId *string `json:"ClientGbId,omitempty" xml:"ClientGbId,omitempty"`
	// The local SIP service IP address.
	//
	// example:
	//
	// 192.168.0.1
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// The local password.
	//
	// example:
	//
	// admin123
	ClientPassword *string `json:"ClientPassword,omitempty" xml:"ClientPassword,omitempty"`
	// The local SIP service port.
	//
	// example:
	//
	// 5160
	ClientPort *int64 `json:"ClientPort,omitempty" xml:"ClientPort,omitempty"`
	// The local username.
	//
	// example:
	//
	// user01
	ClientUsername *string `json:"ClientUsername,omitempty" xml:"ClientUsername,omitempty"`
	// The time when the parent platform was created.
	//
	// example:
	//
	// 2018-12-10T21:00:00Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The description of the parent platform.
	//
	// example:
	//
	// 级联平台描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The GB ID of the parent platform.
	//
	// example:
	//
	// 31000*****2170123451
	GbId *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	// The ID of the parent platform.
	//
	// example:
	//
	// 359*****374-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The SIP service IP address of the parent platform.
	//
	// example:
	//
	// 10.10.10.10
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The name of the parent platform.
	//
	// example:
	//
	// 国标级联平台测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The SIP service port of the parent platform.
	//
	// example:
	//
	// 5060
	Port *int64 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol of the parent platform. Valid values:
	//
	// - gb28181 (GB standard)
	//
	// example:
	//
	// gb28181
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the parent platform. Valid values:
	//
	// - on (online)
	//
	// - off (offline)
	//
	// - failed (failed)
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeParentPlatformResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeParentPlatformResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParentPlatformResponseBody) GetAutoStart() *bool {
	return s.AutoStart
}

func (s *DescribeParentPlatformResponseBody) GetClientAuth() *bool {
	return s.ClientAuth
}

func (s *DescribeParentPlatformResponseBody) GetClientGbId() *string {
	return s.ClientGbId
}

func (s *DescribeParentPlatformResponseBody) GetClientIp() *string {
	return s.ClientIp
}

func (s *DescribeParentPlatformResponseBody) GetClientPassword() *string {
	return s.ClientPassword
}

func (s *DescribeParentPlatformResponseBody) GetClientPort() *int64 {
	return s.ClientPort
}

func (s *DescribeParentPlatformResponseBody) GetClientUsername() *string {
	return s.ClientUsername
}

func (s *DescribeParentPlatformResponseBody) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeParentPlatformResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeParentPlatformResponseBody) GetGbId() *string {
	return s.GbId
}

func (s *DescribeParentPlatformResponseBody) GetId() *string {
	return s.Id
}

func (s *DescribeParentPlatformResponseBody) GetIp() *string {
	return s.Ip
}

func (s *DescribeParentPlatformResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeParentPlatformResponseBody) GetPort() *int64 {
	return s.Port
}

func (s *DescribeParentPlatformResponseBody) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeParentPlatformResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeParentPlatformResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeParentPlatformResponseBody) SetAutoStart(v bool) *DescribeParentPlatformResponseBody {
	s.AutoStart = &v
	return s
}

func (s *DescribeParentPlatformResponseBody) SetClientAuth(v bool) *DescribeParentPlatformResponseBody {
	s.ClientAuth = &v
	return s
}

func (s *DescribeParentPlatformResponseBody) SetClientGbId(v string) *DescribeParentPlatformResponseBody {
	s.ClientGbId = &v
	return s
}

func (s *DescribeParentPlatformResponseBody) SetClientIp(v string) *DescribeParentPlatformResponseBody {
	s.ClientIp = &v
	return s
}

func (s *DescribeParentPlatformResponseBody) SetClientPassword(v string) *DescribeParentPlatformResponseBody {
	s.ClientPassword = &v
	return s
}

func (s *DescribeParentPlatformResponseBody) SetClientPort(v int64) *DescribeParentPlatformResponseBody {
	s.ClientPort = &v
	return s
}

func (s *DescribeParentPlatformResponseBody) SetClientUsername(v string) *DescribeParentPlatformResponseBody {
	s.ClientUsername = &v
	return s
}

func (s *DescribeParentPlatformResponseBody) SetCreatedTime(v string) *DescribeParentPlatformResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeParentPlatformResponseBody) SetDescription(v string) *DescribeParentPlatformResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeParentPlatformResponseBody) SetGbId(v string) *DescribeParentPlatformResponseBody {
	s.GbId = &v
	return s
}

func (s *DescribeParentPlatformResponseBody) SetId(v string) *DescribeParentPlatformResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeParentPlatformResponseBody) SetIp(v string) *DescribeParentPlatformResponseBody {
	s.Ip = &v
	return s
}

func (s *DescribeParentPlatformResponseBody) SetName(v string) *DescribeParentPlatformResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeParentPlatformResponseBody) SetPort(v int64) *DescribeParentPlatformResponseBody {
	s.Port = &v
	return s
}

func (s *DescribeParentPlatformResponseBody) SetProtocol(v string) *DescribeParentPlatformResponseBody {
	s.Protocol = &v
	return s
}

func (s *DescribeParentPlatformResponseBody) SetRequestId(v string) *DescribeParentPlatformResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeParentPlatformResponseBody) SetStatus(v string) *DescribeParentPlatformResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeParentPlatformResponseBody) Validate() error {
	return dara.Validate(s)
}
