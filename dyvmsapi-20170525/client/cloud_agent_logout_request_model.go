// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudAgentLogoutRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCno(v string) *CloudAgentLogoutRequest
	GetCno() *string
	SetEnterpriseId(v int64) *CloudAgentLogoutRequest
	GetEnterpriseId() *int64
	SetIgnoreOffline(v int64) *CloudAgentLogoutRequest
	GetIgnoreOffline() *int64
	SetIsKickout(v int64) *CloudAgentLogoutRequest
	GetIsKickout() *int64
	SetRemoveBinding(v int64) *CloudAgentLogoutRequest
	GetRemoveBinding() *int64
}

type CloudAgentLogoutRequest struct {
	// 座席工号；取值3-10位正整数
	//
	// This parameter is required.
	//
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 是否忽略重复下线报错；取值：0:不忽略 1:忽略 默认为0，不忽略
	//
	// example:
	//
	// 0
	IgnoreOffline *int64 `json:"IgnoreOffline,omitempty" xml:"IgnoreOffline,omitempty"`
	// 是否给座席发生kickout事件；取值： 0:不发送， 1:发送 默认为1，发送
	//
	// example:
	//
	// 1
	IsKickout *int64 `json:"IsKickout,omitempty" xml:"IsKickout,omitempty"`
	// 取值： 0:不解除绑定电话， 1:解除绑定电话 默认为0
	//
	// example:
	//
	// 0
	RemoveBinding *int64 `json:"RemoveBinding,omitempty" xml:"RemoveBinding,omitempty"`
}

func (s CloudAgentLogoutRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudAgentLogoutRequest) GoString() string {
	return s.String()
}

func (s *CloudAgentLogoutRequest) GetCno() *string {
	return s.Cno
}

func (s *CloudAgentLogoutRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudAgentLogoutRequest) GetIgnoreOffline() *int64 {
	return s.IgnoreOffline
}

func (s *CloudAgentLogoutRequest) GetIsKickout() *int64 {
	return s.IsKickout
}

func (s *CloudAgentLogoutRequest) GetRemoveBinding() *int64 {
	return s.RemoveBinding
}

func (s *CloudAgentLogoutRequest) SetCno(v string) *CloudAgentLogoutRequest {
	s.Cno = &v
	return s
}

func (s *CloudAgentLogoutRequest) SetEnterpriseId(v int64) *CloudAgentLogoutRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudAgentLogoutRequest) SetIgnoreOffline(v int64) *CloudAgentLogoutRequest {
	s.IgnoreOffline = &v
	return s
}

func (s *CloudAgentLogoutRequest) SetIsKickout(v int64) *CloudAgentLogoutRequest {
	s.IsKickout = &v
	return s
}

func (s *CloudAgentLogoutRequest) SetRemoveBinding(v int64) *CloudAgentLogoutRequest {
	s.RemoveBinding = &v
	return s
}

func (s *CloudAgentLogoutRequest) Validate() error {
	return dara.Validate(s)
}
